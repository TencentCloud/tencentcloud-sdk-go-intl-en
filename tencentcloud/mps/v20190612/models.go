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

package v20190612

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
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
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitempty" name:"ClassificationConfigure"`

	// Control parameter of intelligent tagging task.
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitempty" name:"TagConfigure"`

	// Control parameter of intelligent cover generating task.
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitempty" name:"CoverConfigure"`

	// Control parameter of intelligent frame-specific tagging task.
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitempty" name:"FrameTagConfigure"`

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AIRecognitionTemplateItem struct {

	// Unique ID of a video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Name of a video content recognition template.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of a video content recognition template.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Face recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// Full text recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// Text keyword recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// Full speech recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// Speech keyword recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AdaptiveDynamicStreamingInfoItem struct {

	// Adaptive bitrate streaming specification.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Container format. Valid values: HLS, MPEG-DASH.
	Package *string `json:"Package,omitempty" name:"Package"`

	// Playback address.
	Path *string `json:"Path,omitempty" name:"Path"`

	// Storage location of adaptive bitrate streaming files.
	Storage *TaskOutputStorage `json:"Storage,omitempty" name:"Storage"`
}

type AdaptiveDynamicStreamingTaskInput struct {

	// Adaptive bitrate streaming template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of up to 10 image or text watermarks.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// Target bucket of an output file after being transcoded to adaptive bitrate streaming. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// The relative or absolute output path of the manifest file after being transcoded to adaptive bitrate streaming. If this parameter is left empty, a relative path in the following format will be used by default: `{inputName}_adaptiveDynamicStreaming_{definition}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// The relative output path of the substream file after being transcoded to adaptive bitrate streaming. If this parameter is left empty, a relative path in the following format will be used by default: `{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}.{format}`.
	SubStreamObjectName *string `json:"SubStreamObjectName,omitempty" name:"SubStreamObjectName"`

	// The relative output path of the segment file after being transcoded to adaptive bitrate streaming (in HLS format only). If this parameter is left empty, a relative path in the following format will be used by default: `{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}_{segmentNumber}.{format}`.
	SegmentObjectName *string `json:"SegmentObjectName,omitempty" name:"SegmentObjectName"`
}

type AdaptiveDynamicStreamingTemplate struct {

	// Unique ID of an adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of an adaptive bitrate streaming template.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of an adaptive bitrate streaming template.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Adaptive bitrate streaming format. Valid values:
	// <li>HLS;</li>
	// <li>MPEG-DASH.</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Parameter information of input streams for transcoding to adaptive bitrate streaming. Up to 10 streams can be input.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitempty" name:"StreamInfos"`

	// Whether to prohibit transcoding from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitempty" name:"DisableHigherVideoResolution"`

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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

	// Whether to remove video stream. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	RemoveVideo *uint64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`
}

type AiAnalysisResult struct {

	// Task type. Valid values:
	// <li>Classification: intelligent categorization</li>
	// <li>Cover: intelligent cover generating</li>
	// <li>Tag: intelligent tagging</li>
	// <li>FrameTag: intelligent frame-specific tagging</li>
	// <li>Highlight: intelligent highlight generating</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Query result of intelligent categorization task in video content analysis, which is valid if task type is `Classification`.
	ClassificationTask *AiAnalysisTaskClassificationResult `json:"ClassificationTask,omitempty" name:"ClassificationTask"`

	// Query result of intelligent cover generating task in video content analysis, which is valid if task type is `Cover`.
	CoverTask *AiAnalysisTaskCoverResult `json:"CoverTask,omitempty" name:"CoverTask"`

	// Query result of intelligent tagging task in video content analysis, which is valid if task type is `Tag`.
	TagTask *AiAnalysisTaskTagResult `json:"TagTask,omitempty" name:"TagTask"`

	// Query result of intelligent frame-specific tagging task in video content analysis, which is valid if task type is `FrameTag`.
	FrameTagTask *AiAnalysisTaskFrameTagResult `json:"FrameTagTask,omitempty" name:"FrameTagTask"`
}

type AiAnalysisTaskClassificationInput struct {

	// Intelligent video categorization template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskClassificationOutput struct {

	// List of intelligently generated video categories.
	ClassificationSet []*MediaAiAnalysisClassificationItem `json:"ClassificationSet,omitempty" name:"ClassificationSet"`
}

type AiAnalysisTaskClassificationResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of intelligent categorization task.
	Input *AiAnalysisTaskClassificationInput `json:"Input,omitempty" name:"Input"`

	// Output of intelligent categorization task.
	Output *AiAnalysisTaskClassificationOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskCoverInput struct {

	// Intelligent video cover generating template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskCoverOutput struct {

	// List of intelligently generated covers.
	CoverSet []*MediaAiAnalysisCoverItem `json:"CoverSet,omitempty" name:"CoverSet"`

	// Storage location of intelligently generated cover.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`
}

type AiAnalysisTaskCoverResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of intelligent cover generating task.
	Input *AiAnalysisTaskCoverInput `json:"Input,omitempty" name:"Input"`

	// Output of intelligent cover generating task.
	Output *AiAnalysisTaskCoverOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskFrameTagInput struct {

	// Intelligent frame-specific video tagging template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskFrameTagOutput struct {

	// List of frame-specific video tags.
	SegmentSet []*MediaAiAnalysisFrameTagSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiAnalysisTaskFrameTagResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of intelligent frame-specific tagging task.
	Input *AiAnalysisTaskFrameTagInput `json:"Input,omitempty" name:"Input"`

	// Output of intelligent frame-specific tagging task.
	Output *AiAnalysisTaskFrameTagOutput `json:"Output,omitempty" name:"Output"`
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
	TagSet []*MediaAiAnalysisTagItem `json:"TagSet,omitempty" name:"TagSet"`
}

type AiAnalysisTaskTagResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of intelligent tagging task.
	Input *AiAnalysisTaskTagInput `json:"Input,omitempty" name:"Input"`

	// Output of intelligent tagging task.
	Output *AiAnalysisTaskTagOutput `json:"Output,omitempty" name:"Output"`
}

type AiContentReviewResult struct {

	// Task type. Valid values:
	// <li>Porn: Porn information detection in image</li>
	// <li>Terrorism: Terrorism information detection in image</li>
	// <li>Political: Politically sensitive information detection in image</li>
	// <li>Porn.Asr: ASR-based porn information detection in text</li>
	// <li>Porn.Ocr: OCR-based porn information detection in text</li>
	// <li>Porn.Voice: Porn information detection in speech</li>
	// <li>Political.Asr: ASR-based politically sensitive information detection in text</li>
	// <li>Political.Ocr: OCR-based politically sensitive information detection in text</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Sample rate, which indicates the number of video frames captured per second for audit
	SampleRate *float64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// Audited video duration in seconds.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Query result of an intelligent porn information detection in image task in video content audit, which is valid when task type is `Porn`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PornTask *AiReviewTaskPornResult `json:"PornTask,omitempty" name:"PornTask"`

	// Query result of an intelligent terrorism information detection in image task in video content audit, which is valid when task type is `Terrorism`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TerrorismTask *AiReviewTaskTerrorismResult `json:"TerrorismTask,omitempty" name:"TerrorismTask"`

	// Query result of an intelligent politically sensitive information detection in image task in video content audit, which is valid when task type is `Political`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PoliticalTask *AiReviewTaskPoliticalResult `json:"PoliticalTask,omitempty" name:"PoliticalTask"`

	// Query result of an ASR-based porn information detection in text task in video content audit, which is valid when task type is `Porn.Asr`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PornAsrTask *AiReviewTaskPornAsrResult `json:"PornAsrTask,omitempty" name:"PornAsrTask"`

	// Query result of an OCR-based porn information detection in text task in video content audit, which is valid when task type is `Porn.Ocr`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PornOcrTask *AiReviewTaskPornOcrResult `json:"PornOcrTask,omitempty" name:"PornOcrTask"`

	// Query result of an ASR-based politically sensitive information detection in text task in video content audit, which is valid when task type is `Political.Asr`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PoliticalAsrTask *AiReviewTaskPoliticalAsrResult `json:"PoliticalAsrTask,omitempty" name:"PoliticalAsrTask"`

	// Query result of an OCR-based politically sensitive information detection in text task in video content audit, which is valid when task type is `Political.Ocr`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PoliticalOcrTask *AiReviewTaskPoliticalOcrResult `json:"PoliticalOcrTask,omitempty" name:"PoliticalOcrTask"`

	// Query result of OCR-based terrorism information detection in text task in video content audit, which is valid if task type is `Terrorism.Ocr`.
	TerrorismOcrTask *AiReviewTaskTerrorismOcrResult `json:"TerrorismOcrTask,omitempty" name:"TerrorismOcrTask"`

	// Query result of ASR-based prohibited information detection in speech task in video content audit, which is valid if task type is `Prohibited.Asr`.
	ProhibitedAsrTask *AiReviewTaskProhibitedAsrResult `json:"ProhibitedAsrTask,omitempty" name:"ProhibitedAsrTask"`

	// Query result of OCR-based prohibited information detection in text task in video content audit, which is valid if task type is `Prohibited.Ocr`.
	ProhibitedOcrTask *AiReviewTaskProhibitedOcrResult `json:"ProhibitedOcrTask,omitempty" name:"ProhibitedOcrTask"`
}

type AiContentReviewTaskInput struct {

	// Video content audit template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionResult struct {

	// Task type. Valid values:
	// <li>FaceRecognition: Face recognition,</li>
	// <li>AsrWordsRecognition: Speech keyword recognition,</li>
	// <li>OcrWordsRecognition: Text keyword recognition,</li>
	// <li>AsrFullTextRecognition: Full speech recognition,</li>
	// <li>OcrFullTextRecognition: Full text recognition,</li>
	// <li>HeadTailRecognition: Video opening and ending credits recognition,</li>
	// <li>ObjectRecognition: Object recognition.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Face recognition result, which is valid when `Type` is 
	//  `FaceRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FaceTask *AiRecognitionTaskFaceResult `json:"FaceTask,omitempty" name:"FaceTask"`

	// Speech keyword recognition result, which is valid when `Type` is
	//  `AsrWordsRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrWordsTask *AiRecognitionTaskAsrWordsResult `json:"AsrWordsTask,omitempty" name:"AsrWordsTask"`

	// Full speech recognition result, which is valid when `Type` is
	//  `AsrFullTextRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrFullTextTask *AiRecognitionTaskAsrFullTextResult `json:"AsrFullTextTask,omitempty" name:"AsrFullTextTask"`

	// Text keyword recognition result, which is valid when `Type` is
	//  `OcrWordsRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrWordsTask *AiRecognitionTaskOcrWordsResult `json:"OcrWordsTask,omitempty" name:"OcrWordsTask"`

	// Full text recognition result, which is valid when `Type` is
	//  `OcrFullTextRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrFullTextTask *AiRecognitionTaskOcrFullTextResult `json:"OcrFullTextTask,omitempty" name:"OcrFullTextTask"`
}

type AiRecognitionTaskAsrFullTextResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of a full speech recognition task.
	Input *AiRecognitionTaskAsrFullTextResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of a full speech recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskAsrFullTextResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskAsrFullTextResultInput struct {

	// Full speech recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrFullTextResultOutput struct {

	// List of full speech recognition segments.
	SegmentSet []*AiRecognitionTaskAsrFullTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// Subtitles file address.
	SubtitlePath *string `json:"SubtitlePath,omitempty" name:"SubtitlePath"`

	// Subtitles file storage location.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`
}

type AiRecognitionTaskAsrFullTextSegmentItem struct {

	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Recognized text.
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AiRecognitionTaskAsrWordsResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of a speech keyword recognition task.
	Input *AiRecognitionTaskAsrWordsResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of a speech keyword recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	SegmentSet []*AiRecognitionTaskAsrWordsSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskAsrWordsResultOutput struct {

	// Speech keyword recognition result set.
	ResultSet []*AiRecognitionTaskAsrWordsResultItem `json:"ResultSet,omitempty" name:"ResultSet"`
}

type AiRecognitionTaskAsrWordsSegmentItem struct {

	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type AiRecognitionTaskFaceResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of a face recognition task.
	Input *AiRecognitionTaskFaceResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of a face recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskFaceResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskFaceResultInput struct {

	// Face recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskFaceResultItem struct {

	// Unique ID of a figure.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Figure library type, indicating to which figure library the recognized figure belongs:
	// <li>Default: Default figure library;</li>
	// <li>UserDefine: Custom figure library.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of a figure.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Result set of segments that contain a figure.
	SegmentSet []*AiRecognitionTaskFaceSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskFaceResultOutput struct {

	// Intelligent face recognition result set.
	ResultSet []*AiRecognitionTaskFaceResultItem `json:"ResultSet,omitempty" name:"ResultSet"`
}

type AiRecognitionTaskFaceSegmentItem struct {

	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of a recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskInput struct {

	// Intelligent video recognition template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrFullTextResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of a full text recognition task.
	Input *AiRecognitionTaskOcrFullTextResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of a full text recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskOcrFullTextResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskOcrFullTextResultInput struct {

	// Full text recognition template ID.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrFullTextResultOutput struct {

	// Full text recognition result set.
	SegmentSet []*AiRecognitionTaskOcrFullTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskOcrFullTextSegmentItem struct {

	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Recognition segment result set.
	TextSet []*AiRecognitionTaskOcrFullTextSegmentTextItem `json:"TextSet,omitempty" name:"TextSet"`
}

type AiRecognitionTaskOcrFullTextSegmentTextItem struct {

	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of a recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`

	// Recognized text.
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AiRecognitionTaskOcrWordsResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input information of a text keyword recognition task.
	Input *AiRecognitionTaskOcrWordsResultInput `json:"Input,omitempty" name:"Input"`

	// Output information of a text keyword recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	SegmentSet []*AiRecognitionTaskOcrWordsSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskOcrWordsResultOutput struct {

	// Text keyword recognition result set.
	ResultSet []*AiRecognitionTaskOcrWordsResultItem `json:"ResultSet,omitempty" name:"ResultSet"`
}

type AiRecognitionTaskOcrWordsSegmentItem struct {

	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of a recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`
}

type AiReviewPoliticalAsrTaskInput struct {

	// ID of a politically sensitive information detection template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalAsrTaskOutput struct {

	// Score of the ASR-detected politically sensitive information in text from 0 to 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for the ASR-detected politically sensitive information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain ASR-detected politically sensitive information in text.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiReviewPoliticalOcrTaskInput struct {

	// ID of a politically sensitive information detection template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalOcrTaskOutput struct {

	// Score of the OCR-detected politically sensitive information in text from 0 to 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for the OCR-detected politically sensitive information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain OCR-detected politically sensitive information in text.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiReviewPoliticalTaskInput struct {

	// ID of a politically sensitive information detection template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalTaskOutput struct {

	// Score of the detected politically sensitive information in video from 0 to 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for the detected politically sensitive information. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tags for the results of video politically sensitive information detection. The relationship between the `LabelSet` parameter in the content audit template [controlling tasks of video politically sensitive information detection](https://intl.cloud.tencent.com/document/api/862/37615?from_cn_redirect=1#AiReviewPoliticalTaskOutput) and this parameter is as follows:
	// violation_photo:
	// <li>violation_photo: violating photo.</li>
	// Other values (politician/entertainment/sport/entrepreneur/scholar/celebrity/military):
	// <li>politician: political figure.</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// List of video segments that contain the detected politically sensitive information.
	SegmentSet []*MediaContentReviewPoliticalSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiReviewPornAsrTaskInput struct {

	// ID of a porn information detection template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornAsrTaskOutput struct {

	// Score of the ASR-detected porn information in text from 0 to 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for the ASR-detected porn information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain the ASR-detected porn information in text.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiReviewPornOcrTaskInput struct {

	// ID of a porn information detection template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornOcrTaskOutput struct {

	// Score of the OCR-detected porn information in text from 0 to 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for the OCR-detected porn information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain the OCR-detected porn information in text.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiReviewPornTaskInput struct {

	// ID of a porn information detection template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornTaskOutput struct {

	// Score of the detected porn information in video from 0 to 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for the detected porn information. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tag of the detected porn information in video. Valid values:
	// <li>porn: Porn.</li>
	// <li>sexy: Sexiness.</li>
	// <li>vulgar: Vulgarity.</li>
	// <li>intimacy: Intimacy.</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// List of video segments that contain the detected porn information.
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
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
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
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
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiReviewTaskPoliticalAsrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see the list of [Error Codes](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for an ASR-based politically sensitive information detection in text task during content audit.
	Input *AiReviewPoliticalAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of an ASR-based politically sensitive information detection in text task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPoliticalAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalOcrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for an OCR-based politically sensitive information detection in text task during content audit.
	Input *AiReviewPoliticalOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of an OCR-based politically sensitive information detection in text task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPoliticalOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for a politically sensitive information detection task during content audit.
	Input *AiReviewPoliticalTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of a politically sensitive information detection task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPoliticalTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornAsrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for an ASR-based porn information detection in text task during content audit.
	Input *AiReviewPornAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of an ASR-based porn information detection in text task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPornAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornOcrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for an OCR-based porn information detection in text task during content audit.
	Input *AiReviewPornOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of an OCR-based porn information detection in text task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPornOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for a porn information detection task during content audit.
	Input *AiReviewPornTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of a porn information detection task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPornTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskProhibitedAsrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

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
	Output *AiReviewProhibitedAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskProhibitedOcrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

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
	Output *AiReviewProhibitedOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismOcrResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

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
	Output *AiReviewTerrorismOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for a terrorism information detection task during content audit.
	Input *AiReviewTerrorismTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of a terrorism information detection task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiReviewTerrorismTaskInput struct {

	// ID of a terrorism information detection template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewTerrorismTaskOutput struct {

	// Score of the detected terrorism information in a video from 0 to 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for the detected terrorism information. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tags for detected terrorism information in a video. Valid values:
	// <li>`guns`: weapons and guns</li>
	// <li>`crowd`: crowds</li>
	// <li>`police`: police forces</li>
	// <li>`bloody`: bloodiness</li>
	// <li>`banners`: terrorism flags</li>
	// <li>`militant`: militants</li>
	// <li>`explosion`: explosions and fires</li>
	// <li>`terrorists`: terrorists</li>
	// <li>`scenario`: terrorism images</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// List of video segments that contain the detected terrorism information.
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiSampleFaceInfo struct {

	// Face image ID.
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// Face image address.
	Url *string `json:"Url,omitempty" name:"Url"`
}

type AiSampleFaceOperation struct {

	// Operation type. Valid values: add, delete, reset. The `reset` operation will clear the existing face data of a figure and add `FaceContents` as the specified face data.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Face ID set. This field is required when `Type` is `delete`.
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds"`

	// String set generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) the face image.
	// <li>This field is required when `Type` is `add` or `reset`;</li>
	// <li>Array length limit: 5 images.</li>
	// Note: The image must be a relatively clear full-face photo of a figure in at least 200 * 200 px.
	FaceContents []*string `json:"FaceContents,omitempty" name:"FaceContents"`
}

type AiSampleFailFaceInfo struct {

	// Corresponds to incorrect image subscripts in the `FaceContents` input parameter, starting from 0.
	Index *uint64 `json:"Index,omitempty" name:"Index"`

	// Error code. Valid values:
	// <li>0: Succeeded;</li>
	// <li>Other values: Failed.</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error description.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type AiSamplePerson struct {

	// Figure ID.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Name of a figure.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Figure description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Face information.
	FaceInfoSet []*AiSampleFaceInfo `json:"FaceInfoSet,omitempty" name:"FaceInfoSet"`

	// Figure tag.
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet"`

	// Use case.
	UsageSet []*string `json:"UsageSet,omitempty" name:"UsageSet"`

	// Creation time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AiSampleTagOperation struct {

	// Operation type. Valid values: add, delete, reset.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Tag. Length limit: 128 characters.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type AiSampleWord struct {

	// Keyword.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Keyword tag.
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet"`

	// Keyword use case.
	UsageSet []*string `json:"UsageSet,omitempty" name:"UsageSet"`

	// Creation time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AiSampleWordInfo struct {

	// Keyword. Length limit: 20 characters.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Keyword tag
	// <li>Array length limit: 20 tags;</li>
	// <li>Tag length limit: 128 characters.</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type AnimatedGraphicTaskInput struct {

	// Animated image generating template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Start time of an animated image in a video in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time of an animated image in a video in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Target bucket of a generated animated image file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Output path to a generated animated image file, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_animatedGraphic_{definition}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`
}

type AnimatedGraphicsTemplate struct {

	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
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
	// <li>open: Enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: Disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Animated image format.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Frame rate.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Image quality.
	Quality *float64 `json:"Quality,omitempty" name:"Quality"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AsrFullTextConfigureInfo struct {

	// Switch of a full speech recognition task. Valid values:
	// <li>ON: Enables an intelligent full speech recognition task;</li>
	// <li>OFF: Disables an intelligent full speech recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Format of the generated subtitles file. If this parameter is left empty or an empty string is entered, no subtitles files will be generated. Valid value:
	// <li>vtt: Generates a WebVTT subtitles file.</li>
	SubtitleFormat *string `json:"SubtitleFormat,omitempty" name:"SubtitleFormat"`
}

type AsrFullTextConfigureInfoForUpdate struct {

	// Switch of a full speech recognition task. Valid values:
	// <li>ON: Enables an intelligent full speech recognition task;</li>
	// <li>OFF: Disables an intelligent full speech recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Format of the generated subtitles file. If an empty string is entered, no subtitles files will be generated. Valid value:
	// <li>vtt: Generates a WebVTT subtitles file.</li>
	SubtitleFormat *string `json:"SubtitleFormat,omitempty" name:"SubtitleFormat"`
}

type AsrWordsConfigureInfo struct {

	// Switch of a speech keyword recognition task. Valid values:
	// <li>ON: Enables a speech keyword recognition task;</li>
	// <li>OFF: Disables a speech keyword recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`
}

type AsrWordsConfigureInfoForUpdate struct {

	// Switch of a speech keyword recognition task. Valid values:
	// <li>ON: Enables a speech keyword recognition task;</li>
	// <li>OFF: Disables a speech keyword recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`
}

type AudioTemplateInfo struct {

	// Audio stream codec.
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
	// <li>libmp3lame: more suitable for flv.</li>
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
	// <li>1: Mono</li>
	// <li>2: Dual</li>
	// <li>6: Stereo</li>
	// When the media is packaged in audio format (FLAC, OGG, MP3, M4A), the sound channel cannot be set to stereo.
	// Default value: 2
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type AudioTemplateInfoForUpdate struct {

	// Audio stream codec.
	// When the outer `Container` parameter is `mp3`, the valid value is:
	// <li>libmp3lame.</li>
	// When the outer `Container` parameter is `ogg` or `flac`, the valid value is:
	// <li>flac.</li>
	// When the outer `Container` parameter is `m4a`, the valid values include:
	// <li>libfdk_aac;</li>
	// <li>libmp3lame;</li>
	// <li>ac3.</li>
	// When the outer `Container` parameter is `mp4` or `flv`, the valid values include:
	// <li>libfdk_aac: More suitable for mp4;</li>
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
	// <li>1: Mono</li>
	// <li>2: Dual</li>
	// <li>6: Stereo</li>
	// When the media is packaged in audio format (FLAC, OGG, MP3, M4A), the sound channel cannot be set to stereo.
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
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

type ContentReviewTemplateItem struct {

	// Unique ID of a content audit template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Name of a content audit template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of a content audit template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Porn information detection control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// Terrorism information detection control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// Politically sensitive information detection control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitempty" name:"ProhibitedConfigure"`

	// Custom content audit control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CosFileUploadTrigger struct {

	// Name of the COS bucket bound to a workflow, such as `TopRankVideo-125xxx88`.
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Region of the COS bucket bound to a workflow, such as `ap-chongiqng`.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Input path directory bound to a workflow, such as `/movie/201907/`. If this parameter is left empty, the `/` root directory will be used.
	Dir *string `json:"Dir,omitempty" name:"Dir"`

	// Format list of files that can trigger a workflow, such as ["mp4", "flv", "mov"]. If this parameter is left empty, files in all formats can trigger the workflow.
	Formats []*string `json:"Formats,omitempty" name:"Formats"`
}

type CosInputInfo struct {

	// Name of the COS bucket where a video processing object file is located, such as `TopRankVideo-125xxx88`.
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Region of the COS bucket where a video processing object file is located, such as `ap-chongqing`.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Path to an input object file for video processing, such as `/movie/201907/WildAnimal.mov`.
	Object *string `json:"Object,omitempty" name:"Object"`
}

type CosOutputStorage struct {

	// Name of the target bucket of a video processing output file, such as `TopRankVideo-125xxx88`. If this parameter is left empty, the parameter of the upper folder will be inherited.
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Region of the target bucket of a video processing output file, such as `ap-chongqing`. If this parameter is left empty, the parameter of the upper folder will be inherited.
	Region *string `json:"Region,omitempty" name:"Region"`
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
}

func (r *CreateAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIAnalysisTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ClassificationConfigure")
	delete(f, "TagConfigure")
	delete(f, "CoverConfigure")
	delete(f, "FrameTagConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIAnalysisTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIAnalysisTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest

	// Name of a video content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of a video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Face recognition control parameter.
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// Full text recognition control parameter.
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// Text keyword recognition control parameter.
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// Full speech recognition control parameter.
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// Speech keyword recognition control parameter.
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`
}

func (r *CreateAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIRecognitionTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "FaceConfigure")
	delete(f, "OcrFullTextConfigure")
	delete(f, "OcrWordsConfigure")
	delete(f, "AsrFullTextConfigure")
	delete(f, "AsrWordsConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of a video content recognition template.
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAIRecognitionTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIRecognitionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest

	// Adaptive bitrate streaming format. Valid values:
	// <li>HLS,</li>
	// <li>MPEG-DASH.</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Parameter information of output substreams for transcoding to adaptive bitrate streaming. Up to 10 substreams can be output.
	// Note: the frame rate of each substream must be consistent; otherwise, the frame rate of the first substream is used as the output frame rate.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitempty" name:"StreamInfos"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Whether to prohibit transcoding from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: 0.
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: 0.
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitempty" name:"DisableHigherVideoResolution"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateAdaptiveDynamicStreamingTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdaptiveDynamicStreamingTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Format")
	delete(f, "StreamInfos")
	delete(f, "Name")
	delete(f, "DisableHigherVideoBitrate")
	delete(f, "DisableHigherVideoResolution")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAdaptiveDynamicStreamingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of an adaptive bitrate streaming template.
		Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAdaptiveDynamicStreamingTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdaptiveDynamicStreamingTemplateResponse) FromJsonString(s string) error {
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

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
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
}

func (r *CreateAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Fps")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Quality")
	delete(f, "Name")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAnimatedGraphicsTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateContentReviewTemplateRequest struct {
	*tchttp.BaseRequest

	// Name of an intelligent content recognition template. Length limit: 64 characters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of an intelligent content recognition template. Length limit: 256 characters
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter for porn information
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// Control parameter for terrorism information
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// Control parameter for politically sensitive information
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this parameter is not supported yet.
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitempty" name:"ProhibitedConfigure"`

	// Control parameter for custom intelligent content recognition tasks
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`
}

func (r *CreateContentReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContentReviewTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "PornConfigure")
	delete(f, "TerrorismConfigure")
	delete(f, "PoliticalConfigure")
	delete(f, "ProhibitedConfigure")
	delete(f, "UserDefineConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of an intelligent content recognition template
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateContentReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContentReviewTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest

	// Sampling type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
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

	// Subimage width of an image sprite in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Subimage height of an image sprite in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageSpriteTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SampleType")
	delete(f, "SampleInterval")
	delete(f, "RowCount")
	delete(f, "ColumnCount")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "FillType")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageSpriteTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageSpriteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePersonSampleRequest struct {
	*tchttp.BaseRequest

	// Name of an image. Length limit: 20 characters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Image usage. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: equivalent to 1+2
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Image description. Length limit: 1,024 characters
	Description *string `json:"Description,omitempty" name:"Description"`

	// [Base64](https://tools.ietf.org/html/rfc4648) string converted from an image. Only JPEG and PNG images are supported. Array length limit: 5 images
	// Note: the image must be a relatively clear facial feature photo of one person with a size of at least 200 x 200 pixels.
	FaceContents []*string `json:"FaceContents,omitempty" name:"FaceContents"`

	// Image tag
	// <li>Array length limit: 20 tags</li>
	// <li>Tag length limit: 128 characters</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreatePersonSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Usages")
	delete(f, "Description")
	delete(f, "FaceContents")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Image information
		Person *AiSamplePerson `json:"Person,omitempty" name:"Person"`

		// Information of images that failed the verification by facial feature positioning
		FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitempty" name:"FailFaceInfoSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePersonSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Sampled screencapturing type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Image format. Valid values: jpg; png. Default value: jpg.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSampleSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SampleType")
	delete(f, "SampleInterval")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Comment")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSampleSnapshotTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSampleSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest

	// Name of a time point screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Image format. Valid values: jpg; png. Default value: jpg.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Comment")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotByTimeOffsetTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Container format. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Name of a transcoding template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Whether to remove video data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	// Default value: 0.
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	// Default value: 0.
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter. This field is required when `RemoveVideo` is 0.
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter. This field is required when `RemoveAudio` is 0.
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter. To enable it, please contact your Tencent Cloud sales rep.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitempty" name:"TEHDConfig"`
}

func (r *CreateTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Container")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	delete(f, "VideoTemplate")
	delete(f, "AudioTemplate")
	delete(f, "TEHDConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of a transcoding template.
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// Watermarking type. Valid values:
	// <li>image: Image watermark;</li>
	// <li>text: Text watermark;</li>
	// <li>svg: SVG watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Watermarking template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Origin position. Valid values:
	// <li>TopLeft: The origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>TopRight: The origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>BottomLeft: The origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>BottomRight: The origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
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

	// Image watermarking template. This field is required and valid only when `Type` is `image`.
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is required and valid only when `Type` is `text`.
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is required and valid only when `Type` is `svg`.
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitempty" name:"SvgTemplate"`
}

func (r *CreateWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWatermarkTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "CoordinateOrigin")
	delete(f, "XPos")
	delete(f, "YPos")
	delete(f, "ImageTemplate")
	delete(f, "TextTemplate")
	delete(f, "SvgTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWatermarkTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of a watermarking template.
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWatermarkTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWordSamplesRequest struct {
	*tchttp.BaseRequest

	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// 7. All: ASR- and OCR-based content recognition and inappropriate information detection; equivalent to 1+2+3+4
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Keyword. Array length limit: 100.
	Words []*AiSampleWordInfo `json:"Words,omitempty" name:"Words"`
}

func (r *CreateWordSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWordSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Usages")
	delete(f, "Words")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWordSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWordSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest

	// Workflow name of up to 128 characters, which must be unique for the same user.
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// Triggering rule bound to a workflow. If an uploaded video hits the rule for the object, the workflow will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// Storage location of a video processing output file. If this parameter is left empty, the storage location in `Trigger` will be inherited.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Target directory of a video processing output file, such as `/movie/201907/`. If this parameter is left empty, the file will be outputted to the same directory where the source file is located.
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// Parameter of a video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Event notification configuration for a task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

	// Workflow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TaskPriority *int64 `json:"TaskPriority,omitempty" name:"TaskPriority"`
}

func (r *CreateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowName")
	delete(f, "Trigger")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "TaskNotifyConfig")
	delete(f, "TaskPriority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Workflow ID.
		WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIAnalysisTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIAnalysisTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIAnalysisTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIRecognitionTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIRecognitionTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIRecognitionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteAdaptiveDynamicStreamingTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAdaptiveDynamicStreamingTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAdaptiveDynamicStreamingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAdaptiveDynamicStreamingTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAdaptiveDynamicStreamingTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAnimatedGraphicsTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteContentReviewTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an intelligent content recognition template
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteContentReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContentReviewTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteContentReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContentReviewTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageSpriteTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageSpriteTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageSpriteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePersonSampleRequest struct {
	*tchttp.BaseRequest

	// Image ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *DeletePersonSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePersonSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteSampleSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSampleSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSampleSnapshotTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSampleSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteSnapshotByTimeOffsetTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotByTimeOffsetTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a transcoding template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

func (r *DeleteWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWatermarkTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWatermarkTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWatermarkTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWordSamplesRequest struct {
	*tchttp.BaseRequest

	// Keyword. Array length limit: 100 words.
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`
}

func (r *DeleteWordSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWordSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keywords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWordSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWordSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowRequest struct {
	*tchttp.BaseRequest

	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

func (r *DeleteWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAIAnalysisTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of video content analysis templates. Array length limit: 10.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAIAnalysisTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAnalysisTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIAnalysisTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAIAnalysisTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of video content analysis template details.
		AIAnalysisTemplateSet []*AIAnalysisTemplateItem `json:"AIAnalysisTemplateSet,omitempty" name:"AIAnalysisTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAIAnalysisTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAnalysisTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAIRecognitionTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of video content recognition templates. Array length limit: 10.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAIRecognitionTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIRecognitionTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIRecognitionTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAIRecognitionTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of video content recognition template details.
		AIRecognitionTemplateSet []*AIRecognitionTemplateItem `json:"AIRecognitionTemplateSet,omitempty" name:"AIRecognitionTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAIRecognitionTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIRecognitionTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAdaptiveDynamicStreamingTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of adaptive bitrate streaming templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAdaptiveDynamicStreamingTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdaptiveDynamicStreamingTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdaptiveDynamicStreamingTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAdaptiveDynamicStreamingTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of adaptive bitrate streaming template details.
		AdaptiveDynamicStreamingTemplateSet []*AdaptiveDynamicStreamingTemplate `json:"AdaptiveDynamicStreamingTemplateSet,omitempty" name:"AdaptiveDynamicStreamingTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAdaptiveDynamicStreamingTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdaptiveDynamicStreamingTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAnimatedGraphicsTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of animated image generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAnimatedGraphicsTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAnimatedGraphicsTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAnimatedGraphicsTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAnimatedGraphicsTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of animated image generating template details.
		AnimatedGraphicsTemplateSet []*AnimatedGraphicsTemplate `json:"AnimatedGraphicsTemplateSet,omitempty" name:"AnimatedGraphicsTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAnimatedGraphicsTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAnimatedGraphicsTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContentReviewTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID of intelligent content recognition templates as the filter. Array length limit: 50
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContentReviewTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentReviewTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContentReviewTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContentReviewTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of content audit template details.
		ContentReviewTemplateSet []*ContentReviewTemplateItem `json:"ContentReviewTemplateSet,omitempty" name:"ContentReviewTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContentReviewTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentReviewTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSpriteTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of image sprite generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeImageSpriteTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSpriteTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageSpriteTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSpriteTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of image sprite generating template details.
		ImageSpriteTemplateSet []*ImageSpriteTemplate `json:"ImageSpriteTemplateSet,omitempty" name:"ImageSpriteTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSpriteTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSpriteTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaMetaDataRequest struct {
	*tchttp.BaseRequest

	// Input information of file for metadata getting.
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`
}

func (r *DescribeMediaMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Media metadata.
		MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePersonSamplesRequest struct {
	*tchttp.BaseRequest

	// Type of images to pull. Valid values:
	// <li>UserDefine: custom image library</li>
	// <li>Default: default image library</li>
	// 
	// Default value: UserDefine. Samples in the custom image library will be pulled.
	// Note: you can pull the default image library only using the image name or a combination of the image name and ID, and only one face image is returned.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Image ID. Array length limit: 100
	PersonIds []*string `json:"PersonIds,omitempty" name:"PersonIds"`

	// Image name. Array length limit: 20
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Image tag. Array length limit: 20
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePersonSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "PersonIds")
	delete(f, "Names")
	delete(f, "Tags")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePersonSamplesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Image information
		PersonSet []*AiSamplePerson `json:"PersonSet,omitempty" name:"PersonSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePersonSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSampleSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of sampled screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSampleSnapshotTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSampleSnapshotTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSampleSnapshotTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSampleSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of sampled screencapturing template details.
		SampleSnapshotTemplateSet []*SampleSnapshotTemplate `json:"SampleSnapshotTemplateSet,omitempty" name:"SampleSnapshotTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSampleSnapshotTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSampleSnapshotTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotByTimeOffsetTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of time point screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSnapshotByTimeOffsetTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotByTimeOffsetTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotByTimeOffsetTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotByTimeOffsetTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of time point screencapturing template details.
		SnapshotByTimeOffsetTemplateSet []*SnapshotByTimeOffsetTemplate `json:"SnapshotByTimeOffsetTemplateSet,omitempty" name:"SnapshotByTimeOffsetTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotByTimeOffsetTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotByTimeOffsetTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task type. Currently valid values:
	// <li>WorkflowTask: Video workflow processing task.</li>
	// <li>LiveStreamProcessTask: Live stream processing task.</li>
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// Task status. Valid values:
	// <li>WAITING: Waiting;</li>
	// <li>PROCESSING: Processing;</li>
	// <li>FINISH: Completed.</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// Creation time of a task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// Start time of task execution in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
		BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

		// End time of task execution in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

		// Information of a video processing task. This field has a value only when `TaskType` is `WorkflowTask`.
	// Note: This field may return null, indicating that no valid values can be obtained.
		WorkflowTask *WorkflowTask `json:"WorkflowTask,omitempty" name:"WorkflowTask"`

		// Video editing task information. This field has a value only when `TaskType` is `EditMediaTask`.
		EditMediaTask *EditMediaTask `json:"EditMediaTask,omitempty" name:"EditMediaTask"`

		// Information of a live stream processing task. This field has a value only when `TaskType` is `LiveStreamProcessTask`.
	// Note: This field may return null, indicating that no valid values can be obtained.
		LiveStreamProcessTask *LiveStreamProcessTask `json:"LiveStreamProcessTask,omitempty" name:"LiveStreamProcessTask"`

		// Event notification information of a task.
	// Note: This field may return null, indicating that no valid values can be obtained.
		TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

		// Task flow priority. Value range: [-10, 10].
		TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

		// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
		SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

		// Extended information field, used in specific scenarios.
		ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// Filter: Task status. Valid values: WAITING (waiting), PROCESSING (processing), FINISH (completed).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Scrolling identifier which is used for pulling in batches. If a single request cannot pull all the data entries, the API will return `ScrollToken`, and if the next request carries it, the next pull will start from the next entry.
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Limit")
	delete(f, "ScrollToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task overview list.
		TaskSet []*TaskSimpleInfo `json:"TaskSet,omitempty" name:"TaskSet"`

		// Scrolling identifier. If a request does not return all the data entries, this field indicates the ID of the next entry. If this field is an empty string, there is no more data.
		ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of transcoding templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Container format filter. Valid values:
	// <li>Video: Video container format that can contain both video stream and audio stream;</li>
	// <li>PureAudio: Audio container format that can contain only audio stream.</li>
	ContainerType *string `json:"ContainerType,omitempty" name:"ContainerType"`

	// TESHD filter, which is used to filter common transcoding or ultra-fast HD transcoding templates. Valid values:
	// <li>Common: Common transcoding template;</li>
	// <li>TEHD: TESHD template.</li>
	TEHDType *string `json:"TEHDType,omitempty" name:"TEHDType"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "ContainerType")
	delete(f, "TEHDType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of transcoding template details.
		TranscodeTemplateSet []*TranscodeTemplate `json:"TranscodeTemplateSet,omitempty" name:"TranscodeTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWatermarkTemplatesRequest struct {
	*tchttp.BaseRequest

	// Unique ID filter of watermarking templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Watermark type filter. Valid values:
	// <li>image: Image watermark;</li>
	// <li>text: Text watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries
	// <li>Default value: 10;</li>
	// <li>Maximum value: 100.</li>
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeWatermarkTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWatermarkTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWatermarkTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWatermarkTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of watermarking template details.
		WatermarkTemplateSet []*WatermarkTemplate `json:"WatermarkTemplateSet,omitempty" name:"WatermarkTemplateSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWatermarkTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWatermarkTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWordSamplesRequest struct {
	*tchttp.BaseRequest

	// Keyword filter. Array length limit: 100 words.
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// You can select multiple elements, which are connected by OR logic. If a usage contains any element in this parameter, the keyword sample will be used.
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Tag filter. Array length limit: 20 words.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeWordSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWordSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keywords")
	delete(f, "Usages")
	delete(f, "Tags")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
	// Note: This field may return null, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Keyword information.
	// Note: This field may return null, indicating that no valid values can be obtained.
		WordSet []*AiSampleWord `json:"WordSet,omitempty" name:"WordSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWordSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWordSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkflowsRequest struct {
	*tchttp.BaseRequest

	// Workflow ID filter. Array length limit: 100.
	WorkflowIds []*int64 `json:"WorkflowIds,omitempty" name:"WorkflowIds"`

	// Workflow status. Valid values:
	// <li>Enabled: Enabled,</li>
	// <li>Disabled: Disabled.</li>
	// If this parameter is left empty, the workflow status will not be distinguished.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Paging offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowIds")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Workflow information array.
		WorkflowInfoSet []*WorkflowInfo `json:"WorkflowInfoSet,omitempty" name:"WorkflowInfoSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableWorkflowRequest struct {
	*tchttp.BaseRequest

	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

func (r *DisableWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisableWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditMediaFileInfo struct {

	// Video input information.
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// Start time offset of video clipping in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of video clipping in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type EditMediaRequest struct {
	*tchttp.BaseRequest

	// Information of input video file.
	FileInfos []*EditMediaFileInfo `json:"FileInfos,omitempty" name:"FileInfos"`

	// Target storage of video processing output file.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Target path of video processing output file.
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// Event notification information of task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

	// Task priority. The higher the value, the higher the priority. Value range: -10–10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// The ID used for deduplication. If there was a request with the same ID in the last three days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

func (r *EditMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileInfos")
	delete(f, "OutputStorage")
	delete(f, "OutputObjectPath")
	delete(f, "TaskNotifyConfig")
	delete(f, "TasksPriority")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EditMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Video editing task ID, which can be used to query the status of an editing task.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditMediaTask struct {

	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task status. Valid values:
	// <li>PROCESSING: processing;</li>
	// <li>FINISH: completed.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of video editing task.
	Input *EditMediaTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of video editing task.
	Output *EditMediaTaskOutput `json:"Output,omitempty" name:"Output"`
}

type EditMediaTaskInput struct {

	// Information of input video file.
	FileInfoSet []*EditMediaFileInfo `json:"FileInfoSet,omitempty" name:"FileInfoSet"`
}

type EditMediaTaskOutput struct {

	// Target storage of edited file.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Path of edited video file.
	Path *string `json:"Path,omitempty" name:"Path"`
}

type EnableWorkflowRequest struct {
	*tchttp.BaseRequest

	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

func (r *EnableWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EnableWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteFunctionRequest struct {
	*tchttp.BaseRequest

	// Name of called backend API.
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// API parameter. Parameter format will depend on the actual function definition.
	FunctionArg *string `json:"FunctionArg,omitempty" name:"FunctionArg"`
}

func (r *ExecuteFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionName")
	delete(f, "FunctionArg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Packed string, which will vary according to the custom API.
		Result *string `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceConfigureInfo struct {

	// Switch of a face recognition task. Valid values:
	// <li>ON: Enables an intelligent face recognition task;</li>
	// <li>OFF: Disables an intelligent face recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Face recognition filter score. If this score is reached or exceeded, a recognition result will be returned. Value range: 0-100. Default value: 95.
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// Default figure filter tag, which specifies the default figure tag that needs to be returned. If this parameter is left empty or an empty value is entered, all results of the default figures will be returned. Valid values:
	// <li>entertainment: Entertainment celebrity;</li>
	// <li>sport: Sports celebrity;</li>
	// <li>politician: Politically sensitive figure.</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitempty" name:"DefaultLibraryLabelSet"`

	// Custom figure filter tag, which specifies the custom figure tag that needs to be returned. If this parameter is left empty or an empty value is entered, all results of the custom figures will be returned. Valid values:
	// There can be up to 10 tags, each with a length limit of 16 characters.
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitempty" name:"UserDefineLibraryLabelSet"`

	// Figure library. Valid values:
	// <li>Default: Default figure library;</li>
	// <li>UserDefine: Custom figure library.</li>
	// <li>All: Both default and custom figure libraries will be used.</li>
	// Default value: All (both default and custom figure libraries will be used.)
	FaceLibrary *string `json:"FaceLibrary,omitempty" name:"FaceLibrary"`
}

type FaceConfigureInfoForUpdate struct {

	// Switch of a face recognition task. Valid values:
	// <li>ON: Enables an intelligent face recognition task;</li>
	// <li>OFF: Disables an intelligent face recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Face recognition filter score. If this score is reached or exceeded, a recognition result will be returned. Value range: 0-100.
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// Default figure filter tag, which specifies the default figure tag that needs to be returned. If this parameter is left empty or an empty value is entered, all results of the default figures will be returned. Valid values:
	// <li>entertainment: Entertainment celebrity;</li>
	// <li>sport: Sports celebrity;</li>
	// <li>politician: Politically sensitive figure.</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitempty" name:"DefaultLibraryLabelSet"`

	// Custom figure filter tag, which specifies the custom figure tag that needs to be returned. If this parameter is left empty or an empty value is entered, all results of the custom figures will be returned. Valid values:
	// There can be up to 10 tags, each with a length limit of 16 characters.
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitempty" name:"UserDefineLibraryLabelSet"`

	// Figure library. Valid values:
	// <li>Default: Default figure library;</li>
	// <li>UserDefine: Custom figure library.</li>
	// <li>All: Both default and custom figure libraries will be used.</li>
	FaceLibrary *string `json:"FaceLibrary,omitempty" name:"FaceLibrary"`
}

type FrameTagConfigureInfo struct {

	// Switch of intelligent frame-specific tagging task. Valid values:
	// <li>ON: enables intelligent frame-specific tagging task;</li>
	// <li>OFF: disables intelligent frame-specific tagging task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type FrameTagConfigureInfoForUpdate struct {

	// Switch of intelligent frame-specific tagging task. Valid values:
	// <li>ON: enables intelligent frame-specific tagging task;</li>
	// <li>OFF: disables intelligent frame-specific tagging task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ImageSpriteTaskInput struct {

	// ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Target bucket of a generated image sprite. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Output path to a generated image sprite file, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_imageSprite_{definition}_{number}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// Output path to the WebVTT file after an image sprite is generated, which can only be a relative path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_imageSprite_{definition}.{format}`.
	WebVttObjectName *string `json:"WebVttObjectName,omitempty" name:"WebVttObjectName"`

	// Rule of the `{number}` variable in the image sprite output path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type ImageSpriteTemplate struct {

	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of an image sprite generating template.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subimage width of an image sprite.
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Subimage height of an image sprite.
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

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: Stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: Fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// Template description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type ImageWatermarkInput struct {

	// String generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) a watermark image. JPEG and PNG images are supported.
	ImageContent *string `json:"ImageContent,omitempty" name:"ImageContent"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width. For example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in pixels. For example, `100px` means that `Width` is 100 pixels. Value range: [8, 4096].</li>
	// Default value: 10%.
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height. For example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in pixels. For example, `100px` means that `Height` is 100 pixels. Value range: 0 or [8, 4096].</li>
	// Default value: 0px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
	Height *string `json:"Height,omitempty" name:"Height"`

	// Repeat type of an animated watermark. Valid values:
	// <li>once: no longer appears after watermark playback ends.</li>
	// <li>repeat_last_frame: stays on the last frame after watermark playback ends.</li>
	// <li>repeat (default): repeats the playback until the video ends.</li>
	RepeatType *string `json:"RepeatType,omitempty" name:"RepeatType"`
}

type ImageWatermarkInputForUpdate struct {

	// String generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) a watermark image. JPEG and PNG images are supported.
	ImageContent *string `json:"ImageContent,omitempty" name:"ImageContent"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width. For example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in pixels. For example, `100px` means that `Width` is 100 pixels. Value range: [8, 4096].</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height. For example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in pixels. For example, `100px` means that `Height` is 100 pixels. Value range: 0 or [8, 4096].</li>
	// Default value: 0px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
	Height *string `json:"Height,omitempty" name:"Height"`

	// Repeat type of an animated watermark. Valid values:
	// <li>once: no longer appears after watermark playback ends.</li>
	// <li>repeat_last_frame: stays on the last frame after watermark playback ends.</li>
	// <li>repeat (default): repeats the playback until the video ends.</li>
	RepeatType *string `json:"RepeatType,omitempty" name:"RepeatType"`
}

type ImageWatermarkTemplate struct {

	// Watermark image address.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height. For example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in pixels. For example, `100px` means that `Height` is 100 pixels.</li>
	// `0px` means that `Height` will be proportionally scaled according to the video width.
	Height *string `json:"Height,omitempty" name:"Height"`

	// Repeat type of an animated watermark. Valid values:
	// <li>once: no longer appears after watermark playback ends.</li>
	// <li>repeat_last_frame: stays on the last frame after watermark playback ends.</li>
	// <li>repeat (default): repeats the playback until the video ends.</li>
	RepeatType *string `json:"RepeatType,omitempty" name:"RepeatType"`
}

type LiveStreamAiRecognitionResultInfo struct {

	// Content recognition result list.
	ResultSet []*LiveStreamAiRecognitionResultItem `json:"ResultSet,omitempty" name:"ResultSet"`
}

type LiveStreamAiRecognitionResultItem struct {

	// Result type. Valid values:
	// <li>FaceRecognition: face recognition,</li>
	// <li>AsrWordsRecognition: speech keyword recognition,</li>
	// <li>OcrWordsRecognition: text keyword recognition,</li>
	// <li>AsrFullTextRecognition: full speech recognition,</li>
	// <li>OcrFullTextRecognition: full text recognition.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Face recognition result, which is valid when `Type` is
	// `FaceRecognition`.
	FaceRecognitionResultSet []*LiveStreamFaceRecognitionResult `json:"FaceRecognitionResultSet,omitempty" name:"FaceRecognitionResultSet"`

	// Speech keyword recognition result, which is valid when `Type` is
	// `AsrWordsRecognition`.
	AsrWordsRecognitionResultSet []*LiveStreamAsrWordsRecognitionResult `json:"AsrWordsRecognitionResultSet,omitempty" name:"AsrWordsRecognitionResultSet"`

	// Text keyword recognition result, which is valid when `Type` is
	// `OcrWordsRecognition`.
	OcrWordsRecognitionResultSet []*LiveStreamOcrWordsRecognitionResult `json:"OcrWordsRecognitionResultSet,omitempty" name:"OcrWordsRecognitionResultSet"`

	// Full speech recognition result, which is valid when `Type` is
	// `AsrFullTextRecognition`.
	AsrFullTextRecognitionResultSet []*LiveStreamAsrFullTextRecognitionResult `json:"AsrFullTextRecognitionResultSet,omitempty" name:"AsrFullTextRecognitionResultSet"`

	// Full text recognition result, which is valid when `Type` is
	// `OcrFullTextRecognition`.
	OcrFullTextRecognitionResultSet []*LiveStreamOcrFullTextRecognitionResult `json:"OcrFullTextRecognitionResultSet,omitempty" name:"OcrFullTextRecognitionResultSet"`
}

type LiveStreamAiReviewImagePoliticalResult struct {

	// Start PTS time of a suspected segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitempty" name:"StartPtsTime"`

	// End PTS time of a suspected segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitempty" name:"EndPtsTime"`

	// Score of a suspected politically sensitive segment.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for porn information detection of a suspected segment. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tag of the detected politically sensitive information in video. Valid values:
	// <li>politician: Politically sensitive figure.</li>
	// <li>violation_photo: Violating photo.</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// Name of a politically sensitive figure or violating photo.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Zone coordinates (at the pixel level) of a politically sensitive figure or violating photo: [x1, y1, x2, y2], i.e., the coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type LiveStreamAiReviewImagePornResult struct {

	// Start PTS time of a suspected segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitempty" name:"StartPtsTime"`

	// End PTS time of a suspected segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitempty" name:"EndPtsTime"`

	// Score of a suspected porn segment.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for porn information detection of a suspected segment. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tag of the detected porn information in video. Valid values:
	// <li>porn: Porn.</li>
	// <li>sexy: Sexiness.</li>
	// <li>vulgar: Vulgarity.</li>
	// <li>intimacy: Intimacy.</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type LiveStreamAiReviewImageTerrorismResult struct {

	// Start PTS time of a suspected segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitempty" name:"StartPtsTime"`

	// End PTS time of a suspected segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitempty" name:"EndPtsTime"`

	// Score of a suspected terrorism segment.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for terrorism information detection of a suspected segment. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tag of the detected terrorism information in a video. Valid values:
	// <li>guns: Weapons and guns.</li>
	// <li>crowd: Crowd.</li>
	// <li>police: Police force.</li>
	// <li>bloody: Bloody scenes.</li>
	// <li>banners: Terrorism flags.</li>
	// <li>militant: Militants.</li>
	// <li>explosion: Explosions and fires.</li>
	// <li>terrorists: Terrorists.</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type LiveStreamAiReviewResultInfo struct {

	// List of content audit results.
	ResultSet []*LiveStreamAiReviewResultItem `json:"ResultSet,omitempty" name:"ResultSet"`
}

type LiveStreamAiReviewResultItem struct {

	// Content audit type. Valid values:
	// <li>ImagePorn: Porn information detection in image</li>
	// <li>ImageTerrorism: Terrorism information detection in image</li>
	// <li>ImagePolitical: Politically sensitive information detection in image</li>
	// <li>PornVoice: Porn information detection in speech</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Result of porn information detection in image, which is valid when `Type` is `ImagePorn`.
	ImagePornResultSet []*LiveStreamAiReviewImagePornResult `json:"ImagePornResultSet,omitempty" name:"ImagePornResultSet"`

	// Result of terrorism information detection in image, which is valid when `Type` is `ImageTerrorism`.
	ImageTerrorismResultSet []*LiveStreamAiReviewImageTerrorismResult `json:"ImageTerrorismResultSet,omitempty" name:"ImageTerrorismResultSet"`

	// Result of politically sensitive information detection in image, which is valid when `Type` is `ImagePolitical`.
	ImagePoliticalResultSet []*LiveStreamAiReviewImagePoliticalResult `json:"ImagePoliticalResultSet,omitempty" name:"ImagePoliticalResultSet"`

	// Result of porn information detection in speech, which is valid when `Type` is `PornVoice`.
	VoicePornResultSet []*LiveStreamAiReviewVoicePornResult `json:"VoicePornResultSet,omitempty" name:"VoicePornResultSet"`
}

type LiveStreamAiReviewVoicePornResult struct {

	// Start PTS time of a suspected segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitempty" name:"StartPtsTime"`

	// End PTS time of a suspected segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitempty" name:"EndPtsTime"`

	// Score of a suspected porn segment.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for porn information detection of a suspected segment. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Tag of the detected porn information in video. Valid values:
	// <li>sexual_moan: Sexual moans.</li>
	Label *string `json:"Label,omitempty" name:"Label"`
}

type LiveStreamAsrFullTextRecognitionResult struct {

	// Recognized text.
	Text *string `json:"Text,omitempty" name:"Text"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type LiveStreamAsrWordsRecognitionResult struct {

	// Speech keyword.
	Word *string `json:"Word,omitempty" name:"Word"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type LiveStreamFaceRecognitionResult struct {

	// Unique ID of figure.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Figure name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Figure library type, indicating to which figure library the recognized figure belongs:
	// <li>Default: default figure library</li><li>UserDefine: custom figure library</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`
}

type LiveStreamOcrFullTextRecognitionResult struct {

	// Speech text.
	Text *string `json:"Text,omitempty" name:"Text"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`
}

type LiveStreamOcrWordsRecognitionResult struct {

	// Text keyword.
	Word *string `json:"Word,omitempty" name:"Word"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoords []*int64 `json:"AreaCoords,omitempty" name:"AreaCoords"`
}

type LiveStreamProcessErrorInfo struct {

	// Error code:
	// <li>0: No error;</li>
	// <li>If this parameter is not 0, an error has occurred. Please see the error message (`Message`).</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type LiveStreamProcessTask struct {

	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task flow status. Valid values:
	// <li>PROCESSING: Processing;</li>
	// <li>FINISH: Completed.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Live stream URL.
	Url *string `json:"Url,omitempty" name:"Url"`
}

type LiveStreamTaskNotifyConfig struct {

	// CMQ model. There are two types: `Queue` and `Topic`. Currently, only `Queue` is supported.
	CmqModel *string `json:"CmqModel,omitempty" name:"CmqModel"`

	// CMQ region, such as `sh` and `bj`.
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// This parameter is valid when the model is `Queue`, indicating the name of the CMQ queue for receiving event notifications.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// This parameter is valid when the model is `Topic`, indicating the name of the CMQ topic for receiving event notifications.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ManageTaskRequest struct {
	*tchttp.BaseRequest

	// Operation type. Valid values:
	// <ul>
	// <li>Abort: task termination. Description:
	// <ul><li>If the [task type](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) is live stream processing (`LiveStreamProcessTask`), tasks whose [task status](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) is `WAITING` or `PROCESSING` can be terminated.</li>
	// <li>For other [task types](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0), only tasks whose [task status](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) is `WAITING` can be terminated.</li></ul>
	// </li></ul>
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ManageTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationType")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ManageTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MediaAiAnalysisClassificationItem struct {

	// Name of intelligently generated category.
	Classification *string `json:"Classification,omitempty" name:"Classification"`

	// Confidence of intelligently generated category between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisCoverItem struct {

	// Storage path of intelligently generated cover.
	CoverPath *string `json:"CoverPath,omitempty" name:"CoverPath"`

	// Confidence of intelligently generated cover between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisFrameTagItem struct {

	// Frame-specific tag name.
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 
	CategorySet []*string `json:"CategorySet,omitempty" name:"CategorySet"`

	// Confidence of intelligently generated frame-specific tag between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisFrameTagSegmentItem struct {

	// Start time offset of frame-specific tag.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of frame-specific tag.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// List of tags in time period.
	TagSet []*MediaAiAnalysisFrameTagItem `json:"TagSet,omitempty" name:"TagSet"`
}

type MediaAiAnalysisTagItem struct {

	// Tag name.
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// Confidence of tag between 0 and 100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAnimatedGraphicsItem struct {

	// Storage location of a generated animated image file.
	Storage *TaskOutputStorage `json:"Storage,omitempty" name:"Storage"`

	// Path to a generated animated image file.
	Path *string `json:"Path,omitempty" name:"Path"`

	// ID of an animated image generating template. For more information, please see [Animated Image Generating Parameter Template](https://intl.cloud.tencent.com/document/product/266/33481?from_cn_redirect=1#.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Animated image format, such as gif.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Height of an animated image in px.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Width of an animated image in px.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Bitrate of an animated image in bps.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Size of an animated image in bytes.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// MD5 value of an animated image.
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// Start time offset of an animated image in the video in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of an animated image in the video in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type MediaAudioStreamItem struct {

	// Bitrate of an audio stream in bps.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Sample rate of an audio stream in Hz.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SamplingRate *int64 `json:"SamplingRate,omitempty" name:"SamplingRate"`

	// Audio stream codec, such as aac.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type MediaContentReviewAsrTextSegmentItem struct {

	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of a suspected segment.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for suspected segment audit. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of suspected keywords.
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet"`
}

type MediaContentReviewOcrTextSegmentItem struct {

	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of a suspected segment.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Suggestion for suspected segment audit. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of suspected keywords.
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet"`

	// Zone coordinates (at the pixel level) of suspected text: [x1, y1, x2, y2], i.e., the coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
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

	// Tags for the results of politically sensitive information detection of suspected video segments. The relationship between the `LabelSet` parameter in the content audit template [controlling tasks of video politically sensitive information detection](https://intl.cloud.tencent.com/document/api/862/37615?from_cn_redirect=1#PoliticalImgReviewTemplateInfo) and this parameter is as follows:
	// violation_photo:
	// <li>violation_photo: violating photo.</li>
	// politician:
	// <li>nation_politician: head of state/government;</li>
	// <li>province_politician: province/state leader;</li>
	// <li>bureau_politician: ministry leader;</li>
	// <li>county_politician: county/city leader;</li>
	// <li>rural_politician: town leader;</li>
	// <li>sensitive_politician: politically sensitive figure.</li>
	// <li>foreign_politician: head of a foreign country/government.</li>
	// entertainment:
	// <li>sensitive_entertainment: sensitive entertainment celebrity.</li>
	// sport:
	// <li>sensitive_sport: sensitive sports figure.</li>
	// entrepreneur:
	// <li>sensitive_entrepreneur: sensitive business figure.</li>
	// scholar:
	// <li>sensitive_scholar: sensitive educator.</li>
	// celebrity:
	// <li>sensitive_celebrity: sensitive well-known figure;</li>
	// <li>historical_celebrity: well-known historical figure.</li>
	// military:
	// <li>sensitive_military: militarily sensitive figure.</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// URL of a suspected image (which will not be permanently stored
	//  and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// Zone coordinates (at the pixel level) of a politically sensitive figure or violating photo: [x1, y1, x2, y2], i.e., the coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
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

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type MediaImageSpriteItem struct {

	// Image sprite specification. For more information, please see [Image Sprite Parameter Template](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Subimage height of an image sprite.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Subimage width of an image sprite.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Total number of subimages in each image sprite.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Path to each image sprite.
	ImagePathSet []*string `json:"ImagePathSet,omitempty" name:"ImagePathSet"`

	// Path to a WebVtt file for the position-time relationship among subimages in an image sprite. The WebVtt file indicates the corresponding time points of each subimage and their coordinates in the image sprite, which is typically used by the player for implementing preview.
	WebVttPath *string `json:"WebVttPath,omitempty" name:"WebVttPath"`

	// Storage location of an image sprite file.
	Storage *TaskOutputStorage `json:"Storage,omitempty" name:"Storage"`
}

type MediaInputInfo struct {

	// The type of video processing input object, which supports COS and URL.
	Type *string `json:"Type,omitempty" name:"Type"`

	// This parameter is required and valid when `Type` is `COS`, indicating the information of a COS object for video processing.
	CosInputInfo *CosInputInfo `json:"CosInputInfo,omitempty" name:"CosInputInfo"`

	// This parameter is required and valid when `Type` is `URL`, indicating the information of a URL object for video processing.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	UrlInputInfo *UrlInputInfo `json:"UrlInputInfo,omitempty" name:"UrlInputInfo"`
}

type MediaMetaData struct {

	// Size of an uploaded media file in bytes (which is the sum of size of m3u8 and ts files if the video is in HLS format).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Container, such as m4a and mp4.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Sum of the average bitrate of a video stream and that of an audio stream in bps.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Maximum value of the height of a video stream in px.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Maximum value of the width of a video stream in px.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Video duration in seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Selected angle during video recording in degrees.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// Video stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet"`

	// Audio stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet"`

	// Video duration in seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoDuration *float64 `json:"VideoDuration,omitempty" name:"VideoDuration"`

	// Audio duration in seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioDuration *float64 `json:"AudioDuration,omitempty" name:"AudioDuration"`
}

type MediaProcessTaskAdaptiveDynamicStreamingResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of an adaptive bitrate streaming task.
	Input *AdaptiveDynamicStreamingTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of an adaptive bitrate streaming task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AdaptiveDynamicStreamingInfoItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskAnimatedGraphicResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for an animated image generating task.
	Input *AnimatedGraphicTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of an animated image generating task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaAnimatedGraphicsItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskImageSpriteResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for an image sprite generating task.
	Input *ImageSpriteTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of an image sprite generating task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaImageSpriteItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskInput struct {

	// List of transcoding tasks.
	TranscodeTaskSet []*TranscodeTaskInput `json:"TranscodeTaskSet,omitempty" name:"TranscodeTaskSet"`

	// List of animated image generating tasks.
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInput `json:"AnimatedGraphicTaskSet,omitempty" name:"AnimatedGraphicTaskSet"`

	// List of time point screencapturing tasks.
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTaskSet,omitempty" name:"SnapshotByTimeOffsetTaskSet"`

	// List of sampled screencapturing tasks.
	SampleSnapshotTaskSet []*SampleSnapshotTaskInput `json:"SampleSnapshotTaskSet,omitempty" name:"SampleSnapshotTaskSet"`

	// List of image sprite generating tasks.
	ImageSpriteTaskSet []*ImageSpriteTaskInput `json:"ImageSpriteTaskSet,omitempty" name:"ImageSpriteTaskSet"`

	// List of adaptive bitrate streaming tasks.
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTaskSet,omitempty" name:"AdaptiveDynamicStreamingTaskSet"`
}

type MediaProcessTaskResult struct {

	// Task type. Valid values:
	// <li>Transcode: Transcoding</li>
	// <li>AnimatedGraphics: Animated image generating</li>
	// <li>SnapshotByTimeOffset: Time point screencapturing</li>
	// <li>SampleSnapshot: Sampled screencapturing</li>
	// <li>ImageSprites: Image sprite generating</li>
	// <li>CoverBySnapshot: Screencapturing for cover image</li>
	// <li>AdaptiveDynamicStreaming: Adaptive bitrate streaming</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Query result of a transcoding task, which is valid when task type is `Transcode`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitempty" name:"TranscodeTask"`

	// Query result of an animated image generating task, which is valid when task type is `AnimatedGraphics`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitempty" name:"AnimatedGraphicTask"`

	// Query result of a time point screencapturing task, which is valid when task type is `SnapshotByTimeOffset`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitempty" name:"SnapshotByTimeOffsetTask"`

	// Query result of a sampled screencapturing task, which is valid when task type is `SampleSnapshot`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitempty" name:"SampleSnapshotTask"`

	// Query result of an image sprite generating task, which is valid when task type is `ImageSprite`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitempty" name:"ImageSpriteTask"`

	// Query result of an adaptive bitrate streaming task, which is valid if the task type is `AdaptiveDynamicStreaming`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdaptiveDynamicStreamingTask *MediaProcessTaskAdaptiveDynamicStreamingResult `json:"AdaptiveDynamicStreamingTask,omitempty" name:"AdaptiveDynamicStreamingTask"`
}

type MediaProcessTaskSampleSnapshotResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for a sampled screencapturing task.
	Input *SampleSnapshotTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of a sampled screencapturing task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaSampleSnapshotItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskSnapshotByTimeOffsetResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for a time point screencapturing task.
	Input *SnapshotByTimeOffsetTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of a time point screencapturing task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaSnapshotByTimeOffsetItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskTranscodeResult struct {

	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; otherwise it is failed. For details about the values, see [Error Code List](https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for a transcoding task.
	Input *TranscodeTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of a transcoding task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaTranscodeItem `json:"Output,omitempty" name:"Output"`
}

type MediaSampleSnapshotItem struct {

	// Sampled screenshot specification ID. For more information, please see [Sampled Screencapturing Parameter Template](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Sample type. Valid values:
	// <li>Percent: Samples at the specified percentage interval.</li>
	// <li>Time: Samples at the specified time interval.</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval
	// <li>If `SampleType` is `Percent`, this value means taking a screenshot at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, this value means taking a screenshot at an interval of the specified time (in seconds). The first screenshot is always the first video frame.</li>
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// Storage location of a generated screenshot file.
	Storage *TaskOutputStorage `json:"Storage,omitempty" name:"Storage"`

	// List of paths to generated screenshots.
	ImagePathSet []*string `json:"ImagePathSet,omitempty" name:"ImagePathSet"`

	// List of watermarking template IDs if the screenshots are watermarked.
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition"`
}

type MediaSnapshotByTimeOffsetItem struct {

	// Specification of a time point screenshot. For more information, please see [Parameter Template for Time Point Screencapturing](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Information set of screenshots of the same specification. Each element represents a screenshot.
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet,omitempty" name:"PicInfoSet"`

	// Location of a time point screenshot file.
	Storage *TaskOutputStorage `json:"Storage,omitempty" name:"Storage"`
}

type MediaSnapshotByTimePicInfoItem struct {

	// Time offset corresponding to the screenshot in the video in <font color=red>milliseconds</font>.
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// Path to the screenshot.
	Path *string `json:"Path,omitempty" name:"Path"`

	// List of watermarking template IDs if the screenshots are watermarked.
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition"`
}

type MediaTranscodeItem struct {

	// Target bucket of an output file.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Path to an output video file.
	Path *string `json:"Path,omitempty" name:"Path"`

	// Transcoding specification ID. For more information, please see [Transcoding Parameter Template](https://intl.cloud.tencent.com/document/product/266/33478?from_cn_redirect=1#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Sum of the average bitrate of a video stream and that of an audio stream in bps.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Maximum value of the height of a video stream in px.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Maximum value of the width of a video stream in px.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Total size of a media file in bytes (which is the sum of size of m3u8 and ts files if the video is in HLS format).
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Video duration in seconds.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Container, such as m4a and mp4.
	Container *string `json:"Container,omitempty" name:"Container"`

	// MD5 value of a video.
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// Audio stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet"`

	// Video stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet"`
}

type MediaVideoStreamItem struct {

	// Bitrate of a video stream in bps.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Height of a video stream in px.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Width of a video stream in px.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Video stream codec, such as h264.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Frame rate in Hz.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
}

func (r *ModifyAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAIAnalysisTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ClassificationConfigure")
	delete(f, "TagConfigure")
	delete(f, "CoverConfigure")
	delete(f, "FrameTagConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAIAnalysisTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAIAnalysisTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Name of a video content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of a video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Face recognition control parameter.
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// Full text recognition control parameter.
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// Text keyword recognition control parameter.
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// Full speech recognition control parameter.
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// Speech keyword recognition control parameter.
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`
}

func (r *ModifyAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAIRecognitionTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "FaceConfigure")
	delete(f, "OcrFullTextConfigure")
	delete(f, "OcrWordsConfigure")
	delete(f, "AsrFullTextConfigure")
	delete(f, "AsrWordsConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAIRecognitionTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAIRecognitionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Adaptive bitrate streaming format. Valid values:
	// <li>HLS,</li>
	// <li>MPEG-DASH.</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Whether to prohibit transcoding from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitempty" name:"DisableHigherVideoResolution"`

	// Parameter information of input streams for transcoding to adaptive bitrate streaming. Up to 10 streams can be input.
	// Note: the frame rate of each stream must be consistent; otherwise, the frame rate of the first stream is used as the output frame rate.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitempty" name:"StreamInfos"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *ModifyAdaptiveDynamicStreamingTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAdaptiveDynamicStreamingTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Format")
	delete(f, "DisableHigherVideoBitrate")
	delete(f, "DisableHigherVideoResolution")
	delete(f, "StreamInfos")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAdaptiveDynamicStreamingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAdaptiveDynamicStreamingTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAdaptiveDynamicStreamingTemplateResponse) FromJsonString(s string) error {
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

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
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
}

func (r *ModifyAnimatedGraphicsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAnimatedGraphicsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Fps")
	delete(f, "Quality")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAnimatedGraphicsTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAnimatedGraphicsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContentReviewTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an intelligent content recognition template
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Name of an intelligent content recognition template. Length limit: 64 characters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of an intelligent content recognition template. Length limit: 256 characters
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter for porn information
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// Control parameter for terrorism information
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// Control parameter for politically sensitive information
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this parameter is not supported yet.
	ProhibitedConfigure *ProhibitedConfigureInfoForUpdate `json:"ProhibitedConfigure,omitempty" name:"ProhibitedConfigure"`

	// Control parameter for custom intelligent content recognition tasks
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`
}

func (r *ModifyContentReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContentReviewTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "PornConfigure")
	delete(f, "TerrorismConfigure")
	delete(f, "PoliticalConfigure")
	delete(f, "ProhibitedConfigure")
	delete(f, "UserDefineConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContentReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContentReviewTemplateResponse) FromJsonString(s string) error {
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
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Subimage row count of an image sprite.
	RowCount *uint64 `json:"RowCount,omitempty" name:"RowCount"`

	// Subimage column count of an image sprite.
	ColumnCount *uint64 `json:"ColumnCount,omitempty" name:"ColumnCount"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *ModifyImageSpriteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSpriteTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "SampleType")
	delete(f, "SampleInterval")
	delete(f, "RowCount")
	delete(f, "ColumnCount")
	delete(f, "FillType")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageSpriteTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSpriteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonSampleRequest struct {
	*tchttp.BaseRequest

	// Image ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Name. Length limit: 128 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description. Length limit: 1,024 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Image usage. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: used for content recognition and inappropriate information recognition; equivalent to 1+2
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Information of operations on facial features
	FaceOperationInfo *AiSampleFaceOperation `json:"FaceOperationInfo,omitempty" name:"FaceOperationInfo"`

	// Tag operation information.
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitempty" name:"TagOperationInfo"`
}

func (r *ModifyPersonSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Usages")
	delete(f, "FaceOperationInfo")
	delete(f, "TagOperationInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPersonSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Image information
		Person *AiSamplePerson `json:"Person,omitempty" name:"Person"`

		// Information of images that failed the verification by facial feature positioning.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitempty" name:"FailFaceInfoSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPersonSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Sampled screencapturing type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Image format. Valid values: jpg; png.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySampleSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "SampleType")
	delete(f, "SampleInterval")
	delete(f, "Format")
	delete(f, "Comment")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySampleSnapshotTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySampleSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Name of a time point screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotByTimeOffsetTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	delete(f, "Format")
	delete(f, "Comment")
	delete(f, "FillType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotByTimeOffsetTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotByTimeOffsetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a transcoding template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Container format. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Name of a transcoding template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Whether to remove video data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter.
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter.
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter. To enable it, please contact your Tencent Cloud sales rep.
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitempty" name:"TEHDConfig"`
}

func (r *ModifyTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Container")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	delete(f, "VideoTemplate")
	delete(f, "AudioTemplate")
	delete(f, "TEHDConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTranscodeTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Watermarking template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Origin position. Valid values:
	// <li>TopLeft: The origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>TopRight: The origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>BottomLeft: The origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>BottomRight: The origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
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

	// SVG watermarking template. This field is required when `Type` is `svg` and is invalid when `Type` is `image` or `text`.
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitempty" name:"SvgTemplate"`
}

func (r *ModifyWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWatermarkTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "CoordinateOrigin")
	delete(f, "XPos")
	delete(f, "YPos")
	delete(f, "ImageTemplate")
	delete(f, "TextTemplate")
	delete(f, "SvgTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWatermarkTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Image watermark address. This field is valid only when `ImageTemplate.ImageContent` is non-empty.
		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWatermarkTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWordSampleRequest struct {
	*tchttp.BaseRequest

	// Keyword. Length limit: 128 characters.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// 7. All: equivalent to 1+2+3+4
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Tag operation information.
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitempty" name:"TagOperationInfo"`
}

func (r *ModifyWordSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWordSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keyword")
	delete(f, "Usages")
	delete(f, "TagOperationInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWordSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWordSampleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWordSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWordSampleResponse) FromJsonString(s string) error {
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

type NumberFormat struct {

	// Start value of the `{number}` variable. Default value: 0.
	InitialValue *uint64 `json:"InitialValue,omitempty" name:"InitialValue"`

	// Increment of the `{number}` variable. Default value: 1.
	Increment *uint64 `json:"Increment,omitempty" name:"Increment"`

	// Minimum length of the `{number}` variable. A placeholder will be used if the variable length is below the minimum requirement. Default value: 1.
	MinLength *uint64 `json:"MinLength,omitempty" name:"MinLength"`

	// Placeholder used when the `{number}` variable length is below the minimum requirement. Default value: 0.
	PlaceHolder *string `json:"PlaceHolder,omitempty" name:"PlaceHolder"`
}

type OcrFullTextConfigureInfo struct {

	// Switch of a full text recognition task. Valid values:
	// <li>ON: Enables an intelligent full text recognition task;</li>
	// <li>OFF: Disables an intelligent full text recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type OcrFullTextConfigureInfoForUpdate struct {

	// Switch of a full text recognition task. Valid values:
	// <li>ON: Enables an intelligent full text recognition task;</li>
	// <li>OFF: Disables an intelligent full text recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type OcrWordsConfigureInfo struct {

	// Switch of a text keyword recognition task. Valid values:
	// <li>ON: Enables a text keyword recognition task;</li>
	// <li>OFF: Disables a text keyword recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`
}

type OcrWordsConfigureInfoForUpdate struct {

	// Switch of a text keyword recognition task. Valid values:
	// <li>ON: Enables a text keyword recognition task;</li>
	// <li>OFF: Disables a text keyword recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`
}

type OverrideTranscodeParameter struct {

	// Container format. Valid values: mp4, flv, hls, mp3, flac, ogg, and m4a; mp3, flac, ogg, and m4a are formats of audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Whether to remove video data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	RemoveVideo *uint64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	RemoveAudio *uint64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter.
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter.
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter.
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitempty" name:"TEHDConfig"`
}

type ParseLiveStreamProcessNotificationRequest struct {
	*tchttp.BaseRequest

	// Live stream event notification obtained from CMQ.
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *ParseLiveStreamProcessNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseLiveStreamProcessNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseLiveStreamProcessNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ParseLiveStreamProcessNotificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result type of live stream processing. Valid values:
	// <li>AiReviewResult: Content audit result;</li>
	// <li>ProcessEof: Live stream processing has been completed.</li>
		NotificationType *string `json:"NotificationType,omitempty" name:"NotificationType"`

		// Video processing task ID.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// Information of a live stream processing error, which is valid when `NotificationType` is `ProcessEof`.
	// Note: This field may return null, indicating that no valid values can be obtained.
		ProcessEofInfo *LiveStreamProcessErrorInfo `json:"ProcessEofInfo,omitempty" name:"ProcessEofInfo"`

		// Content audit result, which is valid when `NotificationType` is `AiReviewResult`.
	// Note: This field may return null, indicating that no valid values can be obtained.
		AiReviewResultInfo *LiveStreamAiReviewResultInfo `json:"AiReviewResultInfo,omitempty" name:"AiReviewResultInfo"`

		// Content recognition result, which is valid if `NotificationType` is `AiRecognitionResult`.
		AiRecognitionResultInfo *LiveStreamAiRecognitionResultInfo `json:"AiRecognitionResultInfo,omitempty" name:"AiRecognitionResultInfo"`

		// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
		SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ParseLiveStreamProcessNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseLiveStreamProcessNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseNotificationRequest struct {
	*tchttp.BaseRequest

	// Event notification obtained from CMQ.
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *ParseNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ParseNotificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Supported event type. Valid values:
	// <li>WorkflowTask: Video workflow processing task.</li>
		EventType *string `json:"EventType,omitempty" name:"EventType"`

		// Information of a video processing task. This field has a value only when `TaskType` is `WorkflowTask`.
	// Note: This field may return null, indicating that no valid values can be obtained.
		WorkflowTaskEvent *WorkflowTask `json:"WorkflowTaskEvent,omitempty" name:"WorkflowTaskEvent"`

		// Video editing task information. This field has a value only when `TaskType` is `EditMediaTask`.
		EditMediaTaskEvent *EditMediaTask `json:"EditMediaTaskEvent,omitempty" name:"EditMediaTaskEvent"`

		// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
		SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ParseNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PoliticalAsrReviewTemplateInfo struct {

	// Switch of a politically sensitive information detection in speech task. Valid values:
	// <li>ON: Enables a politically sensitive information detection in speech task;</li>
	// <li>OFF: Disables a politically sensitive information detection in speech task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalAsrReviewTemplateInfoForUpdate struct {

	// Switch of a politically sensitive information detection in speech task. Valid values:
	// <li>ON: Enables a politically sensitive information detection in speech task;</li>
	// <li>OFF: Disables a politically sensitive information detection in speech task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalConfigureInfo struct {

	// Control parameter of politically sensitive information detection in image.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImgReviewInfo *PoliticalImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Control parameter of politically sensitive information detection in speech.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrReviewInfo *PoliticalAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of politically sensitive information detection in text.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *PoliticalOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PoliticalConfigureInfoForUpdate struct {

	// Control parameter of politically sensitive information detection in image.
	ImgReviewInfo *PoliticalImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Control parameter of politically sensitive information detection in speech.
	AsrReviewInfo *PoliticalAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of politically sensitive information detection in text.
	OcrReviewInfo *PoliticalOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PoliticalImgReviewTemplateInfo struct {

	// Switch of a politically sensitive information detection in image task. Valid values:
	// <li>ON: Enables a politically sensitive information detection in image task;</li>
	// <li>OFF: Disables a politically sensitive information detection in image task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter tags for politically sensitive information detection of video images. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. Valid values:
	// <li>violation_photo: violating photo;</li>
	// <li>politician: political figure;</li>
	// <li>entertainment: entertainment celebrity;</li>
	// <li>sport: sports figure;</li>
	// <li>entrepreneur: business figure;</li>
	// <li>scholar: educator;</li>
	// <li>celebrity: well-known figure;</li>
	// <li>military: military figure.</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 97 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 95 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalImgReviewTemplateInfoForUpdate struct {

	// Switch of a politically sensitive information detection in image task. Valid values:
	// <li>ON: Enables a politically sensitive information detection in image task;</li>
	// <li>OFF: Disables a politically sensitive information detection in image task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter tags for politically sensitive information detection of video images. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. Valid values:
	// <li>violation_photo: violating photo;</li>
	// <li>politician: political figure;</li>
	// <li>entertainment: entertainment celebrity;</li>
	// <li>sport: sports figure;</li>
	// <li>entrepreneur: business figure;</li>
	// <li>scholar: educator;</li>
	// <li>celebrity: well-known figure;</li>
	// <li>military: military figure.</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfo struct {

	// Switch of a politically sensitive information detection in text task. Valid values:
	// <li>ON: Enables a politically sensitive information detection in text task;</li>
	// <li>OFF: Disables a politically sensitive information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfoForUpdate struct {

	// Switch of a politically sensitive information detection in text task. Valid values:
	// <li>ON: Enables a politically sensitive information detection in text task;</li>
	// <li>OFF: Disables a politically sensitive information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfo struct {

	// Switch of a porn information detection in speech task. Valid values:
	// <li>ON: Enables a porn information detection in speech task;</li>
	// <li>OFF: Disables a porn information detection in speech task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfoForUpdate struct {

	// Switch of a porn information detection in speech task. Valid values:
	// <li>ON: Enables a porn information detection in speech task;</li>
	// <li>OFF: Disables a porn information detection in speech task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornConfigureInfo struct {

	// Control parameter of porn information detection in image.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImgReviewInfo *PornImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Control parameter of porn information detection in speech.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrReviewInfo *PornAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of porn information detection in text.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *PornOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PornConfigureInfoForUpdate struct {

	// Control parameter of porn information detection in image.
	ImgReviewInfo *PornImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Control parameter of porn information detection in speech.
	AsrReviewInfo *PornAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of porn information detection in text.
	OcrReviewInfo *PornOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PornImgReviewTemplateInfo struct {

	// Switch of a porn information detection in image task. Valid values:
	// <li>ON: Enables a porn information detection in image task;</li>
	// <li>OFF: Disables a porn information detection in image task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter tag for porn information detection in image. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. Valid values:
	// <li>porn: Porn;</li>
	// <li>vulgar: Vulgarity;</li>
	// <li>intimacy: Intimacy;</li>
	// <li>sexy: Sexiness.</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 90 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 0 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornImgReviewTemplateInfoForUpdate struct {

	// Switch of a porn information detection in image task. Valid values:
	// <li>ON: Enables a porn information detection in image task;</li>
	// <li>OFF: Disables a porn information detection in image task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter tag for porn information detection in image. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. Valid values:
	// <li>porn: Porn;</li>
	// <li>vulgar: Vulgarity;</li>
	// <li>intimacy: Intimacy;</li>
	// <li>sexy: Sexiness.</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfo struct {

	// Switch of a porn information detection in text task. Valid values:
	// <li>ON: Enables a porn information detection in text task;</li>
	// <li>OFF: Disables a porn information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfoForUpdate struct {

	// Switch of a porn information detection in text task. Valid values:
	// <li>ON: Enables a porn information detection in text task;</li>
	// <li>OFF: Disables a porn information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type ProcessLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Live stream URL, which must be a live stream file address. RTMP, HLS, and FLV are supported.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Event notification information of a task, which is used to specify the live stream processing result.
	TaskNotifyConfig *LiveStreamTaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

	// Target bucket of a live stream processing output file. This parameter is required if a file will be output.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Target directory of a live stream processing output file, such as `/movie/201909/`. If this parameter is left empty, the `/` directory will be used.
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Type parameter of video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

func (r *ProcessLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "TaskNotifyConfig")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "AiContentReviewTask")
	delete(f, "AiRecognitionTask")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessLiveStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ProcessLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaRequest struct {
	*tchttp.BaseRequest

	// Input information of a file for video processing.
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// Target bucket of a video processing output file. If this parameter is left empty, the storage location in `InputInfo` will be inherited.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Target directory of a video processing output file, such as `/movie/201907/`. If this parameter is left empty, the file will be outputted to the same directory as that in `InputInfo`.
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// Parameter of a video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

	// Task flow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// The ID used for deduplication. If there was a request with the same ID in the last three days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

func (r *ProcessMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "TaskNotifyConfig")
	delete(f, "TasksPriority")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

type ProhibitedAsrReviewTemplateInfoForUpdate struct {

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
	AsrReviewInfo *ProhibitedAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of prohibited information detection in text.
	OcrReviewInfo *ProhibitedOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type ProhibitedConfigureInfoForUpdate struct {

	// Control parameter of prohibited information detection in speech.
	AsrReviewInfo *ProhibitedAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of prohibited information detection in text.
	OcrReviewInfo *ProhibitedOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
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

type ProhibitedOcrReviewTemplateInfoForUpdate struct {

	// Switch of prohibited information detection in text task. Valid values:
	// <li>ON: enables prohibited information detection in text task;</li>
	// <li>OFF: disables prohibited information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type RawImageWatermarkInput struct {

	// Input content of watermark image. JPEG and PNG images are supported.
	ImageContent *MediaInputInfo `json:"ImageContent,omitempty" name:"ImageContent"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.</li>
	// Default value: 10%.
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px.</li>
	// Default value: 0 px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
	Height *string `json:"Height,omitempty" name:"Height"`

	// Repeat type of an animated watermark. Valid values:
	// <li>`once`: no longer appears after watermark playback ends.</li>
	// <li>`repeat_last_frame`: stays on the last frame after watermark playback ends.</li>
	// <li>`repeat` (default): repeats the playback until the video ends.</li>
	RepeatType *string `json:"RepeatType,omitempty" name:"RepeatType"`
}

type RawTranscodeParameter struct {

	// Container. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Whether to remove video data. Valid values:
	// <li>0: retain;</li>
	// <li>1: remove.</li>
	// Default value: 0.
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: retain;</li>
	// <li>1: remove.</li>
	// Default value: 0.
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter. This field is required when `RemoveVideo` is 0.
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter. This field is required when `RemoveAudio` is 0.
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitempty" name:"TEHDConfig"`
}

type RawWatermarkParameter struct {

	// Watermark type. Valid values:
	// <li>image: image watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Origin position, which currently can only be:
	// <li>TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text.</li>
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

	// Image watermark template. This field is required when `Type` is `image` and is invalid when `Type` is `text`.
	ImageTemplate *RawImageWatermarkInput `json:"ImageTemplate,omitempty" name:"ImageTemplate"`
}

type ResetWorkflowRequest struct {
	*tchttp.BaseRequest

	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// Workflow name of up to 128 characters, which must be unique for the same user.
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// Triggering rule bound to a workflow. If an uploaded video hits the rule for the object, the workflow will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// Output configuration of a video processing output file. If this parameter is left empty, the storage location in `Trigger` will be inherited.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Target directory of a video processing output file, such as `/movie/201907/`. If this parameter is left empty, the file will be outputted to the same directory where the source file is located, i.e.; `{inputDir}`.
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// Parameter of a video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Workflow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TaskPriority *int64 `json:"TaskPriority,omitempty" name:"TaskPriority"`

	// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`
}

func (r *ResetWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	delete(f, "WorkflowName")
	delete(f, "Trigger")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "TaskPriority")
	delete(f, "TaskNotifyConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SampleSnapshotTaskInput struct {

	// Sampled screencapturing template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of up to 10 image or text watermarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// Target bucket of a sampled screenshot. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Output path to a generated sampled screenshot, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_sampleSnapshot_{definition}_{number}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// Rule of the `{number}` variable in the sampled screenshot output path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type SampleSnapshotTemplate struct {

	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
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
	// <li>open: Enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: Disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Image format.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Sampled screencapturing type.
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: Stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: Fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: Fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: Fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type SnapshotByTimeOffsetTaskInput struct {

	// ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of screenshot time points in the format of `s` or `%`:
	// <li>If the string ends in `s`, it means that the time point is in seconds; for example, `3.5s` means that the time point is the 3.5th second;</li>
	// <li>If the string ends in `%`, it means that the time point is the specified percentage of the video duration; for example, `10%` means that the time point is 10% of the video duration.</li>
	ExtTimeOffsetSet []*string `json:"ExtTimeOffsetSet,omitempty" name:"ExtTimeOffsetSet"`

	// List of time points of screenshots in <font color=red>seconds</font>.
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitempty" name:"TimeOffsetSet"`

	// List of up to 10 image or text watermarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// Target bucket of a generated time point screenshot file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Output path to a generated time point screenshot, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_snapshotByTimeOffset_{definition}_{number}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// Rule of the `{number}` variable in the time point screenshot output path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type SnapshotByTimeOffsetTemplate struct {

	// Unique ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of a time point screencapturing template.
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
	// <li>open: Enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: Disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Image format.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: Stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: Fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: Fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: Fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
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
}

type TEHDConfig struct {

	// TESHD type. Valid values:
	// <li>TEHD-100: TESHD-100.</li>
	// If this parameter is left empty, TESHD will not be enabled.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Maximum bitrate, which is valid when `Type` is `TESHD`.
	// If this parameter is left empty or 0 is entered, there will be no upper limit for bitrate.
	MaxVideoBitrate *uint64 `json:"MaxVideoBitrate,omitempty" name:"MaxVideoBitrate"`
}

type TEHDConfigForUpdate struct {

	// TESHD type. Valid values:
	// <li>TEHD-100: TESHD-100.</li>
	// If this parameter is left blank, no modification will be made.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Maximum bitrate. If this parameter is left empty, no modification will be made.
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

type TaskNotifyConfig struct {

	// CMQ model. There are two types: `Queue` and `Topic`. Currently, only `Queue` is supported.
	CmqModel *string `json:"CmqModel,omitempty" name:"CmqModel"`

	// CMQ region, such as `sh` and `bj`.
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// This parameter is valid when the model is `Queue`, indicating the name of the CMQ queue for receiving event notifications.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// This parameter is valid when the model is `Topic`, indicating the name of the CMQ topic for receiving event notifications.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Workflow notification method. Valid values: Finish, Change. If this parameter is left empty, `Finish` will be used.
	NotifyMode *string `json:"NotifyMode,omitempty" name:"NotifyMode"`
}

type TaskOutputStorage struct {

	// Storage location type of a video processing output object. Only COS is supported currently.
	Type *string `json:"Type,omitempty" name:"Type"`

	// This parameter is valid and required when `Type` is COS, indicating the location of an output COS object after video processing.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosOutputStorage *CosOutputStorage `json:"CosOutputStorage,omitempty" name:"CosOutputStorage"`
}

type TaskSimpleInfo struct {

	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task type. Valid values:
	// <li> WorkflowTask: Workflow processing task;</li>
	// <li> LiveProcessTask: Live stream processing task.</li>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Creation time of a task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Start time of task execution in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F). If the task has not been started yet, this field will be `0000-00-00T00:00:00Z`.
	BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

	// End time of a task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F). If the task has not been completed yet, this field will be `0000-00-00T00:00:00Z`.
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
}

type TerrorismConfigureInfo struct {

	// Control parameter of a terrorism information detection in image task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImgReviewInfo *TerrorismImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Control parameter of terrorism information detection in text task.
	OcrReviewInfo *TerrorismOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type TerrorismConfigureInfoForUpdate struct {

	// Control parameter of a terrorism information detection in image task.
	ImgReviewInfo *TerrorismImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Control parameter of terrorism information detection in text task.
	OcrReviewInfo *TerrorismOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type TerrorismImgReviewTemplateInfo struct {

	// Switch of a terrorism information detection in image task. Valid values:
	// <li>ON: Enables a terrorism information detection in image task;</li>
	// <li>OFF: Disables a terrorism information detection in image task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter tags for terrorism information detection in images. If a specified tag is detected, the tag is returned. If no filter tag is specified, all detected tags are returned. Valid values:
	// <li>`guns`: weapons and guns</li>
	// <li>`crowd`: crowds</li>
	// <li>`bloody`: bloodiness</li>
	// <li>`police`: police forces</li>
	// <li>`banners`: terrorism flags</li>
	// <li>`militant`: militants</li>
	// <li>`explosion`: explosions and fires</li>
	// <li>`terrorists`: terrorists</li>
	// <li>`scenario`: terrorism images</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 90 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 80 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type TerrorismImgReviewTemplateInfoForUpdate struct {

	// Switch of a terrorism information detection in image task. Valid values:
	// <li>ON: Enables a terrorism information detection in image task;</li>
	// <li>OFF: Disables a terrorism information detection in image task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter tags for terrorism information detection in images. If a specified tag is detected, the tag is returned. If no filter tag is specified, all detected tags are returned. Valid values:
	// <li>`guns`: weapons and guns</li>
	// <li>`crowd`: crowds</li>
	// <li>`bloody`: bloodiness</li>
	// <li>`police`: police forces</li>
	// <li>`banners`: terrorism flags</li>
	// <li>`militant`: militants</li>
	// <li>`explosion`: explosions and fires</li>
	// <li>`terrorists`: terrorists</li>
	// <li>`scenario`: terrorism images</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
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

type TerrorismOcrReviewTemplateInfoForUpdate struct {

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
	// <li>simkai.ttf: Both Chinese and English are supported;</li>
	// <li>arial.ttf: Only English is supported.</li>
	FontType *string `json:"FontType,omitempty" name:"FontType"`

	// Font size in Npx format where N is a numeric value.
	FontSize *string `json:"FontSize,omitempty" name:"FontSize"`

	// Font color in 0xRRGGBB format. Default value: 0xFFFFFF (white).
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// Text transparency. Value range: (0, 1]
	// <li>0: Completely transparent</li>
	// <li>1: Completely opaque</li>
	// Default value: 1.
	FontAlpha *float64 `json:"FontAlpha,omitempty" name:"FontAlpha"`
}

type TextWatermarkTemplateInputForUpdate struct {

	// Font type. Currently, two types are supported:
	// <li>simkai.ttf: Both Chinese and English are supported;</li>
	// <li>arial.ttf: Only English is supported.</li>
	FontType *string `json:"FontType,omitempty" name:"FontType"`

	// Font size in Npx format where N is a numeric value.
	FontSize *string `json:"FontSize,omitempty" name:"FontSize"`

	// Font color in 0xRRGGBB format. Default value: 0xFFFFFF (white).
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// Text transparency. Value range: (0, 1]
	// <li>0: Completely transparent</li>
	// <li>1: Completely opaque</li>
	FontAlpha *float64 `json:"FontAlpha,omitempty" name:"FontAlpha"`
}

type TranscodeTaskInput struct {

	// ID of a video transcoding template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Custom video transcoding parameter, which is valid if `Definition` is 0.
	// This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the transcoding parameter preferably.
	RawParameter *RawTranscodeParameter `json:"RawParameter,omitempty" name:"RawParameter"`

	// Video transcoding custom parameter, which is valid when `Definition` is not 0.
	// When any parameters in this structure are entered, they will be used to override corresponding parameters in templates.
	// This parameter is used in highly customized scenarios. We recommend you only use `Definition` to specify the transcoding parameter.
	// Note: this field may return `null`, indicating that no valid value was found.
	OverrideParameter *OverrideTranscodeParameter `json:"OverrideParameter,omitempty" name:"OverrideParameter"`

	// List of up to 10 image or text watermarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// List of blurs. Up to 10 ones can be supported.
	MosaicSet []*MosaicInput `json:"MosaicSet,omitempty" name:"MosaicSet"`

	// Start time offset of a transcoded video, in seconds.
	// <li>If this parameter is left empty or set to 0, the transcoded video will start at the same time as the original video.</li>
	// <li>If this parameter is set to a positive number (n for example), the transcoded video will start at the nth second of the original video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the transcoded video will start at the nth second before the end of the original video.</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a transcoded video, in seconds.
	// <li>If this parameter is left empty or set to 0, the transcoded video will end at the same time as the original video.</li>
	// <li>If this parameter is set to a positive number (n for example), the transcoded video will end at the nth second of the original video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the transcoded video will end at the nth second before the end of the original video.</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Target bucket of an output file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Path to a primary output file, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_transcode_{definition}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// Path to an output file part (the path to ts during transcoding to HLS), which can only be a relative path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_transcode_{definition}_{number}.{format}`.
	SegmentObjectName *string `json:"SegmentObjectName,omitempty" name:"SegmentObjectName"`

	// Rule of the `{number}` variable in the output path after transcoding.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type TranscodeTemplate struct {

	// Unique ID of a transcoding template.
	Definition *string `json:"Definition,omitempty" name:"Definition"`

	// Container format. Valid values: mp4, flv, hls, mp3, flac, ogg.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Name of a transcoding template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Whether to remove video data. Valid values:
	// <li>0: Retain;</li>
	// <li>1: Remove.</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: Retain;</li>
	// <li>1: Remove.</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter. This field is valid only when `RemoveVideo` is 0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter. This field is valid only when `RemoveAudio` is 0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter. To enable it, please contact your Tencent Cloud sales rep.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitempty" name:"TEHDConfig"`

	// Container format filter. Valid values:
	// <li>Video: Video container format that can contain both video stream and audio stream;</li>
	// <li>PureAudio: Audio container format that can contain only audio stream.</li>
	ContainerType *string `json:"ContainerType,omitempty" name:"ContainerType"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type UrlInputInfo struct {

	// URL of a video.
	Url *string `json:"Url,omitempty" name:"Url"`
}

type UserDefineAsrTextReviewTemplateInfo struct {

	// Switch of a custom speech audit task. Valid values:
	// <li>ON: Enables a custom speech audit task;</li>
	// <li>OFF: Disables a custom speech audit task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom speech filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom speech keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineAsrTextReviewTemplateInfoForUpdate struct {

	// Switch of a custom speech audit task. Valid values:
	// <li>ON: Enables a custom speech audit task;</li>
	// <li>OFF: Disables a custom speech audit task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom speech filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom speech keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineConfigureInfo struct {

	// Control parameter of custom figure audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FaceReviewInfo *UserDefineFaceReviewTemplateInfo `json:"FaceReviewInfo,omitempty" name:"FaceReviewInfo"`

	// Control parameter of custom speech audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of custom text audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type UserDefineConfigureInfoForUpdate struct {

	// Control parameter of custom figure audit.
	FaceReviewInfo *UserDefineFaceReviewTemplateInfoForUpdate `json:"FaceReviewInfo,omitempty" name:"FaceReviewInfo"`

	// Control parameter of custom speech audit.
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Control parameter of custom text audit.
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type UserDefineFaceReviewTemplateInfo struct {

	// Switch of a custom figure audit task. Valid values:
	// <li>ON: Enables a custom figure audit task;</li>
	// <li>OFF: Disables a custom figure audit task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom figure filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for the custom figure library.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 97 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 95 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineFaceReviewTemplateInfoForUpdate struct {

	// Switch of a custom figure audit task. Valid values:
	// <li>ON: Enables a custom figure audit task;</li>
	// <li>OFF: Disables a custom figure audit task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom figure filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for the custom figure library.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfo struct {

	// Switch of a custom text audit task. Valid values:
	// <li>ON: Enables a custom text audit task;</li>
	// <li>OFF: Disables a custom text audit task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom text filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom text keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfoForUpdate struct {

	// Switch of a custom text audit task. Valid values:
	// <li>ON: Enables a custom text audit task;</li>
	// <li>OFF: Disables a custom text audit task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom text filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom text keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet *string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type VideoTemplateInfo struct {

	// Video stream codec. Valid values:
	// <li>libx264: H.264</li>
	// <li>libx265: H.265</li>
	// Currently, a resolution within 640*480p must be specified for H.265.
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Video frame rate in Hz. Value range: [0, 100].
	// If the value is 0, the frame rate will be the same as that of the source video.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Bitrate of a video stream in Kbps. Value range: 0 and [128, 35,000].
	// If the value is 0, the bitrate of the video will be the same as that of the source video.
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Resolution adaption. Valid values:
	// <li>open: Enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: Disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
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

	// Frame interval between I keyframes. Value range: 0 and [1,100000].
	// If this parameter is 0 or left empty, the system will automatically set the GOP length.
	Gop *uint64 `json:"Gop,omitempty" name:"Gop"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// The control factor of video constant bitrate. Value range: [1, 51]
	// If this parameter is specified, CRF (a bitrate control method) will be used for transcoding. (Video bitrate will no longer take effect.)
	// It is not recommended to specify this parameter if there are no special requirements.
	Vcrf *uint64 `json:"Vcrf,omitempty" name:"Vcrf"`
}

type VideoTemplateInfoForUpdate struct {

	// Video stream codec. Valid values:
	// <li>libx264: H.264</li>
	// <li>libx265: H.265</li>
	// Currently, a resolution within 640*480p must be specified for H.265.
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Video frame rate in Hz. Value range: [0, 100].
	// If the value is 0, the frame rate will be the same as that of the source video.
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// Bitrate of a video stream in Kbps. Value range: 0 and [128, 35,000].
	// If the value is 0, the bitrate of the video will be the same as that of the source video.
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Resolution adaption. Valid values:
	// <li>open: Enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: Disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// Maximum value of the width (or long side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// Frame interval between I keyframes. Value range: 0 and [1,100000]. If this parameter is 0, the system will automatically set the GOP length.
	Gop *uint64 `json:"Gop,omitempty" name:"Gop"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// The control factor of video constant bitrate. Value range: [0, 51]. This parameter will be disabled if you enter `0`.
	// It is not recommended to specify this parameter if there are no special requirements.
	Vcrf *uint64 `json:"Vcrf,omitempty" name:"Vcrf"`
}

type WatermarkInput struct {

	// ID of a watermarking template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Custom watermark parameter, which is valid if `Definition` is 0.
	// This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the watermark parameter preferably.
	// Custom watermark parameter is not available for screenshot.
	RawParameter *RawWatermarkParameter `json:"RawParameter,omitempty" name:"RawParameter"`

	// Text content of up to 100 characters. This field is required only when the watermark type is text.
	// Text watermark is not available for screenshot.
	TextContent *string `json:"TextContent,omitempty" name:"TextContent"`

	// SVG content of up to 2,000,000 characters. This field is required only when the watermark type is `SVG`.
	// SVG watermark is not available for screenshot.
	SvgContent *string `json:"SvgContent,omitempty" name:"SvgContent"`

	// Start time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame.
	// <li>If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame;</li>
	// <li>If this value is greater than 0 (e.g., n), the watermark will appear at second n after the first video frame;</li>
	// <li>If this value is smaller than 0 (e.g., -n), the watermark will appear at second n before the last video frame.</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a watermark in seconds.
	// <li>If this parameter is left empty or 0 is entered, the watermark will exist till the last video frame;</li>
	// <li>If this value is greater than 0 (e.g., n), the watermark will exist till second n;</li>
	// <li>If this value is smaller than 0 (e.g., -n), the watermark will exist till second n before the last video frame.</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type WatermarkTemplate struct {

	// Unique ID of a watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Watermark type. Valid values:
	// <li>image: Image watermark;</li>
	// <li>text: Text watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Name of a watermarking template.
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
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImageTemplate *ImageWatermarkTemplate `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is valid only when `Type` is `text`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is valid when `Type` is `svg`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitempty" name:"SvgTemplate"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Origin position. Valid values:
	// <li>topLeft: The origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>topRight: The origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>bottomLeft: The origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>bottomRight: The origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`
}

type WorkflowInfo struct {

	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// Workflow status. Valid values:
	// <li>Enabled: Enabled,</li>
	// <li>Disabled: Disabled.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Input rule bound to a workflow. If an uploaded video hits the rule for the object, the workflow will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// Target storage of a video processing output file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// Parameter of a video processing task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitempty" name:"TaskNotifyConfig"`

	// Task flow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TaskPriority *int64 `json:"TaskPriority,omitempty" name:"TaskPriority"`

	// Target directory of a video processing output file, such as `/movie/201907/`.
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// Creation time of a workflow in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of a workflow in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type WorkflowTask struct {

	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task flow status. Valid values:
	// <li>PROCESSING: Processing;</li>
	// <li>FINISH: Completed.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Disused. Please use `ErrCode` of each specific task.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Disused. Please use `Message` of each specific task.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Information of a target file of video processing.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// Metadata of a source video.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// Execution status and result of a video processing task.
	MediaProcessResultSet []*MediaProcessTaskResult `json:"MediaProcessResultSet,omitempty" name:"MediaProcessResultSet"`

	// Execution status and result of a video content audit task.
	AiContentReviewResultSet []*AiContentReviewResult `json:"AiContentReviewResultSet,omitempty" name:"AiContentReviewResultSet"`

	// Execution status and result of video content analysis task.
	AiAnalysisResultSet []*AiAnalysisResult `json:"AiAnalysisResultSet,omitempty" name:"AiAnalysisResultSet"`

	// Execution status and result of a video content recognition task.
	AiRecognitionResultSet []*AiRecognitionResult `json:"AiRecognitionResultSet,omitempty" name:"AiRecognitionResultSet"`
}

type WorkflowTrigger struct {

	// Trigger type. Only `CosFileUpload` is supported currently.
	Type *string `json:"Type,omitempty" name:"Type"`

	// This parameter is required and valid when `Type` is `CosFileUpload`, indicating the COS trigger rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosFileUploadTrigger *CosFileUploadTrigger `json:"CosFileUploadTrigger,omitempty" name:"CosFileUploadTrigger"`
}
