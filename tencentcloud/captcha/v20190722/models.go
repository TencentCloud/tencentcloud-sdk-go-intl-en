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
	// It must be `9` here.
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
	
	// It must be `9` here.
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
	// `1 OK`: Verification passed
	// `7 captcha no match`: The passed in `Randstr` is invalid. Make sure it is the same as the `Randstr` returned from the frontend.
	// `8 ticket expired`: The `Ticket` has expired. A ticket is valid for five minutes. Please generate a new `Ticket` and `Randstr`.
	// `9 ticket reused`: The specified `Ticket` has been used. Please generate a new `Ticket` and `Randstr`.
	// `15 decrypt fail`: The specified `Ticket` is invalid. Make sure it's the same as the Ticket returned from the frontend. 
	// `16 appid-ticket mismatch`: The specified `CaptchaAppId` is invalid. Make sure it's the same as the `CaptchaAppId` returned from the frontend. You can obtain it from the CAPTCHA console in **Verification management** > **Basic configuration**.
	// `21 diff`. Ticket verification error. Possible reasons: 1) If the ticket contains the `terror` prefix, it's usually the case that a disaster recovery ticket is generated due to the network connection problems of the user. You can choose to ignore it or verify again. 2) If the ticket does not include the `terror` prefix, Captcha detects security risk on this request . You can choose to block it or not.
	// `100 appid-secretkey-ticket mismatch`: Parameter error. 1) Make sure `CaptchaAppId` and `AppSecretKey` are correct. `CaptchaAppId` and `AppSecretKey` in the CAPTACHA console under **Verification management** > **Basic configuration**. 2) Make sure the passed-in `Ticket` is generated by using the passed-in `CaptchaAppId`.
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// Status description and verification error message
	// Note: This field may return `null`, indicating that no valid value was found.
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// This parameter returns the result of imperceptible verification. This parameter is not available for Tencent Cloud International yet.
	// `0`: The request is trusted.
	// `100`: The request is malicious.
	// Note: This field may return `null`, indicating that no valid value was found.
	EvilLevel *int64 `json:"EvilLevel,omitnil,omitempty" name:"EvilLevel"`

	// The timestamp when the frontend obtains the CAPTCHA.
	// Note: This field may return `null`, indicating that no valid value was found.
	GetCaptchaTime *int64 `json:"GetCaptchaTime,omitnil,omitempty" name:"GetCaptchaTime"`

	// Blocking type
	// Note: This field may return null, indicating that no valid values can be obtained.
	EvilBitmap *int64 `json:"EvilBitmap,omitnil,omitempty" name:"EvilBitmap"`

	// The time when the CAPTCHA is submitted.
	SubmitCaptchaTime *int64 `json:"SubmitCaptchaTime,omitnil,omitempty" name:"SubmitCaptchaTime"`

	// Device Risk Category
	DeviceRiskCategory *string `json:"DeviceRiskCategory,omitnil,omitempty" name:"DeviceRiskCategory"`

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