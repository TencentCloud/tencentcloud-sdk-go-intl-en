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

package v20190919

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ApplicationItem struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Tag list.
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`
}

// Predefined struct for user
type ApplyTiwTrialRequestParams struct {

}

type ApplyTiwTrialRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ApplyTiwTrialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyTiwTrialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyTiwTrialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyTiwTrialResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyTiwTrialResponse struct {
	*tchttp.BaseResponse
	Response *ApplyTiwTrialResponseParams `json:"Response"`
}

func (r *ApplyTiwTrialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyTiwTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthParam struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// User ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Signature corresponding to the user ID.
	UserSig *string `json:"UserSig,omitempty" name:"UserSig"`
}

type Canvas struct {
	// Width and height of the mixed stream canvas
	LayoutParams *LayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// Background color, which is black by default. Its format is RGB. for example, "#FF0000" for the red color.
	BackgroundColor *string `json:"BackgroundColor,omitempty" name:"BackgroundColor"`
}

type Concat struct {
	// Whether to enable the video splicing feature
	// If the video splicing feature is enabled, the real-time recording service will splice multiple video clips resulting from the pause into one video.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// Download address of the padding image used during video splicing. If it is not specified, a pure black image is used by default.
	Image *string `json:"Image,omitempty" name:"Image"`
}

// Predefined struct for user
type CreateApplicationRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// SKey required for creating an IM application.
	SKey *string `json:"SKey,omitempty" name:"SKey"`

	// TinyId required for creating an IM application.
	TinyId *string `json:"TinyId,omitempty" name:"TinyId"`

	// List of tags to be bound.
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// SKey required for creating an IM application.
	SKey *string `json:"SKey,omitempty" name:"SKey"`

	// TinyId required for creating an IM application.
	TinyId *string `json:"TinyId,omitempty" name:"TinyId"`

	// List of tags to be bound.
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`
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
	delete(f, "SdkAppId")
	delete(f, "AppName")
	delete(f, "SKey")
	delete(f, "TinyId")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationResponseParams struct {
	// AppId of the customer.
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

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
type CreateOfflineRecordRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID corresponding to the recording task.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// Group ID corresponding to the recording task.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Stream mixing parameters.
	// The Custom parameter is not supported for offline recording tasks.
	MixStream *MixStream `json:"MixStream,omitempty" name:"MixStream"`

	// Whiteboard parameters.
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`
}

type CreateOfflineRecordRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID corresponding to the recording task.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// Group ID corresponding to the recording task.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Stream mixing parameters.
	// The Custom parameter is not supported for offline recording tasks.
	MixStream *MixStream `json:"MixStream,omitempty" name:"MixStream"`

	// Whiteboard parameters.
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`
}

func (r *CreateOfflineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOfflineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "GroupId")
	delete(f, "MixStream")
	delete(f, "Whiteboard")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOfflineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOfflineRecordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOfflineRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateOfflineRecordResponseParams `json:"Response"`
}

func (r *CreateOfflineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOfflineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotTaskRequestParams struct {
	// Whiteboard parameters.
	Whiteboard *SnapshotWhiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// `SdkAppId` of the whiteboard application.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Whiteboard room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// Callback URL to which the whiteboard snapshot result is to be sent.
	CallbackURL *string `json:"CallbackURL,omitempty" name:"CallbackURL"`

	// `COS` bucket in which the generated whiteboard snapshot file is to be stored. If you leave this parameter empty, the generated file will be stored in the public bucket and retained for only three days.
	COS *SnapshotCOS `json:"COS,omitempty" name:"COS"`

	// Whiteboard snapshot mode. Default value: `AllMarks`. Valid values:
	// 
	// `AllMarks`: Full mode. In this mode, a snapshot image is generated based on each whiteboard snapshot mark that is added by calling the `addSnapshotMark` API on the client.
	// 
	// `LatestMarksOnly`: Single-page deduplication mode. In this mode, a snapshot image is generated based only on the latest whiteboard snapshot mark that is added by calling the `addSnapshotMark` API on the client if the API is called multiple times for the same whiteboard snapshot.
	// 
	// **Note: The `LatestMarksOnly` mode takes effect only when the `addSnapshotMark` API is called by using Tencent Interactive Whiteboard SDK v2.6.8 or later. Otherwise, even if this parameter is set to `LatestMarksOnly`, the default `AllMarks` mode is used.**
	SnapshotMode *string `json:"SnapshotMode,omitempty" name:"SnapshotMode"`
}

type CreateSnapshotTaskRequest struct {
	*tchttp.BaseRequest
	
	// Whiteboard parameters.
	Whiteboard *SnapshotWhiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// `SdkAppId` of the whiteboard application.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Whiteboard room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// Callback URL to which the whiteboard snapshot result is to be sent.
	CallbackURL *string `json:"CallbackURL,omitempty" name:"CallbackURL"`

	// `COS` bucket in which the generated whiteboard snapshot file is to be stored. If you leave this parameter empty, the generated file will be stored in the public bucket and retained for only three days.
	COS *SnapshotCOS `json:"COS,omitempty" name:"COS"`

	// Whiteboard snapshot mode. Default value: `AllMarks`. Valid values:
	// 
	// `AllMarks`: Full mode. In this mode, a snapshot image is generated based on each whiteboard snapshot mark that is added by calling the `addSnapshotMark` API on the client.
	// 
	// `LatestMarksOnly`: Single-page deduplication mode. In this mode, a snapshot image is generated based only on the latest whiteboard snapshot mark that is added by calling the `addSnapshotMark` API on the client if the API is called multiple times for the same whiteboard snapshot.
	// 
	// **Note: The `LatestMarksOnly` mode takes effect only when the `addSnapshotMark` API is called by using Tencent Interactive Whiteboard SDK v2.6.8 or later. Otherwise, even if this parameter is set to `LatestMarksOnly`, the default `AllMarks` mode is used.**
	SnapshotMode *string `json:"SnapshotMode,omitempty" name:"SnapshotMode"`
}

func (r *CreateSnapshotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Whiteboard")
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "CallbackURL")
	delete(f, "COS")
	delete(f, "SnapshotMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotTaskResponseParams struct {
	// ID of the whiteboard snapshot task. This parameter is returned only if a task is created successfully.
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSnapshotTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotTaskResponseParams `json:"Response"`
}

func (r *CreateSnapshotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeRequestParams struct {
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// URL of the transcoded document after URL encoding. URL encoding converts characters into a format that can be transmitted over the Internet. For example, URL encoding can convert the document URL http://example.com/Test.pdf into http://example.com/%E6%B5%8B%E8%AF%95.pdf. To improve the success rate of URL parsing, use URL encoding.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Whether the PowerPoint file is static. The default value is False.
	// If IsStaticPPT is False, documents with the .ppt or .pptx extension will be dynamically transcoded to HTML5 pages, and documents with other extensions will be statically transcoded to images. If IsStaticPPT is True, documents with any extensions will be statically transcoded to images.
	IsStaticPPT *bool `json:"IsStaticPPT,omitempty" name:"IsStaticPPT"`

	// Note: This parameter is disused. Use the MinScaleResolution parameter to pass in a resolution. For more information, see [CreateTranscode](https://intl.cloud.tencent.com/document/api/1137/40060?from_cn_redirect=1#SDK).
	// 
	// Minimum resolution of the transcoded document. If no value or null is specified for it or the resolution format is invalid, the original document resolution is used.
	// 
	// Example: 1280x720. Note that the character between the numbers is the letter x.
	MinResolution *string `json:"MinResolution,omitempty" name:"MinResolution"`

	// Resolution of the thumbnail generated for the dynamically transcoded PowerPoint file. If no value or null is specified for it or the resolution format is invalid, no thumbnail will be generated. The resolution format is the same as that of MinResolution.
	// 
	// For static transcoding, this parameter does not work.
	ThumbnailResolution *string `json:"ThumbnailResolution,omitempty" name:"ThumbnailResolution"`

	// Compression format of the transcoded file. If no value or null is specified for it or the specified format is invalid, no compression file will be generated. Currently, the following compression formats are supported:
	// 
	// `zip`: generates a .zip compression package.
	// `tar.gz: generates a .tar.gz compression package.
	CompressFileType *string `json:"CompressFileType,omitempty" name:"CompressFileType"`

	// Internal parameter.
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`

	// Document transcoding priority. This parameter takes effect only for PowerPoint dynamic transcoding. Valid values:<br/>
	// - low: Low transcoding priority. The task can transcode at most 500 MB of data or a 2000-page document, with a download timeout no longer than 10 minutes. Due to resource limits, these tasks may have to queue for a long period of time. Consider this before you use this feature.
	// - null: Normal transcoding priority. The task can transcode at most 200 MB of data or a 500-page document, with a download timeout no longer than 2 minutes.
	// <br/>
	// Note: For static transcoding such as PDF transcoding, each task can transcode at most 200 MB of data regardless of the transcoding priority.
	Priority *string `json:"Priority,omitempty" name:"Priority"`

	// Minimum resolution of the transcoded document. If no value or null is specified for it or the resolution format is invalid, the original document resolution is used.
	// Higher resolution brings clearer visual effect, but also means larger size of the transcoded image resources and longer loading time of the transcoded file. Set this parameter appropriately based on your actual scenario.
	// 
	// Example: 1280x720. Note that the character between the numbers is the letter x.
	MinScaleResolution *string `json:"MinScaleResolution,omitempty" name:"MinScaleResolution"`

	// Specifies whether to enable auto handling of unsupported elements. By default, this feature is disabled.
	// 
	// If auto handling is enabled, the following processes are performed:
	// 1. Inkblots: Remove unsupported inkblots, such as those drawn by using WPS.
	// 2. Auto page flip: Clear the auto page clip settings in the PowerPoint file and set the page flip mode to mouse click.
	// 3. Corrupted audio/videos: Remove the references to corrupted audio/videos in the PowerPoint file.
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitempty" name:"AutoHandleUnsupportedElement"`
}

type CreateTranscodeRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// URL of the transcoded document after URL encoding. URL encoding converts characters into a format that can be transmitted over the Internet. For example, URL encoding can convert the document URL http://example.com/Test.pdf into http://example.com/%E6%B5%8B%E8%AF%95.pdf. To improve the success rate of URL parsing, use URL encoding.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Whether the PowerPoint file is static. The default value is False.
	// If IsStaticPPT is False, documents with the .ppt or .pptx extension will be dynamically transcoded to HTML5 pages, and documents with other extensions will be statically transcoded to images. If IsStaticPPT is True, documents with any extensions will be statically transcoded to images.
	IsStaticPPT *bool `json:"IsStaticPPT,omitempty" name:"IsStaticPPT"`

	// Note: This parameter is disused. Use the MinScaleResolution parameter to pass in a resolution. For more information, see [CreateTranscode](https://intl.cloud.tencent.com/document/api/1137/40060?from_cn_redirect=1#SDK).
	// 
	// Minimum resolution of the transcoded document. If no value or null is specified for it or the resolution format is invalid, the original document resolution is used.
	// 
	// Example: 1280x720. Note that the character between the numbers is the letter x.
	MinResolution *string `json:"MinResolution,omitempty" name:"MinResolution"`

	// Resolution of the thumbnail generated for the dynamically transcoded PowerPoint file. If no value or null is specified for it or the resolution format is invalid, no thumbnail will be generated. The resolution format is the same as that of MinResolution.
	// 
	// For static transcoding, this parameter does not work.
	ThumbnailResolution *string `json:"ThumbnailResolution,omitempty" name:"ThumbnailResolution"`

	// Compression format of the transcoded file. If no value or null is specified for it or the specified format is invalid, no compression file will be generated. Currently, the following compression formats are supported:
	// 
	// `zip`: generates a .zip compression package.
	// `tar.gz: generates a .tar.gz compression package.
	CompressFileType *string `json:"CompressFileType,omitempty" name:"CompressFileType"`

	// Internal parameter.
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`

	// Document transcoding priority. This parameter takes effect only for PowerPoint dynamic transcoding. Valid values:<br/>
	// - low: Low transcoding priority. The task can transcode at most 500 MB of data or a 2000-page document, with a download timeout no longer than 10 minutes. Due to resource limits, these tasks may have to queue for a long period of time. Consider this before you use this feature.
	// - null: Normal transcoding priority. The task can transcode at most 200 MB of data or a 500-page document, with a download timeout no longer than 2 minutes.
	// <br/>
	// Note: For static transcoding such as PDF transcoding, each task can transcode at most 200 MB of data regardless of the transcoding priority.
	Priority *string `json:"Priority,omitempty" name:"Priority"`

	// Minimum resolution of the transcoded document. If no value or null is specified for it or the resolution format is invalid, the original document resolution is used.
	// Higher resolution brings clearer visual effect, but also means larger size of the transcoded image resources and longer loading time of the transcoded file. Set this parameter appropriately based on your actual scenario.
	// 
	// Example: 1280x720. Note that the character between the numbers is the letter x.
	MinScaleResolution *string `json:"MinScaleResolution,omitempty" name:"MinScaleResolution"`

	// Specifies whether to enable auto handling of unsupported elements. By default, this feature is disabled.
	// 
	// If auto handling is enabled, the following processes are performed:
	// 1. Inkblots: Remove unsupported inkblots, such as those drawn by using WPS.
	// 2. Auto page flip: Clear the auto page clip settings in the PowerPoint file and set the page flip mode to mouse click.
	// 3. Corrupted audio/videos: Remove the references to corrupted audio/videos in the PowerPoint file.
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitempty" name:"AutoHandleUnsupportedElement"`
}

func (r *CreateTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Url")
	delete(f, "IsStaticPPT")
	delete(f, "MinResolution")
	delete(f, "ThumbnailResolution")
	delete(f, "CompressFileType")
	delete(f, "ExtraData")
	delete(f, "Priority")
	delete(f, "MinScaleResolution")
	delete(f, "AutoHandleUnsupportedElement")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTranscodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeResponseParams struct {
	// Unique ID of the document transcoding task, which is used to query the task progress and transcoding result
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateTranscodeResponseParams `json:"Response"`
}

func (r *CreateTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoGenerationTaskRequestParams struct {
	// ID of the recording task.
	OnlineRecordTaskId *string `json:"OnlineRecordTaskId,omitempty" name:"OnlineRecordTaskId"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Whiteboard parameters of the recording video generation task, such as the width and height of the whiteboard.
	// 
	// This parameter conflicts with the Whiteboard parameter in the API for starting a recording task. If the two Whiteboard parameters are both specified, the Whiteboard parameter in this API takes priority for recording video generation. If the Whiteboard parameter in this API is not specified, the Whiteboard parameter specified in the API for starting a recording task is used for recording video generation.
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// Video splicing parameters.
	// 
	// This parameter conflicts with the Concat parameter in the API for starting a recording task. If the two Concat parameters are both specified, the Concat parameter in this API takes priority for video splicing. If the Concat parameter in this API is not specified, the Concat parameter specified in the API for starting a recording task is used for video splicing.
	Concat *Concat `json:"Concat,omitempty" name:"Concat"`

	// Video stream mixing parameters.
	// 
	// This parameter conflicts with the MixStream parameter in the API for starting a recording task. If the two MixStream parameters are both specified, the MixStream parameter in this API takes priority for video stream mixing. If the MixStream parameter in this API is not specified, the MixStream parameter specified in the API for starting a recording task is used for video stream mixing.
	MixStream *MixStream `json:"MixStream,omitempty" name:"MixStream"`

	// A group of video generation parameters. It specifies the streams to be generated, whether to disable audio recording for a stream, and whether to record only low-resolution videos, etc.
	// 
	// This parameter conflicts with the RecordControl parameter in the API for starting a recording task. If the two RecordControl parameters are both specified, the RecordControl parameter in this API takes priority for video generation control. If the RecordControl parameter in this API is not specified, the RecordControl parameter specified in the API for starting a recording task is used for video generation control.
	RecordControl *RecordControl `json:"RecordControl,omitempty" name:"RecordControl"`

	// Internal parameter.
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`
}

type CreateVideoGenerationTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the recording task.
	OnlineRecordTaskId *string `json:"OnlineRecordTaskId,omitempty" name:"OnlineRecordTaskId"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Whiteboard parameters of the recording video generation task, such as the width and height of the whiteboard.
	// 
	// This parameter conflicts with the Whiteboard parameter in the API for starting a recording task. If the two Whiteboard parameters are both specified, the Whiteboard parameter in this API takes priority for recording video generation. If the Whiteboard parameter in this API is not specified, the Whiteboard parameter specified in the API for starting a recording task is used for recording video generation.
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// Video splicing parameters.
	// 
	// This parameter conflicts with the Concat parameter in the API for starting a recording task. If the two Concat parameters are both specified, the Concat parameter in this API takes priority for video splicing. If the Concat parameter in this API is not specified, the Concat parameter specified in the API for starting a recording task is used for video splicing.
	Concat *Concat `json:"Concat,omitempty" name:"Concat"`

	// Video stream mixing parameters.
	// 
	// This parameter conflicts with the MixStream parameter in the API for starting a recording task. If the two MixStream parameters are both specified, the MixStream parameter in this API takes priority for video stream mixing. If the MixStream parameter in this API is not specified, the MixStream parameter specified in the API for starting a recording task is used for video stream mixing.
	MixStream *MixStream `json:"MixStream,omitempty" name:"MixStream"`

	// A group of video generation parameters. It specifies the streams to be generated, whether to disable audio recording for a stream, and whether to record only low-resolution videos, etc.
	// 
	// This parameter conflicts with the RecordControl parameter in the API for starting a recording task. If the two RecordControl parameters are both specified, the RecordControl parameter in this API takes priority for video generation control. If the RecordControl parameter in this API is not specified, the RecordControl parameter specified in the API for starting a recording task is used for video generation control.
	RecordControl *RecordControl `json:"RecordControl,omitempty" name:"RecordControl"`

	// Internal parameter.
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`
}

func (r *CreateVideoGenerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoGenerationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OnlineRecordTaskId")
	delete(f, "SdkAppId")
	delete(f, "Whiteboard")
	delete(f, "Concat")
	delete(f, "MixStream")
	delete(f, "RecordControl")
	delete(f, "ExtraData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoGenerationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoGenerationTaskResponseParams struct {
	// ID of the video generation task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVideoGenerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateVideoGenerationTaskResponseParams `json:"Response"`
}

func (r *CreateVideoGenerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoGenerationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomLayout struct {
	// Mixed stream canvas parameter
	Canvas *Canvas `json:"Canvas,omitempty" name:"Canvas"`

	// Stream layout. The layout of each stream cannot exceed the canvas area.
	InputStreamList []*StreamLayout `json:"InputStreamList,omitempty" name:"InputStreamList"`
}

type DataItem struct {
	// Time. The following formats are supported:
	// yyyy-mm
	// yyyy-mm-dd
	// yyyy-mm-dd HH:MM:SS
	Time *string `json:"Time,omitempty" name:"Time"`

	// Values required for drawing charts.
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// Details of the values.
	Details []*Detail `json:"Details,omitempty" name:"Details"`
}

// Predefined struct for user
type DescribeAPIServiceRequestParams struct {
	// Supported services are cos:GetService and cdn:DescribeDomainsConfig.
	Service *string `json:"Service,omitempty" name:"Service"`

	// Request parameters in JSON format.
	Data *string `json:"Data,omitempty" name:"Data"`
}

type DescribeAPIServiceRequest struct {
	*tchttp.BaseRequest
	
	// Supported services are cos:GetService and cdn:DescribeDomainsConfig.
	Service *string `json:"Service,omitempty" name:"Service"`

	// Request parameters in JSON format.
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *DescribeAPIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAPIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIServiceResponseParams struct {
	// Response data in JSON format.
	ResponseData *string `json:"ResponseData,omitempty" name:"ResponseData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAPIServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAPIServiceResponseParams `json:"Response"`
}

func (r *DescribeAPIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationInfosRequestParams struct {

}

type DescribeApplicationInfosRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeApplicationInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationInfosResponseParams struct {
	// Application list.
	ApplicationInfos []*ApplicationItem `json:"ApplicationInfos,omitempty" name:"ApplicationInfos"`

	// Specifies whether all applications are included. The value 0 indicates no and 1 indicates yes.
	AllOption *int64 `json:"AllOption,omitempty" name:"AllOption"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationInfosResponseParams `json:"Response"`
}

func (r *DescribeApplicationInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationUsageRequestParams struct {
	// Start time of the query period. The start time point is included in the query period.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time of the query period. The end time point is not included in the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Subproduct name.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Unit of time increment.
	// - MONTHLY: month
	// - DAILY: day
	// - MINUTELY: minute
	TimeLevel *string `json:"TimeLevel,omitempty" name:"TimeLevel"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// true: Returns the weighted sum of raw usage data.
	// false: Returns the raw usage data.
	IsWeighted *bool `json:"IsWeighted,omitempty" name:"IsWeighted"`
}

type DescribeApplicationUsageRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query period. The start time point is included in the query period.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time of the query period. The end time point is not included in the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Subproduct name.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Unit of time increment.
	// - MONTHLY: month
	// - DAILY: day
	// - MINUTELY: minute
	TimeLevel *string `json:"TimeLevel,omitempty" name:"TimeLevel"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// true: Returns the weighted sum of raw usage data.
	// false: Returns the raw usage data.
	IsWeighted *bool `json:"IsWeighted,omitempty" name:"IsWeighted"`
}

func (r *DescribeApplicationUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "SubProduct")
	delete(f, "TimeLevel")
	delete(f, "SdkAppId")
	delete(f, "IsWeighted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationUsageResponseParams struct {
	// Usage data required for drawing charts.
	Data []*DataItem `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationUsageResponseParams `json:"Response"`
}

func (r *DescribeApplicationUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBoardSDKLogRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID to be used to query logs.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// User ID to be used to query logs.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Query period, which consists of two Unix timestamps in milliseconds. The first indicates the start time and the second the end time.
	TimeRange []*int64 `json:"TimeRange,omitempty" name:"TimeRange"`

	// Interval for aggregating log number queries. Example values: `5m`, `1h`, `4h`
	AggregationInterval *string `json:"AggregationInterval,omitempty" name:"AggregationInterval"`

	// Extra query conditions.
	Query *string `json:"Query,omitempty" name:"Query"`

	// Specifies whether to sort results in ascending order of time.
	Ascending *bool `json:"Ascending,omitempty" name:"Ascending"`

	// Context key used for recursive extraction. Obtain this parameter in the response to the last request.
	Context *string `json:"Context,omitempty" name:"Context"`
}

type DescribeBoardSDKLogRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID to be used to query logs.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// User ID to be used to query logs.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Query period, which consists of two Unix timestamps in milliseconds. The first indicates the start time and the second the end time.
	TimeRange []*int64 `json:"TimeRange,omitempty" name:"TimeRange"`

	// Interval for aggregating log number queries. Example values: `5m`, `1h`, `4h`
	AggregationInterval *string `json:"AggregationInterval,omitempty" name:"AggregationInterval"`

	// Extra query conditions.
	Query *string `json:"Query,omitempty" name:"Query"`

	// Specifies whether to sort results in ascending order of time.
	Ascending *bool `json:"Ascending,omitempty" name:"Ascending"`

	// Context key used for recursive extraction. Obtain this parameter in the response to the last request.
	Context *string `json:"Context,omitempty" name:"Context"`
}

func (r *DescribeBoardSDKLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBoardSDKLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserId")
	delete(f, "TimeRange")
	delete(f, "AggregationInterval")
	delete(f, "Query")
	delete(f, "Ascending")
	delete(f, "Context")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBoardSDKLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBoardSDKLogResponseParams struct {
	// Number of logs queried.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Log details.
	Sources []*string `json:"Sources,omitempty" name:"Sources"`

	// Number of logs queried within each time range after aggregation based on the time range.
	Buckets []*string `json:"Buckets,omitempty" name:"Buckets"`

	// Context key used for recursive extraction. This parameter can be used in the next request.
	Context *string `json:"Context,omitempty" name:"Context"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBoardSDKLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBoardSDKLogResponseParams `json:"Response"`
}

func (r *DescribeBoardSDKLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBoardSDKLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIMApplicationsRequestParams struct {

}

type DescribeIMApplicationsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIMApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIMApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIMApplicationsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIMApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIMApplicationsResponseParams `json:"Response"`
}

func (r *DescribeIMApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineRecordCallbackRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DescribeOfflineRecordCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeOfflineRecordCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineRecordCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOfflineRecordCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineRecordCallbackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOfflineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOfflineRecordCallbackResponseParams `json:"Response"`
}

func (r *DescribeOfflineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineRecordCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineRecordRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the offline recording task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeOfflineRecordRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the offline recording task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeOfflineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOfflineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineRecordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOfflineRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOfflineRecordResponseParams `json:"Response"`
}

func (r *DescribeOfflineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOnlineRecordCallbackRequestParams struct {
	// SdkAppId of the application
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DescribeOnlineRecordCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the application
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeOnlineRecordCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOnlineRecordCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOnlineRecordCallbackResponseParams struct {
	// Callback address of the real-time recording event. If no callback address is set, this field is null.
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// Authentication key of the real-time recording callback
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOnlineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOnlineRecordCallbackResponseParams `json:"Response"`
}

func (r *DescribeOnlineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOnlineRecordRequestParams struct {
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the real-time recording task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the real-time recording task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOnlineRecordResponseParams struct {
	// Recording stop reason
	// - AUTO: Recording automatically stops because no upstream audio/video or whiteboard operation occurs in the room for a long time.
	// - USER_CALL: The API for stopping recording is called.
	// - EXCEPTION: An exception occurred.
	// - FORCE_STOP: Recording is forcibly stopped, which is usually because the recording has been paused for more than 90 minutes or has lasted for more than 24 hours.
	FinishReason *string `json:"FinishReason,omitempty" name:"FinishReason"`

	// ID of the recording task to be queried.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Recording task status
	// - PREPARED: preparing
	// - RECORDING: recording
	// - PAUSED: recording is paused.
	// - STOPPED: recording is stopped, and the recorded video is being processed and uploaded.
	// - FINISHED: the recorded video has been processed and uploaded, and the recording result is generated.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Room ID
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// Group ID of the whiteboard
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// ID of the recording user
	RecordUserId *string `json:"RecordUserId,omitempty" name:"RecordUserId"`

	// Actual recording start time, which is a UNIX timestamp in seconds
	RecordStartTime *int64 `json:"RecordStartTime,omitempty" name:"RecordStartTime"`

	// Actual recording stop time, which is a UNIX timestamp in seconds
	RecordStopTime *int64 `json:"RecordStopTime,omitempty" name:"RecordStopTime"`

	// Total video playback duration, in milliseconds
	TotalTime *int64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// Number of exceptions during recording
	ExceptionCnt *int64 `json:"ExceptionCnt,omitempty" name:"ExceptionCnt"`

	// Duration to be deleted in the spliced video. This parameter is valid only when the video splicing feature is enabled.
	OmittedDurations []*OmittedDuration `json:"OmittedDurations,omitempty" name:"OmittedDurations"`

	// List of recorded videos
	VideoInfos []*VideoInfo `json:"VideoInfos,omitempty" name:"VideoInfos"`


	ReplayUrl *string `json:"ReplayUrl,omitempty" name:"ReplayUrl"`

	// Number of video stream interruptions during recording.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Interrupts []*Interrupt `json:"Interrupts,omitempty" name:"Interrupts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOnlineRecordResponseParams `json:"Response"`
}

func (r *DescribeOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostpaidUsageRequestParams struct {
	// Start time of the query period.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribePostpaidUsageRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query period.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribePostpaidUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostpaidUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePostpaidUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostpaidUsageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePostpaidUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribePostpaidUsageResponseParams `json:"Response"`
}

func (r *DescribePostpaidUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostpaidUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityMetricsRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Start time, which is a Unix timestamp in seconds. The time length cannot exceed seven days.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time, which is a Unix timestamp in seconds. The time length cannot exceed seven days.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Metrics to be queried. Valid values:
	//   - image_load_total_count: The number of image loads.
	//   - image_load_fail_count: The number of image load failures.
	//   - image_load_success_rate: The success rate of image loading, in percentage.
	//   - ppt_load_total_count: The number of PowerPoint file loads.
	//   - ppt_load_fail_count: The number of PowerPoint file load failures.
	//   - ppt_load_success_rate: The success rate of PowerPoint file loading, in percentage.
	//   - verify_sdk_total_count: The number of SDK verifications.
	//   - verify_sdk_fail_count: The number of SDK verification failures.
	//   - verify_sdk_success_rate: The success rate of SDK verification, in percentage.
	//   - verify_sdk_in_one_second_rate: The rate of SDK verification completed within one second, in percentage.
	//   - verify_sdk_cost_avg: The average time taken by each SDK verification, in milliseconds.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Aggregation interval. Valid value: `1h`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type DescribeQualityMetricsRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Start time, which is a Unix timestamp in seconds. The time length cannot exceed seven days.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time, which is a Unix timestamp in seconds. The time length cannot exceed seven days.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Metrics to be queried. Valid values:
	//   - image_load_total_count: The number of image loads.
	//   - image_load_fail_count: The number of image load failures.
	//   - image_load_success_rate: The success rate of image loading, in percentage.
	//   - ppt_load_total_count: The number of PowerPoint file loads.
	//   - ppt_load_fail_count: The number of PowerPoint file load failures.
	//   - ppt_load_success_rate: The success rate of PowerPoint file loading, in percentage.
	//   - verify_sdk_total_count: The number of SDK verifications.
	//   - verify_sdk_fail_count: The number of SDK verification failures.
	//   - verify_sdk_success_rate: The success rate of SDK verification, in percentage.
	//   - verify_sdk_in_one_second_rate: The rate of SDK verification completed within one second, in percentage.
	//   - verify_sdk_cost_avg: The average time taken by each SDK verification, in milliseconds.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Aggregation interval. Valid value: `1h`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeQualityMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQualityMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityMetricsResponseParams struct {
	// Query metrics.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Time series.
	Content []*TimeValue `json:"Content,omitempty" name:"Content"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeQualityMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQualityMetricsResponseParams `json:"Response"`
}

func (r *DescribeQualityMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordSearchRequestParams struct {

}

type DescribeRecordSearchRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRecordSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordSearchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordSearchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordSearchResponseParams struct {
	// The set of queried recording tasks.
	RecordTaskSet []*RecordTaskSearchResult `json:"RecordTaskSet,omitempty" name:"RecordTaskSet"`

	// Number of recording tasks.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordSearchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordSearchResponseParams `json:"Response"`
}

func (r *DescribeRecordSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomListRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Query period, which consists of two Unix timestamps in milliseconds. The first indicates the start time and the second the end time.
	TimeRange []*int64 `json:"TimeRange,omitempty" name:"TimeRange"`

	// Extra query conditions.
	Query *string `json:"Query,omitempty" name:"Query"`

	// Maximum number of data entries to be returned. Default value: 1000.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type DescribeRoomListRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Query period, which consists of two Unix timestamps in milliseconds. The first indicates the start time and the second the end time.
	TimeRange []*int64 `json:"TimeRange,omitempty" name:"TimeRange"`

	// Extra query conditions.
	Query *string `json:"Query,omitempty" name:"Query"`

	// Maximum number of data entries to be returned. Default value: 1000.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

func (r *DescribeRoomListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TimeRange")
	delete(f, "Query")
	delete(f, "MaxSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomListResponseParams struct {
	// List of rooms of the whiteboard.
	RoomList []*RoomListItem `json:"RoomList,omitempty" name:"RoomList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoomListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomListResponseParams `json:"Response"`
}

func (r *DescribeRoomListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotTaskRequestParams struct {
	// ID of the query task.
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// SdkAppId of the task.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DescribeSnapshotTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the query task.
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// SdkAppId of the task.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeSnapshotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotTaskResponseParams struct {
	// Task ID.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// Task status.
	// Running: The task is running.
	// Finished: The task is finished.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Creation time of the task. Unit: seconds.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Completion time of the task. Unit: seconds.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	FinishTime *uint64 `json:"FinishTime,omitempty" name:"FinishTime"`

	// Task result information.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Result *SnapshotResult `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSnapshotTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotTaskResponseParams `json:"Response"`
}

func (r *DescribeSnapshotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTIWDailyUsageRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Subproduct usage to be queried. The following parameters are supported:
	// - sp_tiw_board: The duration of use of whiteboards, in minutes.
	// - sp_tiw_dt: The number of pages dynamically transcoded.
	// - sp_tiw_st: The number of pages statically transcoded.
	// - sp_tiw_ric: The duration of real-time recording, in minutes.
	// 
	// Note: Dynamic transcoding multiplies the number of pages of a document by eight times. Static transcoding does not change the number of pages of a document.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Start date in the format of YYYY-MM-DD. The start date is included in the query period.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in the format of YYYY-MM-DD. The end date is included in the query period. The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeTIWDailyUsageRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Subproduct usage to be queried. The following parameters are supported:
	// - sp_tiw_board: The duration of use of whiteboards, in minutes.
	// - sp_tiw_dt: The number of pages dynamically transcoded.
	// - sp_tiw_st: The number of pages statically transcoded.
	// - sp_tiw_ric: The duration of real-time recording, in minutes.
	// 
	// Note: Dynamic transcoding multiplies the number of pages of a document by eight times. Static transcoding does not change the number of pages of a document.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Start date in the format of YYYY-MM-DD. The start date is included in the query period.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in the format of YYYY-MM-DD. The end date is included in the query period. The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeTIWDailyUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTIWDailyUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SubProduct")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTIWDailyUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTIWDailyUsageResponseParams struct {
	// Usage summary of a specified product during a specified query period.
	Usages []*UsageDataItem `json:"Usages,omitempty" name:"Usages"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTIWDailyUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTIWDailyUsageResponseParams `json:"Response"`
}

func (r *DescribeTIWDailyUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTIWDailyUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTIWRoomDailyUsageRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Subproduct usage to be queried. The following parameters are supported:
	// - sp_tiw_board: The duration of use of whiteboards, in minutes.
	// - sp_tiw_ric: The duration of real-time recording, in minutes.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Start date in the format of YYYY-MM-DD. The start date is included in the query period.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in the format of YYYY-MM-DD. The end date is included in the query period. The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Room IDs to be queried. If you leave this parameter empty, all room IDs are queried.
	RoomIDs []*uint64 `json:"RoomIDs,omitempty" name:"RoomIDs"`

	// Offset for query. Default value: `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries returned per query. Default value: `20`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTIWRoomDailyUsageRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Subproduct usage to be queried. The following parameters are supported:
	// - sp_tiw_board: The duration of use of whiteboards, in minutes.
	// - sp_tiw_ric: The duration of real-time recording, in minutes.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Start date in the format of YYYY-MM-DD. The start date is included in the query period.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in the format of YYYY-MM-DD. The end date is included in the query period. The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Room IDs to be queried. If you leave this parameter empty, all room IDs are queried.
	RoomIDs []*uint64 `json:"RoomIDs,omitempty" name:"RoomIDs"`

	// Offset for query. Default value: `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries returned per query. Default value: `20`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTIWRoomDailyUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTIWRoomDailyUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SubProduct")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomIDs")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTIWRoomDailyUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTIWRoomDailyUsageResponseParams struct {
	// Usage of the specified product per room during the specified query period.
	Usages []*RoomUsageDataItem `json:"Usages,omitempty" name:"Usages"`

	// Number of usage data entries.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTIWRoomDailyUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTIWRoomDailyUsageResponseParams `json:"Response"`
}

func (r *DescribeTIWRoomDailyUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTIWRoomDailyUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeCallbackRequestParams struct {
	// SdkAppId of the application
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DescribeTranscodeCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the application
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeTranscodeCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeCallbackResponseParams struct {
	// Document transcoding callback address
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// Authentication key of the document transcoding callback
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTranscodeCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTranscodeCallbackResponseParams `json:"Response"`
}

func (r *DescribeTranscodeCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeRequestParams struct {
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Unique ID of the document transcoding task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTranscodeRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Unique ID of the document transcoding task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeResponseParams struct {
	// Total number of document pages
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Transcoding progress. Value range: 0 to 100
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Document resolution
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// URL of the transcoding result
	// Dynamic transcoding: link of the HTML5 page transcoded from a PowerPoint file
	// Static transcoding: URL prefix of the image transcoded for each document page. For example, if the URL prefix is `http://example.com/g0jb42ps49vtebjshilb/`, the image URL of the first page is
	// `http://example.com/g0jb42ps49vtebjshilb/1.jpg`, and so on.
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// Current task state
	// - QUEUED: queuing for transcoding
	// - PROCESSING: transcoding is in progress
	// - FINISHED: transcoded
	Status *string `json:"Status,omitempty" name:"Status"`

	// Unique ID of the transcoding task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Document name
	Title *string `json:"Title,omitempty" name:"Title"`

	// URL prefix of the thumbnail. If the URL prefix is `http://example.com/g0jb42ps49vtebjshilb/ `, the thumbnail URL for the first page of the dynamically transcoded PowerPoint file is
	// `http://example.com/g0jb42ps49vtebjshilb/1.jpg`, and so on.
	// 
	// If the document transcoding request carries the ThumbnailResolution parameter and the transcoding type is dynamic transcoding, this parameter is not null. In other cases, this parameter is null.
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" name:"ThumbnailUrl"`

	// Resolution of the thumbnail generated for dynamic transcoding
	ThumbnailResolution *string `json:"ThumbnailResolution,omitempty" name:"ThumbnailResolution"`

	// URL for downloading the transcoded and compressed file. If `CompressFileType` carried in the document transcoding request is null or is not a supported compression format, this parameter is null.
	CompressFileUrl *string `json:"CompressFileUrl,omitempty" name:"CompressFileUrl"`

	// Download URL (for trial) of the resource list.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	ResourceListUrl *string `json:"ResourceListUrl,omitempty" name:"ResourceListUrl"`

	// Document generation mode (for trial).
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Ext *string `json:"Ext,omitempty" name:"Ext"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTranscodeResponseParams `json:"Response"`
}

func (r *DescribeTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeSearchRequestParams struct {

}

type DescribeTranscodeSearchRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTranscodeSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeSearchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeSearchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeSearchResponseParams struct {
	// The set of queried transcoding tasks.
	TranscodeTaskSet []*TranscodeTaskSearchResult `json:"TranscodeTaskSet,omitempty" name:"TranscodeTaskSet"`

	// Number of transcoding tasks.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTranscodeSearchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTranscodeSearchResponseParams `json:"Response"`
}

func (r *DescribeTranscodeSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageSummaryRequestParams struct {
	// Start time of the query period.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Subproducts whose usage you want to query.
	SubProducts []*string `json:"SubProducts,omitempty" name:"SubProducts"`

	// true: Returns weighted data.
	// false: Returns raw data.
	IsWeighted *bool `json:"IsWeighted,omitempty" name:"IsWeighted"`
}

type DescribeUsageSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query period.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Subproducts whose usage you want to query.
	SubProducts []*string `json:"SubProducts,omitempty" name:"SubProducts"`

	// true: Returns weighted data.
	// false: Returns raw data.
	IsWeighted *bool `json:"IsWeighted,omitempty" name:"IsWeighted"`
}

func (r *DescribeUsageSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "SubProducts")
	delete(f, "IsWeighted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageSummaryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsageSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsageSummaryResponseParams `json:"Response"`
}

func (r *DescribeUsageSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserListRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID to be used to query users.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// Query period, which consists of two Unix timestamps in milliseconds. The first indicates the start time and the second the end time.
	TimeRange []*int64 `json:"TimeRange,omitempty" name:"TimeRange"`

	// Extra query conditions.
	Query *string `json:"Query,omitempty" name:"Query"`

	// Maximum number of data entries to be returned. Default value: `1000`.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type DescribeUserListRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID to be used to query users.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// Query period, which consists of two Unix timestamps in milliseconds. The first indicates the start time and the second the end time.
	TimeRange []*int64 `json:"TimeRange,omitempty" name:"TimeRange"`

	// Extra query conditions.
	Query *string `json:"Query,omitempty" name:"Query"`

	// Maximum number of data entries to be returned. Default value: `1000`.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
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
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "TimeRange")
	delete(f, "Query")
	delete(f, "MaxSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserListResponseParams struct {
	// User list of the room.
	UserList []*UserListItem `json:"UserList,omitempty" name:"UserList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeUserResourcesRequestParams struct {

}

type DescribeUserResourcesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserResourcesResponseParams `json:"Response"`
}

func (r *DescribeUserResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserStatusRequestParams struct {

}

type DescribeUserStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserStatusResponseParams struct {
	// AppId of the customer.
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Specifies whether the whiteboard service of the trial or official edition is activated before.
	// 
	// 0: The whiteboard service is not activated.
	// 1: The whiteboard service is activated.
	IsTiwUser *int64 `json:"IsTiwUser,omitempty" name:"IsTiwUser"`

	// Specifies whether the interactive class feature of the trial or official edition is activated before.
	IsSaaSUser *int64 `json:"IsSaaSUser,omitempty" name:"IsSaaSUser"`

	// Specifies whether the user uses the offline recording feature of the whiteboard service.
	IsTiwOfflineRecordUser *int64 `json:"IsTiwOfflineRecordUser,omitempty" name:"IsTiwOfflineRecordUser"`

	// Specifies whether the user is authenticated.
	IsAuthenticated *int64 `json:"IsAuthenticated,omitempty" name:"IsAuthenticated"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserStatusResponseParams `json:"Response"`
}

func (r *DescribeUserStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoGenerationTaskCallbackRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DescribeVideoGenerationTaskCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeVideoGenerationTaskCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoGenerationTaskCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoGenerationTaskCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoGenerationTaskCallbackResponseParams struct {
	// Callback URL for recording video generation.
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// Callback authentication key for recording video generation.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVideoGenerationTaskCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoGenerationTaskCallbackResponseParams `json:"Response"`
}

func (r *DescribeVideoGenerationTaskCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoGenerationTaskCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoGenerationTaskRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the recording video generation task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeVideoGenerationTaskRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the recording video generation task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeVideoGenerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoGenerationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoGenerationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoGenerationTaskResponseParams struct {
	// Group ID corresponding to the task.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Room ID corresponding to the task.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Disused.
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Status of the recording video generation task. Valid values:
	// - QUEUED: Queuing.
	// - PROCESSING: Video generation in progress.
	// - FINISHED: Video generation finished. (To determine whether the task succeeded or failed, check the error code and message.)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Total video playback duration. Unit: milliseconds.
	TotalTime *int64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// Disused. Use the `VideoInfoList` parameter.
	VideoInfos *VideoInfo `json:"VideoInfos,omitempty" name:"VideoInfos"`

	// List of videos generated by the recording video generation tasks.
	VideoInfoList []*VideoInfo `json:"VideoInfoList,omitempty" name:"VideoInfoList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVideoGenerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoGenerationTaskResponseParams `json:"Response"`
}

func (r *DescribeVideoGenerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoGenerationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardApplicationConfigRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Task types to be queried.
	// recording: Real-time recording.
	// transcode: Document transcoding.
	TaskTypes []*string `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// SdkAppIds to be used to query configurations.
	SdkAppIds []*int64 `json:"SdkAppIds,omitempty" name:"SdkAppIds"`
}

type DescribeWhiteboardApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Task types to be queried.
	// recording: Real-time recording.
	// transcode: Document transcoding.
	TaskTypes []*string `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// SdkAppIds to be used to query configurations.
	SdkAppIds []*int64 `json:"SdkAppIds,omitempty" name:"SdkAppIds"`
}

func (r *DescribeWhiteboardApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskTypes")
	delete(f, "SdkAppIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteboardApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardApplicationConfigResponseParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Task-related configurations of the whiteboard application.
	Configs []*WhiteboardApplicationConfig `json:"Configs,omitempty" name:"Configs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteboardApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteboardApplicationConfigResponseParams `json:"Response"`
}

func (r *DescribeWhiteboardApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardBucketConfigRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Task type to be queried.
	// recording: Real-time recording.
	// transcode: Document transcoding.
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

type DescribeWhiteboardBucketConfigRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Task type to be queried.
	// recording: Real-time recording.
	// transcode: Document transcoding.
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *DescribeWhiteboardBucketConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardBucketConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteboardBucketConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardBucketConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteboardBucketConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteboardBucketConfigResponseParams `json:"Response"`
}

func (r *DescribeWhiteboardBucketConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardBucketConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushCallbackRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DescribeWhiteboardPushCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeWhiteboardPushCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteboardPushCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushCallbackResponseParams struct {
	// Callback URL of the whiteboard push event. If no callback URL is set, this parameter is null.
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// Callback authentication key for whiteboard push.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteboardPushCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteboardPushCallbackResponseParams `json:"Response"`
}

func (r *DescribeWhiteboardPushCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the whiteboard push task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeWhiteboardPushRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the whiteboard push task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeWhiteboardPushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteboardPushRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushResponseParams struct {
	// Reason for push stop.
	// - AUTO: Pushing automatically stops because no upstream audio/video or whiteboard operation occurs in the room for a long time.
	// - USER_CALL: The API for stopping pushing is called.
	// - EXCEPTION: An exception occurred.
	FinishReason *string `json:"FinishReason,omitempty" name:"FinishReason"`

	// ID of the whiteboard push task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Status of the push task.
	// - PREPARED: Push in preparation (including entering the room and starting the push service).
	// - PUSHING: Pushing in progress.
	// - STOPPED: Pushing stopped.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Room ID.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// Group ID of the whiteboard.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// User ID of the push task.
	PushUserId *string `json:"PushUserId,omitempty" name:"PushUserId"`

	// Actual push start time, which is a Unix timestamp in seconds.
	PushStartTime *int64 `json:"PushStartTime,omitempty" name:"PushStartTime"`

	// Actual push stop time, which is a Unix timestamp in seconds.
	PushStopTime *int64 `json:"PushStopTime,omitempty" name:"PushStopTime"`

	// Number of exceptions during push.
	ExceptionCnt *int64 `json:"ExceptionCnt,omitempty" name:"ExceptionCnt"`

	// IM timestamp corresponding to the first frame of the whiteboard push video. The timestamp is used for time synchronization between IM messages and the whiteboard push video during playback.
	IMSyncTime *int64 `json:"IMSyncTime,omitempty" name:"IMSyncTime"`

	// Result information of the backup push task.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Backup *string `json:"Backup,omitempty" name:"Backup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteboardPushResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteboardPushResponseParams `json:"Response"`
}

func (r *DescribeWhiteboardPushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushSearchRequestParams struct {

}

type DescribeWhiteboardPushSearchRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWhiteboardPushSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushSearchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteboardPushSearchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushSearchResponseParams struct {
	// The set of queried push tasks.
	WhiteboardPushTaskSet []*WhiteboardPushTaskSearchResult `json:"WhiteboardPushTaskSet,omitempty" name:"WhiteboardPushTaskSet"`

	// Number of push tasks.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteboardPushSearchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteboardPushSearchResponseParams `json:"Response"`
}

func (r *DescribeWhiteboardPushSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Detail struct {
	// Usage metric.
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// Usage weight.
	Weight *float64 `json:"Weight,omitempty" name:"Weight"`

	// Usage value.
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type Interrupt struct {
	// User ID.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Number of video stream interruptions.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type LayoutParams struct {
	// Stream image width. Value range: [2,3000]
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Stream image height. Value range: [2,3000]
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Offset of the top point in the upper-left corner of the current image to the X axis of the top point in the upper-left corner of the canvas. Default value: 0. Value range: [0,3000].
	X *int64 `json:"X,omitempty" name:"X"`

	// Offset of the top point in the upper-left corner of the current image to the Y axis of the top point in the upper-left corner of the canvas. Default value: 0. Value range: [0,3000].
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// Z-axis position of the image. The default value is 0.
	// The Z axis determines the overlap sequence of images. The image with the largest z-axis value is at the top layer.
	ZOrder *int64 `json:"ZOrder,omitempty" name:"ZOrder"`
}

type MixStream struct {
	// Whether stream mixing is enabled
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// Whether audio stream mixing is disabled
	DisableAudio *bool `json:"DisableAudio,omitempty" name:"DisableAudio"`

	// ID of the embedded mixed stream layout template. Valid values: 1 and 2. For more information on the differences of both values, see the sample embedded mixed stream layout template.
	// If the Custom field is not specified, ModelId is required.
	ModelId *int64 `json:"ModelId,omitempty" name:"ModelId"`

	// ID of a teacher account
	// This field is valid only when ModelId is specified.
	// If you specify TeacherID for a user, the user's video stream will be displayed in the first image of the embedded template.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// Custom mixed stream layout parameter
	// If this parameter is available, the ModelId and TeacherId fields will be ignored.
	Custom *CustomLayout `json:"Custom,omitempty" name:"Custom"`
}

// Predefined struct for user
type ModifyApplicationRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type ModifyApplicationRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
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
	delete(f, "SdkAppId")
	delete(f, "AppName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyAutoRenewFlagRequestParams struct {
	// Subproduct ID. To obtain this ID, call the `DescribeUserResources` API and find the subproduct ID of the monthly feature package with the level 1. Usually the value is `sp_tiw_package`.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Resource ID. To obtain this ID, call the `DescribeUserResources` API and find the resource ID of the monthly feature package with the level 1.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Renewal mode. Valid values: `0` (default mode, which is auto-renewal), `1` (auto-renewal), `2` (manual renewal, which is specified by users). If no renewal is required, set it to `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// Subproduct ID. To obtain this ID, call the `DescribeUserResources` API and find the subproduct ID of the monthly feature package with the level 1. Usually the value is `sp_tiw_package`.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Resource ID. To obtain this ID, call the `DescribeUserResources` API and find the resource ID of the monthly feature package with the level 1.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Renewal mode. Valid values: `0` (default mode, which is auto-renewal), `1` (auto-renewal), `2` (manual renewal, which is specified by users). If no renewal is required, set it to `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *ModifyAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubProduct")
	delete(f, "ResourceId")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoRenewFlagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWhiteboardApplicationConfigRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Task-related configurations of the whiteboard application.
	Configs []*WhiteboardApplicationConfig `json:"Configs,omitempty" name:"Configs"`
}

type ModifyWhiteboardApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Task-related configurations of the whiteboard application.
	Configs []*WhiteboardApplicationConfig `json:"Configs,omitempty" name:"Configs"`
}

func (r *ModifyWhiteboardApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWhiteboardApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Configs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWhiteboardApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWhiteboardApplicationConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWhiteboardApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWhiteboardApplicationConfigResponseParams `json:"Response"`
}

func (r *ModifyWhiteboardApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWhiteboardApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWhiteboardBucketConfigRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Task type to be queried.
	// recording: Real-time recording.
	// transcode: Document transcoding.
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Name of the COS bucket.
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// Region of the COS bucket.
	BucketLocation *string `json:"BucketLocation,omitempty" name:"BucketLocation"`

	// Resource prefix of the bucket.
	BucketPrefix *string `json:"BucketPrefix,omitempty" name:"BucketPrefix"`

	// Domain name of the URL of the bucket.
	ResultDomain *string `json:"ResultDomain,omitempty" name:"ResultDomain"`
}

type ModifyWhiteboardBucketConfigRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Task type to be queried.
	// recording: Real-time recording.
	// transcode: Document transcoding.
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Name of the COS bucket.
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// Region of the COS bucket.
	BucketLocation *string `json:"BucketLocation,omitempty" name:"BucketLocation"`

	// Resource prefix of the bucket.
	BucketPrefix *string `json:"BucketPrefix,omitempty" name:"BucketPrefix"`

	// Domain name of the URL of the bucket.
	ResultDomain *string `json:"ResultDomain,omitempty" name:"ResultDomain"`
}

func (r *ModifyWhiteboardBucketConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWhiteboardBucketConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskType")
	delete(f, "BucketName")
	delete(f, "BucketLocation")
	delete(f, "BucketPrefix")
	delete(f, "ResultDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWhiteboardBucketConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWhiteboardBucketConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWhiteboardBucketConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWhiteboardBucketConfigResponseParams `json:"Response"`
}

func (r *ModifyWhiteboardBucketConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWhiteboardBucketConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OmittedDuration struct {
	// Offset of the paused time in the spliced video, in milliseconds
	VideoTime *int64 `json:"VideoTime,omitempty" name:"VideoTime"`

	// Recording pause timestamp, in milliseconds
	PauseTime *int64 `json:"PauseTime,omitempty" name:"PauseTime"`

	// Recording resumption timestamp, in milliseconds
	ResumeTime *int64 `json:"ResumeTime,omitempty" name:"ResumeTime"`
}

// Predefined struct for user
type PauseOnlineRecordRequestParams struct {
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the real-time recording task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type PauseOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the real-time recording task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *PauseOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseOnlineRecordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PauseOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *PauseOnlineRecordResponseParams `json:"Response"`
}

func (r *PauseOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordControl struct {
	// It specifies whether to enable RecordControl. Valid values: true (yes); false (no).
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// A global parameter generally used in conjunction with `StreamControls` that specifies whether to disable recording. Valid values:
	// 
	// true: no stream is recorded.
	// false: all streams are recorded. Default value: false.
	// 
	// The setting in this parameter is applied to all streams. However, if `StreamControls` is passed in, the parameters in `StreamControls` will take precedence.
	DisableRecord *bool `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// A global parameter generally used in conjunction with `StreamControls` that specifies whether to disable audio recording over all streams. Valid values:
	// 
	// true: no audio recording of any streams.
	// false: audio recording of all streams. Default value: false.
	// 
	// The setting in this parameter is applied to all streams. However, if `StreamControls` is passed in, the parameters in `StreamControls` will take precedence.
	DisableAudio *bool `json:"DisableAudio,omitempty" name:"DisableAudio"`

	// A global parameter generally used in conjunction with `StreamControls` that specifies whether to record low-resolution videos only. Valid values:
	// 
	// true: only records low-resolution videos for all streams. Please ensure that the up-streaming end pushes the low-resolution videos. Otherwise, the recorded video may be black.
	// false: high-resolution video recording of all streams. Default value: false.
	// 
	// The setting in this parameter is applied to all streams. However, if `StreamControls` is passed in, the parameters in `StreamControls` will take precedence.
	PullSmallVideo *bool `json:"PullSmallVideo,omitempty" name:"PullSmallVideo"`

	// Parameters over specific streams, which take priority over global configurations. If it’s empty, all streams are recorded according to the global configurations. 
	StreamControls []*StreamControl `json:"StreamControls,omitempty" name:"StreamControls"`
}

type RecordTaskResult struct {
	// `AUTO`: Recording automatically stops. `USER_CALL`: The API for stopping recording is called.
	FinishReason *string `json:"FinishReason,omitempty" name:"FinishReason"`

	// Number of exceptions.
	ExceptionCnt *int64 `json:"ExceptionCnt,omitempty" name:"ExceptionCnt"`

	// Room ID.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// Group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Actual recording start time.
	RecordStartTime *int64 `json:"RecordStartTime,omitempty" name:"RecordStartTime"`

	// Recording end time.
	RecordStopTime *int64 `json:"RecordStopTime,omitempty" name:"RecordStopTime"`

	// Recording duration.
	TotalTime *int64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// List of video information.
	VideoInfos []*VideoInfo `json:"VideoInfos,omitempty" name:"VideoInfos"`

	// Omitted video durations.
	OmittedDurations []*OmittedDuration `json:"OmittedDurations,omitempty" name:"OmittedDurations"`

	// Details.
	Details *string `json:"Details,omitempty" name:"Details"`

	// Task execution error code.
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// Error message.
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

type RecordTaskSearchResult struct {
	// Unique task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Status of the real-time recording task.
	// - PAUSED: Recording paused.
	// - PREPARED: Recording in preparation.
	// - RECORDING: Recording in progress.
	// - STOPPED: Recording stopped.
	// - FINISHED: Recording finished.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Room ID of the real-time recording task.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// Creation time of the task.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Real-time recording result.
	Result *RecordTaskResult `json:"Result,omitempty" name:"Result"`
}

// Predefined struct for user
type ResumeOnlineRecordRequestParams struct {
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the resumed real-time recording task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type ResumeOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the resumed real-time recording task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ResumeOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeOnlineRecordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResumeOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *ResumeOnlineRecordResponseParams `json:"Response"`
}

func (r *ResumeOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoomListItem struct {
	// Room ID.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// The first time when the room ID appeared during the queried period, which is a Unix timestamp in milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// The last time when the room ID appeared during the queried period, which is a Unix timestamp in milliseconds.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Number of users in the room.
	UserNumber *int64 `json:"UserNumber,omitempty" name:"UserNumber"`
}

type RoomUsageDataItem struct {
	// Start date in the format of YYYY-MM-DD.
	Time *string `json:"Time,omitempty" name:"Time"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Subproduct usage information, which is consistent with the corresponding request parameters.
	// - sp_tiw_board: The duration of use of whiteboards.
	// - sp_tiw_ric: The duration of real-time recording.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Usage values.
	// - The unit of sp_tiw_board and sp_tiw_ric is minutes.
	Value *float64 `json:"Value,omitempty" name:"Value"`


	RoomID *uint64 `json:"RoomID,omitempty" name:"RoomID"`
}

// Predefined struct for user
type SetOfflineRecordCallbackRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback URL of the offline recording task result. If it is specified as null, the set callback URL is deleted. The callback URL only supports the HTTP or HTTPS protocol, and therefore the callback URL must start with `http://` or `https://`.
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

type SetOfflineRecordCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback URL of the offline recording task result. If it is specified as null, the set callback URL is deleted. The callback URL only supports the HTTP or HTTPS protocol, and therefore the callback URL must start with `http://` or `https://`.
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *SetOfflineRecordCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOfflineRecordCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetOfflineRecordCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOfflineRecordCallbackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetOfflineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetOfflineRecordCallbackResponseParams `json:"Response"`
}

func (r *SetOfflineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOfflineRecordCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOnlineRecordCallbackKeyRequestParams struct {
	// SdkAppId of the application
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Authentication key for the real-time recording callback. It is a string that can have up to 64 characters. If an empty string is passed in, the existing callback authentication key will be deleted. For more information, please [see here](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

type SetOnlineRecordCallbackKeyRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the application
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Authentication key for the real-time recording callback. It is a string that can have up to 64 characters. If an empty string is passed in, the existing callback authentication key will be deleted. For more information, please [see here](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *SetOnlineRecordCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetOnlineRecordCallbackKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOnlineRecordCallbackKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetOnlineRecordCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *SetOnlineRecordCallbackKeyResponseParams `json:"Response"`
}

func (r *SetOnlineRecordCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOnlineRecordCallbackRequestParams struct {
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback address of the real-time recording task result. If an empty string is passed in, the existing callback address will be deleted. The callback address only supports the HTTP or HTTPS protocol, so the callback address must start with `http://` or `https://`. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40258?from_cn_redirect=1).
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

type SetOnlineRecordCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback address of the real-time recording task result. If an empty string is passed in, the existing callback address will be deleted. The callback address only supports the HTTP or HTTPS protocol, so the callback address must start with `http://` or `https://`. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40258?from_cn_redirect=1).
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *SetOnlineRecordCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetOnlineRecordCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOnlineRecordCallbackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetOnlineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetOnlineRecordCallbackResponseParams `json:"Response"`
}

func (r *SetOnlineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTranscodeCallbackKeyRequestParams struct {
	// SdkAppId of the application
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Authentication key for the document transcoding callback. It is a string that can have up to 64 characters. If an empty string is passed in, the existing callback authentication key will be deleted. For more information about callback authentication, please [see here](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

type SetTranscodeCallbackKeyRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the application
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Authentication key for the document transcoding callback. It is a string that can have up to 64 characters. If an empty string is passed in, the existing callback authentication key will be deleted. For more information about callback authentication, please [see here](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *SetTranscodeCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTranscodeCallbackKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTranscodeCallbackKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTranscodeCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *SetTranscodeCallbackKeyResponseParams `json:"Response"`
}

func (r *SetTranscodeCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTranscodeCallbackRequestParams struct {
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback address for the document transcoding progress. If an empty string is passed in, the existing callback address will be deleted. The callback address only supports the HTTP or HTTPS protocol, so the callback address must start with `http://` or `https://`.
	// For more information about the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40260?from_cn_redirect=1).
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

type SetTranscodeCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback address for the document transcoding progress. If an empty string is passed in, the existing callback address will be deleted. The callback address only supports the HTTP or HTTPS protocol, so the callback address must start with `http://` or `https://`.
	// For more information about the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40260?from_cn_redirect=1).
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *SetTranscodeCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTranscodeCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTranscodeCallbackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTranscodeCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetTranscodeCallbackResponseParams `json:"Response"`
}

func (r *SetTranscodeCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVideoGenerationTaskCallbackKeyRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback authentication key for recording video generation. It is a string that can have up to 64 characters. If an empty string is passed in, the existing callback authentication key is deleted.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

type SetVideoGenerationTaskCallbackKeyRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback authentication key for recording video generation. It is a string that can have up to 64 characters. If an empty string is passed in, the existing callback authentication key is deleted.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *SetVideoGenerationTaskCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVideoGenerationTaskCallbackKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetVideoGenerationTaskCallbackKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVideoGenerationTaskCallbackKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetVideoGenerationTaskCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *SetVideoGenerationTaskCallbackKeyResponseParams `json:"Response"`
}

func (r *SetVideoGenerationTaskCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVideoGenerationTaskCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVideoGenerationTaskCallbackRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback URL of the offline recording task result. If it is specified as null, the set callback URL is deleted. The callback URL only supports the HTTP or HTTPS protocol, and therefore the callback URL must start with `http://` or `https://`.
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

type SetVideoGenerationTaskCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback URL of the offline recording task result. If it is specified as null, the set callback URL is deleted. The callback URL only supports the HTTP or HTTPS protocol, and therefore the callback URL must start with `http://` or `https://`.
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *SetVideoGenerationTaskCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVideoGenerationTaskCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetVideoGenerationTaskCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVideoGenerationTaskCallbackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetVideoGenerationTaskCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetVideoGenerationTaskCallbackResponseParams `json:"Response"`
}

func (r *SetVideoGenerationTaskCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVideoGenerationTaskCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWhiteboardPushCallbackKeyRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback authentication key for whiteboard push. It is a string that can have up to 64 characters. If an empty string is passed in, the existing callback authentication key is deleted. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

type SetWhiteboardPushCallbackKeyRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback authentication key for whiteboard push. It is a string that can have up to 64 characters. If an empty string is passed in, the existing callback authentication key is deleted. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *SetWhiteboardPushCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWhiteboardPushCallbackKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetWhiteboardPushCallbackKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWhiteboardPushCallbackKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetWhiteboardPushCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *SetWhiteboardPushCallbackKeyResponseParams `json:"Response"`
}

func (r *SetWhiteboardPushCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWhiteboardPushCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWhiteboardPushCallbackRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback URL of the whiteboard push task result. If an empty string is passed in, the existing callback URL is deleted. The callback URL only supports the HTTP or HTTPS protocol, and therefore the callback URL must start with `http://` or `https://`. For the callback format, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

type SetWhiteboardPushCallbackRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback URL of the whiteboard push task result. If an empty string is passed in, the existing callback URL is deleted. The callback URL only supports the HTTP or HTTPS protocol, and therefore the callback URL must start with `http://` or `https://`. For the callback format, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *SetWhiteboardPushCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWhiteboardPushCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetWhiteboardPushCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWhiteboardPushCallbackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetWhiteboardPushCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetWhiteboardPushCallbackResponseParams `json:"Response"`
}

func (r *SetWhiteboardPushCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWhiteboardPushCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotCOS struct {
	// UIN of the Tencent Cloud account.
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// COS region.
	Region *string `json:"Region,omitempty" name:"Region"`

	// COS bucket name.
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Root directory for storing whiteboard snapshots.
	TargetDir *string `json:"TargetDir,omitempty" name:"TargetDir"`

	// CDN acceleration domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type SnapshotResult struct {
	// Task execution error code.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// Task execution error message.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// Number of generated snapshot images.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// List of URLs of the snapshot images.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Snapshots []*string `json:"Snapshots,omitempty" name:"Snapshots"`
}

type SnapshotWhiteboard struct {
	// Whiteboard width. Valid range: [0,2560]. Default value: 1280.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Whiteboard height. Valid range: [0,2560]. Default value: 720.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Escaped JSON string of the whiteboard initial parameters, which is passed through to the Tencent Interactive Whiteboard SDK.
	InitParams *string `json:"InitParams,omitempty" name:"InitParams"`
}

// Predefined struct for user
type StartOnlineRecordRequestParams struct {
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the room for recording. Value range: (1, 4294967295)
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// User ID used by the recording service for entering a room. The ID cannot exceed 60 bytes in length. Its format is `tic_record_user_${RoomId}_${Random}`, where `${RoomId}` indicates the ID of the room for recording and `${Random}` is a random string.
	// The ID must be an unused ID in the SDK. The recording service uses the user ID to enter the room for audio, video, and whiteboard recording. If this ID is already used in the SDK, the SDK and recording service will conflict, affecting the recording operation.
	RecordUserId *string `json:"RecordUserId,omitempty" name:"RecordUserId"`

	// Signature corresponding to RecordUserId
	RecordUserSig *string `json:"RecordUserSig,omitempty" name:"RecordUserSig"`

	// (Disused) IM group ID of the whiteboard. By default, it is the same as the room ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Real-time recording video splicing parameter
	Concat *Concat `json:"Concat,omitempty" name:"Concat"`

	// Real-time recording whiteboard parameter, such as the whiteboard width and height
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// Real-time recording stream mixing parameter
	// Notes:
	// 1. The stream mixing feature needs to be enabled separately. If you need the feature, contact TIW customer service.
	// 2. To use the stream mixing feature, the Extras parameter is required and must contain "MIX_STREAM".
	MixStream *MixStream `json:"MixStream,omitempty" name:"MixStream"`

	// List of advanced features used
	// List of possible values:
	// MIX_STREAM - Stream mixing feature
	Extras []*string `json:"Extras,omitempty" name:"Extras"`

	// Whether to return the audio-only recording file of different streams in the result callback. The file format is mp3.
	AudioFileNeeded *bool `json:"AudioFileNeeded,omitempty" name:"AudioFileNeeded"`

	// A group of real-time recording parameters. It specifies the streams to be recorded, whether to disable the audio recording, and whether to record only low-resolution videos, etc.
	RecordControl *RecordControl `json:"RecordControl,omitempty" name:"RecordControl"`


	RecordMode *string `json:"RecordMode,omitempty" name:"RecordMode"`


	ChatGroupId *string `json:"ChatGroupId,omitempty" name:"ChatGroupId"`

	// Recording timeout. Unit: seconds. Valid range: [300,86400]. Default value: 300.
	// 
	// If no upstream audio/video exists or no operation is performed on the whiteboard for the specified period of time, the recording service automatically stops the recording task.
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitempty" name:"AutoStopTimeout"`

	// Internal parameter. You can ignore this parameter.
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`
}

type StartOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the room for recording. Value range: (1, 4294967295)
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// User ID used by the recording service for entering a room. The ID cannot exceed 60 bytes in length. Its format is `tic_record_user_${RoomId}_${Random}`, where `${RoomId}` indicates the ID of the room for recording and `${Random}` is a random string.
	// The ID must be an unused ID in the SDK. The recording service uses the user ID to enter the room for audio, video, and whiteboard recording. If this ID is already used in the SDK, the SDK and recording service will conflict, affecting the recording operation.
	RecordUserId *string `json:"RecordUserId,omitempty" name:"RecordUserId"`

	// Signature corresponding to RecordUserId
	RecordUserSig *string `json:"RecordUserSig,omitempty" name:"RecordUserSig"`

	// (Disused) IM group ID of the whiteboard. By default, it is the same as the room ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Real-time recording video splicing parameter
	Concat *Concat `json:"Concat,omitempty" name:"Concat"`

	// Real-time recording whiteboard parameter, such as the whiteboard width and height
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// Real-time recording stream mixing parameter
	// Notes:
	// 1. The stream mixing feature needs to be enabled separately. If you need the feature, contact TIW customer service.
	// 2. To use the stream mixing feature, the Extras parameter is required and must contain "MIX_STREAM".
	MixStream *MixStream `json:"MixStream,omitempty" name:"MixStream"`

	// List of advanced features used
	// List of possible values:
	// MIX_STREAM - Stream mixing feature
	Extras []*string `json:"Extras,omitempty" name:"Extras"`

	// Whether to return the audio-only recording file of different streams in the result callback. The file format is mp3.
	AudioFileNeeded *bool `json:"AudioFileNeeded,omitempty" name:"AudioFileNeeded"`

	// A group of real-time recording parameters. It specifies the streams to be recorded, whether to disable the audio recording, and whether to record only low-resolution videos, etc.
	RecordControl *RecordControl `json:"RecordControl,omitempty" name:"RecordControl"`

	RecordMode *string `json:"RecordMode,omitempty" name:"RecordMode"`

	ChatGroupId *string `json:"ChatGroupId,omitempty" name:"ChatGroupId"`

	// Recording timeout. Unit: seconds. Valid range: [300,86400]. Default value: 300.
	// 
	// If no upstream audio/video exists or no operation is performed on the whiteboard for the specified period of time, the recording service automatically stops the recording task.
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitempty" name:"AutoStopTimeout"`

	// Internal parameter. You can ignore this parameter.
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`
}

func (r *StartOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "RecordUserId")
	delete(f, "RecordUserSig")
	delete(f, "GroupId")
	delete(f, "Concat")
	delete(f, "Whiteboard")
	delete(f, "MixStream")
	delete(f, "Extras")
	delete(f, "AudioFileNeeded")
	delete(f, "RecordControl")
	delete(f, "RecordMode")
	delete(f, "ChatGroupId")
	delete(f, "AutoStopTimeout")
	delete(f, "ExtraData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartOnlineRecordResponseParams struct {
	// ID of the real-time recording task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *StartOnlineRecordResponseParams `json:"Response"`
}

func (r *StartOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartWhiteboardPushRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID of the whiteboard push task. Valid range: (1,4294967295).
	// 
	// 1. By default, the whiteboard push task uses the RoomId string as the GroupID of the IM group. For example, if the RoomId string is 1234, the GroupID of the IM group is also 1234. Message synchronization requires joining a group. Make sure that an IM group is created before the push starts. Otherwise, the push fails.
	// 2. If neither TRTCRoomId nor TRTCRoomIdStr is specified, the value of RoomId is used as the Tencent Real-Time Communication (TRTC) room ID for the whiteboard push task.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// User ID used by the whiteboard push service for entering the whiteboard room. If `IMAuthParam`and `TRTCAuthParam` are not specified, this user ID is used for operations such as logging in to the IM application, joining an IM group, and entering the room for TRTC whiteboard push.
	// The ID cannot exceed 60 bytes in length and must be an unused ID. The whiteboard push service uses the user ID to enter the room for whiteboard audio/video push. If this ID is already used in another scenario, the other scenario and whiteboard push service will conflict, affecting the pushing operation.
	PushUserId *string `json:"PushUserId,omitempty" name:"PushUserId"`

	// User's IM signature corresponding to the PushUserId.
	PushUserSig *string `json:"PushUserSig,omitempty" name:"PushUserSig"`

	// Whiteboard parameters, such as the width, height, and background color of the whiteboard.
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// Whiteboard push timeout. Unit: seconds. Valid range: [300,259200]. Default value: 1800.
	// 
	// If no operation is performed on the whiteboard for the specified period of time, the whiteboard push service automatically stops whiteboard push.
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitempty" name:"AutoStopTimeout"`

	// Specifies whether to synchronize operations on the primary whiteboard push task to the backup task.
	AutoManageBackup *bool `json:"AutoManageBackup,omitempty" name:"AutoManageBackup"`

	// Parameters of the backup whiteboard push task.
	// 
	// After the backup parameters are specified, the whiteboard push service adds a relayed video stream. That is, there are two video streams on the whiteboard in the same room.
	Backup *WhiteboardPushBackupParam `json:"Backup,omitempty" name:"Backup"`

	// Advanced permission control parameter in TRTC. If the advanced permission control feature is enabled in TRTC, this parameter is required.
	PrivateMapKey *string `json:"PrivateMapKey,omitempty" name:"PrivateMapKey"`

	// Frame rate of the whiteboard push video. Valid range: [0,30]. Default value: 20. Unit: fps.
	VideoFPS *int64 `json:"VideoFPS,omitempty" name:"VideoFPS"`

	// Whiteboard push bitrate. Valid range: [0,2000]. Default value: 1200. Unit: kbps.
	// 
	// The bitrate specified here is a reference value. During actual push, a dynamic bitrate is used. The actual bitrate does not remain the specified value but rather fluctuates around it.
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Specifies whether to automatically record the whiteboard push when the recording mode in TRTC is set to `SubscribeStreamUserIds`.
	// 
	// Default value: `false`. If you need to enable automatic whiteboard push recording, set this parameter to `true`.
	// 
	// If the recording mode in TRTC is set to `Global Auto Recording`, ignore this parameter.
	AutoRecord *bool `json:"AutoRecord,omitempty" name:"AutoRecord"`

	// ID of the whiteboard push recording. The ID is filled in the `userdefinerecordid` parameter in the callback message after cloud recording is complete in TRTC, helping you identify the recording callback and locate the video file in VOD media resource management.
	// 
	// The value can be up to 64 bytes in length and contain letters (a-z and A-Z), digits (0-9), underscores (_), and hyphens (-).
	// 
	// After this parameter is set, automatic whiteboard push recording is enabled regardless of the value of the `AutoRecord` parameter.
	// 
	// Default RecordId generation rule:
	// urlencode(SdkAppID_RoomID_PushUserID)
	// 
	// Example:
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// Therefore: RecordId = 12345678_12345_push_user_1
	UserDefinedRecordId *string `json:"UserDefinedRecordId,omitempty" name:"UserDefinedRecordId"`

	// Specifies whether to enable automatic relay whiteboard push when the relay push mode in TRTC is set to `SubscribeRelayUserIds`.
	// 
	// Default value: `false`. If you need to enable automatic relay whiteboard push, set this parameter to `true`.
	// 
	// If the recording mode in TRTC is set to `Global Auto Relay`, ignore this parameter.
	AutoPublish *bool `json:"AutoPublish,omitempty" name:"AutoPublish"`

	// Stream ID used by TRTC for relay whiteboard push. With this ID, you can stream the user’s audio/video by using the domain name and standard streaming solution (FLV or HLS) in TRTC.
	// 
	// The value can be up to 64 bytes in length and contain letters (a-z and A-Z), digits (0-9), underscores (_), and hyphens (-).
	// 
	// After this parameter is set, automatic relay whiteboard push is enabled regardless of the value of the `AutoPublish` parameter.
	// 
	// Default StreamID generation rule:
	// urlencode(SdkAppID_RoomID_PushUserID_main)
	// 
	// Example:
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// Therefore, StreamID = 12345678_12345_push_user_1_main
	UserDefinedStreamId *string `json:"UserDefinedStreamId,omitempty" name:"UserDefinedStreamId"`

	// Internal parameter. You can ignore this parameter.
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`

	// TRTC room ID of the numeric type. Valid range: (1,4294967295).
	// 
	// If RoomId and TRTCRoomId are both specified, the value of TRTCRoomId takes priority as the room ID for TRTC whiteboard push.
	// 
	// If the TRTCRoomIdStr parameter is specified, this parameter is ignored.
	TRTCRoomId *int64 `json:"TRTCRoomId,omitempty" name:"TRTCRoomId"`

	// TRTC room ID of the string type.
	// 
	// If TRTCRoomIdStr is specified, its value takes priority as the room ID for TRTC whiteboard push.
	TRTCRoomIdStr *string `json:"TRTCRoomIdStr,omitempty" name:"TRTCRoomIdStr"`

	// IM authentication parameters.
	// If the ID of the IM application used by whiteboard messages is different from SdkAppId of the whiteboard application, specify this parameter to provide authentication information of the IM application.
	// 
	// If this parameter is specified, the whiteboard push service preferably uses the SdkAppId value in this parameter as the transmission channel for whiteboard messages. If this parameter is left empty, the SdkAppId value in the common parameters is used.
	IMAuthParam *AuthParam `json:"IMAuthParam,omitempty" name:"IMAuthParam"`

	// TRTC authentication parameters, used for room entrance authentication for TRTC push.
	// If the ID of the TRTC application to which the target TRTC room belongs is different from SdkAppId of the whiteboard application, specify this parameter to provide authentication information of the TRTC application.
	// 
	// If this parameter is specified, the whiteboard push service preferably uses the SdkAppId value in this parameter as the ID of the target TRTC application. If this parameter is left empty, the SdkAppId value in the common parameters is used.
	TRTCAuthParam *AuthParam `json:"TRTCAuthParam,omitempty" name:"TRTCAuthParam"`

	// This parameter is available only to users in the allowlist for trial.
	// 
	// TRTC room entrance mode for whiteboard push. Default value: `TRTCAppSceneVideoCall`.
	// 
	// `TRTCAppSceneVideoCall`: This scenario is most suitable when there are two or more people on a video call. The internal encoders and network protocols are optimized for video smoothness to reduce call latency and lagging.
	// `TRTCAppSceneLIVE`: This scenario is most suitable when there is only one person speaking or performing for an online audience, and occasionally multiple people interact with one another through video. The internal encoders and network protocols are optimized for performance and compatibility to deliver better performance and video clarity.
	TRTCEnterRoomMode *string `json:"TRTCEnterRoomMode,omitempty" name:"TRTCEnterRoomMode"`
}

type StartWhiteboardPushRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID of the whiteboard push task. Valid range: (1,4294967295).
	// 
	// 1. By default, the whiteboard push task uses the RoomId string as the GroupID of the IM group. For example, if the RoomId string is 1234, the GroupID of the IM group is also 1234. Message synchronization requires joining a group. Make sure that an IM group is created before the push starts. Otherwise, the push fails.
	// 2. If neither TRTCRoomId nor TRTCRoomIdStr is specified, the value of RoomId is used as the Tencent Real-Time Communication (TRTC) room ID for the whiteboard push task.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// User ID used by the whiteboard push service for entering the whiteboard room. If `IMAuthParam`and `TRTCAuthParam` are not specified, this user ID is used for operations such as logging in to the IM application, joining an IM group, and entering the room for TRTC whiteboard push.
	// The ID cannot exceed 60 bytes in length and must be an unused ID. The whiteboard push service uses the user ID to enter the room for whiteboard audio/video push. If this ID is already used in another scenario, the other scenario and whiteboard push service will conflict, affecting the pushing operation.
	PushUserId *string `json:"PushUserId,omitempty" name:"PushUserId"`

	// User's IM signature corresponding to the PushUserId.
	PushUserSig *string `json:"PushUserSig,omitempty" name:"PushUserSig"`

	// Whiteboard parameters, such as the width, height, and background color of the whiteboard.
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// Whiteboard push timeout. Unit: seconds. Valid range: [300,259200]. Default value: 1800.
	// 
	// If no operation is performed on the whiteboard for the specified period of time, the whiteboard push service automatically stops whiteboard push.
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitempty" name:"AutoStopTimeout"`

	// Specifies whether to synchronize operations on the primary whiteboard push task to the backup task.
	AutoManageBackup *bool `json:"AutoManageBackup,omitempty" name:"AutoManageBackup"`

	// Parameters of the backup whiteboard push task.
	// 
	// After the backup parameters are specified, the whiteboard push service adds a relayed video stream. That is, there are two video streams on the whiteboard in the same room.
	Backup *WhiteboardPushBackupParam `json:"Backup,omitempty" name:"Backup"`

	// Advanced permission control parameter in TRTC. If the advanced permission control feature is enabled in TRTC, this parameter is required.
	PrivateMapKey *string `json:"PrivateMapKey,omitempty" name:"PrivateMapKey"`

	// Frame rate of the whiteboard push video. Valid range: [0,30]. Default value: 20. Unit: fps.
	VideoFPS *int64 `json:"VideoFPS,omitempty" name:"VideoFPS"`

	// Whiteboard push bitrate. Valid range: [0,2000]. Default value: 1200. Unit: kbps.
	// 
	// The bitrate specified here is a reference value. During actual push, a dynamic bitrate is used. The actual bitrate does not remain the specified value but rather fluctuates around it.
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Specifies whether to automatically record the whiteboard push when the recording mode in TRTC is set to `SubscribeStreamUserIds`.
	// 
	// Default value: `false`. If you need to enable automatic whiteboard push recording, set this parameter to `true`.
	// 
	// If the recording mode in TRTC is set to `Global Auto Recording`, ignore this parameter.
	AutoRecord *bool `json:"AutoRecord,omitempty" name:"AutoRecord"`

	// ID of the whiteboard push recording. The ID is filled in the `userdefinerecordid` parameter in the callback message after cloud recording is complete in TRTC, helping you identify the recording callback and locate the video file in VOD media resource management.
	// 
	// The value can be up to 64 bytes in length and contain letters (a-z and A-Z), digits (0-9), underscores (_), and hyphens (-).
	// 
	// After this parameter is set, automatic whiteboard push recording is enabled regardless of the value of the `AutoRecord` parameter.
	// 
	// Default RecordId generation rule:
	// urlencode(SdkAppID_RoomID_PushUserID)
	// 
	// Example:
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// Therefore: RecordId = 12345678_12345_push_user_1
	UserDefinedRecordId *string `json:"UserDefinedRecordId,omitempty" name:"UserDefinedRecordId"`

	// Specifies whether to enable automatic relay whiteboard push when the relay push mode in TRTC is set to `SubscribeRelayUserIds`.
	// 
	// Default value: `false`. If you need to enable automatic relay whiteboard push, set this parameter to `true`.
	// 
	// If the recording mode in TRTC is set to `Global Auto Relay`, ignore this parameter.
	AutoPublish *bool `json:"AutoPublish,omitempty" name:"AutoPublish"`

	// Stream ID used by TRTC for relay whiteboard push. With this ID, you can stream the user’s audio/video by using the domain name and standard streaming solution (FLV or HLS) in TRTC.
	// 
	// The value can be up to 64 bytes in length and contain letters (a-z and A-Z), digits (0-9), underscores (_), and hyphens (-).
	// 
	// After this parameter is set, automatic relay whiteboard push is enabled regardless of the value of the `AutoPublish` parameter.
	// 
	// Default StreamID generation rule:
	// urlencode(SdkAppID_RoomID_PushUserID_main)
	// 
	// Example:
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// Therefore, StreamID = 12345678_12345_push_user_1_main
	UserDefinedStreamId *string `json:"UserDefinedStreamId,omitempty" name:"UserDefinedStreamId"`

	// Internal parameter. You can ignore this parameter.
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`

	// TRTC room ID of the numeric type. Valid range: (1,4294967295).
	// 
	// If RoomId and TRTCRoomId are both specified, the value of TRTCRoomId takes priority as the room ID for TRTC whiteboard push.
	// 
	// If the TRTCRoomIdStr parameter is specified, this parameter is ignored.
	TRTCRoomId *int64 `json:"TRTCRoomId,omitempty" name:"TRTCRoomId"`

	// TRTC room ID of the string type.
	// 
	// If TRTCRoomIdStr is specified, its value takes priority as the room ID for TRTC whiteboard push.
	TRTCRoomIdStr *string `json:"TRTCRoomIdStr,omitempty" name:"TRTCRoomIdStr"`

	// IM authentication parameters.
	// If the ID of the IM application used by whiteboard messages is different from SdkAppId of the whiteboard application, specify this parameter to provide authentication information of the IM application.
	// 
	// If this parameter is specified, the whiteboard push service preferably uses the SdkAppId value in this parameter as the transmission channel for whiteboard messages. If this parameter is left empty, the SdkAppId value in the common parameters is used.
	IMAuthParam *AuthParam `json:"IMAuthParam,omitempty" name:"IMAuthParam"`

	// TRTC authentication parameters, used for room entrance authentication for TRTC push.
	// If the ID of the TRTC application to which the target TRTC room belongs is different from SdkAppId of the whiteboard application, specify this parameter to provide authentication information of the TRTC application.
	// 
	// If this parameter is specified, the whiteboard push service preferably uses the SdkAppId value in this parameter as the ID of the target TRTC application. If this parameter is left empty, the SdkAppId value in the common parameters is used.
	TRTCAuthParam *AuthParam `json:"TRTCAuthParam,omitempty" name:"TRTCAuthParam"`

	// This parameter is available only to users in the allowlist for trial.
	// 
	// TRTC room entrance mode for whiteboard push. Default value: `TRTCAppSceneVideoCall`.
	// 
	// `TRTCAppSceneVideoCall`: This scenario is most suitable when there are two or more people on a video call. The internal encoders and network protocols are optimized for video smoothness to reduce call latency and lagging.
	// `TRTCAppSceneLIVE`: This scenario is most suitable when there is only one person speaking or performing for an online audience, and occasionally multiple people interact with one another through video. The internal encoders and network protocols are optimized for performance and compatibility to deliver better performance and video clarity.
	TRTCEnterRoomMode *string `json:"TRTCEnterRoomMode,omitempty" name:"TRTCEnterRoomMode"`
}

func (r *StartWhiteboardPushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartWhiteboardPushRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "PushUserId")
	delete(f, "PushUserSig")
	delete(f, "Whiteboard")
	delete(f, "AutoStopTimeout")
	delete(f, "AutoManageBackup")
	delete(f, "Backup")
	delete(f, "PrivateMapKey")
	delete(f, "VideoFPS")
	delete(f, "VideoBitrate")
	delete(f, "AutoRecord")
	delete(f, "UserDefinedRecordId")
	delete(f, "AutoPublish")
	delete(f, "UserDefinedStreamId")
	delete(f, "ExtraData")
	delete(f, "TRTCRoomId")
	delete(f, "TRTCRoomIdStr")
	delete(f, "IMAuthParam")
	delete(f, "TRTCAuthParam")
	delete(f, "TRTCEnterRoomMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartWhiteboardPushRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartWhiteboardPushResponseParams struct {
	// Push task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Result parameters of the backup task.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Backup *string `json:"Backup,omitempty" name:"Backup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartWhiteboardPushResponse struct {
	*tchttp.BaseResponse
	Response *StartWhiteboardPushResponseParams `json:"Response"`
}

func (r *StartWhiteboardPushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartWhiteboardPushResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopOnlineRecordRequestParams struct {
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the recording task to stop
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type StopOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the recording task to stop
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopOnlineRecordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *StopOnlineRecordResponseParams `json:"Response"`
}

func (r *StopOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWhiteboardPushRequestParams struct {
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the whiteboard push task to be stopped.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type StopWhiteboardPushRequest struct {
	*tchttp.BaseRequest
	
	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the whiteboard push task to be stopped.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopWhiteboardPushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWhiteboardPushRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopWhiteboardPushRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWhiteboardPushResponseParams struct {
	// Parameters of the backup task.
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Backup *string `json:"Backup,omitempty" name:"Backup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopWhiteboardPushResponse struct {
	*tchttp.BaseResponse
	Response *StopWhiteboardPushResponseParams `json:"Response"`
}

func (r *StopWhiteboardPushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWhiteboardPushResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StreamControl struct {
	// Video stream ID
	// Description of the possible video stream ID values:
	// 1. `tic_record_user`: the whiteboard video stream
	// 2. `tic_substream`: the auxiliary video stream
	// 3. Specific user ID: the video stream of the specified user
	// 
	// The actual recording uses the prefix match of the video stream ID. The real stream becomes the specified stream once its ID prefix matches with the stream ID.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Whether to disable recording over the stream.
	// 
	// true: does not record this stream. This stream will not be included in the final recording file.
	// false: records this stream. This stream will be included in the final recording file.
	// 
	// Default value: false
	DisableRecord *bool `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// Whether to disable the audio recording of the stream.
	// 
	// true: does not record the audio of the stream. In the final recording file, this stream will be soundless.
	// false: the stream has both video and audio recording.
	// 
	// Default value: false
	DisableAudio *bool `json:"DisableAudio,omitempty" name:"DisableAudio"`

	// Whether to only record low-resolution stream videos.
	// 
	// true: records only low-resolution videos. In this case, please make sure that the client pushes low-resolution videos upstream. Otherwise, the recorded video may be black. 
	// false: records only high-resolution videos.
	// 
	// Default value: false
	PullSmallVideo *bool `json:"PullSmallVideo,omitempty" name:"PullSmallVideo"`
}

type StreamLayout struct {
	// Stream layout configuration
	LayoutParams *LayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// Video stream ID
	// Description of the possible video stream ID values:
	// 1. tic_record_user: the current picture is used to display the whiteboard video stream.
	// 2. tic_substream: the current picture is used to display the auxiliary video stream.
	// 3. Specific user ID: the current picture is used to display the video stream of a specific user.
	// 4.Left empty: the current picture is vacant for new video stream.
	InputStreamId *string `json:"InputStreamId,omitempty" name:"InputStreamId"`

	// Background color in RGB format, such as "#FF0000" for red. The default color is black. 
	BackgroundColor *string `json:"BackgroundColor,omitempty" name:"BackgroundColor"`

	// Video filling mode.
	// 
	// 0: self-adaption mode. Scales the video proportionally to completely display it in the specified area. In this mode, there may be black bars.
	// 1: full-screen mode. Scales the video to make it fill the entire specified area. In this mode, no black bars will appear, but the video may not be displayed fully.
	FillMode *int64 `json:"FillMode,omitempty" name:"FillMode"`
}

type Tag struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TimeValue struct {
	// Unix timestamp in seconds.
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// Current value of the queried metric.
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TranscodeTaskResult struct {
	// Transcoding result URL.
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// Target resolution.
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// Title (usually the document name).
	Title *string `json:"Title,omitempty" name:"Title"`

	// Number of transcoded pages.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// URL prefix of the thumbnail. If the URL prefix is `http://example.com/g0jb42ps49vtebjshilb/`, the thumbnail URL for the first page of the dynamically transcoded PowerPoint file is
	// `http://example.com/g0jb42ps49vtebjshilb/1.jpg`, and so on.
	// 
	// If the document transcoding request carries the ThumbnailResolution parameter and the transcoding type is dynamic transcoding, this parameter is not null. In other cases, this parameter is null.
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" name:"ThumbnailUrl"`

	// Resolution of the thumbnail generated for dynamic transcoding
	ThumbnailResolution *string `json:"ThumbnailResolution,omitempty" name:"ThumbnailResolution"`

	// URL for downloading the transcoded and compressed file. If `CompressFileType` carried in the document transcoding request is null or is not a supported compression format, this parameter is null.
	CompressFileUrl *string `json:"CompressFileUrl,omitempty" name:"CompressFileUrl"`

	// Task execution error code.
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// Task execution error message.
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

type TranscodeTaskSearchResult struct {
	// Creation time of the task.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Unique task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Current task status.
	// - QUEUED: Queuing for transcoding.
	// - PROCESSING: Transcoding in progress.
	// - FINISHED: Transcoding finished.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Original name of the transcoded document.
	OriginalFilename *string `json:"OriginalFilename,omitempty" name:"OriginalFilename"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Transcoding task result.
	Result *TranscodeTaskResult `json:"Result,omitempty" name:"Result"`

	// Specifies whether the transcoding is static.
	IsStatic *bool `json:"IsStatic,omitempty" name:"IsStatic"`
}

type UsageDataItem struct {
	// Start date in the format of YYYY-MM-DD.
	Time *string `json:"Time,omitempty" name:"Time"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Subproduct usage information, which is consistent with the corresponding request parameters.
	// - sp_tiw_board: The duration of use of whiteboards.
	// - sp_tiw_dt: The number of pages dynamically transcoded.
	// - sp_tiw_st: The number of pages statically transcoded.
	// - sp_tiw_ric: The duration of real-time recording.
	SubProduct *string `json:"SubProduct,omitempty" name:"SubProduct"`

	// Usage values.
	// - The unit of sp_tiw_dt and sp_tiw_st is pages.
	// - The unit of sp_tiw_board and sp_tiw_ric is minutes.
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type UserListItem struct {
	// User ID in the room.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The first time when the user ID appeared during the queried period, which is a Unix timestamp in milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// The last time when the user ID appeared during the queried period, which is a Unix timestamp in milliseconds.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type VideoInfo struct {
	// Video playback start time, in milliseconds
	VideoPlayTime *int64 `json:"VideoPlayTime,omitempty" name:"VideoPlayTime"`

	// Video size, in bytes
	VideoSize *int64 `json:"VideoSize,omitempty" name:"VideoSize"`

	// Video format
	VideoFormat *string `json:"VideoFormat,omitempty" name:"VideoFormat"`

	// Video playback duration, in milliseconds
	VideoDuration *int64 `json:"VideoDuration,omitempty" name:"VideoDuration"`

	// Video file URL
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// Video file ID
	VideoId *string `json:"VideoId,omitempty" name:"VideoId"`

	// Video stream type - 0: camera video - 1: screen-sharing video - 2: whiteboard video - 3: mixed stream video - 4: audio-only (mp3)
	VideoType *int64 `json:"VideoType,omitempty" name:"VideoType"`

	// ID of the user to which the camera video or screen-sharing video belongs (whiteboard video: null, mixed stream video: tic_mixstream_<Room ID>_<Mixed stream layout type>, auxiliary video: tic_substream_user ID)
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Width of the video resolution.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Height of the video resolution.
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type Whiteboard struct {
	// Whiteboard video width in the real-time recording result. The value must be equal to or greater than 2. Default value: 1280.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Whiteboard video height in the real-time recording result. The value must be equal to or greater than 2. Default value: 960.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Whiteboard initialization parameter, which is passed through to the whiteboard SDK
	InitParam *string `json:"InitParam,omitempty" name:"InitParam"`
}

type WhiteboardApplicationConfig struct {
	// Task type.
	// 
	// recording: Real-time recording.
	// transcode: Document transcoding.
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Bucket name.
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// Region of the bucket.
	BucketLocation *string `json:"BucketLocation,omitempty" name:"BucketLocation"`

	// Resource prefix of the bucket.
	BucketPrefix *string `json:"BucketPrefix,omitempty" name:"BucketPrefix"`

	// Destination CDN domain name.
	ResultDomain *string `json:"ResultDomain,omitempty" name:"ResultDomain"`

	// Callback URL.
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// Callback authentication key.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// IM admin user ID.
	AdminUserId *string `json:"AdminUserId,omitempty" name:"AdminUserId"`

	// IM admin user signature.
	AdminUserSig *string `json:"AdminUserSig,omitempty" name:"AdminUserSig"`
}

type WhiteboardPushBackupParam struct {
	// User ID used by the whiteboard push service for entering a room,
	// The ID must be an unused ID in the SDK. The whiteboard push service uses the user ID to enter the room for whiteboard push. If this ID is already used in the SDK, the SDK and recording service will conflict, affecting the pushing operation.
	PushUserId *string `json:"PushUserId,omitempty" name:"PushUserId"`

	// Signature corresponding to the PushUserId ID.
	PushUserSig *string `json:"PushUserSig,omitempty" name:"PushUserSig"`
}

type WhiteboardPushResult struct {
	// `AUTO`: Pushing automatically stops. `USER_CALL`: The API for stopping pushing is called.
	FinishReason *string `json:"FinishReason,omitempty" name:"FinishReason"`

	// Number of exceptions.
	ExceptionCnt *int64 `json:"ExceptionCnt,omitempty" name:"ExceptionCnt"`

	// Room ID.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// IM group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Actual push start time.
	PushStartTime *int64 `json:"PushStartTime,omitempty" name:"PushStartTime"`

	// Push end time.
	PushStopTime *int64 `json:"PushStopTime,omitempty" name:"PushStopTime"`

	// IM timestamp corresponding to the first frame of the whiteboard push video. The timestamp is used for time synchronization between IM messages and the whiteboard push video during playback.
	IMSyncTime *int64 `json:"IMSyncTime,omitempty" name:"IMSyncTime"`

	// Task execution error code.
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// Error message.
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

type WhiteboardPushTaskSearchResult struct {
	// Unique task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Status of the whiteboard push task.
	// - PREPARED: Push in preparation.
	// - PUSHING: Pushing in progress.
	// - STOPPED: Pushing stopped.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Room ID of the whiteboard push task.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// Creation time of the task.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// SdkAppId of the whiteboard application.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Whiteboard push result.
	Result *WhiteboardPushResult `json:"Result,omitempty" name:"Result"`

	// User ID of the whiteboard push task.
	PushUserId *string `json:"PushUserId,omitempty" name:"PushUserId"`
}