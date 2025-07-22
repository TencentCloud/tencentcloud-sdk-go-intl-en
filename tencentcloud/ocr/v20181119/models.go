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

package v20181119

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AirTransport struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// E-ticket No.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Check code
	CheckCode *string `json:"CheckCode,omitnil,omitempty" name:"CheckCode"`

	// Serial number
	SerialNumber *string `json:"SerialNumber,omitnil,omitempty" name:"SerialNumber"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Agent code
	AgentCode *string `json:"AgentCode,omitnil,omitempty" name:"AgentCode"`

	// First line of the agent code
	AgentCodeFirst *string `json:"AgentCodeFirst,omitnil,omitempty" name:"AgentCodeFirst"`

	// Second line of the agent code
	AgentCodeSecond *string `json:"AgentCodeSecond,omitnil,omitempty" name:"AgentCodeSecond"`

	// Name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// ID card number
	UserID *string `json:"UserID,omitnil,omitempty" name:"UserID"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil,omitempty" name:"Issuer"`

	// Fare
	Fare *string `json:"Fare,omitnil,omitempty" name:"Fare"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Fuel surcharge
	FuelSurcharge *string `json:"FuelSurcharge,omitnil,omitempty" name:"FuelSurcharge"`

	// Aviation Development Fund
	AirDevelopmentFund *string `json:"AirDevelopmentFund,omitnil,omitempty" name:"AirDevelopmentFund"`

	// Insurance
	Insurance *string `json:"Insurance,omitnil,omitempty" name:"Insurance"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Domestic or international tag
	DomesticInternationalTag *string `json:"DomesticInternationalTag,omitnil,omitempty" name:"DomesticInternationalTag"`

	// Not-valid-before date
	DateStart *string `json:"DateStart,omitnil,omitempty" name:"DateStart"`

	// Not-valid-after date
	DateEnd *string `json:"DateEnd,omitnil,omitempty" name:"DateEnd"`

	// Endorsements/Restrictions
	Endorsement *string `json:"Endorsement,omitnil,omitempty" name:"Endorsement"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Items
	FlightItems []*FlightItem `json:"FlightItems,omitnil,omitempty" name:"FlightItems"`
}

// Predefined struct for user
type BankCardOCRRequestParams struct {
	// Base64-encoded value of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 7 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the bank card image data after preprocessing (precise cropping and alignment). Default value: `false`
	RetBorderCutImage *bool `json:"RetBorderCutImage,omitnil,omitempty" name:"RetBorderCutImage"`

	// Whether to return the card number image data after slicing. Default value: `false`
	RetCardNoImage *bool `json:"RetCardNoImage,omitnil,omitempty" name:"RetCardNoImage"`

	// Whether to enable photocopy check. If the input image is a bank card photocopy, an alarm will be returned. Default value: `false`
	EnableCopyCheck *bool `json:"EnableCopyCheck,omitnil,omitempty" name:"EnableCopyCheck"`

	// Whether to enable photograph check. If the input image is a bank card photograph, an alarm will be returned. Default value: `false`
	EnableReshootCheck *bool `json:"EnableReshootCheck,omitnil,omitempty" name:"EnableReshootCheck"`

	// Whether to enable obscured border check. If the input image is a bank card with obscured border, an alarm will be returned. Default value: `false`
	EnableBorderCheck *bool `json:"EnableBorderCheck,omitnil,omitempty" name:"EnableBorderCheck"`

	// Whether to return the image quality value, which measures how clear an image is. Default value: `false`
	EnableQualityValue *bool `json:"EnableQualityValue,omitnil,omitempty" name:"EnableQualityValue"`
}

type BankCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 7 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the bank card image data after preprocessing (precise cropping and alignment). Default value: `false`
	RetBorderCutImage *bool `json:"RetBorderCutImage,omitnil,omitempty" name:"RetBorderCutImage"`

	// Whether to return the card number image data after slicing. Default value: `false`
	RetCardNoImage *bool `json:"RetCardNoImage,omitnil,omitempty" name:"RetCardNoImage"`

	// Whether to enable photocopy check. If the input image is a bank card photocopy, an alarm will be returned. Default value: `false`
	EnableCopyCheck *bool `json:"EnableCopyCheck,omitnil,omitempty" name:"EnableCopyCheck"`

	// Whether to enable photograph check. If the input image is a bank card photograph, an alarm will be returned. Default value: `false`
	EnableReshootCheck *bool `json:"EnableReshootCheck,omitnil,omitempty" name:"EnableReshootCheck"`

	// Whether to enable obscured border check. If the input image is a bank card with obscured border, an alarm will be returned. Default value: `false`
	EnableBorderCheck *bool `json:"EnableBorderCheck,omitnil,omitempty" name:"EnableBorderCheck"`

	// Whether to return the image quality value, which measures how clear an image is. Default value: `false`
	EnableQualityValue *bool `json:"EnableQualityValue,omitnil,omitempty" name:"EnableQualityValue"`
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
	CardNo *string `json:"CardNo,omitnil,omitempty" name:"CardNo"`

	// Bank information
	BankInfo *string `json:"BankInfo,omitnil,omitempty" name:"BankInfo"`

	// Expiration date. Format: 07/2023
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// Card type
	CardType *string `json:"CardType,omitnil,omitempty" name:"CardType"`

	// Card name
	CardName *string `json:"CardName,omitnil,omitempty" name:"CardName"`

	// Sliced image data
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BorderCutImage *string `json:"BorderCutImage,omitnil,omitempty" name:"BorderCutImage"`

	// Card number image data
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CardNoImage *string `json:"CardNoImage,omitnil,omitempty" name:"CardNoImage"`

	// Warning code:
	// -9110: the bank card date is invalid. 
	// -9111: the bank card border is incomplete. 
	// -9112: the bank card image is reflective.
	// -9113: the bank card image is a photocopy.
	// -9114: the bank card image is a photograph.
	// Multiple warning codes may be returned at a time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WarningCode []*int64 `json:"WarningCode,omitnil,omitempty" name:"WarningCode"`

	// Image quality value, which is returned when `EnableQualityValue` is set to `true`. The smaller the value, the less clear the image is. Value range: 0−100 (a threshold greater than or equal to 50 is recommended.)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	QualityValue *int64 `json:"QualityValue,omitnil,omitempty" name:"QualityValue"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Departure time
	TimeGetOn *string `json:"TimeGetOn,omitnil,omitempty" name:"TimeGetOn"`

	// Departure date
	DateGetOn *string `json:"DateGetOn,omitnil,omitempty" name:"DateGetOn"`

	// Departure station
	StationGetOn *string `json:"StationGetOn,omitnil,omitempty" name:"StationGetOn"`

	// Destination
	StationGetOff *string `json:"StationGetOff,omitnil,omitempty" name:"StationGetOff"`

	// Fare
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Consumption type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// ID card number
	UserID *string `json:"UserID,omitnil,omitempty" name:"UserID"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Departure place
	PlaceGetOn *string `json:"PlaceGetOn,omitnil,omitempty" name:"PlaceGetOn"`

	// Check-in gate
	GateNumber *string `json:"GateNumber,omitnil,omitempty" name:"GateNumber"`

	// Fare category
	TicketType *string `json:"TicketType,omitnil,omitempty" name:"TicketType"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil,omitempty" name:"VehicleType"`

	// Seat No.
	SeatNumber *string `json:"SeatNumber,omitnil,omitempty" name:"SeatNumber"`

	// Fleet number
	TrainNumber *string `json:"TrainNumber,omitnil,omitempty" name:"TrainNumber"`
}

type ConfigAdvanced struct {
	// Single attribute configuration of a template.
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`
}

type Coord struct {
	// Horizontal coordinate
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Vertical coordinate
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`
}

type DetectedWordCoordPoint struct {
	// Coordinates of a word’s four corners in a clockwise order on the input image, starting from the upper-left corner
	WordCoordinate []*Coord `json:"WordCoordinate,omitnil,omitempty" name:"WordCoordinate"`
}

type DetectedWords struct {
	// Confidence. Value range: 0–100
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// A possible character
	Character *string `json:"Character,omitnil,omitempty" name:"Character"`
}

// Predefined struct for user
type ExtractDocMultiRequestParams struct {
	// The Url address of the image. supported image formats: PNG, JPG, JPEG, WORD, EXCEL. GIF format is not currently supported. supported image size: no more than 10M after Base64 encoding. image download time should not exceed 3 seconds. supported image pixels: between 20-10000px. images stored in tencent cloud's Url ensure higher download speed and stability. it is recommended to store images in tencent cloud. the speed and stability of non-tencent cloud storage urls may be impacted.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The Base64 value of the image. supported image formats: PNG, JPG, JPEG, WORD, EXCEL. GIF format is not currently supported. supported image size: no more than 10M after encoding the downloaded image with Base64. image download time: no more than 3 seconds. supported image pixels: between 20-10000px. either ImageUrl or ImageBase64 must be provided. if both are provided, only use ImageUrl.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Specifies the page number of the PDF to be recognized. only single page recognition is supported. valid when uploading a PDF file with the IsPdf parameter set to true. default value is the first 3 pages.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// Specifies the field names to be returned by the customized structuring feature. for example, if the customer wants to add the recognition result of two fields, name and gender, manually input ItemNames=["name","gender"].
	ItemNames []*string `json:"ItemNames,omitnil,omitempty" name:"ItemNames"`

	// true: only custom field.
	// False: default value field + custom field.
	// Default true.
	ItemNamesShowMode *bool `json:"ItemNamesShowMode,omitnil,omitempty" name:"ItemNamesShowMode"`

	// Whether the full-text field recognition is enabled.
	ReturnFullText *bool `json:"ReturnFullText,omitnil,omitempty" name:"ReturnFullText"`

	// Configuration ID support: 
	// -- General
	// -- InvoiceEng
	// -- WayBillEng
	// -- CustomsDeclaration
	// -- WeightNote
	// -- MedicalMeter
	// -- BillOfLading
	// -- EntrustmentBook
	// -- Statement
	// -- BookingConfirmation
	// -- AirWayBill
	// -- Table
	// -- SteelLabel
	// -- CarInsurance
	// -- MultiRealEstateCertificate
	// -- MultiRealEstateMaterial
	// -- HongKongUtilityBill
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Whether the full-text field coordinate value recognition is enabled.
	EnableCoord *bool `json:"EnableCoord,omitnil,omitempty" name:"EnableCoord"`

	// Whether parent-child key recognition is enabled. the option is selected by default.
	OutputParentKey *bool `json:"OutputParentKey,omitnil,omitempty" name:"OutputParentKey"`

	// Single attribute configuration of a template.
	ConfigAdvanced *ConfigAdvanced `json:"ConfigAdvanced,omitnil,omitempty" name:"ConfigAdvanced"`

	// When cn, the added key is chinese.  
	// When set to en, the added key is english.
	OutputLanguage *string `json:"OutputLanguage,omitnil,omitempty" name:"OutputLanguage"`
}

type ExtractDocMultiRequest struct {
	*tchttp.BaseRequest
	
	// The Url address of the image. supported image formats: PNG, JPG, JPEG, WORD, EXCEL. GIF format is not currently supported. supported image size: no more than 10M after Base64 encoding. image download time should not exceed 3 seconds. supported image pixels: between 20-10000px. images stored in tencent cloud's Url ensure higher download speed and stability. it is recommended to store images in tencent cloud. the speed and stability of non-tencent cloud storage urls may be impacted.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The Base64 value of the image. supported image formats: PNG, JPG, JPEG, WORD, EXCEL. GIF format is not currently supported. supported image size: no more than 10M after encoding the downloaded image with Base64. image download time: no more than 3 seconds. supported image pixels: between 20-10000px. either ImageUrl or ImageBase64 must be provided. if both are provided, only use ImageUrl.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Specifies the page number of the PDF to be recognized. only single page recognition is supported. valid when uploading a PDF file with the IsPdf parameter set to true. default value is the first 3 pages.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// Specifies the field names to be returned by the customized structuring feature. for example, if the customer wants to add the recognition result of two fields, name and gender, manually input ItemNames=["name","gender"].
	ItemNames []*string `json:"ItemNames,omitnil,omitempty" name:"ItemNames"`

	// true: only custom field.
	// False: default value field + custom field.
	// Default true.
	ItemNamesShowMode *bool `json:"ItemNamesShowMode,omitnil,omitempty" name:"ItemNamesShowMode"`

	// Whether the full-text field recognition is enabled.
	ReturnFullText *bool `json:"ReturnFullText,omitnil,omitempty" name:"ReturnFullText"`

	// Configuration ID support: 
	// -- General
	// -- InvoiceEng
	// -- WayBillEng
	// -- CustomsDeclaration
	// -- WeightNote
	// -- MedicalMeter
	// -- BillOfLading
	// -- EntrustmentBook
	// -- Statement
	// -- BookingConfirmation
	// -- AirWayBill
	// -- Table
	// -- SteelLabel
	// -- CarInsurance
	// -- MultiRealEstateCertificate
	// -- MultiRealEstateMaterial
	// -- HongKongUtilityBill
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Whether the full-text field coordinate value recognition is enabled.
	EnableCoord *bool `json:"EnableCoord,omitnil,omitempty" name:"EnableCoord"`

	// Whether parent-child key recognition is enabled. the option is selected by default.
	OutputParentKey *bool `json:"OutputParentKey,omitnil,omitempty" name:"OutputParentKey"`

	// Single attribute configuration of a template.
	ConfigAdvanced *ConfigAdvanced `json:"ConfigAdvanced,omitnil,omitempty" name:"ConfigAdvanced"`

	// When cn, the added key is chinese.  
	// When set to en, the added key is english.
	OutputLanguage *string `json:"OutputLanguage,omitnil,omitempty" name:"OutputLanguage"`
}

func (r *ExtractDocMultiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExtractDocMultiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "PdfPageNumber")
	delete(f, "ItemNames")
	delete(f, "ItemNamesShowMode")
	delete(f, "ReturnFullText")
	delete(f, "ConfigId")
	delete(f, "EnableCoord")
	delete(f, "OutputParentKey")
	delete(f, "ConfigAdvanced")
	delete(f, "OutputLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExtractDocMultiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExtractDocMultiResponseParams struct {
	// Image rotation angle (angle system). the text's horizontal direction is 0. clockwise is positive; counterclockwise is negative.
	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// Configures the structured text info.
	StructuralList []*GroupInfo `json:"StructuralList,omitnil,omitempty" name:"StructuralList"`

	// Restore text information.
	WordList []*WordItem `json:"WordList,omitnil,omitempty" name:"WordList"`

	// Number of sample identification fields.
	TokenNum *int64 `json:"TokenNum,omitnil,omitempty" name:"TokenNum"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExtractDocMultiResponse struct {
	*tchttp.BaseResponse
	Response *ExtractDocMultiResponseParams `json:"Response"`
}

func (r *ExtractDocMultiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExtractDocMultiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlightItem struct {
	// Departure terminal
	TerminalGetOn *string `json:"TerminalGetOn,omitnil,omitempty" name:"TerminalGetOn"`

	// Arrival terminal
	TerminalGetOff *string `json:"TerminalGetOff,omitnil,omitempty" name:"TerminalGetOff"`

	// Carrier
	Carrier *string `json:"Carrier,omitnil,omitempty" name:"Carrier"`

	// Flight number
	FlightNumber *string `json:"FlightNumber,omitnil,omitempty" name:"FlightNumber"`

	// Class
	Seat *string `json:"Seat,omitnil,omitempty" name:"Seat"`

	// Departure date
	DateGetOn *string `json:"DateGetOn,omitnil,omitempty" name:"DateGetOn"`

	// Departure time
	TimeGetOn *string `json:"TimeGetOn,omitnil,omitempty" name:"TimeGetOn"`

	// Departure city
	StationGetOn *string `json:"StationGetOn,omitnil,omitempty" name:"StationGetOn"`

	// Arrival city
	StationGetOff *string `json:"StationGetOff,omitnil,omitempty" name:"StationGetOff"`

	// Baggage allowance
	Allow *string `json:"Allow,omitnil,omitempty" name:"Allow"`

	// Fare category
	FareBasis *string `json:"FareBasis,omitnil,omitempty" name:"FareBasis"`
}

// Predefined struct for user
type GeneralAccurateOCRRequestParams struct {
	// Base64-encoded value of image.
	// The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitnil,omitempty" name:"IsWords"`

	// Whether to slice the input image to enhance the recognition effects for scenarios where the whole image is big, but the size of a single character is small (e.g., test papers). This feature is disabled by default.
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitnil,omitempty" name:"EnableDetectSplit"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitnil,omitempty" name:"IsPdf"`

	// Number of a PDF page that needs to be recognized. Currently, only one single page can be recognized. This parameter takes effect only if a PDF file is uploaded and `IsPdf` is set to `true`. Default value: `1`
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`
}

type GeneralAccurateOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image.
	// The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitnil,omitempty" name:"IsWords"`

	// Whether to slice the input image to enhance the recognition effects for scenarios where the whole image is big, but the size of a single character is small (e.g., test papers). This feature is disabled by default.
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitnil,omitempty" name:"EnableDetectSplit"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitnil,omitempty" name:"IsPdf"`

	// Number of a PDF page that needs to be recognized. Currently, only one single page can be recognized. This parameter takes effect only if a PDF file is uploaded and `IsPdf` is set to `true`. Default value: `1`
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`
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
	TextDetections []*TextDetection `json:"TextDetections,omitnil,omitempty" name:"TextDetections"`

	// Image rotation angle in degrees. 0°: The horizontal direction of the text on the image; a positive value: rotate clockwise; a negative value: rotate counterclockwise.
	Angel *float64 `json:"Angel,omitnil,omitempty" name:"Angel"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image/PDF. (This field is not supported outside Chinese mainland)
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Reserved field.
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

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
	LanguageType *string `json:"LanguageType,omitnil,omitempty" name:"LanguageType"`

	// Whether to enable PDF recognition. Default value: false. After this feature is enabled, both images and PDF files can be recognized at the same time.
	IsPdf *bool `json:"IsPdf,omitnil,omitempty" name:"IsPdf"`

	// Page number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of the `IsPdf` parameter is `true`. Default value: 1.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitnil,omitempty" name:"IsWords"`
}

type GeneralBasicOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image/PDF.
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image/PDF. (This field is not supported outside Chinese mainland)
	// The image/PDF cannot exceed 7 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Reserved field.
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

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
	LanguageType *string `json:"LanguageType,omitnil,omitempty" name:"LanguageType"`

	// Whether to enable PDF recognition. Default value: false. After this feature is enabled, both images and PDF files can be recognized at the same time.
	IsPdf *bool `json:"IsPdf,omitnil,omitempty" name:"IsPdf"`

	// Page number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of the `IsPdf` parameter is `true`. Default value: 1.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// Whether to return the character information. Default value: `false`
	IsWords *bool `json:"IsWords,omitnil,omitempty" name:"IsWords"`
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
	TextDetections []*TextDetection `json:"TextDetections,omitnil,omitempty" name:"TextDetections"`

	// Detected language. For more information on the supported languages, please see the description of the `LanguageType` input parameter.
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// Image rotation angle in degrees. 0°: The horizontal direction of the text on the image; a positive value: rotate clockwise; a negative value: rotate counterclockwise.
	Angel *float64 `json:"Angel,omitnil,omitempty" name:"Angel"`

	// Total number of PDF pages to be returned if the image is a PDF. Default value: 0.
	PdfPageSize *int64 `json:"PdfPageSize,omitnil,omitempty" name:"PdfPageSize"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Specification
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// Unit
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// Unit price
	Price *string `json:"Price,omitnil,omitempty" name:"Price"`

	// Amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`

	// Tax amount
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`
}

type GroupInfo struct {
	// The elements in each line.
	Groups []*LineInfo `json:"Groups,omitnil,omitempty" name:"Groups"`
}

// Predefined struct for user
type HKIDCardOCRRequestParams struct {
	// Whether to return identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// Whether to check for authenticity.
	//
	// Deprecated: DetectFake is deprecated.
	DetectFake *bool `json:"DetectFake,omitnil,omitempty" name:"DetectFake"`

	// Base64 string of the image
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported yet.
	// Supported image size: The downloaded image cannot exceed 7 MB after being Base64-encoded, and it cannot take longer than 3 seconds to download the image.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type HKIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// Whether to return identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// Whether to check for authenticity.
	DetectFake *bool `json:"DetectFake,omitnil,omitempty" name:"DetectFake"`

	// Base64 string of the image
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported yet.
	// Supported image size: The downloaded image cannot exceed 7 MB after being Base64-encoded, and it cannot take longer than 3 seconds to download the image.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
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
	delete(f, "ReturnHeadImage")
	delete(f, "DetectFake")
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
	CnName *string `json:"CnName,omitnil,omitempty" name:"CnName"`

	// Name in English
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// Telecode for the name in Chinese
	TelexCode *string `json:"TelexCode,omitnil,omitempty" name:"TelexCode"`

	// Gender. Valid values: Male, Female
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Date of birth
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Permanent identity card.
	// 0: non-permanent;
	// 1: permanent;
	// -1: unknown.
	Permanent *int64 `json:"Permanent,omitnil,omitempty" name:"Permanent"`

	// Identity card number
	IdNum *string `json:"IdNum,omitnil,omitempty" name:"IdNum"`

	// Document symbol, i.e., the symbol under the date of birth, such as "***AZ"
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// First issue date
	FirstIssueDate *string `json:"FirstIssueDate,omitnil,omitempty" name:"FirstIssueDate"`

	// Last receipt date
	CurrentIssueDate *string `json:"CurrentIssueDate,omitnil,omitempty" name:"CurrentIssueDate"`

	// Authenticity check.
	// 0: unable to judge (because the image is blurred, incomplete, reflective, too dark, etc.);
	// 1: forged;
	// 2: authentic.
	// Note: this field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: FakeDetectResult is deprecated.
	FakeDetectResult *int64 `json:"FakeDetectResult,omitnil,omitempty" name:"FakeDetectResult"`

	// Base64-encoded identity photo
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeadImage *string `json:"HeadImage,omitnil,omitempty" name:"HeadImage"`

	// Multiple alarm codes. If the ID card is spoofed, photocopied, or photoshopped, the corresponding alarm code will be returned.
	// -9102: Alarm for photocopied document
	// -9103: Alarm for spoofed document
	//
	// Deprecated: WarningCode is deprecated.
	WarningCode []*int64 `json:"WarningCode,omitnil,omitempty" name:"WarningCode"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9104 Alarm for PS certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// `FRONT`: The side with the profile photo.
	// `BACK`: The side with the national emblem.
	// If this parameter is not specified, the system will automatically determine the ID card side.
	CardSide *string `json:"CardSide,omitnil,omitempty" name:"CardSide"`
}

type HmtResidentPermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// `FRONT`: The side with the profile photo.
	// `BACK`: The side with the national emblem.
	// If this parameter is not specified, the system will automatically determine the ID card side.
	CardSide *string `json:"CardSide,omitnil,omitempty" name:"CardSide"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Date of birth
	Birth *string `json:"Birth,omitnil,omitempty" name:"Birth"`

	// Address
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// ID card number
	IdCardNo *string `json:"IdCardNo,omitnil,omitempty" name:"IdCardNo"`

	// 0: Front side.
	// 1: Back side.
	CardType *int64 `json:"CardType,omitnil,omitempty" name:"CardType"`

	// Validity period
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// Issuing authority
	Authority *string `json:"Authority,omitnil,omitempty" name:"Authority"`

	// Number of issues
	VisaNum *string `json:"VisaNum,omitnil,omitempty" name:"VisaNum"`

	// Permit number
	PassNo *string `json:"PassNo,omitnil,omitempty" name:"PassNo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// `FRONT`: The side with the profile photo.
	// `BACK`: The side with the national emblem.
	// If this parameter is not specified, the system will automatically determine the ID card side.
	CardSide *string `json:"CardSide,omitnil,omitempty" name:"CardSide"`

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
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

type IDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// `FRONT`: The side with the profile photo.
	// `BACK`: The side with the national emblem.
	// If this parameter is not specified, the system will automatically determine the ID card side.
	CardSide *string `json:"CardSide,omitnil,omitempty" name:"CardSide"`

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
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Gender (profile photo side)
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Ethnicity (profile photo side)
	Nation *string `json:"Nation,omitnil,omitempty" name:"Nation"`

	// Date of birth (profile photo side)
	Birth *string `json:"Birth,omitnil,omitempty" name:"Birth"`

	// Address (profile photo side)
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// ID number (profile photo side)
	IdNum *string `json:"IdNum,omitnil,omitempty" name:"IdNum"`

	// Issuing authority (national emblem side)
	Authority *string `json:"Authority,omitnil,omitempty" name:"Authority"`

	// Validity period (national emblem side)
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

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
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

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
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// The coordinates of the four vertices of the rotated image.
	Polygon *Polygon `json:"Polygon,omitnil,omitempty" name:"Polygon"`

	// The rotation angle of the recognized image in the image with multiple types of invoices/tickets.
	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// The recognized content.
	SingleInvoiceInfos *SingleInvoiceItem `json:"SingleInvoiceInfos,omitnil,omitempty" name:"SingleInvoiceInfos"`

	// The number of the page on which the recognized invoice is in the image or PDF file, starting from 1 by default.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// The detailed invoice type. See the description of `SubType`.
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// The invoice description. See the description of `TypeDescription`.
	TypeDescription *string `json:"TypeDescription,omitnil,omitempty" name:"TypeDescription"`

	// The image file after cropping, encoded in Base64. This is returned if `EnableCutImage` is set to `true`.
	CutImage *string `json:"CutImage,omitnil,omitempty" name:"CutImage"`

	// The description of the detailed invoice type. See the description of `SubType`.
	SubTypeDescription *string `json:"SubTypeDescription,omitnil,omitempty" name:"SubTypeDescription"`
}

type ItemCoord struct {
	// X-coordinate of top-left point.
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Y-coordinate of top-left point.
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// Width
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Height
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type ItemInfo struct {
	// The key information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Key *Key `json:"Key,omitnil,omitempty" name:"Key"`

	// The value information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *Value `json:"Value,omitnil,omitempty" name:"Value"`
}

type Key struct {
	// The name of the recognized field.
	AutoName *string `json:"AutoName,omitnil,omitempty" name:"AutoName"`

	// The name of a defined field (the key passed in).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`
}

type LicensePlateInfo struct {
	// The recognized license plate number.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// The confidence score (0–100).
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The bounding box coordinates of the text line in the original image.
	Rect *Rect `json:"Rect,omitnil,omitempty" name:"Rect"`

	// The recognized license plate color, which currently includes "white", "black", "blue", "green", "yellow", "yellow-green", and "temporary plate".
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`
}

// Predefined struct for user
type LicensePlateOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type LicensePlateOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
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
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// The confidence score (0–100).
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The bounding box coordinates of the text line in the original image.
	Rect *Rect `json:"Rect,omitnil,omitempty" name:"Rect"`

	// The recognized license plate color, which currently includes "white", "black", "blue", "green", "yellow", "yellow-green", and "temporary plate".
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`

	// The vehicle license plate information.
	LicensePlateInfos []*LicensePlateInfo `json:"LicensePlateInfos,omitnil,omitempty" name:"LicensePlateInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Lines []*ItemInfo `json:"Lines,omitnil,omitempty" name:"Lines"`
}

// Predefined struct for user
type MLIDCardOCRRequestParams struct {
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL of an image. (This field is not available outside the Chinese mainland.)
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. Storing images in Tencent Cloud URLs can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to return an image. Default value: `false`.
	RetImage *bool `json:"RetImage,omitnil,omitempty" name:"RetImage"`
}

type MLIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image cannot exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL of an image. (This field is not available outside the Chinese mainland.)
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. Storing images in Tencent Cloud URLs can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to return an image. Default value: `false`.
	RetImage *bool `json:"RetImage,omitnil,omitempty" name:"RetImage"`
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
	delete(f, "BackImageBase64")
	delete(f, "ImageUrl")
	delete(f, "BackImageUrl")
	delete(f, "RetImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MLIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MLIDCardOCRResponseParams struct {
	// ID number
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Full name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Address
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Alarm codes
	// -9103 Alarm for photographed certificate
	// -9102 Alarm for photocopied certificate
	// -9106 Alarm for covered certificate
	// -9107 Alarm for blurry image
	//
	// Deprecated: Warn is deprecated.
	Warn []*int64 `json:"Warn,omitnil,omitempty" name:"Warn"`

	// Identity photo
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// This is an extended field, 
	// with the confidence of a field recognition result returned in the following format.
	// {
	//   Field name:{
	//     Confidence:0.9999
	//   }
	// }
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// Certificate type
	// MyKad  ID card
	// MyPR    Permanent resident card
	// MyTentera   Military identity card
	// MyKAS    Temporary ID card
	// POLIS  Police card
	// IKAD   Work permit
	// MyKid   Kid card
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Date of birth. This field is available only for work permits (i-Kad) and ID cards (MyKad).
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Number on the back of Malaysia ID card 
	MyKadNumber *string `json:"MyKadNumber,omitnil,omitempty" name:"MyKadNumber"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9104 Alarm for PS certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Base64-encoded value of image. The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 500x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported. It is recommended that the card part occupies more than 2/3 area of the image.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Whether to return an image. 
	// Default value: false.
	RetImage *bool `json:"RetImage,omitnil,omitempty" name:"RetImage"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG, BMP, PDF.
	// Supported image size: the downloaded image cannot exceed 7 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type MLIDPassportOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image. The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 500x800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported. It is recommended that the card part occupies more than 2/3 area of the image.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Whether to return an image. 
	// Default value: false.
	RetImage *bool `json:"RetImage,omitnil,omitempty" name:"RetImage"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG, BMP, PDF.
	// Supported image size: the downloaded image cannot exceed 7 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
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
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MLIDPassportOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MLIDPassportOCRResponseParams struct {
	// Passport ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Date of birth
	DateOfBirth *string `json:"DateOfBirth,omitnil,omitempty" name:"DateOfBirth"`

	// Gender (F: female, M: male)
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Expiration date
	DateOfExpiration *string `json:"DateOfExpiration,omitnil,omitempty" name:"DateOfExpiration"`

	// Issuing country
	IssuingCountry *string `json:"IssuingCountry,omitnil,omitempty" name:"IssuingCountry"`

	// Country/region code
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Alarm codes
	// -9102 Alarm for photocopy on a paper document (including black & white and color ones)
	// -9103 Alarm for photocopy on an electronic device
	// -9106 Alarm for covered card
	Warn []*int64 `json:"Warn,omitnil,omitempty" name:"Warn"`

	// Identity photo
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Extended field:
	// {
	//     ID:{
	//         Confidence:0.9999
	//     },
	//     Name:{
	//         Confidence:0.9996
	//     }
	// }
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// The first row of the machine-readable zone (MRZ) at the bottom
	CodeSet *string `json:"CodeSet,omitnil,omitempty" name:"CodeSet"`

	// The second row of the MRZ at the bottom
	CodeCrc *string `json:"CodeCrc,omitnil,omitempty" name:"CodeCrc"`

	// The surname.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Surname *string `json:"Surname,omitnil,omitempty" name:"Surname"`

	// The given name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GivenName *string `json:"GivenName,omitnil,omitempty" name:"GivenName"`

	// Type (in Machine Readable Zone)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Document content in Information Zone
	PassportRecognizeInfos *PassportRecognizeInfos `json:"PassportRecognizeInfos,omitnil,omitempty" name:"PassportRecognizeInfos"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9104 Alarm for PS certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Time
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Check code
	CheckCode *string `json:"CheckCode,omitnil,omitempty" name:"CheckCode"`

	// Ciphertext
	Ciphertext *string `json:"Ciphertext,omitnil,omitempty" name:"Ciphertext"`

	// Category
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// Amount before tax
	PretaxAmount *string `json:"PretaxAmount,omitnil,omitempty" name:"PretaxAmount"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Industry
	IndustryClass *string `json:"IndustryClass,omitnil,omitempty" name:"IndustryClass"`

	// Seller's name
	Seller *string `json:"Seller,omitnil,omitempty" name:"Seller"`

	// Seller's taxpayer identification number
	SellerTaxID *string `json:"SellerTaxID,omitnil,omitempty" name:"SellerTaxID"`

	// Seller's address and phone number
	SellerAddrTel *string `json:"SellerAddrTel,omitnil,omitempty" name:"SellerAddrTel"`

	// Seller's bank account number
	SellerBankAccount *string `json:"SellerBankAccount,omitnil,omitempty" name:"SellerBankAccount"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil,omitempty" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil,omitempty" name:"BuyerTaxID"`

	// Buyer's address and phone number
	BuyerAddrTel *string `json:"BuyerAddrTel,omitnil,omitempty" name:"BuyerAddrTel"`

	// Buyer's bank account number
	BuyerBankAccount *string `json:"BuyerBankAccount,omitnil,omitempty" name:"BuyerBankAccount"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil,omitempty" name:"CompanySealMark"`

	// Whether it is a general machine-printed invoice issued by Zhejiang or Guangdong province (0: No, 1: Yes)
	ElectronicMark *int64 `json:"ElectronicMark,omitnil,omitempty" name:"ElectronicMark"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil,omitempty" name:"Issuer"`

	// Payee
	Receiptor *string `json:"Receiptor,omitnil,omitempty" name:"Receiptor"`

	// Reviewer
	Reviewer *string `json:"Reviewer,omitnil,omitempty" name:"Reviewer"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Operator's payment information
	PaymentInfo *string `json:"PaymentInfo,omitnil,omitempty" name:"PaymentInfo"`

	// Operator-assigned invoice pickup user
	TicketPickupUser *string `json:"TicketPickupUser,omitnil,omitempty" name:"TicketPickupUser"`

	// Operator's merchant number
	MerchantNumber *string `json:"MerchantNumber,omitnil,omitempty" name:"MerchantNumber"`

	// Operator's order number
	OrderNumber *string `json:"OrderNumber,omitnil,omitempty" name:"OrderNumber"`

	// Items
	GeneralMachineItems []*GeneralMachineItem `json:"GeneralMachineItems,omitnil,omitempty" name:"GeneralMachineItems"`
}

// Predefined struct for user
type MainlandPermitOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the ID photo. By default, the ID photo is not returned.
	RetProfile *bool `json:"RetProfile,omitnil,omitempty" name:"RetProfile"`
}

type MainlandPermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the ID photo. By default, the ID photo is not returned.
	RetProfile *bool `json:"RetProfile,omitnil,omitempty" name:"RetProfile"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Name in English
	EnglishName *string `json:"EnglishName,omitnil,omitempty" name:"EnglishName"`

	// Gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Date of birth
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Issuing authority
	IssueAuthority *string `json:"IssueAuthority,omitnil,omitempty" name:"IssueAuthority"`

	// Validity period
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// ID number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Place of issue
	IssueAddress *string `json:"IssueAddress,omitnil,omitempty" name:"IssueAddress"`

	// Number of issues
	IssueNumber *string `json:"IssueNumber,omitnil,omitempty" name:"IssueNumber"`

	// Document type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Base64-encoded profile photo, which is returned only when `RetProfile` is set to `True`
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

	Title *string `json:"Title,omitnil,omitempty" name:"Title"`


	Code *string `json:"Code,omitnil,omitempty" name:"Code"`


	Number *string `json:"Number,omitnil,omitempty" name:"Number"`


	Total *string `json:"Total,omitnil,omitempty" name:"Total"`


	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`


	Date *string `json:"Date,omitnil,omitempty" name:"Date"`


	CheckCode *string `json:"CheckCode,omitnil,omitempty" name:"CheckCode"`


	Place *string `json:"Place,omitnil,omitempty" name:"Place"`


	Reviewer *string `json:"Reviewer,omitnil,omitempty" name:"Reviewer"`
}

type MotorVehicleSaleInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Amount before tax
	PretaxAmount *string `json:"PretaxAmount,omitnil,omitempty" name:"PretaxAmount"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`

	// Seller's name
	Seller *string `json:"Seller,omitnil,omitempty" name:"Seller"`

	// Seller's company code
	SellerTaxID *string `json:"SellerTaxID,omitnil,omitempty" name:"SellerTaxID"`

	// Seller's phone number
	SellerTel *string `json:"SellerTel,omitnil,omitempty" name:"SellerTel"`

	// Seller's address
	SellerAddress *string `json:"SellerAddress,omitnil,omitempty" name:"SellerAddress"`

	// Seller's account opening bank
	SellerBank *string `json:"SellerBank,omitnil,omitempty" name:"SellerBank"`

	// Seller's bank account number
	SellerBankAccount *string `json:"SellerBankAccount,omitnil,omitempty" name:"SellerBankAccount"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil,omitempty" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil,omitempty" name:"BuyerTaxID"`

	// Buyer's ID number/organization code
	BuyerID *string `json:"BuyerID,omitnil,omitempty" name:"BuyerID"`

	// Tax authority
	TaxAuthorities *string `json:"TaxAuthorities,omitnil,omitempty" name:"TaxAuthorities"`

	// Code of the tax authority
	TaxAuthoritiesCode *string `json:"TaxAuthoritiesCode,omitnil,omitempty" name:"TaxAuthoritiesCode"`

	// VIN
	VIN *string `json:"VIN,omitnil,omitempty" name:"VIN"`

	// Vehicle model
	VehicleModel *string `json:"VehicleModel,omitnil,omitempty" name:"VehicleModel"`

	// Engine No.
	VehicleEngineCode *string `json:"VehicleEngineCode,omitnil,omitempty" name:"VehicleEngineCode"`

	// No. of the certificate of conformity
	CertificateNumber *string `json:"CertificateNumber,omitnil,omitempty" name:"CertificateNumber"`

	// Inspection No.
	InspectionNumber *string `json:"InspectionNumber,omitnil,omitempty" name:"InspectionNumber"`

	// Machine No.
	MachineID *string `json:"MachineID,omitnil,omitempty" name:"MachineID"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil,omitempty" name:"VehicleType"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil,omitempty" name:"CompanySealMark"`

	// Tonnage
	Tonnage *string `json:"Tonnage,omitnil,omitempty" name:"Tonnage"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Form type
	FormType *string `json:"FormType,omitnil,omitempty" name:"FormType"`

	// Form name
	FormName *string `json:"FormName,omitnil,omitempty" name:"FormName"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil,omitempty" name:"Issuer"`

	// Tax payment voucher number
	TaxNum *string `json:"TaxNum,omitnil,omitempty" name:"TaxNum"`

	// Passenger capacity
	MaxPeopleNum *string `json:"MaxPeopleNum,omitnil,omitempty" name:"MaxPeopleNum"`

	// Origin
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// Machine-printed invoice code
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// Machine-printed invoice number
	MachineNumber *string `json:"MachineNumber,omitnil,omitempty" name:"MachineNumber"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`
}

type NonTaxIncomeBill struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Check code
	CheckCode *string `json:"CheckCode,omitnil,omitempty" name:"CheckCode"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`

	// Payer's name
	Buyer *string `json:"Buyer,omitnil,omitempty" name:"Buyer"`

	// Payer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil,omitempty" name:"BuyerTaxID"`

	// Payee's name
	Seller *string `json:"Seller,omitnil,omitempty" name:"Seller"`

	// Payee's company name
	SellerCompany *string `json:"SellerCompany,omitnil,omitempty" name:"SellerCompany"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Currency
	CurrencyCode *string `json:"CurrencyCode,omitnil,omitempty" name:"CurrencyCode"`

	// Reviewer
	Reviewer *string `json:"Reviewer,omitnil,omitempty" name:"Reviewer"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Other information
	OtherInfo *string `json:"OtherInfo,omitnil,omitempty" name:"OtherInfo"`

	// Payment code
	PaymentCode *string `json:"PaymentCode,omitnil,omitempty" name:"PaymentCode"`

	// Collecting organization's code
	ReceiveUnitCode *string `json:"ReceiveUnitCode,omitnil,omitempty" name:"ReceiveUnitCode"`

	// Collecting organization's name
	Receiver *string `json:"Receiver,omitnil,omitempty" name:"Receiver"`

	// Operator
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Payer's account
	PayerAccount *string `json:"PayerAccount,omitnil,omitempty" name:"PayerAccount"`

	// Payer's account opening bank
	PayerBank *string `json:"PayerBank,omitnil,omitempty" name:"PayerBank"`

	// Payee's account
	ReceiverAccount *string `json:"ReceiverAccount,omitnil,omitempty" name:"ReceiverAccount"`

	// Payee's account opening bank
	ReceiverBank *string `json:"ReceiverBank,omitnil,omitempty" name:"ReceiverBank"`

	// Items
	NonTaxItems []*NonTaxItem `json:"NonTaxItems,omitnil,omitempty" name:"NonTaxItems"`
}

type NonTaxItem struct {
	// Item code
	ItemID *string `json:"ItemID,omitnil,omitempty" name:"ItemID"`

	// Item name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Unit
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// Standard
	Standard *string `json:"Standard,omitnil,omitempty" name:"Standard"`

	// Amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`
}

type OtherInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// List
	OtherInvoiceListItems []*OtherInvoiceItem `json:"OtherInvoiceListItems,omitnil,omitempty" name:"OtherInvoiceListItems"`

	// Table
	OtherInvoiceTableItems []*OtherInvoiceList `json:"OtherInvoiceTableItems,omitnil,omitempty" name:"OtherInvoiceTableItems"`
}

type OtherInvoiceItem struct {
	// Field name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Field value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type OtherInvoiceList struct {
	// List
	OtherInvoiceItemList []*OtherInvoiceItem `json:"OtherInvoiceItemList,omitnil,omitempty" name:"OtherInvoiceItemList"`
}

type PassportRecognizeInfos struct {

	Type *string `json:"Type,omitnil,omitempty" name:"Type"`


	IssuingCountry *string `json:"IssuingCountry,omitnil,omitempty" name:"IssuingCountry"`


	PassportID *string `json:"PassportID,omitnil,omitempty" name:"PassportID"`


	Surname *string `json:"Surname,omitnil,omitempty" name:"Surname"`


	GivenName *string `json:"GivenName,omitnil,omitempty" name:"GivenName"`


	Name *string `json:"Name,omitnil,omitempty" name:"Name"`


	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`


	DateOfBirth *string `json:"DateOfBirth,omitnil,omitempty" name:"DateOfBirth"`


	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`


	DateOfIssuance *string `json:"DateOfIssuance,omitnil,omitempty" name:"DateOfIssuance"`


	DateOfExpiration *string `json:"DateOfExpiration,omitnil,omitempty" name:"DateOfExpiration"`


	Signature *string `json:"Signature,omitnil,omitempty" name:"Signature"`


	IssuePlace *string `json:"IssuePlace,omitnil,omitempty" name:"IssuePlace"`


	IssuingAuthority *string `json:"IssuingAuthority,omitnil,omitempty" name:"IssuingAuthority"`
}

// Predefined struct for user
type PermitOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the profile photo, default is false.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
}

type PermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the profile photo, default is false.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
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
	delete(f, "CropPortrait")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PermitOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PermitOCRResponseParams struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Name in English
	EnglishName *string `json:"EnglishName,omitnil,omitempty" name:"EnglishName"`

	// ID number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Validity period
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// Issuing authority
	IssueAuthority *string `json:"IssueAuthority,omitnil,omitempty" name:"IssueAuthority"`

	// Place of issue
	IssueAddress *string `json:"IssueAddress,omitnil,omitempty" name:"IssueAddress"`

	// Date of birth
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Base64 of profile picture
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// return type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	LeftTop *Coord `json:"LeftTop,omitnil,omitempty" name:"LeftTop"`

	// The coordinates of the upper-right vertex.
	RightTop *Coord `json:"RightTop,omitnil,omitempty" name:"RightTop"`

	// The coordinates of the lower-left vertex.
	RightBottom *Coord `json:"RightBottom,omitnil,omitempty" name:"RightBottom"`

	// The coordinates of the lower-right vertex.
	LeftBottom *Coord `json:"LeftBottom,omitnil,omitempty" name:"LeftBottom"`
}

type QuotaInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil,omitempty" name:"CompanySealMark"`
}

// Predefined struct for user
type RecognizeBrazilDriverLicenseOCRRequestParams struct {
	// The Base64 value of the image. It is required that the image after Base64 encoding should not exceed 7M, the resolution is recommended to be 500*800 or above, and PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The Base64 value of the image. It is required that the image after Base64 encoding should not exceed 7M, the resolution is recommended to be 500*800 or above, and PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL address of the image. It is required that the image after Base64 encoding should not exceed 7M, the resolution is recommended to be 500*800 or above, and PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. It is recommended that images be stored in Tencent Cloud to ensure higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The URL address of the image. It is required that the image after Base64 encoding should not exceed 7M, the resolution is recommended to be 500*800 or above, and PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. It is recommended that images be stored in Tencent Cloud to ensure higher download speed and stability.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Picture switch. The default is false, and the base64 encoding of the avatar photo is not returned. When set to true, the base64 encoding of the portrait photo is returned.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
}

type RecognizeBrazilDriverLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64 value of the image. It is required that the image after Base64 encoding should not exceed 7M, the resolution is recommended to be 500*800 or above, and PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The Base64 value of the image. It is required that the image after Base64 encoding should not exceed 7M, the resolution is recommended to be 500*800 or above, and PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL address of the image. It is required that the image after Base64 encoding should not exceed 7M, the resolution is recommended to be 500*800 or above, and PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. It is recommended that images be stored in Tencent Cloud to ensure higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The URL address of the image. It is required that the image after Base64 encoding should not exceed 7M, the resolution is recommended to be 500*800 or above, and PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. It is recommended that images be stored in Tencent Cloud to ensure higher download speed and stability.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Picture switch. The default is false, and the base64 encoding of the avatar photo is not returned. When set to true, the base64 encoding of the portrait photo is returned.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
}

func (r *RecognizeBrazilDriverLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilDriverLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "BackImageBase64")
	delete(f, "ImageUrl")
	delete(f, "BackImageUrl")
	delete(f, "CropPortrait")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeBrazilDriverLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeBrazilDriverLicenseOCRResponseParams struct {
	// Name
	NOME *string `json:"NOME,omitnil,omitempty" name:"NOME"`

	// Driving license type
	CatHab *string `json:"CatHab,omitnil,omitempty" name:"CatHab"`

	// CNH number
	CNHNumber *string `json:"CNHNumber,omitnil,omitempty" name:"CNHNumber"`

	// Valid date
	VALIDADE *string `json:"VALIDADE,omitnil,omitempty" name:"VALIDADE"`

	// Qualification
	QUALIFICATION *string `json:"QUALIFICATION,omitnil,omitempty" name:"QUALIFICATION"`

	// ID number
	IDENTIDADE *string `json:"IDENTIDADE,omitnil,omitempty" name:"IDENTIDADE"`

	// CPF
	CPF *string `json:"CPF,omitnil,omitempty" name:"CPF"`

	// Birthday
	NASCIMENTO *string `json:"NASCIMENTO,omitnil,omitempty" name:"NASCIMENTO"`

	// Membership
	MEMBERSHIP *string `json:"MEMBERSHIP,omitnil,omitempty" name:"MEMBERSHIP"`

	// Registration number
	REGISTRO *string `json:"REGISTRO,omitnil,omitempty" name:"REGISTRO"`

	// Remark
	OBSERVATIONS *string `json:"OBSERVATIONS,omitnil,omitempty" name:"OBSERVATIONS"`

	// Issue date
	IssueDate *string `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`

	// Issue location
	LOCAL *string `json:"LOCAL,omitnil,omitempty" name:"LOCAL"`

	// Number in the back of the card
	BackNumber *string `json:"BackNumber,omitnil,omitempty" name:"BackNumber"`

	// Field confidence
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// PortraitImage base64
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeBrazilDriverLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeBrazilDriverLicenseOCRResponseParams `json:"Response"`
}

func (r *RecognizeBrazilDriverLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilDriverLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeBrazilIDCardOCRRequestParams struct {
	// Base64 value of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. Image download time does not exceed 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. URLs of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. An ImageUrl and ImageBase64 must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download time does not exceed 3 seconds. The URL of the image stored in Tencent Cloud can ensure higher download speed and stability. It is recommended to store the image in Tencent Cloud. The speed and stability of the URL stored outside Tencent Cloud may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizeBrazilIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64 value of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. Image download time does not exceed 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. URLs of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. An ImageUrl and ImageBase64 must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download time does not exceed 3 seconds. The URL of the image stored in Tencent Cloud can ensure higher download speed and stability. It is recommended to store the image in Tencent Cloud. The speed and stability of the URL stored outside Tencent Cloud may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

func (r *RecognizeBrazilIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "BackImageBase64")
	delete(f, "BackImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeBrazilIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeBrazilIDCardOCRResponseParams struct {
	// Name
	Nome *string `json:"Nome,omitnil,omitempty" name:"Nome"`

	// Family information
	MemberShip *string `json:"MemberShip,omitnil,omitempty" name:"MemberShip"`

	// Birthday
	DataNascimento *string `json:"DataNascimento,omitnil,omitempty" name:"DataNascimento"`

	// Issuing agency
	IssuingAgency *string `json:"IssuingAgency,omitnil,omitempty" name:"IssuingAgency"`

	// blood type
	Fatorrh *string `json:"Fatorrh,omitnil,omitempty" name:"Fatorrh"`

	// Birth place
	NaturalIDade *string `json:"NaturalIDade,omitnil,omitempty" name:"NaturalIDade"`

	// Additional information 
	Observations *string `json:"Observations,omitnil,omitempty" name:"Observations"`

	// CPF
	CPF *string `json:"CPF,omitnil,omitempty" name:"CPF"`

	// DNI
	DNI *string `json:"DNI,omitnil,omitempty" name:"DNI"`

	// universal registration
	RegistroGeral *string `json:"RegistroGeral,omitnil,omitempty" name:"RegistroGeral"`

	// Issue date
	DispatchDate *string `json:"DispatchDate,omitnil,omitempty" name:"DispatchDate"`

	// address
	Registro *string `json:"Registro,omitnil,omitempty" name:"Registro"`

	// Portrait image
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// Original identity information
	DocOrigem *string `json:"DocOrigem,omitnil,omitempty" name:"DocOrigem"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeBrazilIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeBrazilIDCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeBrazilIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeBrazilRNEOCRRequestParams struct {
	// Base64 value of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image downloading time should not exceed 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. URLs of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. An ImageUrl and ImageBase64 must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download time does not exceed 3 seconds. The URL of the image stored in Tencent Cloud can ensure higher download speed and stability. It is recommended to store the image in Tencent Cloud. The speed and stability of the URL stored outside Tencent Cloud may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizeBrazilRNEOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64 value of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image downloading time should not exceed 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. URLs of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. An ImageUrl and ImageBase64 must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download time does not exceed 3 seconds. The URL of the image stored in Tencent Cloud can ensure higher download speed and stability. It is recommended to store the image in Tencent Cloud. The speed and stability of the URL stored outside Tencent Cloud may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

func (r *RecognizeBrazilRNEOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilRNEOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "BackImageBase64")
	delete(f, "BackImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeBrazilRNEOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeBrazilRNEOCRResponseParams struct {
	// RNE
	RNE *string `json:"RNE,omitnil,omitempty" name:"RNE"`

	// Classification
	CLASSIFICATION *string `json:"CLASSIFICATION,omitnil,omitempty" name:"CLASSIFICATION"`

	// Valid date
	VALIDADE *string `json:"VALIDADE,omitnil,omitempty" name:"VALIDADE"`

	// Name
	NOME *string `json:"NOME,omitnil,omitempty" name:"NOME"`

	// Family information
	Membership *string `json:"Membership,omitnil,omitempty" name:"Membership"`

	// Nationality
	NACIONALIDADE *string `json:"NACIONALIDADE,omitnil,omitempty" name:"NACIONALIDADE"`

	// Place of Birth
	NATURALIDADE *string `json:"NATURALIDADE,omitnil,omitempty" name:"NATURALIDADE"`

	// Issuing agency
	IssuingAgency *string `json:"IssuingAgency,omitnil,omitempty" name:"IssuingAgency"`

	// Birthday
	DateOfBirth *string `json:"DateOfBirth,omitnil,omitempty" name:"DateOfBirth"`

	// Gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Date of entry
	EntryDate *string `json:"EntryDate,omitnil,omitempty" name:"EntryDate"`

	// VIA
	VIA *string `json:"VIA,omitnil,omitempty" name:"VIA"`

	// Dispatch date
	DispatchDate *string `json:"DispatchDate,omitnil,omitempty" name:"DispatchDate"`

	// MRZ
	MRZ *string `json:"MRZ,omitnil,omitempty" name:"MRZ"`

	// PortraitImage base64
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeBrazilRNEOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeBrazilRNEOCRResponseParams `json:"Response"`
}

func (r *RecognizeBrazilRNEOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilRNEOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeBrazilRNMOCRRequestParams struct {
	// Base64 value of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image downloading time should not exceed 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. URLs of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. An ImageUrl and ImageBase64 must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download time does not exceed 3 seconds. The URL of the image stored in Tencent Cloud can ensure higher download speed and stability. It is recommended to store the image in Tencent Cloud. The speed and stability of the URL stored outside Tencent Cloud may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizeBrazilRNMOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64 value of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image downloading time should not exceed 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. URLs of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. An ImageUrl and ImageBase64 must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download time does not exceed 3 seconds. The URL of the image stored in Tencent Cloud can ensure higher download speed and stability. It is recommended to store the image in Tencent Cloud. The speed and stability of the URL stored outside Tencent Cloud may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

func (r *RecognizeBrazilRNMOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilRNMOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "BackImageBase64")
	delete(f, "BackImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeBrazilRNMOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeBrazilRNMOCRResponseParams struct {
	// Last name
	SOBRENOME *string `json:"SOBRENOME,omitnil,omitempty" name:"SOBRENOME"`

	// First name
	NOME *string `json:"NOME,omitnil,omitempty" name:"NOME"`

	// Date of Birth
	DATADENASCIMENTO *string `json:"DATADENASCIMENTO,omitnil,omitempty" name:"DATADENASCIMENTO"`

	// Gender
	SEXO *string `json:"SEXO,omitnil,omitempty" name:"SEXO"`

	// Parents name
	MEMBERSHIP *string `json:"MEMBERSHIP,omitnil,omitempty" name:"MEMBERSHIP"`

	// Nationality
	NACIONALIDADE *string `json:"NACIONALIDADE,omitnil,omitempty" name:"NACIONALIDADE"`

	// Expiry Date
	VALIDADE *string `json:"VALIDADE,omitnil,omitempty" name:"VALIDADE"`

	// RNM
	RNM *string `json:"RNM,omitnil,omitempty" name:"RNM"`

	// CPF
	CPF *string `json:"CPF,omitnil,omitempty" name:"CPF"`


	CLASSIFICATION *string `json:"CLASSIFICATION,omitnil,omitempty" name:"CLASSIFICATION"`


	PRAZODERESIDENCIA *string `json:"PRAZODERESIDENCIA,omitnil,omitempty" name:"PRAZODERESIDENCIA"`

	// Issue Date
	ISSUANCE *string `json:"ISSUANCE,omitnil,omitempty" name:"ISSUANCE"`

	// Legal basis
	AMPAROLEGAL *string `json:"AMPAROLEGAL,omitnil,omitempty" name:"AMPAROLEGAL"`

	// Machine readable zone code
	MRZ *string `json:"MRZ,omitnil,omitempty" name:"MRZ"`

	// PortraitImage
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// PortraitImage(Back)
	PortraitImageBack *string `json:"PortraitImageBack,omitnil,omitempty" name:"PortraitImageBack"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeBrazilRNMOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeBrazilRNMOCRResponseParams `json:"Response"`
}

func (r *RecognizeBrazilRNMOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilRNMOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeGeneralInvoiceRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, JPEG, and PDF. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Supported image pixels: 20 to 10,000
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, JPEG, and PDF. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Supported image pixels: 20 to 10,000
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

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
	Types []*int64 `json:"Types,omitnil,omitempty" name:"Types"`

	// Whether to enable recognition of other invoices. If you enable this feature, other invoices can be recognized. Default value: `true`.	
	EnableOther *bool `json:"EnableOther,omitnil,omitempty" name:"EnableOther"`

	// Whether to enable PDF recognition. If you enable this feature, both images and PDF files can be recognized. Default value: `true`.
	EnablePdf *bool `json:"EnablePdf,omitnil,omitempty" name:"EnablePdf"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `EnablePdf` is `true`. Default value: 1.
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// Whether to enable multi-page PDF recognition. If you enable this feature, multiple pages of a PDF file can be recognized, and the recognition results of a maximum of the first 30 pages can be returned. After you enable this feature, input parameters `EnablePdf` and `PdfPageNumber` are invalid. Default value: `false`.
	EnableMultiplePage *bool `json:"EnableMultiplePage,omitnil,omitempty" name:"EnableMultiplePage"`

	// Whether to return the Base64-encoded value of the cropped image. Default value: `false`.
	EnableCutImage *bool `json:"EnableCutImage,omitnil,omitempty" name:"EnableCutImage"`
}

type RecognizeGeneralInvoiceRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, JPEG, and PDF. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Supported image pixels: 20 to 10,000
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, JPEG, and PDF. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Supported image pixels: 20 to 10,000
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

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
	Types []*int64 `json:"Types,omitnil,omitempty" name:"Types"`

	// Whether to enable recognition of other invoices. If you enable this feature, other invoices can be recognized. Default value: `true`.	
	EnableOther *bool `json:"EnableOther,omitnil,omitempty" name:"EnableOther"`

	// Whether to enable PDF recognition. If you enable this feature, both images and PDF files can be recognized. Default value: `true`.
	EnablePdf *bool `json:"EnablePdf,omitnil,omitempty" name:"EnablePdf"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `EnablePdf` is `true`. Default value: 1.
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// Whether to enable multi-page PDF recognition. If you enable this feature, multiple pages of a PDF file can be recognized, and the recognition results of a maximum of the first 30 pages can be returned. After you enable this feature, input parameters `EnablePdf` and `PdfPageNumber` are invalid. Default value: `false`.
	EnableMultiplePage *bool `json:"EnableMultiplePage,omitnil,omitempty" name:"EnableMultiplePage"`

	// Whether to return the Base64-encoded value of the cropped image. Default value: `false`.
	EnableCutImage *bool `json:"EnableCutImage,omitnil,omitempty" name:"EnableCutImage"`
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
	MixedInvoiceItems []*InvoiceItem `json:"MixedInvoiceItems,omitnil,omitempty" name:"MixedInvoiceItems"`

	// Total number of pages in the PDF file.
	TotalPDFCount *int64 `json:"TotalPDFCount,omitnil,omitempty" name:"TotalPDFCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// The scene, which defaults to `V1`.
	// Valid values:
	// V1
	// V2
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`
}

type RecognizeIndonesiaIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// The scene, which defaults to `V1`.
	// Valid values:
	// V1
	// V2
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`
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
	NIK *string `json:"NIK,omitnil,omitempty" name:"NIK"`

	// The full name.
	Nama *string `json:"Nama,omitnil,omitempty" name:"Nama"`

	// The place and date of birth.
	TempatTglLahir *string `json:"TempatTglLahir,omitnil,omitempty" name:"TempatTglLahir"`

	// The gender.
	JenisKelamin *string `json:"JenisKelamin,omitnil,omitempty" name:"JenisKelamin"`

	// The blood type.
	GolDarah *string `json:"GolDarah,omitnil,omitempty" name:"GolDarah"`

	// The address.
	Alamat *string `json:"Alamat,omitnil,omitempty" name:"Alamat"`

	// The street.
	RTRW *string `json:"RTRW,omitnil,omitempty" name:"RTRW"`

	// The village.
	KelDesa *string `json:"KelDesa,omitnil,omitempty" name:"KelDesa"`

	// The region.
	Kecamatan *string `json:"Kecamatan,omitnil,omitempty" name:"Kecamatan"`

	// The religion.
	Agama *string `json:"Agama,omitnil,omitempty" name:"Agama"`

	// The marital status.
	StatusPerkawinan *string `json:"StatusPerkawinan,omitnil,omitempty" name:"StatusPerkawinan"`

	// The occupation.
	Perkerjaan *string `json:"Perkerjaan,omitnil,omitempty" name:"Perkerjaan"`

	// The nationality.
	KewargaNegaraan *string `json:"KewargaNegaraan,omitnil,omitempty" name:"KewargaNegaraan"`

	// The expiry date.
	BerlakuHingga *string `json:"BerlakuHingga,omitnil,omitempty" name:"BerlakuHingga"`

	// The issue date.
	IssuedDate *string `json:"IssuedDate,omitnil,omitempty" name:"IssuedDate"`

	// The photo.
	Photo *string `json:"Photo,omitnil,omitempty" name:"Photo"`

	// The province, which is supported when the value of `Scene` is `V2`.
	Provinsi *string `json:"Provinsi,omitnil,omitempty" name:"Provinsi"`

	// The city, which is supported when the value of `Scene` is `V2`.
	Kota *string `json:"Kota,omitnil,omitempty" name:"Kota"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizeKoreanDrivingLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
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
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// The license number.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// The resident registration number.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// The license class type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// The name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The renewal period.
	AptitudeTesDate *string `json:"AptitudeTesDate,omitnil,omitempty" name:"AptitudeTesDate"`

	// The issue date.
	DateOfIssue *string `json:"DateOfIssue,omitnil,omitempty" name:"DateOfIssue"`

	// The Base64-encoded identity photo.
	Photo *string `json:"Photo,omitnil,omitempty" name:"Photo"`

	// The gender.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// The birth date in the format of dd/mm/yyyy.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizeKoreanIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
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
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// The address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// The name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The issue date.
	DateOfIssue *string `json:"DateOfIssue,omitnil,omitempty" name:"DateOfIssue"`

	// The Base64-encoded identity photo.
	Photo *string `json:"Photo,omitnil,omitempty" name:"Photo"`

	// The gender.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// The birth date in the format of dd/mm/yyyy.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type RecognizeMacaoIDCardOCRRequestParams struct {
	// The URL address of the image. 
	// Supported image formats: PNG, JPG, JPEG. Not support GIF yet.
	// Supported image size: The downloaded image should not exceed 7M. The image download takes no more than 3 seconds.Storing images in Tencent Cloud URLs can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG. The GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. Storing images in Tencent Cloud URLs can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Base64 value of the image.Supported image formats: PNG, JPG, JPEG. Not support GIF yet.
	// Supported image size: The downloaded image should not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds.
	// One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG. The GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The following optional fields are of string type and are empty by default: 
	// RetImage: whether to return the processed image (base64 encrypted string); the value and meaning of RetImage are as follows: 1.portrait Return portrait image data 2."" Do not return image data SDK setting method reference: Config = Json.stringify({"RetImage":"portrait"}) API 3.0 Explorer setting method reference: Config = {"RetImage":"portrait" }
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

type RecognizeMacaoIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The URL address of the image. 
	// Supported image formats: PNG, JPG, JPEG. Not support GIF yet.
	// Supported image size: The downloaded image should not exceed 7M. The image download takes no more than 3 seconds.Storing images in Tencent Cloud URLs can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG. The GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. Storing images in Tencent Cloud URLs can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Base64 value of the image.Supported image formats: PNG, JPG, JPEG. Not support GIF yet.
	// Supported image size: The downloaded image should not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds.
	// One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG. The GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The following optional fields are of string type and are empty by default: 
	// RetImage: whether to return the processed image (base64 encrypted string); the value and meaning of RetImage are as follows: 1.portrait Return portrait image data 2."" Do not return image data SDK setting method reference: Config = Json.stringify({"RetImage":"portrait"}) API 3.0 Explorer setting method reference: Config = {"RetImage":"portrait" }
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *RecognizeMacaoIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMacaoIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "BackImageUrl")
	delete(f, "ImageBase64")
	delete(f, "BackImageBase64")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeMacaoIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeMacaoIDCardOCRResponseParams struct {
	// Chinese last name
	CnLastName *string `json:"CnLastName,omitnil,omitempty" name:"CnLastName"`

	// English last name
	EnLastName *string `json:"EnLastName,omitnil,omitempty" name:"EnLastName"`

	// Last name code
	LastNameCode *string `json:"LastNameCode,omitnil,omitempty" name:"LastNameCode"`

	// Chinese first name
	CnFirstName *string `json:"CnFirstName,omitnil,omitempty" name:"CnFirstName"`

	// English first name
	EnFirstName *string `json:"EnFirstName,omitnil,omitempty" name:"EnFirstName"`

	// First name code
	FirstNameCode *string `json:"FirstNameCode,omitnil,omitempty" name:"FirstNameCode"`

	// ID Number
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Birthday(DD-MM-YYYY)
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// First issue Date (DD-MM-YYYY)
	FirstIssueDate *string `json:"FirstIssueDate,omitnil,omitempty" name:"FirstIssueDate"`

	// Issue date (DD-MM-YYYY)
	CurrentIssueDate *string `json:"CurrentIssueDate,omitnil,omitempty" name:"CurrentIssueDate"`

	// Validity period (DD-MM-YYYY)
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// ID symbol
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// Height (unit: meters)
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// Processed image (Base64)
	RetImage *string `json:"RetImage,omitnil,omitempty" name:"RetImage"`

	// Image rotation angle, the horizontal direction of the text is 0, clockwise is positive, counterclockwise is negative
	Angle *string `json:"Angle,omitnil,omitempty" name:"Angle"`

	// Resident type.
	ResidentType *string `json:"ResidentType,omitnil,omitempty" name:"ResidentType"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9104 Alarm for PS certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeMacaoIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeMacaoIDCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeMacaoIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMacaoIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeMainlandIDCardOCRRequestParams struct {
	// The Base64 value of the image. The image is required to be no larger than 7M after Base64 encoding, and the resolution is recommended to be 500*800 or above. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL address of the image. The image is required to be no larger than 7M after Base64 encoding, and the resolution is recommended to be 500*800 or above. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. It is recommended that images be stored in Tencent Cloud to ensure higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// FRONT: The side of the ID card with the photo (portrait side), BACK: The side of the ID card with the national emblem (national emblem side). If this parameter is not filled in, the front and back of the ID card will be automatically determined for you.
	CardSide *string `json:"CardSide,omitnil,omitempty" name:"CardSide"`

	// Whether to return the ID card portrait, the default is false
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`

	// Whether to enable ID card photo cropping (removing excess edges outside the ID, automatically correcting the shooting angle), the default value is false
	CropIdCard *bool `json:"CropIdCard,omitnil,omitempty" name:"CropIdCard"`
}

type RecognizeMainlandIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64 value of the image. The image is required to be no larger than 7M after Base64 encoding, and the resolution is recommended to be 500*800 or above. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL address of the image. The image is required to be no larger than 7M after Base64 encoding, and the resolution is recommended to be 500*800 or above. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. It is recommended that images be stored in Tencent Cloud to ensure higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// FRONT: The side of the ID card with the photo (portrait side), BACK: The side of the ID card with the national emblem (national emblem side). If this parameter is not filled in, the front and back of the ID card will be automatically determined for you.
	CardSide *string `json:"CardSide,omitnil,omitempty" name:"CardSide"`

	// Whether to return the ID card portrait, the default is false
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`

	// Whether to enable ID card photo cropping (removing excess edges outside the ID, automatically correcting the shooting angle), the default value is false
	CropIdCard *bool `json:"CropIdCard,omitnil,omitempty" name:"CropIdCard"`
}

func (r *RecognizeMainlandIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMainlandIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CardSide")
	delete(f, "CropPortrait")
	delete(f, "CropIdCard")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeMainlandIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeMainlandIDCardOCRResponseParams struct {
	// Name((portrait side))
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Sex((portrait side))
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Nation((portrait side))
	Nation *string `json:"Nation,omitnil,omitempty" name:"Nation"`

	// Brithday((portrait side))
	Birth *string `json:"Birth,omitnil,omitempty" name:"Birth"`

	// Address(portrait side)
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// ID number (portrait side)
	IdNum *string `json:"IdNum,omitnil,omitempty" name:"IdNum"`

	// Card authority(national emblem side)
	Authority *string `json:"Authority,omitnil,omitempty" name:"Authority"`

	// Card valid date (national emblem side)
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9104 Alarm for PS certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// Portrait image base64
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// ID card photo cropping results base64
	IdCardImage *string `json:"IdCardImage,omitnil,omitempty" name:"IdCardImage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeMainlandIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeMainlandIDCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeMainlandIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMainlandIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeMexicoVTIDRequestParams struct {
	// Base64 value of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image downloading time should not exceed 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. URLs of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizeMexicoVTIDRequest struct {
	*tchttp.BaseRequest
	
	// Base64 value of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image downloading time should not exceed 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of the image. Supported image formats: PNG, JPG, JPEG, GIF format is not supported yet. Supported image size: The downloaded image should not exceed 7M after Base64 encoding. Image download time should not exceed 3 seconds. URLs of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

func (r *RecognizeMexicoVTIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMexicoVTIDRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeMexicoVTIDRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeMexicoVTIDResponseParams struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Sex
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Address
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Vote PIN Code
	VotePIN *string `json:"VotePIN,omitnil,omitempty" name:"VotePIN"`

	// Unique Population Registry Code
	CURP *string `json:"CURP,omitnil,omitempty" name:"CURP"`

	// Birthday
	Birth *string `json:"Birth,omitnil,omitempty" name:"Birth"`

	// Section Number
	SECCION *string `json:"SECCION,omitnil,omitempty" name:"SECCION"`

	// IssueDate
	IssueDate *string `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`

	// ValidDate
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// State
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Locality
	Locality *string `json:"Locality,omitnil,omitempty" name:"Locality"`

	// Edition
	EMISION *string `json:"EMISION,omitnil,omitempty" name:"EMISION"`

	// Portrait image base64
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeMexicoVTIDResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeMexicoVTIDResponseParams `json:"Response"`
}

func (r *RecognizeMexicoVTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMexicoVTIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesDrivingLicenseOCRRequestParams struct {
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizePhilippinesDrivingLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
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
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil,omitempty" name:"HeadPortrait"`

	// The full name.
	Name *TextDetectionResult `json:"Name,omitnil,omitempty" name:"Name"`

	// The last name.
	LastName *TextDetectionResult `json:"LastName,omitnil,omitempty" name:"LastName"`

	// The first name.
	FirstName *TextDetectionResult `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// The middle name.
	MiddleName *TextDetectionResult `json:"MiddleName,omitnil,omitempty" name:"MiddleName"`

	// The nationality.
	Nationality *TextDetectionResult `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// The gender.
	Sex *TextDetectionResult `json:"Sex,omitnil,omitempty" name:"Sex"`

	// The address.
	Address *TextDetectionResult `json:"Address,omitnil,omitempty" name:"Address"`

	// The license No.
	LicenseNo *TextDetectionResult `json:"LicenseNo,omitnil,omitempty" name:"LicenseNo"`

	// The expiration date.
	ExpiresDate *TextDetectionResult `json:"ExpiresDate,omitnil,omitempty" name:"ExpiresDate"`

	// The agency code.
	AgencyCode *TextDetectionResult `json:"AgencyCode,omitnil,omitempty" name:"AgencyCode"`

	// The date of birth.
	Birthday *TextDetectionResult `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type RecognizePhilippinesSssIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
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
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil,omitempty" name:"HeadPortrait"`

	// The common reference number (CRN).
	LicenseNumber *TextDetectionResult `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// The full name.
	FullName *TextDetectionResult `json:"FullName,omitnil,omitempty" name:"FullName"`

	// The date of birth.
	Birthday *TextDetectionResult `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type RecognizePhilippinesTinIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
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
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil,omitempty" name:"HeadPortrait"`

	// The tax identification number (TIN).
	LicenseNumber *TextDetectionResult `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// The name.
	FullName *TextDetectionResult `json:"FullName,omitnil,omitempty" name:"FullName"`

	// The address.
	Address *TextDetectionResult `json:"Address,omitnil,omitempty" name:"Address"`

	// The birth date.
	Birthday *TextDetectionResult `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The issue date.
	IssueDate *TextDetectionResult `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizePhilippinesUMIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
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
	Surname *TextDetectionResult `json:"Surname,omitnil,omitempty" name:"Surname"`

	// The middle name.
	MiddleName *TextDetectionResult `json:"MiddleName,omitnil,omitempty" name:"MiddleName"`

	// The given name.
	GivenName *TextDetectionResult `json:"GivenName,omitnil,omitempty" name:"GivenName"`

	// The address.
	Address *TextDetectionResult `json:"Address,omitnil,omitempty" name:"Address"`

	// The date of birth.
	Birthday *TextDetectionResult `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The common reference number (CRN).
	CRN *TextDetectionResult `json:"CRN,omitnil,omitempty" name:"CRN"`

	// The gender.
	Sex *TextDetectionResult `json:"Sex,omitnil,omitempty" name:"Sex"`

	// The Base64-encoded identity photo.
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil,omitempty" name:"HeadPortrait"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type RecognizePhilippinesVoteIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// Whether to return the identity photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// The Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
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
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil,omitempty" name:"HeadPortrait"`

	// The voter's identification number (VIN).
	VIN *TextDetectionResult `json:"VIN,omitnil,omitempty" name:"VIN"`

	// The first name.
	FirstName *TextDetectionResult `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// The last name.
	LastName *TextDetectionResult `json:"LastName,omitnil,omitempty" name:"LastName"`

	// The date of birth.
	Birthday *TextDetectionResult `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The civil status.
	CivilStatus *TextDetectionResult `json:"CivilStatus,omitnil,omitempty" name:"CivilStatus"`

	// The citizenship.
	Citizenship *TextDetectionResult `json:"Citizenship,omitnil,omitempty" name:"Citizenship"`

	// The address.
	Address *TextDetectionResult `json:"Address,omitnil,omitempty" name:"Address"`

	// The precinct.
	PrecinctNo *TextDetectionResult `json:"PrecinctNo,omitnil,omitempty" name:"PrecinctNo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type RecognizeSingaporeIDCardOCRRequestParams struct {
	// The Base64 value of the image. The image is required to be no larger than 7M after Base64 encoding, and the resolution is recommended to be 500*800 or above. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL address of the image. The image is required to be no larger than 7M after Base64 encoding, and the resolution is recommended to be 500*800 or above. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. It is recommended that images be stored in Tencent Cloud to ensure higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizeSingaporeIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64 value of the image. The image is required to be no larger than 7M after Base64 encoding, and the resolution is recommended to be 500*800 or above. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL address of the image. The image is required to be no larger than 7M after Base64 encoding, and the resolution is recommended to be 500*800 or above. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupies at least 2/3 of the picture. It is recommended that images be stored in Tencent Cloud to ensure higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return portrait photos.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

func (r *RecognizeSingaporeIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeSingaporeIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeSingaporeIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeSingaporeIDCardOCRResponseParams struct {
	// Chinese name
	ChName *string `json:"ChName,omitnil,omitempty" name:"ChName"`

	// English name
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Birth Country
	CountryOfBirth *string `json:"CountryOfBirth,omitnil,omitempty" name:"CountryOfBirth"`

	// Brithday
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Address(back side)
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// License number
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Nationality(back side)
	Race *string `json:"Race,omitnil,omitempty" name:"Race"`

	// NRIC code(back side)
	NRICCode *string `json:"NRICCode,omitnil,omitempty" name:"NRICCode"`

	// Post code(back side)
	PostCode *string `json:"PostCode,omitnil,omitempty" name:"PostCode"`

	// Date of Expiration(back side)
	DateOfExpiration *string `json:"DateOfExpiration,omitnil,omitempty" name:"DateOfExpiration"`

	// Date of issue(back side)
	DateOfIssue *string `json:"DateOfIssue,omitnil,omitempty" name:"DateOfIssue"`

	// Head image 
	Photo *string `json:"Photo,omitnil,omitempty" name:"Photo"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9104 Alarm for PS certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeSingaporeIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeSingaporeIDCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeSingaporeIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeSingaporeIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeTableAccurateOCRRequestParams struct {
	// The Base64-encoded value of an image.
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600 x 800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// Supported image pixels: 20 to 10,000
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image or PDF file.
	// The image or PDF file cannot exceed 7 MB after being Base64-encoded. A resolution above 600 x 800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// Supported image pixels: 20 to 10,000
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `IsPdf` is `true`. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`
}

type RecognizeTableAccurateOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image.
	// The image cannot exceed 7 MB after being Base64-encoded. A resolution above 600 x 800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// Supported image pixels: 20 to 10,000
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image or PDF file.
	// The image or PDF file cannot exceed 7 MB after being Base64-encoded. A resolution above 600 x 800 is recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported.
	// Supported image pixels: 20 to 10,000
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `IsPdf` is `true`. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`
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
	TableDetections []*TableInfo `json:"TableDetections,omitnil,omitempty" name:"TableDetections"`

	// Base64-encoded Excel data.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// The total number of pages in the PDF file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PdfPageSize *int64 `json:"PdfPageSize,omitnil,omitempty" name:"PdfPageSize"`

	// Image rotation angle in degrees. 0°: The horizontal direction of the text on the image; a negative value: rotate counterclockwise. Value range: -360° to 0°.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG. GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG. GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. Storing images in Tencent Cloud URLs can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to crop the profile photo. The default value is `false`, meaning not to return the Base64-encoded value of the profile photo on the Thai identity card.
	// When this parameter is set to `true`, the Base64-encoded value of the profile photo on the Thai identity card after rotation correction is returned.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
}

type RecognizeThaiIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Base64 value of the image on the back of the card. Supported image formats: PNG, JPG, JPEG. GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. One of ImageUrl and ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The URL address of the image on the back of the card. Supported image formats: PNG, JPG, JPEG. GIF format is not supported yet. Supported image size: The downloaded image does not exceed 7M after Base64 encoding. The image download takes no more than 3 seconds. Storing images in Tencent Cloud URLs can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Whether to crop the profile photo. The default value is `false`, meaning not to return the Base64-encoded value of the profile photo on the Thai identity card.
	// When this parameter is set to `true`, the Base64-encoded value of the profile photo on the Thai identity card after rotation correction is returned.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
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
	delete(f, "BackImageBase64")
	delete(f, "ImageUrl")
	delete(f, "BackImageUrl")
	delete(f, "CropPortrait")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeThaiIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeThaiIDCardOCRResponseParams struct {
	// ID card number
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Name in Thai
	ThaiName *string `json:"ThaiName,omitnil,omitempty" name:"ThaiName"`

	// Name in English
	EnFirstName *string `json:"EnFirstName,omitnil,omitempty" name:"EnFirstName"`

	// Name in English
	EnLastName *string `json:"EnLastName,omitnil,omitempty" name:"EnLastName"`

	// Date of issue in Thai
	IssueDate *string `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`

	// Expiration date in Thai
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

	// Date of issue in English
	EnIssueDate *string `json:"EnIssueDate,omitnil,omitempty" name:"EnIssueDate"`

	// Expiration date in English
	EnExpirationDate *string `json:"EnExpirationDate,omitnil,omitempty" name:"EnExpirationDate"`

	// Date of birth in Thai
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Date of birth in English
	EnBirthday *string `json:"EnBirthday,omitnil,omitempty" name:"EnBirthday"`

	// Religion
	Religion *string `json:"Religion,omitnil,omitempty" name:"Religion"`

	// Serial number
	SerialNumber *string `json:"SerialNumber,omitnil,omitempty" name:"SerialNumber"`

	// Address
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// LaserID in the back of the card.
	LaserID *string `json:"LaserID,omitnil,omitempty" name:"LaserID"`

	// Identity photo
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// Field confidence:
	//  { "ID": { "Confidence": 0.9999 }, "ThaiName": { "Confidence": 0.9996 } }
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type RecognizeThaiPinkCardRequestParams struct {
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to crop the profile photo. The default value is `false`, meaning not to return the Base64-encoded value of the profile photo on the Thai identity card.
	// When this parameter is set to `true`, the Base64-encoded value of the profile photo on the Thai identity card after rotation correction is returned.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
}

type RecognizeThaiPinkCardRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to crop the profile photo. The default value is `false`, meaning not to return the Base64-encoded value of the profile photo on the Thai identity card.
	// When this parameter is set to `true`, the Base64-encoded value of the profile photo on the Thai identity card after rotation correction is returned.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
}

func (r *RecognizeThaiPinkCardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeThaiPinkCardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CropPortrait")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeThaiPinkCardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeThaiPinkCardResponseParams struct {
	// Country
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// ID number
	IDNumber *string `json:"IDNumber,omitnil,omitempty" name:"IDNumber"`

	// Name in Thai
	ThaiName *string `json:"ThaiName,omitnil,omitempty" name:"ThaiName"`

	// Name in English
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// Surname in Thai
	ThaiSurName *string `json:"ThaiSurName,omitnil,omitempty" name:"ThaiSurName"`

	// Date of birth in Thai
	ThaiDOB *string `json:"ThaiDOB,omitnil,omitempty" name:"ThaiDOB"`

	// Date of birth in English
	EnDOB *string `json:"EnDOB,omitnil,omitempty" name:"EnDOB"`

	// Photo number
	PhotoNumber *string `json:"PhotoNumber,omitnil,omitempty" name:"PhotoNumber"`

	// Address in Thai
	ThaiAddress *string `json:"ThaiAddress,omitnil,omitempty" name:"ThaiAddress"`

	// Date of issue in Thai
	ThaiDateOfIssue *string `json:"ThaiDateOfIssue,omitnil,omitempty" name:"ThaiDateOfIssue"`

	// Date of issue in English
	DateOfIssue *string `json:"DateOfIssue,omitnil,omitempty" name:"DateOfIssue"`

	// Expiration date in Thai
	ThaiDateOfExpiry *string `json:"ThaiDateOfExpiry,omitnil,omitempty" name:"ThaiDateOfExpiry"`

	// Expiration date in English
	DateOfExpiry *string `json:"DateOfExpiry,omitnil,omitempty" name:"DateOfExpiry"`

	// Issuing agency
	IssuingAgency *string `json:"IssuingAgency,omitnil,omitempty" name:"IssuingAgency"`

	// Ref number
	RefNumber *string `json:"RefNumber,omitnil,omitempty" name:"RefNumber"`

	// Field confidence:
	//  { "ID": { "Confidence": 0.9999 }, "ThaiName": { "Confidence": 0.9996 } }
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// Identity photo
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate,
	// -9102 Alarm for photocopied certificate,
	// -9103 Alarm for photographed certificate,
	// -9107 Alarm for reflective certificate,
	// -9108 Alarm for blurry image,
	// -9109 This capability is not enabled.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeThaiPinkCardResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeThaiPinkCardResponseParams `json:"Response"`
}

func (r *RecognizeThaiPinkCardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeThaiPinkCardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rect struct {
	// X-coordinate of top-left point
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Y-coordinate of top-left point
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// Width
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Height
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type SealInfo struct {
	// Seal body information
	SealBody *string `json:"SealBody,omitnil,omitempty" name:"SealBody"`

	// Seal coordinates
	Location *Rect `json:"Location,omitnil,omitempty" name:"Location"`

	// Other text content
	OtherTexts []*string `json:"OtherTexts,omitnil,omitempty" name:"OtherTexts"`

	// Seal shape. Valid values:
	// 0: Round
	// 1: Oval
	// 2: Rectangle
	// 3: Diamond
	// 4: Triangle
	SealShape *string `json:"SealShape,omitnil,omitempty" name:"SealShape"`
}

// Predefined struct for user
type SealOCRRequestParams struct {
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type SealOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of an image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. The image cannot exceed 7 MB after being Base64-encoded. A resolution above 500 x 800 is recommended. PNG, JPG, JPEG, and BMP formats are supported. It is recommended that the card part occupy more than 2/3 area of the image. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
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
	SealBody *string `json:"SealBody,omitnil,omitempty" name:"SealBody"`

	// Seal coordinates
	Location *Rect `json:"Location,omitnil,omitempty" name:"Location"`

	// Other text content
	OtherTexts []*string `json:"OtherTexts,omitnil,omitempty" name:"OtherTexts"`

	// All seal information
	SealInfos []*SealInfo `json:"SealInfos,omitnil,omitempty" name:"SealInfos"`

	// Seal shape. Valid values:
	// 0: Round
	// 1: Oval
	// 2: Rectangle
	// 3: Diamond
	// 4: Triangle
	SealShape *string `json:"SealShape,omitnil,omitempty" name:"SealShape"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Date
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Time
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Departure station
	StationGetOn *string `json:"StationGetOn,omitnil,omitempty" name:"StationGetOn"`

	// Destination
	StationGetOff *string `json:"StationGetOff,omitnil,omitempty" name:"StationGetOff"`

	// Fare
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Currency
	CurrencyCode *string `json:"CurrencyCode,omitnil,omitempty" name:"CurrencyCode"`
}

type SingleInvoiceItem struct {
	// Special VAT invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatSpecialInvoice *VatInvoiceInfo `json:"VatSpecialInvoice,omitnil,omitempty" name:"VatSpecialInvoice"`

	// General VAT invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatCommonInvoice *VatInvoiceInfo `json:"VatCommonInvoice,omitnil,omitempty" name:"VatCommonInvoice"`

	// Electronic general VAT invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicCommonInvoice *VatInvoiceInfo `json:"VatElectronicCommonInvoice,omitnil,omitempty" name:"VatElectronicCommonInvoice"`

	// Electronic special VAT invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicSpecialInvoice *VatInvoiceInfo `json:"VatElectronicSpecialInvoice,omitnil,omitempty" name:"VatElectronicSpecialInvoice"`

	// Blockchain electronic invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicInvoiceBlockchain *VatInvoiceInfo `json:"VatElectronicInvoiceBlockchain,omitnil,omitempty" name:"VatElectronicInvoiceBlockchain"`

	// Electronic general VAT invoice (toll)
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicInvoiceToll *VatInvoiceInfo `json:"VatElectronicInvoiceToll,omitnil,omitempty" name:"VatElectronicInvoiceToll"`

	// Electronic invoice (special)
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicSpecialInvoiceFull *VatElectronicInfo `json:"VatElectronicSpecialInvoiceFull,omitnil,omitempty" name:"VatElectronicSpecialInvoiceFull"`

	// Electronic invoice (general)
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatElectronicInvoiceFull *VatElectronicInfo `json:"VatElectronicInvoiceFull,omitnil,omitempty" name:"VatElectronicInvoiceFull"`

	// General machine-printed invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	MachinePrintedInvoice *MachinePrintedInvoice `json:"MachinePrintedInvoice,omitnil,omitempty" name:"MachinePrintedInvoice"`

	// Bus ticket
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusInvoice *BusInvoice `json:"BusInvoice,omitnil,omitempty" name:"BusInvoice"`

	// Ship ticket
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShippingInvoice *ShippingInvoice `json:"ShippingInvoice,omitnil,omitempty" name:"ShippingInvoice"`

	// Toll receipt
	// Note: This field may return null, indicating that no valid values can be obtained.
	TollInvoice *TollInvoice `json:"TollInvoice,omitnil,omitempty" name:"TollInvoice"`

	// Other invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	OtherInvoice *OtherInvoice `json:"OtherInvoice,omitnil,omitempty" name:"OtherInvoice"`

	// Motor vehicle sales invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	MotorVehicleSaleInvoice *MotorVehicleSaleInvoice `json:"MotorVehicleSaleInvoice,omitnil,omitempty" name:"MotorVehicleSaleInvoice"`

	// Used car invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsedCarPurchaseInvoice *UsedCarPurchaseInvoice `json:"UsedCarPurchaseInvoice,omitnil,omitempty" name:"UsedCarPurchaseInvoice"`

	// General VAT invoice (roll)
	// Note: This field may return null, indicating that no valid values can be obtained.
	VatInvoiceRoll *VatInvoiceRoll `json:"VatInvoiceRoll,omitnil,omitempty" name:"VatInvoiceRoll"`

	// Taxi receipt
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaxiTicket *TaxiTicket `json:"TaxiTicket,omitnil,omitempty" name:"TaxiTicket"`

	// Quota invoice
	// Note: This field may return null, indicating that no valid values can be obtained.
	QuotaInvoice *QuotaInvoice `json:"QuotaInvoice,omitnil,omitempty" name:"QuotaInvoice"`

	// Itinerary/Receipt of e-ticket for air transportation
	// Note: This field may return null, indicating that no valid values can be obtained.
	AirTransport *AirTransport `json:"AirTransport,omitnil,omitempty" name:"AirTransport"`

	// Non-tax revenue general receipt
	// Note: This field may return null, indicating that no valid values can be obtained.
	NonTaxIncomeGeneralBill *NonTaxIncomeBill `json:"NonTaxIncomeGeneralBill,omitnil,omitempty" name:"NonTaxIncomeGeneralBill"`

	// Non-tax revenue unified payment voucher
	// Note: This field may return null, indicating that no valid values can be obtained.
	NonTaxIncomeElectronicBill *NonTaxIncomeBill `json:"NonTaxIncomeElectronicBill,omitnil,omitempty" name:"NonTaxIncomeElectronicBill"`

	// Train ticket
	// Note: This field may return null, indicating that no valid values can be obtained.
	TrainTicket *TrainTicket `json:"TrainTicket,omitnil,omitempty" name:"TrainTicket"`


	MedicalOutpatientInvoice *MedicalInvoice `json:"MedicalOutpatientInvoice,omitnil,omitempty" name:"MedicalOutpatientInvoice"`


	MedicalHospitalizedInvoice *MedicalInvoice `json:"MedicalHospitalizedInvoice,omitnil,omitempty" name:"MedicalHospitalizedInvoice"`
}

// Predefined struct for user
type SmartStructuralOCRV2RequestParams struct {
	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitnil,omitempty" name:"IsPdf"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `IsPdf` is `true`. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// The names of the fields you want to return for the structured information recognition.
	// For example, if you want to return only the recognition result of the "Name" and "Gender" fields, set this parameter as follows:
	// ItemNames=["Name","Gender"]
	ItemNames []*string `json:"ItemNames,omitnil,omitempty" name:"ItemNames"`

	// Whether to enable recognition of all fields.
	ReturnFullText *bool `json:"ReturnFullText,omitnil,omitempty" name:"ReturnFullText"`
}

type SmartStructuralOCRV2Request struct {
	*tchttp.BaseRequest
	
	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Whether to enable PDF recognition. Default value: `false`. If you enable this feature, both images and PDF files can be recognized.
	IsPdf *bool `json:"IsPdf,omitnil,omitempty" name:"IsPdf"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF and the value of `IsPdf` is `true`. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// The names of the fields you want to return for the structured information recognition.
	// For example, if you want to return only the recognition result of the "Name" and "Gender" fields, set this parameter as follows:
	// ItemNames=["Name","Gender"]
	ItemNames []*string `json:"ItemNames,omitnil,omitempty" name:"ItemNames"`

	// Whether to enable recognition of all fields.
	ReturnFullText *bool `json:"ReturnFullText,omitnil,omitempty" name:"ReturnFullText"`
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
	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// The structural information (key-value).
	StructuralList []*GroupInfo `json:"StructuralList,omitnil,omitempty" name:"StructuralList"`

	// The recognized text information.
	WordList []*WordItem `json:"WordList,omitnil,omitempty" name:"WordList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type SmartStructuralProRequestParams struct {
	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF `. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// The names of the fields you want to return for the structured information recognition.
	// For example, if you want to return only the recognition result of the "Name" and "Gender" fields, set this parameter as follows:
	// ItemNames=["Name","Gender"]
	ItemNames []*string `json:"ItemNames,omitnil,omitempty" name:"ItemNames"`

	// Whether to enable recognition of all fields.
	ReturnFullText *bool `json:"ReturnFullText,omitnil,omitempty" name:"ReturnFullText"`

	// Configuration ID support: General 
	// -- General scenarios; InvoiceEng 
	// -- Ocean bill of lading, international invoice template; 
	// -- Ocean shipment order template; WayBillEng 
	// -- CustomsDeclaration
	// -- WeightNote
	// -- MedicalMeter
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Enable recognition of coordinate values in full-text fields
	EnableCoord *bool `json:"EnableCoord,omitnil,omitempty" name:"EnableCoord"`
}

type SmartStructuralProRequest struct {
	*tchttp.BaseRequest
	
	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The number of the PDF page that needs to be recognized. Only one single PDF page can be recognized. This parameter is valid if the uploaded file is a PDF `. Default value: `1`.
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// The names of the fields you want to return for the structured information recognition.
	// For example, if you want to return only the recognition result of the "Name" and "Gender" fields, set this parameter as follows:
	// ItemNames=["Name","Gender"]
	ItemNames []*string `json:"ItemNames,omitnil,omitempty" name:"ItemNames"`

	// Whether to enable recognition of all fields.
	ReturnFullText *bool `json:"ReturnFullText,omitnil,omitempty" name:"ReturnFullText"`

	// Configuration ID support: General 
	// -- General scenarios; InvoiceEng 
	// -- Ocean bill of lading, international invoice template; 
	// -- Ocean shipment order template; WayBillEng 
	// -- CustomsDeclaration
	// -- WeightNote
	// -- MedicalMeter
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Enable recognition of coordinate values in full-text fields
	EnableCoord *bool `json:"EnableCoord,omitnil,omitempty" name:"EnableCoord"`
}

func (r *SmartStructuralProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartStructuralProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "PdfPageNumber")
	delete(f, "ItemNames")
	delete(f, "ReturnFullText")
	delete(f, "ConfigId")
	delete(f, "EnableCoord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SmartStructuralProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SmartStructuralProResponseParams struct {
	// The rotation angle (degrees) of the text on the image. 0: The text is horizontal. Positive value: The text is rotated clockwise. Negative value: The text is rotated counterclockwise.
	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// The structural information (key-value).
	StructuralList []*GroupInfo `json:"StructuralList,omitnil,omitempty" name:"StructuralList"`

	// The recognized text information.
	WordList []*WordItem `json:"WordList,omitnil,omitempty" name:"WordList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SmartStructuralProResponse struct {
	*tchttp.BaseResponse
	Response *SmartStructuralProResponseParams `json:"Response"`
}

func (r *SmartStructuralProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartStructuralProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableCellInfo struct {
	// Column index of the upper-left corner of the cell
	ColTl *int64 `json:"ColTl,omitnil,omitempty" name:"ColTl"`

	// Row index of the upper-left corner of the cell
	RowTl *int64 `json:"RowTl,omitnil,omitempty" name:"RowTl"`

	// Column index of the lower-right corner of the cell
	ColBr *int64 `json:"ColBr,omitnil,omitempty" name:"ColBr"`

	// Row index of the lower-right corner of the cell
	RowBr *int64 `json:"RowBr,omitnil,omitempty" name:"RowBr"`

	// Recognized string text within the cell. If there are multiple lines, they are separated by "\n".
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Cell type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Cell confidence
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Four-point coordinates of the cell in the image
	Polygon []*Coord `json:"Polygon,omitnil,omitempty" name:"Polygon"`
}

type TableInfo struct {
	// Cell content
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Cells []*TableCellInfo `json:"Cells,omitnil,omitempty" name:"Cells"`

	// Type of text in the image. Valid values:
	// 0: Non-table text
	// 1: Text in a bordered table
	// 2: Text in a borderless table
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// The coordinates of the four vertices (upper-left, upper-right, lower-right, and lower-left) of the table body.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableCoordPoint []*Coord `json:"TableCoordPoint,omitnil,omitempty" name:"TableCoordPoint"`
}

// Predefined struct for user
type TableOCRRequestParams struct {
	// Base64-encoded value of image.
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported at present.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type TableOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded value of image.
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported at present.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. (This field is not supported outside Chinese mainland)
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 3 MB after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
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
	TextDetections []*TextTable `json:"TextDetections,omitnil,omitempty" name:"TextDetections"`

	// Base64-encoded Excel data.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Start time
	TimeGetOn *string `json:"TimeGetOn,omitnil,omitempty" name:"TimeGetOn"`

	// End time
	TimeGetOff *string `json:"TimeGetOff,omitnil,omitempty" name:"TimeGetOff"`

	// Unit price
	Price *string `json:"Price,omitnil,omitempty" name:"Price"`

	// Distance
	Mileage *string `json:"Mileage,omitnil,omitempty" name:"Mileage"`

	// Total amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Invoice place
	Place *string `json:"Place,omitnil,omitempty" name:"Place"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// License plate number
	LicensePlate *string `json:"LicensePlate,omitnil,omitempty" name:"LicensePlate"`

	// Fuel surcharge
	FuelFee *string `json:"FuelFee,omitnil,omitempty" name:"FuelFee"`

	// Booking fee
	BookingCallFee *string `json:"BookingCallFee,omitnil,omitempty" name:"BookingCallFee"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil,omitempty" name:"CompanySealMark"`
}

type TextDetection struct {
	// Recognized text line content.
	DetectedText *string `json:"DetectedText,omitnil,omitempty" name:"DetectedText"`

	// Confidence. Value range: 0–100.
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Text line coordinates, which are represented as 4 vertex coordinates.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Polygon []*Coord `json:"Polygon,omitnil,omitempty" name:"Polygon"`

	// Extended field.
	// The paragraph information `Parag` returned by the `GeneralBasicOcr` API contains `ParagNo`.
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// Pixel coordinates of the text line in the image after rotation correction, which is in the format of `(X-coordinate of top-left point, Y-coordinate of top-left point, width, height)`.
	ItemPolygon *ItemCoord `json:"ItemPolygon,omitnil,omitempty" name:"ItemPolygon"`

	// Information about a character, including the character itself and its confidence. Supported APIs: `GeneralBasicOCR`, `GeneralAccurateOCR`
	Words []*DetectedWords `json:"Words,omitnil,omitempty" name:"Words"`

	// Coordinates of a word’s four corners on the input image. Supported APIs: `GeneralBasicOCR`, `GeneralAccurateOCR`
	WordCoordPoint []*DetectedWordCoordPoint `json:"WordCoordPoint,omitnil,omitempty" name:"WordCoordPoint"`
}

type TextDetectionResult struct {
	// The recognized text line content.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// The coordinates, represented in the coordinates of the four points.
	Polygon []*Coord `json:"Polygon,omitnil,omitempty" name:"Polygon"`
}

type TextTable struct {
	// Column index of the top-left corner of the cell.
	ColTl *int64 `json:"ColTl,omitnil,omitempty" name:"ColTl"`

	// Row index of the top-left corner of the cell.
	RowTl *int64 `json:"RowTl,omitnil,omitempty" name:"RowTl"`

	// Column index of the bottom-right corner of the cell.
	ColBr *int64 `json:"ColBr,omitnil,omitempty" name:"ColBr"`

	// Row index of the bottom-right corner of the cell.
	RowBr *int64 `json:"RowBr,omitnil,omitempty" name:"RowBr"`

	// Cell text
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Cell type. Valid values: body, header, footer
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Confidence. Value range: 0–100
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Text line coordinates, which are represented as 4 vertex coordinates.
	Polygon []*Coord `json:"Polygon,omitnil,omitempty" name:"Polygon"`

	// Extended field
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`
}

type TollInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Date
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Time
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Entrance
	Entrance *string `json:"Entrance,omitnil,omitempty" name:"Entrance"`

	// Exit
	Exit *string `json:"Exit,omitnil,omitempty" name:"Exit"`

	// Highway mark (0: No, 1: Yes)
	HighwayMark *int64 `json:"HighwayMark,omitnil,omitempty" name:"HighwayMark"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`
}

type TrainTicket struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Departure date
	DateGetOn *string `json:"DateGetOn,omitnil,omitempty" name:"DateGetOn"`

	// Departure time
	TimeGetOn *string `json:"TimeGetOn,omitnil,omitempty" name:"TimeGetOn"`

	// Passenger's name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Departure station
	StationGetOn *string `json:"StationGetOn,omitnil,omitempty" name:"StationGetOn"`

	// Destination
	StationGetOff *string `json:"StationGetOff,omitnil,omitempty" name:"StationGetOff"`

	// Seat class
	Seat *string `json:"Seat,omitnil,omitempty" name:"Seat"`

	// Total amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Serial number
	SerialNumber *string `json:"SerialNumber,omitnil,omitempty" name:"SerialNumber"`

	// ID card number
	UserID *string `json:"UserID,omitnil,omitempty" name:"UserID"`

	// Check-in gate
	GateNumber *string `json:"GateNumber,omitnil,omitempty" name:"GateNumber"`

	// Fleet number
	TrainNumber *string `json:"TrainNumber,omitnil,omitempty" name:"TrainNumber"`

	// Handling fee
	HandlingFee *string `json:"HandlingFee,omitnil,omitempty" name:"HandlingFee"`

	// Original ticket price
	OriginalFare *string `json:"OriginalFare,omitnil,omitempty" name:"OriginalFare"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`

	// Seat No.
	SeatNumber *string `json:"SeatNumber,omitnil,omitempty" name:"SeatNumber"`

	// Ticket pickup address
	PickUpAddress *string `json:"PickUpAddress,omitnil,omitempty" name:"PickUpAddress"`

	// Ticket change information
	TicketChange *string `json:"TicketChange,omitnil,omitempty" name:"TicketChange"`

	// Additional fare
	AdditionalFare *string `json:"AdditionalFare,omitnil,omitempty" name:"AdditionalFare"`

	// Receipt No.
	ReceiptNumber *string `json:"ReceiptNumber,omitnil,omitempty" name:"ReceiptNumber"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Whether it is for reimbursement only (0: No, 1: Yes)
	ReimburseOnlyMark *int64 `json:"ReimburseOnlyMark,omitnil,omitempty" name:"ReimburseOnlyMark"`
}

type UsedCarPurchaseInvoice struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`

	// Seller's name
	Seller *string `json:"Seller,omitnil,omitempty" name:"Seller"`

	// Seller's phone number
	SellerTel *string `json:"SellerTel,omitnil,omitempty" name:"SellerTel"`

	// Seller's company code/personal ID card number
	SellerTaxID *string `json:"SellerTaxID,omitnil,omitempty" name:"SellerTaxID"`

	// Seller's address
	SellerAddress *string `json:"SellerAddress,omitnil,omitempty" name:"SellerAddress"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil,omitempty" name:"Buyer"`

	// Buyer's company code/personal ID card number
	BuyerID *string `json:"BuyerID,omitnil,omitempty" name:"BuyerID"`

	// Buyer's address
	BuyerAddress *string `json:"BuyerAddress,omitnil,omitempty" name:"BuyerAddress"`

	// Buyer's phone number
	BuyerTel *string `json:"BuyerTel,omitnil,omitempty" name:"BuyerTel"`

	// Company (used car market) name
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// Company's taxpayer identification number
	CompanyTaxID *string `json:"CompanyTaxID,omitnil,omitempty" name:"CompanyTaxID"`

	// Company's account opening bank and account number
	CompanyBankAccount *string `json:"CompanyBankAccount,omitnil,omitempty" name:"CompanyBankAccount"`

	// Company's phone number
	CompanyTel *string `json:"CompanyTel,omitnil,omitempty" name:"CompanyTel"`

	// Company's address
	CompanyAddress *string `json:"CompanyAddress,omitnil,omitempty" name:"CompanyAddress"`

	// Name of the transfer-to department of motor vehicles
	TransferAdministrationName *string `json:"TransferAdministrationName,omitnil,omitempty" name:"TransferAdministrationName"`

	// License plate number
	LicensePlate *string `json:"LicensePlate,omitnil,omitempty" name:"LicensePlate"`

	// Registration certificate No.
	RegistrationNumber *string `json:"RegistrationNumber,omitnil,omitempty" name:"RegistrationNumber"`

	// VIN
	VIN *string `json:"VIN,omitnil,omitempty" name:"VIN"`

	// Vehicle model
	VehicleModel *string `json:"VehicleModel,omitnil,omitempty" name:"VehicleModel"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil,omitempty" name:"VehicleType"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Form type
	FormType *string `json:"FormType,omitnil,omitempty" name:"FormType"`

	// Form name
	FormName *string `json:"FormName,omitnil,omitempty" name:"FormName"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil,omitempty" name:"CompanySealMark"`
}

type Value struct {
	// The value of the recognized field.
	AutoContent *string `json:"AutoContent,omitnil,omitempty" name:"AutoContent"`

	// The coordinates of the four vertices.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Coord *Polygon `json:"Coord,omitnil,omitempty" name:"Coord"`
}

type VatElectronicInfo struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Amount before tax
	PretaxAmount *string `json:"PretaxAmount,omitnil,omitempty" name:"PretaxAmount"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`

	// Seller's name
	Seller *string `json:"Seller,omitnil,omitempty" name:"Seller"`

	// Seller's taxpayer identification number
	SellerTaxID *string `json:"SellerTaxID,omitnil,omitempty" name:"SellerTaxID"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil,omitempty" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil,omitempty" name:"BuyerTaxID"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil,omitempty" name:"Issuer"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Subtotal amount
	SubTotal *string `json:"SubTotal,omitnil,omitempty" name:"SubTotal"`

	// Subtotal tax
	SubTax *string `json:"SubTax,omitnil,omitempty" name:"SubTax"`

	// Detailed items of an electronic invoice
	VatElectronicItems []*VatElectronicItemInfo `json:"VatElectronicItems,omitnil,omitempty" name:"VatElectronicItems"`
}

type VatElectronicItemInfo struct {
	// Item name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// Specification
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// Unit price
	Price *string `json:"Price,omitnil,omitempty" name:"Price"`

	// Amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`

	// Tax amount
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Unit
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil,omitempty" name:"VehicleType"`

	// Vehicle No.
	VehicleBrand *string `json:"VehicleBrand,omitnil,omitempty" name:"VehicleBrand"`

	// Departure place
	DeparturePlace *string `json:"DeparturePlace,omitnil,omitempty" name:"DeparturePlace"`

	// Destination
	ArrivalPlace *string `json:"ArrivalPlace,omitnil,omitempty" name:"ArrivalPlace"`

	// Name of the transported goods. It is returned only for a goods transport service invoice.
	TransportItemsName *string `json:"TransportItemsName,omitnil,omitempty" name:"TransportItemsName"`

	// Location of the construction service. It is returned only for a construction invoice.
	PlaceOfBuildingService *string `json:"PlaceOfBuildingService,omitnil,omitempty" name:"PlaceOfBuildingService"`

	// Name of the construction project. It is returned only for a construction invoice.
	BuildingName *string `json:"BuildingName,omitnil,omitempty" name:"BuildingName"`

	// Property or real estate ownership certificate No. It is returned only for a real estate operation and leasing service invoice.
	EstateNumber *string `json:"EstateNumber,omitnil,omitempty" name:"EstateNumber"`

	// Unit of area. It is returned only for a real estate operation and leasing service invoice.
	AreaUnit *string `json:"AreaUnit,omitnil,omitempty" name:"AreaUnit"`
}

type VatInvoiceInfo struct {
	// Check code
	CheckCode *string `json:"CheckCode,omitnil,omitempty" name:"CheckCode"`

	// Form type
	FormType *string `json:"FormType,omitnil,omitempty" name:"FormType"`

	// Vehicle and vessel tax
	TravelTax *string `json:"TravelTax,omitnil,omitempty" name:"TravelTax"`

	// Buyer's address and phone number
	BuyerAddrTel *string `json:"BuyerAddrTel,omitnil,omitempty" name:"BuyerAddrTel"`

	// Buyer's bank account number
	BuyerBankAccount *string `json:"BuyerBankAccount,omitnil,omitempty" name:"BuyerBankAccount"`

	// Company seal content
	CompanySealContent *string `json:"CompanySealContent,omitnil,omitempty" name:"CompanySealContent"`

	// Tax authority seal content
	TaxSealContent *string `json:"TaxSealContent,omitnil,omitempty" name:"TaxSealContent"`

	// Service type
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Whether there is an agent (0: No, 1: Yes)
	AgentMark *int64 `json:"AgentMark,omitnil,omitempty" name:"AgentMark"`

	// Whether there is a toll (0: No, 1: Yes)
	TransitMark *int64 `json:"TransitMark,omitnil,omitempty" name:"TransitMark"`

	// Whether there is refined oil (0: No, 1: Yes)
	OilMark *int64 `json:"OilMark,omitnil,omitempty" name:"OilMark"`

	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Machine-printed invoice number
	NumberConfirm *string `json:"NumberConfirm,omitnil,omitempty" name:"NumberConfirm"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`

	// Amount before tax
	PretaxAmount *string `json:"PretaxAmount,omitnil,omitempty" name:"PretaxAmount"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Machine No.
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// Ciphertext
	Ciphertext *string `json:"Ciphertext,omitnil,omitempty" name:"Ciphertext"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Seller's name
	Seller *string `json:"Seller,omitnil,omitempty" name:"Seller"`

	// Seller's taxpayer identification number
	SellerTaxID *string `json:"SellerTaxID,omitnil,omitempty" name:"SellerTaxID"`

	// Seller's address and phone number
	SellerAddrTel *string `json:"SellerAddrTel,omitnil,omitempty" name:"SellerAddrTel"`

	// Seller's bank account number
	SellerBankAccount *string `json:"SellerBankAccount,omitnil,omitempty" name:"SellerBankAccount"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil,omitempty" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil,omitempty" name:"BuyerTaxID"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil,omitempty" name:"CompanySealMark"`

	// Issuer
	Issuer *string `json:"Issuer,omitnil,omitempty" name:"Issuer"`

	// Reviewer
	Reviewer *string `json:"Reviewer,omitnil,omitempty" name:"Reviewer"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// Information about VAT invoice items
	VatInvoiceItemInfos []*VatInvoiceItemInfo `json:"VatInvoiceItemInfos,omitnil,omitempty" name:"VatInvoiceItemInfos"`

	// Machine-printed invoice code
	CodeConfirm *string `json:"CodeConfirm,omitnil,omitempty" name:"CodeConfirm"`

	// Payee
	Receiptor *string `json:"Receiptor,omitnil,omitempty" name:"Receiptor"`


	ElectronicFullMark *int64 `json:"ElectronicFullMark,omitnil,omitempty" name:"ElectronicFullMark"`


	ElectronicFullNumber *string `json:"ElectronicFullNumber,omitnil,omitempty" name:"ElectronicFullNumber"`


	FormName *string `json:"FormName,omitnil,omitempty" name:"FormName"`
}

type VatInvoiceItemInfo struct {
	// Item name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Specification
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// Unit
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// Unit price
	Price *string `json:"Price,omitnil,omitempty" name:"Price"`

	// Amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`

	// Tax amount
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Start date
	DateStart *string `json:"DateStart,omitnil,omitempty" name:"DateStart"`

	// End date
	DateEnd *string `json:"DateEnd,omitnil,omitempty" name:"DateEnd"`

	// License plate number
	LicensePlate *string `json:"LicensePlate,omitnil,omitempty" name:"LicensePlate"`

	// Vehicle type
	VehicleType *string `json:"VehicleType,omitnil,omitempty" name:"VehicleType"`
}

type VatInvoiceRoll struct {
	// Invoice title
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Invoice code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Invoice number
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Machine-printed invoice number
	NumberConfirm *string `json:"NumberConfirm,omitnil,omitempty" name:"NumberConfirm"`

	// Date of issue
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Check code
	CheckCode *string `json:"CheckCode,omitnil,omitempty" name:"CheckCode"`

	// Seller's name
	Seller *string `json:"Seller,omitnil,omitempty" name:"Seller"`

	// Seller's taxpayer identification number
	SellerTaxID *string `json:"SellerTaxID,omitnil,omitempty" name:"SellerTaxID"`

	// Buyer's name
	Buyer *string `json:"Buyer,omitnil,omitempty" name:"Buyer"`

	// Buyer's taxpayer identification number
	BuyerTaxID *string `json:"BuyerTaxID,omitnil,omitempty" name:"BuyerTaxID"`

	// Category
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// Total amount (in figures)
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Total amount (in words)
	TotalCn *string `json:"TotalCn,omitnil,omitempty" name:"TotalCn"`

	// Invoice type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// City
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Whether there is a company seal (0: No, 1: Yes)
	CompanySealMark *int64 `json:"CompanySealMark,omitnil,omitempty" name:"CompanySealMark"`

	// Whether there is a QR code (0: No, 1: Yes)
	QRCodeMark *int64 `json:"QRCodeMark,omitnil,omitempty" name:"QRCodeMark"`

	// Service type
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Company seal content
	CompanySealContent *string `json:"CompanySealContent,omitnil,omitempty" name:"CompanySealContent"`

	// Tax authority seal content
	TaxSealContent *string `json:"TaxSealContent,omitnil,omitempty" name:"TaxSealContent"`

	// Items
	VatRollItems []*VatRollItem `json:"VatRollItems,omitnil,omitempty" name:"VatRollItems"`
}

type VatRollItem struct {
	// Item name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Quantity
	Quantity *string `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// Unit price
	Price *string `json:"Price,omitnil,omitempty" name:"Price"`

	// Amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`
}

// Predefined struct for user
type VinOCRRequestParams struct {
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type VinOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided. If both are provided, only `ImageUrl` is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image.
	// Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported.
	// Supported image size: The downloaded image after Base64 encoding can be up to 7 MB. The download time of the image cannot exceed 3s.
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
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
	Vin *string `json:"Vin,omitnil,omitempty" name:"Vin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DetectedText *string `json:"DetectedText,omitnil,omitempty" name:"DetectedText"`

	// The coordinates of the four vertices.
	Coord *Polygon `json:"Coord,omitnil,omitempty" name:"Coord"`

	// Description.
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// Specifies the four-point coordinate of the word.
	WordCoord []*WordPolygon `json:"WordCoord,omitnil,omitempty" name:"WordCoord"`
}

type WordPolygon struct {
	// The text content.
	DetectedText *string `json:"DetectedText,omitnil,omitempty" name:"DetectedText"`

	// The coordinates of the four vertices.
	Coord *Polygon `json:"Coord,omitnil,omitempty" name:"Coord"`
}