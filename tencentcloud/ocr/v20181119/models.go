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

// Predefined struct for user
type PermitOCRRequestParams struct {
	// The Base64 value of the image. Supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. Supported image size: the downloaded image after Base64 encoding is no more than 7M. Image download time is not more than 3 seconds. Either ImageUrl or ImageBase64 must be provided. If both are provided, only use ImageUrl.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. Supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. Supported image size: no more than 7M after Base64 encoding. Image download time: no more than 3 seconds. We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the avatar image. Default is false.
	CropPortrait *bool `json:"CropPortrait,omitnil,omitempty" name:"CropPortrait"`
}

type PermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// The Base64 value of the image. Supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. Supported image size: the downloaded image after Base64 encoding is no more than 7M. Image download time is not more than 3 seconds. Either ImageUrl or ImageBase64 must be provided. If both are provided, only use ImageUrl.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// The URL of the image. Supported image formats: PNG, JPG, JPEG. GIF format is not currently supported. Supported image size: no more than 7M after Base64 encoding. Image download time: no more than 3 seconds. We recommend that you store the image in Tencent Cloud for higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Whether to return the avatar image. Default is false.
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

	// Issuing place
	IssueAddress *string `json:"IssueAddress,omitnil,omitempty" name:"IssueAddress"`

	// Date of birth
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// base64 of the avatar image
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// Return type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

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