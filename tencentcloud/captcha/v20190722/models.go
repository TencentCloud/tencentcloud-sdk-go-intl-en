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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeCaptchaResultRequestParams struct {
	// It must be `9` here. You can configure the CAPTCHA types in the console.
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// The user verification ticket returned by the frontend callback function
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// The user public IP obtained from the customer backend server
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// A random string returned by the frontend callback function
	Randstr *string `json:"Randstr,omitempty" name:"Randstr"`

	// CAPTCHA's app ID. Log in to the [Captcha console](https://console.cloud.tencent.com/captcha/graphical) and you can view the CaptchaAppId in the "Key" column of the CAPTCHA list.
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// CAPTCHA's app key. Log in to the [Captcha console](https://console.cloud.tencent.com/captcha/graphical) and you can view the AppSecretKey in the "Key" column of the CAPTCHA list. AppSecretKey is the key for CAPTCHA ticket verification performed by the server. Please keep it confidential and do not disclose it to any third parties.
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// Business ID, which is used to differentiate statistical data when this service is used in multiple businesses of a website or an app.
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// Scenario ID, which is used to differentiate statistical data when this service is used in multiple scenarios of a website or an app.
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// MAC address or unique identifier of a device
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// Mobile equipment identity number
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// Indicates whether to return the time when the frontend obtains the CAPTCHA. Valid values: 1 (return the time) and others.
	NeedGetCaptchaTime *int64 `json:"NeedGetCaptchaTime,omitempty" name:"NeedGetCaptchaTime"`
}

type DescribeCaptchaResultRequest struct {
	*tchttp.BaseRequest
	
	// It must be `9` here. You can configure the CAPTCHA types in the console.
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// The user verification ticket returned by the frontend callback function
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// The user public IP obtained from the customer backend server
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// A random string returned by the frontend callback function
	Randstr *string `json:"Randstr,omitempty" name:"Randstr"`

	// CAPTCHA's app ID. Log in to the [Captcha console](https://console.cloud.tencent.com/captcha/graphical) and you can view the CaptchaAppId in the "Key" column of the CAPTCHA list.
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// CAPTCHA's app key. Log in to the [Captcha console](https://console.cloud.tencent.com/captcha/graphical) and you can view the AppSecretKey in the "Key" column of the CAPTCHA list. AppSecretKey is the key for CAPTCHA ticket verification performed by the server. Please keep it confidential and do not disclose it to any third parties.
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// Business ID, which is used to differentiate statistical data when this service is used in multiple businesses of a website or an app.
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// Scenario ID, which is used to differentiate statistical data when this service is used in multiple scenarios of a website or an app.
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// MAC address or unique identifier of a device
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// Mobile equipment identity number
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// Indicates whether to return the time when the frontend obtains the CAPTCHA. Valid values: 1 (return the time) and others.
	NeedGetCaptchaTime *int64 `json:"NeedGetCaptchaTime,omitempty" name:"NeedGetCaptchaTime"`
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
	//  
	//  
	//  
	//  
	//  
	//  
	//  
	//  
	//  
	//  
	//  
	//  
	//  
	//  
	//  
	//   
	//   
	//  
	//  
	//  
	//  
	//  
	//   
	//  
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// Status description and verification error message
	// Note: This field may return `null`, indicating that no valid value was found.
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// This parameter returns the result of imperceptible verification. This parameter is not available for Tencent Cloud International yet.
	// `0`: The request is trusted.
	// `100`: The request is malicious.
	// Note: This field may return `null`, indicating that no valid value was found.
	EvilLevel *int64 `json:"EvilLevel,omitempty" name:"EvilLevel"`

	// The timestamp when the frontend obtains the CAPTCHA.
	// Note: This field may return `null`, indicating that no valid value was found.
	GetCaptchaTime *int64 `json:"GetCaptchaTime,omitempty" name:"GetCaptchaTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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