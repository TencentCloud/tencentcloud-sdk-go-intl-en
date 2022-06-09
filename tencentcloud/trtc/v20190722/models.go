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

package v20190722

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AudioParams struct {

	// The audio sample rate.
	// 1: 48000 Hz (default)
	// 2: 44100 Hz
	// 3: 16000 Hz
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// The number of sound channels.
	// 1: Mono-channel
	// 2: Dual-channel (default)
	Channel *uint64 `json:"Channel,omitempty" name:"Channel"`

	// The audio bitrate (bps). Value range: [32000, 128000]. Default: 64000.
	BitRate *uint64 `json:"BitRate,omitempty" name:"BitRate"`
}

type CloudStorage struct {

	// The cloud storage provider.
	// 0: Tencent Cloud COS. The storage services of other providers are not supported currently.
	Vendor *uint64 `json:"Vendor,omitempty" name:"Vendor"`

	// The region of cloud storage.
	Region *string `json:"Region,omitempty" name:"Region"`

	// The storage bucket.
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// The access_key of the cloud storage account.
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// The secret_key of the cloud storage account.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// The bucket to save data, which is an array of strings that can contain letters (a-z and A-Z), numbers (0-9), underscores (_), and hyphens (-). For example, if the value of this parameter is `["prefix1", "prefix2"]`, the recording file `xxx.m3u8` will be saved as `prefix1/prefix2/TaskId/xxx.m3u8`.
	FileNamePrefix []*string `json:"FileNamePrefix,omitempty" name:"FileNamePrefix"`
}

type CloudVod struct {

	// The Tencent Cloud VOD parameters.
	TencentVod *TencentVod `json:"TencentVod,omitempty" name:"TencentVod"`
}

type CreateCloudRecordingRequest struct {
	*tchttp.BaseRequest

	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The [room ID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// The [user ID](https://intl.cloud.tencent.com/document/product/647/37714) of the recording robot in the TRTC room, which cannot be the same as a user ID already in use. We recommend you include this user ID in the room ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The signature (similar to login password) required for the recording robot to enter the room. For information on how to calculate the signature, see [What is UserSig?](https://intl.cloud.tencent.com/document/product/647/38104). |
	UserSig *string `json:"UserSig,omitempty" name:"UserSig"`

	// The on-cloud recording parameters.
	RecordParams *RecordParams `json:"RecordParams,omitempty" name:"RecordParams"`

	// The cloud storage parameters.
	StorageParams *StorageParams `json:"StorageParams,omitempty" name:"StorageParams"`

	// The type of the TRTC room ID, which must be the same as the ID type of the room whose streams are recorded.
	// 0: String
	// 1: 32-bit integer (default)
	RoomIdType *uint64 `json:"RoomIdType,omitempty" name:"RoomIdType"`

	// The stream mixing parameters, which are valid if the mixed-stream recording mode is used.
	MixTranscodeParams *MixTranscodeParams `json:"MixTranscodeParams,omitempty" name:"MixTranscodeParams"`

	// The layout parameters, which are valid if the mixed-stream recording mode is used.
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitempty" name:"MixLayoutParams"`

	// The amount of time (in hours) during which API requests can be made after recording starts. Calculation starts when a recording task is started (when the recording task ID is returned). Once the period elapses, the query, modification, and stop recording APIs can no longer be called, but the recording task will continue. The default value is `72` (three days), and the maximum and minimum values allowed are `720` (30 days) and `6` respectively. If you do not set this parameter, the query, modification, and stop recording APIs can be called within 72 hours after recording starts.
	ResourceExpiredHour *uint64 `json:"ResourceExpiredHour,omitempty" name:"ResourceExpiredHour"`

	// The permission ticket for a TRTC room. This parameter is required if advanced permission control is enabled in the console, in which case the TRTC backend will verify users’ [PrivateMapKey](https://intl.cloud.tencent.com/document/product/647/32240?from_cn_redirect=1), which include an encrypted room ID and permission bit list. A user providing only `UserSig` and not `PrivateMapKey` will be unable to enter the room.
	PrivateMapKey *string `json:"PrivateMapKey,omitempty" name:"PrivateMapKey"`
}

func (r *CreateCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserId")
	delete(f, "UserSig")
	delete(f, "RecordParams")
	delete(f, "StorageParams")
	delete(f, "RoomIdType")
	delete(f, "MixTranscodeParams")
	delete(f, "MixLayoutParams")
	delete(f, "ResourceExpiredHour")
	delete(f, "PrivateMapKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The task ID assigned by the recording service, which uniquely identifies a recording process and becomes invalid after a recording task ends. After a recording task starts, if you want to perform other actions on the task, you need to specify the task ID when making API requests.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudRecordingRequest struct {
	*tchttp.BaseRequest

	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The task ID assigned by the recording service, which uniquely identifies a recording process and becomes invalid after a recording task ends.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudRecordingRequest struct {
	*tchttp.BaseRequest

	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique ID of the recording task.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The status of the on-cloud recording task.
	// Idle: The task is idle.
	// InProgress: The task is in progress.
	// Exited: The task is being ended.
		Status *string `json:"Status,omitempty" name:"Status"`

		// The information of the recording files.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		StorageFileList []*StorageFile `json:"StorageFileList,omitempty" name:"StorageFileList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DismissRoomByStrRoomIdRequest struct {
	*tchttp.BaseRequest

	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DismissRoomByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DismissRoomByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DismissRoomByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DismissRoomByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DismissRoomRequest struct {
	*tchttp.BaseRequest

	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room number.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DismissRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DismissRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DismissRoomResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DismissRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MixLayout struct {

	// The Y axis of the window’s top-left corner. Value range: [0, 1920]. The value cannot be larger than the canvas height.
	Top *uint64 `json:"Top,omitempty" name:"Top"`

	// The X axis of the window’s top-left corner. Value range: [0, 1920]. The value cannot be larger than the canvas width.
	Left *uint64 `json:"Left,omitempty" name:"Left"`

	// The relative width of the window. Value range: [0, 1920]. The sum of the values of this parameter and `Left` cannot exceed the canvas width.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// The relative height of the window. Value range: [0, 1920]. The sum of the values of this parameter and `Top` cannot exceed the canvas height.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// The user ID (string) of the anchor whose video is shown in the window. If you do not set this parameter, anchors’ videos will be shown in their room entry sequence.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The degree of transparency of the canvas. Value range: [0, 255]. 0 means fully opaque, and 255 means fully transparent.
	Alpha *uint64 `json:"Alpha,omitempty" name:"Alpha"`

	// 0: Stretch. In this mode, the image is stretched to fill the space available. The whole image is visible after scaling. However, if the original aspect ratio is different from the target, the image may be distorted.
	// 
	// 1: Crop (default). In this mode, if the original aspect ratio is different from the target, the image will be cropped according to the target before being stretched to fill the space available. The image will not be distorted.
	// 
	// 2: Blank. This mode stretches the image while keeping its original aspect ratio. If the original aspect ratio is different from the target, there may be blank spaces to the top and bottom or to the left and right of the window.
	// 
	// 3: Smart stretch. This mode is similar to the crop mode, except that it restricts cropping to 20% of the image’s width or height at most.
	RenderMode *uint64 `json:"RenderMode,omitempty" name:"RenderMode"`

	// The type of the stream subscribed to.
	// 0: Primary stream (default)
	// 1: Substream
	MediaId *uint64 `json:"MediaId,omitempty" name:"MediaId"`

	// The image layer. 0 is the default value and means the bottommost layer.
	ImageLayer *uint64 `json:"ImageLayer,omitempty" name:"ImageLayer"`

	// The download URL of the background image for a window. The image must be in JPG or PNG format and cannot be larger than 5 MB. If the image’s aspect ratio is different from that of the window, the image will be rendered according to the value of `RenderMode`.
	SubBackgroundImage *string `json:"SubBackgroundImage,omitempty" name:"SubBackgroundImage"`
}

type MixLayoutParams struct {

	// Layout mode:
	// 1: Floating
	// 2: Screen sharing
	// 3: Grid (default)
	// 4: Custom
	// 
	// Floating: By default, the video of the first anchor (you can also specify an anchor) who enters the room is scaled to fill the screen. When other anchors enter the room, their videos appear smaller and are superimposed over the large video from left to right starting from the bottom of the canvas according to their room entry sequence. If the total number of videos is 17 or less, there will be four windows in each row (4 x 4); if it is greater than 17, there will be five windows in each row (5 x 5). Up to 25 videos can be displayed. A user who publishes only audio will still be displayed in one window.
	// 
	// Screen sharing: The video of a specified anchor occupies a larger part of the canvas on the left side (if you do not specify an anchor, the left window will display the canvas background). The videos of other anchors are smaller and are positioned on the right side. If the total number of videos is 17 or less, the small videos are positioned from top to bottom in up to two columns on the right side, with eight videos per column at most. If there are more than 17 videos, the additional videos are positioned at the bottom of the canvas from left to right. Up to 25 videos can be displayed. A user who publishes only audio will still be displayed in one window.
	// 
	// Grid: The videos of anchors are scaled and positioned automatically according to the total number of anchors in a room. Each video has the same size. Up to 25 videos can be displayed.
	// 
	// Custom: Specify the layout of videos by using the `MixLayoutList` parameter.
	MixLayoutMode *uint64 `json:"MixLayoutMode,omitempty" name:"MixLayoutMode"`

	// The custom layout details. This parameter is valid if `MixLayoutMode` is set to `4`. Up to 25 videos can be displayed.
	MixLayoutList []*MixLayout `json:"MixLayoutList,omitempty" name:"MixLayoutList"`

	// The background color, which is a hexadecimal value (starting with "#", followed by the color value) converted from an 8-bit RGB value. For example, the RGB value of orange is `R:255 G:165 B:0`, and its hexadecimal value is `#FFA500`. The default color is black.
	BackGroundColor *string `json:"BackGroundColor,omitempty" name:"BackGroundColor"`

	// The user whose video is displayed in the big window. This parameter is valid if `MixLayoutMode` is set to `1` (floating) or `2` (screen sharing). If it is left empty, the first anchor entering the room is displayed in the big window in the floating mode and the canvas background is displayed in the screen sharing mode.
	MaxResolutionUserId *string `json:"MaxResolutionUserId,omitempty" name:"MaxResolutionUserId"`

	// The stream type.
	// 0: Primary stream (default)
	// 1: Substream (screen sharing stream)
	// This parameter specifies the type of the stream displayed in the big window. If it appears in `MixLayoutList`, it indicates the type of the stream of a specified user.
	MediaId *uint64 `json:"MediaId,omitempty" name:"MediaId"`

	// The download URL of the background image for the canvas, which must be in JPG or PNG format and cannot be larger than 5 MB.
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" name:"BackgroundImageUrl"`

	// `1` means to use placeholders, and `0` (default) means to not use placeholders. If this parameter is set to `1`, when a user is not publishing video, a placeholder image will be displayed in the window reserved for the user.
	PlaceHolderMode *uint64 `json:"PlaceHolderMode,omitempty" name:"PlaceHolderMode"`

	// The render mode to use when the aspect ratio of a video is different from that of the window. This parameter is defined the same as `RenderMode` in `MixLayoufList`.
	BackgroundImageRenderMode *uint64 `json:"BackgroundImageRenderMode,omitempty" name:"BackgroundImageRenderMode"`

	// The download URL of the default background image for a window. The image must be in JPG or PNG format and cannot be larger than 5 MB. If the image’s aspect ratio is different from that of the window, the image will be rendered according to the value of `RenderMode`.
	DefaultSubBackgroundImage *string `json:"DefaultSubBackgroundImage,omitempty" name:"DefaultSubBackgroundImage"`

	// The watermark layout. Up to 25 watermarks are supported.
	WaterMarkList []*WaterMark `json:"WaterMarkList,omitempty" name:"WaterMarkList"`
}

type MixTranscodeParams struct {

	// The video transcoding parameters for recording. If you set this parameter, you must specify all its fields. If you do not set it, the default will be used.
	VideoParams *VideoParams `json:"VideoParams,omitempty" name:"VideoParams"`

	// The audio transcoding parameters for recording. If you set this parameter, you must specify all its fields. If you do not set it, the default will be used.
	AudioParams *AudioParams `json:"AudioParams,omitempty" name:"AudioParams"`
}

type ModifyCloudRecordingRequest struct {
	*tchttp.BaseRequest

	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The new stream mixing layout to use.
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitempty" name:"MixLayoutParams"`

	// The allowlist/blocklist for stream subscription.
	SubscribeStreamUserIds *SubscribeStreamUserIds `json:"SubscribeStreamUserIds,omitempty" name:"SubscribeStreamUserIds"`
}

func (r *ModifyCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	delete(f, "MixLayoutParams")
	delete(f, "SubscribeStreamUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The task ID assigned by the recording service, which uniquely identifies a recording process and becomes invalid after a recording task ends.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordParams struct {

	// The recording mode.
	// 1: Single-stream recording. Records the audio and video of each subscribed user (`UserId`) in a room and saves the recording files (M3U8/TS) to the cloud.
	// 2: Mixed-stream recording. Mixes the audios and videos of subscribed users (`UserId`) in a room, records the mixed stream, and saves the recording files (M3U8/TS) to the cloud.
	RecordMode *uint64 `json:"RecordMode,omitempty" name:"RecordMode"`

	// The time period (seconds) to wait after there are no anchors in a room to stop recording automatically. The value cannot be smaller than 5 or larger than 86400 (24 hours). Default value: 30.
	MaxIdleTime *uint64 `json:"MaxIdleTime,omitempty" name:"MaxIdleTime"`

	// The media type of the streams to record.
	// 0: Audio and video streams (default)
	// 1: Audio streams only
	// 2: Video streams only
	StreamType *uint64 `json:"StreamType,omitempty" name:"StreamType"`

	// The allowlist/blocklist for stream subscription.
	SubscribeStreamUserIds *SubscribeStreamUserIds `json:"SubscribeStreamUserIds,omitempty" name:"SubscribeStreamUserIds"`

	// The format of recording files. 0 (default): HLS; 1: HLS + MP4 (recorded in HLS and converted to MP4).
	OutputFormat *uint64 `json:"OutputFormat,omitempty" name:"OutputFormat"`
}

type RemoveUserByStrRoomIdRequest struct {
	*tchttp.BaseRequest

	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// List of up to 10 users to be removed
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *RemoveUserByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveUserByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveUserByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveUserRequest struct {
	*tchttp.BaseRequest

	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room number.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// List of up to 10 users to be removed.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *RemoveUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageFile struct {

	// The user whose stream is recorded into the file. In the mixed-stream recording mode, this parameter will be empty.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The filename.
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// The type of the media recorded.
	// video
	// audio
	// audio_video
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TrackType *string `json:"TrackType,omitempty" name:"TrackType"`

	// The start time (Unix timestamp) of the recording file.
	BeginTimeStamp *uint64 `json:"BeginTimeStamp,omitempty" name:"BeginTimeStamp"`
}

type StorageParams struct {

	// The cloud storage information.
	CloudStorage *CloudStorage `json:"CloudStorage,omitempty" name:"CloudStorage"`

	// The VOD information.
	CloudVod *CloudVod `json:"CloudVod,omitempty" name:"CloudVod"`
}

type SubscribeStreamUserIds struct {

	// The allowlist for audio subscription. For example, `["1", "2", "3"]` means to only subscribe to the audios of users 1, 2, and 3, and ["1.*$"] means to only subscribe to the audios of users whose ID prefix is `1`. If this parameter is left empty, the audios of all anchors in the room will be received. The array can contain at most 32 elements.
	SubscribeAudioUserIds []*string `json:"SubscribeAudioUserIds,omitempty" name:"SubscribeAudioUserIds"`

	// The blocklist for audio subscription. For example, `["1", "2", "3"]` means to not subscribe to the audios of users 1, 2, and 3, and `["1.*$"]` means to not subscribe to users whose ID prefix is `1`. If this parameter is left empty, the audios of all anchors in the room will be received. The array can contain at most 32 elements.
	UnSubscribeAudioUserIds []*string `json:"UnSubscribeAudioUserIds,omitempty" name:"UnSubscribeAudioUserIds"`

	// The allowlist for video subscription. For example, `["1", "2", "3"]` means to only subscribe to the videos of users 1, 2, and 3, and `["1.*$"]` means to only subscribe to the videos of users whose ID prefix is `1`. If this parameter is left empty, the videos of all anchors in the room will be received. The array can contain at most 32 elements.
	SubscribeVideoUserIds []*string `json:"SubscribeVideoUserIds,omitempty" name:"SubscribeVideoUserIds"`

	// The blocklist for video subscription. For example, `["1", "2", "3"]` means to not subscribe to the videos of users 1, 2, and 3, and `["1.*$"]` means to not subscribe to the videos of users whose ID prefix is `1`. If this parameter is left empty, the videos of all anchors in the room will be received. The array can contain at most 32 elements.
	UnSubscribeVideoUserIds []*string `json:"UnSubscribeVideoUserIds,omitempty" name:"UnSubscribeVideoUserIds"`
}

type TencentVod struct {

	// The operation to perform on the media uploaded. The value of this parameter is the name of a task flow template. You can create a custom task flow template in Tencent Cloud VOD.
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// The expiration time of the media file, which is a time period (seconds) from the current time. For example, `86400` means to save the media file for one day. To save the file permanently, set this parameter to `0`.
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// The storage region. Set this parameter if you have special requirements on the storage region.
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// The category ID, which is returned after you create a category by calling an API. You can use categories to manage media files.
	// The default value is `0`, which means others.
	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

	// The VOD subapplication ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// The task flow context, which is passed through after the task is completed.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The upload context, which is passed through after upload is completed.
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`
}

type VideoParams struct {

	// The video width in pixels. The value of this parameter cannot be larger than 1920, and the result of multiplying `Width` and `Height` cannot exceed 1920 x 1080. The default value is `360`.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// The video height in pixels. The value of this parameter cannot be larger than 1920, and the result of multiplying `Width` and `Height` cannot exceed 1920 x 1080. The default value is `640`.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// The video frame rate. Value range: [1, 60]. Default: 15.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// The video bitrate (bps). Value range: [64000, 8192000]. Default: 550000.
	BitRate *uint64 `json:"BitRate,omitempty" name:"BitRate"`

	// The keyframe interval (seconds). Default value: 10.
	Gop *uint64 `json:"Gop,omitempty" name:"Gop"`
}

type WaterMark struct {

	// The watermark type. 0 (default): image; 1: text (not supported yet).
	WaterMarkType *uint64 `json:"WaterMarkType,omitempty" name:"WaterMarkType"`

	// The information of watermark images. This parameter is required if the watermark type is image.
	WaterMarkImage *WaterMarkImage `json:"WaterMarkImage,omitempty" name:"WaterMarkImage"`
}

type WaterMarkImage struct {

	// The download URLs of the watermark images, which must be in JPG or PNG format and cannot be larger than 5 MB.
	WaterMarkUrl *string `json:"WaterMarkUrl,omitempty" name:"WaterMarkUrl"`

	// The Y axis of the image's top-left corner. Value range: [0, 2560]. The value cannot be larger than the canvas height.
	Top *uint64 `json:"Top,omitempty" name:"Top"`

	// The X axis of the image’s top-left corner. Value range: [0, 2560]. The value cannot be larger than the canvas width.
	Left *uint64 `json:"Left,omitempty" name:"Left"`

	// The relative width of the image. Value range: [0, 2560]. The sum of the values of this parameter and `Left` cannot exceed the canvas width.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// The relative height of the image. Value range: [0, 2560]. The sum of the values of this parameter and `Top` cannot exceed the canvas height.
	Height *uint64 `json:"Height,omitempty" name:"Height"`
}
