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

package v20250106

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type APPOverview struct {
	// No data available
	// Note: This field may return null, indicating that no valid values can be obtained.
	BaseData *APPOverviewData `json:"BaseData,omitnil,omitempty" name:"BaseData"`

	// Superapp overview summary data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Overview *APPOverviewData `json:"Overview,omitnil,omitempty" name:"Overview"`

	// No data available
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageData *APPOverviewData `json:"PageData,omitnil,omitempty" name:"PageData"`

	// No data available
	// Note: This field may return null, indicating that no valid values can be obtained.
	Payment *APPOverviewData `json:"Payment,omitnil,omitempty" name:"Payment"`
}

type APPOverviewData struct {
	// Superapp active device count
	AllActiveDeviceNum *int64 `json:"AllActiveDeviceNum,omitnil,omitempty" name:"AllActiveDeviceNum"`

	// Superapp new device count
	AllNewDeviceNum *int64 `json:"AllNewDeviceNum,omitnil,omitempty" name:"AllNewDeviceNum"`

	// Mini program team count
	CorpNum *int64 `json:"CorpNum,omitnil,omitempty" name:"CorpNum"`

	// Mini game active device count
	GameActiveDeviceNum *int64 `json:"GameActiveDeviceNum,omitnil,omitempty" name:"GameActiveDeviceNum"`

	// Mini game new device count
	GameNewDeviceNum *int64 `json:"GameNewDeviceNum,omitnil,omitempty" name:"GameNewDeviceNum"`

	// Created mini program count
	MiniAppNum *int64 `json:"MiniAppNum,omitnil,omitempty" name:"MiniAppNum"`

	// Created mini game count
	MngNum *int64 `json:"MngNum,omitnil,omitempty" name:"MngNum"`

	// Mini program new device count
	NewDeviceNum *int64 `json:"NewDeviceNum,omitnil,omitempty" name:"NewDeviceNum"`

	// Released mini program count
	OnlineMiniAppNum *int64 `json:"OnlineMiniAppNum,omitnil,omitempty" name:"OnlineMiniAppNum"`

	// Released mini game count
	OnlineMngNum *int64 `json:"OnlineMngNum,omitnil,omitempty" name:"OnlineMngNum"`

	// Mini program active device count
	VisitNum *int64 `json:"VisitNum,omitnil,omitempty" name:"VisitNum"`

	// Data refresh timestamp
	FlushTime *string `json:"FlushTime,omitnil,omitempty" name:"FlushTime"`
}

type AccessAnalysisDetail struct {
	// Number of active devices
	ActiveCount *int64 `json:"ActiveCount,omitnil,omitempty" name:"ActiveCount"`

	// Average visit duration per user
	AvgDevice *string `json:"AvgDevice,omitnil,omitempty" name:"AvgDevice"`

	// Average visit duration per session
	AvgOnce *string `json:"AvgOnce,omitnil,omitempty" name:"AvgOnce"`

	// Average opens per user
	AvgOpenCount *string `json:"AvgOpenCount,omitnil,omitempty" name:"AvgOpenCount"`

	// Date
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// Data time
	FlushTime *string `json:"FlushTime,omitnil,omitempty" name:"FlushTime"`

	// Number of new devices
	NewCount *int64 `json:"NewCount,omitnil,omitempty" name:"NewCount"`

	// Number of opens
	OpenCount *int64 `json:"OpenCount,omitnil,omitempty" name:"OpenCount"`

	// Number of shares
	TotalShareNum *int64 `json:"TotalShareNum,omitnil,omitempty" name:"TotalShareNum"`

	// Cumulative users
	TotalUserNum *int64 `json:"TotalUserNum,omitnil,omitempty" name:"TotalUserNum"`
}

type AccessAnalysisOverview struct {
	// Overview of visit analysis data
	// Note: This field may return null, indicating that no valid values can be obtained.
	BaseData *AccessAnalysisDetail `json:"BaseData,omitnil,omitempty" name:"BaseData"`

	// This API does not respond.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Overview *Overview `json:"Overview,omitnil,omitempty" name:"Overview"`

	// This API does not respond.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageData *AccessAnalysisDetail `json:"PageData,omitnil,omitempty" name:"PageData"`

	// This API does not respond.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Payment *AccessAnalysisDetail `json:"Payment,omitnil,omitempty" name:"Payment"`
}

type AdTrendChart struct {
	// Estimated revenue
	// Note: This field may return null, indicating that no valid values can be obtained.
	EstimatedEarnings []*AnalysisData `json:"EstimatedEarnings,omitnil,omitempty" name:"EstimatedEarnings"`

	// Requests
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestsNumber []*AnalysisData `json:"RequestsNumber,omitnil,omitempty" name:"RequestsNumber"`

	// Impressions
	// Note: This field may return null, indicating that no valid values can be obtained.
	Impressions []*AnalysisData `json:"Impressions,omitnil,omitempty" name:"Impressions"`

	// Effective Cost Per Mille
	// Note: This field may return null, indicating that no valid values can be obtained.
	ECPM []*AnalysisData `json:"ECPM,omitnil,omitempty" name:"ECPM"`

	// Taps
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClicksNumber []*AnalysisData `json:"ClicksNumber,omitnil,omitempty" name:"ClicksNumber"`
}

// Predefined struct for user
type AddTeamMemberRequestParams struct {
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Members to be added
	MemberList []*CreateTeamMemberInfoReq `json:"MemberList,omitnil,omitempty" name:"MemberList"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type AddTeamMemberRequest struct {
	*tchttp.BaseRequest
	
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Members to be added
	MemberList []*CreateTeamMemberInfoReq `json:"MemberList,omitnil,omitempty" name:"MemberList"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *AddTeamMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTeamMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "MemberList")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddTeamMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddTeamMemberResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddTeamMemberResponse struct {
	*tchttp.BaseResponse
	Response *AddTeamMemberResponseParams `json:"Response"`
}

func (r *AddTeamMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTeamMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvertDataOverview struct {
	// Estimated revenue
	// Note: This field may return null, indicating that no valid values can be obtained.
	EstimatedEarnings *string `json:"EstimatedEarnings,omitnil,omitempty" name:"EstimatedEarnings"`

	// Requests
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestsNumber *int64 `json:"RequestsNumber,omitnil,omitempty" name:"RequestsNumber"`

	// Impressions
	// Note: This field may return null, indicating that no valid values can be obtained.
	Impressions *int64 `json:"Impressions,omitnil,omitempty" name:"Impressions"`

	// Taps
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClicksNumber *int64 `json:"ClicksNumber,omitnil,omitempty" name:"ClicksNumber"`

	// Impressions
	// Note: This field may return null, indicating that no valid values can be obtained.
	ECPM *string `json:"ECPM,omitnil,omitempty" name:"ECPM"`
}

type AnalysisAdvertOverview struct {
	// Advertising overview
	// Note: This field may return null, indicating that no valid values can be obtained.
	OverviewData *AdvertDataOverview `json:"OverviewData,omitnil,omitempty" name:"OverviewData"`
}

type AnalysisData struct {
	// Time
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// Value
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`
}

type ApplicationConfigInfo struct {
	// Superapp configuration type: 1 Non-production, 2 Production
	ApplicationType *int64 `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// Superapp package name
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// Superapp URL
	AppURL *string `json:"AppURL,omitnil,omitempty" name:"AppURL"`

	// Superapp configuration ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type ApprovalItem struct {
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Approval result. 2: Rejected;
	// 3: Approved
	ApprovalResult *int64 `json:"ApprovalResult,omitnil,omitempty" name:"ApprovalResult"`

	// Approval notes. It’s required when the request is rejected.
	ApprovalNote *string `json:"ApprovalNote,omitnil,omitempty" name:"ApprovalNote"`
}

type BooleanInfo struct {
	// Bool type response object
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`
}

// Predefined struct for user
type ConfigureMNPPreviewRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// 1: Set; 2: Cancel
	ActionType *int64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Path to the preview page
	PreivewEntrancePath *string `json:"PreivewEntrancePath,omitnil,omitempty" name:"PreivewEntrancePath"`
}

type ConfigureMNPPreviewRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// 1: Set; 2: Cancel
	ActionType *int64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Path to the preview page
	PreivewEntrancePath *string `json:"PreivewEntrancePath,omitnil,omitempty" name:"PreivewEntrancePath"`
}

func (r *ConfigureMNPPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureMNPPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "ActionType")
	delete(f, "MNPVersionId")
	delete(f, "PlatformId")
	delete(f, "PreivewEntrancePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureMNPPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureMNPPreviewResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfigureMNPPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ConfigureMNPPreviewResponseParams `json:"Response"`
}

func (r *ConfigureMNPPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureMNPPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationConfigRequestParams struct {
	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp type. 1: Test; 2: Formal
	ApplicationType *int64 `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// Superapp operating system. 2 Android 3 iOS
	ApplicationPlatformType *int64 `json:"ApplicationPlatformType,omitnil,omitempty" name:"ApplicationPlatformType"`

	// Package name: corresponds to packageName on Android and bundleId on iOS
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// Superapp URL
	AppURL *string `json:"AppURL,omitnil,omitempty" name:"AppURL"`
}

type CreateApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp type. 1: Test; 2: Formal
	ApplicationType *int64 `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// Superapp operating system. 2 Android 3 iOS
	ApplicationPlatformType *int64 `json:"ApplicationPlatformType,omitnil,omitempty" name:"ApplicationPlatformType"`

	// Package name: corresponds to packageName on Android and bundleId on iOS
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// Superapp URL
	AppURL *string `json:"AppURL,omitnil,omitempty" name:"AppURL"`
}

func (r *CreateApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PlatformId")
	delete(f, "ApplicationType")
	delete(f, "ApplicationPlatformType")
	delete(f, "AppKey")
	delete(f, "AppURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationConfigResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationConfigResponseParams `json:"Response"`
}

func (r *CreateApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationRequestParams struct {
	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Logo address
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Introduction
	Intro *string `json:"Intro,omitnil,omitempty" name:"Intro"`

	// Application type. 1: Test; 2: Formal
	//
	// Deprecated: ApplicationType is deprecated.
	ApplicationType *int64 `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// Android app package name
	//
	// Deprecated: AndroidAppKey is deprecated.
	AndroidAppKey *string `json:"AndroidAppKey,omitnil,omitempty" name:"AndroidAppKey"`

	// iOS App bundleId
	//
	// Deprecated: IosAppKey is deprecated.
	IosAppKey *string `json:"IosAppKey,omitnil,omitempty" name:"IosAppKey"`

	// Remarks
	//
	// Deprecated: Remark is deprecated.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Scheme
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Logo address
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Introduction
	Intro *string `json:"Intro,omitnil,omitempty" name:"Intro"`

	// Application type. 1: Test; 2: Formal
	ApplicationType *int64 `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// Android app package name
	AndroidAppKey *string `json:"AndroidAppKey,omitnil,omitempty" name:"AndroidAppKey"`

	// iOS App bundleId
	IosAppKey *string `json:"IosAppKey,omitnil,omitempty" name:"IosAppKey"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Scheme
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`
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
	delete(f, "Logo")
	delete(f, "PlatformId")
	delete(f, "TeamId")
	delete(f, "Intro")
	delete(f, "ApplicationType")
	delete(f, "AndroidAppKey")
	delete(f, "IosAppKey")
	delete(f, "Remark")
	delete(f, "Scheme")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationResponseParams struct {
	// Response data
	Data *ResourceIdStringInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CreateApplicationSensitiveAPIReq struct {
	// API name
	APIName *string `json:"APIName,omitnil,omitempty" name:"APIName"`

	// API description
	APIDesc *string `json:"APIDesc,omitnil,omitempty" name:"APIDesc"`

	// API type. 1: system; 2: custom.
	APIType *int64 `json:"APIType,omitnil,omitempty" name:"APIType"`
}

// Predefined struct for user
type CreateApplicationSensitiveAPIRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// List of newly-added APIs
	APIList []*CreateApplicationSensitiveAPIReq `json:"APIList,omitnil,omitempty" name:"APIList"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type CreateApplicationSensitiveAPIRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// List of newly-added APIs
	APIList []*CreateApplicationSensitiveAPIReq `json:"APIList,omitnil,omitempty" name:"APIList"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *CreateApplicationSensitiveAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationSensitiveAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "APIList")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationSensitiveAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationSensitiveAPIResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationSensitiveAPIResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationSensitiveAPIResponseParams `json:"Response"`
}

func (r *CreateApplicationSensitiveAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationSensitiveAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainParam struct {
	// Array of domain name URLs
	DomainUrlList []*string `json:"DomainUrlList,omitnil,omitempty" name:"DomainUrlList"`

	// Domain type. 1: requests domain; 2: WebView load domain, 3: sockets domain; 4: File upload; 5: File download
	DomainType *int64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`
}

// Predefined struct for user
type CreateGlobalDomainACLRequestParams struct {
	// Domain name list
	DomainUrlList []*string `json:"DomainUrlList,omitnil,omitempty" name:"DomainUrlList"`

	// Domain type. 1: Allowed; 2: Blocked
	DomainType *int64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type CreateGlobalDomainACLRequest struct {
	*tchttp.BaseRequest
	
	// Domain name list
	DomainUrlList []*string `json:"DomainUrlList,omitnil,omitempty" name:"DomainUrlList"`

	// Domain type. 1: Allowed; 2: Blocked
	DomainType *int64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *CreateGlobalDomainACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalDomainACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainUrlList")
	delete(f, "DomainType")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlobalDomainACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalDomainACLResponseParams struct {
	// Response data
	Data *CreateGlobalDomainResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlobalDomainACLResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlobalDomainACLResponseParams `json:"Response"`
}

func (r *CreateGlobalDomainACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalDomainACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGlobalDomainResp struct {
	// Result
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// Indicates the duplicate domain name.
	RepeatUrls []*string `json:"RepeatUrls,omitnil,omitempty" name:"RepeatUrls"`

	// Lists allowed domain names.
	ExistsWhiteUrls []*string `json:"ExistsWhiteUrls,omitnil,omitempty" name:"ExistsWhiteUrls"`

	// Indicates the domain name already exists in the blocklist.
	ExistsBlackUrls []*string `json:"ExistsBlackUrls,omitnil,omitempty" name:"ExistsBlackUrls"`
}

// Predefined struct for user
type CreateMNPApprovalRequestParams struct {
	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// submit: Submit an approval request; cancel: Cancel the approval request
	ApplyAction *string `json:"ApplyAction,omitnil,omitempty" name:"ApplyAction"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type CreateMNPApprovalRequest struct {
	*tchttp.BaseRequest
	
	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// submit: Submit an approval request; cancel: Cancel the approval request
	ApplyAction *string `json:"ApplyAction,omitnil,omitempty" name:"ApplyAction"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *CreateMNPApprovalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPApprovalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPVersionId")
	delete(f, "ApplyAction")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMNPApprovalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMNPApprovalResp struct {
	// Bool type response object
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// Approval No.
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`
}

// Predefined struct for user
type CreateMNPApprovalResponseParams struct {
	// Response data
	Data *CreateMNPApprovalResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMNPApprovalResponse struct {
	*tchttp.BaseResponse
	Response *CreateMNPApprovalResponseParams `json:"Response"`
}

func (r *CreateMNPApprovalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMNPDomainACLRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Domain name list
	Domain []*CreateDomainParam `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type CreateMNPDomainACLRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Domain name list
	Domain []*CreateDomainParam `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *CreateMNPDomainACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPDomainACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "Domain")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMNPDomainACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMNPDomainACLResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMNPDomainACLResponse struct {
	*tchttp.BaseResponse
	Response *CreateMNPDomainACLResponseParams `json:"Response"`
}

func (r *CreateMNPDomainACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPDomainACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMNPRequestParams struct {
	// Mini program type
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Mini app icon
	MNPIcon *string `json:"MNPIcon,omitnil,omitempty" name:"MNPIcon"`

	// Mini program introduction
	MNPIntro *string `json:"MNPIntro,omitnil,omitempty" name:"MNPIntro"`

	// Mini program description
	MNPDesc *string `json:"MNPDesc,omitnil,omitempty" name:"MNPDesc"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

type CreateMNPRequest struct {
	*tchttp.BaseRequest
	
	// Mini program type
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Mini app icon
	MNPIcon *string `json:"MNPIcon,omitnil,omitempty" name:"MNPIcon"`

	// Mini program introduction
	MNPIntro *string `json:"MNPIntro,omitnil,omitempty" name:"MNPIntro"`

	// Mini program description
	MNPDesc *string `json:"MNPDesc,omitnil,omitempty" name:"MNPDesc"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

func (r *CreateMNPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPType")
	delete(f, "MNPName")
	delete(f, "MNPIcon")
	delete(f, "MNPIntro")
	delete(f, "MNPDesc")
	delete(f, "PlatformId")
	delete(f, "TeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMNPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMNPResponseParams struct {
	// Response mini program ID
	Data *ResourceIdStringInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMNPResponse struct {
	*tchttp.BaseResponse
	Response *CreateMNPResponseParams `json:"Response"`
}

func (r *CreateMNPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMNPSensitiveAPIPermissionApprovalRequestParams struct {
	// API Id
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// Reason for application
	ApplyReason *string `json:"ApplyReason,omitnil,omitempty" name:"ApplyReason"`

	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type CreateMNPSensitiveAPIPermissionApprovalRequest struct {
	*tchttp.BaseRequest
	
	// API Id
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// Reason for application
	ApplyReason *string `json:"ApplyReason,omitnil,omitempty" name:"ApplyReason"`

	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *CreateMNPSensitiveAPIPermissionApprovalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPSensitiveAPIPermissionApprovalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "APIId")
	delete(f, "ApplyReason")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMNPSensitiveAPIPermissionApprovalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMNPSensitiveAPIPermissionApprovalResponseParams struct {
	// Response data
	Data *ResourceIdStringInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMNPSensitiveAPIPermissionApprovalResponse struct {
	*tchttp.BaseResponse
	Response *CreateMNPSensitiveAPIPermissionApprovalResponseParams `json:"Response"`
}

func (r *CreateMNPSensitiveAPIPermissionApprovalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPSensitiveAPIPermissionApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMNPVersionRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Version number
	MNPVersion *string `json:"MNPVersion,omitnil,omitempty" name:"MNPVersion"`

	// Address of the mini program package. You can export the package from IDE and upload it to a file server.
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Version introduction
	MNPVersionIntro *string `json:"MNPVersionIntro,omitnil,omitempty" name:"MNPVersionIntro"`
}

type CreateMNPVersionRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Version number
	MNPVersion *string `json:"MNPVersion,omitnil,omitempty" name:"MNPVersion"`

	// Address of the mini program package. You can export the package from IDE and upload it to a file server.
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Version introduction
	MNPVersionIntro *string `json:"MNPVersionIntro,omitnil,omitempty" name:"MNPVersionIntro"`
}

func (r *CreateMNPVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "MNPVersion")
	delete(f, "FileUrl")
	delete(f, "PlatformId")
	delete(f, "MNPVersionIntro")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMNPVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMNPVersionResp struct {
	// Specifies the ID of the task to create a mini program version.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CreateMNPVersionResponseParams struct {
	// Response data
	Data *CreateMNPVersionResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMNPVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateMNPVersionResponseParams `json:"Response"`
}

func (r *CreateMNPVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMNPVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresetKeyRequestParams struct {

}

type CreatePresetKeyRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreatePresetKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresetKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePresetKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresetKeyResponseParams struct {
	// Response data
	Data *PresetResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePresetKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreatePresetKeyResponseParams `json:"Response"`
}

func (r *CreatePresetKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresetKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTeamMemberInfoReq struct {
	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// User role ID
	UserRoleId *int64 `json:"UserRoleId,omitnil,omitempty" name:"UserRoleId"`
}

// Predefined struct for user
type CreateTeamRequestParams struct {
	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Admin name
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// Permission assigned to the team. 1: Mini program; 2: Application (only one of these types is supported)
	TeamRoleTypeList []*int64 `json:"TeamRoleTypeList,omitnil,omitempty" name:"TeamRoleTypeList"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Platform ID, required for API call
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Associated team ID
	RelatedTeamId *string `json:"RelatedTeamId,omitnil,omitempty" name:"RelatedTeamId"`
}

type CreateTeamRequest struct {
	*tchttp.BaseRequest
	
	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Admin name
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// Permission assigned to the team. 1: Mini program; 2: Application (only one of these types is supported)
	TeamRoleTypeList []*int64 `json:"TeamRoleTypeList,omitnil,omitempty" name:"TeamRoleTypeList"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Platform ID, required for API call
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Associated team ID
	RelatedTeamId *string `json:"RelatedTeamId,omitnil,omitempty" name:"RelatedTeamId"`
}

func (r *CreateTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTeamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamName")
	delete(f, "AdminUserId")
	delete(f, "TeamRoleTypeList")
	delete(f, "Remark")
	delete(f, "PlatformId")
	delete(f, "RelatedTeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTeamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTeamResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTeamResponse struct {
	*tchttp.BaseResponse
	Response *CreateTeamResponseParams `json:"Response"`
}

func (r *CreateTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTeamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// User account
	UserAccount *string `json:"UserAccount,omitnil,omitempty" name:"UserAccount"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// User account type. 2: Platform admin; 3: Member.
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Account password. Use CreatePresetKey to get the public key to encrypt the password.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Call CreatePresetKey to get the keyID from RequestId
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// User account
	UserAccount *string `json:"UserAccount,omitnil,omitempty" name:"UserAccount"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// User account type. 2: Platform admin; 3: Member.
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Account password. Use CreatePresetKey to get the public key to encrypt the password.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Call CreatePresetKey to get the keyID from RequestId
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserAccount")
	delete(f, "UserName")
	delete(f, "AccountType")
	delete(f, "Password")
	delete(f, "KeyId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// Response data, user ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *ResourceIdStringInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DeleteApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
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
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteApplicationSensitiveAPIRequestParams struct {
	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DeleteApplicationSensitiveAPIRequest struct {
	*tchttp.BaseRequest
	
	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DeleteApplicationSensitiveAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationSensitiveAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "APIId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationSensitiveAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationSensitiveAPIResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteApplicationSensitiveAPIResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationSensitiveAPIResponseParams `json:"Response"`
}

func (r *DeleteApplicationSensitiveAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationSensitiveAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalDomainRequestParams struct {
	// Domain ID
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DeleteGlobalDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain ID
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DeleteGlobalDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGlobalDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalDomainResponseParams struct {
	// Response data
	Data *GlobalDomainDeleteResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGlobalDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGlobalDomainResponseParams `json:"Response"`
}

func (r *DeleteGlobalDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMNPRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DeleteMNPRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DeleteMNPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMNPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMNPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMNPResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMNPResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMNPResponseParams `json:"Response"`
}

func (r *DeleteMNPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMNPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTeamMemberRequestParams struct {
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DeleteTeamMemberRequest struct {
	*tchttp.BaseRequest
	
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DeleteTeamMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTeamMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "UserId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTeamMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTeamMemberResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTeamMemberResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTeamMemberResponseParams `json:"Response"`
}

func (r *DeleteTeamMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTeamMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTeamRequestParams struct {
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DeleteTeamRequest struct {
	*tchttp.BaseRequest
	
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DeleteTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTeamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTeamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTeamResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTeamResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTeamResponseParams `json:"Response"`
}

func (r *DeleteTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTeamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPPDataDetailLineChartRequestParams struct {
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId (optional):
	// app_minigame_num: Number of created mini games
	// app_online_miniapp_num: Number of available mini programs
	// app_miniapp_num: Number of created mini programs
	// app_related_corp_num: Mini program team data
	// app_online_minigame_num: Number of available mini games
	// app_active_device_num: Number of active devices
	// app_new_device_num: Number of new devices
	IndexIds []*string `json:"IndexIds,omitnil,omitempty" name:"IndexIds"`

	// Input parameter base64 string: {"BeginDate":"20251122","EndDate":"20251128"}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`

	// Superapp ID
	ApplicationIds []*string `json:"ApplicationIds,omitnil,omitempty" name:"ApplicationIds"`
}

type DescribeAPPDataDetailLineChartRequest struct {
	*tchttp.BaseRequest
	
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId (optional):
	// app_minigame_num: Number of created mini games
	// app_online_miniapp_num: Number of available mini programs
	// app_miniapp_num: Number of created mini programs
	// app_related_corp_num: Mini program team data
	// app_online_minigame_num: Number of available mini games
	// app_active_device_num: Number of active devices
	// app_new_device_num: Number of new devices
	IndexIds []*string `json:"IndexIds,omitnil,omitempty" name:"IndexIds"`

	// Input parameter base64 string: {"BeginDate":"20251122","EndDate":"20251128"}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`

	// Superapp ID
	ApplicationIds []*string `json:"ApplicationIds,omitnil,omitempty" name:"ApplicationIds"`
}

func (r *DescribeAPPDataDetailLineChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPPDataDetailLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformId")
	delete(f, "ReportId")
	delete(f, "IndexIds")
	delete(f, "QueryData")
	delete(f, "ApplicationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAPPDataDetailLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPPDataDetailLineChartResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*ReportDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAPPDataDetailLineChartResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAPPDataDetailLineChartResponseParams `json:"Response"`
}

func (r *DescribeAPPDataDetailLineChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPPDataDetailLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPPDataOverviewRequestParams struct {
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Start time in YYYYMMDD format
	DataTime *int64 `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// Superapp ID
	ApplicationIds []*string `json:"ApplicationIds,omitnil,omitempty" name:"ApplicationIds"`
}

type DescribeAPPDataOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Start time in YYYYMMDD format
	DataTime *int64 `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// Superapp ID
	ApplicationIds []*string `json:"ApplicationIds,omitnil,omitempty" name:"ApplicationIds"`
}

func (r *DescribeAPPDataOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPPDataOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformId")
	delete(f, "DataTime")
	delete(f, "ApplicationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAPPDataOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPPDataOverviewResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *APPOverview `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAPPDataOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAPPDataOverviewResponseParams `json:"Response"`
}

func (r *DescribeAPPDataOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPPDataOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdvertisingLineChartRequestParams struct {
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// // 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	//   1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`
}

type DescribeAdvertisingLineChartRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// // 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	//   1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`
}

func (r *DescribeAdvertisingLineChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdvertisingLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	delete(f, "AdUnitType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdvertisingLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdvertisingLineChartResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *AdTrendChart `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAdvertisingLineChartResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAdvertisingLineChartResponseParams `json:"Response"`
}

func (r *DescribeAdvertisingLineChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdvertisingLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdvertisingOverviewRequestParams struct {
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// // 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`
}

type DescribeAdvertisingOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// // 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`
}

func (r *DescribeAdvertisingOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdvertisingOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	delete(f, "AdUnitType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdvertisingOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdvertisingOverviewResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *MNPAdvertisingOverview `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAdvertisingOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAdvertisingOverviewResponseParams `json:"Response"`
}

func (r *DescribeAdvertisingOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdvertisingOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationConfigFileRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Application platform. 2: Android; 3: iOS
	//
	// Deprecated: AppType is deprecated.
	AppType *int64 `json:"AppType,omitnil,omitempty" name:"AppType"`
}

type DescribeApplicationConfigFileRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Application platform. 2: Android; 3: iOS
	AppType *int64 `json:"AppType,omitnil,omitempty" name:"AppType"`
}

func (r *DescribeApplicationConfigFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationConfigFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PlatformId")
	delete(f, "AppType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationConfigFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationConfigFileResponseParams struct {
	// Response data
	Data *DownloadApplicationConfigResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationConfigFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationConfigFileResponseParams `json:"Response"`
}

func (r *DescribeApplicationConfigFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationConfigFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationConfigInfo struct {
	// Android configuration list
	AndroidConfig []*ApplicationConfigInfo `json:"AndroidConfig,omitnil,omitempty" name:"AndroidConfig"`

	// iOS configuration list
	IosConfig []*ApplicationConfigInfo `json:"IosConfig,omitnil,omitempty" name:"IosConfig"`
}

// Predefined struct for user
type DescribeApplicationConfigInfosRequestParams struct {
	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type DescribeApplicationConfigInfosRequest struct {
	*tchttp.BaseRequest
	
	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationConfigInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationConfigInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationConfigInfosResponseParams struct {
	// Response data
	Data *DescribeApplicationConfigInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationConfigInfosResponseParams `json:"Response"`
}

func (r *DescribeApplicationConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationListData struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// App Id.
	AppIdentityId *int64 `json:"AppIdentityId,omitnil,omitempty" name:"AppIdentityId"`

	// Name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Icon
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Android app package name
	AndroidAppKey *string `json:"AndroidAppKey,omitnil,omitempty" name:"AndroidAppKey"`

	// iOS App bundleId
	IosAppKey *string `json:"IosAppKey,omitnil,omitempty" name:"IosAppKey"`

	// Creator
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Specifies the updater.
	UpdateUser *string `json:"UpdateUser,omitnil,omitempty" name:"UpdateUser"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Introduction
	Intro *string `json:"Intro,omitnil,omitempty" name:"Intro"`

	// Team Id.
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Number of sensitive apis.
	SensitiveApiCount *int64 `json:"SensitiveApiCount,omitnil,omitempty" name:"SensitiveApiCount"`

	// Application type. 1: Test; 2: Formal
	ApplicationType *int64 `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`
}

// Predefined struct for user
type DescribeApplicationListRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Keywords for search (app name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

type DescribeApplicationListRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Keywords for search (app name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

func (r *DescribeApplicationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "Keyword")
	delete(f, "TeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationListResp struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List information
	DataList []*DescribeApplicationListData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeApplicationListResponseParams struct {
	// Response data
	Data *DescribeApplicationListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationListResponseParams `json:"Response"`
}

func (r *DescribeApplicationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationResp struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Product ID.
	AppIdentityId *int64 `json:"AppIdentityId,omitnil,omitempty" name:"AppIdentityId"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Specifies the application icon.
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Android package name
	AndroidAppKey *string `json:"AndroidAppKey,omitnil,omitempty" name:"AndroidAppKey"`

	// iOS bundleId
	IosAppKey *string `json:"IosAppKey,omitnil,omitempty" name:"IosAppKey"`

	// Creator
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Specifies the updater.
	UpdateUser *string `json:"UpdateUser,omitnil,omitempty" name:"UpdateUser"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Describes the application description.
	Intro *string `json:"Intro,omitnil,omitempty" name:"Intro"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Number of sensitive apis.
	SensitiveApiCount *int64 `json:"SensitiveApiCount,omitnil,omitempty" name:"SensitiveApiCount"`

	// Application type. 1: Test; 2: Formal
	ApplicationType *int64 `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// Specifies the application Scheme.
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`
}

// Predefined struct for user
type DescribeApplicationResponseParams struct {
	// Response data
	Data *DescribeApplicationResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationResponseParams `json:"Response"`
}

func (r *DescribeApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationSensitiveAPIListData struct {
	// APIID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// API name
	APIName *string `json:"APIName,omitnil,omitempty" name:"APIName"`

	// API request method
	APIMethod *string `json:"APIMethod,omitnil,omitempty" name:"APIMethod"`

	// API description
	APIDesc *string `json:"APIDesc,omitnil,omitempty" name:"APIDesc"`

	// Creator
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Updater
	UpdateUser *string `json:"UpdateUser,omitnil,omitempty" name:"UpdateUser"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Specifies the application icon.
	ApplicationLogo *string `json:"ApplicationLogo,omitnil,omitempty" name:"ApplicationLogo"`

	// API type. 1: system; 2: custom.
	APIType *int64 `json:"APIType,omitnil,omitempty" name:"APIType"`

	// API status. 0: public; 1: restricted.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeApplicationSensitiveAPIListRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Keywords for search (API name or method)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

type DescribeApplicationSensitiveAPIListRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Keywords for search (API name or method)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

func (r *DescribeApplicationSensitiveAPIListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationSensitiveAPIListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "ApplicationId")
	delete(f, "Keyword")
	delete(f, "TeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationSensitiveAPIListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationSensitiveAPIListResp struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List information
	DataList []*DescribeApplicationSensitiveAPIListData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeApplicationSensitiveAPIListResponseParams struct {
	// Response data
	Data *DescribeApplicationSensitiveAPIListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationSensitiveAPIListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationSensitiveAPIListResponseParams `json:"Response"`
}

func (r *DescribeApplicationSensitiveAPIListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationSensitiveAPIListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainInfoParam struct {
	// Multiple domain separators ';'.
	DomainUrl *string `json:"DomainUrl,omitnil,omitempty" name:"DomainUrl"`

	// Domain type 1-requests domain 2-business domain.
	DomainType *int64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`
}

// Predefined struct for user
type DescribeGlobalDomainACLRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Domain type. 1: Allowed; 2: Blocked
	DomainTypes []*int64 `json:"DomainTypes,omitnil,omitempty" name:"DomainTypes"`

	// Domain names to be queried
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeGlobalDomainACLRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Domain type. 1: Allowed; 2: Blocked
	DomainTypes []*int64 `json:"DomainTypes,omitnil,omitempty" name:"DomainTypes"`

	// Domain names to be queried
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeGlobalDomainACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalDomainACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "DomainTypes")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalDomainACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalDomainACLResponseParams struct {
	// Response data
	Data *DescribeGlobalDomainsResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalDomainACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalDomainACLResponseParams `json:"Response"`
}

func (r *DescribeGlobalDomainACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalDomainACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGlobalDomainsListData struct {
	// Domain ID
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// Domain name
	DomainUrl *string `json:"DomainUrl,omitnil,omitempty" name:"DomainUrl"`

	// Type. 1: allowlist; 2: blocklist.
	DomainType *int64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`

	// Creator
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Specifies the updater.
	UpdateUser *string `json:"UpdateUser,omitnil,omitempty" name:"UpdateUser"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DescribeGlobalDomainsResp struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Describes the domain information.
	DataList []*DescribeGlobalDomainsListData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeGlobalOverviewDataSummaryRequestParams struct {
	// string: Overview
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Date in YYYYMMDD format
	DataTime *int64 `json:"DataTime,omitnil,omitempty" name:"DataTime"`
}

type DescribeGlobalOverviewDataSummaryRequest struct {
	*tchttp.BaseRequest
	
	// string: Overview
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Date in YYYYMMDD format
	DataTime *int64 `json:"DataTime,omitnil,omitempty" name:"DataTime"`
}

func (r *DescribeGlobalOverviewDataSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalOverviewDataSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "DataTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalOverviewDataSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalOverviewDataSummaryResponseParams struct {
	// Data display at the top of the page
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *AccessAnalysisOverview `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalOverviewDataSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalOverviewDataSummaryResponseParams `json:"Response"`
}

func (r *DescribeGlobalOverviewDataSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalOverviewDataSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalOverviewReportDetailRequestParams struct {
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Parameter value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// Parameter value: app_num|corp_num|miniapp_num|miniapp_visit_pv|minigame_num|minigame_visit_pv
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Base64 string containing start and end time: {"BeginDate":20251111,"EndDate":20251125}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

type DescribeGlobalOverviewReportDetailRequest struct {
	*tchttp.BaseRequest
	
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Parameter value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// Parameter value: app_num|corp_num|miniapp_num|miniapp_visit_pv|minigame_num|minigame_visit_pv
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Base64 string containing start and end time: {"BeginDate":20251111,"EndDate":20251125}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

func (r *DescribeGlobalOverviewReportDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalOverviewReportDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformId")
	delete(f, "ReportId")
	delete(f, "IndexId")
	delete(f, "QueryData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalOverviewReportDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalOverviewReportDetailResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*ReportDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalOverviewReportDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalOverviewReportDetailResponseParams `json:"Response"`
}

func (r *DescribeGlobalOverviewReportDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalOverviewReportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAccessAnalysisDetailRequestParams struct {
	// Start time in YYYYMMDD format
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	// End time in YYYYMMDD format
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`
}

type DescribeMNGAccessAnalysisDetailRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	// End time in YYYYMMDD format
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`
}

func (r *DescribeMNGAccessAnalysisDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAccessAnalysisDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "Platform")
	delete(f, "TimeEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGAccessAnalysisDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAccessAnalysisDetailResponseParams struct {
	// Data display at the top of the page
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*AccessAnalysisDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGAccessAnalysisDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGAccessAnalysisDetailResponseParams `json:"Response"`
}

func (r *DescribeMNGAccessAnalysisDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAccessAnalysisDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAccessAnalysisLineChartRequestParams struct {
	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId (optional):
	// active_device_num: Number of active users
	// new_device_num: Number of new users
	// total_user_num: Total number of users
	// share_num: Number of shares
	// miniapp_open_num: Number of mini game opens
	// avg_device_open_num: Average opens per user
	// avg_device_open_duration: Average visit duration per user
	// avg_count_open_duration: Average visit duration per session
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"DataType":"1","Platform":0,"BeginDate":20251118,"EndDate":20251124}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

type DescribeMNGAccessAnalysisLineChartRequest struct {
	*tchttp.BaseRequest
	
	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId (optional):
	// active_device_num: Number of active users
	// new_device_num: Number of new users
	// total_user_num: Total number of users
	// share_num: Number of shares
	// miniapp_open_num: Number of mini game opens
	// avg_device_open_num: Average opens per user
	// avg_device_open_duration: Average visit duration per user
	// avg_count_open_duration: Average visit duration per session
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"DataType":"1","Platform":0,"BeginDate":20251118,"EndDate":20251124}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

func (r *DescribeMNGAccessAnalysisLineChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAccessAnalysisLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "ReportId")
	delete(f, "IndexId")
	delete(f, "QueryData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGAccessAnalysisLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAccessAnalysisLineChartResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*ReportDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGAccessAnalysisLineChartResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGAccessAnalysisLineChartResponseParams `json:"Response"`
}

func (r *DescribeMNGAccessAnalysisLineChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAccessAnalysisLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAccessAnalysisOverviewRequestParams struct {
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time
	TimeEnd *uint64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 1 Production data, 0 Non-production data
	ProdData *int64 `json:"ProdData,omitnil,omitempty" name:"ProdData"`

	// Operating system: 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNGAccessAnalysisOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time
	TimeEnd *uint64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 1 Production data, 0 Non-production data
	ProdData *int64 `json:"ProdData,omitnil,omitempty" name:"ProdData"`

	// Operating system: 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNGAccessAnalysisOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAccessAnalysisOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "TimeEnd")
	delete(f, "ProdData")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGAccessAnalysisOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAccessAnalysisOverviewResponseParams struct {
	// Data display at the top of the page
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *AccessAnalysisOverview `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGAccessAnalysisOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGAccessAnalysisOverviewResponseParams `json:"Response"`
}

func (r *DescribeMNGAccessAnalysisOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAccessAnalysisOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGActiveUserRealTimeStatisticsRequestParams struct {
	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"Platform":0,"DataType":"1","BeginDate":"20251125","EndDate":"20251125"}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

type DescribeMNGActiveUserRealTimeStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"Platform":0,"DataType":"1","BeginDate":"20251125","EndDate":"20251125"}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

func (r *DescribeMNGActiveUserRealTimeStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGActiveUserRealTimeStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "ReportId")
	delete(f, "IndexId")
	delete(f, "QueryData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGActiveUserRealTimeStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGActiveUserRealTimeStatisticsResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*ReportDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGActiveUserRealTimeStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGActiveUserRealTimeStatisticsResponseParams `json:"Response"`
}

func (r *DescribeMNGActiveUserRealTimeStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGActiveUserRealTimeStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAdvertisingDetailRequestParams struct {
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`

	// // 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNGAdvertisingDetailRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`

	// // 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNGAdvertisingDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAdvertisingDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "TimeEnd")
	delete(f, "AdUnitType")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGAdvertisingDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAdvertisingDetailResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*OverviewDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGAdvertisingDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGAdvertisingDetailResponseParams `json:"Response"`
}

func (r *DescribeMNGAdvertisingDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAdvertisingDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAdvertisingLineChartRequestParams struct {
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`

	// // 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNGAdvertisingLineChartRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`

	// // 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNGAdvertisingLineChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAdvertisingLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "TimeEnd")
	delete(f, "AdUnitType")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGAdvertisingLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAdvertisingLineChartResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *AdTrendChart `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGAdvertisingLineChartResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGAdvertisingLineChartResponseParams `json:"Response"`
}

func (r *DescribeMNGAdvertisingLineChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAdvertisingLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAdvertisingOverviewRequestParams struct {
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`

	// // 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNGAdvertisingOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`

	// // 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNGAdvertisingOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAdvertisingOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "TimeEnd")
	delete(f, "AdUnitType")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGAdvertisingOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGAdvertisingOverviewResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *AnalysisAdvertOverview `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGAdvertisingOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGAdvertisingOverviewResponseParams `json:"Response"`
}

func (r *DescribeMNGAdvertisingOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGAdvertisingOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGMAUDataDetailRequestParams struct {
	// Type
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program team ID
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

type DescribeMNGMAUDataDetailRequest struct {
	*tchttp.BaseRequest
	
	// Type
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program team ID
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

func (r *DescribeMNGMAUDataDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGMAUDataDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "ApplicationId")
	delete(f, "MNPId")
	delete(f, "MNPTeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGMAUDataDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGMAUDataDetailResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*MAUDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGMAUDataDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGMAUDataDetailResponseParams `json:"Response"`
}

func (r *DescribeMNGMAUDataDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGMAUDataDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGMAULineChartRequestParams struct {
	// Type: 0 Non-production data, 1 Production data 
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program team ID
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`
}

type DescribeMNGMAULineChartRequest struct {
	*tchttp.BaseRequest
	
	// Type: 0 Non-production data, 1 Production data 
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program team ID
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`
}

func (r *DescribeMNGMAULineChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGMAULineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "ApplicationId")
	delete(f, "MNPTeamId")
	delete(f, "MNPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGMAULineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGMAULineChartResponseParams struct {
	// Data display at the top of the page
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*MNGMAULineChartData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGMAULineChartResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGMAULineChartResponseParams `json:"Response"`
}

func (r *DescribeMNGMAULineChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGMAULineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGMAUMonthlyComparisonMetricCardRequestParams struct {
	// Start time in YYYYMMDD format
	SourceMonth *int64 `json:"SourceMonth,omitnil,omitempty" name:"SourceMonth"`

	// Type: 0 Non-production data, 1 Production data 
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TargetMonth *int64 `json:"TargetMonth,omitnil,omitempty" name:"TargetMonth"`

	// Superapp ID starting with App
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program team ID, required. When provided, the query is performed based on the mini program team.
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

type DescribeMNGMAUMonthlyComparisonMetricCardRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	SourceMonth *int64 `json:"SourceMonth,omitnil,omitempty" name:"SourceMonth"`

	// Type: 0 Non-production data, 1 Production data 
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TargetMonth *int64 `json:"TargetMonth,omitnil,omitempty" name:"TargetMonth"`

	// Superapp ID starting with App
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program team ID, required. When provided, the query is performed based on the mini program team.
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

func (r *DescribeMNGMAUMonthlyComparisonMetricCardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGMAUMonthlyComparisonMetricCardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceMonth")
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "TargetMonth")
	delete(f, "ApplicationId")
	delete(f, "MNPId")
	delete(f, "MNPTeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGMAUMonthlyComparisonMetricCardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGMAUMonthlyComparisonMetricCardResponseParams struct {
	// Data display at the top of the page
	Data *MAUIndicatorCard `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGMAUMonthlyComparisonMetricCardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGMAUMonthlyComparisonMetricCardResponseParams `json:"Response"`
}

func (r *DescribeMNGMAUMonthlyComparisonMetricCardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGMAUMonthlyComparisonMetricCardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGPaymentLineChartRequestParams struct {
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: payment_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// Valid values:
	// mng_paid_amount: Virtual payment amount,
	// paid_user_num: Number of paying users,
	// new_paid_user_num: Number of new paying users,
	// new_paid_user_amount: Revenue from new paying users,
	// new_paid_user_ratio: Percentage of new users who made payments,
	// arppu: Average revenue per paying user (ARPPU),
	// mng_refund_num: Number of refund orders
	// mng_refund_amount: Refund amount
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"Platform":0,"MnpId":"mgcp5xc2yzw8aixv","BeginDate":20251028,"EndDate":20251126,"DataType":"1","PaymentType":2}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

type DescribeMNGPaymentLineChartRequest struct {
	*tchttp.BaseRequest
	
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: payment_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// Valid values:
	// mng_paid_amount: Virtual payment amount,
	// paid_user_num: Number of paying users,
	// new_paid_user_num: Number of new paying users,
	// new_paid_user_amount: Revenue from new paying users,
	// new_paid_user_ratio: Percentage of new users who made payments,
	// arppu: Average revenue per paying user (ARPPU),
	// mng_refund_num: Number of refund orders
	// mng_refund_amount: Refund amount
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"Platform":0,"MnpId":"mgcp5xc2yzw8aixv","BeginDate":20251028,"EndDate":20251126,"DataType":"1","PaymentType":2}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

func (r *DescribeMNGPaymentLineChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGPaymentLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformId")
	delete(f, "ReportId")
	delete(f, "IndexId")
	delete(f, "QueryData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGPaymentLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGPaymentLineChartResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*ReportDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGPaymentLineChartResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGPaymentLineChartResponseParams `json:"Response"`
}

func (r *DescribeMNGPaymentLineChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGPaymentLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGPaymentOverviewRequestParams struct {
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Operating system: 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNGPaymentOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Operating system: 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNGPaymentOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGPaymentOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "DataType")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGPaymentOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGPaymentOverviewResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *MNGPaymentOverview `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGPaymentOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGPaymentOverviewResponseParams `json:"Response"`
}

func (r *DescribeMNGPaymentOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGPaymentOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGPaymentReportDetailRequestParams struct {
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Operating system: 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNGPaymentReportDetailRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Operating system: 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNGPaymentReportDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGPaymentReportDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "DataType")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGPaymentReportDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGPaymentReportDetailResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*MNGPaymentOverview `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGPaymentReportDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGPaymentReportDetailResponseParams `json:"Response"`
}

func (r *DescribeMNGPaymentReportDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGPaymentReportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGPaymentRetentionAnalysisRequestParams struct {
	// Start time in YYYYMMDD format
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time in YYYYMMDD format
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNGPaymentRetentionAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time in YYYYMMDD format
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 0 All, 2-Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNGPaymentRetentionAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGPaymentRetentionAnalysisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "DataType")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGPaymentRetentionAnalysisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGPaymentRetentionAnalysisResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*PaymentActiveRetention `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGPaymentRetentionAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGPaymentRetentionAnalysisResponseParams `json:"Response"`
}

func (r *DescribeMNGPaymentRetentionAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGPaymentRetentionAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGRetentionDataRequestParams struct {
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 1 Production data, 0 Non-production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *uint64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNGRetentionDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 1 Production data, 0 Non-production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *uint64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNGRetentionDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGRetentionDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "DataType")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNGRetentionDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNGRetentionDataResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*RetentionData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNGRetentionDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNGRetentionDataResponseParams `json:"Response"`
}

func (r *DescribeMNGRetentionDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNGRetentionDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPAccessAnalysisOverviewRequestParams struct {
	// Start time: 20251123 (example)
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time: 20251123 (example)
	TimeEnd *uint64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 1 Production data, 0 Non-production data
	ProdData *int64 `json:"ProdData,omitnil,omitempty" name:"ProdData"`

	// Operating system: 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNPAccessAnalysisOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Start time: 20251123 (example)
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time: 20251123 (example)
	TimeEnd *uint64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 1 Production data, 0 Non-production data
	ProdData *int64 `json:"ProdData,omitnil,omitempty" name:"ProdData"`

	// Operating system: 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNPAccessAnalysisOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPAccessAnalysisOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "TimeEnd")
	delete(f, "ProdData")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPAccessAnalysisOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPAccessAnalysisOverviewResponseParams struct {
	// Data display at the top of the page
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *AccessAnalysisOverview `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPAccessAnalysisOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPAccessAnalysisOverviewResponseParams `json:"Response"`
}

func (r *DescribeMNPAccessAnalysisOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPAccessAnalysisOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPActiveUserRealTimeStatisticsRequestParams struct {
	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId 
	// realtime_pv_num or realtime_uv_num
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"Platform":0,"DataType":"1","BeginDate":"20251125","EndDate":"20251125"}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

type DescribeMNPActiveUserRealTimeStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId 
	// realtime_pv_num or realtime_uv_num
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"Platform":0,"DataType":"1","BeginDate":"20251125","EndDate":"20251125"}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

func (r *DescribeMNPActiveUserRealTimeStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPActiveUserRealTimeStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "ReportId")
	delete(f, "IndexId")
	delete(f, "QueryData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPActiveUserRealTimeStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPActiveUserRealTimeStatisticsResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*ReportDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPActiveUserRealTimeStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPActiveUserRealTimeStatisticsResponseParams `json:"Response"`
}

func (r *DescribeMNPActiveUserRealTimeStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPActiveUserRealTimeStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPAdvertisingDetailRequestParams struct {
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// // 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`
}

type DescribeMNPAdvertisingDetailRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	TimeBegin *string `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// // 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	// //1-BANNER  2-REWARDED
	AdUnitType *string `json:"AdUnitType,omitnil,omitempty" name:"AdUnitType"`
}

func (r *DescribeMNPAdvertisingDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPAdvertisingDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	delete(f, "AdUnitType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPAdvertisingDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPAdvertisingDetailResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*MAUDetailData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPAdvertisingDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPAdvertisingDetailResponseParams `json:"Response"`
}

func (r *DescribeMNPAdvertisingDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPAdvertisingDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPAllStageVersionsRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeMNPAllStageVersionsRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeMNPAllStageVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPAllStageVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPAllStageVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPAllStageVersionsResponseParams struct {
	// Response data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DescribeMPAllStageVersionsResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPAllStageVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPAllStageVersionsResponseParams `json:"Response"`
}

func (r *DescribeMNPAllStageVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPAllStageVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPApprovalListData struct {
	// Approval ticket ID
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Approval status. valid values: 1 (processing), 2 (rejected), 3 (approved), 4 (cancelled).
	ApprovalStatus *int64 `json:"ApprovalStatus,omitnil,omitempty" name:"ApprovalStatus"`

	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program version.
	MNPVersion *string `json:"MNPVersion,omitnil,omitempty" name:"MNPVersion"`

	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Applicant
	ApplyUser *string `json:"ApplyUser,omitnil,omitempty" name:"ApplyUser"`

	// Application time
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Mini program icon
	MNPIcon *string `json:"MNPIcon,omitnil,omitempty" name:"MNPIcon"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Specifies the application icon.
	ApplicationLogo *string `json:"ApplicationLogo,omitnil,omitempty" name:"ApplicationLogo"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Mini program review qr code.
	MNPQrCodeUrl *string `json:"MNPQrCodeUrl,omitnil,omitempty" name:"MNPQrCodeUrl"`

	// Mini program type
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Specifies the reviewer.
	ApprovalUser *string `json:"ApprovalUser,omitnil,omitempty" name:"ApprovalUser"`

	// Approval time.
	ApprovalTime *string `json:"ApprovalTime,omitnil,omitempty" name:"ApprovalTime"`

	// Approval notes
	ApprovalNote *string `json:"ApprovalNote,omitnil,omitempty" name:"ApprovalNote"`
}

// Predefined struct for user
type DescribeMNPApprovalListRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Approval status. 1: Processing; 2: Rejected; 3: Approved; 4 Cancelled
	ApprovalStatusList []*int64 `json:"ApprovalStatusList,omitnil,omitempty" name:"ApprovalStatusList"`

	// Keywords of the mini program name to search
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

type DescribeMNPApprovalListRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Approval status. 1: Processing; 2: Rejected; 3: Approved; 4 Cancelled
	ApprovalStatusList []*int64 `json:"ApprovalStatusList,omitnil,omitempty" name:"ApprovalStatusList"`

	// Keywords of the mini program name to search
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

func (r *DescribeMNPApprovalListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPApprovalListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "ApprovalStatusList")
	delete(f, "Keyword")
	delete(f, "ApplicationId")
	delete(f, "TeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPApprovalListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPApprovalListResp struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List information
	DataList []*DescribeMNPApprovalListData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeMNPApprovalListResponseParams struct {
	// Response data
	Data *DescribeMNPApprovalListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPApprovalListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPApprovalListResponseParams `json:"Response"`
}

func (r *DescribeMNPApprovalListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPApprovalListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPCategoryRequestParams struct {
	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeMNPCategoryRequest struct {
	*tchttp.BaseRequest
	
	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeMNPCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPCategoryResponseParams struct {
	// Response data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*MNPTypeDefine `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPCategoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPCategoryResponseParams `json:"Response"`
}

func (r *DescribeMNPCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPDomainACLRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeMNPDomainACLRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeMNPDomainACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPDomainACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPDomainACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPDomainACLResponseParams struct {
	// Response data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DescribeDomainInfoParam `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPDomainACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPDomainACLResponseParams `json:"Response"`
}

func (r *DescribeMNPDomainACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPDomainACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPListData struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program icon
	MNPIcon *string `json:"MNPIcon,omitnil,omitempty" name:"MNPIcon"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Name of the associated team
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Mini program type
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Specifies the mini program listing status. valid values: 1 (submitted), 2 (removed).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Mini program introduction
	MNPIntro *string `json:"MNPIntro,omitnil,omitempty" name:"MNPIntro"`

	// Creator
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Specifies the updater.
	UpdateUser *string `json:"UpdateUser,omitnil,omitempty" name:"UpdateUser"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Effective status of the bound application. valid values: 1 (not effective), 2 (effective).
	EffectStatus *int64 `json:"EffectStatus,omitnil,omitempty" name:"EffectStatus"`

	// Specifies the ID of the application bound with the mini program.
	EffectMNPVersionId *int64 `json:"EffectMNPVersionId,omitnil,omitempty" name:"EffectMNPVersionId"`

	// Specifies the effective version number of the application bound to the mini program.
	EffectMNPVersion *string `json:"EffectMNPVersion,omitnil,omitempty" name:"EffectMNPVersion"`
}

// Predefined struct for user
type DescribeMNPListRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Keywords for search (mini program name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type DescribeMNPListRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Keywords for search (mini program name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

func (r *DescribeMNPListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "Keyword")
	delete(f, "TeamId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPListResp struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List information
	DataList []*DescribeMNPListData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeMNPListResponseParams struct {
	// Response data
	Data *DescribeMNPListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPListResponseParams `json:"Response"`
}

func (r *DescribeMNPListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPMAUDataDetailRequestParams struct {
	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Program team ID, -1 means not provided
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

type DescribeMNPMAUDataDetailRequest struct {
	*tchttp.BaseRequest
	
	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Program team ID, -1 means not provided
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

func (r *DescribeMNPMAUDataDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPMAUDataDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "ApplicationId")
	delete(f, "MNPId")
	delete(f, "MNPTeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPMAUDataDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPMAUDataDetailResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*MAUDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPMAUDataDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPMAUDataDetailResponseParams `json:"Response"`
}

func (r *DescribeMNPMAUDataDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPMAUDataDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPMAULineChartRequestParams struct {
	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program team ID
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

type DescribeMNPMAULineChartRequest struct {
	*tchttp.BaseRequest
	
	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program team ID
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

func (r *DescribeMNPMAULineChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPMAULineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "ApplicationId")
	delete(f, "MNPId")
	delete(f, "MNPTeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPMAULineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPMAULineChartResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*MAUChartData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPMAULineChartResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPMAULineChartResponseParams `json:"Response"`
}

func (r *DescribeMNPMAULineChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPMAULineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPMAUMetricCardRequestParams struct {
	// Start time in YYYYMMDD format
	SourceMonth *int64 `json:"SourceMonth,omitnil,omitempty" name:"SourceMonth"`

	// Type: 0 Non-production data, 1 Production data 
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TargetMonth *int64 `json:"TargetMonth,omitnil,omitempty" name:"TargetMonth"`

	// Superapp ID starting with App
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program team ID, required. When provided, the query is performed based on the mini program team.
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

type DescribeMNPMAUMetricCardRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	SourceMonth *int64 `json:"SourceMonth,omitnil,omitempty" name:"SourceMonth"`

	// Type: 0 Non-production data, 1 Production data 
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// End time in YYYYMMDD format
	TargetMonth *int64 `json:"TargetMonth,omitnil,omitempty" name:"TargetMonth"`

	// Superapp ID starting with App
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Mini program appid, required. When provided, the query is performed based on the mini program.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program team ID, required. When provided, the query is performed based on the mini program team.
	MNPTeamId *int64 `json:"MNPTeamId,omitnil,omitempty" name:"MNPTeamId"`
}

func (r *DescribeMNPMAUMetricCardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPMAUMetricCardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceMonth")
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "TargetMonth")
	delete(f, "ApplicationId")
	delete(f, "MNPId")
	delete(f, "MNPTeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPMAUMetricCardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPMAUMetricCardResponseParams struct {
	// Data display at the top of the page
	Data *MAUIndicatorCard `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPMAUMetricCardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPMAUMetricCardResponseParams `json:"Response"`
}

func (r *DescribeMNPMAUMetricCardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPMAUMetricCardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPManagerDetailData struct {
	// Mini program type 
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Mini program icon
	MNPIcon *string `json:"MNPIcon,omitnil,omitempty" name:"MNPIcon"`

	// Mini program introduction
	MNPIntro *string `json:"MNPIntro,omitnil,omitempty" name:"MNPIntro"`

	// Mini program description
	MNPDesc *string `json:"MNPDesc,omitnil,omitempty" name:"MNPDesc"`

	// Creation time, timestamp.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Creator
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// Access status. 1: not connected; 2: connected.
	AccessStatus *int64 `json:"AccessStatus,omitnil,omitempty" name:"AccessStatus"`

	// Name of the associated team
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Specifies the mini program listing status. valid values: 1 (submitted), 2 (removed).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeMNPOfflinePackageURLRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeMNPOfflinePackageURLRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeMNPOfflinePackageURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPOfflinePackageURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPOfflinePackageURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPOfflinePackageURLResponseParams struct {
	// Response data
	Data *StringData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPOfflinePackageURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPOfflinePackageURLResponseParams `json:"Response"`
}

func (r *DescribeMNPOfflinePackageURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPOfflinePackageURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPPageAnalysisDetailRequestParams struct {
	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Start time in YYYYMMDD format
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// End time in YYYYMMDD format
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNPPageAnalysisDetailRequest struct {
	*tchttp.BaseRequest
	
	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Start time in YYYYMMDD format
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// End time in YYYYMMDD format
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNPPageAnalysisDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPPageAnalysisDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "MNPId")
	delete(f, "TimeBegin")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPPageAnalysisDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPPageAnalysisDetailResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*VisitData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPPageAnalysisDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPPageAnalysisDetailResponseParams `json:"Response"`
}

func (r *DescribeMNPPageAnalysisDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPPageAnalysisDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPPreviewRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeMNPPreviewRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeMNPPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "MNPVersionId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPPreviewResp struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Mini program description
	MNPDesc *string `json:"MNPDesc,omitnil,omitempty" name:"MNPDesc"`

	// Mini program version.
	MNPVersion *string `json:"MNPVersion,omitnil,omitempty" name:"MNPVersion"`

	// Describes the mini program version.
	MNPVersionIntro *string `json:"MNPVersionIntro,omitnil,omitempty" name:"MNPVersionIntro"`

	// Specifies the mini program qr code.
	QRCodeUrl *string `json:"QRCodeUrl,omitnil,omitempty" name:"QRCodeUrl"`

	// Specifies the path to the preview version.
	PreviewEntrancePath *string `json:"PreviewEntrancePath,omitnil,omitempty" name:"PreviewEntrancePath"`

	// Specifies the qr code content.
	QRCodeContent *string `json:"QRCodeContent,omitnil,omitempty" name:"QRCodeContent"`
}

// Predefined struct for user
type DescribeMNPPreviewResponseParams struct {
	// Response data
	Data *DescribeMNPPreviewResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPPreviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPPreviewResponseParams `json:"Response"`
}

func (r *DescribeMNPPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPReleasedVersionHistoryRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeMNPReleasedVersionHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeMNPReleasedVersionHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPReleasedVersionHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPReleasedVersionHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPReleasedVersionHistoryResponseParams struct {
	// Response data
	Data *DescribeRevertOnlineVersionPageResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPReleasedVersionHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPReleasedVersionHistoryResponseParams `json:"Response"`
}

func (r *DescribeMNPReleasedVersionHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPReleasedVersionHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPReportDataLineChartRequestParams struct {
	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId (optional):
	// active_device_num: Active users
	// new_device_num: New users
	// total_user_num: Total users
	// share_num: Number of shares
	// miniapp_open_num: Number of opens
	// avg_device_open_num: Average opens per user
	// avg_device_open_duration: Average duration per user
	// avg_count_open_duration: Average duration per session
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"DataType":"1","Platform":0,"BeginDate":20251118,"EndDate":20251124}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

type DescribeMNPReportDataLineChartRequest struct {
	*tchttp.BaseRequest
	
	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: mnp_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId (optional):
	// active_device_num: Active users
	// new_device_num: New users
	// total_user_num: Total users
	// share_num: Number of shares
	// miniapp_open_num: Number of opens
	// avg_device_open_num: Average opens per user
	// avg_device_open_duration: Average duration per user
	// avg_count_open_duration: Average duration per session
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"DataType":"1","Platform":0,"BeginDate":20251118,"EndDate":20251124}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

func (r *DescribeMNPReportDataLineChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPReportDataLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "ReportId")
	delete(f, "IndexId")
	delete(f, "QueryData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPReportDataLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPReportDataLineChartResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*ReportDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPReportDataLineChartResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPReportDataLineChartResponseParams `json:"Response"`
}

func (r *DescribeMNPReportDataLineChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPReportDataLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPReportDetailRequestParams struct {
	// Start time in YYYYMMDD format
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	// End time in YYYYMMDD format
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`
}

type DescribeMNPReportDetailRequest struct {
	*tchttp.BaseRequest
	
	// Start time in YYYYMMDD format
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	// End time in YYYYMMDD format
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`
}

func (r *DescribeMNPReportDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPReportDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "DataType")
	delete(f, "PlatformId")
	delete(f, "Platform")
	delete(f, "TimeEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPReportDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPReportDetailResponseParams struct {
	// Data display at the top of the page
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*AccessAnalysisDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPReportDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPReportDetailResponseParams `json:"Response"`
}

func (r *DescribeMNPReportDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPReportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeMNPRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeMNPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPResponseParams struct {
	// Response data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *DescribeMNPManagerDetailData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPResponseParams `json:"Response"`
}

func (r *DescribeMNPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPRetentionDataRequestParams struct {
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 1 Production data, 0 Non-production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeMNPRetentionDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 1 Production data, 0 Non-production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeMNPRetentionDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPRetentionDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "DataType")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPRetentionDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPRetentionDataResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*RetentionData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPRetentionDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPRetentionDataResponseParams `json:"Response"`
}

func (r *DescribeMNPRetentionDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPRetentionDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPSensitiveAPIPermissionApprovalData struct {
	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// API method.
	APIMethod *string `json:"APIMethod,omitnil,omitempty" name:"APIMethod"`

	// Reason for application
	ApplyReason *string `json:"ApplyReason,omitnil,omitempty" name:"ApplyReason"`

	// Reason for rejection.
	RejectReason *string `json:"RejectReason,omitnil,omitempty" name:"RejectReason"`

	// Approval status. valid values: 20 (rejected), 30 (approved).
	ApprovalStatus *int64 `json:"ApprovalStatus,omitnil,omitempty" name:"ApprovalStatus"`

	// API feature description.
	APIDesc *string `json:"APIDesc,omitnil,omitempty" name:"APIDesc"`

	// API type. 1: system; 2: custom.
	APIType *int64 `json:"APIType,omitnil,omitempty" name:"APIType"`
}

type DescribeMNPSensitiveAPIPermissionApprovalListData struct {
	// Approval ID
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Sensitive API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// API name
	APIName *string `json:"APIName,omitnil,omitempty" name:"APIName"`

	// API request method
	APIMethod *string `json:"APIMethod,omitnil,omitempty" name:"APIMethod"`

	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Applicant
	ApplyUser *string `json:"ApplyUser,omitnil,omitempty" name:"ApplyUser"`

	// Application time
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// Application notes
	ApplyNote *string `json:"ApplyNote,omitnil,omitempty" name:"ApplyNote"`

	// Approval status. 1: Processing; 20: Rejected; 30: Approved
	ApprovalStatus *int64 `json:"ApprovalStatus,omitnil,omitempty" name:"ApprovalStatus"`

	// Specifies the review user.
	ApprovalUser *string `json:"ApprovalUser,omitnil,omitempty" name:"ApprovalUser"`

	// Approval time.
	ApprovalTime *string `json:"ApprovalTime,omitnil,omitempty" name:"ApprovalTime"`

	// Approval notes
	ApprovalNote *string `json:"ApprovalNote,omitnil,omitempty" name:"ApprovalNote"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Specifies the application icon.
	ApplicationLogo *string `json:"ApplicationLogo,omitnil,omitempty" name:"ApplicationLogo"`

	// API type. 1: system; 2: custom.
	APIType *int64 `json:"APIType,omitnil,omitempty" name:"APIType"`

	// API feature description.
	APIDesc *string `json:"APIDesc,omitnil,omitempty" name:"APIDesc"`
}

// Predefined struct for user
type DescribeMNPSensitiveAPIPermissionApprovalListRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Approval status. 1: Processing; 20: Rejected; 30: Approved
	ApprovalStatusList []*int64 `json:"ApprovalStatusList,omitnil,omitempty" name:"ApprovalStatusList"`

	// Keywords for search (API name, API method or application name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeMNPSensitiveAPIPermissionApprovalListRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Approval status. 1: Processing; 20: Rejected; 30: Approved
	ApprovalStatusList []*int64 `json:"ApprovalStatusList,omitnil,omitempty" name:"ApprovalStatusList"`

	// Keywords for search (API name, API method or application name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeMNPSensitiveAPIPermissionApprovalListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPSensitiveAPIPermissionApprovalListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "ApprovalStatusList")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPSensitiveAPIPermissionApprovalListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPSensitiveAPIPermissionApprovalListResp struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List information
	DataList []*DescribeMNPSensitiveAPIPermissionApprovalListData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeMNPSensitiveAPIPermissionApprovalListResponseParams struct {
	// Response data
	Data *DescribeMNPSensitiveAPIPermissionApprovalListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPSensitiveAPIPermissionApprovalListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPSensitiveAPIPermissionApprovalListResponseParams `json:"Response"`
}

func (r *DescribeMNPSensitiveAPIPermissionApprovalListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPSensitiveAPIPermissionApprovalListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPSensitiveAPIPermissionApprovalRequestParams struct {
	// Approval request number
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeMNPSensitiveAPIPermissionApprovalRequest struct {
	*tchttp.BaseRequest
	
	// Approval request number
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeMNPSensitiveAPIPermissionApprovalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPSensitiveAPIPermissionApprovalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApprovalNo")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPSensitiveAPIPermissionApprovalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPSensitiveAPIPermissionApprovalResponseParams struct {
	// Response data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *DescribeMNPSensitiveAPIPermissionApprovalData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPSensitiveAPIPermissionApprovalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPSensitiveAPIPermissionApprovalResponseParams `json:"Response"`
}

func (r *DescribeMNPSensitiveAPIPermissionApprovalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPSensitiveAPIPermissionApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPSensitiveAPIPermissionListData struct {
	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// API name.
	APIName *string `json:"APIName,omitnil,omitempty" name:"APIName"`

	// API request method
	APIMethod *string `json:"APIMethod,omitnil,omitempty" name:"APIMethod"`

	// API status.
	APIStatus *int64 `json:"APIStatus,omitnil,omitempty" name:"APIStatus"`

	// API application status.
	APIApplyStatus *int64 `json:"APIApplyStatus,omitnil,omitempty" name:"APIApplyStatus"`

	// Reason for rejection.
	RejectReason *string `json:"RejectReason,omitnil,omitempty" name:"RejectReason"`

	// Approval ID
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Specifies the application icon.
	ApplicationIcon *string `json:"ApplicationIcon,omitnil,omitempty" name:"ApplicationIcon"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// API type. 1: system; 2: custom.
	APIType *int64 `json:"APIType,omitnil,omitempty" name:"APIType"`

	// API feature description.
	APIDesc *string `json:"APIDesc,omitnil,omitempty" name:"APIDesc"`
}

// Predefined struct for user
type DescribeMNPSensitiveAPIPermissionListRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Keywords for search (API name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeMNPSensitiveAPIPermissionListRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Keywords for search (API name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeMNPSensitiveAPIPermissionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPSensitiveAPIPermissionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "ApplicationId")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPSensitiveAPIPermissionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPSensitiveAPIPermissionListResp struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List data
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataList []*DescribeMNPSensitiveAPIPermissionListData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeMNPSensitiveAPIPermissionListResponseParams struct {
	// Response parameters.
	Data *DescribeMNPSensitiveAPIPermissionListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPSensitiveAPIPermissionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPSensitiveAPIPermissionListResponseParams `json:"Response"`
}

func (r *DescribeMNPSensitiveAPIPermissionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPSensitiveAPIPermissionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMNPVersionRequestParams struct {
	// ID of the task to create a mini program version
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeMNPVersionRequest struct {
	*tchttp.BaseRequest
	
	// ID of the task to create a mini program version
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeMNPVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMNPVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMNPVersionResp struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 1: Pending; 20: Running; 30: Failed; 60: Succeeded
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task status message
	TaskMsg *string `json:"TaskMsg,omitnil,omitempty" name:"TaskMsg"`

	// Mini program version ID (returned when compilation succeeds)
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`
}

// Predefined struct for user
type DescribeMNPVersionResponseParams struct {
	// Response data
	Data *DescribeMNPVersionResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMNPVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMNPVersionResponseParams `json:"Response"`
}

func (r *DescribeMNPVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMNPVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMPAllStageVersionsResp struct {
	// Mini program ID.
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Specifies the mini program version primary key id.
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Specifies the mini program avatar.
	MNPIcon *string `json:"MNPIcon,omitnil,omitempty" name:"MNPIcon"`

	// Mini program type
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Mini program introduction
	MNPIntro *string `json:"MNPIntro,omitnil,omitempty" name:"MNPIntro"`

	// Mini program description
	MNPDesc *string `json:"MNPDesc,omitnil,omitempty" name:"MNPDesc"`

	// Specifies the developer.
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// Developer creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Mini program version.
	MNPVersion *string `json:"MNPVersion,omitnil,omitempty" name:"MNPVersion"`

	// Describes version features.
	MNPVersionIntro *string `json:"MNPVersionIntro,omitnil,omitempty" name:"MNPVersionIntro"`

	// Development Platform Online.
	Phase *string `json:"Phase,omitnil,omitempty" name:"Phase"`

	// 0 pending review; 1 under review; 2 review rejection; 3 pass review; 4 review cancellation.
	ApprovalStatus *int64 `json:"ApprovalStatus,omitnil,omitempty" name:"ApprovalStatus"`

	// Approval ticket ID
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Specifies whether it is a trial version.
	// Specifies the version type. valid values: 0 (non-preview version); 1 (trial version).
	ShowCase *int64 `json:"ShowCase,omitnil,omitempty" name:"ShowCase"`

	// Version number to roll back to.
	RollbackVersion *int64 `json:"RollbackVersion,omitnil,omitempty" name:"RollbackVersion"`

	// Indicates the release status.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Specifies the current main status of the version. valid values: "0" (pending review), "1" (under review), "2" (review rejection), "3" (pass review), "4" (review cancellation).
	VersionCurrentStatus *int64 `json:"VersionCurrentStatus,omitnil,omitempty" name:"VersionCurrentStatus"`
}

// Predefined struct for user
type DescribePaymentDataDetailRequestParams struct {
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPIds []*string `json:"MNPIds,omitnil,omitempty" name:"MNPIds"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Operating system: 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribePaymentDataDetailRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPIds []*string `json:"MNPIds,omitnil,omitempty" name:"MNPIds"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Operating system: 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribePaymentDataDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePaymentDataDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPIds")
	delete(f, "PlatformId")
	delete(f, "DataType")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePaymentDataDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePaymentDataDetailResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*PaymentDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePaymentDataDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePaymentDataDetailResponseParams `json:"Response"`
}

func (r *DescribePaymentDataDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePaymentDataDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePaymentDataLineChartRequestParams struct {
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: payment_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId (optional):
	// order_user_num: Number of users placing orders
	// order_num: Total orders
	// total_amount: Total amount
	// order_unpaid_num: Total unpaid orders
	// unpaid_amount: Unpaid amount
	// order_paid_num: Total paid orders
	// paid_amount: Amount paid
	// order_refund_num: Total refunded orders
	// refund_amount: Total amount refunded
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"Platform":0,"MnpIds":["mp9e7qluz4i3z3km"],"BeginDate":20251031,"EndDate":20251129,"DataType":"1","PaymentType":1}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

type DescribePaymentDataLineChartRequest struct {
	*tchttp.BaseRequest
	
	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Fixed value: payment_data_analysis
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// IndexId (optional):
	// order_user_num: Number of users placing orders
	// order_num: Total orders
	// total_amount: Total amount
	// order_unpaid_num: Total unpaid orders
	// unpaid_amount: Unpaid amount
	// order_paid_num: Total paid orders
	// paid_amount: Amount paid
	// order_refund_num: Total refunded orders
	// refund_amount: Total amount refunded
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// Input parameter base64 string: {"Platform":0,"MnpIds":["mp9e7qluz4i3z3km"],"BeginDate":20251031,"EndDate":20251129,"DataType":"1","PaymentType":1}
	QueryData *string `json:"QueryData,omitnil,omitempty" name:"QueryData"`
}

func (r *DescribePaymentDataLineChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePaymentDataLineChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformId")
	delete(f, "ReportId")
	delete(f, "IndexId")
	delete(f, "QueryData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePaymentDataLineChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePaymentDataLineChartResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*ReportDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePaymentDataLineChartResponse struct {
	*tchttp.BaseResponse
	Response *DescribePaymentDataLineChartResponseParams `json:"Response"`
}

func (r *DescribePaymentDataLineChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePaymentDataLineChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePaymentDataOverviewRequestParams struct {
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPIds []*string `json:"MNPIds,omitnil,omitempty" name:"MNPIds"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Operating system: 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribePaymentDataOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	TimeBegin *int64 `json:"TimeBegin,omitnil,omitempty" name:"TimeBegin"`

	// Mini program appid
	MNPIds []*string `json:"MNPIds,omitnil,omitempty" name:"MNPIds"`

	// Tenant ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Type: 0 Non-production data, 1 Production data
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// End time
	TimeEnd *int64 `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Operating system: 0 All, 2 Android, 3 iOS
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribePaymentDataOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePaymentDataOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeBegin")
	delete(f, "MNPIds")
	delete(f, "PlatformId")
	delete(f, "DataType")
	delete(f, "TimeEnd")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePaymentDataOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePaymentDataOverviewResponseParams struct {
	// Data display
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *MNPPaymentOverview `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePaymentDataOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribePaymentDataOverviewResponseParams `json:"Response"`
}

func (r *DescribePaymentDataOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePaymentDataOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRevertOnlineVersionPageResp struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List information
	DataList []*QueryOnlineVersionResp `json:"DataList,omitnil,omitempty" name:"DataList"`
}

type DescribeRoleListData struct {
	// Role ID
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// Role name.
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Role type 1-preset role 2-custom role.
	RoleType *int64 `json:"RoleType,omitnil,omitempty" name:"RoleType"`
}

// Predefined struct for user
type DescribeRoleListRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Keywords for search (role name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

type DescribeRoleListRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Keywords for search (role name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

func (r *DescribeRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "Keyword")
	delete(f, "TeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleListResp struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List data
	DataList []*DescribeRoleListData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeRoleListResponseParams struct {
	// Response data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *DescribeRoleListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRoleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoleListResponseParams `json:"Response"`
}

func (r *DescribeRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamDetailResp struct {
	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Team role type 1-mini program team 2-application team
	TeamRoleType *int64 `json:"TeamRoleType,omitnil,omitempty" name:"TeamRoleType"`

	// Administrator account
	AdminUserAccount *string `json:"AdminUserAccount,omitnil,omitempty" name:"AdminUserAccount"`

	// Creator
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Number of team members
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`

	// Number of bound mini program teams
	BindMiniTeamCount *int64 `json:"BindMiniTeamCount,omitnil,omitempty" name:"BindMiniTeamCount"`

	// Name of the bound team
	BindTeamName *string `json:"BindTeamName,omitnil,omitempty" name:"BindTeamName"`

	// Team registration link
	RegisterLink *string `json:"RegisterLink,omitnil,omitempty" name:"RegisterLink"`

	// Application name. It Is required when querying details of a mini program team.
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Team expiration time. 0 means never expire.
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Team status. valid values: 1: normal; 2: disabled; 3: expired.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeTeamListInfoResp struct {
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Administrator user ID
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// Administrator account
	AdminUserAccount *string `json:"AdminUserAccount,omitnil,omitempty" name:"AdminUserAccount"`

	// Administrator username
	AdminUserName *string `json:"AdminUserName,omitnil,omitempty" name:"AdminUserName"`

	// Number of team members
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`

	// Team registration link
	RegisterLink *string `json:"RegisterLink,omitnil,omitempty" name:"RegisterLink"`

	// Team permission type
	TeamRoleTypeList []*int64 `json:"TeamRoleTypeList,omitnil,omitempty" name:"TeamRoleTypeList"`

	// Associated team ID
	RelatedTeamId *int64 `json:"RelatedTeamId,omitnil,omitempty" name:"RelatedTeamId"`

	// Team expiration time. 0 means never expire.
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Team status. valid values: 1: normal; 2: disabled; 3: expired.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeTeamListRequestParams struct {
	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team name to be queried
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeTeamListRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team name to be queried
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeTeamListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTeamListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTeamListResponseParams struct {
	// Response data
	Data *DescribeTeamPageResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTeamListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTeamListResponseParams `json:"Response"`
}

func (r *DescribeTeamListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamMemberInfoResp struct {
	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// User account
	UserAccount *string `json:"UserAccount,omitnil,omitempty" name:"UserAccount"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Specifies the team role name.
	TeamRoleName *string `json:"TeamRoleName,omitnil,omitempty" name:"TeamRoleName"`

	// Specifies the team role ID.
	TeamRoleId *int64 `json:"TeamRoleId,omitnil,omitempty" name:"TeamRoleId"`

	// Whether it is editable
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`
}

type DescribeTeamMemberListPageResp struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List information
	DataList []*DescribeTeamMemberInfoResp `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeTeamMemberListRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Keywords for search (user name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Role ID
	RoleIds []*int64 `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type DescribeTeamMemberListRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Keywords for search (user name)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Role ID
	RoleIds []*int64 `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *DescribeTeamMemberListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamMemberListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "TeamId")
	delete(f, "Keyword")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTeamMemberListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTeamMemberListResponseParams struct {
	// Response data
	Data *DescribeTeamMemberListPageResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTeamMemberListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTeamMemberListResponseParams `json:"Response"`
}

func (r *DescribeTeamMemberListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamMemberListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamPageResp struct {
	// Total count
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List information
	DataList []*DescribeTeamListInfoResp `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeTeamRequestParams struct {
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeTeamRequest struct {
	*tchttp.BaseRequest
	
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTeamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTeamResponseParams struct {
	// Response data
	Data *DescribeTeamDetailResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTeamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTeamResponseParams `json:"Response"`
}

func (r *DescribeTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTempSecret4UploadFile2CosRequestParams struct {
	// Service name
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// File suffix
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`
}

type DescribeTempSecret4UploadFile2CosRequest struct {
	*tchttp.BaseRequest
	
	// Service name
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// File suffix
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`
}

func (r *DescribeTempSecret4UploadFile2CosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTempSecret4UploadFile2CosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessName")
	delete(f, "Suffix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTempSecret4UploadFile2CosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTempSecret4UploadFile2CosResponseParams struct {
	// Response data
	Data *UploadFileTempSecret `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTempSecret4UploadFile2CosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTempSecret4UploadFile2CosResponseParams `json:"Response"`
}

func (r *DescribeTempSecret4UploadFile2CosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTempSecret4UploadFile2CosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDetailResp struct {
	// User iD.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// User account
	UserAccount *string `json:"UserAccount,omitnil,omitempty" name:"UserAccount"`

	// User account
	// 1 - super admin 2 - platform admin 3 - ordinary member. not passing indicates all.
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type DescribeUserListData struct {
	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// User account
	UserAccount *string `json:"UserAccount,omitnil,omitempty" name:"UserAccount"`

	// Account type. 1: super administrator; 2: platform administrator; 3: ordinary member.
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Status. 1: normal; 2: disabled.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`
}

// Predefined struct for user
type DescribeUserListRequestParams struct {
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Keywords for search (username or account)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// User account 1 - Super admin 2 - Platform admin 3 - Member
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

type DescribeUserListRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Keywords for search (username or account)
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// User account 1 - Super admin 2 - Platform admin 3 - Member
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

func (r *DescribeUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PlatformId")
	delete(f, "Keyword")
	delete(f, "AccountType")
	delete(f, "TeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserListResp struct {
	// Total number of entries
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List data
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataList []*DescribeUserListData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type DescribeUserListResponseParams struct {
	// Response data
	Data *DescribeUserListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserListResponseParams `json:"Response"`
}

func (r *DescribeUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRequestParams struct {
	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *DescribeUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResponseParams struct {
	// Response data
	Data *DescribeUserDetailResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserResponseParams `json:"Response"`
}

func (r *DescribeUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApplicationSensitiveAPIRequestParams struct {
	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`
}

type DisableApplicationSensitiveAPIRequest struct {
	*tchttp.BaseRequest
	
	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`
}

func (r *DisableApplicationSensitiveAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableApplicationSensitiveAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformId")
	delete(f, "APIId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableApplicationSensitiveAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApplicationSensitiveAPIResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableApplicationSensitiveAPIResponse struct {
	*tchttp.BaseResponse
	Response *DisableApplicationSensitiveAPIResponseParams `json:"Response"`
}

func (r *DisableApplicationSensitiveAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableApplicationSensitiveAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadApplicationConfigResp struct {
	// Configuration information in Base64 format
	// Note: This field may return null, indicating that no valid values can be obtained.
	File *string `json:"File,omitnil,omitempty" name:"File"`
}

// Predefined struct for user
type EnableApplicationSensitiveAPIRequestParams struct {
	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`
}

type EnableApplicationSensitiveAPIRequest struct {
	*tchttp.BaseRequest
	
	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`
}

func (r *EnableApplicationSensitiveAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableApplicationSensitiveAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformId")
	delete(f, "APIId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableApplicationSensitiveAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableApplicationSensitiveAPIResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableApplicationSensitiveAPIResponse struct {
	*tchttp.BaseResponse
	Response *EnableApplicationSensitiveAPIResponseParams `json:"Response"`
}

func (r *EnableApplicationSensitiveAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableApplicationSensitiveAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalDomainDeleteResp struct {
	// Result.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`
}

type GlobalDomainModifyRespResp struct {
	// 0: success; 1: allowed domains exist; 2: blocked domains exist.
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`
}

type MAUChartData struct {
	// Year and month: 2024-12
	DataTime *int64 `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// Value
	MAUCount *int64 `json:"MAUCount,omitnil,omitempty" name:"MAUCount"`

	// Data update time, only available when MAUCount is greater than 0.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type MAUDetail struct {
	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// MAU details
	DataList []*MAULineChartData `json:"DataList,omitnil,omitempty" name:"DataList"`
}

type MAUDetailData struct {
	// Date
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Mini program type
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Estimated revenue
	EstimatedEarnings *string `json:"EstimatedEarnings,omitnil,omitempty" name:"EstimatedEarnings"`

	// Requests
	RequestsNumber *int64 `json:"RequestsNumber,omitnil,omitempty" name:"RequestsNumber"`

	// Impressions
	Impressions *int64 `json:"Impressions,omitnil,omitempty" name:"Impressions"`

	// Effective Cost Per Mille
	ECPM *string `json:"ECPM,omitnil,omitempty" name:"ECPM"`

	// Taps
	ClicksNumber *int64 `json:"ClicksNumber,omitnil,omitempty" name:"ClicksNumber"`
}

type MAUIndicatorCard struct {
	// Growth rate (targetData - sourceData) / sourceData, returns 0 when SourceMAUNum is 0
	ComparisonRatio *string `json:"ComparisonRatio,omitnil,omitempty" name:"ComparisonRatio"`

	// 1 Increase
	// 2 Decrease
	// Returns 0 when SourceMAUNum is 0
	ComparisonResult *int64 `json:"ComparisonResult,omitnil,omitempty" name:"ComparisonResult"`

	// Last month's MAU data
	SourceMAUNum *int64 `json:"SourceMAUNum,omitnil,omitempty" name:"SourceMAUNum"`

	// This month's MAU data
	TargetMAUNum *int64 `json:"TargetMAUNum,omitnil,omitempty" name:"TargetMAUNum"`

	// Data timestamp
	FlushTime *int64 `json:"FlushTime,omitnil,omitempty" name:"FlushTime"`
}

type MAULineChartData struct {
	// Year-month date data
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataTime *int64 `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// MAU data
	// Note: This field may return null, indicating that no valid values can be obtained.
	MAUCount *string `json:"MAUCount,omitnil,omitempty" name:"MAUCount"`
}

type MNGMAULineChartData struct {
	// Year-month-date data
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataTime *int64 `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// MAU data
	// Note: This field may return null, indicating that no valid values can be obtained.
	MAUCount *int64 `json:"MAUCount,omitnil,omitempty" name:"MAUCount"`

	// Update time
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type MNGPaymentOverview struct {
	// Paid revenue / Number of paying users * 100%
	ARPPu *string `json:"ARPPu,omitnil,omitempty" name:"ARPPu"`

	// Data time in YYYYMMDD format
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// New paying user ratio = NewUserNum / OrderUserNum * 100%
	NewPaidUseRatio *string `json:"NewPaidUseRatio,omitnil,omitempty" name:"NewPaidUseRatio"`

	// Number of new paying users
	NewPaidUserNum *int64 `json:"NewPaidUserNum,omitnil,omitempty" name:"NewPaidUserNum"`

	// Total payment amount from new users
	NewUserPaidAmount *string `json:"NewUserPaidAmount,omitnil,omitempty" name:"NewUserPaidAmount"`

	// Total payment amount
	PaidAmount *string `json:"PaidAmount,omitnil,omitempty" name:"PaidAmount"`

	// Number of paying users
	PaidUserNum *int64 `json:"PaidUserNum,omitnil,omitempty" name:"PaidUserNum"`

	// Refund amount
	RefundAmount *string `json:"RefundAmount,omitnil,omitempty" name:"RefundAmount"`

	// Number of refund orders
	RefundNum *int64 `json:"RefundNum,omitnil,omitempty" name:"RefundNum"`

	// Update time (timestamp in seconds)
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type MNPAdOverview struct {
	// Estimated revenue
	// Note: This field may return null, indicating that no valid values can be obtained.
	EstimatedEarnings *string `json:"EstimatedEarnings,omitnil,omitempty" name:"EstimatedEarnings"`

	// Effective Cost Per Mille
	// Note: This field may return null, indicating that no valid values can be obtained.
	ECPM *string `json:"ECPM,omitnil,omitempty" name:"ECPM"`

	// Requests
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestsNumber *int64 `json:"RequestsNumber,omitnil,omitempty" name:"RequestsNumber"`

	// Impressions
	// Note: This field may return null, indicating that no valid values can be obtained.
	Impressions *int64 `json:"Impressions,omitnil,omitempty" name:"Impressions"`

	// Taps
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClicksNumber *int64 `json:"ClicksNumber,omitnil,omitempty" name:"ClicksNumber"`
}

type MNPAdvertisingOverview struct {
	// Data
	// Note: This field may return null, indicating that no valid values can be obtained.
	OverviewData *MNPAdOverview `json:"OverviewData,omitnil,omitempty" name:"OverviewData"`
}

type MNPPaymentOverview struct {
	// Number of mini programs involved in the order
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrderMNPNum *int64 `json:"OrderMNPNum,omitnil,omitempty" name:"OrderMNPNum"`

	// Total orders
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrderNum *int64 `json:"OrderNum,omitnil,omitempty" name:"OrderNum"`

	// Total paid orders
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrderPaidNum *int64 `json:"OrderPaidNum,omitnil,omitempty" name:"OrderPaidNum"`

	// Total refunded orders
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrderRefundNum *int64 `json:"OrderRefundNum,omitnil,omitempty" name:"OrderRefundNum"`

	// Total unpaid orders
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrderUnpaidNum *int64 `json:"OrderUnpaidNum,omitnil,omitempty" name:"OrderUnpaidNum"`

	// Total order users
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrderUserNum *int64 `json:"OrderUserNum,omitnil,omitempty" name:"OrderUserNum"`

	// Total paying users
	// Note: This field may return null, indicating that no valid values can be obtained.
	PaidUserNum *int64 `json:"PaidUserNum,omitnil,omitempty" name:"PaidUserNum"`

	// Amount paid
	// Note: This field may return null, indicating that no valid values can be obtained.
	PaidAmount *string `json:"PaidAmount,omitnil,omitempty" name:"PaidAmount"`

	// Total amount refunded
	// Note: This field may return null, indicating that no valid values can be obtained.
	RefundAmount *string `json:"RefundAmount,omitnil,omitempty" name:"RefundAmount"`

	// Total order amount
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalAmount *string `json:"TotalAmount,omitnil,omitempty" name:"TotalAmount"`

	// Unpaid amount
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnpaidAmount *string `json:"UnpaidAmount,omitnil,omitempty" name:"UnpaidAmount"`

	// Timestamp
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Data date
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`
}

type MNPTypeDefine struct {
	// Specifies the mini program category name.
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// Mini program category value.
	TypeValue []*string `json:"TypeValue,omitnil,omitempty" name:"TypeValue"`

	// Category ID.
	TypeId *int64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// Creation time
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Creator
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// Indicates whether it is a system category.
	IsSystem *bool `json:"IsSystem,omitnil,omitempty" name:"IsSystem"`
}

// Predefined struct for user
type ModifyApplicationConfigRequestParams struct {
	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp configuration ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Package name: corresponds to packageName on Android and bundleId on iOS
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// Superapp URL
	AppURL *string `json:"AppURL,omitnil,omitempty" name:"AppURL"`
}

type ModifyApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// Superapp ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Superapp configuration ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Package name: corresponds to packageName on Android and bundleId on iOS
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// Superapp URL
	AppURL *string `json:"AppURL,omitnil,omitempty" name:"AppURL"`
}

func (r *ModifyApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PlatformId")
	delete(f, "Id")
	delete(f, "AppKey")
	delete(f, "AppURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationConfigResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationConfigResponseParams `json:"Response"`
}

func (r *ModifyApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Application introduction
	Intro *string `json:"Intro,omitnil,omitempty" name:"Intro"`

	// Icon
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Android app package name
	//
	// Deprecated: AndroidAppKey is deprecated.
	AndroidAppKey *string `json:"AndroidAppKey,omitnil,omitempty" name:"AndroidAppKey"`

	// iOS App bundleId
	//
	// Deprecated: IosAppKey is deprecated.
	IosAppKey *string `json:"IosAppKey,omitnil,omitempty" name:"IosAppKey"`

	// Remarks
	//
	// Deprecated: Remark is deprecated.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Scheme
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`
}

type ModifyApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Application introduction
	Intro *string `json:"Intro,omitnil,omitempty" name:"Intro"`

	// Icon
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Android app package name
	AndroidAppKey *string `json:"AndroidAppKey,omitnil,omitempty" name:"AndroidAppKey"`

	// iOS App bundleId
	IosAppKey *string `json:"IosAppKey,omitnil,omitempty" name:"IosAppKey"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Scheme
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`
}

func (r *ModifyApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationName")
	delete(f, "Intro")
	delete(f, "Logo")
	delete(f, "PlatformId")
	delete(f, "AndroidAppKey")
	delete(f, "IosAppKey")
	delete(f, "Remark")
	delete(f, "Scheme")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationResponseParams `json:"Response"`
}

func (r *ModifyApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalDomainRequestParams struct {
	// Domain ID
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// Domain name
	DomainUrl *string `json:"DomainUrl,omitnil,omitempty" name:"DomainUrl"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type ModifyGlobalDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain ID
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// Domain name
	DomainUrl *string `json:"DomainUrl,omitnil,omitempty" name:"DomainUrl"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *ModifyGlobalDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "DomainUrl")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGlobalDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalDomainResponseParams struct {
	// Response data
	Data *GlobalDomainModifyRespResp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGlobalDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGlobalDomainResponseParams `json:"Response"`
}

func (r *ModifyGlobalDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMNPDomainRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Domain list
	Domain []*CreateDomainParam `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type ModifyMNPDomainRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Domain list
	Domain []*CreateDomainParam `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *ModifyMNPDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMNPDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMNPDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMNPDomainResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMNPDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMNPDomainResponseParams `json:"Response"`
}

func (r *ModifyMNPDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMNPDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMNPRequestParams struct {
	// Mini program type
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Mini program introduction
	MNPIntro *string `json:"MNPIntro,omitnil,omitempty" name:"MNPIntro"`

	// Mini program description
	MNPDesc *string `json:"MNPDesc,omitnil,omitempty" name:"MNPDesc"`

	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Mini program icon
	MNPIcon *string `json:"MNPIcon,omitnil,omitempty" name:"MNPIcon"`
}

type ModifyMNPRequest struct {
	*tchttp.BaseRequest
	
	// Mini program type
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Mini program name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Mini program introduction
	MNPIntro *string `json:"MNPIntro,omitnil,omitempty" name:"MNPIntro"`

	// Mini program description
	MNPDesc *string `json:"MNPDesc,omitnil,omitempty" name:"MNPDesc"`

	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Mini program icon
	MNPIcon *string `json:"MNPIcon,omitnil,omitempty" name:"MNPIcon"`
}

func (r *ModifyMNPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMNPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPType")
	delete(f, "MNPName")
	delete(f, "MNPIntro")
	delete(f, "MNPDesc")
	delete(f, "MNPId")
	delete(f, "PlatformId")
	delete(f, "MNPIcon")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMNPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMNPResponseParams struct {
	// Response data
	Data *ResourceIdInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMNPResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMNPResponseParams `json:"Response"`
}

func (r *ModifyMNPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMNPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTeamMemberRequestParams struct {
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Role ID
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type ModifyTeamMemberRequest struct {
	*tchttp.BaseRequest
	
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Role ID
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *ModifyTeamMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTeamMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "UserId")
	delete(f, "RoleId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTeamMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTeamMemberResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTeamMemberResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTeamMemberResponseParams `json:"Response"`
}

func (r *ModifyTeamMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTeamMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTeamRequestParams struct {
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team administrator
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`
}

type ModifyTeamRequest struct {
	*tchttp.BaseRequest
	
	// Team ID
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// Team name
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Team administrator
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`
}

func (r *ModifyTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTeamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "TeamName")
	delete(f, "PlatformId")
	delete(f, "AdminUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTeamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTeamResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTeamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTeamResponseParams `json:"Response"`
}

func (r *ModifyTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTeamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRequestParams struct {
	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Account type 2 - Platform admin 3 - Member
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Account type 2 - Platform admin 3 - Member
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *ModifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserName")
	delete(f, "AccountType")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserResponseParams `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Overview struct {
	// Superapps
	AppNum *int64 `json:"AppNum,omitnil,omitempty" name:"AppNum"`

	// Teams
	CorpNum *int64 `json:"CorpNum,omitnil,omitempty" name:"CorpNum"`

	// Refresh time
	FlushTime *string `json:"FlushTime,omitnil,omitempty" name:"FlushTime"`

	// Mini programs
	MiniAppNum *int64 `json:"MiniAppNum,omitnil,omitempty" name:"MiniAppNum"`

	// Mini games
	MiniGameNum *int64 `json:"MiniGameNum,omitnil,omitempty" name:"MiniGameNum"`

	// Mini game visits
	MiniGameVisitNum *int64 `json:"MiniGameVisitNum,omitnil,omitempty" name:"MiniGameVisitNum"`

	// Mini program updates
	UpdateNum *int64 `json:"UpdateNum,omitnil,omitempty" name:"UpdateNum"`

	// Mini program visits
	VisitNum *int64 `json:"VisitNum,omitnil,omitempty" name:"VisitNum"`
}

type OverviewDetail struct {
	// Time
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// Name
	MNPName *string `json:"MNPName,omitnil,omitempty" name:"MNPName"`

	// Category
	MNPType *string `json:"MNPType,omitnil,omitempty" name:"MNPType"`

	// Estimated revenue
	EstimatedEarnings *string `json:"EstimatedEarnings,omitnil,omitempty" name:"EstimatedEarnings"`

	// Requests
	RequestsNumber *int64 `json:"RequestsNumber,omitnil,omitempty" name:"RequestsNumber"`

	// Impressions
	Impressions *int64 `json:"Impressions,omitnil,omitempty" name:"Impressions"`

	// Effective Cost Per Mille
	ECPM *string `json:"ECPM,omitnil,omitempty" name:"ECPM"`

	// Taps
	ClicksNumber *int64 `json:"ClicksNumber,omitnil,omitempty" name:"ClicksNumber"`
}

type PaymentActiveRetention struct {
	// Day 1 active retention of paying users
	OneDayRetentionUsers *int64 `json:"OneDayRetentionUsers,omitnil,omitempty" name:"OneDayRetentionUsers"`

	// Day 2 active retention of paying users
	TwoDayRetentionUsers *int64 `json:"TwoDayRetentionUsers,omitnil,omitempty" name:"TwoDayRetentionUsers"`

	// Day 3 active retention of paying users
	ThreeDayRetentionUsers *int64 `json:"ThreeDayRetentionUsers,omitnil,omitempty" name:"ThreeDayRetentionUsers"`

	// Day 4 active retention of paying users
	FourDayRetentionUsers *int64 `json:"FourDayRetentionUsers,omitnil,omitempty" name:"FourDayRetentionUsers"`

	// Day 5 active retention of paying users
	FiveDayRetentionUsers *int64 `json:"FiveDayRetentionUsers,omitnil,omitempty" name:"FiveDayRetentionUsers"`

	// Day 6 active retention of paying users
	SixDayRetentionUsers *int64 `json:"SixDayRetentionUsers,omitnil,omitempty" name:"SixDayRetentionUsers"`

	// Day 7 active retention of paying users
	SevenDayRetentionUsers *int64 `json:"SevenDayRetentionUsers,omitnil,omitempty" name:"SevenDayRetentionUsers"`

	// Day 14 active retention of paying users
	FourteenDayRetentionUsers *int64 `json:"FourteenDayRetentionUsers,omitnil,omitempty" name:"FourteenDayRetentionUsers"`

	// Day 15 active retention of paying users
	FifteenDayRetentionUsers *int64 `json:"FifteenDayRetentionUsers,omitnil,omitempty" name:"FifteenDayRetentionUsers"`

	// Day 30 active retention of paying users
	ThirtyDayRetentionUsers *int64 `json:"ThirtyDayRetentionUsers,omitnil,omitempty" name:"ThirtyDayRetentionUsers"`

	// Number of paying users
	PaymentUserNum *int64 `json:"PaymentUserNum,omitnil,omitempty" name:"PaymentUserNum"`

	// Data time in YYYYMMDD format
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`
}

type PaymentDetail struct {
	// Date in YYYYMMDD format
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// Number of mini programs involved in the order
	OrderMNPNum *int64 `json:"OrderMNPNum,omitnil,omitempty" name:"OrderMNPNum"`

	// Total orders
	OrderNum *int64 `json:"OrderNum,omitnil,omitempty" name:"OrderNum"`

	// Paid orders
	OrderPaidNum *int64 `json:"OrderPaidNum,omitnil,omitempty" name:"OrderPaidNum"`

	// Total refunded orders
	OrderRefundNum *int64 `json:"OrderRefundNum,omitnil,omitempty" name:"OrderRefundNum"`

	// Unpaid orders
	OrderUnpaidNum *int64 `json:"OrderUnpaidNum,omitnil,omitempty" name:"OrderUnpaidNum"`

	// Number of users placing orders (openid)
	OrderUserNum *int64 `json:"OrderUserNum,omitnil,omitempty" name:"OrderUserNum"`

	// Amount paid
	PaidAmount *string `json:"PaidAmount,omitnil,omitempty" name:"PaidAmount"`

	// Amount refunded
	RefundAmount *string `json:"RefundAmount,omitnil,omitempty" name:"RefundAmount"`

	// Total order amount
	TotalAmount *string `json:"TotalAmount,omitnil,omitempty" name:"TotalAmount"`

	// Unpaid amount
	UnpaidAmount *string `json:"UnpaidAmount,omitnil,omitempty" name:"UnpaidAmount"`

	// Data update timestamp
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type PresetResp struct {
	// RSA encryption public key.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

// Predefined struct for user
type ProcessMNPApprovalRequestParams struct {
	// Approval ID
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Approval details
	ApprovalItems []*ApprovalItem `json:"ApprovalItems,omitnil,omitempty" name:"ApprovalItems"`
}

type ProcessMNPApprovalRequest struct {
	*tchttp.BaseRequest
	
	// Approval ID
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Approval details
	ApprovalItems []*ApprovalItem `json:"ApprovalItems,omitnil,omitempty" name:"ApprovalItems"`
}

func (r *ProcessMNPApprovalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMNPApprovalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApprovalNo")
	delete(f, "PlatformId")
	delete(f, "ApprovalItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMNPApprovalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMNPApprovalResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProcessMNPApprovalResponse struct {
	*tchttp.BaseResponse
	Response *ProcessMNPApprovalResponseParams `json:"Response"`
}

func (r *ProcessMNPApprovalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMNPApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMNPSensitiveAPIPermissionApprovalRequestParams struct {
	// Approval ID
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Approval status. 20: Rejected; 30: Approved
	ApprovalStatus *int64 `json:"ApprovalStatus,omitnil,omitempty" name:"ApprovalStatus"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Approval notes
	ApprovalNote *string `json:"ApprovalNote,omitnil,omitempty" name:"ApprovalNote"`
}

type ProcessMNPSensitiveAPIPermissionApprovalRequest struct {
	*tchttp.BaseRequest
	
	// Approval ID
	ApprovalNo *string `json:"ApprovalNo,omitnil,omitempty" name:"ApprovalNo"`

	// Approval status. 20: Rejected; 30: Approved
	ApprovalStatus *int64 `json:"ApprovalStatus,omitnil,omitempty" name:"ApprovalStatus"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Approval notes
	ApprovalNote *string `json:"ApprovalNote,omitnil,omitempty" name:"ApprovalNote"`
}

func (r *ProcessMNPSensitiveAPIPermissionApprovalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMNPSensitiveAPIPermissionApprovalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApprovalNo")
	delete(f, "ApprovalStatus")
	delete(f, "PlatformId")
	delete(f, "ApprovalNote")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMNPSensitiveAPIPermissionApprovalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMNPSensitiveAPIPermissionApprovalResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProcessMNPSensitiveAPIPermissionApprovalResponse struct {
	*tchttp.BaseResponse
	Response *ProcessMNPSensitiveAPIPermissionApprovalResponseParams `json:"Response"`
}

func (r *ProcessMNPSensitiveAPIPermissionApprovalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMNPSensitiveAPIPermissionApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOnlineVersionResp struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Version number
	MNPVersion *string `json:"MNPVersion,omitnil,omitempty" name:"MNPVersion"`

	// Version ID.
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Version remarks.
	MNPVersionNote *string `json:"MNPVersionNote,omitnil,omitempty" name:"MNPVersionNote"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ReleaseMNPVersionRequestParams struct {
	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type ReleaseMNPVersionRequest struct {
	*tchttp.BaseRequest
	
	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *ReleaseMNPVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseMNPVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPVersionId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseMNPVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseMNPVersionResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleaseMNPVersionResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseMNPVersionResponseParams `json:"Response"`
}

func (r *ReleaseMNPVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseMNPVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveMNPRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

type RemoveMNPRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

func (r *RemoveMNPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveMNPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "PlatformId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveMNPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveMNPResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveMNPResponse struct {
	*tchttp.BaseResponse
	Response *RemoveMNPResponseParams `json:"Response"`
}

func (r *RemoveMNPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveMNPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportDataResult struct {
	// Base64-encoded result data
	DataResult *string `json:"DataResult,omitnil,omitempty" name:"DataResult"`

	// Executed SQL
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExecSql *string `json:"ExecSql,omitnil,omitempty" name:"ExecSql"`

	// Execution time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExecTime *int64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// Query index ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`
}

type ResourceIdInfo struct {
	// Specifies the resource ID returned by the business.
	// 0: no trial version available.
	// A trial version is currently available and uploaded by the current user.
	// 2: a trial version is currently available and uploaded by another user.
	ResourceId *int64 `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type ResourceIdStringInfo struct {
	// The ID of the resource returned
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type RetentionData struct {
	// Number of active users
	ActiveUserCount *int64 `json:"ActiveUserCount,omitnil,omitempty" name:"ActiveUserCount"`

	// Day 1 retention of active users
	OneDayRetentionUsers *int64 `json:"OneDayRetentionUsers,omitnil,omitempty" name:"OneDayRetentionUsers"`

	// Day 2 retention of active users
	TwoDayRetentionUsers *int64 `json:"TwoDayRetentionUsers,omitnil,omitempty" name:"TwoDayRetentionUsers"`

	// Day 3 retention of active users
	ThreeDayRetentionUsers *int64 `json:"ThreeDayRetentionUsers,omitnil,omitempty" name:"ThreeDayRetentionUsers"`

	// Day 4 retention of active users
	FourDayRetentionUsers *int64 `json:"FourDayRetentionUsers,omitnil,omitempty" name:"FourDayRetentionUsers"`

	// Day 5 retention of active users
	FiveDayRetentionUsers *int64 `json:"FiveDayRetentionUsers,omitnil,omitempty" name:"FiveDayRetentionUsers"`

	// Day 6 retention of active users
	SixDayRetentionUsers *int64 `json:"SixDayRetentionUsers,omitnil,omitempty" name:"SixDayRetentionUsers"`

	// Day 7 retention of active users
	SevenDayRetentionUsers *int64 `json:"SevenDayRetentionUsers,omitnil,omitempty" name:"SevenDayRetentionUsers"`

	// Day 14 retention of active users
	FourteenDayRetentionUsers *int64 `json:"FourteenDayRetentionUsers,omitnil,omitempty" name:"FourteenDayRetentionUsers"`

	// Day 30 retention of active users
	ThirtyDayRetentionUsers *int64 `json:"ThirtyDayRetentionUsers,omitnil,omitempty" name:"ThirtyDayRetentionUsers"`

	// Number of new users
	NewUserCount *int64 `json:"NewUserCount,omitnil,omitempty" name:"NewUserCount"`

	// Day 1 retention of new users
	OneDayRetentionNewUsers *int64 `json:"OneDayRetentionNewUsers,omitnil,omitempty" name:"OneDayRetentionNewUsers"`

	// Day 2 retention of new users
	TwoDayRetentionNewUsers *int64 `json:"TwoDayRetentionNewUsers,omitnil,omitempty" name:"TwoDayRetentionNewUsers"`

	// Day 3 retention of new users
	ThreeDayRetentionNewUsers *int64 `json:"ThreeDayRetentionNewUsers,omitnil,omitempty" name:"ThreeDayRetentionNewUsers"`

	// Day 4 retention of new users
	FourDayRetentionNewUsers *int64 `json:"FourDayRetentionNewUsers,omitnil,omitempty" name:"FourDayRetentionNewUsers"`

	// Day 5 retention of new users
	FiveDayRetentionNewUsers *int64 `json:"FiveDayRetentionNewUsers,omitnil,omitempty" name:"FiveDayRetentionNewUsers"`

	// Day 6 retention of new users
	SixDayRetentionNewUsers *int64 `json:"SixDayRetentionNewUsers,omitnil,omitempty" name:"SixDayRetentionNewUsers"`

	// Day 7 retention of new users
	SevenDayRetentionNewUsers *int64 `json:"SevenDayRetentionNewUsers,omitnil,omitempty" name:"SevenDayRetentionNewUsers"`

	// Day 14 retention of new users
	FourteenDayRetentionNewUsers *int64 `json:"FourteenDayRetentionNewUsers,omitnil,omitempty" name:"FourteenDayRetentionNewUsers"`

	// Day 30 retention of new users
	ThirtyDayRetentionNewUsers *int64 `json:"ThirtyDayRetentionNewUsers,omitnil,omitempty" name:"ThirtyDayRetentionNewUsers"`

	// Data time in YYYYMMDD format
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`
}

// Predefined struct for user
type RollbackMNPVersionRequestParams struct {
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Mini program version number
	MNPVersion *string `json:"MNPVersion,omitnil,omitempty" name:"MNPVersion"`
}

type RollbackMNPVersionRequest struct {
	*tchttp.BaseRequest
	
	// Mini program ID
	MNPId *string `json:"MNPId,omitnil,omitempty" name:"MNPId"`

	// Mini program version ID
	MNPVersionId *int64 `json:"MNPVersionId,omitnil,omitempty" name:"MNPVersionId"`

	// Platform ID
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`

	// Mini program version number
	MNPVersion *string `json:"MNPVersion,omitnil,omitempty" name:"MNPVersion"`
}

func (r *RollbackMNPVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackMNPVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MNPId")
	delete(f, "MNPVersionId")
	delete(f, "PlatformId")
	delete(f, "MNPVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackMNPVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackMNPVersionResponseParams struct {
	// Response data
	Data *BooleanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackMNPVersionResponse struct {
	*tchttp.BaseResponse
	Response *RollbackMNPVersionResponseParams `json:"Response"`
}

func (r *RollbackMNPVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackMNPVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StringData struct {
	// string type response data.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

type UploadFileTempSecret struct {
	// Bucket
	// Note: This field may return null, indicating that no valid values can be obtained.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Destination of upload
	// Note: This field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Temporary secret ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TempSecretId *string `json:"TempSecretId,omitnil,omitempty" name:"TempSecretId"`

	// Temporary secret key
	// Note: This field may return null, indicating that no valid values can be obtained.
	TempSecretKey *string `json:"TempSecretKey,omitnil,omitempty" name:"TempSecretKey"`

	// Token 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// Whether to enable global acceleration. Valid values: 0 (no), 1 (yes)
	AccelerateEnable *int64 `json:"AccelerateEnable,omitnil,omitempty" name:"AccelerateEnable"`
}

type VisitData struct {
	// Number of visits
	VisitCount *int64 `json:"VisitCount,omitnil,omitempty" name:"VisitCount"`

	// Average pages per device - visit_page_count / active_device_num
	AvgDeviceVisitDeep *string `json:"AvgDeviceVisitDeep,omitnil,omitempty" name:"AvgDeviceVisitDeep"`

	// Pages per visit - visit_page_count / miniapp_open_num
	AvgCountVisitDeep *string `json:"AvgCountVisitDeep,omitnil,omitempty" name:"AvgCountVisitDeep"`

	// Average visit duration - miniapp_total_duration / visit_page_count
	AvgPageVisitDuration *string `json:"AvgPageVisitDuration,omitnil,omitempty" name:"AvgPageVisitDuration"`

	// Average visit duration per session
	// miniapp_total_duration/miniapp_open_num
	AvgCountVisitDuration *string `json:"AvgCountVisitDuration,omitnil,omitempty" name:"AvgCountVisitDuration"`

	// Refresh time in YYYYMMDD format
	DataTime *int64 `json:"DataTime,omitnil,omitempty" name:"DataTime"`
}