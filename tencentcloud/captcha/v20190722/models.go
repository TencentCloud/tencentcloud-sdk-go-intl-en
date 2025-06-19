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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeCaptchaResultRequestParams struct {
	// Fill with fixed value: 9.
	CaptchaType *uint64 `json:"CaptchaType,omitnil,omitempty" name:"CaptchaType"`

	// The user verification ticket returned by the frontend callback function
	Ticket *string `json:"Ticket,omitnil,omitempty" name:"Ticket"`

	// The user public IP obtained from the customer backend server
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// A random string returned by the frontend callback function
	Randstr *string `json:"Randstr,omitnil,omitempty" name:"Randstr"`

	// CAPTCHA's app ID. Log in to the [Captcha console](https://console.cloud.tencent.com/captcha/graphical) and you can view the CaptchaAppId in the "Key" column of the CAPTCHA list.
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// CAPTCHA's app key. Log in to the [Captcha console](https://console.cloud.tencent.com/captcha/graphical) and you can view the AppSecretKey in the "Key" column of the CAPTCHA list. AppSecretKey is the key for CAPTCHA ticket verification performed by the server. Please keep it confidential and do not disclose it to any third parties.
	AppSecretKey *string `json:"AppSecretKey,omitnil,omitempty" name:"AppSecretKey"`

	// Reserved field.
	BusinessId *uint64 `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// Reserved field.
	SceneId *uint64 `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// MAC address or unique identifier of a device
	MacAddress *string `json:"MacAddress,omitnil,omitempty" name:"MacAddress"`

	// Mobile equipment identity number
	Imei *string `json:"Imei,omitnil,omitempty" name:"Imei"`

	// Indicates whether to return the time when the frontend obtains the CAPTCHA. Valid values: 1 (return the time) and others.
	NeedGetCaptchaTime *int64 `json:"NeedGetCaptchaTime,omitnil,omitempty" name:"NeedGetCaptchaTime"`
}

type DescribeCaptchaResultRequest struct {
	*tchttp.BaseRequest
	
	// Fill with fixed value: 9.
	CaptchaType *uint64 `json:"CaptchaType,omitnil,omitempty" name:"CaptchaType"`

	// The user verification ticket returned by the frontend callback function
	Ticket *string `json:"Ticket,omitnil,omitempty" name:"Ticket"`

	// The user public IP obtained from the customer backend server
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// A random string returned by the frontend callback function
	Randstr *string `json:"Randstr,omitnil,omitempty" name:"Randstr"`

	// CAPTCHA's app ID. Log in to the [Captcha console](https://console.cloud.tencent.com/captcha/graphical) and you can view the CaptchaAppId in the "Key" column of the CAPTCHA list.
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// CAPTCHA's app key. Log in to the [Captcha console](https://console.cloud.tencent.com/captcha/graphical) and you can view the AppSecretKey in the "Key" column of the CAPTCHA list. AppSecretKey is the key for CAPTCHA ticket verification performed by the server. Please keep it confidential and do not disclose it to any third parties.
	AppSecretKey *string `json:"AppSecretKey,omitnil,omitempty" name:"AppSecretKey"`

	// Reserved field.
	BusinessId *uint64 `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// Reserved field.
	SceneId *uint64 `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// MAC address or unique identifier of a device
	MacAddress *string `json:"MacAddress,omitnil,omitempty" name:"MacAddress"`

	// Mobile equipment identity number
	Imei *string `json:"Imei,omitnil,omitempty" name:"Imei"`

	// Indicates whether to return the time when the frontend obtains the CAPTCHA. Valid values: 1 (return the time) and others.
	NeedGetCaptchaTime *int64 `json:"NeedGetCaptchaTime,omitnil,omitempty" name:"NeedGetCaptchaTime"`
}

func (r *DescribeCaptchaResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaType")
	delete(f, "Ticket")
	delete(f, "UserIp")
	delete(f, "Randstr")
	delete(f, "CaptchaAppId")
	delete(f, "AppSecretKey")
	delete(f, "BusinessId")
	delete(f, "SceneId")
	delete(f, "MacAddress")
	delete(f, "Imei")
	delete(f, "NeedGetCaptchaTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaResultResponseParams struct {
	// OK indicates verification passed.
	// 7 captcha no match. the passed in Randstr is invalid. please check if the Randstr is consistent with the Randstr returned by the frontend.
	// The passed-in ticket has expired (the valid period of the ticket is 5 minutes). generate the ticket and Randstr again for validation.
	// The passed-in ticket is reused. generate the ticket and Randstr again for verification.
	// 15 decrypt fail. the passed-in Ticket is invalid. please check if the Ticket is consistent with the Ticket returned by the frontend.
	// 16 appid-ticket mismatch. the passed in CaptchaAppId is incorrect. please check if the CaptchaAppId is consistent with the CaptchaAppId passed in by the frontend, and ensure that the CaptchaAppId is obtained from the verification code console [verification management] -> [basic configuration].
	// 21 diff invoice verification exception. possible reasons: (1) if the Ticket contains the trerror prefix, generally because the user has a poor network connection, resulting in the frontend's automatic disaster recovery and generation of a disaster recovery Ticket. the business side may skip or post-process as needed. (2) if the Ticket does not include the trerror prefix, it is because the security risk of the request was detected by the CAPTCHA-intl risk control system. the business side may intercept as needed.
	// 100 appid-secretkey-ticket mismatch. parameter validation error. (1) please check whether the CaptchaAppId and AppSecretKey are correct. the CaptchaAppId and AppSecretKey need to be obtained from verification code console > verification management > basic configuration. (2) please check whether the passed-in ticket is generated by the passed-in CaptchaAppId.
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// Status description and verification error message.
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// In invisible verification mode, this parameter returns the verification result.
	// EvilLevel=0 indicates that the request is not malicious.
	// The parameter EvilLevel = 100 indicates that the request is malicious.
	EvilLevel *int64 `json:"EvilLevel,omitnil,omitempty" name:"EvilLevel"`

	// Frontend retrieval time of the captcha-intl, timestamp format.
	GetCaptchaTime *int64 `json:"GetCaptchaTime,omitnil,omitempty" name:"GetCaptchaTime"`

	// Blocking type
	// Note: This field may return null, indicating that no valid values can be obtained.
	EvilBitmap *int64 `json:"EvilBitmap,omitnil,omitempty" name:"EvilBitmap"`

	// The time when the CAPTCHA is submitted.
	SubmitCaptchaTime *int64 `json:"SubmitCaptchaTime,omitnil,omitempty" name:"SubmitCaptchaTime"`

	// Device risk category.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeviceRiskCategory *string `json:"DeviceRiskCategory,omitnil,omitempty" name:"DeviceRiskCategory"`

	// CAPTCHA-Intl score.
	// Note:The score ranges from 0 to 100 (e.g., 20, 70, 90).
	// A higher score indicates a greater probability that the interaction was initiated by a bot or represents a bot attack.
	// A lower score indicates a greater probability that the interaction was performed by a real human user.
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCaptchaResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaResultResponseParams `json:"Response"`
}

func (r *DescribeCaptchaResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}