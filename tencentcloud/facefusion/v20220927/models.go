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

package v20220927

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type FaceRect struct {
	// Top-left X-axis coordinate of the face box
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Top-left Y-axis coordinate of the face box
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// Face box width
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Face box height
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

// Predefined struct for user
type FuseFaceRequestParams struct {
	// Activity ID. Check the ID in the <a href="https://console.cloud.tencent.com/facefusion" target="_blank"> Face Fusion console</a>.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Material ID. Check the ID in the <a href="https://console.cloud.tencent.com/facefusion" target="_blank"> Face Fusion console</a>.
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// Image return method (url or base64). You cannot use both methods at the same time. The URL is valid for 7 days.
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// Face position information on the user face image and material template image. No more than 6 entries.
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// Switch indicating whether to add a synthesis logo to the fusion result image. Default value: 1.
	// 1: add logo
	// 0: do not add logo
	// Other values: add logo
	// It is recommended to use an obvious logo to indicate that the result image uses face fusion technology and is synthesized by AI.
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// Logo content settings
	// By default, the text "Synthesized by AI" is added to the bottom right corner of the fusion result image. You can also use other texts.
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// Fusion parameter.
	FuseParam *FuseParam `json:"FuseParam,omitnil,omitempty" name:"FuseParam"`
}

type FuseFaceRequest struct {
	*tchttp.BaseRequest
	
	// Activity ID. Check the ID in the <a href="https://console.cloud.tencent.com/facefusion" target="_blank"> Face Fusion console</a>.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Material ID. Check the ID in the <a href="https://console.cloud.tencent.com/facefusion" target="_blank"> Face Fusion console</a>.
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// Image return method (url or base64). You cannot use both methods at the same time. The URL is valid for 7 days.
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// Face position information on the user face image and material template image. No more than 6 entries.
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// Switch indicating whether to add a synthesis logo to the fusion result image. Default value: 1.
	// 1: add logo
	// 0: do not add logo
	// Other values: add logo
	// It is recommended to use an obvious logo to indicate that the result image uses face fusion technology and is synthesized by AI.
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// Logo content settings
	// By default, the text "Synthesized by AI" is added to the bottom right corner of the fusion result image. You can also use other texts.
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// Fusion parameter.
	FuseParam *FuseParam `json:"FuseParam,omitnil,omitempty" name:"FuseParam"`
}

func (r *FuseFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "RspImgType")
	delete(f, "MergeInfos")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "FuseParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FuseFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FuseFaceResponseParams struct {
	// When RspImgType is set to url, return the URL (valid for 7 days). When RspImgType is set to base64, return the Base64 code.
	FusedImage *string `json:"FusedImage,omitnil,omitempty" name:"FusedImage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FuseFaceResponse struct {
	*tchttp.BaseResponse
	Response *FuseFaceResponseParams `json:"Response"`
}

func (r *FuseFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FuseParam struct {
	// Image encoding parameter
	ImageCodecParam *ImageCodecParam `json:"ImageCodecParam,omitnil,omitempty" name:"ImageCodecParam"`
}

type ImageCodecParam struct {
	// Metadata. The number of metadata entries cannot exceed 1.
	MetaData []*MetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`
}

type LogoParam struct {
	// Coordinates of the logo image in the fusion result image. The logo image will be stretched according to the coordinates.
	LogoRect *FaceRect `json:"LogoRect,omitnil,omitempty" name:"LogoRect"`

	// Logo image URL
	// ●Either the Base64 code or URL must be provided. If both are provided, URL prevails.
	// ●Supported image format: JPG or PNG
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// Logo image Base64 code
	// ●Either the Base64 code or URL must be provided. If both are provided, URL prevails.
	// ●Supported image format: JPG or PNG
	LogoImage *string `json:"LogoImage,omitnil,omitempty" name:"LogoImage"`
}

type MergeInfo struct {
	// Enter the image Base64 code.
	// ●Either the Base64 code or URL must be provided. If both are provided, URL prevails.
	// ●Material image limitation: face size in the image greater than 34×34 pixels; image size greater than 64×64 pixels. (After encoding, the image size may increase by about 30%. It is recommended to control the image size reasonably.)
	// ●Supported image format: JPG or PNG
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Enter the image URL.
	// ●Either the Base64 code or URL must be provided. If both are provided, URL prevails.
	// ●Material image limitation: face size in the image greater than 34×34 pixels; image size greater than 64×64 pixels. (After encoding, the image size may increase by about 30%. It is recommended to control the image size reasonably.)
	// ●Supported image format: JPG or PNG
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Face position information (face box) on the uploaded image
	// Width and height are no less than 30.
	InputImageFaceRect *FaceRect `json:"InputImageFaceRect,omitnil,omitempty" name:"InputImageFaceRect"`

	// Material face ID. If this parameter is left blank, the largest face is used by default.
	TemplateFaceID *string `json:"TemplateFaceID,omitnil,omitempty" name:"TemplateFaceID"`

	// Face position information (face box) on the template. If this parameter is left blank, the largest face is used by default. This parameter applies to scenes where custom material templates are used for fusion.
	// Width and height are no less than 30.
	TemplateFaceRect *FaceRect `json:"TemplateFaceRect,omitnil,omitempty" name:"TemplateFaceRect"`
}

type MetaData struct {
	// Metadata key
	MetaKey *string `json:"MetaKey,omitnil,omitempty" name:"MetaKey"`

	// Metadata value
	MetaValue *string `json:"MetaValue,omitnil,omitempty" name:"MetaValue"`
}