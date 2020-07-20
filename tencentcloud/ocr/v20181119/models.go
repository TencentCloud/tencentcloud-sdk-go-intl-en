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

	// Base64-encoded value of image.
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 7 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of the image.
	// Supported image formats: PNG, JPG, JPEG. GIF is currently not supported.
	// Supported image size: the downloaded image cannot exceed 7 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
	// It is recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
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

		// Expiration date
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

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

type GeneralBasicOCRRequest struct {
	*tchttp.BaseRequest

	// Base64-encoded value of image.
	// The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// Either `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// URL address of image.
	// The image cannot exceed 7 MB in size after being Base64-encoded. A resolution above 600x800 is recommended. PNG, JPG, JPEG, and BMP formats are supported.
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Reserved field.
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// Language to be recognized.
	// The language can be automatically recognized or manually specified. Chinese-English mix (`zh`) is selected by default. Mixed characters in English and each supported language can be recognized together.
	// Valid values:
	// zh\auto\jap\kor\nspa\fre\ger\por\nvie\may\rus\ita\nhol\swe\fin\dan\nnor\hun\tha\lat
	// Value meanings:
	// Chinese-English mix, automatic recognition, Japanese, Korean,
	// Spanish, French, German, Portuguese,
	// Vietnamese, Malay, Russian, Italian,
	// Dutch, Swedish, Finnish, Danish,
	// Norwegian, Hungarian, Thai, Latin.
	LanguageType *string `json:"LanguageType,omitempty" name:"LanguageType"`
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

		// Image rotation angle in degrees. 0° indicates horizontal text, a positive value indicates clockwise rotation, and a negative value indicates anticlockwise rotation. For more information, please see <a href="https://cloud.tencent.com/document/product/866/45139">How to Correct Tilted Text</a>
		Angel *float64 `json:"Angel,omitempty" name:"Angel"`

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

type ItemCoord struct {

	// X-coordinate of top-left point
	X *int64 `json:"X,omitempty" name:"X"`

	// Y-coordinate of top-left point
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

	// Base64-encoded value of the image. The image cannot exceed 7 MB in size after being Base64-encoded. A resolution of 500x800 or above is recommended. Supported formats include PNG, JPG, JPEG, and BMP. It is recommended that the card part occupies more than 2/3 of the image.
	// Either the `ImageUrl` or `ImageBase64` of the image must be provided; if both are provided, only `ImageUrl` will be used.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// Whether to return an image
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

type TextDetection struct {

	// Recognized text line content
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// Confidence. Value range: 0–100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// Text line coordinates, which are represented as 4 vertex coordinates
	// Note: this field may return null, indicating that no valid values can be obtained.
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// Extended field.
	// The paragraph information `Parag` returned by the `GeneralBasicOcr` API contains `ParagNo`.
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// Pixel coordinates of the text line in the image after rotation correction, which is in the format of `(X-coordinate of top-left point, Y-coordinate of top-left point, width, height)`
	ItemPolygon *ItemCoord `json:"ItemPolygon,omitempty" name:"ItemPolygon"`
}
