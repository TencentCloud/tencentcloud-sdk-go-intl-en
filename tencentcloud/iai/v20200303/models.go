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

package v20200303

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AnalyzeFaceRequest struct {
	*tchttp.BaseRequest

	// Detection mode. 0: detect all faces that appear; 1: detect the largest face. Default value: 0. The facial feature localization information (facial keypoints) of up to 10 faces can be returned.
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0.  
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default.  
	// Different algorithm model versions correspond to different face recognition algorithms. The new version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("AnalyzeFaceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Width of requested image.
		ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

		// Height of requested image.
		ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

		// Specific information of facial feature localization (facial keypoints).
		FaceShapeSet []*FaceShape `json:"FaceShapeSet,omitempty" name:"FaceShapeSet" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Candidate struct {

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Face ID
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// Match score of candidate. 
	// 
	// In a face base library containing 10,000 faces, the 1%, 0.1%, and 0.01% FARs correspond to scores of 70, 80, and 90, respectively;
	// In a face base library containing 100,000 faces, the 1%, 0.1%, and 0.01% FARs correspond to scores of 80, 90, and 100, respectively;
	// In a face base library containing 300,000 faces, the 1% and 0.1% FARs correspond to scores of 85 and 95, respectively.
	// 
	// Generally, the score of 80 is suitable for most scenarios. We recommend choosing an appropriate score based on the actual situation, preferably no more than 90.
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// Person name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// Person gender
	// Note: this field may return null, indicating that no valid values can be obtained.
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// List of groups containing this person and their description fields
	// Note: this field may return null, indicating that no valid values can be obtained.
	PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitempty" name:"PersonGroupInfos" list`
}

type CompareFaceRequest struct {
	*tchttp.BaseRequest

	// Base64-encoded data of image A, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	ImageA *string `json:"ImageA,omitempty" name:"ImageA"`

	// Base64-encoded data of image B, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	ImageB *string `json:"ImageB,omitempty" name:"ImageB"`

	// URL of image A. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` of image A must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	UrlA *string `json:"UrlA,omitempty" name:"UrlA"`

	// URL of image B. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` of image B must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	UrlB *string `json:"UrlB,omitempty" name:"UrlB"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0. 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default. 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("CompareFaceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CompareFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Face similarity score between two images.
	// The returned similarity score varies by algorithm version. 
	// If you need to verify whether the faces in the two images are the same person, then the 0.1%, 0.01%, and 0.001% FARs on v3.0 correspond to scores of 40, 50, and 60, respectively. Generally, if the score is above 50, it can be judged that they are the same person. 
	// The 0.1%, 0.01%, and 0.001% FARs on v2.0 correspond to scores of 70, 80, and 90, respectively. Generally, if the score is above 80, it can be judged that they are the same person. 
	// If you need to verify whether the faces in the two images are the same person, we recommend using the `VerifyFace` API.
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyPersonRequest struct {
	*tchttp.BaseRequest

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// List of groups to join. The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyPersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return errors.New("CopyPersonRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CopyPersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of groups successfully added to.
		SucGroupNum *uint64 `json:"SucGroupNum,omitempty" name:"SucGroupNum"`

		// List of groups successfully added to.
		SucGroupIds []*string `json:"SucGroupIds,omitempty" name:"SucGroupIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFaceRequest struct {
	*tchttp.BaseRequest

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// A person can have up to 5 face images.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Images []*string `json:"Images,omitempty" name:"Images" list`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	// A person can have up to 5 face images.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// Only faces whose similarity to an existing face of the person is above the value of `FaceMatchThreshold` can be added successfully. 
	// Default value: 60. Value range: [0,100].
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("CreateFaceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of successfully added faces
		SucFaceNum *uint64 `json:"SucFaceNum,omitempty" name:"SucFaceNum"`

		// List of IDs of successfully added faces
		SucFaceIds []*string `json:"SucFaceIds,omitempty" name:"SucFaceIds" list`

		// Adding result for each face image. -1101: no face detected; -1102: image decoding failed; 
	// -1601: the image quality control requirement is not met; -1604: the face similarity is not above `FaceMatchThreshold`. 
	// Other non-zero values: algorithm service exception. 
	// The order of `RetCode` values is the same as the order of `Images` or `Urls` in the input parameter.
		RetCode []*int64 `json:"RetCode,omitempty" name:"RetCode" list`

		// Indexes of successfully added faces. The order of indexes is the same as the order of `Images` or `Urls` in the input parameter. 
	// For example, if there are 3 URLs in `Urls`, and the second URL fails, then the value of `SucIndexes` will be [0,2].
		SucIndexes []*uint64 `json:"SucIndexes,omitempty" name:"SucIndexes" list`

		// Frame positions of successfully added faces. The order is the same as the order of `Images` or `Urls` in the input parameter.
		SucFaceRects []*FaceRect `json:"SucFaceRects,omitempty" name:"SucFaceRects" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// Group name, which is modifiable, must be unique, and can contain 1 to 60 characters.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Group ID, which is unmodifiable, must be unique, and can contain letters, digits, and special symbols (-%@#&_) of up to 64 B.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Custom group description field that describes the person attributes in the group, which will be applied to all persons in the group. 
	// Up to 5 ones can be created. 
	// Each custom description field can contain 1 to 30 characters. 
	// The custom description field must be unique in the group. 
	// Example: if you set the "custom description field" of a group to ["student ID","employee ID","mobile number"], 
	// then all the persons in the group will have description fields named "student ID", "employee ID", and "mobile number". 
	// You can enter content in the corresponding field to register a person's student ID, employee ID, and mobile number.
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions" list`

	// Group remarks, which can contain 0 to 40 characters.
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0.
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default. 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("CreateGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePersonRequest struct {
	*tchttp.BaseRequest

	// ID of the group to join, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Person name, which can contain 1 to 60 characters and is modifiable and repeatable.
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// Person ID, which is unmodifiable, must be unique under a Tencent Cloud account, and can contain letters, digits, and special symbols (-%@#&_) of up to 64 B.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 0: empty; 1: male; 2: female.
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// Content of person description field, which is a `key-value` pair, can contain 0 to 60 characters, and is modifiable and repeatable.
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

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
	UniquePersonControl *uint64 `json:"UniquePersonControl,omitempty" name:"UniquePersonControl"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("CreatePersonRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of face image.
		FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

		// Position of detected face frame.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

		// `PersonId` of suspected duplicate person. 
	// This parameter is meaningful only if the `UniquePersonControl` parameter is not 0 and there is a suspected duplicate person in the group.
		SimilarPersonId *string `json:"SimilarPersonId,omitempty" name:"SimilarPersonId"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFaceRequest struct {
	*tchttp.BaseRequest

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// List of IDs of the faces to be deleted. The array element value is the `FaceId` returned by the `CreateFace` API.
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "FaceIds")
	if len(f) > 0 {
		return errors.New("DeleteFaceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of successfully deleted faces
		SucDeletedNum *uint64 `json:"SucDeletedNum,omitempty" name:"SucDeletedNum"`

		// List of IDs of successfully deleted faces
		SucFaceIds []*string `json:"SucFaceIds,omitempty" name:"SucFaceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("DeleteGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePersonFromGroupRequest struct {
	*tchttp.BaseRequest

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonFromGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("DeletePersonFromGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePersonFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonFromGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePersonRequest struct {
	*tchttp.BaseRequest

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return errors.New("DeletePersonRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectFaceRequest struct {
	*tchttp.BaseRequest

	// Maximum number of processable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 120. 
	// This parameter is used to control the number of faces in the image to be detected. The smaller the value, the faster the processing.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px.
	// Default value: 34. We recommend keeping it at or above 34.
	// Faces below the `MinFaceSize` value will not be detected.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Whether the face attribute information (FaceAttributesInfo) needs to be returned. 0: no; 1: yes. Default value: 0. 
	// If the value is not 1, it will be deemed as no need to return, and `FaceAttributesInfo` is meaningless in this case.  
	// The face attribute information of up to 5 largest faces in the image will be returned, and `FaceAttributesInfo` of the 6th and rest faces is meaningless.  
	// Extracting face attribute information is quite time-consuming. If face attribute information is not required, we recommend disabling this feature to speed up face detection.
	NeedFaceAttributes *uint64 `json:"NeedFaceAttributes,omitempty" name:"NeedFaceAttributes"`

	// Whether to enable quality detection. 0: no; 1: yes. Default value: 0. 
	// If the value is not 1, it will be deemed not to perform quality detection.
	// The face quality score information of up to 30 largest faces in the image will be returned, and `FaceQualityInfo` of the 31st and rest faces is meaningless.  
	// We recommend enabling this feature for the face adding operation.
	NeedQualityDetection *uint64 `json:"NeedQualityDetection,omitempty" name:"NeedQualityDetection"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0.  
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default. 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DetectFaceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetectFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Width of requested image.
		ImageWidth *int64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

		// Height of requested image.
		ImageHeight *int64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

		// Face information list, including face coordinate information, attribute information (if needed), and quality score information (if needed).
		FaceInfos []*FaceInfo `json:"FaceInfos,omitempty" name:"FaceInfos" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectLiveFaceRequest struct {
	*tchttp.BaseRequest

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats (the aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless).
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// (The aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless.) 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0.  
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default. 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DetectLiveFaceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetectLiveFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Liveness score. Value range: [0,100]. The score is generally between 80 and 100, but 0 is also a common value. As a recommendation, when the score is greater than 87, it can be judged that the person in the image is alive. You can adjust the threshold according to your specific scenario.
	// This field is meaningful only if `FaceModelVersion` is 2.0.
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// Whether liveness detection is passed.
	// This field is meaningful only if `FaceModelVersion` is 3.0.
		IsLiveness *bool `json:"IsLiveness,omitempty" name:"IsLiveness"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLiveFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceAttributesInfo struct {

	// Gender. The gender is female for the value range [0,49] and male for the value range [50,100]. The closer the value to 0 or 100, the higher the confidence. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// Age. Value range: [0,100]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// Expression. Value range: [0 (normal)–50 (smile)–100 (laugh)]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Expression *int64 `json:"Expression,omitempty" name:"Expression"`

	// Whether glasses are present. Valid values: true, false. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Glass *bool `json:"Glass,omitempty" name:"Glass"`

	// Vertical offset in degrees. Value range: [-30,30]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless. 
	// We recommend selecting images in the [-10,10] range for adding faces.
	Pitch *int64 `json:"Pitch,omitempty" name:"Pitch"`

	// Horizontal offset in degrees. Value range: [-30,30]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless. 
	// We recommend selecting images in the [-10,10] range for adding faces.
	Yaw *int64 `json:"Yaw,omitempty" name:"Yaw"`

	// Horizontal rotation in degrees. Value range: [-180,180]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.  
	// We recommend selecting images in the [-20,20] range for adding faces.
	Roll *int64 `json:"Roll,omitempty" name:"Roll"`

	// Beauty. Value range: [0,100]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Beauty *int64 `json:"Beauty,omitempty" name:"Beauty"`

	// Whether hat is present. Valid values: true, false. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Hat *bool `json:"Hat,omitempty" name:"Hat"`

	// Whether mask is present. Valid values: true, false. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mask *bool `json:"Mask,omitempty" name:"Mask"`

	// Hair information, including length, bang, and color. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Hair *FaceHairAttributesInfo `json:"Hair,omitempty" name:"Hair"`

	// Whether the eyes are open. Valid values: true, false. As long as there is more than one eye closed, `false` will be returned. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EyeOpen *bool `json:"EyeOpen,omitempty" name:"EyeOpen"`
}

type FaceHairAttributesInfo struct {

	// 0: shaved head, 1: short hair, 2: medium hair, 3: long hair, 4: braid
	// Note: this field may return null, indicating that no valid values can be obtained.
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 0: with bangs, 1: no bangs
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bang *int64 `json:"Bang,omitempty" name:"Bang"`

	// 0: black, 1: golden, 2: brown, 3: gray
	// Note: this field may return null, indicating that no valid values can be obtained.
	Color *int64 `json:"Color,omitempty" name:"Color"`
}

type FaceInfo struct {

	// Horizontal coordinate of the top-left vertex of the face frame.
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the `completeness` score meets the requirement.
	X *int64 `json:"X,omitempty" name:"X"`

	// Vertical coordinate of the top-left vertex of the face frame. 
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the `completeness` score meets the requirement.
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// Face frame width.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Face frame height.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Face attributes, including gender, age, expression, 
	// beauty, glass, mask, hair, and pose (pitch, roll, yaw). Valid information will be returned only if `NeedFaceAttributes` is set to 1.
	FaceAttributesInfo *FaceAttributesInfo `json:"FaceAttributesInfo,omitempty" name:"FaceAttributesInfo"`

	// Face quality information, including score, sharpness, brightness, and completeness. Valid information will be returned only if `NeedFaceDetection` is set to 1.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceQualityInfo *FaceQualityInfo `json:"FaceQualityInfo,omitempty" name:"FaceQualityInfo"`
}

type FaceQualityCompleteness struct {

	// Eyebrow completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,80], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Eyebrow *int64 `json:"Eyebrow,omitempty" name:"Eyebrow"`

	// Eye completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,80], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Eye *int64 `json:"Eye,omitempty" name:"Eye"`

	// Nose completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,60], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Nose *int64 `json:"Nose,omitempty" name:"Nose"`

	// Cheek completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,70], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cheek *int64 `json:"Cheek,omitempty" name:"Cheek"`

	// Mouth completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,50], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mouth *int64 `json:"Mouth,omitempty" name:"Mouth"`

	// Chin completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,70], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Chin *int64 `json:"Chin,omitempty" name:"Chin"`
}

type FaceQualityInfo struct {

	// Quality score. Value range: [0,100]. It comprehensively evaluates whether the image quality is suitable for face recognition; the higher the score, the higher the quality. 
	// In normal cases, you only need to use `Score` as the overall quality standard score. Specific item scores such as `Sharpness`, `Brightness`, `Completeness` are for reference only.
	// Reference range: [0,40]: poor; [40,60]: fine; [60,80]: good; [80,100]: excellent. 
	// We recommend selecting images with a score above 70 for adding faces.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// Sharpness. Value range: [0,100]. It evaluates the sharpness of the image. The higher the score, the sharper the image. 
	// Reference range: [0,40]: very blurry; [40,60]: blurry; [60,80]: fine; [80,100]: sharp. 
	// We recommend selecting images with a score above 80 for adding faces.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Sharpness *int64 `json:"Sharpness,omitempty" name:"Sharpness"`

	// Brightness. Value range: [0,100]. The brighter the image, the higher the score. 
	// Reference range: [0,30]: dark; [30,70]: normal; [70,100]: bright. 
	// We recommend selecting images in the [30,70] range for adding faces.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Brightness *int64 `json:"Brightness,omitempty" name:"Brightness"`

	// Completeness of facial features, which assesses the completeness of the eyebrows, eyes, nose, cheeks, mouth, and chin.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Completeness *FaceQualityCompleteness `json:"Completeness,omitempty" name:"Completeness"`
}

type FaceRect struct {

	// Horizontal coordinate of the top-left vertex of face frame. 
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the completeness score meets the requirement.
	X *int64 `json:"X,omitempty" name:"X"`

	// Vertical coordinate of the top-left vertex of face frame. 
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the completeness score meets the requirement.
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// Face width
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Face height
	Height *uint64 `json:"Height,omitempty" name:"Height"`
}

type FaceShape struct {

	// 21 points that describe the face contour.
	FaceProfile []*Point `json:"FaceProfile,omitempty" name:"FaceProfile" list`

	// 8 points that describe the left eye.
	LeftEye []*Point `json:"LeftEye,omitempty" name:"LeftEye" list`

	// 8 points that describe the right eye.
	RightEye []*Point `json:"RightEye,omitempty" name:"RightEye" list`

	// 8 points that describe the left eyebrow.
	LeftEyeBrow []*Point `json:"LeftEyeBrow,omitempty" name:"LeftEyeBrow" list`

	// 8 points that describe the right eyebrow.
	RightEyeBrow []*Point `json:"RightEyeBrow,omitempty" name:"RightEyeBrow" list`

	// 22 points that describe the mouth.
	Mouth []*Point `json:"Mouth,omitempty" name:"Mouth" list`

	// 13 points that describe the nose.
	Nose []*Point `json:"Nose,omitempty" name:"Nose" list`

	// 1 point that describes the left pupil.
	LeftPupil []*Point `json:"LeftPupil,omitempty" name:"LeftPupil" list`

	// 1 point that describes the right pupil.
	RightPupil []*Point `json:"RightPupil,omitempty" name:"RightPupil" list`
}

type GetGroupInfoRequest struct {
	*tchttp.BaseRequest

	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("GetGroupInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Group name
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// Group ID
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// Custom group description field
		GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions" list`

		// Group remarks
		Tag *string `json:"Tag,omitempty" name:"Tag"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// Group creation time and date (`CreationTimestamp`), whose value is the number of milliseconds between the UNIX epoch time and the group creation time.
		CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest

	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("GetGroupListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned group information
		GroupInfos []*GroupInfo `json:"GroupInfos,omitempty" name:"GroupInfos" list`

		// Total number of groups
	// Note: this field may return null, indicating that no valid values can be obtained.
		GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPersonBaseInfoRequest struct {
	*tchttp.BaseRequest

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return errors.New("GetPersonBaseInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetPersonBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Person name
		PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

		// Person gender. 0: empty; 1: male; 2: female.
		Gender *int64 `json:"Gender,omitempty" name:"Gender"`

		// List of the IDs of included faces
		FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPersonGroupInfoRequest struct {
	*tchttp.BaseRequest

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetPersonGroupInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of groups containing this person and their description fields
		PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitempty" name:"PersonGroupInfos" list`

		// Total number of groups
	// Note: this field may return null, indicating that no valid values can be obtained.
		GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

		// Algorithm model version used by the Face Recognition service.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPersonListNumRequest struct {
	*tchttp.BaseRequest

	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("GetPersonListNumRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetPersonListNumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of persons
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// Number of faces
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPersonListRequest struct {
	*tchttp.BaseRequest

	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetPersonListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetPersonListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned person information
		PersonInfos []*PersonInfo `json:"PersonInfos,omitempty" name:"PersonInfos" list`

		// Number of persons in the group
	// Note: this field may return null, indicating that no valid values can be obtained.
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// Number of faces in the group
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// Algorithm model version used for face recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupCandidate struct {

	// Group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Most matching candidate recognized
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates" list`
}

type GroupExDescriptionInfo struct {

	// Custom group description field index, whose value starts from 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupExDescriptionIndex *uint64 `json:"GroupExDescriptionIndex,omitempty" name:"GroupExDescriptionIndex"`

	// Content of the custom group description field to be updated
	GroupExDescription *string `json:"GroupExDescription,omitempty" name:"GroupExDescription"`
}

type GroupInfo struct {

	// Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Custom group description field
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions" list`

	// Group remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// Algorithm model version used for face recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Group creation time and date (`CreationTimestamp`), whose value is the number of milliseconds between the UNIX epoch time and the group creation time. 
	// The UNIX epoch time is 00:00:00, Thursday, January 1, 1970, Coordinated Universal Time (UTC). For more information, please see the UNIX time document.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest

	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Custom description field of the group to be modified, which is a `key-value` pair.
	GroupExDescriptionInfos []*GroupExDescriptionInfo `json:"GroupExDescriptionInfos,omitempty" name:"GroupExDescriptionInfos" list`

	// Group remarks
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonBaseInfoRequest struct {
	*tchttp.BaseRequest

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Name of the person to be modified
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// Gender of the person to be modified. 1: male; 2: female.
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "PersonName")
	delete(f, "Gender")
	if len(f) > 0 {
		return errors.New("ModifyPersonBaseInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonGroupInfoRequest struct {
	*tchttp.BaseRequest

	// Group ID, which is the `GroupId` in the `CreateGroup` API.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Person ID, which is the `PersonId` in the `CreatePerson` API.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Custom description field of the person to be modified, which is a `key-value` pair.
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifyPersonGroupInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PersonExDescriptionInfo struct {

	// Person description field index, whose value starts from 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PersonExDescriptionIndex *uint64 `json:"PersonExDescriptionIndex,omitempty" name:"PersonExDescriptionIndex"`

	// Content of the person description field to be updated
	PersonExDescription *string `json:"PersonExDescription,omitempty" name:"PersonExDescription"`
}

type PersonGroupInfo struct {

	// ID of the group that contains this person
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Content of person description field
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitempty" name:"PersonExDescriptions" list`
}

type PersonInfo struct {

	// Person name
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Person gender
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// Content of person description field
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitempty" name:"PersonExDescriptions" list`

	// List of contained face images
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`

	// Person creation time and date (`CreationTimestamp`), whose value is the number of milliseconds between the UNIX epoch time and the group creation time. 
	// The UNIX epoch time is 00:00:00, Thursday, January 1, 1970, Coordinated Universal Time (UTC). For more information, please see the UNIX time document.
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type Point struct {

	// X coordinate
	X *int64 `json:"X,omitempty" name:"X"`

	// Y coordinate
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type Result struct {

	// Most matching candidate recognized
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates" list`

	// Position of detected face frame
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// Status return code of detected face image. 0: normal. 
	// -1601: the image quality control requirement is not met; in this case, `Candidate` is empty.
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
}

type ResultsReturnsByGroup struct {

	// Position of detected face frame
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// Recognition result.
	GroupCandidates []*GroupCandidate `json:"GroupCandidates,omitempty" name:"GroupCandidates" list`

	// Status return code of detected face image. 0: normal. 
	// -1601: the image quality control requirement is not met; in this case, `Candidate` is empty.
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
}

type SearchFacesRequest struct {
	*tchttp.BaseRequest

	// List of groups to be searched in (up to 100). The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10. 
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized. 
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. Face images whose value is below 34 cannot be recognized. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Number of the most similar persons returned for one single recognized face image. Default value: 5. Maximum value: 100. 
	// For example, if `MaxFaceNum` is 1 and `MaxPersonNum` is 8, information of the top 8 most similar persons will be returned.
	// The greater the value, the longer the processing time. We recommend setting a value below 10.
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is above the `FaceMatchThreshold` value. Default value: 0.
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("SearchFacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchFacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recognition result.
		Results []*Result `json:"Results,omitempty" name:"Results" list`

		// Number of faces included in searched groups.
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFacesReturnsByGroupRequest struct {
	*tchttp.BaseRequest

	// List of groups to be searched in (up to 60). The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Detected faces, which is corresponding to the maximum number of returned most matching persons. Default value: 5. Maximum value: 10.  
	// For example, if `MaxFaceNum` is 3 and `MaxPersonNum` is 5, up to 15 (3 * 5) persons will be returned.
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitempty" name:"MaxPersonNumPerGroup"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is greater than or equal to the `FaceMatchThreshold` value.
	// Default value: 0.
	// Value range: [0.0,100.0).
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("SearchFacesReturnsByGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchFacesReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of faces included in searched groups.
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// Recognition result.
		ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitempty" name:"ResultsReturnsByGroup" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesReturnsByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsRequest struct {
	*tchttp.BaseRequest

	// List of groups to be searched in (up to 100). The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Number of the most similar persons returned for one single recognized face image. Default value: 5. Maximum value: 100.
	// For example, if `MaxFaceNum` is 1 and `MaxPersonNum` is 8, information of the top 8 most similar persons will be returned.
	// The greater the value, the longer the processing time. We recommend setting a value below 10.
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is greater than or equal to the `FaceMatchThreshold` value. Default value: 0. Value range: [0.0,100.0).
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("SearchPersonsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recognition result.
		Results []*Result `json:"Results,omitempty" name:"Results" list`

		// Number of the persons included in searched groups. If the quality of all faces in the input image does not meet the requirement, 0 will be returned.
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// Algorithm model version used for face recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsReturnsByGroupRequest struct {
	*tchttp.BaseRequest

	// List of groups to be searched in (up to 60). The array element value is the `GroupId` in the `CreateGroup` API.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. We recommend setting this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Detected faces, which is corresponding to the maximum number of returned most matching persons. Default value: 5. Maximum value: 10.  
	// For example, if `MaxFaceNum` is 3, `MaxPersonNumPerGroup` is 5, and the `GroupIds` length is 3, up to 45 (3 * 5 * 3) persons will be returned.
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitempty" name:"MaxPersonNumPerGroup"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is above the `FaceMatchThreshold` value. Default value: 0.
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default.
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("SearchPersonsReturnsByGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of the persons included in searched groups. If the quality of all faces in the input image does not meet the requirement, 0 will be returned.
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// Recognition result.
		ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitempty" name:"ResultsReturnsByGroup" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsReturnsByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyFaceRequest struct {
	*tchttp.BaseRequest

	// ID of the person to be verified. For more information on `PersonId`, please see the group management APIs.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("VerifyFaceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VerifyFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Similarity between given face image and `PersonId`. If there are multiple faces under the `PersonId`, the score of the highest similarity will be returned.
	// 
	// The returned similarity score varies by algorithm version.
	// If you need to verify whether the faces in the two images are the same person, then the 0.1%, 0.01%, and 0.001% FARs on v3.0 correspond to scores of 40, 50, and 60, respectively. Generally, if the score is above 50, it can be judged that they are the same person.
	// The 0.1%, 0.01%, and 0.001% FARs on v2.0 correspond to scores of 70, 80, and 90, respectively. Generally, if the score is above 80, it can be judged that they are the same person.
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// Whether the person in the image matches the `PersonId`.
		IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

		// Algorithm model version used for face recognition in the group where the `Person` is, which is set when the group is created. For more information, please see [Algorithm Model Version](https://intl.cloud.tencent.com/document/product/867/40042?from_cn_redirect=1)
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyPersonRequest struct {
	*tchttp.BaseRequest

	// ID of the person to be verified. For more information on `PersonId`, please see the group management APIs.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Base64-encoded data of the image.
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL 
	// The long side cannot exceed 4,000 px for images in JPG format or 2,000 px for images in other formats.
	//  Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// We recommend storing the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// PNG, JPG, JPEG, and BMP images are supported, while GIF images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of the following problems: excessive brightness, excessive dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high-quality requirement. The image has one to two of the following problems: excessive brightness, excessive dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high-quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("VerifyPersonRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VerifyPersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Similarity between given face image and `PersonId`. If there are multiple faces under the `PersonId`, their information will be fused for verification.
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// Whether the person in the image matches the `PersonId`.
		IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

		// Algorithm model version used for face recognition in the group where the `Person` is, which is set when the group is created. For more information, please see [Algorithm Model Version](https://intl.cloud.tencent.com/document/product/867/40042?from_cn_redirect=1)
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
