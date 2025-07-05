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

package v20201229

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Device struct {
	// This field indicates the IP address of the business user's device and supports recording both **IPv4 and IPv6** addresses. It needs to be used together with the `IpType` parameter.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// This field indicates the MAC address of the business user, which makes it easier to identify and manage devices. Its format and value are the same as those of a standard MAC address.
	Mac *string `json:"Mac,omitnil,omitempty" name:"Mac"`

	// *In beta test. Stay tuned.*
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// *In beta test. Stay tuned.*
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// This field indicates the international mobile equipment identity **(IMEI)** number of the business user's device, which can be used to identify each mobile communication device such as mobile phone for easier device identification and management.<br>Note: the format is **15–17 digits**.
	IMEI *string `json:"IMEI,omitnil,omitempty" name:"IMEI"`

	// **For iOS devices**: this field indicates the identifier for advertisers **(IDFA)** of the business user, which is provided by Apple to identify the user and contains a hexadecimal string of 32 digits and letters.<br>
	// Note: as iOS 14 has been updated by Apple to allow users to manually enable or disable IDFA since 2021, the validity of this string may be reduced.
	IDFA *string `json:"IDFA,omitnil,omitempty" name:"IDFA"`

	// **For iOS devices**: this field indicates the identifier for vendors **(IDFV)** of the business user, which is provided by Apple to identify the app vendor and contains a hexadecimal string of 32 digits and letters. It can be used to uniquely identify a device.
	IDFV *string `json:"IDFV,omitnil,omitempty" name:"IDFV"`

	// This field indicates the type of the recorded IP address. Valid values: **0** (IPv4 address), **1** (IPv6 address). It needs to be used together with the `IpType` parameter.
	IpType *uint64 `json:"IpType,omitnil,omitempty" name:"IpType"`
}

// Predefined struct for user
type ImageModerationRequestParams struct {
	// This field indicates the specific number of the policy, which is used for API scheduling and can be configured in the CMS console. If the `Biztype` parameter is passed in, a moderation policy will be used based on the business scenario; otherwise, the default moderation policy will be used.<br>Note: `Biztype` can contain 3-32 digits, letters, and underscores; different `Biztype` values are associated with different business scenarios and moderation policies, so you need to verify the `Biztype` before calling this API.
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// This field indicates the data ID assigned by you to the object to be detected for easier file identification and management.<br>It **can contain up to 64 letters, digits, and special symbols (_-@#)**.
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// This field indicates the Base64 encoding of the image to be detected. The image **size cannot exceed 5 MB**. **A resolution of 256x256 or higher is recommended**; otherwise, the recognition effect may be affected.<br>Note: **you must enter a value for either this field or `FileUrl`**.
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// URL of the image to moderate. It supports PNG, JPG, JPEG, BMP, GIF AND WEBP files. The file **cannot exceed 5 MB** and the resolution should not below **256*246**. The default timeout period is 3 seconds. Note that **redirection URLs may be blocked by security policies**. In this case, an error message will return. For example, if an HTTP request gets the 302 code, the error `ResourceUnavailable.ImageDownloadError` is returned. <br>**Either `FileUrl` or `FileContent` must be specified. 
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// **For GIF/long image detection only**. This field indicates the GIF frame capturing frequency (the image interval for capturing a frame for detection). For long images, you should round the width:height ratio as the total number of images to be split. The default value is 0, where only the first frame of the GIF image will be detected, and the long image will not be split.<br>Note: the `Interval` and `MaxFrames` parameters need to be used in combination; for example, if `Interval` is `3` and `MaxFrames` is `400`, the GIF/long image will be detected once every two frames for up to 400 frames.
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// **For GIF/long image detection only**. This field indicates the maximum number of frames that can be captured. The default value is 1, where only the first frame of the input GIF image will be detected, and the long image will not be split (which may cause a processing timeout).<br>Note: the `Interval` and `MaxFrames` parameters need to be used in combination; for example, if `Interval` is `3` and `MaxFrames` is `400`, the GIF/long image will be detected once every two frames for up to 400 frames.
	MaxFrames *int64 `json:"MaxFrames,omitnil,omitempty" name:"MaxFrames"`

	// This field indicates the information of the user that corresponds to the object to be detected. It can be passed in to identify the user involved in the violation.
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// This field indicates the information of the device that corresponds to the object to be detected. It can be passed in to identify the device involved in the violation.
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`
}

type ImageModerationRequest struct {
	*tchttp.BaseRequest
	
	// This field indicates the specific number of the policy, which is used for API scheduling and can be configured in the CMS console. If the `Biztype` parameter is passed in, a moderation policy will be used based on the business scenario; otherwise, the default moderation policy will be used.<br>Note: `Biztype` can contain 3-32 digits, letters, and underscores; different `Biztype` values are associated with different business scenarios and moderation policies, so you need to verify the `Biztype` before calling this API.
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// This field indicates the data ID assigned by you to the object to be detected for easier file identification and management.<br>It **can contain up to 64 letters, digits, and special symbols (_-@#)**.
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// This field indicates the Base64 encoding of the image to be detected. The image **size cannot exceed 5 MB**. **A resolution of 256x256 or higher is recommended**; otherwise, the recognition effect may be affected.<br>Note: **you must enter a value for either this field or `FileUrl`**.
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// URL of the image to moderate. It supports PNG, JPG, JPEG, BMP, GIF AND WEBP files. The file **cannot exceed 5 MB** and the resolution should not below **256*246**. The default timeout period is 3 seconds. Note that **redirection URLs may be blocked by security policies**. In this case, an error message will return. For example, if an HTTP request gets the 302 code, the error `ResourceUnavailable.ImageDownloadError` is returned. <br>**Either `FileUrl` or `FileContent` must be specified. 
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// **For GIF/long image detection only**. This field indicates the GIF frame capturing frequency (the image interval for capturing a frame for detection). For long images, you should round the width:height ratio as the total number of images to be split. The default value is 0, where only the first frame of the GIF image will be detected, and the long image will not be split.<br>Note: the `Interval` and `MaxFrames` parameters need to be used in combination; for example, if `Interval` is `3` and `MaxFrames` is `400`, the GIF/long image will be detected once every two frames for up to 400 frames.
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// **For GIF/long image detection only**. This field indicates the maximum number of frames that can be captured. The default value is 1, where only the first frame of the input GIF image will be detected, and the long image will not be split (which may cause a processing timeout).<br>Note: the `Interval` and `MaxFrames` parameters need to be used in combination; for example, if `Interval` is `3` and `MaxFrames` is `400`, the GIF/long image will be detected once every two frames for up to 400 frames.
	MaxFrames *int64 `json:"MaxFrames,omitnil,omitempty" name:"MaxFrames"`

	// This field indicates the information of the user that corresponds to the object to be detected. It can be passed in to identify the user involved in the violation.
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// This field indicates the information of the device that corresponds to the object to be detected. It can be passed in to identify the device involved in the violation.
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`
}

func (r *ImageModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageModerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizType")
	delete(f, "DataId")
	delete(f, "FileContent")
	delete(f, "FileUrl")
	delete(f, "Interval")
	delete(f, "MaxFrames")
	delete(f, "User")
	delete(f, "Device")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageModerationResponseParams struct {
	// This field is used to return the operation suggestion for the `Label` tag. When you get the determination result, the returned value indicates the operation suggested by the system. We recommend you handle different types of violations and suggestions according to your business needs. <br>Returned values: **Block**, **Review**, **Pass**.
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// This field is used to return the **maliciousness tag with the highest priority** in the detection result (LabelResults), which represents the moderation result suggested by the model. We recommend you handle different types of violations and suggestions according to your business needs. <br>Returned values: **Normal**: normal; **Porn**: pornographic; **Abuse**: abusive; **Ad**: advertising; **Custom**: custom type of non-compliant content and other offensive, unsafe, or inappropriate types of content.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// This field is used to return the subtag name under the maliciousness tag with the highest priority hit by the detection result, such as *Porn-SexBehavior*. If no subtag is hit, an empty string will be returned.
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// Confidence score of the under the current label. Value range: 0 (**the lowest confidence**) to 100 (**the highest confidence**). For example, *Porn 99* indicates that the image is highly likely to be pornographic, while *Porn 0* indicates that the image is not pornographic.
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// This field is used to return the detailed recognition result for the maliciousness tag hit by the categorization model, such as porn, advertising, or any other offensive, unsafe, or inappropriate type of content.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LabelResults []*LabelResult `json:"LabelResults,omitnil,omitempty" name:"LabelResults"`

	// This field is used to return the detailed detection result of the object detection model, including the tag name hit by the content such as object, advertising logo, or QR code, tag score, coordinate information, scenario recognition result, and operation suggestion. For more information on the returned value, see the description of the `ObjectResults` data structure.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ObjectResults []*ObjectResult `json:"ObjectResults,omitnil,omitempty" name:"ObjectResults"`

	// This field is used to return the detailed text OCR result, including the text coordinate information, text recognition result, and operation suggestion. For more information on the returned value, see the description of the `OcrResults` data structure.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrResults []*OcrResult `json:"OcrResults,omitnil,omitempty" name:"OcrResults"`

	// This field is used to return the result of recognition based on image risk libraries (blocklist and allowlist). For more information on the returned value, see the description of the `LibResults` data structure.<br>Note: currently, **you cannot customize image risk libraries**.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LibResults []*LibResult `json:"LibResults,omitnil,omitempty" name:"LibResults"`

	// This field is used to return the `DataId` in the request parameters that correspond to the detected object.
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// This field is used to return the `BizType` in the request parameters that correspond to the detected object.
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// This field is used to return the additional information (Extra) configured based on your needs. If it is not configured, an empty value will be returned by default.<br>Note: the returned information varies by customer or `Biztype`. If you need to configure this field, submit a ticket or contact the aftersales service for assistance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// This field is used to return the MD5 checksum of the detected object for easier verification of the file integrity.
	FileMD5 *string `json:"FileMD5,omitnil,omitempty" name:"FileMD5"`

	// Image recognition result, including the hit tags, confidence and location.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RecognitionResults []*RecognitionResult `json:"RecognitionResults,omitnil,omitempty" name:"RecognitionResults"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageModerationResponse struct {
	*tchttp.BaseResponse
	Response *ImageModerationResponseParams `json:"Response"`
}

func (r *ImageModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelDetailItem struct {
	// This field is used to return the ID of the recognized object for easier recognition and distinction.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// This field is used to return the hit subtag name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// This field is used to return the hit score of the subtag. Value range: **0–100**; for example, *Porn-SexBehavior 99* indicates that the hit score of the *Porn-SexBehavior* tag for the recognized content is 99.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`
}

type LabelResult struct {
	// This field is used to return the scenario result recognized by the model, such as advertising, pornographic, and harmful.
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// This field is used to return the operation suggestion for the current maliciousness tag. When you get the determination result, the returned value indicates the operation suggested by the system. We recommend you handle different types of violations and suggestions according to your business needs. <br>Returned values: **Block**, **Review**, **Pass**.
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// This field is used to return the maliciousness tag in the detection result.<br>Returned values: **Normal**: normal; **Porn**: pornographic; **Abuse**: abusive; **Ad**: advertising; **Custom**: custom type of non-compliant content and other offensive, unsafe, or inappropriate types of content.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// This field is used to return the detection result for a subtag under the maliciousness tag, such as *Porn-SexBehavior*.
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// Confidence score of the under the current label. Value range: 0 (**the lowest confidence**) to 100 (**the highest confidence**). For example, *Porn 99* indicates that the image is highly likely to be pornographic, while *Porn 0* indicates that the image is not pornographic.
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// This field is used to return the details of the subtag hit by the categorization model, such as number, hit tag name, and score.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Details []*LabelDetailItem `json:"Details,omitnil,omitempty" name:"Details"`
}

type LibDetail struct {
	// This field is used to return the ID of the recognized object for easier recognition and distinction.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// This field is **valid only when `Label` is `Custom` (custom keyword)**. It is used to return the ID of the custom library for easier custom library management and configuration.
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// This field is **valid only when `Label` is `Custom` (custom keyword)**. It is used to return the name of the custom library for easier custom library management and configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// This field is used to return the ID of the recognized image object for easier file management.
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// This field is used to return the maliciousness tag in the detection result.<br>Returned values: **Normal**: normal; **Porn**: pornographic; **Abuse**: abusive; **Ad**: advertising; **Custom**: custom type of non-compliant content and other offensive, unsafe, or inappropriate types of content.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// This field is used to return other custom tags to meet the needs in your customized scenarios. It can be skipped if you have no custom needs.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// This field is used to return the hit score of the model. Value range: **0–100**; for example, *Porn 99* indicates that the hit score of the porn tag for the recognized content is 99.
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`
}

type LibResult struct {
	// This field indicates the scenario recognition result of the model. Default value: Similar.
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// This field is used to return the operation suggestion. When you get the determination result, the returned value indicates the operation suggested by the system. We recommend you handle different types of violations and suggestions according to your business needs. <br>Returned values: **Block**, **Review**, **Pass**.
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// This field is used to return the maliciousness tag in the detection result.<br>Returned values: **Normal**: normal; **Porn**: pornographic; **Abuse**: abusive; **Ad**: advertising; **Custom**: custom type of non-compliant content and other offensive, unsafe, or inappropriate types of content.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// This field is used to return the detection result for a subtag under the maliciousness tag, such as *Porn-SexBehavior*.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// This field is used to return the recognition score of the image search model. Value range: **0–100**. It indicates the score for the similarity between the moderated image **and the samples in the library**. A higher score indicates that the content is more likely to hit a sample in the library of similar images.
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// This field is used to return the detailed result of the comparison with the blocklist/allowlist, such as number, library name, and maliciousness tag. For more information on the returned value, see the description of the [LibDetail](https://intl.cloud.tencent.com/document/product/1125/53274?from_cn_redirect=1#LibDetail) data structure.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Details []*LibDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type Location struct {
	// This parameter is used to return the pixel position of the **abscissa (X) of the top-left corner** of the detection frame. It can be combined with other parameters to uniquely determine the size and position of the detection frame.
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// This parameter is used to return the pixel position of the **ordinate of the top-left corner** (Y) of the detection frame. It can be combined with other parameters to uniquely determine the size and position of the detection frame.
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`

	// This parameter is used to return the **width of the detection frame** (the length starting from the top-left corner and extending to the right on the X axis). It can be combined with other parameters to uniquely determine the size and position of the detection frame.
	Width *float64 `json:"Width,omitnil,omitempty" name:"Width"`

	// This parameter is used to return the **height of the detection frame** (the length starting from the top-left corner and extending down the Y axis). It can be combined with other parameters to uniquely determine the size and position of the detection frame.
	Height *float64 `json:"Height,omitnil,omitempty" name:"Height"`

	// This parameter is used to return the **rotation angle of the detection frame**. Valid values: **0–360** (**degrees**), and the direction is **counterclockwise rotation**. This parameter can be combined with the `X` and `Y` coordinate parameters to uniquely determine the specific position of the detection frame.
	Rotate *float64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`
}

type ObjectDetail struct {
	// This parameter is used to return the ID of the recognized object for easier recognition and distinction.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// This parameter is used to return the hit object tag.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// This parameter is used to return the value or content of the object tag; for example, when the tag is *QR code (QrCode)*, this field will be the URL of the recognized QR code.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// This parameter is used to return the hit score of the object tag. Valid values: **0–100**; for example, *QrCode 99* indicates that it is highly likely that the recognized content will hit the QR code tag.
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// This field is used to return the coordinate position (X and Y coordinates of the top-left corner, length, width, and rotation angle) of the object detection frame for quick location of the object information.
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`

	// This parameter is used to return the hit object subtag.
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`
}

type ObjectResult struct {
	// This field is used to return the recognized object scenario result, such as QR code, logo, and image OCR.
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// This field is used to return the operation suggestion for the current maliciousness tag. When you get the determination result, the returned value indicates the operation suggested by the system. We recommend you handle different types of violations and suggestions according to your business needs. <br>Returned values: **Block**, **Review**, **Pass**.
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// This field is used to return the maliciousness tag in the detection result, which represents the moderation result suggested by the model. We recommend you handle different types of violations and suggestions according to your business needs. <br>Returned values: **Normal**: normal; **Porn**: pornographic; **Abuse**: abusive; **Ad**: advertising; **Custom**: custom type of non-compliant content and other offensive, unsafe, or inappropriate types of content.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// This field is used to return the detection result for a subtag under the current maliciousness tag, such as *Porn-SexBehavior*.
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// This field is used to return the hit score of a subtag under the current maliciousness tag. Value range: **0–100**; for example, *Porn-SexBehavior 99* indicates that the hit score of the *Porn-SexBehavior* tag for the recognized content is 99.
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// This field is used to return the name of the recognized object.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// This field is used to return the details of the recognized object, such as number, hit tag name, and position coordinates. For more information on the returned value, see the description of the [ObjectDetail](https://intl.cloud.tencent.com/document/api/1125/53274?from_cn_redirect=1#ObjectDetail) data structure.
	//  
	// Note: this field may return null, indicating that no valid values can be obtained.
	Details []*ObjectDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type OcrResult struct {
	// This field indicates the recognition scenario. Default value: OCR (image OCR).
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// This field is used to return the operation suggestion for the maliciousness tag with the highest priority. When you get the determination result, the returned value indicates the operation suggested by the system. We recommend you handle different types of violations and suggestions according to your business needs. <br>Returned values: **Block**, **Review**, **Pass**.
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// This field is used to return the maliciousness tag with the highest priority in the OCR detection result, which represents the moderation result suggested by the model. We recommend you handle different types of violations and suggestions according to your business needs. <br>Returned values: **Normal**: normal; **Porn**: pornographic; **Abuse**: abusive; **Ad**: advertising; **Custom**: custom type of non-compliant content and other offensive, unsafe, or inappropriate types of content.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// This field is used to return the detection result for a subtag under the current tag (Label), such as *Porn-SexBehavior*.
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// This field is used to return the confidence under the current tag (Label). Value range: 0 (**the lowest confidence**)–100 (**the highest confidence**), where a higher value indicates that the text is more likely to fall into the category of the current returned tag; for example, *Porn 99* indicates that the text is highly likely to be pornographic, while *Porn 0* indicates that the text is not pornographic.
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// This field is used to return the details of the OCR recognition result, such as text content, tag, and recognition frame position.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Details []*OcrTextDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// This field is used to return the text information recognized by OCR.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type OcrTextDetail struct {
	// This field is used to return the text content recognized by OCR.<br>Note: OCR can recognize text of **up to 5,000 bytes**.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// This field is used to return the maliciousness tag in the detection result.<br>Returned values: **Normal**: normal; **Porn**: pornographic; **Abuse**: abusive; **Ad**: advertising; **Custom**: custom type of non-compliant content and other offensive, unsafe, or inappropriate types of content.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// This field is **valid only when `Label` is `Custom` (custom keyword)**. It is used to return the ID of the custom library for easier custom library management and configuration.
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// This field is **valid only when `Label` is `Custom` (custom keyword)**. It is used to return the name of the custom library for easier custom library management and configuration.
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// This parameter is used to return the hit keyword under the current tag (label).
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// This parameter is used to return the model hit score of the current maliciousness tag. Value range: **0–100**, where a higher value indicates that the current scenario agrees more with the scenario represented by the maliciousness tag.
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// This parameter is used to return the position (X and Y coordinates of the top-left corner, length, width, and rotation angle) of the OCR detection frame in the image for quick location of the recognized text.
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`

	// This parameter is used to return the confidence of the text OCR result. Valid values: **0** (**the lowest confidence**)–**100** (**the highest confidence**), where a higher value indicates that it is more likely that the image contains the recognized text; for example, *Hello 99* indicates that it is highly likely that the text in the OCR recognition frame is "Hello".
	Rate *uint64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// This field is used to return the maliciousness subtag that corresponds to the detection result.
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`
}

type RecognitionResult struct {
	// Value: `Scene`
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// Hit tags under the `Label`
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Tags []*RecognitionTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RecognitionTag struct {
	// Tag name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Confidence score. Value: 1 to 100. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Location information. It returns 0 if there is not location information.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`
}

type User struct {
	// This field indicates the business user ID. After it is specified, the system can optimize the moderation result according to the violation history to facilitate determination when a suspicious violation risk exists.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// This field indicates the nickname of the business user's account.
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// This field indicates the account type of the business user ID.<br>
	// This field can be used together with the ID parameter (UserId) to uniquely identify the account.
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// This field indicates the gender of the business user's account.<br>
	// Valid values: **0** (default value): unknown; **1** (male); **2** (female).
	Gender *uint64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// This field indicates the age of the business user's account.<br>
	// Valid values: integers between **0** (default value, which indicates unknown) and **custom age limit**.
	Age *uint64 `json:"Age,omitnil,omitempty" name:"Age"`

	// This field indicates the level of the business user's account.<br>
	// Valid values: **0** (default value): unknown; **1**: low level; **2**: medium level; **3**: high level. Currently, **the level is not customizable**.
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// This field indicates the mobile number of the business user's account. It supports recording mobile numbers across the world.<br>
	// Note: you need to use a consistent mobile number format, such as area code format (086/+86).
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// This field indicates the profile of the business user. It can contain **up to 5,000 letters and special symbols**.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// This field indicates the access URL of the business user's profile photo in PNG, JPG, JPEG, BMP, GIF, or WEBP format.<br>Note: the profile photo **cannot exceed 5 MB in size**. **A resolution of 256x256 or higher** is recommended. The image download time should be limited to 3 seconds; otherwise, a download timeout will be returned.
	HeadUrl *string `json:"HeadUrl,omitnil,omitempty" name:"HeadUrl"`
}