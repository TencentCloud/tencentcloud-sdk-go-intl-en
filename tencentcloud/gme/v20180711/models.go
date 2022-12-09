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

type DescribeScanResult struct {
	// Business return code
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Unique data ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// Detection completion timestamp
	ScanFinishTime *uint64 `json:"ScanFinishTime,omitempty" name:"ScanFinishTime"`

	// Whether non-compliant information is detected
	HitFlag *bool `json:"HitFlag,omitempty" name:"HitFlag"`

	// Whether it is a stream
	Live *bool `json:"Live,omitempty" name:"Live"`

	// Business return description
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Detection result, which will be returned if `Code` is 0
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ScanPiece []*ScanPiece `json:"ScanPiece,omitempty" name:"ScanPiece"`

	// Detection task submission timestamp
	ScanStartTime *uint64 `json:"ScanStartTime,omitempty" name:"ScanStartTime"`

	// Voice detection scenario, which corresponds to the `Scene` at the time of request
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// Voice detection task ID, which is assigned by the backend
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// File or stream address
	Url *string `json:"Url,omitempty" name:"Url"`

	// Detection task execution result status. Valid values:
	// <li>Start: Task started</li>
	// <li>Success: Task ended successfully</li>
	// <li>Error: An exception occurs</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// ID of the application submitted for detection
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`
}

// Predefined struct for user
type DescribeScanResultListRequestParams struct {
	// Application ID, which is obtained when you create an application in the [GME console](https://console.cloud.tencent.com/gamegme).
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// List of IDs of the tasks to be queried. Up to 100 entries can be added in the ID list.
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// Number of task results to be returned. Default value: 10. Maximum value: 500. This parameter will be ignored for large file tasks where all results will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeScanResultListRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is obtained when you create an application in the [GME console](https://console.cloud.tencent.com/gamegme).
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// List of IDs of the tasks to be queried. Up to 100 entries can be added in the ID list.
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// Number of task results to be returned. Default value: 10. Maximum value: 500. This parameter will be ignored for large file tasks where all results will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeScanResultListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanResultListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "TaskIdList")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanResultListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanResultListResponseParams struct {
	// Result of the voice detection task to be queried
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data []*DescribeScanResult `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeScanResultListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanResultListResponseParams `json:"Response"`
}

func (r *DescribeScanResultListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanResultListResponse) FromJsonString(s string) error {
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

type ScanDetail struct {
	// Violation scenario. For more information, please see the definition of <a href="https://intl.cloud.tencent.com/document/product/607/37622?from_cn_redirect=1#Label_Value">Label</a>.
	Label *string `json:"Label,omitempty" name:"Label"`

	// Confidence score in scenario. Value range: [0.00,100.00]. The higher the score, the more likely the content is non-compliant.
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// Non-compliant keyword
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// Start time offset of keyword from 0 in audio (in milliseconds)
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time offset of keyword from 0 in audio (in milliseconds)
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type ScanPiece struct {
	// Audio retention address, which will be returned for stream detection. The audio will be retained for 30 minutes.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DumpUrl *string `json:"DumpUrl,omitempty" name:"DumpUrl"`

	// Whether non-compliant information is detected
	HitFlag *bool `json:"HitFlag,omitempty" name:"HitFlag"`

	// Main non-compliant content type
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	MainType *string `json:"MainType,omitempty" name:"MainType"`

	// Voice detection details
	ScanDetail []*ScanDetail `json:"ScanDetail,omitempty" name:"ScanDetail"`

	// GME Voice Chat room ID, which is the `RoomId` passed through when the task was submitted.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// GME Voice Chat user ID, which is the `OpenId` passed through when the task was submitted.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// Remarks
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Info *string `json:"Info,omitempty" name:"Info"`

	// Offset time of multipart in stream during stream detection (in milliseconds)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Multipart duration during stream detection
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Multipart detection start time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PieceStartTime *uint64 `json:"PieceStartTime,omitempty" name:"PieceStartTime"`
}

// Predefined struct for user
type ScanVoiceRequestParams struct {
	// Application ID, which is obtained when you create an application in the [GME console - Service Management](https://console.cloud.tencent.com/gamegme).
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Voice detection scenario. It must be `default`. See the <a href="#Label_Value">Label description</a> as the result.
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// Whether it is a live stream. Values: `false` (voice file), `true` (live stream).
	Live *bool `json:"Live,omitempty" name:"Live"`

	// Voice detection task list. Up to 100 tasks can be added in the list. 
	// <li>`DataId`: Unique data ID</li>
	// <li>`Url`: URL-encoded data file URL, which is a pull address if the detected voice is a stream</li>
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// Async callback address for detection result. For more information, please see the <a href="#Callback_Declare">Callback description</a> above. (Note: If this field is empty, the detection result can only be queried by calling the `DescribeScanResultList` API.)
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// Language. Chinese will be used if it is left empty.
	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

type ScanVoiceRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is obtained when you create an application in the [GME console - Service Management](https://console.cloud.tencent.com/gamegme).
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Voice detection scenario. It must be `default`. See the <a href="#Label_Value">Label description</a> as the result.
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// Whether it is a live stream. Values: `false` (voice file), `true` (live stream).
	Live *bool `json:"Live,omitempty" name:"Live"`

	// Voice detection task list. Up to 100 tasks can be added in the list. 
	// <li>`DataId`: Unique data ID</li>
	// <li>`Url`: URL-encoded data file URL, which is a pull address if the detected voice is a stream</li>
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// Async callback address for detection result. For more information, please see the <a href="#Callback_Declare">Callback description</a> above. (Note: If this field is empty, the detection result can only be queried by calling the `DescribeScanResultList` API.)
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// Language. Chinese will be used if it is left empty.
	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *ScanVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "Scenes")
	delete(f, "Live")
	delete(f, "Tasks")
	delete(f, "Callback")
	delete(f, "Lang")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanVoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanVoiceResponseParams struct {
	// Voice moderation result. <li>`DataId`: Corresponding `DataId` in request.</li>
	// <li>`TaskID`: Moderation task ID, which is used to poll the voice detection result.</li>
	Data []*ScanVoiceResult `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanVoiceResponse struct {
	*tchttp.BaseResponse
	Response *ScanVoiceResponseParams `json:"Response"`
}

func (r *ScanVoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVoiceResult struct {
	// Data ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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

type Task struct {
	// Unique data ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// URL-encoded data file URL, which is a pull address if the detected voice is a stream.
	Url *string `json:"Url,omitempty" name:"Url"`

	// GME Voice Chat room ID, which is entered during voice analysis by GME Voice Chat.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// GME Voice Chat user ID, which is entered during voice analysis by GME Voice Chat.
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
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