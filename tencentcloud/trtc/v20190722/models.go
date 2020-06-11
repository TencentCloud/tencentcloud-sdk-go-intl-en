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

type DescribeCallDetailRequest struct {
	*tchttp.BaseRequest

	// Call ID (unique call ID): sdkappid_roomgString (room ID)_createTime (room creation time in UNIX timestamp in seconds). You can get the parameter value through the `DescribeRoomInformation` API which is used to query the room list.
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// Query start time in the format of local UNIX timestamp, such as 1588031999s, which is a point in time in the last 5 days.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of local UNIX timestamp, such as 1588031999s.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// User `sdkappid`
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// User array to query, which contains up to 6 users. If it is left empty, 6 users will be returned by default.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`

	// Metric to query. The user list will be returned if it is left empty; all metrics will be returned if its value is `all`.
	// appCpu: CPU utilization of application;
	// sysCpu: CPU utilization of system;
	// aBit: upstream/downstream audio bitrate;
	// aBlock: audio lag duration;
	// vBit: upstream/downstream video bitrate;
	// vCapFps: video capturing frame rate;
	// vEncFps: video sending frame rate;
	// vDecFps: rendering frame rate;
	// vBlock: video lag duration;
	// aLoss: upstream/downstream audio packet loss;
	// vLoss: upstream/downstream video packet loss;
	// vWidth: upstream/downstream resolution in width;
	// vHeight: upstream/downstream resolution in height.
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

	// Query start time in the format of local UNIX timestamp, such as 1588031999s, which is a point in time in the last 24 hours.
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

	// Page index. If it is left empty, 10 entries will be returned by default.
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`

	// Page size. Maximum value: 100. If it is left empty, 10 entries will be returned by default.
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

	// Number of users in room
	UserNumber *uint64 `json:"UserNumber,omitempty" name:"UserNumber"`

	// Number of times a room has been entered
	UserCount *uint64 `json:"UserCount,omitempty" name:"UserCount"`

	// Number of rooms.
	RoomNumbers *uint64 `json:"RoomNumbers,omitempty" name:"RoomNumbers"`
}

type TimeValue struct {

	// Time
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// Current time value in the format of UNIX timestamp
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type UserInformation struct {

	// Room ID of string type.
	RoomStr *string `json:"RoomStr,omitempty" name:"RoomStr"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The time when the user enters the room
	JoinTs *uint64 `json:"JoinTs,omitempty" name:"JoinTs"`

	// The time when the user exits the room
	LeaveTs *uint64 `json:"LeaveTs,omitempty" name:"LeaveTs"`

	// Device type
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// SDK version number
	SdkVersion *string `json:"SdkVersion,omitempty" name:"SdkVersion"`

	// Client IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// Determine whether a user has left the room
	Finished *bool `json:"Finished,omitempty" name:"Finished"`
}
