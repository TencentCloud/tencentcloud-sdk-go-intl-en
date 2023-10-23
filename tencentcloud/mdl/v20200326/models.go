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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AVTemplate struct {
	// Name of an audio/video transcoding template, which can contain 1-20 case-sensitive letters and digits
	Name *string `json:"Name,omitnil" name:"Name"`

	// Whether video is needed. `0`: not needed; `1`: needed
	NeedVideo *uint64 `json:"NeedVideo,omitnil" name:"NeedVideo"`

	// Video codec. Valid values: `H264`, `H265`. If this parameter is left empty, the original video codec will be used.
	Vcodec *string `json:"Vcodec,omitnil" name:"Vcodec"`

	// Video width. Value range: (0, 3000]. The value must be an integer multiple of 4. If this parameter is left empty, the original video width will be used.
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// Video height. Value range: (0, 3000]. The value must be an integer multiple of 4. If this parameter is left empty, the original video height will be used.
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// Video frame rate. Value range: [1, 240]. If this parameter is left empty, the original frame rate will be used.
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// Whether to enable top speed codec transcoding. Valid values: `CLOSE` (disable), `OPEN` (enable). Default value: `CLOSE`
	TopSpeed *string `json:"TopSpeed,omitnil" name:"TopSpeed"`

	// Compression ratio for top speed codec transcoding. Value range: [0, 50]. The lower the compression ratio, the higher the image quality.
	BitrateCompressionRatio *uint64 `json:"BitrateCompressionRatio,omitnil" name:"BitrateCompressionRatio"`

	// Whether audio is needed. `0`: not needed; `1`: needed
	NeedAudio *int64 `json:"NeedAudio,omitnil" name:"NeedAudio"`

	// Audio codec. Valid value: `AAC` (default)
	Acodec *string `json:"Acodec,omitnil" name:"Acodec"`

	// Audio bitrate. If this parameter is left empty, the original bitrate will be used.
	// Valid values: `6000`, `7000`, `8000`, `10000`, `12000`, `14000`, `16000`, `20000`, `24000`, `28000`, `32000`, `40000`, `48000`, `56000`, `64000`, `80000`, `96000`, `112000`, `128000`, `160000`, `192000`, `224000`, `256000`, `288000`, `320000`, `384000`, `448000`, `512000`, `576000`, `640000`, `768000`, `896000`, `1024000`
	AudioBitrate *uint64 `json:"AudioBitrate,omitnil" name:"AudioBitrate"`

	// Video bitrate. Value range: [50000, 40000000]. The value must be an integer multiple of 1000. If this parameter is left empty, the original bitrate will be used.
	VideoBitrate *uint64 `json:"VideoBitrate,omitnil" name:"VideoBitrate"`

	// Bitrate control mode. Valid values: `CBR`, `ABR` (default)
	RateControlMode *string `json:"RateControlMode,omitnil" name:"RateControlMode"`

	// Watermark ID
	WatermarkId *string `json:"WatermarkId,omitnil" name:"WatermarkId"`

	// Whether to convert audio to text. `0` (default): No; `1`: Yes.
	SmartSubtitles *uint64 `json:"SmartSubtitles,omitnil" name:"SmartSubtitles"`

	// The subtitle settings. Currently, the following subtitles are supported:
	// `eng2eng`: English speech to English text.
	// `eng2chs`: English speech to Chinese text. 
	// `eng2chseng`: English speech to English and Chinese text. 
	// `chs2chs`: Chinese speech to Chinese text.   
	// `chs2eng`: Chinese speech to English text. 
	// `chs2chseng`: Chinese speech to Chinese and English text.
	SubtitleConfiguration *string `json:"SubtitleConfiguration,omitnil" name:"SubtitleConfiguration"`
}

type AttachedInput struct {
	// Input ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Audio selector for the input. There can be 0 to 20 audio selectors.
	// Note: this field may return `null`, indicating that no valid value was found.
	AudioSelectors []*AudioSelectorInfo `json:"AudioSelectors,omitnil" name:"AudioSelectors"`

	// Pull mode. If the input type is `HLS_PULL` or `MP4_PULL`, you can set this parameter to `LOOP` or `ONCE`. `LOOP` is the default value.
	// Note: this field may return `null`, indicating that no valid value was found.
	PullBehavior *string `json:"PullBehavior,omitnil" name:"PullBehavior"`

	// Input failover configuration
	// Note: this field may return `null`, indicating that no valid value was found.
	FailOverSettings *FailOverSettings `json:"FailOverSettings,omitnil" name:"FailOverSettings"`
}

type AudioPidSelectionInfo struct {
	// Audio `Pid`. Default value: 0.
	Pid *uint64 `json:"Pid,omitnil" name:"Pid"`
}

type AudioPipelineInputStatistics struct {
	// Audio FPS.
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// Audio bitrate in bps.
	Rate *uint64 `json:"Rate,omitnil" name:"Rate"`

	// Audio `Pid`, which is available only if the input is `rtp/udp`.
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`
}

type AudioSelectorInfo struct {
	// Audio name, which can contain 1-32 letters, digits, and underscores.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Audio `Pid` selection.
	AudioPidSelection *AudioPidSelectionInfo `json:"AudioPidSelection,omitnil" name:"AudioPidSelection"`
}

type AudioTemplateInfo struct {
	// Only `AttachedInputs.AudioSelectors.Name` can be selected. This parameter is required for RTP_PUSH and UDP_PUSH.
	AudioSelectorName *string `json:"AudioSelectorName,omitnil" name:"AudioSelectorName"`

	// Audio transcoding template name, which can contain 1-20 letters and digits.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Audio codec. Valid value: AAC. Default value: AAC.
	Acodec *string `json:"Acodec,omitnil" name:"Acodec"`

	// Audio bitrate. If this parameter is left empty, the original value will be used.
	// Valid values: 6000, 7000, 8000, 10000, 12000, 14000, 16000, 20000, 24000, 28000, 32000, 40000, 48000, 56000, 64000, 80000, 96000, 112000, 128000, 160000, 192000, 224000, 256000, 288000, 320000, 384000, 448000, 512000, 576000, 640000, 768000, 896000, 1024000
	AudioBitrate *uint64 `json:"AudioBitrate,omitnil" name:"AudioBitrate"`

	// Audio language code, whose length is always 3 characters.
	LanguageCode *string `json:"LanguageCode,omitnil" name:"LanguageCode"`
}

type ChannelAlertInfos struct {
	// Alarm details of pipeline 0 under this channel.
	Pipeline0 []*ChannelPipelineAlerts `json:"Pipeline0,omitnil" name:"Pipeline0"`

	// Alarm details of pipeline 1 under this channel.
	Pipeline1 []*ChannelPipelineAlerts `json:"Pipeline1,omitnil" name:"Pipeline1"`
}

type ChannelInputStatistics struct {
	// Input ID.
	InputId *string `json:"InputId,omitnil" name:"InputId"`

	// Input statistics.
	Statistics *InputStatistics `json:"Statistics,omitnil" name:"Statistics"`
}

type ChannelOutputsStatistics struct {
	// Output group name.
	OutputGroupName *string `json:"OutputGroupName,omitnil" name:"OutputGroupName"`

	// Output group statistics.
	Statistics *OutputsStatistics `json:"Statistics,omitnil" name:"Statistics"`
}

type ChannelPipelineAlerts struct {
	// Alarm start time in UTC time.
	SetTime *string `json:"SetTime,omitnil" name:"SetTime"`

	// Alarm end time in UTC time.
	// This time is available only after the alarm ends.
	ClearTime *string `json:"ClearTime,omitnil" name:"ClearTime"`

	// Alarm type.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Alarm details.
	Message *string `json:"Message,omitnil" name:"Message"`
}

type CreateImageSettings struct {
	// Image file format. Valid values: png, jpg.
	ImageType *string `json:"ImageType,omitnil" name:"ImageType"`

	// Base64 encoded image content
	ImageContent *string `json:"ImageContent,omitnil" name:"ImageContent"`

	// Origin. Valid values: TOP_LEFT, BOTTOM_LEFT, TOP_RIGHT, BOTTOM_RIGHT.
	Location *string `json:"Location,omitnil" name:"Location"`

	// The watermark’s horizontal distance from the origin as a percentage of the video width. Value range: 0-100. Default: 10.
	XPos *int64 `json:"XPos,omitnil" name:"XPos"`

	// The watermark’s vertical distance from the origin as a percentage of the video height. Value range: 0-100. Default: 10.
	YPos *int64 `json:"YPos,omitnil" name:"YPos"`

	// The watermark image’s width as a percentage of the video width. Value range: 0-100. Default: 10.
	// `0` means to scale the width proportionally to the height.
	// You cannot set both `Width` and `Height` to `0`.
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// The watermark image’s height as a percentage of the video height. Value range: 0-100. Default: 10.
	// `0` means to scale the height proportionally to the width.
	// You cannot set both `Width` and `Height` to `0`.
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

// Predefined struct for user
type CreateStreamLiveChannelRequestParams struct {
	// Channel name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Inputs to attach. You can attach 1 to 5 inputs.
	AttachedInputs []*AttachedInput `json:"AttachedInputs,omitnil" name:"AttachedInputs"`

	// Configuration information of the channel’s output groups. Quantity: [1, 10]
	OutputGroups []*StreamLiveOutputGroupsInfo `json:"OutputGroups,omitnil" name:"OutputGroups"`

	// Audio transcoding templates. Quantity: [1, 20]
	AudioTemplates []*AudioTemplateInfo `json:"AudioTemplates,omitnil" name:"AudioTemplates"`

	// Video transcoding templates. Quantity: [1, 10]
	VideoTemplates []*VideoTemplateInfo `json:"VideoTemplates,omitnil" name:"VideoTemplates"`

	// Audio/Video transcoding templates. Quantity: [1, 10]
	AVTemplates []*AVTemplate `json:"AVTemplates,omitnil" name:"AVTemplates"`

	// Event settings
	PlanSettings *PlanSettings `json:"PlanSettings,omitnil" name:"PlanSettings"`

	// The callback settings.
	EventNotifySettings *EventNotifySetting `json:"EventNotifySettings,omitnil" name:"EventNotifySettings"`

	// Complement the last video frame settings.
	InputLossBehavior *InputLossBehaviorInfo `json:"InputLossBehavior,omitnil" name:"InputLossBehavior"`
}

type CreateStreamLiveChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Inputs to attach. You can attach 1 to 5 inputs.
	AttachedInputs []*AttachedInput `json:"AttachedInputs,omitnil" name:"AttachedInputs"`

	// Configuration information of the channel’s output groups. Quantity: [1, 10]
	OutputGroups []*StreamLiveOutputGroupsInfo `json:"OutputGroups,omitnil" name:"OutputGroups"`

	// Audio transcoding templates. Quantity: [1, 20]
	AudioTemplates []*AudioTemplateInfo `json:"AudioTemplates,omitnil" name:"AudioTemplates"`

	// Video transcoding templates. Quantity: [1, 10]
	VideoTemplates []*VideoTemplateInfo `json:"VideoTemplates,omitnil" name:"VideoTemplates"`

	// Audio/Video transcoding templates. Quantity: [1, 10]
	AVTemplates []*AVTemplate `json:"AVTemplates,omitnil" name:"AVTemplates"`

	// Event settings
	PlanSettings *PlanSettings `json:"PlanSettings,omitnil" name:"PlanSettings"`

	// The callback settings.
	EventNotifySettings *EventNotifySetting `json:"EventNotifySettings,omitnil" name:"EventNotifySettings"`

	// Complement the last video frame settings.
	InputLossBehavior *InputLossBehaviorInfo `json:"InputLossBehavior,omitnil" name:"InputLossBehavior"`
}

func (r *CreateStreamLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AttachedInputs")
	delete(f, "OutputGroups")
	delete(f, "AudioTemplates")
	delete(f, "VideoTemplates")
	delete(f, "AVTemplates")
	delete(f, "PlanSettings")
	delete(f, "EventNotifySettings")
	delete(f, "InputLossBehavior")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLiveChannelResponseParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLiveChannelResponseParams `json:"Response"`
}

func (r *CreateStreamLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLiveInputRequestParams struct {
	// Input name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Input type
	// Valid values: `RTMP_PUSH`, `RTP_PUSH`, `UDP_PUSH`, `RTMP_PULL`, `HLS_PULL`, `MP4_PULL`
	Type *string `json:"Type,omitnil" name:"Type"`

	// ID of the input security group to attach
	// You can attach only one security group to an input.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Input settings. For the type `RTMP_PUSH`, `RTMP_PULL`, `HLS_PULL`, or `MP4_PULL`, 1 or 2 inputs of the corresponding type can be configured.
	InputSettings []*InputSettingInfo `json:"InputSettings,omitnil" name:"InputSettings"`
}

type CreateStreamLiveInputRequest struct {
	*tchttp.BaseRequest
	
	// Input name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Input type
	// Valid values: `RTMP_PUSH`, `RTP_PUSH`, `UDP_PUSH`, `RTMP_PULL`, `HLS_PULL`, `MP4_PULL`
	Type *string `json:"Type,omitnil" name:"Type"`

	// ID of the input security group to attach
	// You can attach only one security group to an input.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Input settings. For the type `RTMP_PUSH`, `RTMP_PULL`, `HLS_PULL`, or `MP4_PULL`, 1 or 2 inputs of the corresponding type can be configured.
	InputSettings []*InputSettingInfo `json:"InputSettings,omitnil" name:"InputSettings"`
}

func (r *CreateStreamLiveInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLiveInputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "SecurityGroupIds")
	delete(f, "InputSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLiveInputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLiveInputResponseParams struct {
	// Input ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamLiveInputResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLiveInputResponseParams `json:"Response"`
}

func (r *CreateStreamLiveInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLiveInputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLiveInputSecurityGroupRequestParams struct {
	// Input security group name, which can contain case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Allowlist entries. Quantity: [1, 10]
	Whitelist []*string `json:"Whitelist,omitnil" name:"Whitelist"`
}

type CreateStreamLiveInputSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Input security group name, which can contain case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Allowlist entries. Quantity: [1, 10]
	Whitelist []*string `json:"Whitelist,omitnil" name:"Whitelist"`
}

func (r *CreateStreamLiveInputSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLiveInputSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Whitelist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLiveInputSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLiveInputSecurityGroupResponseParams struct {
	// Security group ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamLiveInputSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLiveInputSecurityGroupResponseParams `json:"Response"`
}

func (r *CreateStreamLiveInputSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLiveInputSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLivePlanRequestParams struct {
	// ID of the channel for which you want to configure an event
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Event configuration
	Plan *PlanReq `json:"Plan,omitnil" name:"Plan"`
}

type CreateStreamLivePlanRequest struct {
	*tchttp.BaseRequest
	
	// ID of the channel for which you want to configure an event
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Event configuration
	Plan *PlanReq `json:"Plan,omitnil" name:"Plan"`
}

func (r *CreateStreamLivePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLivePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Plan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLivePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLivePlanResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamLivePlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLivePlanResponseParams `json:"Response"`
}

func (r *CreateStreamLivePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLivePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLiveWatermarkRequestParams struct {
	// Watermark name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Watermark type. Valid values: STATIC_IMAGE, TEXT.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Watermark image settings. This parameter is valid if `Type` is `STATIC_IMAGE`.
	ImageSettings *CreateImageSettings `json:"ImageSettings,omitnil" name:"ImageSettings"`

	// Watermark text settings. This parameter is valid if `Type` is `TEXT`.
	TextSettings *CreateTextSettings `json:"TextSettings,omitnil" name:"TextSettings"`
}

type CreateStreamLiveWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// Watermark name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Watermark type. Valid values: STATIC_IMAGE, TEXT.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Watermark image settings. This parameter is valid if `Type` is `STATIC_IMAGE`.
	ImageSettings *CreateImageSettings `json:"ImageSettings,omitnil" name:"ImageSettings"`

	// Watermark text settings. This parameter is valid if `Type` is `TEXT`.
	TextSettings *CreateTextSettings `json:"TextSettings,omitnil" name:"TextSettings"`
}

func (r *CreateStreamLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLiveWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "ImageSettings")
	delete(f, "TextSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLiveWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLiveWatermarkResponseParams struct {
	// Watermark ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLiveWatermarkResponseParams `json:"Response"`
}

func (r *CreateStreamLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLiveWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTextSettings struct {
	// Text
	Text *string `json:"Text,omitnil" name:"Text"`

	// Origin. Valid values: TOP_LEFT, BOTTOM_LEFT, TOP_RIGHT, BOTTOM_RIGHT.
	Location *string `json:"Location,omitnil" name:"Location"`

	// The watermark’s horizontal distance from the origin as a percentage of the video width. Value range: 0-100. Default: 10.
	XPos *int64 `json:"XPos,omitnil" name:"XPos"`

	// The watermark’s vertical distance from the origin as a percentage of the video height. Value range: 0-100. Default: 10.
	YPos *int64 `json:"YPos,omitnil" name:"YPos"`

	// Font size. Value range: 25-50.
	FontSize *int64 `json:"FontSize,omitnil" name:"FontSize"`

	// Font color, which is an RGB color value. Default value: 0x000000.
	FontColor *string `json:"FontColor,omitnil" name:"FontColor"`
}

type DashRemuxSettingsInfo struct {
	// Segment duration in ms. Value range: [1000,30000]. Default value: 4000. The value can only be a multiple of 1,000.
	SegmentDuration *uint64 `json:"SegmentDuration,omitnil" name:"SegmentDuration"`

	// Number of segments. Value range: [1,30]. Default value: 5.
	SegmentNumber *uint64 `json:"SegmentNumber,omitnil" name:"SegmentNumber"`

	// Whether to enable multi-period. Valid values: CLOSE/OPEN. Default value: CLOSE.
	PeriodTriggers *string `json:"PeriodTriggers,omitnil" name:"PeriodTriggers"`

	// The HLS package type when the H.265 codec is used. Valid values: `hvc1`, `hev1` (default).
	H265PackageType *string `json:"H265PackageType,omitnil" name:"H265PackageType"`
}

// Predefined struct for user
type DeleteStreamLiveChannelRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DeleteStreamLiveChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteStreamLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLiveChannelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLiveChannelResponseParams `json:"Response"`
}

func (r *DeleteStreamLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLiveInputRequestParams struct {
	// Input ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DeleteStreamLiveInputRequest struct {
	*tchttp.BaseRequest
	
	// Input ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteStreamLiveInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLiveInputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLiveInputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLiveInputResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamLiveInputResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLiveInputResponseParams `json:"Response"`
}

func (r *DeleteStreamLiveInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLiveInputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLiveInputSecurityGroupRequestParams struct {
	// Input security group ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DeleteStreamLiveInputSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Input security group ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteStreamLiveInputSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLiveInputSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLiveInputSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLiveInputSecurityGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamLiveInputSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLiveInputSecurityGroupResponseParams `json:"Response"`
}

func (r *DeleteStreamLiveInputSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLiveInputSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLivePlanRequestParams struct {
	// ID of the channel whose event is to be deleted
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Name of the event to delete
	EventName *string `json:"EventName,omitnil" name:"EventName"`
}

type DeleteStreamLivePlanRequest struct {
	*tchttp.BaseRequest
	
	// ID of the channel whose event is to be deleted
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Name of the event to delete
	EventName *string `json:"EventName,omitnil" name:"EventName"`
}

func (r *DeleteStreamLivePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLivePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "EventName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLivePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLivePlanResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamLivePlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLivePlanResponseParams `json:"Response"`
}

func (r *DeleteStreamLivePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLivePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLiveWatermarkRequestParams struct {
	// Watermark ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DeleteStreamLiveWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// Watermark ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteStreamLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLiveWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLiveWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLiveWatermarkResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLiveWatermarkResponseParams `json:"Response"`
}

func (r *DeleteStreamLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLiveWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeliveryRestrictionsInfo struct {
	// Corresponds to SCTE-35 web_delivery_allowed_flag parameter.
	WebDeliveryAllowed *string `json:"WebDeliveryAllowed,omitnil" name:"WebDeliveryAllowed"`

	// Corresponds to SCTE-35 no_regional_blackout_flag parameter.
	NoRegionalBlackout *string `json:"NoRegionalBlackout,omitnil" name:"NoRegionalBlackout"`

	// Corresponds to SCTE-35 archive_allowed_flag.
	ArchiveAllowed *string `json:"ArchiveAllowed,omitnil" name:"ArchiveAllowed"`

	// Corresponds to SCTE-35 device_restrictions parameter.
	DeviceRestrictions *string `json:"DeviceRestrictions,omitnil" name:"DeviceRestrictions"`
}

type DescribeImageSettings struct {
	// Origin
	Location *string `json:"Location,omitnil" name:"Location"`

	// The watermark image’s horizontal distance from the origin as a percentage of the video width
	XPos *int64 `json:"XPos,omitnil" name:"XPos"`

	// The watermark image’s vertical distance from the origin as a percentage of the video height
	YPos *int64 `json:"YPos,omitnil" name:"YPos"`

	// The watermark image’s width as a percentage of the video width
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// The watermark image’s height as a percentage of the video height
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

// Predefined struct for user
type DescribeStreamLiveChannelAlertsRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`
}

type DescribeStreamLiveChannelAlertsRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`
}

func (r *DescribeStreamLiveChannelAlertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelAlertsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveChannelAlertsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelAlertsResponseParams struct {
	// Alarm information of the channel’s two pipelines
	Infos *ChannelAlertInfos `json:"Infos,omitnil" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveChannelAlertsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveChannelAlertsResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveChannelAlertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelInputStatisticsRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// UTC time, such as `2020-01-01T12:00:00Z`
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time for query, which is 1 hour after `StartTime` by default
	// UTC time, such as `2020-01-01T12:00:00Z`
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Data collection interval. Valid values: `5s`, `1min` (default), `5min`, `15min`
	Period *string `json:"Period,omitnil" name:"Period"`
}

type DescribeStreamLiveChannelInputStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// UTC time, such as `2020-01-01T12:00:00Z`
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time for query, which is 1 hour after `StartTime` by default
	// UTC time, such as `2020-01-01T12:00:00Z`
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Data collection interval. Valid values: `5s`, `1min` (default), `5min`, `15min`
	Period *string `json:"Period,omitnil" name:"Period"`
}

func (r *DescribeStreamLiveChannelInputStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelInputStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveChannelInputStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelInputStatisticsResponseParams struct {
	// Channel input statistics
	Infos []*ChannelInputStatistics `json:"Infos,omitnil" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveChannelInputStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveChannelInputStatisticsResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveChannelInputStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelInputStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelLogsRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Start time for query, which is 1 hour ago by default. You can query logs in the last 7 days.
	// UTC time, such as `2020-01-01T12:00:00Z`
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time for query, which is 1 hour after `StartTime` by default
	// UTC time, such as `2020-01-01T12:00:00Z`
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeStreamLiveChannelLogsRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Start time for query, which is 1 hour ago by default. You can query logs in the last 7 days.
	// UTC time, such as `2020-01-01T12:00:00Z`
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time for query, which is 1 hour after `StartTime` by default
	// UTC time, such as `2020-01-01T12:00:00Z`
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeStreamLiveChannelLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveChannelLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelLogsResponseParams struct {
	// Pipeline push information
	Infos *PipelineLogInfo `json:"Infos,omitnil" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveChannelLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveChannelLogsResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveChannelLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelOutputStatisticsRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// UTC time, such as `2020-01-01T12:00:00Z`
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time for query, which is 1 hour after `StartTime` by default
	// UTC time, such as `2020-01-01T12:00:00Z`
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Data collection interval. Valid values: `5s`, `1min` (default), `5min`, `15min`
	Period *string `json:"Period,omitnil" name:"Period"`
}

type DescribeStreamLiveChannelOutputStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// UTC time, such as `2020-01-01T12:00:00Z`
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time for query, which is 1 hour after `StartTime` by default
	// UTC time, such as `2020-01-01T12:00:00Z`
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Data collection interval. Valid values: `5s`, `1min` (default), `5min`, `15min`
	Period *string `json:"Period,omitnil" name:"Period"`
}

func (r *DescribeStreamLiveChannelOutputStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelOutputStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveChannelOutputStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelOutputStatisticsResponseParams struct {
	// Channel output information
	Infos []*ChannelOutputsStatistics `json:"Infos,omitnil" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveChannelOutputStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveChannelOutputStatisticsResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveChannelOutputStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelOutputStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeStreamLiveChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeStreamLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelResponseParams struct {
	// Channel information
	Info *StreamLiveChannelInfo `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveChannelResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelsRequestParams struct {

}

type DescribeStreamLiveChannelsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamLiveChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveChannelsResponseParams struct {
	// List of channel information
	// Note: this field may return `null`, indicating that no valid value was found.
	Infos []*StreamLiveChannelInfo `json:"Infos,omitnil" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveChannelsResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveInputRequestParams struct {
	// Input ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeStreamLiveInputRequest struct {
	*tchttp.BaseRequest
	
	// Input ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeStreamLiveInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveInputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveInputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveInputResponseParams struct {
	// Input information
	Info *InputInfo `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveInputResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveInputResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveInputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveInputSecurityGroupRequestParams struct {
	// Input security group ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeStreamLiveInputSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Input security group ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeStreamLiveInputSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveInputSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveInputSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveInputSecurityGroupResponseParams struct {
	// Input security group information
	Info *InputSecurityGroupInfo `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveInputSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveInputSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveInputSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveInputSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveInputSecurityGroupsRequestParams struct {

}

type DescribeStreamLiveInputSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamLiveInputSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveInputSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveInputSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveInputSecurityGroupsResponseParams struct {
	// List of input security group information
	Infos []*InputSecurityGroupInfo `json:"Infos,omitnil" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveInputSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveInputSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveInputSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveInputSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveInputsRequestParams struct {

}

type DescribeStreamLiveInputsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamLiveInputsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveInputsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveInputsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveInputsResponseParams struct {
	// List of input information
	// Note: this field may return `null`, indicating that no valid value was found.
	Infos []*InputInfo `json:"Infos,omitnil" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveInputsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveInputsResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveInputsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveInputsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLivePlansRequestParams struct {
	// ID of the channel whose events you want to query
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`
}

type DescribeStreamLivePlansRequest struct {
	*tchttp.BaseRequest
	
	// ID of the channel whose events you want to query
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`
}

func (r *DescribeStreamLivePlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLivePlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLivePlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLivePlansResponseParams struct {
	// List of event information
	// Note: this field may return `null`, indicating that no valid value was found.
	Infos []*PlanResp `json:"Infos,omitnil" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLivePlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLivePlansResponseParams `json:"Response"`
}

func (r *DescribeStreamLivePlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLivePlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveRegionsRequestParams struct {

}

type DescribeStreamLiveRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamLiveRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveRegionsResponseParams struct {
	// StreamLive region information
	Info *StreamLiveRegionInfo `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveRegionsResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveTranscodeDetailRequestParams struct {
	// The query start time (UTC+8) in the format of yyyy-MM-dd.
	// You can only query data in the last month (not including the current day).
	StartDayTime *string `json:"StartDayTime,omitnil" name:"StartDayTime"`

	// The query end time (UTC+8) in the format of yyyy-MM-dd.
	// You can only query data in the last month (not including the current day).
	EndDayTime *string `json:"EndDayTime,omitnil" name:"EndDayTime"`

	// The channel ID (optional).
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// The number of pages. Default value: 1.
	// The value cannot exceed 100.
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// The number of records per page. Default value: 10.
	// Value range: 1-1000.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeStreamLiveTranscodeDetailRequest struct {
	*tchttp.BaseRequest
	
	// The query start time (UTC+8) in the format of yyyy-MM-dd.
	// You can only query data in the last month (not including the current day).
	StartDayTime *string `json:"StartDayTime,omitnil" name:"StartDayTime"`

	// The query end time (UTC+8) in the format of yyyy-MM-dd.
	// You can only query data in the last month (not including the current day).
	EndDayTime *string `json:"EndDayTime,omitnil" name:"EndDayTime"`

	// The channel ID (optional).
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// The number of pages. Default value: 1.
	// The value cannot exceed 100.
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// The number of records per page. Default value: 10.
	// Value range: 1-1000.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeStreamLiveTranscodeDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveTranscodeDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDayTime")
	delete(f, "EndDayTime")
	delete(f, "ChannelId")
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveTranscodeDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveTranscodeDetailResponseParams struct {
	// A list of the transcoding information.
	Infos []*DescribeTranscodeDetailInfo `json:"Infos,omitnil" name:"Infos"`

	// The number of the current page.
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// The number of records per page.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// The total number of records.
	TotalNum *int64 `json:"TotalNum,omitnil" name:"TotalNum"`

	// The total number of pages.
	TotalPage *int64 `json:"TotalPage,omitnil" name:"TotalPage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveTranscodeDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveTranscodeDetailResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveTranscodeDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveTranscodeDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveWatermarkRequestParams struct {
	// Watermark ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeStreamLiveWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// Watermark ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeStreamLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveWatermarkResponseParams struct {
	// Watermark information
	Info *DescribeWatermarkInfo `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveWatermarkResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveWatermarksRequestParams struct {

}

type DescribeStreamLiveWatermarksRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamLiveWatermarksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveWatermarksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLiveWatermarksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLiveWatermarksResponseParams struct {
	// List of watermark information
	Infos []*DescribeWatermarkInfo `json:"Infos,omitnil" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamLiveWatermarksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLiveWatermarksResponseParams `json:"Response"`
}

func (r *DescribeStreamLiveWatermarksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLiveWatermarksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTextSettings struct {
	// Text
	Text *string `json:"Text,omitnil" name:"Text"`

	// Origin
	Location *string `json:"Location,omitnil" name:"Location"`

	// The watermark image’s horizontal distance from the origin as a percentage of the video width
	XPos *int64 `json:"XPos,omitnil" name:"XPos"`

	// The watermark image’s vertical distance from the origin as a percentage of the video height
	YPos *int64 `json:"YPos,omitnil" name:"YPos"`

	// Font size
	FontSize *int64 `json:"FontSize,omitnil" name:"FontSize"`

	// Font color
	FontColor *string `json:"FontColor,omitnil" name:"FontColor"`
}

type DescribeTranscodeDetailInfo struct {
	// The channel ID.
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// The start time (UTC+8) of transcoding in the format of yyyy-MM-dd HH:mm:ss.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end time (UTC+8) of transcoding in the format of yyyy-MM-dd HH:mm:ss.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The duration (s) of transcoding.
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// The encoding method.
	// Examples:
	// `liveprocessor_H264`: Live transcoding-H264
	// `liveprocessor_H265`: Live transcoding-H265
	// `topspeed_H264`: Top speed codec-H264
	// `topspeed_H265`: Top speed codec-H265
	ModuleCodec *string `json:"ModuleCodec,omitnil" name:"ModuleCodec"`

	// The target bitrate (Kbps).
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// The transcoding type.
	Type *string `json:"Type,omitnil" name:"Type"`

	// The push domain name.
	PushDomain *string `json:"PushDomain,omitnil" name:"PushDomain"`

	// The target resolution.
	Resolution *string `json:"Resolution,omitnil" name:"Resolution"`
}

type DescribeWatermarkInfo struct {
	// Watermark ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Watermark name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Watermark type. Valid values: STATIC_IMAGE, TEXT.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Watermark image settings. This parameter is valid if `Type` is `STATIC_IMAGE`.
	// Note: This field may return `null`, indicating that no valid value was found.
	ImageSettings *DescribeImageSettings `json:"ImageSettings,omitnil" name:"ImageSettings"`

	// Watermark text settings. This parameter is valid if `Type` is `TEXT`.
	// Note: This field may return `null`, indicating that no valid value was found.
	TextSettings *DescribeTextSettings `json:"TextSettings,omitnil" name:"TextSettings"`

	// Last modified time (UTC+0) of the watermark, in the format of `2020-01-01T12:00:00Z`
	// Note: This field may return `null`, indicating that no valid value was found.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// List of channel IDs the watermark is bound to
	// Note: This field may return `null`, indicating that no valid value was found.
	AttachedChannels []*string `json:"AttachedChannels,omitnil" name:"AttachedChannels"`
}

type DestinationInfo struct {
	// Relay destination address. Length limit: [1,512].
	OutputUrl *string `json:"OutputUrl,omitnil" name:"OutputUrl"`

	// Authentication key. Length limit: [1,128].
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthKey *string `json:"AuthKey,omitnil" name:"AuthKey"`

	// Authentication username. Length limit: [1,128].
	// Note: this field may return null, indicating that no valid values can be obtained.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Authentication password. Length limit: [1,128].
	// Note: this field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitnil" name:"Password"`
}

type DrmKey struct {
	// DRM key, which is a 32-bit hexadecimal string.
	// Note: uppercase letters in the string will be automatically converted to lowercase ones.
	Key *string `json:"Key,omitnil" name:"Key"`

	// Required for Widevine encryption. Valid values: SD, HD, UHD1, UHD2, AUDIO, ALL.
	// ALL refers to all tracks. If this parameter is set to ALL, no other tracks can be added.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Track *string `json:"Track,omitnil" name:"Track"`

	// Required for Widevine encryption. It is a 32-bit hexadecimal string.
	// Note: uppercase letters in the string will be automatically converted to lowercase ones.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`

	// Required when FairPlay uses the AES encryption method. It is a 32-bit hexadecimal string.
	// For more information about this parameter, please see: 
	// https://tools.ietf.org/html/rfc3826
	// Note: uppercase letters in the string will be automatically converted to lowercase ones.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Iv *string `json:"Iv,omitnil" name:"Iv"`

	// The URI of the license server when AES-128 is used. This parameter may be empty.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	KeyUri *string `json:"KeyUri,omitnil" name:"KeyUri"`
}

type DrmSettingsInfo struct {
	// Whether to enable DRM encryption. Valid values: `CLOSE` (disable), `OPEN` (enable). Default value: `CLOSE`
	// DRM encryption is supported only for HLS, DASH, HLS_ARCHIVE, DASH_ARCHIVE, HLS_MEDIAPACKAGE, and DASH_MEDIAPACKAGE outputs.
	State *string `json:"State,omitnil" name:"State"`

	// Valid values: `CustomDRMKeys` (default value), `SDMCDRM`
	// `CustomDRMKeys` means encryption keys customized by users.
	// `SDMCDRM` means the DRM key management system of SDMC.
	Scheme *string `json:"Scheme,omitnil" name:"Scheme"`

	// If `Scheme` is set to `CustomDRMKeys`, this parameter is required.
	// If `Scheme` is set to `SDMCDRM`, this parameter is optional. It supports digits, letters, hyphens, and underscores and must contain 1 to 36 characters. If it is not specified, the value of `ChannelId` will be used.
	ContentId *string `json:"ContentId,omitnil" name:"ContentId"`

	// The key customized by the content user, which is required when `Scheme` is set to CustomDRMKeys.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Keys []*DrmKey `json:"Keys,omitnil" name:"Keys"`

	// SDMC key configuration. This parameter is used when `Scheme` is set to `SDMCDRM`.
	// Note: This field may return `null`, indicating that no valid value was found.
	SDMCSettings *SDMCSettingsInfo `json:"SDMCSettings,omitnil" name:"SDMCSettings"`

	// The DRM type. Valid values: `FAIRPLAY`, `WIDEVINE`, `AES128`. For HLS, this can be `FAIRPLAY` or `AES128`. For DASH, this can only be `WIDEVINE`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DrmType *string `json:"DrmType,omitnil" name:"DrmType"`
}

type EventNotifySetting struct {
	// The callback configuration for push events.
	PushEventSettings *PushEventSetting `json:"PushEventSettings,omitnil" name:"PushEventSettings"`
}

type EventSettingsDestinationReq struct {
	// URL of the COS bucket to save recording files
	Url *string `json:"Url,omitnil" name:"Url"`
}

type EventSettingsDestinationResp struct {
	// URL of the COS bucket where recording files are saved
	Url *string `json:"Url,omitnil" name:"Url"`
}

type EventSettingsReq struct {
	// Valid values: `INPUT_SWITCH`, `TIMED_RECORD`, SCTE35_TIME_SIGNAL, SCTE35_SPLICE_INSERT, SCTE35_RETURN_TO_NETWORK. If it is not specified, `INPUT_SWITCH` will be used.
	EventType *string `json:"EventType,omitnil" name:"EventType"`

	// ID of the input to attach, which is required if `EventType` is `INPUT_SWITCH`
	InputAttachment *string `json:"InputAttachment,omitnil" name:"InputAttachment"`

	// Name of the output group to attach. This parameter is required if `EventType` is `TIMED_RECORD`.
	OutputGroupName *string `json:"OutputGroupName,omitnil" name:"OutputGroupName"`

	// Name of the manifest file for timed recording, which must end with `.m3u8` for HLS and `.mpd` for DASH. This parameter is required if `EventType` is `TIMED_RECORD`.
	ManifestName *string `json:"ManifestName,omitnil" name:"ManifestName"`

	// URL of the COS bucket to save recording files. This parameter is required if `EventType` is `TIMED_RECORD`. It may contain 1 or 2 URLs. The first URL corresponds to pipeline 0 and the second pipeline 1.
	Destinations []*EventSettingsDestinationReq `json:"Destinations,omitnil" name:"Destinations"`

	// SCTE-35 configuration information.
	SCTE35SegmentationDescriptor []*SegmentationDescriptorInfo `json:"SCTE35SegmentationDescriptor,omitnil" name:"SCTE35SegmentationDescriptor"`

	// A 32-bit unique segmentation event identifier.Only one occurrence of a given segmentation_event_id value shall be active at any one time.
	SpliceEventID *uint64 `json:"SpliceEventID,omitnil" name:"SpliceEventID"`

	// The duration of the segment in 90kHz ticks.It used to  give the splicer an indication of when the break will be over and when the network In Point will occur. If not specifyed,the splice_insert will continue when enter a return_to_network to end the splice_insert at the appropriate time.
	SpliceDuration *uint64 `json:"SpliceDuration,omitnil" name:"SpliceDuration"`

	// Meta information plan configuration.
	TimedMetadataSetting *TimedMetadataInfo `json:"TimedMetadataSetting,omitnil" name:"TimedMetadataSetting"`
}

type EventSettingsResp struct {
	// Valid values: INPUT_SWITCH, TIMED_RECORD, SCTE35_TIME_SIGNAL, SCTE35_SPLICE_INSERT, SCTE35_RETURN_TO_NETWORK.
	EventType *string `json:"EventType,omitnil" name:"EventType"`

	// ID of the input attached, which is not empty if `EventType` is `INPUT_SWITCH`
	InputAttachment *string `json:"InputAttachment,omitnil" name:"InputAttachment"`

	// Name of the output group attached. This parameter is not empty if `EventType` is `TIMED_RECORD`.
	OutputGroupName *string `json:"OutputGroupName,omitnil" name:"OutputGroupName"`

	// Name of the manifest file for timed recording, which ends with `.m3u8` for HLS and `.mpd` for DASH. This parameter is not empty if `EventType` is `TIMED_RECORD`.
	ManifestName *string `json:"ManifestName,omitnil" name:"ManifestName"`

	// URL of the COS bucket where recording files are saved. This parameter is not empty if `EventType` is `TIMED_RECORD`. It may contain 1 or 2 URLs. The first URL corresponds to pipeline 0 and the second pipeline 1.
	Destinations []*EventSettingsDestinationResp `json:"Destinations,omitnil" name:"Destinations"`

	// SCTE-35 configuration information.
	SCTE35SegmentationDescriptor []*SegmentationDescriptorRespInfo `json:"SCTE35SegmentationDescriptor,omitnil" name:"SCTE35SegmentationDescriptor"`

	// A 32-bit unique segmentation event identifier.Only one occurrence of a given segmentation_event_id value shall be active at any one time.
	SpliceEventID *uint64 `json:"SpliceEventID,omitnil" name:"SpliceEventID"`

	// The duration of the segment in 90kHz ticks.It used to  give the splicer an indication of when the break will be over and when the network In Point will occur. If not specifyed,the splice_insert will continue when enter a return_to_network to end the splice_insert at the appropriate time.
	SpliceDuration *string `json:"SpliceDuration,omitnil" name:"SpliceDuration"`

	// Meta information plan configuration.
	TimedMetadataSetting *TimedMetadataInfo `json:"TimedMetadataSetting,omitnil" name:"TimedMetadataSetting"`
}

type FailOverSettings struct {
	// ID of the backup input
	// Note: this field may return `null`, indicating that no valid value was found.
	SecondaryInputId *string `json:"SecondaryInputId,omitnil" name:"SecondaryInputId"`

	// The wait time (ms) for triggering failover after the primary input becomes unavailable. Value range: [1000, 86400000]. Default value: `3000`
	LossThreshold *int64 `json:"LossThreshold,omitnil" name:"LossThreshold"`

	// Failover policy. Valid values: `CURRENT_PREFERRED` (default), `PRIMARY_PREFERRED`
	RecoverBehavior *string `json:"RecoverBehavior,omitnil" name:"RecoverBehavior"`
}

type HlsRemuxSettingsInfo struct {
	// Segment duration in ms. Value range: [1000,30000]. Default value: 4000. The value can only be a multiple of 1,000.
	SegmentDuration *uint64 `json:"SegmentDuration,omitnil" name:"SegmentDuration"`

	// Number of segments. Value range: [1,30]. Default value: 5.
	SegmentNumber *uint64 `json:"SegmentNumber,omitnil" name:"SegmentNumber"`

	// Whether to enable PDT insertion. Valid values: CLOSE/OPEN. Default value: CLOSE.
	PdtInsertion *string `json:"PdtInsertion,omitnil" name:"PdtInsertion"`

	// PDT duration in seconds. Value range: (0,3000]. Default value: 600.
	PdtDuration *uint64 `json:"PdtDuration,omitnil" name:"PdtDuration"`

	// Audio/Video packaging scheme. Valid values: `SEPARATE`, `MERGE`. Default value is: SEPARATE.
	Scheme *string `json:"Scheme,omitnil" name:"Scheme"`

	// The segment type. Valid values: `ts` (default), `fmp4`.
	// Currently, fMP4 segments do not support DRM or time shifting.
	SegmentType *string `json:"SegmentType,omitnil" name:"SegmentType"`

	// The HLS package type when the H.265 codec is used. Valid values: `hvc1`, `hev1` (default).
	H265PackageType *string `json:"H265PackageType,omitnil" name:"H265PackageType"`


	LowLatency *uint64 `json:"LowLatency,omitnil" name:"LowLatency"`


	PartialSegmentDuration *uint64 `json:"PartialSegmentDuration,omitnil" name:"PartialSegmentDuration"`


	PartialSegmentPlaySite *uint64 `json:"PartialSegmentPlaySite,omitnil" name:"PartialSegmentPlaySite"`
}

type InputInfo struct {
	// Input region.
	Region *string `json:"Region,omitnil" name:"Region"`

	// Input ID.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Input name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Input type.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Array of security groups associated with input.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Array of channels associated with input.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AttachedChannels []*string `json:"AttachedChannels,omitnil" name:"AttachedChannels"`

	// Input configuration array.
	InputSettings []*InputSettingInfo `json:"InputSettings,omitnil" name:"InputSettings"`
}

type InputLossBehaviorInfo struct {
	// The time to fill in the last video frame, unit ms, range 0-1000000, 1000000 means always inserting, default 0 means filling in black screen frame.
	RepeatLastFrameMs *uint64 `json:"RepeatLastFrameMs,omitnil" name:"RepeatLastFrameMs"`

	// Fill frame type, COLOR means solid color filling, IMAGE means picture filling, the default is COLOR.
	InputLossImageType *string `json:"InputLossImageType,omitnil" name:"InputLossImageType"`

	// When the type is COLOR, the corresponding rgb value
	ColorRGB *string `json:"ColorRGB,omitnil" name:"ColorRGB"`

	// When the type is IMAGE, the corresponding image url value
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type InputSecurityGroupInfo struct {
	// Input security group ID.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Input security group name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// List of allowlist entries.
	Whitelist []*string `json:"Whitelist,omitnil" name:"Whitelist"`

	// List of bound input streams.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OccupiedInputs []*string `json:"OccupiedInputs,omitnil" name:"OccupiedInputs"`

	// Input security group address.
	Region *string `json:"Region,omitnil" name:"Region"`
}

type InputSettingInfo struct {
	// Application name, which is valid if `Type` is `RTMP_PUSH` and can contain 1-32 letters and digits
	// Note: This field may return `null`, indicating that no valid value was found.
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// Stream name, which is valid if `Type` is `RTMP_PUSH` and can contain 1-32 letters and digits
	// Note: This field may return `null`, indicating that no valid value was found.
	StreamName *string `json:"StreamName,omitnil" name:"StreamName"`

	// Source URL, which is valid if `Type` is `RTMP_PULL`, `HLS_PULL`, or `MP4_PULL` and can contain 1-512 characters
	// Note: This field may return `null`, indicating that no valid value was found.
	SourceUrl *string `json:"SourceUrl,omitnil" name:"SourceUrl"`

	// RTP/UDP input address, which does not need to be entered for the input parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InputAddress *string `json:"InputAddress,omitnil" name:"InputAddress"`

	// Source type for stream pulling and relaying. To pull content from private-read COS buckets under the current account, set this parameter to `TencentCOS`; otherwise, leave it empty.
	// Note: this field may return `null`, indicating that no valid value was found.
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// Delayed time (ms) for playback, which is valid if `Type` is `RTMP_PUSH`
	// Value range: 0 (default) or 10000-600000
	// The value must be a multiple of 1,000.
	// Note: This field may return `null`, indicating that no valid value was found.
	DelayTime *int64 `json:"DelayTime,omitnil" name:"DelayTime"`

	// The domain of an SRT_PUSH address. If this is a request parameter, you don’t need to specify it.
	// Note: This field may return `null`, indicating that no valid value was found.
	InputDomain *string `json:"InputDomain,omitnil" name:"InputDomain"`

	// The username, which is used for authentication.
	// Note: This field may return `null`, indicating that no valid value was found.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// The password, which is used for authentication.
	// Note: This field may return `null`, indicating that no valid value was found.
	Password *string `json:"Password,omitnil" name:"Password"`
}

type InputStatistics struct {
	// Input statistics of pipeline 0.
	Pipeline0 []*PipelineInputStatistics `json:"Pipeline0,omitnil" name:"Pipeline0"`

	// Input statistics of pipeline 1.
	Pipeline1 []*PipelineInputStatistics `json:"Pipeline1,omitnil" name:"Pipeline1"`
}

type InputStreamInfo struct {
	// The input stream address.
	InputAddress *string `json:"InputAddress,omitnil" name:"InputAddress"`

	// The input stream path.
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// The input stream name.
	StreamName *string `json:"StreamName,omitnil" name:"StreamName"`

	// The input stream status. `1` indicates the stream is active.
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type LogInfo struct {
	// Log type.
	// It contains the value of `StreamStart` which refers to the push information.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Time when the log is printed.
	Time *string `json:"Time,omitnil" name:"Time"`

	// Log details.
	Message *LogMessageInfo `json:"Message,omitnil" name:"Message"`
}

type LogMessageInfo struct {
	// Push information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StreamInfo *StreamInfo `json:"StreamInfo,omitnil" name:"StreamInfo"`
}

// Predefined struct for user
type ModifyStreamLiveChannelRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Channel name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Inputs to attach. You can attach 1 to 5 inputs.
	AttachedInputs []*AttachedInput `json:"AttachedInputs,omitnil" name:"AttachedInputs"`

	// Configuration information of the channel’s output groups. Quantity: [1, 10]
	OutputGroups []*StreamLiveOutputGroupsInfo `json:"OutputGroups,omitnil" name:"OutputGroups"`

	// Audio transcoding templates. Quantity: [1, 20]
	AudioTemplates []*AudioTemplateInfo `json:"AudioTemplates,omitnil" name:"AudioTemplates"`

	// Video transcoding templates. Quantity: [1, 10]
	VideoTemplates []*VideoTemplateInfo `json:"VideoTemplates,omitnil" name:"VideoTemplates"`

	// Audio/Video transcoding templates. Quantity: [1, 10]
	AVTemplates []*AVTemplate `json:"AVTemplates,omitnil" name:"AVTemplates"`

	// Event settings
	PlanSettings *PlanSettings `json:"PlanSettings,omitnil" name:"PlanSettings"`

	// The callback settings.
	EventNotifySettings *EventNotifySetting `json:"EventNotifySettings,omitnil" name:"EventNotifySettings"`

	// Complement the last video frame settings.
	InputLossBehavior *InputLossBehaviorInfo `json:"InputLossBehavior,omitnil" name:"InputLossBehavior"`
}

type ModifyStreamLiveChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Channel name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Inputs to attach. You can attach 1 to 5 inputs.
	AttachedInputs []*AttachedInput `json:"AttachedInputs,omitnil" name:"AttachedInputs"`

	// Configuration information of the channel’s output groups. Quantity: [1, 10]
	OutputGroups []*StreamLiveOutputGroupsInfo `json:"OutputGroups,omitnil" name:"OutputGroups"`

	// Audio transcoding templates. Quantity: [1, 20]
	AudioTemplates []*AudioTemplateInfo `json:"AudioTemplates,omitnil" name:"AudioTemplates"`

	// Video transcoding templates. Quantity: [1, 10]
	VideoTemplates []*VideoTemplateInfo `json:"VideoTemplates,omitnil" name:"VideoTemplates"`

	// Audio/Video transcoding templates. Quantity: [1, 10]
	AVTemplates []*AVTemplate `json:"AVTemplates,omitnil" name:"AVTemplates"`

	// Event settings
	PlanSettings *PlanSettings `json:"PlanSettings,omitnil" name:"PlanSettings"`

	// The callback settings.
	EventNotifySettings *EventNotifySetting `json:"EventNotifySettings,omitnil" name:"EventNotifySettings"`

	// Complement the last video frame settings.
	InputLossBehavior *InputLossBehaviorInfo `json:"InputLossBehavior,omitnil" name:"InputLossBehavior"`
}

func (r *ModifyStreamLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "AttachedInputs")
	delete(f, "OutputGroups")
	delete(f, "AudioTemplates")
	delete(f, "VideoTemplates")
	delete(f, "AVTemplates")
	delete(f, "PlanSettings")
	delete(f, "EventNotifySettings")
	delete(f, "InputLossBehavior")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLiveChannelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLiveChannelResponseParams `json:"Response"`
}

func (r *ModifyStreamLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLiveInputRequestParams struct {
	// Input ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Input name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// List of the IDs of the security groups to attach
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Input settings
	// For the type `RTMP_PUSH`, `RTMP_PULL`, `HLS_PULL`, or `MP4_PULL`, 1 or 2 inputs of the corresponding type can be configured.
	// This parameter can be left empty for RTP_PUSH and UDP_PUSH inputs.
	// Note: If this parameter is not specified or empty, the original input settings will be used.
	InputSettings []*InputSettingInfo `json:"InputSettings,omitnil" name:"InputSettings"`
}

type ModifyStreamLiveInputRequest struct {
	*tchttp.BaseRequest
	
	// Input ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Input name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// List of the IDs of the security groups to attach
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Input settings
	// For the type `RTMP_PUSH`, `RTMP_PULL`, `HLS_PULL`, or `MP4_PULL`, 1 or 2 inputs of the corresponding type can be configured.
	// This parameter can be left empty for RTP_PUSH and UDP_PUSH inputs.
	// Note: If this parameter is not specified or empty, the original input settings will be used.
	InputSettings []*InputSettingInfo `json:"InputSettings,omitnil" name:"InputSettings"`
}

func (r *ModifyStreamLiveInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLiveInputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "SecurityGroupIds")
	delete(f, "InputSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLiveInputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLiveInputResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamLiveInputResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLiveInputResponseParams `json:"Response"`
}

func (r *ModifyStreamLiveInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLiveInputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLiveInputSecurityGroupRequestParams struct {
	// Input security group ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Input security group name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Allowlist entries (max: 10)
	Whitelist []*string `json:"Whitelist,omitnil" name:"Whitelist"`
}

type ModifyStreamLiveInputSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Input security group ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Input security group name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the region level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Allowlist entries (max: 10)
	Whitelist []*string `json:"Whitelist,omitnil" name:"Whitelist"`
}

func (r *ModifyStreamLiveInputSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLiveInputSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "Whitelist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLiveInputSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLiveInputSecurityGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamLiveInputSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLiveInputSecurityGroupResponseParams `json:"Response"`
}

func (r *ModifyStreamLiveInputSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLiveInputSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLiveWatermarkRequestParams struct {
	// Watermark ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Watermark name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Watermark image settings. This parameter is valid if `Type` is `STATIC_IMAGE`.
	ImageSettings *CreateImageSettings `json:"ImageSettings,omitnil" name:"ImageSettings"`

	// Watermark text settings. This parameter is valid if `Type` is `TEXT`.
	TextSettings *CreateTextSettings `json:"TextSettings,omitnil" name:"TextSettings"`
}

type ModifyStreamLiveWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// Watermark ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Watermark name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Watermark image settings. This parameter is valid if `Type` is `STATIC_IMAGE`.
	ImageSettings *CreateImageSettings `json:"ImageSettings,omitnil" name:"ImageSettings"`

	// Watermark text settings. This parameter is valid if `Type` is `TEXT`.
	TextSettings *CreateTextSettings `json:"TextSettings,omitnil" name:"TextSettings"`
}

func (r *ModifyStreamLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLiveWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "ImageSettings")
	delete(f, "TextSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLiveWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLiveWatermarkResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLiveWatermarkResponseParams `json:"Response"`
}

func (r *ModifyStreamLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLiveWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutputInfo struct {
	// Output name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Audio transcoding template name array.
	// Quantity limit: [0,1] for RTMP; [0,20] for others.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioTemplateNames []*string `json:"AudioTemplateNames,omitnil" name:"AudioTemplateNames"`

	// Video transcoding template name array. Quantity limit: [0,1].
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoTemplateNames []*string `json:"VideoTemplateNames,omitnil" name:"VideoTemplateNames"`

	// SCTE-35 information configuration.
	Scte35Settings *Scte35SettingsInfo `json:"Scte35Settings,omitnil" name:"Scte35Settings"`

	// Audio/Video transcoding template name. If `HlsRemuxSettings.Scheme` is `MERGE`, there is 1 audio/video transcoding template. Otherwise, this parameter is empty.
	// Note: this field may return `null`, indicating that no valid value was found.
	AVTemplateNames []*string `json:"AVTemplateNames,omitnil" name:"AVTemplateNames"`

	// Meta information controls configuration.
	TimedMetadataSettings *TimedMetadataSettingInfo `json:"TimedMetadataSettings,omitnil" name:"TimedMetadataSettings"`
}

type OutputsStatistics struct {
	// Output information of pipeline 0.
	Pipeline0 []*PipelineOutputStatistics `json:"Pipeline0,omitnil" name:"Pipeline0"`

	// Output information of pipeline 1.
	Pipeline1 []*PipelineOutputStatistics `json:"Pipeline1,omitnil" name:"Pipeline1"`
}

type PipelineInputStatistics struct {
	// Data timestamp in seconds.
	Timestamp *uint64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// Input bandwidth in bps.
	NetworkIn *uint64 `json:"NetworkIn,omitnil" name:"NetworkIn"`

	// Video information array.
	// For `rtp/udp` input, the quantity is the number of `Pid` of the input video.
	// For other inputs, the quantity is 1.
	Video []*VideoPipelineInputStatistics `json:"Video,omitnil" name:"Video"`

	// Audio information array.
	// For `rtp/udp` input, the quantity is the number of `Pid` of the input audio.
	// For other inputs, the quantity is 1.
	Audio []*AudioPipelineInputStatistics `json:"Audio,omitnil" name:"Audio"`
}

type PipelineLogInfo struct {
	// Log information of pipeline 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pipeline0 []*LogInfo `json:"Pipeline0,omitnil" name:"Pipeline0"`

	// Log information of pipeline 1.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pipeline1 []*LogInfo `json:"Pipeline1,omitnil" name:"Pipeline1"`
}

type PipelineOutputStatistics struct {
	// Timestamp.
	// In seconds, indicating data time.
	Timestamp *uint64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// Output bandwidth in bps.
	NetworkOut *uint64 `json:"NetworkOut,omitnil" name:"NetworkOut"`
}

type PlanReq struct {
	// Event name
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// Event trigger time settings
	TimingSettings *TimingSettingsReq `json:"TimingSettings,omitnil" name:"TimingSettings"`

	// Event configuration
	EventSettings *EventSettingsReq `json:"EventSettings,omitnil" name:"EventSettings"`
}

type PlanResp struct {
	// Event name
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// Event trigger time settings
	TimingSettings *TimingSettingsResp `json:"TimingSettings,omitnil" name:"TimingSettings"`

	// Event configuration
	EventSettings *EventSettingsResp `json:"EventSettings,omitnil" name:"EventSettings"`
}

type PlanSettings struct {
	// Timed recording settings
	// Note: This field may return `null`, indicating that no valid value was found.
	TimedRecordSettings *TimedRecordSettings `json:"TimedRecordSettings,omitnil" name:"TimedRecordSettings"`
}

type PushEventSetting struct {
	// The callback URL (required).
	NotifyUrl *string `json:"NotifyUrl,omitnil" name:"NotifyUrl"`

	// The callback key (optional).
	NotifyKey *string `json:"NotifyKey,omitnil" name:"NotifyKey"`
}

type QueryDispatchInputInfo struct {
	// The input ID.
	InputID *string `json:"InputID,omitnil" name:"InputID"`

	// The input name.
	InputName *string `json:"InputName,omitnil" name:"InputName"`

	// The input protocol.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// The stream status of the input.
	InputStreamInfoList []*InputStreamInfo `json:"InputStreamInfoList,omitnil" name:"InputStreamInfoList"`
}

// Predefined struct for user
type QueryInputStreamStateRequestParams struct {
	// The StreamLive input ID.
	Id *string `json:"Id,omitnil" name:"Id"`
}

type QueryInputStreamStateRequest struct {
	*tchttp.BaseRequest
	
	// The StreamLive input ID.
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *QueryInputStreamStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInputStreamStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryInputStreamStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryInputStreamStateResponseParams struct {
	// The information of the StreamLive input queried.
	Info *QueryDispatchInputInfo `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryInputStreamStateResponse struct {
	*tchttp.BaseResponse
	Response *QueryInputStreamStateResponseParams `json:"Response"`
}

func (r *QueryInputStreamStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInputStreamStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// Region name
	Name *string `json:"Name,omitnil" name:"Name"`
}

type SDMCSettingsInfo struct {
	// User ID in the SDMC DRM system
	Uid *string `json:"Uid,omitnil" name:"Uid"`

	// Tracks of the SDMC DRM system. This parameter is valid for DASH output groups.
	// `1`: audio
	// `2`: SD
	// `4`: HD
	// `8`: UHD1
	// `16`: UHD2
	// 
	// Default value: `31` (audio + SD + HD + UHD1 + UHD2)
	Tracks *int64 `json:"Tracks,omitnil" name:"Tracks"`

	// Key ID in the SDMC DRM system; required
	SecretId *string `json:"SecretId,omitnil" name:"SecretId"`

	// Key in the SDMC DRM system; required
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// Key request URL of the SDMC DRM system, which is `https://uat.multidrm.tv/cpix/2.2/getcontentkey` by default
	Url *string `json:"Url,omitnil" name:"Url"`

	// Token name in an SDMC key request URL, which is `token` by default
	TokenName *string `json:"TokenName,omitnil" name:"TokenName"`
}

type Scte35SettingsInfo struct {
	// Whether to pass through SCTE-35 information. Valid values: NO_PASSTHROUGH/PASSTHROUGH. Default value: NO_PASSTHROUGH.
	Behavior *string `json:"Behavior,omitnil" name:"Behavior"`
}

type SegmentationDescriptorInfo struct {
	// A 32-bit unique segmentation event identifier. Only one occurrence of a given segmentation_event_id value shall be active at any one time.
	EventID *uint64 `json:"EventID,omitnil" name:"EventID"`

	// Indicates that a previously sent segmentation event, identified by segmentation_event_id, has been cancelled.
	EventCancelIndicator *uint64 `json:"EventCancelIndicator,omitnil" name:"EventCancelIndicator"`

	// Distribution configuration.
	DeliveryRestrictions *DeliveryRestrictionsInfo `json:"DeliveryRestrictions,omitnil" name:"DeliveryRestrictions"`

	// The duration of the segment in 90kHz ticks. indicat when the segment will be over and when the next segmentation message will occur.Shall be 0 for end messages.the time signal will continue until insert a cancellation message when not specify the duration.
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`

	// Corresponds to SCTE-35 segmentation_upid_type parameter.
	UPIDType *uint64 `json:"UPIDType,omitnil" name:"UPIDType"`

	// Corresponds to SCTE-35 segmentation_upid. 
	UPID *string `json:"UPID,omitnil" name:"UPID"`

	// Corresponds to SCTE-35 segmentation_type_id.
	TypeID *uint64 `json:"TypeID,omitnil" name:"TypeID"`

	// Corresponds to SCTE-35 segment_num。This field provides support for numbering segments within a given collection of segments.
	Num *uint64 `json:"Num,omitnil" name:"Num"`

	// Corresponds to SCTE-35 segment_expected.This field provides a count of the expected number of individual segments within a collection of segments.
	Expected *uint64 `json:"Expected,omitnil" name:"Expected"`

	// Corresponds to SCTE-35 sub_segment_num.This field provides identification for a specific sub-segment within a collection of sub-segments.
	SubSegmentNum *uint64 `json:"SubSegmentNum,omitnil" name:"SubSegmentNum"`

	// Corresponds to SCTE-35 sub_segments_expected.This field provides a count of the expected number of individual sub-segments within the collection of sub-segments.
	SubSegmentsExpected *uint64 `json:"SubSegmentsExpected,omitnil" name:"SubSegmentsExpected"`
}

type SegmentationDescriptorRespInfo struct {
	// A 32-bit unique segmentation event identifier. Only one occurrence of a given segmentation_event_id value shall be active at any one time.
	EventID *uint64 `json:"EventID,omitnil" name:"EventID"`

	// Indicates that a previously sent segmentation event, identified by segmentation_event_id, has been cancelled.
	EventCancelIndicator *uint64 `json:"EventCancelIndicator,omitnil" name:"EventCancelIndicator"`

	// Distribution configuration.
	DeliveryRestrictions *DeliveryRestrictionsInfo `json:"DeliveryRestrictions,omitnil" name:"DeliveryRestrictions"`

	// The duration of the segment in 90kHz ticks. indicat when the segment will be over and when the next segmentation message will occur.Shall be 0 for end messages.the time signal will continue until insert a cancellation message when not specify the duration.
	Duration *string `json:"Duration,omitnil" name:"Duration"`

	// Corresponds to SCTE-35 segmentation_upid_type parameter.
	UPIDType *uint64 `json:"UPIDType,omitnil" name:"UPIDType"`

	// Corresponds to SCTE-35 segmentation_upid. 
	UPID *string `json:"UPID,omitnil" name:"UPID"`

	// Corresponds to SCTE-35 segmentation_type_id.
	TypeID *uint64 `json:"TypeID,omitnil" name:"TypeID"`

	// Corresponds to SCTE-35 segment_num。This field provides support for numbering segments within a given collection of segments.
	Num *uint64 `json:"Num,omitnil" name:"Num"`

	// Corresponds to SCTE-35 segment_expected.This field provides a count of the expected number of individual segments within a collection of segments.
	Expected *uint64 `json:"Expected,omitnil" name:"Expected"`

	// Corresponds to SCTE-35 sub_segment_num.This field provides identification for a specific sub-segment within a collection of sub-segments.
	SubSegmentNum *uint64 `json:"SubSegmentNum,omitnil" name:"SubSegmentNum"`

	// Corresponds to SCTE-35 sub_segments_expected.This field provides a count of the expected number of individual sub-segments within the collection of sub-segments.
	SubSegmentsExpected *uint64 `json:"SubSegmentsExpected,omitnil" name:"SubSegmentsExpected"`
}

// Predefined struct for user
type StartStreamLiveChannelRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type StartStreamLiveChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *StartStreamLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartStreamLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStreamLiveChannelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartStreamLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *StartStreamLiveChannelResponseParams `json:"Response"`
}

func (r *StartStreamLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamLiveChannelRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type StopStreamLiveChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *StopStreamLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopStreamLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamLiveChannelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopStreamLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *StopStreamLiveChannelResponseParams `json:"Response"`
}

func (r *StopStreamLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StreamAudioInfo struct {
	// Audio `Pid`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// Audio codec.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// Audio frame rate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`

	// Audio bitrate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Rate *int64 `json:"Rate,omitnil" name:"Rate"`

	// Audio sample rate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SampleRate *int64 `json:"SampleRate,omitnil" name:"SampleRate"`
}

type StreamInfo struct {
	// Client IP.
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`

	// Video information of pushed streams.
	Video []*StreamVideoInfo `json:"Video,omitnil" name:"Video"`

	// Audio information of pushed streams.
	Audio []*StreamAudioInfo `json:"Audio,omitnil" name:"Audio"`

	// SCTE-35 information of pushed streams.
	Scte35 []*StreamScte35Info `json:"Scte35,omitnil" name:"Scte35"`
}

type StreamLiveChannelInfo struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Channel status
	State *string `json:"State,omitnil" name:"State"`

	// Information of attached inputs
	AttachedInputs []*AttachedInput `json:"AttachedInputs,omitnil" name:"AttachedInputs"`

	// Information of output groups
	OutputGroups []*StreamLiveOutputGroupsInfo `json:"OutputGroups,omitnil" name:"OutputGroups"`

	// Channel name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Audio transcoding templates
	// Note: this field may return `null`, indicating that no valid value was found.
	AudioTemplates []*AudioTemplateInfo `json:"AudioTemplates,omitnil" name:"AudioTemplates"`

	// Video transcoding templates
	// Note: this field may return `null`, indicating that no valid value was found.
	VideoTemplates []*VideoTemplateInfo `json:"VideoTemplates,omitnil" name:"VideoTemplates"`

	// Audio/Video transcoding templates
	// Note: this field may return `null`, indicating that no valid value was found.
	AVTemplates []*AVTemplate `json:"AVTemplates,omitnil" name:"AVTemplates"`

	// Event settings
	// Note: This field may return `null`, indicating that no valid value was found.
	PlanSettings *PlanSettings `json:"PlanSettings,omitnil" name:"PlanSettings"`

	// The callback settings.
	// Note: This field may return `null`, indicating that no valid value was found.
	EventNotifySettings *EventNotifySetting `json:"EventNotifySettings,omitnil" name:"EventNotifySettings"`

	// Supplement the last video frame configuration settings.
	InputLossBehavior *InputLossBehaviorInfo `json:"InputLossBehavior,omitnil" name:"InputLossBehavior"`
}

type StreamLiveOutputGroupsInfo struct {
	// Output group name, which can contain 1-32 case-sensitive letters, digits, and underscores and must be unique at the channel level
	Name *string `json:"Name,omitnil" name:"Name"`

	// Output protocol
	// Valid values: `HLS`, `DASH`, `HLS_ARCHIVE`, `HLS_STREAM_PACKAGE`, `DASH_STREAM_PACKAGE`
	Type *string `json:"Type,omitnil" name:"Type"`

	// Output information
	// If the type is RTMP or RTP, only one output is allowed; if it is HLS or DASH, 1-10 outputs are allowed.
	Outputs []*OutputInfo `json:"Outputs,omitnil" name:"Outputs"`

	// Relay destinations. Quantity: [1, 2]
	Destinations []*DestinationInfo `json:"Destinations,omitnil" name:"Destinations"`

	// HLS protocol configuration information, which takes effect only for HLS/HLS_ARCHIVE outputs
	// Note: this field may return `null`, indicating that no valid value was found.
	HlsRemuxSettings *HlsRemuxSettingsInfo `json:"HlsRemuxSettings,omitnil" name:"HlsRemuxSettings"`

	// DRM configuration information
	// Note: this field may return `null`, indicating that no valid value was found.
	DrmSettings *DrmSettingsInfo `json:"DrmSettings,omitnil" name:"DrmSettings"`

	// DASH protocol configuration information, which takes effect only for DASH/DASH_ARCHIVE outputs
	// Note: this field may return `null`, indicating that no valid value was found.
	DashRemuxSettings *DashRemuxSettingsInfo `json:"DashRemuxSettings,omitnil" name:"DashRemuxSettings"`

	// StreamPackage configuration information, which is required if the output type is StreamPackage
	// Note: this field may return `null`, indicating that no valid value was found.
	StreamPackageSettings *StreamPackageSettingsInfo `json:"StreamPackageSettings,omitnil" name:"StreamPackageSettings"`

	// Time-shift configuration information
	// Note: This field may return `null`, indicating that no valid value was found.
	TimeShiftSettings *TimeShiftSettingsInfo `json:"TimeShiftSettings,omitnil" name:"TimeShiftSettings"`
}

type StreamLiveRegionInfo struct {
	// List of StreamLive regions
	Regions []*RegionInfo `json:"Regions,omitnil" name:"Regions"`
}

type StreamPackageSettingsInfo struct {
	// Channel ID in StreamPackage
	Id *string `json:"Id,omitnil" name:"Id"`
}

type StreamScte35Info struct {
	// SCTE-35 `Pid`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`
}

type StreamVideoInfo struct {
	// Video `Pid`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// Video codec.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Codec *string `json:"Codec,omitnil" name:"Codec"`

	// Video frame rate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Fps *int64 `json:"Fps,omitnil" name:"Fps"`

	// Video bitrate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Rate *int64 `json:"Rate,omitnil" name:"Rate"`

	// Video width.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// Video height.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type TimeShiftSettingsInfo struct {
	// Whether to enable time shifting. Valid values: `OPEN`; `CLOSE`
	// Note: This field may return `null`, indicating that no valid value was found.
	State *string `json:"State,omitnil" name:"State"`

	// Domain name bound for time shifting
	// Note: This field may return `null`, indicating that no valid value was found.
	PlayDomain *string `json:"PlayDomain,omitnil" name:"PlayDomain"`

	// Allowable time-shift period (s). Value range: [600, 1209600]. Default value: 300
	// Note: This field may return `null`, indicating that no valid value was found.
	StartoverWindow *int64 `json:"StartoverWindow,omitnil" name:"StartoverWindow"`
}

type TimedMetadataInfo struct {
	// Base64-encoded id3 metadata information, with a maximum limit of 1024 characters.
	ID3 *string `json:"ID3,omitnil" name:"ID3"`
}

type TimedMetadataSettingInfo struct {
	// Whether to transparently transmit ID3 information, optional values: 0:NO_PASSTHROUGH, 1:PASSTHROUGH, default 0.
	Behavior *uint64 `json:"Behavior,omitnil" name:"Behavior"`
}

type TimedRecordSettings struct {
	// Whether to automatically delete finished recording events. Valid values: `CLOSE`, `OPEN`. If this parameter is left empty, `CLOSE` will be used.
	// If it is set to `OPEN`, a recording event will be deleted 7 days after it is finished.
	// Note: This field may return `null`, indicating that no valid value was found.
	AutoClear *string `json:"AutoClear,omitnil" name:"AutoClear"`
}

type TimingSettingsReq struct {
	// Event trigger type. Valid values: `FIXED_TIME`, `IMMEDIATE`. This parameter is required if `EventType` is `INPUT_SWITCH`.
	StartType *string `json:"StartType,omitnil" name:"StartType"`

	// This parameter is required if `EventType` is `INPUT_SWITCH` and `StartType` is `FIXED_TIME`.
	// It must be in UTC format, e.g., `2020-01-01T12:00:00Z`.
	Time *string `json:"Time,omitnil" name:"Time"`

	// This parameter is required if `EventType` is `TIMED_RECORD`.
	// It specifies the recording start time in UTC format (e.g., `2020-01-01T12:00:00Z`) and must be at least 1 minute later than the current time.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// This parameter is required if `EventType` is `TIMED_RECORD`.
	// It specifies the recording end time in UTC format (e.g., `2020-01-01T12:00:00Z`) and must be at least 1 minute later than the recording start time.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type TimingSettingsResp struct {
	// Event trigger type
	StartType *string `json:"StartType,omitnil" name:"StartType"`

	// Not empty if `StartType` is `FIXED_TIME`
	// UTC time, such as `2020-01-01T12:00:00Z`
	Time *string `json:"Time,omitnil" name:"Time"`

	// This parameter cannot be empty if `EventType` is `TIMED_RECORD`.
	// It indicates the start time for recording in UTC format (e.g., `2020-01-01T12:00:00Z`) and must be at least 1 minute later than the current time.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// This parameter cannot be empty if `EventType` is `TIMED_RECORD`.
	// It indicates the end time for recording in UTC format (e.g., `2020-01-01T12:00:00Z`) and must be at least 1 minute later than the start time for recording.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type VideoPipelineInputStatistics struct {
	// Video FPS.
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// Video bitrate in bps.
	Rate *uint64 `json:"Rate,omitnil" name:"Rate"`

	// Video `Pid`, which is available only if the input is `rtp/udp`.
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`
}

type VideoTemplateInfo struct {
	// Video transcoding template name, which can contain 1-20 letters and digits.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Video codec. Valid values: H264/H265. If this parameter is left empty, the original value will be used.
	Vcodec *string `json:"Vcodec,omitnil" name:"Vcodec"`

	// Video bitrate. Value range: [50000,40000000]. The value can only be a multiple of 1,000. If this parameter is left empty, the original value will be used.
	VideoBitrate *uint64 `json:"VideoBitrate,omitnil" name:"VideoBitrate"`

	// Video width. Value range: (0,3000]. The value can only be a multiple of 4. If this parameter is left empty, the original value will be used.
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// Video height. Value range: (0,3000]. The value can only be a multiple of 4. If this parameter is left empty, the original value will be used.
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// Video frame rate. Value range: [1,240]. If this parameter is left empty, the original value will be used.
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// Whether to enable top speed codec. Valid value: CLOSE/OPEN. Default value: CLOSE.
	TopSpeed *string `json:"TopSpeed,omitnil" name:"TopSpeed"`

	// Top speed codec compression ratio. Value range: [0,50]. The lower the compression ratio, the higher the image quality.
	BitrateCompressionRatio *uint64 `json:"BitrateCompressionRatio,omitnil" name:"BitrateCompressionRatio"`

	// Bitrate control mode. Valid values: `CBR`, `ABR` (default)
	RateControlMode *string `json:"RateControlMode,omitnil" name:"RateControlMode"`

	// Watermark ID
	// Note: This field may return `null`, indicating that no valid value was found.
	WatermarkId *string `json:"WatermarkId,omitnil" name:"WatermarkId"`
}