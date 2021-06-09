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

type CreateTranscodeRequest struct {
	*tchttp.BaseRequest

	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Address of the file for transcoding
	Url *string `json:"Url,omitempty" name:"Url"`

	// Whether the PowerPoint file is static. The default value is False.
	// If IsStaticPPT is False, documents with the .ppt or .pptx extension will be dynamically transcoded to HTML5 pages, and documents with other extensions will be statically transcoded to images. If IsStaticPPT is True, documents with any extensions will be statically transcoded to images.
	IsStaticPPT *bool `json:"IsStaticPPT,omitempty" name:"IsStaticPPT"`

	// Minimum resolution of the transcoded document. If no value or null is specified for it or the resolution format is invalid, the original document resolution is used.
	// 
	//  
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
}

func (r *CreateTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTranscodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of the document transcoding task, which is used to query the task progress and transcoding result
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomLayout struct {

	// Mixed stream canvas parameter
	Canvas *Canvas `json:"Canvas,omitempty" name:"Canvas"`

	// Stream layout. The layout of each stream cannot exceed the canvas area.
	InputStreamList []*StreamLayout `json:"InputStreamList,omitempty" name:"InputStreamList"`
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

// It is highly **NOT** recommended to use this function
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

type DescribeOnlineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Callback address of the real-time recording event. If no callback address is set, this field is null.
		Callback *string `json:"Callback,omitempty" name:"Callback"`

		// Authentication key of the real-time recording callback
		CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOnlineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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

type DescribeOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recording stop reason
	// - AUTO: recording automatically stops because no upstream audio/video or whiteboard operation occurs in the room for a long time.
	// - USER_CALL: the API for stopping recording is called.
	// - EXCEPTION: an exception occurred during recording.
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

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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

type DescribeTranscodeCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Document transcoding callback address
		Callback *string `json:"Callback,omitempty" name:"Callback"`

		// Authentication key of the document transcoding callback
		CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTranscodeCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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

type DescribeTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type OmittedDuration struct {

	// Offset of the paused time in the spliced video, in milliseconds
	VideoTime *int64 `json:"VideoTime,omitempty" name:"VideoTime"`

	// Recording pause timestamp, in milliseconds
	PauseTime *int64 `json:"PauseTime,omitempty" name:"PauseTime"`

	// Recording resumption timestamp, in milliseconds
	ResumeTime *int64 `json:"ResumeTime,omitempty" name:"ResumeTime"`
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

// It is highly **NOT** recommended to use this function
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

type PauseOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

// It is highly **NOT** recommended to use this function
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

type ResumeOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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

type SetOnlineRecordCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetOnlineRecordCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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

type SetOnlineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetOnlineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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

type SetTranscodeCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetTranscodeCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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

type SetTranscodeCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetTranscodeCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartOnlineRecordRequest struct {
	*tchttp.BaseRequest

	// SdkAppId of the customer
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// ID of the room for recording. Value range: (1, 4294967295)
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// User ID used by the real-time recording service for entering a room. Its format is `tic_record_user_${RoomId}_${Random}`, where `${RoomId}` indicates the ID of the room for recording and `${Random}` is a random string.
	// The ID must be an unused ID in the SDK. The real-time recording service uses the user ID to enter the room for audio, video, and whiteboard recording. If this ID is already used in the SDK, the SDK and recording service will conflict, affecting the recording operation.
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
}

func (r *StartOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of the real-time recording task
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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

type StopOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopOnlineRecordResponse) FromJsonString(s string) error {
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

	// Whiteboard video width in the real-time recording result. The default value is 1280.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Whiteboard video height in the real-time recording result. The default value is 960.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Whiteboard initialization parameter, which is passed through to the whiteboard SDK
	InitParam *string `json:"InitParam,omitempty" name:"InitParam"`
}
