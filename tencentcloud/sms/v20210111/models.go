// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20210111

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AddSignStatus struct {
	// Signature ID.
	SignId *uint64 `json:"SignId,omitnil,omitempty" name:"SignId"`
}

// Predefined struct for user
type AddSmsSignRequestParams struct {
	// Signature name.
	// Note: you cannot apply for an approved or pending signature again.
	SignName *string `json:"SignName,omitnil,omitempty" name:"SignName"`

	// Signature type. Each of these types is followed by their `DocumentType` (identity certificate type) option:
	// 0: company. Valid values of `DocumentType` include 0 and 1.
	// 1: app. Valid values of `DocumentType` include 0, 1, 2, 3, and 4.
	// 2: website. Valid values of `DocumentType` include 0, 1, 2, 3, and 5.
	// 3: WeChat Official Account. Valid values of `DocumentType` include 0, 1, 2, 3, and 8.
	// 4: trademark. Valid values of `DocumentType` include 7.
	// 5: government/public institution/other. Valid values of `DocumentType` include 2 and 3.
	// 6: WeChat Mini Program. Valid values of `DocumentType` include 0, 1, 2, 3, and 6.
	// Note: the identity certificate type must be selected according to the correspondence; otherwise, the review will fail.
	SignType *uint64 `json:"SignType,omitnil,omitempty" name:"SignType"`

	// Identity certificate type:
	// 0: three-in-one licence.
	// 1: business license.
	// 2: organization code certificate.
	// 3: social credit code certificate.
	// 4: screenshot of application backend management (for personal app).
	// 5: screenshot of website ICP filing backend (for personal website).
	// 6: screenshot of WeChat Mini Program settings page (for personal WeChat Mini Program).
	// 7: trademark registration certificate.
	// 8: screenshot of WeChat Official Account settings page (for personal WeChat Official Account).
	DocumentType *uint64 `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Signature purpose:
	// 0: for personal use.
	// 1: for others.
	SignPurpose *uint64 `json:"SignPurpose,omitnil,omitempty" name:"SignPurpose"`

	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	ProofImage *string `json:"ProofImage,omitnil,omitempty" name:"ProofImage"`

	// Power of attorney, which should be submitted if `SignPurpose` is for use by others.
	// You should Base64-encode the image first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	// Note: this field will take effect only when `SignPurpose` is 1 (for user by others).
	CommissionImage *string `json:"CommissionImage,omitnil,omitempty" name:"CommissionImage"`

	// Signature application remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type AddSmsSignRequest struct {
	*tchttp.BaseRequest
	
	// Signature name.
	// Note: you cannot apply for an approved or pending signature again.
	SignName *string `json:"SignName,omitnil,omitempty" name:"SignName"`

	// Signature type. Each of these types is followed by their `DocumentType` (identity certificate type) option:
	// 0: company. Valid values of `DocumentType` include 0 and 1.
	// 1: app. Valid values of `DocumentType` include 0, 1, 2, 3, and 4.
	// 2: website. Valid values of `DocumentType` include 0, 1, 2, 3, and 5.
	// 3: WeChat Official Account. Valid values of `DocumentType` include 0, 1, 2, 3, and 8.
	// 4: trademark. Valid values of `DocumentType` include 7.
	// 5: government/public institution/other. Valid values of `DocumentType` include 2 and 3.
	// 6: WeChat Mini Program. Valid values of `DocumentType` include 0, 1, 2, 3, and 6.
	// Note: the identity certificate type must be selected according to the correspondence; otherwise, the review will fail.
	SignType *uint64 `json:"SignType,omitnil,omitempty" name:"SignType"`

	// Identity certificate type:
	// 0: three-in-one licence.
	// 1: business license.
	// 2: organization code certificate.
	// 3: social credit code certificate.
	// 4: screenshot of application backend management (for personal app).
	// 5: screenshot of website ICP filing backend (for personal website).
	// 6: screenshot of WeChat Mini Program settings page (for personal WeChat Mini Program).
	// 7: trademark registration certificate.
	// 8: screenshot of WeChat Official Account settings page (for personal WeChat Official Account).
	DocumentType *uint64 `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Signature purpose:
	// 0: for personal use.
	// 1: for others.
	SignPurpose *uint64 `json:"SignPurpose,omitnil,omitempty" name:"SignPurpose"`

	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	ProofImage *string `json:"ProofImage,omitnil,omitempty" name:"ProofImage"`

	// Power of attorney, which should be submitted if `SignPurpose` is for use by others.
	// You should Base64-encode the image first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	// Note: this field will take effect only when `SignPurpose` is 1 (for user by others).
	CommissionImage *string `json:"CommissionImage,omitnil,omitempty" name:"CommissionImage"`

	// Signature application remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *AddSmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SignName")
	delete(f, "SignType")
	delete(f, "DocumentType")
	delete(f, "International")
	delete(f, "SignPurpose")
	delete(f, "ProofImage")
	delete(f, "CommissionImage")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSmsSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSmsSignResponseParams struct {
	// Signature addition response
	AddSignStatus *AddSignStatus `json:"AddSignStatus,omitnil,omitempty" name:"AddSignStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddSmsSignResponse struct {
	*tchttp.BaseResponse
	Response *AddSmsSignResponseParams `json:"Response"`
}

func (r *AddSmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSmsTemplateRequestParams struct {
	// Template name.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template content.
	TemplateContent *string `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// SMS type. 1: Marketing SMS, 2: Notification SMS, 3: OTP SMS.
	// Note: To further improve the quality of SMS service and increase the approval rate of SMS templates, Tencent Cloud SMS template types will be optimized to "OTP", "Notification", and "Marketing" from May 16, 2024. You can Refer to [Optimization of Tencent Cloud SMS Template Types](https://www.tencentcloud.com/document/product/382/60410). Customers who newly register for SMS service must strictly refer to the new SMS type to apply for SMS template.
	SmsType *uint64 `json:"SmsType,omitnil,omitempty" name:"SmsType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Template remarks, such as reason for application and use case.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type AddSmsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template name.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template content.
	TemplateContent *string `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// SMS type. 1: Marketing SMS, 2: Notification SMS, 3: OTP SMS.
	// Note: To further improve the quality of SMS service and increase the approval rate of SMS templates, Tencent Cloud SMS template types will be optimized to "OTP", "Notification", and "Marketing" from May 16, 2024. You can Refer to [Optimization of Tencent Cloud SMS Template Types](https://www.tencentcloud.com/document/product/382/60410). Customers who newly register for SMS service must strictly refer to the new SMS type to apply for SMS template.
	SmsType *uint64 `json:"SmsType,omitnil,omitempty" name:"SmsType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Template remarks, such as reason for application and use case.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *AddSmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TemplateContent")
	delete(f, "SmsType")
	delete(f, "International")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSmsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSmsTemplateResponseParams struct {
	// SMS template addition response body
	AddTemplateStatus *AddTemplateStatus `json:"AddTemplateStatus,omitnil,omitempty" name:"AddTemplateStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddSmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *AddSmsTemplateResponseParams `json:"Response"`
}

func (r *AddSmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSmsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddTemplateStatus struct {
	// Template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type CallbackStatusStatistics struct {
	// Messages with receipt.
	CallbackCount *uint64 `json:"CallbackCount,omitnil,omitempty" name:"CallbackCount"`

	// Successfully submitted SMS messages.
	RequestSuccessCount *uint64 `json:"RequestSuccessCount,omitnil,omitempty" name:"RequestSuccessCount"`

	// Failed receipts.
	CallbackFailCount *uint64 `json:"CallbackFailCount,omitnil,omitempty" name:"CallbackFailCount"`

	// Successful receipts.
	CallbackSuccessCount *uint64 `json:"CallbackSuccessCount,omitnil,omitempty" name:"CallbackSuccessCount"`

	// Carrier's internal error.
	InternalErrorCount *uint64 `json:"InternalErrorCount,omitnil,omitempty" name:"InternalErrorCount"`

	// Invalid numbers.
	InvalidNumberCount *uint64 `json:"InvalidNumberCount,omitnil,omitempty" name:"InvalidNumberCount"`

	// Errors such as out-of-service or power-off.
	ShutdownErrorCount *uint64 `json:"ShutdownErrorCount,omitnil,omitempty" name:"ShutdownErrorCount"`

	// Blocked mobile numbers.
	BlackListCount *uint64 `json:"BlackListCount,omitnil,omitempty" name:"BlackListCount"`

	// Carrier rate limit hits.
	FrequencyLimitCount *uint64 `json:"FrequencyLimitCount,omitnil,omitempty" name:"FrequencyLimitCount"`
}

// Predefined struct for user
type CallbackStatusStatisticsRequestParams struct {
	// Start time in the format of `yyyymmddhh` accurate to the hour, such as 2021050113 (13:00 on May 1, 2021).
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time in the format of `yyyymmddhh` accurate to the hour, such as 2021050118 (18:00 on May 1, 2021).
	// Note: `EndTime` must be after `BeginTime`, and the difference should not exceed 32 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Upper limit.
	// Note: this parameter is currently fixed at 0.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type CallbackStatusStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Start time in the format of `yyyymmddhh` accurate to the hour, such as 2021050113 (13:00 on May 1, 2021).
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time in the format of `yyyymmddhh` accurate to the hour, such as 2021050118 (18:00 on May 1, 2021).
	// Note: `EndTime` must be after `BeginTime`, and the difference should not exceed 32 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Upper limit.
	// Note: this parameter is currently fixed at 0.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *CallbackStatusStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallbackStatusStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "SmsSdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallbackStatusStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallbackStatusStatisticsResponseParams struct {
	// Receipt statistics response body.
	CallbackStatusStatistics *CallbackStatusStatistics `json:"CallbackStatusStatistics,omitnil,omitempty" name:"CallbackStatusStatistics"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CallbackStatusStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *CallbackStatusStatisticsResponseParams `json:"Response"`
}

func (r *CallbackStatusStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallbackStatusStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSignStatus struct {
	// Deletion status information.
	DeleteStatus *string `json:"DeleteStatus,omitnil,omitempty" name:"DeleteStatus"`

	// Deleted time in seconds in the format of UNIX timestamp.
	DeleteTime *uint64 `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`
}

// Predefined struct for user
type DeleteSmsSignRequestParams struct {
	// ID of the signature to be deleted.
	SignId *uint64 `json:"SignId,omitnil,omitempty" name:"SignId"`
}

type DeleteSmsSignRequest struct {
	*tchttp.BaseRequest
	
	// ID of the signature to be deleted.
	SignId *uint64 `json:"SignId,omitnil,omitempty" name:"SignId"`
}

func (r *DeleteSmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmsSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSmsSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmsSignResponseParams struct {
	// Signature deletion response
	DeleteSignStatus *DeleteSignStatus `json:"DeleteSignStatus,omitnil,omitempty" name:"DeleteSignStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSmsSignResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSmsSignResponseParams `json:"Response"`
}

func (r *DeleteSmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmsSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmsTemplateRequestParams struct {
	// ID of the template to be deleted.
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteSmsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the template to be deleted.
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteSmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSmsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmsTemplateResponseParams struct {
	// Template deletion response
	DeleteTemplateStatus *DeleteTemplateStatus `json:"DeleteTemplateStatus,omitnil,omitempty" name:"DeleteTemplateStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSmsTemplateResponseParams `json:"Response"`
}

func (r *DeleteSmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateStatus struct {
	// Deletion status information.
	DeleteStatus *string `json:"DeleteStatus,omitnil,omitempty" name:"DeleteStatus"`

	// Deleted time in seconds in the format of UNIX timestamp.
	DeleteTime *uint64 `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`
}

// Predefined struct for user
type DescribePhoneNumberInfoRequestParams struct {
	// A parameter used to query mobile numbers in E.164 format (+[country/region code][subscriber number]). Up to 200 mobile numbers can be queried at a time.
	// Take the number +8613711112222 as an example. “86” is the country code (with a “+” sign in its front) and “13711112222” is the subscriber number.
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitnil,omitempty" name:"PhoneNumberSet"`
}

type DescribePhoneNumberInfoRequest struct {
	*tchttp.BaseRequest
	
	// A parameter used to query mobile numbers in E.164 format (+[country/region code][subscriber number]). Up to 200 mobile numbers can be queried at a time.
	// Take the number +8613711112222 as an example. “86” is the country code (with a “+” sign in its front) and “13711112222” is the subscriber number.
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitnil,omitempty" name:"PhoneNumberSet"`
}

func (r *DescribePhoneNumberInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePhoneNumberInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PhoneNumberSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePhoneNumberInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePhoneNumberInfoResponseParams struct {
	// A parameter used to obtain mobile number information.
	PhoneNumberInfoSet []*PhoneNumberInfo `json:"PhoneNumberInfoSet,omitnil,omitempty" name:"PhoneNumberInfoSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePhoneNumberInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePhoneNumberInfoResponseParams `json:"Response"`
}

func (r *DescribePhoneNumberInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePhoneNumberInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSignListStatus struct {
	// Signature ID.
	SignId *uint64 `json:"SignId,omitnil,omitempty" name:"SignId"`

	// Whether it is Global SMS. 0: Mainland China SMS; 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Signature application status. Valid values: 0: approved; 1: under review.
	// -1: application rejected or failed.
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// Review reply, i.e., response given by the reviewer, which is usually the reason for rejection.
	ReviewReply *string `json:"ReviewReply,omitnil,omitempty" name:"ReviewReply"`

	// Signature name.
	SignName *string `json:"SignName,omitnil,omitempty" name:"SignName"`

	// Application submission time in the format of UNIX timestamp in seconds.
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type DescribeSmsSignListRequestParams struct {
	// Signature ID array.
	// Note: the maximum length of the array is 100 by default.
	SignIdSet []*uint64 `json:"SignIdSet,omitnil,omitempty" name:"SignIdSet"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`
}

type DescribeSmsSignListRequest struct {
	*tchttp.BaseRequest
	
	// Signature ID array.
	// Note: the maximum length of the array is 100 by default.
	SignIdSet []*uint64 `json:"SignIdSet,omitnil,omitempty" name:"SignIdSet"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`
}

func (r *DescribeSmsSignListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsSignListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SignIdSet")
	delete(f, "International")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmsSignListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsSignListResponseParams struct {
	// Response for getting signature information
	DescribeSignListStatusSet []*DescribeSignListStatus `json:"DescribeSignListStatusSet,omitnil,omitempty" name:"DescribeSignListStatusSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSmsSignListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmsSignListResponseParams `json:"Response"`
}

func (r *DescribeSmsSignListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsSignListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsTemplateListRequestParams struct {
	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Array of template IDs. If the array is empty, the template list information will be queried by default. You need to use the `Limit` and `Offset` fields to set the query range.
	// <dx-alert infotype="notice" title="Note">The max array length is 100 by default.</dx-alert>
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitnil,omitempty" name:"TemplateIdSet"`

	// Upper limit. Maximum value: 100.
	// Note: it is 0 by default and is enabled when `TemplateIdSet` is empty.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	// Note: it is 0 by default and is enabled when `TemplateIdSet` is empty.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSmsTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Array of template IDs. If the array is empty, the template list information will be queried by default. You need to use the `Limit` and `Offset` fields to set the query range.
	// <dx-alert infotype="notice" title="Note">The max array length is 100 by default.</dx-alert>
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitnil,omitempty" name:"TemplateIdSet"`

	// Upper limit. Maximum value: 100.
	// Note: it is 0 by default and is enabled when `TemplateIdSet` is empty.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	// Note: it is 0 by default and is enabled when `TemplateIdSet` is empty.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSmsTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "International")
	delete(f, "TemplateIdSet")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmsTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsTemplateListResponseParams struct {
	// Response for getting SMS template information
	DescribeTemplateStatusSet []*DescribeTemplateListStatus `json:"DescribeTemplateStatusSet,omitnil,omitempty" name:"DescribeTemplateStatusSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSmsTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmsTemplateListResponseParams `json:"Response"`
}

func (r *DescribeSmsTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateListStatus struct {
	// Template ID.
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Whether it is Global SMS. 0: Mainland China SMS; 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Template application status. Valid values: 0: approved and effective; 1: under review; 2: approved but to be effective; -1: application rejected or failed.
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// Review reply, i.e., response given by the reviewer, which is usually the reason for rejection.
	ReviewReply *string `json:"ReviewReply,omitnil,omitempty" name:"ReviewReply"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Application submission time in the format of UNIX timestamp in seconds.
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Template content.
	TemplateContent *string `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`
}

type ModifySignStatus struct {
	// Signature ID.
	SignId *uint64 `json:"SignId,omitnil,omitempty" name:"SignId"`
}

// Predefined struct for user
type ModifySmsSignRequestParams struct {
	// ID of the signature to be modified.
	SignId *uint64 `json:"SignId,omitnil,omitempty" name:"SignId"`

	// Signature name.
	SignName *string `json:"SignName,omitnil,omitempty" name:"SignName"`

	// Signature type. Each of these types is followed by their `DocumentType` (identity certificate type) option:
	// 0: company. Valid values of `DocumentType` include 0 and 1.
	// 1: app. Valid values of `DocumentType` include 0, 1, 2, 3, and 4.
	// 2: website. Valid values of `DocumentType` include 0, 1, 2, 3, and 5.
	// 3: WeChat Official Account. Valid values of `DocumentType` include 0, 1, 2, 3, and 8.
	// 4: trademark. Valid values of `DocumentType` include 7.
	// 5: government/public institution/other. Valid values of `DocumentType` include 2 and 3.
	// 6: WeChat Mini Program. Valid values of `DocumentType` include 0, 1, 2, 3, and 6.
	// Note: the identity certificate type must be selected according to the correspondence; otherwise, the review will fail.
	SignType *uint64 `json:"SignType,omitnil,omitempty" name:"SignType"`

	// Identity certificate type:
	// 0: three-in-one.
	// 1: business license.
	// 2: organization code certificate.
	// 3: social credit code certificate.
	// 4: screenshot of application backend management (for personal app).
	// 5: screenshot of website ICP filing backend (for personal website).
	// 6: screenshot of WeChat Mini Program settings page (for personal WeChat Mini Program).
	// 7: trademark registration certificate.
	// 8: screenshot of WeChat Official Account settings page (for personal WeChat Official Account).
	DocumentType *uint64 `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// A parameter used to specify whether it is Global SMS:
	// `0`: Chinese mainland SMS.
	// `1`: Global SMS.
	// Note: the value of this parameter must be consistent with the `International` value of the signature to be modified. This parameter cannot be used to directly change a Chinese mainland signature to an international signature.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Signature purpose:
	// 0: for personal use.
	// 1: for others.
	SignPurpose *uint64 `json:"SignPurpose,omitnil,omitempty" name:"SignPurpose"`

	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	ProofImage *string `json:"ProofImage,omitnil,omitempty" name:"ProofImage"`

	// Power of attorney, which should be submitted if `SignPurpose` is for use by others.
	// You should Base64-encode the image first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	// Note: this field will take effect only when `SignPurpose` is 1 (for user by others).
	CommissionImage *string `json:"CommissionImage,omitnil,omitempty" name:"CommissionImage"`

	// Signature application remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifySmsSignRequest struct {
	*tchttp.BaseRequest
	
	// ID of the signature to be modified.
	SignId *uint64 `json:"SignId,omitnil,omitempty" name:"SignId"`

	// Signature name.
	SignName *string `json:"SignName,omitnil,omitempty" name:"SignName"`

	// Signature type. Each of these types is followed by their `DocumentType` (identity certificate type) option:
	// 0: company. Valid values of `DocumentType` include 0 and 1.
	// 1: app. Valid values of `DocumentType` include 0, 1, 2, 3, and 4.
	// 2: website. Valid values of `DocumentType` include 0, 1, 2, 3, and 5.
	// 3: WeChat Official Account. Valid values of `DocumentType` include 0, 1, 2, 3, and 8.
	// 4: trademark. Valid values of `DocumentType` include 7.
	// 5: government/public institution/other. Valid values of `DocumentType` include 2 and 3.
	// 6: WeChat Mini Program. Valid values of `DocumentType` include 0, 1, 2, 3, and 6.
	// Note: the identity certificate type must be selected according to the correspondence; otherwise, the review will fail.
	SignType *uint64 `json:"SignType,omitnil,omitempty" name:"SignType"`

	// Identity certificate type:
	// 0: three-in-one.
	// 1: business license.
	// 2: organization code certificate.
	// 3: social credit code certificate.
	// 4: screenshot of application backend management (for personal app).
	// 5: screenshot of website ICP filing backend (for personal website).
	// 6: screenshot of WeChat Mini Program settings page (for personal WeChat Mini Program).
	// 7: trademark registration certificate.
	// 8: screenshot of WeChat Official Account settings page (for personal WeChat Official Account).
	DocumentType *uint64 `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// A parameter used to specify whether it is Global SMS:
	// `0`: Chinese mainland SMS.
	// `1`: Global SMS.
	// Note: the value of this parameter must be consistent with the `International` value of the signature to be modified. This parameter cannot be used to directly change a Chinese mainland signature to an international signature.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Signature purpose:
	// 0: for personal use.
	// 1: for others.
	SignPurpose *uint64 `json:"SignPurpose,omitnil,omitempty" name:"SignPurpose"`

	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	ProofImage *string `json:"ProofImage,omitnil,omitempty" name:"ProofImage"`

	// Power of attorney, which should be submitted if `SignPurpose` is for use by others.
	// You should Base64-encode the image first, remove the prefix `data:image/jpeg;base64,` from the resulted string, and then use it as the value of this parameter.
	// Note: this field will take effect only when `SignPurpose` is 1 (for user by others).
	CommissionImage *string `json:"CommissionImage,omitnil,omitempty" name:"CommissionImage"`

	// Signature application remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifySmsSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SignId")
	delete(f, "SignName")
	delete(f, "SignType")
	delete(f, "DocumentType")
	delete(f, "International")
	delete(f, "SignPurpose")
	delete(f, "ProofImage")
	delete(f, "CommissionImage")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySmsSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySmsSignResponseParams struct {
	// Signature modification response
	ModifySignStatus *ModifySignStatus `json:"ModifySignStatus,omitnil,omitempty" name:"ModifySignStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySmsSignResponse struct {
	*tchttp.BaseResponse
	Response *ModifySmsSignResponseParams `json:"Response"`
}

func (r *ModifySmsSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySmsTemplateRequestParams struct {
	// ID of the template to be modified.
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// New template name.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// New template content.
	TemplateContent *string `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// SMS type. 1: Marketing SMS, 2: Notification SMS, 3: OTP SMS.
	// Note: To further improve the quality of SMS service and increase the approval rate of SMS templates, Tencent Cloud SMS template types will be optimized to "OTP", "Notification", and "Marketing" from May 16, 2024. You can Refer to [Optimization of Tencent Cloud SMS Template Types](https://www.tencentcloud.com/document/product/382/60410). Customers who newly register for SMS service must strictly refer to the new SMS type to apply for SMS template.
	SmsType *uint64 `json:"SmsType,omitnil,omitempty" name:"SmsType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Template remarks, such as reason for application and use case.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifySmsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the template to be modified.
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// New template name.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// New template content.
	TemplateContent *string `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// SMS type. 1: Marketing SMS, 2: Notification SMS, 3: OTP SMS.
	// Note: To further improve the quality of SMS service and increase the approval rate of SMS templates, Tencent Cloud SMS template types will be optimized to "OTP", "Notification", and "Marketing" from May 16, 2024. You can Refer to [Optimization of Tencent Cloud SMS Template Types](https://www.tencentcloud.com/document/product/382/60410). Customers who newly register for SMS service must strictly refer to the new SMS type to apply for SMS template.
	SmsType *uint64 `json:"SmsType,omitnil,omitempty" name:"SmsType"`

	// Whether it is Global SMS:
	// 0: Mainland China SMS.
	// 1: Global SMS.
	International *uint64 `json:"International,omitnil,omitempty" name:"International"`

	// Template remarks, such as reason for application and use case.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifySmsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "TemplateContent")
	delete(f, "SmsType")
	delete(f, "International")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySmsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySmsTemplateResponseParams struct {
	// Template parameter modification response
	ModifyTemplateStatus *ModifyTemplateStatus `json:"ModifyTemplateStatus,omitnil,omitempty" name:"ModifyTemplateStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySmsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifySmsTemplateResponseParams `json:"Response"`
}

func (r *ModifySmsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateStatus struct {
	// Template ID.
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type PhoneNumberInfo struct {
	// Error code for mobile number information query. `Ok` will be returned if the query is successful.
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Description of the error code for mobile number information query.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Country (or region) code.
	NationCode *string `json:"NationCode,omitnil,omitempty" name:"NationCode"`

	// Subscriber number in normal format such as 13711112222, without any prefix (country or region code).
	SubscriberNumber *string `json:"SubscriberNumber,omitnil,omitempty" name:"SubscriberNumber"`

	// The standardized mobile number in E.164 format after parsing, which is consistent with the parsed number for SMS message delivery. If the parsing fails, the original number will be returned.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Country or region code such as CN and US. If the country or region code cannot be identified, `DEF` will be returned by default.
	IsoCode *string `json:"IsoCode,omitnil,omitempty" name:"IsoCode"`

	// Country code or region name such as China. For more information, see [Global SMS Price Overview](https://intl.cloud.tencent.com/document/product/382/18051?from_cn_redirect=1#.E6.97.A5.E7.BB.93.E5.90.8E.E4.BB.98.E8.B4.B9.3Ca-id.3D.22post-payment.22.3E.3C.2Fa.3E)
	IsoName *string `json:"IsoName,omitnil,omitempty" name:"IsoName"`
}

type PullSmsReplyStatus struct {
	// SMS code number extension, which is not activated by default. If you need to activate it, please contact [SMS Helper](https://intl.cloud.tencent.com/document/product/382/3773?from_cn_redirect=1#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81).
	ExtendCode *string `json:"ExtendCode,omitnil,omitempty" name:"ExtendCode"`

	// Country (or region) code.
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// Mobile number in the E.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// SMS signature name.
	SignName *string `json:"SignName,omitnil,omitempty" name:"SignName"`

	// User reply.
	ReplyContent *string `json:"ReplyContent,omitnil,omitempty" name:"ReplyContent"`

	// Reply time in seconds in the format of UNIX timestamp.
	ReplyTime *uint64 `json:"ReplyTime,omitnil,omitempty" name:"ReplyTime"`

	// User's mobile number in a common format such as 13711112222.
	SubscriberNumber *string `json:"SubscriberNumber,omitnil,omitempty" name:"SubscriberNumber"`
}

// Predefined struct for user
type PullSmsReplyStatusByPhoneNumberRequestParams struct {
	// Pull start time in seconds in the format of UNIX timestamp.
	// Note: the data for the last 7 days can be pulled at most.
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Target mobile number in the E.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Pull end time in seconds in the format of UNIX timestamp.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type PullSmsReplyStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest
	
	// Pull start time in seconds in the format of UNIX timestamp.
	// Note: the data for the last 7 days can be pulled at most.
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Target mobile number in the E.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Pull end time in seconds in the format of UNIX timestamp.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *PullSmsReplyStatusByPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsReplyStatusByPhoneNumberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PhoneNumber")
	delete(f, "SmsSdkAppId")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsReplyStatusByPhoneNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsReplyStatusByPhoneNumberResponseParams struct {
	// Reply status response set.
	PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitnil,omitempty" name:"PullSmsReplyStatusSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PullSmsReplyStatusByPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *PullSmsReplyStatusByPhoneNumberResponseParams `json:"Response"`
}

func (r *PullSmsReplyStatusByPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsReplyStatusByPhoneNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsReplyStatusRequestParams struct {
	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`
}

type PullSmsReplyStatusRequest struct {
	*tchttp.BaseRequest
	
	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`
}

func (r *PullSmsReplyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsReplyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "SmsSdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsReplyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsReplyStatusResponseParams struct {
	// Reply status response set.
	PullSmsReplyStatusSet []*PullSmsReplyStatus `json:"PullSmsReplyStatusSet,omitnil,omitempty" name:"PullSmsReplyStatusSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PullSmsReplyStatusResponse struct {
	*tchttp.BaseResponse
	Response *PullSmsReplyStatusResponseParams `json:"Response"`
}

func (r *PullSmsReplyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsReplyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PullSmsSendStatus struct {
	// Actual time of SMS receipt by user in seconds in the format of UNIX timestamp.
	UserReceiveTime *uint64 `json:"UserReceiveTime,omitnil,omitempty" name:"UserReceiveTime"`

	// Country (or region) code.
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// User's mobile number in a common format such as 13711112222.
	SubscriberNumber *string `json:"SubscriberNumber,omitnil,omitempty" name:"SubscriberNumber"`

	// Mobile number in the E.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// ID of the current delivery.
	SerialNo *string `json:"SerialNo,omitnil,omitempty" name:"SerialNo"`

	// Whether the SMS message is actually received. Valid values: SUCCESS (success), FAIL (failure).
	ReportStatus *string `json:"ReportStatus,omitnil,omitempty" name:"ReportStatus"`

	// Description of SMS receipt by user.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// User session content, which is the same as the `SessionContext` in the request and is empty by default. If you need to activate it, contact [SMS Helper](https://intl.cloud.tencent.com/document/product/382/3773?from_cn_redirect=1#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81).
	// Note: This field may return null, indicating that no valid values can be obtained.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

// Predefined struct for user
type PullSmsSendStatusByPhoneNumberRequestParams struct {
	// Pull start time in seconds in the format of UNIX timestamp.
	// Note: the data for the last 7 days can be pulled at most.
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Target mobile number in the E.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Pull end time in seconds in the format of UNIX timestamp.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type PullSmsSendStatusByPhoneNumberRequest struct {
	*tchttp.BaseRequest
	
	// Pull start time in seconds in the format of UNIX timestamp.
	// Note: the data for the last 7 days can be pulled at most.
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Target mobile number in the E.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Pull end time in seconds in the format of UNIX timestamp.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *PullSmsSendStatusByPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsSendStatusByPhoneNumberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PhoneNumber")
	delete(f, "SmsSdkAppId")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsSendStatusByPhoneNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsSendStatusByPhoneNumberResponseParams struct {
	// Delivery status response set.
	PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitnil,omitempty" name:"PullSmsSendStatusSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PullSmsSendStatusByPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *PullSmsSendStatusByPhoneNumberResponseParams `json:"Response"`
}

func (r *PullSmsSendStatusByPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsSendStatusByPhoneNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsSendStatusRequestParams struct {
	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`
}

type PullSmsSendStatusRequest struct {
	*tchttp.BaseRequest
	
	// Maximum number of pulled entries. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`
}

func (r *PullSmsSendStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsSendStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "SmsSdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullSmsSendStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullSmsSendStatusResponseParams struct {
	// Delivery status response set.
	PullSmsSendStatusSet []*PullSmsSendStatus `json:"PullSmsSendStatusSet,omitnil,omitempty" name:"PullSmsSendStatusSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PullSmsSendStatusResponse struct {
	*tchttp.BaseResponse
	Response *PullSmsSendStatusResponseParams `json:"Response"`
}

func (r *PullSmsSendStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullSmsSendStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportConversionRequestParams struct {
	// The SMS SdkAppId generated after an application is created in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as “1400006666”.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// The serial number returned for a message sent.
	SerialNo *string `json:"SerialNo,omitnil,omitempty" name:"SerialNo"`

	// The recipient’s reply time in seconds in the format of UNIX timestamp.
	ConversionTime *uint64 `json:"ConversionTime,omitnil,omitempty" name:"ConversionTime"`
}

type ReportConversionRequest struct {
	*tchttp.BaseRequest
	
	// The SMS SdkAppId generated after an application is created in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as “1400006666”.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// The serial number returned for a message sent.
	SerialNo *string `json:"SerialNo,omitnil,omitempty" name:"SerialNo"`

	// The recipient’s reply time in seconds in the format of UNIX timestamp.
	ConversionTime *uint64 `json:"ConversionTime,omitnil,omitempty" name:"ConversionTime"`
}

func (r *ReportConversionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportConversionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SmsSdkAppId")
	delete(f, "SerialNo")
	delete(f, "ConversionTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportConversionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportConversionResponseParams struct {
	// Response packet for conversion rate reporting.
	ReportConversionStatus *ReportConversionStatus `json:"ReportConversionStatus,omitnil,omitempty" name:"ReportConversionStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReportConversionResponse struct {
	*tchttp.BaseResponse
	Response *ReportConversionResponseParams `json:"Response"`
}

func (r *ReportConversionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportConversionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportConversionStatus struct {
	// Error code. `ok` is returned if the conversion rate is successfully reported.
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Error code description.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type SendSmsRequestParams struct {
	// Target mobile number in the E.164 standard in the format of +[country/region code][mobile number]. Up to 200 mobile numbers are supported in one request (which should be all Chinese mainland mobile numbers or all global mobile numbers). For example, +60198890000, which has a + sign followed by 60 (country/region code) and then by 198890000 (mobile number).
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitnil,omitempty" name:"PhoneNumberSet"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 2400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Template ID, which can be viewed on the **Body Templates** page in [Global SMS](https://console.cloud.tencent.com/smsv2/isms-template). You must enter the ID of an approved template.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// SMS signature information which is encoded in UTF-8. You must enter an approved signature (such as Tencent Cloud). The signing information can be viewed on the **Signatures** page in [Global SMS](https://console.cloud.tencent.com/smsv2/isms-sign).
	SignName *string `json:"SignName,omitnil,omitempty" name:"SignName"`

	// Template parameter. If there is no template parameter, leave this field empty.
	// <dx-alert infotype="notice" title="Note">The number of template parameters should be consistent with that of the template variables of `TemplateId`.</dx-alert>
	TemplateParamSet []*string `json:"TemplateParamSet,omitnil,omitempty" name:"TemplateParamSet"`

	// SMS code number extension, which is not activated by default. If you need to activate it, you can contact [SMS Helper](https://intl.cloud.tencent.com/document/product/382/3773?from_cn_redirect=1#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81).
	ExtendCode *string `json:"ExtendCode,omitnil,omitempty" name:"ExtendCode"`

	// User session content, which can carry context information such as user-side ID and will be returned as-is by the server. Note that the length must be less than 512 bytes.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// For Global SMS, if you have applied for a separate `SenderId`, this parameter is required. By default, the public `SenderId` is used, in which case you don't need to enter this parameter.
	// Note: If your monthly usage reaches the specified threshold, you can apply for an independent `SenderId`. For more information, contact [SMS Helper](https://intl.cloud.tencent.com/document/product/382/3773?from_cn_redirect=1#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81).
	SenderId *string `json:"SenderId,omitnil,omitempty" name:"SenderId"`
}

type SendSmsRequest struct {
	*tchttp.BaseRequest
	
	// Target mobile number in the E.164 standard in the format of +[country/region code][mobile number]. Up to 200 mobile numbers are supported in one request (which should be all Chinese mainland mobile numbers or all global mobile numbers). For example, +60198890000, which has a + sign followed by 60 (country/region code) and then by 198890000 (mobile number).
	PhoneNumberSet []*string `json:"PhoneNumberSet,omitnil,omitempty" name:"PhoneNumberSet"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 2400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Template ID, which can be viewed on the **Body Templates** page in [Global SMS](https://console.cloud.tencent.com/smsv2/isms-template). You must enter the ID of an approved template.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// SMS signature information which is encoded in UTF-8. You must enter an approved signature (such as Tencent Cloud). The signing information can be viewed on the **Signatures** page in [Global SMS](https://console.cloud.tencent.com/smsv2/isms-sign).
	SignName *string `json:"SignName,omitnil,omitempty" name:"SignName"`

	// Template parameter. If there is no template parameter, leave this field empty.
	// <dx-alert infotype="notice" title="Note">The number of template parameters should be consistent with that of the template variables of `TemplateId`.</dx-alert>
	TemplateParamSet []*string `json:"TemplateParamSet,omitnil,omitempty" name:"TemplateParamSet"`

	// SMS code number extension, which is not activated by default. If you need to activate it, you can contact [SMS Helper](https://intl.cloud.tencent.com/document/product/382/3773?from_cn_redirect=1#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81).
	ExtendCode *string `json:"ExtendCode,omitnil,omitempty" name:"ExtendCode"`

	// User session content, which can carry context information such as user-side ID and will be returned as-is by the server. Note that the length must be less than 512 bytes.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// For Global SMS, if you have applied for a separate `SenderId`, this parameter is required. By default, the public `SenderId` is used, in which case you don't need to enter this parameter.
	// Note: If your monthly usage reaches the specified threshold, you can apply for an independent `SenderId`. For more information, contact [SMS Helper](https://intl.cloud.tencent.com/document/product/382/3773?from_cn_redirect=1#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81).
	SenderId *string `json:"SenderId,omitnil,omitempty" name:"SenderId"`
}

func (r *SendSmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PhoneNumberSet")
	delete(f, "SmsSdkAppId")
	delete(f, "TemplateId")
	delete(f, "SignName")
	delete(f, "TemplateParamSet")
	delete(f, "ExtendCode")
	delete(f, "SessionContext")
	delete(f, "SenderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendSmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendSmsResponseParams struct {
	// SMS delivery status.
	SendStatusSet []*SendStatus `json:"SendStatusSet,omitnil,omitempty" name:"SendStatusSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendSmsResponse struct {
	*tchttp.BaseResponse
	Response *SendSmsResponseParams `json:"Response"`
}

func (r *SendSmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendStatus struct {
	// Delivery serial number.
	SerialNo *string `json:"SerialNo,omitnil,omitempty" name:"SerialNo"`

	// Mobile number in the E.164 standard (+[country/region code][mobile number]), such as +8613711112222, which has a + sign followed by 86 (country/region code) and then by 13711112222 (mobile number).
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Number of billable SMS messages. For billing rules, see Billing Policy.
	Fee *uint64 `json:"Fee,omitnil,omitempty" name:"Fee"`

	// User session content.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// SMS request error code. For specific meanings, see [Error Codes](https://intl.cloud.tencent.com/zh/document/product/382/40536#6.-error-code). `Ok` will be returned for successful delivery.
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// SMS request error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Country/Region code, such as CN and US. For unrecognized country/region codes, `DEF` is returned by default. For the specific list of supported values, please see [Global SMS Price Overview](https://intl.cloud.tencent.com/document/product/382/18051?from_cn_redirect=1).
	IsoCode *string `json:"IsoCode,omitnil,omitempty" name:"IsoCode"`
}

type SendStatusStatistics struct {
	// Billable SMS message quantity; for example, in 100 successfully submitted SMS messages, if 20 ones are long messages (over 80 characters) and split into two messages each, then the billable quantity will be 80 * 1 + 20 * 2 = 120.
	FeeCount *uint64 `json:"FeeCount,omitnil,omitempty" name:"FeeCount"`

	// Submitted SMS messages.
	RequestCount *uint64 `json:"RequestCount,omitnil,omitempty" name:"RequestCount"`

	// Successfully submitted SMS messages.
	RequestSuccessCount *uint64 `json:"RequestSuccessCount,omitnil,omitempty" name:"RequestSuccessCount"`
}

// Predefined struct for user
type SendStatusStatisticsRequestParams struct {
	// Start time in the format of `yyyymmddhh` accurate to the hour, such as 2021050113 (13:00 on May 1, 2021).
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time in the format of `yyyymmddhh` accurate to the hour, such as 2021050118 (18:00 on May 1, 2021).
	// Note: `EndTime` must be after `BeginTime`.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Upper limit.
	// Note: this parameter is currently fixed at 0.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type SendStatusStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Start time in the format of `yyyymmddhh` accurate to the hour, such as 2021050113 (13:00 on May 1, 2021).
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time in the format of `yyyymmddhh` accurate to the hour, such as 2021050118 (18:00 on May 1, 2021).
	// Note: `EndTime` must be after `BeginTime`.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitnil,omitempty" name:"SmsSdkAppId"`

	// Upper limit.
	// Note: this parameter is currently fixed at 0.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	// Note: this parameter is currently fixed at 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *SendStatusStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendStatusStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "SmsSdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendStatusStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendStatusStatisticsResponseParams struct {
	// Delivery statistics response body.
	SendStatusStatistics *SendStatusStatistics `json:"SendStatusStatistics,omitnil,omitempty" name:"SendStatusStatistics"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendStatusStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *SendStatusStatisticsResponseParams `json:"Response"`
}

func (r *SendStatusStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendStatusStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}