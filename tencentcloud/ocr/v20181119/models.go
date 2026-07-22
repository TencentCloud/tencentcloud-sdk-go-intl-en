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

type AddressInfo struct {
	// Country.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// Postal code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// Sub-region or state/province.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Subdivision *string `json:"Subdivision,omitnil,omitempty" name:"Subdivision"`

	// District or county.
	// Note: This field may return null, indicating that no valid values can be obtained.
	District *string `json:"District,omitnil,omitempty" name:"District"`

	// City name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Subdistrict or township.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Subdistrict *string `json:"Subdistrict,omitnil,omitempty" name:"Subdistrict"`

	// Formatted complete address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormattedAddress *string `json:"FormattedAddress,omitnil,omitempty" name:"FormattedAddress"`

	// First line of the address bar.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LineOne *string `json:"LineOne,omitnil,omitempty" name:"LineOne"`

	// Second line of the address bar.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LineTwo *string `json:"LineTwo,omitnil,omitempty" name:"LineTwo"`

	// Specifies the third line of the address bar.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LineThree *string `json:"LineThree,omitnil,omitempty" name:"LineThree"`

	// Specifies the fourth line of the address bar.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LineFour *string `json:"LineFour,omitnil,omitempty" name:"LineFour"`

	// Specifies the fifth line of the address bar.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LineFive *string `json:"LineFive,omitnil,omitempty" name:"LineFive"`
}

type AnalyzedLog struct {
	// <p>Indexes of the procedure.</p><p>Enumeration value:</p><ul><li>L1_IMAGE_QUALITY: Image quality detection</li><li>L2_RULE_ENGINE: Rule verification</li><li>L3_LLM_JUDGE: Large model judgment</li></ul>
	StepKey *string `json:"StepKey,omitnil,omitempty" name:"StepKey"`

	// <p>Compliant and non-compliant are final states; to be determined is an intermediate state. Judgment status of each layer: Compliant Non-Compliant Pending.</p>
	Decision *string `json:"Decision,omitnil,omitempty" name:"Decision"`

	// <p>Reason for the current layer judgment</p>
	DecisionMessage *string `json:"DecisionMessage,omitnil,omitempty" name:"DecisionMessage"`

	// <p>Time taken by the current layer</p><p>Unit: ms</p>
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`
}

// Predefined struct for user
type ApplyCardVerificationExternalRequestParams struct {
	// Country/Region of the document. For the full list of supported countries/regions, refer to the API description.
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Document type. Supported values: ID_CARD, PASSPORT, DRIVING_LICENSE, RESIDENCE_PERMIT (only supported in certain countries/regions, including Australia, Canada, Germany, New Zealand, Nigeria, Singapore).
	CardType *string `json:"CardType,omitnil,omitempty" name:"CardType"`

	// Base64-encoded image of the document front.
	// Supported image formats: PNG, JPG/JPEG (GIF not supported).
	// Supported image size: The downloaded image after Base64 encoding must not exceed 2 MB. Image download time must not exceed 5 seconds.
	// Supported image resolution: Between 256*256 and 4096*4096 pixels.
	// Note: You must provide either ImageUrlFront or ImageBase64Front. If both are provided, only ImageUrlFront is used.
	ImageBase64Front *string `json:"ImageBase64Front,omitnil,omitempty" name:"ImageBase64Front"`

	// The Base64 value of the reverse side of the document. Supported image formats: PNG, JPG/JPEG. 
	// Supported image size: the downloaded image after Base64 encoding must be no more than 2M. Image download time must be no more than 5 seconds. 
	// Supported image resolution: between 256 \* 256 and 4096 \* 4096. For some documents, either ImageUrlBack or ImageBase64Back must be provided. If both are provided, only ImageUrlBack is used.
	ImageBase64Back *string `json:"ImageBase64Back,omitnil,omitempty" name:"ImageBase64Back"`

	// URL of the document front image.
	// Supported image formats: PNG, JPG/JPEG (GIF not supported).
	// Supported image size: The downloaded image after Base64 encoding must not exceed 2 MB. Image download time must not exceed 5 seconds.
	// Supported image resolution: Between 256*256 and 4096*4096 pixels.
	// Note: You must provide either ImageUrlFront or ImageBase64Front. If both are provided, only ImageUrlFront is used.
	ImageUrlFront *string `json:"ImageUrlFront,omitnil,omitempty" name:"ImageUrlFront"`

	// URL of the document back image.
	// Supported image formats: PNG, JPG/JPEG (GIF not supported).
	// Supported image size: The downloaded image after Base64 encoding must not exceed 2 MB. Image download time must not exceed 5 seconds.
	// Supported image resolution: Between 256*256 and 4096*4096 pixels.
	// Note: For some documents, you must provide either ImageUrlBack or ImageBase64Back. If both are provided, only ImageUrlBack is used.
	ImageUrlBack *string `json:"ImageUrlBack,omitnil,omitempty" name:"ImageUrlBack"`

	// Whether to crop and return the face image from the document. Default: false.
	// If set to true, the image constraints are:
	// - Size after Base64 encoding must not exceed 5 MB.
	// - Maximum pixel width/height: 4000 for JPG, 2000 for other formats.
	// - Minimum pixel width/height: 64.
	// - Supported formats: PNG, JPG, JPEG, BMP (GIF not supported).
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type ApplyCardVerificationExternalRequest struct {
	*tchttp.BaseRequest
	
	// Country/Region of the document. For the full list of supported countries/regions, refer to the API description.
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Document type. Supported values: ID_CARD, PASSPORT, DRIVING_LICENSE, RESIDENCE_PERMIT (only supported in certain countries/regions, including Australia, Canada, Germany, New Zealand, Nigeria, Singapore).
	CardType *string `json:"CardType,omitnil,omitempty" name:"CardType"`

	// Base64-encoded image of the document front.
	// Supported image formats: PNG, JPG/JPEG (GIF not supported).
	// Supported image size: The downloaded image after Base64 encoding must not exceed 2 MB. Image download time must not exceed 5 seconds.
	// Supported image resolution: Between 256*256 and 4096*4096 pixels.
	// Note: You must provide either ImageUrlFront or ImageBase64Front. If both are provided, only ImageUrlFront is used.
	ImageBase64Front *string `json:"ImageBase64Front,omitnil,omitempty" name:"ImageBase64Front"`

	// The Base64 value of the reverse side of the document. Supported image formats: PNG, JPG/JPEG. 
	// Supported image size: the downloaded image after Base64 encoding must be no more than 2M. Image download time must be no more than 5 seconds. 
	// Supported image resolution: between 256 \* 256 and 4096 \* 4096. For some documents, either ImageUrlBack or ImageBase64Back must be provided. If both are provided, only ImageUrlBack is used.
	ImageBase64Back *string `json:"ImageBase64Back,omitnil,omitempty" name:"ImageBase64Back"`

	// URL of the document front image.
	// Supported image formats: PNG, JPG/JPEG (GIF not supported).
	// Supported image size: The downloaded image after Base64 encoding must not exceed 2 MB. Image download time must not exceed 5 seconds.
	// Supported image resolution: Between 256*256 and 4096*4096 pixels.
	// Note: You must provide either ImageUrlFront or ImageBase64Front. If both are provided, only ImageUrlFront is used.
	ImageUrlFront *string `json:"ImageUrlFront,omitnil,omitempty" name:"ImageUrlFront"`

	// URL of the document back image.
	// Supported image formats: PNG, JPG/JPEG (GIF not supported).
	// Supported image size: The downloaded image after Base64 encoding must not exceed 2 MB. Image download time must not exceed 5 seconds.
	// Supported image resolution: Between 256*256 and 4096*4096 pixels.
	// Note: For some documents, you must provide either ImageUrlBack or ImageBase64Back. If both are provided, only ImageUrlBack is used.
	ImageUrlBack *string `json:"ImageUrlBack,omitnil,omitempty" name:"ImageUrlBack"`

	// Whether to crop and return the face image from the document. Default: false.
	// If set to true, the image constraints are:
	// - Size after Base64 encoding must not exceed 5 MB.
	// - Maximum pixel width/height: 4000 for JPG, 2000 for other formats.
	// - Minimum pixel width/height: 64.
	// - Supported formats: PNG, JPG, JPEG, BMP (GIF not supported).
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

func (r *ApplyCardVerificationExternalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCardVerificationExternalRequest) FromJsonString(s string) error {
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
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyCardVerificationExternalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyCardVerificationExternalResponseParams struct {
	// Unique token for the verification process, used to retrieve the result.
	CardVerificationToken *string `json:"CardVerificationToken,omitnil,omitempty" name:"CardVerificationToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyCardVerificationExternalResponse struct {
	*tchttp.BaseResponse
	Response *ApplyCardVerificationExternalResponseParams `json:"Response"`
}

func (r *ApplyCardVerificationExternalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCardVerificationExternalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BrazilCardInfo struct {
	// RNE document.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RNE *BrazilRNEInfo `json:"RNE,omitnil,omitempty" name:"RNE"`

	// Specifies the document.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RNM *BrazilRNMInfo `json:"RNM,omitnil,omitempty" name:"RNM"`

	// Driver license document.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DriverLicense *BrazilDriverLicenseInfo `json:"DriverLicense,omitnil,omitempty" name:"DriverLicense"`

	// ID card document.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IDCard *BrazilIDCardInfo `json:"IDCard,omitnil,omitempty" name:"IDCard"`
}

type BrazilDriverLicenseInfo struct {
	// Name.
	NOME *string `json:"NOME,omitnil,omitempty" name:"NOME"`

	// Specifies the driver's license type.
	CatHab *string `json:"CatHab,omitnil,omitempty" name:"CatHab"`

	// Driver’s license id.
	CNHNumber *string `json:"CNHNumber,omitnil,omitempty" name:"CNHNumber"`

	// Expiration date.
	VALIDADE *string `json:"VALIDADE,omitnil,omitempty" name:"VALIDADE"`

	// Indicates the qualification.
	QUALIFICATION *string `json:"QUALIFICATION,omitnil,omitempty" name:"QUALIFICATION"`

	// Identity card number.
	IDENTIDADE *string `json:"IDENTIDADE,omitnil,omitempty" name:"IDENTIDADE"`

	// Tax number of the person.
	CPF *string `json:"CPF,omitnil,omitempty" name:"CPF"`

	// Date of birth.
	NASCIMENTO *string `json:"NASCIMENTO,omitnil,omitempty" name:"NASCIMENTO"`

	// Membership status.
	MEMBERSHIP *string `json:"MEMBERSHIP,omitnil,omitempty" name:"MEMBERSHIP"`

	// Registration number.
	REGISTRO *string `json:"REGISTRO,omitnil,omitempty" name:"REGISTRO"`

	// Remarks.
	OBSERVATIONS *string `json:"OBSERVATIONS,omitnil,omitempty" name:"OBSERVATIONS"`

	// Issue date.
	IssueDate *string `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`

	// Issuing location.
	LOCAL *string `json:"LOCAL,omitnil,omitempty" name:"LOCAL"`

	// Record number.
	BackNumber *string `json:"BackNumber,omitnil,omitempty" name:"BackNumber"`

	// Specifies the avatar in base64 format.
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`
}

type BrazilIDCardInfo struct {
	// Name.
	Nome *string `json:"Nome,omitnil,omitempty" name:"Nome"`

	// Parent information.
	MemberShip *string `json:"MemberShip,omitnil,omitempty" name:"MemberShip"`

	// Date of birth.
	DataNascimento *string `json:"DataNascimento,omitnil,omitempty" name:"DataNascimento"`

	// Issuing organization.
	IssuingAgency *string `json:"IssuingAgency,omitnil,omitempty" name:"IssuingAgency"`

	// Specifies the blood type.
	Fatorrh *string `json:"Fatorrh,omitnil,omitempty" name:"Fatorrh"`

	// Birthplace.
	NaturalIDade *string `json:"NaturalIDade,omitnil,omitempty" name:"NaturalIDade"`

	// Remarks.
	Observations *string `json:"Observations,omitnil,omitempty" name:"Observations"`

	// CPF Type
	CPF *string `json:"CPF,omitnil,omitempty" name:"CPF"`

	// DNI Type
	DNI *string `json:"DNI,omitnil,omitempty" name:"DNI"`

	// Common registration.
	RegistroGeral *string `json:"RegistroGeral,omitnil,omitempty" name:"RegistroGeral"`

	// Issue date. valid values: dd/mm/yyyy.
	DispatchDate *string `json:"DispatchDate,omitnil,omitempty" name:"DispatchDate"`

	// Address.
	Registro *string `json:"Registro,omitnil,omitempty" name:"Registro"`

	// Specifies the avatar in Base64 format of the id card.
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// Original identity information.
	DocOrigem *string `json:"DocOrigem,omitnil,omitempty" name:"DocOrigem"`
}

type BrazilRNEInfo struct {
	// RNE
	RNE *string `json:"RNE,omitnil,omitempty" name:"RNE"`

	// CLASSIFICAÇÃO(CLASSIFICATION)
	CLASSIFICATION *string `json:"CLASSIFICATION,omitnil,omitempty" name:"CLASSIFICATION"`

	// VALIDADE
	VALIDADE *string `json:"VALIDADE,omitnil,omitempty" name:"VALIDADE"`

	// NOME
	NOME *string `json:"NOME,omitnil,omitempty" name:"NOME"`

	// FILIAÇÃO(MEMBERSHIP)
	Membership *string `json:"Membership,omitnil,omitempty" name:"Membership"`

	// NACIONALIDADE
	NACIONALIDADE *string `json:"NACIONALIDADE,omitnil,omitempty" name:"NACIONALIDADE"`

	// NATURALIDADE(PAÍS)
	NATURALIDADE *string `json:"NATURALIDADE,omitnil,omitempty" name:"NATURALIDADE"`

	// ORGÃO EXPEDIDOR(IssuingAgency)
	IssuingAgency *string `json:"IssuingAgency,omitnil,omitempty" name:"IssuingAgency"`

	// DATA DE NASCIMENTO(DateOfBirth)
	DateOfBirth *string `json:"DateOfBirth,omitnil,omitempty" name:"DateOfBirth"`

	// SEXO
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// DATA DE ENTRADA(EntryDate)
	EntryDate *string `json:"EntryDate,omitnil,omitempty" name:"EntryDate"`

	// VIA
	VIA *string `json:"VIA,omitnil,omitempty" name:"VIA"`

	// DATA DE EXPEDIÇÃO(DispatchDate)
	DispatchDate *string `json:"DispatchDate,omitnil,omitempty" name:"DispatchDate"`

	// MRZ
	MRZ *string `json:"MRZ,omitnil,omitempty" name:"MRZ"`

	// PortraitImage
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`
}

type BrazilRNMInfo struct {
	// SOBRENOME Type
	SOBRENOME *string `json:"SOBRENOME,omitnil,omitempty" name:"SOBRENOME"`

	// NOME Type
	NOME *string `json:"NOME,omitnil,omitempty" name:"NOME"`

	// DATA DE NASCIMENTO
	DATADENASCIMENTO *string `json:"DATADENASCIMENTO,omitnil,omitempty" name:"DATADENASCIMENTO"`

	// SEXO F
	SEXO *string `json:"SEXO,omitnil,omitempty" name:"SEXO"`

	// FILIAÇÃO(MEMBERSHIP)
	MEMBERSHIP *string `json:"MEMBERSHIP,omitnil,omitempty" name:"MEMBERSHIP"`

	// NACIONALIDADE Type
	NACIONALIDADE *string `json:"NACIONALIDADE,omitnil,omitempty" name:"NACIONALIDADE"`

	// VALIDADE Type
	VALIDADE *string `json:"VALIDADE,omitnil,omitempty" name:"VALIDADE"`

	// RNM Type
	RNM *string `json:"RNM,omitnil,omitempty" name:"RNM"`

	// CPF Type
	CPF *string `json:"CPF,omitnil,omitempty" name:"CPF"`

	// CLASSIFICAÇÃO(CLASSIFICATION)
	CLASSIFICATION *string `json:"CLASSIFICATION,omitnil,omitempty" name:"CLASSIFICATION"`

	// PRAZO DE RESIDENCIA
	PRAZODERESIDENCIA *string `json:"PRAZODERESIDENCIA,omitnil,omitempty" name:"PRAZODERESIDENCIA"`

	// EMISSÃO(ISSUANCE)
	ISSUANCE *string `json:"ISSUANCE,omitnil,omitempty" name:"ISSUANCE"`

	// AMPARO LEGAL(LegalHelp)
	AMPAROLEGAL *string `json:"AMPAROLEGAL,omitnil,omitempty" name:"AMPAROLEGAL"`

	// MRZCode
	MRZ *string `json:"MRZ,omitnil,omitempty" name:"MRZ"`

	// Portrait Image
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// PortraitImage(Back)
	PortraitImageBack *string `json:"PortraitImageBack,omitnil,omitempty" name:"PortraitImageBack"`
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

type CoordsItem struct {
	// Coordinates of four points in the image.
	Polygon *Polygon `json:"Polygon,omitnil,omitempty" name:"Polygon"`

	// Coordinates of two points in the image.
	Coords *ItemCoord `json:"Coords,omitnil,omitempty" name:"Coords"`
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
	// -- OverseasCheques
	// -- RegistrationCertificate
	// -- GridPhoto
	// -- SignaturePage
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
	// -- OverseasCheques
	// -- RegistrationCertificate
	// -- GridPhoto
	// -- SignaturePage
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

// Predefined struct for user
type GeneralAccurateOCRRequestParams struct {
	// <p>The Base64 value of the image/PDF. The image size after Base64 encoding must be no more than 10M, with a resolution of 600*800 or higher recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported. Either ImageUrl or ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.</p>
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. 
	// The image cannot exceed 10 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// <p>Whether to enable original image slicing detection. Once enabled, it improves recognition accuracy in scenarios where "the overall image area is large but the single character occupies a small proportion" (for example: exam paper). Default: disabled. Note: Only supported when ConfigID is configured as OCR.</p>
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitnil,omitempty" name:"EnableDetectSplit"`

	// <p>Whether PDF recognition is enabled. The default value is false. Once enabled, it can simultaneously support image and PDF recognition.</p>
	IsPdf *bool `json:"IsPdf,omitnil,omitempty" name:"IsPdf"`

	// <p>The corresponding page number of the PDF page to be recognized. Only single page recognition is supported. Valid at that time when the upload file is a PDF and the IsPdf parameter value is true. The default value is 1.</p>
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// <p>Text Detection Switch, true by default. Set to false to directly perform single-line recognition, suitable for image scenarios containing only forward single-line text.</p>
	EnableDetectText *bool `json:"EnableDetectText,omitnil,omitempty" name:"EnableDetectText"`

	// <p>Configuration ID support: OCR -- common scenarios MulOCR -- multilingual scenario. Default value: OCR.</p>
	ConfigID *string `json:"ConfigID,omitnil,omitempty" name:"ConfigID"`
}

type GeneralAccurateOCRRequest struct {
	*tchttp.BaseRequest
	
	// <p>The Base64 value of the image/PDF. The image size after Base64 encoding must be no more than 10M, with a resolution of 600*800 or higher recommended. PNG, JPG, JPEG, BMP, and PDF formats are supported. Either ImageUrl or ImageBase64 of the image must be provided. If both are provided, only ImageUrl will be used.</p>
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// URL address of image. 
	// The image cannot exceed 10 MB after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// <p>Whether to enable original image slicing detection. Once enabled, it improves recognition accuracy in scenarios where "the overall image area is large but the single character occupies a small proportion" (for example: exam paper). Default: disabled. Note: Only supported when ConfigID is configured as OCR.</p>
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitnil,omitempty" name:"EnableDetectSplit"`

	// <p>Whether PDF recognition is enabled. The default value is false. Once enabled, it can simultaneously support image and PDF recognition.</p>
	IsPdf *bool `json:"IsPdf,omitnil,omitempty" name:"IsPdf"`

	// <p>The corresponding page number of the PDF page to be recognized. Only single page recognition is supported. Valid at that time when the upload file is a PDF and the IsPdf parameter value is true. The default value is 1.</p>
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil,omitempty" name:"PdfPageNumber"`

	// <p>Text Detection Switch, true by default. Set to false to directly perform single-line recognition, suitable for image scenarios containing only forward single-line text.</p>
	EnableDetectText *bool `json:"EnableDetectText,omitnil,omitempty" name:"EnableDetectText"`

	// <p>Configuration ID support: OCR -- common scenarios MulOCR -- multilingual scenario. Default value: OCR.</p>
	ConfigID *string `json:"ConfigID,omitnil,omitempty" name:"ConfigID"`
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
	delete(f, "EnableDetectSplit")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	delete(f, "EnableDetectText")
	delete(f, "ConfigID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneralAccurateOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GeneralAccurateOCRResponseParams struct {
	// <p>Detected text information, including row content, confidence degree, text line coordinate, and rotation corrected coordinate. For specific content, please click the left-side link.</p>
	TextDetections []*TextDetection `json:"TextDetections,omitnil,omitempty" name:"TextDetections"`

	// Image rotation angle in degrees. Zero degrees: The horizontal direction of the text on the image; a positive value: rotate clockwise; a negative value: rotate counterclockwise.
	//
	// Deprecated: Angel is deprecated.
	Angel *float64 `json:"Angel,omitnil,omitempty" name:"Angel"`

	// <p>Image rotation angle (angle system), the text's horizontal direction is Zero degrees; clockwise is positive, counterclockwise is negative. Click to view <a href="https://www.tencentcloud.com/document/product/866/45139?from_cn_redirect=1">How to correct tilt text</a></p>
	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Image rotation angle in degrees. 0: The horizontal direction of the text on the image; a positive value: rotate clockwise; a negative value: rotate counterclockwise.
	//
	// Deprecated: Angel is deprecated.
	Angel *float64 `json:"Angel,omitnil,omitempty" name:"Angel"`

	// Total number of PDF pages to be returned if the image is a PDF. Default value: 0.
	PdfPageSize *int64 `json:"PdfPageSize,omitnil,omitempty" name:"PdfPageSize"`


	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

type GeneralCard struct {
	// ID number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LicenseNumber *string `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// Personal number. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	PersonalNumber *string `json:"PersonalNumber,omitnil,omitempty" name:"PersonalNumber"`

	// Full name on the document.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Full name in local language.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullNameLocal *string `json:"FullNameLocal,omitnil,omitempty" name:"FullNameLocal"`

	// First name or given name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// First name in local language.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstNameLocal *string `json:"FirstNameLocal,omitnil,omitempty" name:"FirstNameLocal"`

	// Middle name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MiddleName *string `json:"MiddleName,omitnil,omitempty" name:"MiddleName"`

	// Middle name in local language.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MiddleNameLocal *string `json:"MiddleNameLocal,omitnil,omitempty" name:"MiddleNameLocal"`

	// Last name or surname.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// Last name in local language.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastNameLocal *string `json:"LastNameLocal,omitnil,omitempty" name:"LastNameLocal"`

	// Gender on the document.
	// - M: man.
	// - F: woman.
	// - X: other gender identity.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Date of birth.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Birth place.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BirthPlace *string `json:"BirthPlace,omitnil,omitempty" name:"BirthPlace"`

	// Issue date.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedDate *string `json:"IssuedDate,omitnil,omitempty" name:"IssuedDate"`

	// Issuing authority.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedAuthority *string `json:"IssuedAuthority,omitnil,omitempty" name:"IssuedAuthority"`

	// Place of issue.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedPlace *string `json:"IssuedPlace,omitnil,omitempty" name:"IssuedPlace"`

	// Issuing country.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedCountry *string `json:"IssuedCountry,omitnil,omitempty" name:"IssuedCountry"`

	// Country code of issue, ISO Alpha-3 format.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuedCountryCode *string `json:"IssuedCountryCode,omitnil,omitempty" name:"IssuedCountryCode"`

	// Expiry date.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

	// First line of the Machine Readable Zone (MRZ).
	// Note: This field may return null, indicating that no valid values can be obtained.
	MRZLine1 *string `json:"MRZLine1,omitnil,omitempty" name:"MRZLine1"`

	// Second line of the Machine Readable Zone (MRZ).
	// Note: This field may return null, indicating that no valid values can be obtained.
	MRZLine2 *string `json:"MRZLine2,omitnil,omitempty" name:"MRZLine2"`

	// Document nationality, following ISO 3166 country coding specification.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Address information on the document.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *AddressInfo `json:"Address,omitnil,omitempty" name:"Address"`

	// Religion (if displayed on the document).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Religion *string `json:"Religion,omitnil,omitempty" name:"Religion"`

	// Type of document.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Blood type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BloodType *string `json:"BloodType,omitnil,omitempty" name:"BloodType"`

	// Height.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// Weight.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Weight *string `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Vehicle class authorized on the driver license (e.g., A, B, C).
	// Note: This field may return null, indicating that no valid values can be obtained.
	VehicleClass *string `json:"VehicleClass,omitnil,omitempty" name:"VehicleClass"`

	// Restrictions on the driver license.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Restrictions *string `json:"Restrictions,omitnil,omitempty" name:"Restrictions"`

	// Endorsements or additional records on the driver license.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Endorsement *string `json:"Endorsement,omitnil,omitempty" name:"Endorsement"`

	// Supplementary fields (varies by document type).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Others *string `json:"Others,omitnil,omitempty" name:"Others"`

	// First line of the passport MRZ (Machine Readable Zone).
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: PassportCodeFirst is deprecated.
	PassportCodeFirst *string `json:"PassportCodeFirst,omitnil,omitempty" name:"PassportCodeFirst"`

	// Second line of the passport MRZ (Machine Readable Zone).
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: PassportCodeSecond is deprecated.
	PassportCodeSecond *string `json:"PassportCodeSecond,omitnil,omitempty" name:"PassportCodeSecond"`

	// Expiry date.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: DueDate is deprecated.
	DueDate *string `json:"DueDate,omitnil,omitempty" name:"DueDate"`

	// Age. 0 means no valid info.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: Age is deprecated.
	Age *string `json:"Age,omitnil,omitempty" name:"Age"`

	// Registration number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: RegistrationNumber is deprecated.
	RegistrationNumber *string `json:"RegistrationNumber,omitnil,omitempty" name:"RegistrationNumber"`
}

// Predefined struct for user
type GetCardVerificationExternalResultRequestParams struct {
	// Initiates the recognition interface and returns a unique token.
	CardVerificationToken *string `json:"CardVerificationToken,omitnil,omitempty" name:"CardVerificationToken"`
}

type GetCardVerificationExternalResultRequest struct {
	*tchttp.BaseRequest
	
	// Initiates the recognition interface and returns a unique token.
	CardVerificationToken *string `json:"CardVerificationToken,omitnil,omitempty" name:"CardVerificationToken"`
}

func (r *GetCardVerificationExternalResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCardVerificationExternalResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CardVerificationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCardVerificationExternalResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCardVerificationExternalResultResponseParams struct {
	// Verification status. Valid values: 
	// PROCESSING
	// ABNORMAL
	// COMPLETED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Anti-counterfeiting information. 
	// - ScreenshotSuspected: The image is a screenshot.
	// - RetakeSuspected: The image is taken from another screen.
	// - PaperCopy: The image is a black and white, or color photocopy.
	// - FakeSuspected: The image of the card, or the information on the card has been edited or altered.
	// - PoorImageQuality: The image is bad quality.
	// - InformationVerificationFailed: Information verification failed based on OCR recognition results
	// - TooManyCards: Multiple cards present in the frame.
	// - IncompleteCard: Captured document is incomplete.
	// - OtherWarning: Document's authenticity is not verified for various reasons.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	WarnInfo []*string `json:"WarnInfo,omitnil,omitempty" name:"WarnInfo"`

	// Country or region of the document.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: Nationality is deprecated.
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Front-side document recognition results. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	CardInfo *GeneralCard `json:"CardInfo,omitnil,omitempty" name:"CardInfo"`

	// Back-side document recognition results.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackCardInfo *GeneralCard `json:"BackCardInfo,omitnil,omitempty" name:"BackCardInfo"`

	// The token passed in the input parameters.
	CardVerificationToken *string `json:"CardVerificationToken,omitnil,omitempty" name:"CardVerificationToken"`

	// Base64-encoded head image from the document. If ReturnHeadImage was set to false or not provided in the request, this field returns an empty string. If ReturnHeadImage was set to true and this field returns an empty string, indicating a failure to extract the head image extraction failed. Please check the input document photo.
	HeadImageBase64 *string `json:"HeadImageBase64,omitnil,omitempty" name:"HeadImageBase64"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCardVerificationExternalResultResponse struct {
	*tchttp.BaseResponse
	Response *GetCardVerificationExternalResultResponseParams `json:"Response"`
}

func (r *GetCardVerificationExternalResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCardVerificationExternalResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Telecode of the name in Chinese
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

	// Date of first issue
	FirstIssueDate *string `json:"FirstIssueDate,omitnil,omitempty" name:"FirstIssueDate"`

	// Date of last receipt
	CurrentIssueDate *string `json:"CurrentIssueDate,omitnil,omitempty" name:"CurrentIssueDate"`

	// Authenticity check.
	// 0: unable to judge (because the image is blurred, incomplete, reflective, too dark, etc.);
	// 1: forged;
	// 2: authentic.
	// Note: this field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: FakeDetectResult is deprecated.
	FakeDetectResult *int64 `json:"FakeDetectResult,omitnil,omitempty" name:"FakeDetectResult"`

	// Base64-encoded large primary portrait photo from the left side of the new-generation Hong Kong Identity Card.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeadImage *string `json:"HeadImage,omitnil,omitempty" name:"HeadImage"`

	// Base64-encoded small secondary portrait photo from the right side of the new-generation Hong Kong Identity Card.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SmallHeadImage *string `json:"SmallHeadImage,omitnil,omitempty" name:"SmallHeadImage"`

	// This field is deprecated and will always return an empty array. Usage is not recommended.
	//
	// Deprecated: WarningCode is deprecated.
	WarningCode []*int64 `json:"WarningCode,omitnil,omitempty" name:"WarningCode"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate
	// -9102 Alarm for photocopied certificate
	// -9103 Alarm for photographed certificate
	// -9104 Alarm for tamper certificate
	// -9107 Alarm for reflective certificate
	// -9108 Alarm for blurry certificate 
	// -9109 This capability is not enabled. Please contact customer support to activate the alert service.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// Text information incorporated within the laser-perforated see-through window on the new-generation Hong Kong Identity Card.
	WindowEmbeddedText *string `json:"WindowEmbeddedText,omitnil,omitempty" name:"WindowEmbeddedText"`

	// Versions of the Hong Kong Identity Card: HKID-2018, HKID-2003.
	HKIDVersion *string `json:"HKIDVersion,omitnil,omitempty" name:"HKIDVersion"`

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

	// This field is deprecated and will always return an empty array. Usage is not recommended.
	//
	// Deprecated: Warn is deprecated.
	Warn []*int64 `json:"Warn,omitnil,omitempty" name:"Warn"`

	// Identity photo
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// This field is deprecated and will always return "1". Usage is not recommended.
	//
	// Deprecated: AdvancedInfo is deprecated.
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// Certificate type: 
	// - MyKad: ID card 
	// - MyPR: Permanent resident card 
	// - MyTentera: Military identity card 
	// - MyKAS: Temporary ID card 
	// - POLIS: Police card 
	// - IKAD: Work permit 
	// - MyKid: Child ID card
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Date of birth. This field is available only for work permits (i-Kad) and ID cards (MyKad).
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Number on the back of the Malaysian ID card
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
	// Base64-encoded image data. The image must be no larger than 7 MB after Base64 encoding. A resolution of at least 500x800 is recommended. Supported image formats: PNG, JPG, JPEG, BMP, and PDF. The document should occupy more than 2/3 of the image area.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Whether to return an image. 
	// Default value: false.
	RetImage *bool `json:"RetImage,omitnil,omitempty" name:"RetImage"`

	// URL of the image. The downloaded image must be no larger than 7 MB after Base64 encoding. A resolution of at least 500x800 is recommended. Supported image formats: PNG, JPG, JPEG, BMP, and PDF. The document should occupy more than 2/3 of the image area. Image download must complete within 3 seconds. We recommend storing images in Tencent Cloud for higher download speed and stability. The speed and stability of URLs from non-Tencent Cloud storage may be affected. Note: This field is not supported outside the Chinese mainland region.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type MLIDPassportOCRRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded image data. The image must be no larger than 7 MB after Base64 encoding. A resolution of at least 500x800 is recommended. Supported image formats: PNG, JPG, JPEG, BMP, and PDF. The document should occupy more than 2/3 of the image area.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Whether to return an image. 
	// Default value: false.
	RetImage *bool `json:"RetImage,omitnil,omitempty" name:"RetImage"`

	// URL of the image. The downloaded image must be no larger than 7 MB after Base64 encoding. A resolution of at least 500x800 is recommended. Supported image formats: PNG, JPG, JPEG, BMP, and PDF. The document should occupy more than 2/3 of the image area. Image download must complete within 3 seconds. We recommend storing images in Tencent Cloud for higher download speed and stability. The speed and stability of URLs from non-Tencent Cloud storage may be affected. Note: This field is not supported outside the Chinese mainland region.
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

	// Nationality code (MRZ field)
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// This field is deprecated and will always return an empty array. Usage is not recommended.
	//
	// Deprecated: Warn is deprecated.
	Warn []*int64 `json:"Warn,omitnil,omitempty" name:"Warn"`

	// Base64-encoded identity photo
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// This field is deprecated and will always return "1". Usage is not recommended.
	//
	// Deprecated: AdvancedInfo is deprecated.
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

	// Document content in the visual zone
	PassportRecognizeInfos *PassportRecognizeInfos `json:"PassportRecognizeInfos,omitnil,omitempty" name:"PassportRecognizeInfos"`

	// Warning information for the document. This field applies only to international site requests and will return an empty array for domestic site requests. Valid warning codes: 
	// -9101 (incomplete card border), 
	// -9102 (photocopied document), 
	// -9103 (re-photographed document), -9104 (PS-altered document), 
	// -9107 (reflective document), 
	// -9108 (blurry image), 
	// -9109 (warning capability not enabled).
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// The number of cards detected in the input image.(Currently supported only in ap-bangkok region)
	CardCount *int64 `json:"CardCount,omitnil,omitempty" name:"CardCount"`

	// Whether the passport information is complete.
	IsComplete *bool `json:"IsComplete,omitnil,omitempty" name:"IsComplete"`

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

	// The side of the document. Valid values: FRONT (front side, default),
	// BACK (back side, only supported for Mainland Travel Permit for inbound visits). 
	// Note: Back side recognition is only supported for the "Mainland Travel Permit for Hong Kong and Macao Residents" , and is not supported for Hong Kong, Macao, or Taiwan passes.
	CardSide *string `json:"CardSide,omitnil,omitempty" name:"CardSide"`
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

	// The side of the document. Valid values: FRONT (front side, default),
	// BACK (back side, only supported for Mainland Travel Permit for inbound visits). 
	// Note: Back side recognition is only supported for the "Mainland Travel Permit for Hong Kong and Macao Residents" , and is not supported for Hong Kong, Macao, or Taiwan passes.
	CardSide *string `json:"CardSide,omitnil,omitempty" name:"CardSide"`
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
	delete(f, "CardSide")
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

	// Document type, such as: Mainland Travel Permit for Taiwan Residents, Mainland Travel Permit for Hong Kong and Macao Residents, or Exit-Entry Permit for Travelling to and from Hong Kong and Macao.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Base64-encoded profile photo, which is returned only when `RetProfile` is set to `True`
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// Nationality of the document holder.
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Information on the back of the document. 
	// Note: Only supported for the back side of the Mainland Travel Permit for Hong Kong and Macao Residents.
	MainlandTravelPermitBackInfos *MainlandTravelPermitBackInfos `json:"MainlandTravelPermitBackInfos,omitnil,omitempty" name:"MainlandTravelPermitBackInfos"`

	// Card Warning Information
	// 
	// -9102 Alarm for photocopied certificate
	// -9103 Alarm for photographed certificate
	// -9104 Alarm for tamper certificate
	// -9109 This capability is not enabled.This capability is not enabled. Please contact customer support to activate the alert service.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

type MainlandTravelPermitBackInfos struct {

	Type *string `json:"Type,omitnil,omitempty" name:"Type"`


	Name *string `json:"Name,omitnil,omitempty" name:"Name"`


	IDNumber *string `json:"IDNumber,omitnil,omitempty" name:"IDNumber"`


	HistoryNumber *string `json:"HistoryNumber,omitnil,omitempty" name:"HistoryNumber"`
}

// Predefined struct for user
type PODAuditAIRequestParams struct {
	// <p>The Base64 value of the image/PDF. The Base64 must be no more than 10M, resolution is recommended to be 600*800 or higher, and supports PNG, JPG, JPEG, BMP, PDF formats. Either ImageUrl or ImageBase64 of the image must be provided. If both are provided, only use ImageUrl. Example value: /9j/4AAQSkZJRg.....s97n//2Q==</p>
	ImageBase64List []*string `json:"ImageBase64List,omitnil,omitempty" name:"ImageBase64List"`

	// <p>The Url address of the image/PDF. The image after Base64 encoding should be no more than 10M, with a resolution of 600*800 or higher, and supports PNG, JPG, JPEG, BMP, and PDF formats. The image download time should not exceed 3 seconds. Images stored in Tencent Cloud's Url can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud. The speed and stability of non-Tencent Cloud storage URLs may be impacted to a certain extent. Example value: https://ocr-demo-1254418846.cos.ap-guangzhou.myqcloud.com/general/GeneralAccurateOCR/GeneralAccurateOCR1.JPG</p>
	ImageUrlList []*string `json:"ImageUrlList,omitnil,omitempty" name:"ImageUrlList"`

	// <p>Waybill number is used for pod rule review assistance</p>
	WaybillNumber *string `json:"WaybillNumber,omitnil,omitempty" name:"WaybillNumber"`

	// <p>No      The acknowledge type, 0 is selected by default</p><p>Enumeration value:</p><ul><li>0: Doorstep/yard</li><li>1: Parcel reception room</li><li>2: Myself/others acknowledge</li><li>3: Front desk/reception</li><li>4: Express delivery collection point</li><li>5: Express cabinet</li></ul>
	SignType *int64 `json:"SignType,omitnil,omitempty" name:"SignType"`
}

type PODAuditAIRequest struct {
	*tchttp.BaseRequest
	
	// <p>The Base64 value of the image/PDF. The Base64 must be no more than 10M, resolution is recommended to be 600*800 or higher, and supports PNG, JPG, JPEG, BMP, PDF formats. Either ImageUrl or ImageBase64 of the image must be provided. If both are provided, only use ImageUrl. Example value: /9j/4AAQSkZJRg.....s97n//2Q==</p>
	ImageBase64List []*string `json:"ImageBase64List,omitnil,omitempty" name:"ImageBase64List"`

	// <p>The Url address of the image/PDF. The image after Base64 encoding should be no more than 10M, with a resolution of 600*800 or higher, and supports PNG, JPG, JPEG, BMP, and PDF formats. The image download time should not exceed 3 seconds. Images stored in Tencent Cloud's Url can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud. The speed and stability of non-Tencent Cloud storage URLs may be impacted to a certain extent. Example value: https://ocr-demo-1254418846.cos.ap-guangzhou.myqcloud.com/general/GeneralAccurateOCR/GeneralAccurateOCR1.JPG</p>
	ImageUrlList []*string `json:"ImageUrlList,omitnil,omitempty" name:"ImageUrlList"`

	// <p>Waybill number is used for pod rule review assistance</p>
	WaybillNumber *string `json:"WaybillNumber,omitnil,omitempty" name:"WaybillNumber"`

	// <p>No      The acknowledge type, 0 is selected by default</p><p>Enumeration value:</p><ul><li>0: Doorstep/yard</li><li>1: Parcel reception room</li><li>2: Myself/others acknowledge</li><li>3: Front desk/reception</li><li>4: Express delivery collection point</li><li>5: Express cabinet</li></ul>
	SignType *int64 `json:"SignType,omitnil,omitempty" name:"SignType"`
}

func (r *PODAuditAIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PODAuditAIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64List")
	delete(f, "ImageUrlList")
	delete(f, "WaybillNumber")
	delete(f, "SignType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PODAuditAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PODAuditAIResponseParams struct {
	// <p>0 means non-compliance 1 means compliance</p>
	AuditorDecision *int64 `json:"AuditorDecision,omitnil,omitempty" name:"AuditorDecision"`

	// <p>Reason code for non-compliance. If there are multiple, return a list of multiple codes.</p><p>Enumeration value:</p><ul><li>100: Wrong delivery address</li><li>101: No house number</li><li>104: Single question</li><li>200: No package</li><li>202: Privacy leakage</li></ul>
	FailCode []*string `json:"FailCode,omitnil,omitempty" name:"FailCode"`

	// <p>Entire approval result analysis content</p>
	ResultAnalysis *string `json:"ResultAnalysis,omitnil,omitempty" name:"ResultAnalysis"`

	// <p>Analysis result logs of each layer, including time taken, judgment reason, and judgment conclusion</p>
	AnalyzedLogs []*AnalyzedLog `json:"AnalyzedLogs,omitnil,omitempty" name:"AnalyzedLogs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PODAuditAIResponse struct {
	*tchttp.BaseResponse
	Response *PODAuditAIResponseParams `json:"Response"`
}

func (r *PODAuditAIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PODAuditAIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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


	BirthPlace *string `json:"BirthPlace,omitnil,omitempty" name:"BirthPlace"`
}

// Predefined struct for user
type PermitOCRRequestParams struct {
	// The Base64-encoded value of the image. Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported. Supported image size: The downloaded image after Base64 encoding cannot exceed 7 MB. The download time of the image cannot exceed 3 seconds. Either ImageUrl or ImageBase64 of the image must be provided. If both are provided, only ImageUrl is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported. Supported image size: The downloaded image after Base64 encoding cannot exceed 7 MB. The download time of the image cannot exceed 3 seconds. We recommend that you store the image in Tencent Cloud for higher download speed and stability. The download speed and stability of images stored outside Tencent Cloud may be compromised.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the ID photo. The default value is false.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
}

type PermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64-encoded value of the image. Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported. Supported image size: The downloaded image after Base64 encoding cannot exceed 7 MB. The download time of the image cannot exceed 3 seconds. Either ImageUrl or ImageBase64 of the image must be provided. If both are provided, only ImageUrl is used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. Supported image formats: PNG, JPG, and JPEG. GIF is currently not supported. Supported image size: The downloaded image after Base64 encoding cannot exceed 7 MB. The download time of the image cannot exceed 3 seconds. We recommend that you store the image in Tencent Cloud for higher download speed and stability. The download speed and stability of images stored outside Tencent Cloud may be compromised.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the ID photo. The default value is false.
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
	// Name in Chinese
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// English name
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

	// Base64-encoded profile photo of the document holder.
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// Document type, such as: Exit-Entry Permit for Travelling to and from Hong Kong and Macao, or Exit-Entry Permit for Travelling to and from Taiwan.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate
	// -9102 Alarm for photocopied certificate
	// -9103 Alarm for photographed certificate
	// -9104 Alarm for tamper certificate
	// -9107 Alarm for reflective certificate
	// -9108 Alarm for blurry certificate 
	// -9109 This capability is not enabled. Please contact customer support to activate the alert service.
	WarnCardInfos []*int64 `json:"WarnCardInfos,omitnil,omitempty" name:"WarnCardInfos"`

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

// Predefined struct for user
type RecognizeBrazilCommonOCRRequestParams struct {
	// The Base64 value of the image. supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. supported image size: no more than 7M after encoding. image download time: no more than 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The Url of the image. supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. supported image size: no more than 7M after Base64 encoding. image download time is no more than 3 seconds. urls stored in tencent cloud guarantee higher download speed and stability. it is advisable to store images in tencent cloud. urls not stored in tencent cloud may possibly be impacted in speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The Base64 value of the back side of the card. supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. supported image size: no more than 7M after encoding. image download time: not more than 3 seconds. either ImageUrl or ImageBase64 must be provided. if both are provided, only use ImageUrl.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The Url address of the back side of the card. supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. supported image size: no more than 7M after Base64 encoding. image download time is no more than 3 seconds. urls stored in tencent cloud guarantee higher download speed and stability. it is recommended to store images in tencent cloud. speed and stability of non-tencent cloud storage urls may be impacted.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Specifies whether to return the portrait photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

type RecognizeBrazilCommonOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64 value of the image. supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. supported image size: no more than 7M after encoding. image download time: no more than 3 seconds.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The Url of the image. supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. supported image size: no more than 7M after Base64 encoding. image download time is no more than 3 seconds. urls stored in tencent cloud guarantee higher download speed and stability. it is advisable to store images in tencent cloud. urls not stored in tencent cloud may possibly be impacted in speed and stability.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The Base64 value of the back side of the card. supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. supported image size: no more than 7M after encoding. image download time: not more than 3 seconds. either ImageUrl or ImageBase64 must be provided. if both are provided, only use ImageUrl.
	BackImageBase64 *string `json:"BackImageBase64,omitnil,omitempty" name:"BackImageBase64"`

	// The Url address of the back side of the card. supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. supported image size: no more than 7M after Base64 encoding. image download time is no more than 3 seconds. urls stored in tencent cloud guarantee higher download speed and stability. it is recommended to store images in tencent cloud. speed and stability of non-tencent cloud storage urls may be impacted.
	BackImageUrl *string `json:"BackImageUrl,omitnil,omitempty" name:"BackImageUrl"`

	// Specifies whether to return the portrait photo.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`
}

func (r *RecognizeBrazilCommonOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilCommonOCRRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeBrazilCommonOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeBrazilCommonOCRResponseParams struct {
	// Specifies the type of document in brazil. valid values: 
	// 1. RNE 
	// 2. RNM 
	// 3. IDCard 
	// 4. DrivingLicense.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// The recognized content of the Brazilian document.
	Result *BrazilCardInfo `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeBrazilCommonOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeBrazilCommonOCRResponseParams `json:"Response"`
}

func (r *RecognizeBrazilCommonOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeBrazilCommonOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Version of the driver's license image. 
	// Valid values: 
	// OLD (old version), 
	// NEW (new version). 
	// The default value is OLD.
	LicenceVersion *string `json:"LicenceVersion,omitnil,omitempty" name:"LicenceVersion"`
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

	// Version of the driver's license image. 
	// Valid values: 
	// OLD (old version), 
	// NEW (new version). 
	// The default value is OLD.
	LicenceVersion *string `json:"LicenceVersion,omitnil,omitempty" name:"LicenceVersion"`
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
	delete(f, "LicenceVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeBrazilDriverLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeBrazilDriverLicenseOCRResponseParams struct {
	// Name of the license holder.
	NOME *string `json:"NOME,omitnil,omitempty" name:"NOME"`

	// Driver's license category.
	CatHab *string `json:"CatHab,omitnil,omitempty" name:"CatHab"`

	// Driver's license number (CNH).
	CNHNumber *string `json:"CNHNumber,omitnil,omitempty" name:"CNHNumber"`

	// Validity date (valid until).
	VALIDADE *string `json:"VALIDADE,omitnil,omitempty" name:"VALIDADE"`

	// Qualification information.
	QUALIFICATION *string `json:"QUALIFICATION,omitnil,omitempty" name:"QUALIFICATION"`

	// ID number (Identity document number).
	IDENTIDADE *string `json:"IDENTIDADE,omitnil,omitempty" name:"IDENTIDADE"`

	// CPF
	CPF *string `json:"CPF,omitnil,omitempty" name:"CPF"`

	// Date of birth.
	NASCIMENTO *string `json:"NASCIMENTO,omitnil,omitempty" name:"NASCIMENTO"`

	// Membership
	MEMBERSHIP *string `json:"MEMBERSHIP,omitnil,omitempty" name:"MEMBERSHIP"`

	// Registration number
	REGISTRO *string `json:"REGISTRO,omitnil,omitempty" name:"REGISTRO"`

	// Remarks
	OBSERVATIONS *string `json:"OBSERVATIONS,omitnil,omitempty" name:"OBSERVATIONS"`

	// Date of issue.
	IssueDate *string `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`

	// Place of issue.
	LOCAL *string `json:"LOCAL,omitnil,omitempty" name:"LOCAL"`

	// Registration number on the back of the card.
	BackNumber *string `json:"BackNumber,omitnil,omitempty" name:"BackNumber"`

	// This field is deprecated and will always return "1". Usage is not recommended.
	//
	// Deprecated: AdvancedInfo is deprecated.
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

	// Date of birth
	DataNascimento *string `json:"DataNascimento,omitnil,omitempty" name:"DataNascimento"`

	// Issuing agency
	IssuingAgency *string `json:"IssuingAgency,omitnil,omitempty" name:"IssuingAgency"`

	// Blood type
	Fatorrh *string `json:"Fatorrh,omitnil,omitempty" name:"Fatorrh"`

	// Birth place
	NaturalIDade *string `json:"NaturalIDade,omitnil,omitempty" name:"NaturalIDade"`

	// Additional information 
	Observations *string `json:"Observations,omitnil,omitempty" name:"Observations"`

	// CPF
	CPF *string `json:"CPF,omitnil,omitempty" name:"CPF"`

	// DNI
	DNI *string `json:"DNI,omitnil,omitempty" name:"DNI"`

	// General registry (Registro Geral)
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
	// The RNE (Registro Nacional de Estrangeiros) number.
	RNE *string `json:"RNE,omitnil,omitempty" name:"RNE"`

	// The classification of the RNE document.
	CLASSIFICATION *string `json:"CLASSIFICATION,omitnil,omitempty" name:"CLASSIFICATION"`

	// The validity period (expiry date) of the RNE document.
	VALIDADE *string `json:"VALIDADE,omitnil,omitempty" name:"VALIDADE"`

	// The full name.
	NOME *string `json:"NOME,omitnil,omitempty" name:"NOME"`

	// Family information (parents' names).
	Membership *string `json:"Membership,omitnil,omitempty" name:"Membership"`

	// Nationality
	NACIONALIDADE *string `json:"NACIONALIDADE,omitnil,omitempty" name:"NACIONALIDADE"`

	// Place of Birth
	NATURALIDADE *string `json:"NATURALIDADE,omitnil,omitempty" name:"NATURALIDADE"`

	// The issuing agency.
	IssuingAgency *string `json:"IssuingAgency,omitnil,omitempty" name:"IssuingAgency"`

	// Date of birth.
	DateOfBirth *string `json:"DateOfBirth,omitnil,omitempty" name:"DateOfBirth"`

	// Gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// The date of entry into Brazil.
	EntryDate *string `json:"EntryDate,omitnil,omitempty" name:"EntryDate"`

	// The VIA (document version/sequence number).
	VIA *string `json:"VIA,omitnil,omitempty" name:"VIA"`

	// The issue date.
	DispatchDate *string `json:"DispatchDate,omitnil,omitempty" name:"DispatchDate"`

	// The machine readable zone (MRZ) code.
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

	// Residence category
	CLASSIFICATION *string `json:"CLASSIFICATION,omitnil,omitempty" name:"CLASSIFICATION"`

	// Residence validity term
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
type RecognizeDetectCardCoordsRequestParams struct {
	// URL of the image. Supported image formats: PNG, JPG, JPEG. GIF is not supported. The downloaded image must be no larger than 7 MB after Base64 encoding. Image download must complete within 3 seconds. Images stored in Tencent Cloud offer higher download speed and stability. We recommend storing images in Tencent Cloud. The speed and stability of URLs from non-Tencent Cloud storage may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Base64-encoded image data. Supported image formats: PNG, JPG, JPEG. GIF is not supported. The image must be no larger than 7 MB after Base64 encoding. Image download must complete within 3 seconds. You must provide either ImageUrl or ImageBase64. If both are provided, ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`
}

type RecognizeDetectCardCoordsRequest struct {
	*tchttp.BaseRequest
	
	// URL of the image. Supported image formats: PNG, JPG, JPEG. GIF is not supported. The downloaded image must be no larger than 7 MB after Base64 encoding. Image download must complete within 3 seconds. Images stored in Tencent Cloud offer higher download speed and stability. We recommend storing images in Tencent Cloud. The speed and stability of URLs from non-Tencent Cloud storage may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Base64-encoded image data. Supported image formats: PNG, JPG, JPEG. GIF is not supported. The image must be no larger than 7 MB after Base64 encoding. Image download must complete within 3 seconds. You must provide either ImageUrl or ImageBase64. If both are provided, ImageUrl will be used.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`
}

func (r *RecognizeDetectCardCoordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeDetectCardCoordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeDetectCardCoordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeDetectCardCoordsResponseParams struct {
	// Coordinate information of the detected four corners of the card.
	ItemList []*CoordsItem `json:"ItemList,omitnil,omitempty" name:"ItemList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeDetectCardCoordsResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeDetectCardCoordsResponseParams `json:"Response"`
}

func (r *RecognizeDetectCardCoordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeDetectCardCoordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeIndonesiaIDCardOCRRequestParams struct {
	// The Base64 value of the image. Supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. Supported image size: the downloaded image after Base64 encoding is no more than 7M. Image download time is not more than 3 seconds. Either ImageUrl or ImageBase64 must be provided. If both are provided, only use ImageUrl.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The Url address of the image. 
	// Supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. 
	// Supported image size: the downloaded image after Base64 encoding is no more than 7M. Image download time is no more than 3 seconds. 
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the portrait photo.
	// If selected true, image restrictions are: Image size after encoding must not exceed 5M, jpg format long side pixel cannot exceed 4000, other formats image long edge pixel maximum of 2000. Short side pixel of all format images not less than 64.
	// Support PNG, jpg, JPEG, BMP, no support for GIF images.
	// If portrait matting fails, return an empty string.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// Scene parameter, default value is V1
	// Available values:
	// V1
	// V2
	//
	// Deprecated: Scene is deprecated.
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`
}

type RecognizeIndonesiaIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64 value of the image. Supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. Supported image size: the downloaded image after Base64 encoding is no more than 7M. Image download time is not more than 3 seconds. Either ImageUrl or ImageBase64 must be provided. If both are provided, only use ImageUrl.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The Url address of the image. 
	// Supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. 
	// Supported image size: the downloaded image after Base64 encoding is no more than 7M. Image download time is no more than 3 seconds. 
	// We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// For a non-Tencent Cloud URL, the download speed and stability may be affected.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the portrait photo.
	// If selected true, image restrictions are: Image size after encoding must not exceed 5M, jpg format long side pixel cannot exceed 4000, other formats image long edge pixel maximum of 2000. Short side pixel of all format images not less than 64.
	// Support PNG, jpg, JPEG, BMP, no support for GIF images.
	// If portrait matting fails, return an empty string.
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil,omitempty" name:"ReturnHeadImage"`

	// Scene parameter, default value is V1
	// Available values:
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

	// The neighborhood/community unit (RT/RW).
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

	// The portraitImage.
	Photo *string `json:"Photo,omitnil,omitempty" name:"Photo"`

	// The province.
	Provinsi *string `json:"Provinsi,omitnil,omitempty" name:"Provinsi"`

	// The city.
	Kota *string `json:"Kota,omitnil,omitempty" name:"Kota"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate
	// -9102 Alarm for photocopied certificate
	// -9103 Alarm for photographed certificate
	// -9104 Alarm for tamper certificate
	// -9107 Alarm for reflective certificate
	// -9108 Alarm for blurry image
	// -8101 Alarm for document information format verification
	// -8102 Alarm for document information consistency verification
	// -9109 This capability is not enabled. Please contact customer support to activate the alert service
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
	// Last name in Chinese
	CnLastName *string `json:"CnLastName,omitnil,omitempty" name:"CnLastName"`

	// Last name in English
	EnLastName *string `json:"EnLastName,omitnil,omitempty" name:"EnLastName"`

	// Telecode of the last name in Chinese
	LastNameCode *string `json:"LastNameCode,omitnil,omitempty" name:"LastNameCode"`

	// First name in Chinese
	CnFirstName *string `json:"CnFirstName,omitnil,omitempty" name:"CnFirstName"`

	// First name in English
	EnFirstName *string `json:"EnFirstName,omitnil,omitempty" name:"EnFirstName"`

	// Telecode of the first name in Chinese
	FirstNameCode *string `json:"FirstNameCode,omitnil,omitempty" name:"FirstNameCode"`

	// Identity card number
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Date of birth (DD-MM-YYYY)
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Date of first issue (DD-MM-YYYY)
	FirstIssueDate *string `json:"FirstIssueDate,omitnil,omitempty" name:"FirstIssueDate"`

	// Date of issue (DD-MM-YYYY)
	CurrentIssueDate *string `json:"CurrentIssueDate,omitnil,omitempty" name:"CurrentIssueDate"`

	// Validity period (DD-MM-YYYY)
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// Document symbol
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// Height (unit: meters)
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// Processed image (Base64)
	RetImage *string `json:"RetImage,omitnil,omitempty" name:"RetImage"`

	// This field is deprecated and will always return null. Usage is not recommended.
	//
	// Deprecated: Angle is deprecated.
	Angle *string `json:"Angle,omitnil,omitempty" name:"Angle"`

	// Resident type. 
	// Valid values: Permanent Resident Identity Card, Non-permanent Resident Identity Card.
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

	// Gender (portrait side)
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Ethnicity (portrait side)
	Nation *string `json:"Nation,omitnil,omitempty" name:"Nation"`

	// Date of birth (portrait side)
	Birth *string `json:"Birth,omitnil,omitempty" name:"Birth"`

	// Address(portrait side)
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// ID number (portrait side)
	IdNum *string `json:"IdNum,omitnil,omitempty" name:"IdNum"`

	// Issuing authority (national emblem side)
	Authority *string `json:"Authority,omitnil,omitempty" name:"Authority"`

	// Validity period (national emblem side)
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// Card Warning Information
	// 
	// -9101 Alarm for covered certificate
	// -9102 Alarm for photocopied certificate
	// -9103 Alarm for photographed certificate
	// -9104 Alarm for tamper certificate
	// -9107 Alarm for reflective certificate
	// -9108 Alarm for blurry certificate 
	// -9109 This capability is not enabled. Please contact customer support to activate the alert service.
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
	// The full name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Gender.
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Address
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Vote PIN Code
	VotePIN *string `json:"VotePIN,omitnil,omitempty" name:"VotePIN"`

	// Unique Population Registry Code
	CURP *string `json:"CURP,omitnil,omitempty" name:"CURP"`

	// Date of birth.
	Birth *string `json:"Birth,omitnil,omitempty" name:"Birth"`

	// Section Number
	SECCION *string `json:"SECCION,omitnil,omitempty" name:"SECCION"`

	// IssueDate
	IssueDate *string `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`

	// The validity period (expiration date).
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

	// Full name of the license holder.
	Name *TextDetectionResult `json:"Name,omitnil,omitempty" name:"Name"`

	// Last name / surname of the license holder.
	LastName *TextDetectionResult `json:"LastName,omitnil,omitempty" name:"LastName"`

	// First name of the license holder.
	FirstName *TextDetectionResult `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// Middle name of the license holder.
	MiddleName *TextDetectionResult `json:"MiddleName,omitnil,omitempty" name:"MiddleName"`

	// Nationality of the license holder.
	Nationality *TextDetectionResult `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// Gender
	Sex *TextDetectionResult `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Address of the license holder.
	Address *TextDetectionResult `json:"Address,omitnil,omitempty" name:"Address"`

	// Driver's license number.
	LicenseNo *TextDetectionResult `json:"LicenseNo,omitnil,omitempty" name:"LicenseNo"`

	// Expiration date of the driver's license.
	ExpiresDate *TextDetectionResult `json:"ExpiresDate,omitnil,omitempty" name:"ExpiresDate"`

	// Code of the issuing agency.
	AgencyCode *TextDetectionResult `json:"AgencyCode,omitnil,omitempty" name:"AgencyCode"`

	// Date of birth of the license holder.
	Birthday *TextDetectionResult `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The license number (SSSID number).
	LicenseNumber *TextDetectionResult `json:"LicenseNumber,omitnil,omitempty" name:"LicenseNumber"`

	// The full name.
	FullName *TextDetectionResult `json:"FullName,omitnil,omitempty" name:"FullName"`

	// The date of birth.
	Birthday *TextDetectionResult `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The full name.
	FullName *TextDetectionResult `json:"FullName,omitnil,omitempty" name:"FullName"`

	// The address.
	Address *TextDetectionResult `json:"Address,omitnil,omitempty" name:"Address"`

	// The birth date.
	Birthday *TextDetectionResult `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The issue date.
	IssueDate *TextDetectionResult `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The last name.
	LastName *TextDetectionResult `json:"LastName,omitnil,omitempty" name:"LastName"`

	// The first name.
	FirstName *TextDetectionResult `json:"FirstName,omitnil,omitempty" name:"FirstName"`


	MiddleName *TextDetectionResult `json:"MiddleName,omitnil,omitempty" name:"MiddleName"`

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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Gender
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// Country of birth
	CountryOfBirth *string `json:"CountryOfBirth,omitnil,omitempty" name:"CountryOfBirth"`

	// Date of birth
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// Address (back side)
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// ID number
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Nationality(back side)
	Race *string `json:"Race,omitnil,omitempty" name:"Race"`

	// NRIC code(back side)
	NRICCode *string `json:"NRICCode,omitnil,omitempty" name:"NRICCode"`

	// Postal code (back side)
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

	// First name in English
	EnFirstName *string `json:"EnFirstName,omitnil,omitempty" name:"EnFirstName"`

	// Last name in English
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

	// Laser ID on the back of the card.
	LaserID *string `json:"LaserID,omitnil,omitempty" name:"LaserID"`

	// Identity photo
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

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

	// This field is deprecated and will always return "1". Usage is not recommended.
	//
	// Deprecated: AdvancedInfo is deprecated.
	AdvancedInfo *string `json:"AdvancedInfo,omitnil,omitempty" name:"AdvancedInfo"`

	// The number of cards detected in the input image provided via ImageBase64 parameter.(Currently supported only in ap-bangkok region)
	CardCount *int64 `json:"CardCount,omitnil,omitempty" name:"CardCount"`

	// The card information field complete or not
	// true: complete; false: incomplete
	IsComplete *bool `json:"IsComplete,omitnil,omitempty" name:"IsComplete"`

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

	// This field is deprecated and will always return "1". Usage is not recommended.
	//
	// Deprecated: AdvancedInfo is deprecated.
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

	// This field is deprecated and will always return an empty array. Usage is not recommended.
	//
	// Deprecated: Polygon is deprecated.
	Polygon []*Coord `json:"Polygon,omitnil,omitempty" name:"Polygon"`
}

type Value struct {
	// The value of the recognized field.
	AutoContent *string `json:"AutoContent,omitnil,omitempty" name:"AutoContent"`

	// The coordinates of the four vertices.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Coord *Polygon `json:"Coord,omitnil,omitempty" name:"Coord"`
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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