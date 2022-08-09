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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AppStatisticsItem struct {
	// Voice chat statistics
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealtimeSpeechStatisticsItem *RealTimeSpeechStatisticsItem `json:"RealtimeSpeechStatisticsItem,omitempty" name:"RealtimeSpeechStatisticsItem"`

	// Voice messaging statistics
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoiceMessageStatisticsItem *VoiceMessageStatisticsItem `json:"VoiceMessageStatisticsItem,omitempty" name:"VoiceMessageStatisticsItem"`

	// Phrase filtering statistics
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoiceFilterStatisticsItem *VoiceFilterStatisticsItem `json:"VoiceFilterStatisticsItem,omitempty" name:"VoiceFilterStatisticsItem"`

	// Statistical period
	Date *string `json:"Date,omitempty" name:"Date"`
}

type ApplicationDataStatistics struct {
	// Application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// DAU data
	DauDataNum *uint64 `json:"DauDataNum,omitempty" name:"DauDataNum"`

	// DAU in Chinese mainland
	DauDataMainland []*StatisticsItem `json:"DauDataMainland,omitempty" name:"DauDataMainland"`

	// DAU outside Chinese mainland
	DauDataOversea []*StatisticsItem `json:"DauDataOversea,omitempty" name:"DauDataOversea"`

	// Total DAU
	DauDataSum []*StatisticsItem `json:"DauDataSum,omitempty" name:"DauDataSum"`

	// Number of voice chat metrics
	DurationDataNum *uint64 `json:"DurationDataNum,omitempty" name:"DurationDataNum"`

	// Duration of voice chat in Chinese mainland in minutes
	DurationDataMainland []*StatisticsItem `json:"DurationDataMainland,omitempty" name:"DurationDataMainland"`

	// Duration of voice chat outside Chinese mainland in minutes
	DurationDataOversea []*StatisticsItem `json:"DurationDataOversea,omitempty" name:"DurationDataOversea"`

	// Total duration of voice chat in minutes
	DurationDataSum []*StatisticsItem `json:"DurationDataSum,omitempty" name:"DurationDataSum"`

	// PCU data
	PcuDataNum *uint64 `json:"PcuDataNum,omitempty" name:"PcuDataNum"`

	// PCU in Chinese mainland
	PcuDataMainland []*StatisticsItem `json:"PcuDataMainland,omitempty" name:"PcuDataMainland"`

	// PCU outside Chinese mainland
	PcuDataOversea []*StatisticsItem `json:"PcuDataOversea,omitempty" name:"PcuDataOversea"`

	// Total PCU
	PcuDataSum []*StatisticsItem `json:"PcuDataSum,omitempty" name:"PcuDataSum"`
}

// Predefined struct for user
type CreateAppRequestParams struct {
	// Application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Tencent Cloud project ID. Default value: 0, which means the default project
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// List of engines to be supported. All values are selected by default.
	EngineList []*string `json:"EngineList,omitempty" name:"EngineList"`

	// Service region list. All values are selected by default.
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// Configuration information of voice chat
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// Configuration information of voice messaging and speech-to-text
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// Configuration information of phrase analysis
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitempty" name:"VoiceFilterConf"`

	// List of tags to be added
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// Application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Tencent Cloud project ID. Default value: 0, which means the default project
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// List of engines to be supported. All values are selected by default.
	EngineList []*string `json:"EngineList,omitempty" name:"EngineList"`

	// Service region list. All values are selected by default.
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// Configuration information of voice chat
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// Configuration information of voice messaging and speech-to-text
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// Configuration information of phrase analysis
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
	// App ID, automatically generated by the backend.
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// App name, the input of `AppName`.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Project ID, the input of `ProjectId`.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// App key, used to initialize GME SDK.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Timestamp, indicating when the service is created.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Configuration information of voice chat
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// Configuration information of voice messaging and speech-to-text
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// Configuration information of phrase analysis
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

// Predefined struct for user
type DescribeAppStatisticsRequestParams struct {
	// GME application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Data start date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// Data end date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// List of services to be queried. Valid values: RealTimeSpeech, VoiceMessage, VoiceFilter
	Services []*string `json:"Services,omitempty" name:"Services"`
}

type DescribeAppStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// GME application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Data start date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// Data end date (GMT+8) in the format of yyyy-mm-dd, such as 2018-07-13
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// List of services to be queried. Valid values: RealTimeSpeech, VoiceMessage, VoiceFilter
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
	// App usage statistics
	AppStatistics []*AppStatisticsItem `json:"AppStatistics,omitempty" name:"AppStatistics"`
}

// Predefined struct for user
type DescribeAppStatisticsResponseParams struct {
	// App usage statistics
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

	// Data start date in the format of yyyy-mm-dd, such as 2018-07-13
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// Data end date in the format of yyyy-mm-dd, such as 2018-07-13
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeApplicationDataRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Data start date in the format of yyyy-mm-dd, such as 2018-07-13
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// Data end date in the format of yyyy-mm-dd, such as 2018-07-13
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
	// Note: this field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Detection result, which will be returned if `Code` is 0
	// Note: this field may return null, indicating that no valid values can be obtained.
	ScanPiece []*ScanPiece `json:"ScanPiece,omitempty" name:"ScanPiece"`

	// Detection task submission timestamp
	ScanStartTime *uint64 `json:"ScanStartTime,omitempty" name:"ScanStartTime"`

	// Speech detection scenario, which corresponds to the `Scene` at the time of request
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// Speech detection task ID, which is assigned by the backend
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// File or stream address
	Url *string `json:"Url,omitempty" name:"Url"`

	// Detection task execution result task. Valid values:
	// <li>Start: task started</li>
	// <li>Success: task ended successfully</li>
	// <li>Error: exceptional</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// ID of the application submitted for detection
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`
}

// Predefined struct for user
type DescribeScanResultListRequestParams struct {
	// Application ID, which is the `AppID` obtained when you create an application in the [console](https://console.cloud.tencent.com/gamegme)
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// List of IDs of the tasks to be queried. Up to 100 entries can be added in the ID list.
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// Number of task results to be returned. Default value: 10. Maximum value: 500. This parameter will be ignored for large file tasks where all results will be returned
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeScanResultListRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is the `AppID` obtained when you create an application in the [console](https://console.cloud.tencent.com/gamegme)
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// List of IDs of the tasks to be queried. Up to 100 entries can be added in the ID list.
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// Number of task results to be returned. Default value: 10. Maximum value: 500. This parameter will be ignored for large file tasks where all results will be returned
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
	// Result of the speech detection task to be queried
	// Note: this field may return null, indicating that no valid values can be obtained.
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
	// Application ID, which is generated and returned by the backend after application creation.
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Application status. Valid values: open, close
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyAppStatusRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is generated and returned by the backend after application creation.
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Application status. Valid values: open, close
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
	// GME app ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// App status. Valid values: `open`, `close`
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

type RealTimeSpeechStatisticsItem struct {
	// DAU in Mainland China
	MainLandDau *uint64 `json:"MainLandDau,omitempty" name:"MainLandDau"`

	// PCU in Mainland China
	MainLandPcu *uint64 `json:"MainLandPcu,omitempty" name:"MainLandPcu"`

	// Total duration of use in Mainland China in minutes
	MainLandDuration *uint64 `json:"MainLandDuration,omitempty" name:"MainLandDuration"`

	// DAU outside Mainland China
	OverseaDau *uint64 `json:"OverseaDau,omitempty" name:"OverseaDau"`

	// PCU outside Mainland China
	OverseaPcu *uint64 `json:"OverseaPcu,omitempty" name:"OverseaPcu"`

	// Total duration of use outside Mainland China in minutes
	OverseaDuration *uint64 `json:"OverseaDuration,omitempty" name:"OverseaDuration"`
}

type RealtimeSpeechConf struct {
	// Voice chat status. Valid values: open, close
	Status *string `json:"Status,omitempty" name:"Status"`

	// Voice chat sound quality. Valid value: `high`
	Quality *string `json:"Quality,omitempty" name:"Quality"`
}

type ScanDetail struct {
	// Violation scenario. For more information, please see the definition of <a href="https://intl.cloud.tencent.com/document/product/607/37622?from_cn_redirect=1#Label_Value">Label</a>
	Label *string `json:"Label,omitempty" name:"Label"`

	// Confidence score in scenario. Value range: [0.00,100.00]. The higher the score, the more likely the content is non-compliant
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// Non-compliant keyword
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// Start time offset in milliseconds from 0 of keyword in audio
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time offset in milliseconds from 0 of keyword in audio
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type ScanPiece struct {
	// Audio retention address, which will be returned for stream detection. The audio will be retained for 30 minutes
	// Note: this field may return null, indicating that no valid values can be obtained.
	DumpUrl *string `json:"DumpUrl,omitempty" name:"DumpUrl"`

	// Whether non-compliant information is detected
	HitFlag *bool `json:"HitFlag,omitempty" name:"HitFlag"`

	// Main non-compliant content type
	// Note: this field may return null, indicating that no valid values can be obtained.
	MainType *string `json:"MainType,omitempty" name:"MainType"`

	// Speech detection details
	ScanDetail []*ScanDetail `json:"ScanDetail,omitempty" name:"ScanDetail"`

	// GME voice chat room ID, which is the `RoomId` passed through when the task was submitted
	// Note: this field may return null, indicating that no valid values can be obtained.
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// GME voice chat user ID, which is the `OpenId` passed through when the task was submitted
	// Note: this field may return null, indicating that no valid values can be obtained.
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Info *string `json:"Info,omitempty" name:"Info"`

	// Offset time in milliseconds of segment in stream during stream detection
	// Note: this field may return null, indicating that no valid values can be obtained.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Segment duration during stream detection
	// Note: this field may return null, indicating that no valid values can be obtained.
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Segment detection start time
	// Note: this field may return null, indicating that no valid values can be obtained.
	PieceStartTime *uint64 `json:"PieceStartTime,omitempty" name:"PieceStartTime"`
}

// Predefined struct for user
type ScanVoiceRequestParams struct {
	// Application ID, which is the `AppID` obtained when you create an application in [Console > Service Management](https://console.cloud.tencent.com/gamegme)
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Speech detection scenario. The value of this parameter is currently required to be `default`. Preset scenarios: abusive, pornographic, advertising, and other prohibited scenarios. For specific values, please see the <a href="#Label_Value">Label description</a> above.
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// Whether it is a live stream. false: audio file detection, true: audio stream detection.
	Live *bool `json:"Live,omitempty" name:"Live"`

	// Speech detection task list. Up to 100 tasks can be added in the list. The structure contains:
	// <li>DataId: unique data ID</li>
	// <li>Url: URL-encoded data file URL, which is a pull address if the detected speech is a stream</li>
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// Async callback address for detection result. For more information, please see the <a href="#Callback_Declare">callback description</a> above. (Note: if this field is empty, the detection result can only be obtained by calling the `DescribeScanResultList` API.)
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// The language. `jp` represents Japanese
	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

type ScanVoiceRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is the `AppID` obtained when you create an application in [Console > Service Management](https://console.cloud.tencent.com/gamegme)
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Speech detection scenario. The value of this parameter is currently required to be `default`. Preset scenarios: abusive, pornographic, advertising, and other prohibited scenarios. For specific values, please see the <a href="#Label_Value">Label description</a> above.
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// Whether it is a live stream. false: audio file detection, true: audio stream detection.
	Live *bool `json:"Live,omitempty" name:"Live"`

	// Speech detection task list. Up to 100 tasks can be added in the list. The structure contains:
	// <li>DataId: unique data ID</li>
	// <li>Url: URL-encoded data file URL, which is a pull address if the detected speech is a stream</li>
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// Async callback address for detection result. For more information, please see the <a href="#Callback_Declare">callback description</a> above. (Note: if this field is empty, the detection result can only be obtained by calling the `DescribeScanResultList` API.)
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// The language. `jp` represents Japanese
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
	// Speech detection return. The `Data` field is a JSON array where each element contains: <li>DataId: corresponding `DataId` in request.</li>
	// <li>TaskID: detection task ID, which is used to poll the speech detection result.</li>
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

	// Statistics
	Data *uint64 `json:"Data,omitempty" name:"Data"`
}

type Tag struct {
	// Tag key
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Task struct {
	// Unique data ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// URL-encoded data file URL, which is a pull address if the detected speech is a stream
	Url *string `json:"Url,omitempty" name:"Url"`

	// GME voice chat room ID, which is entered during speech detection by GME voice chat
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// GME voice chat user ID, which is entered during speech detection by GME voice chat
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

type VoiceFilterConf struct {
	// Phrase filtering status. Valid values: open, close
	Status *string `json:"Status,omitempty" name:"Status"`
}

type VoiceFilterStatisticsItem struct {
	// Total duration of phrase filtering
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type VoiceMessageConf struct {
	// Voice messaging and speech-to-text status. Valid values: open, close
	Status *string `json:"Status,omitempty" name:"Status"`

	// Language supported for voice messaging and speech-to-text. Valid values: all (all languages), cnen (Chinese and English). Default value: cnen
	Language *string `json:"Language,omitempty" name:"Language"`
}

type VoiceMessageStatisticsItem struct {
	// DAU of voice messaging and speech-to-text
	Dau *uint64 `json:"Dau,omitempty" name:"Dau"`
}