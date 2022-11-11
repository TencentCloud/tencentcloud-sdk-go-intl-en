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

type AgentParams struct {
	// The [user ID](https://intl.cloud.tencent.com/document/product/647/37714) of the relaying robot in the TRTC room, which cannot be the same as a user ID already in use. We recommend you include the room ID in this user ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The signature (similar to a login password) required for the relaying robot to enter the room. For information on how to calculate the signature, see [What is UserSig?](https://intl.cloud.tencent.com/document/product/647/38104). |
	UserSig *string `json:"UserSig,omitempty" name:"UserSig"`

	// The timeout period (seconds) for relaying to stop automatically after all the users whose streams are mixed leave the room. The value cannot be smaller than 5 or larger than 86400 (24 hours). Default value: 30.
	MaxIdleTime *uint64 `json:"MaxIdleTime,omitempty" name:"MaxIdleTime"`
}

type AudioEncode struct {
	// The audio sample rate (Hz). Valid values: 48000, 44100, 32000, 24000, 16000, 8000.
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// The number of sound channels. Valid values: 1 (mono), 2 (dual).
	Channel *uint64 `json:"Channel,omitempty" name:"Channel"`

	// The audio bitrate (Kbps). Value range: 8-500.
	BitRate *uint64 `json:"BitRate,omitempty" name:"BitRate"`

	// The audio codec. Valid values: 0 (LC-AAC), 1 (HE-AAC), 2 (HE-AACv2). The default value is 0. If this parameter is set to 2, `Channel` must be 2. If it is set to 1 or 2, `SampleRate` can only be 48000, 44100, 32000, 24000, or 16000.
	Codec *uint64 `json:"Codec,omitempty" name:"Codec"`
}

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

// Predefined struct for user
type CreateCloudRecordingRequestParams struct {
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The [room ID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// The [user ID](https://www.tencentcloud.com/document/product/647/37714#userid) of the recording robot in the TRTC room, which cannot be identical to the user IDs of anchors in the room or other recording robots. To distinguish this user ID from others, we recommend you include the room ID in the user ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The signature (similar to a login password) required for the recording robot to enter the room. Each user ID corresponds to a signature. For information on how to calculate the signature, see [What is UserSig?](https://intl.cloud.tencent.com/document/product/647/38104).
	UserSig *string `json:"UserSig,omitempty" name:"UserSig"`

	// The on-cloud recording parameters.
	RecordParams *RecordParams `json:"RecordParams,omitempty" name:"RecordParams"`

	// The cloud storage information of the recording file. Currently, you can only save recording files to Tencent Cloud VOD.
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

type CreateCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The [room ID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// The [user ID](https://www.tencentcloud.com/document/product/647/37714#userid) of the recording robot in the TRTC room, which cannot be identical to the user IDs of anchors in the room or other recording robots. To distinguish this user ID from others, we recommend you include the room ID in the user ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The signature (similar to a login password) required for the recording robot to enter the room. Each user ID corresponds to a signature. For information on how to calculate the signature, see [What is UserSig?](https://intl.cloud.tencent.com/document/product/647/38104).
	UserSig *string `json:"UserSig,omitempty" name:"UserSig"`

	// The on-cloud recording parameters.
	RecordParams *RecordParams `json:"RecordParams,omitempty" name:"RecordParams"`

	// The cloud storage information of the recording file. Currently, you can only save recording files to Tencent Cloud VOD.
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

// Predefined struct for user
type CreateCloudRecordingResponseParams struct {
	// The task ID assigned by the recording service, which uniquely identifies a recording process and becomes invalid after a recording task ends. After a recording task starts, if you want to perform other actions on the task, you need to specify the task ID when making API requests.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudRecordingResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteCloudRecordingRequestParams struct {
	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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

// Predefined struct for user
type DeleteCloudRecordingResponseParams struct {
	// The task ID assigned by the recording service, which uniquely identifies a recording process and becomes invalid after a recording task ends.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudRecordingResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCloudRecordingRequestParams struct {
	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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

// Predefined struct for user
type DescribeCloudRecordingResponseParams struct {
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
}

type DescribeCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudRecordingResponseParams `json:"Response"`
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

// Predefined struct for user
type DismissRoomByStrRoomIdRequestParams struct {
	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
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

// Predefined struct for user
type DismissRoomByStrRoomIdResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DismissRoomByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *DismissRoomByStrRoomIdResponseParams `json:"Response"`
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

// Predefined struct for user
type DismissRoomRequestParams struct {
	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room number.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
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

// Predefined struct for user
type DismissRoomResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DismissRoomResponse struct {
	*tchttp.BaseResponse
	Response *DismissRoomResponseParams `json:"Response"`
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

type MaxVideoUser struct {
	// The stream information.
	UserMediaStream *UserMediaStream `json:"UserMediaStream,omitempty" name:"UserMediaStream"`
}

type McuAudioParams struct {
	// The audio encoding parameters.
	AudioEncode *AudioEncode `json:"AudioEncode,omitempty" name:"AudioEncode"`

	// The audio mix allowlist. For the `StartPublishCdnStream` API, if you do not pass this parameter or leave it empty, the audios of all anchors will be mixed. For the `UpdatePublishCdnStream` API, if you do not pass this parameter, no changes will be made to the current allowlist; if you pass in an empty string, the audios of all anchors will be mixed.
	// In cases where `SubscribeAudioList` and `UnSubscribeAudioList` are used at the same time, you need to specify both parameters. If you pass neither `SubscribeAudioList` nor `UnSubscribeAudioList`, no changes will be made. If a user is included in both parameters, the user’s audio will not be mixed.
	SubscribeAudioList []*McuUserInfoParams `json:"SubscribeAudioList,omitempty" name:"SubscribeAudioList"`

	// The audio mix blocklist. If you do not pass this parameter or leave it empty, there won’t be a blocklist. For the `UpdatePublishCdnStream` API, if you do not pass this parameter, no changes will be made to the current blocklist; if you pass in an empty string, the blocklist will be reset.
	// In cases where `SubscribeAudioList` and `UnSubscribeAudioList` are used at the same time, you need to specify both parameters. If you pass neither `SubscribeAudioList` nor `UnSubscribeAudioList`, no changes will be made. If a user is included in both parameters, the user’s audio will not be mixed.
	UnSubscribeAudioList []*McuUserInfoParams `json:"UnSubscribeAudioList,omitempty" name:"UnSubscribeAudioList"`
}

type McuCustomCrop struct {
	// The horizontal offset (pixels) of the starting point for cropping. This parameter must be greater than 0.
	LocationX *uint64 `json:"LocationX,omitempty" name:"LocationX"`

	// The vertical offset (pixels) of the starting point for cropping. This parameter must be greater than 0.
	LocationY *uint64 `json:"LocationY,omitempty" name:"LocationY"`

	// The video width (pixels) after cropping. The sum of this parameter and `LocationX` cannot be greater than 10000.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// The video height (pixels) after cropping. The sum of this parameter and `LocationY` cannot be greater than 10000.
	Height *uint64 `json:"Height,omitempty" name:"Height"`
}

type McuFeedBackRoomParams struct {
	// The room ID.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// The ID type of the room to which streams are relayed. `0` indicates integer, and `1` indicates string.
	RoomIdType *uint64 `json:"RoomIdType,omitempty" name:"RoomIdType"`

	// The [user ID](https://intl.cloud.tencent.com/document/product/647/37714) of the relaying robot in the TRTC room, which cannot be the same as a user ID already in use. We recommend you include the room ID in this user ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The signature (similar to login password) required for the relaying robot to enter the room. For information on how to calculate the signature, see [What is UserSig?](https://intl.cloud.tencent.com/document/product/647/38104).
	UserSig *string `json:"UserSig,omitempty" name:"UserSig"`
}

type McuLayout struct {
	// The information of the stream that is displayed. If you do not pass this parameter, TRTC will display the videos of anchors in the room according to their room entry sequence.
	UserMediaStream *UserMediaStream `json:"UserMediaStream,omitempty" name:"UserMediaStream"`

	// The video width (pixels). If you do not pass this parameter, 0 will be used.
	ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// The video height (pixels). If you do not pass this parameter, 0 will be used.
	ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// The horizontal offset (pixels) of the video. The sum of `LocationX` and `ImageWidth` cannot exceed the width of the canvas. If you do not pass this parameter, 0 will be used.
	LocationX *uint64 `json:"LocationX,omitempty" name:"LocationX"`

	// The vertical offset of the video. The sum of `LocationY` and `ImageHeight` cannot exceed the height of the canvas. If you do not pass this parameter, 0 will be used.
	LocationY *uint64 `json:"LocationY,omitempty" name:"LocationY"`

	// The image layer of the video. If you do not pass this parameter, 0 will be used.
	ZOrder *uint64 `json:"ZOrder,omitempty" name:"ZOrder"`

	// The rendering mode of the video. 0 (the video is scaled and the excess parts are cropped), 1 (the video is scaled), 2 (the video is scaled and the blank spaces are filled with black bars). If you do not pass this parameter, 0 will be used.
	RenderMode *uint64 `json:"RenderMode,omitempty" name:"RenderMode"`

	// The background color of the video. Below are the values for some common colors:
	// Red: 0xcc0033
	// Yellow: 0xcc9900
	// Green: 0xcccc33
	// Blue: 0x99CCFF
	// Black: 0x000000
	// White: 0xFFFFFF
	// Grey: 0x999999
	BackGroundColor *string `json:"BackGroundColor,omitempty" name:"BackGroundColor"`

	// The URL of the background image for the video. This parameter allows you to specify an image to display when the user’s camera is turned off or before the user enters the room. If the dimensions of the image specified are different from those of the video window, the image will be stretched to fit the space. This parameter has a higher priority than `BackGroundColor`.
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" name:"BackgroundImageUrl"`

	// Custom cropping.
	CustomCrop *McuCustomCrop `json:"CustomCrop,omitempty" name:"CustomCrop"`
}

type McuLayoutParams struct {
	// The layout mode. Valid values: 1 (floating), 2 (screen sharing), 3 (grid), 4 (custom). Floating, screen sharing, and grid are dynamic layouts. Custom layouts are static layouts.
	MixLayoutMode *uint64 `json:"MixLayoutMode,omitempty" name:"MixLayoutMode"`

	// Whether to display users who publish only audio. 0: No; 1: Yes. This parameter is valid only if a dynamic layout is used. If you do not pass this parameter, 0 will be used.
	PureAudioHoldPlaceMode *uint64 `json:"PureAudioHoldPlaceMode,omitempty" name:"PureAudioHoldPlaceMode"`

	// The details of a custom layout.
	MixLayoutList []*McuLayout `json:"MixLayoutList,omitempty" name:"MixLayoutList"`

	// The information of the large video in screen sharing or floating layout mode.
	MaxVideoUser *MaxVideoUser `json:"MaxVideoUser,omitempty" name:"MaxVideoUser"`
}

type McuLayoutVolume struct {
	// The application data, which will be embedded in the `app_data` field of the custom SEI. It must be shorter than 4,096 characters.
	AppData *string `json:"AppData,omitempty" name:"AppData"`

	// The payload type of the SEI message. The default is 100. Value range: 100-254 (244 is used internally by Tencent Cloud for timestamps).
	PayloadType *uint64 `json:"PayloadType,omitempty" name:"PayloadType"`
}

type McuPassThrough struct {
	// The payload of the pass-through SEI.
	PayloadContent *string `json:"PayloadContent,omitempty" name:"PayloadContent"`

	// The payload type of the SEI message. Value range: 5 and 100-254 (244 is used internally by Tencent Cloud for timestamps).
	PayloadType *uint64 `json:"PayloadType,omitempty" name:"PayloadType"`

	// This parameter is required only if `PayloadType` is 5. It must be a 32-character hexadecimal string. If `PayloadType` is not 5, this parameter will be ignored.
	PayloadUuid *string `json:"PayloadUuid,omitempty" name:"PayloadUuid"`
}

type McuPublishCdnParam struct {
	// The URLs of the CDNs to relay to.
	PublishCdnUrl *string `json:"PublishCdnUrl,omitempty" name:"PublishCdnUrl"`

	// Whether to relay to Tencent Cloud’s CDN. 0: Third-party CDN; 1 (default): Tencent Cloud’s CDN. Relaying to a third-party CDN will incur fees. To avoid unexpected charges, we recommend you pass in a specific value. For details, see the API document.
	IsTencentCdn *uint64 `json:"IsTencentCdn,omitempty" name:"IsTencentCdn"`
}

type McuSeiParams struct {
	// The audio volume layout SEI.
	LayoutVolume *McuLayoutVolume `json:"LayoutVolume,omitempty" name:"LayoutVolume"`

	// The pass-through SEI.
	PassThrough *McuPassThrough `json:"PassThrough,omitempty" name:"PassThrough"`
}

type McuUserInfoParams struct {
	// The user information.
	UserInfo *MixUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
}

type McuVideoParams struct {
	// The video encoding parameters.
	VideoEncode *VideoEncode `json:"VideoEncode,omitempty" name:"VideoEncode"`

	// The layout parameters.
	LayoutParams *McuLayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// The canvas color. Below are the values for some common colors:
	// Red: 0xcc0033
	// Yellow: 0xcc9900
	// Green: 0xcccc33
	// Blue: 0x99CCFF
	// Black: 0x000000
	// White: 0xFFFFFF
	// Grey: 0x999999
	BackGroundColor *string `json:"BackGroundColor,omitempty" name:"BackGroundColor"`

	// The URL of the background image for the canvas. This parameter has a higher priority than `BackGroundColor`.
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" name:"BackgroundImageUrl"`

	// The watermark information for the mixed stream.
	WaterMarkList []*McuWaterMarkParams `json:"WaterMarkList,omitempty" name:"WaterMarkList"`
}

type McuWaterMarkImage struct {
	// The URL of the watermark image, which must be in PNG, JPG, or JPEG format and cannot exceed 5 MB.
	WaterMarkUrl *string `json:"WaterMarkUrl,omitempty" name:"WaterMarkUrl"`

	// The watermark width (pixels).
	WaterMarkWidth *uint64 `json:"WaterMarkWidth,omitempty" name:"WaterMarkWidth"`

	// The watermark height (pixels).
	WaterMarkHeight *uint64 `json:"WaterMarkHeight,omitempty" name:"WaterMarkHeight"`

	// The horizontal offset (pixels) of the watermark.
	LocationX *uint64 `json:"LocationX,omitempty" name:"LocationX"`

	// The vertical offset (pixels) of the watermark.
	LocationY *uint64 `json:"LocationY,omitempty" name:"LocationY"`

	// The image layer of the watermark. If you do not pass this parameter, 0 will be used.
	ZOrder *uint64 `json:"ZOrder,omitempty" name:"ZOrder"`
}

type McuWaterMarkParams struct {
	// The watermark type. The default is 0, which indicates an image watermark.
	WaterMarkType *uint64 `json:"WaterMarkType,omitempty" name:"WaterMarkType"`

	// The watermark image information. This parameter is required if `WaterMarkType` is 0.
	WaterMarkImage *McuWaterMarkImage `json:"WaterMarkImage,omitempty" name:"WaterMarkImage"`
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

type MixUserInfo struct {
	// User ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// If a dynamic layout is used, the value of this parameter should be the ID of the main room. If a custom layout is used, the value of this parameter should be the same as the room ID in `MixLayoutList`.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// The type of the `RoomId` parameter. 0: integer; 1: string.
	RoomIdType *uint64 `json:"RoomIdType,omitempty" name:"RoomIdType"`
}

// Predefined struct for user
type ModifyCloudRecordingRequestParams struct {
	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The new stream mixing layout to use.
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitempty" name:"MixLayoutParams"`

	// The allowlist/blocklist for stream subscription.
	SubscribeStreamUserIds *SubscribeStreamUserIds `json:"SubscribeStreamUserIds,omitempty" name:"SubscribeStreamUserIds"`
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

// Predefined struct for user
type ModifyCloudRecordingResponseParams struct {
	// The task ID assigned by the recording service, which uniquely identifies a recording process and becomes invalid after a recording task ends.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudRecordingResponseParams `json:"Response"`
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
	// 1: Single-stream recording. Records the audio and video of each subscribed user (`UserId`) in a room and saves the recording files to the cloud.
	// 2: Mixed-stream recording. Mixes the audios and videos of subscribed users (`UserId`) in a room, records the mixed stream, and saves the recording files to the cloud.
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

	// The output format. 0 (default): HLS; 1: HLS + MP4 (recorded in HLS and converted to MP4). This parameter is invalid if you save recording files to VOD. To specify the format of files saved to VOD, use `MediaType` of `TencentVod`.
	OutputFormat *uint64 `json:"OutputFormat,omitempty" name:"OutputFormat"`

	// Whether to merge the audio and video of a user in the single-stream recording mode. 0 (default): Do not mix the audio and video; 1: Mix the audio and video into one TS file. You don’t need to specify this parameter for mixed-stream recording, which merges audios and videos by default.
	AvMerge *uint64 `json:"AvMerge,omitempty" name:"AvMerge"`
}

// Predefined struct for user
type RemoveUserByStrRoomIdRequestParams struct {
	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// List of up to 10 users to be removed
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
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

// Predefined struct for user
type RemoveUserByStrRoomIdResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveUserByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserByStrRoomIdResponseParams `json:"Response"`
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

// Predefined struct for user
type RemoveUserRequestParams struct {
	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room number.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// List of up to 10 users to be removed.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
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

// Predefined struct for user
type RemoveUserResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveUserResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserResponseParams `json:"Response"`
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

// Predefined struct for user
type SetUserBlockedByStrRoomIdRequestParams struct {
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The room ID (string).
	StrRoomId *string `json:"StrRoomId,omitempty" name:"StrRoomId"`

	// The user ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Whether to disable the user’s audio and video. 0: Enable; 1: Disable.
	IsMute *uint64 `json:"IsMute,omitempty" name:"IsMute"`
}

type SetUserBlockedByStrRoomIdRequest struct {
	*tchttp.BaseRequest
	
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The room ID (string).
	StrRoomId *string `json:"StrRoomId,omitempty" name:"StrRoomId"`

	// The user ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Whether to disable the user’s audio and video. 0: Enable; 1: Disable.
	IsMute *uint64 `json:"IsMute,omitempty" name:"IsMute"`
}

func (r *SetUserBlockedByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserBlockedByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StrRoomId")
	delete(f, "UserId")
	delete(f, "IsMute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetUserBlockedByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetUserBlockedByStrRoomIdResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetUserBlockedByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *SetUserBlockedByStrRoomIdResponseParams `json:"Response"`
}

func (r *SetUserBlockedByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserBlockedByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetUserBlockedRequestParams struct {
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The room ID (number).
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The user ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Whether to disable the user’s audio and video. 0: Enable; 1: Disable.
	IsMute *uint64 `json:"IsMute,omitempty" name:"IsMute"`
}

type SetUserBlockedRequest struct {
	*tchttp.BaseRequest
	
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The room ID (number).
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The user ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Whether to disable the user’s audio and video. 0: Enable; 1: Disable.
	IsMute *uint64 `json:"IsMute,omitempty" name:"IsMute"`
}

func (r *SetUserBlockedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserBlockedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserId")
	delete(f, "IsMute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetUserBlockedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetUserBlockedResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetUserBlockedResponse struct {
	*tchttp.BaseResponse
	Response *SetUserBlockedResponseParams `json:"Response"`
}

func (r *SetUserBlockedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserBlockedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SingleSubscribeParams struct {
	// The stream information.
	UserMediaStream *UserMediaStream `json:"UserMediaStream,omitempty" name:"UserMediaStream"`
}

// Predefined struct for user
type StartPublishCdnStreamRequestParams struct {
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The ID of the room whose streams are relayed (the main room).
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// The type of the `RoomId` parameter, which must be the same as the ID type of the room whose streams are relayed. 0: integer; 1: string.
	RoomIdType *uint64 `json:"RoomIdType,omitempty" name:"RoomIdType"`

	// The information of the relaying robot in the room.
	AgentParams *AgentParams `json:"AgentParams,omitempty" name:"AgentParams"`

	// Whether to transcode the streams. 0: No; 1: Yes.
	WithTranscoding *uint64 `json:"WithTranscoding,omitempty" name:"WithTranscoding"`

	// The audio encoding parameters for relaying.
	AudioParams *McuAudioParams `json:"AudioParams,omitempty" name:"AudioParams"`

	// The video encoding parameters for relaying. If you do not pass this parameter, only audio will be relayed.
	VideoParams *McuVideoParams `json:"VideoParams,omitempty" name:"VideoParams"`

	// The information of a single stream relayed. When you relay a single stream, set `WithTranscoding` to 0.
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitempty" name:"SingleSubscribeParams"`

	// The CDN information.
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitempty" name:"PublishCdnParams"`

	// The stream mixing SEI parameters.
	SeiParams *McuSeiParams `json:"SeiParams,omitempty" name:"SeiParams"`

	// The information of the room to which streams are relayed.
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitempty" name:"FeedBackRoomParams"`
}

type StartPublishCdnStreamRequest struct {
	*tchttp.BaseRequest
	
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The ID of the room whose streams are relayed (the main room).
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// The type of the `RoomId` parameter, which must be the same as the ID type of the room whose streams are relayed. 0: integer; 1: string.
	RoomIdType *uint64 `json:"RoomIdType,omitempty" name:"RoomIdType"`

	// The information of the relaying robot in the room.
	AgentParams *AgentParams `json:"AgentParams,omitempty" name:"AgentParams"`

	// Whether to transcode the streams. 0: No; 1: Yes.
	WithTranscoding *uint64 `json:"WithTranscoding,omitempty" name:"WithTranscoding"`

	// The audio encoding parameters for relaying.
	AudioParams *McuAudioParams `json:"AudioParams,omitempty" name:"AudioParams"`

	// The video encoding parameters for relaying. If you do not pass this parameter, only audio will be relayed.
	VideoParams *McuVideoParams `json:"VideoParams,omitempty" name:"VideoParams"`

	// The information of a single stream relayed. When you relay a single stream, set `WithTranscoding` to 0.
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitempty" name:"SingleSubscribeParams"`

	// The CDN information.
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitempty" name:"PublishCdnParams"`

	// The stream mixing SEI parameters.
	SeiParams *McuSeiParams `json:"SeiParams,omitempty" name:"SeiParams"`

	// The information of the room to which streams are relayed.
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitempty" name:"FeedBackRoomParams"`
}

func (r *StartPublishCdnStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishCdnStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "RoomIdType")
	delete(f, "AgentParams")
	delete(f, "WithTranscoding")
	delete(f, "AudioParams")
	delete(f, "VideoParams")
	delete(f, "SingleSubscribeParams")
	delete(f, "PublishCdnParams")
	delete(f, "SeiParams")
	delete(f, "FeedBackRoomParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishCdnStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishCdnStreamResponseParams struct {
	// The task ID, which is generated by the Tencent Cloud server. You need to pass in the task ID when making a request to update or stop a relaying task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartPublishCdnStreamResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishCdnStreamResponseParams `json:"Response"`
}

func (r *StartPublishCdnStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishCdnStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishCdnStreamRequestParams struct {
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type StopPublishCdnStreamRequest struct {
	*tchttp.BaseRequest
	
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopPublishCdnStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishCdnStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopPublishCdnStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishCdnStreamResponseParams struct {
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopPublishCdnStreamResponse struct {
	*tchttp.BaseResponse
	Response *StopPublishCdnStreamResponseParams `json:"Response"`
}

func (r *StopPublishCdnStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishCdnStreamResponse) FromJsonString(s string) error {
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
	// The account information for third-party cloud storage. This parameter is not available currently. Please use `CloudVod` instead to save files to Tencent Cloud VOD.
	CloudStorage *CloudStorage `json:"CloudStorage,omitempty" name:"CloudStorage"`

	// The account information for saving files to Tencent Cloud VOD. This parameter is required. Currently, you can only save files to Tencent Cloud VOD.
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

	// The format of recording files saved to VOD. 0 (default): MP4; 1: HLS.
	MediaType *uint64 `json:"MediaType,omitempty" name:"MediaType"`
}

// Predefined struct for user
type UpdatePublishCdnStreamRequestParams struct {
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The sequence of a request. This parameter ensures the requests to change the parameters of the same relaying task are in the correct order. It increases each time a new request is made.
	SequenceNumber *uint64 `json:"SequenceNumber,omitempty" name:"SequenceNumber"`

	// Whether to transcode the streams. 0: No; 1: Yes.
	WithTranscoding *uint64 `json:"WithTranscoding,omitempty" name:"WithTranscoding"`

	// Pass this parameter to change the users whose audios are mixed. If you do not pass this parameter, no changes will be made.
	AudioParams *McuAudioParams `json:"AudioParams,omitempty" name:"AudioParams"`

	// Pass this parameter to change video parameters other than the codec, including the video layout, background image, background color, and watermark information. This parameter is valid only if streams are transcoded. If you do not pass it, no changes will be made.
	VideoParams *McuVideoParams `json:"VideoParams,omitempty" name:"VideoParams"`

	// Pass this parameter to change the single stream that is relayed. This parameter is valid only if streams are not transcoded. If you do not pass this parameter, no changes will be made.
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitempty" name:"SingleSubscribeParams"`

	// Pass this parameter to change the CDNs to relay to. If you do not pass this parameter, no changes will be made.
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitempty" name:"PublishCdnParams"`

	// The stream mixing SEI parameters.
	SeiParams *McuSeiParams `json:"SeiParams,omitempty" name:"SeiParams"`

	// The information of the room to which streams are relayed.
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitempty" name:"FeedBackRoomParams"`
}

type UpdatePublishCdnStreamRequest struct {
	*tchttp.BaseRequest
	
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The sequence of a request. This parameter ensures the requests to change the parameters of the same relaying task are in the correct order. It increases each time a new request is made.
	SequenceNumber *uint64 `json:"SequenceNumber,omitempty" name:"SequenceNumber"`

	// Whether to transcode the streams. 0: No; 1: Yes.
	WithTranscoding *uint64 `json:"WithTranscoding,omitempty" name:"WithTranscoding"`

	// Pass this parameter to change the users whose audios are mixed. If you do not pass this parameter, no changes will be made.
	AudioParams *McuAudioParams `json:"AudioParams,omitempty" name:"AudioParams"`

	// Pass this parameter to change video parameters other than the codec, including the video layout, background image, background color, and watermark information. This parameter is valid only if streams are transcoded. If you do not pass it, no changes will be made.
	VideoParams *McuVideoParams `json:"VideoParams,omitempty" name:"VideoParams"`

	// Pass this parameter to change the single stream that is relayed. This parameter is valid only if streams are not transcoded. If you do not pass this parameter, no changes will be made.
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitempty" name:"SingleSubscribeParams"`

	// Pass this parameter to change the CDNs to relay to. If you do not pass this parameter, no changes will be made.
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitempty" name:"PublishCdnParams"`

	// The stream mixing SEI parameters.
	SeiParams *McuSeiParams `json:"SeiParams,omitempty" name:"SeiParams"`

	// The information of the room to which streams are relayed.
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitempty" name:"FeedBackRoomParams"`
}

func (r *UpdatePublishCdnStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePublishCdnStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	delete(f, "SequenceNumber")
	delete(f, "WithTranscoding")
	delete(f, "AudioParams")
	delete(f, "VideoParams")
	delete(f, "SingleSubscribeParams")
	delete(f, "PublishCdnParams")
	delete(f, "SeiParams")
	delete(f, "FeedBackRoomParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePublishCdnStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePublishCdnStreamResponseParams struct {
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdatePublishCdnStreamResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePublishCdnStreamResponseParams `json:"Response"`
}

func (r *UpdatePublishCdnStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePublishCdnStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserMediaStream struct {
	// The user information.
	UserInfo *MixUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// The stream type. 0: Camera; 1: Screen sharing. If you do not pass this parameter, 0 will be used.
	StreamType *uint64 `json:"StreamType,omitempty" name:"StreamType"`
}

type VideoEncode struct {
	// The width of the output stream (pixels). This parameter is required if audio and video are relayed. Value range: [0, 1920].
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// The height of the output stream (pixels). This parameter is required if audio and video are relayed. Value range: [0, 1080].
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// The frame rate (fps) of the output stream. This parameter is required if audio and video are relayed. Value range: [0, 60].
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// The bitrate (Kbps) of the output stream. This parameter is required if audio and video are relayed. Value range: [0, 10000].
	BitRate *uint64 `json:"BitRate,omitempty" name:"BitRate"`

	// The GOP (seconds) of the output stream. This parameter is required if audio and video are relayed. Value range: [1, 5].
	Gop *uint64 `json:"Gop,omitempty" name:"Gop"`
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