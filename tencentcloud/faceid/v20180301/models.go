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

	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`


	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`


	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`


	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`


	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`


	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`


	Age *string `json:"Age,omitnil,omitempty" name:"Age"`


	IssuedCountry *string `json:"IssuedCountry,omitnil,omitempty" name:"IssuedCountry"`


	Field1 *string `json:"Field1,omitnil,omitempty" name:"Field1"`


	Field2 *string `json:"Field2,omitnil,omitempty" name:"Field2"`
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