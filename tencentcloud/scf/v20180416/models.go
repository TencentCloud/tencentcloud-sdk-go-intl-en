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

package v20180416

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AccessInfo struct {

	// Domain name
	Host *string `json:"Host,omitempty" name:"Host"`

	// VIP
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type Alias struct {

	// Master version of alias
	FunctionVersion *string `json:"FunctionVersion,omitempty" name:"FunctionVersion"`

	// Alias name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Routing information of alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	RoutingConfig *RoutingConfig `json:"RoutingConfig,omitempty" name:"RoutingConfig"`

	// Description
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Update time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`
}

type Code struct {

	// COS bucket name
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// COS object path
	CosObjectName *string `json:"CosObjectName,omitempty" name:"CosObjectName"`

	// It contains a function code file and its dependencies in the ZIP format. When you use this API, the ZIP file needs to be encoded with Base64. Up to 20 MB is supported.
	ZipFile *string `json:"ZipFile,omitempty" name:"ZipFile"`

	// COS region. For Beijing regions, you need to import `ap-beijing`. For Beijing Region 1, you need to input `ap-beijing-1`. For other regions, no import is required.
	CosBucketRegion *string `json:"CosBucketRegion,omitempty" name:"CosBucketRegion"`

	// `DemoId` is required if Demo is used for the creation.
	DemoId *string `json:"DemoId,omitempty" name:"DemoId"`

	// `TempCosObjectName` is required if TempCos is used for the creation.
	TempCosObjectName *string `json:"TempCosObjectName,omitempty" name:"TempCosObjectName"`

	// Git address
	GitUrl *string `json:"GitUrl,omitempty" name:"GitUrl"`

	// Git user name
	GitUserName *string `json:"GitUserName,omitempty" name:"GitUserName"`

	// Git password
	GitPassword *string `json:"GitPassword,omitempty" name:"GitPassword"`

	// Git password after encryption. In general, this value is not required.
	GitPasswordSecret *string `json:"GitPasswordSecret,omitempty" name:"GitPasswordSecret"`

	// Git branch
	GitBranch *string `json:"GitBranch,omitempty" name:"GitBranch"`

	// Code path in Git repository
	GitDirectory *string `json:"GitDirectory,omitempty" name:"GitDirectory"`

	// Version to be pulled
	GitCommitId *string `json:"GitCommitId,omitempty" name:"GitCommitId"`

	// Git user name after encryption. In general, this value is not required.
	GitUserNameSecret *string `json:"GitUserNameSecret,omitempty" name:"GitUserNameSecret"`
}

type CopyFunctionRequest struct {
	*tchttp.BaseRequest

	// Name of the function to be replicated
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Name of the new function
	NewFunctionName *string `json:"NewFunctionName,omitempty" name:"NewFunctionName"`

	// Namespace of the function to be replicated. The default value is `default`.
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Namespace of the replicated function. The default value is default.
	TargetNamespace *string `json:"TargetNamespace,omitempty" name:"TargetNamespace"`

	// Description of the new function
	Description *string `json:"Description,omitempty" name:"Description"`

	// Region of the target of the function replication. If the value is not set, the current region is used by default.
	TargetRegion *string `json:"TargetRegion,omitempty" name:"TargetRegion"`

	// It specifies whether to replace the function with the same name in the target namespace. The default option is `FALSE`.
	// (Note: The `TRUE` option results in deletion of the function in the target namespace. Please operate with caution.)
	// TRUE: Replaces the function.
	// FALSE: Does not replace the function.
	Override *bool `json:"Override,omitempty" name:"Override"`

	// It specifies whether to replicate the function attributes, including environment variables, memory, timeout, function description, labels, and VPC. The default value is `TRUE`.
	// TRUE: Replicates the function configuration.
	// FALSE: Does not replicate the function configuration.
	CopyConfiguration *bool `json:"CopyConfiguration,omitempty" name:"CopyConfiguration"`
}

func (r *CopyFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAliasRequest struct {
	*tchttp.BaseRequest

	// Alias name, which must be unique in the function, can contain 1–64 letters, digits, `_`, and `-`, and must begin with a letter
	Name *string `json:"Name,omitempty" name:"Name"`

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Master version of alias
	FunctionVersion *string `json:"FunctionVersion,omitempty" name:"FunctionVersion"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Request routing configuration of alias
	RoutingConfig *RoutingConfig `json:"RoutingConfig,omitempty" name:"RoutingConfig"`

	// Alias description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFunctionRequest struct {
	*tchttp.BaseRequest

	// Name of the new function. The name can contain 2 to 60 characters, including English letters, digits, hyphens (-), and underscores (_). The name must start with a letter and cannot end with a hyphen or underscore.
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Function code. Note: You cannot specify `Cos` and `ZipFile` at the same time.
	Code *Code `json:"Code,omitempty" name:"Code"`

	// Name of the handler, which is in the “file name.handler name” form. Use periods (.) to separate a file name and function name. The file name and function name must start and end with a letter and can contain 2 to 60 characters, including letters, digits, hyphens (-), and underscores (_).
	Handler *string `json:"Handler,omitempty" name:"Handler"`

	// Function description. It can contain up to 1,000 characters including letters, digits, spaces, commas (,), periods (.), and Chinese characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Memory size available for function during execution. Default value: 128 MB. Value range: 64 or 128–3,072 MB in increments of 128 MB
	MemorySize *int64 `json:"MemorySize,omitempty" name:"MemorySize"`

	// Maximum execution duration of function in seconds. Value range: 1–900 seconds. Default value: 3 seconds
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// Function environment variable
	Environment *Environment `json:"Environment,omitempty" name:"Environment"`

	// Function runtime environment. Valid values: Python2.7, Python3.6, Nodejs6.10, Nodejs8.9, Nodejs10.15, Nodejs12.16, PHP5, PHP7, Golang1 and Java8. Default value: Python2.7
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// Function VPC configuration
	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Role bound to the function
	Role *string `json:"Role,omitempty" name:"Role"`

	// CLS Logset ID to which the function logs are shipped
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// CLS Topic ID to which the function logs are shipped
	ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

	// Function type. The default value is `Event`. Enter `Event` if you need to create a trigger function. Enter `HTTP` if you need to create an HTTP function service.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Code source, including ZipFile, Cos, Demo, TempCos, and Git. This field is required if the source is Git.
	CodeSource *string `json:"CodeSource,omitempty" name:"CodeSource"`

	// List of layer versions to be associate with the function. Layers will be overwritten sequentially in the order in the list.
	Layers []*LayerVersionSimple `json:"Layers,omitempty" name:"Layers" list`

	// Dead letter queue parameter
	DeadLetterConfig *DeadLetterConfig `json:"DeadLetterConfig,omitempty" name:"DeadLetterConfig"`

	// Public network access configuration
	PublicNetConfig *PublicNetConfigIn `json:"PublicNetConfig,omitempty" name:"PublicNetConfig"`
}

func (r *CreateFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest

	// Namespace name
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Namespace description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTriggerRequest struct {
	*tchttp.BaseRequest

	// Name of the function bound to the new trigger
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Name of a new trigger. For a timer trigger, the name can contain up to 100 letters, digits, dashes, and underscores; for a COS trigger, it should be an access domain name of the corresponding COS bucket applicable to the XML API (e.g., 5401-5ff414-12345.cos.ap-shanghai.myqcloud.com); for other triggers, please see the descriptions of parameters bound to the specific trigger.
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// Trigger type. Currently, COS, CMQ, timer, and ckafka triggers are supported.
	Type *string `json:"Type,omitempty" name:"Type"`

	// For parameters of triggers, see [Trigger Description](https://cloud.tencent.com/document/product/583/39901)
	TriggerDesc *string `json:"TriggerDesc,omitempty" name:"TriggerDesc"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Function version
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// Initial enabling status of the trigger. `OPEN` indicates enabled, and `CLOSE` indicates disabled.
	Enable *string `json:"Enable,omitempty" name:"Enable"`
}

func (r *CreateTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTriggerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTriggerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Trigger information
		TriggerInfo *Trigger `json:"TriggerInfo,omitempty" name:"TriggerInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTriggerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeadLetterConfig struct {

	// Dead letter queue mode
	Type *string `json:"Type,omitempty" name:"Type"`

	// Dead letter queue name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Tag form of a dead letter queue topic mode
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type DeleteAliasRequest struct {
	*tchttp.BaseRequest

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Alias name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFunctionRequest struct {
	*tchttp.BaseRequest

	// Name of the function to be deleted
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLayerVersionRequest struct {
	*tchttp.BaseRequest

	// Layer name
	LayerName *string `json:"LayerName,omitempty" name:"LayerName"`

	// Version number
	LayerVersion *int64 `json:"LayerVersion,omitempty" name:"LayerVersion"`
}

func (r *DeleteLayerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLayerVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLayerVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLayerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLayerVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest

	// Namespace name
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTriggerRequest struct {
	*tchttp.BaseRequest

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Name of the trigger to be deleted
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// Type of the trigger to be deleted. Currently, COS, CMQ, timer, and ckafka triggers are supported.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// This field is required if a COS trigger is to be deleted. It stores the data {"event":"cos:ObjectCreated:*"} in the JSON format. The data content of this field is in the same format as that of SetTrigger. This field is optional if a scheduled trigger or CMQ trigger is to be deleted.
	TriggerDesc *string `json:"TriggerDesc,omitempty" name:"TriggerDesc"`

	// Function version information
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`
}

func (r *DeleteTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTriggerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTriggerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTriggerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EipConfigIn struct {

	// Status of the EIP. Values: ['ENABLE','DISABLE']
	EipStatus *string `json:"EipStatus,omitempty" name:"EipStatus"`
}

type EipConfigOut struct {

	// Whether it is a fixed IP. Valid values: ["ENABLE","DISABLE"]
	EipStatus *string `json:"EipStatus,omitempty" name:"EipStatus"`

	// IP list
	// Note: This field may return null, indicating that no valid values can be obtained.
	EipAddress []*string `json:"EipAddress,omitempty" name:"EipAddress" list`
}

type EipOutConfig struct {

	// It specifies whether the IP is fixed. The value is `TRUE` or `FALSE`.
	EipFixed *string `json:"EipFixed,omitempty" name:"EipFixed"`

	// IP list
	Eips []*string `json:"Eips,omitempty" name:"Eips" list`
}

type Environment struct {

	// Environment variable array
	Variables []*Variable `json:"Variables,omitempty" name:"Variables" list`
}

type Filter struct {

	// Fields to be filtered
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter values of the field
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Function struct {

	// Modification time
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// Creation time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Running
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Function ID
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// Namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Function status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Function status details
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Function description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Function tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Function type. The value is `HTTP` or `Event`.
	Type *string `json:"Type,omitempty" name:"Type"`
}

type FunctionLog struct {

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Return value after the function is executed
	RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`

	// RequestId corresponding to the executed function
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// Start time of the function execution
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Function execution result. `0` indicates successful execution and other values indicate failure.
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`

	// It specifies whether the function invocation is finished. `1` indicates execution completion and other values indicate that exceptions occurred during the invocation.
	InvokeFinished *int64 `json:"InvokeFinished,omitempty" name:"InvokeFinished"`

	// Duration of the function execution. The unit is millisecond (ms).
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Function billing duration. The unit is millisecond (ms). The value is rounded up to a multiple of 100 ms.
	BillDuration *int64 `json:"BillDuration,omitempty" name:"BillDuration"`

	// Actual memory size used during the function execution. The unit is byte.
	MemUsage *int64 `json:"MemUsage,omitempty" name:"MemUsage"`

	// Function execution logs
	Log *string `json:"Log,omitempty" name:"Log"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Log source
	Source *string `json:"Source,omitempty" name:"Source"`

	// Number of retries
	RetryNum *uint64 `json:"RetryNum,omitempty" name:"RetryNum"`
}

type FunctionVersion struct {

	// Function version name
	Version *string `json:"Version,omitempty" name:"Version"`

	// Version description
	// Note: This field may return null, indicating that no valid values is found.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The creation time
	// Note: This field may return null, indicating that no valid value was found.
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Update time
	// Note: This field may return null, indicating that no valid value was found.
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`
}

type GetAliasRequest struct {
	*tchttp.BaseRequest

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Alias name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *GetAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Master version of alias
		FunctionVersion *string `json:"FunctionVersion,omitempty" name:"FunctionVersion"`

		// Alias name
		Name *string `json:"Name,omitempty" name:"Name"`

		// Routing information of alias
		RoutingConfig *RoutingConfig `json:"RoutingConfig,omitempty" name:"RoutingConfig"`

		// Alias description
	// Note: this field may return null, indicating that no valid values can be obtained.
		Description *string `json:"Description,omitempty" name:"Description"`

		// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
		AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

		// Update time
	// Note: this field may return null, indicating that no valid values can be obtained.
		ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFunctionAddressRequest struct {
	*tchttp.BaseRequest

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Function version
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *GetFunctionAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFunctionAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cos address of the function
		Url *string `json:"Url,omitempty" name:"Url"`

		// SHA256 code of the function
		CodeSha256 *string `json:"CodeSha256,omitempty" name:"CodeSha256"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFunctionAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFunctionLogsRequest struct {
	*tchttp.BaseRequest

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Data offset. The addition of `Offset` and `Limit` cannot exceed 10,000.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Length of the return data. The addition of `Offset` and `Limit` cannot exceed 10,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// It specifies whether to sort the logs in an ascending or descending order. The value is `desc` or `asc`.
	Order *string `json:"Order,omitempty" name:"Order"`

	// It specifies the sorting order of the logs based on a specified field, such as `function_name`, `duration`, `mem_usage`, and `start_time`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Log filter used to identify whether to return logs of successful or failed requests. `filter.RetCode=not0` indicates that only the logs of failed requests will be returned. `filter.RetCode=is0` indicates that only the logs of successful requests will be returned. If this parameter is left blank, all logs will be returned. 
	Filter *LogFilter `json:"Filter,omitempty" name:"Filter"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Function version
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// RequestId corresponding to the executed function
	FunctionRequestId *string `json:"FunctionRequestId,omitempty" name:"FunctionRequestId"`

	// Query date, for example, 2017-05-16 20:00:00. The date must be within one day of the end time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query date, for example, 2017-05-16 20:59:59. The date must be within one day of the start time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service log related parameter. `Offset` on the first page is a null string. Enter other pages based on SearchContext in the response field.
	SearchContext *LogSearchContext `json:"SearchContext,omitempty" name:"SearchContext"`
}

func (r *GetFunctionLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFunctionLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of function logs
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Function log information
		Data []*FunctionLog `json:"Data,omitempty" name:"Data" list`

		// Parameter on the log service page
		SearchContext *LogSearchContext `json:"SearchContext,omitempty" name:"SearchContext"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFunctionLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFunctionRequest struct {
	*tchttp.BaseRequest

	// Name of the function to obtain details
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Function version number
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// It indicates whether to display the code. `TRUE` means displaying the code, and `FALSE` means hiding the code. The code will not be displayed for entry files exceeding 1 MB.
	ShowCode *string `json:"ShowCode,omitempty" name:"ShowCode"`
}

func (r *GetFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Latest modification time of the function
		ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

		// Function code
		CodeInfo *string `json:"CodeInfo,omitempty" name:"CodeInfo"`

		// Function description
		Description *string `json:"Description,omitempty" name:"Description"`

		// Function trigger list
		Triggers []*Trigger `json:"Triggers,omitempty" name:"Triggers" list`

		// Function entry
		Handler *string `json:"Handler,omitempty" name:"Handler"`

		// Function code size
		CodeSize *int64 `json:"CodeSize,omitempty" name:"CodeSize"`

		// Function timeout
		Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

		// Function version
		FunctionVersion *string `json:"FunctionVersion,omitempty" name:"FunctionVersion"`

		// Maximum available memory of the function
		MemorySize *int64 `json:"MemorySize,omitempty" name:"MemorySize"`

		// Function running environment
		Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

		// Function name
		FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

		// Function VPC
		VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

		// Whether to use GPU
		UseGpu *string `json:"UseGpu,omitempty" name:"UseGpu"`

		// Function environment variable
		Environment *Environment `json:"Environment,omitempty" name:"Environment"`

		// Whether the code is correct
		CodeResult *string `json:"CodeResult,omitempty" name:"CodeResult"`

		// Code error information
		CodeError *string `json:"CodeError,omitempty" name:"CodeError"`

		// Error code
		ErrNo *int64 `json:"ErrNo,omitempty" name:"ErrNo"`

		// Function namespace
		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

		// Role bound to the function
		Role *string `json:"Role,omitempty" name:"Role"`

		// Whether to install dependencies automatically
		InstallDependency *string `json:"InstallDependency,omitempty" name:"InstallDependency"`

		// Function status
		Status *string `json:"Status,omitempty" name:"Status"`

		// Status description
		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

		// CLS logset to which logs are shipped
		ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

		// CLS Topic to which logs are shipped
		ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

		// Function ID
		FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

		// Function tag list
		Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

		// EipConfig configuration
		EipConfig *EipOutConfig `json:"EipConfig,omitempty" name:"EipConfig"`

		// Domain name information
		AccessInfo *AccessInfo `json:"AccessInfo,omitempty" name:"AccessInfo"`

		// Function type. The value is `HTTP` or `Event`.
		Type *string `json:"Type,omitempty" name:"Type"`

		// Whether to enable L5
		L5Enable *string `json:"L5Enable,omitempty" name:"L5Enable"`

		// Version information of a layer associated with a function
		Layers []*LayerVersionInfo `json:"Layers,omitempty" name:"Layers" list`

		// Information of a dead letter queue associated with a function
		DeadLetterConfig *DeadLetterConfig `json:"DeadLetterConfig,omitempty" name:"DeadLetterConfig"`

		// Function creation time
		AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

		// Public network access configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
		PublicNetConfig *PublicNetConfigOut `json:"PublicNetConfig,omitempty" name:"PublicNetConfig"`

		// Whether Ons is enabled
	// Note: This field may return null, indicating that no valid value was found.
		OnsEnable *string `json:"OnsEnable,omitempty" name:"OnsEnable"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLayerVersionRequest struct {
	*tchttp.BaseRequest

	// Layer name
	LayerName *string `json:"LayerName,omitempty" name:"LayerName"`

	// Version number
	LayerVersion *int64 `json:"LayerVersion,omitempty" name:"LayerVersion"`
}

func (r *GetLayerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLayerVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLayerVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Compatible runtimes
		CompatibleRuntimes []*string `json:"CompatibleRuntimes,omitempty" name:"CompatibleRuntimes" list`

		// SHA256 encoding of file on layer version
		CodeSha256 *string `json:"CodeSha256,omitempty" name:"CodeSha256"`

		// Download address of file on layer version
		Location *string `json:"Location,omitempty" name:"Location"`

		// Version creation time
		AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

		// Version description
		Description *string `json:"Description,omitempty" name:"Description"`

		// License information
		LicenseInfo *string `json:"LicenseInfo,omitempty" name:"LicenseInfo"`

		// Version number
		LayerVersion *int64 `json:"LayerVersion,omitempty" name:"LayerVersion"`

		// Layer name
		LayerName *string `json:"LayerName,omitempty" name:"LayerName"`

		// Current status of specific layer version. Valid values:
	// Active: normal
	// Publishing: publishing
	// PublishFailed: publishing failed
	// Deleted: deleted
		Status *string `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLayerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLayerVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InvokeRequest struct {
	*tchttp.BaseRequest

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// The value is `RequestResponse` (synchronous) or `Event` (asynchronous). The default value is synchronous.
	InvocationType *string `json:"InvocationType,omitempty" name:"InvocationType"`

	// Version number of the triggered function
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// Function running parameter, which is in the JSON format. Maximum parameter size is 1 MB.
	ClientContext *string `json:"ClientContext,omitempty" name:"ClientContext"`

	// If this field is specified for a synchronous invocation, the return value will contain a 4-KB log. The value is `None` (default) or `Tail`. If the value is `Tail`, `logMsg` in the return parameter will contain the corresponding function execution log.
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// Namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Traffic routing config in json format, e.g., {"k":"v"}. Please note that both "k" and "v" must be strings. Up to 1024 bytes allowed.
	RoutingKey *string `json:"RoutingKey,omitempty" name:"RoutingKey"`
}

func (r *InvokeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InvokeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InvokeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Function execution result
		Result *Result `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InvokeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InvokeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LayerVersionInfo struct {

	// Runtime applicable to a version
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompatibleRuntimes []*string `json:"CompatibleRuntimes,omitempty" name:"CompatibleRuntimes" list`

	// Creation time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Version description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// License information
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseInfo *string `json:"LicenseInfo,omitempty" name:"LicenseInfo"`

	// Version number
	LayerVersion *int64 `json:"LayerVersion,omitempty" name:"LayerVersion"`

	// Layer name
	LayerName *string `json:"LayerName,omitempty" name:"LayerName"`

	// The status of the layer version. Values can be: 
	// `Active`: normal
	// `Publishing`: publishing
	// `PublishFailed`: failed to publish
	// `Deleted`: deleted
	Status *string `json:"Status,omitempty" name:"Status"`
}

type LayerVersionSimple struct {

	// Layer name
	LayerName *string `json:"LayerName,omitempty" name:"LayerName"`

	// Version number
	LayerVersion *int64 `json:"LayerVersion,omitempty" name:"LayerVersion"`
}

type ListAliasesRequest struct {
	*tchttp.BaseRequest

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// If this parameter is provided, only aliases associated with this function version will be returned
	FunctionVersion *string `json:"FunctionVersion,omitempty" name:"FunctionVersion"`

	// Data offset. Default value: 0
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListAliasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAliasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAliasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alias list
		Aliases []*Alias `json:"Aliases,omitempty" name:"Aliases" list`

		// Total number of aliases
	// Note: this field may return null, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAliasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAliasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListFunctionsRequest struct {
	*tchttp.BaseRequest

	// It specifies whether to return the results in ascending or descending order. The value is `ASC` or `DESC`.
	Order *string `json:"Order,omitempty" name:"Order"`

	// It specifies the sorting order of the results according to a specified field, such as `AddTime`, `ModTime`, and `FunctionName`.
	Orderby *string `json:"Orderby,omitempty" name:"Orderby"`

	// Data offset. The default value is `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Return data length. The default value is `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// It specifies whether to support fuzzy matching for the function name.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Function description. Fuzzy search is supported.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Filters
	// - tag:tag-key - String - Required: No - Filtering criteria based on tag-key - value pairs. Replace `tag-key` with a specific tag-key.
	// 
	// The maximum number of `Filters` for each request is 10, and that of `Filter.Values` is 5.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *ListFunctionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFunctionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListFunctionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Function list
		Functions []*Function `json:"Functions,omitempty" name:"Functions" list`

		// Total number
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListFunctionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFunctionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListLayerVersionsRequest struct {
	*tchttp.BaseRequest

	// Layer name
	LayerName *string `json:"LayerName,omitempty" name:"LayerName"`

	// Compatible runtimes
	CompatibleRuntime []*string `json:"CompatibleRuntime,omitempty" name:"CompatibleRuntime" list`
}

func (r *ListLayerVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListLayerVersionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListLayerVersionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Layer version list
		LayerVersions []*LayerVersionInfo `json:"LayerVersions,omitempty" name:"LayerVersions" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLayerVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListLayerVersionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListLayersRequest struct {
	*tchttp.BaseRequest

	// Compatible runtimes
	CompatibleRuntime *string `json:"CompatibleRuntime,omitempty" name:"CompatibleRuntime"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query key, which fuzzily matches the name
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *ListLayersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListLayersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListLayersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Layer list
		Layers []*LayerVersionInfo `json:"Layers,omitempty" name:"Layers" list`

		// Total number of layers
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLayersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListLayersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListNamespacesRequest struct {
	*tchttp.BaseRequest

	// Return data length. The default value is `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Data offset. The default value is `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// It specifies the sorting order of the results according to a specified field, such as `Name` and `Updatetime`.
	Orderby *string `json:"Orderby,omitempty" name:"Orderby"`

	// It specifies whether to return the results in ascending or descending order. The value is `ASC` or `DESC`.
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ListNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListNamespacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Namespace details
		Namespaces []*Namespace `json:"Namespaces,omitempty" name:"Namespaces" list`

		// Number of return namespaces
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListNamespacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTriggersRequest struct {
	*tchttp.BaseRequest

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Namespace. Default value: default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Data offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Indicates by which field to sort the returned results. Valid values: AddTime, ModTime. Default value: ModTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Indicates whether the returned results are sorted in ascending or descending order. Valid values: ASC, DESC. Default value: DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// * Qualifier:
	// Function version, i.e., alias
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *ListTriggersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTriggersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTriggersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of triggers
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Trigger list
		Triggers []*TriggerInfo `json:"Triggers,omitempty" name:"Triggers" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTriggersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTriggersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListVersionByFunctionRequest struct {
	*tchttp.BaseRequest

	// Function Name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// The namespace where the function locates
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Data offset. The default value is `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Return data length. The default value is `20`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// It specifies whether to return the results in ascending or descending order. The value is `ASC` or `DESC`.
	Order *string `json:"Order,omitempty" name:"Order"`

	// It specifies the sorting order of the results according to a specified field, such as `AddTime`, `ModTime`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListVersionByFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListVersionByFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListVersionByFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Function version
		FunctionVersion []*string `json:"FunctionVersion,omitempty" name:"FunctionVersion" list`

		// Function version list
	// Note: This field may return null, indicating that no valid values is found.
		Versions []*FunctionVersion `json:"Versions,omitempty" name:"Versions" list`

		// Total number of function versions
	// Note: This field may return null, indicating that no valid value was found.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListVersionByFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListVersionByFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LogFilter struct {

	// Values of `filter.RetCode` include:
	// not0, indicating that only logs of failed execution will be returned.
	// is0, indicating that only logs of successful execution will be returned.
	// TimeLimitExceeded, indicating that logs of function invocations which timed out will be returned.
	// ResourceLimitExceeded, indicating that logs of function invocations during which resources exceeded the upper limit will be returned.
	// UserCodeException, indicating that logs of function invocations during which a user code error occurred will be returned.
	// Blank, indicating that all logs will be returned.
	RetCode *string `json:"RetCode,omitempty" name:"RetCode"`
}

type LogSearchContext struct {

	// Offset.
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Log record number
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Log keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Log type. The value is `Application` (default) or `Platform`.
	Type *string `json:"Type,omitempty" name:"Type"`
}

type Namespace struct {

	// Creation time of the namespace
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// Modification time of the namespace
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Namespace description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Namespace name
	Name *string `json:"Name,omitempty" name:"Name"`

	// The default value is default. TCB indicates that the namespace is developed and created through the mini-program cloud.
	Type *string `json:"Type,omitempty" name:"Type"`
}

type PublicNetConfigIn struct {

	// Whether to enable public network access. Valid values: ['DISABLE', 'ENABLE']
	PublicNetStatus *string `json:"PublicNetStatus,omitempty" name:"PublicNetStatus"`

	// EIP configuration
	EipConfig *EipConfigIn `json:"EipConfig,omitempty" name:"EipConfig"`
}

type PublicNetConfigOut struct {

	// Whether to enable public network access. Valid values: ['DISABLE', 'ENABLE']
	PublicNetStatus *string `json:"PublicNetStatus,omitempty" name:"PublicNetStatus"`

	// EIP configuration
	EipConfig *EipConfigOut `json:"EipConfig,omitempty" name:"EipConfig"`
}

type PublishLayerVersionRequest struct {
	*tchttp.BaseRequest

	// Layer name, which can contain 1–64 English letters, digits, hyphens, and underscores, must begin with a letter, and cannot end with a hyphen or underscore
	LayerName *string `json:"LayerName,omitempty" name:"LayerName"`

	// Runtimes compatible with layer. Multiple choices are allowed. The valid values of this parameter correspond to the valid values of the `Runtime` of the function.
	CompatibleRuntimes []*string `json:"CompatibleRuntimes,omitempty" name:"CompatibleRuntimes" list`

	// Layer file source or content
	Content *Code `json:"Content,omitempty" name:"Content"`

	// Layer version description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Software license of layer
	LicenseInfo *string `json:"LicenseInfo,omitempty" name:"LicenseInfo"`
}

func (r *PublishLayerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishLayerVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishLayerVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Version number of the layer created in this request
		LayerVersion *int64 `json:"LayerVersion,omitempty" name:"LayerVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishLayerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishLayerVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishVersionRequest struct {
	*tchttp.BaseRequest

	// Name of the released function
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Function description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *PublishVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Function version
		FunctionVersion *string `json:"FunctionVersion,omitempty" name:"FunctionVersion"`

		// Code size
		CodeSize *int64 `json:"CodeSize,omitempty" name:"CodeSize"`

		// Maximum available memory
		MemorySize *int64 `json:"MemorySize,omitempty" name:"MemorySize"`

		// Function description
		Description *string `json:"Description,omitempty" name:"Description"`

		// Function entry
		Handler *string `json:"Handler,omitempty" name:"Handler"`

		// Function timeout
		Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

		// Function running environment
		Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

		// Function namespace
		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Result struct {

	// It indicates the log output during the function execution. Null is returned for asynchronous invocations.
	Log *string `json:"Log,omitempty" name:"Log"`

	// It indicates the response from the executed function. Null is returned for asynchronous invocations.
	RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`

	// It indicates the error message of the executed function. Null is returned for asynchronous invocations.
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// It indicates the memory size (in bytes) when the function is running. Null is returned for asynchronous invocations.
	MemUsage *int64 `json:"MemUsage,omitempty" name:"MemUsage"`

	// It indicates the duration (in milliseconds) required for running the function. Null is returned for asynchronous invocations.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// It indicates the billing duration (in milliseconds) for the function. Null is returned for asynchronous invocations.
	BillDuration *int64 `json:"BillDuration,omitempty" name:"BillDuration"`

	// ID of the executed function
	FunctionRequestId *string `json:"FunctionRequestId,omitempty" name:"FunctionRequestId"`

	// `0` indicates successful execution. Null is returned for asynchronous invocations.
	InvokeResult *int64 `json:"InvokeResult,omitempty" name:"InvokeResult"`
}

type RoutingConfig struct {

	// Additional version with random weight-based routing
	AdditionalVersionWeights []*VersionWeight `json:"AdditionalVersionWeights,omitempty" name:"AdditionalVersionWeights" list`

	// Additional version with rule-based routing
	AddtionVersionMatchs []*VersionMatch `json:"AddtionVersionMatchs,omitempty" name:"AddtionVersionMatchs" list`
}

type Tag struct {

	// Tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Trigger struct {

	// Latest modification time of the trigger
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// Trigger type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Detailed trigger configuration
	TriggerDesc *string `json:"TriggerDesc,omitempty" name:"TriggerDesc"`

	// Trigger name
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// Creation time of the trigger
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Enabling switch
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Custom parameter
	CustomArgument *string `json:"CustomArgument,omitempty" name:"CustomArgument"`

	// Trigger status
	AvailableStatus *string `json:"AvailableStatus,omitempty" name:"AvailableStatus"`
}

type TriggerInfo struct {

	// Enablement switch
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// Function version or alias
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// Trigger name
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// Trigger type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Detailed configuration of trigger
	TriggerDesc *string `json:"TriggerDesc,omitempty" name:"TriggerDesc"`

	// Whether the trigger is available
	AvailableStatus *string `json:"AvailableStatus,omitempty" name:"AvailableStatus"`

	// Custom parameter
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomArgument *string `json:"CustomArgument,omitempty" name:"CustomArgument"`

	// Trigger creation time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Trigger last modified time
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`
}

type UpdateAliasRequest struct {
	*tchttp.BaseRequest

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Alias name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Master version of alias
	FunctionVersion *string `json:"FunctionVersion,omitempty" name:"FunctionVersion"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Routing information of alias, which is required if you need to specify an additional version for the alias.
	RoutingConfig *RoutingConfig `json:"RoutingConfig,omitempty" name:"RoutingConfig"`

	// Alias description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFunctionCodeRequest struct {
	*tchttp.BaseRequest

	// Function handler name, which is in the `file name.function name` form. Use a period (.) to separate a file name and function name. The file name and function name must start and end with letters and contain 2-60 characters, including letters, digits, underscores (_), and hyphens (-).
	Handler *string `json:"Handler,omitempty" name:"Handler"`

	// Name of the function to be modified
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// COS bucket name
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// COS object path
	CosObjectName *string `json:"CosObjectName,omitempty" name:"CosObjectName"`

	// It contains a function code file and its dependencies in the ZIP format. When you use this API, the ZIP file needs to be encoded with Base64. Up to 20 MB is supported.
	ZipFile *string `json:"ZipFile,omitempty" name:"ZipFile"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// COS region. Note: Beijing includes ap-beijing and ap-beijing-1.
	CosBucketRegion *string `json:"CosBucketRegion,omitempty" name:"CosBucketRegion"`

	// Function environment
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// It specifies whether to synchronously release a new version during the update. The default value is `FALSE`, indicating not to release a new version.
	Publish *string `json:"Publish,omitempty" name:"Publish"`

	// Function code
	Code *Code `json:"Code,omitempty" name:"Code"`

	// Source mode of code. Valid values: `ZipFile`, `Cos`, `Inline`, `TempCos` and `Git`. This field must be specified if the source is Git
	CodeSource *string `json:"CodeSource,omitempty" name:"CodeSource"`
}

func (r *UpdateFunctionCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFunctionCodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFunctionCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFunctionCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFunctionCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFunctionConfigurationRequest struct {
	*tchttp.BaseRequest

	// Name of the function to be modified
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Function description. It can contain up to 1,000 characters, including letters, digits, spaces, commas (,), periods (.), and Chinese characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Memory size available for function during execution. Default value: 128 MB. Value range: 64 or 128–3,072 MB in increments of 128 MB.
	MemorySize *int64 `json:"MemorySize,omitempty" name:"MemorySize"`

	// Maximum execution duration of function in seconds. Value range: 1–900 seconds. Default value: 3 seconds
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// Function runtime environment. Valid values: Python2.7, Python3.6, Nodejs6.10, Nodejs8.9, Nodejs10.15, Nodejs12.16, PHP5, PHP7, Golang1 and Java8
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// Function environment variable
	Environment *Environment `json:"Environment,omitempty" name:"Environment"`

	// Function namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Function VPC configuration
	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

	// Role bound to the function
	Role *string `json:"Role,omitempty" name:"Role"`

	// CLS logset ID to which logs are shipped
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// CLS Topic ID to which logs are shipped
	ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

	// It specifies whether to synchronously release a new version during the update. The default value is `FALSE`, indicating not to release a new version.
	Publish *string `json:"Publish,omitempty" name:"Publish"`

	// Whether to enable L5 access. TRUE: enable; FALSE: not enable
	L5Enable *string `json:"L5Enable,omitempty" name:"L5Enable"`

	// List of layer versions that bound with the function. Files with the same name will be overridden by the bound layer versions according to the ascending order in the list. 
	Layers []*LayerVersionSimple `json:"Layers,omitempty" name:"Layers" list`

	// Information of a dead letter queue associated with a function
	DeadLetterConfig *DeadLetterConfig `json:"DeadLetterConfig,omitempty" name:"DeadLetterConfig"`

	// Public network access configuration
	PublicNetConfig *PublicNetConfigIn `json:"PublicNetConfig,omitempty" name:"PublicNetConfig"`
}

func (r *UpdateFunctionConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFunctionConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFunctionConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFunctionConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFunctionConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNamespaceRequest struct {
	*tchttp.BaseRequest

	// Namespace name
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Namespace description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Variable struct {

	// Variable name
	Key *string `json:"Key,omitempty" name:"Key"`

	// Variable value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type VersionMatch struct {

	// Function version name
	Version *string `json:"Version,omitempty" name:"Version"`

	// Matching rule key. When the API is called, pass in the `key` to route the request to the specified version based on the matching rule
	// Header method:
	// Enter "invoke.headers.User" for `key` and pass in `RoutingKey:{"User":"value"}` when invoking a function through `invoke` for invocation based on rule matching
	Key *string `json:"Key,omitempty" name:"Key"`

	// Match method. Valid values:
	// range: range match
	// exact: exact string match
	Method *string `json:"Method,omitempty" name:"Method"`

	// Rule requirements for range match:
	// It should be described in an open or closed range, i.e., `(a,b)` or `[a,b]`, where both a and b are integers
	// Rule requirements for exact match:
	// Exact string match
	Expression *string `json:"Expression,omitempty" name:"Expression"`
}

type VersionWeight struct {

	// Function version name
	Version *string `json:"Version,omitempty" name:"Version"`

	// Version weight
	Weight *float64 `json:"Weight,omitempty" name:"Weight"`
}

type VpcConfig struct {

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
