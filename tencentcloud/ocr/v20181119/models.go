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
	// -9103 Alarm for photographed document
	// -9102 Alarm for photocopied document
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

func (r *MLIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MLIDCardOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MLIDPassportOCRRequest struct {
	*tchttp.BaseRequest

	// Base64-encoded value of an image.
	// Supported image formats: PNG, JPG, JPEG. GIF is not supported at present.
	// Supported image size: the downloaded image cannot exceed 3 MB in size after being Base64-encoded. The download time of the image cannot exceed 3 seconds.
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
	// -9103 Alarm for photographed document
	// -9102 Alarm for photocopied document
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
