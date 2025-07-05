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

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
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
	// Person ID
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Face ID, which is valid only when returned by the `SearchFaces` or `SearchFacesReturnsByGroup` API. User search APIs use facial feature fusion to search for users, for which this field is meaningless.
	FaceId *string `json:"FaceId,omitnil,omitempty" name:"FaceId"`

	// Match score of candidate. 
	// 
	// In a face base library containing 10,000 faces, the 1%, 0.1%, and 0.01% FARs correspond to scores of 70, 80, and 90, respectively;
	// In a face base library containing 100,000 faces, the 1%, 0.1%, and 0.01% FARs correspond to scores of 80, 90, and 100, respectively;
	// In a face base library containing 300,000 faces, the 1% and 0.1% FARs correspond to scores of 85 and 95, respectively.
	// 
	// Generally, the score of 80 is suitable for most scenarios. We recommend choosing an appropriate score based on the actual situation, preferably no more than 90.
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Person name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// Person gender
	// Note: this field may return null, indicating that no valid values can be obtained.
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// List of groups containing this person and their description fields
	// Note: this field may return null, indicating that no valid values can be obtained.
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

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
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
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// A person can have up to 5 face images.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	// A person can have up to 5 face images.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// Only faces whose similarity to an existing face of the person is above the value of `FaceMatchThreshold` can be added successfully. 
	// Default value: 60. Value range: [0,100].
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

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

type CreateFaceRequest struct {
	*tchttp.BaseRequest
	
	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// A person can have up to 5 face images.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	// A person can have up to 5 face images.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// Only faces whose similarity to an existing face of the person is above the value of `FaceMatchThreshold` can be added successfully. 
	// Default value: 60. Value range: [0,100].
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFaceResponseParams struct {
	// Number of successfully added faces
	SucFaceNum *uint64 `json:"SucFaceNum,omitnil,omitempty" name:"SucFaceNum"`

	// List of IDs of successfully added faces
	SucFaceIds []*string `json:"SucFaceIds,omitnil,omitempty" name:"SucFaceIds"`

	// Adding result for each face image. -1101: no face detected; -1102: image decoding failed; 
	// -1601: the image quality control requirement is not met; -1604: the face similarity is not above `FaceMatchThreshold`. 
	// Other non-zero values: algorithm service exception. 
	// The order of `RetCode` values is the same as the order of `Images` or `Urls` in the input parameter.
	RetCode []*int64 `json:"RetCode,omitnil,omitempty" name:"RetCode"`

	// Indexes of successfully added faces. The order of indexes is the same as the order of `Images` or `Urls` in the input parameter. 
	// For example, if there are 3 URLs in `Urls`, and the second URL fails, then the value of `SucIndexes` will be [0,2].
	SucIndexes []*uint64 `json:"SucIndexes,omitnil,omitempty" name:"SucIndexes"`

	// Frame positions of successfully added faces. The order is the same as the order of `Images` or `Urls` in the input parameter.
	SucFaceRects []*FaceRect `json:"SucFaceRects,omitnil,omitempty" name:"SucFaceRects"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
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
	// ID of the group to join, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Person name, which can contain 1 to 60 characters and is modifiable and repeatable.
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// Person ID, which is unmodifiable, must be unique under a Tencent Cloud account, and can contain letters, digits, and special symbols (-%@#&_) of up to 64 B.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 0: empty; 1: male; 2: female.
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// Content of person description field, which is a `key-value` pair, can contain 0 to 60 characters, and is modifiable and repeatable.
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitnil,omitempty" name:"PersonExDescriptionInfos"`

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

	// This parameter is used to control the judgment whether the face contained in the image in `Image` or `Url` corresponds to an existing person in the group. 
	// If it is judged that a duplicate person exists in the group, no new person will be created, and information of the suspected duplicate person will be returned. 
	// Otherwise, the new person will be created. 
	// 0: do not judge, i.e., the person will be created no matter whether a duplicate person exists in the group. 
	// 1: low duplicate person judgment requirement (1% FAR); 
	// 2: average duplicate person judgment requirement (0.1% FAR); 
	// 3: high duplicate person judgment requirement (0.01% FAR); 
	// 4: very high duplicate person judgment requirement (0.001% FAR). 
	// Default value: 0.  
	// Note: the higher the requirement, the lower the probability of duplicate person. The FARs corresponding to different requirements are for reference only and can be adjusted as needed.
	UniquePersonControl *uint64 `json:"UniquePersonControl,omitnil,omitempty" name:"UniquePersonControl"`

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

type CreatePersonRequest struct {
	*tchttp.BaseRequest
	
	// ID of the group to join, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Person name, which can contain 1 to 60 characters and is modifiable and repeatable.
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// Person ID, which is unmodifiable, must be unique under a Tencent Cloud account, and can contain letters, digits, and special symbols (-%@#&_) of up to 64 B.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 0: empty; 1: male; 2: female.
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// Content of person description field, which is a `key-value` pair, can contain 0 to 60 characters, and is modifiable and repeatable.
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitnil,omitempty" name:"PersonExDescriptionInfos"`

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

	// This parameter is used to control the judgment whether the face contained in the image in `Image` or `Url` corresponds to an existing person in the group. 
	// If it is judged that a duplicate person exists in the group, no new person will be created, and information of the suspected duplicate person will be returned. 
	// Otherwise, the new person will be created. 
	// 0: do not judge, i.e., the person will be created no matter whether a duplicate person exists in the group. 
	// 1: low duplicate person judgment requirement (1% FAR); 
	// 2: average duplicate person judgment requirement (0.1% FAR); 
	// 3: high duplicate person judgment requirement (0.01% FAR); 
	// 4: very high duplicate person judgment requirement (0.001% FAR). 
	// Default value: 0.  
	// Note: the higher the requirement, the lower the probability of duplicate person. The FARs corresponding to different requirements are for reference only and can be adjusted as needed.
	UniquePersonControl *uint64 `json:"UniquePersonControl,omitnil,omitempty" name:"UniquePersonControl"`

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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonResponseParams struct {
	// Unique ID of face image.
	FaceId *string `json:"FaceId,omitnil,omitempty" name:"FaceId"`

	// Position of detected face frame.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceRect *FaceRect `json:"FaceRect,omitnil,omitempty" name:"FaceRect"`

	// `PersonId` of suspected duplicate person. 
	// This parameter is meaningful only if the `UniquePersonControl` parameter is not 0 and there is a suspected duplicate person in the group.
	SimilarPersonId *string `json:"SimilarPersonId,omitnil,omitempty" name:"SimilarPersonId"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
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
	// A image base64 data.
	//  - The size after base64 encoding cannot exceed 5M. 
	// - The long side pixels of jpg format cannot exceed 4000, and the long side pixels of pictures in other formats cannot exceed 2000. The short side of images in all formats must be no less than 64 pixels. 
	// - If the picture contains multiple faces, only the face with the highest confidence is selected. - Supports PNG, JPG, JPEG, BMP, but does not support GIF images.
	ImageA *string `json:"ImageA,omitnil,omitempty" name:"ImageA"`

	// B image base64 data.
	//  - The size after base64 encoding cannot exceed 5M. 
	// - The long side pixels of jpg format cannot exceed 4000, and the long side pixels of pictures in other formats cannot exceed 2000. The short side of images in all formats must be no less than 64 pixels. 
	// - If the picture contains multiple faces, only the face with the highest confidence is selected. - Supports PNG, JPG, JPEG, BMP, but does not support GIF images.
	ImageB *string `json:"ImageB,omitnil,omitempty" name:"ImageB"`

	// A URL for the image. 
	// - The size of the corresponding image after base64 encoding cannot exceed 5M. 
	// - The long side pixels of jpg format cannot exceed 4000, and the long side pixels of pictures in other formats cannot exceed 2000. The short side of images in all formats must be no less than 64 pixels. 
	// - A The URL and Image of the picture must be provided. If both are provided, only the Url will be used. 
	// - Images stored in Tencent Cloud's Url can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. 
	// - The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent. 
	// - If the picture contains multiple faces, only the face with the largest face area is selected. 
	// - Supports PNG, JPG, JPEG, BMP, but does not support GIF images.
	UrlA *string `json:"UrlA,omitnil,omitempty" name:"UrlA"`

	// B The URL of the image. 
	// - The size of the corresponding image after base64 encoding cannot exceed 5M. 
	// - The long side pixels of jpg format cannot exceed 4000, and the long side pixels of pictures in other formats cannot exceed 2000. The short side of images in all formats must be no less than 64 pixels. 
	// - B The URL and Image of the picture must be provided. If both are provided, only the Url will be used. 
	// - Images stored in Tencent Cloud's Url can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. 
	// - The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent. 
	// - If the picture contains multiple faces, only the face with the largest face area is selected. 
	// - Supports PNG, JPG, JPEG, BMP, but does not support GIF images.
	UrlB *string `json:"UrlB,omitnil,omitempty" name:"UrlB"`

	// Image quality control. 
	// - Value range: 0: No control; 1: Lower quality requirements, the image is very blurry, and the eyes, nose, and mouth cover at least one or more of them; 2: General quality requirements, the image is bright, Dark, blurry or generally blurred, eyebrows covered, cheeks covered, chin covered, at least three of them; 3: High quality requirements, the image may be brighter, darker, generally blurry, eyebrows blocked, cheeks blocked, chin blocked, one or two of them; 4: Very high quality requirements, all dimensions are the best or the most , there is a slight problem in one dimension; default is 0. 
	// - If the image quality does not meet the requirements, the returned result will prompt that the image quality test does not meet the requirements.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`
}

type DetectFaceSimilarityRequest struct {
	*tchttp.BaseRequest
	
	// A image base64 data.
	//  - The size after base64 encoding cannot exceed 5M. 
	// - The long side pixels of jpg format cannot exceed 4000, and the long side pixels of pictures in other formats cannot exceed 2000. The short side of images in all formats must be no less than 64 pixels. 
	// - If the picture contains multiple faces, only the face with the highest confidence is selected. - Supports PNG, JPG, JPEG, BMP, but does not support GIF images.
	ImageA *string `json:"ImageA,omitnil,omitempty" name:"ImageA"`

	// B image base64 data.
	//  - The size after base64 encoding cannot exceed 5M. 
	// - The long side pixels of jpg format cannot exceed 4000, and the long side pixels of pictures in other formats cannot exceed 2000. The short side of images in all formats must be no less than 64 pixels. 
	// - If the picture contains multiple faces, only the face with the highest confidence is selected. - Supports PNG, JPG, JPEG, BMP, but does not support GIF images.
	ImageB *string `json:"ImageB,omitnil,omitempty" name:"ImageB"`

	// A URL for the image. 
	// - The size of the corresponding image after base64 encoding cannot exceed 5M. 
	// - The long side pixels of jpg format cannot exceed 4000, and the long side pixels of pictures in other formats cannot exceed 2000. The short side of images in all formats must be no less than 64 pixels. 
	// - A The URL and Image of the picture must be provided. If both are provided, only the Url will be used. 
	// - Images stored in Tencent Cloud's Url can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. 
	// - The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent. 
	// - If the picture contains multiple faces, only the face with the largest face area is selected. 
	// - Supports PNG, JPG, JPEG, BMP, but does not support GIF images.
	UrlA *string `json:"UrlA,omitnil,omitempty" name:"UrlA"`

	// B The URL of the image. 
	// - The size of the corresponding image after base64 encoding cannot exceed 5M. 
	// - The long side pixels of jpg format cannot exceed 4000, and the long side pixels of pictures in other formats cannot exceed 2000. The short side of images in all formats must be no less than 64 pixels. 
	// - B The URL and Image of the picture must be provided. If both are provided, only the Url will be used. 
	// - Images stored in Tencent Cloud's Url can ensure higher download speed and stability. It is recommended that images be stored in Tencent Cloud. 
	// - The URL speed and stability of non-Tencent cloud storage may be affected to a certain extent. 
	// - If the picture contains multiple faces, only the face with the largest face area is selected. 
	// - Supports PNG, JPG, JPEG, BMP, but does not support GIF images.
	UrlB *string `json:"UrlB,omitnil,omitempty" name:"UrlB"`

	// Image quality control. 
	// - Value range: 0: No control; 1: Lower quality requirements, the image is very blurry, and the eyes, nose, and mouth cover at least one or more of them; 2: General quality requirements, the image is bright, Dark, blurry or generally blurred, eyebrows covered, cheeks covered, chin covered, at least three of them; 3: High quality requirements, the image may be brighter, darker, generally blurry, eyebrows blocked, cheeks blocked, chin blocked, one or two of them; 4: Very high quality requirements, all dimensions are the best or the most , there is a slight problem in one dimension; default is 0. 
	// - If the image quality does not meet the requirements, the returned result will prompt that the image quality test does not meet the requirements.
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
	// The value range is [0.00, 100.00]. It is recommended that when the similarity is greater than or equal to 70, the person can be judged to be the same person, and the threshold can be adjusted according to the specific scenario (the false pass rate for a threshold of 70 is one in 1,000, and the false pass rate for a threshold of 80 is one in 10,000).
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
	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in .jpg format or 2,000 px for images in other formats. 
	// The recommended image aspect ratio is 3:4 (generally, the aspect ratio of images taken by mobile phones).
	// The face must be greater than 100*100 px in size.
	// Supported image formats are PNG, JPG, JPEG, and BMP. GIF is currently not supported.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in .jpg format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// The recommended image aspect ratio is 3:4 (generally, the aspect ratio of images taken by mobile phones).
	// The face must be greater than 100*100 px in size.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	// Supported image formats are PNG, JPG, JPEG, and BMP. GIF is currently not supported.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Algorithm model version used for face recognition. Valid value: `3.0`.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`
}

type DetectLiveFaceAccurateRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in .jpg format or 2,000 px for images in other formats. 
	// The recommended image aspect ratio is 3:4 (generally, the aspect ratio of images taken by mobile phones).
	// The face must be greater than 100*100 px in size.
	// Supported image formats are PNG, JPG, JPEG, and BMP. GIF is currently not supported.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in .jpg format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// The recommended image aspect ratio is 3:4 (generally, the aspect ratio of images taken by mobile phones).
	// The face must be greater than 100*100 px in size.
	// We recommend you store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. The download speed and stability of non-Tencent Cloud URLs may be low.
	// Supported image formats are PNG, JPG, JPEG, and BMP. GIF is currently not supported.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Algorithm model version used for face recognition. Valid value: `3.0`.
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
	// Liveness score. Value range: [0, 100]. You can set several thresholds such as 5, 10, 40, 70 and 90 to determine whether the image is photographed. We recommend you use the threshold of 40.
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
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

// Predefined struct for user
type DetectLiveFaceRequestParams struct {
	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats (the aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless).
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// (The aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless.) 
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
}

type DetectLiveFaceRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats (the aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless).
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// (The aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless.) 
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
}

func (r *DetectLiveFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLiveFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FaceModelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectLiveFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLiveFaceResponseParams struct {
	// Liveness score. Value range: [0,100]. The score is generally between 80 and 100, but 0 is also a common value. As a recommendation, when the score is greater than 87, it can be judged that the person in the image is alive. You can adjust the threshold according to your specific scenario.
	// This field is meaningful only if `FaceModelVersion` is 2.0.
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// Whether liveness detection is passed.
	// This field is meaningful only if `FaceModelVersion` is 3.0.
	IsLiveness *bool `json:"IsLiveness,omitnil,omitempty" name:"IsLiveness"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetectLiveFaceResponse struct {
	*tchttp.BaseResponse
	Response *DetectLiveFaceResponseParams `json:"Response"`
}

func (r *DetectLiveFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLiveFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// Horizontal coordinate of the top-left vertex of face frame. 
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the completeness score meets the requirement.
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Vertical coordinate of the top-left vertex of face frame. 
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the completeness score meets the requirement.
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// Face width
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Face height
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
	// Person description field index, whose value starts from 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PersonExDescriptionIndex *uint64 `json:"PersonExDescriptionIndex,omitnil,omitempty" name:"PersonExDescriptionIndex"`

	// Content of the person description field to be updated
	PersonExDescription *string `json:"PersonExDescription,omitnil,omitempty" name:"PersonExDescription"`
}

type PersonGroupInfo struct {
	// ID of the group that contains this person
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Content of person description field
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
	// Most matching candidate recognized
	Candidates []*Candidate `json:"Candidates,omitnil,omitempty" name:"Candidates"`

	// Position of detected face frame
	FaceRect *FaceRect `json:"FaceRect,omitnil,omitempty" name:"FaceRect"`

	// The status return code for the face image detected. Valid values: `0` - normal; `-1601` - the image does not meet the quality requirements, in which case `Candidate` is empty; `-1604` - the face similarity is not higher than `FaceMatchThreshold`.
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
	// List of groups to be searched in (up to 100). The array element value is the `GroupId` in the `CreateGroup` API.
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

	// Minimum height and width of face in px. Default value: 34. Face images whose value is below 34 cannot be recognized. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Number of the most similar persons returned for one single recognized face image. Default value: 5. Maximum value: 100. 
	// For example, if `MaxFaceNum` is 1 and `MaxPersonNum` is 8, information of the top 8 most similar persons will be returned.
	// The greater the value, the longer the processing time. We recommend setting a value below 10.
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil,omitempty" name:"MaxPersonNum"`

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

	// In the output parameter `Score`, the result will be returned only if the result value is above the `FaceMatchThreshold` value. Default value: 0.
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

type SearchFacesRequest struct {
	*tchttp.BaseRequest
	
	// List of groups to be searched in (up to 100). The array element value is the `GroupId` in the `CreateGroup` API.
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

	// Minimum height and width of face in px. Default value: 34. Face images whose value is below 34 cannot be recognized. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Number of the most similar persons returned for one single recognized face image. Default value: 5. Maximum value: 100. 
	// For example, if `MaxFaceNum` is 1 and `MaxPersonNum` is 8, information of the top 8 most similar persons will be returned.
	// The greater the value, the longer the processing time. We recommend setting a value below 10.
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil,omitempty" name:"MaxPersonNum"`

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

	// In the output parameter `Score`, the result will be returned only if the result value is above the `FaceMatchThreshold` value. Default value: 0.
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchFacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFacesResponseParams struct {
	// Recognition result.
	Results []*Result `json:"Results,omitnil,omitempty" name:"Results"`

	// Number of faces included in searched groups.
	FaceNum *uint64 `json:"FaceNum,omitnil,omitempty" name:"FaceNum"`

	// Algorithm model version used for face recognition.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
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
	// List of groups to be searched in (up to 100). The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

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
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Number of the most similar persons returned for one single recognized face image. Default value: 5. Maximum value: 100.
	// For example, if `MaxFaceNum` is 1 and `MaxPersonNum` is 8, information of the top 8 most similar persons will be returned.
	// The greater the value, the longer the processing time. We recommend setting a value below 10.
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil,omitempty" name:"MaxPersonNum"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is greater than or equal to the `FaceMatchThreshold` value. Default value: 0. Value range: [0.0,100.0).
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
}

type SearchPersonsRequest struct {
	*tchttp.BaseRequest
	
	// List of groups to be searched in (up to 100). The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

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
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitnil,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitnil,omitempty" name:"MinFaceSize"`

	// Number of the most similar persons returned for one single recognized face image. Default value: 5. Maximum value: 100.
	// For example, if `MaxFaceNum` is 1 and `MaxPersonNum` is 8, information of the top 8 most similar persons will be returned.
	// The greater the value, the longer the processing time. We recommend setting a value below 10.
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil,omitempty" name:"MaxPersonNum"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitnil,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is greater than or equal to the `FaceMatchThreshold` value. Default value: 0. Value range: [0.0,100.0).
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitnil,omitempty" name:"FaceMatchThreshold"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitnil,omitempty" name:"NeedPersonInfo"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitnil,omitempty" name:"NeedRotateDetection"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchPersonsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPersonsResponseParams struct {
	// Recognition result.
	Results []*Result `json:"Results,omitnil,omitempty" name:"Results"`

	// Number of the persons included in searched groups. If the quality of all faces in the input image does not meet the requirement, 0 will be returned.
	PersonNum *uint64 `json:"PersonNum,omitnil,omitempty" name:"PersonNum"`

	// Algorithm model version used for face recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceModelVersion *string `json:"FaceModelVersion,omitnil,omitempty" name:"FaceModelVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
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