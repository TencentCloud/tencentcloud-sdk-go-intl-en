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

package v20200303

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AnalyzeFaceRequestParams struct {
	// Detection mode. 0: detect all faces that appear; 1: detect the largest face. Default value: 0. The facial feature localization information (facial keypoints) of up to 10 faces can be returned.
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Algorithm model version used by the Face Recognition service.
	// 
	// Currently, `2.0` and `3.0` are supported.
	// 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API, `2.0` will be used by default.
	// 
	// The parameter can be set only to `3.0` for accounts that purchase the service after November 26, 2020.
	// 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

type AnalyzeFaceRequest struct {
	*tchttp.BaseRequest
	
	// Detection mode. 0: detect all faces that appear; 1: detect the largest face. Default value: 0. The facial feature localization information (facial keypoints) of up to 10 faces can be returned.
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Algorithm model version used by the Face Recognition service.
	// 
	// Currently, `2.0` and `3.0` are supported.
	// 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API, `2.0` will be used by default.
	// 
	// The parameter can be set only to `3.0` for accounts that purchase the service after November 26, 2020.
	// 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

func (r *AnalyzeFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FaceModelVersion")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AnalyzeFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AnalyzeFaceResponseParams struct {
	// Width of requested image.
	ImageWidth *uint64 `json:"ImageWidth,omitnil,omitempty" name:"ImageWidth"`

	// Height of requested image.
	ImageHeight *uint64 `json:"ImageHeight,omitnil,omitempty" name:"ImageHeight"`

	// Specific information of facial feature localization (facial keypoints).
	FaceShapeSet []*FaceShape `json:"FaceShapeSet,omitnil,omitempty" name:"FaceShapeSet"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AnalyzeFaceResponse struct {
	*tchttp.BaseResponse
	Response *AnalyzeFaceResponseParams `json:"Response"`
}

func (r *AnalyzeFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeItem struct {
	// Attribute value
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Probability of recognizing `Type`, which indicates the probability of correct recognition. Value range: [0,1].
	Probability *float64 `json:"Probability,omitnil,omitempty" name:"Probability"`
}

type Candidate struct {
	// Person ID.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Face ID.
	// -Valid only in the API response of SearchFaces/SearchFacesReturnsByGroup.
	// -Person search class interfaces use the fusion feature method for retrieval. This field is meaningless.
	FaceId *string `json:"FaceId,omitnil,omitempty" name:"FaceId"`

	// Match score of the candidate. 
	// -In a face base library of 10,000 faces, a false acceptance rate of 1% corresponds to a score of 70 points, 0.1% corresponds to 80 points, and 0.01% corresponds to 90 points.
	// -For a face base library of 100,000, a false acceptance rate of 1% corresponds to a score of 80 points, 0.1% corresponds to 90 points, and 0.01% corresponds to 100 points.
	// -With a face base library of 300,000, a false acceptance rate of 1% corresponds to a score of 85, and a false acceptance rate of 0.1% corresponds to a score of 95.
	// -Generally, a score of around 80 applies to most scenarios. It is advisable not to exceed 90. You can choose a suitable score according to the actual situation.
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Person name.
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// Gender of personnel.
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// List of person libraries containing this person and the content of the description field.
	PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitnil,omitempty" name:"PersonGroupInfos"`
}

// Predefined struct for user
type CompareFaceRequestParams struct {
	// Base64-encoded data of image A, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	ImageA *string `json:"ImageA,omitnil,omitempty" name:"ImageA"`

	// Base64-encoded data of image B, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	ImageB *string `json:"ImageB,omitnil,omitempty" name:"ImageB"`

	// URL of image A. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` of image A must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	UrlA *string `json:"UrlA,omitnil,omitempty" name:"UrlA"`

	// URL of image B. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` of image B must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	UrlB *string `json:"UrlB,omitnil,omitempty" name:"UrlB"`

	// Algorithm model version used by the Face Recognition service.
	// 
	// Currently, `2.0` and `3.0` are supported.
	// 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API, `2.0` will be used by default.
	// 
	// The parameter can be set only to `3.0` for accounts that purchase the service after November 26, 2020.
	// 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

type CompareFaceRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded data of image A, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	ImageA *string `json:"ImageA,omitnil,omitempty" name:"ImageA"`

	// Base64-encoded data of image B, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	ImageB *string `json:"ImageB,omitnil,omitempty" name:"ImageB"`

	// URL of image A. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` of image A must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	UrlA *string `json:"UrlA,omitnil,omitempty" name:"UrlA"`

	// URL of image B. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` of image B must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	UrlB *string `json:"UrlB,omitnil,omitempty" name:"UrlB"`

	// Algorithm model version used by the Face Recognition service.
	// 
	// Currently, `2.0` and `3.0` are supported.
	// 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API, `2.0` will be used by default.
	// 
	// The parameter can be set only to `3.0` for accounts that purchase the service after November 26, 2020.
	// 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

func (r *CompareFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageA")
	delete(f, "ImageB")
	delete(f, "UrlA")
	delete(f, "UrlB")
	delete(f, "FaceModelVersion")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompareFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompareFaceResponseParams struct {
	// Face similarity score between two images.
	// The returned similarity score varies by algorithm version. 
	// If you need to verify whether the faces in the two images are the same person, then the 0.1%, 0.01%, and 0.001% FARs on v3.0 correspond to scores of 40, 50, and 60, respectively. Generally, if the score is above 50, it can be judged that they are the same person. 
	// The 0.1%, 0.01%, and 0.001% FARs on v2.0 correspond to scores of 70, 80, and 90, respectively. Generally, if the score is above 80, it can be judged that they are the same person. 
	// If you need to verify whether the faces in the two images are the same person, we recommend using the `VerifyFace` API.
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CompareFaceResponse struct {
	*tchttp.BaseResponse
	Response *CompareFaceResponseParams `json:"Response"`
}

func (r *CompareFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyPersonRequestParams struct {
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// List of groups to join. The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type CopyPersonRequest struct {
	*tchttp.BaseRequest
	
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// List of groups to join. The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *CopyPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyPersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyPersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyPersonResponseParams struct {
	// Number of groups successfully added to.
	SucGroupNum *uint64 `json:"SucGroupNum,omitnil,omitempty" name:"SucGroupNum"`

	// List of groups successfully added to.
	SucGroupIds []*string `json:"SucGroupIds,omitnil,omitempty" name:"SucGroupIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyPersonResponse struct {
	*tchttp.BaseResponse
	Response *CopyPersonResponseParams `json:"Response"`
}

func (r *CopyPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFaceRequestParams struct {
	// <p>Person ID, which is the PersonId in the person creation API.</p>
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// <p>base64 data of the image.</p><ul><li>base64 encoded size must not exceed 5M.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the image long edge must not exceed 2000. The short side pixel of images in all formats must be not less than 64.</li><li>The total number of faces must not exceed 5.</li><li>If the image contains several faces, only the face with the maximum area is selected.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>Image Url.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel must not exceed 4000. For other Image formats, the long side pixel must not exceed 2000. The short side pixel of images in all formats must be not less than 64.</li><li>Either Url or Image must be provided. If both are provided, only Url is used.</li><li>Images stored in Tencent Cloud URLs can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of non-Tencent Cloud Storage URLs may be affected to a certain extent.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li><li>The total number of human faces must not exceed 5.</li><li>If the Image contains several faces, only the face with the maximum area is selected.</li></ul>
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// <p>Only faces with a similarity to the person's existing faces exceeding the FaceMatchThreshold can be added successfully. </p><ul><li>Default value: 60. </li><li>Value range: [0,100]. </li></ul>
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// <p>Image quality control. </p><ul><li>Value range: <br>0: no control; <br>1: low quality requirement, where the image is super blurry, or at least one or more of the following exist: eye occlusion, nose occlusion, or mouth occlusion; <br>2: normal quality requirement, where the image is bright, dark, blurry or generally blurry, or at least three of the following exist: eyebrow obstruction, cheek obstruction, or chin obstruction; <br>3: high quality requirement, where the image is bright, dark, generally blurry, or one to two of the following exist: eyebrow obstruction, cheek obstruction, or chin obstruction; <br>4: super high quality requirement, where each dimension is at its best or at most one dimension has a minor problem; </li><li>Default: 0. </li><li>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// <p>Whether to enable image rotation recognition support.</p><ul><li>Value range: 0 means not enabled, 1 means enabled.</li><li>Default value: 0.</li><li>This parameter is used when the human face in the image is rotated and the image has no exif information. If image rotation recognition support is not enabled, the human face in the image cannot be correctly detected or recognized.</li><li>If you confirm that the image contains exif information or that the human face in the input image will not be rotated, do not enable this parameter. Once enabled, the overall duration may increase by several hundred milliseconds.</li></ul>
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// <p>Whether full-body encryption is required for the request information. Supported encryption algorithm: AES-256-CBC. Users with encryption requirements can use this parameter.</p>
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// <p>Encrypted ciphertext. The data format before encryption is as follows: {&quot;Url&quot;:&quot;<strong>&quot;,&quot;Image&quot;:&quot;</strong>&quot;}。</p>
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

type CreateFaceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Person ID, which is the PersonId in the person creation API.</p>
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// <p>base64 data of the image.</p><ul><li>base64 encoded size must not exceed 5M.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the image long edge must not exceed 2000. The short side pixel of images in all formats must be not less than 64.</li><li>The total number of faces must not exceed 5.</li><li>If the image contains several faces, only the face with the maximum area is selected.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>Image Url.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel must not exceed 4000. For other Image formats, the long side pixel must not exceed 2000. The short side pixel of images in all formats must be not less than 64.</li><li>Either Url or Image must be provided. If both are provided, only Url is used.</li><li>Images stored in Tencent Cloud URLs can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of non-Tencent Cloud Storage URLs may be affected to a certain extent.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li><li>The total number of human faces must not exceed 5.</li><li>If the Image contains several faces, only the face with the maximum area is selected.</li></ul>
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// <p>Only faces with a similarity to the person's existing faces exceeding the FaceMatchThreshold can be added successfully. </p><ul><li>Default value: 60. </li><li>Value range: [0,100]. </li></ul>
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// <p>Image quality control. </p><ul><li>Value range: <br>0: no control; <br>1: low quality requirement, where the image is super blurry, or at least one or more of the following exist: eye occlusion, nose occlusion, or mouth occlusion; <br>2: normal quality requirement, where the image is bright, dark, blurry or generally blurry, or at least three of the following exist: eyebrow obstruction, cheek obstruction, or chin obstruction; <br>3: high quality requirement, where the image is bright, dark, generally blurry, or one to two of the following exist: eyebrow obstruction, cheek obstruction, or chin obstruction; <br>4: super high quality requirement, where each dimension is at its best or at most one dimension has a minor problem; </li><li>Default: 0. </li><li>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// <p>Whether to enable image rotation recognition support.</p><ul><li>Value range: 0 means not enabled, 1 means enabled.</li><li>Default value: 0.</li><li>This parameter is used when the human face in the image is rotated and the image has no exif information. If image rotation recognition support is not enabled, the human face in the image cannot be correctly detected or recognized.</li><li>If you confirm that the image contains exif information or that the human face in the input image will not be rotated, do not enable this parameter. Once enabled, the overall duration may increase by several hundred milliseconds.</li></ul>
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// <p>Whether full-body encryption is required for the request information. Supported encryption algorithm: AES-256-CBC. Users with encryption requirements can use this parameter.</p>
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// <p>Encrypted ciphertext. The data format before encryption is as follows: {&quot;Url&quot;:&quot;<strong>&quot;,&quot;Image&quot;:&quot;</strong>&quot;}。</p>
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

func (r *CreateFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Images")
	delete(f, "Urls")
	delete(f, "FaceMatchThreshold")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	delete(f, "Encryption")
	delete(f, "EncryptedBody")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFaceResponseParams struct {
	// <p>Number of faces that joined successfully.</p>
	SucFaceNum *uint64 `json:"SucFaceNum,omitnil,omitempty" name:"SucFaceNum"`

	// <p>List of face IDs that joined successfully.</p>
	SucFaceIds []*string `json:"SucFaceIds,omitnil,omitempty" name:"SucFaceIds"`

	// <p>Addition result of each face image.</p><ul><li>Other non-zero values represent algorithm service exceptions, with the following value ranges:<br>-1101 represents no face detected.<br>-1102 represents image decoding failure.<br>-1109 represents image size too large or too small.<br>-1601 represents the image quality control requirement is not met.<br>-1604 represents face similarity does not exceed FaceMatchThreshold. </li><li>The order of RetCode is consistent with the sequence of Images or Urls in the input parameters.</li></ul>
	RetCode []*int64 `json:"RetCode,omitnil,omitempty" name:"RetCode"`

	// <p>Index of successfully added faces.</p><ul><li>The index order is consistent with the sequence of Images or Urls in the input parameter.</li><li>For example, if there are 3 Urls in Urls and the second url fails, the value of SucIndexes is [0,2].</li></ul>
	SucIndexes []*uint64 `json:"SucIndexes,omitnil,omitempty" name:"SucIndexes"`

	// <p>Face bounding box position of successfully joined faces. The sequence is consistent with the sequence of Images or Urls in the input parameter.</p>
	SucFaceRects []*FaceRect `json:"SucFaceRects,omitnil,omitempty" name:"SucFaceRects"`

	// <p>Algorithm model version used for face recognition.</p>
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateFaceResponseParams `json:"Response"`
}

func (r *CreateFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupRequestParams struct {
	// Group name, which is modifiable, must be unique, and can contain 1 to 60 characters.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Group ID, which is unmodifiable, must be unique, and can contain letters, digits, and special symbols (-%@#&_) of up to 64 B.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Custom group description field that describes the person attributes in the group, which will be applied to all persons in the group. 
	// Up to 5 ones can be created. 
	// Each custom description field can contain 1 to 30 characters. 
	// The custom description field must be unique in the group. 
	// Example: if you set the "custom description field" of a group to ["student ID","employee ID","mobile number"], 
	// then all the persons in the group will have description fields named "student ID", "employee ID", and "mobile number". 
	// You can enter content in the corresponding field to register a person's student ID, employee ID, and mobile number.
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitnil,omitempty" name:"GroupExDescriptions"`

	// Group remarks, which can contain 0 to 40 characters.
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Algorithm model version used by the Face Recognition service.
	// 
	// Currently, `2.0` and `3.0` are supported.
	// 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API, `2.0` will be used by default.
	// 
	// The parameter can be set only to `3.0` for accounts that purchase the service after November 26, 2020.
	// 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group name, which is modifiable, must be unique, and can contain 1 to 60 characters.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Group ID, which is unmodifiable, must be unique, and can contain letters, digits, and special symbols (-%@#&_) of up to 64 B.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Custom group description field that describes the person attributes in the group, which will be applied to all persons in the group. 
	// Up to 5 ones can be created. 
	// Each custom description field can contain 1 to 30 characters. 
	// The custom description field must be unique in the group. 
	// Example: if you set the "custom description field" of a group to ["student ID","employee ID","mobile number"], 
	// then all the persons in the group will have description fields named "student ID", "employee ID", and "mobile number". 
	// You can enter content in the corresponding field to register a person's student ID, employee ID, and mobile number.
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitnil,omitempty" name:"GroupExDescriptions"`

	// Group remarks, which can contain 0 to 40 characters.
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Algorithm model version used by the Face Recognition service.
	// 
	// Currently, `2.0` and `3.0` are supported.
	// 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API, `2.0` will be used by default.
	// 
	// The parameter can be set only to `3.0` for accounts that purchase the service after November 26, 2020.
	// 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupId")
	delete(f, "GroupExDescriptions")
	delete(f, "Tag")
	delete(f, "FaceModelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupResponseParams struct {
	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupResponseParams `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonRequestParams struct {
	// <p>ID of the personnel database to be added, which is the GroupId in the API for creating a personnel database.</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>Person name. [1, 60] characters, modifiable, repeatable.</p>
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// <p>Person ID.</p><ul><li>Cannot be modified or repeated under a single Tencent Cloud account.</li><li>Supports English letters, digits, -%@#&amp;_, length limited to 64B.</li></ul>
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// <p>Gender.</p><ul><li>Value range: 0 represents empty, 1 represents man, 2 represents woman.</li></ul>
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// <p>Person description field content, key-value. [0, 60] characters, modifiable, repeatable.</p>
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitnil,omitempty" name:"PersonExDescriptionInfos"`

	// <p>base64 image data.</p><ul><li>base64-encoded data must not exceed 5 MB.</li><li>For jpg images, the long edge must not exceed 4000 pixels. For other formats, the long edge must not exceed 2000 pixels. The short edge of images in all formats must be at least 64 pixels.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>Image Url.</p><ul><li>The corresponding Image size after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000. The short side pixel of images in all formats must not be less than 64.</li><li>Either Url or Image must be provided. If both are provided, only Url is used.</li><li>Images stored in Tencent Cloud's Url can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>This parameter is used to control whether the human face contained in the image specified by Image or Url is determined to be a suspected duplicate person already existing in the person library. </p><ul><li>If determined to be the same person already in the person library, no new person will be created, and the personnel information of the suspected duplicate person will be returned. </li><li>If not, the person will be created. </li><li>Value range:<br>0: No judgment is performed, and the person will be stored regardless of whether there is a suspected duplicate person in the library;<br>1: Low duplicate person judgment requirement (1% false recognition rate);<br>2: Normal duplicate person judgment requirement (0.1% false recognition rate);<br>3: High duplicate person judgment requirement (0.01% false recognition rate);<br>4: Very high duplicate person judgment requirement (0.001% false recognition rate). </li><li>Default value: 0. </li><li>Note: The higher the requirement, the lower the probability of suspected duplicate persons. The false recognition rates for different requirements are for reference only and can be adjusted according to the actual situation.</li></ul>
	UniquePersonControl *uint64 `json:"UniquePersonControl,omitnil,omitempty" name:"UniquePersonControl"`

	// <p>Image quality control. </p><ul><li>Value range: <br>0: no control; <br>1: low quality requirement, where the image is very blurry, or at least one or more of the eyes, nose, or mouth are occluded; <br>2: general quality requirement, where the image is bright, dark, blurry or generally blurry, with at least three of the following: eyebrow occlusion, cheek occlusion, or chin occlusion; <br>3: high quality requirement, where the image is bright, dark, generally blurry, with one or two of the following: eyebrow occlusion, cheek occlusion, or chin occlusion; <br>4: very high quality requirement, where each dimension is optimal or at most one dimension has a minor problem; </li><li>Default: 0. </li><li>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// <p>Whether to enable image rotation recognition support.</p><ul><li>Value range: 0 for disabled, 1 for enabled.</li><li>Default value: 0.</li><li>This parameter is used to correctly detect and recognize human faces in an image when the faces are rotated and the image has no exif information. If image rotation recognition support is not enabled in this case, the faces cannot be detected or recognized.</li><li>If you confirm that the image contains exif information or that the human faces in the input image will not be rotated, do not enable this parameter. Once enabled, the overall duration may increase by several hundred milliseconds.</li></ul>
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// <p>Whether to encrypt the entire request body. Supported encryption algorithm: AES-256-CBC. Users with encryption requirements can use this parameter.</p>
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// <p>Encrypted ciphertext. The data format before encryption is as follows: {&quot;Url&quot;:&quot;<strong>&quot;,&quot;Image&quot;:&quot;</strong>&quot;}.</p>
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

type CreatePersonRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the personnel database to be added, which is the GroupId in the API for creating a personnel database.</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>Person name. [1, 60] characters, modifiable, repeatable.</p>
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// <p>Person ID.</p><ul><li>Cannot be modified or repeated under a single Tencent Cloud account.</li><li>Supports English letters, digits, -%@#&amp;_, length limited to 64B.</li></ul>
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// <p>Gender.</p><ul><li>Value range: 0 represents empty, 1 represents man, 2 represents woman.</li></ul>
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// <p>Person description field content, key-value. [0, 60] characters, modifiable, repeatable.</p>
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitnil,omitempty" name:"PersonExDescriptionInfos"`

	// <p>base64 image data.</p><ul><li>base64-encoded data must not exceed 5 MB.</li><li>For jpg images, the long edge must not exceed 4000 pixels. For other formats, the long edge must not exceed 2000 pixels. The short edge of images in all formats must be at least 64 pixels.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>Image Url.</p><ul><li>The corresponding Image size after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000. The short side pixel of images in all formats must not be less than 64.</li><li>Either Url or Image must be provided. If both are provided, only Url is used.</li><li>Images stored in Tencent Cloud's Url can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>This parameter is used to control whether the human face contained in the image specified by Image or Url is determined to be a suspected duplicate person already existing in the person library. </p><ul><li>If determined to be the same person already in the person library, no new person will be created, and the personnel information of the suspected duplicate person will be returned. </li><li>If not, the person will be created. </li><li>Value range:<br>0: No judgment is performed, and the person will be stored regardless of whether there is a suspected duplicate person in the library;<br>1: Low duplicate person judgment requirement (1% false recognition rate);<br>2: Normal duplicate person judgment requirement (0.1% false recognition rate);<br>3: High duplicate person judgment requirement (0.01% false recognition rate);<br>4: Very high duplicate person judgment requirement (0.001% false recognition rate). </li><li>Default value: 0. </li><li>Note: The higher the requirement, the lower the probability of suspected duplicate persons. The false recognition rates for different requirements are for reference only and can be adjusted according to the actual situation.</li></ul>
	UniquePersonControl *uint64 `json:"UniquePersonControl,omitnil,omitempty" name:"UniquePersonControl"`

	// <p>Image quality control. </p><ul><li>Value range: <br>0: no control; <br>1: low quality requirement, where the image is very blurry, or at least one or more of the eyes, nose, or mouth are occluded; <br>2: general quality requirement, where the image is bright, dark, blurry or generally blurry, with at least three of the following: eyebrow occlusion, cheek occlusion, or chin occlusion; <br>3: high quality requirement, where the image is bright, dark, generally blurry, with one or two of the following: eyebrow occlusion, cheek occlusion, or chin occlusion; <br>4: very high quality requirement, where each dimension is optimal or at most one dimension has a minor problem; </li><li>Default: 0. </li><li>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// <p>Whether to enable image rotation recognition support.</p><ul><li>Value range: 0 for disabled, 1 for enabled.</li><li>Default value: 0.</li><li>This parameter is used to correctly detect and recognize human faces in an image when the faces are rotated and the image has no exif information. If image rotation recognition support is not enabled in this case, the faces cannot be detected or recognized.</li><li>If you confirm that the image contains exif information or that the human faces in the input image will not be rotated, do not enable this parameter. Once enabled, the overall duration may increase by several hundred milliseconds.</li></ul>
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// <p>Whether to encrypt the entire request body. Supported encryption algorithm: AES-256-CBC. Users with encryption requirements can use this parameter.</p>
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// <p>Encrypted ciphertext. The data format before encryption is as follows: {&quot;Url&quot;:&quot;<strong>&quot;,&quot;Image&quot;:&quot;</strong>&quot;}.</p>
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

func (r *CreatePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "PersonName")
	delete(f, "PersonId")
	delete(f, "Gender")
	delete(f, "PersonExDescriptionInfos")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "UniquePersonControl")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	delete(f, "Encryption")
	delete(f, "EncryptedBody")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonResponseParams struct {
	// <p>Unique ID of the face image.</p>
	FaceId *string `json:"FaceId,omitnil,omitempty" name:"FaceId"`

	// <p>Position of the detected face frame.</p>
	FaceRect *FaceRect `json:"FaceRect,omitnil,omitempty" name:"FaceRect"`

	// <p>PersonId of the suspected duplicate person. </p><ul><li>This parameter is meaningful only when the UniquePersonControl parameter is not 0 and there is a suspected duplicate person in the person library.</li></ul>
	SimilarPersonId *string `json:"SimilarPersonId,omitnil,omitempty" name:"SimilarPersonId"`

	// <p>Algorithm model version used for face recognition.</p>
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePersonResponse struct {
	*tchttp.BaseResponse
	Response *CreatePersonResponseParams `json:"Response"`
}

func (r *CreatePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFaceRequestParams struct {
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// List of IDs of the faces to be deleted. The array element value is the `FaceId` returned by the `CreateFace` API.
	FaceIds []*string `json:"FaceIds,omitnil,omitempty" name:"FaceIds"`
}

type DeleteFaceRequest struct {
	*tchttp.BaseRequest
	
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// List of IDs of the faces to be deleted. The array element value is the `FaceId` returned by the `CreateFace` API.
	FaceIds []*string `json:"FaceIds,omitnil,omitempty" name:"FaceIds"`
}

func (r *DeleteFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "FaceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFaceResponseParams struct {
	// Number of successfully deleted faces
	SucDeletedNum *uint64 `json:"SucDeletedNum,omitnil,omitempty" name:"SucDeletedNum"`

	// List of IDs of successfully deleted faces
	SucFaceIds []*string `json:"SucFaceIds,omitnil,omitempty" name:"SucFaceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFaceResponseParams `json:"Response"`
}

func (r *DeleteFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonFromGroupRequestParams struct {
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeletePersonFromGroupRequest struct {
	*tchttp.BaseRequest
	
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeletePersonFromGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonFromGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonFromGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonFromGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePersonFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeletePersonFromGroupResponseParams `json:"Response"`
}

func (r *DeletePersonFromGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonFromGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonRequestParams struct {
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

type DeletePersonRequest struct {
	*tchttp.BaseRequest
	
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

func (r *DeletePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePersonResponse struct {
	*tchttp.BaseResponse
	Response *DeletePersonResponseParams `json:"Response"`
}

func (r *DeletePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFaceAttributesRequestParams struct {
	// Maximum number of processable faces. 
	// Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 120. 
	// This parameter is used to control the number of faces in the image to be detected. The smaller the value, the faster the processing.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats. 
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. 
	// The image cannot exceed 5 MB in size after being Base64-encoded. 
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low. 
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Whether to return attributes such as age, gender, and emotion. 
	// Valid values (case-insensitive): None, Age, Beauty, Emotion, Eye, Eyebrow, Gender, Hair, Hat, Headpose, Mask, Mouth, Moustache, Nose, Shape, Skin, Smile. 
	//   
	// `None` indicates that no attributes need to be returned, which is the default value; that is, if the `FaceAttributesType` attribute is empty, the values of all attributes will be `0`.
	// You need to combine the attributes into a string and separate them by comma. The sequence of the attributes is not limited. 
	// For more information on the attributes, see the output parameters as described below. 
	// The face attribute information of up to 5 largest faces in the image will be returned, and `AttributesInfo` of the 6th and rest faces is meaningless.
	FaceAttributesType *string `json:"FaceAttributesType,omitnil,omitempty" name:"FaceAttributesType"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image is not rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// Algorithm model version used by the Face Recognition service. You can enter only `3.0` for this API.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`
}

type DetectFaceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Maximum number of processable faces. 
	// Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 120. 
	// This parameter is used to control the number of faces in the image to be detected. The smaller the value, the faster the processing.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats. 
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. 
	// The image cannot exceed 5 MB in size after being Base64-encoded. 
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low. 
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Whether to return attributes such as age, gender, and emotion. 
	// Valid values (case-insensitive): None, Age, Beauty, Emotion, Eye, Eyebrow, Gender, Hair, Hat, Headpose, Mask, Mouth, Moustache, Nose, Shape, Skin, Smile. 
	//   
	// `None` indicates that no attributes need to be returned, which is the default value; that is, if the `FaceAttributesType` attribute is empty, the values of all attributes will be `0`.
	// You need to combine the attributes into a string and separate them by comma. The sequence of the attributes is not limited. 
	// For more information on the attributes, see the output parameters as described below. 
	// The face attribute information of up to 5 largest faces in the image will be returned, and `AttributesInfo` of the 6th and rest faces is meaningless.
	FaceAttributesType *string `json:"FaceAttributesType,omitnil,omitempty" name:"FaceAttributesType"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image is not rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// Algorithm model version used by the Face Recognition service. You can enter only `3.0` for this API.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`
}

func (r *DetectFaceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxFaceNum")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FaceAttributesType")
	delete(f, "NeedRotateDetection")
	delete(f, "FaceModelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectFaceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFaceAttributesResponseParams struct {
	// Width of requested image.
	ImageWidth *uint64 `json:"ImageWidth,omitnil,omitempty" name:"ImageWidth"`

	// Height of requested image.
	ImageHeight *uint64 `json:"ImageHeight,omitnil,omitempty" name:"ImageHeight"`

	// Face information list.
	FaceDetailInfos []*FaceDetailInfo `json:"FaceDetailInfos,omitnil,omitempty" name:"FaceDetailInfos"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetectFaceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DetectFaceAttributesResponseParams `json:"Response"`
}

func (r *DetectFaceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFaceRequestParams struct {
	// Maximum number of processable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 120. 
	// This parameter is used to control the number of faces in the image to be detected. The smaller the value, the faster the processing.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px.
	// Default value: 34. We recommend keeping it at or above 34.
	// Faces below the `MinFaceSize` value will not be detected.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Whether the face attribute information (FaceAttributesInfo) needs to be returned. 0: no; 1: yes. Default value: 0. 
	// If the value is not 1, it will be deemed as no need to return, and `FaceAttributesInfo` is meaningless in this case.  
	// The face attribute information of up to 5 largest faces in the image will be returned, and `FaceAttributesInfo` of the 6th and rest faces is meaningless.  
	// Extracting face attribute information is quite time-consuming. If face attribute information is not required, we recommend disabling this feature to speed up face detection.
	NeedFaceAttributes *uint64 `json:"NeedFaceAttributes,omitnil,omitempty" name:"NeedFaceAttributes"`

	// Whether to enable quality detection. 0: no; 1: yes. Default value: 0. 
	// If the value is not 1, it will be deemed not to perform quality detection.
	// The face quality score information of up to 30 largest faces in the image will be returned, and `FaceQualityInfo` of the 31st and rest faces is meaningless.  
	// We recommend enabling this feature for the face adding operation.
	NeedQualityDetection *uint64 `json:"NeedQualityDetection,omitnil,omitempty" name:"NeedQualityDetection"`

	// Algorithm model version used by the Face Recognition service.
	// 
	// Currently, `2.0` and `3.0` are supported.
	// 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API, `2.0` will be used by default.
	// 
	// The parameter can be set only to `3.0` for accounts that purchase the service after November 26, 2020.
	// 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

type DetectFaceRequest struct {
	*tchttp.BaseRequest
	
	// Maximum number of processable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 120. 
	// This parameter is used to control the number of faces in the image to be detected. The smaller the value, the faster the processing.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px.
	// Default value: 34. We recommend keeping it at or above 34.
	// Faces below the `MinFaceSize` value will not be detected.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Whether the face attribute information (FaceAttributesInfo) needs to be returned. 0: no; 1: yes. Default value: 0. 
	// If the value is not 1, it will be deemed as no need to return, and `FaceAttributesInfo` is meaningless in this case.  
	// The face attribute information of up to 5 largest faces in the image will be returned, and `FaceAttributesInfo` of the 6th and rest faces is meaningless.  
	// Extracting face attribute information is quite time-consuming. If face attribute information is not required, we recommend disabling this feature to speed up face detection.
	NeedFaceAttributes *uint64 `json:"NeedFaceAttributes,omitnil,omitempty" name:"NeedFaceAttributes"`

	// Whether to enable quality detection. 0: no; 1: yes. Default value: 0. 
	// If the value is not 1, it will be deemed not to perform quality detection.
	// The face quality score information of up to 30 largest faces in the image will be returned, and `FaceQualityInfo` of the 31st and rest faces is meaningless.  
	// We recommend enabling this feature for the face adding operation.
	NeedQualityDetection *uint64 `json:"NeedQualityDetection,omitnil,omitempty" name:"NeedQualityDetection"`

	// Algorithm model version used by the Face Recognition service.
	// 
	// Currently, `2.0` and `3.0` are supported.
	// 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API, `2.0` will be used by default.
	// 
	// The parameter can be set only to `3.0` for accounts that purchase the service after November 26, 2020.
	// 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

func (r *DetectFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "NeedFaceAttributes")
	delete(f, "NeedQualityDetection")
	delete(f, "FaceModelVersion")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFaceResponseParams struct {
	// Width of requested image.
	ImageWidth *int64 `json:"ImageWidth,omitnil,omitempty" name:"ImageWidth"`

	// Height of requested image.
	ImageHeight *int64 `json:"ImageHeight,omitnil,omitempty" name:"ImageHeight"`

	// Face information list, including face coordinate information, attribute information (if needed), and quality score information (if needed).
	FaceInfos []*FaceInfo `json:"FaceInfos,omitnil,omitempty" name:"FaceInfos"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetectFaceResponse struct {
	*tchttp.BaseResponse
	Response *DetectFaceResponseParams `json:"Response"`
}

func (r *DetectFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFaceSimilarityRequestParams struct {
	// <p>A base64-encoded image.</p><ul><li>The size after base64 encoding must not exceed 5 MB.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the long side pixel must not exceed 2000. The short side pixel for all formats must be greater than or equal to 64.</li><li>If the image contains several faces, only the face with the highest confidence degree is selected.</li><li>Supports PNG, jpg, JPEG, and BMP formats. GIF images are not supported.</li></ul>
	ImageA *string `json:"ImageA,omitnil,omitempty" name:"ImageA"`

	// <p>base64 data of image B.</p><ul><li>The size after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000. The short side pixel of images in all formats must not be less than 64.</li><li>If the image contains several faces, only the face with the highest confidence degree is selected.</li><li>Supports PNG, jpg, JPEG, and BMP. GIF images are not supported.</li></ul>
	ImageB *string `json:"ImageB,omitnil,omitempty" name:"ImageB"`

	// <p>Url of Image A.</p><ul><li>Size of the corresponding Image after base64 encoding must not exceed 5 MB.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the long side pixel cannot exceed 2000. The short side pixel for all formats must be not less than 64.</li><li>Either the Url or the Image of picture A must be provided. If both are provided, only the Url is used.</li><li>URLs of images stored in Tencent Cloud guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>Speed and stability of URLs not stored in Tencent Cloud may be impacted.</li><li>If the Image contains several faces, only the human face with the maximum area is selected.</li><li>Supports PNG, jpg, jpg, JPEG, BMP. GIF images are not supported.</li></ul>
	UrlA *string `json:"UrlA,omitnil,omitempty" name:"UrlA"`

	// <p>Url of Image B.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5 MB.</li><li>For jpg format, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000. For all formats, the short side pixel must be not less than 64.</li><li>Either the Url or the Image of picture B must be provided. If both are provided, only the Url is used.</li><li>URLs of images stored in Tencent Cloud guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain degree.</li><li>If the Image contains several faces, only the face with the maximum area is selected.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	UrlB *string `json:"UrlB,omitnil,omitempty" name:"UrlB"`

	// <p>Image quality control. </p><ul><li><p>Value ranges from 0 to 4:<br>0: No control;<br>1: Low quality requirement. The image has at least one of the following: super blurry, eye occlusion, nose occlusion, or mouth occlusion;<br>2: Normal quality requirement. The image has at least three of the following: too bright, dark, blurry, normal blur, eyebrow obstruction, cheek obstruction, or chin obstruction;<br>3: Relatively high quality requirement. The image has one or two of the following: too bright, dark, normal blur, eyebrow obstruction, cheek obstruction, or chin obstruction;<br>4: High quality requirement. All dimensions are the best or most, with a minor problem in one dimension;<br>Default is 0. </p></li><li><p>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</p></li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`
}

type DetectFaceSimilarityRequest struct {
	*tchttp.BaseRequest
	
	// <p>A base64-encoded image.</p><ul><li>The size after base64 encoding must not exceed 5 MB.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the long side pixel must not exceed 2000. The short side pixel for all formats must be greater than or equal to 64.</li><li>If the image contains several faces, only the face with the highest confidence degree is selected.</li><li>Supports PNG, jpg, JPEG, and BMP formats. GIF images are not supported.</li></ul>
	ImageA *string `json:"ImageA,omitnil,omitempty" name:"ImageA"`

	// <p>base64 data of image B.</p><ul><li>The size after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000. The short side pixel of images in all formats must not be less than 64.</li><li>If the image contains several faces, only the face with the highest confidence degree is selected.</li><li>Supports PNG, jpg, JPEG, and BMP. GIF images are not supported.</li></ul>
	ImageB *string `json:"ImageB,omitnil,omitempty" name:"ImageB"`

	// <p>Url of Image A.</p><ul><li>Size of the corresponding Image after base64 encoding must not exceed 5 MB.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the long side pixel cannot exceed 2000. The short side pixel for all formats must be not less than 64.</li><li>Either the Url or the Image of picture A must be provided. If both are provided, only the Url is used.</li><li>URLs of images stored in Tencent Cloud guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>Speed and stability of URLs not stored in Tencent Cloud may be impacted.</li><li>If the Image contains several faces, only the human face with the maximum area is selected.</li><li>Supports PNG, jpg, jpg, JPEG, BMP. GIF images are not supported.</li></ul>
	UrlA *string `json:"UrlA,omitnil,omitempty" name:"UrlA"`

	// <p>Url of Image B.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5 MB.</li><li>For jpg format, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000. For all formats, the short side pixel must be not less than 64.</li><li>Either the Url or the Image of picture B must be provided. If both are provided, only the Url is used.</li><li>URLs of images stored in Tencent Cloud guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain degree.</li><li>If the Image contains several faces, only the face with the maximum area is selected.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	UrlB *string `json:"UrlB,omitnil,omitempty" name:"UrlB"`

	// <p>Image quality control. </p><ul><li><p>Value ranges from 0 to 4:<br>0: No control;<br>1: Low quality requirement. The image has at least one of the following: super blurry, eye occlusion, nose occlusion, or mouth occlusion;<br>2: Normal quality requirement. The image has at least three of the following: too bright, dark, blurry, normal blur, eyebrow obstruction, cheek obstruction, or chin obstruction;<br>3: Relatively high quality requirement. The image has one or two of the following: too bright, dark, normal blur, eyebrow obstruction, cheek obstruction, or chin obstruction;<br>4: High quality requirement. All dimensions are the best or most, with a minor problem in one dimension;<br>Default is 0. </p></li><li><p>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</p></li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`
}

func (r *DetectFaceSimilarityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceSimilarityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageA")
	delete(f, "ImageB")
	delete(f, "UrlA")
	delete(f, "UrlB")
	delete(f, "QualityControl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectFaceSimilarityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFaceSimilarityResponseParams struct {
	// <p>Value ranges from 0.00 to 100.00.<br>It is recommended to judge as the same person when the similarity is equal to or greater than 70 (a fixed threshold of 70 is used, with a false acceptance rate of 1 in 10,000, and the threshold cannot be adjusted).</p>
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetectFaceSimilarityResponse struct {
	*tchttp.BaseResponse
	Response *DetectFaceSimilarityResponseParams `json:"Response"`
}

func (r *DetectFaceSimilarityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceSimilarityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLiveFaceAccurateRequestParams struct {
	// <p>base64 data of the image.</p><ul><li>The size after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the long side pixel must not exceed 2000.</li><li>The aspect ratio of the image should be close to 3:4. The ratio from mobile phone shooting is best.</li><li>The human face size must be greater than 100X100 pixels.</li><li>Format support: PNG, jpg, JPEG, and BMP. GIF images are not supported.</li></ul>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>The Url of the Image.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the long side pixel must not exceed 2000.</li><li>Either Url or Image must be provided. If both are provided, only use Url.</li><li>The aspect ratio of the Image should be close to 3:4, which is the best ratio for mobile phone shooting.</li><li>The human face size must be greater than 100X100 pixels.</li><li>An Image stored in Tencent Cloud via Url can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected.</li><li>Supported Image formats include PNG, jpg, JPEG, and BMP. GIF images are not supported.</li></ul>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Algorithm model version used by the face recognition service. Currently, the input parameter supports "3.0".</p>
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`
}

type DetectLiveFaceAccurateRequest struct {
	*tchttp.BaseRequest
	
	// <p>base64 data of the image.</p><ul><li>The size after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the long side pixel must not exceed 2000.</li><li>The aspect ratio of the image should be close to 3:4. The ratio from mobile phone shooting is best.</li><li>The human face size must be greater than 100X100 pixels.</li><li>Format support: PNG, jpg, JPEG, and BMP. GIF images are not supported.</li></ul>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>The Url of the Image.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel must not exceed 4000. For other formats, the long side pixel must not exceed 2000.</li><li>Either Url or Image must be provided. If both are provided, only use Url.</li><li>The aspect ratio of the Image should be close to 3:4, which is the best ratio for mobile phone shooting.</li><li>The human face size must be greater than 100X100 pixels.</li><li>An Image stored in Tencent Cloud via Url can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud. The speed and stability of URLs not stored in Tencent Cloud may be affected.</li><li>Supported Image formats include PNG, jpg, JPEG, and BMP. GIF images are not supported.</li></ul>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Algorithm model version used by the face recognition service. Currently, the input parameter supports "3.0".</p>
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`
}

func (r *DetectLiveFaceAccurateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLiveFaceAccurateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FaceModelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectLiveFaceAccurateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLiveFaceAccurateResponseParams struct {
	// <p>Liveness authenticity score.</p><ul><li>Value range: [0,100].</li><li>Determine whether it is a rephotograph based on the threshold range corresponding to the liveness score.</li><li>Currently, the threshold can be divided into [5,10,40,70,90], with a recommended threshold of 40.</li><li>The higher the score, the closer to a real person and the lower the rephotograph risk.</li></ul>
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>Algorithm model version used for face recognition.</p>
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetectLiveFaceAccurateResponse struct {
	*tchttp.BaseResponse
	Response *DetectLiveFaceAccurateResponseParams `json:"Response"`
}

func (r *DetectLiveFaceAccurateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLiveFaceAccurateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Encryption struct {
	// <p>When using the encryption service, fill in the field to be encrypted. In this API, you can fill in one or more encrypted fields.</p>
	EncryptList []*string `json:"EncryptList,omitnil,omitempty" name:"EncryptList"`

	// <p>Encrypted symmetric key. For key generation and use, see the <a href="https://www.tencentcloud.com/document/product/1007/47180?from_cn_redirect=1">data encryption</a> document.</p>
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// <p>For users with encryption requirements, import the initial vector for CBC encryption (a custom string of 16 characters).</p>
	Iv *string `json:"Iv,omitnil,omitempty" name:"Iv"`

	// <p>Encryption algorithm (supports 'AES-256-CBC'). If not specified, it defaults to 'AES-256-CBC'</p>
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// <p>Message digest generated by the SM4-GCM algorithm (used when validating message integrity), not supported.</p>
	TagList []*string `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type Eye struct {
	// Whether glasses are worn.
	// The `Type` values of the `AttributeItem` include: 0: no glasses; 1: general glasses; 2: sunglasses.
	Glass *AttributeItem `json:"Glass,omitnil,omitempty" name:"Glass"`

	// Whether the eyes are open.
	// The `Type` values of the `AttributeItem` include: 0: open; 1: closed.
	EyeOpen *AttributeItem `json:"EyeOpen,omitnil,omitempty" name:"EyeOpen"`

	// Whether the person has double eyelids.
	// The `Type` values of the `AttributeItem` include: 0: no; 1: yes.
	EyelidType *AttributeItem `json:"EyelidType,omitnil,omitempty" name:"EyelidType"`

	// Eye size.
	// The `Type` values of the `AttributeItem` include: 0: small eyes; 1: general eyes; 2: big eyes.
	EyeSize *AttributeItem `json:"EyeSize,omitnil,omitempty" name:"EyeSize"`
}

type Eyebrow struct {
	// Eyebrow thickness.
	// The `Type` values of the `AttributeItem` include: 0: light; 1: thick.
	EyebrowDensity *AttributeItem `json:"EyebrowDensity,omitnil,omitempty" name:"EyebrowDensity"`

	// Eyebrow curve.
	// The `Type` values of the `AttributeItem` include: 0: flat; 1: curved.
	EyebrowCurve *AttributeItem `json:"EyebrowCurve,omitnil,omitempty" name:"EyebrowCurve"`

	// Eyebrow length.
	// The `Type` values of the `AttributeItem` include: 0: short; 1: long.
	EyebrowLength *AttributeItem `json:"EyebrowLength,omitnil,omitempty" name:"EyebrowLength"`
}

type FaceAttributesInfo struct {
	// Gender. The gender is female for the value range [0,49] and male for the value range [50,100]. The closer the value to 0 or 100, the higher the confidence. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// Age. Value range: [0,100]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Age *int64 `json:"Age,omitnil,omitempty" name:"Age"`

	// Expression. Value range: [0 (normal)–50 (smile)–100 (laugh)]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Expression *int64 `json:"Expression,omitnil,omitempty" name:"Expression"`

	// Whether glasses are present. Valid values: true, false. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Glass *bool `json:"Glass,omitnil,omitempty" name:"Glass"`

	// Vertical offset in degrees. Value range: [-30,30]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless. 
	// We recommend selecting images in the [-10,10] range for adding faces.
	Pitch *int64 `json:"Pitch,omitnil,omitempty" name:"Pitch"`

	// Horizontal offset in degrees. Value range: [-30,30]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless. 
	// We recommend selecting images in the [-10,10] range for adding faces.
	Yaw *int64 `json:"Yaw,omitnil,omitempty" name:"Yaw"`

	// Horizontal rotation in degrees. Value range: [-180,180]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.  
	// We recommend selecting images in the [-20,20] range for adding faces.
	Roll *int64 `json:"Roll,omitnil,omitempty" name:"Roll"`

	// Beauty. Value range: [0,100]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Beauty *int64 `json:"Beauty,omitnil,omitempty" name:"Beauty"`

	// Whether hat is present. Valid values: true, false. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Hat *bool `json:"Hat,omitnil,omitempty" name:"Hat"`

	// Whether mask is present. Valid values: true, false. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mask *bool `json:"Mask,omitnil,omitempty" name:"Mask"`

	// Hair information, including length, bang, and color. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Hair *FaceHairAttributesInfo `json:"Hair,omitnil,omitempty" name:"Hair"`

	// Whether the eyes are open. Valid values: true, false. As long as there is more than one eye closed, `false` will be returned. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EyeOpen *bool `json:"EyeOpen,omitnil,omitempty" name:"EyeOpen"`
}

type FaceDetailAttributesInfo struct {
	// Age. Value range: [0,65], where 65 indicates 65 years old or above. 
	// If `FaceAttributesType` does not include `Age` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Age *int64 `json:"Age,omitnil,omitempty" name:"Age"`

	// Beauty score. Value range: [0,100]. 
	// If `FaceAttributesType` does not include `Beauty` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Beauty *int64 `json:"Beauty,omitnil,omitempty" name:"Beauty"`

	// Emotion, including relaxed, happy, surprised, angry, sad, disgusted, and scared. 
	// The `Type` values of the `AttributeItem` include: 0: relaxed; 1: happy; 2: surprised; 3: angry; 4: sad; 5: disgusted; 6: scared.
	// If `FaceAttributesType` does not include `Emotion` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Emotion *AttributeItem `json:"Emotion,omitnil,omitempty" name:"Emotion"`

	// Eye information, including whether glasses are worn, whether eyes are closed, whether the person has double eyelids, and the eye size. 
	// If `FaceAttributesType` does not include `Eye` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Eye *Eye `json:"Eye,omitnil,omitempty" name:"Eye"`

	// Eyebrow information, including whether the eyebrows are thick, curved, or long. 
	// If `FaceAttributesType` does not include `Eyebrow` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Eyebrow *Eyebrow `json:"Eyebrow,omitnil,omitempty" name:"Eyebrow"`

	// Gender information. 
	// The `Type` values of the `AttributeItem` include: 0: male; 1: female.	
	// If `FaceAttributesType` does not include `Gender` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Gender *AttributeItem `json:"Gender,omitnil,omitempty" name:"Gender"`

	// Hair information, including length, bang, and color. 
	// If `FaceAttributesType` does not include `Hair` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Hair *Hair `json:"Hair,omitnil,omitempty" name:"Hair"`

	// Hat information, including whether a hat is worn, hat style, and hat color. 
	// If `FaceAttributesType` does not include `Hat` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Hat *Hat `json:"Hat,omitnil,omitempty" name:"Hat"`

	// Pose information, including the face pitch, yaw, and roll. 
	// If `FaceAttributesType` does not include `Headpose` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	HeadPose *HeadPose `json:"HeadPose,omitnil,omitempty" name:"HeadPose"`

	// Mask information. 
	// The `Type` values of the `AttributeItem` include: 0: no mask; 1: the mask is worn and does not cover the face; 2: the mask is worn and covers the chin; 3: the mask is worn and covers the mouth; 4: the mask is worn properly.
	// If `FaceAttributesType` does not include `Mask` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Mask *AttributeItem `json:"Mask,omitnil,omitempty" name:"Mask"`

	// Mouth information, including whether the mouth is open and the lip thickness. 
	// If `FaceAttributesType` does not include `Mouth` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Mouth *Mouth `json:"Mouth,omitnil,omitempty" name:"Mouth"`

	// Beard information.
	// The `Type` values of the `AttributeItem` include: 0: no beard; 1: beard detected. 
	// If `FaceAttributesType` does not include `Moustache` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Moustache *AttributeItem `json:"Moustache,omitnil,omitempty" name:"Moustache"`

	// Nose information. 
	// The `Type` values of the `AttributeItem` include: 0: upturned nose; 1: aquiline nose; 2: general nose; 3: bulbous nose.
	// If `FaceAttributesType` does not include `Nose` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Nose *AttributeItem `json:"Nose,omitnil,omitempty" name:"Nose"`

	// Face shape information. 
	// The `Type` values of the `AttributeItem` include: 0: square; 1: triangular; 2: oval; 3: heart-shaped; 4: round.
	// If `FaceAttributesType` does not include `Shape` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Shape *AttributeItem `json:"Shape,omitnil,omitempty" name:"Shape"`

	// Skin color information. 
	// The `Type` values of the `AttributeItem` include: 0: yellow; 1: brown; 2: black; 3: white.
	// If `FaceAttributesType` does not include `Skin` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Skin *AttributeItem `json:"Skin,omitnil,omitempty" name:"Skin"`

	// Smile level. Value range: [0,100]. 
	// If `FaceAttributesType` does not include `Smile` or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Smile *int64 `json:"Smile,omitnil,omitempty" name:"Smile"`
}

type FaceDetailInfo struct {
	// Position of the detected face frame.
	FaceRect *FaceRect `json:"FaceRect,omitnil,omitempty" name:"FaceRect"`

	// Face attribute information. According to the types specified in `FaceAttributesType`, the following face attributes will be returned: age (Age), beauty score (Beauty), 
	// emotion (Emotion), eye information (Eye), eyebrow information (Eyebrow), gender (Gender), 
	// hair information (Hair), hat information (Hat), pose (Headpose), mask information (Mask), mouth information (Mouse), beard information (Moustache), 
	// nose information (Nose), face shape (Shape), skin color (Skin), and smile information (Smile), etc.  
	// If no types are specified in `FaceAttributesType`, the detailed items returned by `FaceDetaiAttributesInfo` will be meaningless.
	FaceDetailAttributesInfo *FaceDetailAttributesInfo `json:"FaceDetailAttributesInfo,omitnil,omitempty" name:"FaceDetailAttributesInfo"`
}

type FaceHairAttributesInfo struct {
	// 0: shaved head, 1: short hair, 2: medium hair, 3: long hair, 4: braid
	// Note: this field may return null, indicating that no valid values can be obtained.
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 0: with bangs, 1: no bangs
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bang *int64 `json:"Bang,omitnil,omitempty" name:"Bang"`

	// 0: black, 1: golden, 2: brown, 3: gray
	// Note: this field may return null, indicating that no valid values can be obtained.
	Color *int64 `json:"Color,omitnil,omitempty" name:"Color"`
}

type FaceInfo struct {
	// Horizontal coordinate of the top-left vertex of the face frame.
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the `completeness` score meets the requirement.
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Vertical coordinate of the top-left vertex of the face frame. 
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the `completeness` score meets the requirement.
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// Face frame width.
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Face frame height.
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Face attributes, including gender, age, expression, 
	// beauty, glass, mask, hair, and pose (pitch, roll, yaw). Valid information will be returned only if `NeedFaceAttributes` is set to 1.
	FaceAttributesInfo *FaceAttributesInfo `json:"FaceAttributesInfo,omitnil,omitempty" name:"FaceAttributesInfo"`

	// Face quality information, including score, sharpness, brightness, and completeness. Valid information will be returned only if `NeedFaceDetection` is set to 1.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceQualityInfo *FaceQualityInfo `json:"FaceQualityInfo,omitnil,omitempty" name:"FaceQualityInfo"`
}

type FaceQualityCompleteness struct {
	// Eyebrow completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,80], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Eyebrow *int64 `json:"Eyebrow,omitnil,omitempty" name:"Eyebrow"`

	// Eye completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,80], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Eye *int64 `json:"Eye,omitnil,omitempty" name:"Eye"`

	// Nose completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,60], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Nose *int64 `json:"Nose,omitnil,omitempty" name:"Nose"`

	// Cheek completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,70], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cheek *int64 `json:"Cheek,omitnil,omitempty" name:"Cheek"`

	// Mouth completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,50], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mouth *int64 `json:"Mouth,omitnil,omitempty" name:"Mouth"`

	// Chin completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,70], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Chin *int64 `json:"Chin,omitnil,omitempty" name:"Chin"`
}

type FaceQualityInfo struct {
	// Quality score. Value range: [0,100]. It comprehensively evaluates whether the image quality is suitable for face recognition; the higher the score, the higher the quality. 
	// In normal cases, you only need to use `Score` as the overall quality standard score. Specific item scores such as `Sharpness`, `Brightness`, `Completeness` are for reference only.
	// Reference range: [0,40]: poor; [40,60]: fine; [60,80]: good; [80,100]: excellent. 
	// We recommend selecting images with a score above 70 for adding faces.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Sharpness. Value range: [0,100]. It evaluates the sharpness of the image. The higher the score, the sharper the image. 
	// Reference range: [0,40]: very blurry; [40,60]: blurry; [60,80]: fine; [80,100]: sharp. 
	// We recommend selecting images with a score above 80 for adding faces.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Sharpness *int64 `json:"Sharpness,omitnil,omitempty" name:"Sharpness"`

	// Brightness. Value range: [0,100]. The brighter the image, the higher the score. 
	// Reference range: [0,30]: dark; [30,70]: normal; [70,100]: bright. 
	// We recommend selecting images in the [30,70] range for adding faces.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Brightness *int64 `json:"Brightness,omitnil,omitempty" name:"Brightness"`

	// Completeness of facial features, which assesses the completeness of the eyebrows, eyes, nose, cheeks, mouth, and chin.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Completeness *FaceQualityCompleteness `json:"Completeness,omitnil,omitempty" name:"Completeness"`
}

type FaceRect struct {
	// Top-left X-axis coordinate of the face box 
	// -The face frame contains the facial feature locations and is expanded on this basis. If the face frame exceeds the image range, it can cause negative coordinate values. 
	// -If you need to capture a complete face, you can set negative coordinates to 0 when the completeness meets the requirements.
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Top-left Y-axis coordinate of the face box 
	// -The face frame contains the facial feature locations and is expanded on this basis. If the face frame exceeds the image range, it can cause negative coordinate values. 
	// -If you need to capture a complete face, you can set negative coordinates to 0 when the completeness meets the requirements.
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// Face width.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Face height.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type FaceShape struct {
	// 21 points that describe the face contour.
	FaceProfile []*Point `json:"FaceProfile,omitnil,omitempty" name:"FaceProfile"`

	// 8 points that describe the left eye.
	LeftEye []*Point `json:"LeftEye,omitnil,omitempty" name:"LeftEye"`

	// 8 points that describe the right eye.
	RightEye []*Point `json:"RightEye,omitnil,omitempty" name:"RightEye"`

	// 8 points that describe the left eyebrow.
	LeftEyeBrow []*Point `json:"LeftEyeBrow,omitnil,omitempty" name:"LeftEyeBrow"`

	// 8 points that describe the right eyebrow.
	RightEyeBrow []*Point `json:"RightEyeBrow,omitnil,omitempty" name:"RightEyeBrow"`

	// 22 points that describe the mouth.
	Mouth []*Point `json:"Mouth,omitnil,omitempty" name:"Mouth"`

	// 13 points that describe the nose.
	Nose []*Point `json:"Nose,omitnil,omitempty" name:"Nose"`

	// 1 point that describes the left pupil.
	LeftPupil []*Point `json:"LeftPupil,omitnil,omitempty" name:"LeftPupil"`

	// 1 point that describes the right pupil.
	RightPupil []*Point `json:"RightPupil,omitnil,omitempty" name:"RightPupil"`
}

// Predefined struct for user
type GetGroupInfoRequestParams struct {
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type GetGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *GetGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupInfoResponseParams struct {
	// Group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Custom group description field
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitnil,omitempty" name:"GroupExDescriptions"`

	// Group remarks
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// Group creation time and date (`CreationTimestamp`), whose value is the number of milliseconds between the UNIX epoch time and the group creation time.
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitnil,omitempty" name:"CreationTimestamp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupInfoResponseParams `json:"Response"`
}

func (r *GetGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListRequestParams struct {
	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListResponseParams struct {
	// Returned group information
	GroupInfos []*GroupInfo `json:"GroupInfos,omitnil,omitempty" name:"GroupInfos"`

	// Total number of groups
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupNum *uint64 `json:"GroupNum,omitnil,omitempty" name:"GroupNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupListResponseParams `json:"Response"`
}

func (r *GetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonBaseInfoRequestParams struct {
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

type GetPersonBaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

func (r *GetPersonBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPersonBaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonBaseInfoResponseParams struct {
	// Person name
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// Person gender. 0: empty; 1: male; 2: female.
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// List of the IDs of included faces
	FaceIds []*string `json:"FaceIds,omitnil,omitempty" name:"FaceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPersonBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetPersonBaseInfoResponseParams `json:"Response"`
}

func (r *GetPersonBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonGroupInfoRequestParams struct {
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetPersonGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetPersonGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPersonGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonGroupInfoResponseParams struct {
	// List of groups containing this person and their description fields
	PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitnil,omitempty" name:"PersonGroupInfos"`

	// Total number of groups
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupNum *uint64 `json:"GroupNum,omitnil,omitempty" name:"GroupNum"`

	// Algorithm model version used by the Face Recognition service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetPersonGroupInfoResponseParams `json:"Response"`
}

func (r *GetPersonGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListNumRequestParams struct {
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type GetPersonListNumRequest struct {
	*tchttp.BaseRequest
	
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *GetPersonListNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPersonListNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListNumResponseParams struct {
	// Number of persons
	PersonNum *uint64 `json:"PersonNum,omitnil,omitempty" name:"PersonNum"`

	// Number of faces
	FaceNum *uint64 `json:"FaceNum,omitnil,omitempty" name:"FaceNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPersonListNumResponse struct {
	*tchttp.BaseResponse
	Response *GetPersonListNumResponseParams `json:"Response"`
}

func (r *GetPersonListNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListRequestParams struct {
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetPersonListRequest struct {
	*tchttp.BaseRequest
	
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetPersonListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPersonListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListResponseParams struct {
	// Returned person information
	PersonInfos []*PersonInfo `json:"PersonInfos,omitnil,omitempty" name:"PersonInfos"`

	// Number of persons in the group
	// Note: this field may return null, indicating that no valid values can be obtained.
	PersonNum *uint64 `json:"PersonNum,omitnil,omitempty" name:"PersonNum"`

	// Number of faces in the group
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceNum *uint64 `json:"FaceNum,omitnil,omitempty" name:"FaceNum"`

	// Algorithm model version used for face recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPersonListResponse struct {
	*tchttp.BaseResponse
	Response *GetPersonListResponseParams `json:"Response"`
}

func (r *GetPersonListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupCandidate struct {
	// Group ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Most matching candidate recognized
	Candidates []*Candidate `json:"Candidates,omitnil,omitempty" name:"Candidates"`
}

type GroupExDescriptionInfo struct {
	// Custom group description field index, whose value starts from 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupExDescriptionIndex *uint64 `json:"GroupExDescriptionIndex,omitnil,omitempty" name:"GroupExDescriptionIndex"`

	// Content of the custom group description field to be updated
	GroupExDescription *string `json:"GroupExDescription,omitnil,omitempty" name:"GroupExDescription"`
}

type GroupInfo struct {
	// Group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Custom group description field
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitnil,omitempty" name:"GroupExDescriptions"`

	// Group remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Algorithm model version used for face recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// Group creation time and date (`CreationTimestamp`), whose value is the number of milliseconds between the UNIX epoch time and the group creation time. 
	// The UNIX epoch time is 00:00:00, Thursday, January 1, 1970, Coordinated Universal Time (UTC). For more information, please see the UNIX time document.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitnil,omitempty" name:"CreationTimestamp"`
}

type Hair struct {
	// Hair length information.
	// The `Type` values of the `AttributeItem` include: 0: bald, 1: short hair, 2: medium hair, 3: long hair, 4: braid.
	Length *AttributeItem `json:"Length,omitnil,omitempty" name:"Length"`

	// Bang information.
	// The `Type` values of the `AttributeItem` include: 0: no bang; 1: bang detected.
	Bang *AttributeItem `json:"Bang,omitnil,omitempty" name:"Bang"`

	// Hair color information.
	// The `Type` values of the `AttributeItem` include: 0: black; 1: golden; 2: brown; 3: gray.
	Color *AttributeItem `json:"Color,omitnil,omitempty" name:"Color"`
}

type Hat struct {
	// Hat wearing status information.
	// The `Type` values of the `AttributeItem` include: 0: no hat; 1: general hat; 2: helmet; 3: security guard hat.
	Style *AttributeItem `json:"Style,omitnil,omitempty" name:"Style"`

	// Hat color.
	// The `Type` values of the `AttributeItem` include: 0: no hat; 1: red; 2: yellow; 3: blue; 4: black; 5: gray; 6: mixed colors.
	Color *AttributeItem `json:"Color,omitnil,omitempty" name:"Color"`
}

type HeadPose struct {
	// Pitch. Value range: [-30,30].
	Pitch *int64 `json:"Pitch,omitnil,omitempty" name:"Pitch"`

	// Yaw. Value range: [-30,30].
	Yaw *int64 `json:"Yaw,omitnil,omitempty" name:"Yaw"`

	// Roll. Value range: [-180,180].
	Roll *int64 `json:"Roll,omitnil,omitempty" name:"Roll"`
}

// Predefined struct for user
type ModifyGroupRequestParams struct {
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Custom description field of the group to be modified, which is a `key-value` pair.
	GroupExDescriptionInfos []*GroupExDescriptionInfo `json:"GroupExDescriptionInfos,omitnil,omitempty" name:"GroupExDescriptionInfos"`

	// Group remarks
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Custom description field of the group to be modified, which is a `key-value` pair.
	GroupExDescriptionInfos []*GroupExDescriptionInfo `json:"GroupExDescriptionInfos,omitnil,omitempty" name:"GroupExDescriptionInfos"`

	// Group remarks
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "GroupExDescriptionInfos")
	delete(f, "Tag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGroupResponseParams `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonGroupInfoRequestParams struct {
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Custom description field of the person to be modified, which is a `key-value` pair.
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitnil,omitempty" name:"PersonExDescriptionInfos"`
}

type ModifyPersonGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Custom description field of the person to be modified, which is a `key-value` pair.
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitnil,omitempty" name:"PersonExDescriptionInfos"`
}

func (r *ModifyPersonGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "PersonId")
	delete(f, "PersonExDescriptionInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPersonGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonGroupInfoResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPersonGroupInfoResponseParams `json:"Response"`
}

func (r *ModifyPersonGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Mouth struct {
	// Whether the mouth is open.
	// The `Type` values of the `AttributeItem` include: 0: closed; 1: open.
	MouthOpen *AttributeItem `json:"MouthOpen,omitnil,omitempty" name:"MouthOpen"`
}

type PersonExDescriptionInfo struct {
	// Person description field Index, starts from 0.
	PersonExDescriptionIndex *uint64 `json:"PersonExDescriptionIndex,omitnil,omitempty" name:"PersonExDescriptionIndex"`

	// Person description field content that needs to be updated.
	PersonExDescription *string `json:"PersonExDescription,omitnil,omitempty" name:"PersonExDescription"`
}

type PersonGroupInfo struct {
	// Personnel database ID that contains this person.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Person description field content.
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitnil,omitempty" name:"PersonExDescriptions"`
}

type PersonInfo struct {
	// Person name
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// Person ID
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Person gender
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// Content of person description field
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitnil,omitempty" name:"PersonExDescriptions"`

	// List of contained face images
	FaceIds []*string `json:"FaceIds,omitnil,omitempty" name:"FaceIds"`

	// Person creation time, measured in the number of milliseconds elapsed since the Unix epoch 
	// The Unix epoch is 00:00:00, Thursday, January 1, 1970, Coordinated Universal Time (UTC). For more information, please see the Unix time document.
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitnil,omitempty" name:"CreationTimestamp"`
}

type Point struct {
	// X coordinate
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Y coordinate
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`
}

type Result struct {
	// The most similar candidate identified.
	Candidates []*Candidate `json:"Candidates,omitnil,omitempty" name:"Candidates"`

	// Face bounding box position detected.
	FaceRect *FaceRect `json:"FaceRect,omitnil,omitempty" name:"FaceRect"`

	// Status return code of the detected face image.
	// - 0 indicates normal. 
	// -1601 indicates the image quality control requirement is not met. At this point, the Candidate content is empty.
	RetCode *int64 `json:"RetCode,omitnil,omitempty" name:"RetCode"`
}

type ResultsReturnsByGroup struct {
	// Position of detected face frame
	FaceRect *FaceRect `json:"FaceRect,omitnil,omitempty" name:"FaceRect"`

	// Recognition result.
	GroupCandidates []*GroupCandidate `json:"GroupCandidates,omitnil,omitempty" name:"GroupCandidates"`

	// Status return code of detected face image. 0: normal. 
	// -1601: the image quality control requirement is not met; in this case, `Candidate` is empty.
	RetCode *int64 `json:"RetCode,omitnil,omitempty" name:"RetCode"`
}

// Predefined struct for user
type SearchFacesRequestParams struct {
	// <p>List of person libraries to search, with a maximum of 100.</p><ul><li>The array element value is the GroupId in the person library creation API.</li><li>You cannot search person libraries with different algorithm model versions (FaceModelVersion) at the same time.</li></ul>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>base64 data of the image.</p><ul><li>base64-encoded data must not exceed 5 MB.</li><li>For jpg images, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000.</li><li>The short side pixel of images in all formats must be at least 64.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>The Url of the Image.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5M.</li><li>The long side pixel of a jpg Image cannot exceed 4000, and that of other formats cannot exceed 2000.</li><li>The short side pixel of images in all formats must be not less than 64.</li><li>Either Url or Image must be provided. If both are provided, only Url is used.</li><li>An Image stored in Tencent Cloud via Url can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.</li><li>PNG, jpg, JPEG, and BMP are supported, while GIF images are not supported.</li></ul>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Maximum number of faces to be recognized.</p><ul><li>Default value: 1 (only the face with the maximum area in the image is detected). Maximum value: 10. </li><li>MaxFaceNum is used to set the number of faces to search when the input image to be recognized contains several faces. </li><li>For example, if the image in the input Image or Url contains several faces and MaxFaceNum is set to 5, the 5 faces with the maximum area in the image will be recognized.</li></ul>
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// <p>Minimum size of the face length and width, unit: pixel.</p><ul><li>Default value: 34.</li><li>Face images smaller than 34 cannot be recognized.</li><li>It is recommended to set it to 80.</li></ul>
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// <p>Number of most similar persons returned for a single recognized human face.</p><ul><li>Default value: 5. Maximum value: 100. </li><li>For example, if MaxFaceNum is 1 and MaxPersonNum is 8, the Top 8 similar personnel information is returned.</li><li>A larger value means a longer time to process.</li><li>It is recommended not to exceed 10.</li></ul>
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil,omitempty" name:"MaxPersonNum"`

	// <p>Whether to return detailed personnel information.</p><ul><li>Value range: 0 to disable, 1 to enable.</li><li>Default value: 0.</li><li>Other values other than 0 and 1 default to 0.</li></ul>
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// <p>Image quality control. </p><ul><li>Value range: <br>0: No control. <br>1: Low quality requirement. The image is super blurry, with at least one or more of the following: eye, nose, or mouth occlusion. <br>2: General quality requirement. The image is bright, dark, blurry or generally blurry, with at least three of the following: eyebrow occlusion, cheek occlusion, or chin occlusion. <br>3: High quality requirement. The image is bright, dark, generally blurry, with one or two of the following: eyebrow occlusion, cheek occlusion, or chin occlusion. <br>4: Very high quality requirement. Each dimension is optimal or at most one dimension has a minor problem.</li><li>Default: 0. </li><li>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// <p>In the output parameter Score, only results exceeding the FaceMatchThreshold value will be returned. The default is 0.</p>
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// <p>Whether to enable image rotation recognition support.</p><ul><li>Value range: 0 for not enabled, 1 for enabled.</li><li>Default value: 0.</li><li>This parameter is used to correctly detect and recognize human faces in the input image when the faces are rotated and the image has no exif information. If image rotation recognition support is not enabled, the faces cannot be detected or recognized correctly.</li><li>If you confirm that the image contains exif information or that the human faces in the input image will not be rotated, do not enable this parameter. Once enabled, the overall duration may increase by several hundred milliseconds.</li></ul>
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// <p>Whether to encrypt the entire request body. Supported encryption algorithms: AES-256-CBC. Users with encryption requirements can use this parameter.</p>
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// <p>Encrypted ciphertext. The data format before encryption is as follows: {&quot;Url&quot;:&quot;<strong>&quot;,&quot;Image&quot;:&quot;</strong>&quot;}.</p>
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

type SearchFacesRequest struct {
	*tchttp.BaseRequest
	
	// <p>List of person libraries to search, with a maximum of 100.</p><ul><li>The array element value is the GroupId in the person library creation API.</li><li>You cannot search person libraries with different algorithm model versions (FaceModelVersion) at the same time.</li></ul>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>base64 data of the image.</p><ul><li>base64-encoded data must not exceed 5 MB.</li><li>For jpg images, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000.</li><li>The short side pixel of images in all formats must be at least 64.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>The Url of the Image.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5M.</li><li>The long side pixel of a jpg Image cannot exceed 4000, and that of other formats cannot exceed 2000.</li><li>The short side pixel of images in all formats must be not less than 64.</li><li>Either Url or Image must be provided. If both are provided, only Url is used.</li><li>An Image stored in Tencent Cloud via Url can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of URLs not stored in Tencent Cloud may be affected to a certain extent.</li><li>PNG, jpg, JPEG, and BMP are supported, while GIF images are not supported.</li></ul>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Maximum number of faces to be recognized.</p><ul><li>Default value: 1 (only the face with the maximum area in the image is detected). Maximum value: 10. </li><li>MaxFaceNum is used to set the number of faces to search when the input image to be recognized contains several faces. </li><li>For example, if the image in the input Image or Url contains several faces and MaxFaceNum is set to 5, the 5 faces with the maximum area in the image will be recognized.</li></ul>
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// <p>Minimum size of the face length and width, unit: pixel.</p><ul><li>Default value: 34.</li><li>Face images smaller than 34 cannot be recognized.</li><li>It is recommended to set it to 80.</li></ul>
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// <p>Number of most similar persons returned for a single recognized human face.</p><ul><li>Default value: 5. Maximum value: 100. </li><li>For example, if MaxFaceNum is 1 and MaxPersonNum is 8, the Top 8 similar personnel information is returned.</li><li>A larger value means a longer time to process.</li><li>It is recommended not to exceed 10.</li></ul>
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil,omitempty" name:"MaxPersonNum"`

	// <p>Whether to return detailed personnel information.</p><ul><li>Value range: 0 to disable, 1 to enable.</li><li>Default value: 0.</li><li>Other values other than 0 and 1 default to 0.</li></ul>
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// <p>Image quality control. </p><ul><li>Value range: <br>0: No control. <br>1: Low quality requirement. The image is super blurry, with at least one or more of the following: eye, nose, or mouth occlusion. <br>2: General quality requirement. The image is bright, dark, blurry or generally blurry, with at least three of the following: eyebrow occlusion, cheek occlusion, or chin occlusion. <br>3: High quality requirement. The image is bright, dark, generally blurry, with one or two of the following: eyebrow occlusion, cheek occlusion, or chin occlusion. <br>4: Very high quality requirement. Each dimension is optimal or at most one dimension has a minor problem.</li><li>Default: 0. </li><li>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// <p>In the output parameter Score, only results exceeding the FaceMatchThreshold value will be returned. The default is 0.</p>
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// <p>Whether to enable image rotation recognition support.</p><ul><li>Value range: 0 for not enabled, 1 for enabled.</li><li>Default value: 0.</li><li>This parameter is used to correctly detect and recognize human faces in the input image when the faces are rotated and the image has no exif information. If image rotation recognition support is not enabled, the faces cannot be detected or recognized correctly.</li><li>If you confirm that the image contains exif information or that the human faces in the input image will not be rotated, do not enable this parameter. Once enabled, the overall duration may increase by several hundred milliseconds.</li></ul>
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// <p>Whether to encrypt the entire request body. Supported encryption algorithms: AES-256-CBC. Users with encryption requirements can use this parameter.</p>
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// <p>Encrypted ciphertext. The data format before encryption is as follows: {&quot;Url&quot;:&quot;<strong>&quot;,&quot;Image&quot;:&quot;</strong>&quot;}.</p>
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

func (r *SearchFacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "MaxPersonNum")
	delete(f, "NeedPersonInfo")
	delete(f, "QualityControl")
	delete(f, "FaceMatchThreshold")
	delete(f, "NeedRotateDetection")
	delete(f, "Encryption")
	delete(f, "EncryptedBody")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchFacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFacesResponseParams struct {
	// <p>Recognition result.</p>
	Results []*Result `json:"Results,omitnil,omitempty" name:"Results"`

	// <p>Number of faces in the searched group.</p>
	FaceNum *uint64 `json:"FaceNum,omitnil,omitempty" name:"FaceNum"`

	// <p>Algorithm model version used for face recognition.</p>
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchFacesResponse struct {
	*tchttp.BaseResponse
	Response *SearchFacesResponseParams `json:"Response"`
}

func (r *SearchFacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFacesReturnsByGroupRequestParams struct {
	// List of groups to be searched in (up to 60). The array element value is the `GroupId` in the `CreateGroup` API.
	// You cannot search for groups using different algorithm model versions (`FaceModelVersion`) at a time.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Detected faces, which is corresponding to the maximum number of returned most matching persons. Default value: 5. Maximum value: 10.  
	// For example, if `MaxFaceNum` is 3, `MaxPersonNumPerGroup` is 5, and the `GroupIds` length is 3, up to 45 (3 * 5 * 3) persons will be returned.
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitnil,omitempty" name:"MaxPersonNumPerGroup"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is greater than or equal to the `FaceMatchThreshold` value.
	// Default value: 0.
	// Value range: [0.0,100.0).
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

type SearchFacesReturnsByGroupRequest struct {
	*tchttp.BaseRequest
	
	// List of groups to be searched in (up to 60). The array element value is the `GroupId` in the `CreateGroup` API.
	// You cannot search for groups using different algorithm model versions (`FaceModelVersion`) at a time.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Detected faces, which is corresponding to the maximum number of returned most matching persons. Default value: 5. Maximum value: 10.  
	// For example, if `MaxFaceNum` is 3, `MaxPersonNumPerGroup` is 5, and the `GroupIds` length is 3, up to 45 (3 * 5 * 3) persons will be returned.
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitnil,omitempty" name:"MaxPersonNumPerGroup"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is greater than or equal to the `FaceMatchThreshold` value.
	// Default value: 0.
	// Value range: [0.0,100.0).
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchFacesReturnsByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesReturnsByGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "MaxPersonNumPerGroup")
	delete(f, "NeedPersonInfo")
	delete(f, "QualityControl")
	delete(f, "FaceMatchThreshold")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchFacesReturnsByGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFacesReturnsByGroupResponseParams struct {
	// Number of faces included in searched groups.
	FaceNum *uint64 `json:"FaceNum,omitnil,omitempty" name:"FaceNum"`

	// Recognition result.
	ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitnil,omitempty" name:"ResultsReturnsByGroup"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchFacesReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *SearchFacesReturnsByGroupResponseParams `json:"Response"`
}

func (r *SearchFacesReturnsByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesReturnsByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPersonsRequestParams struct {
	// <p>List of person libraries to query, with a capacity limit of 100. The value of an array element is the GroupId in the person library creation API.</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>Image base64 data.</p><ul><li>Size after base64 encoding must not exceed 5M.</li><li>Long side pixel of jpg images cannot exceed 4000, and that of other formats cannot exceed 2000.</li><li>Short side pixel of images in all formats must not be less than 64.</li><li>If an image contains several faces, only the face with the maximum area is selected.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>Image Url.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000.</li><li>The short side pixel of images in all formats must be not less than 64.</li><li>Either Url or Image must be provided. If both are provided, only Url is used.</li><li>Urls of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of Urls not stored in Tencent Cloud may be affected to a certain extent.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Maximum number of human faces to be recognized.</p><ul><li>Default value: 1 (only the face with the maximum area in the image is detected). Maximum value: 10.</li><li>MaxFaceNum is used to set the number of faces to search when the input image to be recognized contains several faces.</li><li>For example, if the image in the input Image or Url contains several faces and MaxFaceNum is set to 5, the 5 faces with the maximum area in the image will be recognized.</li></ul>
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// <p>Minimum length and width of a human face, unit: pixel.</p><ul><li>Default value: 34.</li><li>Less than 34 will affect search accuracy.</li><li>Set it to 80.</li></ul>
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// <p>Number of most similar persons returned for a single recognized human face.</p><ul><li>Default value: 5. Maximum value: 100.</li><li>For example, if MaxFaceNum is 1 and MaxPersonNum is 8, the information of the Top 8 similar persons is returned.</li><li>A larger value means a longer time to process. It is advisable not to exceed 10.</li></ul>
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil,omitempty" name:"MaxPersonNum"`

	// <p>Image quality control. </p><ul><li>Value range:<br>0: No control.<br>1: Low quality requirement. The image is super blurry, with at least one of the following: eye, nose, or mouth occlusion.<br>2: Normal quality requirement. The image is bright, dark, blurry or generally blurry, with at least three of the following: eyebrow occlusion, cheek occlusion, or chin occlusion.<br>3: High quality requirement. The image is bright, dark, generally blurry, with one or two of the following: eyebrow occlusion, cheek occlusion, or chin occlusion.<br>4: Very high quality requirement. Each dimension is optimal or at most one dimension has a minor problem.</li><li>Default: 0. </li><li>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// <p>In the output parameter Score, only results equal to or greater than the FaceMatchThreshold value will be returned.</p><ul><li>Default value: 0.</li><li>Value range: [0.0,100.0).</li></ul>
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// <p>Whether to return specific personnel information.</p><ul><li>Value range: 0 to disable, 1 to enable.</li><li>Default value: 0.</li><li>Other values other than 0 and 1 default to 0.</li></ul>
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// <p>Whether to enable image rotation recognition support.</p><ul><li>Value range: 0 for not enabled, 1 for enabled.</li><li>Default value: 0.</li><li>This parameter is used when the human face in the image is rotated and the image has no exif information. If image rotation recognition support is not enabled, the human face in the image cannot be correctly detected and recognized.</li><li>If you confirm that the image contains exif information or that the human face in the input image will not be rotated, do not enable this parameter. Once enabled, the overall duration may increase by several hundred milliseconds.</li></ul>
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// <p>Whether full request body encryption is required. Supported encryption algorithm: AES-256-CBC. Users with encryption requirements can use this parameter.</p>
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// <p>Encrypted ciphertext. The data format before encryption is as follows: {&quot;Url&quot;:&quot;<strong>&quot;,&quot;Image&quot;:&quot;</strong>&quot;}.</p>
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

type SearchPersonsRequest struct {
	*tchttp.BaseRequest
	
	// <p>List of person libraries to query, with a capacity limit of 100. The value of an array element is the GroupId in the person library creation API.</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>Image base64 data.</p><ul><li>Size after base64 encoding must not exceed 5M.</li><li>Long side pixel of jpg images cannot exceed 4000, and that of other formats cannot exceed 2000.</li><li>Short side pixel of images in all formats must not be less than 64.</li><li>If an image contains several faces, only the face with the maximum area is selected.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>Image Url.</p><ul><li>The size of the corresponding Image after base64 encoding must not exceed 5M.</li><li>For jpg format, the long side pixel cannot exceed 4000. For other formats, the long side pixel cannot exceed 2000.</li><li>The short side pixel of images in all formats must be not less than 64.</li><li>Either Url or Image must be provided. If both are provided, only Url is used.</li><li>Urls of images stored in Tencent Cloud can guarantee higher download speed and stability. It is recommended to store images in Tencent Cloud.</li><li>The speed and stability of Urls not stored in Tencent Cloud may be affected to a certain extent.</li><li>PNG, jpg, JPEG, and BMP are supported. GIF images are not supported.</li></ul>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Maximum number of human faces to be recognized.</p><ul><li>Default value: 1 (only the face with the maximum area in the image is detected). Maximum value: 10.</li><li>MaxFaceNum is used to set the number of faces to search when the input image to be recognized contains several faces.</li><li>For example, if the image in the input Image or Url contains several faces and MaxFaceNum is set to 5, the 5 faces with the maximum area in the image will be recognized.</li></ul>
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// <p>Minimum length and width of a human face, unit: pixel.</p><ul><li>Default value: 34.</li><li>Less than 34 will affect search accuracy.</li><li>Set it to 80.</li></ul>
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// <p>Number of most similar persons returned for a single recognized human face.</p><ul><li>Default value: 5. Maximum value: 100.</li><li>For example, if MaxFaceNum is 1 and MaxPersonNum is 8, the information of the Top 8 similar persons is returned.</li><li>A larger value means a longer time to process. It is advisable not to exceed 10.</li></ul>
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil,omitempty" name:"MaxPersonNum"`

	// <p>Image quality control. </p><ul><li>Value range:<br>0: No control.<br>1: Low quality requirement. The image is super blurry, with at least one of the following: eye, nose, or mouth occlusion.<br>2: Normal quality requirement. The image is bright, dark, blurry or generally blurry, with at least three of the following: eyebrow occlusion, cheek occlusion, or chin occlusion.<br>3: High quality requirement. The image is bright, dark, generally blurry, with one or two of the following: eyebrow occlusion, cheek occlusion, or chin occlusion.<br>4: Very high quality requirement. Each dimension is optimal or at most one dimension has a minor problem.</li><li>Default: 0. </li><li>If the image quality fails to satisfy the requirement, the returned result will prompt that the image quality detection does not meet the requirement.</li></ul>
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// <p>In the output parameter Score, only results equal to or greater than the FaceMatchThreshold value will be returned.</p><ul><li>Default value: 0.</li><li>Value range: [0.0,100.0).</li></ul>
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// <p>Whether to return specific personnel information.</p><ul><li>Value range: 0 to disable, 1 to enable.</li><li>Default value: 0.</li><li>Other values other than 0 and 1 default to 0.</li></ul>
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// <p>Whether to enable image rotation recognition support.</p><ul><li>Value range: 0 for not enabled, 1 for enabled.</li><li>Default value: 0.</li><li>This parameter is used when the human face in the image is rotated and the image has no exif information. If image rotation recognition support is not enabled, the human face in the image cannot be correctly detected and recognized.</li><li>If you confirm that the image contains exif information or that the human face in the input image will not be rotated, do not enable this parameter. Once enabled, the overall duration may increase by several hundred milliseconds.</li></ul>
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`

	// <p>Whether full request body encryption is required. Supported encryption algorithm: AES-256-CBC. Users with encryption requirements can use this parameter.</p>
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// <p>Encrypted ciphertext. The data format before encryption is as follows: {&quot;Url&quot;:&quot;<strong>&quot;,&quot;Image&quot;:&quot;</strong>&quot;}.</p>
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

func (r *SearchPersonsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "MaxPersonNum")
	delete(f, "QualityControl")
	delete(f, "FaceMatchThreshold")
	delete(f, "NeedPersonInfo")
	delete(f, "NeedRotateDetection")
	delete(f, "Encryption")
	delete(f, "EncryptedBody")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchPersonsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPersonsResponseParams struct {
	// <p>Recognition result.</p>
	Results []*Result `json:"Results,omitnil,omitempty" name:"Results"`

	// <p>Number of personnel included in the searched group. If all faces in the input image fail to meet the quality requirement, 0 is returned.</p>
	PersonNum *uint64 `json:"PersonNum,omitnil,omitempty" name:"PersonNum"`

	// <p>Algorithm model version used for face recognition.</p>
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchPersonsResponse struct {
	*tchttp.BaseResponse
	Response *SearchPersonsResponseParams `json:"Response"`
}

func (r *SearchPersonsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPersonsReturnsByGroupRequestParams struct {
	// List of groups to be searched in (up to 60). The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Detected faces, which is corresponding to the maximum number of returned most matching persons. Default value: 5. Maximum value: 10.  
	// For example, if `MaxFaceNum` is 3, `MaxPersonNumPerGroup` is 5, and the `GroupIds` length is 3, up to 45 (3 * 5 * 3) persons will be returned.
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitnil,omitempty" name:"MaxPersonNumPerGroup"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is above the `FaceMatchThreshold` value. Default value: 0.
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

type SearchPersonsReturnsByGroupRequest struct {
	*tchttp.BaseRequest
	
	// List of groups to be searched in (up to 60). The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Detected faces, which is corresponding to the maximum number of returned most matching persons. Default value: 5. Maximum value: 10.  
	// For example, if `MaxFaceNum` is 3, `MaxPersonNumPerGroup` is 5, and the `GroupIds` length is 3, up to 45 (3 * 5 * 3) persons will be returned.
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitnil,omitempty" name:"MaxPersonNumPerGroup"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is above the `FaceMatchThreshold` value. Default value: 0.
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchPersonsReturnsByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsReturnsByGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "MaxPersonNumPerGroup")
	delete(f, "QualityControl")
	delete(f, "FaceMatchThreshold")
	delete(f, "NeedPersonInfo")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchPersonsReturnsByGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPersonsReturnsByGroupResponseParams struct {
	// Number of the persons included in searched groups. If the quality of all faces in the input image does not meet the requirement, 0 will be returned.
	PersonNum *uint64 `json:"PersonNum,omitnil,omitempty" name:"PersonNum"`

	// Recognition result.
	ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitnil,omitempty" name:"ResultsReturnsByGroup"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchPersonsReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *SearchPersonsReturnsByGroupResponseParams `json:"Response"`
}

func (r *SearchPersonsReturnsByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsReturnsByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyFaceRequestParams struct {
	// ID of the person to be verified. For more information on `PersonId`, please see the group management APIs.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

type VerifyFaceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the person to be verified. For more information on `PersonId`, please see the group management APIs.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

func (r *VerifyFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyFaceResponseParams struct {
	// Similarity between given face image and `PersonId`. If there are multiple faces under the `PersonId`, the score of the highest similarity will be returned.
	// 
	// The returned similarity score varies by algorithm version.
	// If you need to verify whether the faces in the two images are the same person, then the 0.1%, 0.01%, and 0.001% FARs on v3.0 correspond to scores of 40, 50, and 60, respectively. Generally, if the score is above 50, it can be judged that they are the same person.
	// The 0.1%, 0.01%, and 0.001% FARs on v2.0 correspond to scores of 70, 80, and 90, respectively. Generally, if the score is above 80, it can be judged that they are the same person.
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Whether the person is the one in the image. The fixed threshold score is 60. If you want to adjust the threshold more flexibly, you can take the returned `Score` parameter value for judgment.
	IsMatch *bool `json:"IsMatch,omitnil,omitempty" name:"IsMatch"`

	// Algorithm model version used for face recognition in the group where the `Person` is, which is set when the group is created.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyFaceResponse struct {
	*tchttp.BaseResponse
	Response *VerifyFaceResponseParams `json:"Response"`
}

func (r *VerifyFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPersonRequestParams struct {
	// ID of the person to be verified. For more information on `PersonId`, please see the group management APIs.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Base64-encoded data of the image.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL 
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	//  Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

type VerifyPersonRequest struct {
	*tchttp.BaseRequest
	
	// ID of the person to be verified. For more information on `PersonId`, please see the group management APIs.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Base64-encoded data of the image.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL 
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	//  Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

func (r *VerifyPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyPersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPersonResponseParams struct {
	// Similarity between given face image and `PersonId`. If there are multiple faces under the `PersonId`, their information will be fused for verification.
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Whether the person in the image matches the `PersonId`.
	IsMatch *bool `json:"IsMatch,omitnil,omitempty" name:"IsMatch"`

	// Algorithm model version used for face recognition in the group where the `Person` is, which is set when the group is created.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyPersonResponse struct {
	*tchttp.BaseResponse
	Response *VerifyPersonResponseParams `json:"Response"`
}

func (r *VerifyPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}