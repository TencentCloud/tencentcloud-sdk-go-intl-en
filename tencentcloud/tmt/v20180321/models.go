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

package v20180321

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type BoundingBox struct {
	// <p>x-coordinate of the top-left corner</p>
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// <p>y-coordinate of the top-left corner</p>
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// <p>Width.</p><p>Unit: px.</p>
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// <p>High.</p><p>Unit: px.</p>
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type Coord struct {
	// X coordinate
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Y-axis coordinate
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`
}

// Predefined struct for user
type ImageTranslateLLMRequestParams struct {
	// <p>Base64 string of the image data, no more than 9M after Base64 encoding. A resolution of 600*800 or higher is recommended. PNG, JPG, and JPEG formats are supported.</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Target language, supported languages:</p><ul><li>Chinese: zh</li><li>Traditional (Taiwan): zh-TW</li><li>Traditional (Hong Kong (China)): zh-HK</li><li>English: en</li><li>Japanese: ja</li><li>Korean: ko</li><li>Thai: th</li><li>Vietnamese: vi</li><li>Russian: ru</li><li>German: de</li><li>French: fr</li><li>Arabic: ar</li><li>Spanish: es</li><li>Italian: it</li><li>Indonesian: id</li><li>Malay language: ms</li><li>Portuguese: pt</li><li>Turkish: tr<br>-</li></ul>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>Enter image Url. When using a Url, the Data parameter requires the input of "". Image restrictions: less than 10MB, resolution recommendation 600*800 or higher, format support jpg, jpeg, png.</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Invocation method.</p><p>Enumeration value:</p><ul><li>0: End-to-end image translation large model pro version</li><li>1: End-to-end image translation large model lite version</li></ul><p>Default value: 0</p>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type ImageTranslateLLMRequest struct {
	*tchttp.BaseRequest
	
	// <p>Base64 string of the image data, no more than 9M after Base64 encoding. A resolution of 600*800 or higher is recommended. PNG, JPG, and JPEG formats are supported.</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Target language, supported languages:</p><ul><li>Chinese: zh</li><li>Traditional (Taiwan): zh-TW</li><li>Traditional (Hong Kong (China)): zh-HK</li><li>English: en</li><li>Japanese: ja</li><li>Korean: ko</li><li>Thai: th</li><li>Vietnamese: vi</li><li>Russian: ru</li><li>German: de</li><li>French: fr</li><li>Arabic: ar</li><li>Spanish: es</li><li>Italian: it</li><li>Indonesian: id</li><li>Malay language: ms</li><li>Portuguese: pt</li><li>Turkish: tr<br>-</li></ul>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>Enter image Url. When using a Url, the Data parameter requires the input of "". Image restrictions: less than 10MB, resolution recommendation 600*800 or higher, format support jpg, jpeg, png.</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Invocation method.</p><p>Enumeration value:</p><ul><li>0: End-to-end image translation large model pro version</li><li>1: End-to-end image translation large model lite version</li></ul><p>Default value: 0</p>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

func (r *ImageTranslateLLMRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageTranslateLLMRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Target")
	delete(f, "Url")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageTranslateLLMRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageTranslateLLMResponseParams struct {
	// <p>Base64 string of the image data. The output format is JPG.</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Primary source language.</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Target translation language.</p>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>All original text in the image.</p>
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// <p>All translations in the image.</p>
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// <p>Image angle counterclockwise, value range 0-359</p>
	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// <p>Translation detailed information</p>
	TransDetails []*TransDetail `json:"TransDetails,omitnil,omitempty" name:"TransDetails"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageTranslateLLMResponse struct {
	*tchttp.BaseResponse
	Response *ImageTranslateLLMResponseParams `json:"Response"`
}

func (r *ImageTranslateLLMResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageTranslateLLMResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RotateParagraphRect struct {
	// Paragraph text coordinates
	Coord []*Coord `json:"Coord,omitnil,omitempty" name:"Coord"`

	// Rotation angle
	TiltAngle *float64 `json:"TiltAngle,omitnil,omitempty" name:"TiltAngle"`

	// Whether the paragraph text information is valid
	Valid *bool `json:"Valid,omitnil,omitempty" name:"Valid"`
}

type TransDetail struct {
	// <p>Original text of the current row</p>
	SourceLineText *string `json:"SourceLineText,omitnil,omitempty" name:"SourceLineText"`

	// <p>Translation of the current row</p>
	TargetLineText *string `json:"TargetLineText,omitnil,omitempty" name:"TargetLineText"`

	// <p>Paragraph text box location</p>
	BoundingBox *BoundingBox `json:"BoundingBox,omitnil,omitempty" name:"BoundingBox"`

	// <p>Row count</p>
	LinesCount *int64 `json:"LinesCount,omitnil,omitempty" name:"LinesCount"`

	// <p>Line height.</p><p>Unit: px.</p>
	LineHeight *int64 `json:"LineHeight,omitnil,omitempty" name:"LineHeight"`

	// <p>The spam_code field is 0 in a normal paragraph; if the spam_code field exists and its value is above 0 (1: hit garbage check; 2: hit security policy; 3: another.), then the security check hit is filtered.</p>
	SpamCode *int64 `json:"SpamCode,omitnil,omitempty" name:"SpamCode"`

	// <p>Rotation information of paragraph text. Coordinates are valid only when valid is true.</p>
	RotateParagraphRect *RotateParagraphRect `json:"RotateParagraphRect,omitnil,omitempty" name:"RotateParagraphRect"`
}