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

package v20180808

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type APIDoc struct {
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`

	// API document name
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// API document build status
	ApiDocStatus *string `json:"ApiDocStatus,omitempty" name:"ApiDocStatus"`
}

type APIDocInfo struct {
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`

	// API document name
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// API document build status
	ApiDocStatus *string `json:"ApiDocStatus,omitempty" name:"ApiDocStatus"`

	// Number of APIs with API documents
	ApiCount *int64 `json:"ApiCount,omitempty" name:"ApiCount"`

	// Number of views of API document
	ViewCount *int64 `json:"ViewCount,omitempty" name:"ViewCount"`

	// Number of releases of API document
	ReleaseCount *int64 `json:"ReleaseCount,omitempty" name:"ReleaseCount"`

	// API document access URI
	ApiDocUri *string `json:"ApiDocUri,omitempty" name:"ApiDocUri"`

	// API document password for sharing
	SharePassword *string `json:"SharePassword,omitempty" name:"SharePassword"`

	// API document update time
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment information
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// ID of the API for which to generate the API document
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// Service name
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Name of the API for which to generate the API document
	ApiNames []*string `json:"ApiNames,omitempty" name:"ApiNames"`
}

type APIDocs struct {
	// Number of API documents
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Basic information of API document
	APIDocSet []*APIDoc `json:"APIDocSet,omitempty" name:"APIDocSet"`
}

type ApiAppApiInfo struct {
	// Application name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// Application ID
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// API ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Service ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Binding authorization time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthorizedTime *string `json:"AuthorizedTime,omitempty" name:"AuthorizedTime"`

	// API region
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`

	// Authorized binding environment
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
}

type ApiAppApiInfos struct {
	// Quantity
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Information array of the API bound to the application
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAppApiSet []*ApiAppApiInfo `json:"ApiAppApiSet,omitempty" name:"ApiAppApiSet"`
}

type ApiAppInfo struct {
	// Application name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// Application ID
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Application SECRET
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAppSecret *string `json:"ApiAppSecret,omitempty" name:"ApiAppSecret"`

	// Application description
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Modification time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Application KEY
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAppKey *string `json:"ApiAppKey,omitempty" name:"ApiAppKey"`
}

type ApiAppInfos struct {
	// Number of applications
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Application information array
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAppSet []*ApiAppInfo `json:"ApiAppSet,omitempty" name:"ApiAppSet"`
}

type ApiEnvironmentStrategy struct {
	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Custom API name.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API path, such as `/path`.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API method, such as `GET`.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Environment throttling information.
	EnvironmentStrategySet []*EnvironmentStrategy `json:"EnvironmentStrategySet,omitempty" name:"EnvironmentStrategySet"`
}

type ApiEnvironmentStrategyStataus struct {
	// Number of throttling policies bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of throttling policies bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiEnvironmentStrategySet []*ApiEnvironmentStrategy `json:"ApiEnvironmentStrategySet,omitempty" name:"ApiEnvironmentStrategySet"`
}

type ApiIdStatus struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API description
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API path.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API method.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Service creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Service modification time.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Unique VPC ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API protocol.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Whether to enable debugging after purchase.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Authorization type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// API business type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Unique ID of associated authorization API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// List of business APIs associated with authorization API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds"`

	// OAuth configuration information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Token storage position, which is an OAuth 2.0 API request.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`
}

type ApiInfo struct {
	// Unique service ID of API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Service name of API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Service description of API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Unique API ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API type. Valid values: NORMAL (general API), TSF (microservice API).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API frontend request type, such as HTTP, HTTPS, or HTTP and HTTPS.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// API authentication type. Valid values: SECRET (key pair authentication), NONE (no authentication), OAUTH.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// OAuth API type. Valid values: NORMAL (business API), OAUTH (authorization API).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Unique ID of the authorization API associated with OAuth business API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// OAuth configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Whether to enable debugging after purchase (reserved field for the marketplace).
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Request frontend configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// Return type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Sample response for successful custom response configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// Sample response for failed custom response configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// Custom error code configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseErrorCodes []*ErrorCodes `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// Frontend request parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

	// API backend service timeout period in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API backend service type. Valid values: HTTP, MOCK, TSF, CLB, SCF, WEBSOCKET, TARGET (in beta test).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// API backend service configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// API backend service parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// Constant parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

	// Returned message of API backend Mock, which is required if `ServiceType` is `Mock`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// SCF function name, which takes effect if the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// SCF function namespace, which takes effect if the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// SCF function version, which takes effect if the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// Whether integrated response is enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// SCF WebSocket registration function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// SCF WebSocket registration function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// SCF WebSocket cleanup function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// SCF WebSocket cleanup function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// SCF WebSocket cleanup function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// WebSocket callback address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InternalDomain *string `json:"InternalDomain,omitempty" name:"InternalDomain"`

	// SCF WebSocket transfer function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// SCF WebSocket transfer function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// List of microservices bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MicroServices []*MicroService `json:"MicroServices,omitempty" name:"MicroServices"`

	// Microservice details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MicroServicesInfo []*int64 `json:"MicroServicesInfo,omitempty" name:"MicroServicesInfo"`

	// Load balancing configuration of microservice.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// Health check configuration of microservice.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// Whether to enable CORS.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// Information of tags bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Environment information published for API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Environments []*string `json:"Environments,omitempty" name:"Environments"`

	// Whether to enable Base64 encoding. This parameter takes effect only when the backend is SCF.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// Whether to trigger Base64 encoding by header. This parameter takes effect only when the backend is SCF.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsBase64Trigger *bool `json:"IsBase64Trigger,omitempty" name:"IsBase64Trigger"`

	// Header trigger rules. The number of rules cannot exceed 10.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Base64EncodedTriggerRules []*Base64EncodedTriggerRule `json:"Base64EncodedTriggerRules,omitempty" name:"Base64EncodedTriggerRules"`
}

type ApiInfoSummary struct {
	// Total number of APIs that can use this plugin
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Information of the APIs that can use this plugin
	ApiSet []*AvailableApiInfo `json:"ApiSet,omitempty" name:"ApiSet"`
}

type ApiKey struct {
	// Created API key ID.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Created API key.
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`

	// Key type. Valid values: auto, manual.
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// Custom key name.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Key status. 0: disabled. 1: enabled.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ApiKeysStatus struct {
	// Number of eligible API keys.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API key list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiKeySet []*ApiKey `json:"ApiKeySet,omitempty" name:"ApiKeySet"`
}

type ApiRequestConfig struct {
	// path
	Path *string `json:"Path,omitempty" name:"Path"`

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ApiUsagePlan struct {
	// Unique service ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API method.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Unique usage plan ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Usage plan name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Usage plan description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Service environment bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Used quota.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

	// Total number of requests allowed. `-1` indicates no limit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Request QPS upper limit. `-1` indicates no limit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// Usage plan creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time of usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Service name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type ApiUsagePlanSet struct {
	// Total number of usage plans bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of usage plans bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiUsagePlanList []*ApiUsagePlan `json:"ApiUsagePlanList,omitempty" name:"ApiUsagePlanList"`
}

type ApisStatus struct {
	// Number of eligible APIs.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API list.
	ApiIdStatusSet []*DesApisStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`
}

// Predefined struct for user
type AttachPluginRequestParams struct {
	// ID of the plugin to be bound
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API environment
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs bound with the plugin
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type AttachPluginRequest struct {
	*tchttp.BaseRequest
	
	// ID of the plugin to be bound
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API environment
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs bound with the plugin
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *AttachPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachPluginResponseParams struct {
	// Whether binding succeeded.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachPluginResponse struct {
	*tchttp.BaseResponse
	Response *AttachPluginResponseParams `json:"Response"`
}

func (r *AttachPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedApiInfo struct {
	// ID of the service to which the API belongs
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the service to which the API belongs
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Description of the service to which the API belongs
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API name
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API description
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// Environment of the API bound with the plugin
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Time when the plugin was bound to the API
	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`
}

type AttachedApiSummary struct {
	// Number of APIs bound with the plugin
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Information of the API bound with the plugin
	AttachedApis []*AttachedApiInfo `json:"AttachedApis,omitempty" name:"AttachedApis"`
}

type AvailableApiInfo struct {
	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API name
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API type
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API path
	Path *string `json:"Path,omitempty" name:"Path"`

	// API method
	Method *string `json:"Method,omitempty" name:"Method"`

	// Whether the API is bound with another plugin
	// Note: this field may return null, indicating that no valid values can be obtained.
	AttachedOtherPlugin *bool `json:"AttachedOtherPlugin,omitempty" name:"AttachedOtherPlugin"`

	// Whether the API is bound with the current plugin
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsAttached *bool `json:"IsAttached,omitempty" name:"IsAttached"`
}

type Base64EncodedTriggerRule struct {
	// Header for triggering encoding. Valid values are `Accept` and `Content_Type`, corresponding to the `Accept` and `Content-Type` headers in the data stream request, respectively.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Array of header values that can trigger the encoding. Each element in the array can be up to 40 characters, including digits, letters, and special characters (`.`, `+`, `*`, `-`, `/`, and `_`). 
	// 
	// For example, [
	//     "application/x-vpeg005",
	//     "application/xhtml+xml",
	//     "application/vnd.ms-project",
	//     "application/vnd.rn-rn_music_package"
	// ] are valid.
	Value []*string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type BindApiAppRequestParams struct {
	// Unique ID of the application to be bound.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Environment to be bound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the API to be bound.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type BindApiAppRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the application to be bound.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Environment to be bound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the API to be bound.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *BindApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "Environment")
	delete(f, "ServiceId")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindApiAppResponseParams struct {
	// Whether binding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindApiAppResponse struct {
	*tchttp.BaseResponse
	Response *BindApiAppResponseParams `json:"Response"`
}

func (r *BindApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindApiInfo struct {
	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Unique ID of the service
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API name
	// Note: This field may return `null`, indicating that no valid value was found.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Service name
	// Note: This field may return `null`, indicating that no valid value was found.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Bound At
	BindTime *string `json:"BindTime,omitempty" name:"BindTime"`
}

// Predefined struct for user
type BindEnvironmentRequestParams struct {
	// List of unique IDs of the usage plans to be bound.
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`

	// Binding type. Valid values: API, SERVICE. Default value: SERVICE.
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// Environment to be bound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID array, which is required if `bindType` is `API`.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type BindEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// List of unique IDs of the usage plans to be bound.
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`

	// Binding type. Valid values: API, SERVICE. Default value: SERVICE.
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// Environment to be bound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID array, which is required if `bindType` is `API`.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *BindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanIds")
	delete(f, "BindType")
	delete(f, "Environment")
	delete(f, "ServiceId")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindEnvironmentResponseParams struct {
	// Whether binding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *BindEnvironmentResponseParams `json:"Response"`
}

func (r *BindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindIPStrategyRequestParams struct {
	// Unique service ID of the IP policy to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the IP policy to be bound.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Environment to be bound to IP policy.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs to be bound to IP policy.
	BindApiIds []*string `json:"BindApiIds,omitempty" name:"BindApiIds"`
}

type BindIPStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of the IP policy to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the IP policy to be bound.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Environment to be bound to IP policy.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs to be bound to IP policy.
	BindApiIds []*string `json:"BindApiIds,omitempty" name:"BindApiIds"`
}

func (r *BindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "EnvironmentName")
	delete(f, "BindApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindIPStrategyResponseParams struct {
	// Whether binding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *BindIPStrategyResponseParams `json:"Response"`
}

func (r *BindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSecretIdsRequestParams struct {
	// Unique ID of the usage plan to be bound.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Array of IDs of the keys to be bound.
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

type BindSecretIdsRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the usage plan to be bound.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Array of IDs of the keys to be bound.
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

func (r *BindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSecretIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "AccessKeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindSecretIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSecretIdsResponseParams struct {
	// Whether binding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *BindSecretIdsResponseParams `json:"Response"`
}

func (r *BindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSecretIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSubDomainRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name to be bound.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Protocol supported by service. Valid values: http, https, http&https.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Network type. Valid values: OUTER, INNER.
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Whether the default path mapping is used. The default value is `true`. If the value is `false`, the custom path mapping will be used and `PathMappingSet` will be required in this case.
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// Default domain name.
	NetSubDomain *string `json:"NetSubDomain,omitempty" name:"NetSubDomain"`

	// Unique certificate ID of the custom domain name to be bound. The certificate can be uploaded if `Protocol` is `https` or `http&https`.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Custom domain name path mapping. It can contain up to 3 `Environment` values which can be set to only `test`, `prepub`, and `release`, respectively.
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`

	// Whether to force HTTP requests to redirect to HTTPS. Default value: `false`. When this parameter is `true`, API Gateway will redirect all requests using the custom domain name over the HTTP protocol to the HTTPS protocol for forwarding.
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
}

type BindSubDomainRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name to be bound.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Protocol supported by service. Valid values: http, https, http&https.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Network type. Valid values: OUTER, INNER.
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Whether the default path mapping is used. The default value is `true`. If the value is `false`, the custom path mapping will be used and `PathMappingSet` will be required in this case.
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// Default domain name.
	NetSubDomain *string `json:"NetSubDomain,omitempty" name:"NetSubDomain"`

	// Unique certificate ID of the custom domain name to be bound. The certificate can be uploaded if `Protocol` is `https` or `http&https`.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Custom domain name path mapping. It can contain up to 3 `Environment` values which can be set to only `test`, `prepub`, and `release`, respectively.
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`

	// Whether to force HTTP requests to redirect to HTTPS. Default value: `false`. When this parameter is `true`, API Gateway will redirect all requests using the custom domain name over the HTTP protocol to the HTTPS protocol for forwarding.
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
}

func (r *BindSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSubDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	delete(f, "Protocol")
	delete(f, "NetType")
	delete(f, "IsDefaultMapping")
	delete(f, "NetSubDomain")
	delete(f, "CertificateId")
	delete(f, "PathMappingSet")
	delete(f, "IsForcedHttps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindSubDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSubDomainResponseParams struct {
	// Whether binding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *BindSubDomainResponseParams `json:"Response"`
}

func (r *BindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuildAPIDocRequestParams struct {
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

type BuildAPIDocRequest struct {
	*tchttp.BaseRequest
	
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *BuildAPIDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildAPIDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BuildAPIDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuildAPIDocResponseParams struct {
	// Whether the operation succeeded
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BuildAPIDocResponse struct {
	*tchttp.BaseResponse
	Response *BuildAPIDocResponseParams `json:"Response"`
}

func (r *BuildAPIDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConstantParameter struct {
	// Constant parameter name, which is used only if `ServiceType` is `HTTP`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Constant parameter description, which is used only if `ServiceType` is `HTTP`.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Constant parameter position, which is used only if `ServiceType` is `HTTP`.
	Position *string `json:"Position,omitempty" name:"Position"`

	// Default value of constant parameter, which is used only if `ServiceType` is `HTTP`.
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
}

type CosConfig struct {
	// Specifies how the backend COS bucket is called by the API. The frontend request method and Action can be:
	// GET：GetObject
	// PUT：PutObject
	// POST：PostObject、AppendObject
	// HEAD： HeadObject
	// DELETE： DeleteObject
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Backend COS bucket of the API
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// Whether to enable the backend COS signature for the API. It defaults to `false`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Authorization *bool `json:"Authorization,omitempty" name:"Authorization"`

	// The path matching mode of the backend COS service
	// `BackEndPath`: Match the backend path
	// `FullPath`: Match the full path
	// 
	// Default: `BackEndPath`
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PathMatchMode *string `json:"PathMatchMode,omitempty" name:"PathMatchMode"`
}

// Predefined struct for user
type CreateAPIDocRequestParams struct {
	// API document name
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// Service name
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// List of APIs for which to generate documents
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type CreateAPIDocRequest struct {
	*tchttp.BaseRequest
	
	// API document name
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// Service name
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// List of APIs for which to generate documents
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *CreateAPIDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAPIDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocName")
	delete(f, "ServiceId")
	delete(f, "Environment")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAPIDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAPIDocResponseParams struct {
	// Basic information of API document
	Result *APIDoc `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAPIDocResponse struct {
	*tchttp.BaseResponse
	Response *CreateAPIDocResponseParams `json:"Response"`
}

func (r *CreateAPIDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiAppRequestParams struct {
	// Custom application name.
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// Application description
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
}

type CreateApiAppRequest struct {
	*tchttp.BaseRequest
	
	// Custom application name.
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// Application description
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
}

func (r *CreateApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppName")
	delete(f, "ApiAppDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiAppResponseParams struct {
	// New application details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiAppInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApiAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiAppResponseParams `json:"Response"`
}

func (r *CreateApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiKeyRequestParams struct {
	// Custom key name.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Key type. Valid values: auto, manual (custom key). Default value: auto.
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// Custom key ID, which is required if `AccessKeyType` is `manual`. It can contain 5–50 letters, digits, and underscores.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Custom key, which is required if `AccessKeyType` is `manual`. It can contain 10–50 letters, digits, and underscores.
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

type CreateApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// Custom key name.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Key type. Valid values: auto, manual (custom key). Default value: auto.
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// Custom key ID, which is required if `AccessKeyType` is `manual`. It can contain 5–50 letters, digits, and underscores.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Custom key, which is required if `AccessKeyType` is `manual`. It can contain 10–50 letters, digits, and underscores.
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *CreateApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "AccessKeyType")
	delete(f, "AccessKeyId")
	delete(f, "AccessKeySecret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiKeyResponseParams struct {
	// New key details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiKey `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiKeyResponseParams `json:"Response"`
}

func (r *CreateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiRequestParams struct {
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API backend service type. Valid values: HTTP, MOCK, TSF, SCF, WEBSOCKET, TARGET (in beta test).
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// API backend service timeout period in seconds.
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API frontend request protocol. Valid values: HTTPS, WEBSOCKET.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Request frontend configuration.
	RequestConfig *ApiRequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// Custom API name.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Custom API description.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API type. Valid values: NORMAL (general API), TSF (microservice API). Default value: NORMAL.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API authentication type. Valid values: SECRET (key pair authentication), NONE (no authentication), OAUTH, APP (application authentication). Default value: NONE.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// Whether to enable CORS.
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// Constant parameter.
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

	// Frontend request parameter.
	RequestParameters []*RequestParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

	// This field is valid if `AuthType` is `OAUTH`. NORMAL: business API; OAUTH: authorization API.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Returned message of API backend Mock, which is required if `ServiceType` is `Mock`.
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// List of microservices bound to API.
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`

	// Load balancing configuration of microservice.
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// Health check configuration of microservice.
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// `target` type backend resource information (in beta test).
	TargetServices []*TargetServicesReq `json:"TargetServices,omitempty" name:"TargetServices"`

	// `target` type load balancing configuration (in beta test).
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// `target` health check configuration (in beta test).
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// SCF function name, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// SCF WebSocket registration function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// SCF WebSocket cleanup function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// SCF WebSocket transfer function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// SCF function namespace, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// SCF function version, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// SCF WebSocket registration function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// SCF WebSocket transfer function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// SCF WebSocket cleanup function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// SCF WebSocket cleanup function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// Whether to enable response integration, which takes effect if the backend type is `SCF`.
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// Billing after debugging starts (reserved field for marketplace).
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Whether to delete the error codes for custom response configuration. If the value is left empty or `False`, the error codes will not be deleted. If the value is `True`, all custom response configuration error codes of the API will be deleted.
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// Return type.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Sample response for successful custom response configuration.
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// Sample response for failed custom response configuration.
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// API backend service configuration.
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// Unique ID of associated authorization API, which takes effect only if `AuthType` is `OAUTH` and `ApiBusinessType` is `NORMAL`. It is the unique ID of the OAuth 2.0 authorization API bound to the business API.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// API backend service parameter.
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// OAuth configuration, which takes effect if `AuthType` is `OAUTH`.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Custom error code configuration.
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// TSF Serverless namespace ID (in beta test).
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// User type.
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// Whether to enable Base64 encoding. This parameter takes effect only when the backend is SCF.
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// Event bus ID.
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// SCF function type, which takes effect if the backend type is `SCF`. Valid values: `EVENT` and `HTTP`.
	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`

	// EIAM application type.
	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`

	// EIAM application authentication type. Valid values: `AuthenticationOnly`, `Authentication`, `Authorization`.
	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`

	// Validity of the EIAM application token. Unit: second. Default value: `7200`.
	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`

	// EIAM application ID.
	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`

	// Owner of the resource
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

type CreateApiRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API backend service type. Valid values: HTTP, MOCK, TSF, SCF, WEBSOCKET, TARGET (in beta test).
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// API backend service timeout period in seconds.
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API frontend request protocol. Valid values: HTTPS, WEBSOCKET.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Request frontend configuration.
	RequestConfig *ApiRequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// Custom API name.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Custom API description.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API type. Valid values: NORMAL (general API), TSF (microservice API). Default value: NORMAL.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API authentication type. Valid values: SECRET (key pair authentication), NONE (no authentication), OAUTH, APP (application authentication). Default value: NONE.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// Whether to enable CORS.
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// Constant parameter.
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

	// Frontend request parameter.
	RequestParameters []*RequestParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

	// This field is valid if `AuthType` is `OAUTH`. NORMAL: business API; OAUTH: authorization API.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Returned message of API backend Mock, which is required if `ServiceType` is `Mock`.
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// List of microservices bound to API.
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`

	// Load balancing configuration of microservice.
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// Health check configuration of microservice.
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// `target` type backend resource information (in beta test).
	TargetServices []*TargetServicesReq `json:"TargetServices,omitempty" name:"TargetServices"`

	// `target` type load balancing configuration (in beta test).
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// `target` health check configuration (in beta test).
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// SCF function name, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// SCF WebSocket registration function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// SCF WebSocket cleanup function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// SCF WebSocket transfer function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// SCF function namespace, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// SCF function version, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// SCF WebSocket registration function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// SCF WebSocket transfer function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// SCF WebSocket cleanup function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// SCF WebSocket cleanup function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// Whether to enable response integration, which takes effect if the backend type is `SCF`.
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// Billing after debugging starts (reserved field for marketplace).
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Whether to delete the error codes for custom response configuration. If the value is left empty or `False`, the error codes will not be deleted. If the value is `True`, all custom response configuration error codes of the API will be deleted.
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// Return type.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Sample response for successful custom response configuration.
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// Sample response for failed custom response configuration.
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// API backend service configuration.
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// Unique ID of associated authorization API, which takes effect only if `AuthType` is `OAUTH` and `ApiBusinessType` is `NORMAL`. It is the unique ID of the OAuth 2.0 authorization API bound to the business API.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// API backend service parameter.
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// OAuth configuration, which takes effect if `AuthType` is `OAUTH`.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Custom error code configuration.
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// TSF Serverless namespace ID (in beta test).
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// User type.
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// Whether to enable Base64 encoding. This parameter takes effect only when the backend is SCF.
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// Event bus ID.
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// SCF function type, which takes effect if the backend type is `SCF`. Valid values: `EVENT` and `HTTP`.
	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`

	// EIAM application type.
	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`

	// EIAM application authentication type. Valid values: `AuthenticationOnly`, `Authentication`, `Authorization`.
	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`

	// Validity of the EIAM application token. Unit: second. Default value: `7200`.
	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`

	// EIAM application ID.
	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`

	// Owner of the resource
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *CreateApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceType")
	delete(f, "ServiceTimeout")
	delete(f, "Protocol")
	delete(f, "RequestConfig")
	delete(f, "ApiName")
	delete(f, "ApiDesc")
	delete(f, "ApiType")
	delete(f, "AuthType")
	delete(f, "EnableCORS")
	delete(f, "ConstantParameters")
	delete(f, "RequestParameters")
	delete(f, "ApiBusinessType")
	delete(f, "ServiceMockReturnMessage")
	delete(f, "MicroServices")
	delete(f, "ServiceTsfLoadBalanceConf")
	delete(f, "ServiceTsfHealthCheckConf")
	delete(f, "TargetServices")
	delete(f, "TargetServicesLoadBalanceConf")
	delete(f, "TargetServicesHealthCheckConf")
	delete(f, "ServiceScfFunctionName")
	delete(f, "ServiceWebsocketRegisterFunctionName")
	delete(f, "ServiceWebsocketCleanupFunctionName")
	delete(f, "ServiceWebsocketTransportFunctionName")
	delete(f, "ServiceScfFunctionNamespace")
	delete(f, "ServiceScfFunctionQualifier")
	delete(f, "ServiceWebsocketRegisterFunctionNamespace")
	delete(f, "ServiceWebsocketRegisterFunctionQualifier")
	delete(f, "ServiceWebsocketTransportFunctionNamespace")
	delete(f, "ServiceWebsocketTransportFunctionQualifier")
	delete(f, "ServiceWebsocketCleanupFunctionNamespace")
	delete(f, "ServiceWebsocketCleanupFunctionQualifier")
	delete(f, "ServiceScfIsIntegratedResponse")
	delete(f, "IsDebugAfterCharge")
	delete(f, "IsDeleteResponseErrorCodes")
	delete(f, "ResponseType")
	delete(f, "ResponseSuccessExample")
	delete(f, "ResponseFailExample")
	delete(f, "ServiceConfig")
	delete(f, "AuthRelationApiId")
	delete(f, "ServiceParameters")
	delete(f, "OauthConfig")
	delete(f, "ResponseErrorCodes")
	delete(f, "TargetNamespaceId")
	delete(f, "UserType")
	delete(f, "IsBase64Encoded")
	delete(f, "EventBusId")
	delete(f, "ServiceScfFunctionType")
	delete(f, "EIAMAppType")
	delete(f, "EIAMAuthType")
	delete(f, "TokenTimeout")
	delete(f, "EIAMAppId")
	delete(f, "Owner")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiResponseParams struct {
	// API information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *CreateApiRsp `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApiResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiResponseParams `json:"Response"`
}

func (r *CreateApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApiRsp struct {
	// API ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Path
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// Request method
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Status of the import task
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Details of the error
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// API name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
}

// Predefined struct for user
type CreateIPStrategyRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom policy name.
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// Policy type. Valid values: WHITE (allowlist), BLACK (blocklist).
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// Policy details. Multiple IPs are separated with \n.
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

type CreateIPStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom policy name.
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// Policy type. Valid values: WHITE (allowlist), BLACK (blocklist).
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// Policy details. Multiple IPs are separated with \n.
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *CreateIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyName")
	delete(f, "StrategyType")
	delete(f, "StrategyData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIPStrategyResponseParams struct {
	// New IP policy details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *IPStrategy `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateIPStrategyResponseParams `json:"Response"`
}

func (r *CreateIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePluginRequestParams struct {
	// Custom plugin name. A plugin name should contain 2-50 characters out of a-z, A-Z, 0-9, and _, which must begin with a letter and end with a letter or a number.
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// Plugin type. Valid values: `IPControl`, `TrafficControl`, `Cors`, `CustomReq`, `CustomAuth`, `Routing`, `TrafficControlByParameter`, `CircuitBreaker`, `ProxyCache`
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// Plugin definition statement in json format
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`

	// Plugin description within 200 characters
	Description *string `json:"Description,omitempty" name:"Description"`

	// Label
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreatePluginRequest struct {
	*tchttp.BaseRequest
	
	// Custom plugin name. A plugin name should contain 2-50 characters out of a-z, A-Z, 0-9, and _, which must begin with a letter and end with a letter or a number.
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// Plugin type. Valid values: `IPControl`, `TrafficControl`, `Cors`, `CustomReq`, `CustomAuth`, `Routing`, `TrafficControlByParameter`, `CircuitBreaker`, `ProxyCache`
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// Plugin definition statement in json format
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`

	// Plugin description within 200 characters
	Description *string `json:"Description,omitempty" name:"Description"`

	// Label
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreatePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginName")
	delete(f, "PluginType")
	delete(f, "PluginData")
	delete(f, "Description")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePluginResponseParams struct {
	// Details of the new plugin
	Result *Plugin `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePluginResponse struct {
	*tchttp.BaseResponse
	Response *CreatePluginResponseParams `json:"Response"`
}

func (r *CreatePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceRequestParams struct {
	// Custom service name.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Service frontend request type, such as `http`, `https`, and `http&https`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Custom service description.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Dedicated cluster name, which is used to specify the dedicated cluster where the service is to be created.
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// Network type list, which is used to specify the supported network types. INNER: private network access; OUTER: public network access. Default value: OUTER.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// IP version number. Valid values: IPv4, IPv6. Default value: IPv4.
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// Cluster name, which is reserved and used by the `tsf serverless` type.
	SetServerName *string `json:"SetServerName,omitempty" name:"SetServerName"`

	// User type, which is reserved and can be used by `serverless` users.
	AppIdType *string `json:"AppIdType,omitempty" name:"AppIdType"`

	// Tag information.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Dedicated instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// VPC attribute
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest
	
	// Custom service name.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Service frontend request type, such as `http`, `https`, and `http&https`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Custom service description.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Dedicated cluster name, which is used to specify the dedicated cluster where the service is to be created.
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// Network type list, which is used to specify the supported network types. INNER: private network access; OUTER: public network access. Default value: OUTER.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// IP version number. Valid values: IPv4, IPv6. Default value: IPv4.
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// Cluster name, which is reserved and used by the `tsf serverless` type.
	SetServerName *string `json:"SetServerName,omitempty" name:"SetServerName"`

	// User type, which is reserved and can be used by `serverless` users.
	AppIdType *string `json:"AppIdType,omitempty" name:"AppIdType"`

	// Tag information.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Dedicated instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// VPC attribute
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
}

func (r *CreateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceName")
	delete(f, "Protocol")
	delete(f, "ServiceDesc")
	delete(f, "ExclusiveSetName")
	delete(f, "NetTypes")
	delete(f, "IpVersion")
	delete(f, "SetServerName")
	delete(f, "AppIdType")
	delete(f, "Tags")
	delete(f, "InstanceId")
	delete(f, "UniqVpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceResponseParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom service name.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Custom service description.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Default public domain name.
	OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

	// Default VPC domain name.
	InnerSubDomain *string `json:"InnerSubDomain,omitempty" name:"InnerSubDomain"`

	// Service creation time in the format of `YYYY-MM-DDThh:mm:ssZ` according to ISO 8601 standard. UTC time is used.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Network type list. INNER: private network access; OUTER: public network access.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// IP version number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceResponseParams `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUpstreamRequestParams struct {
	// Backend protocol. Valid values: `HTTP`, `HTTPS`
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Load balancing algorithm. Valid value: `ROUND-ROBIN`
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Unique VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Upstream name
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// Upstream description
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// Upstream access type. Valid values: `IP_PORT`, `K8S`
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// Retry attempts. It defaults to `3`.
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// The Host request header that forwarded from the gateway to backend
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`

	// Backend nodes
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// Label
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Health check configuration
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// Configuration of TKE service
	K8sService []*K8sService `json:"K8sService,omitempty" name:"K8sService"`
}

type CreateUpstreamRequest struct {
	*tchttp.BaseRequest
	
	// Backend protocol. Valid values: `HTTP`, `HTTPS`
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Load balancing algorithm. Valid value: `ROUND-ROBIN`
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Unique VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Upstream name
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// Upstream description
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// Upstream access type. Valid values: `IP_PORT`, `K8S`
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// Retry attempts. It defaults to `3`.
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// The Host request header that forwarded from the gateway to backend
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`

	// Backend nodes
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// Label
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Health check configuration
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// Configuration of TKE service
	K8sService []*K8sService `json:"K8sService,omitempty" name:"K8sService"`
}

func (r *CreateUpstreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUpstreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Scheme")
	delete(f, "Algorithm")
	delete(f, "UniqVpcId")
	delete(f, "UpstreamName")
	delete(f, "UpstreamDescription")
	delete(f, "UpstreamType")
	delete(f, "Retries")
	delete(f, "UpstreamHost")
	delete(f, "Nodes")
	delete(f, "Tags")
	delete(f, "HealthChecker")
	delete(f, "K8sService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUpstreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUpstreamResponseParams struct {
	// The unique upstream ID returned
	// Note: This field may return `NULL`, indicating that no valid value was found.
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUpstreamResponse struct {
	*tchttp.BaseResponse
	Response *CreateUpstreamResponseParams `json:"Response"`
}

func (r *CreateUpstreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsagePlanRequestParams struct {
	// Custom usage plan name.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Custom usage plan description.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is `-1`, which indicates no limit.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is `-1`, which indicates no limit.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

type CreateUsagePlanRequest struct {
	*tchttp.BaseRequest
	
	// Custom usage plan name.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Custom usage plan description.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is `-1`, which indicates no limit.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is `-1`, which indicates no limit.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *CreateUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanName")
	delete(f, "UsagePlanDesc")
	delete(f, "MaxRequestNum")
	delete(f, "MaxRequestNumPreSec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsagePlanResponseParams struct {
	// Usage plan details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateUsagePlanResponseParams `json:"Response"`
}

func (r *CreateUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAPIDocRequestParams struct {
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

type DeleteAPIDocRequest struct {
	*tchttp.BaseRequest
	
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *DeleteAPIDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAPIDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAPIDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAPIDocResponseParams struct {
	// Whether the operation succeeded
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAPIDocResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAPIDocResponseParams `json:"Response"`
}

func (r *DeleteAPIDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiAppRequestParams struct {
	// Unique application ID.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

type DeleteApiAppRequest struct {
	*tchttp.BaseRequest
	
	// Unique application ID.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

func (r *DeleteApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiAppResponseParams struct {
	// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApiAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApiAppResponseParams `json:"Response"`
}

func (r *DeleteApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiKeyRequestParams struct {
	// ID of the key to be deleted.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

type DeleteApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the key to be deleted.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DeleteApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiKeyResponseParams struct {
	// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApiKeyResponseParams `json:"Response"`
}

func (r *DeleteApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiRequestParams struct {
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type DeleteApiRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DeleteApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiResponseParams struct {
	// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApiResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApiResponseParams `json:"Response"`
}

func (r *DeleteApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIPStrategyRequestParams struct {
	// Unique service ID of the IP policy to be deleted.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the IP policy to be deleted.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

type DeleteIPStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of the IP policy to be deleted.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the IP policy to be deleted.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DeleteIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIPStrategyResponseParams struct {
	// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIPStrategyResponseParams `json:"Response"`
}

func (r *DeleteIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePluginRequestParams struct {
	// ID of the plugin to be deleted
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
}

type DeletePluginRequest struct {
	*tchttp.BaseRequest
	
	// ID of the plugin to be deleted
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
}

func (r *DeletePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePluginResponseParams struct {
	// Result of the deletion action
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePluginResponse struct {
	*tchttp.BaseResponse
	Response *DeletePluginResponseParams `json:"Response"`
}

func (r *DeletePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceRequestParams struct {
	// Unique ID of the service to be deleted.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// A parameter used to set to skip the deletion precondition verification (only supported for services on dedicated instances).
	SkipVerification *int64 `json:"SkipVerification,omitempty" name:"SkipVerification"`
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be deleted.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// A parameter used to set to skip the deletion precondition verification (only supported for services on dedicated instances).
	SkipVerification *int64 `json:"SkipVerification,omitempty" name:"SkipVerification"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SkipVerification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceResponseParams struct {
	// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceResponseParams `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceSubDomainMappingRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name bound to service.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Name of the environment whose mapping is to be deleted. Valid values: test (test environment), prepub (pre-release environment), release (release environment).
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type DeleteServiceSubDomainMappingRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name bound to service.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Name of the environment whose mapping is to be deleted. Valid values: test (test environment), prepub (pre-release environment), release (release environment).
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DeleteServiceSubDomainMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceSubDomainMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceSubDomainMappingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceSubDomainMappingResponseParams struct {
	// Whether the path mapping of the custom domain name is successfully deleted.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteServiceSubDomainMappingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceSubDomainMappingResponseParams `json:"Response"`
}

func (r *DeleteServiceSubDomainMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceSubDomainMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUpstreamRequestParams struct {
	// ID of the upstream to be deleted
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
}

type DeleteUpstreamRequest struct {
	*tchttp.BaseRequest
	
	// ID of the upstream to be deleted
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
}

func (r *DeleteUpstreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUpstreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UpstreamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUpstreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUpstreamResponseParams struct {
	// ID of the deleted upstream
	// Note: This field may return `NULL`, indicating that no valid value was found.
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUpstreamResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUpstreamResponseParams `json:"Response"`
}

func (r *DeleteUpstreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsagePlanRequestParams struct {
	// Unique ID of the usage plan to be deleted.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

type DeleteUsagePlanRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the usage plan to be deleted.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DeleteUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsagePlanResponseParams struct {
	// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsagePlanResponseParams `json:"Response"`
}

func (r *DeleteUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DemoteServiceUsagePlanRequestParams struct {
	// Usage plan ID.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Unique ID of the service to be demoted.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name.
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type DemoteServiceUsagePlanRequest struct {
	*tchttp.BaseRequest
	
	// Usage plan ID.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Unique ID of the service to be demoted.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name.
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DemoteServiceUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DemoteServiceUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "ServiceId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DemoteServiceUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DemoteServiceUsagePlanResponseParams struct {
	// Whether demotion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DemoteServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DemoteServiceUsagePlanResponseParams `json:"Response"`
}

func (r *DemoteServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DemoteServiceUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DesApisStatus struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Custom API description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// VPC ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Unique VPC ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API type. Valid values: NORMAL (general API), TSF (microservice API).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API protocol.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Whether to enable debugging after purchase (reserved field for the marketplace)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// API authentication type. Valid values: `SECRET` (key pair authentication), `NONE` (no authentication), `OAUTH`, and `EIAM`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// OAuth API type, which is valid if `AuthType` is `OAUTH`. Valid values: NORMAL (business API), OAUTH (authorization API).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Unique ID of associated authorization API, which takes effect only if `AuthType` is `OAUTH` and `ApiBusinessType` is `NORMAL`. It is the unique ID of the OAuth 2.0 authorization API bound to the business API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// OAuth configuration information, which takes effect if `AuthType` is `OAUTH`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// List of business APIs associated with authorization API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds"`

	// Information of tags associated with API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// API path, such as `/path`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API request method, such as `GET`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`
}

// Predefined struct for user
type DescribeAPIDocDetailRequestParams struct {
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

type DescribeAPIDocDetailRequest struct {
	*tchttp.BaseRequest
	
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *DescribeAPIDocDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIDocDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAPIDocDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIDocDetailResponseParams struct {
	// API document details
	Result *APIDocInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAPIDocDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAPIDocDetailResponseParams `json:"Response"`
}

func (r *DescribeAPIDocDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIDocDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIDocsRequestParams struct {
	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAPIDocsRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAPIDocsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIDocsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAPIDocsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIDocsResponseParams struct {
	// API document list information
	Result *APIDocs `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAPIDocsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAPIDocsResponseParams `json:"Response"`
}

func (r *DescribeAPIDocsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIDocsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllPluginApisRequestParams struct {
	// ID of the service to be queried
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// ID of the plugin to be queried
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Environment information
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAllPluginApisRequest struct {
	*tchttp.BaseRequest
	
	// ID of the service to be queried
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// ID of the plugin to be queried
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Environment information
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAllPluginApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllPluginApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "PluginId")
	delete(f, "EnvironmentName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllPluginApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllPluginApisResponseParams struct {
	// List of APIs that ca use this plugin
	Result *ApiInfoSummary `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllPluginApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllPluginApisResponseParams `json:"Response"`
}

func (r *DescribeAllPluginApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllPluginApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppBindApisStatusRequestParams struct {
	// Application ID
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: ApiId, ApiName, ServiceId, Environment, KeyWord (match with `name` or ID).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApiAppBindApisStatusRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: ApiId, ApiName, ServiceId, Environment, KeyWord (match with `name` or ID).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiAppBindApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppBindApisStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiAppBindApisStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppBindApisStatusResponseParams struct {
	// List of APIs bound to the application.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiAppApiInfos `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiAppBindApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiAppBindApisStatusResponseParams `json:"Response"`
}

func (r *DescribeApiAppBindApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppBindApisStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppRequestParams struct {
	// Application ID.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

type DescribeApiAppRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

func (r *DescribeApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppResponseParams struct {
	// Application details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiAppInfos `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiAppResponseParams `json:"Response"`
}

func (r *DescribeApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppsStatusRequestParams struct {
	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: ApiAppId, ApiAppName, KeyWord (match with `name` or ID).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApiAppsStatusRequest struct {
	*tchttp.BaseRequest
	
	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: ApiAppId, ApiAppName, KeyWord (match with `name` or ID).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiAppsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiAppsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppsStatusResponseParams struct {
	// Application list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiAppInfos `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiAppsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiAppsStatusResponseParams `json:"Response"`
}

func (r *DescribeApiAppsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiBindApiAppsStatusRequestParams struct {
	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Array of API IDs
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: ApiAppId, Environment, KeyWord (match with `name` or ID).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApiBindApiAppsStatusRequest struct {
	*tchttp.BaseRequest
	
	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Array of API IDs
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: ApiAppId, Environment, KeyWord (match with `name` or ID).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiBindApiAppsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiBindApiAppsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiBindApiAppsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiBindApiAppsStatusResponseParams struct {
	// List of APIs bound to the application.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiAppApiInfos `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiBindApiAppsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiBindApiAppsStatusResponseParams `json:"Response"`
}

func (r *DescribeApiBindApiAppsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiBindApiAppsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiEnvironmentStrategyRequestParams struct {
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment list.
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment list.
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiEnvironmentStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentNames")
	delete(f, "ApiId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiEnvironmentStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiEnvironmentStrategyResponseParams struct {
	// Details of policies bound to API
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiEnvironmentStrategyStataus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiEnvironmentStrategyResponseParams `json:"Response"`
}

func (r *DescribeApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiForApiAppRequestParams struct {
	// Unique service ID of the API
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API region
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

type DescribeApiForApiAppRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of the API
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API region
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

func (r *DescribeApiForApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiForApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiId")
	delete(f, "ApiRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiForApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiForApiAppResponseParams struct {
	// API details.
	Result *ApiInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiForApiAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiForApiAppResponseParams `json:"Response"`
}

func (r *DescribeApiForApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiForApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyRequestParams struct {
	// API key ID.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

type DescribeApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// API key ID.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DescribeApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyResponseParams struct {
	// Key details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiKey `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiKeyResponseParams `json:"Response"`
}

func (r *DescribeApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeysStatusRequestParams struct {
	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: AccessKeyId, AccessKeySecret, SecretName, NotUsagePlanId, Status, KeyWord (match with `name` or `path`).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApiKeysStatusRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: AccessKeyId, AccessKeySecret, SecretName, NotUsagePlanId, Status, KeyWord (match with `name` or `path`).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiKeysStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeysStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiKeysStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeysStatusResponseParams struct {
	// Key list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiKeysStatus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiKeysStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiKeysStatusResponseParams `json:"Response"`
}

func (r *DescribeApiKeysStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeysStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiRequestParams struct {
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type DescribeApiRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DescribeApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiResponseParams struct {
	// API details.
	Result *ApiInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiResponseParams `json:"Response"`
}

func (r *DescribeApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiUsagePlanRequestParams struct {
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeApiUsagePlanRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiUsagePlanResponseParams struct {
	// List of usage plans bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiUsagePlanSet `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiUsagePlanResponseParams `json:"Response"`
}

func (r *DescribeApiUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApisStatusRequestParams struct {
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// API filter. Valid values: ApiId, ApiName, ApiPath, ApiType, AuthRelationApiId, AuthType, ApiBuniessType, NotUsagePlanId, Environment, Tags (whose values are the list of `$tag_key:tag_value`), TagKeys (whose values are the list of tag keys).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApisStatusRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// API filter. Valid values: ApiId, ApiName, ApiPath, ApiType, AuthRelationApiId, AuthType, ApiBuniessType, NotUsagePlanId, Environment, Tags (whose values are the list of `$tag_key:tag_value`), TagKeys (whose values are the list of tag keys).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApisStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApisStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApisStatusResponseParams struct {
	// List of API details.
	Result *ApisStatus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApisStatusResponseParams `json:"Response"`
}

func (r *DescribeApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApisStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategyApisStatusRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique policy ID.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Policy environment.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: ApiPath, ApiName, KeyWord (fuzzy search by `Path` and `Name`).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeIPStrategyApisStatusRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique policy ID.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Policy environment.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: ApiPath, ApiName, KeyWord (fuzzy search by `Path` and `Name`).
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPStrategyApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategyApisStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "EnvironmentName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPStrategyApisStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategyApisStatusResponseParams struct {
	// List of APIs bound to environment.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *IPStrategyApiStatus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIPStrategyApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPStrategyApisStatusResponseParams `json:"Response"`
}

func (r *DescribeIPStrategyApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategyApisStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategyRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique IP policy ID.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Environment associated with policy.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter, which is a reserved field. Filtering is not supported currently.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeIPStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique IP policy ID.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Environment associated with policy.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter, which is a reserved field. Filtering is not supported currently.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "EnvironmentName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategyResponseParams struct {
	// IP policy details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *IPStrategy `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPStrategyResponseParams `json:"Response"`
}

func (r *DescribeIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategysStatusRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Filter. Valid values: StrategyName.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeIPStrategysStatusRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Filter. Valid values: StrategyName.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPStrategysStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategysStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPStrategysStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategysStatusResponseParams struct {
	// List of eligible policies.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *IPStrategysStatus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIPStrategysStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPStrategysStatusResponseParams `json:"Response"`
}

func (r *DescribeIPStrategysStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategysStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogSearchRequestParams struct {
	// Log start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Reserved field
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Number of logs to be returned at a time. Maximum value: 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Subsequent content can be obtained based on the `ConText` returned last time. Up to 10,000 data entries can be obtained
	ConText *string `json:"ConText,omitempty" name:"ConText"`

	// Sorting by time. Valid values: asc (ascending), desc (descending). Default value: desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// Reserved field
	Query *string `json:"Query,omitempty" name:"Query"`

	// Search criterion. Valid values:
	// req_id: "="
	// api_id: "="
	// cip: "="
	// uip: ":"
	// err_msg: ":"
	// rsp_st: "=", "!=", ":", ">", "<"
	// req_t: ">=", "<="
	// 
	// Note:
	// ":" indicates included, and "!=" indicates not equal to. For the meanings of fields, please see the `LogSet` description of the output parameter
	LogQuerys []*LogQuery `json:"LogQuerys,omitempty" name:"LogQuerys"`
}

type DescribeLogSearchRequest struct {
	*tchttp.BaseRequest
	
	// Log start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Reserved field
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Number of logs to be returned at a time. Maximum value: 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Subsequent content can be obtained based on the `ConText` returned last time. Up to 10,000 data entries can be obtained
	ConText *string `json:"ConText,omitempty" name:"ConText"`

	// Sorting by time. Valid values: asc (ascending), desc (descending). Default value: desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// Reserved field
	Query *string `json:"Query,omitempty" name:"Query"`

	// Search criterion. Valid values:
	// req_id: "="
	// api_id: "="
	// cip: "="
	// uip: ":"
	// err_msg: ":"
	// rsp_st: "=", "!=", ":", ">", "<"
	// req_t: ">=", "<="
	// 
	// Note:
	// ":" indicates included, and "!=" indicates not equal to. For the meanings of fields, please see the `LogSet` description of the output parameter
	LogQuerys []*LogQuery `json:"LogQuerys,omitempty" name:"LogQuerys"`
}

func (r *DescribeLogSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSearchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ServiceId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "ConText")
	delete(f, "Sort")
	delete(f, "Query")
	delete(f, "LogQuerys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogSearchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogSearchResponseParams struct {
	// Cursor for getting more search results. If the value is `""`, there will be no subsequent results
	ConText *string `json:"ConText,omitempty" name:"ConText"`

	// The returned result contains any number of logs, which are in the following format:
	// '[$app_id][$env_name][$service_id][$http_host][$api_id][$uri][$scheme][rsp_st:$status][ups_st:$upstream_status]'
	// '[cip:$remote_addr][uip:$upstream_addr][vip:$server_addr][rsp_len:$bytes_sent][req_len:$request_length]'
	// '[req_t:$request_time][ups_rsp_t:$upstream_response_time][ups_conn_t:$upstream_connect_time][ups_head_t:$upstream_header_time]’
	// '[err_msg:$err_msg][tcp_rtt:$tcpinfo_rtt][$pid][$time_local][req_id:$request_id]';
	// 
	// Note:
	// app_id: user ID.
	// env_name: environment name.
	// service_id: service ID.
	// http_host: domain name.
	// api_id: API ID.
	// uri: request path.
	// scheme: HTTP/HTTPS protocol.
	// rsp_st: request response status code.
	// ups_st: backend business server response status code (if the request is passed through to the backend, this variable will not be empty. If the request is blocked in API Gateway, this variable will be displayed as `-`).
	// cip: client IP.
	// uip: backend business service (upstream) IP.
	// vip: VIP requested to be accessed.
	// rsp_len: response length.
	// req_len: request length.
	// req_t: total request response time.
	// ups_rsp_t: total backend response time (time between connection establishment by API Gateway and backend response receipt).
	// ups_conn_t: time when the backend business server is successfully connected to.
	// ups_head_t: time when the backend response header arrives.
	// err_msg: error message.
	// tcp_rtt: client TCP connection information. RTT (Round Trip Time) consists of three parts: link propagation delay, end system processing delay, and queuing delay in router cache.
	// req_id: request ID.
	LogSet []*string `json:"LogSet,omitempty" name:"LogSet"`

	// Number of logs returned for one search (`TotalCount <= Limit`)
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogSearchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogSearchResponseParams `json:"Response"`
}

func (r *DescribeLogSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginApisRequestParams struct {
	// ID of the plugin to be queried
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribePluginApisRequest struct {
	*tchttp.BaseRequest
	
	// ID of the plugin to be queried
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePluginApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginApisResponseParams struct {
	// List of APIs bound with the plugin
	Result *AttachedApiSummary `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePluginApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginApisResponseParams `json:"Response"`
}

func (r *DescribePluginApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginRequestParams struct {
	// ID of the plugin to be queried
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribePluginRequest struct {
	*tchttp.BaseRequest
	
	// ID of the plugin to be queried
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginResponseParams struct {
	// Plugin details
	Result *Plugin `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePluginResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginResponseParams `json:"Response"`
}

func (r *DescribePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentListRequestParams struct {
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeServiceEnvironmentListRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceEnvironmentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentListResponseParams struct {
	// Details of environments bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ServiceEnvironmentSet `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceEnvironmentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceEnvironmentListResponseParams `json:"Response"`
}

func (r *DescribeServiceEnvironmentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentReleaseHistoryRequestParams struct {
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeServiceEnvironmentReleaseHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentReleaseHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceEnvironmentReleaseHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentReleaseHistoryResponseParams struct {
	// Service release history.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ServiceReleaseHistory `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceEnvironmentReleaseHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceEnvironmentReleaseHistoryResponseParams `json:"Response"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentReleaseHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentStrategyRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceEnvironmentStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentStrategyResponseParams struct {
	// Throttling policy list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ServiceEnvironmentStrategyStatus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceEnvironmentStrategyResponseParams `json:"Response"`
}

func (r *DescribeServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceForApiAppRequestParams struct {
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Service region.
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

type DescribeServiceForApiAppRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Service region.
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

func (r *DescribeServiceForApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceForApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceForApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceForApiAppResponseParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Service environment list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments"`

	// Service name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Service description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Protocol supported by service. Valid values: http, https, http&https.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Service creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Service modification time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Self-Deployed cluster name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// Network type list. INNER: private network access; OUTER: public network access.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// Subdomain name for private network access.
	InternalSubDomain *string `json:"InternalSubDomain,omitempty" name:"InternalSubDomain"`

	// Subdomain name for public network access.
	OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

	// Service port number for HTTP access over private network.
	InnerHttpPort *int64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`

	// Port number for HTTPS access over private network.
	InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`

	// Total number of APIs.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiTotalCount *int64 `json:"ApiTotalCount,omitempty" name:"ApiTotalCount"`

	// API list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiIdStatusSet []*ApiIdStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`

	// Total number of usage plans.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanTotalCount *int64 `json:"UsagePlanTotalCount,omitempty" name:"UsagePlanTotalCount"`

	// Usage plan array.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanList []*UsagePlan `json:"UsagePlanList,omitempty" name:"UsagePlanList"`

	// IP version.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// Service user type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// Reserved field.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SetId *int64 `json:"SetId,omitempty" name:"SetId"`

	// Tag bound to the service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceForApiAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceForApiAppResponseParams `json:"Response"`
}

func (r *DescribeServiceForApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceForApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceReleaseVersionRequestParams struct {
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeServiceReleaseVersionRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceReleaseVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceReleaseVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceReleaseVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceReleaseVersionResponseParams struct {
	// Service release version list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ServiceReleaseVersion `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceReleaseVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceReleaseVersionResponseParams `json:"Response"`
}

func (r *DescribeServiceReleaseVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceReleaseVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceRequestParams struct {
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DescribeServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceResponseParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Service environment list.
	AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments"`

	// Service name.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Service description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Protocol supported by service. Valid values: http, https, http&https.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Service creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Service modification time.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Dedicated cluster name.
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// Network type list. INNER: private network access; OUTER: public network access.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// Subdomain name for private network access.
	InternalSubDomain *string `json:"InternalSubDomain,omitempty" name:"InternalSubDomain"`

	// Subdomain name for public network access.
	OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

	// Service port number for HTTP access over private network.
	InnerHttpPort *int64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`

	// Port number for HTTPS access over private network.
	InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`

	// Total number of APIs.
	ApiTotalCount *int64 `json:"ApiTotalCount,omitempty" name:"ApiTotalCount"`

	// API list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiIdStatusSet []*ApiIdStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`

	// Total number of usage plans.
	UsagePlanTotalCount *int64 `json:"UsagePlanTotalCount,omitempty" name:"UsagePlanTotalCount"`

	// Usage plan array.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanList []*UsagePlan `json:"UsagePlanList,omitempty" name:"UsagePlanList"`

	// IP version.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// Service user type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// Reserved field.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SetId *int64 `json:"SetId,omitempty" name:"SetId"`

	// Tags bound to a service.
	// Note: this field may return null, indicating that no valid values found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Dedicated instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Dedicated instance name
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Cluster type
	// Note: this field may return null, indicating that no valid values can be obtained.
	SetType *string `json:"SetType,omitempty" name:"SetType"`

	// Cluster type for service deployment
	// Note: this field may return null, indicating that no valid values found.
	DeploymentType *string `json:"DeploymentType,omitempty" name:"DeploymentType"`

	// Whether the service if for special usage. Valid values: `DEFAULT` (general usage), `HTTP_DNS`.
	// Note: This field may return `NULL`, indicating that no valid value was found.
	SpecialUse *string `json:"SpecialUse,omitempty" name:"SpecialUse"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceResponseParams `json:"Response"`
}

func (r *DescribeServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceSubDomainMappingsRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name bound to service.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

type DescribeServiceSubDomainMappingsRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name bound to service.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *DescribeServiceSubDomainMappingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceSubDomainMappingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceSubDomainMappingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceSubDomainMappingsResponseParams struct {
	// Custom path mapping list.
	Result *ServiceSubDomainMappings `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceSubDomainMappingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceSubDomainMappingsResponseParams `json:"Response"`
}

func (r *DescribeServiceSubDomainMappingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceSubDomainMappingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceSubDomainsRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeServiceSubDomainsRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceSubDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceSubDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceSubDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceSubDomainsResponseParams struct {
	// Custom service domain name query.
	Result *DomainSets `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceSubDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceSubDomainsResponseParams `json:"Response"`
}

func (r *DescribeServiceSubDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceSubDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceUsagePlanRequestParams struct {
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeServiceUsagePlanRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceUsagePlanResponseParams struct {
	// List of usage plans bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ServiceUsagePlanSet `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceUsagePlanResponseParams `json:"Response"`
}

func (r *DescribeServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServicesStatusRequestParams struct {
	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: `ServiceId`, `ServiceName`, `NotUsagePlanId`, `Environment`, `IpVersion`, `InstanceId`, `NetType`, `EIAMAppId`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeServicesStatusRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. Valid values: `ServiceId`, `ServiceName`, `NotUsagePlanId`, `Environment`, `IpVersion`, `InstanceId`, `NetType`, `EIAMAppId`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeServicesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServicesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServicesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServicesStatusResponseParams struct {
	// Service list query result.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ServicesStatus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServicesStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServicesStatusResponseParams `json:"Response"`
}

func (r *DescribeServicesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServicesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpstreamBindApis struct {
	// Total number
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Information of bound APIs
	BindApiSet []*BindApiInfo `json:"BindApiSet,omitempty" name:"BindApiSet"`
}

// Predefined struct for user
type DescribeUpstreamBindApisRequestParams struct {
	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The starting position of paging
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Upstream ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// Filters the results by `ServiceId` and `ApiId`
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeUpstreamBindApisRequest struct {
	*tchttp.BaseRequest
	
	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The starting position of paging
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Upstream ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// Filters the results by `ServiceId` and `ApiId`
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUpstreamBindApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamBindApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "UpstreamId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpstreamBindApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpstreamBindApisResponseParams struct {
	// Query results
	Result *DescribeUpstreamBindApis `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUpstreamBindApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpstreamBindApisResponseParams `json:"Response"`
}

func (r *DescribeUpstreamBindApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamBindApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpstreamInfo struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of query result
	UpstreamSet []*UpstreamInfo `json:"UpstreamSet,omitempty" name:"UpstreamSet"`
}

// Predefined struct for user
type DescribeUpstreamsRequestParams struct {
	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The starting position of paging
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filters. Valid values: `UpstreamId` and `UpstreamName`
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeUpstreamsRequest struct {
	*tchttp.BaseRequest
	
	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The starting position of paging
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filters. Valid values: `UpstreamId` and `UpstreamName`
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUpstreamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpstreamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpstreamsResponseParams struct {
	// Query results
	Result *DescribeUpstreamInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUpstreamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpstreamsResponseParams `json:"Response"`
}

func (r *DescribeUpstreamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanEnvironmentsRequestParams struct {
	// Unique ID of the usage plan to be queried.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Binding type. Valid values: API, SERVICE. Default value: SERVICE.
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeUsagePlanEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the usage plan to be queried.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Binding type. Valid values: API, SERVICE. Default value: SERVICE.
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "BindType")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsagePlanEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanEnvironmentsResponseParams struct {
	// Usage plan binding details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *UsagePlanEnvironmentStatus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsagePlanEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsagePlanEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeUsagePlanEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanRequestParams struct {
	// Unique ID of the usage plan to be queried.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

type DescribeUsagePlanRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the usage plan to be queried.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DescribeUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanResponseParams struct {
	// Usage plan details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsagePlanResponseParams `json:"Response"`
}

func (r *DescribeUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanSecretIdsRequestParams struct {
	// Unique ID of bound usage plan.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeUsagePlanSecretIdsRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of bound usage plan.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanSecretIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsagePlanSecretIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanSecretIdsResponseParams struct {
	// List of keys bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *UsagePlanBindSecretStatus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsagePlanSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsagePlanSecretIdsResponseParams `json:"Response"`
}

func (r *DescribeUsagePlanSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanSecretIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlansStatusRequestParams struct {
	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Usage plan filter. Valid values: UsagePlanId, UsagePlanName, NotServiceId, NotApiId, Environment.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeUsagePlansStatusRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Usage plan filter. Valid values: UsagePlanId, UsagePlanName, NotServiceId, NotApiId, Environment.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUsagePlansStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlansStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsagePlansStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlansStatusResponseParams struct {
	// Usage plan list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *UsagePlansStatus `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsagePlansStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsagePlansStatusResponseParams `json:"Response"`
}

func (r *DescribeUsagePlansStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlansStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachPluginRequestParams struct {
	// ID of the plugin to be unbound
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API environment
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// ID of the API to unbind from the plugin
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type DetachPluginRequest struct {
	*tchttp.BaseRequest
	
	// ID of the plugin to be unbound
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API environment
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// ID of the API to unbind from the plugin
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DetachPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachPluginResponseParams struct {
	// Whether unbinding succeeded.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachPluginResponse struct {
	*tchttp.BaseResponse
	Response *DetachPluginResponseParams `json:"Response"`
}

func (r *DetachPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApiKeyRequestParams struct {
	// ID of the key to be disabled.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

type DisableApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the key to be disabled.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DisableApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApiKeyResponseParams struct {
	// Whether the key is successfully disabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DisableApiKeyResponseParams `json:"Response"`
}

func (r *DisableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DocumentSDK struct {
	// Download link of generated file. Generated documents will be stored in COS.
	DocumentURL *string `json:"DocumentURL,omitempty" name:"DocumentURL"`

	// Download link of generated SDK file. Generated SDK files will be stored in COS.
	SdkURL *string `json:"SdkURL,omitempty" name:"SdkURL"`
}

type DomainSetList struct {
	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Domain name resolution status. `1`: normal, `0`: failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Certificate ID.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Whether the default path mapping is used.
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// Custom domain name protocol type.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Network type. Valid values: INNER, OUTER.
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Whether to force HTTP requests to redirect to HTTPS. Default value: `false`. When this parameter is `true`, API Gateway will redirect all requests using the custom domain name over the HTTP protocol to the HTTPS protocol for forwarding.
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`

	// ICP filing status
	RegistrationStatus *bool `json:"RegistrationStatus,omitempty" name:"RegistrationStatus"`
}

type DomainSets struct {
	// Number of custom domain names under service
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Custom service domain name list.
	DomainSet []*DomainSetList `json:"DomainSet,omitempty" name:"DomainSet"`
}

// Predefined struct for user
type EnableApiKeyRequestParams struct {
	// ID of the key to be enabled.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

type EnableApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the key to be enabled.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *EnableApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableApiKeyResponseParams struct {
	// Whether the key is successfully enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *EnableApiKeyResponseParams `json:"Response"`
}

func (r *EnableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Environment struct {
	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Access path.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Release status. 1: published. 0: not published.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Running version.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

type EnvironmentStrategy struct {
	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Throttling value
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`

	// Maximum quota value
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxQuota *int64 `json:"MaxQuota,omitempty" name:"MaxQuota"`
}

type ErrorCodes struct {
	// Custom response configuration error code.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Custom response configuration error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Custom response configuration error code remarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Custom error code conversion.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`

	// Whether to enable error code conversion.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type Filter struct {
	// Field to be filtered.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter value of field.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type GenerateApiDocumentRequestParams struct {
	// Unique service ID of the document to be created.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment of the service for which to create an SDK.
	GenEnvironment *string `json:"GenEnvironment,omitempty" name:"GenEnvironment"`

	// Programming language of the SDK to be created. Currently, only Python and JavaScript are supported.
	GenLanguage *string `json:"GenLanguage,omitempty" name:"GenLanguage"`
}

type GenerateApiDocumentRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of the document to be created.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment of the service for which to create an SDK.
	GenEnvironment *string `json:"GenEnvironment,omitempty" name:"GenEnvironment"`

	// Programming language of the SDK to be created. Currently, only Python and JavaScript are supported.
	GenLanguage *string `json:"GenLanguage,omitempty" name:"GenLanguage"`
}

func (r *GenerateApiDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateApiDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "GenEnvironment")
	delete(f, "GenLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateApiDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateApiDocumentResponseParams struct {
	// API document and SDK link.
	Result *DocumentSDK `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateApiDocumentResponse struct {
	*tchttp.BaseResponse
	Response *GenerateApiDocumentResponseParams `json:"Response"`
}

func (r *GenerateApiDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateApiDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheckConf struct {
	// Whether health check is enabled.
	IsHealthCheck *bool `json:"IsHealthCheck,omitempty" name:"IsHealthCheck"`

	// Health check threshold.
	RequestVolumeThreshold *int64 `json:"RequestVolumeThreshold,omitempty" name:"RequestVolumeThreshold"`

	// Window size.
	SleepWindowInMilliseconds *int64 `json:"SleepWindowInMilliseconds,omitempty" name:"SleepWindowInMilliseconds"`

	// Threshold percentage.
	ErrorThresholdPercentage *int64 `json:"ErrorThresholdPercentage,omitempty" name:"ErrorThresholdPercentage"`
}

type IPStrategy struct {
	// Unique policy ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Custom policy name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// Policy type. Valid values: WHITE (allowlist), BLACK (blocklist).
	// Note: this field may return null, indicating that no valid values can be obtained.
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// IP list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Modification time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of APIs bound to policy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindApiTotalCount *int64 `json:"BindApiTotalCount,omitempty" name:"BindApiTotalCount"`

	// Bound API details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindApis []*DesApisStatus `json:"BindApis,omitempty" name:"BindApis"`
}

type IPStrategyApi struct {
	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Custom API name.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API type. Valid values: NORMAL (general API), TSF (microservice API).
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API path, such as `/path`.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API request method, such as `GET`.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Unique ID of another policy bound to API.
	OtherIPStrategyId *string `json:"OtherIPStrategyId,omitempty" name:"OtherIPStrategyId"`

	// Environment bound to API.
	OtherEnvironmentName *string `json:"OtherEnvironmentName,omitempty" name:"OtherEnvironmentName"`
}

type IPStrategyApiStatus struct {
	// Number of APIs bound to environment.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Details of APIs bound to environment.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiIdStatusSet []*IPStrategyApi `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`
}

type IPStrategysStatus struct {
	// Number of policies.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Policy list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StrategySet []*IPStrategy `json:"StrategySet,omitempty" name:"StrategySet"`
}

type K8sLabel struct {
	// Key of the label
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value of the label
	Value *string `json:"Value,omitempty" name:"Value"`
}

type K8sService struct {
	// Weight
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// K8s cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace of the container
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Name of the service
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Service port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// The additional Label of the Pod
	ExtraLabels []*K8sLabel `json:"ExtraLabels,omitempty" name:"ExtraLabels"`

	// (Optional) Custom name of the service
	Name *string `json:"Name,omitempty" name:"Name"`
}

type LogQuery struct {
	// Search field
	Name *string `json:"Name,omitempty" name:"Name"`

	// Operator
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Search value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type MicroService struct {
	// Microservice cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Microservice namespace ID.
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Microservice name.
	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

type MicroServiceReq struct {
	// Microservice cluster.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Microservice namespace.
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Microservice name.
	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

// Predefined struct for user
type ModifyAPIDocRequestParams struct {
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`

	// API document name
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// Service name
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// List of APIs for which to generate documents
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type ModifyAPIDocRequest struct {
	*tchttp.BaseRequest
	
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`

	// API document name
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// Service name
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// List of APIs for which to generate documents
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *ModifyAPIDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAPIDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	delete(f, "ApiDocName")
	delete(f, "ServiceId")
	delete(f, "Environment")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAPIDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAPIDocResponseParams struct {
	// Basic information of API document
	Result *APIDoc `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAPIDocResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAPIDocResponseParams `json:"Response"`
}

func (r *ModifyAPIDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiAppRequestParams struct {
	// Unique application ID.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Modified application name
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// Modified application description
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
}

type ModifyApiAppRequest struct {
	*tchttp.BaseRequest
	
	// Unique application ID.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Modified application name
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// Modified application description
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
}

func (r *ModifyApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "ApiAppName")
	delete(f, "ApiAppDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiAppResponseParams struct {
	// Whether modification succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApiAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiAppResponseParams `json:"Response"`
}

func (r *ModifyApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiEnvironmentStrategyRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Throttling value.
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// API list.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type ModifyApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Throttling value.
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// API list.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *ModifyApiEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiEnvironmentStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Strategy")
	delete(f, "EnvironmentName")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiEnvironmentStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiEnvironmentStrategyResponseParams struct {
	// Whether modification succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiEnvironmentStrategyResponseParams `json:"Response"`
}

func (r *ModifyApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiIncrementRequestParams struct {
	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Authorization type of the API to be modified (you can select `OAUTH`, i.e., authorization API)
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// Public key value to be modified by OAuth API
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// OAuth API redirect address
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

type ModifyApiIncrementRequest struct {
	*tchttp.BaseRequest
	
	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Authorization type of the API to be modified (you can select `OAUTH`, i.e., authorization API)
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// Public key value to be modified by OAuth API
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// OAuth API redirect address
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

func (r *ModifyApiIncrementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiIncrementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiId")
	delete(f, "BusinessType")
	delete(f, "PublicKey")
	delete(f, "LoginRedirectUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiIncrementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiIncrementResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApiIncrementResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiIncrementResponseParams `json:"Response"`
}

func (r *ModifyApiIncrementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiIncrementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiRequestParams struct {
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API backend service type. Valid values: HTTP, MOCK, TSF, CLB, SCF, WEBSOCKET, TARGET (in beta test).
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Request frontend configuration.
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Custom API name.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Custom API description.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API type. Valid values: NORMAL, TSF. Default value: NORMAL.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API authentication type. Valid values: SECRET, NONE, OAUTH, APP. Default value: NONE.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// Whether signature authentication is required. True: yes; False: no. This parameter is to be disused.
	AuthRequired *bool `json:"AuthRequired,omitempty" name:"AuthRequired"`

	// API backend service timeout period in seconds.
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API frontend request type, such as HTTP, HTTPS, or HTTP and HTTPS.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Whether to enable CORS. True: yes; False: no.
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// Constant parameter.
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

	// Frontend request parameter.
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

	// This field is valid if `AuthType` is `OAUTH`. NORMAL: business API; OAUTH: authorization API.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Returned message of API backend Mock, which is required if `ServiceType` is `Mock`.
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// List of microservices bound to API.
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`

	// Load balancing configuration of microservice.
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// Health check configuration of microservice.
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// `target` type load balancing configuration (in beta test).
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// `target` health check configuration (in beta test).
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// SCF function name, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// SCF WebSocket registration function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// SCF WebSocket cleanup function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// SCF WebSocket transfer function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// SCF function namespace, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// SCF function version, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// SCF WebSocket registration function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// SCF WebSocket transfer function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// SCF WebSocket cleanup function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// SCF WebSocket cleanup function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// Whether to enable response integration, which takes effect if the backend type is `SCF`.
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// Billing after debugging starts (reserved field for marketplace).
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Tag.
	TagSpecifications *Tag `json:"TagSpecifications,omitempty" name:"TagSpecifications"`

	// Whether to delete the error codes for custom response configuration. If the value is left empty or `False`, the error codes will not be deleted. If the value is `True`, all custom response configuration error codes of the API will be deleted.
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// Return type.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Sample response for successful custom response configuration.
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// Sample response for failed custom response configuration.
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// API backend service configuration.
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// Unique ID of associated authorization API, which takes effect only if `AuthType` is `OAUTH` and `ApiBusinessType` is `NORMAL`. It is the unique ID of the OAuth 2.0 authorization API bound to the business API.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// API backend service parameter.
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// OAuth configuration, which takes effect if `AuthType` is `OAUTH`.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Custom error code configuration.
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// Whether to enable Base64 encoding. This parameter takes effect only when the backend is SCF.
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// Whether to trigger Base64 encoding by header. This parameter takes effect only when the backend is SCF.
	IsBase64Trigger *bool `json:"IsBase64Trigger,omitempty" name:"IsBase64Trigger"`

	// Header trigger rules. The number of rules cannot exceed 10.
	Base64EncodedTriggerRules []*Base64EncodedTriggerRule `json:"Base64EncodedTriggerRules,omitempty" name:"Base64EncodedTriggerRules"`

	// Event bus ID.
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// SCF function type, which takes effect when the backend type is `SCF`. Valid values: `EVENT` and `HTTP`.
	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`

	// EIAM application type.
	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`

	// EIAM application authentication type. Valid values: `AuthenticationOnly`, `Authentication`, `Authorization`.
	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`

	// Validity of the EIAM application token. Unit: second. Default value: `7200`.
	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`

	// EIAM application ID.
	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`
}

type ModifyApiRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API backend service type. Valid values: HTTP, MOCK, TSF, CLB, SCF, WEBSOCKET, TARGET (in beta test).
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Request frontend configuration.
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Custom API name.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Custom API description.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API type. Valid values: NORMAL, TSF. Default value: NORMAL.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API authentication type. Valid values: SECRET, NONE, OAUTH, APP. Default value: NONE.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// Whether signature authentication is required. True: yes; False: no. This parameter is to be disused.
	AuthRequired *bool `json:"AuthRequired,omitempty" name:"AuthRequired"`

	// API backend service timeout period in seconds.
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API frontend request type, such as HTTP, HTTPS, or HTTP and HTTPS.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Whether to enable CORS. True: yes; False: no.
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// Constant parameter.
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

	// Frontend request parameter.
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

	// This field is valid if `AuthType` is `OAUTH`. NORMAL: business API; OAUTH: authorization API.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Returned message of API backend Mock, which is required if `ServiceType` is `Mock`.
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// List of microservices bound to API.
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`

	// Load balancing configuration of microservice.
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// Health check configuration of microservice.
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// `target` type load balancing configuration (in beta test).
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// `target` health check configuration (in beta test).
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// SCF function name, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// SCF WebSocket registration function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// SCF WebSocket cleanup function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// SCF WebSocket transfer function, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// SCF function namespace, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// SCF function version, which takes effect if the backend type is `SCF`.
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// SCF WebSocket registration function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// SCF WebSocket transfer function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// SCF WebSocket transfer function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// SCF WebSocket cleanup function namespace, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// SCF WebSocket cleanup function version, which takes effect if the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// Whether to enable response integration, which takes effect if the backend type is `SCF`.
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// Billing after debugging starts (reserved field for marketplace).
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Tag.
	TagSpecifications *Tag `json:"TagSpecifications,omitempty" name:"TagSpecifications"`

	// Whether to delete the error codes for custom response configuration. If the value is left empty or `False`, the error codes will not be deleted. If the value is `True`, all custom response configuration error codes of the API will be deleted.
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// Return type.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Sample response for successful custom response configuration.
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// Sample response for failed custom response configuration.
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// API backend service configuration.
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// Unique ID of associated authorization API, which takes effect only if `AuthType` is `OAUTH` and `ApiBusinessType` is `NORMAL`. It is the unique ID of the OAuth 2.0 authorization API bound to the business API.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// API backend service parameter.
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// OAuth configuration, which takes effect if `AuthType` is `OAUTH`.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Custom error code configuration.
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// Whether to enable Base64 encoding. This parameter takes effect only when the backend is SCF.
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// Whether to trigger Base64 encoding by header. This parameter takes effect only when the backend is SCF.
	IsBase64Trigger *bool `json:"IsBase64Trigger,omitempty" name:"IsBase64Trigger"`

	// Header trigger rules. The number of rules cannot exceed 10.
	Base64EncodedTriggerRules []*Base64EncodedTriggerRule `json:"Base64EncodedTriggerRules,omitempty" name:"Base64EncodedTriggerRules"`

	// Event bus ID.
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// SCF function type, which takes effect when the backend type is `SCF`. Valid values: `EVENT` and `HTTP`.
	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`

	// EIAM application type.
	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`

	// EIAM application authentication type. Valid values: `AuthenticationOnly`, `Authentication`, `Authorization`.
	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`

	// Validity of the EIAM application token. Unit: second. Default value: `7200`.
	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`

	// EIAM application ID.
	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`
}

func (r *ModifyApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceType")
	delete(f, "RequestConfig")
	delete(f, "ApiId")
	delete(f, "ApiName")
	delete(f, "ApiDesc")
	delete(f, "ApiType")
	delete(f, "AuthType")
	delete(f, "AuthRequired")
	delete(f, "ServiceTimeout")
	delete(f, "Protocol")
	delete(f, "EnableCORS")
	delete(f, "ConstantParameters")
	delete(f, "RequestParameters")
	delete(f, "ApiBusinessType")
	delete(f, "ServiceMockReturnMessage")
	delete(f, "MicroServices")
	delete(f, "ServiceTsfLoadBalanceConf")
	delete(f, "ServiceTsfHealthCheckConf")
	delete(f, "TargetServicesLoadBalanceConf")
	delete(f, "TargetServicesHealthCheckConf")
	delete(f, "ServiceScfFunctionName")
	delete(f, "ServiceWebsocketRegisterFunctionName")
	delete(f, "ServiceWebsocketCleanupFunctionName")
	delete(f, "ServiceWebsocketTransportFunctionName")
	delete(f, "ServiceScfFunctionNamespace")
	delete(f, "ServiceScfFunctionQualifier")
	delete(f, "ServiceWebsocketRegisterFunctionNamespace")
	delete(f, "ServiceWebsocketRegisterFunctionQualifier")
	delete(f, "ServiceWebsocketTransportFunctionNamespace")
	delete(f, "ServiceWebsocketTransportFunctionQualifier")
	delete(f, "ServiceWebsocketCleanupFunctionNamespace")
	delete(f, "ServiceWebsocketCleanupFunctionQualifier")
	delete(f, "ServiceScfIsIntegratedResponse")
	delete(f, "IsDebugAfterCharge")
	delete(f, "TagSpecifications")
	delete(f, "IsDeleteResponseErrorCodes")
	delete(f, "ResponseType")
	delete(f, "ResponseSuccessExample")
	delete(f, "ResponseFailExample")
	delete(f, "ServiceConfig")
	delete(f, "AuthRelationApiId")
	delete(f, "ServiceParameters")
	delete(f, "OauthConfig")
	delete(f, "ResponseErrorCodes")
	delete(f, "IsBase64Encoded")
	delete(f, "IsBase64Trigger")
	delete(f, "Base64EncodedTriggerRules")
	delete(f, "EventBusId")
	delete(f, "ServiceScfFunctionType")
	delete(f, "EIAMAppType")
	delete(f, "EIAMAuthType")
	delete(f, "EIAMAppId")
	delete(f, "TokenTimeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApiResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiResponseParams `json:"Response"`
}

func (r *ModifyApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIPStrategyRequestParams struct {
	// Unique service ID of the policy to be modified.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the policy to be modified.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Details of the policy to be modified.
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

type ModifyIPStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID of the policy to be modified.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the policy to be modified.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Details of the policy to be modified.
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *ModifyIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "StrategyData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIPStrategyResponseParams struct {
	// Whether modification succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIPStrategyResponseParams `json:"Response"`
}

func (r *ModifyIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPluginRequestParams struct {
	// ID of the plugin to be modified
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Plugin name to be modified. A plugin name can contain up to 50 characters out of `a-z`, `A-Z`, `0-9`, and `_`, which must begin with a letter and end with a letter or a number.
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// Plugin description to be modified. A description is within 200 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Plugin definition statement to be modified. The json format is supported.
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`
}

type ModifyPluginRequest struct {
	*tchttp.BaseRequest
	
	// ID of the plugin to be modified
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Plugin name to be modified. A plugin name can contain up to 50 characters out of `a-z`, `A-Z`, `0-9`, and `_`, which must begin with a letter and end with a letter or a number.
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// Plugin description to be modified. A description is within 200 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Plugin definition statement to be modified. The json format is supported.
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`
}

func (r *ModifyPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "PluginName")
	delete(f, "Description")
	delete(f, "PluginData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPluginResponseParams struct {
	// Whether modification succeeded.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPluginResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPluginResponseParams `json:"Response"`
}

func (r *ModifyPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceEnvironmentStrategyRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Throttling value.
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Environment list.
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`
}

type ModifyServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Throttling value.
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Environment list.
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`
}

func (r *ModifyServiceEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Strategy")
	delete(f, "EnvironmentNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceEnvironmentStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceEnvironmentStrategyResponseParams struct {
	// Whether modification succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceEnvironmentStrategyResponseParams `json:"Response"`
}

func (r *ModifyServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceRequestParams struct {
	// Unique ID of the service to be modified.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Service name after modification.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Service description after modification.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Service frontend request type after modification, such as `http`, `https`, and `http&https`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Network type list, which is used to specify the supported network types. INNER: private network access; OUTER: public network access. Default value: OUTER.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
}

type ModifyServiceRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be modified.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Service name after modification.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Service description after modification.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Service frontend request type after modification, such as `http`, `https`, and `http&https`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Network type list, which is used to specify the supported network types. INNER: private network access; OUTER: public network access. Default value: OUTER.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
}

func (r *ModifyServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceName")
	delete(f, "ServiceDesc")
	delete(f, "Protocol")
	delete(f, "NetTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceResponseParams `json:"Response"`
}

func (r *ModifyServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubDomainRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name whose path mapping is to be modified.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Whether to change to the default path mapping. true: use the default path mapping; false: use the custom path mapping.
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// Certificate ID, which is required if the HTTPS protocol is included.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Custom domain name protocol type after modification. Valid values: http, https, http&https.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Path mapping list after modification.
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`

	// Network type. Valid values: INNER, OUTER.
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Whether to force HTTP requests to redirect to HTTPS. Default value: `false`. When this parameter is `true`, API Gateway will redirect all requests using the custom domain name over the HTTP protocol to the HTTPS protocol for forwarding.
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
}

type ModifySubDomainRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name whose path mapping is to be modified.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Whether to change to the default path mapping. true: use the default path mapping; false: use the custom path mapping.
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// Certificate ID, which is required if the HTTPS protocol is included.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Custom domain name protocol type after modification. Valid values: http, https, http&https.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Path mapping list after modification.
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`

	// Network type. Valid values: INNER, OUTER.
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Whether to force HTTP requests to redirect to HTTPS. Default value: `false`. When this parameter is `true`, API Gateway will redirect all requests using the custom domain name over the HTTP protocol to the HTTPS protocol for forwarding.
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
}

func (r *ModifySubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	delete(f, "IsDefaultMapping")
	delete(f, "CertificateId")
	delete(f, "Protocol")
	delete(f, "PathMappingSet")
	delete(f, "NetType")
	delete(f, "IsForcedHttps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubDomainResponseParams struct {
	// Whether the custom domain name is successfully modified.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubDomainResponseParams `json:"Response"`
}

func (r *ModifySubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUpstreamRequestParams struct {
	// Unique upstream ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// Upstream name
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// Upstream description
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// Backend protocol. Valid values: `HTTP`, `HTTPS`
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Upstream access type. Valid values: `IP_PORT`, `K8S`
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// Load balancing algorithm. Valid value: `ROUND_ROBIN`
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Unique VPC ID.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Retry attempts. It defaults to `3`.
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// Gateway forwarding to the upstream Host request header
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`

	// List of backend nodes
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// Health check configuration
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// Configuration of TKE service
	K8sService []*K8sService `json:"K8sService,omitempty" name:"K8sService"`
}

type ModifyUpstreamRequest struct {
	*tchttp.BaseRequest
	
	// Unique upstream ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// Upstream name
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// Upstream description
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// Backend protocol. Valid values: `HTTP`, `HTTPS`
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Upstream access type. Valid values: `IP_PORT`, `K8S`
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// Load balancing algorithm. Valid value: `ROUND_ROBIN`
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Unique VPC ID.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Retry attempts. It defaults to `3`.
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// Gateway forwarding to the upstream Host request header
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`

	// List of backend nodes
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// Health check configuration
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// Configuration of TKE service
	K8sService []*K8sService `json:"K8sService,omitempty" name:"K8sService"`
}

func (r *ModifyUpstreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUpstreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UpstreamId")
	delete(f, "UpstreamName")
	delete(f, "UpstreamDescription")
	delete(f, "Scheme")
	delete(f, "UpstreamType")
	delete(f, "Algorithm")
	delete(f, "UniqVpcId")
	delete(f, "Retries")
	delete(f, "UpstreamHost")
	delete(f, "Nodes")
	delete(f, "HealthChecker")
	delete(f, "K8sService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUpstreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUpstreamResponseParams struct {
	// Return modified upstream information
	// Note: This field may return `NULL`, indicating that no valid value was found.
	Result *UpstreamInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUpstreamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUpstreamResponseParams `json:"Response"`
}

func (r *ModifyUpstreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUsagePlanRequestParams struct {
	// Unique usage plan ID.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Custom usage plan name after modification.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Custom usage plan description after modification.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is `-1`, which indicates no limit.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is `-1`, which indicates no limit.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

type ModifyUsagePlanRequest struct {
	*tchttp.BaseRequest
	
	// Unique usage plan ID.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Custom usage plan name after modification.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Custom usage plan description after modification.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is `-1`, which indicates no limit.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is `-1`, which indicates no limit.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *ModifyUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "UsagePlanName")
	delete(f, "UsagePlanDesc")
	delete(f, "MaxRequestNum")
	delete(f, "MaxRequestNumPreSec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUsagePlanResponseParams struct {
	// Usage plan details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUsagePlanResponseParams `json:"Response"`
}

func (r *ModifyUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OauthConfig struct {
	// Public key for user token verification.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// Token delivery location.
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`

	// Redirect address, which is used to guide user logins.
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

type PathMapping struct {
	// Path.
	Path *string `json:"Path,omitempty" name:"Path"`

	// Release environment. Valid values: test, prepub, release.
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type Plugin struct {
	// Plugin ID
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Plugin name
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// Plugin type
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// Plugin definition statement
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`

	// Plugin description
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Plugin creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Plugin modification time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Total number of APIs bound with the plugin
	// Note: this field may return null, indicating that no valid values can be obtained.
	AttachedApiTotalCount *int64 `json:"AttachedApiTotalCount,omitempty" name:"AttachedApiTotalCount"`

	// Information of the API bound with the plugin
	// Note: this field may return null, indicating that no valid values can be obtained.
	AttachedApis []*AttachedApiInfo `json:"AttachedApis,omitempty" name:"AttachedApis"`
}

type ReleaseService struct {
	// Release remarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// Published version ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
}

// Predefined struct for user
type ReleaseServiceRequestParams struct {
	// Unique ID of the service to be published.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the environment to be published. Valid values: test (test environment), prepub (pre-release environment), release (release environment).
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Release description.
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// `apiId` list, which is reserved. Full API release is used by default.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type ReleaseServiceRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be published.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the environment to be published. Valid values: test (test environment), prepub (pre-release environment), release (release environment).
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Release description.
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// `apiId` list, which is reserved. Full API release is used by default.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *ReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "ReleaseDesc")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseServiceResponseParams struct {
	// Release information.
	Result *ReleaseService `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseServiceResponseParams `json:"Response"`
}

func (r *ReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReqParameter struct {
	// API frontend parameter name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Position of the API frontend parameter, such as the header. Supported values: `header`, `query`, and `path`.
	Position *string `json:"Position,omitempty" name:"Position"`

	// API frontend parameter type, such as `String` and `int`.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Default value of API frontend parameter.
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Whether the API frontend parameter is required. True: yes; False: no.
	Required *bool `json:"Required,omitempty" name:"Required"`

	// API frontend parameter remarks.
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type RequestConfig struct {
	// API path, such as `/path`.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API request method, such as `GET`.
	Method *string `json:"Method,omitempty" name:"Method"`
}

type RequestParameter struct {
	// Request parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Parameter position
	Position *string `json:"Position,omitempty" name:"Position"`

	// Parameter type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Default value
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Whether it is required
	Required *bool `json:"Required,omitempty" name:"Required"`
}

// Predefined struct for user
type ResetAPIDocPasswordRequestParams struct {
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

type ResetAPIDocPasswordRequest struct {
	*tchttp.BaseRequest
	
	// API document ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *ResetAPIDocPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAPIDocPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAPIDocPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAPIDocPasswordResponseParams struct {
	// Basic information of API document
	Result *APIDoc `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetAPIDocPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAPIDocPasswordResponseParams `json:"Response"`
}

func (r *ResetAPIDocPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAPIDocPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResponseErrorCodeReq struct {
	// Custom response configuration error code.
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Custom response configuration error message.
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Custom response configuration error code remarks.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Custom error code conversion.
	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`

	// Whether to enable error code conversion.
	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type Service struct {
	// Port for HTTPS access over private network.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`

	// Custom service description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Service frontend request type, such as `http`, `https`, and `http&https`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Network types supported by service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// Dedicated cluster name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// Unique service ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// IP version.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// List of published environments, such as test, prepub, and release.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments"`

	// Custom service name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Public domain name assigned by the system for this service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Port for HTTP access over private network.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InnerHttpPort *uint64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`

	// Private domain name automatically assigned by the system for this service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InnerSubDomain *string `json:"InnerSubDomain,omitempty" name:"InnerSubDomain"`

	// Billing status of service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TradeIsolateStatus *int64 `json:"TradeIsolateStatus,omitempty" name:"TradeIsolateStatus"`

	// Tags bound to a service.
	// Note: this field may return null, indicating that no valid values found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Dedicated instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Cluster type
	// Note: this field may return null, indicating that no valid values can be obtained.
	SetType *string `json:"SetType,omitempty" name:"SetType"`

	// Cluster type for service deployment
	// Note: this field may return null, indicating that no valid values found.
	DeploymentType *string `json:"DeploymentType,omitempty" name:"DeploymentType"`
}

type ServiceConfig struct {
	// The backend type. It’s available when `vpc` is enabled. Values: `clb`, `cvm` and `upstream`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Unique VPC ID.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API backend service URL, which is required if `ServiceType` is `HTTP`.
	Url *string `json:"Url,omitempty" name:"Url"`

	// API backend service path, such as `/path`, which is required if `ServiceType` is `HTTP`. The frontend and backend paths can be different.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API backend service request method, such as `GET`, which is required if `ServiceType` is `HTTP`. The frontend and backend methods can be different
	Method *string `json:"Method,omitempty" name:"Method"`

	// It’s required for `upstream`.
	// Note: This field may return `NULL`, indicating that no valid value was found.
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// API backend COS configuration. It’s required if the `ServiceType` is ·`COS`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CosConfig *CosConfig `json:"CosConfig,omitempty" name:"CosConfig"`
}

type ServiceEnvironmentSet struct {
	// Total number of environments bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of environments bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnvironmentList []*Environment `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
}

type ServiceEnvironmentStrategy struct {
	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Access service environment URL.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Release status.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Published version number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Throttling value.
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Maximum quota value
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxStrategy *int64 `json:"MaxStrategy,omitempty" name:"MaxStrategy"`
}

type ServiceEnvironmentStrategyStatus struct {
	// Number of throttling policies.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Throttling policy list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnvironmentList []*ServiceEnvironmentStrategy `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
}

type ServiceParameter struct {
	// API backend service parameter name, which is used only if `ServiceType` is `HTTP`. The frontend and backend parameter names can be different.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Position of API backend service parameter, such as `head`, which is used only if `ServiceType` is `HTTP`. The positions of frontend and backend parameters can be different.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Position *string `json:"Position,omitempty" name:"Position"`

	// Position of the API frontend parameter corresponding to the backend service parameter, such as `head`, which is used only if `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelevantRequestParameterPosition *string `json:"RelevantRequestParameterPosition,omitempty" name:"RelevantRequestParameterPosition"`

	// Name of the API frontend parameter corresponding to the backend service parameter, which is used only if `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelevantRequestParameterName *string `json:"RelevantRequestParameterName,omitempty" name:"RelevantRequestParameterName"`

	// Default value of API backend service parameter, which is used only if `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// API backend service parameter remarks, which is used only if `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelevantRequestParameterDesc *string `json:"RelevantRequestParameterDesc,omitempty" name:"RelevantRequestParameterDesc"`

	// API backend service parameter type, which is used only if `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelevantRequestParameterType *string `json:"RelevantRequestParameterType,omitempty" name:"RelevantRequestParameterType"`
}

type ServiceReleaseHistory struct {
	// Total number of published versions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Historical version list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList"`
}

type ServiceReleaseHistoryInfo struct {
	// Version ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Version description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// Version release time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
}

type ServiceReleaseVersion struct {
	// Total number of published versions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Release version list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList"`
}

type ServiceSubDomainMappings struct {
	// Whether the default path mapping is used. true: use the default path mapping; false: use the custom path mapping (`PathMappingSet` is required in this case).
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// Custom path mapping list.
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`
}

type ServiceUsagePlanSet struct {
	// Total number of usage plans bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of usage plans bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceUsagePlanList []*ApiUsagePlan `json:"ServiceUsagePlanList,omitempty" name:"ServiceUsagePlanList"`
}

type ServicesStatus struct {
	// Total number of services in list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Service list details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceSet []*Service `json:"ServiceSet,omitempty" name:"ServiceSet"`
}

type Tag struct {
	// Tag key.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TargetServicesReq struct {
	// VM IP
	VmIp *string `json:"VmIp,omitempty" name:"VmIp"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VM Port
	VmPort *int64 `json:"VmPort,omitempty" name:"VmPort"`

	// IP of the host where the CVM instance resides
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// Docker IP
	DockerIp *string `json:"DockerIp,omitempty" name:"DockerIp"`
}

type TsfLoadBalanceConfResp struct {
	// Whether load balancing is enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsLoadBalance *bool `json:"IsLoadBalance,omitempty" name:"IsLoadBalance"`

	// Load balancing method.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Whether session persistence is enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionStickRequired *bool `json:"SessionStickRequired,omitempty" name:"SessionStickRequired"`

	// Session persistence timeout period.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionStickTimeout *int64 `json:"SessionStickTimeout,omitempty" name:"SessionStickTimeout"`
}

// Predefined struct for user
type UnBindEnvironmentRequestParams struct {
	// Binding type. Valid values: API, SERVICE. Default value: SERVICE.
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// List of unique IDs of the usage plans to be bound.
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`

	// Service environment to be unbound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be unbound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID array, which is required if `BindType` is `API`.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type UnBindEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// Binding type. Valid values: API, SERVICE. Default value: SERVICE.
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// List of unique IDs of the usage plans to be bound.
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`

	// Service environment to be unbound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be unbound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID array, which is required if `BindType` is `API`.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *UnBindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BindType")
	delete(f, "UsagePlanIds")
	delete(f, "Environment")
	delete(f, "ServiceId")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindEnvironmentResponseParams struct {
	// Whether unbinding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *UnBindEnvironmentResponseParams `json:"Response"`
}

func (r *UnBindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindIPStrategyRequestParams struct {
	// Unique ID of the service to be unbound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the IP policy to be unbound.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Environment to be unbound.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs to be unbound.
	UnBindApiIds []*string `json:"UnBindApiIds,omitempty" name:"UnBindApiIds"`
}

type UnBindIPStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be unbound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the IP policy to be unbound.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Environment to be unbound.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs to be unbound.
	UnBindApiIds []*string `json:"UnBindApiIds,omitempty" name:"UnBindApiIds"`
}

func (r *UnBindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "EnvironmentName")
	delete(f, "UnBindApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindIPStrategyResponseParams struct {
	// Whether unbinding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *UnBindIPStrategyResponseParams `json:"Response"`
}

func (r *UnBindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindSecretIdsRequestParams struct {
	// Unique ID of the usage plan to be unbound.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Array of IDs of the keys to be unbound.
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

type UnBindSecretIdsRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the usage plan to be unbound.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Array of IDs of the keys to be unbound.
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

func (r *UnBindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindSecretIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "AccessKeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindSecretIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindSecretIdsResponseParams struct {
	// Whether unbinding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *UnBindSecretIdsResponseParams `json:"Response"`
}

func (r *UnBindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindSecretIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindSubDomainRequestParams struct {
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name to be unbound.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

type UnBindSubDomainRequest struct {
	*tchttp.BaseRequest
	
	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name to be unbound.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *UnBindSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindSubDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindSubDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindSubDomainResponseParams struct {
	// Whether the custom domain name is successfully unbound.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *UnBindSubDomainResponseParams `json:"Response"`
}

func (r *UnBindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindSubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnReleaseServiceRequestParams struct {
	// Unique ID of the service to be deactivated.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the environment to be deactivated. Valid values: test (test environment), prepub (pre-release environment), release (release environment).
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs to be deactivated, which is a reserved field.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type UnReleaseServiceRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be deactivated.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the environment to be deactivated. Valid values: test (test environment), prepub (pre-release environment), release (release environment).
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs to be deactivated, which is a reserved field.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *UnReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnReleaseServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnReleaseServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnReleaseServiceResponseParams struct {
	// Whether deactivation succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *UnReleaseServiceResponseParams `json:"Response"`
}

func (r *UnReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnReleaseServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindApiAppRequestParams struct {
	// Unique ID of the application to be bound.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Environment to be bound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the API to be bound.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type UnbindApiAppRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the application to be bound.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Environment to be bound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the API to be bound.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *UnbindApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "Environment")
	delete(f, "ServiceId")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindApiAppResponseParams struct {
	// Whether unbinding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindApiAppResponse struct {
	*tchttp.BaseResponse
	Response *UnbindApiAppResponseParams `json:"Response"`
}

func (r *UnbindApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiAppKeyRequestParams struct {
	// Unique application ID.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Application Key.
	ApiAppKey *string `json:"ApiAppKey,omitempty" name:"ApiAppKey"`

	// Application Secret.
	ApiAppSecret *string `json:"ApiAppSecret,omitempty" name:"ApiAppSecret"`
}

type UpdateApiAppKeyRequest struct {
	*tchttp.BaseRequest
	
	// Unique application ID.
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Application Key.
	ApiAppKey *string `json:"ApiAppKey,omitempty" name:"ApiAppKey"`

	// Application Secret.
	ApiAppSecret *string `json:"ApiAppSecret,omitempty" name:"ApiAppSecret"`
}

func (r *UpdateApiAppKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiAppKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "ApiAppKey")
	delete(f, "ApiAppSecret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApiAppKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiAppKeyResponseParams struct {
	// Whether update succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApiAppKeyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApiAppKeyResponseParams `json:"Response"`
}

func (r *UpdateApiAppKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiAppKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiKeyRequestParams struct {
	// ID of the key to be changed.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Key to be updated, which is required when a custom key is updated. It can contain 10–50 letters, digits, and underscores.
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

type UpdateApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the key to be changed.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Key to be updated, which is required when a custom key is updated. It can contain 10–50 letters, digits, and underscores.
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *UpdateApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	delete(f, "AccessKeySecret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiKeyResponseParams struct {
	// Key details after change.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *ApiKey `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApiKeyResponseParams `json:"Response"`
}

func (r *UpdateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServiceRequestParams struct {
	// Unique ID of the service to be switch.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the environment to be switched to. Valid values: test (test environment), prepub (pre-release environment), release (release environment).
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of the version to be switched to.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Switch description.
	UpdateDesc *string `json:"UpdateDesc,omitempty" name:"UpdateDesc"`
}

type UpdateServiceRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the service to be switch.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the environment to be switched to. Valid values: test (test environment), prepub (pre-release environment), release (release environment).
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of the version to be switched to.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Switch description.
	UpdateDesc *string `json:"UpdateDesc,omitempty" name:"UpdateDesc"`
}

func (r *UpdateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "VersionName")
	delete(f, "UpdateDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServiceResponseParams struct {
	// Whether the version is successfully switched.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateServiceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateServiceResponseParams `json:"Response"`
}

func (r *UpdateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpstreamHealthChecker struct {
	// Specifies whether to enable active health check
	EnableActiveCheck *bool `json:"EnableActiveCheck,omitempty" name:"EnableActiveCheck"`

	// Specifies whether the enable passive health check
	EnablePassiveCheck *bool `json:"EnablePassiveCheck,omitempty" name:"EnablePassiveCheck"`

	// The HTTP status code that indicates that the upstream is healthy
	HealthyHttpStatus *string `json:"HealthyHttpStatus,omitempty" name:"HealthyHttpStatus"`

	// The HTTP status code that indicates that the upstream is unhealthy
	UnhealthyHttpStatus *string `json:"UnhealthyHttpStatus,omitempty" name:"UnhealthyHttpStatus"`

	// The threshold on consecutive TCP errors. Range: [0, 254]. `0` indicates not to check TCP.
	TcpFailureThreshold *uint64 `json:"TcpFailureThreshold,omitempty" name:"TcpFailureThreshold"`

	// The threshold on consecutive timeouts. Range: [0, 254]. `0` indicates not to check TCP.
	TimeoutThreshold *uint64 `json:"TimeoutThreshold,omitempty" name:"TimeoutThreshold"`

	// The threshold on consecutive HTTP errors. Range: [0, 254]. `0` indicates not to check HTTP.
	HttpFailureThreshold *uint64 `json:"HttpFailureThreshold,omitempty" name:"HttpFailureThreshold"`

	// The path for active health check. It defaults to `/`.
	ActiveCheckHttpPath *string `json:"ActiveCheckHttpPath,omitempty" name:"ActiveCheckHttpPath"`

	// The timeout period for active health check in seconds. Default: `5`. 
	ActiveCheckTimeout *uint64 `json:"ActiveCheckTimeout,omitempty" name:"ActiveCheckTimeout"`

	// The interval for active health check in seconds. Default: `5`. 
	ActiveCheckInterval *uint64 `json:"ActiveCheckInterval,omitempty" name:"ActiveCheckInterval"`

	// Header of the active health check request
	ActiveRequestHeader []*UpstreamHealthCheckerReqHeaders `json:"ActiveRequestHeader,omitempty" name:"ActiveRequestHeader"`

	// The period for an abnormal to recover automatically in seconds. If only the passive health check is enabled, it must be greater than 0. Otherwise the abnormal nodes can not recovered automatically. The default value is 30 seconds. 
	UnhealthyTimeout *uint64 `json:"UnhealthyTimeout,omitempty" name:"UnhealthyTimeout"`
}

type UpstreamHealthCheckerReqHeaders struct {

}

type UpstreamInfo struct {
	// Unique upstream ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// Upstream name
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// Upstream description
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// Backend protocol. Valid values: `HTTP`, `HTTPS`
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Load balancing algorithm. Valid value: `ROUND_ROBIN`
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Unique VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Number of retry attempts
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// Backend nodes
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// Creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Label
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Health check configuration
	// Note: This field may return `null`, indicating that no valid value was found.
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// Upstream type. Valid values: `IP_PORT`, `K8S`
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// Configuration of TKE service
	// Note: This field may return `NULL`, indicating that no valid value was found.
	K8sServices []*K8sService `json:"K8sServices,omitempty" name:"K8sServices"`

	// The Host header that the gateway forwards to the upstream
	// Note: This field may return `NULL`, indicating that no valid value was found.
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`
}

type UpstreamNode struct {
	// IP or domain name
	Host *string `json:"Host,omitempty" name:"Host"`

	// The port number. Range: [0, 65535]
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Value range: [0, 100]. `0` refers to disable it.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// CVM Instance ID
	// Note: This field may return `NULL`, indicating that no valid value was found.
	VmInstanceId *string `json:"VmInstanceId,omitempty" name:"VmInstanceId"`

	// Tag
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// Health status of the node. Values: `OFF`, `HEALTHY`, `UNHEALTHY` and `NO_DATA`. It’s not required for creating and editing actions. It only supports VPC upstreams.
	// Note: This field may return `NULL`, indicating that no valid value was found.
	Healthy *string `json:"Healthy,omitempty" name:"Healthy"`

	// TKE container name
	// Note: This field may return `NULL`, indicating that no valid value was found.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// TKE namespace
	// Note: This field may return `NULL`, indicating that no valid value was found.
	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`

	// ID of the TKE cluster
	// Note: This field may return `null`, indicating that no valid value was found.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node source. Valid value: `K8S`
	// Note: This field may return `NULL`, indicating that no valid value was found.
	Source *string `json:"Source,omitempty" name:"Source"`

	// The unique service name in API Gateway
	// Note: This field may return `null`, indicating that no valid value was found.
	UniqueServiceName *string `json:"UniqueServiceName,omitempty" name:"UniqueServiceName"`
}

type UsagePlan struct {
	// Environment name.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique usage plan ID.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Usage plan name.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Usage plan description.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Usage plan QPS. `-1` indicates no limit.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// Usage plan time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Usage plan modification time.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type UsagePlanBindEnvironment struct {
	// Environment name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Unique service ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

type UsagePlanBindSecret struct {
	// Key ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Key name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Key status. 0: disabled. 1: enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type UsagePlanBindSecretStatus struct {
	// Number of keys bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of key details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccessKeyList []*UsagePlanBindSecret `json:"AccessKeyList,omitempty" name:"AccessKeyList"`
}

type UsagePlanEnvironment struct {
	// Unique ID of bound service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API method.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Name of bound environment.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Used quota.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

	// Maximum number of requests.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Maximum number of requests per second.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Service name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type UsagePlanEnvironmentStatus struct {
	// Number of environments of the service bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Environment status of services bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnvironmentList []*UsagePlanEnvironment `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
}

type UsagePlanInfo struct {
	// Unique usage plan ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Usage plan name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Usage plan description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Number of initialization calls.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InitQuota *int64 `json:"InitQuota,omitempty" name:"InitQuota"`

	// Limit of requests per second.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// Maximum number of calls.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Whether to hide.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsHide *int64 `json:"IsHide,omitempty" name:"IsHide"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Number of bound keys.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindSecretIdTotalCount *int64 `json:"BindSecretIdTotalCount,omitempty" name:"BindSecretIdTotalCount"`

	// Details of bound keys.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindSecretIds []*string `json:"BindSecretIds,omitempty" name:"BindSecretIds"`

	// Number of bound environments.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindEnvironmentTotalCount *int64 `json:"BindEnvironmentTotalCount,omitempty" name:"BindEnvironmentTotalCount"`

	// Details of bound environments.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindEnvironments []*UsagePlanBindEnvironment `json:"BindEnvironments,omitempty" name:"BindEnvironments"`
}

type UsagePlanStatusInfo struct {
	// Unique usage plan ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Custom usage plan name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Custom usage plan description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Maximum number of requests per second.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// Total number of requests allowed. `-1` indicates no limit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type UsagePlansStatus struct {
	// Number of eligible usage plans.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Usage plan list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanStatusSet []*UsagePlanStatusInfo `json:"UsagePlanStatusSet,omitempty" name:"UsagePlanStatusSet"`
}