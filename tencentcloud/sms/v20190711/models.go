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

package v20190711

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AddSignStatus struct {

	// Signature ID.
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// Signature application ID.
	SignApplyId *uint64 `json:"SignApplyId,omitempty" name:"SignApplyId"`
}

type AddSmsSignRequest struct {
	*tchttp.BaseRequest

	// Signature name.
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// Signature type. Each of these types is followed by their `DocumentType` (identity document type) option:
	// 0: company (0, 1, 2, 3).
	// 1: app (0, 1, 2, 3, 4).
	// 2: website (0, 1, 2, 3, 5).
	// 3: WeChat Official Account or WeChat Mini Program (0, 1, 2, 3, 6).
	// 4: trademark (7).
	// 5: governmental/public institution or others (2, 3).
	// Note: the identity document type must be selected according to the correspondence; otherwise, the review will fail.
	SignType *uint64 `json:"SignType,omitempty" name:"SignType"`

	// Identity document type:
	// 0: 3-in-1 license.
	// 1: business license.
	// 2: organization code certificate.
	// 3: certificate of unified social credit code.
	// 4: screenshot of application backend management (for personal app).
	// 5: screenshot of website ICP filing backend (for personal website).
	// 6: screenshot of WeChat Mini Program settings page (for personal WeChat Mini Program).
	// 7: trademark registration certificate.
	DocumentType *uint64 `json:"DocumentType,omitempty" name:"DocumentType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitempty" name:"International"`

	// Signature use:
	// 0: for self-use.
	// 1: for others.
	UsedMethod *uint64 `json:"UsedMethod,omitempty" name:"UsedMethod"`

	// You should Base64-encode the image of the identity document corresponding to the signature first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	ProofImage *string `json:"ProofImage,omitempty" name:"ProofImage"`

	// Authorization letter, which should be submitted if `UsedMethod` is for others.
	// You should Base64-encode the image first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	// Note: this field will take effect only when `UsedMethod` is 1 (for others).
	CommissionImage *string `json:"CommissionImage,omitempty" name:"CommissionImage"`

	// Signature application remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *AddSmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSmsSignRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSmsSignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Signature addition response
		AddSignStatus *AddSignStatus `json:"AddSignStatus,omitempty" name:"AddSignStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSmsSignResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSmsTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Template content.
	TemplateContent *string `json:"TemplateContent,omitempty" name:"TemplateContent"`

	// SMS type. 0: ordinary SMS, 1: marketing SMS.
	SmsType *uint64 `json:"SmsType,omitempty" name:"SmsType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitempty" name:"International"`

	// Template remarks, such as reason for application and use case.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *AddSmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSmsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SMS template addition response packet body
		AddTemplateStatus *AddTemplateStatus `json:"AddTemplateStatus,omitempty" name:"AddTemplateStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSmsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTemplateStatus struct {

	// Template parameter
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type CallbackStatusStatistics struct {

	// SMS receipts.
	CallbackCount *uint64 `json:"CallbackCount,omitempty" name:"CallbackCount"`

	// Successfully submitted SMS messages.
	RequestSuccessCount *uint64 `json:"RequestSuccessCount,omitempty" name:"RequestSuccessCount"`

	// Failed SMS receipts.
	CallbackFailCount *uint64 `json:"CallbackFailCount,omitempty" name:"CallbackFailCount"`

	// Successful SMS receipts.
	CallbackSuccessCount *uint64 `json:"CallbackSuccessCount,omitempty" name:"CallbackSuccessCount"`

	// Internal carrier errors.
	InternalErrorCount *uint64 `json:"InternalErrorCount,omitempty" name:"InternalErrorCount"`

	// Invalid or empty mobile numbers.
	InvalidNumberCount *uint64 `json:"InvalidNumberCount,omitempty" name:"InvalidNumberCount"`

	// Errors such as out-of-service or power-off.
	ShutdownErrorCount *uint64 `json:"ShutdownErrorCount,omitempty" name:"ShutdownErrorCount"`

	// Blacklisted mobile numbers.
	BlackListCount *uint64 `json:"BlackListCount,omitempty" name:"BlackListCount"`

	// Carrier frequency limit hits.
	FrequencyLimitCount *uint64 `json:"FrequencyLimitCount,omitempty" name:"FrequencyLimitCount"`
}

type CallbackStatusStatisticsRequest struct {
	*tchttp.BaseRequest

	// Start time of pull in the format of `yyyymmddhh` accurate to the hour.
	StartDateTime *uint64 `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// End time of pull in the format of `yyyymmddhh` accurate to the hour.
	// Note: `EndDataTime` must be later than `StartDateTime`.
	EndDataTime *uint64 `json:"EndDataTime,omitempty" name:"EndDataTime"`

	// SMS `SdkAppid` actually generated after an application is added in the [SMS Console](https://console.cloud.tencent.com/sms/smslist), such as 1400006666.
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// Upper limit.
	// Note: this parameter is currently fixed at 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *CallbackStatusStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CallbackStatusStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CallbackStatusStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Receipt statistics response packet body.
		CallbackStatusStatistics *CallbackStatusStatistics `json:"CallbackStatusStatistics,omitempty" name:"CallbackStatusStatistics"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CallbackStatusStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CallbackStatusStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSignStatus struct {

	// Deletion status information.
	DeleteStatus *string `json:"DeleteStatus,omitempty" name:"DeleteStatus"`

	// Deletion time in seconds in the format of UNIX timestamp.
	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`
}

type DeleteSmsSignRequest struct {
	*tchttp.BaseRequest

	// ID of signature to be deleted.
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`
}

func (r *DeleteSmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSmsSignRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSmsSignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Signature deletion response.
		DeleteSignStatus *DeleteSignStatus `json:"DeleteSignStatus,omitempty" name:"DeleteSignStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSmsSignResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSmsTemplateRequest struct {
	*tchttp.BaseRequest

	// ID of template to be deleted.
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteSmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSmsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template deletion response.
		DeleteTemplateStatus *DeleteTemplateStatus `json:"DeleteTemplateStatus,omitempty" name:"DeleteTemplateStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSmsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateStatus struct {

	// Deletion status information.
	DeleteStatus *string `json:"DeleteStatus,omitempty" name:"DeleteStatus"`

	// Deletion time in seconds in the format of UNIX timestamp.
	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`
}

type DescribeSignListStatus struct {

	// Signature ID
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// Whether it is Global SMS. Valid values:
	// 0: Mainland China SMS.
	// 1: Global SMS
	International *uint64 `json:"International,omitempty" name:"International"`

	// Signature application status. Valid values:
	// 0: approved.
	// -1: rejected or failed.
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// Review reply, i.e., response given by the reviewer, which is usually the reason for rejection.
	ReviewReply *string `json:"ReviewReply,omitempty" name:"ReviewReply"`

	// Signature name.
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// Application submission time in the format of UNIX timestamp in seconds.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeSmsSignListRequest struct {
	*tchttp.BaseRequest

	// Signature ID array.
	SignIdSet []*uint64 `json:"SignIdSet,omitempty" name:"SignIdSet" list`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitempty" name:"International"`
}

func (r *DescribeSmsSignListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsSignListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsSignListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Response for getting signature information
		DescribeSignListStatusSet []*DescribeSignListStatus `json:"DescribeSignListStatusSet,omitempty" name:"DescribeSignListStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmsSignListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsSignListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsTemplateListRequest struct {
	*tchttp.BaseRequest

	// Template ID array.
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitempty" name:"TemplateIdSet" list`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitempty" name:"International"`
}

func (r *DescribeSmsTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsTemplateListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Response for getting SMS signature information
		DescribeTemplateStatusSet []*DescribeTemplateListStatus `json:"DescribeTemplateStatusSet,omitempty" name:"DescribeTemplateStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmsTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsTemplateListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateListStatus struct {

	// Template ID
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Whether it is Global SMS. Valid values:
	// 0: Mainland China SMS.
	// 1: Global SMS
	International *uint64 `json:"International,omitempty" name:"International"`

	// Signature application status. Valid values:
	// 0: approved.
	// -1: rejected or failed.
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// Review reply, i.e., response given by the reviewer, which is usually the reason for rejection.
	ReviewReply *string `json:"ReviewReply,omitempty" name:"ReviewReply"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Application submission time in the format of UNIX timestamp in seconds.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ModifySignStatus struct {

	// Signature ID
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// Signature modification application ID
	SignApplyId *string `json:"SignApplyId,omitempty" name:"SignApplyId"`
}

type ModifySmsSignRequest struct {
	*tchttp.BaseRequest

	// ID of signature to be modified.
	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`

	// Signature name.
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// Signature type. Each of these types is followed by their `DocumentType` (identity document type) option:
	// 0: company (0, 1, 2, 3).
	// 1: app (0, 1, 2, 3, 4).
	// 2: website (0, 1, 2, 3, 5).
	// 3: WeChat Official Account or WeChat Mini Program (0, 1, 2, 3, 6).
	// 4: trademark (7).
	// 5: governmental/public institution or others (2, 3).
	// Note: the identity document type must be selected according to the correspondence; otherwise, the review will fail.
	SignType *uint64 `json:"SignType,omitempty" name:"SignType"`

	// Identity document type:
	// 0: 3-in-1 license.
	// 1: business license.
	// 2: organization code certificate.
	// 3: certificate of unified social credit code.
	// 4: screenshot of application backend management (for personal app).
	// 5: screenshot of website ICP filing backend (for personal website).
	// 6: screenshot of WeChat Mini Program settings page (for personal WeChat Mini Program).
	// 7: trademark registration certificate.
	DocumentType *uint64 `json:"DocumentType,omitempty" name:"DocumentType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitempty" name:"International"`

	// Signature use:
	// 0: for self-use.
	// 1: for others.
	UsedMethod *uint64 `json:"UsedMethod,omitempty" name:"UsedMethod"`

	// You should Base64-encode the image of the identity document corresponding to the signature first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	ProofImage *string `json:"ProofImage,omitempty" name:"ProofImage"`

	// Authorization letter, which should be submitted if `UsedMethod` is for others.
	// You should Base64-encode the image first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	// Note: this field will take effect only when `UsedMethod` is 1 (for others).
	CommissionImage *string `json:"CommissionImage,omitempty" name:"CommissionImage"`

	// Signature application remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifySmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySmsSignRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySmsSignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Signature modification response
		ModifySignStatus *ModifySignStatus `json:"ModifySignStatus,omitempty" name:"ModifySignStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySmsSignResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySmsTemplateRequest struct {
	*tchttp.BaseRequest

	// ID of template to be modified.
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// New template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// New template content.
	TemplateContent *string `json:"TemplateContent,omitempty" name:"TemplateContent"`

	// SMS type. 0: ordinary SMS, 1: marketing SMS.
	SmsType *uint64 `json:"SmsType,omitempty" name:"SmsType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitempty" name:"International"`

	// Template remarks, such as reason for application and use case.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifySmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySmsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template parameter modification response
		ModifyTemplateStatus *ModifyTemplateStatus `json:"ModifyTemplateStatus,omitempty" name:"ModifyTemplateStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySmsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateStatus struct {

	// Template parameter
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type PullSmsReplyStatus struct {

	// SMS code number extension, which is not activated by default. If you need to activate it, please contact [SMS Helper](https://cloud.tencent.com/document/product/382/3773).
	ExtendCode *string `json:"ExtendCode,omitempty" name:"ExtendCode"`

	// Country (or region) code.
	NationCode *string `json:"NationCode,omitempty" name:"NationCode"`

	// Mobile number in the e.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// SMS signature.
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// User reply.
	ReplyContent *string `json:"ReplyContent,omitempty" name:"ReplyContent"`

	// Reply time (e.g., 2019-10-08 17:18:37).
	ReplyTime *string `json:"ReplyTime,omitempty" name:"ReplyTime"`

	// Reply time in seconds in the format of UNIX timestamp.
	ReplyUnixTime *uint64 `json:"ReplyUnixTime,omitempty" name:"ReplyUnixTime"`
}

type PullSmsReplyStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest

	// Pull start time in seconds in the format of UNIX timestamp.
	SendDateTime *uint64 `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Target mobile number in the e.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// SMS `SdkAppid` actually generated after an application is added in the [SMS Console](https://console.cloud.tencent.com/sms/smslist), such as 1400006666.
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

func (r *PullSmsReplyStatusByPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsReplyStatusByPhoneNumberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsReplyStatusByPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Reply status response set.
		PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitempty" name:"PullSmsReplyStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullSmsReplyStatusByPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsReplyStatusByPhoneNumberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsReplyStatusRequest struct {
	*tchttp.BaseRequest

	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// SMS `SdkAppid` actually generated after an application is added in the [SMS Console](https://console.cloud.tencent.com/sms/smslist), such as 1400006666.
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

func (r *PullSmsReplyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsReplyStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsReplyStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Reply status response set.
		PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitempty" name:"PullSmsReplyStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullSmsReplyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsReplyStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatus struct {

	// Actual time of SMS receipt by user.
	UserReceiveTime *string `json:"UserReceiveTime,omitempty" name:"UserReceiveTime"`

	// Actual time of SMS receipt by user in seconds in the format of UNIX timestamp.
	UserReceiveUnixTime *uint64 `json:"UserReceiveUnixTime,omitempty" name:"UserReceiveUnixTime"`

	// Country (or region) code.
	NationCode *string `json:"NationCode,omitempty" name:"NationCode"`

	// Mobile number in the e.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PurePhoneNumber *string `json:"PurePhoneNumber,omitempty" name:"PurePhoneNumber"`

	// Mobile number in a common format such as 13711112222.
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// ID of the current delivery.
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// Whether the SMS message is actually received. Valid values: SUCCESS (success), FAIL (failure).
	ReportStatus *string `json:"ReportStatus,omitempty" name:"ReportStatus"`

	// Description of SMS receipt by user.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type PullSmsSendStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest

	// Pull start time in seconds in the format of UNIX timestamp.
	SendDateTime *uint64 `json:"SendDateTime,omitempty" name:"SendDateTime"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Target mobile number in the e.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// SMS `SdkAppid` actually generated after an application is added in the [SMS Console](https://console.cloud.tencent.com/sms/smslist), such as 1400006666.
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

func (r *PullSmsSendStatusByPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsSendStatusByPhoneNumberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatusByPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Delivery status response set.
		PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitempty" name:"PullSmsSendStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullSmsSendStatusByPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsSendStatusByPhoneNumberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatusRequest struct {
	*tchttp.BaseRequest

	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// SMS `SdkAppid` actually generated after an application is added in the [SMS Console](https://console.cloud.tencent.com/sms/smslist), such as 1400006666.
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`
}

func (r *PullSmsSendStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsSendStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Delivery status response set.
		PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitempty" name:"PullSmsSendStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullSmsSendStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullSmsSendStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendSmsRequest struct {
	*tchttp.BaseRequest

	// Target mobile number in the e.164 standard in the format of +[country/region code][mobile number]. Up to 200 mobile numbers are supported in one request (which should be all Mainland China mobile numbers or all global mobile numbers).
	// Example: +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitempty" name:"PhoneNumberSet" list`

	// Template ID. You must enter the ID of an approved template, which can be viewed in the [SMS Console](https://console.cloud.tencent.com/sms/smslist).
	TemplateID *string `json:"TemplateID,omitempty" name:"TemplateID"`

	// SMS `SdkAppid` actually generated after an application is added in the [SMS Console](https://console.cloud.tencent.com/sms/smslist), such as 1400006666.
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// The content of SMS signature should be encoded in UTF-8. You must enter an approved signature, which can be viewed in the [SMS Console](https://console.cloud.tencent.com/sms/smslist). Note: this parameter is required for Mainland China SMS.
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// Template parameter. If there is no template parameter, leave this parameter blank.
	TemplateParamSet []*string `json:"TemplateParamSet,omitempty" name:"TemplateParamSet" list`

	// SMS code number extension, which is not activated by default. If you need to activate it, please contact [SMS Helper](https://cloud.tencent.com/document/product/382/3773).
	ExtendCode *string `json:"ExtendCode,omitempty" name:"ExtendCode"`

	// User session content, which can carry context information such as user-side ID and will be returned as-is by the server.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// `senderid` for Global SMS, which is not activated by default. If you need to activate it, please contact [SMS Helper](https://cloud.tencent.com/document/product/382/3773) for assistance. This parameter should be empty for Mainland China SMS.
	SenderId *string `json:"SenderId,omitempty" name:"SenderId"`
}

func (r *SendSmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendSmsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendSmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SMS delivery status.
		SendStatusSet []*SendStatus `json:"SendStatusSet,omitempty" name:"SendStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendSmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendSmsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendStatus struct {

	// Delivery serial number.
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// Mobile number in the e.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// Number of billable SMS messages. For billing rules, please see [Billing Policy](https://cloud.tencent.com/document/product/382/36135).
	Fee *uint64 `json:"Fee,omitempty" name:"Fee"`

	// User session content.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// SMS request error code. For specific meanings, please see Error Codes.
	Code *string `json:"Code,omitempty" name:"Code"`

	// SMS request error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Country code or region code, such as CN and US. If the country code or region code is not obtained, the returned value will be 'DEF' by default. For more information on the supported list, see price overview for non-Mainland China regions.
	IsoCode *string `json:"IsoCode,omitempty" name:"IsoCode"`
}

type SendStatusStatistics struct {

	// Billable SMS message quantity; for example, in 100 successfully submitted SMS messages, if 20 are long messages (over 80 characters) and split into two messages each, then the billable quantity will be 80 * 1 + 20 * 2 = 120.
	FeeCount *uint64 `json:"FeeCount,omitempty" name:"FeeCount"`

	// Submitted SMS messages.
	RequestCount *uint64 `json:"RequestCount,omitempty" name:"RequestCount"`

	// Successfully submitted SMS messages.
	RequestSuccessCount *uint64 `json:"RequestSuccessCount,omitempty" name:"RequestSuccessCount"`
}

type SendStatusStatisticsRequest struct {
	*tchttp.BaseRequest

	// Start time of pull in the format of `yyyymmddhh` accurate to the hour.
	StartDateTime *uint64 `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// End time of pull in the format of `yyyymmddhh` accurate to the hour
	// Note: `EndDataTime` must be later than `StartDateTime`.
	EndDataTime *uint64 `json:"EndDataTime,omitempty" name:"EndDataTime"`

	// SMS `SdkAppid` actually generated after an application is added in the [SMS Console](https://console.cloud.tencent.com/sms/smslist), such as 1400006666.
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// Upper limit.
	// Note: this parameter is currently fixed at 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SendStatusStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendStatusStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendStatusStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Delivery statistics response packet.
		SendStatusStatistics *SendStatusStatistics `json:"SendStatusStatistics,omitempty" name:"SendStatusStatistics"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendStatusStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendStatusStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SmsPackagesStatistics struct {

	// Package creation time in standard time format, such as 2019-10-08 17:18:37.
	PackageCreateTime *string `json:"PackageCreateTime,omitempty" name:"PackageCreateTime"`

	// Package creation time in seconds in the format of UNIX timestamp.
	PackageCreateUnixTime *uint64 `json:"PackageCreateUnixTime,omitempty" name:"PackageCreateUnixTime"`

	// Package effective time in standard time format, such as 2019-10-08 17:18:37.
	PackageEffectiveTime *string `json:"PackageEffectiveTime,omitempty" name:"PackageEffectiveTime"`

	// Package effective time in seconds in the format of UNIX timestamp.
	PackageEffectiveUnixTime *uint64 `json:"PackageEffectiveUnixTime,omitempty" name:"PackageEffectiveUnixTime"`

	// Package expiration time in standard time format, such as 2019-10-08 17:18:37.
	PackageExpiredTime *string `json:"PackageExpiredTime,omitempty" name:"PackageExpiredTime"`

	// Package expiration time in seconds in the format of UNIX timestamp.
	PackageExpiredUnixTime *uint64 `json:"PackageExpiredUnixTime,omitempty" name:"PackageExpiredUnixTime"`

	// Number of SMS messages in package.
	AmountOfPackage *uint64 `json:"AmountOfPackage,omitempty" name:"AmountOfPackage"`

	// 0: gifted package. 1: purchased package.
	TypeOfPackage *uint64 `json:"TypeOfPackage,omitempty" name:"TypeOfPackage"`

	// Package ID.
	PackageId *uint64 `json:"PackageId,omitempty" name:"PackageId"`

	// Current usage.
	CurrentUsage *uint64 `json:"CurrentUsage,omitempty" name:"CurrentUsage"`
}

type SmsPackagesStatisticsRequest struct {
	*tchttp.BaseRequest

	// SMS `SdkAppid` actually generated after an application is added in the [SMS Console](https://console.cloud.tencent.com/sms/smslist), such as 1400006666.
	SmsSdkAppid *string `json:"SmsSdkAppid,omitempty" name:"SmsSdkAppid"`

	// Upper limit (number of packages to be pulled).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SmsPackagesStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SmsPackagesStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SmsPackagesStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Delivery statistics response packet body.
		SmsPackagesStatisticsSet []*SmsPackagesStatistics `json:"SmsPackagesStatisticsSet,omitempty" name:"SmsPackagesStatisticsSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SmsPackagesStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SmsPackagesStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
