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

package v20180717

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AIAnalysisTemplateItem struct {

	// Unique ID of intelligent analysis template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Intelligent analysis template name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Intelligent analysis template description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter of intelligent categorization task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitempty" name:"ClassificationConfigure"`

	// Control parameter of intelligent tagging task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitempty" name:"TagConfigure"`

	// Control parameter of intelligent cover generating task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitempty" name:"CoverConfigure"`

	// Control parameter of intelligent frame-specific tagging task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitempty" name:"FrameTagConfigure"`

	// Control parameter of an intelligent highlight generating task.
	HighlightConfigure *HighlightsConfigureInfo `json:"HighlightConfigure,omitempty" name:"HighlightConfigure"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AIRecognitionTemplateItem struct {

	// Unique ID of video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Video content recognition template name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Video content recognition template description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter of opening and closing credits recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeadTailConfigure *HeadTailConfigureInfo `json:"HeadTailConfigure,omitempty" name:"HeadTailConfigure"`

	// Control parameter of splitting recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SegmentConfigure *SegmentConfigureInfo `json:"SegmentConfigure,omitempty" name:"SegmentConfigure"`

	// Face recognition control parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// Full text recognition control parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// Text keyword recognition control parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// Full speech recognition control parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// Speech keyword recognition control parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`

	// Control parameter of object recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ObjectConfigure *ObjectConfigureInfo `json:"ObjectConfigure,omitempty" name:"ObjectConfigure"`

	// Screencapturing interval in seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AdaptiveDynamicStreamingInfoItem struct {

	// Adaptive bitrate streaming specification.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Container format. Valid values: hls, dash.
	Package *string `json:"Package,omitempty" name:"Package"`

	// Encryption type.
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// Playback address.
	Url *string `json:"Url,omitempty" name:"Url"`
}

type AdaptiveDynamicStreamingTaskInput struct {

	// Adaptive bitrate streaming template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of up to 10 image or text watermarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type AdaptiveDynamicStreamingTemplate struct {

	// Unique ID of a transcoding to adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of a transcoding to adaptive bitrate streaming template.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of a transcoding to adaptive bitrate streaming template.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Container type. Valid values:
	// <li>hls; </li>
	// <li>dash.</li>
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// DRM type. Valid values:
	// <li>FairPlay;</li>
	// <li>SimpleAES;</li>
	// <li>Widevine.</li>
	// If this field is a blank string, DRM will not be performed on the video.
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// List of video track templates.
	VideoTrackTemplateSet []*VideoTrackTemplateInfo `json:"VideoTrackTemplateSet,omitempty" name:"VideoTrackTemplateSet" list`

	// List of audio track templates.
	AudioTrackTemplateSet []*AudioTrackTemplateInfo `json:"AudioTrackTemplateSet,omitempty" name:"AudioTrackTemplateSet" list`

	// Adaptive bitstream format. Valid value:
	// <li>HLS.</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Parameter information of input stream for adaptive bitrate streaming. Up to 10 streams can be input.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitempty" name:"StreamInfos" list`

	// Whether to prohibit transcoding from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitempty" name:"DisableHigherVideoResolution"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AdaptiveStreamTemplate struct {

	// Video parameter information.
	Video *VideoTemplateInfo `json:"Video,omitempty" name:"Video"`

	// Audio parameter information.
	Audio *AudioTemplateInfo `json:"Audio,omitempty" name:"Audio"`

	// Whether to remove audio stream. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	RemoveAudio *uint64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`
}

type AiAnalysisResult struct {

	// Task type. Valid values:
	// <li>Classification: intelligent categorization</li>
	// <li>Cover: intelligent cover generating</li>
	// <li>Tag: intelligent tagging</li>
	// <li>FrameTag: intelligent frame tagging</li>
	// <li>Highlight: intelligent highlight generating</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Query result of intelligent categorization task in video content analysis, which is valid if task type is `Classification`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassificationTask *AiAnalysisTaskClassificationResult `json:"ClassificationTask,omitempty" name:"ClassificationTask"`

	// Query result of intelligent cover generating task in video content analysis, which is valid if task type is `Cover`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoverTask *AiAnalysisTaskCoverResult `json:"CoverTask,omitempty" name:"CoverTask"`

	// Query result of intelligent tagging task in video content analysis, which is valid if task type is `Tag`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagTask *AiAnalysisTaskTagResult `json:"TagTask,omitempty" name:"TagTask"`

	// Query result of intelligent frame-specific tagging task in video content analysis, which is valid if task type is `FrameTag`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FrameTagTask *AiAnalysisTaskFrameTagResult `json:"FrameTagTask,omitempty" name:"FrameTagTask"`

	// Query result of an intelligent highlight generating task in video content analysis, which is valid when task type is `Highlight`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HighlightTask *AiAnalysisTaskHighlightResult `json:"HighlightTask,omitempty" name:"HighlightTask"`
}

type AiAnalysisTaskClassificationInput struct {

	// Intelligent video categorization template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskClassificationOutput struct {

	// List of intelligently generated video categories.
	ClassificationSet []*MediaAiAnalysisClassificationItem `json:"ClassificationSet,omitempty" name:"ClassificationSet" list`
}

type AiAnalysisTaskClassificationResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of intelligent categorization task.
	Input *AiAnalysisTaskClassificationInput `json:"Input,omitempty" name:"Input"`

	// Output of intelligent categorization task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskClassificationOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskCoverInput struct {

	// Intelligent video cover generating template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskCoverOutput struct {

	// List of intelligently generated covers.
	CoverSet []*MediaAiAnalysisCoverItem `json:"CoverSet,omitempty" name:"CoverSet" list`
}

type AiAnalysisTaskCoverResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of intelligent cover generating task.
	Input *AiAnalysisTaskCoverInput `json:"Input,omitempty" name:"Input"`

	// Output of intelligent cover generating task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskCoverOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskFrameTagInput struct {

	// Intelligent frame-specific video tagging template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskFrameTagOutput struct {

	// List of frame-specific video tags.
	SegmentSet []*MediaAiAnalysisFrameTagSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiAnalysisTaskFrameTagResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of intelligent frame-specific tagging task.
	Input *AiAnalysisTaskFrameTagInput `json:"Input,omitempty" name:"Input"`

	// Output of intelligent frame-specific tagging task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskFrameTagOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskHighlightInput struct {

	// ID of an intelligent highlight generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskHighlightOutput struct {

	// List of intelligently generated highlights.
	HighlightSet []*MediaAiAnalysisHighlightItem `json:"HighlightSet,omitempty" name:"HighlightSet" list`
}

type AiAnalysisTaskHighlightResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for an intelligent highlight generating task.
	Input *AiAnalysisTaskHighlightInput `json:"Input,omitempty" name:"Input"`

	// Output of an intelligent highlight generating task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskHighlightOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskInput struct {

	// Video content analysis template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskTagInput struct {

	// Intelligent video tagging template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskTagOutput struct {

	// List of intelligently generated video tags.
	TagSet []*MediaAiAnalysisTagItem `json:"TagSet,omitempty" name:"TagSet" list`
}

type AiAnalysisTaskTagResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of intelligent tagging task.
	Input *AiAnalysisTaskTagInput `json:"Input,omitempty" name:"Input"`

	// Output of intelligent tagging task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskTagOutput `json:"Output,omitempty" name:"Output"`
}

type AiContentReviewResult struct {

	// Task type. Valid values:
	// <li>Porn: porn information detection in image</li>
	// <li>Terrorism: terrorism information detection in image</li>
	// <li>Political: politically sensitive information detection in image</li>
	// <li>Porn.Asr: ASR-based porn information detection in speech</li>
	// <li>Porn.Ocr: OCR-based porn information detection in text</li>
	// <li>Political.Asr: ASR-based politically sensitive information detection in speech</li>
	// <li>Political.Ocr: OCR-based politically sensitive information detection in text</li>
	// <li>Terrorism.Ocr: OCR-based terrorism information in text</li>
	// <li>Prohibited.Asr: ASR-based prohibited information detection in speech</li>
	// <li>Prohibited.Ocr: OCR-based prohibited information detection in text</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Query result of intelligent porn information detection in video image task in video content audit, which is valid when task type is `Porn`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PornTask *AiReviewTaskPornResult `json:"PornTask,omitempty" name:"PornTask"`

	// Query result of intelligent terrorism information detection in video image task in video content audit, which is valid when task type is `Terrorism`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TerrorismTask *AiReviewTaskTerrorismResult `json:"TerrorismTask,omitempty" name:"TerrorismTask"`

	// Query result of intelligent politically sensitive information detection in video image task in video content audit, which is valid when task type is `Political`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PoliticalTask *AiReviewTaskPoliticalResult `json:"PoliticalTask,omitempty" name:"PoliticalTask"`

	// Query result of ASR-based porn information detection in speech task in video content audit, which is valid when task type is `Porn.Asr`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PornAsrTask *AiReviewTaskPornAsrResult `json:"PornAsrTask,omitempty" name:"PornAsrTask"`

	// Query result of OCR-based porn information detection in text task in video content audit, which is valid when task type is `Porn.Ocr`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PornOcrTask *AiReviewTaskPornOcrResult `json:"PornOcrTask,omitempty" name:"PornOcrTask"`

	// Query result of ASR-based politically sensitive information detection in speech task in video content audit, which is valid when task type is `Political.Asr`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PoliticalAsrTask *AiReviewTaskPoliticalAsrResult `json:"PoliticalAsrTask,omitempty" name:"PoliticalAsrTask"`

	// Query result of OCR-based politically sensitive information detection in text task in video content audit, which is valid when task type is `Political.Ocr`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PoliticalOcrTask *AiReviewTaskPoliticalOcrResult `json:"PoliticalOcrTask,omitempty" name:"PoliticalOcrTask"`

	// Query result of OCR-based terrorism information detection in text task in video content audit, which is valid when task type is `Terrorism.Ocr`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TerrorismOcrTask *AiReviewTaskTerrorismOcrResult `json:"TerrorismOcrTask,omitempty" name:"TerrorismOcrTask"`

	// Query result of ASR-based prohibited information detection in speech task in video content audit, which is valid when task type is `Prohibited.Asr`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProhibitedAsrTask *AiReviewTaskProhibitedAsrResult `json:"ProhibitedAsrTask,omitempty" name:"ProhibitedAsrTask"`

	// Query result of OCR-based prohibited information detection in text task in video content audit, which is valid when task type is `Prohibited.Ocr`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProhibitedOcrTask *AiReviewTaskProhibitedOcrResult `json:"ProhibitedOcrTask,omitempty" name:"ProhibitedOcrTask"`
}

type AiContentReviewTaskInput struct {

	// Video content audit template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionResult struct {

	// Task type. Valid values:
	// <li>FaceRecognition: face recognition,</li>
	// <li>AsrWordsRecognition: speech keyword recognition,</li>
	// <li>OcrWordsRecognition: text keyword recognition,</li>
	// <li>AsrFullTextRecognition: full speech recognition,</li>
	// <li>OcrFullTextRecognition: full text recognition,</li>
	// <li>HeadTailRecognition: video opening and ending credits recognition,</li>
	// <li>ObjectRecognition: object recognition.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Video opening and ending credits recognition result, which is valid when `Type` is
	//  `HeadTailRecognition`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeadTailTask *AiRecognitionTaskHeadTailResult `json:"HeadTailTask,omitempty" name:"HeadTailTask"`

	// Video splitting recognition result, which is valid when `Type` is
	//  `SegmentRecognition`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SegmentTask *AiRecognitionTaskSegmentResult `json:"SegmentTask,omitempty" name:"SegmentTask"`

	// Face recognition result, which is valid when `Type` is 
	//  `FaceRecognition`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceTask *AiRecognitionTaskFaceResult `json:"FaceTask,omitempty" name:"FaceTask"`

	// Speech keyword recognition result, which is valid when `Type` is
	//  `AsrWordsRecognition`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AsrWordsTask *AiRecognitionTaskAsrWordsResult `json:"AsrWordsTask,omitempty" name:"AsrWordsTask"`

	// Full speech recognition result, which is valid when `Type` is
	//  `AsrFullTextRecognition`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AsrFullTextTask *AiRecognitionTaskAsrFullTextResult `json:"AsrFullTextTask,omitempty" name:"AsrFullTextTask"`

	// Text keyword recognition result, which is valid when `Type` is
	//  `OcrWordsRecognition`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrWordsTask *AiRecognitionTaskOcrWordsResult `json:"OcrWordsTask,omitempty" name:"OcrWordsTask"`

	// Full text recognition result, which is valid when `Type` is
	//  `OcrFullTextRecognition`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrFullTextTask *AiRecognitionTaskOcrFullTextResult `json:"OcrFullTextTask,omitempty" name:"OcrFullTextTask"`

	// Object recognition result, which is valid when `Type` is
	//  `ObjectRecognition`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ObjectTask *AiRecognitionTaskObjectResult `json:"ObjectTask,omitempty" name:"ObjectTask"`
}

type AiRecognitionTaskAsrFullTextResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of full speech recognition task.
	Input *AiRecognitionTaskAsrFullTextResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of full speech recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskAsrFullTextResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskAsrFullTextResultInput struct {

	// Full speech recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrFullTextResultOutput struct {

	// List of full speech recognition segments.
	SegmentSet []*AiRecognitionTaskAsrFullTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`

	// Subtitles file URL.
	SubtitleUrl *string `json:"SubtitleUrl,omitempty" name:"SubtitleUrl"`
}

type AiRecognitionTaskAsrFullTextSegmentItem struct {

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Recognized text.
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AiRecognitionTaskAsrWordsResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of speech keyword recognition task.
	Input *AiRecognitionTaskAsrWordsResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of speech keyword recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskAsrWordsResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskAsrWordsResultInput struct {

	// Speech keyword recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrWordsResultItem struct {

	// Speech keyword.
	Word *string `json:"Word,omitempty" name:"Word"`

	// List of time segments that contain the speech keyword.
	SegmentSet []*AiRecognitionTaskAsrWordsSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskAsrWordsResultOutput struct {

	// Speech keyword recognition result set.
	ResultSet []*AiRecognitionTaskAsrWordsResultItem `json:"ResultSet,omitempty" name:"ResultSet" list`
}

type AiRecognitionTaskAsrWordsSegmentItem struct {

	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type AiRecognitionTaskFaceResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of face recognition task.
	Input *AiRecognitionTaskFaceResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of face recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskFaceResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskFaceResultInput struct {

	// Face recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskFaceResultItem struct {

	// Unique ID of figure.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Figure library type, indicating to which figure library the recognized figure belongs:
	// <li>Default: default figure library;</li>
	// <li>UserDefine: custom figure library.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Figure name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Result set of segments that contain a figure.
	SegmentSet []*AiRecognitionTaskFaceSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskFaceResultOutput struct {

	// Intelligent face recognition result set.
	ResultSet []*AiRecognitionTaskFaceResultItem `json:"ResultSet,omitempty" name:"ResultSet" list`
}

type AiRecognitionTaskFaceSegmentItem struct {

	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`
}

type AiRecognitionTaskHeadTailResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of video opening and ending credits recognition task.
	Input *AiRecognitionTaskHeadTailResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of video opening and ending credits recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskHeadTailResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskHeadTailResultInput struct {

	// Video opening and ending credits recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskHeadTailResultOutput struct {

	// Confidence of recognized opening credits. Value range: 0–100.
	HeadConfidence *float64 `json:"HeadConfidence,omitempty" name:"HeadConfidence"`

	// End time point of video opening credits in seconds.
	HeadTimeOffset *float64 `json:"HeadTimeOffset,omitempty" name:"HeadTimeOffset"`

	// Confidence of recognized closing credits. Value range: 0–100.
	TailConfidence *float64 `json:"TailConfidence,omitempty" name:"TailConfidence"`

	// Start time point of video closing credits in seconds.
	TailTimeOffset *float64 `json:"TailTimeOffset,omitempty" name:"TailTimeOffset"`
}

type AiRecognitionTaskInput struct {

	// Intelligent video recognition template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskObjectResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of object recognition task.
	Input *AiRecognitionTaskObjectResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of object recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskObjectResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskObjectResultInput struct {

	// Object recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskObjectResultItem struct {

	// Name of recognized object.
	Name *string `json:"Name,omitempty" name:"Name"`

	// List of segments that contain an object.
	SegmentSet []*AiRecognitionTaskObjectSeqmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskObjectResultOutput struct {

	// Result set of intelligent object recognition.
	ResultSet []*AiRecognitionTaskObjectResultItem `json:"ResultSet,omitempty" name:"ResultSet" list`
}

type AiRecognitionTaskObjectSeqmentItem struct {

	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`
}

type AiRecognitionTaskOcrFullTextResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of full text recognition task.
	Input *AiRecognitionTaskOcrFullTextResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of full text recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskOcrFullTextResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskOcrFullTextResultInput struct {

	// Full text recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrFullTextResultOutput struct {

	// Full text recognition result set.
	SegmentSet []*AiRecognitionTaskOcrFullTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskOcrFullTextSegmentItem struct {

	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Recognition segment result set.
	TextSet []*AiRecognitionTaskOcrFullTextSegmentTextItem `json:"TextSet,omitempty" name:"TextSet" list`
}

type AiRecognitionTaskOcrFullTextSegmentTextItem struct {

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`

	// Recognized text.
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AiRecognitionTaskOcrWordsResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of text keyword recognition task.
	Input *AiRecognitionTaskOcrWordsResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of text keyword recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskOcrWordsResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskOcrWordsResultInput struct {

	// Text keyword recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrWordsResultItem struct {

	// Text keyword.
	Word *string `json:"Word,omitempty" name:"Word"`

	// List of segments that contain a text keyword.
	SegmentSet []*AiRecognitionTaskOcrWordsSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskOcrWordsResultOutput struct {

	// Text keyword recognition result set.
	ResultSet []*AiRecognitionTaskOcrWordsResultItem `json:"ResultSet,omitempty" name:"ResultSet" list`
}

type AiRecognitionTaskOcrWordsSegmentItem struct {

	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`
}

type AiRecognitionTaskSegmentResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of video splitting task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Input *AiRecognitionTaskSegmentResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of video splitting task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskSegmentResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskSegmentResultInput struct {

	// Video splitting template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskSegmentResultOutput struct {

	// List of split video segments.
	SegmentSet []*AiRecognitionTaskSegmentSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskSegmentSegmentItem struct {

	// File ID, which is valid only when a VOD file is processed and the subsegments generated through segmentation are also VOD files.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Split video segment URL.
	SegmentUrl *string `json:"SegmentUrl,omitempty" name:"SegmentUrl"`

	// Confidence of split segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Start time offset of split segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of split segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Split cover image URL.
	CovImgUrl *string `json:"CovImgUrl,omitempty" name:"CovImgUrl"`

	// Special field, which should be ignored.
	SpecialInfo *string `json:"SpecialInfo,omitempty" name:"SpecialInfo"`
}

type AiReviewPoliticalAsrTaskInput struct {

	// Politically sensitive information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalAsrTaskOutput struct {

	// Score of ASR-detected politically sensitive information in speech between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for ASR-detected politically sensitive information in speech. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain ASR-detected politically sensitive information in speech.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPoliticalOcrTaskInput struct {

	// Politically sensitive information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalOcrTaskOutput struct {

	// Score of OCR-detected politically sensitive information in text between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for OCR-detected politically sensitive information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain OCR-detected politically sensitive information in text.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPoliticalTaskInput struct {

	// Politically sensitive information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalTaskOutput struct {

	// Score of detected politically sensitive information in video between 0 and 100.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for detected politically sensitive information. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tag of detected politically sensitive information in video. Valid values:
	// <li>politician: politically sensitive figure.</li>
	// <li>violation_photo: violating photo.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Label *string `json:"Label,omitempty" name:"Label"`

	// List of video segments that contain the detected politically sensitive information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SegmentSet []*MediaContentReviewPoliticalSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPornAsrTaskInput struct {

	// Porn information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornAsrTaskOutput struct {

	// Score of ASR-detected porn information in speech between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for ASR-detected porn information in speech. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain the ASR-detected porn information in speech.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPornOcrTaskInput struct {

	// Porn information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornOcrTaskOutput struct {

	// Score of OCR-detected porn information in text between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for OCR-detected porn information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain the OCR-detected porn information in text.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPornTaskInput struct {

	// Porn information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornTaskOutput struct {

	// Score of detected porn information in video between 0 and 100.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for detected porn information. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tag of detected porn information in video. Valid values:
	// <li>porn: porn.</li>
	// <li>sexy: sexiness.</li>
	// <li>vulgar: vulgarity.</li>
	// <li>intimacy: intimacy.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Label *string `json:"Label,omitempty" name:"Label"`

	// List of video segments that contain the detected porn information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewProhibitedAsrTaskInput struct {

	// Prohibited information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewProhibitedAsrTaskOutput struct {

	// Score of ASR-detected prohibited information in speech between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for ASR-detected prohibited information in speech. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain the ASR-detected prohibited information in speech.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewProhibitedOcrTaskInput struct {

	// Prohibited information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewProhibitedOcrTaskOutput struct {

	// Score of OCR-detected prohibited information in text between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for OCR-detected prohibited information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain the OCR-detected prohibited information in text.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewTaskPoliticalAsrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of ASR-based politically sensitive information detection in speech task in content audit.
	Input *AiReviewPoliticalAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of ASR-based politically sensitive information detection in speech task in content audit.
	Output *AiReviewPoliticalAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalOcrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of OCR-based politically sensitive information detection in text task in content audit.
	Input *AiReviewPoliticalOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of OCR-based politically sensitive information detection in text task in content audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPoliticalOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of politically sensitive information detection task in content audit.
	Input *AiReviewPoliticalTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of politically sensitive information detection task in content audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPoliticalTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornAsrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of ASR-based porn information detection in speech task in content audit.
	Input *AiReviewPornAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of ASR-based porn information detection in speech task in content audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPornAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornOcrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of OCR-based porn information detection in text task in content audit.
	Input *AiReviewPornOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of OCR-based porn information detection in text task in content audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPornOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of porn information detection task in content audit.
	Input *AiReviewPornTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of porn information detection task in content audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPornTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskProhibitedAsrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of ASR-based prohibited information detection in speech task in content audit
	Input *AiReviewProhibitedAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of ASR-based prohibited information detection in speech task in content audit
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiReviewProhibitedAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskProhibitedOcrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of OCR-based prohibited information detection in text task in content audit
	Input *AiReviewProhibitedOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of OCR-based prohibited information detection in text task in content audit
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiReviewProhibitedOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismOcrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of OCR-based terrorism information detection in text task in content audit.
	Input *AiReviewTerrorismOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of OCR-based terrorism information detection in text task in content audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiReviewTerrorismOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of terrorism information detection task in content audit.
	Input *AiReviewTerrorismTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of terrorism information detection task in content audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AiReviewTerrorismTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTerrorismOcrTaskInput struct {

	// Terrorism information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewTerrorismOcrTaskOutput struct {

	// Score of OCR-detected terrorism information in text between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for OCR-detected terrorism information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain OCR-detected terrorism information in text.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewTerrorismTaskInput struct {

	// Terrorism information detection template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewTerrorismTaskOutput struct {

	// Score of detected terrorism information in a video between 0 and 100.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for detected terrorism information. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tag of detected terrorism information in a video. Valid values:
	// <li>guns: weapons and guns.</li>
	// <li>crowd: crowd.</li>
	// <li>police: police force.</li>
	// <li>bloody: bloody scenes.</li>
	// <li>banners: terrorism flags.</li>
	// <li>militant: militants.</li>
	// <li>explosion: explosions and fires.</li>
	// <li>terrorists: terrorists.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Label *string `json:"Label,omitempty" name:"Label"`

	// List of video segments that contain the detected terrorism information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AnimatedGraphicTaskInput struct {

	// Animated image generating template ID
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Start time of animated image in video in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time of animated image in video in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type AnimatedGraphicsTemplate struct {

	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of an animated image generating template.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of an animated image generating template.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Maximum value of the width (or long side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Animated image format.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Frame rate.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Image quality.
	Quality *float64 `json:"Quality,omitempty" name:"Quality"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ApplyUploadRequest struct {
	*tchttp.BaseRequest

	// Media type. For the detailed valid values, please see [Upload Overview](/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B).
	MediaType *string `json:"MediaType,omitempty" name:"MediaType"`

	// Media name.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// Cover type. For the detailed valid values, please see [Upload Overview](/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B).
	CoverType *string `json:"CoverType,omitempty" name:"CoverType"`

	// Subsequent task operation on a media file, i.e., after a media file is uploaded, task flow operations will be initiated automatically. This parameter value is a task flow template name. VOD supports [creating task flow templates](/document/product/266/33819) and naming the templates.
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Expiration time of a media file in ISO 8601 format. For more information, please see [Notes on ISO Date Format](/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Specifies upload region. This is only applicable to users that have special requirements for the upload region.
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](/document/product/266/7812) API.
	// <li>Default value: 0, which means "Other".</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Source context, which is used to pass through the user request information. The [upload callback](/document/product/266/7830) API will return the value of this field. It can contain up to 250 characters.
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`

	// Session context, which is used to pass through the user request information. If the `Procedure` parameter is specified, the [task flow status change callback](/document/product/266/9636) API will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ApplyUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Storage bucket, which is used as the `bucket_name` in the URL of the upload API.
		StorageBucket *string `json:"StorageBucket,omitempty" name:"StorageBucket"`

		// Storage region, which is used as the `Region` in the `Host` of the upload API.
		StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

		// VOD session, which is used to confirm the `VodSessionKey` parameter of the upload API.
		VodSessionKey *string `json:"VodSessionKey,omitempty" name:"VodSessionKey"`

		// Media storage path, which is used as the `Key` of the stored media of the upload API.
	// Note: this field may return null, indicating that no valid values can be obtained.
		MediaStoragePath *string `json:"MediaStoragePath,omitempty" name:"MediaStoragePath"`

		// Cover storage path, which is used as the `Key` of the stored cover of the upload API.
	// Note: this field may return null, indicating that no valid values can be obtained.
		CoverStoragePath *string `json:"CoverStoragePath,omitempty" name:"CoverStoragePath"`

		// Temporary credential, which is used for authentication of the upload API.
		TempCertificate *TempCertificate `json:"TempCertificate,omitempty" name:"TempCertificate"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsrFullTextConfigureInfo struct {

	// Switch of full speech recognition task. Valid values:
	// <li>ON: enables intelligent full speech recognition task;</li>
	// <li>OFF: disables intelligent full speech recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Format of generated subtitles file. If this parameter is left empty or a blank string is entered, no subtitles files will be generated. Valid value:
	// <li>vtt: generates a WebVTT subtitles file.</li>
	SubtitleFormat *string `json:"SubtitleFormat,omitempty" name:"SubtitleFormat"`
}

type AsrWordsConfigureInfo struct {

	// Switch of speech keyword recognition task. Valid values:
	// <li>ON: enables speech keyword recognition task;</li>
	// <li>OFF: disables speech keyword recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`
}

type AudioTemplateInfo struct {

	// Audio stream encoder.
	// When the outer `Container` parameter is `mp3`, the valid value is:
	// <li>libmp3lame.</li>
	// When the outer `Container` parameter is `ogg` or `flac`, the valid value is:
	// <li>flac.</li>
	// When the outer `Container` parameter is `m4a`, the valid values include:
	// <li>libfdk_aac;</li>
	// <li>libmp3lame;</li>
	// <li>ac3.</li>
	// When the outer `Container` parameter is `mp4` or `flv`, the valid values include:
	// <li>libfdk_aac: more suitable for mp4;</li>
	// <li>libmp3lame: More suitable for flv;</li>
	// <li>mp2.</li>
	// When the outer `Container` parameter is `hls`, the valid values include:
	// <li>libfdk_aac;</li>
	// <li>libmp3lame.</li>
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Audio stream bitrate in Kbps. Value range: 0 and [26, 256].
	// If the value is 0, the bitrate of the audio stream will be the same as that of the original audio.
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Audio stream sample rate. Valid values:
	// <li>32,000</li>
	// <li>44,100</li>
	// <li>48,000</li>
	// In Hz.
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// Audio channel system. Valid values:
	// <li>1: mono</li>
	// <li>2: dual</li>
	// <li>6: stereo</li>
	// Default value: 2.
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type AudioTemplateInfoForUpdate struct {

	// Audio stream encoder.
	// When the outer `Container` parameter is `mp3`, the valid value is:
	// <li>libmp3lame.</li>
	// When the outer `Container` parameter is `ogg` or `flac`, the valid value is:
	// <li>flac.</li>
	// When the outer `Container` parameter is `m4a`, the valid values include:
	// <li>libfdk_aac;</li>
	// <li>libmp3lame;</li>
	// <li>ac3.</li>
	// When the outer `Container` parameter is `mp4` or `flv`, the valid values include:
	// <li>libfdk_aac: more suitable for mp4;</li>
	// <li>libmp3lame: More suitable for flv;</li>
	// <li>mp2.</li>
	// When the outer `Container` parameter is `hls`, the valid values include:
	// <li>libfdk_aac;</li>
	// <li>libmp3lame.</li>
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Audio stream bitrate in Kbps. Value range: 0 and [26, 256]. If the value is 0, the bitrate of the audio stream will be the same as that of the original audio.
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Audio stream sample rate. Valid values:
	// <li>32,000</li>
	// <li>44,100</li>
	// <li>48,000</li>
	// In Hz.
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// Audio channel system. Valid values:
	// <li>1: mono</li>
	// <li>2: dual</li>
	// <li>6: stereo</li>
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type AudioTrackItem struct {

	// Source of media file for audio material, which can be an ID of a VOD file or URL of another file.
	SourceMedia *string `json:"SourceMedia,omitempty" name:"SourceMedia"`

	// Start time of audio segment in material file in seconds. Default value: 0, which means to start capturing from the beginning position of the material.
	SourceMediaStartTime *float64 `json:"SourceMediaStartTime,omitempty" name:"SourceMediaStartTime"`

	// Audio segment duration in seconds. By default, the length of the material will be used, which means that the entire material will be captured.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Operation on audio segment, such as volume adjustment.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioOperations []*AudioTransform `json:"AudioOperations,omitempty" name:"AudioOperations" list`
}

type AudioTrackTemplateInfo struct {

	// Unique ID of a template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Audio track encoder.
	// When `Container` is `mp3`, the valid value is:
	// <li>libmp3lame.</li>
	// When `Container` is `ogg` or `flac`, the valid value is:
	// <li>flac.</li>
	// When `Container` is `m4a`, the valid values include:
	// <li>libfdk_aac;</li>
	// <li>libmp3lame;</li>
	// <li>ac3.</li>
	// When the video track's `Container` is `mp4` or `flv`, the valid values include:
	// <li>libfdk_aac: more suitable for mp4;</li>
	// <li>libmp3lame: more suitable for flv;</li>
	// <li>mp2.</li>
	// When the video track's `Container` is `hls`, the valid values include:
	// <li>libfdk_aac;</li>
	// <li>libmp3lame.</li>
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Audio stream bitrate in Kbps. Value range: 0 and [26, 256].
	// If the value is 0, the bitrate of the audio stream will be the same as that of the original audio.
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Audio stream sample rate. Valid values:
	// <li>32,000</li>
	// <li>44,100</li>
	// <li>48,000</li>
	// In Hz.
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Audio channel system. Valid values:
	// <li>1: Mono</li>
	// <li>2: Dual</li>
	// <li>6: Stereo</li>
	// Default value: 2.
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AudioTransform struct {

	// Audio operation type. Valid values:
	// <li>Volume: volume adjustment.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Volume adjustment parameter, which is valid if `Type` is `Volume`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VolumeParam *AudioVolumeParam `json:"VolumeParam,omitempty" name:"VolumeParam"`
}

type AudioVolumeParam struct {

	// Whether to mute. Valid values: 0, 1.
	// <li>0: not muted.</li>
	// <li>1: muted.</li>
	// Default value: 0.
	Mute *int64 `json:"Mute,omitempty" name:"Mute"`

	// Audio gain. Value range: 0–10.
	// <li>If the value is greater than 1, the volume will be increased.</li>
	// <li>If the value is smaller than 1, the volume will be decreased.</li>
	// <li>0 and 1: no change.</li>
	// Default value: 0.
	Gain *float64 `json:"Gain,omitempty" name:"Gain"`
}

type Canvas struct {

	// Background color. Valid values:
	// <li>Black: black background</li>
	// <li>White: white background</li>
	// Default value: Black.
	Color *string `json:"Color,omitempty" name:"Color"`

	// Canvas width, which is the width of the output video. Value range: 0–4096 px.
	// Default value: 0, which means that the value is the same as the video width of the first video segment in the first video track.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Canvas height, which is the height (or long side) of the output video. Value range: 0–4096 px.
	// Default value: 0, which means that the value is the same as the video height of the first video segment in the first video track.
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type ClassificationConfigureInfo struct {

	// Switch of intelligent categorization task. Valid values:
	// <li>ON: enables intelligent categorization task;</li>
	// <li>OFF: disables intelligent categorization task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ClassificationConfigureInfoForUpdate struct {

	// Switch of intelligent categorization task. Valid values:
	// <li>ON: enables intelligent categorization task;</li>
	// <li>OFF: disables intelligent categorization task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ClipFileInfo2017 struct {

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Output target file ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Output target file address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// Output target file type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type ClipTask2017 struct {

	// Video clipping task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// ID of source file for video clipping task.
	SrcFileId *string `json:"SrcFileId,omitempty" name:"SrcFileId"`

	// Information of file output by video clipping.
	FileInfo *ClipFileInfo2017 `json:"FileInfo,omitempty" name:"FileInfo"`
}

type CommitUploadRequest struct {
	*tchttp.BaseRequest

	// VOD session, which takes the returned value (VodSessionKey) of the `ApplyUpload` API.
	VodSessionKey *string `json:"VodSessionKey,omitempty" name:"VodSessionKey"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CommitUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommitUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of media file.
		FileId *string `json:"FileId,omitempty" name:"FileId"`

		// Media playback address.
	// Note: this field may return null, indicating that no valid values can be obtained.
		MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

		// Media cover address.
	// Note: this field may return null, indicating that no valid values can be obtained.
		CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommitUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ComposeMediaOutput struct {

	// Filename of up to 64 characters.
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Description, which can contain up to 128 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](/document/product/266/7812) API.
	// <li>Default value: 0, which means "Other".</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Expiration time of output media file in ISO 8601 format, after which the file will be deleted. Files will never expire by default. For more information, please see [Notes on ISO Date Format](https://cloud.tencent.com/document/product/266/11732#I).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Container. Valid values: mp4, mp3. mp3 is for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Information of output video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoStream *OutputVideoStream `json:"VideoStream,omitempty" name:"VideoStream"`

	// Information of output audio.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioStream *OutputAudioStream `json:"AudioStream,omitempty" name:"AudioStream"`

	// Whether to remove video data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	// Default value: 0.
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	// Default value: 0.
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`
}

type ComposeMediaRequest struct {
	*tchttp.BaseRequest

	// List of input media tracks, i.e., information of multiple tracks composed of video, audio, image, and other materials. Multiple input tracks are aligned with the output media file on the time axis. The materials of each track at the same time point on the time axis will be superimposed. Specifically, videos or images will be superimposed for video image by track order, where a material with a higher track order will be more on top, while audio materials will be mixed.
	Tracks []*MediaTrack `json:"Tracks,omitempty" name:"Tracks" list`

	// Information of output media file.
	Output *ComposeMediaOutput `json:"Output,omitempty" name:"Output"`

	// Canvas used for composing video file.
	Canvas *Canvas `json:"Canvas,omitempty" name:"Canvas"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ComposeMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ComposeMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ComposeMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Media file composing task ID, which can be used to query the status of composing task (with task type being `MakeMedia`).
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ComposeMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ComposeMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ComposeMediaTask struct {

	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task flow status. Valid values:
	// <li>PROCESSING: processing;</li>
	// <li>FINISH: completed.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of media file composing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Input *ComposeMediaTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of media file composing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *ComposeMediaTaskOutput `json:"Output,omitempty" name:"Output"`
}

type ComposeMediaTaskInput struct {

	// List of input media tracks, i.e., information of multiple tracks composed of video, audio, image, and other materials.
	Tracks []*MediaTrack `json:"Tracks,omitempty" name:"Tracks" list`

	// Canvas used for composing video file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Canvas *Canvas `json:"Canvas,omitempty" name:"Canvas"`

	// Information of output media file.
	Output *ComposeMediaOutput `json:"Output,omitempty" name:"Output"`
}

type ComposeMediaTaskOutput struct {

	// File type, such as mp4 and mp3.
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// Media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Media file playback address.
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// Filename of up to 64 characters.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](/document/product/266/7812) API.
	// <li>Default value: 0, which means "Other".</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Expiration time of output media file in ISO 8601 format, after which the file will be deleted. Files will never expire by default. For more information, please see [Notes on ISO Date Format](https://cloud.tencent.com/document/product/266/11732#I).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type ConcatFileInfo2017 struct {

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// ID of source file for video splicing.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Address of source file for video splicing.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// Format of source file for video splicing.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type ConcatTask2017 struct {

	// Video splicing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Information of source file for video splicing.
	FileInfoSet []*ConcatFileInfo2017 `json:"FileInfoSet,omitempty" name:"FileInfoSet" list`
}

type ConfirmEventsRequest struct {
	*tchttp.BaseRequest

	// Event handler, i.e., the `EventSet. EventHandle` field in the output parameters of the [event notification pulling](/document/product/266/33433) API.
	// Array length limit: 16.
	EventHandles []*string `json:"EventHandles,omitempty" name:"EventHandles" list`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ConfirmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConfirmEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConfirmEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConfirmEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ContentReviewTemplateItem struct {

	// Unique ID of content audit template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Content audit template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Content audit template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Porn information detection control parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// Terrorism information detection control parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// Politically sensitive information detection control parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this parameter is not supported yet.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitempty" name:"ProhibitedConfigure"`

	// Custom content audit control parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// Switch controlling whether to add audit result to review list (for human review).
	// <li>ON: yes;</li>
	// <li>OFF: no.</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitempty" name:"ReviewWallSwitch"`

	// Frame capturing interval in seconds. If this parameter is left empty, 1 second will be used by default. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CoverBySnapshotTaskInput struct {

	// Time point screencapturing template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Screencapturing mode. Valid values:
	// <li>Time: screencaptures by time point</li>
	// <li>Percent: screencaptures by percentage</li>
	PositionType *string `json:"PositionType,omitempty" name:"PositionType"`

	// Screenshot position:
	// <li>For time point screencapturing, this means to take a screenshot at a specified time point (in seconds) and use it as the cover</li>
	// <li>For percentage screencapturing, this value means to take a screenshot at a specified percentage of the video duration and use it as the cover</li>
	PositionValue *float64 `json:"PositionValue,omitempty" name:"PositionValue"`

	// List of up to 10 image or text watermarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type CoverBySnapshotTaskOutput struct {

	// Cover URL.
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`
}

type CoverConfigureInfo struct {

	// Switch of intelligent cover generating task. Valid values:
	// <li>ON: enables intelligent cover generating task;</li>
	// <li>OFF: disables intelligent cover generating task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CoverConfigureInfoForUpdate struct {

	// Switch of intelligent cover generating task. Valid values:
	// <li>ON: enables intelligent cover generating task;</li>
	// <li>OFF: disables intelligent cover generating task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CreateAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest

	// Video content analysis template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Video content analysis template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter of intelligent categorization task.
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitempty" name:"ClassificationConfigure"`

	// Control parameter of intelligent tagging task.
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitempty" name:"TagConfigure"`

	// Control parameter of intelligent cover generating task.
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitempty" name:"CoverConfigure"`

	// Control parameter of intelligent frame-specific tagging task.
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitempty" name:"FrameTagConfigure"`

	// Control parameter of an intelligent highlight generating task.
	HighlightConfigure *HighlightsConfigureInfo `json:"HighlightConfigure,omitempty" name:"HighlightConfigure"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAIAnalysisTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of video content analysis template.
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAIAnalysisTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAIAnalysisTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest

	// Video frame rate in Hz. Value range: [1, 30].
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Maximum value of the width (or long side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Animated image format. Valid values: gif; webp. Default value: gif.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Image quality. Value range: [1, 100]. Default value: 75.
	Quality *float64 `json:"Quality,omitempty" name:"Quality"`

	// Name of an animated image generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of an animated image generating template.
		Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAnimatedGraphicsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClassRequest struct {
	*tchttp.BaseRequest

	// Parent category ID. For a first-level category, enter `-1`.
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// Category name. Length limit: 1–64 characters.
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Category ID
		ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageSpriteTask2017 struct {

	// Image sprite generating task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// ID of generated image sprite file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Image sprite specification. For more information, please see [Image Sprite Generating Template](https://cloud.tencent.com/document/product/266/33480#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Total number of subimages in image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Address of output image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageSpriteUrlSet []*string `json:"ImageSpriteUrlSet,omitempty" name:"ImageSpriteUrlSet" list`

	// Address of WebVtt file for the position-time relationship among subimages in an image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WebVttUrl *string `json:"WebVttUrl,omitempty" name:"WebVttUrl"`
}

type CreateImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest

	// Sampling type. Valid values:
	// <li>Percent: by percent.</li>
	// <li>Time: by time interval.</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Subimage row count of an image sprite.
	RowCount *uint64 `json:"RowCount,omitempty" name:"RowCount"`

	// Subimage column count of an image sprite.
	ColumnCount *uint64 `json:"ColumnCount,omitempty" name:"ColumnCount"`

	// Name of an image sprite generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Maximum value of the width (or long side) of a subimage in an image sprite in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a subimage in an image sprite in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

func (r *CreateImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageSpriteTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of an image sprite generating template.
		Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageSpriteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageSpriteTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProcedureTemplateRequest struct {
	*tchttp.BaseRequest

	// Task flow name (up to 20 characters).
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter of video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Parameter of AI-based content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Parameter of AI-based content analysis task.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of AI-based content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProcedureTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProcedureTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Sampled screencapturing type. Valid values:
	// <li>Percent: by percent.</li>
	// <li>Time: by time interval.</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Image format. Valid values: jpg, png. Default value: jpg.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

func (r *CreateSampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSampleSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of a sampled screencapturing template.
		Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSampleSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSampleSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest

	// Name of a time point screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Image format. Valid values: jpg, png. Default value: jpg.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

func (r *CreateSnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of a time point screencapturing template.
		Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSnapshotByTimeOffsetTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Container. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Transcoding template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Whether to remove video data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	// Default value: 0.
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	// Default value: 0.
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter. This field is required when `RemoveVideo` is 0.
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter. This field is required when `RemoveAudio` is 0.
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitempty" name:"TEHDConfig"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of transcoding template.
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// Watermarking type. Valid values:
	// <li>image: image watermark;</li>
	// <li>text: text watermark;</li>
	// <li>svg: SVG watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Watermarking template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Origin position. Valid values:
	// <li>TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>TopRight: the origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>BottomLeft: the origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>BottomRight: the origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	// Default value: TopLeft.
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the watermark will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width;</li>
	// <li>If the string ends in px, the `XPos` of the watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.</li>
	// Default value: 0 px.
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the watermark will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height;</li>
	// <li>If the string ends in px, the `YPos` of the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.</li>
	// Default value: 0 px.
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// Image watermarking template. This field is required when `Type` is `image` and is invalid when `Type` is `text`.
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is required when `Type` is `text` and is invalid when `Type` is `image`.
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is required when `Type` is `svg` and is invalid when `Type` is `image` or `text`.
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitempty" name:"SvgTemplate"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWatermarkTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of watermarking template.
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// Watermark image address. This field is valid only when `Type` is `image`.
		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWatermarkTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAIAnalysisTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAIAnalysisTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAIAnalysisTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAIRecognitionTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAIRecognitionTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAIRecognitionTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAnimatedGraphicsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageSpriteTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageSpriteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageSpriteTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProcedureTemplateRequest struct {
	*tchttp.BaseRequest

	// Task flow name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProcedureTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProcedureTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteSampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSampleSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSampleSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSampleSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a specified time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteSnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotByTimeOffsetTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWatermarkTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWatermarkTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAIAnalysisTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of video content analysis templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAIAnalysisTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAIAnalysisTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAIAnalysisTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of video content analysis template details.
		AIAnalysisTemplateSet []*AIAnalysisTemplateItem `json:"AIAnalysisTemplateSet,omitempty" name:"AIAnalysisTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAIAnalysisTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAIAnalysisTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAIRecognitionTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of video content recognition templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAIRecognitionTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAIRecognitionTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAIRecognitionTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of video content recognition template details.
		AIRecognitionTemplateSet []*AIRecognitionTemplateItem `json:"AIRecognitionTemplateSet,omitempty" name:"AIRecognitionTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAIRecognitionTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAIRecognitionTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAdaptiveDynamicStreamingTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of transcoding to adaptive bitrate streaming templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAdaptiveDynamicStreamingTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAdaptiveDynamicStreamingTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAdaptiveDynamicStreamingTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of transcoding to adaptive bitrate streaming template details.
		AdaptiveDynamicStreamingTemplateSet []*AdaptiveDynamicStreamingTemplate `json:"AdaptiveDynamicStreamingTemplateSet,omitempty" name:"AdaptiveDynamicStreamingTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAdaptiveDynamicStreamingTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAdaptiveDynamicStreamingTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAnimatedGraphicsTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of animated image generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAnimatedGraphicsTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAnimatedGraphicsTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAnimatedGraphicsTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of animated image generating template details.
		AnimatedGraphicsTemplateSet []*AnimatedGraphicsTemplate `json:"AnimatedGraphicsTemplateSet,omitempty" name:"AnimatedGraphicsTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAnimatedGraphicsTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAnimatedGraphicsTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAudioTrackTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAudioTrackTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAudioTrackTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAudioTrackTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of audio track template details.
		AudioTrackTemplateSet []*AudioTrackTemplateInfo `json:"AudioTrackTemplateSet,omitempty" name:"AudioTrackTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAudioTrackTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAudioTrackTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCDNUsageDataRequest struct {
	*tchttp.BaseRequest

	// Start date in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I). The end date must be after the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// CDN statistics type. Valid values:
	// <li>Flux: traffic in bytes.</li>
	// <li>Bandwidth: bandwidth in bps.</li>
	DataType *string `json:"DataType,omitempty" name:"DataType"`

	// Time granularity of usage data in minutes. Valid values:
	// <li>5: 5-minute granularity, which returns the details at the 5-minute granularity within the specified time range.</li>
	// <li>60: 1-hour granularity, which returns the details at the 1-hour granularity within the specified time range.</li>
	// <li>1440: 1-day granularity, which returns the details at the 1-day granularity within the specified time range.</li>
	// Default value: 1440. Data at the 1-day granularity will be returned.
	// When the value of this field is 1, the total usage of all subapplications (including primary application) are queried by an admin.
	DataInterval *uint64 `json:"DataInterval,omitempty" name:"DataInterval"`

	// List of domain names. The usage data of up to 20 domain names can be queried at a time. You can specify multiple domain names and query their combined usage data. The usage data of all domain names will be returned by default.
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames" list`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	// When the value of this field is 1, the total usage of all subapplications (including primary application) are queried by an admin. In this case, only 1-day granularity is supported.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeCDNUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCDNUsageDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCDNUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time granularity in minutes.
		DataInterval *int64 `json:"DataInterval,omitempty" name:"DataInterval"`

		// CDN statistics.
		Data []*StatDataItem `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCDNUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCDNUsageDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentReviewTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of content audit templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeContentReviewTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentReviewTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentReviewTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of content audit template details.
		ContentReviewTemplateSet []*ContentReviewTemplateItem `json:"ContentReviewTemplateSet,omitempty" name:"ContentReviewTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContentReviewTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentReviewTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSpriteTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of image sprite generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeImageSpriteTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSpriteTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSpriteTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of image sprite generating template details.
		ImageSpriteTemplateSet []*ImageSpriteTemplate `json:"ImageSpriteTemplateSet,omitempty" name:"ImageSpriteTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSpriteTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSpriteTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaInfosRequest struct {
	*tchttp.BaseRequest

	// List of media file IDs. N starts from 0 and can be up to 19.
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds" list`

	// Specifies information entry that needs to be returned by all media files. Multiple entries can be specified simultaneously. N starts from 0. If this field is left empty, all information entries will be returned by default. Valid values:
	// <li>basicInfo (basic video information).</li>
	// <li>metaData (video metadata).</li>
	// <li>transcodeInfo (result information of video transcoding).</li>
	// <li>animatedGraphicsInfo (result information of animated image generating task).</li>
	// <li>imageSpriteInfo (image sprite information).</li>
	// <li>snapshotByTimeOffsetInfo (time point screenshot information).</li>
	// <li>sampleSnapshotInfo (sampled screenshot information).</li>
	// <li>keyFrameDescInfo (timestamp information).</li>
	// <li>adaptiveDynamicStreamingInfo (information of adaptive bitrate streaming).</li>
	// <li>miniProgramReviewInfo (WeChat Mini Program audit information).</li>
	Filters []*string `json:"Filters,omitempty" name:"Filters" list`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeMediaInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Media file information list.
	// Note: this field may return null, indicating that no valid values can be obtained.
		MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet" list`

		// List of IDs of files that do not exist.
	// Note: this field may return null, indicating that no valid values can be obtained.
		NotExistFileIdSet []*string `json:"NotExistFileIdSet,omitempty" name:"NotExistFileIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaProcessUsageDataRequest struct {
	*tchttp.BaseRequest

	// Start date in [ISO date format](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F). The end date must be on or after the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Type of video processing task to be queried. Valid value: Transcode. Default value: Transcode.
	// <li>Transcode: transcoding</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeMediaProcessUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaProcessUsageDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaProcessUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Overview of video processing statistics, which displays the overview and details of queried tasks.
		MediaProcessDataSet []*TaskStatData `json:"MediaProcessDataSet,omitempty" name:"MediaProcessDataSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaProcessUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaProcessUsageDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcedureTemplatesRequest struct {
	*tchttp.BaseRequest

	// Name filter of task flow template. Array length limit: 100.
	Names []*string `json:"Names,omitempty" name:"Names" list`

	// Filter of task flow template types. Valid values:
	// <li>Preset: preset task flow template;</li>
	// <li>Custom: custom task flow template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeProcedureTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcedureTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcedureTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of task flow template details.
		ProcedureTemplateSet []*ProcedureTemplate `json:"ProcedureTemplateSet,omitempty" name:"ProcedureTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcedureTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcedureTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReviewDetailsRequest struct {
	*tchttp.BaseRequest

	// Start date in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I). The end date must be after the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeReviewDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReviewDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReviewDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of initiated content audits.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Total content audit duration.
		TotalDuration *int64 `json:"TotalDuration,omitempty" name:"TotalDuration"`

		// Data of content audit duration, which is collected once per day.
		Data []*StatDataItem `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReviewDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReviewDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSampleSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of sampled screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeSampleSnapshotTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSampleSnapshotTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSampleSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of sampled screencapturing template details.
		SampleSnapshotTemplateSet []*SampleSnapshotTemplate `json:"SampleSnapshotTemplateSet,omitempty" name:"SampleSnapshotTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSampleSnapshotTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSampleSnapshotTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotByTimeOffsetTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of time point screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeSnapshotByTimeOffsetTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSnapshotByTimeOffsetTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotByTimeOffsetTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of time point screencapturing template details.
		SnapshotByTimeOffsetTemplateSet []*SnapshotByTimeOffsetTemplate `json:"SnapshotByTimeOffsetTemplateSet,omitempty" name:"SnapshotByTimeOffsetTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotByTimeOffsetTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSnapshotByTimeOffsetTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageDataRequest struct {
	*tchttp.BaseRequest

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeStorageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStorageDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of current media files.
		MediaCount *uint64 `json:"MediaCount,omitempty" name:"MediaCount"`

		// Total current storage capacity in bytes.
		TotalStorage *uint64 `json:"TotalStorage,omitempty" name:"TotalStorage"`

		// Current Standard_IA storage capacity in bytes.
		InfrequentStorage *uint64 `json:"InfrequentStorage,omitempty" name:"InfrequentStorage"`

		// Current Standard storage capacity in bytes.
		StandardStorage *uint64 `json:"StandardStorage,omitempty" name:"StandardStorage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStorageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStorageDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageDetailsRequest struct {
	*tchttp.BaseRequest

	// Start time in ISO 8601 format. For more information, please see [Notes on ISO Date Format](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in ISO 8601 format, which must be after the start time. For more information, please see [Notes on ISO Date Format](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Query time interval. Valid values:
	// <li>Minute: once per minute.</li>
	// <li>Hour: once per hour.</li>
	// <li>Day: once per day.</li>
	// The default value is determined by the time span. `Minute` will be used if the time span is less than 1 hour, `Hour` if less than or equal to 7 days, and `Day` if more than 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Storage class to be queried. Valid values:
	// <li>TotalStorage: total storage capacity.</li>
	// <li>StandardStorage: Standard storage.</li>
	// <li>InfrequentStorage: Standard_IA storage.</li>
	// Default value: TotalStorage.
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	// When the value of this field is 1, the total usage of all subapplications (including primary application) are queried by an admin.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeStorageDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStorageDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Storage statistics. One data entry per minute/hour/day.
		Data []*StatDataItem `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStorageDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStorageDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubAppIdsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSubAppIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubAppIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubAppIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Subapplication information set.
		SubAppIdInfoSet []*SubAppIdInfo `json:"SubAppIdInfoSet,omitempty" name:"SubAppIdInfoSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubAppIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubAppIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task type. Valid values:
	// <li>Procedure: video processing task;</li>
	// <li>EditMedia: video editing task;</li>
	// <li>WechatPublish: release on WeChat task;</li>
	// <li>WechatMiniProgramPublish: release on WeChat Mini Program task;</li>
	// <li>ComposeMedia: media file composing task;</li>
	// <li>PullUpload: media file pulling for upload task.</li>
	// 
	// Task types compatible with v2017:
	// <li>Transcode: transcoding task;</li>
	// <li>SnapshotByTimeOffset: screencapturing task</li>
	// <li>Concat: video splicing task;</li>
	// <li>Clip: video clipping task;</li>
	// <li>ImageSprites: image sprite generating task.</li>
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// Task status. Valid values:
	// <li>WAITING: waiting;</li>
	// <li>PROCESSING: processing;</li>
	// <li>FINISH: completed.</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// Creation time of task in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// Start time of task execution in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
		BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

		// End time of task execution in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

		// Video processing task information. This field has a value only when `TaskType` is `Procedure`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ProcedureTask *ProcedureTask `json:"ProcedureTask,omitempty" name:"ProcedureTask"`

		// Video editing task information. This field has a value only when `TaskType` is `EditMedia`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		EditMediaTask *EditMediaTask `json:"EditMediaTask,omitempty" name:"EditMediaTask"`

		// Release on WeChat task information. This field has a value only when `TaskType` is `WechatPublish`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		WechatPublishTask *WechatPublishTask `json:"WechatPublishTask,omitempty" name:"WechatPublishTask"`

		// Media file composing task information. This field has a value only when `TaskType` is `ComposeMedia`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ComposeMediaTask *ComposeMediaTask `json:"ComposeMediaTask,omitempty" name:"ComposeMediaTask"`

		// Media file pulling for upload task information. This field has a value only when `TaskType` is `PullUpload`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		PullUploadTask *PullUploadTask `json:"PullUploadTask,omitempty" name:"PullUploadTask"`

		// Video transcoding task information. This field has a value only when `TaskType` is `Transcode`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TranscodeTask *TranscodeTask2017 `json:"TranscodeTask,omitempty" name:"TranscodeTask"`

		// Time point screencapturing task information. This field has a value only when `TaskType` is `SnapshotByTimeOffset`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		SnapshotByTimeOffsetTask *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetTask,omitempty" name:"SnapshotByTimeOffsetTask"`

		// Video splicing task information. This field has a value only when `TaskType` is `Concat`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ConcatTask *ConcatTask2017 `json:"ConcatTask,omitempty" name:"ConcatTask"`

		// Video clipping task information. This field has a value only when `TaskType` is `Clip`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ClipTask *ClipTask2017 `json:"ClipTask,omitempty" name:"ClipTask"`

		// Image sprite creating task information. This field has a value only when `TaskType` is `ImageSprite`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		CreateImageSpriteTask *CreateImageSpriteTask2017 `json:"CreateImageSpriteTask,omitempty" name:"CreateImageSpriteTask"`

		// Release on WeChat Mini Program task information. This field has a value only when `TaskType` is `WechatMiniProgramPublish`.
	// Note: this field may return null, indicating that no valid values can be obtained.
		WechatMiniProgramPublishTask *WechatMiniProgramPublishTask `json:"WechatMiniProgramPublishTask,omitempty" name:"WechatMiniProgramPublishTask"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of transcoding templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Container filter. Valid values:
	// <li>Video: video container that can contain both video stream and audio stream;</li>
	// <li>PureAudio: audio container that can contain only audio stream.</li>
	ContainerType *string `json:"ContainerType,omitempty" name:"ContainerType"`

	// TESHD filter, which is used to filter common transcoding or ultra-fast HD transcoding templates. Valid values:
	// <li>Common: Common transcoding template;</li>
	// <li>TEHD: TESHD template.</li>
	TEHDType *string `json:"TEHDType,omitempty" name:"TEHDType"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of transcoding template details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TranscodeTemplateSet []*TranscodeTemplate `json:"TranscodeTemplateSet,omitempty" name:"TranscodeTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoTrackTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeVideoTrackTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVideoTrackTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoTrackTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of video track template details.
		VideoTrackTemplateSet []*VideoTrackTemplateInfo `json:"VideoTrackTemplateSet,omitempty" name:"VideoTrackTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVideoTrackTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVideoTrackTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWatermarkTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of watermarking templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// Watermark type filter. Valid values:
	// <li>image: image watermark;</li>
	// <li>text: text watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries
	// <li>Default value: 10;</li>
	// <li>Maximum value: 100.</li>
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeWatermarkTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWatermarkTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWatermarkTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of watermarking template details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		WatermarkTemplateSet []*WatermarkTemplate `json:"WatermarkTemplateSet,omitempty" name:"WatermarkTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWatermarkTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWatermarkTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditMediaFileInfo struct {

	// Video ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Start time offset of video clipping in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of video clipping in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type EditMediaOutputConfig struct {

	// Output filename of up to 64 characters, which is generated by the system by default.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// Output file format. Valid values: mp4, hls. Default value: mp4.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](/document/product/266/7812) API.
	// <li>Default value: 0, which means "Other".</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Expiration time of output media file in ISO 8601 format, after which the file will be deleted. Files will never expire by default. For more information, please see [Notes on ISO Date Format](https://cloud.tencent.com/document/product/266/11732#I).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type EditMediaRequest struct {
	*tchttp.BaseRequest

	// Input video type. Valid values: File, Stream.
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// Information of input video file, which is required if `InputType` is `File`.
	FileInfos []*EditMediaFileInfo `json:"FileInfos,omitempty" name:"FileInfos" list`

	// Input stream information, which is required if `InputType` is `Stream`.
	StreamInfos []*EditMediaStreamInfo `json:"StreamInfos,omitempty" name:"StreamInfos" list`

	// Editing template ID. Valid values: 10, 20. If this parameter is left empty, template 10 will be used.
	// <li>10: the input with the highest resolution will be used as the benchmark;</li>
	// <li>20: the input with the highest bitrate will be used as the benchmark;</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// [Task flow template](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF) name, which should be entered if you want to perform a task flow on the generated new video.
	ProcedureName *string `json:"ProcedureName,omitempty" name:"ProcedureName"`

	// Configuration of file generated after editing.
	OutputConfig *EditMediaOutputConfig `json:"OutputConfig,omitempty" name:"OutputConfig"`

	// Identifies the source context which is used to pass through the user request information. The `EditMediaComplete` callback and task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Task priority. The higher the value, the higher the priority. Value range: -10–10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// ID used for task deduplication. If there was a request with the same ID in the last day, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *EditMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Video editing task ID, which can be used to query the status of editing task (with task type being `EditMedia`).
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditMediaStreamInfo struct {

	// ID of recorded stream
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Start time of stream clipping in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of stream clipping in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type EditMediaTask struct {

	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task flow status. Valid values:
	// <li>PROCESSING: processing;</li>
	// <li>FINISH: completed.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of video editing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Input *EditMediaTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of video editing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *EditMediaTaskOutput `json:"Output,omitempty" name:"Output"`

	// If a video processing flow is specified when a video editing task is initiated, this field will be the ID of the task flow.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type EditMediaTaskInput struct {

	// Input video source type. Valid values: File, Stream.
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// Information of input video file. This field has a value only when `InputType` is `File`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileInfoSet []*EditMediaFileInfo `json:"FileInfoSet,omitempty" name:"FileInfoSet" list`

	// Input stream information. This field has a value only when `InputType` is `Stream`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StreamInfoSet []*EditMediaStreamInfo `json:"StreamInfoSet,omitempty" name:"StreamInfoSet" list`
}

type EditMediaTaskOutput struct {

	// File type, such as mp4 and flv.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// Media file playback address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// Media file ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Output filename of up to 64 characters, which is generated by the system by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](/document/product/266/7812) API.
	// <li>Default value: 0, which means "Other".</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Expiration time of output media file in ISO 8601 format, after which the file will be deleted. Files will never expire by default. For more information, please see [Notes on ISO Date Format](https://cloud.tencent.com/document/product/266/11732#I).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type EmptyTrackItem struct {

	// Duration in seconds.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
}

type EventContent struct {

	// Event handler. The caller must call `ConfirmEvents` to confirm that the message has been received, and the confirmation is valid for 30 seconds. After the confirmation expires, the event can be obtained again.
	EventHandle *string `json:"EventHandle,omitempty" name:"EventHandle"`

	// <b>Supported event type:</b>
	// <li>NewFileUpload: video upload completion;</li>
	// <li>ProcedureStateChanged: task flow status change;</li>
	// <li>FileDeleted: video deletion completion;</li>
	// <li>PullComplete: video pull for upload completion;</li>
	// <li>EditMediaComplete: video editing completion;</li>
	// <li>WechatPublishComplete: release on WeChat completion;</li>
	// <li>ComposeMediaComplete: media file composing completion;</li>
	// <li>WechatMiniProgramPublishComplete: release on WeChat Mini Program completion.</li>
	// <b>Event types compatible with v2017:</b>
	// <li>TranscodeComplete: video transcoding completion;</li>
	// <li>ConcatComplete: video splicing completion;</li>
	// <li>ClipComplete: video clipping completion;</li>
	// <li>CreateImageSpriteComplete: image sprite generating completion;</li>
	// <li>CreateSnapshotByTimeOffsetComplete: time point screencapturing completion.</li>
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// Video upload completion event, which is valid if the event type is `NewFileUpload`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileUploadEvent *FileUploadTask `json:"FileUploadEvent,omitempty" name:"FileUploadEvent"`

	// Task flow status change event, which is valid if the event type is `ProcedureStateChanged`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProcedureStateChangeEvent *ProcedureTask `json:"ProcedureStateChangeEvent,omitempty" name:"ProcedureStateChangeEvent"`

	// File deletion event, which is valid if the event type is `FileDeleted`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileDeleteEvent *FileDeleteTask `json:"FileDeleteEvent,omitempty" name:"FileDeleteEvent"`

	// Video pull for upload completion event, which is valid if the event type is `PullComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PullCompleteEvent *PullUploadTask `json:"PullCompleteEvent,omitempty" name:"PullCompleteEvent"`

	// Video editing completion event, which is valid if the event type is `EditMediaComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EditMediaCompleteEvent *EditMediaTask `json:"EditMediaCompleteEvent,omitempty" name:"EditMediaCompleteEvent"`

	// Release on WeChat completion event, which is valid if the event type is `WechatPublishComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WechatPublishCompleteEvent *WechatPublishTask `json:"WechatPublishCompleteEvent,omitempty" name:"WechatPublishCompleteEvent"`

	// Video transcoding completion event, which is valid if the event type is `TranscodeComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TranscodeCompleteEvent *TranscodeTask2017 `json:"TranscodeCompleteEvent,omitempty" name:"TranscodeCompleteEvent"`

	// Video splicing completion event, which is valid if the event type is `ConcatComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConcatCompleteEvent *ConcatTask2017 `json:"ConcatCompleteEvent,omitempty" name:"ConcatCompleteEvent"`

	// Video clipping completion event, which is valid if the event type is `ClipComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClipCompleteEvent *ClipTask2017 `json:"ClipCompleteEvent,omitempty" name:"ClipCompleteEvent"`

	// Image sprite generating completion event, which is valid if the event type is `CreateImageSpriteComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateImageSpriteCompleteEvent *CreateImageSpriteTask2017 `json:"CreateImageSpriteCompleteEvent,omitempty" name:"CreateImageSpriteCompleteEvent"`

	// Time point screencapturing completion event, which is valid when the event type is `CreateSnapshotByTimeOffsetComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetCompleteEvent *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetCompleteEvent,omitempty" name:"SnapshotByTimeOffsetCompleteEvent"`

	// Media file composing task completion event, which is valid when the event type is `ComposeMediaComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ComposeMediaCompleteEvent *ComposeMediaTask `json:"ComposeMediaCompleteEvent,omitempty" name:"ComposeMediaCompleteEvent"`

	// Release on WeChat Mini Program task completion event, which is valid if the event type is `WechatMiniProgramPublishComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WechatMiniProgramPublishCompleteEvent *WechatMiniProgramPublishTask `json:"WechatMiniProgramPublishCompleteEvent,omitempty" name:"WechatMiniProgramPublishCompleteEvent"`
}

type ExecuteFunctionRequest struct {
	*tchttp.BaseRequest

	// Name of called backend API.
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// API parameter. For specific parameter format, negotiate with the backend before calling.
	FunctionArg *string `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ExecuteFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExecuteFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String generated by packaging processing result. For specifications, negotiate with the backend.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *string `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FaceConfigureInfo struct {

	// Switch of face recognition task. Valid values:
	// <li>ON: enables intelligent face recognition task;</li>
	// <li>OFF: disables intelligent face recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Face recognition filter score. If this score is reached or exceeded, a recognition result will be returned. Value range: 0–100. Default value: 95.
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// Default figure filter tag, which specifies the default figure tag that needs to be returned. If this parameter is left empty or a blank value is entered, all results of the default figures will be returned. Valid values:
	// <li>entertainment: entertainment celebrity;</li>
	// <li>sport: sports celebrity;</li>
	// <li>politician: politically sensitive figure.</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitempty" name:"DefaultLibraryLabelSet" list`

	// Custom figure filter tag, which specifies the custom figure tag that needs to be returned. If this parameter is left empty or a blank value is entered, all results of the custom figures will be returned. Valid values:
	// There can be up to 10 tags, each with a length limit of 16 characters.
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitempty" name:"UserDefineLibraryLabelSet" list`

	// Figure library. Valid values:
	// <li>Default: default figure library;</li>
	// <li>UserDefine: custom figure library.</li>
	// <li>All: both default and custom figure libraries will be used.</li>
	// Default value: All (both default and custom figure libraries will be used.)
	FaceLibrary *string `json:"FaceLibrary,omitempty" name:"FaceLibrary"`
}

type FileDeleteTask struct {

	// List of IDs of deleted files.
	FileIdSet []*string `json:"FileIdSet,omitempty" name:"FileIdSet" list`
}

type FileUploadTask struct {

	// Unique file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Basic information of media file generated after upload is completed.
	MediaBasicInfo *MediaBasicInfo `json:"MediaBasicInfo,omitempty" name:"MediaBasicInfo"`

	// If a video processing flow is specified when a video is uploaded, this field will be the ID of the task flow.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`

	// Metadata, such as size, duration, video stream information, audio stream information, etc.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`
}

type ForbidMediaDistributionRequest struct {
	*tchttp.BaseRequest

	// List of media files. Up to 20 ones can be submitted at a time.
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds" list`

	// forbid: forbids, recover: unblocks.
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ForbidMediaDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidMediaDistributionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidMediaDistributionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of IDs of files that do not exist.
		NotExistFileIdSet []*string `json:"NotExistFileIdSet,omitempty" name:"NotExistFileIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForbidMediaDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidMediaDistributionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FrameTagConfigureInfo struct {

	// Switch of intelligent frame-specific tagging task. Valid values:
	// <li>ON: enables intelligent frame-specific tagging task;</li>
	// <li>OFF: disables intelligent frame-specific tagging task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Frame capturing interval in seconds. If this parameter is left empty, 1 second will be used by default. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
}

type FrameTagConfigureInfoForUpdate struct {

	// Switch of intelligent frame-specific tagging task. Valid values:
	// <li>ON: enables intelligent frame-specific tagging task;</li>
	// <li>OFF: disables intelligent frame-specific tagging task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Frame capturing interval in seconds. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
}

type HeadTailConfigureInfo struct {

	// Switch of video opening and ending credits recognition task. Valid values:
	// <li>ON: enables video opening and ending credits recognition task;</li>
	// <li>OFF: disables video opening and ending credits recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type HighlightSegmentItem struct {

	// Confidence.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Start time offset of a segment.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a segment.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type HighlightsConfigureInfo struct {

	// Switch of an intelligent highlight generating task. Valid values:
	// <li>ON: enable an intelligent highlight generating task;</li>
	// <li>OFF: disable an intelligent highlight generating task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type HighlightsConfigureInfoForUpdate struct {

	// Switch of an intelligent highlight generating task. Valid values:
	// <li>ON: enable an intelligent highlight generating task;</li>
	// <li>OFF: disable an intelligent highlight generating task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ImageSpriteTaskInput struct {

	// Image sprite generating template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type ImageSpriteTemplate struct {

	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of an image sprite generating template.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Maximum value of the width (or long side) of a subimage in an image sprite in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a subimage in an image sprite in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Sampling type.
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Subimage row count of an image sprite.
	RowCount *uint64 `json:"RowCount,omitempty" name:"RowCount"`

	// Subimage column count of an image sprite.
	ColumnCount *uint64 `json:"ColumnCount,omitempty" name:"ColumnCount"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type ImageTransform struct {

	// Type. Valid values:
	// <li> Rotate: image rotation.</li>
	// <li> Flip: image flipping.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Rotation angle of image with its center point as origin. Value range: 0–360. This parameter is valid if `Type` is `Rotate`.
	RotateAngle *float64 `json:"RotateAngle,omitempty" name:"RotateAngle"`

	// Image flipping action. Valid values:
	// <li>Horizental: horizontal flipping, i.e., horizontally mirrored.</li>
	// <li>Vertical: vertical flipping, i.e., vertically mirrored.</li>
	// This is valid if `Type` is `Flip`.
	Flip *string `json:"Flip,omitempty" name:"Flip"`
}

type ImageWatermarkInput struct {

	// String generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) a watermark image. JPEG and PNG images are supported.
	ImageContent *string `json:"ImageContent,omitempty" name:"ImageContent"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.</li>
	// Default value: 10%.
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.</li>
	// Default value: 0 px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
	Height *string `json:"Height,omitempty" name:"Height"`
}

type ImageWatermarkInputForUpdate struct {

	// String generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) a watermark image. JPEG and PNG images are supported.
	ImageContent *string `json:"ImageContent,omitempty" name:"ImageContent"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.</li>
	// `0px` means that `Height` will be proportionally scaled according to the video width.
	Height *string `json:"Height,omitempty" name:"Height"`
}

type ImageWatermarkTemplate struct {

	// Watermark image address.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px;</li>
	// `0px` means that `Height` will be proportionally scaled according to the video width.
	Height *string `json:"Height,omitempty" name:"Height"`
}

type LiveRealTimeClipRequest struct {
	*tchttp.BaseRequest

	// [LVB code](https://cloud.tencent.com/document/product/267/5959) of a stream.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Start time of stream clipping in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of stream clipping in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Whether to clip persistently. 0: no, 1: yes. Default: no.
	IsPersistence *int64 `json:"IsPersistence,omitempty" name:"IsPersistence"`

	// Storage expiration time of video generated by persistent clipping in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I). `9999-12-31T23:59:59Z` means `never expire`. After the expiration, the media file and its related resources (such as transcoding results and image sprites) will be permanently deleted. This parameter will be valid only when `IsPersistence` is 1. By default, the video will never expire.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// VOD task flow processing for video generated by persistent clipping. For more information, please see [Specifying Task Flow After Upload](https://cloud.tencent.com/document/product/266/9759). This parameter will be valid only when `IsPersistence` is 1.
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Whether the metadata of clipped video needs to be returned. 0: no, 1: yes. Default value: no.
	MetaDataRequired *uint64 `json:"MetaDataRequired,omitempty" name:"MetaDataRequired"`

	// Domain name used for live clipping. Time shifting must be enabled in LVB.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Reserved field. Do not enter a value for it.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *LiveRealTimeClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveRealTimeClipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LiveRealTimeClipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Playback URL of clipped video.
		Url *string `json:"Url,omitempty" name:"Url"`

		// Unique media file ID of video generated by persistent clipping.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FileId *string `json:"FileId,omitempty" name:"FileId"`

		// Task flow ID of video generated by persistent clipping.
	// Note: this field may return null, indicating that no valid values can be obtained.
		VodTaskId *string `json:"VodTaskId,omitempty" name:"VodTaskId"`

		// Metadata of clipped video.
	// Note: this field may return null, indicating that no valid values can be obtained.
		MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveRealTimeClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveRealTimeClipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MediaAdaptiveDynamicStreamingInfo struct {

	// Information array of adaptive bitrate streaming.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdaptiveDynamicStreamingSet []*AdaptiveDynamicStreamingInfoItem `json:"AdaptiveDynamicStreamingSet,omitempty" name:"AdaptiveDynamicStreamingSet" list`
}

type MediaAiAnalysisClassificationItem struct {

	// Name of intelligently generated category.
	Classification *string `json:"Classification,omitempty" name:"Classification"`

	// Confidence of intelligently generated category between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisCoverItem struct {

	// Address of intelligently generated cover.
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// Confidence of intelligently generated cover between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisFrameTagItem struct {

	// Frame-specific tag name.
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// Confidence of intelligently generated frame-specific tag between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisFrameTagSegmentItem struct {

	// Start time offset of frame-specific tag.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of frame-specific tag.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// List of tags in time period.
	TagSet []*MediaAiAnalysisFrameTagItem `json:"TagSet,omitempty" name:"TagSet" list`
}

type MediaAiAnalysisHighlightItem struct {

	// Address of an intelligently generated highlight.
	HighlightUrl *string `json:"HighlightUrl,omitempty" name:"HighlightUrl"`

	// Address of an intelligently generated highlight cover.
	CovImgUrl *string `json:"CovImgUrl,omitempty" name:"CovImgUrl"`

	// Confidence of an intelligently generated highlight between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Duration of an intelligently generated highlight.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// List of intelligently generated highlight subsegments, which together form a highlight.
	SegmentSet []*HighlightSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type MediaAiAnalysisTagItem struct {

	// Tag name.
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// Confidence of tag between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAnimatedGraphicsInfo struct {

	// Result information of animated image generating task
	// Note: this field may return null, indicating that no valid values can be obtained.
	AnimatedGraphicsSet []*MediaAnimatedGraphicsItem `json:"AnimatedGraphicsSet,omitempty" name:"AnimatedGraphicsSet" list`
}

type MediaAnimatedGraphicsItem struct {

	// Address of generated animated image.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Animated image generating template ID. For more information, please see [Animated Image Generating Parameter Template](https://cloud.tencent.com/document/product/266/33481#.3Cspan-id-.3D-.22zdt.22.3E.3C.2Fspan.3E.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Animated image format, such as gif.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Height of animated image in px.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Width of animated image in px.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Bitrate of animated image in bps.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Size of animated image in bytes.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// MD5 value of an animated image.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// Start time offset of animated image in video in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of animated image in video in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type MediaAudioStreamItem struct {

	// Bitrate of audio stream in bps.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Sample rate of audio stream in Hz.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SamplingRate *int64 `json:"SamplingRate,omitempty" name:"SamplingRate"`

	// Audio stream encoder, such as aac.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type MediaBasicInfo struct {

	// Media filename.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Media file description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time of media file in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last update time of media file (by an operation that triggers updating of media file information such as modifying video attributes or initiating video processing) in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Expiration time of media file in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I). After the expiration, the media file and its related resources (such as transcoding results and image sprites) will be permanently deleted. `9999-12-31T23:59:59Z` means "never expire".
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Category ID of media file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Category name of media file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// Category path to media file separated by "-", such as "new first-level category - new second-level category".
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// Cover image address of media file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// Media file container, such as mp4 and flv.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// URL of source media file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

	// Source information of media file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SourceInfo *MediaSourceData `json:"SourceInfo,omitempty" name:"SourceInfo"`

	// Storage region of media file, such as ap-guangzhou. For more information, please see [Region List](https://cloud.tencent.com/document/api/213/15692#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8).
	// Note: this field may return null, indicating that no valid values can be obtained.
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// Tag information of media file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`

	// Unique ID of LVB recording file
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vid *string `json:"Vid,omitempty" name:"Vid"`
}

type MediaContentReviewAsrTextSegmentItem struct {

	// Start time offset of suspected segment in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of suspected segment in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of suspected segment.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for suspected segment audit. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of suspected keywords.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet" list`
}

type MediaContentReviewOcrTextSegmentItem struct {

	// Start time offset of suspected segment in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of suspected segment in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of suspected segment.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for suspected segment audit. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of suspected keywords.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet" list`

	// Zone coordinates (at the pixel level) of suspected text: [x1, y1, x2, y2], i.e., the coordinates of the top-left and bottom-right corners.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// Expiration time of suspected image URL in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type MediaContentReviewPoliticalSegmentItem struct {

	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Score of a suspected politically sensitive segment.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for politically sensitive information detection of a suspected segment. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Name of a politically sensitive figure or violating photo.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Tag of politically sensitive information detection result of a suspected segment.
	Label *string `json:"Label,omitempty" name:"Label"`

	// URL of a suspected image (which will not be permanently stored
	//  and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// Zone coordinates (at the pixel level) of a politically sensitive figure or violating photo: [x1, y1, x2, y2], i.e., the coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`

	// This field has been disused. Please use `PicUrlExpireTime`.
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`

	// Expiration time of suspected image URL in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type MediaContentReviewSegmentItem struct {

	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Score of a suspected porn segment.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Tag of porn information detection result of a suspected segment.
	Label *string `json:"Label,omitempty" name:"Label"`

	// Suggestion for porn information detection of a suspected segment. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// URL of a suspected image (which will not be permanently stored
	//  and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// This field has been disused. Please use `PicUrlExpireTime`.
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`

	// Expiration time of suspected image URL in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type MediaImageSpriteInfo struct {

	// Information set of image sprites with specified specifications. Each element represents a set of image sprites with the same specification.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageSpriteSet []*MediaImageSpriteItem `json:"ImageSpriteSet,omitempty" name:"ImageSpriteSet" list`
}

type MediaImageSpriteItem struct {

	// Image sprite specification. For more information, please see [Image Sprite Parameter Template](https://cloud.tencent.com/document/product/266/33480#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Subimage height of image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Subimage width of image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Total number of subimages in each image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Address of each image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet" list`

	// Address of WebVtt file for the position-time relationship among subimages in an image sprite. The WebVtt file indicates the corresponding time points of each subimage and their coordinates in the image sprite, which is typically used by the player for implementing preview.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WebVttUrl *string `json:"WebVttUrl,omitempty" name:"WebVttUrl"`
}

type MediaInfo struct {

	// Basic information, such as video name, category, playback address, and cover image.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BasicInfo *MediaBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// Metadata, such as size, duration, video stream information, and audio stream information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// Result information of transcoding, such as address, specification, bitrate, and resolution of the videos in various bitrates generated by transcoding a video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TranscodeInfo *MediaTranscodeInfo `json:"TranscodeInfo,omitempty" name:"TranscodeInfo"`

	// Result information of animated image generating, i.e., relevant information of an animated image (such as .gif) generated from a video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AnimatedGraphicsInfo *MediaAnimatedGraphicsInfo `json:"AnimatedGraphicsInfo,omitempty" name:"AnimatedGraphicsInfo"`

	// Sampled screenshot information, i.e., relevant information of a sampled screenshot generated from a video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SampleSnapshotInfo *MediaSampleSnapshotInfo `json:"SampleSnapshotInfo,omitempty" name:"SampleSnapshotInfo"`

	// Image sprite information, i.e., relevant information of image sprite generated from video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageSpriteInfo *MediaImageSpriteInfo `json:"ImageSpriteInfo,omitempty" name:"ImageSpriteInfo"`

	// Time point screenshot information, i.e., information of each time point screenshot generated from a video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetInfo *MediaSnapshotByTimeOffsetInfo `json:"SnapshotByTimeOffsetInfo,omitempty" name:"SnapshotByTimeOffsetInfo"`

	// Timestamp information, i.e., information of each timestamp set for a video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyFrameDescInfo *MediaKeyFrameDescInfo `json:"KeyFrameDescInfo,omitempty" name:"KeyFrameDescInfo"`

	// Adaptive bitrate streaming information, such as specification, encryption type, and container format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdaptiveDynamicStreamingInfo *MediaAdaptiveDynamicStreamingInfo `json:"AdaptiveDynamicStreamingInfo,omitempty" name:"AdaptiveDynamicStreamingInfo"`

	// WeChat Mini Program audit information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MiniProgramReviewInfo *MediaMiniProgramReviewInfo `json:"MiniProgramReviewInfo,omitempty" name:"MiniProgramReviewInfo"`

	// Unique ID of media file.
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

type MediaInputInfo struct {

	// Video URL.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Video name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Custom video ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

type MediaKeyFrameDescInfo struct {

	// Information array of video timestamps.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyFrameDescSet []*MediaKeyFrameDescItem `json:"KeyFrameDescSet,omitempty" name:"KeyFrameDescSet" list`
}

type MediaKeyFrameDescItem struct {

	// Offset time of video timestamp in seconds.
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// Content string of timestamp containing 1–128 characters.
	Content *string `json:"Content,omitempty" name:"Content"`
}

type MediaMetaData struct {

	// Size of uploaded media file in bytes (which is the sum of size of m3u8 and ts files if the video is in HLS format).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Container, such as m4a and mp4.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Sum of the average bitrate of a video stream and that of an audio stream in bps.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Maximum value of the height of a video stream in px.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Maximum value of the width of a video stream in px.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Video duration in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Selected angle during video recording in degrees.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// Video stream information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet" list`

	// Audio stream information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet" list`

	// Video duration in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoDuration *float64 `json:"VideoDuration,omitempty" name:"VideoDuration"`

	// Audio duration in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioDuration *float64 `json:"AudioDuration,omitempty" name:"AudioDuration"`
}

type MediaMiniProgramReviewElem struct {

	// Audit type. 
	// <li>Porn: porn image,</li>
	// <li>Porn.Ocr: porn text,</li>
	// <li>Porn.Asr: porn speech,</li>
	// <li>Terrorism: terrorism image,</li>
	// <li>Political: politically sensitive image,</li>
	// <li>Political.Ocr: politically sensitive text</li>
	// <li>Political.Asr: politically sensitive speech</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Audit suggestion.
	// <li>pass: normal,</li>
	// <li>block: violating,</li>
	// <li>review: suspected of violation.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Confidence of audit result between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaMiniProgramReviewInfo struct {

	// Audit information list.
	MiniProgramReviewList []*MediaMiniProgramReviewInfoItem `json:"MiniProgramReviewList,omitempty" name:"MiniProgramReviewList" list`
}

type MediaMiniProgramReviewInfoItem struct {

	// Template ID, which is the ID of the transcoding template corresponding to the video published on WeChat Mini Program. 0 represents the source video.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Video metadata.
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// Video playback address for WeChat Mini Program audit
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Status of video release on WeChat Mini Program
	// <li>Pass: succeeded.</li>
	// <li>Rejected: rejected.</li>
	ReviewResult *string `json:"ReviewResult,omitempty" name:"ReviewResult"`

	// WeChat Mini Program audit element.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReviewSummary []*MediaMiniProgramReviewElem `json:"ReviewSummary,omitempty" name:"ReviewSummary" list`
}

type MediaOutputInfo struct {

	// Region of the bucket where an output file is stored, such as ap-guangzhou.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Bucket of output file.
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Path to output file, which must end in "/".
	Dir *string `json:"Dir,omitempty" name:"Dir"`
}

type MediaProcessTaskAdaptiveDynamicStreamingResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of adaptive bitrate streaming task.
	Input *AdaptiveDynamicStreamingTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of adaptive bitrate streaming task.
	Output *AdaptiveDynamicStreamingInfoItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskAnimatedGraphicResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of animated image generating task.
	Input *AnimatedGraphicTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of animated image generating task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *MediaAnimatedGraphicsItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskCoverBySnapshotResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of cover generating task.
	Input *CoverBySnapshotTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of cover generating task.
	Output *CoverBySnapshotTaskOutput `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskImageSpriteResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of image sprite generating task.
	Input *ImageSpriteTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of image sprite generating task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *MediaImageSpriteItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskInput struct {

	// List of transcoding tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TranscodeTaskSet []*TranscodeTaskInput `json:"TranscodeTaskSet,omitempty" name:"TranscodeTaskSet" list`

	// List of animated image generating tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInput `json:"AnimatedGraphicTaskSet,omitempty" name:"AnimatedGraphicTaskSet" list`

	// List of time point screencapturing tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTaskSet,omitempty" name:"SnapshotByTimeOffsetTaskSet" list`

	// List of sampled screencapturing tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SampleSnapshotTaskSet []*SampleSnapshotTaskInput `json:"SampleSnapshotTaskSet,omitempty" name:"SampleSnapshotTaskSet" list`

	// List of image sprite generating tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageSpriteTaskSet []*ImageSpriteTaskInput `json:"ImageSpriteTaskSet,omitempty" name:"ImageSpriteTaskSet" list`

	// List of cover generating tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoverBySnapshotTaskSet []*CoverBySnapshotTaskInput `json:"CoverBySnapshotTaskSet,omitempty" name:"CoverBySnapshotTaskSet" list`

	// List of adaptive bitrate streaming tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTaskSet,omitempty" name:"AdaptiveDynamicStreamingTaskSet" list`
}

type MediaProcessTaskResult struct {

	// Task type. Valid values:
	// <li>Transcode: transcoding</li>
	// <li>AnimatedGraphics: animated image generating</li>
	// <li>SnapshotByTimeOffset: time point screencapturing</li>
	// <li>SampleSnapshot: sampled screencapturing</li>
	// <li>ImageSprites: image sprite generating</li>
	// <li>CoverBySnapshot: Screencapturing for cover image</li>
	// <li>AdaptiveDynamicStreaming: adaptive bitrate streaming</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Query result of transcoding task, which is valid when task type is `Transcode`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitempty" name:"TranscodeTask"`

	// Query result of animated image generating task, which is valid when task type is `AnimatedGraphics`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitempty" name:"AnimatedGraphicTask"`

	// Query result of time point screencapturing task, which is valid when task type is `SnapshotByTimeOffset`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitempty" name:"SnapshotByTimeOffsetTask"`

	// Query result of sampled screencapturing task, which is valid when task type is `SampleSnapshot`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitempty" name:"SampleSnapshotTask"`

	// Query result of image sprite generating task, which is valid when task type is `ImageSprite`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitempty" name:"ImageSpriteTask"`

	// Query result of cover generating task, which is valid if task type is `CoverBySnapshot`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoverBySnapshotTask *MediaProcessTaskCoverBySnapshotResult `json:"CoverBySnapshotTask,omitempty" name:"CoverBySnapshotTask"`

	// Query result of adaptive bitrate streaming, which is valid if task type is `AdaptiveDynamicStreaming`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdaptiveDynamicStreamingTask *MediaProcessTaskAdaptiveDynamicStreamingResult `json:"AdaptiveDynamicStreamingTask,omitempty" name:"AdaptiveDynamicStreamingTask"`
}

type MediaProcessTaskSampleSnapshotResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of sampled screencapturing task.
	Input *SampleSnapshotTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of sampled screencapturing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *MediaSampleSnapshotItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskSnapshotByTimeOffsetResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of time point screencapturing task.
	Input *SnapshotByTimeOffsetTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of time point screencapturing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *MediaSnapshotByTimeOffsetItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskTranscodeResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of transcoding task.
	Input *TranscodeTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of transcoding task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *MediaTranscodeItem `json:"Output,omitempty" name:"Output"`
}

type MediaSampleSnapshotInfo struct {

	// Information set of sampled screenshots with the specified specifications. Each element represents a set of sampled screenshots with the same specification.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SampleSnapshotSet []*MediaSampleSnapshotItem `json:"SampleSnapshotSet,omitempty" name:"SampleSnapshotSet" list`
}

type MediaSampleSnapshotItem struct {

	// Sampled screenshot specification ID. For more information, please see [Sampled Screencapturing Parameter Template](https://cloud.tencent.com/document/product/266/33480#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Sample type. Valid values:
	// <li>Percent: samples at a specified percentage interval.</li>
	// <li>Time: samples at a specified time interval.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval
	// <li>If `SampleType` is `Percent`, this value means taking a screenshot at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, this value means taking a screenshot at an interval of the specified time (in seconds). The first screenshot is always the first video frame.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// List of URLs of generated screenshots.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet" list`

	// List of watermarking template IDs if the screenshots are watermarked.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition" list`
}

type MediaSnapshotByTimeOffsetInfo struct {

	// Information set of time point screenshots with a specified specification. Currently, there can be only one set of screenshots for each specification.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetSet []*MediaSnapshotByTimeOffsetItem `json:"SnapshotByTimeOffsetSet,omitempty" name:"SnapshotByTimeOffsetSet" list`
}

type MediaSnapshotByTimeOffsetItem struct {

	// Specification of a time point screenshot. For more information, please see [Parameter Template for Time Point Screencapturing](https://cloud.tencent.com/document/product/266/33480#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Information set of screenshots of the same specification. Each element represents a screenshot.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet,omitempty" name:"PicInfoSet" list`
}

type MediaSnapshotByTimePicInfoItem struct {

	// Time offset corresponding to the screenshot in the video in <font color=red>milliseconds</font>.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// Screenshot URL.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// List of watermarking template IDs if the screenshots are watermarked.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition" list`
}

type MediaSourceData struct {

	// Source category of media file:
	// <li>Record: recording, such as LVB recording and LVB time shifting recording.</li>
	// <li>Upload: upload, such as pull for upload, upload from server, and UCG upload from client.</li>
	// <li>VideoProcessing: video processing, such as video splicing and video clipping.</li>
	// <li>Unknown: unknown source.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// Field passed through when a file is created.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`
}

type MediaTrack struct {

	// Track type. Valid values:
	// <ul>
	// <li>Video: video track, which is composed of the following items: <ul><li>VideoTrackItem</li><li>MediaTransitionItem</li> <li>EmptyTrackItem</li></ul> </li>
	// <li>Audio: audio track, which is composed of the following items: <ul><li>AudioTrackItem</li><li>MediaTransitionItem</li><li>EmptyTrackItem</li></ul></li>
	// <li>Sticker: sticker track, which is composed of the following items: <ul><li> StickerTrackItem</li><li>EmptyTrackItem</li></ul></li>	
	// </ul>
	Type *string `json:"Type,omitempty" name:"Type"`

	// List of media segments on track.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TrackItems []*MediaTrackItem `json:"TrackItems,omitempty" name:"TrackItems" list`
}

type MediaTrackItem struct {

	// Segment type. Valid values:
	// <li>Video: video segment.</li>
	// <li>Audio: audio segment.</li>
	// <li>Sticker: sticker segment.</li>
	// <li>Transition: transition.</li>
	// <li>Empty: empty segment.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Video segment, which is valid if `Type` is `Video`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoItem *VideoTrackItem `json:"VideoItem,omitempty" name:"VideoItem"`

	// Audio segment, which is valid if `Type` is `Audio`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioItem *AudioTrackItem `json:"AudioItem,omitempty" name:"AudioItem"`

	// Sticker segment, which is valid if `Type` is `Sticker`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StickerItem *StickerTrackItem `json:"StickerItem,omitempty" name:"StickerItem"`

	// Transition, which is valid if `Type` is `Transition`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TransitionItem *MediaTransitionItem `json:"TransitionItem,omitempty" name:"TransitionItem"`

	// Empty segment, which is valid if `Type` is `Empty`. It is used as placeholder on time axis. <li>If you want a period of silence between two audio segments, you can use `EmptyTrackItem` to hold the place.</li>
	// <li>Use `EmptyTrackItem` as a placeholder to locate an item.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	EmptyItem *EmptyTrackItem `json:"EmptyItem,omitempty" name:"EmptyItem"`
}

type MediaTranscodeInfo struct {

	// Information set of transcoding with each specification. Each element represents a result of transcoding with a specification.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TranscodeSet []*MediaTranscodeItem `json:"TranscodeSet,omitempty" name:"TranscodeSet" list`
}

type MediaTranscodeItem struct {

	// Address of output video file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Transcoding specification ID. For more information, please see [Transcoding Parameter Template](https://cloud.tencent.com/document/product/266/33478#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Sum of the average bitrate of a video stream and that of an audio stream in bps.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Maximum value of the height of a video stream in px.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Maximum value of the width of a video stream in px.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Total size of a media file in bytes (which is the sum of size of m3u8 and ts files if the video is in HLS format).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Video duration in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Container, such as m4a and mp4.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Container *string `json:"Container,omitempty" name:"Container"`

	// MD5 value of video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// Audio stream information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet" list`

	// Video stream information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet" list`
}

type MediaTransitionItem struct {

	// Transition duration in seconds. For two media segments that use a transition, the start time of the second segment on the track will be automatically set to the end time of the first segment minus the transition duration.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// List of transition operations. Up to one video image or audio transition operation is supported.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Transitions []*TransitionOpertion `json:"Transitions,omitempty" name:"Transitions" list`
}

type MediaVideoStreamItem struct {

	// Bitrate of video stream in bps.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Height of video stream in px.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Width of video stream in px.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Video stream encoder, such as h264.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Frame rate in Hz.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`
}

type ModifyAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Video content analysis template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Video content analysis template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter of intelligent categorization task.
	ClassificationConfigure *ClassificationConfigureInfoForUpdate `json:"ClassificationConfigure,omitempty" name:"ClassificationConfigure"`

	// Control parameter of intelligent tagging task.
	TagConfigure *TagConfigureInfoForUpdate `json:"TagConfigure,omitempty" name:"TagConfigure"`

	// Control parameter of intelligent cover generating task.
	CoverConfigure *CoverConfigureInfoForUpdate `json:"CoverConfigure,omitempty" name:"CoverConfigure"`

	// Control parameter of intelligent frame-specific tagging task.
	FrameTagConfigure *FrameTagConfigureInfoForUpdate `json:"FrameTagConfigure,omitempty" name:"FrameTagConfigure"`

	// Control parameter of an intelligent highlight generating task.
	HighlightConfigure *HighlightsConfigureInfoForUpdate `json:"HighlightConfigure,omitempty" name:"HighlightConfigure"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAIAnalysisTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAIAnalysisTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAIAnalysisTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Name of an animated image generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Maximum value of the width (or long side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Animated image format. Valid values: gif, webp.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Video frame rate in Hz. Value range: [1, 30].
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Image quality. Value range: [1, 100]. Default value: 75.
	Quality *float64 `json:"Quality,omitempty" name:"Quality"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAnimatedGraphicsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Name of an image sprite generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subimage width of an image sprite in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Subimage height of an image sprite in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Sampling type. Valid values:
	// <li>Percent: by percent.</li>
	// <li>Time: by time interval.</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Subimage row count of an image sprite.
	RowCount *uint64 `json:"RowCount,omitempty" name:"RowCount"`

	// Subimage column count of an image sprite.
	ColumnCount *uint64 `json:"ColumnCount,omitempty" name:"ColumnCount"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

func (r *ModifyImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSpriteTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageSpriteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSpriteTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Sampled screencapturing type. Valid values:
	// <li>Percent: by percent.</li>
	// <li>Time: by time interval.</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Image format. Valid values: jpg, png.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

func (r *ModifySampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySampleSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySampleSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySampleSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a specified time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Name of a time point screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Image format. Valid values: jpg, png.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// ID of a [subapplication](/document/product/266/14574) in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

func (r *ModifySnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySnapshotByTimeOffsetTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubAppIdInfoRequest struct {
	*tchttp.BaseRequest

	// Subapplication ID.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Subapplication name. Length limit: 40 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subapplication overview. Length limit: 300 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifySubAppIdInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubAppIdInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubAppIdInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubAppIdInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubAppIdStatusRequest struct {
	*tchttp.BaseRequest

	// Subapplication ID.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Subapplication status. Valid values:
	// <li>On: enabled</li>
	// <li>Off: disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifySubAppIdStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubAppIdStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubAppIdStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubAppIdStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubAppIdStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of transcoding template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Container. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Transcoding template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 bytes.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Whether to remove video data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter.
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter.
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter.
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitempty" name:"TEHDConfig"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Watermarking template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Origin position. Valid values:
	// <li>TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>TopRight: the origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>BottomLeft: the origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>BottomRight: the origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the watermark will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width;</li>
	// <li>If the string ends in px, the `XPos` of the watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.</li>
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the watermark will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height;</li>
	// <li>If the string ends in px, the `YPos` of the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.</li>
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// Image watermarking template. This field is valid only for image watermarking templates.
	ImageTemplate *ImageWatermarkInputForUpdate `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is valid only for text watermarking templates.
	TextTemplate *TextWatermarkTemplateInputForUpdate `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is only valid for SVG watermarking templates.
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitempty" name:"SvgTemplate"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWatermarkTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Image watermark address. This field has a value only when `ImageTemplate.ImageContent` is not empty.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWatermarkTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MosaicInput struct {

	// Origin position, which currently can only be:
	// <li>TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the blur is in the top-left corner of the image or text.</li>
	// Default value: TopLeft.
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the blur relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the blur will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width;</li>
	// <li>If the string ends in px, the `XPos` of the blur will be the specified px; for example, `100px` means that `XPos` is 100 px.</li>
	// Default value: 0 px.
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// Vertical position of the origin of blur relative to the origin of coordinates of video. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the blur will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height;</li>
	// <li>If the string ends in px, the `YPos` of the blur will be the specified px; for example, `100px` means that `YPos` is 100 px.</li>
	// Default value: 0 px.
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// Blur width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the blur will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the blur will be in px; for example, `100px` means that `Width` is 100 px.</li>
	// Default value: 10%.
	Width *string `json:"Width,omitempty" name:"Width"`

	// Blur height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the blur will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the blur will be in px; for example, `100px` means that `Height` is 100 px.</li>
	// Default value: 10%.
	Height *string `json:"Height,omitempty" name:"Height"`

	// Start time offset of blur in seconds. If this parameter is left empty or 0 is entered, the blur will appear upon the first video frame.
	// <li>If this parameter is left empty or 0 is entered, the blur will appear upon the first video frame;</li>
	// <li>If this value is greater than 0 (e.g., n), the blur will appear at second n after the first video frame;</li>
	// <li>If this value is smaller than 0 (e.g., -n), the blur will appear at second n before the last video frame.</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of blur in seconds.
	// <li>If this parameter is left empty or 0 is entered, the blur will exist till the last video frame;</li>
	// <li>If this value is greater than 0 (e.g., n), the blur will exist till second n;</li>
	// <li>If this value is smaller than 0 (e.g., -n), the blur will exist till second n before the last video frame.</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type ObjectConfigureInfo struct {

	// Switch of object recognition task. Valid values:
	// <li>ON: enables intelligent object recognition task;</li>
	// <li>OFF: disables intelligent object recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Object library. Valid values:
	// <li>Default: default object library;</li>
	// <li>UserDefine: custom object library.</li>
	// <li>All: both default and custom object libraries will be used.</li>
	// Default value: All, i.e., both default and custom object libraries will be used.
	ObjectLibrary *string `json:"ObjectLibrary,omitempty" name:"ObjectLibrary"`
}

type OcrFullTextConfigureInfo struct {

	// Switch of full text recognition task. Valid values:
	// <li>ON: enables intelligent full text recognition task;</li>
	// <li>OFF: disables intelligent full text recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type OcrWordsConfigureInfo struct {

	// Switch of text keyword recognition task. Valid values:
	// <li>ON: enables text keyword recognition task;</li>
	// <li>OFF: disables text keyword recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`
}

type OutputAudioStream struct {

	// Audio stream encoder. Valid values:
	// <li>libfdk_aac: suitable for mp4 files.</li>
	// Default value: libfdk_aac.
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Audio stream sample rate. Valid values:
	// <li>16,000</li>
	// <li>32,000</li>
	// <li>44,100</li>
	// <li>48,000</li>
	// In Hz.
	// Default value: 16,000.
	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// Number of sound channels. Valid values:
	// <li>1: mono.</li>
	// <li>2: dual</li>
	// Default value: 2.
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type OutputVideoStream struct {

	// Video stream encoder. Valid values:
	// <li>libx264: H.264</li>
	// Default value: libx264.
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Video frame rate in Hz. Value range: [0, 60].
	// Default value: 0, which means that the value is the same as the video frame rate of the first video segment in the first video track.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`
}

type PoliticalAsrReviewTemplateInfo struct {

	// Switch of politically sensitive information detection in speech task. Valid values:
	// <li>ON: enables politically sensitive information detection in speech task;</li>
	// <li>OFF: disables politically sensitive information detection in speech task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalConfigureInfo struct {

	// Control parameter of politically sensitive information detection in video image.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImgReviewInfo *PoliticalImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Control parameter of politically sensitive information detection in speech.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AsrReviewInfo *PoliticalAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of politically sensitive information detection in text.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *PoliticalOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PoliticalImgReviewTemplateInfo struct {

	// Switch of politically sensitive information detection in video image task. Valid values:
	// <li>ON: enables politically sensitive information detection in video image task;</li>
	// <li>OFF: disables politically sensitive information detection in video image task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter tag for politically sensitive information detection in video image. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. Valid values:
	// <li>violation_photo: violating photo;</li>
	// <li>politician: politically sensitive figure;</li>
	// <li>entertainment: entertainment celebrity;</li>
	// <li>sport: sports celebrity.</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 97 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 95 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfo struct {

	// Switch of politically sensitive information detection in text task. Valid values:
	// <li>ON: enables politically sensitive information detection in text task;</li>
	// <li>OFF: disables politically sensitive information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfo struct {

	// Switch of porn information detection in speech task. Valid values:
	// <li>ON: enables porn information detection in speech task;</li>
	// <li>OFF: disables porn information detection in speech task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornConfigureInfo struct {

	// Control parameter of porn information detection in video image.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImgReviewInfo *PornImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Control parameter of porn information detection in speech.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AsrReviewInfo *PornAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of porn information detection in text.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *PornOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PornImgReviewTemplateInfo struct {

	// Switch of porn information detection in video image task. Valid values:
	// <li>ON: enables porn information detection in video image task;</li>
	// <li>OFF: disables porn information detection in video image task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter tag for porn information detection in video image. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. Valid values:
	// <li>porn: porn;</li>
	// <li>vulgar: vulgarity;</li>
	// <li>intimacy: intimacy;</li>
	// <li>sexy: sexiness.</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 90 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 0 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfo struct {

	// Switch of porn information detection in text task. Valid values:
	// <li>ON: enables porn information detection in text task;</li>
	// <li>OFF: disables porn information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type ProcedureTask struct {

	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task flow status. Valid values:
	// <li>PROCESSING: processing;</li>
	// <li>FINISH: completed.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Disused. Please use `ErrCode` of each specific task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Disused. Please use `Message` of each specific task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Media file ID.
	// <li>If the task flow is initiated by [ProcessMedia](https://cloud.tencent.com/document/product/266/33427), this field means the `FileId` in [MediaInfo](https://cloud.tencent.com/document/product/266/31773#MediaInfo);</li>
	// <li>If the task flow is initiated by [ProcessMediaByUrl](https://cloud.tencent.com/document/product/266/33426), this field means the `Id` in [MediaInputInfo](https://cloud.tencent.com/document/product/266/31773#MediaInputInfo).</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Media filename
	// <li>If the task flow is initiated by [ProcessMedia](https://cloud.tencent.com/document/product/266/33427), this field means the `BasicInfo.Name` in [MediaInfo](https://cloud.tencent.com/document/product/266/31773#MediaInfo);</li>
	// <li>If the task flow is initiated by [ProcessMediaByUrl](https://cloud.tencent.com/document/product/266/33426), this field means the `Name` in [MediaInputInfo](https://cloud.tencent.com/document/product/266/31773#MediaInputInfo).</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Media file address
	// <li>If the task flow is initiated by [ProcessMedia](https://cloud.tencent.com/document/product/266/33427), this field means the `BasicInfo.MediaUrl` in [MediaInfo](https://cloud.tencent.com/document/product/266/31773#MediaInfo);</li>
	// <li>If the task flow is initiated by [ProcessMediaByUrl](https://cloud.tencent.com/document/product/266/33426), this field means the `Url` in [MediaInputInfo](https://cloud.tencent.com/document/product/266/31773#MediaInputInfo).</li>
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// Source video metadata.
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// Execution status and result of video processing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MediaProcessResultSet []*MediaProcessTaskResult `json:"MediaProcessResultSet,omitempty" name:"MediaProcessResultSet" list`

	// Execution status and result of video content audit task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AiContentReviewResultSet []*AiContentReviewResult `json:"AiContentReviewResultSet,omitempty" name:"AiContentReviewResultSet" list`

	// Execution status and result of video content analysis task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AiAnalysisResultSet []*AiAnalysisResult `json:"AiAnalysisResultSet,omitempty" name:"AiAnalysisResultSet" list`

	// Execution status and result of video content recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AiRecognitionResultSet []*AiRecognitionResult `json:"AiRecognitionResultSet,omitempty" name:"AiRecognitionResultSet" list`

	// Task flow priority. Value range: [-10, 10].
	// Note: this field may return null, indicating that no valid values can be obtained.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for change in task flow status.
	// <li>Finish: an event notification will be initiated only after the task flow is completely executed;</li>
	// <li>Change: an event notification will be initiated as soon as the status of a subtask in the task flow changes; </li>
	// <li>None: no callback for the task flow will be accepted.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type ProcedureTemplate struct {

	// Task flow name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type of a task flow template. Valid values:
	// <li>Preset: preset task flow template;</li>
	// <li>Custom: custom task flow template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Parameter of video processing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Parameter of AI-based content audit task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Parameter of AI-based content analysis task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of AI-based content recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Parameter of a release on WeChat Mini Program task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MiniProgramPublishTask *WechatMiniProgramPublishTaskInput `json:"MiniProgramPublishTask,omitempty" name:"MiniProgramPublishTask"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ProcessMediaByProcedureRequest struct {
	*tchttp.BaseRequest

	// Media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// [Task flow template](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF) name.
	ProcedureName *string `json:"ProcedureName,omitempty" name:"ProcedureName"`

	// Task flow priority. The higher the value, the higher the priority. Value range: -10–10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for task flow status change. Valid values: Finish, Change, None. If this parameter is left empty, `Finish` will be used.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// ID used for deduplication. If there was a request with the same ID on the last day, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ProcessMediaByProcedureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByProcedureRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaByProcedureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaByProcedureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByProcedureResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaByUrlRequest struct {
	*tchttp.BaseRequest

	// Information of input video, including video's URL, name, and custom ID.
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// Information of COS path to output file.
	OutputInfo *MediaOutputInfo `json:"OutputInfo,omitempty" name:"OutputInfo"`

	// Type parameter of video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Task flow priority. The higher the value, the higher the priority. Value range: -10–10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for task flow status change. Valid values: Finish, Change, None. If this parameter is left empty, `Finish` will be used.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ProcessMediaByUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaByUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaByUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaRequest struct {
	*tchttp.BaseRequest

	// Media file ID, i.e., the globally unique ID of a file in VOD assigned by the VOD backend after successful upload. This field can be obtained through the [video upload completion event notification](/document/product/266/7830) or [VOD Console](https://console.cloud.tencent.com/vod/media).
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Parameter of video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Type parameter of video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Task flow priority. The higher the value, the higher the priority. Value range: -10–10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for task flow status change. Valid values: Finish, Change, None. If this parameter is left empty, `Finish` will be used.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ProcessMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProhibitedAsrReviewTemplateInfo struct {

	// Switch of prohibited information detection in speech task. Valid values:
	// <li>ON: enables prohibited information detection in speech task;</li>
	// <li>OFF: disables prohibited information detection in speech task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type ProhibitedConfigureInfo struct {

	// Control parameter of prohibited information detection in speech.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AsrReviewInfo *ProhibitedAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of prohibited information detection in text.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *ProhibitedOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type ProhibitedOcrReviewTemplateInfo struct {

	// Switch of prohibited information detection in text task. Valid values:
	// <li>ON: enables prohibited information detection in text task;</li>
	// <li>OFF: disables prohibited information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PullEventsRequest struct {
	*tchttp.BaseRequest

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *PullEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of events.
	// Note: this field may return null, indicating that no valid values can be obtained.
		EventSet []*EventContent `json:"EventSet,omitempty" name:"EventSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullUploadRequest struct {
	*tchttp.BaseRequest

	// URL of the media to be pulled. Media files in HLS and Dash formats cannot be pulled currently.
	// For the supported extensions, please see [Media Types](https://cloud.tencent.com/document/product/266/9760#.E5.AA.92.E4.BD.93.E7.B1.BB.E5.9E.8B).
	MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

	// Media name.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// URL of video cover to be pulled. Only gif, jpeg, and png formats are supported.
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// Subsequent task for media. For more information, please see [Specifying Task Flow After Upload](https://cloud.tencent.com/document/product/266/9759).
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Expiration time of media file in ISO 8601 format. For more information, please see [Notes on ISO Date Format](https://cloud.tencent.com/document/product/266/11732#I).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Specifies upload region. This is only applicable to users that have special requirements for the upload region (currently, only Beijing, Shanghai, and Chongqing regions are supported).
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [CreateClass](https://cloud.tencent.com/document/product/266/7812) API.
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// The source context which is used to pass through the user request information. After `Procedure` is specified, the task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *PullUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Video pull for upload task ID, which can be used to query the status of pull for upload task.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullUploadTask struct {

	// Pull for upload task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task flow status. Valid values:
	// <li>PROCESSING: processing;</li>
	// <li>FINISH: completed.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// ID of video generated after pull for upload is completed.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Basic information of media file generated after pull for upload is completed.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MediaBasicInfo *MediaBasicInfo `json:"MediaBasicInfo,omitempty" name:"MediaBasicInfo"`

	// Playback address generated after pull for upload is completed.
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// If a video processing flow is specified when a video is pulled for upload, this parameter will be the ID of the task flow.
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type PushUrlCacheRequest struct {
	*tchttp.BaseRequest

	// List of prefetched URLs. Up to 20 ones can be specified at a time.
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *PushUrlCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PushUrlCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PushUrlCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetProcedureTemplateRequest struct {
	*tchttp.BaseRequest

	// Task flow name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter of video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Parameter of AI-based content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Parameter of AI-based content analysis task.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of AI-based content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ResetProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetProcedureTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetProcedureTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SampleSnapshotTaskInput struct {

	// Sampled screencapturing template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of up to 10 image or text watermarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type SampleSnapshotTemplate struct {

	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of a sampled screencapturing template.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Image format.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Sampled screencapturing type.
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type SearchMediaRequest struct {
	*tchttp.BaseRequest

	// Search text, which fuzzily matches the media file name or description. The more matching items and the higher the match rate, the higher-ranked the result. It can contain up to 64 characters.
	Text *string `json:"Text,omitempty" name:"Text"`

	// Tag set, which matches any element in the set.
	// <li>Tag length limit: 8 characters.</li>
	// <li>Array length limit: 10.</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// Category ID set, which matches the categories of the specified IDs and all subcategories. Array length limit: 10.
	ClassIds []*int64 `json:"ClassIds,omitempty" name:"ClassIds" list`

	// Start time in the creation time range.
	// <li>After or at the start time.</li>
	// <li>In ISO 8601 format. For more information, please see [Notes on ISO Date Format](https://cloud.tencent.com/document/product/266/11732#I).</li>
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the creation time range.
	// <li>Before the end time.</li>
	// <li>In ISO 8601 format. For more information, please see [Notes on ISO Date Format](https://cloud.tencent.com/document/product/266/11732#I).</li>
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Media file source. For valid values, please see [SourceType](https://cloud.tencent.com/document/product/266/31773#MediaSourceData).
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// [LVB code](https://cloud.tencent.com/document/product/267/5959) of a stream.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Unique ID of LVB recording file.
	Vid *string `json:"Vid,omitempty" name:"Vid"`

	// Sorting order.
	// <li>Valid value of `Sort.Field`: CreateTime</li>
	// <li>If `Text` is specified for the search, the results will be sorted by the match rate, and this field will not take effect</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// <div id="p_offset">Start offset of a paged return. Default value: 0. Entries from No. "Offset" to No. "Offset + Limit - 1" will be returned.
	// <li>Value range: "Offset + Limit" cannot be more than 5,000. (For more information, please see <a href="#maxResultsDesc">Limit on the Number of Results Returned by API</a>)</li></div>
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// <div id="p_limit">Number of entries returned by a paged query. Default value: 10. Entries from No. "Offset" to No. "Offset + Limit - 1" will be returned.
	// <li>Value range: "Offset + Limit" cannot be more than 5,000. (For more information, please see <a href="#maxResultsDesc">Limit on the Number of Results Returned by API</a>)</li></div>
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *SearchMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
	// <li>Maximum value: 5000. If the number of eligible entries is greater than 5,000, this field will return 5,000 instead of the actual number.</li>
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of media file information, only including the basic information (BasicInfo).
	// Note: this field may return null, indicating that no valid values can be obtained.
		MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SegmentConfigureInfo struct {

	// Switch of video splitting recognition task. Valid values:
	// <li>ON: enables intelligent video splitting recognition task;</li>
	// <li>OFF: disables intelligent video splitting recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type SimpleHlsClipRequest struct {
	*tchttp.BaseRequest

	// URL of the HLS video in VOD that needs to be clipped.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Start offset time of clipping in seconds. Default value: 0, which means to clip from the beginning of the video. A negative number indicates how many seconds from the end of the video clipping will start at. For example, -10 means that clipping will start at the 10th second from the end.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End offset time of clipping in seconds. Default value: 0, which means to clip till the end of the video. A negative number indicates how many seconds from the end of the video clipping will end. For example, -10 means that clipping will end at the 10th second from the end.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *SimpleHlsClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SimpleHlsClipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SimpleHlsClipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Address of clipped video.
		Url *string `json:"Url,omitempty" name:"Url"`

		// Metadata of clipped video. Currently, `Size`, `Rotate`, `VideoDuration`, and `AudioDuration` fields use default value with no actual data.
		MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SimpleHlsClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SimpleHlsClipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SnapshotByTimeOffset2017 struct {

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Specific time point of screenshot in milliseconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeOffset *uint64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// Address of output screenshot file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`
}

type SnapshotByTimeOffsetTask2017 struct {

	// Screencapturing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Screenshot file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// screenshot specification. For more information, please see [Parameter Template for Time Point Screencapturing](https://cloud.tencent.com/document/product/266/33480#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Result information of screencapturing.
	SnapshotInfoSet []*SnapshotByTimeOffset2017 `json:"SnapshotInfoSet,omitempty" name:"SnapshotInfoSet" list`
}

type SnapshotByTimeOffsetTaskInput struct {

	// Time point screencapturing template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of time points for screencapturing in <font color=red>milliseconds</font>.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitempty" name:"TimeOffsetSet" list`

	// List of up to 10 image or text watermarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type SnapshotByTimeOffsetTemplate struct {

	// Unique ID of a specified time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of a specified time point screencapturing template.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Image format.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type SortBy struct {

	// Sort by field
	Field *string `json:"Field,omitempty" name:"Field"`

	// Sorting order. Valid values: Asc (ascending), Desc (descending)
	Order *string `json:"Order,omitempty" name:"Order"`
}

type SpecificationDataItem struct {

	// Task specification.
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// Statistics.
	Data []*TaskStatDataItem `json:"Data,omitempty" name:"Data" list`
}

type StatDataItem struct {

	// Start time of data time range in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I). For example, if the time granularity is 1-day, `2018-12-01T00:00:00+08:00` represents the time range between December 1, 2018 (inclusive) and December 2, 2018 (not inclusive).
	// <li>For data at hourly level, `2019-08-22T00:00:00+08:00` indicates the statistics between 00:00 and 01:00 AM on August 22, 2019.</li>
	// <li>For data at daily level, `2019-08-22T00:00:00+08:00` indicates statistics on August 22, 2019.</li>
	Time *string `json:"Time,omitempty" name:"Time"`

	// Data size.
	// <li>Storage capacity in bytes.</li>
	// <li>Transcoding duration in seconds.</li>
	// <li>Traffic in bytes.</li>
	// <li>Bandwidth in bps.</li>
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type StickerTrackItem struct {

	// Source of media file for sticker material, which can be an ID of a VOD file or URL of another file.
	SourceMedia *string `json:"SourceMedia,omitempty" name:"SourceMedia"`

	// Sticker duration in seconds.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Start time of sticker on track in seconds.
	StartTime *float64 `json:"StartTime,omitempty" name:"StartTime"`

	// Origin position. Valid values:
	// <li> Center: the origin of coordinates is the center position, such as the center of canvas.</li>
	// Default value: Center.
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the sticker relative to the origin of the canvas. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the sticker will be at the position of the specified percentage of the canvas width; for example, `10%` means that `XPos` is 10% of the canvas width.</li><li>If the string ends in px, the `XPos` of the sticker will be in px; for example, `100px` means that `XPos` is 100 px.</li>
	// Default value: 0 px.
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// The vertical position of the origin of the sticker relative to the origin of the canvas. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the sticker will be at the position of the specified percentage of the canvas height; for example, `10%` means that `YPos` is 10% of the canvas height.</li>
	// <li>If the string ends in px, the `YPos` of the sticker will be in px; for example, `100px` means that `YPos` is 100 px.</li>
	// Default value: 0 px.
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// Sticker width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the sticker will be the specified percentage of the canvas width; for example, `10%` means that `Width` is 10% of the canvas width.</li>
	// <li>If the string ends in px, the `Width` of the sticker will be in px; for example, `100px` means that `Width` is 100 px.</li>
	// <li>If both `Width` and `Height` are empty, then they will be the `Width` and `Height` of the sticker material, respectively.</li>
	// <li>If `Width` is empty (0), but `Height` is not empty, `Width` will be proportionally scaled.</li>
	// <li>If `Width` is not empty, but `Height` is empty, `Height` will be proportionally scaled.</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// Sticker height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the sticker will be the specified percentage of the canvas height; for example, `10%` means that `Height` is 10% of the canvas height.</li>
	// <li>If the string ends in px, the `Height` of the sticker will be in px; for example, `100px` means that `Height` is 100 px.</li>
	// <li>If both `Width` and `Height` are empty, then they will be the `Width` and `Height` of the sticker material, respectively.</li>
	// <li>If `Width` is empty, but `Height` is not empty, `Width` will be proportionally scaled.</li>
	// <li>If `Width` is not empty, but `Height` is empty, `Height` will be proportionally scaled.</li>
	Height *string `json:"Height,omitempty" name:"Height"`

	// Operation on sticker such as image rotation.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageOperations []*ImageTransform `json:"ImageOperations,omitempty" name:"ImageOperations" list`
}

type SubAppIdInfo struct {

	// Subapplication ID.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Subapplication name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subapplication overview.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Subapplication creation time of task in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Subapplication status. Valid values:
	// <li>On: enabled;</li>
	// <li>Off: disabled.</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type SvgWatermarkInput struct {

	// Watermark width, which supports six formats of px, %, W%, H%, S%, and L%:
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px; if `0px` is entered
	//  and `Height` is not `0px`, the watermark width will be proportionally scaled based on the source SVG image; if `0px` is entered for both `Width` and `Height`, the watermark width will be the width of the source SVG image;</li>
	// <li>If the string ends in `W%`, the `Width` of the watermark will be the specified percentage of the video width; for example, `10W%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in `H%`, the `Width` of the watermark will be the specified percentage of the video height; for example, `10H%` means that `Width` is 10% of the video height;</li>
	// <li>If the string ends in `S%`, the `Width` of the watermark will be the specified percentage of the short side of the video; for example, `10S%` means that `Width` is 10% of the short side of the video;</li>
	// <li>If the string ends in `L%`, the `Width` of the watermark will be the specified percentage of the long side of the video; for example, `10L%` means that `Width` is 10% of the long side of the video;</li>
	// <li>If the string ends in %, the meaning is the same as `W%`.</li>
	// Default value: 10W%.
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height, which supports six formats of px, %, W%, H%, S%, and L%:
	// <li>If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px; if `0px` is entered
	//  and `Width` is not `0px`, the watermark height will be proportionally scaled based on the source SVG image; if `0px` is entered for both `Width` and `Height`, the watermark height will be the height of the source SVG image;</li>
	// <li>If the string ends in `W%`, the `Height` of the watermark will be the specified percentage of the video width; for example, `10W%` means that `Height` is 10% of the video width;</li>
	// <li>If the string ends in `H%`, the `Height` of the watermark will be the specified percentage of the video height; for example, `10H%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in `S%`, the `Height` of the watermark will be the specified percentage of the short side of the video; for example, `10S%` means that `Height` is 10% of the short side of the video;</li>
	// <li>If the string ends in `L%`, the `Height` of the watermark will be the specified percentage of the long side of the video; for example, `10L%` means that `Height` is 10% of the long side of the video;</li>
	// <li>If the string ends in %, the meaning is the same as `H%`.</li>
	// Default value: 0 px.
	Height *string `json:"Height,omitempty" name:"Height"`
}

type SvgWatermarkInputForUpdate struct {

	// Watermark width, which supports six formats of px, %, W%, H%, S%, and L%:
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px; if `0px` is entered
	//  and `Height` is not `0px`, the watermark width will be proportionally scaled based on the source SVG image; if `0px` is entered for both `Width` and `Height`, the watermark width will be the width of the source SVG image;</li>
	// <li>If the string ends in `W%`, the `Width` of the watermark will be the specified percentage of the video width; for example, `10W%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in `H%`, the `Width` of the watermark will be the specified percentage of the video height; for example, `10H%` means that `Width` is 10% of the video height;</li>
	// <li>If the string ends in `S%`, the `Width` of the watermark will be the specified percentage of the short side of the video; for example, `10S%` means that `Width` is 10% of the short side of the video;</li>
	// <li>If the string ends in `L%`, the `Width` of the watermark will be the specified percentage of the long side of the video; for example, `10L%` means that `Width` is 10% of the long side of the video;</li>
	// <li>If the string ends in %, the meaning is the same as `W%`.</li>
	// Default value: 10W%.
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height, which supports six formats of px, %, W%, H%, S%, and L%:
	// <li>If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px; if `0px` is entered
	//  and `Width` is not `0px`, the watermark height will be proportionally scaled based on the source SVG image; if `0px` is entered for both `Width` and `Height`, the watermark height will be the height of the source SVG image;</li>
	// <li>If the string ends in `W%`, the `Height` of the watermark will be the specified percentage of the video width; for example, `10W%` means that `Height` is 10% of the video width;</li>
	// <li>If the string ends in `H%`, the `Height` of the watermark will be the specified percentage of the video height; for example, `10H%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in `S%`, the `Height` of the watermark will be the specified percentage of the short side of the video; for example, `10S%` means that `Height` is 10% of the short side of the video;</li>
	// <li>If the string ends in `L%`, the `Height` of the watermark will be the specified percentage of the long side of the video; for example, `10L%` means that `Height` is 10% of the long side of the video;</li>
	// <li>If the string ends in %, the meaning is the same as `H%`.
	// Default value: 0 px.
	Height *string `json:"Height,omitempty" name:"Height"`

	// Watermark cycle configuration, which is used to configure watermarks so that they will be displayed and hidden periodically.
	// Primary use case: watermarks can be added at various positions in a video, which are displayed and hidden periodically to prevent them from being covered.
	// For example, watermarks A, B, C, and D are set in the top-left corner, top-right corner, bottom-right corner, and bottom-left corner of a video, respectively. After the first video frame, { A will be displayed for 5s -> B for 5s -> C for 5s -> D for 5s } -> A for 5s -> B for 5s -> ... Only one watermark will be visible at any time.
	// Within the braces ({}) is a major cycle composed of four watermarks, namely, A, B, C, and D, which lasts for 20 seconds in a cycle.
	// Watermarks A, B, C, and D are displayed periodically for 5 seconds and hidden for 15 seconds each in a fixed order.
	// This configuration item is used to describe the cycle configuration of a single watermark.
	CycleConfig *WatermarkCycleConfigForUpdate `json:"CycleConfig,omitempty" name:"CycleConfig"`
}

type TEHDConfig struct {

	// TESHD type. Valid values:
	// <li>TEHD-100: TESHD-100.</li>
	// If this parameter is left blank, TESHD will not be enabled.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Maximum bitrate, which is valid when `Type` is `TESHD`.
	// If this parameter is left blank or 0 is entered, there will be no upper limit for bitrate.
	MaxVideoBitrate *uint64 `json:"MaxVideoBitrate,omitempty" name:"MaxVideoBitrate"`
}

type TEHDConfigForUpdate struct {

	// TESHD type. Valid values:
	// <li>TEHD-100: TESHD-100.</li>
	// If this parameter is left blank, no modification will be made.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Maximum bitrate. If this parameter is left blank, no modification will be made.
	MaxVideoBitrate *uint64 `json:"MaxVideoBitrate,omitempty" name:"MaxVideoBitrate"`
}

type TagConfigureInfo struct {

	// Switch of intelligent tagging task. Valid values:
	// <li>ON: enables intelligent tagging task;</li>
	// <li>OFF: disables intelligent tagging task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type TagConfigureInfoForUpdate struct {

	// Switch of intelligent tagging task. Valid values:
	// <li>ON: enables intelligent tagging task;</li>
	// <li>OFF: disables intelligent tagging task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type TaskStatData struct {

	// Task type
	// <li>Transcode: transcoding</li>
	// <li>Snapshot: screencapturing</li>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Task statistics overview.
	// <li>Transcode: usage in seconds</li>
	Summary []*TaskStatDataItem `json:"Summary,omitempty" name:"Summary" list`

	// Detailed statistics of tasks with different specifications.
	// Transcoding specification:
	// <li>Remuxing: remuxing</li>
	// <li>Audio: audio transcoding</li>
	// <li>Standard.H264.SD: H.264 SD transcoding</li>
	// <li>Standard.H264.HD: H.264 HD transcoding</li>
	// <li>Standard.H264.FHD: H.264 FHD transcoding</li>
	// <li>Standard.H264.2K: H.264 2K transcoding</li>
	// <li>Standard.H264.4K: H.264 4K transcoding</li>
	// <li>Standard.H265.SD: H.265 SD transcoding</li>
	// <li>Standard.H265.HD: H.265 HD transcoding</li>
	// <li>Standard.H265.FHD: H.265 FHD transcoding</li>
	// <li>Standard.H265.2K: H.265 2K transcoding</li>
	// <li>Standard.H265.4K: H.265 4K transcoding</li>
	// <li>TESHD-10.H264.SD: H.264 SD TESHD transcoding</li>
	// <li>TESHD-10.H264.HD: H.264 HD TESHD transcoding</li>
	// <li>TESHD-10.H264.FHD: H.264 FHD TESHD transcoding</li>
	// <li>TESHD-10.H264.2K: H.264 2K TESHD transcoding</li>
	// <li>TESHD-10.H264.4K: H.264 4K TESHD transcoding</li>
	// <li>TESHD-10.H265.SD: H.265 SD TESHD transcoding</li>
	// <li>TESHD-10.H265.HD: H.265 HD TESHD transcoding</li>
	// <li>TESHD-10.H265.FHD: H.265 FHD TESHD transcoding</li>
	// <li>TESHD-10.H265.2K: H.265 2K TESHD transcoding</li>
	// <li>TESHD-10.H265.4K: H.265 4K TESHD transcoding</li>
	// <li>Edit.Audio: audio editing</li>
	// <li>Edit.H264.SD: H.264 SD video editing</li>
	// <li>Edit.H264.HD: H.264 HD video editing</li>
	// <li>Edit.H264.FHD: H.264 FHD video editing</li>
	// <li>Edit.H264.2K: H.264 2K video editing</li>
	// <li>Edit.H264.4K: H.264 4K video editing</li>
	// <li>Edit.H265.SD: H.265 SD video editing</li>
	// <li>Edit.H265.HD: H.265 HD video editing</li>
	// <li>Edit.H265.FHD: H.265 FHD video editing</li>
	// <li>Edit.H265.2K: H.265 2K video editing</li>
	// <li>Edit.H265.4K: H.265 4K video editing</li>
	Details []*SpecificationDataItem `json:"Details,omitempty" name:"Details" list`
}

type TaskStatDataItem struct {

	// Start time of data time range in [ISO date format](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F). For example, if the time granularity is 1-day, `2018-12-01T00:00:00+08:00` represents the time range between December 1, 2018 (inclusive) and December 2, 2018 (not inclusive).
	Time *string `json:"Time,omitempty" name:"Time"`

	// Number of tasks.
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Task usage.
	Usage *int64 `json:"Usage,omitempty" name:"Usage"`
}

type TempCertificate struct {

	// Temporary security certificate ID.
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// Temporary security certificate `Key`.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Token value.
	Token *string `json:"Token,omitempty" name:"Token"`

	// Certificate expiration time. A Unix timestamp will be returned which is accurate down to the second.
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}

type TerrorismConfigureInfo struct {

	// Control parameter of terrorism information detection in video image task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImgReviewInfo *TerrorismImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Control parameter of terrorism information detection in text task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *TerrorismOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type TerrorismImgReviewTemplateInfo struct {

	// Switch of terrorism information detection in video image task. Valid values:
	// <li>ON: enables terrorism information detection in video image task;</li>
	// <li>OFF: disables terrorism information detection in video image task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter tag for terrorism information detection in video image. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. Valid values:
	// <li>guns: weapons and guns;</li>
	// <li>crowd: crowd;</li>
	// <li>bloody: bloody scenes;</li>
	// <li>police: police force;</li>
	// <li>banners: terrorism flags;</li>
	// <li>militant: militants;</li>
	// <li>explosion: explosions and fires;</li>
	// <li>terrorists: terrorists.</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 90 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 80 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type TerrorismOcrReviewTemplateInfo struct {

	// Switch of terrorism information detection in text task. Valid values:
	// <li>ON: enables terrorism information detection in text task;</li>
	// <li>OFF: disables terrorism information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type TextWatermarkTemplateInput struct {

	// Font type. Currently, two types are supported:
	// <li>simkai.ttf: both Chinese and English are supported;</li>
	// <li>arial.ttf: only English is supported.</li>
	FontType *string `json:"FontType,omitempty" name:"FontType"`

	// Font size in Npx format where N is a numeric value.
	FontSize *string `json:"FontSize,omitempty" name:"FontSize"`

	// Font color in 0xRRGGBB format. Default value: 0xFFFFFF (white).
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// Text transparency. Value range: (0, 1]
	// <li>0: completely transparent</li>
	// <li>1: completely opaque</li>
	// Default value: 1.
	FontAlpha *float64 `json:"FontAlpha,omitempty" name:"FontAlpha"`
}

type TextWatermarkTemplateInputForUpdate struct {

	// Font type. Currently, two types are supported:
	// <li>simkai.ttf: both Chinese and English are supported;</li>
	// <li>arial.ttf: only English is supported.</li>
	FontType *string `json:"FontType,omitempty" name:"FontType"`

	// Font size in Npx format where N is a numeric value.
	FontSize *string `json:"FontSize,omitempty" name:"FontSize"`

	// Font color in 0xRRGGBB format. Default value: 0xFFFFFF (white).
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// Text transparency. Value range: (0, 1]
	// <li>0: completely transparent</li>
	// <li>1: completely opaque</li>
	FontAlpha *float64 `json:"FontAlpha,omitempty" name:"FontAlpha"`
}

type TranscodePlayInfo2017 struct {

	// Playback address.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Transcoding specification ID. For more information, please see [Transcoding Parameter Template](https://cloud.tencent.com/document/product/266/33478#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Sum of the average bitrate of a video stream and that of an audio stream in bps.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Maximum value of the height of a video stream in px.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Maximum value of the width of a video stream in px.
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type TranscodeTask2017 struct {

	// Transcoding task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// ID of transcoded file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Name of transcoded file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Video duration in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Cover address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// Playback information generated after video transcoding.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlayInfoSet []*TranscodePlayInfo2017 `json:"PlayInfoSet,omitempty" name:"PlayInfoSet" list`
}

type TranscodeTaskInput struct {

	// Video transcoding template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of up to 10 image or text watermarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`

	// List of blurs. Up to 10 ones can be supported.
	MosaicSet []*MosaicInput `json:"MosaicSet,omitempty" name:"MosaicSet" list`
}

type TranscodeTemplate struct {

	// Unique ID of transcoding template.
	Definition *string `json:"Definition,omitempty" name:"Definition"`

	// Container. Valid values: mp4, flv, hls, mp3, flac, ogg.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Transcoding template name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Whether to remove video data. Valid values:
	// <li>0: retain;</li>
	// <li>1: remove.</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: retain;</li>
	// <li>1: remove.</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter. This field is valid only when `RemoveVideo` is 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter. This field is valid only when `RemoveAudio` is 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitempty" name:"TEHDConfig"`

	// Container filter. Valid values:
	// <li>Video: video container that can contain both video stream and audio stream;</li>
	// <li>PureAudio: audio container that can contain only audio stream.</li>
	ContainerType *string `json:"ContainerType,omitempty" name:"ContainerType"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TransitionOpertion struct {

	// Transition type. Valid values:
	// <ul>
	// <li>Video image transition operation, which is used for transition with video image between two video segments:
	// <ul>
	// <li>ImageFadeInFadeOut: video image fade-in/fade-out.</li>
	// <li>BowTieHorizontal: horizontal bow.</li>
	// <li>BowTieVertical: vertical bow.</li>
	// <li>ButterflyWaveScrawler: waggling.</li>
	// <li>Cannabisleaf: maple leaf.</li>
	// <li> Circle: curved circling.</li>
	// <li>CircleCrop: circle gathering.</li>
	// <li>Circleopen: elliptic gathering.</li>
	// <li>Crosswarp: horizontal warping.</li>
	// <li>Cube: cube.</li>
	// <li>DoomScreenTransition: curtain.</li>
	// <li>Doorway: doorway.</li>
	// <li>Dreamy: wave.</li>
	// <li>DreamyZoom: horizontal gathering.</li>
	// <li>FilmBurn: evening glow.</li>
	// <li>GlitchMemories: joggling.</li>
	// <li>Heart: heart.</li>
	// <li>InvertedPageCurl: page turning.</li>
	// <li>Luma: corroding.</li>
	// <li>Mosaic: grid.</li>
	// <li>Pinwheel: pinwheel.</li>
	// <li>PolarFunction: elliptic diffusing.</li>
	// <li>PolkaDotsCurtain: curved diffusing.</li>
	// <li>Radial: radar scan.</li>
	// <li>RotateScaleFade: vertical rotating.</li>
	// <li>Squeeze: vertical gathering.</li>
	// <li>Swap: zooming in.</li>
	// <li>Swirl: swirling.</li>
	// <li>UndulatingBurnOutSwirl: water spreading.</li>
	// <li>Windowblinds: blinds.</li>
	// <li>WipeDown: collapsing down.</li>
	// <li>WipeLeft: collapsing to the left.</li>
	// <li>WipeRight: collapsing to the right.</li>
	// <li>WipeUp: collapsing up.</li>
	// <li>ZoomInCircles: ripples.</li>
	// </ul>
	// </li>
	// <li>Audio transition operation, which is used for transition between two audio segments:
	// <ul>
	// <li>AudioFadeInFadeOut: audio fade-in/fade-out.</li>
	// </ul>
	// </li>
	// </ul>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type UserDefineAsrTextReviewTemplateInfo struct {

	// Switch of custom speech audit task. Valid values:
	// <li>ON: enables custom speech audit task;</li>
	// <li>OFF: disables custom speech audit task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom speech filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom speech keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineConfigureInfo struct {

	// Control parameter of custom figure audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceReviewInfo *UserDefineFaceReviewTemplateInfo `json:"FaceReviewInfo,omitempty" name:"FaceReviewInfo"`

	// Control parameter of custom speech audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of custom text audit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type UserDefineFaceReviewTemplateInfo struct {

	// Switch of custom figure audit task. Valid values:
	// <li>ON: enables custom figure audit task;</li>
	// <li>OFF: disables custom figure audit task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom figure filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for the custom figure library.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 97 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 95 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfo struct {

	// Switch of custom text audit task. Valid values:
	// <li>ON: enables custom text audit task;</li>
	// <li>OFF: disables custom text audit task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom text filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom text keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type VideoTemplateInfo struct {

	// Video stream encoder. Valid values:
	// <li>libx264: H.264</li>
	// <li>libx265: H.265</li>
	// <li>av1: AOMedia Video 1</li>
	// Currently, a resolution within 640x480 must be specified for H.265. and the `av1` container only supports mp4.
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Video frame rate in Hz. Value range: [0, 60].
	// If the value is 0, the frame rate will be the same as that of the source video.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Bitrate of video stream in Kbps. Value range: 0 and [128, 35,000].
	// If the value is 0, the bitrate of the video will be the same as that of the source video.
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Maximum value of the width (or long side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type VideoTemplateInfoForUpdate struct {

	// Video stream encoder. Valid values:
	// <li>libx264: H.264</li>
	// <li>libx265: H.265</li>
	// <li>av1: AOMedia Video 1</li>
	// Currently, a resolution within 640x480 must be specified for H.265. and the `av1` container only supports mp4.
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Video frame rate in Hz. Value range: [0, 60].
	// If the value is 0, the frame rate will be the same as that of the source video.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Bitrate of video stream in Kbps. Value range: 0 and [128, 35,000].
	// If the value is 0, the bitrate of the video will be the same as that of the source video.
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Maximum value of the width (or long side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type VideoTrackItem struct {

	// Source of media material for video segment, which can be an ID of a VOD file or URL of another file.
	SourceMedia *string `json:"SourceMedia,omitempty" name:"SourceMedia"`

	// Start time of video segment in material file in seconds. Default value: 0.
	SourceMediaStartTime *float64 `json:"SourceMediaStartTime,omitempty" name:"SourceMediaStartTime"`

	// Video segment duration in seconds. By default, the length of the video material will be used, which means that the entire material will be captured. If the source file is an image, `Duration` needs to be greater than 0.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Video origin position. Valid values:
	// <li> Center: the origin of coordinates is the center position, such as the center of canvas.</li>
	// Default value: Center.
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the video segment relative to the origin of the canvas. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the video segment will be at the position of the specified percentage of the canvas width; for example, `10%` means that `XPos` is 10% of the canvas width.</li>
	// <li>If the string ends in px, the `XPos` of the video segment will be in px; for example, `100px` means that `XPos` is 100 px.</li>
	// Default value: 0 px.
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// The vertical position of the origin of the video segment relative to the origin of the canvas. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the video segment will be at the position of the specified percentage of the canvas height; for example, `10%` means that `YPos` is 10% of the canvas height.</li>
	// <li>If the string ends in px, the `YPos` of the video segment will be in px; for example, `100px` means that `YPos` is 100 px.</li>
	// Default value: 0 px.
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// Video segment width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the video segment will be the specified percentage of the canvas width; for example, `10%` means that `Width` is 10% of the canvas width.</li>
	// <li>If the string ends in px, the `Width` of the video segment will be in px; for example, `100px` means that `Width` is 100 px.</li>
	// <li>If both `Width` and `Height` are empty, then they will be the `Width` and `Height` of the video material, respectively.</li>
	// <li>If `Width` is empty, but `Height` is not empty, `Width` will be proportionally scaled.</li>
	// <li>If `Width` is not empty, but `Height` is empty, `Height` will be proportionally scaled.</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// Video segment height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the video segment will be the specified percentage of the canvas height; for example, `10%` means that `Height` is 10% of the canvas height;
	// </li><li>If the string ends in px, the `Height` of the video segment will be in px; for example, `100px` means that `Height` is 100 px.</li>
	// <li>If both `Width` and `Height` are empty, then they will be the `Width` and `Height` of the video material, respectively.</li>
	// <li>If `Width` is empty, but `Height` is not empty, `Width` will be proportionally scaled.</li>
	// <li>If `Width` is not empty, but `Height` is empty, `Height` will be proportionally scaled.</li>
	Height *string `json:"Height,omitempty" name:"Height"`

	// Operation on video image such as image rotation.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageOperations []*ImageTransform `json:"ImageOperations,omitempty" name:"ImageOperations" list`

	// Operation on audio such as muting.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioOperations []*AudioTransform `json:"AudioOperations,omitempty" name:"AudioOperations" list`
}

type VideoTrackTemplateInfo struct {

	// Unique ID of a video track template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Encoder. Valid values:
	// <li>libx264: H.264</li>
	// <li>libx265: H.265</li>
	// Currently, a resolution within 640*480p must be specified for H.265.
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Video frame rate in Hz. Value range: [0, 60].
	// If the value is 0, the frame rate will be the same as that of the source video.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Bitrate of a video stream in Kbps. Value range: 0 and [128, 35,000].
	// If the value is 0, the bitrate of the video will be the same as that of the source video.
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Maximum value of the width (or long side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type WatermarkCycleConfigForUpdate struct {

	// Playback time point in seconds when a watermark appears in a video for the first time.
	StartTime *float64 `json:"StartTime,omitempty" name:"StartTime"`

	// Display duration of a watermark in a watermark cycle in seconds.
	DisplayDuration *float64 `json:"DisplayDuration,omitempty" name:"DisplayDuration"`

	// Duration of a watermark cycle in seconds.
	// If 0 is entered, a watermark will last for only one cycle (i.e., visible for `DisplayDuration` seconds throughout the video).
	CycleDuration *float64 `json:"CycleDuration,omitempty" name:"CycleDuration"`
}

type WatermarkInput struct {

	// Watermarking template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Text content of up to 100 characters. This needs to be entered only when the watermark type is text.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TextContent *string `json:"TextContent,omitempty" name:"TextContent"`

	// SVG content of up to 2,000,000 characters. This needs to be entered only when the watermark type is `SVG`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SvgContent *string `json:"SvgContent,omitempty" name:"SvgContent"`

	// Start time offset of a watermark in seconds. If this parameter is left blank or 0 is entered, the watermark will appear upon the first video frame.
	// <li>If this parameter is left blank or 0 is entered, the watermark will appear upon the first video frame;</li>
	// <li>If this value is greater than 0 (e.g., n), the watermark will appear at second n after the first video frame;</li>
	// <li>If this value is smaller than 0 (e.g., -n), the watermark will appear at second n before the last video frame.</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a watermark in seconds.
	// <li>If this parameter is left blank or 0 is entered, the watermark will exist till the last video frame;</li>
	// <li>If this value is greater than 0 (e.g., n), the watermark will exist till second n;</li>
	// <li>If this value is smaller than 0 (e.g., -n), the watermark will exist till second n before the last video frame.</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type WatermarkTemplate struct {

	// Unique ID of watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Watermark type. Valid values:
	// <li>image: image watermark;</li>
	// <li>text: text watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Watermarking template name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Horizontal position of the origin of the watermark image relative to the origin of the video.
	// <li>If the string ends in %, the `Left` edge of the watermark will be at the position of the specified percentage of the video width; for example, `10%` means that the `Left` edge is at 10% of the video width;</li>
	// <li>If the string ends in px, the `Left` edge of the watermark will be at the position of the specified px of the video width; for example, `100px` means that the `Left` edge is at the position of 100 px.</li>
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// Vertical position of the origin of the watermark image relative to the origin of the video.
	// <li>If the string ends in %, the `Top` edge of the watermark will beat the position of the specified percentage of the video height; for example, `10%` means that the `Top` edge is at 10% of the video height;</li>
	// <li>If the string ends in px, the `Top` edge of the watermark will be at the position of the specified px of the video height; for example, `100px` means that the `Top` edge is at the position of 100 px.</li>
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// Image watermarking template. This field is valid only when `Type` is `image`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageTemplate *ImageWatermarkTemplate `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is valid only when `Type` is `text`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is valid when `Type` is `svg`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitempty" name:"SvgTemplate"`

	// Creation time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://cloud.tencent.com/document/product/266/11732#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Origin position. Valid values:
	// <li>topLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>topRight: the origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>bottomLeft: the origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>bottomRight: the origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`
}

type WeChatMiniProgramPublishRequest struct {
	*tchttp.BaseRequest

	// Media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// ID of the transcoding template corresponding to the published video. 0 represents the source video.
	SourceDefinition *int64 `json:"SourceDefinition,omitempty" name:"SourceDefinition"`

	// [Subapplication](/document/product/266/14574) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *WeChatMiniProgramPublishRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WeChatMiniProgramPublishRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WeChatMiniProgramPublishResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WeChatMiniProgramPublishResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WeChatMiniProgramPublishResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WechatMiniProgramPublishTask struct {

	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task status. Valid values:
	// WAITING: waiting;
	// PROCESSING: processing;
	// FINISH: completed.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// ID of published video file.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// ID of the transcoding template corresponding to the published video. 0 represents the source video.
	SourceDefinition *uint64 `json:"SourceDefinition,omitempty" name:"SourceDefinition"`

	// Status of video release on WeChat Mini Program. Valid values:
	// <li>Pass: successfully published;</li>
	// <li>Failed: failed to publish;</li>
	// <li>Rejected: rejected.</li>
	PublishResult *string `json:"PublishResult,omitempty" name:"PublishResult"`
}

type WechatMiniProgramPublishTaskInput struct {

	// ID of the transcoding template corresponding to the published video. 0 represents the source video.
	SourceDefinition *uint64 `json:"SourceDefinition,omitempty" name:"SourceDefinition"`
}

type WechatPublishTask struct {

	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task status. Valid values:
	// WAITING: waiting;
	// PROCESSING: processing;
	// FINISH: completed.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// ID of published video file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Release on WeChat template ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// ID of the transcoding template corresponding to the published video. 0 represents the source video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SourceDefinition *uint64 `json:"SourceDefinition,omitempty" name:"SourceDefinition"`

	// Release on WeChat status. Valid values:
	// <li>FAIL: failure;</li>
	// <li>SUCCESS: success;</li>
	// <li>AUDITNOTPASS: rejected</li>
	// <li>NOTTRIGGERED: release on WeChat not initiated yet.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	WechatStatus *string `json:"WechatStatus,omitempty" name:"WechatStatus"`

	// WeChat `Vid`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WechatVid *string `json:"WechatVid,omitempty" name:"WechatVid"`

	// WeChat address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WechatUrl *string `json:"WechatUrl,omitempty" name:"WechatUrl"`
}
