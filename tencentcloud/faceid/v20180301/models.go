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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type ApplyLivenessTokenRequestParams struct {
	// Enumerated value. Valid values: `1`, `2`, `3`, and `4`.
	// Their meanings are as follows:
	// 1 - silent
	// 2 - blinking
	// 3 - light
	// 4 - blinking + light (default)
	SecureLevel *string `json:"SecureLevel,omitempty" name:"SecureLevel"`
}

type ApplyLivenessTokenRequest struct {
	*tchttp.BaseRequest
	
	// Enumerated value. Valid values: `1`, `2`, `3`, and `4`.
	// Their meanings are as follows:
	// 1 - silent
	// 2 - blinking
	// 3 - light
	// 4 - blinking + light (default)
	SecureLevel *string `json:"SecureLevel,omitempty" name:"SecureLevel"`
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
	SdkToken *string `json:"SdkToken,omitempty" name:"SdkToken"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Whether ID card authentication is required. If not, only document OCR will be performed. Currently, authentication is available only when the value of `IdCardType` is `HK`.
	NeedVerifyIdCard *bool `json:"NeedVerifyIdCard,omitempty" name:"NeedVerifyIdCard"`

	// The verification mode. Valid values:
	// 1: OCR + liveness detection + face comparison
	// 2: Liveness detection + face comparison
	// 3: Liveness detection
	// Default value: 1
	CheckMode *int64 `json:"CheckMode,omitempty" name:"CheckMode"`

	// The security level of the verification. Valid values:
	// 1: Video-based liveness detection
	// 2: Motion-based liveness detection
	// 3: Reflection-based liveness detection
	// 4: Motion- and reflection-based liveness detection
	// Default value: 4
	SecurityLevel *int64 `json:"SecurityLevel,omitempty" name:"SecurityLevel"`

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
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// The Base64-encoded value of the photo to compare, which is required only when `CheckMode` is set to `2`.
	CompareImage *string `json:"CompareImage,omitempty" name:"CompareImage"`

	// Whether to forbid the modification of the OCR result by users. Default value: `false` (modification allowed). (Currently, this parameter is not applied.)
	DisableChangeOcrResult *bool `json:"DisableChangeOcrResult,omitempty" name:"DisableChangeOcrResult"`

	// Whether to disable the OCR warnings. Default value: `false` (not disable), where OCR warnings are enabled and the OCR result will not be returned if there is a warning.
	// This feature applies only to Hong Kong (China) identity cards, Malaysian identity cards, and passports.
	DisableCheckOcrWarnings *bool `json:"DisableCheckOcrWarnings,omitempty" name:"DisableCheckOcrWarnings"`

	// A passthrough field, which is returned together with the verification result and can contain up to 1,024 bits.
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

type ApplySdkVerificationTokenRequest struct {
	*tchttp.BaseRequest
	
	// Whether ID card authentication is required. If not, only document OCR will be performed. Currently, authentication is available only when the value of `IdCardType` is `HK`.
	NeedVerifyIdCard *bool `json:"NeedVerifyIdCard,omitempty" name:"NeedVerifyIdCard"`

	// The verification mode. Valid values:
	// 1: OCR + liveness detection + face comparison
	// 2: Liveness detection + face comparison
	// 3: Liveness detection
	// Default value: 1
	CheckMode *int64 `json:"CheckMode,omitempty" name:"CheckMode"`

	// The security level of the verification. Valid values:
	// 1: Video-based liveness detection
	// 2: Motion-based liveness detection
	// 3: Reflection-based liveness detection
	// 4: Motion- and reflection-based liveness detection
	// Default value: 4
	SecurityLevel *int64 `json:"SecurityLevel,omitempty" name:"SecurityLevel"`

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
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// The Base64-encoded value of the photo to compare, which is required only when `CheckMode` is set to `2`.
	CompareImage *string `json:"CompareImage,omitempty" name:"CompareImage"`

	// Whether to forbid the modification of the OCR result by users. Default value: `false` (modification allowed). (Currently, this parameter is not applied.)
	DisableChangeOcrResult *bool `json:"DisableChangeOcrResult,omitempty" name:"DisableChangeOcrResult"`

	// Whether to disable the OCR warnings. Default value: `false` (not disable), where OCR warnings are enabled and the OCR result will not be returned if there is a warning.
	// This feature applies only to Hong Kong (China) identity cards, Malaysian identity cards, and passports.
	DisableCheckOcrWarnings *bool `json:"DisableCheckOcrWarnings,omitempty" name:"DisableCheckOcrWarnings"`

	// A passthrough field, which is returned together with the verification result and can contain up to 1,024 bits.
	Extra *string `json:"Extra,omitempty" name:"Extra"`
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
	delete(f, "NeedVerifyIdCard")
	delete(f, "CheckMode")
	delete(f, "SecurityLevel")
	delete(f, "IdCardType")
	delete(f, "CompareImage")
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
	SdkToken *string `json:"SdkToken,omitempty" name:"SdkToken"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// The Base64-encoded string (max 8 MB in size) of the photo to be compared.
	CompareImageBase64 *string `json:"CompareImageBase64,omitempty" name:"CompareImageBase64"`

	// The web callback URL to redirect to after the verification is completed, including the protocol, hostname, and path. Example: `https://www.tencentcloud.com/products/faceid`.
	// After the verification process is completed, the `BizToken` of this process will be spliced to the callback URL in the format of `https://www.tencentcloud.com/products/faceid?token={BizToken}` before redirect.
	RedirectURL *string `json:"RedirectURL,omitempty" name:"RedirectURL"`

	// The passthrough parameter of the business, max 1,000 characters, which will be returned in `GetWebVerificationResultIntl`.
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

type ApplyWebVerificationBizTokenIntlRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded string (max 8 MB in size) of the photo to be compared.
	CompareImageBase64 *string `json:"CompareImageBase64,omitempty" name:"CompareImageBase64"`

	// The web callback URL to redirect to after the verification is completed, including the protocol, hostname, and path. Example: `https://www.tencentcloud.com/products/faceid`.
	// After the verification process is completed, the `BizToken` of this process will be spliced to the callback URL in the format of `https://www.tencentcloud.com/products/faceid?token={BizToken}` before redirect.
	RedirectURL *string `json:"RedirectURL,omitempty" name:"RedirectURL"`

	// The passthrough parameter of the business, max 1,000 characters, which will be returned in `GetWebVerificationResultIntl`.
	Extra *string `json:"Extra,omitempty" name:"Extra"`
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
	delete(f, "CompareImageBase64")
	delete(f, "RedirectURL")
	delete(f, "Extra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyWebVerificationBizTokenIntlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyWebVerificationBizTokenIntlResponseParams struct {
	// The URL of this verification process, which will be returned to the frontend of the browser for starting the process.
	VerificationUrl *string `json:"VerificationUrl,omitempty" name:"VerificationUrl"`

	// The token identifying this web-based verification process, valid for 7,200s after issuance. It is required for getting the result after the verification process is completed.
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// The COS URL of the image for face comparison, which can be obtained with one of the following methods:
	// 1. Call the `CreateUploadUrl` API to generate a URL and call it again after the image is successfully uploaded.
	// 2. Use an existing COS URL. For a private bucket, grant the download permission with a pre-signed URL. The corresponding COS bucket must be in the same region as the input parameter `Region`.
	CompareImageUrl *string `json:"CompareImageUrl,omitempty" name:"CompareImageUrl"`

	// The MD5 hash values of the image for face comparison (CompareImageUrl).
	CompareImageMd5 *string `json:"CompareImageMd5,omitempty" name:"CompareImageMd5"`
}

type ApplyWebVerificationTokenRequest struct {
	*tchttp.BaseRequest
	
	// The web redirect URL after the verification is completed.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// The COS URL of the image for face comparison, which can be obtained with one of the following methods:
	// 1. Call the `CreateUploadUrl` API to generate a URL and call it again after the image is successfully uploaded.
	// 2. Use an existing COS URL. For a private bucket, grant the download permission with a pre-signed URL. The corresponding COS bucket must be in the same region as the input parameter `Region`.
	CompareImageUrl *string `json:"CompareImageUrl,omitempty" name:"CompareImageUrl"`

	// The MD5 hash values of the image for face comparison (CompareImageUrl).
	CompareImageMd5 *string `json:"CompareImageMd5,omitempty" name:"CompareImageMd5"`
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
	VerificationUrl *string `json:"VerificationUrl,omitempty" name:"VerificationUrl"`

	// The token used to identify a web-based verification process. It is valid for 7,200s and can be used to get the verification result after the process is completed.
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type CardVerifyResult struct {
	// Whether the authentication or OCR process is successful.
	IsPass *bool `json:"IsPass,omitempty" name:"IsPass"`

	// The download URL of the video used for identity document verification, which is valid for 10 minutes. This parameter is returned only if video-based identity document verification is enabled.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CardVideo *FileInfo `json:"CardVideo,omitempty" name:"CardVideo"`

	// The download URL of the identity document image, which is valid for 10 minutes.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CardImage *FileInfo `json:"CardImage,omitempty" name:"CardImage"`

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
	CardInfoOcrJson *FileInfo `json:"CardInfoOcrJson,omitempty" name:"CardInfoOcrJson"`

	// The request ID of a single process.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// The description of the final verification result.
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// The liveness algorithm package generated during this SDK-based liveness detection.
	LiveData *FileInfo `json:"LiveData,omitempty" name:"LiveData"`

	// The download URL of the video used for verification, which is valid for 10 minutes.
	LiveVideo *FileInfo `json:"LiveVideo,omitempty" name:"LiveVideo"`

	// The liveness detection result code.
	// 0: Success.
	// 1001: Failed to call the liveness detection engine.
	// 1004: Face detection failed.
	LiveErrorCode *string `json:"LiveErrorCode,omitempty" name:"LiveErrorCode"`

	// The description of the liveness detection result.
	LiveErrorMsg *string `json:"LiveErrorMsg,omitempty" name:"LiveErrorMsg"`

	// The download URL of the face screenshot during verification, which is valid for 10 minutes.
	// Note: This field may return null, indicating that no valid value can be obtained.
	BestFrame *FileInfo `json:"BestFrame,omitempty" name:"BestFrame"`

	// The download URL of the profile photo screenshot from the identity document, which is valid for 10 minutes.
	ProfileImage *FileInfo `json:"ProfileImage,omitempty" name:"ProfileImage"`

	// The face comparison result code.
	// 0: Success.
	// 2004: The uploaded face image is too large or too small.
	// 2012: The face is not fully exposed.
	// 2013: No face is detected.
	// 2014: The resolution of the uploaded image is too low . Please upload a new one.
	// 2015: Face comparison failed.
	// 2016: The similarity did not reach the passing standard.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CompareErrorCode *string `json:"CompareErrorCode,omitempty" name:"CompareErrorCode"`

	// The description of the face comparison result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompareErrorMsg *string `json:"CompareErrorMsg,omitempty" name:"CompareErrorMsg"`

	// The similarity score of face comparison.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sim *float64 `json:"Sim,omitempty" name:"Sim"`

	// This parameter is disused.
	IsNeedCharge *bool `json:"IsNeedCharge,omitempty" name:"IsNeedCharge"`

	// The identity document photo info edited by the user. Currently, this parameter is not applied.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CardInfoInputJson *FileInfo `json:"CardInfoInputJson,omitempty" name:"CardInfoInputJson"`

	// The request ID of this verification process.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

// Predefined struct for user
type CreateUploadUrlRequestParams struct {
	// Target API
	TargetAction *string `json:"TargetAction,omitempty" name:"TargetAction"`
}

type CreateUploadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Target API
	TargetAction *string `json:"TargetAction,omitempty" name:"TargetAction"`
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
	UploadUrl *string `json:"UploadUrl,omitempty" name:"UploadUrl"`

	// The resource URL obtained after this upload is completed and to be passed in where it is required later.
	ResourceUrl *string `json:"ResourceUrl,omitempty" name:"ResourceUrl"`

	// The point in time when the upload/download link expires, which is a 10-bit Unix timestamp.
	ExpiredTimestamp *int64 `json:"ExpiredTimestamp,omitempty" name:"ExpiredTimestamp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LiveDataUrl *string `json:"LiveDataUrl,omitempty" name:"LiveDataUrl"`

	// MD5 hash value (32-bit) of the liveness detection data package generated by the SDK, which is used to verify the LiveData consistency.
	LiveDataMd5 *string `json:"LiveDataMd5,omitempty" name:"LiveDataMd5"`

	// URL of the target image for comparison
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// MD5 hash value (32-bit) of the target image for comparison, which is used to verify the `Image` consistency.
	ImageMd5 *string `json:"ImageMd5,omitempty" name:"ImageMd5"`
}

type DetectReflectLivenessAndCompareRequest struct {
	*tchttp.BaseRequest
	
	// URL of the liveness detection data package generated by the SDK
	LiveDataUrl *string `json:"LiveDataUrl,omitempty" name:"LiveDataUrl"`

	// MD5 hash value (32-bit) of the liveness detection data package generated by the SDK, which is used to verify the LiveData consistency.
	LiveDataMd5 *string `json:"LiveDataMd5,omitempty" name:"LiveDataMd5"`

	// URL of the target image for comparison
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// MD5 hash value (32-bit) of the target image for comparison, which is used to verify the `Image` consistency.
	ImageMd5 *string `json:"ImageMd5,omitempty" name:"ImageMd5"`
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
	BestFrameUrl *string `json:"BestFrameUrl,omitempty" name:"BestFrameUrl"`

	// MD5 hash value (32-bit) of the best screenshot of the video after successful verification, which is used to verify the `BestFrame` consistency.
	BestFrameMd5 *string `json:"BestFrameMd5,omitempty" name:"BestFrameMd5"`

	// Service error code. `Success` will be returned for success. For error information, see the `FailedOperation` section in the error code list below.
	Result *string `json:"Result,omitempty" name:"Result"`

	// Service result description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Similarity. Value range: [0.00, 100.00]. As a recommendation, when the similarity is greater than or equal to 70, it can be determined that the two faces are of the same person. You can adjust the threshold according to your specific scenario (the FAR at the threshold of 70 is 0.1%, and FAR at the threshold of 80 is 0.01%).
	Sim *float64 `json:"Sim,omitempty" name:"Sim"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Url *string `json:"Url,omitempty" name:"Url"`

	// The 32-bit MD5 checksum of the file
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// The file size
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

// Predefined struct for user
type GenerateReflectSequenceRequestParams struct {
	// The resource URL of the data package generated by the SDK.
	DeviceDataUrl *string `json:"DeviceDataUrl,omitempty" name:"DeviceDataUrl"`

	// The MD5 hash value of the data package generated by the SDK.
	DeviceDataMd5 *string `json:"DeviceDataMd5,omitempty" name:"DeviceDataMd5"`

	// 1 - silent
	// 2 - blinking
	// 3 - light
	// 4 - blinking + light (default)
	SecurityLevel *string `json:"SecurityLevel,omitempty" name:"SecurityLevel"`
}

type GenerateReflectSequenceRequest struct {
	*tchttp.BaseRequest
	
	// The resource URL of the data package generated by the SDK.
	DeviceDataUrl *string `json:"DeviceDataUrl,omitempty" name:"DeviceDataUrl"`

	// The MD5 hash value of the data package generated by the SDK.
	DeviceDataMd5 *string `json:"DeviceDataMd5,omitempty" name:"DeviceDataMd5"`

	// 1 - silent
	// 2 - blinking
	// 3 - light
	// 4 - blinking + light (default)
	SecurityLevel *string `json:"SecurityLevel,omitempty" name:"SecurityLevel"`
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
	ReflectSequenceUrl *string `json:"ReflectSequenceUrl,omitempty" name:"ReflectSequenceUrl"`

	// The MD5 hash value of the light sequence, which is used to check whether the light sequence is altered.
	ReflectSequenceMd5 *string `json:"ReflectSequenceMd5,omitempty" name:"ReflectSequenceMd5"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type GetFaceIdResultIntlRequestParams struct {
	// The ID of the SDK-based liveness detection and face comparison process, which is generated when the `GetFaceIdTokenIntl` API is called.	
	SdkToken *string `json:"SdkToken,omitempty" name:"SdkToken"`
}

type GetFaceIdResultIntlRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the SDK-based liveness detection and face comparison process, which is generated when the `GetFaceIdTokenIntl` API is called.	
	SdkToken *string `json:"SdkToken,omitempty" name:"SdkToken"`
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
	Result *string `json:"Result,omitempty" name:"Result"`

	// The description of the verification result.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The best frame screenshot (in Base64) obtained during the verification.
	BestFrame *string `json:"BestFrame,omitempty" name:"BestFrame"`

	// The video file (Base64) for verification.
	Video *string `json:"Video,omitempty" name:"Video"`

	// The similarity, with a value range of 0-100. A greater value indicates higher similarity. This parameter is returned only in the `compare` (liveness detection and face comparison) mode.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

	// The pass-through parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CheckMode *string `json:"CheckMode,omitempty" name:"CheckMode"`

	// The verification security level. Valid values:
	// `1`: Video-based liveness detection.
	// `2`: Motion-based liveness detection.
	// `3`: Reflection-based liveness detection.
	// `4`: Motion- and reflection-based liveness detection.
	// Default value: `4`.
	SecureLevel *string `json:"SecureLevel,omitempty" name:"SecureLevel"`

	// The photo (in Base64) to compare. This parameter is required when the value of `CheckMode` is `compare`.
	Image *string `json:"Image,omitempty" name:"Image"`

	// The pass-through parameter, which can be omitted if there are no special requirements.
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

type GetFaceIdTokenIntlRequest struct {
	*tchttp.BaseRequest
	
	// The detection mode. Valid values:
	// `liveness`: Liveness detection only.
	// `compare`: Liveness detection and face comparison.
	// Default value: `liveness`.
	CheckMode *string `json:"CheckMode,omitempty" name:"CheckMode"`

	// The verification security level. Valid values:
	// `1`: Video-based liveness detection.
	// `2`: Motion-based liveness detection.
	// `3`: Reflection-based liveness detection.
	// `4`: Motion- and reflection-based liveness detection.
	// Default value: `4`.
	SecureLevel *string `json:"SecureLevel,omitempty" name:"SecureLevel"`

	// The photo (in Base64) to compare. This parameter is required when the value of `CheckMode` is `compare`.
	Image *string `json:"Image,omitempty" name:"Image"`

	// The pass-through parameter, which can be omitted if there are no special requirements.
	Extra *string `json:"Extra,omitempty" name:"Extra"`
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
	SdkToken *string `json:"SdkToken,omitempty" name:"SdkToken"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkToken *string `json:"SdkToken,omitempty" name:"SdkToken"`
}

type GetLivenessResultRequest struct {
	*tchttp.BaseRequest
	
	// The token used to identify an SDK-based verification process.
	SdkToken *string `json:"SdkToken,omitempty" name:"SdkToken"`
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
	Result *string `json:"Result,omitempty" name:"Result"`

	// The description of the final verification result.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The face screenshot.
	BestFrame *FileInfo `json:"BestFrame,omitempty" name:"BestFrame"`

	// The video for the detection.
	Video *FileInfo `json:"Video,omitempty" name:"Video"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SdkToken *string `json:"SdkToken,omitempty" name:"SdkToken"`
}

type GetSdkVerificationResultRequest struct {
	*tchttp.BaseRequest
	
	// The token used to identify an SDK-based verification process.
	SdkToken *string `json:"SdkToken,omitempty" name:"SdkToken"`
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
	Result *string `json:"Result,omitempty" name:"Result"`

	// The verification result description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The charge count.
	ChargeCount *int64 `json:"ChargeCount,omitempty" name:"ChargeCount"`

	// The results of multiple OCR processes (in order). The result of the final process is used as the valid result.
	CardVerifyResults []*CardVerifyResult `json:"CardVerifyResults,omitempty" name:"CardVerifyResults"`

	// The results of multiple liveness detection processes (in order). The result of the final process is used as the valid result.
	CompareResults []*CompareResult `json:"CompareResults,omitempty" name:"CompareResults"`

	// Data passed through in the process of getting the token.
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`
}

type GetWebVerificationResultIntlRequest struct {
	*tchttp.BaseRequest
	
	// The token for the web-based verification, which is generated using the `ApplyWebVerificationBizTokenIntl` API.
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`
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
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// The description of the final verification result.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// The detailed verification result list of this process. Retries are allowed, so a verification process may have several entries of results.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	VerificationDetailList []*VerificationDetail `json:"VerificationDetailList,omitempty" name:"VerificationDetailList"`

	// The Base64-encoded string of the video collected from the video stream. Retries are allowed, and this field returns only the data collected in the last verification. If no video is collected, null is returned.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// The Base64-encoded string of the best face screenshot u200dcollected from the video stream. Retries are allowed, and this field returns only the data collected in the last verification. If no best face screenshot is collected, null is returned.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`
}

type GetWebVerificationResultRequest struct {
	*tchttp.BaseRequest
	
	// The token for the web-based verification, which is generated with the `ApplyWebVerificationToken` API.
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`
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
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// The description of the final verification result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// The temporary URL of the best face screenshot collected from the video stream. It is valid for 10 minutes. Download the image if needed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoBestFrameUrl *string `json:"VideoBestFrameUrl,omitempty" name:"VideoBestFrameUrl"`

	// The MD5 hash value of the best face screenshot collected from the video stream. It can be used to check whether the image content is consistent with the file content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoBestFrameMd5 *string `json:"VideoBestFrameMd5,omitempty" name:"VideoBestFrameMd5"`

	// The details list of this verification process.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VerificationDetailList []*VerificationDetail `json:"VerificationDetailList,omitempty" name:"VerificationDetailList"`

	// The temporary URL of the video collected from the video stream. It is valid for 10 minutes. Download the video if needed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// The MD5 hash value of the video collected from the video stream. It can be used to check whether the video content is consistent with the file content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoMd5 *string `json:"VideoMd5,omitempty" name:"VideoMd5"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type LivenessCompareRequestParams struct {
	// Liveness detection type. Valid values: LIP/ACTION/SILENT.
	// LIP: numeric mode; ACTION: motion mode; SILENT: silent mode. You need to select a mode to input.
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// Base64 string of the image for face comparison.
	// The size of the Base64-encoded image data can be up to 3 MB. JPG and PNG formats are supported.
	// Please use the standard Base64 encoding scheme (with the "=" padding). For the encoding conventions, please see RFC 4648.
	// 
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageBase64` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL of the image for face comparison. The size of the downloaded image after Base64 encoding can be up to 3 MB. JPG and PNG formats are supported.
	// 
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageBase64` will be used.
	// 
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Lip mode: set this parameter to a custom 4-digit verification code.
	// Action mode: set this parameter to a custom action sequence (e.g., `2,1` or `1,2`).
	// Silent mode: do not pass in this parameter.
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`

	// Optional configuration (a JSON string)
	// {
	// "BestFrameNum": 2  // Return multiple best screenshots. Value range: 2−10
	// }
	Optional *string `json:"Optional,omitempty" name:"Optional"`

	// Base64 string of the video for liveness detection.
	// The size of the Base64-encoded video data can be up to 8 MB. MP4, AVI, and FLV formats are supported.
	// Please use the standard Base64 encoding scheme (with the "=" padding). For the encoding conventions, please see RFC 4648.
	// 
	// Either the `VideoUrl` or `VideoBase64` of the video must be provided. If both are provided, only `VideoBase64` will be used.
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// URL of the video for liveness detection. The size of the downloaded video after Base64 encoding can be up to 8 MB. It takes no more than 4 seconds to download. MP4, AVI, and FLV formats are supported.
	// 
	// Either the `VideoUrl` or `VideoBase64` of the video must be provided. If both are provided, only `VideoBase64` will be used.
	// 
	// We recommend you store the video in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`
}

type LivenessCompareRequest struct {
	*tchttp.BaseRequest
	
	// Liveness detection type. Valid values: LIP/ACTION/SILENT.
	// LIP: numeric mode; ACTION: motion mode; SILENT: silent mode. You need to select a mode to input.
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// Base64 string of the image for face comparison.
	// The size of the Base64-encoded image data can be up to 3 MB. JPG and PNG formats are supported.
	// Please use the standard Base64 encoding scheme (with the "=" padding). For the encoding conventions, please see RFC 4648.
	// 
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageBase64` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL of the image for face comparison. The size of the downloaded image after Base64 encoding can be up to 3 MB. JPG and PNG formats are supported.
	// 
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageBase64` will be used.
	// 
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Lip mode: set this parameter to a custom 4-digit verification code.
	// Action mode: set this parameter to a custom action sequence (e.g., `2,1` or `1,2`).
	// Silent mode: do not pass in this parameter.
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`

	// Optional configuration (a JSON string)
	// {
	// "BestFrameNum": 2  // Return multiple best screenshots. Value range: 2−10
	// }
	Optional *string `json:"Optional,omitempty" name:"Optional"`

	// Base64 string of the video for liveness detection.
	// The size of the Base64-encoded video data can be up to 8 MB. MP4, AVI, and FLV formats are supported.
	// Please use the standard Base64 encoding scheme (with the "=" padding). For the encoding conventions, please see RFC 4648.
	// 
	// Either the `VideoUrl` or `VideoBase64` of the video must be provided. If both are provided, only `VideoBase64` will be used.
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// URL of the video for liveness detection. The size of the downloaded video after Base64 encoding can be up to 8 MB. It takes no more than 4 seconds to download. MP4, AVI, and FLV formats are supported.
	// 
	// Either the `VideoUrl` or `VideoBase64` of the video must be provided. If both are provided, only `VideoBase64` will be used.
	// 
	// We recommend you store the video in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`
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
	BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

	// Similarity. Value range: [0.00, 100.00]. As a recommendation, when the similarity is greater than or equal to 70, it can be determined that the two faces are of the same person. You can adjust the threshold according to your specific scenario (the FAR at the threshold of 70 is 0.1%, and FAR at the threshold of 80 is 0.01%).
	Sim *float64 `json:"Sim,omitempty" name:"Sim"`

	// Service error code. `Success` will be returned for success. For error information, please see the `FailedOperation` section in the error code list below.
	Result *string `json:"Result,omitempty" name:"Result"`

	// Service result description.
	Description *string `json:"Description,omitempty" name:"Description"`


	BestFrameList []*string `json:"BestFrameList,omitempty" name:"BestFrameList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type VerificationDetail struct {
	// The final result of this verification. `0` indicates that the person is the same as that in the photo.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// The description of the final verification result.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// The result of this liveness detection process. `0` indicates success.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	LivenessErrorCode *int64 `json:"LivenessErrorCode,omitempty" name:"LivenessErrorCode"`

	// The result description of this liveness detection process.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LivenessErrorMsg *string `json:"LivenessErrorMsg,omitempty" name:"LivenessErrorMsg"`

	// The result of this comparison process. `0` indicates that the person in the best face screenshot collected from the video stream is the same as that in the uploaded image for comparison.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	CompareErrorCode *int64 `json:"CompareErrorCode,omitempty" name:"CompareErrorCode"`

	// The result description of this comparison process.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	CompareErrorMsg *string `json:"CompareErrorMsg,omitempty" name:"CompareErrorMsg"`

	// The timestamp (ms) of this verification process.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReqTimestamp *uint64 `json:"ReqTimestamp,omitempty" name:"ReqTimestamp"`

	// The similarity of the best face screenshot collected from the video stream and the uploaded image for comparison in this verification process. Value range: [0.00, 100.00]. By default, the person in the screenshot is determined to be the same person in the image if the similarity is greater than or equal to 70.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

	// Unique ID of this verification process.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Seq *string `json:"Seq,omitempty" name:"Seq"`
}

// Predefined struct for user
type VideoLivenessCompareRequestParams struct {
	// The URL of the photo for face comparison. The downloaded image after Base64 encoding can be up to 3 MB and must be in JPG or PNG.
	// 
	// The image must be stored in a COS bucket in the region where the FaceID service resides to ensure a higher download speed and better stability. You can generate an image URL by using `CreateUploadUrl` or purchase the COS service.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// The 32-bit MD5 checksum of the image for comparison
	ImageMd5 *string `json:"ImageMd5,omitempty" name:"ImageMd5"`

	// The URL of the video for liveness detection. The downloaded video after Base64 encoding can be up to 8 MB and must be in MP4, AVI, or FLV. It takes no more than 4s to download the video.
	// 
	// The video must be stored in a COS bucket in the region where the FaceID service resides to ensure a higher download speed and better stability. You can generate a video URL by using `CreateUploadUrl` or purchase the COS service.
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// The 32-bit MD5 checksum of the video
	VideoMd5 *string `json:"VideoMd5,omitempty" name:"VideoMd5"`

	// The liveness detection type. Valid values: `LIP`, `ACTION`, and `SILENT`.
	// `LIP`: Numeric mode; `ACTION`: Motion mode; `SILENT`: silent mode. Select one of them.
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// LIP parameter: Pass in a custom 4-digit verification code.
	// ACTION parameter: Pass in a custom action sequence (`2,1` or `1,2`).
	// SILENT parameter: Null.
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`
}

type VideoLivenessCompareRequest struct {
	*tchttp.BaseRequest
	
	// The URL of the photo for face comparison. The downloaded image after Base64 encoding can be up to 3 MB and must be in JPG or PNG.
	// 
	// The image must be stored in a COS bucket in the region where the FaceID service resides to ensure a higher download speed and better stability. You can generate an image URL by using `CreateUploadUrl` or purchase the COS service.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// The 32-bit MD5 checksum of the image for comparison
	ImageMd5 *string `json:"ImageMd5,omitempty" name:"ImageMd5"`

	// The URL of the video for liveness detection. The downloaded video after Base64 encoding can be up to 8 MB and must be in MP4, AVI, or FLV. It takes no more than 4s to download the video.
	// 
	// The video must be stored in a COS bucket in the region where the FaceID service resides to ensure a higher download speed and better stability. You can generate a video URL by using `CreateUploadUrl` or purchase the COS service.
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// The 32-bit MD5 checksum of the video
	VideoMd5 *string `json:"VideoMd5,omitempty" name:"VideoMd5"`

	// The liveness detection type. Valid values: `LIP`, `ACTION`, and `SILENT`.
	// `LIP`: Numeric mode; `ACTION`: Motion mode; `SILENT`: silent mode. Select one of them.
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// LIP parameter: Pass in a custom 4-digit verification code.
	// ACTION parameter: Pass in a custom action sequence (`2,1` or `1,2`).
	// SILENT parameter: Null.
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`
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
	Sim *float64 `json:"Sim,omitempty" name:"Sim"`

	// The service error code. `Success` will be returned for success. For error information, see the `FailedOperation` section in the error code list below.
	Result *string `json:"Result,omitempty" name:"Result"`

	// The service result description
	Description *string `json:"Description,omitempty" name:"Description"`

	// The best video screenshot after successful verification
	// Note: This field may return null, indicating that no valid values can be obtained.
	BestFrame *FileInfo `json:"BestFrame,omitempty" name:"BestFrame"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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