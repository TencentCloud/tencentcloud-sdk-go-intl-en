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

package v20180801

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AddDelayLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Delay time in seconds, up to 600s.
	DelayTime *uint64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// Expiration time of the configured delayed playback in UTC format, such as 2018-11-29T19:00:00Z.
	// Notes:
	// 1. The configuration will expire after 7 days by default and can last up to 7 days.
	// 2. The Beijing time is in UTC+8. This value should be in the format as required by ISO 8601. For more information, please see [ISO Date and Time Format](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *AddDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDelayLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDelayLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDelayLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Domain name type.
	// 0: push domain name.
	// 1: playback domain name.
	DomainType *uint64 `json:"DomainType,omitempty" name:"DomainType"`

	// Pull domain name type:
	// 1: Mainland China.
	// 2: global.
	// 3: outside Mainland China.
	// Default value: 1.
	PlayType *uint64 `json:"PlayType,omitempty" name:"PlayType"`

	// Whether it is LCB:
	// 0: LVB,
	// 1: LCB.
	// Default value: 0.
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`
}

func (r *AddLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// Watermark image URL.
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// Watermark name.
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// Display position: X-axis offset.
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// Display position: Y-axis offset.
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// Watermark width or its percentage of the live streaming video width. It is recommended to just specify either height or width as the other will be scaled proportionally to avoid distortions.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Watermark height or its percentage of the live streaming video width. It is recommended to just specify either height or width as the other will be scaled proportionally to avoid distortions.
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

func (r *AddLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Watermark ID.
		WatermarkId *uint64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// Certificate ID, which can be obtained through the `CreateLiveCert` API.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Status. 0: off, 1: on.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *BindLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CallBackRuleInfo struct {

	// Rule creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Rule update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type CallBackTemplateInfo struct {

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Stream starting callback URL.
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// Stream ending callback URL.
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// Stream mixing callback URL.
	StreamMixNotifyUrl *string `json:"StreamMixNotifyUrl,omitempty" name:"StreamMixNotifyUrl"`

	// Recording callback URL.
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// Screencapturing callback URL.
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// Porn detection callback URL.
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`

	// Callback authentication key.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

type CertInfo struct {

	// Certificate ID.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// Certificate name.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time in UTC format.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Certificate content.
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// Certificate type.
	// 0: Tencent Cloud-hosted certificate
	// 1: user-added certificate.
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// Certificate expiration time in UTC format.
	CertExpireTime *string `json:"CertExpireTime,omitempty" name:"CertExpireTime"`

	// List of domain names that use this certificate.
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList" list`
}

type CreateLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveCallbackRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCallbackRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name, which is a non-empty string.
	// Maximum length: 255 bytes.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Description.
	// Maximum length: 1,024 bytes.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Stream starting callback URL,
	// Protocol-related document: [Event Message Notification](/document/product/267/32744).
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// Stream ending callback URL,
	// Protocol-related document: [Event Message Notification](/document/product/267/32744).
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// Recording callback URL,
	// Protocol-related document: [Event Message Notification](/document/product/267/32744).
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// Screencapturing callback URL,
	// Protocol-related document: [Event Message Notification](/document/product/267/32744).
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// Porn detection callback URL,
	// Protocol-related document: [Event Message Notification](/document/product/267/32741).
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`

	// Callback key, which is shared by callback URLs. For more information on authentication callback, please see the callback format document
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *CreateLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCertRequest struct {
	*tchttp.BaseRequest

	// Certificate type. 0: user-added certificate, 1: Tencent Cloud-hosted certificate.
	// Note: if the certificate type is 0, `HttpsCrt` and `HttpsKey` are required;
	// If the certificate type is 1, the certificate corresponding to `CloudCertId` will be used first. If `CloudCertId` is empty, `HttpsCrt` and `HttpsKey` will be used.
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// Certificate name.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Certificate content, i.e., public key.
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// Private key.
	HttpsKey *string `json:"HttpsKey,omitempty" name:"HttpsKey"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Tencent Cloud-hosted certificate ID.
	CloudCertId *string `json:"CloudCertId,omitempty" name:"CloudCertId"`
}

func (r *CreateLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate ID
		CertId *int64 `json:"CertId,omitempty" name:"CertId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name. This parameter must be set for multi-domain name push.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Recording start time, which is China standard time and should be URL-encoded (RFC3986). For example, the encoding of 2017-01-01 10:10:01 is 2017-01-01+10%3a10%3a01.
	// In scheduled recording mode, this field must be set; in real-time video recording mode, this field is ignored.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Recording end time, which is China standard time and should be URL-encoded (RFC3986). For example, the encoding of 2017-01-01 10:30:01 is 2017-01-01+10%3a30%3a01.
	// In scheduled recording mode, this field must be set; in real-time video recording mode, this field is optional. If the recording is set to real-time video recording mode through the Highlight parameter, the end time set should not be more than 30 minutes after the current time. If the set end time is more than 30 minutes after the current time, earlier than the current time or left blank, the actual end time will be 30 minutes after the current time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Recording type.
	// "video": Audio-video recording **(default)**.
	// "audio": audio recording.
	// In both scheduled and real-time video recording modes, this parameter is valid and is not case sensitive.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Recording file format. Valid values:
	// "flv" **(default)**, "hls", "mp4", "aac", "mp3".
	// In both scheduled and real-time video recording modes, this parameter is valid and is not case sensitive.
	FileFormat *string `json:"FileFormat,omitempty" name:"FileFormat"`

	// Mark for enabling real-time video recording mode.
	// 0: Real-time video recording mode is not enabled, i.e., the scheduled recording mode is used **(default)**. See [Sample 1](#.E7.A4.BA.E4.BE.8B1-.E5.88.9B.E5.BB.BA.E5.AE.9A.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1).
	// 1: Real-time video recording mode is enabled. See [Sample 2](#.E7.A4.BA.E4.BE.8B2-.E5.88.9B.E5.BB.BA.E5.AE.9E.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1).
	Highlight *int64 `json:"Highlight,omitempty" name:"Highlight"`

	// Mark for enabling A+B=C mixed stream recording.
	// 0: A+B=C mixed stream recording is not enabled **(default)**.
	// 1: A+B=C mixed stream recording is enabled.
	// In both scheduled and real-time video recording modes, this parameter is valid.
	MixStream *int64 `json:"MixStream,omitempty" name:"MixStream"`

	// Recording stream parameter. The following parameters are supported currently:
	// record_interval: Recording interval in seconds. Value range: 1,800–7,200
	// storage_time: Recording file duration in seconds
	// eg. record_interval=3600&storage_time=2592000
	// Note: The parameter needs url encode.
	// In both scheduled and real-time video recording modes, this parameter is valid.
	StreamParam *string `json:"StreamParam,omitempty" name:"StreamParam"`
}

func (r *CreateLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID, which uniquely identifies the recording task globally.
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	// Note: If the parameter is a non-empty string, the rule will be only applicable to the particular stream.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *CreateLiveRecordRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name, which is a non-empty string.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Message description
	Description *string `json:"Description,omitempty" name:"Description"`

	// FLV recording parameter, which is set when FLV recording is enabled.
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// HLS recording parameter, which is set when HLS recording is enabled.
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// Mp4 recording parameter, which is set when Mp4 recording is enabled.
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// AAC recording parameter, which is set when AAC recording is enabled.
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`

	// 0: LVB,
	// 1: LCB.
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// HLS-specific recording parameter.
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitempty" name:"HlsSpecialParam"`

	// Mp3 recording parameter, which is set when Mp3 recording is enabled.
	Mp3Param *RecordParam `json:"Mp3Param,omitempty" name:"Mp3Param"`
}

func (r *CreateLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	// Note: if this parameter is a non-empty string, the rule will take effect only for the particular stream.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *CreateLiveSnapshotRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveSnapshotRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name, which is a non-empty string.
	// Maximum length: 255 bytes.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// COS `AppId`.
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// COS bucket name.
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// COS region.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Description.
	// Maximum length: 1,024 bytes.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Screencapturing interval in seconds. Default value: 10s.
	// Value range: 5–600s.
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// Screenshot width. Default value: 0 (original width).
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Screenshot height. Default value: 0 (original height).
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Whether to enable porn detection. 0: no, 1: yes. Default value: 0
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`
}

func (r *CreateLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default. If you only bind a domain name, leave this parameter empty.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name. If only the domain name or path is bound, leave this parameter blank.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Designates an existing template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveTranscodeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveTranscodeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name, such as 900 900p. This can be only a combination of letters and digits.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Video bitrate. Value range: 100–8,000.
	// Note: The bitrate must be a multiple of 100.
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Video encoding format. Valid values: h264, h265. Default value: h264.
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// Audio encoding in ACC format. Default value: original audio format.
	// Note: This parameter will take effect later.
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// Audio bitrate. Value range: 0–500. Default value: 0.
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// Template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Width. Default value: 0.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Whether to keep the video. 0: no; 1: yes. Default value: 1.
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// Whether to keep the audio. 0: no; 1: yes. Default value: 1.
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// Height. Default value: 0.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Frame rate. Default value: 0.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// Keyframe interval in seconds. Original interval by default
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// Whether to rotate. 0: no; 1: yes. Default value: 0.
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// Encoding quality:
	// baseline/main/high. Default value: baseline.
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// Whether to not exceed the original bitrate. 0: no; 1: yes. Default value: 0.
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// Whether to not exceed the original height. 0: no; 1: yes. Default value: 0.
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// Whether to not exceed the original frame rate. 0: no; 1: yes. Default value: 0.
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// Whether it is a top speed codec template. 0: no, 1: yes. Default value: 0.
	AiTransCode *int64 `json:"AiTransCode,omitempty" name:"AiTransCode"`

	// `VideoBitrate` minus top speed codec bitrate. Value range: 0.1–0.5.
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitempty" name:"AdaptBitratePercent"`
}

func (r *CreateLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Watermark ID, which is the `WatermarkId` returned by the [AddLiveWatermark](/document/product/267/30154) API.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveWatermarkRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveWatermarkRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveWatermarkRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveWatermarkRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveWatermarkRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelayInfo struct {

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Delay time in seconds.
	DelayInterval *uint64 `json:"DelayInterval,omitempty" name:"DelayInterval"`

	// Creation time in UTC format.
	// Note: Beijing time is 8 hours ahead of UTC.
	// Example: 2019-06-18T12:00:00Z (20:00:00, June 18, 2019, Beijing time).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Expiration time in UTC format.
	// Note: Beijing time is 8 hours ahead of UTC.
	// Example: 2019-06-18T12:00:00Z (20:00:00, June 18, 2019, Beijing time).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Current status,
	// -1: Expired,
	// 1: Effective.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DeleteLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DeleteLiveCallbackRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveCallbackRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCertRequest struct {
	*tchttp.BaseRequest

	// Certificate ID.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`
}

func (r *DeleteLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name to be deleted.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Type. 0: push, 1: playback.
	DomainType *uint64 `json:"DomainType,omitempty" name:"DomainType"`
}

func (r *DeleteLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Task ID, which uniquely identifies the recording task globally.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	// Domain name+AppName+StreamName uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	// Domain name+AppName+StreamName uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	// Domain name+AppName+StreamName uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveRecordRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveSnapshotRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveSnapshotRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	// For transcoding at the domain name level, domain name+AppName+StreamName uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	// Domain name+AppName+StreamName+TemplateId uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	// Domain name+AppName+StreamName+TemplateId uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Template ID.
	// Domain name+AppName+StreamName+TemplateId uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveTranscodeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveTranscodeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// Watermark ID.
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

func (r *DeleteLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveWatermarkRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveWatermarkRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCallbackRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule information list.
		Rules []*CallBackRuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Callback template information.
		Template *CallBackTemplateInfo `json:"Template,omitempty" name:"Template"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCallbackTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template information list.
		Templates []*CallBackTemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertRequest struct {
	*tchttp.BaseRequest

	// Certificate ID.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`
}

func (r *DescribeLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate information.
		CertInfo *CertInfo `json:"CertInfo,omitempty" name:"CertInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate information list.
		CertInfoSet []*CertInfo `json:"CertInfoSet,omitempty" name:"CertInfoSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDelayInfoListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveDelayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDelayInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDelayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Delayed playback information list.
		DelayInfoList []*DelayInfo `json:"DelayInfoList,omitempty" name:"DelayInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDelayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDelayInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate information.
		DomainCertInfo *DomainCertInfo `json:"DomainCertInfo,omitempty" name:"DomainCertInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Domain name information.
		DomainInfo *DomainInfo `json:"DomainInfo,omitempty" name:"DomainInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainsRequest struct {
	*tchttp.BaseRequest

	// Filter by domain name status. 0: disabled, 1: enabled.
	DomainStatus *uint64 `json:"DomainStatus,omitempty" name:"DomainStatus"`

	// Filter by domain name type. 0: push. 1: playback
	DomainType *uint64 `json:"DomainType,omitempty" name:"DomainType"`

	// Number of entries per page. Value range: 10–100. Default value: 10.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Page number to get. Value range: 1–100000. Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 0: LVB, 1: LCB. Default value: 0.
	IsDelayLive *uint64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// Domain name prefix.
	DomainPrefix *string `json:"DomainPrefix,omitempty" name:"DomainPrefix"`
}

func (r *DescribeLiveDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of results.
		AllCount *uint64 `json:"AllCount,omitempty" name:"AllCount"`

		// List of domain name details.
		DomainList []*DomainInfo `json:"DomainList,omitempty" name:"DomainList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveForbidStreamListRequest struct {
	*tchttp.BaseRequest

	// Page number to get. Default value: 1.
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Maximum value: 100. 
	// Value: any integer between 1 and 100.
	// Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeLiveForbidStreamListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveForbidStreamListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveForbidStreamListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of eligible ones.
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *int64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Page number.
		PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries displayed per page.
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// List of forbidden streams.
		ForbidStreamList []*ForbidStreamInfo `json:"ForbidStreamList,omitempty" name:"ForbidStreamList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveForbidStreamListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveForbidStreamListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePlayAuthKeyRequest struct {
	*tchttp.BaseRequest

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLivePlayAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePlayAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePlayAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Playback authentication key information.
		PlayAuthKeyInfo *PlayAuthKeyInfo `json:"PlayAuthKeyInfo,omitempty" name:"PlayAuthKeyInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLivePlayAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePlayAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePushAuthKeyRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLivePushAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePushAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePushAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Push authentication key information.
		PushAuthKeyInfo *PushAuthKeyInfo `json:"PushAuthKeyInfo,omitempty" name:"PushAuthKeyInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLivePushAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePushAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveRecordRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of rules.
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recording template information.
		Template *RecordTemplateInfo `json:"Template,omitempty" name:"Template"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplatesRequest struct {
	*tchttp.BaseRequest

	// Whether it is an LCB template
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`
}

func (r *DescribeLiveRecordTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recording template information list.
		Templates []*RecordTemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveSnapshotRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule list.
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Screencapturing template information.
		Template *SnapshotTemplateInfo `json:"Template,omitempty" name:"Template"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveSnapshotTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Screencapturing template list.
		Templates []*SnapshotTemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamEventListRequest struct {
	*tchttp.BaseRequest

	// Start time. 
	// In UTC format, such as 2018-12-29T19:00:00Z.
	// This supports querying the history of 60 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time.
	// In UTC format, such as 2018-12-29T20:00:00Z.
	// This cannot be after the current time and cannot be more than 30 days after the start time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name; query with wildcard (*) is not supported; fuzzy match by default.
	// The IsStrict field can be used to change to exact query.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Page number to get.
	// Default value: 1.
	// Note: Currently, query for up to 10,000 entries is supported.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page.
	// Maximum value: 100.
	// Value range: any integer between 1 and 100.
	// Default value: 10.
	// Note: Currently, query for up to 10,000 entries is supported.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Whether to filter. No filtering by default.
	// 0: No filtering at all.
	// 1: Filter out the failing streams and return only the successful ones.
	IsFilter *int64 `json:"IsFilter,omitempty" name:"IsFilter"`

	// Whether to query exactly. Fuzzy match by default.
	// 0: Fuzzy match.
	// 1: Exact query.
	// Note: This parameter takes effect when StreamName is used.
	IsStrict *int64 `json:"IsStrict,omitempty" name:"IsStrict"`

	// Whether to display in ascending order by end time. Descending order by default.
	// 0: Descending.
	// 1: Ascending.
	IsAsc *int64 `json:"IsAsc,omitempty" name:"IsAsc"`
}

func (r *DescribeLiveStreamEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamEventListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of streaming events.
		EventList []*StreamEventInfo `json:"EventList,omitempty" name:"EventList" list`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Total number of eligible ones.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamEventListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineListRequest struct {
	*tchttp.BaseRequest

	// Push domain name. If you use multiple paths, enter the `DomainName`.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default. If you use multiple paths, enter the `AppName`.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Page number to get. Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Maximum value: 100. 
	// Value: any integer between 10 and 100.
	// Default value: 10.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Stream name, which is used for exact query.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamOnlineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of eligible ones.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries displayed per page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Active push information list.
		OnlineInfo []*StreamOnlineInfo `json:"OnlineInfo,omitempty" name:"OnlineInfo" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamOnlineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListRequest struct {
	*tchttp.BaseRequest

	// Your push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// End time.
	// In UTC format, such as 2016-06-30T19:00:00Z.
	// This cannot be after the current time.
	// Note: The difference between EndTime and StartTime cannot be greater than 30 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time. 
	// In UTC format, such as 2016-06-29T19:00:00Z.
	// This supports querying data in the past 60 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default. Fuzzy match is not supported.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Page number to get.
	// Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page.
	// Maximum value: 100.
	// Value range: any integer between 1 and 100.
	// Default value: 10.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Stream name, which supports fuzzy match.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamPublishedListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPublishedListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Push record information.
		PublishInfo []*StreamName `json:"PublishInfo,omitempty" name:"PublishInfo" list`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Total number of eligible ones.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamPublishedListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPublishedListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamStateRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Your push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamStateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Stream status,
	// active: active
	// inactive: Inactive
	// forbid: forbidden.
		StreamState *string `json:"StreamState,omitempty" name:"StreamState"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamStateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveTranscodeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of transcoding rules.
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template information.
		Template *TemplateInfo `json:"Template,omitempty" name:"Template"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of transcoding templates.
		Templates []*TemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// Watermark ID.
	WatermarkId *uint64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

func (r *DescribeLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Watermark information.
		Watermark *WatermarkInfo `json:"Watermark,omitempty" name:"Watermark"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveWatermarkRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Watermarking rule list.
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarkRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarksRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveWatermarksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of watermarks.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Watermark information list.
		WatermarkList []*WatermarkInfo `json:"WatermarkList,omitempty" name:"WatermarkList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DomainCertInfo struct {

	// Certificate ID.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// Certificate name.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time in UTC format.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Certificate content.
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// Certificate type.
	// 0: user-added certificate
	// 1: Tencent Cloud-hosted certificate.
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// Certificate expiration time in UTC format.
	CertExpireTime *string `json:"CertExpireTime,omitempty" name:"CertExpireTime"`

	// Domain name that uses this certificate.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Certificate status
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DomainInfo struct {

	// LVB domain name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Domain name type. 0: push, 1: playback
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Domain name status. 0: disabled, 1: enabled.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether there is a CNAME record pointing to a fixed rule. 0: no, 1: yes.
	BCName *uint64 `json:"BCName,omitempty" name:"BCName"`

	// Domain name corresponding to the CNAME record
	TargetDomain *string `json:"TargetDomain,omitempty" name:"TargetDomain"`

	// Playback region. This parameter is valid only if `Type` is 1.
	// 1: Mainland China, 2: global, 3: outside Mainland China.
	PlayType *int64 `json:"PlayType,omitempty" name:"PlayType"`

	// 0: LVB,
	// 1: LCB.
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// Information of currently used CNAME record
	CurrentCName *string `json:"CurrentCName,omitempty" name:"CurrentCName"`

	// Whether it is a leased domain name
	RentTag *int64 `json:"RentTag,omitempty" name:"RentTag"`

	// Expiration time of leased domain name
	RentExpireTime *string `json:"RentExpireTime,omitempty" name:"RentExpireTime"`

	// 
	IsMiniProgramLive *int64 `json:"IsMiniProgramLive,omitempty" name:"IsMiniProgramLive"`
}

type DropLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Your acceleration domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DropLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DropLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DropLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableLiveDomainRequest struct {
	*tchttp.BaseRequest

	// LVB domain name to be enabled.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *EnableLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveDomainRequest struct {
	*tchttp.BaseRequest

	// LVB domain name to be disabled.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *ForbidLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForbidLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Your push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Time to resume the stream in UTC format, such as 2018-11-29T19:00:00Z.
	// Notes:
	// 1. The duration of forbidding is 7 days by default and can be up to 90 days.
	// 2. The Beijing time is in UTC+8. This value should be in the format as required by ISO 8601. For more information, please see [ISO Date and Time Format](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	ResumeTime *string `json:"ResumeTime,omitempty" name:"ResumeTime"`

	// Reason for forbidding.
	// Note: Be sure to enter the reason for forbidding to avoid any faulty operations.
	// Length limit: 2,048 bytes.
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *ForbidLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForbidLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidStreamInfo struct {

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Forbidding expiration time.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type HlsSpecialParam struct {

	// HLS timeout period.
	FlowContinueDuration *uint64 `json:"FlowContinueDuration,omitempty" name:"FlowContinueDuration"`
}

type ModifyLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Stream starting callback URL.
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// Stream ending callback URL.
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// Recording callback URL.
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// Screencapturing callback URL.
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// Porn detection callback URL.
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`

	// Callback key, which is shared by callback URLs. For more information on authentication callback, please see the callback format document.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *ModifyLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCertRequest struct {
	*tchttp.BaseRequest

	// Certificate ID.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Certificate type. 0: user-added certificate, 1: Tencent Cloud-hosted certificate.
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// Certificate name.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Certificate content, i.e., public key.
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// Private key.
	HttpsKey *string `json:"HttpsKey,omitempty" name:"HttpsKey"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Certificate ID.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// Status. 0: off, 1: on.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayAuthKeyRequest struct {
	*tchttp.BaseRequest

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Whether to enable. 0: disabled; 1: enabled.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Authentication key.
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// Validity period in seconds.
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`

	// Authentication backkey.
	AuthBackKey *string `json:"AuthBackKey,omitempty" name:"AuthBackKey"`
}

func (r *ModifyLivePlayAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePlayAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayDomainRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Pull domain name type. 1: Mainland China. 2: global, 3: outside Mainland China
	PlayType *int64 `json:"PlayType,omitempty" name:"PlayType"`
}

func (r *ModifyLivePlayDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePlayDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePushAuthKeyRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Whether to enable. 0: disabled; 1: enabled.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Master authentication key.
	MasterAuthKey *string `json:"MasterAuthKey,omitempty" name:"MasterAuthKey"`

	// Backup authentication key.
	BackupAuthKey *string `json:"BackupAuthKey,omitempty" name:"BackupAuthKey"`

	// Validity period in seconds.
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`
}

func (r *ModifyLivePushAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePushAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePushAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePushAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePushAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Message description
	Description *string `json:"Description,omitempty" name:"Description"`

	// FLV recording parameter, which is set when FLV recording is enabled.
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// HLS recording parameter, which is set when HLS recording is enabled.
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// Mp4 recording parameter, which is set when Mp4 recording is enabled.
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// AAC recording parameter, which is set when AAC recording is enabled.
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`

	// Custom HLS recording parameter.
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitempty" name:"HlsSpecialParam"`

	// Mp3 recording parameter, which is set when Mp3 recording is enabled.
	Mp3Param *RecordParam `json:"Mp3Param,omitempty" name:"Mp3Param"`
}

func (r *ModifyLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	// Maximum length: 255 bytes.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Description.
	// Maximum length: 1,024 bytes.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Screencapturing interval in seconds. Default value: 10s.
	// Value range: 5–600s.
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// Screenshot width. Default value: 0 (original width).
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Screenshot height. Default value: 0 (original height).
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Whether to enable porn detection. 0: no, 1: yes.
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// COS `AppId`.
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// COS bucket name.
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// COS region.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
}

func (r *ModifyLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Video encoding format:
	// h264/h265.
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// Audio encoding format:
	// aac/mp3.
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// Audio bitrate. Value range: 0–500. Default value: 0.
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// Template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Video bitrate. Value range: 100–8,000
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Width. Value range: 0–3,000
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Whether to keep the video. 0: no; 1: yes. Default value: 1.
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// Whether to keep the audio. 0: no; 1: yes. Default value: 1.
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// Height. Value range: 0–3,000
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Frame rate. Value range: 0–200
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// Keyframe interval in seconds. Value range: 0–50
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// Rotation angle. Valid values: 0, 90, 180, 270
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// Encoding quality:
	// baseline/main/high.
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// Whether to not exceed the original bitrate. 0: no; 1: yes. Default value: 0.
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// Whether to not exceed the original height. 0: no; 1: yes. Default value: 0.
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// Whether to not exceed the original frame rate. 0: no; 1: yes. Default value: 0.
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// VideoBitrate minus TESHD bitrate. Value range: 0.1–0.5.
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitempty" name:"AdaptBitratePercent"`
}

func (r *ModifyLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PlayAuthKeyInfo struct {

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Whether to enable. 0: disabled; 1: enabled.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Authentication key.
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// Validity period in seconds.
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`

	// Authentication BackKey.
	AuthBackKey *string `json:"AuthBackKey,omitempty" name:"AuthBackKey"`
}

type PublishTime struct {

	// Push time
	// In UTC format, for example: 2018-06-29T19:00:00Z.
	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
}

type PushAuthKeyInfo struct {

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Whether to enable. 0: disabled; 1: enabled.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Master authentication key.
	MasterAuthKey *string `json:"MasterAuthKey,omitempty" name:"MasterAuthKey"`

	// Backup authentication key.
	BackupAuthKey *string `json:"BackupAuthKey,omitempty" name:"BackupAuthKey"`

	// Validity period in seconds.
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`
}

type RecordParam struct {

	// Recording interval.
	// In seconds. Default value: 1,800.
	// Value range: 300–7,200.
	// This parameter is not valid for HLS, and a file is generated from push start to push end when HLS is recorded.
	RecordInterval *int64 `json:"RecordInterval,omitempty" name:"RecordInterval"`

	// Recording storage duration.
	// In seconds. Value range: 0–93,312,000.
	// 0 represents permanent storage.
	StorageTime *int64 `json:"StorageTime,omitempty" name:"StorageTime"`

	// Whether to enable recording in the current format. 0: no; 1: yes. Default value: 0.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// VOD subapplication ID.
	VodSubAppId *int64 `json:"VodSubAppId,omitempty" name:"VodSubAppId"`
}

type RecordTemplateInfo struct {

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Message description
	Description *string `json:"Description,omitempty" name:"Description"`

	// FLV recording parameter.
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// HLS recording parameter.
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// Mp4 recording parameter.
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// AAC recording parameter.
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`

	// 0: LVB,
	// 1: LCB.
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// Custom HLS recording parameter.
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitempty" name:"HlsSpecialParam"`

	// Mp3 recording parameter.
	Mp3Param *RecordParam `json:"Mp3Param,omitempty" name:"Mp3Param"`
}

type ResumeDelayLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *ResumeDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeDelayLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeDelayLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeDelayLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Your acceleration domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *ResumeLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {

	// Rule creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Rule update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

type SnapshotTemplateInfo struct {

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Screencapturing interval in seconds. Value range: 5–300s.
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// Screenshot width. Value range: 0–3000. 0: original width and fit to the original aspect ratio
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Screenshot height. Value range: 0–2,000. 0: original height and fit to the original aspect ratio
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Whether to enable porn detection. 0: no, 1: yes.
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// COS `AppId`.
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// COS bucket name.
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// COS region.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Template description
	Description *string `json:"Description,omitempty" name:"Description"`
}

type StopLiveRecordRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Task ID, which uniquely identifies the recording task globally.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopLiveRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopLiveRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StreamEventInfo struct {

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Push start time.
	// In UTC format.
	// Example: 2019-01-07T12:00:00Z.
	StreamStartTime *string `json:"StreamStartTime,omitempty" name:"StreamStartTime"`

	// Push end time.
	// In UTC format.
	// Example: 2019-01-07T15:00:00Z.
	StreamEndTime *string `json:"StreamEndTime,omitempty" name:"StreamEndTime"`

	// Stop reason.
	StopReason *string `json:"StopReason,omitempty" name:"StopReason"`

	// Push duration in seconds.
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Host IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// Resolution.
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

type StreamName struct {

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push start time.
	// In UTC format.
	// Example: 2019-01-07T12:00:00Z.
	StreamStartTime *string `json:"StreamStartTime,omitempty" name:"StreamStartTime"`

	// Push end time.
	// In UTC format.
	// Example: 2019-01-07T15:00:00Z.
	StreamEndTime *string `json:"StreamEndTime,omitempty" name:"StreamEndTime"`

	// Stop reason.
	StopReason *string `json:"StopReason,omitempty" name:"StopReason"`

	// Push duration in seconds.
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Host IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// Resolution.
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

type StreamOnlineInfo struct {

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Push time list
	PublishTimeList []*PublishTime `json:"PublishTimeList,omitempty" name:"PublishTimeList" list`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

type TemplateInfo struct {

	// Video encoding format:
	// h264/h265.
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// Video bitrate in Kbps. Value range: 100–8,000
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Audio encoding format: AAC/MP3
	// aac/mp3.
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// Audio bitrate. Value range: 0–500
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// Width. Value range: 0–3,000
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Height. Value range: 0–3,000
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Frame rate. Value range: 0–200
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// Keyframe interval in seconds. Value range: 1–50
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// Rotation angle. Valid values: 0, 90, 180, 270
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// Encoding quality:
	// baseline/main/high.
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// Whether to not exceed the original bitrate. 0: no; 1: yes.
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// Whether to not exceed the original height. 0: no; 1: yes.
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// Whether to not exceed the original frame rate. 0: no; 1: yes.
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// Whether to keep the video. 0: no; 1: yes.
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// Whether to keep the audio. 0: no; 1: yes.
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Template description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether it is a TESHD template. 0: no; 1: yes. Default value: 0.
	AiTransCode *int64 `json:"AiTransCode,omitempty" name:"AiTransCode"`

	// VideoBitrate minus TESHD bitrate. Value range: 0.1–0.5.
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitempty" name:"AdaptBitratePercent"`
}

type UnBindLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *UnBindLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// Watermark ID.
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

	// Watermark image URL.
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// Display position: X-axis offset.
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// Display position: Y-axis offset.
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// Watermark name.
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// Watermark width or its percentage of the live streaming video width. It is recommended to just specify either height or width as the other will be scaled proportionally to avoid distortions.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Watermark height or its percentage of the live streaming video width. It is recommended to just specify either height or width as the other will be scaled proportionally to avoid distortions.
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

func (r *UpdateLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WatermarkInfo struct {

	// Watermark ID.
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

	// Watermark image URL.
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// Display position: X-axis offset.
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// Display position: Y-axis offset.
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// Watermark name.
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// Current status. 0: not used. 1: in use.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Watermark width
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Watermark height
	Height *int64 `json:"Height,omitempty" name:"Height"`
}
