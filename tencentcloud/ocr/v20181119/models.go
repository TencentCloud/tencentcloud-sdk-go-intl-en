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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

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
}

func (r *BankCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCardOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BankCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		WarningCode []*int64 `json:"WarningCode,omitempty" name:"WarningCode" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BankCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCardOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Coord struct {

	// Horizontal coordinate
	X *int64 `json:"X,omitempty" name:"X"`

	// Vertical coordinate
	Y *int64 `json:"Y,omitempty" name:"Y"`
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
}

func (r *GeneralAccurateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralAccurateOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralAccurateOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information on recognized text, including the text line content, confidence, text line coordinates, and text line coordinates after rotation correction. For more information, please click the link on the left.
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// Image rotation angle in degrees. 0° indicates horizontal text, a positive value indicates clockwise rotation, and a negative value indicates anticlockwise rotation.
		Angel *float64 `json:"Angel,omitempty" name:"Angel"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralAccurateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralAccurateOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// Language to be recognized.
	// The language can be automatically recognized or manually specified. Chinese-English mix (`zh`) is selected by default. Mixed characters in English and each supported language can be recognized together.
	// Valid values:
	// zh\auto\jap\kor\
	// spa\fre\ger\por\
	// vie\may\rus\ita\
	// hol\swe\fin\dan\
	// nor\hun\tha\lat\ara
	// Value meanings:
	// Chinese-English mix, automatic recognition, Japanese, Korean,
	// Spanish, French, German, Portuguese,
	// Vietnamese, Malay, Russian, Italian,
	// Dutch, Swedish, Finnish, Danish,
	// Norwegian, Hungarian, Thai, Latin,
	// Arabic.
	LanguageType *string `json:"LanguageType,omitempty" name:"LanguageType"`

	// Whether to enable PDF recognition. Default value: false. After this feature is enabled, both images and PDF files can be recognized at the same time.
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// Page number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of the `IsPdf` parameter is `true`. Default value: 1.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`
}

func (r *GeneralBasicOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralBasicOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralBasicOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of recognized text, including the text line content, confidence, text line coordinates, and text line coordinates after rotation correction. For more information, please click the link on the left.
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// Detected language. For more information on the supported languages, please see the description of the `LanguageType` input parameter.
		Language *string `json:"Language,omitempty" name:"Language"`

		// Image rotation angle in degrees. 0° indicates horizontal text, a positive value indicates clockwise rotation, and a negative value indicates anticlockwise rotation. For more information, please see <a href="https://intl.cloud.tencent.com/document/product/866/45139?from_cn_redirect=1">How to Correct Tilted Text</a>.
		Angel *float64 `json:"Angel,omitempty" name:"Angel"`

		// Total number of PDF pages to be returned if the image is a PDF. Default value: 0.
		PdfPageSize *int64 `json:"PdfPageSize,omitempty" name:"PdfPageSize"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralBasicOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralBasicOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HKIDCardOCRRequest struct {
	*tchttp.BaseRequest

	// Whether to check for authenticity.
	DetectFake *bool `json:"DetectFake,omitempty" name:"DetectFake"`

	// Whether to return identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitempty" name:"ReturnHeadImage"`

	// Base64-encoded value of image.
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
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

func (r *HKIDCardOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HKIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// Multiple alarm codes. If the ID card is spoofed, photocopied, or doctored, the corresponding alarm code will be returned.
	// -9102: alarm for photocopied document
	// -9103: alarm for spoofed document
	// -9104: alarm for doctored document
	// -9105: alarm for forged document
		WarningCode []*int64 `json:"WarningCode,omitempty" name:"WarningCode" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HKIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type MLIDCardOCRRequest struct {
	*tchttp.BaseRequest

	// Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported at present.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of an image. (This field is not supported outside Mainland China)
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported at present.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// It is recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Whether to return an image
	RetImage *bool `json:"RetImage,omitempty" name:"RetImage"`
}

func (r *MLIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MLIDCardOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MLIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Identity card number
		ID *string `json:"ID,omitempty" name:"ID"`

		// Name
		Name *string `json:"Name,omitempty" name:"Name"`

		// Address
		Address *string `json:"Address,omitempty" name:"Address"`

		// Gender
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// Alarm code
	// -9103	Alarm for photographed document
	// -9102	Alarm for photocopied document
	// -9106       Alarm for covered card
		Warn []*int64 `json:"Warn,omitempty" name:"Warn" list`

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

		// Certificate types
	// MyKad: Malaysian Identity Card
	// MyPR: Malaysia Permanent Resident Identity Card
	// MyTentera: Malaysian Armed Forces Identity Card
	// MyKAS: Malaysian Temporary Resident Identity Card
	// POLIS: Royal Malaysia Police Identity Card
	// IKAD: Malaysia Temporary Employment Visit Pass
		Type *string `json:"Type,omitempty" name:"Type"`

		// Date of birth (currently, this field is only supported for IKAD).
		Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MLIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MLIDCardOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *MLIDPassportOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MLIDPassportOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// Nationality
		Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

		// Alarm code
	// -9103 Alarm for spoofed card
	// -9102 Alarm for photocopied card
	// -9106 Alarm for covered card
		Warn []*int64 `json:"Warn,omitempty" name:"Warn" list`

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

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MLIDPassportOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MLIDPassportOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *TableOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TableOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recognized text. For more information, please click the link on the left
		TextDetections []*TextTable `json:"TextDetections,omitempty" name:"TextDetections" list`

		// Base64-encoded Excel data.
		Data *string `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TableOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// Extended field.
	// The paragraph information `Parag` returned by the `GeneralBasicOcr` API contains `ParagNo`.
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// Pixel coordinates of the text line in the image after rotation correction, which is in the format of `(X-coordinate of top-left point, Y-coordinate of top-left point, width, height)`.
	ItemPolygon *ItemCoord `json:"ItemPolygon,omitempty" name:"ItemPolygon"`
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
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// Extended field
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}
