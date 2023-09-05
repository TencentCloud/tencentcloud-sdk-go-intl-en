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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AppStatisticsItem struct {
	// Voice Chat statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RealtimeSpeechStatisticsItem *RealTimeSpeechStatisticsItem `json:"RealtimeSpeechStatisticsItem,omitnil" name:"RealtimeSpeechStatisticsItem"`

	// Voice Message statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	VoiceMessageStatisticsItem *VoiceMessageStatisticsItem `json:"VoiceMessageStatisticsItem,omitnil" name:"VoiceMessageStatisticsItem"`

	// Phrase Filtering statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	VoiceFilterStatisticsItem *VoiceFilterStatisticsItem `json:"VoiceFilterStatisticsItem,omitnil" name:"VoiceFilterStatisticsItem"`

	// Reference period
	Date *string `json:"Date,omitnil" name:"Date"`

	// Recording-to-Text usage statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AudioTextStatisticsItem *AudioTextStatisticsItem `json:"AudioTextStatisticsItem,omitnil" name:"AudioTextStatisticsItem"`

	// Stream-to-Text usage statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StreamTextStatisticsItem *StreamTextStatisticsItem `json:"StreamTextStatisticsItem,omitnil" name:"StreamTextStatisticsItem"`

	// Usage statistics of Voice-to-Text of outside-MLC requests
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	OverseaTextStatisticsItem *OverseaTextStatisticsItem `json:"OverseaTextStatisticsItem,omitnil" name:"OverseaTextStatisticsItem"`

	// Real-time Voice-to-Text usage statistics
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RealtimeTextStatisticsItem *RealtimeTextStatisticsItem `json:"RealtimeTextStatisticsItem,omitnil" name:"RealtimeTextStatisticsItem"`
}

type ApplicationDataStatistics struct {
	// Application ID
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Number of DAU metrics
	DauDataNum *uint64 `json:"DauDataNum,omitnil" name:"DauDataNum"`

	// DAUs in the Chinese mainland
	DauDataMainland []*StatisticsItem `json:"DauDataMainland,omitnil" name:"DauDataMainland"`

	// DAUs outside the Chinese mainland
	DauDataOversea []*StatisticsItem `json:"DauDataOversea,omitnil" name:"DauDataOversea"`

	// Total DAUs
	DauDataSum []*StatisticsItem `json:"DauDataSum,omitnil" name:"DauDataSum"`

	// Number of Voice Chat metrics
	DurationDataNum *uint64 `json:"DurationDataNum,omitnil" name:"DurationDataNum"`

	// Duration of Voice Chat in the Chinese mainland (in minutes)
	DurationDataMainland []*StatisticsItem `json:"DurationDataMainland,omitnil" name:"DurationDataMainland"`

	// Duration of Voice Chat outside the Chinese mainland (in minutes)
	DurationDataOversea []*StatisticsItem `json:"DurationDataOversea,omitnil" name:"DurationDataOversea"`

	// Total duration of Voice Chat (in minutes)
	DurationDataSum []*StatisticsItem `json:"DurationDataSum,omitnil" name:"DurationDataSum"`

	// Number of PCU metrics
	PcuDataNum *uint64 `json:"PcuDataNum,omitnil" name:"PcuDataNum"`

	// PCUs in the Chinese mainland
	PcuDataMainland []*StatisticsItem `json:"PcuDataMainland,omitnil" name:"PcuDataMainland"`

	// PCUs outside the Chinese mainland
	PcuDataOversea []*StatisticsItem `json:"PcuDataOversea,omitnil" name:"PcuDataOversea"`

	// Total PCUs
	PcuDataSum []*StatisticsItem `json:"PcuDataSum,omitnil" name:"PcuDataSum"`
}

type AsrConf struct {
	// Speech-to-Text status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitnil" name:"Status"`
}

type AudioTextStatisticsItem struct {
	// Statistical value (in seconds)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data *float64 `json:"Data,omitnil" name:"Data"`
}

// Predefined struct for user
type CreateAppRequestParams struct {
	// Application name
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// Tencent Cloud project ID. Default value: 0, which means that the default project is used.
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// List of engines to be supported.
	// Valid values: `android`, `ios`, `unity`, `cocos`, `unreal`, `windows`. All values are selected by default.
	EngineList []*string `json:"EngineList,omitnil" name:"EngineList"`

	// List of regions.
	// Valid values: `mainland` (Chinese mainland), `hmt` (Hong Kong, Macao and Taiwan (China)), `sea` (Southeast Asia), `na` (North America), `eu` (Europe), `jpkr` (Japan, Korea and Asia Pacific), `sa` (South America), `oc` (Oceania), `me` (Middle East). All values are selected by default.
	RegionList []*string `json:"RegionList,omitnil" name:"RegionList"`

	// Configuration information of Voice Chat
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitnil" name:"RealtimeSpeechConf"`

	// Configuration information of Voice Messaging
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitnil" name:"VoiceMessageConf"`

	// Configuration information of Voice Analysis Service
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitnil" name:"VoiceFilterConf"`

	// Configuration information of Speech-to-Text
	AsrConf *AsrConf `json:"AsrConf,omitnil" name:"AsrConf"`

	// List of tags to be added
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// Application name
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// Tencent Cloud project ID. Default value: 0, which means that the default project is used.
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// List of engines to be supported.
	// Valid values: `android`, `ios`, `unity`, `cocos`, `unreal`, `windows`. All values are selected by default.
	EngineList []*string `json:"EngineList,omitnil" name:"EngineList"`

	// List of regions.
	// Valid values: `mainland` (Chinese mainland), `hmt` (Hong Kong, Macao and Taiwan (China)), `sea` (Southeast Asia), `na` (North America), `eu` (Europe), `jpkr` (Japan, Korea and Asia Pacific), `sa` (South America), `oc` (Oceania), `me` (Middle East). All values are selected by default.
	RegionList []*string `json:"RegionList,omitnil" name:"RegionList"`

	// Configuration information of Voice Chat
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitnil" name:"RealtimeSpeechConf"`

	// Configuration information of Voice Messaging
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitnil" name:"VoiceMessageConf"`

	// Configuration information of Voice Analysis Service
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitnil" name:"VoiceFilterConf"`

	// Configuration information of Speech-to-Text
	AsrConf *AsrConf `json:"AsrConf,omitnil" name:"AsrConf"`

	// List of tags to be added
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
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
	delete(f, "AsrConf")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppResp struct {
	// Application ID, automatically generated by the backend.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Application name, the input of `AppName`.
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// Project ID, the input of `ProjectId`.
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Application key, used to initialize GME SDK.
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// Timestamp, indicating when the service is created.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Configuration information of Voice Chat
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitnil" name:"RealtimeSpeechConf"`

	// Configuration information of Voice Messaging
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitnil" name:"VoiceMessageConf"`

	// Configuration information of Voice Analysis Service
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitnil" name:"VoiceFilterConf"`

	// Configuration information of Speech-to-Text
	AsrConf *AsrConf `json:"AsrConf,omitnil" name:"AsrConf"`
}

// Predefined struct for user
type CreateAppResponseParams struct {
	// Returned data
	Data *CreateAppResp `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Code *int64 `json:"Code,omitnil" name:"Code"`

	// Description
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`
}

// Predefined struct for user
type DeleteRoomMemberRequestParams struct {
	// ID of the target room
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// List of the members to remove
	Uids []*string `json:"Uids,omitnil" name:"Uids"`

	// Operation type. `1`: Delete a room; `2`: Remove members
	DeleteType *uint64 `json:"DeleteType,omitnil" name:"DeleteType"`

	// Application ID
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`
}

type DeleteRoomMemberRequest struct {
	*tchttp.BaseRequest
	
	// ID of the target room
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// List of the members to remove
	Uids []*string `json:"Uids,omitnil" name:"Uids"`

	// Operation type. `1`: Delete a room; `2`: Remove members
	DeleteType *uint64 `json:"DeleteType,omitnil" name:"DeleteType"`

	// Application ID
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`
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
	DeleteResult *DeleteResult `json:"DeleteResult,omitnil" name:"DeleteResult"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Data start date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13.
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// Data end date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// List of services to be queried. Valid values: `RealTimeSpeech`, `VoiceMessage`, `VoiceFilter`, `SpeechToText`.
	Services []*string `json:"Services,omitnil" name:"Services"`
}

type DescribeAppStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// GME application ID
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Data start date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13.
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// Data end date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// List of services to be queried. Valid values: `RealTimeSpeech`, `VoiceMessage`, `VoiceFilter`, `SpeechToText`.
	Services []*string `json:"Services,omitnil" name:"Services"`
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
	AppStatistics []*AppStatisticsItem `json:"AppStatistics,omitnil" name:"AppStatistics"`
}

// Predefined struct for user
type DescribeAppStatisticsResponseParams struct {
	// Application usage statistics
	Data *DescribeAppStatisticsResp `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Data start date in the format of yyyy-mm-dd, such as 2018-07-13.
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// Data end date in the format of yyyy-mm-dd, such as 2018-07-13.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

type DescribeApplicationDataRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Data start date in the format of yyyy-mm-dd, such as 2018-07-13.
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// Data end date in the format of yyyy-mm-dd, such as 2018-07-13.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
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
	Data *ApplicationDataStatistics `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeRecordInfoRequestParams struct {
	// ID of the ongoing task, which is returned from the `StartRecord` API.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`
}

type DescribeRecordInfoRequest struct {
	*tchttp.BaseRequest
	
	// ID of the ongoing task, which is returned from the `StartRecord` API.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`
}

func (r *DescribeRecordInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordInfoResponseParams struct {
	// Information about the recording task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecordInfo []*RecordInfo `json:"RecordInfo,omitnil" name:"RecordInfo"`

	// Recording mode. Valid values: `1`: single stream; `2`: mixed streams; `3`: single stream and mixed streams.
	RecordMode *uint64 `json:"RecordMode,omitnil" name:"RecordMode"`

	// Room ID.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordInfoResponseParams `json:"Response"`
}

func (r *DescribeRecordInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoRequestParams struct {
	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Room ID.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Room ID.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoResponseParams struct {
	// ID of the ongoing task, which is returned from the `StartRecord` API.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// Recording mode. Valid values: `1`: single stream; `2`: mixed streams; `3`: single stream and mixed streams.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecordMode *uint64 `json:"RecordMode,omitnil" name:"RecordMode"`

	// Allowlist or blocklist for stream subscription.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil" name:"SubscribeRecordUserIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInfoResponseParams `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppStatusRequestParams struct {
	// Application ID, which is generated and returned by the backend after the application creation.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Application status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitnil" name:"Status"`
}

type ModifyAppStatusRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is generated and returned by the backend after the application creation.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Application status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitnil" name:"Status"`
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
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Application status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitnil" name:"Status"`
}

// Predefined struct for user
type ModifyAppStatusResponseParams struct {
	// Returned data
	Data *ModifyAppStatusResp `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type ModifyRecordInfoRequestParams struct {
	// ID of the ongoing task, which is returned from the `StartRecord` API.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// Recording mode. Valid values: `1`: single stream; `2`: mixed streams; `3`: single stream and mixed streams.
	RecordMode *uint64 `json:"RecordMode,omitnil" name:"RecordMode"`

	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Allowlist or blocklist for stream subscription.
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil" name:"SubscribeRecordUserIds"`
}

type ModifyRecordInfoRequest struct {
	*tchttp.BaseRequest
	
	// ID of the ongoing task, which is returned from the `StartRecord` API.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// Recording mode. Valid values: `1`: single stream; `2`: mixed streams; `3`: single stream and mixed streams.
	RecordMode *uint64 `json:"RecordMode,omitnil" name:"RecordMode"`

	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Allowlist or blocklist for stream subscription.
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil" name:"SubscribeRecordUserIds"`
}

func (r *ModifyRecordInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "RecordMode")
	delete(f, "BizId")
	delete(f, "SubscribeRecordUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordInfoResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRecordInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordInfoResponseParams `json:"Response"`
}

func (r *ModifyRecordInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverseaTextStatisticsItem struct {
	// Statistical value (in seconds)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data *float64 `json:"Data,omitnil" name:"Data"`
}

type RealTimeSpeechStatisticsItem struct {
	// DAUs in the Chinese mainland
	MainLandDau *uint64 `json:"MainLandDau,omitnil" name:"MainLandDau"`

	// PCUs in the Chinese mainland
	MainLandPcu *uint64 `json:"MainLandPcu,omitnil" name:"MainLandPcu"`

	// Total duration of use in the Chinese mainland (in minutes)
	MainLandDuration *uint64 `json:"MainLandDuration,omitnil" name:"MainLandDuration"`

	// DAUs outside the Chinese mainland
	OverseaDau *uint64 `json:"OverseaDau,omitnil" name:"OverseaDau"`

	// PCUs outside the Chinese mainland
	OverseaPcu *uint64 `json:"OverseaPcu,omitnil" name:"OverseaPcu"`

	// Total duration of use outside the Chinese mainland (in minutes)
	OverseaDuration *uint64 `json:"OverseaDuration,omitnil" name:"OverseaDuration"`
}

type RealtimeSpeechConf struct {
	// Voice Chat status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitnil" name:"Status"`

	// Voice Chat sound quality type. Valid values: `high` (HD), `ordinary` (SD).
	Quality *string `json:"Quality,omitnil" name:"Quality"`
}

type RealtimeTextStatisticsItem struct {
	// Statistical value (in seconds)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data *float64 `json:"Data,omitnil" name:"Data"`
}

type RecordInfo struct {
	// User ID. The value is `0` in mixed streams recording mode.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// Recording filename.
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// Recording start time, which is a Unix timestamp. Example: 1234567868.
	RecordBeginTime *uint64 `json:"RecordBeginTime,omitnil" name:"RecordBeginTime"`

	// Recording status. Valid values: `2`: recording; `10`: to be transcoded; `11`: transcoding; `12`: uploading; `13`: uploaded; `14`: user notified.
	RecordStatus *uint64 `json:"RecordStatus,omitnil" name:"RecordStatus"`
}

type SceneInfo struct {

}

// Predefined struct for user
type StartRecordRequestParams struct {
	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Room ID.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// Recording mode. Valid values: `1`: single stream; `2`: mixed streams; `3`: single stream and mixed streams.
	RecordMode *uint64 `json:"RecordMode,omitnil" name:"RecordMode"`

	// Allowlist or blocklist for stream subscription. If you do not specify this parameter, the audio streams of all the users in the room are subscribed to by default.
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil" name:"SubscribeRecordUserIds"`
}

type StartRecordRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`

	// Room ID.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// Recording mode. Valid values: `1`: single stream; `2`: mixed streams; `3`: single stream and mixed streams.
	RecordMode *uint64 `json:"RecordMode,omitnil" name:"RecordMode"`

	// Allowlist or blocklist for stream subscription. If you do not specify this parameter, the audio streams of all the users in the room are subscribed to by default.
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil" name:"SubscribeRecordUserIds"`
}

func (r *StartRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomId")
	delete(f, "RecordMode")
	delete(f, "SubscribeRecordUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartRecordResponseParams struct {
	// Task ID.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartRecordResponse struct {
	*tchttp.BaseResponse
	Response *StartRecordResponseParams `json:"Response"`
}

func (r *StartRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticsItem struct {
	// Date in the format of yyyy-mm-dd, such as 2018-07-13
	StatDate *string `json:"StatDate,omitnil" name:"StatDate"`

	// Statistical value
	Data *uint64 `json:"Data,omitnil" name:"Data"`
}

// Predefined struct for user
type StopRecordRequestParams struct {
	// Task ID.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`
}

type StopRecordRequest struct {
	*tchttp.BaseRequest
	
	// Task ID.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// Application ID.
	BizId *uint64 `json:"BizId,omitnil" name:"BizId"`
}

func (r *StopRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRecordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopRecordResponse struct {
	*tchttp.BaseResponse
	Response *StopRecordResponseParams `json:"Response"`
}

func (r *StopRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StreamTextStatisticsItem struct {
	// Usage of the service (in seconds)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data *float64 `json:"Data,omitnil" name:"Data"`
}

type SubscribeRecordUserIds struct {
	// Blocklist for audio subscription. For example, `["1", "2", "3"]` means to not subscribe to the audio streams of users 1, 2, and 3. If this parameter is left empty, the audio streams of all users (max 20) in the room will not be subscribed to.
	// Note: You cannot specify `UnSubscribeAudioUserIds` and `SubscribeAudioUserIds` at the same time.
	UnSubscribeUserIds []*string `json:"UnSubscribeUserIds,omitnil" name:"UnSubscribeUserIds"`

	// Allowlist for audio subscription. For example, `["1", "2", "3"]` means to subscribe to the audio streams of users 1, 2, and 3. If this parameter is left empty, the audio streams of all users (max 20) in the room will be subscribed to.
	// Note: You cannot specify `UnSubscribeAudioUserIds` and `SubscribeAudioUserIds` at the same time.
	SubscribeUserIds []*string `json:"SubscribeUserIds,omitnil" name:"SubscribeUserIds"`
}

type Tag struct {
	// Tag key
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type VoiceFilterConf struct {
	// Phrase Filtering status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitnil" name:"Status"`

	// Scenario configuration information, such as status and callback URL.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SceneInfos []*SceneInfo `json:"SceneInfos,omitnil" name:"SceneInfos"`
}

type VoiceFilterStatisticsItem struct {
	// Total duration of phrase filtering (in minutes)
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`
}

type VoiceMessageConf struct {
	// Voice Message Service status. Valid values: `open`, `close`.
	Status *string `json:"Status,omitnil" name:"Status"`

	// Language supported for Voice Message Service. Valid values: `all` (all languages), `cnen` (Chinese and English). Default value: `cnen`.
	Language *string `json:"Language,omitnil" name:"Language"`
}

type VoiceMessageStatisticsItem struct {
	// DAUs of Voice Message Service
	Dau *uint64 `json:"Dau,omitnil" name:"Dau"`
}