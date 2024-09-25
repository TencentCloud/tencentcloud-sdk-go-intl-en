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

type FuseFaceReviewDetail struct {

	Field *string `json:"Field,omitnil,omitempty" name:"Field"`


	Label *string `json:"Label,omitnil,omitempty" name:"Label"`


	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`


	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`
}

type FuseFaceReviewResult struct {

	Category *string `json:"Category,omitnil,omitempty" name:"Category"`


	Code *string `json:"Code,omitnil,omitempty" name:"Code"`


	CodeDescription *string `json:"CodeDescription,omitnil,omitempty" name:"CodeDescription"`


	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`


	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`


	DetailSet []*FuseFaceReviewDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`
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

// Predefined struct for user
type QueryVideoFaceFusionJobRequestParams struct {
	// Job ID of the video face fusion task
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryVideoFaceFusionJobRequest struct {
	*tchttp.BaseRequest
	
	// Job ID of the video face fusion task
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryVideoFaceFusionJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVideoFaceFusionJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVideoFaceFusionJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVideoFaceFusionJobResponseParams struct {
	// Current task status: queuing, processing, processing failed, or processing completed
	JobStatus *string `json:"JobStatus,omitnil,omitempty" name:"JobStatus"`

	// Video face fusion result
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoFaceFusionOutput *VideoFaceFusionOutput `json:"VideoFaceFusionOutput,omitnil,omitempty" name:"VideoFaceFusionOutput"`

	// Task status code. 1: queuing; 3: processing; 5: processing failed; 7: processing completed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobStatusCode *int64 `json:"JobStatusCode,omitnil,omitempty" name:"JobStatusCode"`

	// Task failure error code
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobErrorCode *string `json:"JobErrorCode,omitnil,omitempty" name:"JobErrorCode"`

	// Task failure error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryVideoFaceFusionJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryVideoFaceFusionJobResponseParams `json:"Response"`
}

func (r *QueryVideoFaceFusionJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVideoFaceFusionJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoFaceFusionJobRequestParams struct {
	// Activity ID. Check it in the video face fusion console.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Material ID. Check it in the video face fusion console.
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// Face position information on the user face image and material template image. Only one entry is allowed.
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 0: inappropriate content recognition not required; 1: inappropriate content recognition required. Default value: 0.
	// Note: Once the inappropriate content recognition service is enabled, you need to decide whether to adjust your business logic based on the returned results. For example, you need to replace the image if the system informs you that the image does not meet the requirements.
	// **<font color=#1E90FF>Note: This field will be deprecated later due to business adjustments. It is not recommended for use.</font>**
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitnil,omitempty" name:"CelebrityIdentify"`

	// Video watermark logo parameter
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// COS pre-signed URL (PUT method). If this parameter is specified, the video after fusion will be uploaded to this URL.
	// **<font color=#1E90FF>Note: If upload to this URL fails, the video will be uploaded to the default address of Tencent Cloud.</font>**
	UserDesignatedUrl *string `json:"UserDesignatedUrl,omitnil,omitempty" name:"UserDesignatedUrl"`

	// User IP address
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// Video metadata field
	MetaData []*MetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`
}

type SubmitVideoFaceFusionJobRequest struct {
	*tchttp.BaseRequest
	
	// Activity ID. Check it in the video face fusion console.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Material ID. Check it in the video face fusion console.
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// Face position information on the user face image and material template image. Only one entry is allowed.
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 0: inappropriate content recognition not required; 1: inappropriate content recognition required. Default value: 0.
	// Note: Once the inappropriate content recognition service is enabled, you need to decide whether to adjust your business logic based on the returned results. For example, you need to replace the image if the system informs you that the image does not meet the requirements.
	// **<font color=#1E90FF>Note: This field will be deprecated later due to business adjustments. It is not recommended for use.</font>**
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitnil,omitempty" name:"CelebrityIdentify"`

	// Video watermark logo parameter
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// COS pre-signed URL (PUT method). If this parameter is specified, the video after fusion will be uploaded to this URL.
	// **<font color=#1E90FF>Note: If upload to this URL fails, the video will be uploaded to the default address of Tencent Cloud.</font>**
	UserDesignatedUrl *string `json:"UserDesignatedUrl,omitnil,omitempty" name:"UserDesignatedUrl"`

	// User IP address
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// Video metadata field
	MetaData []*MetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`
}

func (r *SubmitVideoFaceFusionJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoFaceFusionJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "MergeInfos")
	delete(f, "CelebrityIdentify")
	delete(f, "LogoParam")
	delete(f, "UserDesignatedUrl")
	delete(f, "UserIp")
	delete(f, "MetaData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoFaceFusionJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoFaceFusionJobResponseParams struct {
	// Job ID of the video face fusion task
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Estimated processing time of the video face fusion task, in seconds
	EstimatedProcessTime *float64 `json:"EstimatedProcessTime,omitnil,omitempty" name:"EstimatedProcessTime"`

	// Estimated processing time of the video face fusion task, in seconds
	JobQueueLength *int64 `json:"JobQueueLength,omitnil,omitempty" name:"JobQueueLength"`

	// Inappropriate content recognition result. The element order of this array is the same as that of mergeinfo in the request, with a one-to-one relationship.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReviewResultSet []*FuseFaceReviewResult `json:"ReviewResultSet,omitnil,omitempty" name:"ReviewResultSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitVideoFaceFusionJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitVideoFaceFusionJobResponseParams `json:"Response"`
}

func (r *SubmitVideoFaceFusionJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoFaceFusionJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VideoFaceFusionOutput struct {
	// URL of the video output after video face fusion
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// MD5 value of the video output after video face fusion, which is used for verification
	VideoMD5 *string `json:"VideoMD5,omitnil,omitempty" name:"VideoMD5"`

	// Video width
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Video height
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Frames per second
	FPS *int64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// Video duration, in seconds
	DurationInSec *float64 `json:"DurationInSec,omitnil,omitempty" name:"DurationInSec"`

	// Number of frames
	Frame *int64 `json:"Frame,omitnil,omitempty" name:"Frame"`
}