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

package v20180711

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AppStatisticsItem struct {
	// Voice Chat statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RealtimeSpeechStatisticsItem *RealTimeSpeechStatisticsItem `json:"RealtimeSpeechStatisticsItem,omitempty" name:"RealtimeSpeechStatisticsItem"`

	// Voice Message statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	VoiceMessageStatisticsItem *VoiceMessageStatisticsItem `json:"VoiceMessageStatisticsItem,omitempty" name:"VoiceMessageStatisticsItem"`

	// Phrase Filtering statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	VoiceFilterStatisticsItem *VoiceFilterStatisticsItem `json:"VoiceFilterStatisticsItem,omitempty" name:"VoiceFilterStatisticsItem"`

	// Reference period
	Date *string `json:"Date,omitempty" name:"Date"`

	// Recording-to-Text usage statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AudioTextStatisticsItem *AudioTextStatisticsItem `json:"AudioTextStatisticsItem,omitempty" name:"AudioTextStatisticsItem"`

	// Stream-to-Text usage statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StreamTextStatisticsItem *StreamTextStatisticsItem `json:"StreamTextStatisticsItem,omitempty" name:"StreamTextStatisticsItem"`

	// Usage statistics of Voice-to-Text of outside-MLC requests
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	OverseaTextStatisticsItem *OverseaTextStatisticsItem `json:"OverseaTextStatisticsItem,omitempty" name:"OverseaTextStatisticsItem"`

	// Real-time Voice-to-Text usage statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RealtimeTextStatisticsItem *RealtimeTextStatisticsItem `json:"RealtimeTextStatisticsItem,omitempty" name:"RealtimeTextStatisticsItem"`
}

type ApplicationDataStatistics struct {
	// Application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Number of DAU metrics
	DauDataNum *uint64 `json:"DauDataNum,omitempty" name:"DauDataNum"`

	// DAUs in the Chinese mainland
	DauDataMainland []*StatisticsItem `json:"DauDataMainland,omitempty" name:"DauDataMainland"`

	// DAUs outside the Chinese mainland
	DauDataOversea []*StatisticsItem `json:"DauDataOversea,omitempty" name:"DauDataOversea"`

	// Total DAUs
	DauDataSum []*StatisticsItem `json:"DauDataSum,omitempty" name:"DauDataSum"`

	// Number of Voice Chat metrics
	DurationDataNum *uint64 `json:"DurationDataNum,omitempty" name:"DurationDataNum"`

	// Duration of Voice Chat in the Chinese mainland (in minutes)
	DurationDataMainland []*StatisticsItem `json:"DurationDataMainland,omitempty" name:"DurationDataMainland"`

	// Duration of Voice Chat outside the Chinese mainland (in minutes)
	DurationDataOversea []*StatisticsItem `json:"DurationDataOversea,omitempty" name:"DurationDataOversea"`

	// Total duration of Voice Chat (in minutes)
	DurationDataSum []*StatisticsItem `json:"DurationDataSum,omitempty" name:"DurationDataSum"`

	// Number of PCU metrics
	PcuDataNum *uint64 `json:"PcuDataNum,omitempty" name:"PcuDataNum"`

	// PCUs in the Chinese mainland
	PcuDataMainland []*StatisticsItem `json:"PcuDataMainland,omitempty" name:"PcuDataMainland"`

	// PCUs outside the Chinese mainland
	PcuDataOversea []*StatisticsItem `json:"PcuDataOversea,omitempty" name:"PcuDataOversea"`

	// Total PCUs
	PcuDataSum []*StatisticsItem `json:"PcuDataSum,omitempty" name:"PcuDataSum"`
}

type AudioTextStatisticsItem struct {
	// Statistical value (in seconds)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

// Predefined struct for user
type CreateAppRequestParams struct {
	// Application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Tencent Cloud project ID. Default value: 0, which means that the default project is used.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// List of engines to be supported. All values are selected by default.
	EngineList []*string `json:"EngineList,omitempty" name:"EngineList"`

	// Service region list. All values are selected by default.
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// Configuration information of Voice Chat
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// Configuration information of Voice Message Service
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// Configuration information of Voice Analysis Service
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitempty" name:"VoiceFilterConf"`

	// List of tags to be added
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// Application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Tencent Cloud project ID. Default value: 0, which means that the default project is used.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// List of engines to be supported. All values are selected by default.
	EngineList []*string `json:"EngineList,omitempty" name:"EngineList"`

	// Service region list. All values are selected by default.
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// Configuration information of Voice Chat
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// Configuration information of Voice Message Service
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// Configuration information of Voice Analysis Service
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitempty" name:"VoiceFilterConf"`

	// List of tags to be added
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "ProjectId")
	delete(f, "EngineList")
	delete(f, "RegionList")
	delete(f, "RealtimeSpeechConf")
	delete(f, "VoiceMessageConf")
	delete(f, "VoiceFilterConf")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppResp struct {
	// Application ID, automatically generated by the backend.
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Application name, the input of `AppName`.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Project ID, the input of `ProjectId`.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Application key, used to initialize GME SDK.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Timestamp, indicating when the service is created.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Configuration information of Voice Chat
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// Configuration information of Voice Message Service
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// Configuration information of Voice Analysis Service
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitempty" name:"VoiceFilterConf"`
}

// Predefined struct for user
type CreateAppResponseParams struct {
	// Returned data
	Data *CreateAppResp `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppResponseParams `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResult struct {
	// Status code. `0`: Succeeded. Others: Failed\
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Description
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

// Predefined struct for user
type DeleteRoomMemberRequestParams struct {
	// ID of the target room
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// List of the members to remove
	Uids []*string `json:"Uids,omitempty" name:"Uids"`

	// Operation type. `1`: Delete a room; `2`: Remove members
	DeleteType *uint64 `json:"DeleteType,omitempty" name:"DeleteType"`

	// Application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`
}

type DeleteRoomMemberRequest struct {
	*tchttp.BaseRequest
	
	// ID of the target room
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// List of the members to remove
	Uids []*string `json:"Uids,omitempty" name:"Uids"`

	// Operation type. `1`: Delete a room; `2`: Remove members
	DeleteType *uint64 `json:"DeleteType,omitempty" name:"DeleteType"`

	// Application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`
}

func (r *DeleteRoomMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "Uids")
	delete(f, "DeleteType")
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoomMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoomMemberResponseParams struct {
	// Result of the operation to delete a room or remove a member
	DeleteResult *DeleteResult `json:"DeleteResult,omitempty" name:"DeleteResult"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRoomMemberResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoomMemberResponseParams `json:"Response"`
}

func (r *DeleteRoomMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppStatisticsRequestParams struct {
	// GME application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Data start date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// Data end date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// List of services to be queried. Valid values: `RealTimeSpeech`, `VoiceMessage`, `VoiceFilter`, `SpeechToText`.
	Services []*string `json:"Services,omitempty" name:"Services"`
}

type DescribeAppStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// GME application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Data start date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// Data end date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// List of services to be queried. Valid values: `RealTimeSpeech`, `VoiceMessage`, `VoiceFilter`, `SpeechToText`.
	Services []*string `json:"Services,omitempty" name:"Services"`
}

func (r *DescribeAppStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Services")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppStatisticsResp struct {
	// Application usage statistics
	AppStatistics []*AppStatisticsItem `json:"AppStatistics,omitempty" name:"AppStatistics"`
}

// Predefined struct for user
type DescribeAppStatisticsResponseParams struct {
	// Application usage statistics
	Data *DescribeAppStatisticsResp `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAppStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppStatisticsResponseParams `json:"Response"`
}

func (r *DescribeAppStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationDataRequestParams struct {
	// Application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Data start date in the format of yyyy-mm-dd, such as 2018-07-13.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// Data end date in the format of yyyy-mm-dd, such as 2018-07-13.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeApplicationDataRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Data start date in the format of yyyy-mm-dd, such as 2018-07-13.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// Data end date in the format of yyyy-mm-dd, such as 2018-07-13.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeApplicationDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationDataResponseParams struct {
	// Application statistics
	Data *ApplicationDataStatistics `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationDataResponseParams `json:"Response"`
}

func (r *DescribeApplicationDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppStatusRequestParams struct {
	// Application ID, which is generated and returned by the backend after the application creation.
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Application status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyAppStatusRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is generated and returned by the backend after the application creation.
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Application status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAppStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppStatusResp struct {
	// GME application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Application status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type ModifyAppStatusResponseParams struct {
	// Returned data
	Data *ModifyAppStatusResp `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAppStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppStatusResponseParams `json:"Response"`
}

func (r *ModifyAppStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverseaTextStatisticsItem struct {
	// Statistical value (in seconds)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

type RealTimeSpeechStatisticsItem struct {
	// DAUs in the Chinese mainland
	MainLandDau *uint64 `json:"MainLandDau,omitempty" name:"MainLandDau"`

	// PCUs in the Chinese mainland
	MainLandPcu *uint64 `json:"MainLandPcu,omitempty" name:"MainLandPcu"`

	// Total duration of use in the Chinese mainland (in minutes)
	MainLandDuration *uint64 `json:"MainLandDuration,omitempty" name:"MainLandDuration"`

	// DAUs outside the Chinese mainland
	OverseaDau *uint64 `json:"OverseaDau,omitempty" name:"OverseaDau"`

	// PCUs outside the Chinese mainland
	OverseaPcu *uint64 `json:"OverseaPcu,omitempty" name:"OverseaPcu"`

	// Total duration of use outside the Chinese mainland (in minutes)
	OverseaDuration *uint64 `json:"OverseaDuration,omitempty" name:"OverseaDuration"`
}

type RealtimeSpeechConf struct {
	// Voice Chat status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Voice Chat sound quality. Valid value: `high`.
	Quality *string `json:"Quality,omitempty" name:"Quality"`
}

type RealtimeTextStatisticsItem struct {
	// Statistical value (in seconds)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

type StatisticsItem struct {
	// Date in the format of yyyy-mm-dd, such as 2018-07-13
	StatDate *string `json:"StatDate,omitempty" name:"StatDate"`

	// Statistical value
	Data *uint64 `json:"Data,omitempty" name:"Data"`
}

type StreamTextStatisticsItem struct {
	// Usage of the service (in seconds)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

type Tag struct {
	// Tag key
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type VoiceFilterConf struct {
	// Phrase Filtering status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type VoiceFilterStatisticsItem struct {
	// Total duration of phrase filtering (in minutes)
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type VoiceMessageConf struct {
	// Voice Message Service status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Language supported for Voice Message Service. Valid values: `all` (all languages), `cnen` (Chinese and English). Default value: `cnen`.
	Language *string `json:"Language,omitempty" name:"Language"`
}

type VoiceMessageStatisticsItem struct {
	// DAUs of Voice Message Service
	Dau *uint64 `json:"Dau,omitempty" name:"Dau"`
}