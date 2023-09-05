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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AirTransport struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// E-ticket No.
	Number *string `json:"Number,omitnil" name:"Number"`

	// Check code
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// Serial number
	SerialNumber *string `json:"SerialNumber,omitnil" name:"SerialNumber"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Agent code
	AgentCode *string `json:"AgentCode,omitnil" name:"AgentCode"`

	// First line of the agent code
	AgentCodeFirst *string `json:"AgentCodeFirst,omitnil" name:"AgentCodeFirst"`

	// Second line of the agent code
	AgentCodeSecond *string `json:"AgentCodeSecond,omitnil" name:"AgentCodeSecond"`

	// Name
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// ID card number
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// Fare
	Fare *string `json:"Fare,omitnil" name:"Fare"`

	// Tax
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// Fuel surcharge
	FuelSurcharge *string `json:"FuelSurcharge,omitnil" name:"FuelSurcharge"`

	// Aviation Development Fund
	AirDevelopmentFund *string `json:"AirDevelopmentFund,omitnil" name:"AirDevelopmentFund"`

	// Insurance
	Insurance *string `json:"Insurance,omitnil" name:"Insurance"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Domestic or international tag
	DomesticInternationalTag *string `json:"DomesticInternationalTag,omitnil" name:"DomesticInternationalTag"`

	// Not-valid-before date
	DateStart *string `json:"DateStart,omitnil" name:"DateStart"`

	// Not-valid-after date
	DateEnd *string `json:"DateEnd,omitnil" name:"DateEnd"`

	// Endorsements/Restrictions
	Endorsement *string `json:"Endorsement,omitnil" name:"Endorsement"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Items
	FlightItems []*FlightItem `json:"FlightItems,omitnil" name:"FlightItems"`
}

// Predefined struct for user
type BankCardOCRRequestParams struct {
	// Base64-encoded value of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 7 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the bank card image data after preprocessing (precise cropping and alignment). Default value: `false`
	RetBorderCutImage *bool `json:"RetBorderCutImage,omitnil" name:"RetBorderCutImage"`

	// Whether to return the card number image data after slicing. Default value: `false`
	RetCardNoImage *bool `json:"RetCardNoImage,omitnil" name:"RetCardNoImage"`

	// Whether to enable photocopy check. If the input image is a bank card photocopy, an alarm will be returned. Default value: `false`
	EnableCopyCheck *bool `json:"EnableCopyCheck,omitnil" name:"EnableCopyCheck"`

	// Whether to enable photograph check. If the input image is a bank card photograph, an alarm will be returned. Default value: `false`
	EnableReshootCheck *bool `json:"EnableReshootCheck,omitnil" name:"EnableReshootCheck"`

	// Whether to enable obscured border check. If the input image is a bank card with obscured border, an alarm will be returned. Default value: `false`
	EnableBorderCheck *bool `json:"EnableBorderCheck,omitnil" name:"EnableBorderCheck"`

	// Whether to return the image quality value, which measures how clear an image is. Default value: `false`
	EnableQualityValue *bool `json:"EnableQualityValue,omitnil" name:"EnableQualityValue"`
}

type BankCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 7 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the bank card image data after preprocessing (precise cropping and alignment). Default value: `false`
	RetBorderCutImage *bool `json:"RetBorderCutImage,omitnil" name:"RetBorderCutImage"`

	// Whether to return the card number image data after slicing. Default value: `false`
	RetCardNoImage *bool `json:"RetCardNoImage,omitnil" name:"RetCardNoImage"`

	// Whether to enable photocopy check. If the input image is a bank card photocopy, an alarm will be returned. Default value: `false`
	EnableCopyCheck *bool `json:"EnableCopyCheck,omitnil" name:"EnableCopyCheck"`

	// Whether to enable photograph check. If the input image is a bank card photograph, an alarm will be returned. Default value: `false`
	EnableReshootCheck *bool `json:"EnableReshootCheck,omitnil" name:"EnableReshootCheck"`

	// Whether to enable obscured border check. If the input image is a bank card with obscured border, an alarm will be returned. Default value: `false`
	EnableBorderCheck *bool `json:"EnableBorderCheck,omitnil" name:"EnableBorderCheck"`

	// Whether to return the image quality value, which measures how clear an image is. Default value: `false`
	EnableQualityValue *bool `json:"EnableQualityValue,omitnil" name:"EnableQualityValue"`
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
	CardNo *string `json:"CardNo,omitnil" name:"CardNo"`

	// Bank information
	BankInfo *string `json:"BankInfo,omitnil" name:"BankInfo"`

	// Expiration date. Format: 07/2023
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// Card type
	CardType *string `json:"CardType,omitnil" name:"CardType"`

	// Card name
	CardName *string `json:"CardName,omitnil" name:"CardName"`

	// Sliced image data
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BorderCutImage *string `json:"BorderCutImage,omitnil" name:"BorderCutImage"`

	// Card number image data
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CardNoImage *string `json:"CardNoImage,omitnil" name:"CardNoImage"`

	// Warning code:
	// -9110: the bank card date is invalid. 
	// -9111: the bank card border is incomplete. 
	// -9112: the bank card image is reflective.
	// -9113: the bank card image is a photocopy.
	// -9114: the bank card image is a photograph.
	// Multiple warning codes may be returned at a time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WarningCode []*int64 `json:"WarningCode,omitnil" name:"WarningCode"`

	// Image quality value, which is returned when `EnableQualityValue` is set to `true`. The smaller the value, the less clear the image is. Value range: 0−100 (a threshold greater than or equal to 50 is recommended.)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	QualityValue *int64 `json:"QualityValue,omitnil" name:"QualityValue"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type BusInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Departure time
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// Departure date
	DateGetOn *string `json:"DateGetOn,omitnil" name:"DateGetOn"`

	// Departure station
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// Destination
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// Fare
	Total *string `json:"Total,omitnil" name:"Total"`

	// Name
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Consumption type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// ID card number
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// Province
	Province *string `json:"Province,omitnil" name:"Province"`

	// City
	City *string `json:"City,omitnil" name:"City"`

	// Departure place
	PlaceGetOn *string `json:"PlaceGetOn,omitnil" name:"PlaceGetOn"`

	// Check-in gate
	GateNumber *string `json:"GateNumber,omitnil" name:"GateNumber"`

	// Fare category
	TicketType *string `json:"TicketType,omitnil" name:"TicketType"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// Seat No.
	SeatNumber *string `json:"SeatNumber,omitnil" name:"SeatNumber"`

	// Fleet number
	TrainNumber *string `json:"TrainNumber,omitnil" name:"TrainNumber"`
}

type Coord struct {
	// Horizontal coordinate
	X *int64 `json:"X,omitnil" name:"X"`

	// Vertical coordinate
	Y *int64 `json:"Y,omitnil" name:"Y"`
}

type DetectedWordCoordPoint struct {
	// Coordinates of a word’s four corners in a clockwise order on the input image, starting from the upper-left corner
	WordCoordinate []*Coord `json:"WordCoordinate,omitnil" name:"WordCoordinate"`
}

type DetectedWords struct {
	// Confidence. Value range: 0–100
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// A possible character
	Character *string `json:"Character,omitnil" name:"Character"`
}

type FlightItem struct {
	// Departure terminal
	TerminalGetOn *string `json:"TerminalGetOn,omitnil" name:"TerminalGetOn"`

	// Arrival terminal
	TerminalGetOff *string `json:"TerminalGetOff,omitnil" name:"TerminalGetOff"`

	// Carrier
	Carrier *string `json:"Carrier,omitnil" name:"Carrier"`

	// Flight number
	FlightNumber *string `json:"FlightNumber,omitnil" name:"FlightNumber"`

	// Class
	Seat *string `json:"Seat,omitnil" name:"Seat"`

	// Departure date
	DateGetOn *string `json:"DateGetOn,omitnil" name:"DateGetOn"`

	// Departure time
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// Departure city
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// Arrival city
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// Baggage allowance
	Allow *string `json:"Allow,omitnil" name:"Allow"`

	// Fare category
	FareBasis *string `json:"FareBasis,omitnil" name:"FareBasis"`
}

// Predefined struct for user
type GeneralAccurateOCRRequestParams struct {
	// Base64-encoded value of image.
	// The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitnil" name:"IsWords"`

	// Whether to slice the input image to enhance the recognition effects for scenarios where the whole image is big, but the size of a single character is small (e.g., test papers). This feature is disabled by default.
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitnil" name:"EnableDetectSplit"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// Number of a PDF page that needs to be recognized. Currently, only one single page can be recognized. This parameter takes effect only if a PDF file is uploaded and `IsPdf` is set to `true`. Default value: `1`
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type GeneralAccurateOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image.
	// The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitnil" name:"IsWords"`

	// Whether to slice the input image to enhance the recognition effects for scenarios where the whole image is big, but the size of a single character is small (e.g., test papers). This feature is disabled by default.
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitnil" name:"EnableDetectSplit"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// Number of a PDF page that needs to be recognized. Currently, only one single page can be recognized. This parameter takes effect only if a PDF file is uploaded and `IsPdf` is set to `true`. Default value: `1`
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	TextDetections []*TextDetection `json:"TextDetections,omitnil" name:"TextDetections"`

	// Image rotation angle in degrees. 0°: The horizontal direction of the text on the image; a positive value: rotate clockwise; a negative value: rotate counterclockwise.
	Angel *float64 `json:"Angel,omitnil" name:"Angel"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image/PDF. (This field is not supported outside Chinese mainland)
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Reserved field.
	Scene *string `json:"Scene,omitnil" name:"Scene"`

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
	LanguageType *string `json:"LanguageType,omitnil" name:"LanguageType"`

	// Whether to enable PDF recognition. Default value: false. After this feature is enabled, both images and PDF files can be recognized at the same time.
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// Page number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of the `IsPdf` parameter is `true`. Default value: 1.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitnil" name:"IsWords"`
}

type GeneralBasicOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image/PDF.
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image/PDF. (This field is not supported outside Chinese mainland)
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Reserved field.
	Scene *string `json:"Scene,omitnil" name:"Scene"`

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
	LanguageType *string `json:"LanguageType,omitnil" name:"LanguageType"`

	// Whether to enable PDF recognition. Default value: false. After this feature is enabled, both images and PDF files can be recognized at the same time.
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// Page number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of the `IsPdf` parameter is `true`. Default value: 1.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitnil" name:"IsWords"`
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
	TextDetections []*TextDetection `json:"TextDetections,omitnil" name:"TextDetections"`

	// Detected language. For more information on the supported languages, please see the description of the `LanguageType` input parameter.
	Language *string `json:"Language,omitnil" name:"Language"`

	// Image rotation angle in degrees. 0°: The horizontal direction of the text on the image; a positive value: rotate clockwise; a negative value: rotate counterclockwise.
	Angel *float64 `json:"Angel,omitnil" name:"Angel"`

	// Total number of PDF pages to be returned if the image is a PDF. Default value: 0.
	PdfPageSize *int64 `json:"PdfPageSize,omitnil" name:"PdfPageSize"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type GeneralMachineItem struct {
	// Item name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Specification
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// Unit
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// Unit price
	Price *string `json:"Price,omitnil" name:"Price"`

	// Amount
	Total *string `json:"Total,omitnil" name:"Total"`

	// Tax rate
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// Tax amount
	Tax *string `json:"Tax,omitnil" name:"Tax"`
}

type GroupInfo struct {
	// The elements in each line.
	Groups []*LineInfo `json:"Groups,omitnil" name:"Groups"`
}

// Predefined struct for user
type HKIDCardOCRRequestParams struct {
	// Whether to check for authenticity.
	DetectFake *bool `json:"DetectFake,omitnil" name:"DetectFake"`

	// Whether to return identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// Base64 string of the image
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported yet.
	// Supported image size: The downloaded image cannot exceed 7 MB after being Base64-encoded, and it cannot take longer than 3 seconds to download the image.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type HKIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// Whether to check for authenticity.
	DetectFake *bool `json:"DetectFake,omitnil" name:"DetectFake"`

	// Whether to return identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// Base64 string of the image
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported yet.
	// Supported image size: The downloaded image cannot exceed 7 MB after being Base64-encoded, and it cannot take longer than 3 seconds to download the image.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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
	CnName *string `json:"CnName,omitnil" name:"CnName"`

	// Name in English
	EnName *string `json:"EnName,omitnil" name:"EnName"`

	// Telecode for the name in Chinese
	TelexCode *string `json:"TelexCode,omitnil" name:"TelexCode"`

	// Gender. Valid values: Male, Female
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// Date of birth
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// Permanent identity card.
	// 0: non-permanent;
	// 1: permanent;
	// -1: unknown.
	Permanent *int64 `json:"Permanent,omitnil" name:"Permanent"`

	// Identity card number
	IdNum *string `json:"IdNum,omitnil" name:"IdNum"`

	// Document symbol, i.e., the symbol under the date of birth, such as "***AZ"
	Symbol *string `json:"Symbol,omitnil" name:"Symbol"`

	// First issue date
	FirstIssueDate *string `json:"FirstIssueDate,omitnil" name:"FirstIssueDate"`

	// Last receipt date
	CurrentIssueDate *string `json:"CurrentIssueDate,omitnil" name:"CurrentIssueDate"`

	// Authenticity check.
	// 0: unable to judge (because the image is blurred, incomplete, reflective, too dark, etc.);
	// 1: forged;
	// 2: authentic.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FakeDetectResult *int64 `json:"FakeDetectResult,omitnil" name:"FakeDetectResult"`

	// Base64-encoded identity photo
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeadImage *string `json:"HeadImage,omitnil" name:"HeadImage"`

	// Multiple alarm codes. If the ID card is spoofed, photocopied, or photoshopped, the corresponding alarm code will be returned.
	// -9102: Alarm for photocopied document
	// -9103: Alarm for spoofed document
	// -9104: Alarm for photoshopped document
	WarningCode []*int64 `json:"WarningCode,omitnil" name:"WarningCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type HmtResidentPermitOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// `FRONT`: The side with the profile photo.
	// `BACK`: The side with the national emblem.
	// If this parameter is not specified, the system will automatically determine the ID card side.
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
}

type HmtResidentPermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// `FRONT`: The side with the profile photo.
	// `BACK`: The side with the national emblem.
	// If this parameter is not specified, the system will automatically determine the ID card side.
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
}

func (r *HmtResidentPermitOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HmtResidentPermitOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CardSide")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HmtResidentPermitOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HmtResidentPermitOCRResponseParams struct {
	// Name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Gender
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// Date of birth
	Birth *string `json:"Birth,omitnil" name:"Birth"`

	// Address
	Address *string `json:"Address,omitnil" name:"Address"`

	// ID card number
	IdCardNo *string `json:"IdCardNo,omitnil" name:"IdCardNo"`

	// 0: Front side.
	// 1: Back side.
	CardType *int64 `json:"CardType,omitnil" name:"CardType"`

	// Validity period
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// Issuing authority
	Authority *string `json:"Authority,omitnil" name:"Authority"`

	// Number of issues
	VisaNum *string `json:"VisaNum,omitnil" name:"VisaNum"`

	// Permit number
	PassNo *string `json:"PassNo,omitnil" name:"PassNo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type HmtResidentPermitOCRResponse struct {
	*tchttp.BaseResponse
	Response *HmtResidentPermitOCRResponseParams `json:"Response"`
}

func (r *HmtResidentPermitOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HmtResidentPermitOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IDCardOCRRequestParams struct {
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// `FRONT`: The side with the profile photo.
	// `BACK`: The side with the national emblem.
	// If this parameter is not specified, the system will automatically determine the ID card side.
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`

	// The following parameters are all of `bool` type and default to `false`:
	// `CropIdCard`: Crops the ID card photo (by removing extra edges outside the ID card and automatically correcting the shooting angle).
	// `CropPortrait`: Crops the profile photo (by automatically cutting out the face area in the ID card).
	// `CopyWarn`: Warns about photocopied images.
	// `BorderCheckWarn`: Warns about border and frame occlusions.
	// `ReshootWarn`: Warns about spoofed images.
	// `DetectPsWarn`: Warns about photoshopped images.
	// `TempIdWarn`: Warns about temporary ID cards.
	// `InvalidDateWarn`: Warns about invalid ID card validity periods.
	// `Quality`: Gets the image quality score (by evaluating the blurriness of the image).
	// `MultiCardDetect`: Enables multi-card detection.
	// `ReflectWarn`: Enables glare detection.
	// 
	// Parameter setting method via SDK:
	// Config = Json.stringify({"CropIdCard":true,"CropPortrait":true})
	// Parameter setting method via API 3.0 Explorer:
	// Config = {"CropIdCard":true,"CropPortrait":true}
	Config *string `json:"Config,omitnil" name:"Config"`
}

type IDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// `FRONT`: The side with the profile photo.
	// `BACK`: The side with the national emblem.
	// If this parameter is not specified, the system will automatically determine the ID card side.
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`

	// The following parameters are all of `bool` type and default to `false`:
	// `CropIdCard`: Crops the ID card photo (by removing extra edges outside the ID card and automatically correcting the shooting angle).
	// `CropPortrait`: Crops the profile photo (by automatically cutting out the face area in the ID card).
	// `CopyWarn`: Warns about photocopied images.
	// `BorderCheckWarn`: Warns about border and frame occlusions.
	// `ReshootWarn`: Warns about spoofed images.
	// `DetectPsWarn`: Warns about photoshopped images.
	// `TempIdWarn`: Warns about temporary ID cards.
	// `InvalidDateWarn`: Warns about invalid ID card validity periods.
	// `Quality`: Gets the image quality score (by evaluating the blurriness of the image).
	// `MultiCardDetect`: Enables multi-card detection.
	// `ReflectWarn`: Enables glare detection.
	// 
	// Parameter setting method via SDK:
	// Config = Json.stringify({"CropIdCard":true,"CropPortrait":true})
	// Parameter setting method via API 3.0 Explorer:
	// Config = {"CropIdCard":true,"CropPortrait":true}
	Config *string `json:"Config,omitnil" name:"Config"`
}

func (r *IDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CardSide")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IDCardOCRResponseParams struct {
	// Name (profile photo side)
	Name *string `json:"Name,omitnil" name:"Name"`

	// Gender (profile photo side)
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// Ethnicity (profile photo side)
	Nation *string `json:"Nation,omitnil" name:"Nation"`

	// Date of birth (profile photo side)
	Birth *string `json:"Birth,omitnil" name:"Birth"`

	// Address (profile photo side)
	Address *string `json:"Address,omitnil" name:"Address"`

	// ID number (profile photo side)
	IdNum *string `json:"IdNum,omitnil" name:"IdNum"`

	// Issuing authority (national emblem side)
	Authority *string `json:"Authority,omitnil" name:"Authority"`

	// Validity period (national emblem side)
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// Extended information, which will be returned only when requested. For the input parameters, please see example 3 and example 4.
	// `IdCard`: Base64-encoded content of the cropped ID card photo, which will be returned if `Config.CropIdCard` is set to `true`.
	// `Portrait`: Base64-encoded content of the ID photo on the card, which will be returned if `Config.CropPortrait` is set to `true`.
	// 
	// `Quality`: Image quality score, which will be returned if `Config.Quality` is set to `true`. Value range: 0–100. The lower the score, the blurrier the image. The recommended threshold is ≥ 50.
	// `BorderCodeValue`: Warning threshold score for incomplete ID card borders, which will be returned if `Config.BorderCheckWarn` is set to `true`. Value range: 0–100. The lower the score, the lower the probability of border occlusion. The recommended threshold value is ≤ 50.
	// 
	// `WarnInfos`: Warning information. Warning codes and descriptions are as follows:
	// -9100: The ID card validity period is invalid.
	// -9101: The ID card borders are incomplete.
	// -9102: The ID card image is photocopied.
	// -9103: The ID card image is spoofed. 
	// -9104: The ID card is a temporary one. 
	// -9105: The ID card frame is occluded.
	// -9106: The ID card image is photoshopped.
	// -9107: The ID card image has glares.
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *IDCardOCRResponseParams `json:"Response"`
}

func (r *IDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InvoiceItem struct {
	// The recognition result.
	// `OK`: Recognition is successful.
	// `FailedOperation.UnsupportedInvoice`: Recognition is not supported.
	// `FailedOperation.UnKnowError`: Recognition failed.
	// For the information about other error codes, see the OCR API description for each invoice/ticket.
	Code *string `json:"Code,omitnil" name:"Code"`

	// The type of invoice/ticket to which the recognized image belongs.
	// -1: Unknown
	// 0: Taxi receipt
	// 1: Quota invoice
	// 2: Train ticket
	// 3: VAT invoice
	// 5: Itinerary/Receipt of e-ticket for air transportation
	// 8: General machine-printed invoice
	// 9: Bus ticket
	// 10: Ship ticket
	// 11: VAT invoice (roll)
	// 12: Car sales invoice
	// 13: Toll receipt
	// 15: Non-tax revenue invoice
	// 16: Fully digitalized electronic invoice
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// The coordinates of the four vertices of the rotated image.
	Polygon *Polygon `json:"Polygon,omitnil" name:"Polygon"`

	// The rotation angle of the recognized image in the image with multiple types of invoices/tickets.
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// The recognized content.
	SingleInvoiceInfos *SingleInvoiceItem `json:"SingleInvoiceInfos,omitnil" name:"SingleInvoiceInfos"`

	// The number of the page on which the recognized invoice is in the image or PDF file, starting from 1 by default.
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// The detailed invoice type. See the description of `SubType`.
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// The invoice description. See the description of `TypeDescription`.
	TypeDescription *string `json:"TypeDescription,omitnil" name:"TypeDescription"`

	// The image file after cropping, encoded in Base64. This is returned if `EnableCutImage` is set to `true`.
	CutImage *string `json:"CutImage,omitnil" name:"CutImage"`

	// The description of the detailed invoice type. See the description of `SubType`.
	SubTypeDescription *string `json:"SubTypeDescription,omitnil" name:"SubTypeDescription"`
}

type ItemCoord struct {
	// X-coordinate of top-left point.
	X *int64 `json:"X,omitnil" name:"X"`

	// Y-coordinate of top-left point.
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// Width
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// Height
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type ItemInfo struct {
	// The key information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Key *Key `json:"Key,omitnil" name:"Key"`

	// The value information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *Value `json:"Value,omitnil" name:"Value"`
}

type Key struct {
	// The name of the recognized field.
	AutoName *string `json:"AutoName,omitnil" name:"AutoName"`

	// The name of a defined field (the key passed in).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConfigName *string `json:"ConfigName,omitnil" name:"ConfigName"`
}

type LicensePlateInfo struct {
	// The recognized license plate number.
	Number *string `json:"Number,omitnil" name:"Number"`

	// The confidence score (0–100).
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// The bounding box coordinates of the text line in the original image.
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`

	// The recognized license plate color, which currently includes "white", "black", "blue", "green", "yellow", "yellow-green", and "temporary plate".
	Color *string `json:"Color,omitnil" name:"Color"`
}

// Predefined struct for user
type LicensePlateOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type LicensePlateOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *LicensePlateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LicensePlateOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LicensePlateOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LicensePlateOCRResponseParams struct {
	// The recognized license plate number.
	Number *string `json:"Number,omitnil" name:"Number"`

	// The confidence score (0–100).
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// The bounding box coordinates of the text line in the original image.
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`

	// The recognized license plate color, which currently includes "white", "black", "blue", "green", "yellow", "yellow-green", and "temporary plate".
	Color *string `json:"Color,omitnil" name:"Color"`

	// The vehicle license plate information.
	LicensePlateInfos []*LicensePlateInfo `json:"LicensePlateInfos,omitnil" name:"LicensePlateInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type LicensePlateOCRResponse struct {
	*tchttp.BaseResponse
	Response *LicensePlateOCRResponseParams `json:"Response"`
}

func (r *LicensePlateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LicensePlateOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LineInfo struct {
	// The elements in a line
	Lines []*ItemInfo `json:"Lines,omitnil" name:"Lines"`
}

// Predefined struct for user
type MLIDCardOCRRequestParams struct {
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of an image. (This field is not available outside the Chinese mainland.)
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return an image. Default value: `false`.
	RetImage *bool `json:"RetImage,omitnil" name:"RetImage"`
}

type MLIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of an image. (This field is not available outside the Chinese mainland.)
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return an image. Default value: `false`.
	RetImage *bool `json:"RetImage,omitnil" name:"RetImage"`
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
	ID *string `json:"ID,omitnil" name:"ID"`

	// Full name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Address
	Address *string `json:"Address,omitnil" name:"Address"`

	// Gender
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// Alarm codes
	// -9103 Alarm for photographed certificate
	// -9102 Alarm for photocopied certificate
	// -9106 Alarm for covered certificate
	// -9107 Alarm for blurry image
	Warn []*int64 `json:"Warn,omitnil" name:"Warn"`

	// Identity photo
	Image *string `json:"Image,omitnil" name:"Image"`

	// This is an extended field, 
	// with the confidence of a field recognition result returned in the following format.
	// {
	//   Field name:{
	//     Confidence:0.9999
	//   }
	// }
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// Certificate type
	// MyKad  ID card
	// MyPR    Permanent resident card
	// MyTentera   Military identity card
	// MyKAS    Temporary ID card
	// POLIS  Police card
	// IKAD   Work permit
	// MyKid   Kid card
	Type *string `json:"Type,omitnil" name:"Type"`

	// Date of birth. This field is available only for work permits (i-Kad) and ID cards (MyKad).
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// Whether to return an image. Default value: false.
	RetImage *bool `json:"RetImage,omitnil" name:"RetImage"`
}

type MLIDPassportOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image. The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 500x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies more than 2/3 area of the image.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// Whether to return an image. Default value: false.
	RetImage *bool `json:"RetImage,omitnil" name:"RetImage"`
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
	ID *string `json:"ID,omitnil" name:"ID"`

	// Name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Date of birth
	DateOfBirth *string `json:"DateOfBirth,omitnil" name:"DateOfBirth"`

	// Gender (F: female, M: male)
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// Expiration date
	DateOfExpiration *string `json:"DateOfExpiration,omitnil" name:"DateOfExpiration"`

	// Issuing country
	IssuingCountry *string `json:"IssuingCountry,omitnil" name:"IssuingCountry"`

	// Country/region code
	Nationality *string `json:"Nationality,omitnil" name:"Nationality"`

	// Alarm codes
	// -9103 Alarm for spoofed document
	// -9102 Alarm for photocopied document (including black & white and color ones)
	// -9106 Alarm for covered card
	Warn []*int64 `json:"Warn,omitnil" name:"Warn"`

	// Identity photo
	Image *string `json:"Image,omitnil" name:"Image"`

	// Extended field:
	// {
	//     ID:{
	//         Confidence:0.9999
	//     },
	//     Name:{
	//         Confidence:0.9996
	//     }
	// }
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// The first row of the machine-readable zone (MRZ) at the bottom
	CodeSet *string `json:"CodeSet,omitnil" name:"CodeSet"`

	// The second row of the MRZ at the bottom
	CodeCrc *string `json:"CodeCrc,omitnil" name:"CodeCrc"`

	// The surname.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Surname *string `json:"Surname,omitnil" name:"Surname"`

	// The given name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GivenName *string `json:"GivenName,omitnil" name:"GivenName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type MachinePrintedInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Time
	Time *string `json:"Time,omitnil" name:"Time"`

	// Check code
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// Ciphertext
	Ciphertext *string `json:"Ciphertext,omitnil" name:"Ciphertext"`

	// Category
	Category *string `json:"Category,omitnil" name:"Category"`

	// Amount before tax
	PretaxAmount *string `json:"PretaxAmount,omitnil" name:"PretaxAmount"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// Tax
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// Industry
	IndustryClass *string `json:"IndustryClass,omitnil" name:"IndustryClass"`

	// Seller's name
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// Seller's taxpayer identification number
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// Seller's address and phone number
	SellerAddrTel *string `json:"SellerAddrTel,omitnil" name:"SellerAddrTel"`

	// Seller's bank account number
	SellerBankAccount *string `json:"SellerBankAccount,omitnil" name:"SellerBankAccount"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// Buyer's address and phone number
	BuyerAddrTel *string `json:"BuyerAddrTel,omitnil" name:"BuyerAddrTel"`

	// Buyer's bank account number
	BuyerBankAccount *string `json:"BuyerBankAccount,omitnil" name:"BuyerBankAccount"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil" name:"Province"`

	// City
	City *string `json:"City,omitnil" name:"City"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`

	// Whether it is a general machine-printed invoice issued by Zhejiang or Guangdong province (0: No, 1: Yes)
	ElectronicMark *int64 `json:"ElectronicMark,omitnil" name:"ElectronicMark"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// Payee
	Receiptor *string `json:"Receiptor,omitnil" name:"Receiptor"`

	// Reviewer
	Reviewer *string `json:"Reviewer,omitnil" name:"Reviewer"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Operator's payment information
	PaymentInfo *string `json:"PaymentInfo,omitnil" name:"PaymentInfo"`

	// Operator-assigned invoice pickup user
	TicketPickupUser *string `json:"TicketPickupUser,omitnil" name:"TicketPickupUser"`

	// Operator's merchant number
	MerchantNumber *string `json:"MerchantNumber,omitnil" name:"MerchantNumber"`

	// Operator's order number
	OrderNumber *string `json:"OrderNumber,omitnil" name:"OrderNumber"`

	// Items
	GeneralMachineItems []*GeneralMachineItem `json:"GeneralMachineItems,omitnil" name:"GeneralMachineItems"`
}

// Predefined struct for user
type MainlandPermitOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the ID photo. By default, the ID photo is not returned.
	RetProfile *bool `json:"RetProfile,omitnil" name:"RetProfile"`
}

type MainlandPermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the ID photo. By default, the ID photo is not returned.
	RetProfile *bool `json:"RetProfile,omitnil" name:"RetProfile"`
}

func (r *MainlandPermitOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MainlandPermitOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "RetProfile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MainlandPermitOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MainlandPermitOCRResponseParams struct {
	// Name in Chinese
	Name *string `json:"Name,omitnil" name:"Name"`

	// Name in English
	EnglishName *string `json:"EnglishName,omitnil" name:"EnglishName"`

	// Gender
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// Date of birth
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// Issuing authority
	IssueAuthority *string `json:"IssueAuthority,omitnil" name:"IssueAuthority"`

	// Validity period
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// ID number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Place of issue
	IssueAddress *string `json:"IssueAddress,omitnil" name:"IssueAddress"`

	// Number of issues
	IssueNumber *string `json:"IssueNumber,omitnil" name:"IssueNumber"`

	// Document type
	Type *string `json:"Type,omitnil" name:"Type"`

	// Base64-encoded profile photo, which is returned only when `RetProfile` is set to `True`
	Profile *string `json:"Profile,omitnil" name:"Profile"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type MainlandPermitOCRResponse struct {
	*tchttp.BaseResponse
	Response *MainlandPermitOCRResponseParams `json:"Response"`
}

func (r *MainlandPermitOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MainlandPermitOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MedicalInvoice struct {

	Title *string `json:"Title,omitnil" name:"Title"`


	Code *string `json:"Code,omitnil" name:"Code"`


	Number *string `json:"Number,omitnil" name:"Number"`


	Total *string `json:"Total,omitnil" name:"Total"`


	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`


	Date *string `json:"Date,omitnil" name:"Date"`


	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`


	Place *string `json:"Place,omitnil" name:"Place"`


	Reviewer *string `json:"Reviewer,omitnil" name:"Reviewer"`
}

type MotorVehicleSaleInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Amount before tax
	PretaxAmount *string `json:"PretaxAmount,omitnil" name:"PretaxAmount"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// Seller's name
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// Seller's company code
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// Seller's phone number
	SellerTel *string `json:"SellerTel,omitnil" name:"SellerTel"`

	// Seller's address
	SellerAddress *string `json:"SellerAddress,omitnil" name:"SellerAddress"`

	// Seller's account opening bank
	SellerBank *string `json:"SellerBank,omitnil" name:"SellerBank"`

	// Seller's bank account number
	SellerBankAccount *string `json:"SellerBankAccount,omitnil" name:"SellerBankAccount"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// Buyer's ID number/organization code
	BuyerID *string `json:"BuyerID,omitnil" name:"BuyerID"`

	// Tax authority
	TaxAuthorities *string `json:"TaxAuthorities,omitnil" name:"TaxAuthorities"`

	// Code of the tax authority
	TaxAuthoritiesCode *string `json:"TaxAuthoritiesCode,omitnil" name:"TaxAuthoritiesCode"`

	// VIN
	VIN *string `json:"VIN,omitnil" name:"VIN"`

	// Vehicle model
	VehicleModel *string `json:"VehicleModel,omitnil" name:"VehicleModel"`

	// Engine No.
	VehicleEngineCode *string `json:"VehicleEngineCode,omitnil" name:"VehicleEngineCode"`

	// No. of the certificate of conformity
	CertificateNumber *string `json:"CertificateNumber,omitnil" name:"CertificateNumber"`

	// Inspection No.
	InspectionNumber *string `json:"InspectionNumber,omitnil" name:"InspectionNumber"`

	// Machine No.
	MachineID *string `json:"MachineID,omitnil" name:"MachineID"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil" name:"Province"`

	// City
	City *string `json:"City,omitnil" name:"City"`

	// Tax
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// Tax rate
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`

	// Tonnage
	Tonnage *string `json:"Tonnage,omitnil" name:"Tonnage"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Form type
	FormType *string `json:"FormType,omitnil" name:"FormType"`

	// Form name
	FormName *string `json:"FormName,omitnil" name:"FormName"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// Tax payment voucher number
	TaxNum *string `json:"TaxNum,omitnil" name:"TaxNum"`

	// Passenger capacity
	MaxPeopleNum *string `json:"MaxPeopleNum,omitnil" name:"MaxPeopleNum"`

	// Origin
	Origin *string `json:"Origin,omitnil" name:"Origin"`

	// Machine-printed invoice code
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// Machine-printed invoice number
	MachineNumber *string `json:"MachineNumber,omitnil" name:"MachineNumber"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`
}

type NonTaxIncomeBill struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Check code
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// Payer's name
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// Payer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// Payee's name
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// Payee's company name
	SellerCompany *string `json:"SellerCompany,omitnil" name:"SellerCompany"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Currency
	CurrencyCode *string `json:"CurrencyCode,omitnil" name:"CurrencyCode"`

	// Reviewer
	Reviewer *string `json:"Reviewer,omitnil" name:"Reviewer"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Other information
	OtherInfo *string `json:"OtherInfo,omitnil" name:"OtherInfo"`

	// Payment code
	PaymentCode *string `json:"PaymentCode,omitnil" name:"PaymentCode"`

	// Collecting organization's code
	ReceiveUnitCode *string `json:"ReceiveUnitCode,omitnil" name:"ReceiveUnitCode"`

	// Collecting organization's name
	Receiver *string `json:"Receiver,omitnil" name:"Receiver"`

	// Operator
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// Payer's account
	PayerAccount *string `json:"PayerAccount,omitnil" name:"PayerAccount"`

	// Payer's account opening bank
	PayerBank *string `json:"PayerBank,omitnil" name:"PayerBank"`

	// Payee's account
	ReceiverAccount *string `json:"ReceiverAccount,omitnil" name:"ReceiverAccount"`

	// Payee's account opening bank
	ReceiverBank *string `json:"ReceiverBank,omitnil" name:"ReceiverBank"`

	// Items
	NonTaxItems []*NonTaxItem `json:"NonTaxItems,omitnil" name:"NonTaxItems"`
}

type NonTaxItem struct {
	// Item code
	ItemID *string `json:"ItemID,omitnil" name:"ItemID"`

	// Item name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Unit
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// Standard
	Standard *string `json:"Standard,omitnil" name:"Standard"`

	// Amount
	Total *string `json:"Total,omitnil" name:"Total"`
}

type OtherInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Amount
	Total *string `json:"Total,omitnil" name:"Total"`

	// List
	OtherInvoiceListItems []*OtherInvoiceItem `json:"OtherInvoiceListItems,omitnil" name:"OtherInvoiceListItems"`

	// Table
	OtherInvoiceTableItems []*OtherInvoiceList `json:"OtherInvoiceTableItems,omitnil" name:"OtherInvoiceTableItems"`
}

type OtherInvoiceItem struct {
	// Field name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Field value
	Value *string `json:"Value,omitnil" name:"Value"`
}

type OtherInvoiceList struct {
	// List
	OtherInvoiceItemList []*OtherInvoiceItem `json:"OtherInvoiceItemList,omitnil" name:"OtherInvoiceItemList"`
}

// Predefined struct for user
type PermitOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type PermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *PermitOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PermitOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PermitOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PermitOCRResponseParams struct {
	// Name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Name in English
	EnglishName *string `json:"EnglishName,omitnil" name:"EnglishName"`

	// ID number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Gender
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// Validity period
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// Issuing authority
	IssueAuthority *string `json:"IssueAuthority,omitnil" name:"IssueAuthority"`

	// Place of issue
	IssueAddress *string `json:"IssueAddress,omitnil" name:"IssueAddress"`

	// Date of birth
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PermitOCRResponse struct {
	*tchttp.BaseResponse
	Response *PermitOCRResponseParams `json:"Response"`
}

func (r *PermitOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PermitOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Polygon struct {
	// The coordinates of the upper-left vertex.
	LeftTop *Coord `json:"LeftTop,omitnil" name:"LeftTop"`

	// The coordinates of the upper-right vertex.
	RightTop *Coord `json:"RightTop,omitnil" name:"RightTop"`

	// The coordinates of the lower-left vertex.
	RightBottom *Coord `json:"RightBottom,omitnil" name:"RightBottom"`

	// The coordinates of the lower-right vertex.
	LeftBottom *Coord `json:"LeftBottom,omitnil" name:"LeftBottom"`
}

type QuotaInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil" name:"Province"`

	// City
	City *string `json:"City,omitnil" name:"City"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`
}

// Predefined struct for user
type RecognizeGeneralInvoiceRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, JPEG, and PDF. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Supported image pixels: 20 to 10,000
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, JPEG, and PDF. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Supported image pixels: 20 to 10,000
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// The list of the types of invoices to be recognized. If this parameter is left empty, all types of invoices are recognized.
	// 0: Taxi receipt
	// 1: Quota invoice
	// 2: Train ticket
	// 3: VAT invoice
	// 5: Itinerary/Receipt of e-ticket for air transport
	// 8: General machine-printed invoice
	// 9: Bus ticket
	// 10: Ship ticket
	// 11: VAT invoice (roll)
	// 12: Car sales inovice
	// 13: Toll receipt
	// 15: Non-tax revenue invoice
	// 16: Fully digitalized electronic invoice
	// -1: Other
	// 
	// By default, this parameter is left empty, which means to recognize all types of invoices.
	// When a single type is passed in, the image is recognized based on this type.
	// You can only specify a singe type or all types, but not some types.
	Types []*int64 `json:"Types,omitnil" name:"Types"`

	// Whether to enable recognition of other invoices. If you enable this feature, other invoices can be recognized. Default value: `true`.	
	EnableOther *bool `json:"EnableOther,omitnil" name:"EnableOther"`

	// Whether to enable PDF recognition. If you enable this feature, both images and PDF files can be recognized. Default value: `true`.
	EnablePdf *bool `json:"EnablePdf,omitnil" name:"EnablePdf"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `EnablePdf` is `true`. Default value: 1.
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// Whether to enable multi-page PDF recognition. If you enable this feature, multiple pages of a PDF file can be recognized, and the recognition results of a maximum of the first 30 pages can be returned. After you enable this feature, input parameters `EnablePdf` and `PdfPageNumber` are invalid. Default value: `false`.
	EnableMultiplePage *bool `json:"EnableMultiplePage,omitnil" name:"EnableMultiplePage"`

	// Whether to return the Base64-encoded value of the cropped image. Default value: `false`.
	EnableCutImage *bool `json:"EnableCutImage,omitnil" name:"EnableCutImage"`
}

type RecognizeGeneralInvoiceRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, JPEG, and PDF. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Supported image pixels: 20 to 10,000
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, JPEG, and PDF. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Supported image pixels: 20 to 10,000
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// The list of the types of invoices to be recognized. If this parameter is left empty, all types of invoices are recognized.
	// 0: Taxi receipt
	// 1: Quota invoice
	// 2: Train ticket
	// 3: VAT invoice
	// 5: Itinerary/Receipt of e-ticket for air transport
	// 8: General machine-printed invoice
	// 9: Bus ticket
	// 10: Ship ticket
	// 11: VAT invoice (roll)
	// 12: Car sales inovice
	// 13: Toll receipt
	// 15: Non-tax revenue invoice
	// 16: Fully digitalized electronic invoice
	// -1: Other
	// 
	// By default, this parameter is left empty, which means to recognize all types of invoices.
	// When a single type is passed in, the image is recognized based on this type.
	// You can only specify a singe type or all types, but not some types.
	Types []*int64 `json:"Types,omitnil" name:"Types"`

	// Whether to enable recognition of other invoices. If you enable this feature, other invoices can be recognized. Default value: `true`.	
	EnableOther *bool `json:"EnableOther,omitnil" name:"EnableOther"`

	// Whether to enable PDF recognition. If you enable this feature, both images and PDF files can be recognized. Default value: `true`.
	EnablePdf *bool `json:"EnablePdf,omitnil" name:"EnablePdf"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `EnablePdf` is `true`. Default value: 1.
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// Whether to enable multi-page PDF recognition. If you enable this feature, multiple pages of a PDF file can be recognized, and the recognition results of a maximum of the first 30 pages can be returned. After you enable this feature, input parameters `EnablePdf` and `PdfPageNumber` are invalid. Default value: `false`.
	EnableMultiplePage *bool `json:"EnableMultiplePage,omitnil" name:"EnableMultiplePage"`

	// Whether to return the Base64-encoded value of the cropped image. Default value: `false`.
	EnableCutImage *bool `json:"EnableCutImage,omitnil" name:"EnableCutImage"`
}

func (r *RecognizeGeneralInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeGeneralInvoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Types")
	delete(f, "EnableOther")
	delete(f, "EnablePdf")
	delete(f, "PdfPageNumber")
	delete(f, "EnableMultiplePage")
	delete(f, "EnableCutImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeGeneralInvoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeGeneralInvoiceResponseParams struct {
	// Mixed invoice/ticket recognition result. Please click the link on the left for details.
	MixedInvoiceItems []*InvoiceItem `json:"MixedInvoiceItems,omitnil" name:"MixedInvoiceItems"`

	// Total number of pages in the PDF file.
	TotalPDFCount *int64 `json:"TotalPDFCount,omitnil" name:"TotalPDFCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeGeneralInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeGeneralInvoiceResponseParams `json:"Response"`
}

func (r *RecognizeGeneralInvoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeGeneralInvoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeIndonesiaIDCardOCRRequestParams struct {
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// The scene, which defaults to `V1`.
	// Valid values:
	// V1
	// V2
	Scene *string `json:"Scene,omitnil" name:"Scene"`
}

type RecognizeIndonesiaIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// The scene, which defaults to `V1`.
	// Valid values:
	// V1
	// V2
	Scene *string `json:"Scene,omitnil" name:"Scene"`
}

func (r *RecognizeIndonesiaIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeIndonesiaIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	delete(f, "Scene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeIndonesiaIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeIndonesiaIDCardOCRResponseParams struct {
	// The Single Identity Number.
	NIK *string `json:"NIK,omitnil" name:"NIK"`

	// The full name.
	Nama *string `json:"Nama,omitnil" name:"Nama"`

	// The place and date of birth.
	TempatTglLahir *string `json:"TempatTglLahir,omitnil" name:"TempatTglLahir"`

	// The gender.
	JenisKelamin *string `json:"JenisKelamin,omitnil" name:"JenisKelamin"`

	// The blood type.
	GolDarah *string `json:"GolDarah,omitnil" name:"GolDarah"`

	// The address.
	Alamat *string `json:"Alamat,omitnil" name:"Alamat"`

	// The street.
	RTRW *string `json:"RTRW,omitnil" name:"RTRW"`

	// The village.
	KelDesa *string `json:"KelDesa,omitnil" name:"KelDesa"`

	// The region.
	Kecamatan *string `json:"Kecamatan,omitnil" name:"Kecamatan"`

	// The religion.
	Agama *string `json:"Agama,omitnil" name:"Agama"`

	// The marital status.
	StatusPerkawinan *string `json:"StatusPerkawinan,omitnil" name:"StatusPerkawinan"`

	// The occupation.
	Perkerjaan *string `json:"Perkerjaan,omitnil" name:"Perkerjaan"`

	// The nationality.
	KewargaNegaraan *string `json:"KewargaNegaraan,omitnil" name:"KewargaNegaraan"`

	// The expiry date.
	BerlakuHingga *string `json:"BerlakuHingga,omitnil" name:"BerlakuHingga"`

	// The issue date.
	IssuedDate *string `json:"IssuedDate,omitnil" name:"IssuedDate"`

	// The photo.
	Photo *string `json:"Photo,omitnil" name:"Photo"`

	// The province, which is supported when the value of `Scene` is `V2`.
	Provinsi *string `json:"Provinsi,omitnil" name:"Provinsi"`

	// The city, which is supported when the value of `Scene` is `V2`.
	Kota *string `json:"Kota,omitnil" name:"Kota"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeIndonesiaIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeIndonesiaIDCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeIndonesiaIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeIndonesiaIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeKoreanDrivingLicenseOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

type RecognizeKoreanDrivingLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

func (r *RecognizeKoreanDrivingLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeKoreanDrivingLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeKoreanDrivingLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeKoreanDrivingLicenseOCRResponseParams struct {
	// The ID card number.
	ID *string `json:"ID,omitnil" name:"ID"`

	// The license number.
	LicenseNumber *string `json:"LicenseNumber,omitnil" name:"LicenseNumber"`

	// The resident registration number.
	Number *string `json:"Number,omitnil" name:"Number"`

	// The license class type.
	Type *string `json:"Type,omitnil" name:"Type"`

	// The address.
	Address *string `json:"Address,omitnil" name:"Address"`

	// The name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// The renewal period.
	AptitudeTesDate *string `json:"AptitudeTesDate,omitnil" name:"AptitudeTesDate"`

	// The issue date.
	DateOfIssue *string `json:"DateOfIssue,omitnil" name:"DateOfIssue"`

	// The Base64-encoded identity photo.
	Photo *string `json:"Photo,omitnil" name:"Photo"`

	// The gender.
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// The birth date in the format of dd/mm/yyyy.
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeKoreanDrivingLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeKoreanDrivingLicenseOCRResponseParams `json:"Response"`
}

func (r *RecognizeKoreanDrivingLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeKoreanDrivingLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeKoreanIDCardOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

type RecognizeKoreanIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

func (r *RecognizeKoreanIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeKoreanIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeKoreanIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeKoreanIDCardOCRResponseParams struct {
	// The ID card number.
	ID *string `json:"ID,omitnil" name:"ID"`

	// The address.
	Address *string `json:"Address,omitnil" name:"Address"`

	// The name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// The issue date.
	DateOfIssue *string `json:"DateOfIssue,omitnil" name:"DateOfIssue"`

	// The Base64-encoded identity photo.
	Photo *string `json:"Photo,omitnil" name:"Photo"`

	// The gender.
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// The birth date in the format of dd/mm/yyyy.
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeKoreanIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeKoreanIDCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeKoreanIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeKoreanIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesDrivingLicenseOCRRequestParams struct {
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

type RecognizePhilippinesDrivingLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

func (r *RecognizePhilippinesDrivingLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesDrivingLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesDrivingLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesDrivingLicenseOCRResponseParams struct {
	// The Base64-encoded identity photo.
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// The full name.
	Name *TextDetectionResult `json:"Name,omitnil" name:"Name"`

	// The last name.
	LastName *TextDetectionResult `json:"LastName,omitnil" name:"LastName"`

	// The first name.
	FirstName *TextDetectionResult `json:"FirstName,omitnil" name:"FirstName"`

	// The middle name.
	MiddleName *TextDetectionResult `json:"MiddleName,omitnil" name:"MiddleName"`

	// The nationality.
	Nationality *TextDetectionResult `json:"Nationality,omitnil" name:"Nationality"`

	// The gender.
	Sex *TextDetectionResult `json:"Sex,omitnil" name:"Sex"`

	// The address.
	Address *TextDetectionResult `json:"Address,omitnil" name:"Address"`

	// The license No.
	LicenseNo *TextDetectionResult `json:"LicenseNo,omitnil" name:"LicenseNo"`

	// The expiration date.
	ExpiresDate *TextDetectionResult `json:"ExpiresDate,omitnil" name:"ExpiresDate"`

	// The agency code.
	AgencyCode *TextDetectionResult `json:"AgencyCode,omitnil" name:"AgencyCode"`

	// The date of birth.
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesDrivingLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesDrivingLicenseOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesDrivingLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesDrivingLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesSssIDOCRRequestParams struct {
	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RecognizePhilippinesSssIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *RecognizePhilippinesSssIDOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesSssIDOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReturnHeadImage")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesSssIDOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesSssIDOCRResponseParams struct {
	// The Base64-encoded identity photo.
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// The common reference number (CRN).
	LicenseNumber *TextDetectionResult `json:"LicenseNumber,omitnil" name:"LicenseNumber"`

	// The full name.
	FullName *TextDetectionResult `json:"FullName,omitnil" name:"FullName"`

	// The date of birth.
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesSssIDOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesSssIDOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesSssIDOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesSssIDOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesTinIDOCRRequestParams struct {
	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RecognizePhilippinesTinIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *RecognizePhilippinesTinIDOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesTinIDOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReturnHeadImage")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesTinIDOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesTinIDOCRResponseParams struct {
	// The Base64-encoded identity photo.
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// The tax identification number (TIN).
	LicenseNumber *TextDetectionResult `json:"LicenseNumber,omitnil" name:"LicenseNumber"`

	// The name.
	FullName *TextDetectionResult `json:"FullName,omitnil" name:"FullName"`

	// The address.
	Address *TextDetectionResult `json:"Address,omitnil" name:"Address"`

	// The birth date.
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// The issue date.
	IssueDate *TextDetectionResult `json:"IssueDate,omitnil" name:"IssueDate"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesTinIDOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesTinIDOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesTinIDOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesTinIDOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesUMIDOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

type RecognizePhilippinesUMIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

func (r *RecognizePhilippinesUMIDOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesUMIDOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesUMIDOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesUMIDOCRResponseParams struct {
	// The surname.
	Surname *TextDetectionResult `json:"Surname,omitnil" name:"Surname"`

	// The middle name.
	MiddleName *TextDetectionResult `json:"MiddleName,omitnil" name:"MiddleName"`

	// The given name.
	GivenName *TextDetectionResult `json:"GivenName,omitnil" name:"GivenName"`

	// The address.
	Address *TextDetectionResult `json:"Address,omitnil" name:"Address"`

	// The date of birth.
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// The common reference number (CRN).
	CRN *TextDetectionResult `json:"CRN,omitnil" name:"CRN"`

	// The gender.
	Sex *TextDetectionResult `json:"Sex,omitnil" name:"Sex"`

	// The Base64-encoded identity photo.
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesUMIDOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesUMIDOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesUMIDOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesUMIDOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesVoteIDOCRRequestParams struct {
	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RecognizePhilippinesVoteIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *RecognizePhilippinesVoteIDOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesVoteIDOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReturnHeadImage")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesVoteIDOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesVoteIDOCRResponseParams struct {
	// The Base64-encoded identity photo.
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// The voter's identification number (VIN).
	VIN *TextDetectionResult `json:"VIN,omitnil" name:"VIN"`

	// The first name.
	FirstName *TextDetectionResult `json:"FirstName,omitnil" name:"FirstName"`

	// The last name.
	LastName *TextDetectionResult `json:"LastName,omitnil" name:"LastName"`

	// The date of birth.
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// The civil status.
	CivilStatus *TextDetectionResult `json:"CivilStatus,omitnil" name:"CivilStatus"`

	// The citizenship.
	Citizenship *TextDetectionResult `json:"Citizenship,omitnil" name:"Citizenship"`

	// The address.
	Address *TextDetectionResult `json:"Address,omitnil" name:"Address"`

	// The precinct.
	PrecinctNo *TextDetectionResult `json:"PrecinctNo,omitnil" name:"PrecinctNo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesVoteIDOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesVoteIDOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesVoteIDOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesVoteIDOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeTableAccurateOCRRequestParams struct {
	// The Base64-encoded value of an image.
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600 x 800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// Supported image pixels: 20 to 10,000
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image or PDF file.
	// The image or PDF file cannot exceed 7 MB after being Base64-encoded. A resolution above 600 x 800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// Supported image pixels: 20 to 10,000
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `IsPdf` is `true`. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type RecognizeTableAccurateOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image.
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600 x 800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// Supported image pixels: 20 to 10,000
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image or PDF file.
	// The image or PDF file cannot exceed 7 MB after being Base64-encoded. A resolution above 600 x 800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// Supported image pixels: 20 to 10,000
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `IsPdf` is `true`. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

func (r *RecognizeTableAccurateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTableAccurateOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeTableAccurateOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeTableAccurateOCRResponseParams struct {
	// The recognized text information. Please click the link on the left for details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableDetections []*TableInfo `json:"TableDetections,omitnil" name:"TableDetections"`

	// Base64-encoded Excel data.
	Data *string `json:"Data,omitnil" name:"Data"`

	// The total number of pages in the PDF file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PdfPageSize *int64 `json:"PdfPageSize,omitnil" name:"PdfPageSize"`

	// Image rotation angle in degrees. 0°: The horizontal direction of the text on the image; a negative value: rotate counterclockwise. Value range: -360° to 0°.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeTableAccurateOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeTableAccurateOCRResponseParams `json:"Response"`
}

func (r *RecognizeTableAccurateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTableAccurateOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeThaiIDCardOCRRequestParams struct {
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to crop the profile photo. The default value is `false`, meaning not to return the Base64-encoded value of the profile photo on the Thai identity card.
	// When this parameter is set to `true`, the Base64-encoded value of the profile photo on the Thai identity card after rotation correction is returned.
	CropPortrait *bool `json:"CropPortrait,omitnil" name:"CropPortrait"`
}

type RecognizeThaiIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Whether to crop the profile photo. The default value is `false`, meaning not to return the Base64-encoded value of the profile photo on the Thai identity card.
	// When this parameter is set to `true`, the Base64-encoded value of the profile photo on the Thai identity card after rotation correction is returned.
	CropPortrait *bool `json:"CropPortrait,omitnil" name:"CropPortrait"`
}

func (r *RecognizeThaiIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeThaiIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CropPortrait")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeThaiIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeThaiIDCardOCRResponseParams struct {
	// ID card number
	ID *string `json:"ID,omitnil" name:"ID"`

	// Name in Thai
	ThaiName *string `json:"ThaiName,omitnil" name:"ThaiName"`

	// Name in English
	EnFirstName *string `json:"EnFirstName,omitnil" name:"EnFirstName"`

	// Address
	Address *string `json:"Address,omitnil" name:"Address"`

	// Date of birth
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// Date of issue
	IssueDate *string `json:"IssueDate,omitnil" name:"IssueDate"`

	// Expiration date
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// Name in English
	EnLastName *string `json:"EnLastName,omitnil" name:"EnLastName"`

	// Identity photo
	PortraitImage *string `json:"PortraitImage,omitnil" name:"PortraitImage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeThaiIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeThaiIDCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeThaiIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeThaiIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rect struct {
	// X-coordinate of top-left point
	X *int64 `json:"X,omitnil" name:"X"`

	// Y-coordinate of top-left point
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// Width
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// Height
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type SealInfo struct {
	// Seal body information
	SealBody *string `json:"SealBody,omitnil" name:"SealBody"`

	// Seal coordinates
	Location *Rect `json:"Location,omitnil" name:"Location"`

	// Other text content
	OtherTexts []*string `json:"OtherTexts,omitnil" name:"OtherTexts"`

	// Seal shape. Valid values:
	// 0: Round
	// 1: Oval
	// 2: Rectangle
	// 3: Diamond
	// 4: Triangle
	SealShape *string `json:"SealShape,omitnil" name:"SealShape"`
}

// Predefined struct for user
type SealOCRRequestParams struct {
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type SealOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *SealOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SealOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SealOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SealOCRResponseParams struct {
	// Seal content
	SealBody *string `json:"SealBody,omitnil" name:"SealBody"`

	// Seal coordinates
	Location *Rect `json:"Location,omitnil" name:"Location"`

	// Other text content
	OtherTexts []*string `json:"OtherTexts,omitnil" name:"OtherTexts"`

	// All seal information
	SealInfos []*SealInfo `json:"SealInfos,omitnil" name:"SealInfos"`

	// Seal shape. Valid values:
	// 0: Round
	// 1: Oval
	// 2: Rectangle
	// 3: Diamond
	// 4: Triangle
	SealShape *string `json:"SealShape,omitnil" name:"SealShape"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SealOCRResponse struct {
	*tchttp.BaseResponse
	Response *SealOCRResponseParams `json:"Response"`
}

func (r *SealOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SealOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShippingInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Name
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Date
	Date *string `json:"Date,omitnil" name:"Date"`

	// Time
	Time *string `json:"Time,omitnil" name:"Time"`

	// Departure station
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// Destination
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// Fare
	Total *string `json:"Total,omitnil" name:"Total"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil" name:"Province"`

	// City
	City *string `json:"City,omitnil" name:"City"`

	// Currency
	CurrencyCode *string `json:"CurrencyCode,omitnil" name:"CurrencyCode"`
}

type SingleInvoiceItem struct {
	// Special VAT invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatSpecialInvoice *VatInvoiceInfo `json:"VatSpecialInvoice,omitnil" name:"VatSpecialInvoice"`

	// General VAT invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatCommonInvoice *VatInvoiceInfo `json:"VatCommonInvoice,omitnil" name:"VatCommonInvoice"`

	// Electronic general VAT invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicCommonInvoice *VatInvoiceInfo `json:"VatElectronicCommonInvoice,omitnil" name:"VatElectronicCommonInvoice"`

	// Electronic special VAT invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicSpecialInvoice *VatInvoiceInfo `json:"VatElectronicSpecialInvoice,omitnil" name:"VatElectronicSpecialInvoice"`

	// Blockchain electronic invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicInvoiceBlockchain *VatInvoiceInfo `json:"VatElectronicInvoiceBlockchain,omitnil" name:"VatElectronicInvoiceBlockchain"`

	// Electronic general VAT invoice (toll)
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicInvoiceToll *VatInvoiceInfo `json:"VatElectronicInvoiceToll,omitnil" name:"VatElectronicInvoiceToll"`

	// Electronic invoice (special)
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicSpecialInvoiceFull *VatElectronicInfo `json:"VatElectronicSpecialInvoiceFull,omitnil" name:"VatElectronicSpecialInvoiceFull"`

	// Electronic invoice (general)
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicInvoiceFull *VatElectronicInfo `json:"VatElectronicInvoiceFull,omitnil" name:"VatElectronicInvoiceFull"`

	// General machine-printed invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	MachinePrintedInvoice *MachinePrintedInvoice `json:"MachinePrintedInvoice,omitnil" name:"MachinePrintedInvoice"`

	// Bus ticket
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusInvoice *BusInvoice `json:"BusInvoice,omitnil" name:"BusInvoice"`

	// Ship ticket
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShippingInvoice *ShippingInvoice `json:"ShippingInvoice,omitnil" name:"ShippingInvoice"`

	// Toll receipt
	// Note: This field may return null, indicating that no valid values can be obtained.
	TollInvoice *TollInvoice `json:"TollInvoice,omitnil" name:"TollInvoice"`

	// Other invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	OtherInvoice *OtherInvoice `json:"OtherInvoice,omitnil" name:"OtherInvoice"`

	// Motor vehicle sales invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	MotorVehicleSaleInvoice *MotorVehicleSaleInvoice `json:"MotorVehicleSaleInvoice,omitnil" name:"MotorVehicleSaleInvoice"`

	// Used car invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsedCarPurchaseInvoice *UsedCarPurchaseInvoice `json:"UsedCarPurchaseInvoice,omitnil" name:"UsedCarPurchaseInvoice"`

	// General VAT invoice (roll)
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatInvoiceRoll *VatInvoiceRoll `json:"VatInvoiceRoll,omitnil" name:"VatInvoiceRoll"`

	// Taxi receipt
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaxiTicket *TaxiTicket `json:"TaxiTicket,omitnil" name:"TaxiTicket"`

	// Quota invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	QuotaInvoice *QuotaInvoice `json:"QuotaInvoice,omitnil" name:"QuotaInvoice"`

	// Itinerary/Receipt of e-ticket for air transportation
	// Note: This field may return null, indicating that no valid values can be obtained.
	AirTransport *AirTransport `json:"AirTransport,omitnil" name:"AirTransport"`

	// Non-tax revenue general receipt
	// Note: This field may return null, indicating that no valid values can be obtained.
	NonTaxIncomeGeneralBill *NonTaxIncomeBill `json:"NonTaxIncomeGeneralBill,omitnil" name:"NonTaxIncomeGeneralBill"`

	// Non-tax revenue unified payment voucher
	// Note: This field may return null, indicating that no valid values can be obtained.
	NonTaxIncomeElectronicBill *NonTaxIncomeBill `json:"NonTaxIncomeElectronicBill,omitnil" name:"NonTaxIncomeElectronicBill"`

	// Train ticket
	// Note: This field may return null, indicating that no valid values can be obtained.
	TrainTicket *TrainTicket `json:"TrainTicket,omitnil" name:"TrainTicket"`


	MedicalOutpatientInvoice *MedicalInvoice `json:"MedicalOutpatientInvoice,omitnil" name:"MedicalOutpatientInvoice"`


	MedicalHospitalizedInvoice *MedicalInvoice `json:"MedicalHospitalizedInvoice,omitnil" name:"MedicalHospitalizedInvoice"`
}

// Predefined struct for user
type SmartStructuralOCRV2RequestParams struct {
	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `IsPdf` is `true`. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// The names of the fields you want to return for the structured information recognition.
	// For example, if you want to return only the recognition result of the "Name" and "Gender" fields, set this parameter as follows:
	// ItemNames=["Name","Gender"]
	ItemNames []*string `json:"ItemNames,omitnil" name:"ItemNames"`

	// Whether to enable recognition of all fields.
	ReturnFullText *bool `json:"ReturnFullText,omitnil" name:"ReturnFullText"`
}

type SmartStructuralOCRV2Request struct {
	*tchttp.BaseRequest
	
	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `IsPdf` is `true`. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// The names of the fields you want to return for the structured information recognition.
	// For example, if you want to return only the recognition result of the "Name" and "Gender" fields, set this parameter as follows:
	// ItemNames=["Name","Gender"]
	ItemNames []*string `json:"ItemNames,omitnil" name:"ItemNames"`

	// Whether to enable recognition of all fields.
	ReturnFullText *bool `json:"ReturnFullText,omitnil" name:"ReturnFullText"`
}

func (r *SmartStructuralOCRV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartStructuralOCRV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	delete(f, "ItemNames")
	delete(f, "ReturnFullText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SmartStructuralOCRV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SmartStructuralOCRV2ResponseParams struct {
	// The rotation angle (degrees) of the text on the image. 0: The text is horizontal. Positive value: The text is rotated clockwise. Negative value: The text is rotated counterclockwise.
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// The structural information (key-value).
	StructuralList []*GroupInfo `json:"StructuralList,omitnil" name:"StructuralList"`

	// The recognized text information.
	WordList []*WordItem `json:"WordList,omitnil" name:"WordList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SmartStructuralOCRV2Response struct {
	*tchttp.BaseResponse
	Response *SmartStructuralOCRV2ResponseParams `json:"Response"`
}

func (r *SmartStructuralOCRV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartStructuralOCRV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableCellInfo struct {
	// Column index of the upper-left corner of the cell
	ColTl *int64 `json:"ColTl,omitnil" name:"ColTl"`

	// Row index of the upper-left corner of the cell
	RowTl *int64 `json:"RowTl,omitnil" name:"RowTl"`

	// Column index of the lower-right corner of the cell
	ColBr *int64 `json:"ColBr,omitnil" name:"ColBr"`

	// Row index of the lower-right corner of the cell
	RowBr *int64 `json:"RowBr,omitnil" name:"RowBr"`

	// Recognized string text within the cell. If there are multiple lines, they are separated by "\n".
	Text *string `json:"Text,omitnil" name:"Text"`

	// Cell type
	Type *string `json:"Type,omitnil" name:"Type"`

	// Cell confidence
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// Four-point coordinates of the cell in the image
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`
}

type TableInfo struct {
	// Cell content
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Cells []*TableCellInfo `json:"Cells,omitnil" name:"Cells"`

	// Type of text in the image. Valid values:
	// 0: Non-table text
	// 1: Text in a bordered table
	// 2: Text in a borderless table
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// The coordinates of the four vertices (upper-left, upper-right, lower-right, and lower-left) of the table body.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableCoordPoint []*Coord `json:"TableCoordPoint,omitnil" name:"TableCoordPoint"`
}

// Predefined struct for user
type TableOCRRequestParams struct {
	// Base64-encoded value of image.
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported at present.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type TableOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image.
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported at present.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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
	TextDetections []*TextTable `json:"TextDetections,omitnil" name:"TextDetections"`

	// Base64-encoded Excel data.
	Data *string `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type TaxiTicket struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Start time
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// End time
	TimeGetOff *string `json:"TimeGetOff,omitnil" name:"TimeGetOff"`

	// Unit price
	Price *string `json:"Price,omitnil" name:"Price"`

	// Distance
	Mileage *string `json:"Mileage,omitnil" name:"Mileage"`

	// Total amount
	Total *string `json:"Total,omitnil" name:"Total"`

	// Invoice place
	Place *string `json:"Place,omitnil" name:"Place"`

	// Province
	Province *string `json:"Province,omitnil" name:"Province"`

	// City
	City *string `json:"City,omitnil" name:"City"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// License plate number
	LicensePlate *string `json:"LicensePlate,omitnil" name:"LicensePlate"`

	// Fuel surcharge
	FuelFee *string `json:"FuelFee,omitnil" name:"FuelFee"`

	// Booking fee
	BookingCallFee *string `json:"BookingCallFee,omitnil" name:"BookingCallFee"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`
}

type TextDetection struct {
	// Recognized text line content.
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`

	// Confidence. Value range: 0–100.
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// Text line coordinates, which are represented as 4 vertex coordinates.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`

	// Extended field.
	// The paragraph information `Parag` returned by the `GeneralBasicOcr` API contains `ParagNo`.
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// Pixel coordinates of the text line in the image after rotation correction, which is in the format of `(X-coordinate of top-left point, Y-coordinate of top-left point, width, height)`.
	ItemPolygon *ItemCoord `json:"ItemPolygon,omitnil" name:"ItemPolygon"`

	// Information about a character, including the character itself and its confidence. Supported APIs: `GeneralBasicOCR`, `GeneralAccurateOCR`
	Words []*DetectedWords `json:"Words,omitnil" name:"Words"`

	// Coordinates of a word’s four corners on the input image. Supported APIs: `GeneralBasicOCR`, `GeneralAccurateOCR`
	WordCoordPoint []*DetectedWordCoordPoint `json:"WordCoordPoint,omitnil" name:"WordCoordPoint"`
}

type TextDetectionResult struct {
	// The recognized text line content.
	Value *string `json:"Value,omitnil" name:"Value"`

	// The coordinates, represented in the coordinates of the four points.
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`
}

type TextTable struct {
	// Column index of the top-left corner of the cell.
	ColTl *int64 `json:"ColTl,omitnil" name:"ColTl"`

	// Row index of the top-left corner of the cell.
	RowTl *int64 `json:"RowTl,omitnil" name:"RowTl"`

	// Column index of the bottom-right corner of the cell.
	ColBr *int64 `json:"ColBr,omitnil" name:"ColBr"`

	// Row index of the bottom-right corner of the cell.
	RowBr *int64 `json:"RowBr,omitnil" name:"RowBr"`

	// Cell text
	Text *string `json:"Text,omitnil" name:"Text"`

	// Cell type. Valid values: body, header, footer
	Type *string `json:"Type,omitnil" name:"Type"`

	// Confidence. Value range: 0–100
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// Text line coordinates, which are represented as 4 vertex coordinates.
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`

	// Extended field
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`
}

type TollInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Date
	Date *string `json:"Date,omitnil" name:"Date"`

	// Time
	Time *string `json:"Time,omitnil" name:"Time"`

	// Entrance
	Entrance *string `json:"Entrance,omitnil" name:"Entrance"`

	// Exit
	Exit *string `json:"Exit,omitnil" name:"Exit"`

	// Highway mark (0: No, 1: Yes)
	HighwayMark *int64 `json:"HighwayMark,omitnil" name:"HighwayMark"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`
}

type TrainTicket struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Departure date
	DateGetOn *string `json:"DateGetOn,omitnil" name:"DateGetOn"`

	// Departure time
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// Passenger's name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Departure station
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// Destination
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// Seat class
	Seat *string `json:"Seat,omitnil" name:"Seat"`

	// Total amount
	Total *string `json:"Total,omitnil" name:"Total"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Serial number
	SerialNumber *string `json:"SerialNumber,omitnil" name:"SerialNumber"`

	// ID card number
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// Check-in gate
	GateNumber *string `json:"GateNumber,omitnil" name:"GateNumber"`

	// Fleet number
	TrainNumber *string `json:"TrainNumber,omitnil" name:"TrainNumber"`

	// Handling fee
	HandlingFee *string `json:"HandlingFee,omitnil" name:"HandlingFee"`

	// Original ticket price
	OriginalFare *string `json:"OriginalFare,omitnil" name:"OriginalFare"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// Seat No.
	SeatNumber *string `json:"SeatNumber,omitnil" name:"SeatNumber"`

	// Ticket pickup address
	PickUpAddress *string `json:"PickUpAddress,omitnil" name:"PickUpAddress"`

	// Ticket change information
	TicketChange *string `json:"TicketChange,omitnil" name:"TicketChange"`

	// Additional fare
	AdditionalFare *string `json:"AdditionalFare,omitnil" name:"AdditionalFare"`

	// Receipt No.
	ReceiptNumber *string `json:"ReceiptNumber,omitnil" name:"ReceiptNumber"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Whether it is for reimbursement only (0: No, 1: Yes)
	ReimburseOnlyMark *int64 `json:"ReimburseOnlyMark,omitnil" name:"ReimburseOnlyMark"`
}

type UsedCarPurchaseInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// Seller's name
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// Seller's phone number
	SellerTel *string `json:"SellerTel,omitnil" name:"SellerTel"`

	// Seller's company code/personal ID card number
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// Seller's address
	SellerAddress *string `json:"SellerAddress,omitnil" name:"SellerAddress"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// Buyer's company code/personal ID card number
	BuyerID *string `json:"BuyerID,omitnil" name:"BuyerID"`

	// Buyer's address
	BuyerAddress *string `json:"BuyerAddress,omitnil" name:"BuyerAddress"`

	// Buyer's phone number
	BuyerTel *string `json:"BuyerTel,omitnil" name:"BuyerTel"`

	// Company (used car market) name
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// Company's taxpayer identification number
	CompanyTaxID *string `json:"CompanyTaxID,omitnil" name:"CompanyTaxID"`

	// Company's account opening bank and account number
	CompanyBankAccount *string `json:"CompanyBankAccount,omitnil" name:"CompanyBankAccount"`

	// Company's phone number
	CompanyTel *string `json:"CompanyTel,omitnil" name:"CompanyTel"`

	// Company's address
	CompanyAddress *string `json:"CompanyAddress,omitnil" name:"CompanyAddress"`

	// Name of the transfer-to department of motor vehicles
	TransferAdministrationName *string `json:"TransferAdministrationName,omitnil" name:"TransferAdministrationName"`

	// License plate number
	LicensePlate *string `json:"LicensePlate,omitnil" name:"LicensePlate"`

	// Registration certificate No.
	RegistrationNumber *string `json:"RegistrationNumber,omitnil" name:"RegistrationNumber"`

	// VIN
	VIN *string `json:"VIN,omitnil" name:"VIN"`

	// Vehicle model
	VehicleModel *string `json:"VehicleModel,omitnil" name:"VehicleModel"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil" name:"Province"`

	// City
	City *string `json:"City,omitnil" name:"City"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Form type
	FormType *string `json:"FormType,omitnil" name:"FormType"`

	// Form name
	FormName *string `json:"FormName,omitnil" name:"FormName"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`
}

type Value struct {
	// The value of the recognized field.
	AutoContent *string `json:"AutoContent,omitnil" name:"AutoContent"`

	// The coordinates of the four vertices.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Coord *Polygon `json:"Coord,omitnil" name:"Coord"`
}

type VatElectronicInfo struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Amount before tax
	PretaxAmount *string `json:"PretaxAmount,omitnil" name:"PretaxAmount"`

	// Tax
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// Seller's name
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// Seller's taxpayer identification number
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Subtotal amount
	SubTotal *string `json:"SubTotal,omitnil" name:"SubTotal"`

	// Subtotal tax
	SubTax *string `json:"SubTax,omitnil" name:"SubTax"`

	// Detailed items of an electronic invoice
	VatElectronicItems []*VatElectronicItemInfo `json:"VatElectronicItems,omitnil" name:"VatElectronicItems"`
}

type VatElectronicItemInfo struct {
	// Item name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// Specification
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// Unit price
	Price *string `json:"Price,omitnil" name:"Price"`

	// Amount
	Total *string `json:"Total,omitnil" name:"Total"`

	// Tax rate
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// Tax amount
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// Unit
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// Vehicle No.
	VehicleBrand *string `json:"VehicleBrand,omitnil" name:"VehicleBrand"`

	// Departure place
	DeparturePlace *string `json:"DeparturePlace,omitnil" name:"DeparturePlace"`

	// Destination
	ArrivalPlace *string `json:"ArrivalPlace,omitnil" name:"ArrivalPlace"`

	// Name of the transported goods. It is returned only for a goods transport service invoice.
	TransportItemsName *string `json:"TransportItemsName,omitnil" name:"TransportItemsName"`

	// Location of the construction service. It is returned only for a construction invoice.
	PlaceOfBuildingService *string `json:"PlaceOfBuildingService,omitnil" name:"PlaceOfBuildingService"`

	// Name of the construction project. It is returned only for a construction invoice.
	BuildingName *string `json:"BuildingName,omitnil" name:"BuildingName"`

	// Property or real estate ownership certificate No. It is returned only for a real estate operation and leasing service invoice.
	EstateNumber *string `json:"EstateNumber,omitnil" name:"EstateNumber"`

	// Unit of area. It is returned only for a real estate operation and leasing service invoice.
	AreaUnit *string `json:"AreaUnit,omitnil" name:"AreaUnit"`
}

type VatInvoiceInfo struct {
	// Check code
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// Form type
	FormType *string `json:"FormType,omitnil" name:"FormType"`

	// Vehicle and vessel tax
	TravelTax *string `json:"TravelTax,omitnil" name:"TravelTax"`

	// Buyer's address and phone number
	BuyerAddrTel *string `json:"BuyerAddrTel,omitnil" name:"BuyerAddrTel"`

	// Buyer's bank account number
	BuyerBankAccount *string `json:"BuyerBankAccount,omitnil" name:"BuyerBankAccount"`

	// Company seal content
	CompanySealContent *string `json:"CompanySealContent,omitnil" name:"CompanySealContent"`

	// Tax authority seal content
	TaxSealContent *string `json:"TaxSealContent,omitnil" name:"TaxSealContent"`

	// Service type
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// City
	City *string `json:"City,omitnil" name:"City"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Whether there is an agent (0: No, 1: Yes)
	AgentMark *int64 `json:"AgentMark,omitnil" name:"AgentMark"`

	// Whether there is a toll (0: No, 1: Yes)
	TransitMark *int64 `json:"TransitMark,omitnil" name:"TransitMark"`

	// Whether there is refined oil (0: No, 1: Yes)
	OilMark *int64 `json:"OilMark,omitnil" name:"OilMark"`

	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Machine-printed invoice number
	NumberConfirm *string `json:"NumberConfirm,omitnil" name:"NumberConfirm"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// Amount before tax
	PretaxAmount *string `json:"PretaxAmount,omitnil" name:"PretaxAmount"`

	// Tax
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// Machine No.
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// Ciphertext
	Ciphertext *string `json:"Ciphertext,omitnil" name:"Ciphertext"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Seller's name
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// Seller's taxpayer identification number
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// Seller's address and phone number
	SellerAddrTel *string `json:"SellerAddrTel,omitnil" name:"SellerAddrTel"`

	// Seller's bank account number
	SellerBankAccount *string `json:"SellerBankAccount,omitnil" name:"SellerBankAccount"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// Reviewer
	Reviewer *string `json:"Reviewer,omitnil" name:"Reviewer"`

	// Province
	Province *string `json:"Province,omitnil" name:"Province"`

	// Information about VAT invoice items
	VatInvoiceItemInfos []*VatInvoiceItemInfo `json:"VatInvoiceItemInfos,omitnil" name:"VatInvoiceItemInfos"`

	// Machine-printed invoice code
	CodeConfirm *string `json:"CodeConfirm,omitnil" name:"CodeConfirm"`

	// Payee
	Receiptor *string `json:"Receiptor,omitnil" name:"Receiptor"`


	ElectronicFullMark *int64 `json:"ElectronicFullMark,omitnil" name:"ElectronicFullMark"`


	ElectronicFullNumber *string `json:"ElectronicFullNumber,omitnil" name:"ElectronicFullNumber"`


	FormName *string `json:"FormName,omitnil" name:"FormName"`
}

type VatInvoiceItemInfo struct {
	// Item name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Specification
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// Unit
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// Unit price
	Price *string `json:"Price,omitnil" name:"Price"`

	// Amount
	Total *string `json:"Total,omitnil" name:"Total"`

	// Tax rate
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// Tax amount
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// Start date
	DateStart *string `json:"DateStart,omitnil" name:"DateStart"`

	// End date
	DateEnd *string `json:"DateEnd,omitnil" name:"DateEnd"`

	// License plate number
	LicensePlate *string `json:"LicensePlate,omitnil" name:"LicensePlate"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`
}

type VatInvoiceRoll struct {
	// Invoice title
	Title *string `json:"Title,omitnil" name:"Title"`

	// Invoice code
	Code *string `json:"Code,omitnil" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil" name:"Number"`

	// Machine-printed invoice number
	NumberConfirm *string `json:"NumberConfirm,omitnil" name:"NumberConfirm"`

	// Date of issue
	Date *string `json:"Date,omitnil" name:"Date"`

	// Check code
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// Seller's name
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// Seller's taxpayer identification number
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// Category
	Category *string `json:"Category,omitnil" name:"Category"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// Invoice type
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil" name:"Province"`

	// City
	City *string `json:"City,omitnil" name:"City"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// Service type
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// Company seal content
	CompanySealContent *string `json:"CompanySealContent,omitnil" name:"CompanySealContent"`

	// Tax authority seal content
	TaxSealContent *string `json:"TaxSealContent,omitnil" name:"TaxSealContent"`

	// Items
	VatRollItems []*VatRollItem `json:"VatRollItems,omitnil" name:"VatRollItems"`
}

type VatRollItem struct {
	// Item name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// Unit price
	Price *string `json:"Price,omitnil" name:"Price"`

	// Amount
	Total *string `json:"Total,omitnil" name:"Total"`
}

// Predefined struct for user
type VinOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type VinOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *VinOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VinOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VinOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VinOCRResponseParams struct {
	// The detected VIN.
	Vin *string `json:"Vin,omitnil" name:"Vin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VinOCRResponse struct {
	*tchttp.BaseResponse
	Response *VinOCRResponseParams `json:"Response"`
}

func (r *VinOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VinOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WordItem struct {
	// The text content.
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`

	// The coordinates of the four vertices.
	Coord *Polygon `json:"Coord,omitnil" name:"Coord"`
}