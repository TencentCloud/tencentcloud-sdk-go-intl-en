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

package v20181119

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type BankCardOCRRequestParams struct {
	// Base64-encoded value of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 7 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Whether to return the bank card image data after preprocessing (precise cropping and alignment). Default value: `false`
	RetBorderCutImage *bool `json:"RetBorderCutImage,omitempty" name:"RetBorderCutImage"`

	// Whether to return the card number image data after slicing. Default value: `false`
	RetCardNoImage *bool `json:"RetCardNoImage,omitempty" name:"RetCardNoImage"`

	// Whether to enable photocopy check. If the input image is a bank card photocopy, an alarm will be returned. Default value: `false`
	EnableCopyCheck *bool `json:"EnableCopyCheck,omitempty" name:"EnableCopyCheck"`

	// Whether to enable photograph check. If the input image is a bank card photograph, an alarm will be returned. Default value: `false`
	EnableReshootCheck *bool `json:"EnableReshootCheck,omitempty" name:"EnableReshootCheck"`

	// Whether to enable obscured border check. If the input image is a bank card with obscured border, an alarm will be returned. Default value: `false`
	EnableBorderCheck *bool `json:"EnableBorderCheck,omitempty" name:"EnableBorderCheck"`

	// Whether to return the image quality value, which measures how clear an image is. Default value: `false`
	EnableQualityValue *bool `json:"EnableQualityValue,omitempty" name:"EnableQualityValue"`
}

type BankCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 7 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Whether to return the bank card image data after preprocessing (precise cropping and alignment). Default value: `false`
	RetBorderCutImage *bool `json:"RetBorderCutImage,omitempty" name:"RetBorderCutImage"`

	// Whether to return the card number image data after slicing. Default value: `false`
	RetCardNoImage *bool `json:"RetCardNoImage,omitempty" name:"RetCardNoImage"`

	// Whether to enable photocopy check. If the input image is a bank card photocopy, an alarm will be returned. Default value: `false`
	EnableCopyCheck *bool `json:"EnableCopyCheck,omitempty" name:"EnableCopyCheck"`

	// Whether to enable photograph check. If the input image is a bank card photograph, an alarm will be returned. Default value: `false`
	EnableReshootCheck *bool `json:"EnableReshootCheck,omitempty" name:"EnableReshootCheck"`

	// Whether to enable obscured border check. If the input image is a bank card with obscured border, an alarm will be returned. Default value: `false`
	EnableBorderCheck *bool `json:"EnableBorderCheck,omitempty" name:"EnableBorderCheck"`

	// Whether to return the image quality value, which measures how clear an image is. Default value: `false`
	EnableQualityValue *bool `json:"EnableQualityValue,omitempty" name:"EnableQualityValue"`
}

func (r *BankCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "RetBorderCutImage")
	delete(f, "RetCardNoImage")
	delete(f, "EnableCopyCheck")
	delete(f, "EnableReshootCheck")
	delete(f, "EnableBorderCheck")
	delete(f, "EnableQualityValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BankCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BankCardOCRResponseParams struct {
	// Card number
	CardNo *string `json:"CardNo,omitempty" name:"CardNo"`

	// Bank information
	BankInfo *string `json:"BankInfo,omitempty" name:"BankInfo"`

	// Expiration date. Format: 07/2023
	ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

	// Card type
	CardType *string `json:"CardType,omitempty" name:"CardType"`

	// Card name
	CardName *string `json:"CardName,omitempty" name:"CardName"`

	// Sliced image data
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BorderCutImage *string `json:"BorderCutImage,omitempty" name:"BorderCutImage"`

	// Card number image data
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CardNoImage *string `json:"CardNoImage,omitempty" name:"CardNoImage"`

	// Warning code:
	// -9110: the bank card date is invalid. 
	// -9111: the bank card border is incomplete. 
	// -9112: the bank card image is reflective.
	// -9113: the bank card image is a photocopy.
	// -9114: the bank card image is a photograph.
	// Multiple warning codes may be returned at a time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WarningCode []*int64 `json:"WarningCode,omitempty" name:"WarningCode"`

	// Image quality value, which is returned when `EnableQualityValue` is set to `true`. The smaller the value, the less clear the image is. Value range: 0−100 (a threshold greater than or equal to 50 is recommended.)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	QualityValue *int64 `json:"QualityValue,omitempty" name:"QualityValue"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BankCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *BankCardOCRResponseParams `json:"Response"`
}

func (r *BankCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Coord struct {
	// Horizontal coordinate
	X *int64 `json:"X,omitempty" name:"X"`

	// Vertical coordinate
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type DetectedWordCoordPoint struct {
	// Coordinates of a word’s four corners in a clockwise order on the input image, starting from the upper-left corner
	WordCoordinate []*Coord `json:"WordCoordinate,omitempty" name:"WordCoordinate"`
}

type DetectedWords struct {
	// Confidence. Value range: 0–100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// A possible character
	Character *string `json:"Character,omitempty" name:"Character"`
}

// Predefined struct for user
type GeneralAccurateOCRRequestParams struct {
	// Base64-encoded value of image.
	// The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitempty" name:"IsWords"`

	// Whether to slice the input image to enhance the recognition effects for scenarios where the whole image is big, but the size of a single character is small (e.g., test papers). This feature is disabled by default.
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitempty" name:"EnableDetectSplit"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// Number of a PDF page that needs to be recognized. Currently, only one single page can be recognized. This parameter takes effect only if a PDF file is uploaded and `IsPdf` is set to `true`. Default value: `1`
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`
}

type GeneralAccurateOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image.
	// The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitempty" name:"IsWords"`

	// Whether to slice the input image to enhance the recognition effects for scenarios where the whole image is big, but the size of a single character is small (e.g., test papers). This feature is disabled by default.
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitempty" name:"EnableDetectSplit"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// Number of a PDF page that needs to be recognized. Currently, only one single page can be recognized. This parameter takes effect only if a PDF file is uploaded and `IsPdf` is set to `true`. Default value: `1`
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`
}

func (r *GeneralAccurateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralAccurateOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "IsWords")
	delete(f, "EnableDetectSplit")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneralAccurateOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GeneralAccurateOCRResponseParams struct {
	// Information on recognized text, including the text line content, confidence, text line coordinates, and text line coordinates after rotation correction. For more information, please click the link on the left.
	TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections"`

	// Image rotation angle in degrees. 0°: The horizontal direction of the text on the image; a positive value: rotate clockwise; a negative value: rotate counterclockwise.
	Angel *float64 `json:"Angel,omitempty" name:"Angel"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GeneralAccurateOCRResponse struct {
	*tchttp.BaseResponse
	Response *GeneralAccurateOCRResponseParams `json:"Response"`
}

func (r *GeneralAccurateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralAccurateOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GeneralBasicOCRRequestParams struct {
	// Base64-encoded value of image/PDF.
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image/PDF. (This field is not supported outside Chinese mainland)
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Reserved field.
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// Language to recognize
	// The language can be automatically recognized or manually specified. Chinese-English mix (`zh`) is selected by default. Mixed characters in English and each supported language can be recognized together.
	// Valid values:
	// `zh`: Chinese-English mix
	// `zh_rare`: supports letters, digits, rare Chinese characters, Traditional Chinese characters, special characters, etc.
	// `auto`
	// `mix`: language mix
	// `jap`: Japanese
	// `kor`: Korean
	// `spa`: Spanish
	// `fre`: French
	// `ger`: German
	// `por`: Portuguese
	// `vie`: Vietnamese
	// `may`: Malay
	// `rus`: Russian
	// `ita`: Italian
	// `hol`: Dutch
	// `swe`: Swedish
	// `fin`: Finnish
	// `dan`: Danish
	// `nor`: Norwegian
	// `hun`: Hungarian
	// `tha`: Thai
	// `hi`: Hindi
	// `ara`: Arabic
	LanguageType *string `json:"LanguageType,omitempty" name:"LanguageType"`

	// Whether to enable PDF recognition. Default value: false. After this feature is enabled, both images and PDF files can be recognized at the same time.
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// Page number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of the `IsPdf` parameter is `true`. Default value: 1.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitempty" name:"IsWords"`
}

type GeneralBasicOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image/PDF.
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image/PDF. (This field is not supported outside Chinese mainland)
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Reserved field.
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// Language to recognize
	// The language can be automatically recognized or manually specified. Chinese-English mix (`zh`) is selected by default. Mixed characters in English and each supported language can be recognized together.
	// Valid values:
	// `zh`: Chinese-English mix
	// `zh_rare`: supports letters, digits, rare Chinese characters, Traditional Chinese characters, special characters, etc.
	// `auto`
	// `mix`: language mix
	// `jap`: Japanese
	// `kor`: Korean
	// `spa`: Spanish
	// `fre`: French
	// `ger`: German
	// `por`: Portuguese
	// `vie`: Vietnamese
	// `may`: Malay
	// `rus`: Russian
	// `ita`: Italian
	// `hol`: Dutch
	// `swe`: Swedish
	// `fin`: Finnish
	// `dan`: Danish
	// `nor`: Norwegian
	// `hun`: Hungarian
	// `tha`: Thai
	// `hi`: Hindi
	// `ara`: Arabic
	LanguageType *string `json:"LanguageType,omitempty" name:"LanguageType"`

	// Whether to enable PDF recognition. Default value: false. After this feature is enabled, both images and PDF files can be recognized at the same time.
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// Page number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of the `IsPdf` parameter is `true`. Default value: 1.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitempty" name:"IsWords"`
}

func (r *GeneralBasicOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralBasicOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Scene")
	delete(f, "LanguageType")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	delete(f, "IsWords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneralBasicOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GeneralBasicOCRResponseParams struct {
	// Information of recognized text, including the text line content, confidence, text line coordinates, and text line coordinates after rotation correction. For more information, please click the link on the left.
	TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections"`

	// Detected language. For more information on the supported languages, please see the description of the `LanguageType` input parameter.
	Language *string `json:"Language,omitempty" name:"Language"`

	// Image rotation angle in degrees. 0°: The horizontal direction of the text on the image; a positive value: rotate clockwise; a negative value: rotate counterclockwise.
	Angel *float64 `json:"Angel,omitempty" name:"Angel"`

	// Total number of PDF pages to be returned if the image is a PDF. Default value: 0.
	PdfPageSize *int64 `json:"PdfPageSize,omitempty" name:"PdfPageSize"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GeneralBasicOCRResponse struct {
	*tchttp.BaseResponse
	Response *GeneralBasicOCRResponseParams `json:"Response"`
}

func (r *GeneralBasicOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralBasicOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HKIDCardOCRRequestParams struct {
	// Whether to check for authenticity.
	DetectFake *bool `json:"DetectFake,omitempty" name:"DetectFake"`

	// Whether to return identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitempty" name:"ReturnHeadImage"`

	// Base64 string of the image
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported yet.
	// Supported image size: The downloaded image cannot exceed 7 MB after being Base64-encoded, and it cannot take longer than 3 seconds to download the image.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

type HKIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// Whether to check for authenticity.
	DetectFake *bool `json:"DetectFake,omitempty" name:"DetectFake"`

	// Whether to return identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitempty" name:"ReturnHeadImage"`

	// Base64 string of the image
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported yet.
	// Supported image size: The downloaded image cannot exceed 7 MB after being Base64-encoded, and it cannot take longer than 3 seconds to download the image.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *HKIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HKIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DetectFake")
	delete(f, "ReturnHeadImage")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HKIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HKIDCardOCRResponseParams struct {
	// Name in Chinese
	CnName *string `json:"CnName,omitempty" name:"CnName"`

	// Name in English
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// Telecode for the name in Chinese
	TelexCode *string `json:"TelexCode,omitempty" name:"TelexCode"`

	// Gender. Valid values: Male, Female
	Sex *string `json:"Sex,omitempty" name:"Sex"`

	// Date of birth
	Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

	// Permanent identity card.
	// 0: non-permanent;
	// 1: permanent;
	// -1: unknown.
	Permanent *int64 `json:"Permanent,omitempty" name:"Permanent"`

	// Identity card number
	IdNum *string `json:"IdNum,omitempty" name:"IdNum"`

	// Document symbol, i.e., the symbol under the date of birth, such as "***AZ"
	Symbol *string `json:"Symbol,omitempty" name:"Symbol"`

	// First issue date
	FirstIssueDate *string `json:"FirstIssueDate,omitempty" name:"FirstIssueDate"`

	// Last receipt date
	CurrentIssueDate *string `json:"CurrentIssueDate,omitempty" name:"CurrentIssueDate"`

	// Authenticity check.
	// 0: unable to judge (because the image is blurred, incomplete, reflective, too dark, etc.);
	// 1: forged;
	// 2: authentic.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FakeDetectResult *int64 `json:"FakeDetectResult,omitempty" name:"FakeDetectResult"`

	// Base64-encoded identity photo
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeadImage *string `json:"HeadImage,omitempty" name:"HeadImage"`

	// Multiple alarm codes. If the ID card is spoofed, photocopied, or photoshopped, the corresponding alarm code will be returned.
	// -9102: Alarm for photocopied document
	// -9103: Alarm for spoofed document
	// -9104: Alarm for photoshopped document
	WarningCode []*int64 `json:"WarningCode,omitempty" name:"WarningCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type HKIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *HKIDCardOCRResponseParams `json:"Response"`
}

func (r *HKIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HKIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ItemCoord struct {
	// X-coordinate of top-left point.
	X *int64 `json:"X,omitempty" name:"X"`

	// Y-coordinate of top-left point.
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// Width
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Height
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

// Predefined struct for user
type MLIDCardOCRRequestParams struct {
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// The URL of an image. (This field is not available outside the Chinese mainland.)
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Whether to return an image. Default value: `false`.
	RetImage *bool `json:"RetImage,omitempty" name:"RetImage"`
}

type MLIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// The URL of an image. (This field is not available outside the Chinese mainland.)
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Whether to return an image. Default value: `false`.
	RetImage *bool `json:"RetImage,omitempty" name:"RetImage"`
}

func (r *MLIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MLIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "RetImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MLIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MLIDCardOCRResponseParams struct {
	// ID number
	ID *string `json:"ID,omitempty" name:"ID"`

	// Full name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Address
	Address *string `json:"Address,omitempty" name:"Address"`

	// Gender
	Sex *string `json:"Sex,omitempty" name:"Sex"`

	// Alarm codes
	// -9103 Alarm for photographed certificate
	// -9102 Alarm for photocopied certificate
	// -9106 Alarm for covered certificate
	// -9107 Alarm for blurry image
	Warn []*int64 `json:"Warn,omitempty" name:"Warn"`

	// Identity photo
	Image *string `json:"Image,omitempty" name:"Image"`

	// This is an extended field, 
	// with the confidence of a field recognition result returned in the following format.
	// {
	//   Field name:{
	//     Confidence:0.9999
	//   }
	// }
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// Certificate type
	// MyKad  ID card
	// MyPR    Permanent resident card
	// MyTentera   Military identity card
	// MyKAS    Temporary ID card
	// POLIS  Police card
	// IKAD   Work permit
	// MyKid   Kid card
	Type *string `json:"Type,omitempty" name:"Type"`

	// Date of birth. This field is available only for work permits (i-Kad) and ID cards (MyKad).
	Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MLIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *MLIDCardOCRResponseParams `json:"Response"`
}

func (r *MLIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MLIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MLIDPassportOCRRequestParams struct {
	// Base64-encoded value of image. The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 500x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies more than 2/3 area of the image.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// Whether to return an image. Default value: false.
	RetImage *bool `json:"RetImage,omitempty" name:"RetImage"`
}

type MLIDPassportOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image. The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 500x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies more than 2/3 area of the image.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// Whether to return an image. Default value: false.
	RetImage *bool `json:"RetImage,omitempty" name:"RetImage"`
}

func (r *MLIDPassportOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MLIDPassportOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "RetImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MLIDPassportOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MLIDPassportOCRResponseParams struct {
	// Passport ID
	ID *string `json:"ID,omitempty" name:"ID"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Date of birth
	DateOfBirth *string `json:"DateOfBirth,omitempty" name:"DateOfBirth"`

	// Gender (F: female, M: male)
	Sex *string `json:"Sex,omitempty" name:"Sex"`

	// Expiration date
	DateOfExpiration *string `json:"DateOfExpiration,omitempty" name:"DateOfExpiration"`

	// Issuing country
	IssuingCountry *string `json:"IssuingCountry,omitempty" name:"IssuingCountry"`

	// Country/region code
	Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

	// Alarm codes
	// -9103 Alarm for spoofed document
	// -9102 Alarm for photocopied document (including black & white and color ones)
	// -9106 Alarm for covered card
	Warn []*int64 `json:"Warn,omitempty" name:"Warn"`

	// Identity photo
	Image *string `json:"Image,omitempty" name:"Image"`

	// Extended field:
	// {
	//     ID:{
	//         Confidence:0.9999
	//     },
	//     Name:{
	//         Confidence:0.9996
	//     }
	// }
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// The first row of the machine-readable zone (MRZ) at the bottom
	CodeSet *string `json:"CodeSet,omitempty" name:"CodeSet"`

	// The second row of the MRZ at the bottom
	CodeCrc *string `json:"CodeCrc,omitempty" name:"CodeCrc"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MLIDPassportOCRResponse struct {
	*tchttp.BaseResponse
	Response *MLIDPassportOCRResponseParams `json:"Response"`
}

func (r *MLIDPassportOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MLIDPassportOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TableOCRRequestParams struct {
	// Base64-encoded value of image.
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported at present.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

type TableOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image.
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported at present.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *TableOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TableOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TableOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TableOCRResponseParams struct {
	// Recognized text. For more information, please click the link on the left
	TextDetections []*TextTable `json:"TextDetections,omitempty" name:"TextDetections"`

	// Base64-encoded Excel data.
	Data *string `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TableOCRResponse struct {
	*tchttp.BaseResponse
	Response *TableOCRResponseParams `json:"Response"`
}

func (r *TableOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TableOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextDetection struct {
	// Recognized text line content.
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// Confidence. Value range: 0–100.
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// Text line coordinates, which are represented as 4 vertex coordinates.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon"`

	// Extended field.
	// The paragraph information `Parag` returned by the `GeneralBasicOcr` API contains `ParagNo`.
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// Pixel coordinates of the text line in the image after rotation correction, which is in the format of `(X-coordinate of top-left point, Y-coordinate of top-left point, width, height)`.
	ItemPolygon *ItemCoord `json:"ItemPolygon,omitempty" name:"ItemPolygon"`

	// Information about a character, including the character itself and its confidence. Supported APIs: `GeneralBasicOCR`, `GeneralAccurateOCR`
	Words []*DetectedWords `json:"Words,omitempty" name:"Words"`

	// Coordinates of a word’s four corners on the input image. Supported APIs: `GeneralBasicOCR`, `GeneralAccurateOCR`
	WordCoordPoint []*DetectedWordCoordPoint `json:"WordCoordPoint,omitempty" name:"WordCoordPoint"`
}

type TextTable struct {
	// Column index of the top-left corner of the cell.
	ColTl *int64 `json:"ColTl,omitempty" name:"ColTl"`

	// Row index of the top-left corner of the cell.
	RowTl *int64 `json:"RowTl,omitempty" name:"RowTl"`

	// Column index of the bottom-right corner of the cell.
	ColBr *int64 `json:"ColBr,omitempty" name:"ColBr"`

	// Row index of the bottom-right corner of the cell.
	RowBr *int64 `json:"RowBr,omitempty" name:"RowBr"`

	// Cell text
	Text *string `json:"Text,omitempty" name:"Text"`

	// Cell type. Valid values: body, header, footer
	Type *string `json:"Type,omitempty" name:"Type"`

	// Confidence. Value range: 0–100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// Text line coordinates, which are represented as 4 vertex coordinates.
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon"`

	// Extended field
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}