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

package v20200326

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AttachedInputInfo struct {

	// Media input ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Audio selector for media input. Quantity limit: [0,20]
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioSelectors []*AudioSelectorInfo `json:"AudioSelectors,omitempty" name:"AudioSelectors" list`
}

type AudioPidSelectionInfo struct {

	// Audio `Pid`. Default value: 0.
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

type AudioPipelineInputStatistics struct {

	// Audio FPS.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Audio bitrate in bps.
	Rate *uint64 `json:"Rate,omitempty" name:"Rate"`

	// Audio `Pid`, which is available only if the input is `rtp/udp`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
}

type AudioSelectorInfo struct {

	// Audio name, which can contain 1–32 letters, digits, and underscores.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Audio `Pid` selection.
	AudioPidSelection *AudioPidSelectionInfo `json:"AudioPidSelection,omitempty" name:"AudioPidSelection"`
}

type AudioTemplateInfo struct {

	// Only `AttachedInputs.AudioSelectors.Name` can be selected. This parameter is required for RTP_PUSH and UDP_PUSH.
	AudioSelectorName *string `json:"AudioSelectorName,omitempty" name:"AudioSelectorName"`

	// Audio transcoding template name, which can contain 1–20 letters and digits.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Audio codec. Valid value: AAC. Default value: AAC.
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// Audio bitrate. If this parameter is left empty, the original value will be used.
	// Valid values: 6000, 7000, 8000, 10000, 12000, 14000, 16000, 20000, 24000, 28000, 32000, 40000, 48000, 56000, 64000, 80000, 96000, 112000, 128000, 160000, 192000, 224000, 256000, 288000, 320000, 384000, 448000, 512000, 576000, 640000, 768000, 896000, 1024000
	AudioBitrate *uint64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// Audio language code, whose length is always 3 characters.
	LanguageCode *string `json:"LanguageCode,omitempty" name:"LanguageCode"`
}

type ChannelAlertInfos struct {

	// Alarm details of pipeline 0 under this channel.
	Pipeline0 []*ChannelPipelineAlerts `json:"Pipeline0,omitempty" name:"Pipeline0" list`

	// Alarm details of pipeline 1 under this channel.
	Pipeline1 []*ChannelPipelineAlerts `json:"Pipeline1,omitempty" name:"Pipeline1" list`
}

type ChannelInfo struct {

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel status.
	State *string `json:"State,omitempty" name:"State"`

	// Information of associated input.
	AttachedInputs []*AttachedInputInfo `json:"AttachedInputs,omitempty" name:"AttachedInputs" list`

	// Information of output group.
	OutputGroups []*OutputGroupsInfo `json:"OutputGroups,omitempty" name:"OutputGroups" list`

	// Channel name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Audio transcoding template array.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioTemplates []*AudioTemplateInfo `json:"AudioTemplates,omitempty" name:"AudioTemplates" list`

	// Video transcoding template array.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoTemplates []*VideoTemplateInfo `json:"VideoTemplates,omitempty" name:"VideoTemplates" list`
}

type ChannelInputStatistics struct {

	// Input ID.
	InputId *string `json:"InputId,omitempty" name:"InputId"`

	// Input statistics.
	Statistics *InputStatistics `json:"Statistics,omitempty" name:"Statistics"`
}

type ChannelOutputsStatistics struct {

	// Output group name.
	OutputGroupName *string `json:"OutputGroupName,omitempty" name:"OutputGroupName"`

	// Output group statistics.
	Statistics *OutputsStatistics `json:"Statistics,omitempty" name:"Statistics"`
}

type ChannelPipelineAlerts struct {

	// Alarm start time in UTC time.
	SetTime *string `json:"SetTime,omitempty" name:"SetTime"`

	// Alarm end time in UTC time.
	// This time is available only after the alarm ends.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClearTime *string `json:"ClearTime,omitempty" name:"ClearTime"`

	// Alarm type.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Alarm details.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type CreateMediaLiveChannelRequest struct {
	*tchttp.BaseRequest

	// Channel name, which can contain 1–32 letters, digits, and underscores and must be unique at the region level.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Associated media input. Quantity limit: [1,1].
	AttachedInputs []*AttachedInputInfo `json:"AttachedInputs,omitempty" name:"AttachedInputs" list`

	// Configuration information of channel output groups. Quantity limit: [1,10].
	OutputGroups []*OutputGroupsInfo `json:"OutputGroups,omitempty" name:"OutputGroups" list`

	// Audio transcoding template array. Quantity limit: [1,20].
	AudioTemplates []*AudioTemplateInfo `json:"AudioTemplates,omitempty" name:"AudioTemplates" list`

	// Video transcoding template array. Quantity limit: [1,10].
	VideoTemplates []*VideoTemplateInfo `json:"VideoTemplates,omitempty" name:"VideoTemplates" list`
}

func (r *CreateMediaLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaLiveChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel ID.
		Id *string `json:"Id,omitempty" name:"Id"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMediaLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaLiveChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaLiveInputRequest struct {
	*tchttp.BaseRequest

	// Media input name, which can contain 1–32 letters, digits, and underscores and must be unique at the region level.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Media input type.
	// Valid values: RTMP_PUSH/RTP_PUSH/UDP_PUSH/RTMP_PULL/HLS_PULL/MP4_PULL.
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of the input security group to be bound.
	// Only one security group can be associated.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// Input settings information, two sets of which need to be configured for RTMP_PUSH/RTMP_PULL/HLS_PULL/MP4_PULL.
	InputSettings []*InputSettingInfo `json:"InputSettings,omitempty" name:"InputSettings" list`
}

func (r *CreateMediaLiveInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaLiveInputRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaLiveInputResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Media input ID.
		Id *string `json:"Id,omitempty" name:"Id"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMediaLiveInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaLiveInputResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaLiveInputSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// Input security group name, which can contain letters, digits, and underscores and must be unique at the region level.
	Name *string `json:"Name,omitempty" name:"Name"`

	// List of whitelist entries. Quantity limit: [1,10].
	Whitelist []*string `json:"Whitelist,omitempty" name:"Whitelist" list`
}

func (r *CreateMediaLiveInputSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaLiveInputSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaLiveInputSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group ID.
		Id *string `json:"Id,omitempty" name:"Id"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMediaLiveInputSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaLiveInputSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DashRemuxSettingsInfo struct {

	// Segment duration in ms. Value range: [1000,30000]. Default value: 4000. The value can only be a multiple of 1,000.
	SegmentDuration *uint64 `json:"SegmentDuration,omitempty" name:"SegmentDuration"`

	// Number of segments. Value range: [1,30]. Default value: 5.
	SegmentNumber *uint64 `json:"SegmentNumber,omitempty" name:"SegmentNumber"`

	// Whether to enable multi-period. Valid values: CLOSE/OPEN. Default value: CLOSE.
	PeriodTriggers *string `json:"PeriodTriggers,omitempty" name:"PeriodTriggers"`
}

type DeleteMediaLiveChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteMediaLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaLiveChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaLiveChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaLiveInputRequest struct {
	*tchttp.BaseRequest

	// Media input ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteMediaLiveInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaLiveInputRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaLiveInputResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaLiveInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaLiveInputResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaLiveInputSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// Input security group ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteMediaLiveInputSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaLiveInputSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaLiveInputSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaLiveInputSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaLiveInputSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelAlertsRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *DescribeMediaLiveChannelAlertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelAlertsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelAlertsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm information of two pipelines under this channel.
		Infos *ChannelAlertInfos `json:"Infos,omitempty" name:"Infos"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaLiveChannelAlertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelAlertsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelInputStatisticsRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// Statistics start time, which is one hour ago by default. Maximum value: the last 7 days.
	// UTC time, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time, which is one hour after `StartTime` by default.
	// UTC time, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeMediaLiveChannelInputStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelInputStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelInputStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel input statistics array.
		Infos []*ChannelInputStatistics `json:"Infos,omitempty" name:"Infos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaLiveChannelInputStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelInputStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelOutputStatisticsRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// Statistics start time, which is one hour ago by default. Maximum value: the last 7 days.
	// UTC time, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time, which is one hour after `StartTime` by default.
	// UTC time, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeMediaLiveChannelOutputStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelOutputStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelOutputStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel output information.
		Infos []*ChannelOutputsStatistics `json:"Infos,omitempty" name:"Infos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaLiveChannelOutputStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelOutputStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeMediaLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel information.
		Info *ChannelInfo `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMediaLiveChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveChannelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel list information.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Infos []*ChannelInfo `json:"Infos,omitempty" name:"Infos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaLiveChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveChannelsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveInputRequest struct {
	*tchttp.BaseRequest

	// Media input ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeMediaLiveInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveInputRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveInputResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// MediaLive input information.
		Info *InputInfo `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaLiveInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveInputResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveInputSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// Input security group ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeMediaLiveInputSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveInputSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveInputSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Input security group information.
		Info *InputSecurityGroupInfo `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaLiveInputSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveInputSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveInputSecurityGroupsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMediaLiveInputSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveInputSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveInputSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Input security group information list.
		Infos []*InputSecurityGroupInfo `json:"Infos,omitempty" name:"Infos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaLiveInputSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveInputSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveInputsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMediaLiveInputsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveInputsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaLiveInputsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// MediaLive input information list.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Infos []*InputInfo `json:"Infos,omitempty" name:"Infos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaLiveInputsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaLiveInputsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestinationInfo struct {

	// Relay destination address. Length limit: [1,512].
	OutputUrl *string `json:"OutputUrl,omitempty" name:"OutputUrl"`

	// Authentication key. Length limit: [1,128].
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// Authentication username. Length limit: [1,128].
	// Note: this field may return null, indicating that no valid values can be obtained.
	Username *string `json:"Username,omitempty" name:"Username"`

	// Authentication password. Length limit: [1,128].
	// Note: this field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitempty" name:"Password"`
}

type DrmSettingsInfo struct {

	// Whether to enable DRM encryption. Valid value: CLOSE/OPEN. Default value: CLOSE.
	// Currently, this is supported only for HLS/DASH/HLS_ARCHIVE/DASH_ARCHIVE.
	State *string `json:"State,omitempty" name:"State"`

	// `ContentId` of DRM encryption, which will be automatically created if this parameter is left empty.
	// For more information on the custom creation method, please visit https://cloud.tencent.com/document/product/1000/40960
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`
}

type HlsRemuxSettingsInfo struct {

	// Segment duration in ms. Value range: [1000,30000]. Default value: 4000. The value can only be a multiple of 1,000.
	SegmentDuration *uint64 `json:"SegmentDuration,omitempty" name:"SegmentDuration"`

	// Number of segments. Value range: [1,30]. Default value: 5.
	SegmentNumber *uint64 `json:"SegmentNumber,omitempty" name:"SegmentNumber"`

	// Whether to enable PDT insertion. Valid values: CLOSE/OPEN. Default value: CLOSE.
	PdtInsertion *string `json:"PdtInsertion,omitempty" name:"PdtInsertion"`

	// PDT duration in seconds. Value range: (0,3000]. Default value: 600.
	PdtDuration *uint64 `json:"PdtDuration,omitempty" name:"PdtDuration"`
}

type InputInfo struct {

	// Input region.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Input ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Input name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Input type.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Array of security groups associated with input.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// Array of channels associated with input.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AttachedChannels []*string `json:"AttachedChannels,omitempty" name:"AttachedChannels" list`

	// Input configuration array.
	InputSettings []*InputSettingInfo `json:"InputSettings,omitempty" name:"InputSettings" list`
}

type InputSecurityGroupInfo struct {

	// Input security group ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Input security group name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// List of whitelist entries.
	Whitelist []*string `json:"Whitelist,omitempty" name:"Whitelist" list`

	// List of bound input streams.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OccupiedInputs []*string `json:"OccupiedInputs,omitempty" name:"OccupiedInputs" list`

	// Input security group address.
	Region *string `json:"Region,omitempty" name:"Region"`
}

type InputSettingInfo struct {

	// Application name, which is used for RTMP_PUSH and can contain 1–32 letters and digits.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name, which is used for RTMP_PUSH and can contain 1–32 letters and digits.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Origin-pull URL, which is used for RTMP_PULL/HLS_PULL/MP4_PULL. Length limit: [1,512].
	// Note: this field may return null, indicating that no valid values can be obtained.
	SourceUrl *string `json:"SourceUrl,omitempty" name:"SourceUrl"`

	// RTP/UDP input address, which does not need to be entered for the input parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InputAddress *string `json:"InputAddress,omitempty" name:"InputAddress"`
}

type InputStatistics struct {

	// Input statistics of pipeline 0.
	Pipeline0 []*PipelineInputStatistics `json:"Pipeline0,omitempty" name:"Pipeline0" list`

	// Input statistics of pipeline 1.
	Pipeline1 []*PipelineInputStatistics `json:"Pipeline1,omitempty" name:"Pipeline1" list`
}

type ModifyMediaLiveChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel name, which can contain 1–32 letters, digits, and underscores and must be unique at the region level.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Associated media input. Quantity limit: [1,1].
	AttachedInputs []*AttachedInputInfo `json:"AttachedInputs,omitempty" name:"AttachedInputs" list`

	// Configuration information of channel output groups. Quantity limit: [1,10].
	OutputGroups []*OutputGroupsInfo `json:"OutputGroups,omitempty" name:"OutputGroups" list`

	// Audio transcoding template array. Quantity limit: [1,20].
	AudioTemplates []*AudioTemplateInfo `json:"AudioTemplates,omitempty" name:"AudioTemplates" list`

	// Video transcoding template array. Quantity limit: [1,10].
	VideoTemplates []*VideoTemplateInfo `json:"VideoTemplates,omitempty" name:"VideoTemplates" list`
}

func (r *ModifyMediaLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaLiveChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaLiveChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaLiveInputRequest struct {
	*tchttp.BaseRequest

	// Media input ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Media input name, which can contain 1–32 letters, digits, and underscores and must be unique at the region level.
	Name *string `json:"Name,omitempty" name:"Name"`

	// List of IDs of bound security groups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *ModifyMediaLiveInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaLiveInputRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaLiveInputResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaLiveInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaLiveInputResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaLiveInputSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// Input security group ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Input security group name, which can contain 1–32 letters, digits, and underscores and must be unique at the region level.
	Name *string `json:"Name,omitempty" name:"Name"`

	// List of whitelist entries. Up to 10 entries are allowed.
	Whitelist []*string `json:"Whitelist,omitempty" name:"Whitelist" list`
}

func (r *ModifyMediaLiveInputSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaLiveInputSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaLiveInputSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaLiveInputSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaLiveInputSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OutputGroupsInfo struct {

	// Channel output group name, which can contain 1–32 letters, digits, and underscores and must be unique at the channel level.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Output protocol type.
	// Valid values: HLS/DASH/RTP/RTMP/HLS_ARCHIVE/DASH_ARCHIVE.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Output information.
	// Quantity limit: [1,1] for RTMP/RTP; [1,10] for HLS/DASH.
	Outputs []*OutputInfo `json:"Outputs,omitempty" name:"Outputs" list`

	// Relay destination address. Quantity limit: [1,2].
	Destinations []*DestinationInfo `json:"Destinations,omitempty" name:"Destinations" list`

	// HLS protocol configuration information, which takes effect only for HLS/HLS_ARCHIVE.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HlsRemuxSettings *HlsRemuxSettingsInfo `json:"HlsRemuxSettings,omitempty" name:"HlsRemuxSettings"`

	// DASH protocol configuration information, which takes effect only for DASH/DSAH_ARCHIVE.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DashRemuxSettings *DashRemuxSettingsInfo `json:"DashRemuxSettings,omitempty" name:"DashRemuxSettings"`

	// DRM configuration information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DrmSettings *DrmSettingsInfo `json:"DrmSettings,omitempty" name:"DrmSettings"`
}

type OutputInfo struct {

	// Output name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Audio transcoding template name array.
	// Quantity limit: [0,1] for RTMP; [0,20] for others.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioTemplateNames []*string `json:"AudioTemplateNames,omitempty" name:"AudioTemplateNames" list`

	// Video transcoding template name array. Quantity limit: [0,1].
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoTemplateNames []*string `json:"VideoTemplateNames,omitempty" name:"VideoTemplateNames" list`

	// SCTE-35 information configuration.
	Scte35Settings *Scte35SettingsInfo `json:"Scte35Settings,omitempty" name:"Scte35Settings"`
}

type OutputsStatistics struct {

	// Output information of pipeline 0.
	Pipeline0 []*PipelineOutputStatistics `json:"Pipeline0,omitempty" name:"Pipeline0" list`

	// Output information of pipeline 1.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pipeline1 []*PipelineOutputStatistics `json:"Pipeline1,omitempty" name:"Pipeline1" list`
}

type PipelineInputStatistics struct {

	// Data timestamp in seconds.
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Input bandwidth in bps.
	NetworkIn *uint64 `json:"NetworkIn,omitempty" name:"NetworkIn"`

	// Video information array.
	// For `rtp/udp` input, the quantity is the number of `Pid` of the input video.
	// For other inputs, the quantity is 1.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Video []*VideoPipelineInputStatistics `json:"Video,omitempty" name:"Video" list`

	// Audio information array.
	// For `rtp/udp` input, the quantity is the number of `Pid` of the input audio.
	// For other inputs, the quantity is 1.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Audio []*AudioPipelineInputStatistics `json:"Audio,omitempty" name:"Audio" list`
}

type PipelineOutputStatistics struct {

	// Timestamp.
	// In seconds, indicating data time.
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Output bandwidth in bps.
	NetworkOut *uint64 `json:"NetworkOut,omitempty" name:"NetworkOut"`
}

type Scte35SettingsInfo struct {

	// Whether to pass through SCTE-35 information. Valid values: NO_PASSTHROUGH/PASSTHROUGH. Default value: NO_PASSTHROUGH.
	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`
}

type StartMediaLiveChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *StartMediaLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMediaLiveChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartMediaLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMediaLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMediaLiveChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMediaLiveChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *StopMediaLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMediaLiveChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMediaLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMediaLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMediaLiveChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VideoPipelineInputStatistics struct {

	// Video FPS.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Video bitrate in bps.
	Rate *uint64 `json:"Rate,omitempty" name:"Rate"`

	// Video `Pid`, which is available only if the input is `rtp/udp`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
}

type VideoTemplateInfo struct {

	// Video transcoding template name, which can contain 1–20 letters and digits.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Video codec. Valid values: H264/H265. If this parameter is left empty, the original value will be used.
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// Video bitrate. Value range: [50000,40000000]. The value can only be a multiple of 1,000. If this parameter is left empty, the original value will be used.
	VideoBitrate *uint64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Video width. Value range: (0,3000]. The value can only be a multiple of 4. If this parameter is left empty, the original value will be used.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Video height. Value range: (0,3000]. The value can only be a multiple of 4. If this parameter is left empty, the original value will be used.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Video frame rate. Value range: [1,240]. If this parameter is left empty, the original value will be used.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Whether to enable top speed codec. Valid value: CLOSE/OPEN. Default value: CLOSE.
	TopSpeed *string `json:"TopSpeed,omitempty" name:"TopSpeed"`

	// Top speed codec compression ratio. Value range: [0,50]. The lower the compression ratio, the higher the image quality.
	BitrateCompressionRatio *uint64 `json:"BitrateCompressionRatio,omitempty" name:"BitrateCompressionRatio"`
}
