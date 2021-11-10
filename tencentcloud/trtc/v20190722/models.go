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
	AbnormalEventList []*AbnormalEvent `json:"AbnormalEventList,omitempty" name:"AbnormalEventList"`

	// Report time of the exception event
	EventTime *uint64 `json:"EventTime,omitempty" name:"EventTime"`
}

type CreatePictureRequest struct {
	*tchttp.BaseRequest

	// Application ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Base64-encoded image content
	Content *string `json:"Content,omitempty" name:"Content"`

	// Image file extension
	Suffix *string `json:"Suffix,omitempty" name:"Suffix"`

	// Image height
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Image width
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// X-axis value of the image’s position
	XPosition *uint64 `json:"XPosition,omitempty" name:"XPosition"`

	// Y-axis value of the image’s position
	YPosition *uint64 `json:"YPosition,omitempty" name:"YPosition"`
}

func (r *CreatePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Content")
	delete(f, "Suffix")
	delete(f, "Height")
	delete(f, "Width")
	delete(f, "XPosition")
	delete(f, "YPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePictureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Image ID
		PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTroubleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "TeacherUserId")
	delete(f, "StudentUserId")
	delete(f, "TroubleUserId")
	delete(f, "TroubleType")
	delete(f, "TroubleTime")
	delete(f, "TroubleMsg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTroubleInfoRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTroubleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePictureRequest struct {
	*tchttp.BaseRequest

	// Image ID
	PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`

	// Application ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DeletePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PictureId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePictureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalEventRequest struct {
	*tchttp.BaseRequest

	// User `SDKAppID`, which can be used to query 20 exceptional experience events (in one or more rooms)
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Query start time (s) in the format of Unix timestamp, e.g., 1592448600
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time (s) in the format of Unix timestamp, e.g., 1592449080
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Room ID, which can be used to query up to 20 exceptional experience events in a specific room
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeAbnormalEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned data entries.
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Exceptional experience list.
		AbnormalExperienceList []*AbnormalExperience `json:"AbnormalExperienceList,omitempty" name:"AbnormalExperienceList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCallDetailRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a call: sdkappid_roomgString_createTime. The `roomgString` refers to the room ID, and `createTime` refers to the creation time of a room in the format of UNIX timestamp in seconds, such as 1400353843_218695_1590065777. Its value can be obtained from the `DescribeRoomInformation` API (related document: https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1).
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// Query start time (s) in the format of Unix timestamp (e.g., 1590065777), which must be a time point in the last 14 days. The start and end time for query must not be more than 1 hour apart.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time (s) in the format of Unix timestamp, e.g., 1590065877
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// `SDKAppID` of the users to query, e.g., 1400353843
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// User array to query, which contains up to 6 users. If it is left empty, 6 users will be returned by default.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

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
	DataType []*string `json:"DataType,omitempty" name:"DataType"`

	// Page index starting from 0. If either `PageNumber` or `PageSize` is left empty, 6 data entries will be returned by default.
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`

	// Number of entries per page. If either `PageNumber` or `PageSize` is left empty, 6 data entries will be returned by default. When either `DataType` or `UserId` is not null, `PageSize` is up to 6. When `DataType` and `UserId` are null, `PageSize` is up to 100.
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeCallDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "UserIds")
	delete(f, "DataType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCallDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of returned users
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// User information list
		UserList []*UserInformation `json:"UserList,omitempty" name:"UserList"`

		// Quality data
		Data []*QualityData `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCallDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDetailEventRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a call: sdkappid_roomgString_createTime. The `roomgString` refers to the room ID, and `createTime` refers to the creation time of a room in the format of UNIX timestamp in seconds. Its value can be obtained from the `DescribeRoomInformation` API (related document: https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1).
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// Query start time (s) in the format of Unix timestamp (e.g., 1588055615), which must be a time point in the last 14 days
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time (s) in the format of Unix timestamp, e.g., 1588058615
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetailEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDetailEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDetailEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of returned events. An empty array will be returned if no data can be found.
		Data []*EventList `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDetailEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetailEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleRequest struct {
	*tchttp.BaseRequest

	// `SDKAppID` of the users to query, e.g., 1400188366
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Query start time (s) in the format of Unix timestamp (e.g., 1587571000), which must be a time point in the last 5 days
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time (s) in the format of Unix timestamp, e.g., 1588034999
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeHistoryScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryScaleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHistoryScaleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned data entries
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Returned data
		ScaleList []*ScaleInfomation `json:"ScaleList,omitempty" name:"ScaleList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryScaleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePictureRequest struct {
	*tchttp.BaseRequest

	// Application ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Image ID. If it is left empty, the IDs of all images under the application are returned.
	PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`

	// Number of records per page. `10` is used if it is left empty.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Page number. `1` is used if it is left empty.
	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`
}

func (r *DescribePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PictureId")
	delete(f, "PageSize")
	delete(f, "PageNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePictureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of records returned
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Image information list
		PictureInfo []*PictureInfo `json:"PictureInfo,omitempty" name:"PictureInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordStatisticRequest struct {
	*tchttp.BaseRequest

	// Query start date in the format of YYYY-MM-DD
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end date in the format of YYYY-MM-DD
	// The period queried in a request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Application ID, which is optional. If it is specified, duration statistics for the specified application are returned; otherwise, the total durations of all applications are returned.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeRecordStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordStatisticResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Duration statistics of the queried application(s)
		SdkAppIdUsages []*SdkAppIdRecordUsage `json:"SdkAppIdUsages,omitempty" name:"SdkAppIdUsages"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationRequest struct {
	*tchttp.BaseRequest

	// User `sdkappid`
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Query start time (s) in the format of Unix timestamp (e.g., 1588031999), which must be a time point in the last 14 days
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time (s) in the format of Unix timestamp, e.g., 1588034999
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Room ID in string type
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of data entries displayed on the current page
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// Room information list
		RoomList []*RoomState `json:"RoomList,omitempty" name:"RoomList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoomInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrtcInteractiveTimeRequest struct {
	*tchttp.BaseRequest

	// Query start date in the format of YYYY-MM-DD
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end date in the format of YYYY-MM-DD
	// The period queried in a request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Application ID, which is optional. If it is specified, duration statistics for the specified application are returned; otherwise, the total durations of all applications are returned.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeTrtcInteractiveTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcInteractiveTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrtcInteractiveTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrtcInteractiveTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Duration statistics of the queried application(s)
		Usages []*OneSdkAppIdUsagesInfo `json:"Usages,omitempty" name:"Usages"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrtcInteractiveTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcInteractiveTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrtcMcuTranscodeTimeRequest struct {
	*tchttp.BaseRequest

	// Query start date in the format of YYYY-MM-DD
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end date in the format of YYYY-MM-DD
	// The period queried in a request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Application ID, which is optional. If it is specified, duration statistics for the specified application are returned; otherwise, the total durations of all applications are returned.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeTrtcMcuTranscodeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcMcuTranscodeTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrtcMcuTranscodeTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrtcMcuTranscodeTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Duration statistics of the queried application(s)
		Usages []*OneSdkAppIdTranscodeTimeUsagesInfo `json:"Usages,omitempty" name:"Usages"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrtcMcuTranscodeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcMcuTranscodeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInformationRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a call: sdkappid_roomgString_createTime. The `roomgString` refers to the room ID, and `createTime` refers to the creation time of a room in the format of UNIX timestamp in seconds, such as 1400353843_218695_1590065777. Its value can be obtained from the `DescribeRoomInformation` API (related document: https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1).
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// Query start time (s) in the format of Unix timestamp (e.g., 1590065777), which must be a time point in the last 14 days
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time (s) in the format of Unix timestamp (e.g., 1590065877)
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// `SDKAppID` of the users to query, e.g., 1400353843
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The array of user IDs for query. You can enter up to 6 user IDs. If it is left empty, data of 6 users will be returned.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// Page index starting from 0. If either `PageNumber` or `PageSize` is left empty, 6 data entries will be returned.
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`

	// Number of entries per page. If either `PageNumber` or `PageSize` is left empty, 6 data entries will be returned. `PageSize` is up to 100.
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeUserInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "UserIds")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of users whose information will be returned
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// User information list
	// Note: this field may return `null`, indicating that no valid value was found.
		UserList []*UserInformation `json:"UserList,omitempty" name:"UserList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInformationResponse) FromJsonString(s string) error {
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

type EncodeParams struct {

	// Output audio sample rate (Hz) for On-Cloud MixTranscoding. Valid values: 48000, 44100, 32000, 24000, 16000, 8000
	AudioSampleRate *uint64 `json:"AudioSampleRate,omitempty" name:"AudioSampleRate"`

	// Output audio bitrate (Kbps) for On-Cloud MixTranscoding. Value range: 8-500
	AudioBitrate *uint64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// Number of sound channels of output stream for On-Cloud MixTranscoding. Valid values: 1, 2. 1 represents mono-channel, and 2 represents dual-channel.
	AudioChannels *uint64 `json:"AudioChannels,omitempty" name:"AudioChannels"`

	// Output stream width in pixels for On-Cloud MixTranscoding, which is required for audio/video output. Value range: [0, 1920].
	VideoWidth *uint64 `json:"VideoWidth,omitempty" name:"VideoWidth"`

	// Output stream height in pixels for On-Cloud MixTranscoding, which is required for audio/video output. Value range: [0, 1080].
	VideoHeight *uint64 `json:"VideoHeight,omitempty" name:"VideoHeight"`

	// Output bitrate (Kbps) for On-Cloud MixTranscoding, which is required for audio-video output. Value range: 1-10000
	VideoBitrate *uint64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Output stream frame rate for On-Cloud MixTranscoding in FPS. This parameter is required for audio/video outputs. Value range: [1, 60].
	VideoFramerate *uint64 `json:"VideoFramerate,omitempty" name:"VideoFramerate"`

	// Output stream GOP in seconds for On-Cloud MixTranscoding, which is required for audio/video output. Value range: [1, 5].
	VideoGop *uint64 `json:"VideoGop,omitempty" name:"VideoGop"`

	// Output background color for On-Cloud MixTranscoding. Valid values: decimal integers. Commonly used colors include:
	// Red: 0xff0000, whose decimal number is 16724736
	// Yellow: 0xffff00, whose decimal number is 16776960
	// Green: 0x33cc00, whose decimal number is 3394560
	// Blue: 0x0066ff, whose decimal number is 26367
	// Black: 0x000000, whose decimal number is 0
	// White: 0xFFFFFF, whose decimal number is 16777215
	// Grey: 0x999999, whose decimal number is 10066329
	BackgroundColor *uint64 `json:"BackgroundColor,omitempty" name:"BackgroundColor"`

	// Output stream background image for stream mix. Its value is the ID of image uploaded through the TRTC Console.
	BackgroundImageId *uint64 `json:"BackgroundImageId,omitempty" name:"BackgroundImageId"`

	// Output audio codec for On-Cloud MixTranscoding. Valid values: 0, 1, 2. 0 (default): LC-AAC; 1: HE-AAC; 2: HE-AACv2. If this parameter is set to 2 (HE-AACv2), On-Cloud MixTranscoding can produce only dual-channel streams. If it is set to 1 (HE-AAC) or 2 (HE-AACv2), the valid values for the audio sample rate of output streams are 48000, 44100, 32000, 24000, and 16000.
	AudioCodec *uint64 `json:"AudioCodec,omitempty" name:"AudioCodec"`

	// URL of the background image for the mixed stream, which can be in PNG, JPG, JPEG, or BMP format and does not support the alpha channel. The URL must not exceed 512 bytes. When both `BackgroundImageUrl` and `BackgroundImageId` are specified, the former will be used. The background image must not exceed 10 MB.
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" name:"BackgroundImageUrl"`
}

type EventList struct {

	// Data content
	Content []*EventMessage `json:"Content,omitempty" name:"Content"`

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

	// On-cloud stream mix layout template ID. 0: floating template (default value); 1: grid template; 2: screen sharing template; 3: picture-in-picture template; 4: custom template.
	Template *uint64 `json:"Template,omitempty" name:"Template"`

	// ID of the user in the big image, which takes effect in a screen sharing, floating, or picture-in-picture template.
	MainVideoUserId *string `json:"MainVideoUserId,omitempty" name:"MainVideoUserId"`

	// Stream type of the big image, which takes effect in a screen sharing, floating, or picture-in-picture template. 0: camera; 1: screen sharing. If a web user's stream is displayed in the big image on the left, enter 0 for this parameter.
	MainVideoStreamType *uint64 `json:"MainVideoStreamType,omitempty" name:"MainVideoStreamType"`

	// Layout parameter of the small image, which takes effect in a picture-in-picture template.
	SmallVideoLayoutParams *SmallVideoLayoutParams `json:"SmallVideoLayoutParams,omitempty" name:"SmallVideoLayoutParams"`

	// You can set the layout parameter as 1 or 0 in the screen sharing template. 1: big image on the right and small images on the left, 0: big image on the left and small images on the right. The default value is 0. 
	MainVideoRightAlign *uint64 `json:"MainVideoRightAlign,omitempty" name:"MainVideoRightAlign"`

	// A user list, which takes effect for floating, grid, or screen sharing templates. When the user list has been set, the stream mix output for users in this user list will include both audio and video; the stream mix output for users not in the list will only include audio. Up to 16 users can be set.
	MixVideoUids []*string `json:"MixVideoUids,omitempty" name:"MixVideoUids"`

	// Valid in custom template, used to specify the video image position of a user in mixed streams.
	PresetLayoutConfig []*PresetLayoutConfig `json:"PresetLayoutConfig,omitempty" name:"PresetLayoutConfig"`

	// Valid in custom templates. 1: the placeholding feature is enabled; 0 (default): the feature is disabled. When the feature is enabled, but a user for whom a position is reserved is not sending video data, the position will show the corresponding placeholder image.
	PlaceHolderMode *uint64 `json:"PlaceHolderMode,omitempty" name:"PlaceHolderMode"`

	// Whether an audio-only stream occupies an image spot, which takes effect in a floating, grid, or screen sharing template. Valid values: 0 (default): when a floating or grid template is used, users sending audio only occupy image spots; when a screen sharing template is used, users (except the user whose screen is shared) sending audio only do not occupy image spots; 1: users sending audio only occupy image spots; 2: users sending audio only do not occupy image spots.
	PureAudioHoldPlaceMode *uint64 `json:"PureAudioHoldPlaceMode,omitempty" name:"PureAudioHoldPlaceMode"`

	// Watermark parameters
	WaterMarkParams *WaterMarkParams `json:"WaterMarkParams,omitempty" name:"WaterMarkParams"`
}

type ModifyPictureRequest struct {
	*tchttp.BaseRequest

	// Image ID
	PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`

	// Application ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Image height
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Image width
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// X-axis value of the image’s position
	XPosition *uint64 `json:"XPosition,omitempty" name:"XPosition"`

	// Y-axis value of the image’s position
	YPosition *uint64 `json:"YPosition,omitempty" name:"YPosition"`
}

func (r *ModifyPictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PictureId")
	delete(f, "SdkAppId")
	delete(f, "Height")
	delete(f, "Width")
	delete(f, "XPosition")
	delete(f, "YPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPictureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OneSdkAppIdTranscodeTimeUsagesInfo struct {

	// Array of relaying and transcoding durations
	SdkAppIdTranscodeTimeUsages []*SdkAppIdTrtcMcuTranscodeTimeUsage `json:"SdkAppIdTranscodeTimeUsages,omitempty" name:"SdkAppIdTranscodeTimeUsages"`

	// Number of records returned
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// ID of the application queried. Its value may be an application ID or `total`, which indicates that the total durations of all applications are queried.
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type OneSdkAppIdUsagesInfo struct {

	// Number of records returned for the `SdkAppId`
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// Array of durations
	SdkAppIdTrtcTimeUsages []*SdkAppIdTrtcUsage `json:"SdkAppIdTrtcTimeUsages,omitempty" name:"SdkAppIdTrtcTimeUsages"`

	// Application ID
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type OutputParams struct {

	// Custom live stream ID, which must be different from the ID of relayed stream.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Value range: [0, 1]. If it is 0, live streams are audio and video; if it is 1, live streams are only audio. Default value: 0.
	PureAudioStream *uint64 `json:"PureAudioStream,omitempty" name:"PureAudioStream"`

	// Prefix of custom recording file names. Please enable the recording feature in the TRTC console first. https://intl.cloud.tencent.com/document/product/647/50768?from_cn_redirect=1
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// Whether to record audio only. Valid values: 0, 1. `0`: no meaning; `1`: records into MP3 files. This parameter is not recommended. Instead, you are advised to create an audio-only recording template in the TRTC console.
	RecordAudioOnly *uint64 `json:"RecordAudioOnly,omitempty" name:"RecordAudioOnly"`
}

type PictureInfo struct {

	// Image height
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Image width
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// X-axis value of the image’s position
	XPosition *uint64 `json:"XPosition,omitempty" name:"XPosition"`

	// Y-axis value of the image’s position
	YPosition *uint64 `json:"YPosition,omitempty" name:"YPosition"`

	// Application ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Image ID
	PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`
}

type PresetLayoutConfig struct {

	// Used to assign users to preset positions; if not assigned, users will occupy the positions set in `PresetLayoutConfig` in room entry sequence.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Stream type of the user when a specified user is assigned to the image. 0: camera; 1: screen sharing. Set this parameter to 0 when the small image is occupied by a web user.
	StreamType *uint64 `json:"StreamType,omitempty" name:"StreamType"`

	// Width of the output image in pixels. If this parameter is not set, 0 is used by default.
	ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// Height of the output image in pixels. If this parameter is not set, 0 is used by default.
	ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// X offset of the output image in pixels. The sum of `LocationX` and `ImageWidth` cannot exceed the total width of the mixed stream. If this parameter is not set, 0 is used by default.
	LocationX *uint64 `json:"LocationX,omitempty" name:"LocationX"`

	// Y offset of the output image in pixels. The sum of `LocationY` and `ImageHeight` cannot exceed the total height of the mixed stream. If this parameter is not set, 0 is used by default.
	LocationY *uint64 `json:"LocationY,omitempty" name:"LocationY"`

	// Output order of the image. `0` is used if it is left empty.
	ZOrder *uint64 `json:"ZOrder,omitempty" name:"ZOrder"`

	// Render mode of the output image. 0: cropping; 1: scaling; 2: scaling on a black background. If this parameter is not set, 0 is used by default.
	RenderMode *uint64 `json:"RenderMode,omitempty" name:"RenderMode"`

	// Media type of the mixed stream of the user occupying the current position. 0 (default): audio and video; 1: audio; 2: video. You are advised to specify a user ID when using this parameter.
	MixInputType *uint64 `json:"MixInputType,omitempty" name:"MixInputType"`

	// ID of the placeholder image. If the placeholding feature is enabled, and a user for whom an image position is reserved is not sending video data, the position will show the placeholder image. The ID is generated after the placeholder image is uploaded in the TRTC console. https://intl.cloud.tencent.com/document/product/647/50769?from_cn_redirect=1
	PlaceImageId *uint64 `json:"PlaceImageId,omitempty" name:"PlaceImageId"`
}

type PublishCdnParams struct {

	// Tencent Cloud LVB BizId
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Destination of non-Tencent Cloud CDN relayed push. It is possible to push to only one non-Tencent Cloud CDN address at a time.
	PublishCdnUrls []*string `json:"PublishCdnUrls,omitempty" name:"PublishCdnUrls"`
}

type QualityData struct {

	// Data content
	Content []*TimeValue `json:"Content,omitempty" name:"Content"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Peer ID. An empty value indicates that the returned data is upstream.
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`

	// Data type
	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

type RecordUsage struct {

	// Time point for the statistics, e.g., `2020-09-07` or `2020-09-07 00:05:05`
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// SD video duration (s)
	Class1VideoTime *uint64 `json:"Class1VideoTime,omitempty" name:"Class1VideoTime"`

	// HD video duration (s)
	Class2VideoTime *uint64 `json:"Class2VideoTime,omitempty" name:"Class2VideoTime"`

	// FHD video duration (s)
	Class3VideoTime *uint64 `json:"Class3VideoTime,omitempty" name:"Class3VideoTime"`

	// Audio duration (s)
	AudioTime *uint64 `json:"AudioTime,omitempty" name:"AudioTime"`
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

type SdkAppIdRecordUsage struct {

	// Application ID
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Durations for the period queried
	Usages []*RecordUsage `json:"Usages,omitempty" name:"Usages"`
}

type SdkAppIdTrtcMcuTranscodeTimeUsage struct {

	// Time point for the statistics. e.g., `2020-09-07` or `2020-09-07 00:05:05`
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// Audio duration (s)
	AudioTime *uint64 `json:"AudioTime,omitempty" name:"AudioTime"`

	// SD video duration (s)
	VideoTimeSd *uint64 `json:"VideoTimeSd,omitempty" name:"VideoTimeSd"`

	// HD video duration (s)
	VideoTimeHd *uint64 `json:"VideoTimeHd,omitempty" name:"VideoTimeHd"`

	// FHD video duration (s)
	VideoTimeFhd *uint64 `json:"VideoTimeFhd,omitempty" name:"VideoTimeFhd"`
}

type SdkAppIdTrtcUsage struct {

	// Time point for the statistics. e.g., `2020-09-07` or `2020-09-07 00:05:05`
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// Audio duration (s)
	AudioTime *uint64 `json:"AudioTime,omitempty" name:"AudioTime"`

	// Audio/Video duration (s)
	// This parameter is returned only for users who signed up before October 11, 2019 and have not switched to the [new billing standards](https://intl.cloud.tencent.com/document/product/647/17157?from_cn_redirect=1).
	AudioVideoTime *uint64 `json:"AudioVideoTime,omitempty" name:"AudioVideoTime"`

	// SD video duration (s)
	VideoTimeSd *uint64 `json:"VideoTimeSd,omitempty" name:"VideoTimeSd"`

	// HD video duration (s)
	VideoTimeHd *uint64 `json:"VideoTimeHd,omitempty" name:"VideoTimeHd"`

	// FHD video duration (s)
	VideoTimeHdp *uint64 `json:"VideoTimeHdp,omitempty" name:"VideoTimeHdp"`
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

type StartMCUMixTranscodeByStrRoomIdRequest struct {
	*tchttp.BaseRequest

	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID in string type
	StrRoomId *string `json:"StrRoomId,omitempty" name:"StrRoomId"`

	// On-Cloud MixTranscoding output parameters
	OutputParams *OutputParams `json:"OutputParams,omitempty" name:"OutputParams"`

	// On-Cloud MixTranscoding output encoding parameters
	EncodeParams *EncodeParams `json:"EncodeParams,omitempty" name:"EncodeParams"`

	// On-Cloud MixTranscoding output layout parameters
	LayoutParams *LayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// Relayed push parameters of a non-Tencent Cloud CDN
	PublishCdnParams *PublishCdnParams `json:"PublishCdnParams,omitempty" name:"PublishCdnParams"`
}

func (r *StartMCUMixTranscodeByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StrRoomId")
	delete(f, "OutputParams")
	delete(f, "EncodeParams")
	delete(f, "LayoutParams")
	delete(f, "PublishCdnParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMCUMixTranscodeByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartMCUMixTranscodeByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMCUMixTranscodeByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Relayed push parameters of a non-Tencent Cloud CDN
	PublishCdnParams *PublishCdnParams `json:"PublishCdnParams,omitempty" name:"PublishCdnParams"`
}

func (r *StartMCUMixTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "OutputParams")
	delete(f, "EncodeParams")
	delete(f, "LayoutParams")
	delete(f, "PublishCdnParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMCUMixTranscodeRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeByStrRoomIdRequest struct {
	*tchttp.BaseRequest

	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Room ID in string type
	StrRoomId *string `json:"StrRoomId,omitempty" name:"StrRoomId"`
}

func (r *StopMCUMixTranscodeByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StrRoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMCUMixTranscodeByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMCUMixTranscodeByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeByStrRoomIdResponse) FromJsonString(s string) error {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMCUMixTranscodeRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Room ID
	RoomStr *string `json:"RoomStr,omitempty" name:"RoomStr"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The time when the user enters the room
	JoinTs *uint64 `json:"JoinTs,omitempty" name:"JoinTs"`

	// The time when the user exits the room. If the user is still in the room, the current time will be returned
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

type WaterMarkParams struct {

	// Image ID of the watermark, which is generated after the image is uploaded to the TRTC console
	WaterMarkId *uint64 `json:"WaterMarkId,omitempty" name:"WaterMarkId"`

	// Width (px) of the watermark for On-Cloud MixTranscoding
	WaterMarkWidth *uint64 `json:"WaterMarkWidth,omitempty" name:"WaterMarkWidth"`

	// Height (px) of the watermark for On-Cloud MixTranscoding
	WaterMarkHeight *uint64 `json:"WaterMarkHeight,omitempty" name:"WaterMarkHeight"`

	// Horizontal offset (px) of the watermark
	LocationX *uint64 `json:"LocationX,omitempty" name:"LocationX"`

	// Vertical offset (px) of the watermark
	LocationY *uint64 `json:"LocationY,omitempty" name:"LocationY"`

	// URL of the watermark image for the mixed stream, which can be in PNG, JPG, JPEG, or BMP format and does not support the alpha channel. The URL must not exceed 512 bytes. When both `WaterMarkUrl` and `WaterMarkId` are specified, the former will be used. The watermark image cannot exceed 10 MB.
	WaterMarkUrl *string `json:"WaterMarkUrl,omitempty" name:"WaterMarkUrl"`
}
