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

package v20180301

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Address struct {
	// Nationality.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// Post code.
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// Subregion.
	Subdivision *string `json:"Subdivision,omitnil,omitempty" name:"Subdivision"`

	// City.
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Complete address.
	FormattedAddress *string `json:"FormattedAddress,omitnil,omitempty" name:"FormattedAddress"`

	// The first line of address.
	LineOne *string `json:"LineOne,omitnil,omitempty" name:"LineOne"`

	// The second line of address.
	LineTwo *string `json:"LineTwo,omitnil,omitempty" name:"LineTwo"`

	// The third line of address.
	LineThree *string `json:"LineThree,omitnil,omitempty" name:"LineThree"`

	// The fourth line of address.
	LineFour *string `json:"LineFour,omitnil,omitempty" name:"LineFour"`

	// The fifth line of address.
	LineFive *string `json:"LineFive,omitnil,omitempty" name:"LineFive"`
}

// Predefined struct for user
type ApplyCardVerificationRequestParams struct {
	// Please select the country code of ID document.
	// IDN: Indonesia
	// HKG: Hong Kong
	// THA: Thailand
	// MYS: Malaysia
	// SGP: Singapore
	// JPN: Japan
	// AUTO: Automatic Identification
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Please select the type of ID document. The supported options are:
	// ID_CARD
	// PASSPORT
	// DRIVING_LICENSE
	// AUTO
	CardType *string `json:"CardType,omitnil,omitempty" name:"CardType"`

	// Base64 value for the front of the document. Supported image formats: PNG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 5M after Base64 encoding. 
	// The image download takes no more than 3 seconds. Supported image resolution: 8000*8000. One of ImageUrlFront and ImageBase64 Front of the image must be provided. If both are provided, only ImageUrlFront will be used.
	ImageBase64Front *string `json:"ImageBase64Front,omitnil,omitempty" name:"ImageBase64Front"`

	// Base64 value of the reverse side of the document. Supported image formats: PNG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 5M after Base64 encoding. The image download takes no more than 3 seconds. Maximum supported image resolution: 8000*8000. For some certificates, one of ImageUrlBack and ImageBase64Back must be provided. If both are provided, only ImageUrlBack will be used.
	ImageBase64Back *string `json:"ImageBase64Back,omitnil,omitempty" name:"ImageBase64Back"`

	// The URL value on the back of the certificate. Supported image formats: PNG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 5M after Base64 encoding. The image download takes no more than 3 seconds. Maximum supported image resolution: 8000*8000. One of ImageUrlFront and ImageBase64Front of the image must be provided. If both are provided, only ImageUrlFront will be used.
	ImageUrlFront *string `json:"ImageUrlFront,omitnil,omitempty" name:"ImageUrlFront"`

	// The URL value on the back of the certificate. Supported image formats: PNG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 5M after Base64 encoding. The image download takes no more than 3 seconds. Maximum supported image resolution: 8000*8000. For some certificates, one of ImageUrlBack and ImageBase64Back must be provided. If both are provided, only ImageUrlBack will be used.
	ImageUrlBack *string `json:"ImageUrlBack,omitnil,omitempty" name:"ImageUrlBack"`
}

type ApplyCardVerificationRequest struct {
	*tchttp.BaseRequest
	
	// Please select the country code of ID document.
	// IDN: Indonesia
	// HKG: Hong Kong
	// THA: Thailand
	// MYS: Malaysia
	// SGP: Singapore
	// JPN: Japan
	// AUTO: Automatic Identification
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Please select the type of ID document. The supported options are:
	// ID_CARD
	// PASSPORT
	// DRIVING_LICENSE
	// AUTO
	CardType *string `json:"CardType,omitnil,omitempty" name:"CardType"`

	// Base64 value for the front of the document. Supported image formats: PNG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 5M after Base64 encoding. 
	// The image download takes no more than 3 seconds. Supported image resolution: 8000*8000. One of ImageUrlFront and ImageBase64 Front of the image must be provided. If both are provided, only ImageUrlFront will be used.
	ImageBase64Front *string `json:"ImageBase64Front,omitnil,omitempty" name:"ImageBase64Front"`

	// Base64 value of the reverse side of the document. Supported image formats: PNG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 5M after Base64 encoding. The image download takes no more than 3 seconds. Maximum supported image resolution: 8000*8000. For some certificates, one of ImageUrlBack and ImageBase64Back must be provided. If both are provided, only ImageUrlBack will be used.
	ImageBase64Back *string `json:"ImageBase64Back,omitnil,omitempty" name:"ImageBase64Back"`

	// The URL value on the back of the certificate. Supported image formats: PNG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 5M after Base64 encoding. The image download takes no more than 3 seconds. Maximum supported image resolution: 8000*8000. One of ImageUrlFront and ImageBase64Front of the image must be provided. If both are provided, only ImageUrlFront will be used.
	ImageUrlFront *string `json:"ImageUrlFront,omitnil,omitempty" name:"ImageUrlFront"`

	// The URL value on the back of the certificate. Supported image formats: PNG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 5M after Base64 encoding. The image download takes no more than 3 seconds. Maximum supported image resolution: 8000*8000. For some certificates, one of ImageUrlBack and ImageBase64Back must be provided. If both are provided, only ImageUrlBack will be used.
	ImageUrlBack *string `json:"ImageUrlBack,omitnil,omitempty" name:"ImageUrlBack"`
}

func (r *ApplyCardVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCardVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Nationality")
	delete(f, "CardType")
	delete(f, "ImageBase64Front")
	delete(f, "ImageBase64Back")
	delete(f, "ImageUrlFront")
	delete(f, "ImageUrlBack")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyCardVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyCardVerificationResponseParams struct {
	// The token used to identify an verification process. It can be used to get the verification result after the process is completed.
	CardVerificationToken *string `json:"CardVerificationToken,omitnil,omitempty" name:"CardVerificationToken"`

	// The maximum number of polls for calling the pull result interface polling.
	AsyncCardVerificationMaxPollingTimes *uint64 `json:"AsyncCardVerificationMaxPollingTimes,omitnil,omitempty" name:"AsyncCardVerificationMaxPollingTimes"`

	// The interval for polling when calling the pull result interface (in seconds).
	AsyncCardVerificationPollingWaitTime *uint64 `json:"AsyncCardVerificationPollingWaitTime,omitnil,omitempty" name:"AsyncCardVerificationPollingWaitTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyCardVerificationResponse struct {
	*tchttp.BaseResponse
	Response *ApplyCardVerificationResponseParams `json:"Response"`
}

func (r *ApplyCardVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCardVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyLivenessTokenRequestParams struct {
	// Enumerated value. Valid values: `1`, `2`, `3`, and `4`.
	// Their meanings are as follows:
	// 1 - silent
	// 2 - blinking
	// 3 - light
	// 4 - blinking + light (default)
	SecureLevel *string `json:"SecureLevel,omitnil,omitempty" name:"SecureLevel"`
}

type ApplyLivenessTokenRequest struct {
	*tchttp.BaseRequest
	
	// Enumerated value. Valid values: `1`, `2`, `3`, and `4`.
	// Their meanings are as follows:
	// 1 - silent
	// 2 - blinking
	// 3 - light
	// 4 - blinking + light (default)
	SecureLevel *string `json:"SecureLevel,omitnil,omitempty" name:"SecureLevel"`
}

func (r *ApplyLivenessTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyLivenessTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecureLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyLivenessTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyLivenessTokenResponseParams struct {
	// The token used to identify an SDK-based verification process. It is valid for 10 minutes and can be used to get the verification result after the process is completed.
	SdkToken *string `json:"SdkToken,omitnil,omitempty" name:"SdkToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyLivenessTokenResponse struct {
	*tchttp.BaseResponse
	Response *ApplyLivenessTokenResponseParams `json:"Response"`
}

func (r *ApplyLivenessTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyLivenessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplySdkVerificationTokenRequestParams struct {
	// The verification mode. Valid values:
	// 1: OCR + liveness detection + face comparison
	// 2: Liveness detection + face comparison
	// 3: Liveness detection
	// Default value: 1
	CheckMode *int64 `json:"CheckMode,omitnil,omitempty" name:"CheckMode"`

	// The security level of the verification. Valid values:
	// 1: Video-based liveness detection
	// 2: Motion-based liveness detection
	// 3: Reflection-based liveness detection
	// 4: Motion- and reflection-based liveness detection
	// Default value: 4
	SecurityLevel *int64 `json:"SecurityLevel,omitnil,omitempty" name:"SecurityLevel"`

	// The identity document type. Valid values: 
	// 1. `HK` (default): Identity card of Hong Kong (China)
	// 2. `ML`: Malaysian identity card
	// 3. `IndonesiaIDCard`: Indonesian identity card
	// 4. `PhilippinesVoteID`: Philippine voters ID card
	// 5. `PhilippinesDrivingLicense`: Philippine driver's license
	// 6. `PhilippinesTinID`: Philippine TIN ID card
	// 7. `PhilippinesSSSID`: Philippine SSS ID card
	// 8. `PhilippinesUMID`: Philippine UMID card
	// 9. `MLIDPassport`: Passport issued in Hong Kong/Macao/Taiwan (China) or other countries/regions
	// 10..MacaoIDCard: Macao ID Card
	// 11.ThailandIDCard: Thailand ID Card
	// 12.MainlandIDCard: Mainland ID Card
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// The Base64-encoded value of the photo to compare, which is required only when `CheckMode` is set to `2`.
	CompareImage *string `json:"CompareImage,omitnil,omitempty" name:"CompareImage"`

	// Whether ID card authentication is required. If not, only document OCR will be performed. Currently, authentication is available only when the value of `IdCardType` is `HK`.
	//
	// Deprecated: NeedVerifyIdCard is deprecated.
	NeedVerifyIdCard *bool `json:"NeedVerifyIdCard,omitnil,omitempty" name:"NeedVerifyIdCard"`

	// Whether to forbid the modification of the OCR result by users. Default value: `false` (modification allowed). (Currently, this parameter is not applied.)
	DisableChangeOcrResult *bool `json:"DisableChangeOcrResult,omitnil,omitempty" name:"DisableChangeOcrResult"`

	// Whether to disable the OCR warnings. Default value: `false` (not disable), where OCR warnings are enabled and the OCR result will not be returned if there is a warning.
	// This feature applies only to Hong Kong (China) identity cards, Malaysian identity cards, and passports.
	DisableCheckOcrWarnings *bool `json:"DisableCheckOcrWarnings,omitnil,omitempty" name:"DisableCheckOcrWarnings"`

	// A passthrough field, which is returned together with the verification result and can contain up to 1,024 bits.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`
}

type ApplySdkVerificationTokenRequest struct {
	*tchttp.BaseRequest
	
	// The verification mode. Valid values:
	// 1: OCR + liveness detection + face comparison
	// 2: Liveness detection + face comparison
	// 3: Liveness detection
	// Default value: 1
	CheckMode *int64 `json:"CheckMode,omitnil,omitempty" name:"CheckMode"`

	// The security level of the verification. Valid values:
	// 1: Video-based liveness detection
	// 2: Motion-based liveness detection
	// 3: Reflection-based liveness detection
	// 4: Motion- and reflection-based liveness detection
	// Default value: 4
	SecurityLevel *int64 `json:"SecurityLevel,omitnil,omitempty" name:"SecurityLevel"`

	// The identity document type. Valid values: 
	// 1. `HK` (default): Identity card of Hong Kong (China)
	// 2. `ML`: Malaysian identity card
	// 3. `IndonesiaIDCard`: Indonesian identity card
	// 4. `PhilippinesVoteID`: Philippine voters ID card
	// 5. `PhilippinesDrivingLicense`: Philippine driver's license
	// 6. `PhilippinesTinID`: Philippine TIN ID card
	// 7. `PhilippinesSSSID`: Philippine SSS ID card
	// 8. `PhilippinesUMID`: Philippine UMID card
	// 9. `MLIDPassport`: Passport issued in Hong Kong/Macao/Taiwan (China) or other countries/regions
	// 10..MacaoIDCard: Macao ID Card
	// 11.ThailandIDCard: Thailand ID Card
	// 12.MainlandIDCard: Mainland ID Card
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// The Base64-encoded value of the photo to compare, which is required only when `CheckMode` is set to `2`.
	CompareImage *string `json:"CompareImage,omitnil,omitempty" name:"CompareImage"`

	// Whether ID card authentication is required. If not, only document OCR will be performed. Currently, authentication is available only when the value of `IdCardType` is `HK`.
	NeedVerifyIdCard *bool `json:"NeedVerifyIdCard,omitnil,omitempty" name:"NeedVerifyIdCard"`

	// Whether to forbid the modification of the OCR result by users. Default value: `false` (modification allowed). (Currently, this parameter is not applied.)
	DisableChangeOcrResult *bool `json:"DisableChangeOcrResult,omitnil,omitempty" name:"DisableChangeOcrResult"`

	// Whether to disable the OCR warnings. Default value: `false` (not disable), where OCR warnings are enabled and the OCR result will not be returned if there is a warning.
	// This feature applies only to Hong Kong (China) identity cards, Malaysian identity cards, and passports.
	DisableCheckOcrWarnings *bool `json:"DisableCheckOcrWarnings,omitnil,omitempty" name:"DisableCheckOcrWarnings"`

	// A passthrough field, which is returned together with the verification result and can contain up to 1,024 bits.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`
}

func (r *ApplySdkVerificationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplySdkVerificationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CheckMode")
	delete(f, "SecurityLevel")
	delete(f, "IdCardType")
	delete(f, "CompareImage")
	delete(f, "NeedVerifyIdCard")
	delete(f, "DisableChangeOcrResult")
	delete(f, "DisableCheckOcrWarnings")
	delete(f, "Extra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplySdkVerificationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplySdkVerificationTokenResponseParams struct {
	// The token used to identify an SDK-based verification process. It is valid for 7,200s and can be used to get the verification result after the process is completed.
	SdkToken *string `json:"SdkToken,omitnil,omitempty" name:"SdkToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplySdkVerificationTokenResponse struct {
	*tchttp.BaseResponse
	Response *ApplySdkVerificationTokenResponseParams `json:"Response"`
}

func (r *ApplySdkVerificationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplySdkVerificationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyWebVerificationBizTokenIntlRequestParams struct {
	// The web callback URL to redirect to after the verification is completed, including the protocol, hostname, and path. 
	// After the verification process is completed, the BizToken of this process will be spliced to the callback URL in the format of https://www.tencentcloud.com/products/faceid?token={BizToken} before redirect.
	// Example: https://www.tencentcloud.com/products/faceid.
	RedirectURL *string `json:"RedirectURL,omitnil,omitempty" name:"RedirectURL"`

	// The Base64-encoded string (max 8 MB in size) of the photo to be compared.The Data URI scheme header needs to be removed from the encoded string
	// Example: xhBQAAACBjSFJNAAB6****AAAASUVORK5CYII=
	CompareImageBase64 *string `json:"CompareImageBase64,omitnil,omitempty" name:"CompareImageBase64"`

	// The passthrough parameter of the business, max 1,000 characters, which will be returned in GetWebVerificationResultIntl.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// The parameter control the page configuration.
	// Example: {"AutoSkip": true,"CheckMode": 1,"IdCardType": "HKIDCard"}
	Config *WebVerificationConfigIntl `json:"Config,omitnil,omitempty" name:"Config"`
}

type ApplyWebVerificationBizTokenIntlRequest struct {
	*tchttp.BaseRequest
	
	// The web callback URL to redirect to after the verification is completed, including the protocol, hostname, and path. 
	// After the verification process is completed, the BizToken of this process will be spliced to the callback URL in the format of https://www.tencentcloud.com/products/faceid?token={BizToken} before redirect.
	// Example: https://www.tencentcloud.com/products/faceid.
	RedirectURL *string `json:"RedirectURL,omitnil,omitempty" name:"RedirectURL"`

	// The Base64-encoded string (max 8 MB in size) of the photo to be compared.The Data URI scheme header needs to be removed from the encoded string
	// Example: xhBQAAACBjSFJNAAB6****AAAASUVORK5CYII=
	CompareImageBase64 *string `json:"CompareImageBase64,omitnil,omitempty" name:"CompareImageBase64"`

	// The passthrough parameter of the business, max 1,000 characters, which will be returned in GetWebVerificationResultIntl.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// The parameter control the page configuration.
	// Example: {"AutoSkip": true,"CheckMode": 1,"IdCardType": "HKIDCard"}
	Config *WebVerificationConfigIntl `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *ApplyWebVerificationBizTokenIntlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyWebVerificationBizTokenIntlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RedirectURL")
	delete(f, "CompareImageBase64")
	delete(f, "Extra")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyWebVerificationBizTokenIntlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyWebVerificationBizTokenIntlResponseParams struct {
	// The token identifying this web-based verification process, valid for 7,200s after issuance. It is required for getting the result after the verification process is completed.
	// Example: https://intl.faceid.qq.com/reflect/?token=81EEF678-28EE-4759-A82E-6CBBBE6BC442
	//
	// Deprecated: VerificationUrl is deprecated.
	VerificationUrl *string `json:"VerificationUrl,omitnil,omitempty" name:"VerificationUrl"`

	// The token for the web-based verification, which is generated using the ApplyWebVerificationBizTokenIntl API.
	// Example: 81EEF678-28EE-4759-A82E-6CBBBE6BC442
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`

	// The verification URL to be opened with a browser to start the verification process.
	// Example: https://intl.faceid.qq.com/reflect/?token=81EEF678-28EE-4759-A82E-6CBBBE6BC442
	VerificationURL *string `json:"VerificationURL,omitnil,omitempty" name:"VerificationURL"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyWebVerificationBizTokenIntlResponse struct {
	*tchttp.BaseResponse
	Response *ApplyWebVerificationBizTokenIntlResponseParams `json:"Response"`
}

func (r *ApplyWebVerificationBizTokenIntlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyWebVerificationBizTokenIntlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyWebVerificationTokenRequestParams struct {
	// The web redirect URL after the verification is completed.
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`

	// The COS URL of the image for face comparison, which can be obtained with one of the following methods:
	// 1. Call the `CreateUploadUrl` API to generate a URL and call it again after the image is successfully uploaded.
	// 2. Use an existing COS URL. For a private bucket, grant the download permission with a pre-signed URL. The corresponding COS bucket must be in the same region as the input parameter `Region`.
	CompareImageUrl *string `json:"CompareImageUrl,omitnil,omitempty" name:"CompareImageUrl"`

	// The MD5 hash values of the image for face comparison (CompareImageUrl).
	CompareImageMd5 *string `json:"CompareImageMd5,omitnil,omitempty" name:"CompareImageMd5"`
}

type ApplyWebVerificationTokenRequest struct {
	*tchttp.BaseRequest
	
	// The web redirect URL after the verification is completed.
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`

	// The COS URL of the image for face comparison, which can be obtained with one of the following methods:
	// 1. Call the `CreateUploadUrl` API to generate a URL and call it again after the image is successfully uploaded.
	// 2. Use an existing COS URL. For a private bucket, grant the download permission with a pre-signed URL. The corresponding COS bucket must be in the same region as the input parameter `Region`.
	CompareImageUrl *string `json:"CompareImageUrl,omitnil,omitempty" name:"CompareImageUrl"`

	// The MD5 hash values of the image for face comparison (CompareImageUrl).
	CompareImageMd5 *string `json:"CompareImageMd5,omitnil,omitempty" name:"CompareImageMd5"`
}

func (r *ApplyWebVerificationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyWebVerificationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RedirectUrl")
	delete(f, "CompareImageUrl")
	delete(f, "CompareImageMd5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyWebVerificationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyWebVerificationTokenResponseParams struct {
	// The verification URL to be opened with a browser to start the verification process.
	VerificationUrl *string `json:"VerificationUrl,omitnil,omitempty" name:"VerificationUrl"`

	// The token used to identify a web-based verification process. It is valid for 7,200s and can be used to get the verification result after the process is completed.
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyWebVerificationTokenResponse struct {
	*tchttp.BaseResponse
	Response *ApplyWebVerificationTokenResponseParams `json:"Response"`
}

func (r *ApplyWebVerificationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyWebVerificationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CardInfo struct {
	// Hong Kong ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	HKIDCard *HKIDCard `json:"HKIDCard,omitnil,omitempty" name:"HKIDCard"`

	// Malaysia ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	MLIDCard *MLIDCard `json:"MLIDCard,omitnil,omitempty" name:"MLIDCard"`

	// Philippines VoteID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesVoteID *PhilippinesVoteID `json:"PhilippinesVoteID,omitnil,omitempty" name:"PhilippinesVoteID"`

	// Indonesia ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndonesiaIDCard *IndonesiaIDCard `json:"IndonesiaIDCard,omitnil,omitempty" name:"IndonesiaIDCard"`

	// Philippines Driving License
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesDrivingLicense *PhilippinesDrivingLicense `json:"PhilippinesDrivingLicense,omitnil,omitempty" name:"PhilippinesDrivingLicense"`

	// Philippines TinID
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesTinID *PhilippinesTinID `json:"PhilippinesTinID,omitnil,omitempty" name:"PhilippinesTinID"`

	// Philippines SSSID
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesSSSID *PhilippinesSSSID `json:"PhilippinesSSSID,omitnil,omitempty" name:"PhilippinesSSSID"`

	// Philippines UMID
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesUMID *PhilippinesUMID `json:"PhilippinesUMID,omitnil,omitempty" name:"PhilippinesUMID"`

	// ID Cards of Hong Kong, Macao and Taiwan (China), and International Passport
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternationalIDPassport *InternationalIDPassport `json:"InternationalIDPassport,omitnil,omitempty" name:"InternationalIDPassport"`

	// General license information
	// Note: This field may return null, indicating that no valid values can be obtained.
	GeneralCard *GeneralCard `json:"GeneralCard,omitnil,omitempty" name:"GeneralCard"`

	// Indonesia Driving License
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndonesiaDrivingLicense *IndonesiaDrivingLicense `json:"IndonesiaDrivingLicense,omitnil,omitempty" name:"IndonesiaDrivingLicense"`

	// Thailand ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	ThailandIDCard *ThailandIDCard `json:"ThailandIDCard,omitnil,omitempty" name:"ThailandIDCard"`

	// Singapore ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	SingaporeIDCard *SingaporeIDCard `json:"SingaporeIDCard,omitnil,omitempty" name:"SingaporeIDCard"`

	// Macao ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	MacaoIDCard *MacaoIDCard `json:"MacaoIDCard,omitnil,omitempty" name:"MacaoIDCard"`
}

type CardVerifyResult struct {
	// Whether the authentication or OCR process is successful.
	IsPass *bool `json:"IsPass,omitnil,omitempty" name:"IsPass"`

	// Whether the user modified the card recognition result
	IsEdit *bool `json:"IsEdit,omitnil,omitempty" name:"IsEdit"`

	// The download URL of the video used for identity document verification, which is valid for 10 minutes. This parameter is returned only if video-based identity document verification is enabled.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CardVideo *FileInfo `json:"CardVideo,omitnil,omitempty" name:"CardVideo"`

	// The download URL of the identity document image, which is valid for 10 minutes.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CardImage *FileInfo `json:"CardImage,omitnil,omitempty" name:"CardImage"`

	// The OCR result (in JSON) of the identity document image. If verification or OCR fails, this parameter is left empty. The URL is valid for 10 minutes.
	// (1) Hong Kong (China) identity card
	// When the value of `IdCardType` is `HK`:
	// - CnName (string): Name in Chinese.
	// - EnName (string): Name in English.
	// - TelexCode (string): The code corresponding to the name in Chinese.
	// - Sex (string): Gender. Valid values: `M` (male) and `F` (female).
	// - Birthday (string): Date of birth.
	// - Permanent (int): Whether it is a permanent residence identity card. Valid values: `0` (non-permanent), `1` (permanent), and `-1` (unknown).
	// - IdNum (string): Identity card number.
	// - Symbol (string): The ID symbol below the date of birth, such as "***AZ".
	// - FirstIssueDate (string): Month and year of first registration.
	// - CurrentIssueDate (string): The date of latest issuance.
	// 
	// (2) Malaysian identity card
	// When the value of `IdCardType` is `ML`:
	// - Sex (string): Gender. Valid values: `LELAKI` (male) and `PEREMPUAN` (female).
	// - Birthday (string): Date of birth.
	// - ID (string): Identity card number.
	// - Name (string): Name.
	// - Address (string): Address.
	// - Type (string): Identity document type.
	// 
	// (3) Philippine identity document
	// When the value of `IdCardType` is `PhilippinesVoteID`:
	// - Birthday (string): Date of birth.
	// - Address (string): Address.
	// - LastName (string): Last name.
	// - FirstName (string): First name.
	// - VIN (string): Voter's identification number (VIN).
	// - CivilStatus (string): Civil status.
	// - Citizenship (string): Citizenship.
	// - PrecinctNo (string): Precinct.
	// 
	// When the value of `IdCardType` is `PhilippinesDrivingLicense`:
	// - Sex (string): Gender.
	// - Birthday (string): Date of birth.
	// - Name (string): Name.
	// - Address (string): Address.
	// - LastName (string): Last name.
	// - FirstName (string): First name.
	// - MiddleName (string): Middle name.
	// - Nationality (string): Nationality.
	// - LicenseNo (string): License number.
	// - ExpiresDate (string): Expiration date.
	// - AgencyCode (string): Agency code.
	// 
	// When the value of `IdCardType` is `PhilippinesTinID`:
	// - LicenseNumber (string): Tax identification number (TIN).
	// - FullName (string): Full name.
	// - Address (string): Address.
	// - Birthday (string): Date of birth.
	// - IssueDate (string): Issue date.
	// 
	// When the value of `IdCardType` is `PhilippinesSSSID`:
	// - LicenseNumber (string): Common reference number (CRN).
	// - FullName (string): Full name.
	// - Birthday (string): Date of birth.
	// 
	// When the value of `IdCardType` is `PhilippinesUMID`:
	// - Surname (string): Surname.
	// - MiddleName (string):Middle name.
	// - GivenName (string): Given name.
	// - Sex (string): Gender.
	// - Birthday (string): Date of birth.
	// - Address (string): Address.
	// - CRN (string): Common reference number (CRN).
	// 
	// (4) Indonesian identity card
	// When the value of `IdCardType` is `IndonesiaIDCard`:
	// - NIK (string): Single Identity Number.
	// - Nama (string): Full name.
	// - TempatTglLahir (string): Place and date of birth.
	// - JenisKelamin (string): Gender.
	// - GolDarah (string): Blood type.
	// - Alamat (string): Address.
	// - RTRW (string): Street.
	// - KelDesa (string): Village.
	// - Kecamatan (string): Region.
	// - Agama (string): Religion.
	// - StatusPerkawinan (string): Marital status.
	// - Perkerjaan (string): Occupation.
	// - KewargaNegaraan (string): Nationality.
	// - BerlakuHingga (string): Expiry date.
	// - IssuedDate (string): Issue date.
	// 
	// (5) A passport issued in Hong Kong/Macao/Taiwan (China) or other countries/regions
	// When the value of `IdCardType` is `MLIDPassport`:
	// - FullName (string): Full name.
	// - Surname (string): Surname.
	// - GivenName (string): Given name.
	// - Birthday (string): Date of birth.
	// - Sex (string): Gender. Valid values: `F` (female) and `M` (male).
	// - DateOfExpiration (string): Expiration date.
	// - IssuingCountry (string): Issuing country.
	// - NationalityCode (string): Country/region code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CardInfoOcrJson *FileInfo `json:"CardInfoOcrJson,omitnil,omitempty" name:"CardInfoOcrJson"`

	// The request ID of a single process.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// The recognition results of ID card
	//
	// Deprecated: CardInfo is deprecated.
	CardInfo *CardInfo `json:"CardInfo,omitnil,omitempty" name:"CardInfo"`

	// License OCR result
	NormalCardInfo *NormalCardInfo `json:"NormalCardInfo,omitnil,omitempty" name:"NormalCardInfo"`

	// Card warning information
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`
}

// Predefined struct for user
type CompareFaceLivenessRequestParams struct {
	// Base64 value of photos used for face comparison. 
	// The size of image data encoded by Base64 shall not exceed 3M, only jpg and png are supported. 
	// Please use standard Base64 encoding (use = for padding). Refer to RFC4648 for encoding specifications. 
	// Example values: "/9j/4AAQSk... (total length:61944)KiiK//2Q=="
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Base64 value of photos used for face comparison. 
	// The size of image data encoded by Base64 shall not exceed 3M, only jpg and png are supported. 
	// Please use standard Base64 encoding (use = for padding). Refer to RFC4648 for encoding specifications. 
	// Example values: "/9j/4AAQSk... (total length:61944)KiiK//2Q=="
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// The liveness detection type. Valid values: `LIP`, `ACTION`, and `SILENT`.
	// `LIP`: Numeric mode; `ACTION`: Motion mode; `SILENT`: silent mode. Select one of them.
	// Example value: "SILENT"
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// When the “LivenessType” parameter is “ACTION”, it must be specified.
	// It is used to control the action sequence. Action types: 
	// 1 (open mouth)
	// 2 (blink)
	// 3 (nod)
	// 4 (shake head). 
	// Select one or two from the four actions.
	// Example of passing single action parameter: "1".
	// Example of passing multiple action parameters: "4,2".
	// When the “LivenessType” parameter value is “SILENT”, it shall be unspecified.
	// Example value: ""
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`
}

type CompareFaceLivenessRequest struct {
	*tchttp.BaseRequest
	
	// Base64 value of photos used for face comparison. 
	// The size of image data encoded by Base64 shall not exceed 3M, only jpg and png are supported. 
	// Please use standard Base64 encoding (use = for padding). Refer to RFC4648 for encoding specifications. 
	// Example values: "/9j/4AAQSk... (total length:61944)KiiK//2Q=="
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Base64 value of photos used for face comparison. 
	// The size of image data encoded by Base64 shall not exceed 3M, only jpg and png are supported. 
	// Please use standard Base64 encoding (use = for padding). Refer to RFC4648 for encoding specifications. 
	// Example values: "/9j/4AAQSk... (total length:61944)KiiK//2Q=="
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// The liveness detection type. Valid values: `LIP`, `ACTION`, and `SILENT`.
	// `LIP`: Numeric mode; `ACTION`: Motion mode; `SILENT`: silent mode. Select one of them.
	// Example value: "SILENT"
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// When the “LivenessType” parameter is “ACTION”, it must be specified.
	// It is used to control the action sequence. Action types: 
	// 1 (open mouth)
	// 2 (blink)
	// 3 (nod)
	// 4 (shake head). 
	// Select one or two from the four actions.
	// Example of passing single action parameter: "1".
	// Example of passing multiple action parameters: "4,2".
	// When the “LivenessType” parameter value is “SILENT”, it shall be unspecified.
	// Example value: ""
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`
}

func (r *CompareFaceLivenessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareFaceLivenessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "VideoBase64")
	delete(f, "LivenessType")
	delete(f, "ValidateData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompareFaceLivenessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompareFaceLivenessResponseParams struct {
	// Service error code. When the return value is "Success", it indicates that the liveness detection and face comparison succeeded. It is determined that they are the same person. When the return value is "FailedOperation.CompareLowSimilarity", it indicates that the liveness detection succeeded, and the face comparison similarity is lower than 70 points. It is determined that they are not the same person. For other error cases, please refer to Liveness Face Comparison (Pure API) Error Code (https://www.tencentcloud.com/document/product/1061/55390). 
	// Example Value: "Success".
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// Description of business results. 
	// Example value: "Success"
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// This value is valid when the “Result” parameter is "Success" or "FailedOperation.CompareLowSimilarity." 
	// This value indicates the similarity of face comparison. Value range: [0.00, 100.00]. The false pass rate for threshold 70 is 1 in 1,000, and the false pass rate for threshold 80 is 1 in 1,000. 
	// Example value: 80.00
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// The optimal screenshot of the video after verification is the value encoded by BASE64, jpg format. 
	// Note: This field may return “null”, indicating that no valid value can be obtained. 
	// Example values: "/9j/4AAQSk... (total length:142036)s97n//2Q=="
	BestFrameBase64 *string `json:"BestFrameBase64,omitnil,omitempty" name:"BestFrameBase64"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CompareFaceLivenessResponse struct {
	*tchttp.BaseResponse
	Response *CompareFaceLivenessResponseParams `json:"Response"`
}

func (r *CompareFaceLivenessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareFaceLivenessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompareResult struct {
	// The final verification result code.
	// 0: Success.
	// 1001: Failed to call the liveness detection engine.
	// 1004: Face detection failed.
	// 2004: The uploaded face image is too large or too small.
	// 2012: The face is not fully exposed.
	// 2013: No face is detected.
	// 2014: The resolution of the uploaded image is too low . Please upload a new one.
	// 2015: Face comparison failed.
	// 2016: The similarity did not reach the passing standard.
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// The description of the final verification result.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The liveness algorithm package generated during this SDK-based liveness detection.
	LiveData *FileInfo `json:"LiveData,omitnil,omitempty" name:"LiveData"`

	// The download URL of the video used for verification, which is valid for 10 minutes.
	LiveVideo *FileInfo `json:"LiveVideo,omitnil,omitempty" name:"LiveVideo"`

	// The liveness detection result code.
	// 0: Success.
	// 1001: Failed to call the liveness detection engine.
	// 1004: Face detection failed.
	LiveErrorCode *string `json:"LiveErrorCode,omitnil,omitempty" name:"LiveErrorCode"`

	// The description of the liveness detection result.
	LiveErrorMsg *string `json:"LiveErrorMsg,omitnil,omitempty" name:"LiveErrorMsg"`

	// The download URL of the face screenshot during verification, which is valid for 10 minutes.
	// Note: This field may return null, indicating that no valid value can be obtained.
	BestFrame *FileInfo `json:"BestFrame,omitnil,omitempty" name:"BestFrame"`

	// The download URL of the profile photo screenshot from the identity document, which is valid for 10 minutes.
	ProfileImage *FileInfo `json:"ProfileImage,omitnil,omitempty" name:"ProfileImage"`

	// The face comparison result code.
	// 0: Success.
	// 2004: The uploaded face image is too large or too small.
	// 2012: The face is not fully exposed.
	// 2013: No face is detected.
	// 2014: The resolution of the uploaded image is too low . Please upload a new one.
	// 2015: Face comparison failed.
	// 2016: The similarity did not reach the passing standard.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CompareErrorCode *string `json:"CompareErrorCode,omitnil,omitempty" name:"CompareErrorCode"`

	// The description of the face comparison result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompareErrorMsg *string `json:"CompareErrorMsg,omitnil,omitempty" name:"CompareErrorMsg"`

	// The similarity score of face comparison.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// This parameter is disused.
	IsNeedCharge *bool `json:"IsNeedCharge,omitnil,omitempty" name:"IsNeedCharge"`

	// The identity document photo info edited by the user. Currently, this parameter is not applied.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CardInfoInputJson *FileInfo `json:"CardInfoInputJson,omitnil,omitempty" name:"CardInfoInputJson"`

	// The request ID of this verification process.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

// Predefined struct for user
type CreateUploadUrlRequestParams struct {
	// Target API
	TargetAction *string `json:"TargetAction,omitnil,omitempty" name:"TargetAction"`
}

type CreateUploadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Target API
	TargetAction *string `json:"TargetAction,omitnil,omitempty" name:"TargetAction"`
}

func (r *CreateUploadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUploadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUploadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUploadUrlResponseParams struct {
	// The URL for uploading contents with the `HTTP PUT` method.
	UploadUrl *string `json:"UploadUrl,omitnil,omitempty" name:"UploadUrl"`

	// The resource URL obtained after this upload is completed and to be passed in where it is required later.
	ResourceUrl *string `json:"ResourceUrl,omitnil,omitempty" name:"ResourceUrl"`

	// The point in time when the upload/download link expires, which is a 10-bit Unix timestamp.
	ExpiredTimestamp *int64 `json:"ExpiredTimestamp,omitnil,omitempty" name:"ExpiredTimestamp"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUploadUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateUploadUrlResponseParams `json:"Response"`
}

func (r *CreateUploadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUploadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectReflectLivenessAndCompareRequestParams struct {
	// URL of the liveness detection data package generated by the SDK
	LiveDataUrl *string `json:"LiveDataUrl,omitnil,omitempty" name:"LiveDataUrl"`

	// MD5 hash value (32-bit) of the liveness detection data package generated by the SDK, which is used to verify the LiveData consistency.
	LiveDataMd5 *string `json:"LiveDataMd5,omitnil,omitempty" name:"LiveDataMd5"`

	// URL of the target image for comparison
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// MD5 hash value (32-bit) of the target image for comparison, which is used to verify the `Image` consistency.
	ImageMd5 *string `json:"ImageMd5,omitnil,omitempty" name:"ImageMd5"`
}

type DetectReflectLivenessAndCompareRequest struct {
	*tchttp.BaseRequest
	
	// URL of the liveness detection data package generated by the SDK
	LiveDataUrl *string `json:"LiveDataUrl,omitnil,omitempty" name:"LiveDataUrl"`

	// MD5 hash value (32-bit) of the liveness detection data package generated by the SDK, which is used to verify the LiveData consistency.
	LiveDataMd5 *string `json:"LiveDataMd5,omitnil,omitempty" name:"LiveDataMd5"`

	// URL of the target image for comparison
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// MD5 hash value (32-bit) of the target image for comparison, which is used to verify the `Image` consistency.
	ImageMd5 *string `json:"ImageMd5,omitnil,omitempty" name:"ImageMd5"`
}

func (r *DetectReflectLivenessAndCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectReflectLivenessAndCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LiveDataUrl")
	delete(f, "LiveDataMd5")
	delete(f, "ImageUrl")
	delete(f, "ImageMd5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectReflectLivenessAndCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectReflectLivenessAndCompareResponseParams struct {
	// Temporary URL of the best screenshot (.jpg) of the video after successful verification. Both the screenshot and the URL are valid for two hours only, so you need to download the screenshot within this period.
	BestFrameUrl *string `json:"BestFrameUrl,omitnil,omitempty" name:"BestFrameUrl"`

	// MD5 hash value (32-bit) of the best screenshot of the video after successful verification, which is used to verify the `BestFrame` consistency.
	BestFrameMd5 *string `json:"BestFrameMd5,omitnil,omitempty" name:"BestFrameMd5"`

	// Service error code. `Success` will be returned for success. For error information, see the `FailedOperation` section in the error code list below.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// Service result description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Similarity. Value range: [0.00, 100.00]. As a recommendation, when the similarity is greater than or equal to 70, it can be determined that the two faces are of the same person. You can adjust the threshold according to your specific scenario (the FAR at the threshold of 70 is 0.1%, and FAR at the threshold of 80 is 0.01%).
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetectReflectLivenessAndCompareResponse struct {
	*tchttp.BaseResponse
	Response *DetectReflectLivenessAndCompareResponseParams `json:"Response"`
}

func (r *DetectReflectLivenessAndCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectReflectLivenessAndCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileInfo struct {
	// The URL for downloading the file
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// The 32-bit MD5 checksum of the file
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// The file size
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type GeneralCard struct {
	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Personal number, which is returned when it is a passport
	// Note: This field may return null, indicating that no valid values can be obtained.
	PersonalNumber *string `json:"PersonalNumber,omitnil,omitempty" name:"PersonalNumber"`

	// The first line of passport machine reading code
	// Note: This field may return null, indicating that no valid values can be obtained.
	PassportCodeFirst *string `json:"PassportCodeFirst,omitnil,omitempty" name:"PassportCodeFirst"`

	// The first line of passport machine reading code
	// Note: This field may return null, indicating that no valid values can be obtained.
	PassportCodeSecond *string `json:"PassportCodeSecond,omitnil,omitempty" name:"PassportCodeSecond"`

	// Date of expiry in the format of YYYY-MM-DD
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

	// Valid date in the format of YYYY-MM-DD
	// Note: This field may return null, indicating that no valid values can be obtained.
	DueDate *string `json:"DueDate,omitnil,omitempty" name:"DueDate"`

	// Date of issue in the format of YYYY-MM-DD
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedDate *string `json:"IssuedDate,omitnil,omitempty" name:"IssuedDate"`

	// Issuing authority
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedAuthority *string `json:"IssuedAuthority,omitnil,omitempty" name:"IssuedAuthority"`

	// Issuing country, which is returned following the ISO 3166 country coding specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example: MYS
	IssuedCountry *string `json:"IssuedCountry,omitnil,omitempty" name:"IssuedCountry"`

	// Full Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// First name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// Last name
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// Gender on the license
	// - M: male
	// - F: female
	// - X: other gender
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example: M
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Age. 0 indicates that no valid information is obtained.
	// Example: 0
	Age *string `json:"Age,omitnil,omitempty" name:"Age"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Birth place
	// Note: This field may return null, indicating that no valid values can be obtained.
	BirthPlace *string `json:"BirthPlace,omitnil,omitempty" name:"BirthPlace"`

	// Nationality, which is returned following the ISO 3166 country coding specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example: IND
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Registration number
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegistrationNumber *string `json:"RegistrationNumber,omitnil,omitempty" name:"RegistrationNumber"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *Address `json:"Address,omitnil,omitempty" name:"Address"`
}

// Predefined struct for user
type GenerateReflectSequenceRequestParams struct {
	// The resource URL of the data package generated by the SDK.
	DeviceDataUrl *string `json:"DeviceDataUrl,omitnil,omitempty" name:"DeviceDataUrl"`

	// The MD5 hash value of the data package generated by the SDK.
	DeviceDataMd5 *string `json:"DeviceDataMd5,omitnil,omitempty" name:"DeviceDataMd5"`

	// 1 - silent
	// 2 - blinking
	// 3 - light
	// 4 - blinking + light (default)
	SecurityLevel *string `json:"SecurityLevel,omitnil,omitempty" name:"SecurityLevel"`
}

type GenerateReflectSequenceRequest struct {
	*tchttp.BaseRequest
	
	// The resource URL of the data package generated by the SDK.
	DeviceDataUrl *string `json:"DeviceDataUrl,omitnil,omitempty" name:"DeviceDataUrl"`

	// The MD5 hash value of the data package generated by the SDK.
	DeviceDataMd5 *string `json:"DeviceDataMd5,omitnil,omitempty" name:"DeviceDataMd5"`

	// 1 - silent
	// 2 - blinking
	// 3 - light
	// 4 - blinking + light (default)
	SecurityLevel *string `json:"SecurityLevel,omitnil,omitempty" name:"SecurityLevel"`
}

func (r *GenerateReflectSequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateReflectSequenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceDataUrl")
	delete(f, "DeviceDataMd5")
	delete(f, "SecurityLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateReflectSequenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateReflectSequenceResponseParams struct {
	// The resource URL of the light sequence, which needs to be downloaded and passed through to the SDK to start the identity verification process.
	ReflectSequenceUrl *string `json:"ReflectSequenceUrl,omitnil,omitempty" name:"ReflectSequenceUrl"`

	// The MD5 hash value of the light sequence, which is used to check whether the light sequence is altered.
	ReflectSequenceMd5 *string `json:"ReflectSequenceMd5,omitnil,omitempty" name:"ReflectSequenceMd5"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateReflectSequenceResponse struct {
	*tchttp.BaseResponse
	Response *GenerateReflectSequenceResponseParams `json:"Response"`
}

func (r *GenerateReflectSequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateReflectSequenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCardVerificationResultRequestParams struct {
	// The token used to identify an verification process. It can be used to get the verification result after the process is completed.
	CardVerificationToken *string `json:"CardVerificationToken,omitnil,omitempty" name:"CardVerificationToken"`
}

type GetCardVerificationResultRequest struct {
	*tchttp.BaseRequest
	
	// The token used to identify an verification process. It can be used to get the verification result after the process is completed.
	CardVerificationToken *string `json:"CardVerificationToken,omitnil,omitempty" name:"CardVerificationToken"`
}

func (r *GetCardVerificationResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCardVerificationResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CardVerificationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCardVerificationResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCardVerificationResultResponseParams struct {
	// Pass status. When Warning and Rejected are returned, please check the specific reasons in the WarnInfo structure return. Example values are as follows:
	// Passed
	// Warning
	// Rejected
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Warning information returned by document verification.
	WarnInfo []*string `json:"WarnInfo,omitnil,omitempty" name:"WarnInfo"`

	// Nationality code. Complies with standard ISO 3166-1 alpha-3. 
	// 
	// Example value: IDN
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Card Type. The supported options are:
	// ID_CARD
	// PASSPORT
	// DRIVING_LICENSE
	// AUTO
	// 
	// Example value: ID_CARD
	CardType *string `json:"CardType,omitnil,omitempty" name:"CardType"`

	// Subtype of the ID document.
	CardSubType *string `json:"CardSubType,omitnil,omitempty" name:"CardSubType"`

	// Recognition results of the ID document.
	CardInfo *CardInfo `json:"CardInfo,omitnil,omitempty" name:"CardInfo"`

	// The token used to identify an verification process. It can be used to get the verification result after the process is completed.
	IDVerificationToken *string `json:"IDVerificationToken,omitnil,omitempty" name:"IDVerificationToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCardVerificationResultResponse struct {
	*tchttp.BaseResponse
	Response *GetCardVerificationResultResponseParams `json:"Response"`
}

func (r *GetCardVerificationResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCardVerificationResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdResultIntlRequestParams struct {
	// The ID of the SDK-based liveness detection and face comparison process, which is generated when the `GetFaceIdTokenIntl` API is called.	
	SdkToken *string `json:"SdkToken,omitnil,omitempty" name:"SdkToken"`
}

type GetFaceIdResultIntlRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the SDK-based liveness detection and face comparison process, which is generated when the `GetFaceIdTokenIntl` API is called.	
	SdkToken *string `json:"SdkToken,omitnil,omitempty" name:"SdkToken"`
}

func (r *GetFaceIdResultIntlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdResultIntlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFaceIdResultIntlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdResultIntlResponseParams struct {
	// The return code of the verification result.
	// 0: Succeeded.
	// 1001: System error.
	// 1004: Liveness detection and face comparison failed.
	// 2004: The image passed in is too large or too small.
	// 2012: Several faces were detected.
	// 2013: No face was detected, or the face detected was incomplete.
	// 2014: The image resolution is too low or the quality does not meet the requirements.
	// 2015: Face comparison failed.
	// 2016: The similarity did not reach the standard passing threshold.
	// -999: The verification process wasn't finished.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The description of the verification result.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The best frame screenshot (in Base64) obtained during the verification.
	BestFrame *string `json:"BestFrame,omitnil,omitempty" name:"BestFrame"`

	// The video file (Base64) for verification.
	Video *string `json:"Video,omitnil,omitempty" name:"Video"`

	// The similarity, with a value range of 0-100. A greater value indicates higher similarity. This parameter is returned only in the `compare` (liveness detection and face comparison) mode.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Similarity *float64 `json:"Similarity,omitnil,omitempty" name:"Similarity"`

	// The pass-through parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFaceIdResultIntlResponse struct {
	*tchttp.BaseResponse
	Response *GetFaceIdResultIntlResponseParams `json:"Response"`
}

func (r *GetFaceIdResultIntlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdResultIntlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdTokenIntlRequestParams struct {
	// The detection mode. Valid values:
	// `liveness`: Liveness detection only.
	// `compare`: Liveness detection and face comparison.
	// Default value: `liveness`.
	CheckMode *string `json:"CheckMode,omitnil,omitempty" name:"CheckMode"`

	// The verification security level. Valid values:
	// `1`: Video-based liveness detection.
	// `2`: Motion-based liveness detection.
	// `3`: Reflection-based liveness detection.
	// `4`: Motion- and reflection-based liveness detection.
	// Default value: `4`.
	SecureLevel *string `json:"SecureLevel,omitnil,omitempty" name:"SecureLevel"`

	// The photo (in Base64) to compare. This parameter is required when the value of `CheckMode` is `compare`.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// The pass-through parameter, which can be omitted if there are no special requirements.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`
}

type GetFaceIdTokenIntlRequest struct {
	*tchttp.BaseRequest
	
	// The detection mode. Valid values:
	// `liveness`: Liveness detection only.
	// `compare`: Liveness detection and face comparison.
	// Default value: `liveness`.
	CheckMode *string `json:"CheckMode,omitnil,omitempty" name:"CheckMode"`

	// The verification security level. Valid values:
	// `1`: Video-based liveness detection.
	// `2`: Motion-based liveness detection.
	// `3`: Reflection-based liveness detection.
	// `4`: Motion- and reflection-based liveness detection.
	// Default value: `4`.
	SecureLevel *string `json:"SecureLevel,omitnil,omitempty" name:"SecureLevel"`

	// The photo (in Base64) to compare. This parameter is required when the value of `CheckMode` is `compare`.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// The pass-through parameter, which can be omitted if there are no special requirements.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`
}

func (r *GetFaceIdTokenIntlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdTokenIntlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CheckMode")
	delete(f, "SecureLevel")
	delete(f, "Image")
	delete(f, "Extra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFaceIdTokenIntlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdTokenIntlResponseParams struct {
	// The SDK token, which is used throughout the verification process and to get the verification result.
	SdkToken *string `json:"SdkToken,omitnil,omitempty" name:"SdkToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFaceIdTokenIntlResponse struct {
	*tchttp.BaseResponse
	Response *GetFaceIdTokenIntlResponseParams `json:"Response"`
}

func (r *GetFaceIdTokenIntlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdTokenIntlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLivenessResultRequestParams struct {
	// The token used to identify an SDK-based verification process.
	SdkToken *string `json:"SdkToken,omitnil,omitempty" name:"SdkToken"`
}

type GetLivenessResultRequest struct {
	*tchttp.BaseRequest
	
	// The token used to identify an SDK-based verification process.
	SdkToken *string `json:"SdkToken,omitnil,omitempty" name:"SdkToken"`
}

func (r *GetLivenessResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLivenessResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLivenessResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLivenessResultResponseParams struct {
	// The final verification result.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The description of the final verification result.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The face screenshot.
	BestFrame *FileInfo `json:"BestFrame,omitnil,omitempty" name:"BestFrame"`

	// The video for the detection.
	Video *FileInfo `json:"Video,omitnil,omitempty" name:"Video"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLivenessResultResponse struct {
	*tchttp.BaseResponse
	Response *GetLivenessResultResponseParams `json:"Response"`
}

func (r *GetLivenessResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLivenessResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSdkVerificationResultRequestParams struct {
	// The token used to identify an SDK-based verification process.
	SdkToken *string `json:"SdkToken,omitnil,omitempty" name:"SdkToken"`
}

type GetSdkVerificationResultRequest struct {
	*tchttp.BaseRequest
	
	// The token used to identify an SDK-based verification process.
	SdkToken *string `json:"SdkToken,omitnil,omitempty" name:"SdkToken"`
}

func (r *GetSdkVerificationResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSdkVerificationResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSdkVerificationResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSdkVerificationResultResponseParams struct {
	// The result code of the verification result.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The verification result description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The charge count.
	ChargeCount *int64 `json:"ChargeCount,omitnil,omitempty" name:"ChargeCount"`

	// The results of multiple OCR processes (in order). The result of the final process is used as the valid result.
	CardVerifyResults []*CardVerifyResult `json:"CardVerifyResults,omitnil,omitempty" name:"CardVerifyResults"`

	// The results of multiple liveness detection processes (in order). The result of the final process is used as the valid result.
	CompareResults []*CompareResult `json:"CompareResults,omitnil,omitempty" name:"CompareResults"`

	// Data passed through in the process of getting the token.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSdkVerificationResultResponse struct {
	*tchttp.BaseResponse
	Response *GetSdkVerificationResultResponseParams `json:"Response"`
}

func (r *GetSdkVerificationResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSdkVerificationResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWebVerificationResultIntlRequestParams struct {
	// The token for the web-based verification, which is generated using the `ApplyWebVerificationBizTokenIntl` API.
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`
}

type GetWebVerificationResultIntlRequest struct {
	*tchttp.BaseRequest
	
	// The token for the web-based verification, which is generated using the `ApplyWebVerificationBizTokenIntl` API.
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`
}

func (r *GetWebVerificationResultIntlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWebVerificationResultIntlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWebVerificationResultIntlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWebVerificationResultIntlResponseParams struct {
	// The final result of this verification. `0` indicates that the person is the same as that in the photo.
	// For other error codes, see <a href="https://www.tencentcloud.com/document/product/1061/55390?lang=en&pg=#8a960e1e-39c0-42cb-b181-b3164d77f81e">Liveness Detection and Face Comparison (Mobile HTML5) Error Codes</a>
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// The description of the final verification result.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The detailed verification result list of this process. Retries are allowed, so a verification process may have several entries of results.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	VerificationDetailList []*VerificationDetail `json:"VerificationDetailList,omitnil,omitempty" name:"VerificationDetailList"`

	// The Base64-encoded string of the video collected from the video stream. Retries are allowed, and this field returns only the data collected in the last verification. If no video is collected, null is returned.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// The Base64-encoded string of the best face screenshot u200dcollected from the video stream. Retries are allowed, and this field returns only the data collected in the last verification. If no best face screenshot is collected, null is returned.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	BestFrameBase64 *string `json:"BestFrameBase64,omitnil,omitempty" name:"BestFrameBase64"`

	// Card recognize result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OCRResult []*OCRResult `json:"OCRResult,omitnil,omitempty" name:"OCRResult"`

	// The passthrough parameter of the business, max 1,000 characters, which will be returned in GetWebVerificationResultIntl.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWebVerificationResultIntlResponse struct {
	*tchttp.BaseResponse
	Response *GetWebVerificationResultIntlResponseParams `json:"Response"`
}

func (r *GetWebVerificationResultIntlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWebVerificationResultIntlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWebVerificationResultRequestParams struct {
	// The token for the web-based verification, which is generated with the `ApplyWebVerificationToken` API.
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`
}

type GetWebVerificationResultRequest struct {
	*tchttp.BaseRequest
	
	// The token for the web-based verification, which is generated with the `ApplyWebVerificationToken` API.
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`
}

func (r *GetWebVerificationResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWebVerificationResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWebVerificationResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWebVerificationResultResponseParams struct {
	// The final result of this verification. `0` indicates that the person is the same as that in the photo.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// The description of the final verification result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The temporary URL of the best face screenshot collected from the video stream. It is valid for 10 minutes. Download the image if needed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoBestFrameUrl *string `json:"VideoBestFrameUrl,omitnil,omitempty" name:"VideoBestFrameUrl"`

	// The MD5 hash value of the best face screenshot collected from the video stream. It can be used to check whether the image content is consistent with the file content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoBestFrameMd5 *string `json:"VideoBestFrameMd5,omitnil,omitempty" name:"VideoBestFrameMd5"`

	// The details list of this verification process.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VerificationDetailList []*VerificationDetail `json:"VerificationDetailList,omitnil,omitempty" name:"VerificationDetailList"`

	// The temporary URL of the video collected from the video stream. It is valid for 10 minutes. Download the video if needed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// The MD5 hash value of the video collected from the video stream. It can be used to check whether the video content is consistent with the file content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoMd5 *string `json:"VideoMd5,omitnil,omitempty" name:"VideoMd5"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWebVerificationResultResponse struct {
	*tchttp.BaseResponse
	Response *GetWebVerificationResultResponseParams `json:"Response"`
}

func (r *GetWebVerificationResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWebVerificationResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HKIDCard struct {

	CnName *string `json:"CnName,omitnil,omitempty" name:"CnName"`

	// English name
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example: SAN, Nan
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// Telex code correspondint to the Chinese name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TelexCode *string `json:"TelexCode,omitnil,omitempty" name:"TelexCode"`

	// Gender: "Male-M" or "Female-F"
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example: 01-01-2001
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Permanent resident ID card: 0-non-permanent; 1-permanent; -1-unknown
	// Note: This field may return null, indicating that no valid values can be obtained.
	Permanent *string `json:"Permanent,omitnil,omitempty" name:"Permanent"`

	// ID card number
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example: C000000(E)
	IdNum *string `json:"IdNum,omitnil,omitempty" name:"IdNum"`

	// Lisence symbol, which is the symbol below Birthday. Example: "***AZ"
	// Note: This field may return null, indicating that no valid values can be obtained.
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// The first date of issue
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstIssueDate *string `json:"FirstIssueDate,omitnil,omitempty" name:"FirstIssueDate"`

	// The current date of issue
	// Note: This field may return null, indicating that no valid values can be obtained.
	CurrentIssueDate *string `json:"CurrentIssueDate,omitnil,omitempty" name:"CurrentIssueDate"`
}

type IndonesiaDrivingLicense struct {
	// Last name
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// First name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Expiration date
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

	// Date of issue
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedDate *string `json:"IssuedDate,omitnil,omitempty" name:"IssuedDate"`

	// Issuing country
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedCountry *string `json:"IssuedCountry,omitnil,omitempty" name:"IssuedCountry"`
}

type IndonesiaIDCard struct {
	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	NIK *string `json:"NIK,omitnil,omitempty" name:"NIK"`

	// Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nama *string `json:"Nama,omitnil,omitempty" name:"Nama"`

	// Birth place/Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	TempatTglLahir *string `json:"TempatTglLahir,omitnil,omitempty" name:"TempatTglLahir"`

	// Gender
	// Note: This field may return null, indicating that no valid values can be obtained.
	JenisKelamin *string `json:"JenisKelamin,omitnil,omitempty" name:"JenisKelamin"`

	// Blood type
	// Note: This field may return null, indicating that no valid values can be obtained.
	GolDarah *string `json:"GolDarah,omitnil,omitempty" name:"GolDarah"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Alamat *string `json:"Alamat,omitnil,omitempty" name:"Alamat"`

	// Street
	// Note: This field may return null, indicating that no valid values can be obtained.
	RTRW *string `json:"RTRW,omitnil,omitempty" name:"RTRW"`

	// Village
	// Note: This field may return null, indicating that no valid values can be obtained.
	KelDesa *string `json:"KelDesa,omitnil,omitempty" name:"KelDesa"`

	// Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Kecamatan *string `json:"Kecamatan,omitnil,omitempty" name:"Kecamatan"`

	// Religious beliefs
	// Note: This field may return null, indicating that no valid values can be obtained.
	Agama *string `json:"Agama,omitnil,omitempty" name:"Agama"`

	// Marital status
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusPerkawinan *string `json:"StatusPerkawinan,omitnil,omitempty" name:"StatusPerkawinan"`

	// Job
	// Note: This field may return null, indicating that no valid values can be obtained.
	Perkerjaan *string `json:"Perkerjaan,omitnil,omitempty" name:"Perkerjaan"`

	// Nationality
	// Note: This field may return null, indicating that no valid values can be obtained.
	KewargaNegaraan *string `json:"KewargaNegaraan,omitnil,omitempty" name:"KewargaNegaraan"`

	// ID card validity period
	// Note: This field may return null, indicating that no valid values can be obtained.
	BerlakuHingga *string `json:"BerlakuHingga,omitnil,omitempty" name:"BerlakuHingga"`

	// Date of issue
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedDate *string `json:"IssuedDate,omitnil,omitempty" name:"IssuedDate"`

	// Province
	// Note: This field may return null, indicating that no valid values can be obtained.
	Provinsi *string `json:"Provinsi,omitnil,omitempty" name:"Provinsi"`

	// City
	// Note: This field may return null, indicating that no valid values can be obtained.
	Kota *string `json:"Kota,omitnil,omitempty" name:"Kota"`
}

type InternationalIDPassport struct {
	// Passport ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Full name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Last name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Surname *string `json:"Surname,omitnil,omitempty" name:"Surname"`

	// First name
	// Note: This field may return null, indicating that no valid values can be obtained.
	GivenName *string `json:"GivenName,omitnil,omitempty" name:"GivenName"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Gender (F-Female, M-Male)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Expiration date
	// Note: This field may return null, indicating that no valid values can be obtained.
	DateOfExpiration *string `json:"DateOfExpiration,omitnil,omitempty" name:"DateOfExpiration"`

	// Issuing country
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuingCountry *string `json:"IssuingCountry,omitnil,omitempty" name:"IssuingCountry"`

	// Nationality code
	// Note: This field may return null, indicating that no valid values can be obtained.
	NationalityCode *string `json:"NationalityCode,omitnil,omitempty" name:"NationalityCode"`

	// The first line at the bottom, the MRZ Code sequence
	// Note: This field may return null, indicating that no valid values can be obtained.
	PassportCodeFirst *string `json:"PassportCodeFirst,omitnil,omitempty" name:"PassportCodeFirst"`

	// The second line at the bottom, the MRZ Code sequence
	// Note: This field may return null, indicating that no valid values can be obtained.
	PassportCodeSecond *string `json:"PassportCodeSecond,omitnil,omitempty" name:"PassportCodeSecond"`
}

// Predefined struct for user
type LivenessCompareRequestParams struct {
	// Liveness detection type. Valid values: LIP/ACTION/SILENT.
	// LIP: numeric mode; ACTION: motion mode; SILENT: silent mode. You need to select a mode to input.
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// Base64 string of the image for face comparison.
	// The size of the Base64-encoded image data can be up to 3 MB. JPG and PNG formats are supported.
	// Please use the standard Base64 encoding scheme (with the "=" padding). For the encoding conventions, please see RFC 4648.
	// 
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageBase64` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL of the image for face comparison. The size of the downloaded image after Base64 encoding can be up to 3 MB. JPG and PNG formats are supported.
	// 
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageBase64` will be used.
	// 
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Lip mode: set this parameter to a custom 4-digit verification code.
	// Action mode: set this parameter to a custom action sequence (e.g., `2,1` or `1,2`).
	// Silent mode: do not pass in this parameter.
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`

	// Optional configuration (a JSON string)
	// {
	// "BestFrameNum": 2  // Return multiple best screenshots. Value range: 2−10
	// }
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// Base64 string of the video for liveness detection.
	// The size of the Base64-encoded video data can be up to 8 MB. MP4, AVI, and FLV formats are supported.
	// Please use the standard Base64 encoding scheme (with the "=" padding). For the encoding conventions, please see RFC 4648.
	// 
	// Either the `VideoUrl` or `VideoBase64` of the video must be provided. If both are provided, only `VideoBase64` will be used.
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// URL of the video for liveness detection. The size of the downloaded video after Base64 encoding can be up to 8 MB. It takes no more than 4 seconds to download. MP4, AVI, and FLV formats are supported.
	// 
	// Either the `VideoUrl` or `VideoBase64` of the video must be provided. If both are provided, only `VideoBase64` will be used.
	// 
	// We recommend you store the video in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`
}

type LivenessCompareRequest struct {
	*tchttp.BaseRequest
	
	// Liveness detection type. Valid values: LIP/ACTION/SILENT.
	// LIP: numeric mode; ACTION: motion mode; SILENT: silent mode. You need to select a mode to input.
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// Base64 string of the image for face comparison.
	// The size of the Base64-encoded image data can be up to 3 MB. JPG and PNG formats are supported.
	// Please use the standard Base64 encoding scheme (with the "=" padding). For the encoding conventions, please see RFC 4648.
	// 
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageBase64` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL of the image for face comparison. The size of the downloaded image after Base64 encoding can be up to 3 MB. JPG and PNG formats are supported.
	// 
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageBase64` will be used.
	// 
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Lip mode: set this parameter to a custom 4-digit verification code.
	// Action mode: set this parameter to a custom action sequence (e.g., `2,1` or `1,2`).
	// Silent mode: do not pass in this parameter.
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`

	// Optional configuration (a JSON string)
	// {
	// "BestFrameNum": 2  // Return multiple best screenshots. Value range: 2−10
	// }
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// Base64 string of the video for liveness detection.
	// The size of the Base64-encoded video data can be up to 8 MB. MP4, AVI, and FLV formats are supported.
	// Please use the standard Base64 encoding scheme (with the "=" padding). For the encoding conventions, please see RFC 4648.
	// 
	// Either the `VideoUrl` or `VideoBase64` of the video must be provided. If both are provided, only `VideoBase64` will be used.
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// URL of the video for liveness detection. The size of the downloaded video after Base64 encoding can be up to 8 MB. It takes no more than 4 seconds to download. MP4, AVI, and FLV formats are supported.
	// 
	// Either the `VideoUrl` or `VideoBase64` of the video must be provided. If both are provided, only `VideoBase64` will be used.
	// 
	// We recommend you store the video in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`
}

func (r *LivenessCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LivenessCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LivenessType")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ValidateData")
	delete(f, "Optional")
	delete(f, "VideoBase64")
	delete(f, "VideoUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LivenessCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LivenessCompareResponseParams struct {
	// The best screenshot of the video after successful verification. The photo is Base64-encoded and in JPG format.
	BestFrameBase64 *string `json:"BestFrameBase64,omitnil,omitempty" name:"BestFrameBase64"`

	// Similarity. Value range: [0.00, 100.00]. As a recommendation, when the similarity is greater than or equal to 70, it can be determined that the two faces are of the same person. You can adjust the threshold according to your specific scenario (the FAR at the threshold of 70 is 0.1%, and FAR at the threshold of 80 is 0.01%).
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// Service error code. `Success` will be returned for success. For error information, please see the `FailedOperation` section in the error code list below.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// Service result description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`


	BestFrameList []*string `json:"BestFrameList,omitnil,omitempty" name:"BestFrameList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LivenessCompareResponse struct {
	*tchttp.BaseResponse
	Response *LivenessCompareResponseParams `json:"Response"`
}

func (r *LivenessCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LivenessCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MLIDCard struct {
	// Full Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Gender
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Lisence type
	// MyKad ID card
	// MyPR Permanent resident ID card
	// MyTentera Military ID card
	// MyKAS Temporary ID card
	// POLIS Police ID card
	// IKAD Labor ID card
	// MyKid Juvenile ID card
	// Example: MyKad
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Birthday (Currently, this filed only supports IKAD labor ID card and MyKad ID card)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`
}

type MacaoIDCard struct {
	// First name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// Last name
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Expiration date
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Sex
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Age
	// Note: This field may return null, indicating that no valid values can be obtained.
	Age *string `json:"Age,omitnil,omitempty" name:"Age"`

	// Issued country
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedCountry *string `json:"IssuedCountry,omitnil,omitempty" name:"IssuedCountry"`

	// MRZ1 on card
	// Note: This field may return null, indicating that no valid values can be obtained. 
	Field1 *string `json:"Field1,omitnil,omitempty" name:"Field1"`

	// MRZ2 on card
	// Note: This field may return null, indicating that no valid values can be obtained.
	Field2 *string `json:"Field2,omitnil,omitempty" name:"Field2"`
}

type MainlandIDCard struct {
	// Chinese name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Sex
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Nation
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nation *string `json:"Nation,omitnil,omitempty" name:"Nation"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: Address is deprecated.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormattedAddress *string `json:"FormattedAddress,omitnil,omitempty" name:"FormattedAddress"`
}

type NormalCardInfo struct {
	// Hong Kong ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	HKIDCard *NormalHKIDCard `json:"HKIDCard,omitnil,omitempty" name:"HKIDCard"`

	// Malaysia ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	MLIDCard *NormalMLIDCard `json:"MLIDCard,omitnil,omitempty" name:"MLIDCard"`

	// Philippines VoteID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesVoteID *PhilippinesVoteID `json:"PhilippinesVoteID,omitnil,omitempty" name:"PhilippinesVoteID"`

	// Indonesia ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndonesiaIDCard *NormalIndonesiaIDCard `json:"IndonesiaIDCard,omitnil,omitempty" name:"IndonesiaIDCard"`

	// Philippines Driving License
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesDrivingLicense *PhilippinesDrivingLicense `json:"PhilippinesDrivingLicense,omitnil,omitempty" name:"PhilippinesDrivingLicense"`

	// Philippines TinID
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesTinID *PhilippinesTinID `json:"PhilippinesTinID,omitnil,omitempty" name:"PhilippinesTinID"`

	// Philippines SSSID
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesSSSID *PhilippinesSSSID `json:"PhilippinesSSSID,omitnil,omitempty" name:"PhilippinesSSSID"`

	// Philippines UMID
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhilippinesUMID *PhilippinesUMID `json:"PhilippinesUMID,omitnil,omitempty" name:"PhilippinesUMID"`

	// ID Cards of Hong Kong, Macao and Taiwan (China), and International Passport
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternationalIDPassport *InternationalIDPassport `json:"InternationalIDPassport,omitnil,omitempty" name:"InternationalIDPassport"`

	// General license information
	// Note: This field may return null, indicating that no valid values can be obtained.
	GeneralCard *GeneralCard `json:"GeneralCard,omitnil,omitempty" name:"GeneralCard"`

	// Indonesia Driving License
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndonesiaDrivingLicense *IndonesiaDrivingLicense `json:"IndonesiaDrivingLicense,omitnil,omitempty" name:"IndonesiaDrivingLicense"`

	// Thailand ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	ThailandIDCard *NormalThailandIDCard `json:"ThailandIDCard,omitnil,omitempty" name:"ThailandIDCard"`

	// Singapore ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	SingaporeIDCard *SingaporeIDCard `json:"SingaporeIDCard,omitnil,omitempty" name:"SingaporeIDCard"`

	// Macao ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	MacaoIDCard *MacaoIDCard `json:"MacaoIDCard,omitnil,omitempty" name:"MacaoIDCard"`

	// Mainland ID Card
	// Note: This field may return null, indicating that no valid values can be obtained.
	MainlandIDCard *MainlandIDCard `json:"MainlandIDCard,omitnil,omitempty" name:"MainlandIDCard"`
}

type NormalHKIDCard struct {
	// Chinese name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChineseName *string `json:"ChineseName,omitnil,omitempty" name:"ChineseName"`

	// English name
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example: SAN, Nan
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Telex code correspondint to the Chinese name
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegistrationNumber *string `json:"RegistrationNumber,omitnil,omitempty" name:"RegistrationNumber"`

	// Gender: "Male-M" or "Female-F"
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example: 01-01-2001
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Permanent resident ID card: 0-non-permanent; 1-permanent; -1-unknown
	// Note: This field may return null, indicating that no valid values can be obtained.
	Permanent *string `json:"Permanent,omitnil,omitempty" name:"Permanent"`

	// ID card number
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example: C000000(E)
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Lisence symbol, which is the symbol below Birthday. Example: "***AZ"
	// Note: This field may return null, indicating that no valid values can be obtained.
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// The first date of issue
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedDate *string `json:"IssuedDate,omitnil,omitempty" name:"IssuedDate"`

	// The current date of issue
	// Note: This field may return null, indicating that no valid values can be obtained.
	CurrentIssueDate *string `json:"CurrentIssueDate,omitnil,omitempty" name:"CurrentIssueDate"`
}

type NormalIndonesiaIDCard struct {
	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Birth place/Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Gender
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Blood type
	// Note: This field may return null, indicating that no valid values can be obtained.
	BloodType *string `json:"BloodType,omitnil,omitempty" name:"BloodType"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormattedAddress *string `json:"FormattedAddress,omitnil,omitempty" name:"FormattedAddress"`

	// Street
	// Note: This field may return null, indicating that no valid values can be obtained.
	Street *string `json:"Street,omitnil,omitempty" name:"Street"`

	// Village
	// Note: This field may return null, indicating that no valid values can be obtained.
	Village *string `json:"Village,omitnil,omitempty" name:"Village"`

	// Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Religious beliefs
	// Note: This field may return null, indicating that no valid values can be obtained.
	Religion *string `json:"Religion,omitnil,omitempty" name:"Religion"`

	// Marital status
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaritalStatus *string `json:"MaritalStatus,omitnil,omitempty" name:"MaritalStatus"`

	// Job
	// Note: This field may return null, indicating that no valid values can be obtained.
	Occupation *string `json:"Occupation,omitnil,omitempty" name:"Occupation"`

	// Nationality
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// ID card validity period
	// Note: This field may return null, indicating that no valid values can be obtained.
	DueDate *string `json:"DueDate,omitnil,omitempty" name:"DueDate"`

	// Date of issue
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedDate *string `json:"IssuedDate,omitnil,omitempty" name:"IssuedDate"`

	// Province
	// Note: This field may return null, indicating that no valid values can be obtained.
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// City
	// Note: This field may return null, indicating that no valid values can be obtained.
	City *string `json:"City,omitnil,omitempty" name:"City"`
}

type NormalMLIDCard struct {
	// Full Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Gender
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormattedAddress *string `json:"FormattedAddress,omitnil,omitempty" name:"FormattedAddress"`

	// Lisence type
	// MyKad ID card
	// MyPR Permanent resident ID card
	// MyTentera Military ID card
	// MyKAS Temporary ID card
	// POLIS Police ID card
	// IKAD Labor ID card
	// MyKid Juvenile ID card
	// Example: MyKad
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Birthday (Currently, this filed only supports IKAD labor ID card and MyKad ID card)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`
}

type NormalThailandIDCard struct {
	// LicenseNumber
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Thailand name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Last name
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// First name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormattedAddress *string `json:"FormattedAddress,omitnil,omitempty" name:"FormattedAddress"`

	// Expiration date
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

	// Issued date
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedDate *string `json:"IssuedDate,omitnil,omitempty" name:"IssuedDate"`

	// Registration number 
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegistrationNumber *string `json:"RegistrationNumber,omitnil,omitempty" name:"RegistrationNumber"`

	// Religion
	// Note: This field may return null, indicating that no valid values can be obtained.
	Religion *string `json:"Religion,omitnil,omitempty" name:"Religion"`

	// Birthday in Thai
	// Note: This field may return null, indicating that no valid values can be obtained.
	ThaiBirthday *string `json:"ThaiBirthday,omitnil,omitempty" name:"ThaiBirthday"`

	// Expiration date in Thai
	// Note: This field may return null, indicating that no valid values can be obtained.
	ThaiExpirationDate *string `json:"ThaiExpirationDate,omitnil,omitempty" name:"ThaiExpirationDate"`

	// Issued date in Thai
	// Note: This field may return null, indicating that no valid values can be obtained.
	ThaiIssueDate *string `json:"ThaiIssueDate,omitnil,omitempty" name:"ThaiIssueDate"`
}

type OCRResult struct {
	// Is the indentity verification or OCR process passed
	IsPass *bool `json:"IsPass,omitnil,omitempty" name:"IsPass"`

	// The Base64 of ID card image
	// Note: This field may return null, indicating that no valid values can be obtained.
	CardImageBase64 *string `json:"CardImageBase64,omitnil,omitempty" name:"CardImageBase64"`

	// OCR result of the ID card.
	CardInfo *CardInfo `json:"CardInfo,omitnil,omitempty" name:"CardInfo"`

	// The request id
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// Base64 of cropped image of ID card
	CardCutImageBase64 *string `json:"CardCutImageBase64,omitnil,omitempty" name:"CardCutImageBase64"`

	// Base64 of the cropped image on the reverse side of the ID card
	CardBackCutImageBase64 *string `json:"CardBackCutImageBase64,omitnil,omitempty" name:"CardBackCutImageBase64"`
}

type PhilippinesDrivingLicense struct {
	// Full Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Last name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// First name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// Middle name
	// Note: This field may return null, indicating that no valid values can be obtained.
	MiddleName *string `json:"MiddleName,omitnil,omitempty" name:"MiddleName"`

	// Nationality
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Gender
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNo *string `json:"LicenseNo,omitnil,omitempty" name:"LicenseNo"`

	// Date of expiry
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpiresDate *string `json:"ExpiresDate,omitnil,omitempty" name:"ExpiresDate"`

	// Agency code
	// Note: This field may return null, indicating that no valid values can be obtained.
	AgencyCode *string `json:"AgencyCode,omitnil,omitempty" name:"AgencyCode"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`
}

type PhilippinesSSSID struct {
	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Full name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`
}

type PhilippinesTinID struct {
	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Full name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Date of issue
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssueDate *string `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`
}

type PhilippinesUMID struct {
	// Surname
	// Note: This field may return null, indicating that no valid values can be obtained.
	Surname *string `json:"Surname,omitnil,omitempty" name:"Surname"`

	// Middle Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	MiddleName *string `json:"MiddleName,omitnil,omitempty" name:"MiddleName"`

	// First name
	// Note: This field may return null, indicating that no valid values can be obtained.
	GivenName *string `json:"GivenName,omitnil,omitempty" name:"GivenName"`

	// Gender
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// CRN code
	// Note: This field may return null, indicating that no valid values can be obtained.
	CRN *string `json:"CRN,omitnil,omitempty" name:"CRN"`
}

type PhilippinesVoteID struct {
	// VIN of Philippines VoteID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VIN *string `json:"VIN,omitnil,omitempty" name:"VIN"`

	// First name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// Last name
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Civil status
	// Note: This field may return null, indicating that no valid values can be obtained.
	CivilStatus *string `json:"CivilStatus,omitnil,omitempty" name:"CivilStatus"`

	// Nationality
	// Note: This field may return null, indicating that no valid values can be obtained.
	Citizenship *string `json:"Citizenship,omitnil,omitempty" name:"Citizenship"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrecinctNo *string `json:"PrecinctNo,omitnil,omitempty" name:"PrecinctNo"`
}

type SingaporeIDCard struct {
	// Chinese name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChName *string `json:"ChName,omitnil,omitempty" name:"ChName"`

	// English name
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Gender
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Country of birth
	// Note: This field may return null, indicating that no valid values can be obtained.
	CountryOfBirth *string `json:"CountryOfBirth,omitnil,omitempty" name:"CountryOfBirth"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Address (on the back)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Nationality (on the back)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Race *string `json:"Race,omitnil,omitempty" name:"Race"`

	//  NRIC number (on the back)
	// Note: This field may return null, indicating that no valid values can be obtained.
	NRICCode *string `json:"NRICCode,omitnil,omitempty" name:"NRICCode"`

	// Post number (on the front)
	// Note: This field may return null, indicating that no valid values can be obtained.
	PostCode *string `json:"PostCode,omitnil,omitempty" name:"PostCode"`

	// Date of expiry (on the back)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DateOfExpiration *string `json:"DateOfExpiration,omitnil,omitempty" name:"DateOfExpiration"`

	// Date of issue (on the back)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DateOfIssue *string `json:"DateOfIssue,omitnil,omitempty" name:"DateOfIssue"`
}

type ThailandIDCard struct {
	// Last name
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// First name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// License number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Birthday
	// Note: This field may return null, indicating that no valid values can be obtained.
	DateOfBirth *string `json:"DateOfBirth,omitnil,omitempty" name:"DateOfBirth"`

	// Date of expiry
	// Note: This field may return null, indicating that no valid values can be obtained.
	DateOfExpiry *string `json:"DateOfExpiry,omitnil,omitempty" name:"DateOfExpiry"`

	// Date of issue
	// Note: This field may return null, indicating that no valid values can be obtained.
	DateOfIssue *string `json:"DateOfIssue,omitnil,omitempty" name:"DateOfIssue"`

	// Issuing country
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedCountry *string `json:"IssuedCountry,omitnil,omitempty" name:"IssuedCountry"`
}

type VerificationDetail struct {
	// The final result of this verification. `0` indicates that the person is the same as that in the photo.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// The description of the final verification result.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The result of this liveness detection process. `0` indicates success.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	LivenessErrorCode *int64 `json:"LivenessErrorCode,omitnil,omitempty" name:"LivenessErrorCode"`

	// The result description of this liveness detection process.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LivenessErrorMsg *string `json:"LivenessErrorMsg,omitnil,omitempty" name:"LivenessErrorMsg"`

	// The result of this comparison process. `0` indicates that the person in the best face screenshot collected from the video stream is the same as that in the uploaded image for comparison.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	CompareErrorCode *int64 `json:"CompareErrorCode,omitnil,omitempty" name:"CompareErrorCode"`

	// The result description of this comparison process.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	CompareErrorMsg *string `json:"CompareErrorMsg,omitnil,omitempty" name:"CompareErrorMsg"`

	// The timestamp (ms) of this verification process.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReqTimestamp *uint64 `json:"ReqTimestamp,omitnil,omitempty" name:"ReqTimestamp"`

	// The similarity of the best face screenshot collected from the video stream and the uploaded image for comparison in this verification process. Value range: [0.00, 100.00]. By default, the person in the screenshot is determined to be the same person in the image if the similarity is greater than or equal to 70.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Similarity *float64 `json:"Similarity,omitnil,omitempty" name:"Similarity"`

	// Unique ID of this verification process.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Seq *string `json:"Seq,omitnil,omitempty" name:"Seq"`
}

// Predefined struct for user
type VideoLivenessCompareRequestParams struct {
	// The URL of the photo for face comparison. The downloaded image after Base64 encoding can be up to 3 MB and must be in JPG or PNG.
	// 
	// The image must be stored in a COS bucket in the region where the FaceID service resides to ensure a higher download speed and better stability. You can generate an image URL by using `CreateUploadUrl` or purchase the COS service.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The 32-bit MD5 checksum of the image for comparison
	ImageMd5 *string `json:"ImageMd5,omitnil,omitempty" name:"ImageMd5"`

	// The URL of the video for liveness detection. The downloaded video after Base64 encoding can be up to 8 MB and must be in MP4, AVI, or FLV. It takes no more than 4s to download the video.
	// 
	// The video must be stored in a COS bucket in the region where the FaceID service resides to ensure a higher download speed and better stability. You can generate a video URL by using `CreateUploadUrl` or purchase the COS service.
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// The 32-bit MD5 checksum of the video
	VideoMd5 *string `json:"VideoMd5,omitnil,omitempty" name:"VideoMd5"`

	// The liveness detection type. Valid values: `LIP`, `ACTION`, and `SILENT`.
	// `LIP`: Numeric mode; `ACTION`: Motion mode; `SILENT`: silent mode. Select one of them.
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// LIP parameter: Pass in a custom 4-digit verification code.
	// ACTION parameter: Pass in a custom action sequence (`2,1` or `1,2`).
	// SILENT parameter: Null.
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`
}

type VideoLivenessCompareRequest struct {
	*tchttp.BaseRequest
	
	// The URL of the photo for face comparison. The downloaded image after Base64 encoding can be up to 3 MB and must be in JPG or PNG.
	// 
	// The image must be stored in a COS bucket in the region where the FaceID service resides to ensure a higher download speed and better stability. You can generate an image URL by using `CreateUploadUrl` or purchase the COS service.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The 32-bit MD5 checksum of the image for comparison
	ImageMd5 *string `json:"ImageMd5,omitnil,omitempty" name:"ImageMd5"`

	// The URL of the video for liveness detection. The downloaded video after Base64 encoding can be up to 8 MB and must be in MP4, AVI, or FLV. It takes no more than 4s to download the video.
	// 
	// The video must be stored in a COS bucket in the region where the FaceID service resides to ensure a higher download speed and better stability. You can generate a video URL by using `CreateUploadUrl` or purchase the COS service.
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// The 32-bit MD5 checksum of the video
	VideoMd5 *string `json:"VideoMd5,omitnil,omitempty" name:"VideoMd5"`

	// The liveness detection type. Valid values: `LIP`, `ACTION`, and `SILENT`.
	// `LIP`: Numeric mode; `ACTION`: Motion mode; `SILENT`: silent mode. Select one of them.
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// LIP parameter: Pass in a custom 4-digit verification code.
	// ACTION parameter: Pass in a custom action sequence (`2,1` or `1,2`).
	// SILENT parameter: Null.
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`
}

func (r *VideoLivenessCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VideoLivenessCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageMd5")
	delete(f, "VideoUrl")
	delete(f, "VideoMd5")
	delete(f, "LivenessType")
	delete(f, "ValidateData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VideoLivenessCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VideoLivenessCompareResponseParams struct {
	// The similarity. Value range: [0.00, 100.00]. As a recommendation, when the similarity is greater than or equal to 70, it can be determined that the two persons are of the same person. You can adjust the threshold according to your specific scenario (the FARs at the thresholds of 70 and 80 are 0.1% and 0.01%, respectively).
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// The service error code. `Success` will be returned for success. For error information, see the `FailedOperation` section in the error code list below.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The service result description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The best video screenshot after successful verification
	// Note: This field may return null, indicating that no valid values can be obtained.
	BestFrame *FileInfo `json:"BestFrame,omitnil,omitempty" name:"BestFrame"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VideoLivenessCompareResponse struct {
	*tchttp.BaseResponse
	Response *VideoLivenessCompareResponseParams `json:"Response"`
}

func (r *VideoLivenessCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VideoLivenessCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebVerificationConfigIntl struct {
	// When starting verification, whether to skip the starting verification page. If true, enter the verification process directly. The default is false. This configuration will not take effect if the downgrade policy is triggered.
	AutoSkipStartPage *bool `json:"AutoSkipStartPage,omitnil,omitempty" name:"AutoSkipStartPage"`

	// When the verification passed, whether to skip the result page and automatically jump to RedirectURL. The default value is false.
	// Example value: false
	AutoSkip *bool `json:"AutoSkip,omitnil,omitempty" name:"AutoSkip"`

	// Detection mode, parameter values are as follows:
	// 1: OCR+living detection & face comparison;
	// 2: Living detection & face comparison;
	// 3: Living detection;
	// The default value is 2.
	// Example value: 3
	CheckMode *int64 `json:"CheckMode,omitnil,omitempty" name:"CheckMode"`

	// The type of lisence used for verification. The following types are supported.
	// 1.HKIDCard: Hong Kong (China) ID card
	// 2.MLIDCard: Malaysia ID card
	// 3.IndonesiaIDCard: Indonesia ID card
	// 4.PhilippinesVoteID: Philippines VoteID card
	// 5.PhilippinesDrivingLicense: Philippines driving license
	// 6.PhilippinesTinID: Philippines TinID card
	// 7.PhilippinesSSSID: Philippines SSSID card
	// 8.PhilippinesUMID: Philippines UMID card
	// 9.InternationalIDPassport: ID cards of Hong Kong, Macao and Taiwan (China), and international passport.
	// Example: HKIDCard
	IDCardType *string `json:"IDCardType,omitnil,omitempty" name:"IDCardType"`

	// Whether to turn off document alarms, the default is false (the alarm detection function is turned on). When enabled, the identity authentication process will be intercepted based on the alarm status of the certificate. If you need to use the document authentication function, please contact us.
	DisableCheckOcrWarnings *bool `json:"DisableCheckOcrWarnings,omitnil,omitempty" name:"DisableCheckOcrWarnings"`

	// Liveness security level: 1:Silent mode;2:Action mode;3:Lighting mode;4:Action+Lighting mode;default value is 3
	SecurityLevel *int64 `json:"SecurityLevel,omitnil,omitempty" name:"SecurityLevel"`

	// Whether to skip the agreement page, the default is false. When SkipPrivacyPolicy is false, the agreement page will be displayed and the privacy agreement needs to be checked; when SkipPrivacyPolicy is true, the agreement page will be skipped and the liveness process will be entered directly without checking the privacy agreement page.
	SkipPrivacyPolicy *bool `json:"SkipPrivacyPolicy,omitnil,omitempty" name:"SkipPrivacyPolicy"`

	// The default value is false. If it is false, the original ID image will be displayed. If it is true, the cut ID image will be displayed.
	IdCardCutReturn *bool `json:"IdCardCutReturn,omitnil,omitempty" name:"IdCardCutReturn"`

	// Front-end theme color, in the format of RGB hexadecimal color code. The default value is "#2d72+1". If the format is incorrect, the default value color will be used.
	ThemeColor *string `json:"ThemeColor,omitnil,omitempty" name:"ThemeColor"`

	// International language, the default value is en (English). Currently supported: th: Thai; en: English;
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// Automatic downgrade mode, with the following parameter values: 1: Downgrade to silent live mode; 2: Disable downgrade mode. The default value is 1.
	AutoDowngrade *int64 `json:"AutoDowngrade,omitnil,omitempty" name:"AutoDowngrade"`
}