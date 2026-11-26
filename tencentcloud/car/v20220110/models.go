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

package v20220110

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type ApplicationBaseInfo struct {
	// Application window usage mode (ApplicationDesktop: Captures the entire desktop; ApplicationWindow: Captures only the application window).
	WindowUseType *string `json:"WindowUseType,omitnil,omitempty" name:"WindowUseType"`

	// Application window name.
	WindowName *string `json:"WindowName,omitnil,omitempty" name:"WindowName"`

	// Application window class name.
	WindowClassName *string `json:"WindowClassName,omitnil,omitempty" name:"WindowClassName"`

	// Application window capture mode (HOOK: default mode; DDA_CUT: compatible mode).
	WindowCaptureMode *string `json:"WindowCaptureMode,omitnil,omitempty" name:"WindowCaptureMode"`
}

type ApplicationConcurrentPackage struct {
	// <p>Concurrency pack ID</p>
	ConcurrentId *string `json:"ConcurrentId,omitnil,omitempty" name:"ConcurrentId"`
}

type ApplicationProject struct {
	// <p>Project ID.</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Project name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Project description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Concurrency type required for project operation.<br>S1: Small-scale cloud application rendering concurrency<br>M1: Medium cloud application rendering concurrency<br>L1: Large cloud application rendering concurrency<br>L2: Large cloud application rendering concurrency<br>XL2: Ultra-large cloud application rendering concurrency<br>MM1_HD: Performance Cloud ARM concurrency - high-definition<br>MM1_FHD: Performance Cloud ARM concurrency - ultra-high-definition</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Cloud Application ID</p>
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// <p>Pre-launch</p>
	IsPreload *bool `json:"IsPreload,omitnil,omitempty" name:"IsPreload"`

	// <p>Configured concurrency</p>
	Amount *int64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// <p>Number of concurrencies in use</p>
	Using *int64 `json:"Using,omitnil,omitempty" name:"Using"`

	// <p>Application status: NoConcurrent - no concurrency package configured, Online - launched, other statuses are cloud application status: applicationCreating - creating, applicationCreateFail - creation failed, applicationDeleting - deleting, applicationNoConfigured - startup parameter not configured</p>
	ApplicationStatus *string `json:"ApplicationStatus,omitnil,omitempty" name:"ApplicationStatus"`

	// <p>Application startup parameter</p>
	ApplicationParams *string `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Application name</p>
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// <p>Resolution, format as "width x height", for example 1920x1080</p>
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// <p>Project type.<br>SHARED: Shared by all applications<br>EXCLUSIVE (default value): Dedicated for one application</p>
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// <p>Purpose.<br>EXPERIENCE: Experience</p>
	Purpose *string `json:"Purpose,omitnil,omitempty" name:"Purpose"`

	// <p>Application distribution area (Standard zone: ap-chinese-mainland: Chinese mainland, na-north-america: North America, eu-frankfurt: Frankfurt, ap-mumbai: Mumbai, ap-tokyo: Tokyo, ap-seoul: Seoul, ap-singapore: Singapore, ap-bangkok: Bangkok, ap-hongkong: Hong Kong (China). Fusion region: me-middle-east-fusion: Middle East, na-north-america-fusion: North America, sa-south-america-fusion: South America, ap-tokyo-fusion: Tokyo, ap-seoul-fusion: Seoul, eu-frankfurt-fusion: Frankfurt, ap-singapore-fusion: Singapore, ap-hongkong-fusion: Hong Kong (China))</p>
	ApplicationRegions []*string `json:"ApplicationRegions,omitnil,omitempty" name:"ApplicationRegions"`

	// <p>Concurrent region (Standard zone: ap-chinese-mainland: Chinese mainland, na-north-america: North America, eu-frankfurt: Frankfurt, ap-mumbai: Mumbai, ap-tokyo: Tokyo, ap-seoul: Seoul, ap-singapore: Singapore, ap-bangkok: Bangkok, ap-hongkong: Hong Kong (China), Fusion region: me-middle-east-fusion: Middle East, na-north-america-fusion: North America, sa-south-america-fusion: South America, ap-tokyo-fusion: Tokyo, ap-seoul-fusion: Seoul, eu-frankfurt-fusion: Frankfurt, ap-singapore-fusion: Singapore, ap-hongkong-fusion: Hong Kong (China))</p>
	ConcurrentRegions []*string `json:"ConcurrentRegions,omitnil,omitempty" name:"ConcurrentRegions"`

	// <p>Project category.<br>DESKTOP: On PC, default value<br>MOBILE: Mobile terminal</p>
	ProjectCategory *string `json:"ProjectCategory,omitnil,omitempty" name:"ProjectCategory"`
}

// Predefined struct for user
type ApplyConcurrentRequestParams struct {
	// Unique user ID, which is customized by you and is not parsed by CAR. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Public IP address of the user's client, which is used for nearby scheduling.
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Application version ID. If the application of the current version is requested, you do not need to fill in this field. If the application of the other versions is requested, you need to specify the version through this field.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// Application ID, which is used only by the multi-application project to specify applications. For a single-application project, this parameter is ignored, and the application bound to the project will be used.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type ApplyConcurrentRequest struct {
	*tchttp.BaseRequest
	
	// Unique user ID, which is customized by you and is not parsed by CAR. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Public IP address of the user's client, which is used for nearby scheduling.
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Application version ID. If the application of the current version is requested, you do not need to fill in this field. If the application of the other versions is requested, you need to specify the version through this field.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// Application ID, which is used only by the multi-application project to specify applications. For a single-application project, this parameter is ignored, and the application bound to the project will be used.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

func (r *ApplyConcurrentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConcurrentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserIp")
	delete(f, "ProjectId")
	delete(f, "ApplicationVersionId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyConcurrentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyConcurrentResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyConcurrentResponse struct {
	*tchttp.BaseResponse
	Response *ApplyConcurrentResponseParams `json:"Response"`
}

func (r *ApplyConcurrentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConcurrentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackgroundImage struct {
	// ID of the COS file.
	COSFileId *string `json:"COSFileId,omitnil,omitempty" name:"COSFileId"`

	// Download URL.
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Creation Time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type BindConcurrentPackagesToProjectRequestParams struct {
	// Concurrency pack ID list.
	ConcurrentIds []*string `json:"ConcurrentIds,omitnil,omitempty" name:"ConcurrentIds"`

	// Cloud application project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BindConcurrentPackagesToProjectRequest struct {
	*tchttp.BaseRequest
	
	// Concurrency pack ID list.
	ConcurrentIds []*string `json:"ConcurrentIds,omitnil,omitempty" name:"ConcurrentIds"`

	// Cloud application project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *BindConcurrentPackagesToProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindConcurrentPackagesToProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConcurrentIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindConcurrentPackagesToProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindConcurrentPackagesToProjectResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindConcurrentPackagesToProjectResponse struct {
	*tchttp.BaseResponse
	Response *BindConcurrentPackagesToProjectResponseParams `json:"Response"`
}

func (r *BindConcurrentPackagesToProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindConcurrentPackagesToProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProjectRequestParams struct {
	// Project name, which is user-defined.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Bound application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Concurrency type required for project operation.S1: concurrency for rendering small cloud applications.M1: concurrency for rendering medium cloud applications.L1: concurrency for rendering large cloud applications.L2: concurrency for rendering large cloud applications.XL2: concurrency for rendering extra large cloud applications.MM1_HD: concurrency for performance-based cloud ARM (HD).MM1_FHD: concurrency for performance-based cloud ARM (FHD).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether to enable warm-up. The default value is false.
	IsPreload *bool `json:"IsPreload,omitnil,omitempty" name:"IsPreload"`

	// Application startup parameters.
	ApplicationParams *string `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// Resolution, in the format of widthxheight, such as 1920x1080.
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Project type.SHARED: shared by all applications.EXCLUSIVE (default value): dedicated for one application.
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// Frame rate.
	FPS *int64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// Waiting time for application pre-launch.
	PreloadDuration *string `json:"PreloadDuration,omitnil,omitempty" name:"PreloadDuration"`

	// Waiting time for reconnection.
	ReconnectTimeout *string `json:"ReconnectTimeout,omitnil,omitempty" name:"ReconnectTimeout"`

	// Minimum bitrate, in Mbps.
	MinBitrate *int64 `json:"MinBitrate,omitnil,omitempty" name:"MinBitrate"`

	// Maximum bitrate, in Mbps.
	MaxBitrate *int64 `json:"MaxBitrate,omitnil,omitempty" name:"MaxBitrate"`

	// Upstream audio options.DisableMixIntoStreamPush: not mixing upstream audio in streaming.
	UpstreamAudioOption *string `json:"UpstreamAudioOption,omitnil,omitempty" name:"UpstreamAudioOption"`

	// Video encoding configuration.
	VideoEncodeConfig *VideoEncodeConfig `json:"VideoEncodeConfig,omitnil,omitempty" name:"VideoEncodeConfig"`

	// Upper limit of the XR application resolution.If the project concurrency type is L or L2, the upper limit is 5000; if the project concurrency type is XL2, the upper limit is 6000.
	XRMaxWidth *uint64 `json:"XRMaxWidth,omitnil,omitempty" name:"XRMaxWidth"`

	// ID of the background image COS file.
	BackgroundImageCOSFileId *string `json:"BackgroundImageCOSFileId,omitnil,omitempty" name:"BackgroundImageCOSFileId"`

	// Project category.DESKTOP: desktop (default value).MOBILE: mobile.
	ProjectCategory *string `json:"ProjectCategory,omitnil,omitempty" name:"ProjectCategory"`

	// Disabled code list.
	DisableVideoCodecs []*string `json:"DisableVideoCodecs,omitnil,omitempty" name:"DisableVideoCodecs"`
}

type CreateApplicationProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project name, which is user-defined.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Bound application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Concurrency type required for project operation.S1: concurrency for rendering small cloud applications.M1: concurrency for rendering medium cloud applications.L1: concurrency for rendering large cloud applications.L2: concurrency for rendering large cloud applications.XL2: concurrency for rendering extra large cloud applications.MM1_HD: concurrency for performance-based cloud ARM (HD).MM1_FHD: concurrency for performance-based cloud ARM (FHD).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether to enable warm-up. The default value is false.
	IsPreload *bool `json:"IsPreload,omitnil,omitempty" name:"IsPreload"`

	// Application startup parameters.
	ApplicationParams *string `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// Resolution, in the format of widthxheight, such as 1920x1080.
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Project type.SHARED: shared by all applications.EXCLUSIVE (default value): dedicated for one application.
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// Frame rate.
	FPS *int64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// Waiting time for application pre-launch.
	PreloadDuration *string `json:"PreloadDuration,omitnil,omitempty" name:"PreloadDuration"`

	// Waiting time for reconnection.
	ReconnectTimeout *string `json:"ReconnectTimeout,omitnil,omitempty" name:"ReconnectTimeout"`

	// Minimum bitrate, in Mbps.
	MinBitrate *int64 `json:"MinBitrate,omitnil,omitempty" name:"MinBitrate"`

	// Maximum bitrate, in Mbps.
	MaxBitrate *int64 `json:"MaxBitrate,omitnil,omitempty" name:"MaxBitrate"`

	// Upstream audio options.DisableMixIntoStreamPush: not mixing upstream audio in streaming.
	UpstreamAudioOption *string `json:"UpstreamAudioOption,omitnil,omitempty" name:"UpstreamAudioOption"`

	// Video encoding configuration.
	VideoEncodeConfig *VideoEncodeConfig `json:"VideoEncodeConfig,omitnil,omitempty" name:"VideoEncodeConfig"`

	// Upper limit of the XR application resolution.If the project concurrency type is L or L2, the upper limit is 5000; if the project concurrency type is XL2, the upper limit is 6000.
	XRMaxWidth *uint64 `json:"XRMaxWidth,omitnil,omitempty" name:"XRMaxWidth"`

	// ID of the background image COS file.
	BackgroundImageCOSFileId *string `json:"BackgroundImageCOSFileId,omitnil,omitempty" name:"BackgroundImageCOSFileId"`

	// Project category.DESKTOP: desktop (default value).MOBILE: mobile.
	ProjectCategory *string `json:"ProjectCategory,omitnil,omitempty" name:"ProjectCategory"`

	// Disabled code list.
	DisableVideoCodecs []*string `json:"DisableVideoCodecs,omitnil,omitempty" name:"DisableVideoCodecs"`
}

func (r *CreateApplicationProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ApplicationId")
	delete(f, "Type")
	delete(f, "IsPreload")
	delete(f, "ApplicationParams")
	delete(f, "Resolution")
	delete(f, "ProjectType")
	delete(f, "FPS")
	delete(f, "PreloadDuration")
	delete(f, "ReconnectTimeout")
	delete(f, "MinBitrate")
	delete(f, "MaxBitrate")
	delete(f, "UpstreamAudioOption")
	delete(f, "VideoEncodeConfig")
	delete(f, "XRMaxWidth")
	delete(f, "BackgroundImageCOSFileId")
	delete(f, "ProjectCategory")
	delete(f, "DisableVideoCodecs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProjectResponseParams struct {
	// Generated project ID.Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationProjectResponseParams `json:"Response"`
}

func (r *CreateApplicationProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationRequestParams struct {
	// Application name.
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Application type (Application3D: cloud 3D; ApplicationXR: cloud XR; ApplicationAPK: cloud APK).
	ApplicationType *string `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application name.
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Application type (Application3D: cloud 3D; ApplicationXR: cloud XR; ApplicationAPK: cloud APK).
	ApplicationType *string `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`
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
	delete(f, "ApplicationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationResponseParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

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

// Predefined struct for user
type CreateApplicationSnapshotRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application download address (if the version is created by file upload, this parameter is an empty string).
	ApplicationDownloadUrl *string `json:"ApplicationDownloadUrl,omitnil,omitempty" name:"ApplicationDownloadUrl"`
}

type CreateApplicationSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application download address (if the version is created by file upload, this parameter is an empty string).
	ApplicationDownloadUrl *string `json:"ApplicationDownloadUrl,omitnil,omitempty" name:"ApplicationDownloadUrl"`
}

func (r *CreateApplicationSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationDownloadUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationSnapshotResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationSnapshotResponseParams `json:"Response"`
}

func (r *CreateApplicationSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationVersionRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application file name (desktop applications should be files in zip/rar/7z format, and mobile applications should be files in apk format).
	ApplicationFileName *string `json:"ApplicationFileName,omitnil,omitempty" name:"ApplicationFileName"`

	// Region for application version distribution.
	ApplicationVersionRegions []*string `json:"ApplicationVersionRegions,omitnil,omitempty" name:"ApplicationVersionRegions"`

	// Application update method.
	ApplicationVersionUpdateMode *string `json:"ApplicationVersionUpdateMode,omitnil,omitempty" name:"ApplicationVersionUpdateMode"`
}

type CreateApplicationVersionRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application file name (desktop applications should be files in zip/rar/7z format, and mobile applications should be files in apk format).
	ApplicationFileName *string `json:"ApplicationFileName,omitnil,omitempty" name:"ApplicationFileName"`

	// Region for application version distribution.
	ApplicationVersionRegions []*string `json:"ApplicationVersionRegions,omitnil,omitempty" name:"ApplicationVersionRegions"`

	// Application update method.
	ApplicationVersionUpdateMode *string `json:"ApplicationVersionUpdateMode,omitnil,omitempty" name:"ApplicationVersionUpdateMode"`
}

func (r *CreateApplicationVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationFileName")
	delete(f, "ApplicationVersionRegions")
	delete(f, "ApplicationVersionUpdateMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationVersionResponseParams struct {
	// Application version data (new).
	Version *UserApplicationVersion `json:"Version,omitnil,omitempty" name:"Version"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationVersionResponseParams `json:"Response"`
}

func (r *CreateApplicationVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMCPTokenRequestParams struct {
	// <p>Token expiration time</p><p>Value range: [0, 2592000]</p><p>Unit: seconds</p><p>Default value: 86400</p><p>If this parameter is not specified, the default timeout is 1 day. The expiration time is the current time + the timeout period.</p>
	ExpireSeconds *int64 `json:"ExpireSeconds,omitnil,omitempty" name:"ExpireSeconds"`

	// <p>Token name (comment)</p><p>Format: Only a combination of Chinese characters, English letters, digits, and hyphens (&quot;-&quot;) is supported.</p><p>Maximum length: 16 characters</p>
	TokenName *string `json:"TokenName,omitnil,omitempty" name:"TokenName"`

	// <p>User Id</p><p>Format: Only a combination of letters (A-Za-z), digits (0-9), and underscores (_) is supported.</p><p>This UserId is always used to operate cloud rendering concurrent resources. Maximum length: 64 characters</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type CreateMCPTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>Token expiration time</p><p>Value range: [0, 2592000]</p><p>Unit: seconds</p><p>Default value: 86400</p><p>If this parameter is not specified, the default timeout is 1 day. The expiration time is the current time + the timeout period.</p>
	ExpireSeconds *int64 `json:"ExpireSeconds,omitnil,omitempty" name:"ExpireSeconds"`

	// <p>Token name (comment)</p><p>Format: Only a combination of Chinese characters, English letters, digits, and hyphens (&quot;-&quot;) is supported.</p><p>Maximum length: 16 characters</p>
	TokenName *string `json:"TokenName,omitnil,omitempty" name:"TokenName"`

	// <p>User Id</p><p>Format: Only a combination of letters (A-Za-z), digits (0-9), and underscores (_) is supported.</p><p>This UserId is always used to operate cloud rendering concurrent resources. Maximum length: 64 characters</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *CreateMCPTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMCPTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExpireSeconds")
	delete(f, "TokenName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMCPTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMCPTokenResponseParams struct {
	// <p>TokenId</p>
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// <p>MCP Token</p><p>Keep the current value safe. It is required for subsequent communication with the MCP Server. Once lost, it cannot be viewed again.</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMCPTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateMCPTokenResponseParams `json:"Response"`
}

func (r *CreateMCPTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMCPTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionRequestParams struct {
	// Unique user ID, which is customized by you and is not parsed by CAR. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Public IP address of the user's client, which is used for nearby scheduling.
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// Client-side session information, which is obtained from the SDK. If `RunMode` is `RunWithoutClient`, this parameter can be empty.
	ClientSession *string `json:"ClientSession,omitnil,omitempty" name:"ClientSession"`

	// On-cloud running mode.RunWithoutClient: Keeps the application running on the cloud even when there are no client connections.Empty string (default): Keeps the application running on the cloud only when there are client connections.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Application startup parameters.This parameter is effective for multi-application projects.
	// This parameter is effective for single-application projects with prelaunch disabled.This parameter is ineffective for single-application projects with prelaunch enabled.
	// Note: When this parameter is effective, it will be appended to the startup parameters of application or project configuration in the console.
	// For example, for a single-application project with prelaunch disabled, if its startup parameter `bar` is `0` for project configuration in the console and the `ApplicationParameters` parameter `foo` is `1`, the actual application startup parameters will be `bar=0 and foo=1`.
	ApplicationParameters *string `json:"ApplicationParameters,omitnil,omitempty" name:"ApplicationParameters"`

	// [Multi-person Interaction] Homeowner's user ID, which is required in multi-person interaction mode.
	// If the user is the homeowner, HostUserID must be the same as UserID.
	// If the user is not the homeowner, HostUserID must be the homeowner's HostUserID.
	HostUserId *string `json:"HostUserId,omitnil,omitempty" name:"HostUserId"`

	// [Multi-person Interaction] Role.
	// Player: a user who can operate the application via keyboard, mouse, etc.
	// Viewer: a user who can only watch the video in the room but cannot operate the application.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

type CreateSessionRequest struct {
	*tchttp.BaseRequest
	
	// Unique user ID, which is customized by you and is not parsed by CAR. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Public IP address of the user's client, which is used for nearby scheduling.
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// Client-side session information, which is obtained from the SDK. If `RunMode` is `RunWithoutClient`, this parameter can be empty.
	ClientSession *string `json:"ClientSession,omitnil,omitempty" name:"ClientSession"`

	// On-cloud running mode.RunWithoutClient: Keeps the application running on the cloud even when there are no client connections.Empty string (default): Keeps the application running on the cloud only when there are client connections.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Application startup parameters.This parameter is effective for multi-application projects.
	// This parameter is effective for single-application projects with prelaunch disabled.This parameter is ineffective for single-application projects with prelaunch enabled.
	// Note: When this parameter is effective, it will be appended to the startup parameters of application or project configuration in the console.
	// For example, for a single-application project with prelaunch disabled, if its startup parameter `bar` is `0` for project configuration in the console and the `ApplicationParameters` parameter `foo` is `1`, the actual application startup parameters will be `bar=0 and foo=1`.
	ApplicationParameters *string `json:"ApplicationParameters,omitnil,omitempty" name:"ApplicationParameters"`

	// [Multi-person Interaction] Homeowner's user ID, which is required in multi-person interaction mode.
	// If the user is the homeowner, HostUserID must be the same as UserID.
	// If the user is not the homeowner, HostUserID must be the homeowner's HostUserID.
	HostUserId *string `json:"HostUserId,omitnil,omitempty" name:"HostUserId"`

	// [Multi-person Interaction] Role.
	// Player: a user who can operate the application via keyboard, mouse, etc.
	// Viewer: a user who can only watch the video in the room but cannot operate the application.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

func (r *CreateSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserIp")
	delete(f, "ClientSession")
	delete(f, "RunMode")
	delete(f, "ApplicationParameters")
	delete(f, "HostUserId")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionResponseParams struct {
	// Server-side session information, which is returned to the SDK.
	ServerSession *string `json:"ServerSession,omitnil,omitempty" name:"ServerSession"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateSessionResponseParams `json:"Response"`
}

func (r *CreateSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProjectsRequestParams struct {
	// ID list of cloud application projects.
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`
}

type DeleteApplicationProjectsRequest struct {
	*tchttp.BaseRequest
	
	// ID list of cloud application projects.
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`
}

func (r *DeleteApplicationProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProjectsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteApplicationProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationProjectsResponseParams `json:"Response"`
}

func (r *DeleteApplicationProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type DeleteApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationResponseParams struct {
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
type DeleteApplicationVersionRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application version ID.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`
}

type DeleteApplicationVersionRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application version ID.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`
}

func (r *DeleteApplicationVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationVersionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteApplicationVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationVersionResponseParams `json:"Response"`
}

func (r *DeleteApplicationVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationFileInfoRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application file path list.
	FilePathList []*string `json:"FilePathList,omitnil,omitempty" name:"FilePathList"`
}

type DescribeApplicationFileInfoRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application file path list.
	FilePathList []*string `json:"FilePathList,omitnil,omitempty" name:"FilePathList"`
}

func (r *DescribeApplicationFileInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationFileInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "FilePathList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationFileInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationFileInfoResponseParams struct {
	// Application file data list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileInfoList []*UserApplicationFileInfo `json:"FileInfoList,omitnil,omitempty" name:"FileInfoList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationFileInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationFileInfoResponseParams `json:"Response"`
}

func (r *DescribeApplicationFileInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationFileInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationListRequestParams struct {
	// Application list offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Application quantity limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Application category (DESKTOP: desktop; MOBILE: mobile).
	ApplicationCategory *string `json:"ApplicationCategory,omitnil,omitempty" name:"ApplicationCategory"`
}

type DescribeApplicationListRequest struct {
	*tchttp.BaseRequest
	
	// Application list offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Application quantity limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Application category (DESKTOP: desktop; MOBILE: mobile).
	ApplicationCategory *string `json:"ApplicationCategory,omitnil,omitempty" name:"ApplicationCategory"`
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
	delete(f, "Filters")
	delete(f, "ApplicationCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationListResponseParams struct {
	// Application list information.
	UserApplicationList []*UserApplicationInfo `json:"UserApplicationList,omitnil,omitempty" name:"UserApplicationList"`

	// Total number of applications.
	ApplicationTotal *int64 `json:"ApplicationTotal,omitnil,omitempty" name:"ApplicationTotal"`

	// Mobile application list information.
	UserMobileApplicationList []*UserMobileApplicationInfo `json:"UserMobileApplicationList,omitnil,omitempty" name:"UserMobileApplicationList"`

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
type DescribeApplicationPathListRequestParams struct {
	// Cloud application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Cloud application version ID.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`
}

type DescribeApplicationPathListRequest struct {
	*tchttp.BaseRequest
	
	// Cloud application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Cloud application version ID.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`
}

func (r *DescribeApplicationPathListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationPathListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationPathListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationPathListResponseParams struct {
	// Application .exe file path list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PathList []*string `json:"PathList,omitnil,omitempty" name:"PathList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationPathListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationPathListResponseParams `json:"Response"`
}

func (r *DescribeApplicationPathListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationPathListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProjectAdvancedConfigRequestParams struct {
	// Application project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeApplicationProjectAdvancedConfigRequest struct {
	*tchttp.BaseRequest
	
	// Application project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeApplicationProjectAdvancedConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProjectAdvancedConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProjectAdvancedConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProjectAdvancedConfigResponseParams struct {
	// Application startup parameters.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationParams *string `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// Resolution, in the format of widthxheight, such as 1920x1080.Note: This field may return null, indicating that no valid values can be obtained.
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Frame rate. Valid values: 0, 30, 60.Note: This field may return null, indicating that no valid values can be obtained.
	FPS *int64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// Minimum bitrate, in Mbps.Note: This field may return null, indicating that no valid values can be obtained.
	MinBitrate *int64 `json:"MinBitrate,omitnil,omitempty" name:"MinBitrate"`

	// Maximum bitrate, in Mbps.Note: This field may return null, indicating that no valid values can be obtained.
	MaxBitrate *int64 `json:"MaxBitrate,omitnil,omitempty" name:"MaxBitrate"`

	// Waiting time for application pre-launch.Note: This field may return null, indicating that no valid values can be obtained.
	PreloadDuration *string `json:"PreloadDuration,omitnil,omitempty" name:"PreloadDuration"`

	// Waiting time for reconnection.Note: This field may return null, indicating that no valid values can be obtained.
	ReconnectTimeout *string `json:"ReconnectTimeout,omitnil,omitempty" name:"ReconnectTimeout"`

	// Upstream audio options.DisableMixIntoStreamPush: not mixing upstream audio in streaming.Note: This field may return null, indicating that no valid values can be obtained.
	UpstreamAudioOption *string `json:"UpstreamAudioOption,omitnil,omitempty" name:"UpstreamAudioOption"`

	// Video encoding configuration.Note: This field may return null, indicating that no valid values can be obtained.
	VideoEncodeConfig *VideoEncodeConfig `json:"VideoEncodeConfig,omitnil,omitempty" name:"VideoEncodeConfig"`

	// Upper limit of the XR application resolution.If the project concurrency type is L or L2, the upper limit is 5000; if the project concurrency type is XL2, the upper limit is 6000.Note: This field may return null, indicating that no valid values can be obtained.
	XRMaxWidth *uint64 `json:"XRMaxWidth,omitnil,omitempty" name:"XRMaxWidth"`

	// Background image information.Note: This field may return null, indicating that no valid values can be obtained.
	BackgroundImage *BackgroundImage `json:"BackgroundImage,omitnil,omitempty" name:"BackgroundImage"`

	// Disabled code list.Note: This field may return null, indicating that no valid values can be obtained.
	DisableVideoCodecs []*string `json:"DisableVideoCodecs,omitnil,omitempty" name:"DisableVideoCodecs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationProjectAdvancedConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationProjectAdvancedConfigResponseParams `json:"Response"`
}

func (r *DescribeApplicationProjectAdvancedConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProjectAdvancedConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProjectsRequestParams struct {
	// Subscript.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Project category.DESKTOP: desktop (default value).MOBILE: mobile.
	ProjectCategory *string `json:"ProjectCategory,omitnil,omitempty" name:"ProjectCategory"`
}

type DescribeApplicationProjectsRequest struct {
	*tchttp.BaseRequest
	
	// Subscript.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Project category.DESKTOP: desktop (default value).MOBILE: mobile.
	ProjectCategory *string `json:"ProjectCategory,omitnil,omitempty" name:"ProjectCategory"`
}

func (r *DescribeApplicationProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "ProjectCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProjectsResponseParams struct {
	// Project list.Note: This field may return null, indicating that no valid values can be obtained.
	Projects []*ApplicationProject `json:"Projects,omitnil,omitempty" name:"Projects"`

	// Total number.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationProjectsResponseParams `json:"Response"`
}

func (r *DescribeApplicationProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationStatusRequestParams struct {
	// Application ID list.
	ApplicationIdList []*string `json:"ApplicationIdList,omitnil,omitempty" name:"ApplicationIdList"`
}

type DescribeApplicationStatusRequest struct {
	*tchttp.BaseRequest
	
	// Application ID list.
	ApplicationIdList []*string `json:"ApplicationIdList,omitnil,omitempty" name:"ApplicationIdList"`
}

func (r *DescribeApplicationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationStatusResponseParams struct {
	// Application status list.
	StatusList []*UserApplicationStatus `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationStatusResponseParams `json:"Response"`
}

func (r *DescribeApplicationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationVersionRequestParams struct {
	// Application version ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type DescribeApplicationVersionRequest struct {
	*tchttp.BaseRequest
	
	// Application version ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationVersionResponseParams struct {
	// List of application versions.
	Versions []*UserApplicationVersion `json:"Versions,omitnil,omitempty" name:"Versions"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationVersionResponseParams `json:"Response"`
}

func (r *DescribeApplicationVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrentCountRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeConcurrentCountRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeConcurrentCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrentCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrentCountResponseParams struct {
	// Total number of concurrencies.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The number of concurrent executions, including all non-idle concurrent executions such as those in prelaunch, connected, waiting for reconnection, and to be cleaned up or recovered. Therefore, refreshing projects or disconnecting user connections with concurrency packages will affect this value.
	Running *uint64 `json:"Running,omitnil,omitempty" name:"Running"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrentCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrentCountResponseParams `json:"Response"`
}

func (r *DescribeConcurrentCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrentPackagesRequestParams struct {
	// Subscript.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter List
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeConcurrentPackagesRequest struct {
	*tchttp.BaseRequest
	
	// Subscript.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter List
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeConcurrentPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrentPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrentPackagesResponseParams struct {
	// Total number.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Concurrency pack list.Note: This field may return null, indicating that no valid values can be obtained.
	ConcurrentPackages []*ApplicationConcurrentPackage `json:"ConcurrentPackages,omitnil,omitempty" name:"ConcurrentPackages"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrentPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrentPackagesResponseParams `json:"Response"`
}

func (r *DescribeConcurrentPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrentSummaryRequestParams struct {

}

type DescribeConcurrentSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeConcurrentSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrentSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrentSummaryResponseParams struct {
	// Total number of prepaid (monthly subscription) concurrencies.
	PrepaidConcurrentTotal *uint64 `json:"PrepaidConcurrentTotal,omitnil,omitempty" name:"PrepaidConcurrentTotal"`

	// Remaining duration of an hourly package.
	HourlyRemainDuration *string `json:"HourlyRemainDuration,omitnil,omitempty" name:"HourlyRemainDuration"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrentSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrentSummaryResponseParams `json:"Response"`
}

func (r *DescribeConcurrentSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosCredentialRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application file name (the file must be a compressed package with a zip/rar/7z file name extension).
	ApplicationFileName *string `json:"ApplicationFileName,omitnil,omitempty" name:"ApplicationFileName"`
}

type DescribeCosCredentialRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application file name (the file must be a compressed package with a zip/rar/7z file name extension).
	ApplicationFileName *string `json:"ApplicationFileName,omitnil,omitempty" name:"ApplicationFileName"`
}

func (r *DescribeCosCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosCredentialResponseParams struct {
	// Cos SecretID
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// Cos SecretKey
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// Cos SessionToken
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// Cos Bucket
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// Cos Region
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// COS operation path.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Start time of the COS key.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Expiration time of the COS key.
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosCredentialResponseParams `json:"Response"`
}

func (r *DescribeCosCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMCPTokensRequestParams struct {
	// <p>Offset.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page</p><p>Value range: [0, 200]</p><p>Default value: 20</p><p>If 0 is passed in, the default value is used.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter list</p><p>Only filtering by TokenId and TokenName is supported. The number of Filters cannot exceed 10, the number of Values in each Filter cannot exceed 20, the length of a single Value cannot exceed 64 characters, and duplicate Filter names are not allowed.</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeMCPTokensRequest struct {
	*tchttp.BaseRequest
	
	// <p>Offset.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page</p><p>Value range: [0, 200]</p><p>Default value: 20</p><p>If 0 is passed in, the default value is used.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter list</p><p>Only filtering by TokenId and TokenName is supported. The number of Filters cannot exceed 10, the number of Values in each Filter cannot exceed 20, the length of a single Value cannot exceed 64 characters, and duplicate Filter names are not allowed.</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeMCPTokensRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMCPTokensRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMCPTokensRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMCPTokensResponseParams struct {
	// <p>Token list</p>
	Tokens []*MCPTokenInfo `json:"Tokens,omitnil,omitempty" name:"Tokens"`

	// <p>Number of Tokens</p><p>Unit: count</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMCPTokensResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMCPTokensResponseParams `json:"Response"`
}

func (r *DescribeMCPTokensResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMCPTokensResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySessionRequestParams struct {
	// Unique user ID, which is customized by you and is not parsed by CAR. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DestroySessionRequest struct {
	*tchttp.BaseRequest
	
	// Unique user ID, which is customized by you and is not parsed by CAR. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DestroySessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroySessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySessionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroySessionResponse struct {
	*tchttp.BaseResponse
	Response *DestroySessionResponseParams `json:"Response"`
}

func (r *DestroySessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Filter field name (ApplicationId: application ID; ApplicationName: application name; ApplicationRunStatus: running status; ApplicationUpdateStatus: update status).
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value set (When the filter name is ApplicationRunStatus, the values can be [ApplicationDeleting: application deletion in progress; ApplicationCreateFail: application creation failed; ApplicationCreating: application creation in progress; ApplicationRunning: normal running; ApplicationNoConfigured: main execution program path not configured]. When the filter name is ApplicationUpdateStatus, the values can be [ApplicationUpdateCreating: version creation in progress; ApplicationUpdateCreateFail: version creation failed; ApplicationUpdateNoReleased: version to be released; ApplicationUpdateReleased: version release completed; ApplicationUpdateNormal: none]).
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type MCPTokenInfo struct {
	// <p>TokenId</p><p>Used to manage Token data. It can be used to extend the validity period of the corresponding Token.</p>
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// <p>Token name (comment)</p>
	TokenName *string `json:"TokenName,omitnil,omitempty" name:"TokenName"`

	// <p>User Id</p><p>This UserId is always used to operate cloud rendering concurrent resources. (If this parameter is empty, no restriction is applied.)</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Permission scope</p><p>Default value: [&quot;*&quot;]</p><p>Used to specify operation permissions. By default, all permissions are granted.</p>
	Scopes []*string `json:"Scopes,omitnil,omitempty" name:"Scopes"`

	// <p>Token status</p><p>Valid values:</p><ul><li>ACTIVE: Normal status. The token has not been revoked, but it may not be within its validity period.</li><li>REVOKED: The token has been revoked and is permanently invalid.</li></ul><p>Default value: ACTIVE</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Token creation time</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>Token expiration time</p>
	ExpiresTime *string `json:"ExpiresTime,omitnil,omitempty" name:"ExpiresTime"`
}

// Predefined struct for user
type ModifyApplicationBaseInfoRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application program execution path.
	ApplicationExePath *string `json:"ApplicationExePath,omitnil,omitempty" name:"ApplicationExePath"`

	// Application process list.
	ApplicationInterList *string `json:"ApplicationInterList,omitnil,omitempty" name:"ApplicationInterList"`

	// Application basic data.
	ApplicationBaseInfo *ApplicationBaseInfo `json:"ApplicationBaseInfo,omitnil,omitempty" name:"ApplicationBaseInfo"`

	// Application startup parameters.
	ApplicationParams *string `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// Application name.
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Application repository information list.
	ApplicationStores []*UserApplicationStore `json:"ApplicationStores,omitnil,omitempty" name:"ApplicationStores"`
}

type ModifyApplicationBaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application program execution path.
	ApplicationExePath *string `json:"ApplicationExePath,omitnil,omitempty" name:"ApplicationExePath"`

	// Application process list.
	ApplicationInterList *string `json:"ApplicationInterList,omitnil,omitempty" name:"ApplicationInterList"`

	// Application basic data.
	ApplicationBaseInfo *ApplicationBaseInfo `json:"ApplicationBaseInfo,omitnil,omitempty" name:"ApplicationBaseInfo"`

	// Application startup parameters.
	ApplicationParams *string `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// Application name.
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Application repository information list.
	ApplicationStores []*UserApplicationStore `json:"ApplicationStores,omitnil,omitempty" name:"ApplicationStores"`
}

func (r *ModifyApplicationBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationExePath")
	delete(f, "ApplicationInterList")
	delete(f, "ApplicationBaseInfo")
	delete(f, "ApplicationParams")
	delete(f, "ApplicationName")
	delete(f, "ApplicationStores")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationBaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationBaseInfoResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApplicationBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationBaseInfoResponseParams `json:"Response"`
}

func (r *ModifyApplicationBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProjectRequestParams struct {
	// Project ID returned by cloud.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Concurrency type required for project operation.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether to Enable Pre-launch.
	IsPreload *bool `json:"IsPreload,omitnil,omitempty" name:"IsPreload"`

	// Application startup parameters.
	ApplicationParams *string `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// Cloud application project description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Resolution, in the format of widthxheight, such as 1920x1080.
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Frame rate.
	FPS *int64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// Waiting time for pre-launch.
	PreloadDuration *string `json:"PreloadDuration,omitnil,omitempty" name:"PreloadDuration"`

	// Waiting time for reconnection.
	ReconnectTimeout *string `json:"ReconnectTimeout,omitnil,omitempty" name:"ReconnectTimeout"`

	// Minimum bitrate, in Mbps.
	MinBitrate *int64 `json:"MinBitrate,omitnil,omitempty" name:"MinBitrate"`

	// Maximum bitrate, in Mbps.
	MaxBitrate *int64 `json:"MaxBitrate,omitnil,omitempty" name:"MaxBitrate"`

	// Upstream audio options.DisableMixIntoStreamPush: not mixing upstream audio in streaming.
	UpstreamAudioOption *string `json:"UpstreamAudioOption,omitnil,omitempty" name:"UpstreamAudioOption"`

	// Video encoding configuration.
	VideoEncodeConfig *VideoEncodeConfig `json:"VideoEncodeConfig,omitnil,omitempty" name:"VideoEncodeConfig"`

	// Upper limit of the XR application resolution.If the project concurrency type is L or L2, the upper limit is 5000; if the project concurrency type is XL2, the upper limit is 6000.
	XRMaxWidth *uint64 `json:"XRMaxWidth,omitnil,omitempty" name:"XRMaxWidth"`

	// ID of the background image COS file.
	BackgroundImageCOSFileId *string `json:"BackgroundImageCOSFileId,omitnil,omitempty" name:"BackgroundImageCOSFileId"`

	// Disabled code list.
	DisableVideoCodecs []*string `json:"DisableVideoCodecs,omitnil,omitempty" name:"DisableVideoCodecs"`
}

type ModifyApplicationProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project ID returned by cloud.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Concurrency type required for project operation.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether to Enable Pre-launch.
	IsPreload *bool `json:"IsPreload,omitnil,omitempty" name:"IsPreload"`

	// Application startup parameters.
	ApplicationParams *string `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// Cloud application project description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Resolution, in the format of widthxheight, such as 1920x1080.
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Frame rate.
	FPS *int64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// Waiting time for pre-launch.
	PreloadDuration *string `json:"PreloadDuration,omitnil,omitempty" name:"PreloadDuration"`

	// Waiting time for reconnection.
	ReconnectTimeout *string `json:"ReconnectTimeout,omitnil,omitempty" name:"ReconnectTimeout"`

	// Minimum bitrate, in Mbps.
	MinBitrate *int64 `json:"MinBitrate,omitnil,omitempty" name:"MinBitrate"`

	// Maximum bitrate, in Mbps.
	MaxBitrate *int64 `json:"MaxBitrate,omitnil,omitempty" name:"MaxBitrate"`

	// Upstream audio options.DisableMixIntoStreamPush: not mixing upstream audio in streaming.
	UpstreamAudioOption *string `json:"UpstreamAudioOption,omitnil,omitempty" name:"UpstreamAudioOption"`

	// Video encoding configuration.
	VideoEncodeConfig *VideoEncodeConfig `json:"VideoEncodeConfig,omitnil,omitempty" name:"VideoEncodeConfig"`

	// Upper limit of the XR application resolution.If the project concurrency type is L or L2, the upper limit is 5000; if the project concurrency type is XL2, the upper limit is 6000.
	XRMaxWidth *uint64 `json:"XRMaxWidth,omitnil,omitempty" name:"XRMaxWidth"`

	// ID of the background image COS file.
	BackgroundImageCOSFileId *string `json:"BackgroundImageCOSFileId,omitnil,omitempty" name:"BackgroundImageCOSFileId"`

	// Disabled code list.
	DisableVideoCodecs []*string `json:"DisableVideoCodecs,omitnil,omitempty" name:"DisableVideoCodecs"`
}

func (r *ModifyApplicationProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "IsPreload")
	delete(f, "ApplicationParams")
	delete(f, "Description")
	delete(f, "Resolution")
	delete(f, "FPS")
	delete(f, "PreloadDuration")
	delete(f, "ReconnectTimeout")
	delete(f, "MinBitrate")
	delete(f, "MaxBitrate")
	delete(f, "UpstreamAudioOption")
	delete(f, "VideoEncodeConfig")
	delete(f, "XRMaxWidth")
	delete(f, "BackgroundImageCOSFileId")
	delete(f, "DisableVideoCodecs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProjectResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApplicationProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProjectResponseParams `json:"Response"`
}

func (r *ModifyApplicationProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationVersionRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application version ID.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// Application version name.
	ApplicationVersionName *string `json:"ApplicationVersionName,omitnil,omitempty" name:"ApplicationVersionName"`
}

type ModifyApplicationVersionRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application version ID.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// Application version name.
	ApplicationVersionName *string `json:"ApplicationVersionName,omitnil,omitempty" name:"ApplicationVersionName"`
}

func (r *ModifyApplicationVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationVersionId")
	delete(f, "ApplicationVersionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationVersionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApplicationVersionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationVersionResponseParams `json:"Response"`
}

func (r *ModifyApplicationVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConcurrentPackageRequestParams struct {
	// Concurrency pack ID.
	ConcurrentId *string `json:"ConcurrentId,omitnil,omitempty" name:"ConcurrentId"`

	// Concurrency pack name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ModifyConcurrentPackageRequest struct {
	*tchttp.BaseRequest
	
	// Concurrency pack ID.
	ConcurrentId *string `json:"ConcurrentId,omitnil,omitempty" name:"ConcurrentId"`

	// Concurrency pack name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *ModifyConcurrentPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConcurrentPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConcurrentId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConcurrentPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConcurrentPackageResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConcurrentPackageResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConcurrentPackageResponseParams `json:"Response"`
}

func (r *ModifyConcurrentPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConcurrentPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMobileApplicationInfoRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application name.
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`
}

type ModifyMobileApplicationInfoRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application name.
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`
}

func (r *ModifyMobileApplicationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMobileApplicationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMobileApplicationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMobileApplicationInfoResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMobileApplicationInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMobileApplicationInfoResponseParams `json:"Response"`
}

func (r *ModifyMobileApplicationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMobileApplicationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewMCPTokenRequestParams struct {
	// <p>TokenId</p>
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// <p>Token expiration time</p><p>Value range: [0, 2592000]</p><p>Default value: 86400</p><p>Unit: seconds. If not specified, the default is 1 day. The new expiration time is the current time plus ExpireSeconds.</p>
	ExpireSeconds *int64 `json:"ExpireSeconds,omitnil,omitempty" name:"ExpireSeconds"`
}

type RenewMCPTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>TokenId</p>
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// <p>Token expiration time</p><p>Value range: [0, 2592000]</p><p>Default value: 86400</p><p>Unit: seconds. If not specified, the default is 1 day. The new expiration time is the current time plus ExpireSeconds.</p>
	ExpireSeconds *int64 `json:"ExpireSeconds,omitnil,omitempty" name:"ExpireSeconds"`
}

func (r *RenewMCPTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewMCPTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TokenId")
	delete(f, "ExpireSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewMCPTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewMCPTokenResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewMCPTokenResponse struct {
	*tchttp.BaseResponse
	Response *RenewMCPTokenResponseParams `json:"Response"`
}

func (r *RenewMCPTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewMCPTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetConcurrentPackagesRequestParams struct {
	// Concurrency pack ID array.
	ConcurrentPackageIds []*string `json:"ConcurrentPackageIds,omitnil,omitempty" name:"ConcurrentPackageIds"`
}

type ResetConcurrentPackagesRequest struct {
	*tchttp.BaseRequest
	
	// Concurrency pack ID array.
	ConcurrentPackageIds []*string `json:"ConcurrentPackageIds,omitnil,omitempty" name:"ConcurrentPackageIds"`
}

func (r *ResetConcurrentPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetConcurrentPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConcurrentPackageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetConcurrentPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetConcurrentPackagesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetConcurrentPackagesResponse struct {
	*tchttp.BaseResponse
	Response *ResetConcurrentPackagesResponseParams `json:"Response"`
}

func (r *ResetConcurrentPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetConcurrentPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetApplicationVersionOnlineRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application version ID.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`
}

type SetApplicationVersionOnlineRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application version ID.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`
}

func (r *SetApplicationVersionOnlineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetApplicationVersionOnlineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetApplicationVersionOnlineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetApplicationVersionOnlineResponseParams struct {
	// List of application versions.
	Versions []*UserApplicationVersion `json:"Versions,omitnil,omitempty" name:"Versions"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetApplicationVersionOnlineResponse struct {
	*tchttp.BaseResponse
	Response *SetApplicationVersionOnlineResponseParams `json:"Response"`
}

func (r *SetApplicationVersionOnlineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetApplicationVersionOnlineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamRequestParams struct {
	// Unique user ID, which is customized by you and is not parsed by CAR. It will be used as the `StreamId` for streaming. For example, if the bound streaming domain is **abc.livepush.myqcloud.com**, the streaming address will be **rtmp://abc.livepush.myqcloud.com/live/UserId?txSecret=xxx&txTime=xxx**.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Streaming parameter, which is a custom parameter carried during streaming.
	PublishStreamArgs *string `json:"PublishStreamArgs,omitnil,omitempty" name:"PublishStreamArgs"`
}

type StartPublishStreamRequest struct {
	*tchttp.BaseRequest
	
	// Unique user ID, which is customized by you and is not parsed by CAR. It will be used as the `StreamId` for streaming. For example, if the bound streaming domain is **abc.livepush.myqcloud.com**, the streaming address will be **rtmp://abc.livepush.myqcloud.com/live/UserId?txSecret=xxx&txTime=xxx**.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Streaming parameter, which is a custom parameter carried during streaming.
	PublishStreamArgs *string `json:"PublishStreamArgs,omitnil,omitempty" name:"PublishStreamArgs"`
}

func (r *StartPublishStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PublishStreamArgs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartPublishStreamResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishStreamResponseParams `json:"Response"`
}

func (r *StartPublishStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamWithURLRequestParams struct {
	// Unique user ID, which is customized by you and is not parsed by CAR.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Streaming address. Only RTMP is supported for streaming currently.
	PublishStreamURL *string `json:"PublishStreamURL,omitnil,omitempty" name:"PublishStreamURL"`
}

type StartPublishStreamWithURLRequest struct {
	*tchttp.BaseRequest
	
	// Unique user ID, which is customized by you and is not parsed by CAR.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Streaming address. Only RTMP is supported for streaming currently.
	PublishStreamURL *string `json:"PublishStreamURL,omitnil,omitempty" name:"PublishStreamURL"`
}

func (r *StartPublishStreamWithURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamWithURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PublishStreamURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishStreamWithURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamWithURLResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartPublishStreamWithURLResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishStreamWithURLResponseParams `json:"Response"`
}

func (r *StartPublishStreamWithURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamWithURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishStreamRequestParams struct {
	// Unique user ID, which is customized by you and is not parsed by CAR. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type StopPublishStreamRequest struct {
	*tchttp.BaseRequest
	
	// Unique user ID, which is customized by you and is not parsed by CAR. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *StopPublishStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopPublishStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishStreamResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopPublishStreamResponse struct {
	*tchttp.BaseResponse
	Response *StopPublishStreamResponseParams `json:"Response"`
}

func (r *StopPublishStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// <p>Tag key</p>
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// <p>Tag value</p>
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UnbindConcurrentPackagesFromProjectRequestParams struct {
	// Concurrency pack ID list.
	ConcurrentIds []*string `json:"ConcurrentIds,omitnil,omitempty" name:"ConcurrentIds"`

	// Cloud application project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type UnbindConcurrentPackagesFromProjectRequest struct {
	*tchttp.BaseRequest
	
	// Concurrency pack ID list.
	ConcurrentIds []*string `json:"ConcurrentIds,omitnil,omitempty" name:"ConcurrentIds"`

	// Cloud application project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *UnbindConcurrentPackagesFromProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindConcurrentPackagesFromProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConcurrentIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindConcurrentPackagesFromProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindConcurrentPackagesFromProjectResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindConcurrentPackagesFromProjectResponse struct {
	*tchttp.BaseResponse
	Response *UnbindConcurrentPackagesFromProjectResponseParams `json:"Response"`
}

func (r *UnbindConcurrentPackagesFromProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindConcurrentPackagesFromProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserApplicationFileInfo struct {
	// Application file path.
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// File status. NO_EXIST: not EXIST; EXIST: exist.
	FileState *string `json:"FileState,omitnil,omitempty" name:"FileState"`
}

type UserApplicationInfo struct {
	// <p>Application Id</p>
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// <p>Application name</p>
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// <p>Application type (Cloud 3D: Application3D, Cloud XR: ApplicationXR, Cloud Web: ApplicationWeb)</p>
	ApplicationType *string `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// <p>Program execution path</p>
	ApplicationExePath *string `json:"ApplicationExePath,omitnil,omitempty" name:"ApplicationExePath"`

	// <p>Application process list</p>
	ApplicationInterList *string `json:"ApplicationInterList,omitnil,omitempty" name:"ApplicationInterList"`

	// <p>Application startup parameter</p>
	ApplicationParams *string `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// <p>Application running status (ApplicationDeleting: application deletion in progress; ApplicationCreateFail: application creation failed; ApplicationCreating: application creation in progress; ApplicationRunning: normal running; ApplicationNoConfigured: main execution program path not configured).</p>
	ApplicationRunStatus *string `json:"ApplicationRunStatus,omitnil,omitempty" name:"ApplicationRunStatus"`

	// <p>Application update status (ApplicationUpdateCreating: version creation in progress; ApplicationUpdateCreateFail: version creation failed; ApplicationUpdateNoReleased: version to be released; ApplicationUpdateReleased: version release completed; ApplicationUpdateNormal: none)</p>
	ApplicationUpdateStatus *string `json:"ApplicationUpdateStatus,omitnil,omitempty" name:"ApplicationUpdateStatus"`

	// <p>Application creation time</p>
	ApplicationCreateTime *string `json:"ApplicationCreateTime,omitnil,omitempty" name:"ApplicationCreateTime"`

	// <p>Application version list</p>
	ApplicationVersions []*UserApplicationVersion `json:"ApplicationVersions,omitnil,omitempty" name:"ApplicationVersions"`

	// <p>Application basic data</p>
	ApplicationBaseInfo *ApplicationBaseInfo `json:"ApplicationBaseInfo,omitnil,omitempty" name:"ApplicationBaseInfo"`

	// <p>Application update progress</p>
	ApplicationUpdateProgress *int64 `json:"ApplicationUpdateProgress,omitnil,omitempty" name:"ApplicationUpdateProgress"`

	// <p>Application nature (PUBLIC: public application; PRIVATE: user application)</p>
	ApplicationNature *string `json:"ApplicationNature,omitnil,omitempty" name:"ApplicationNature"`

	// <p>Application repository list</p>
	ApplicationStores []*UserApplicationStore `json:"ApplicationStores,omitnil,omitempty" name:"ApplicationStores"`

	// <p>Application tag</p>
	ApplicationTags []*Tag `json:"ApplicationTags,omitnil,omitempty" name:"ApplicationTags"`
}

type UserApplicationStatus struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application running status (ApplicationDeleting: application deletion in progress; ApplicationCreateFail: application creation failed; ApplicationCreating: application creation in progress; ApplicationRunning: normal running; ApplicationNoConfigured: main execution program path not configured; ApplicationNoPackage: no available package).
	ApplicationRunStatus *string `json:"ApplicationRunStatus,omitnil,omitempty" name:"ApplicationRunStatus"`

	// Application update status (ApplicationUpdateCreating: version creation in progress; ApplicationUpdateCreateFail: version creation failed; ApplicationUpdateNoReleased: version to be released; ApplicationUpdateReleased: version release completed; ApplicationUpdateNormal: none).
	ApplicationUpdateStatus *string `json:"ApplicationUpdateStatus,omitnil,omitempty" name:"ApplicationUpdateStatus"`

	// Application update progress.
	ApplicationUpdateProgress *int64 `json:"ApplicationUpdateProgress,omitnil,omitempty" name:"ApplicationUpdateProgress"`
}

type UserApplicationStore struct {
	// <p>Cos bucket name</p>
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// <p>Cos Bucket Region</p>
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// <p>Repository type. LOG: application logs; ARCHIVE: application archive.</p>
	StoreType *string `json:"StoreType,omitnil,omitempty" name:"StoreType"`

	// <p>Repository status. ON: enabled; OFF: disabled.</p>
	StoreState *string `json:"StoreState,omitnil,omitempty" name:"StoreState"`

	// <p>Repository path</p><p>Primary path data. If StoreType is ARCHIVE and SubStorePaths is configured, the system default value will be used.</p>
	StorePath *string `json:"StorePath,omitnil,omitempty" name:"StorePath"`

	// <p>Subpath list</p><p>Subpath data (multi-path pattern) supports only ARCHIVE StoreType; up to 25 data entries.</p>
	SubStorePaths []*string `json:"SubStorePaths,omitnil,omitempty" name:"SubStorePaths"`
}

type UserApplicationVersion struct {
	// <p>Application version Id</p>
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// <p>Application version size</p>
	ApplicationVersionSize *int64 `json:"ApplicationVersionSize,omitnil,omitempty" name:"ApplicationVersionSize"`

	// <p>Application version status (Uploading: Creating, Creating: Creating, CreateFailed: Failed to create, Deleting: Deleting, Inuse: Current version, Normal: To be released, Usable: Available)</p>
	ApplicationVersionStatus *string `json:"ApplicationVersionStatus,omitnil,omitempty" name:"ApplicationVersionStatus"`

	// <p>Application version name</p>
	ApplicationVersionName *string `json:"ApplicationVersionName,omitnil,omitempty" name:"ApplicationVersionName"`

	// <p>Application version creation time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Application version distribution regions (<br>Standard regions:<br>ap-chinese-mainland: Chinese mainland,<br>na-north-america: North America,<br>eu-frankfurt: Frankfurt,<br>ap-mumbai: Mumbai,<br>ap-tokyo: Tokyo,<br>ap-seoul: Seoul,<br>ap-singapore: Singapore,<br>ap-bangkok: Bangkok,<br>ap-hongkong: Hong Kong (China),<br>Fusion regions:<br>me-middle-east-fusion: Middle East<br>na-north-america-fusion: North America<br>sa-south-america-fusion: South America<br>ap-tokyo-fusion: Tokyo<br>ap-seoul-fusion: Seoul<br>eu-frankfurt-fusion: Frankfurt<br>ap-singapore-fusion: Singapore<br>ap-hongkong-fusion: Hong Kong (China)<br>)</p>
	ApplicationVersionRegions []*string `json:"ApplicationVersionRegions,omitnil,omitempty" name:"ApplicationVersionRegions"`

	// <p>Application version update mode: <br>FULL: Full update, <br>INCREMENT: Incremental update</p>
	ApplicationVersionUpdateMode *string `json:"ApplicationVersionUpdateMode,omitnil,omitempty" name:"ApplicationVersionUpdateMode"`
}

type UserMobileApplicationInfo struct {
	// <p>Application Id</p>
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// <p>Application name</p>
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// <p>App type (Cloud APK: ApplicationAPK)</p>
	ApplicationType *string `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// <p>Application running status (ApplicationRunning: normal running; ApplicationNoPackage: no available package)</p>
	ApplicationRunStatus *string `json:"ApplicationRunStatus,omitnil,omitempty" name:"ApplicationRunStatus"`

	// <p>Application update status (ApplicationUpdateCreating: version creation in progress; ApplicationUpdateCreateFail: version creation failed; ApplicationUpdateNoReleased: version to be released; ApplicationUpdateReleased: version release completed; ApplicationUpdateNormal: none)</p>
	ApplicationUpdateStatus *string `json:"ApplicationUpdateStatus,omitnil,omitempty" name:"ApplicationUpdateStatus"`

	// <p>Application creation time</p>
	ApplicationCreateTime *string `json:"ApplicationCreateTime,omitnil,omitempty" name:"ApplicationCreateTime"`

	// <p>Application version list</p>
	ApplicationVersions []*UserApplicationVersion `json:"ApplicationVersions,omitnil,omitempty" name:"ApplicationVersions"`

	// <p>Application nature (PUBLIC: public application; PRIVATE: user application)</p>
	ApplicationNature *string `json:"ApplicationNature,omitnil,omitempty" name:"ApplicationNature"`

	// <p>Application tag</p>
	ApplicationTags []*Tag `json:"ApplicationTags,omitnil,omitempty" name:"ApplicationTags"`
}

type VideoEncodeConfig struct {
	// Push streaming GOP length, in seconds.
	StreamPushGOPSeconds *uint64 `json:"StreamPushGOPSeconds,omitnil,omitempty" name:"StreamPushGOPSeconds"`
}