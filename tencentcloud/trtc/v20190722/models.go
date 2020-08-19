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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AbnormalEvent struct {

	// Exception event ID. For specific values, please see Appendix. Exceptional Experience ID Mapping Table.
	AbnormalEventId *uint64 `json:"AbnormalEventId,omitempty" name:"AbnormalEventId"`

	// Remote user ID. If this parameter is left empty, it indicates that the exception event is not caused by the remote user.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`
}

type AbnormalExperience struct {

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Exceptional experience ID
	ExperienceId *uint64 `json:"ExperienceId,omitempty" name:"ExperienceId"`

	// Room ID in string type
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// Exception event array
	AbnormalEventList []*AbnormalEvent `json:"AbnormalEventList,omitempty" name:"AbnormalEventList" list`

	// Report time of the exception event
	EventTime *uint64 `json:"EventTime,omitempty" name:"EventTime"`
}

type CreateTroubleInfoRequest struct {
	*tchttp.BaseRequest

	// Application ID
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// Teacher user ID
	TeacherUserId *string `json:"TeacherUserId,omitempty" name:"TeacherUserId"`

	// Student user ID
	StudentUserId *string `json:"StudentUserId,omitempty" name:"StudentUserId"`

	// ID of the user (teacher or student) with exception.
	TroubleUserId *string `json:"TroubleUserId,omitempty" name:"TroubleUserId"`

	// Exception type.
	// 1: exceptional video
	// 2: exceptional audio
	// 3: exceptional video and audio
	// 5: exceptional room entry
	// 4: course switch
	// 6: help seeking
	// 7: problem feedback
	// 8: complaint
	TroubleType *uint64 `json:"TroubleType,omitempty" name:"TroubleType"`

	// UNIX timestamp when the exception occurred in seconds.
	TroubleTime *uint64 `json:"TroubleTime,omitempty" name:"TroubleTime"`

	// Exception details
	TroubleMsg *string `json:"TroubleMsg,omitempty" name:"TroubleMsg"`
}

func (r *CreateTroubleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTroubleInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTroubleInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTroubleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTroubleInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalEventRequest struct {
	*tchttp.BaseRequest

	// User `SDKAppID`, which can be used to query 20 exceptional experience events (in one or more rooms)
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Query start time
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Room ID, which can be used to query up to 20 exceptional experience events in a specific room
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeAbnormalEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAbnormalEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned data entries.
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Exceptional experience list.
		AbnormalExperienceList []*AbnormalExperience `json:"AbnormalExperienceList,omitempty" name:"AbnormalExperienceList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAbnormalEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCallDetailRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a call: sdkappid_roomgString_createTime. The `roomgString` refers to the room ID, and `createTime` refers to the creation time of a room in the format of UNIX timestamp in seconds, such as 1400353843_218695_1590065777. Its value can be obtained from the `DescribeRoomInformation` API (related document: https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1).
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// Query start time in the format of local UNIX timestamp, such as 1588031999s, which is a point in time in the last 5 days.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of local UNIX timestamp, such as 1588031999s.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// User `sdkappid`, such as 1400188366.
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// User array to query, which contains up to 6 users. If it is left empty, 6 users will be returned by default.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`

	// Metric to query. The user list will be returned if it is left empty; all metrics will be returned if its value is `all`.
	// appCpu: CPU utilization of the application;
	// sysCpu: CPU utilization of the system;
	// aBit: upstream/downstream audio bitrate;
	// aBlock: audio lag duration;
	// bigvBit: upstream/downstream video bitrate;
	// bigvCapFps: frame rate for capturing videos;
	// bigvEncFps: frame rate for sending videos;
	// bigvDecFps: rendering frame rate;
	// bigvBlock: video lag duration;
	// aLoss: upstream/downstream audio packet loss;
	// bigvLoss: upstream/downstream video packet loss;
	// bigvWidth: upstream/downstream resolution in width;
	// bigvHeight: upstream/downstream resolution in height.
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeCallDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCallDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCallDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of returned users
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// User information list
		UserList []*UserInformation `json:"UserList,omitempty" name:"UserList" list`

		// Quality data
		Data []*QualityData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCallDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCallDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDetailEventRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a call: sdkappid_roomgString_createTime. The `roomgString` refers to the room ID, and `createTime` refers to the creation time of a room in the format of UNIX timestamp in seconds. Its value can be obtained from the `DescribeRoomInformation` API (related document: https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1).
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// Query start time in the format of local UNIX timestamp, such as 1588031999s, which is a point in time in the last 5 days.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of local UNIX timestamp, such as 1588031999s.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Room ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeDetailEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDetailEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDetailEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of returned events
		Data []*EventList `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDetailEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDetailEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleRequest struct {
	*tchttp.BaseRequest

	// User `sdkappid`
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Query start time in the format of local UNIX timestamp, such as 1588031999s, which is a point in time in the last 5 days.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of local UNIX timestamp, such as 1588031999s.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeHistoryScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHistoryScaleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned data entries
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Returned data
		ScaleList []*ScaleInfomation `json:"ScaleList,omitempty" name:"ScaleList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHistoryScaleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeNetworkRequest struct {
	*tchttp.BaseRequest

	// Query start time in the format of UNIX timestamp, such as 1588031999s, which is a point in time in the last 24 hours.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of local UNIX timestamp, such as 1588031999s.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// User `sdkappid`
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Type of data to query
	// sendLossRateRaw: upstream packet loss rate;
	// recvLossRateRaw: downstream packet loss rate.
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeRealtimeNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeNetworkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data returned by query
		Data []*RealtimeData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeNetworkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeQualityRequest struct {
	*tchttp.BaseRequest

	// Query start time in the format of local UNIX timestamp, such as 1588031999s, which is a point in time in the last 24 hours.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of local UNIX timestamp, such as 1588031999s.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// User `sdkappid`
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Type of data to query
	// enterTotalSuccPercent: room entry success rate;
	// fistFreamInSecRate: instant playback rate of the first frame;
	// blockPercent: video lag rate;
	// audioBlockPercent: audio lag rate.
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeRealtimeQualityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeQualityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeQualityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Type of returned data
		Data []*RealtimeData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeQualityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeQualityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeScaleRequest struct {
	*tchttp.BaseRequest

	// Query start time in the format of local UNIX timestamp, such as 1588031999s, which is a point in time in the last 24 hours.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of local UNIX timestamp, such as 1588031999s.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// User `sdkappid`
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Type of data to query
	// `UserNum: number of users in call;
	// RoomNum: number of rooms.
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeRealtimeScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeScaleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned data array
		Data []*RealtimeData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeScaleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationRequest struct {
	*tchttp.BaseRequest

	// User `sdkappid`
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Query start time in the format of local UNIX timestamp, such as 1588031999s, which is a point in time in the last 5 days.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of local UNIX timestamp, such as 1588031999s.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Room ID of uint type
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// Page index starting from 0 (if either `PageNumber` or `PageSize` is left empty, 10 data entries will be returned by default)
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`

	// Number of entries per page (if either `PageNumber` or `PageSize` is left empty, 10 data entries will be returned by default. Maximum value: 100)
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeRoomInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoomInformationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of returned data entries.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// Room information list
		RoomList []*RoomState `json:"RoomList,omitempty" name:"RoomList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoomInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoomInformationResponse) FromJsonString(s string) error {
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

func (r *DismissRoomRequest) FromJsonString(s string) error {
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

func (r *DismissRoomResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EncodeParams struct {

	// Output stream audio sample rate for stream mix. Valid values: 48000, 44100, 32000,24000, 22050, 16000, 12000, 11025, 8000.
	AudioSampleRate *uint64 `json:"AudioSampleRate,omitempty" name:"AudioSampleRate"`

	// Output stream audio bitrate in Kbps for On-Cloud MixTranscoding. Value range: [8, 500].
	AudioBitrate *uint64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// Number of output stream audio sound channels for On-Cloud MixTranscoding. Value range: [1, 2].
	AudioChannels *uint64 `json:"AudioChannels,omitempty" name:"AudioChannels"`

	// Output stream width in pixels for On-Cloud MixTranscoding, which is required for audio/video output. Value range: [0, 1920].
	VideoWidth *uint64 `json:"VideoWidth,omitempty" name:"VideoWidth"`

	// Output stream height in pixels for On-Cloud MixTranscoding, which is required for audio/video output. Value range: [0, 1080].
	VideoHeight *uint64 `json:"VideoHeight,omitempty" name:"VideoHeight"`

	// Output stream bitrate in Kbps for On-Cloud MixTranscoding, which is required for audio/video output. Value range: [1, 10000].
	VideoBitrate *uint64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Output stream frame rate for On-Cloud MixTranscoding, which is required for audio/video output. Value range: [6, 12, 15, 24, 30, 48, 60]. If the frame rate lies outside the valid value range, it will be automatically modified to a value within the range.
	VideoFramerate *uint64 `json:"VideoFramerate,omitempty" name:"VideoFramerate"`

	// Output stream GOP in seconds for On-Cloud MixTranscoding, which is required for audio/video output. Value range: [1, 5].
	VideoGop *uint64 `json:"VideoGop,omitempty" name:"VideoGop"`

	// Output stream background color for On-Cloud MixTranscoding.
	BackgroundColor *uint64 `json:"BackgroundColor,omitempty" name:"BackgroundColor"`

	// Output stream background image for stream mix. Its value is the ID of image uploaded through the TRTC Console.
	BackgroundImageId *uint64 `json:"BackgroundImageId,omitempty" name:"BackgroundImageId"`
}

type EventList struct {

	// Data content
	Content []*EventMessage `json:"Content,omitempty" name:"Content" list`

	// Sender `userId`
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`
}

type EventMessage struct {

	// Video stream type:
	// 0: non-video event;
	// 2: big image;
	// 3: small image;
	// 7: relayed stream image.
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Event reporting time in the format of UNIX timestamp, such as 1589891188801ms
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// Event ID. Events divide into SDK events and WebRTC events. For more information, please see Appendix - Event ID Mapping Table at https://intl.cloud.tencent.com/document/product/647/44916?from_cn_redirect=1
	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`

	// First event parameter, such as video resolution width
	ParamOne *int64 `json:"ParamOne,omitempty" name:"ParamOne"`

	// Second event parameter, such as video resolution height
	ParamTwo *int64 `json:"ParamTwo,omitempty" name:"ParamTwo"`
}

type LayoutParams struct {

	// On-cloud stream mix layout template ID. 0: floating template (default value); 1: grid template; 2: screen sharing template; 3: picture-in-picture template.
	Template *uint64 `json:"Template,omitempty" name:"Template"`

	// ID of the user in the big image, which takes effect in a screen sharing, floating, or picture-in-picture template.
	MainVideoUserId *string `json:"MainVideoUserId,omitempty" name:"MainVideoUserId"`

	// Stream type of the big image, which takes effect in a screen sharing, floating, or picture-in-picture template. 0: camera; 1: screen sharing. If a web user's stream is displayed in the big image on the left, enter 0 for this parameter.
	MainVideoStreamType *uint64 `json:"MainVideoStreamType,omitempty" name:"MainVideoStreamType"`

	// Layout parameter of the small image, which takes effect in a picture-in-picture template.
	SmallVideoLayoutParams *SmallVideoLayoutParams `json:"SmallVideoLayoutParams,omitempty" name:"SmallVideoLayoutParams"`
}

type OutputParams struct {

	// Custom live stream ID, which must be different from the ID of relayed stream.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Value range: [0, 1]. If it is 0, live streams are audio and video; if it is 1, live streams are only audio. Default value: 0.
	PureAudioStream *uint64 `json:"PureAudioStream,omitempty" name:"PureAudioStream"`

	// Custom recording file name
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// Value range: [0, 1]. If it is 0, the recording template configured in the console will be used; if it is 1, streams are recorded as .mp3 files.
	RecordAudioOnly *uint64 `json:"RecordAudioOnly,omitempty" name:"RecordAudioOnly"`
}

type QualityData struct {

	// Data content
	Content []*TimeValue `json:"Content,omitempty" name:"Content" list`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Peer ID. An empty value indicates that the returned data is upstream.
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`

	// Data type
	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

type RealtimeData struct {

	// Returned data
	Content []*TimeValue `json:"Content,omitempty" name:"Content" list`

	// Data type field
	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

type RemoveUserRequest struct {
	*tchttp.BaseRequest

	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room number.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// List of up to 10 users to be removed.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`
}

func (r *RemoveUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveUserRequest) FromJsonString(s string) error {
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

func (r *RemoveUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RoomState struct {

	// Call ID (unique call ID)
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// Room ID of string type
	RoomString *string `json:"RoomString,omitempty" name:"RoomString"`

	// Room creation time
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Room termination time
	DestroyTime *uint64 `json:"DestroyTime,omitempty" name:"DestroyTime"`

	// Whether the room is terminated
	IsFinished *bool `json:"IsFinished,omitempty" name:"IsFinished"`

	// Room creator ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type ScaleInfomation struct {

	// Start time for each day
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// Number of users in room. If a user enters the room for multiple times, the user will be counted as one user.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserNumber *uint64 `json:"UserNumber,omitempty" name:"UserNumber"`

	// Number of room entries. Every time when a user enters the room, it will be counted as one room entry.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserCount *uint64 `json:"UserCount,omitempty" name:"UserCount"`

	// Number of rooms under `sdkappid` on a day
	// Note: this field may return null, indicating that no valid values can be obtained.
	RoomNumbers *uint64 `json:"RoomNumbers,omitempty" name:"RoomNumbers"`
}

type SmallVideoLayoutParams struct {

	// ID of the user in the small image.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Stream type of the small image. 0: camera; 1: screen sharing. If a web user's stream is displayed in the small image, enter 0 for this parameter.
	StreamType *uint64 `json:"StreamType,omitempty" name:"StreamType"`

	// Output width of the small image in pixels. If this parameter is left empty, 0 will be used by default.
	ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// Output height of the small image in pixels. If this parameter is left empty, 0 will be used by default.
	ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// Output X-axis offset of the small image in pixels. The sum of `LocationX` and `ImageWidth` cannot exceed the total width of the output mixed stream. If this parameter is left empty, 0 will be used by default.
	LocationX *uint64 `json:"LocationX,omitempty" name:"LocationX"`

	// Output Y-axis offset of the small image in pixels. The sum of `LocationY` and `ImageHeight` cannot exceed the total height of the output mixed stream. If this parameter is left empty, 0 will be used by default.
	LocationY *uint64 `json:"LocationY,omitempty" name:"LocationY"`
}

type StartMCUMixTranscodeRequest struct {
	*tchttp.BaseRequest

	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// On-Cloud MixTranscoding output control parameters.
	OutputParams *OutputParams `json:"OutputParams,omitempty" name:"OutputParams"`

	// On-Cloud MixTranscoding output encoding parameters.
	EncodeParams *EncodeParams `json:"EncodeParams,omitempty" name:"EncodeParams"`

	// On-Cloud MixTranscoding output layout parameters.
	LayoutParams *LayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`
}

func (r *StartMCUMixTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMCUMixTranscodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartMCUMixTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMCUMixTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMCUMixTranscodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeRequest struct {
	*tchttp.BaseRequest

	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *StopMCUMixTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMCUMixTranscodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMCUMixTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMCUMixTranscodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TimeValue struct {

	// Time in the format of UNIX timestamp, such as 1590065877s.
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// Parameter value returned in the current time. For example, if `bigvCapFps` is set to 0 when the current time is 1590065877s (UNIX timestamp), the value of this parameter is 0.
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type UserInformation struct {

	// 
	RoomStr *string `json:"RoomStr,omitempty" name:"RoomStr"`

	// 
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 
	JoinTs *uint64 `json:"JoinTs,omitempty" name:"JoinTs"`

	// The time when the user exits the room. If the user is still in the room, the current time will be returned
	LeaveTs *uint64 `json:"LeaveTs,omitempty" name:"LeaveTs"`

	// 
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 
	SdkVersion *string `json:"SdkVersion,omitempty" name:"SdkVersion"`

	// 
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 
	Finished *bool `json:"Finished,omitempty" name:"Finished"`
}
