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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
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

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AccelerateAreaInfo struct {
	// Acceleration region. Valid values:
	// <li>Chinese Mainland</li>
	// <li>Outside Chinese Mainland</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// Reason why acceleration is disabled by Tencent Cloud. Valid values:
	// <li>ForLegalReasons: legal reasons</li>
	// <li>ForOverdueBills: overdue payment</li>
	TencentDisableReason *string `json:"TencentDisableReason,omitempty" name:"TencentDisableReason"`

	// CNAME of the acceleration domain name
	TencentEdgeDomain *string `json:"TencentEdgeDomain,omitempty" name:"TencentEdgeDomain"`
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

	// File size (bytes)
	// <li>If the file is an HLS file, the value of this parameter is the sum of the size of the M3U8 and TS files.</li>
	// <li>If the file is a DASH file, the value of this parameter is the sum of the size of the MPD and segment files.</li>
	// <li><font color=red>Note</font>: For adaptive bitrate streaming files generated before 2022-01-10T16:00:00Z, the value of this parameter is `0`.</li>
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

type AdaptiveDynamicStreamingTaskInput struct {
	// Adaptive bitrate streaming template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of up to 10 image or text watermarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// Digital watermark.
	TraceWatermark *TraceWatermarkInput `json:"TraceWatermark,omitempty" name:"TraceWatermark"`

	// List of subtitle IDs (maximum: 16)
	SubtitleSet []*string `json:"SubtitleSet,omitempty" name:"SubtitleSet"`
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

	// Adaptive bitstream format. Valid value:
	// <li>HLS.</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// The DRM type. Valid values:
	// <li>SimpleAES</li>
	// <li>Widevine</li>
	// <li>FairPlay</li>
	// If this parameter is an empty string, it indicates that the video is not protected by DRM.
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// Parameter information of input stream for adaptive bitrate streaming. Up to 10 streams can be input.
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

	// Whether to remove a video stream. Valid values:
	// <li>0: no</li>
	// <li>1: yes</li>
	RemoveVideo *uint64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// TESHD transcoding parameters
	// Note: This field may return `null`, indicating that no valid value was found.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitempty" name:"TEHDConfig"`
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
	// List of intelligently generated video categories
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `ClassificationSetFileUrl`.
	ClassificationSet []*MediaAiAnalysisClassificationItem `json:"ClassificationSet,omitempty" name:"ClassificationSet"`

	// URL to the file for intelligently generated video categories. The file is in JSON format and has the same data structure as `ClassificationSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `ClassificationSetFileUrlExpireTime`.
	ClassificationSetFileUrl *string `json:"ClassificationSetFileUrl,omitempty" name:"ClassificationSetFileUrl"`

	// Expiration time of the URL to the file for intelligently generated video categories, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	ClassificationSetFileUrlExpireTime *string `json:"ClassificationSetFileUrlExpireTime,omitempty" name:"ClassificationSetFileUrlExpireTime"`
}

type AiAnalysisTaskClassificationResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	// List of intelligently generated thumbnails
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `CoverSetFileUrl`.
	CoverSet []*MediaAiAnalysisCoverItem `json:"CoverSet,omitempty" name:"CoverSet"`

	// URL to the file for intelligently generated thumbnails. The file is in JSON format and has the same data structure as `CoverSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `CoverSetFileUrlExpireTime`.
	CoverSetFileUrl *string `json:"CoverSetFileUrl,omitempty" name:"CoverSetFileUrl"`

	// Expiration time of the URL to the file for intelligently generated thumbnails, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	CoverSetFileUrlExpireTime *string `json:"CoverSetFileUrlExpireTime,omitempty" name:"CoverSetFileUrlExpireTime"`
}

type AiAnalysisTaskCoverResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	// List of frame-specific video tags
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaAiAnalysisFrameTagSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for frame-specific video tags. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for frame-specific video tags, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiAnalysisTaskFrameTagResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	// List of intelligently generated highlights
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `HighlightSetFileUrl`.
	HighlightSet []*MediaAiAnalysisHighlightItem `json:"HighlightSet,omitempty" name:"HighlightSet"`

	// URL to the file for intelligently generated highlights. The file is in JSON format and has the same data structure as `HighlightSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `HighlightSetFileUrlExpireTime`.
	HighlightSetFileUrl *string `json:"HighlightSetFileUrl,omitempty" name:"HighlightSetFileUrl"`

	// Expiration time of the URL to the file for intelligently generated highlights, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	HighlightSetFileUrlExpireTime *string `json:"HighlightSetFileUrlExpireTime,omitempty" name:"HighlightSetFileUrlExpireTime"`
}

type AiAnalysisTaskHighlightResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	// List of intelligently generated video tags
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `TagSetFileUrl`.
	TagSet []*MediaAiAnalysisTagItem `json:"TagSet,omitempty" name:"TagSet"`

	// URL to the file for intelligently generated video tags. The file is in JSON format and has the same data structure as `TagSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `TagSetFileUrlExpireTime`.
	TagSetFileUrl *string `json:"TagSetFileUrl,omitempty" name:"TagSetFileUrl"`

	// Expiration time of the URL to the file for intelligently generated video tags, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	TagSetFileUrlExpireTime *string `json:"TagSetFileUrlExpireTime,omitempty" name:"TagSetFileUrlExpireTime"`
}

type AiAnalysisTaskTagResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	// <li>`Porn`: porn information recognition in images</li>
	// <li>`Terrorism`: terrorism information recognition in images</li>
	// <li>`Political`: politically sensitive information recognition in images</li>
	// <li>`Porn.Asr`: ASR-based porn information recognition in speech</li>
	// <li>`Porn.Ocr`: OCR-based porn information recognition in text</li>
	// <li>`Political.Asr`: ASR-based politically sensitive information recognition in speech</li>
	// <li>`Political.Ocr`: OCR-based politically sensitive information recognition in text</li>
	// <li>`Terrorism.Ocr`: OCR-based terrorism information recognition in text</li>
	// <li>`Prohibited.Asr`: ASR-based prohibited information recognition in speech</li>
	// <li>`Prohibited.Ocr`: OCR-based prohibited information recognition in text</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Result for intelligent recognition of pornographic content in images. This parameter is valid when `Type` is `Porn`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	PornTask *AiReviewTaskPornResult `json:"PornTask,omitempty" name:"PornTask"`

	// Result for intelligent recognition of terrorism content in images. This parameter is valid when `Type` is `Terrorism`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	TerrorismTask *AiReviewTaskTerrorismResult `json:"TerrorismTask,omitempty" name:"TerrorismTask"`

	// Result for intelligent recognition of politically sensitive content in images. This parameter is valid when `Type` is `Political`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	PoliticalTask *AiReviewTaskPoliticalResult `json:"PoliticalTask,omitempty" name:"PoliticalTask"`

	// Result for ASR-based recognition of pornographic content. This parameter is valid when `Type` is `Porn.Asr`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	PornAsrTask *AiReviewTaskPornAsrResult `json:"PornAsrTask,omitempty" name:"PornAsrTask"`

	// Result for OCR-based recognition of pornographic content. This parameter is valid when `Type` is `Porn.Ocr`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	PornOcrTask *AiReviewTaskPornOcrResult `json:"PornOcrTask,omitempty" name:"PornOcrTask"`

	// Result for ASR-based recognition of politically sensitive content. This parameter is valid when `Type` is `Political.Asr`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	PoliticalAsrTask *AiReviewTaskPoliticalAsrResult `json:"PoliticalAsrTask,omitempty" name:"PoliticalAsrTask"`

	// Result for OCR-based recognition of politically sensitive content. This parameter is valid when `Type` is `Political.Ocr`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	PoliticalOcrTask *AiReviewTaskPoliticalOcrResult `json:"PoliticalOcrTask,omitempty" name:"PoliticalOcrTask"`

	// Result for OCR-based recognition of terrorism content. This parameter is valid when `Type` is `Terrorism.Ocr`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	TerrorismOcrTask *AiReviewTaskTerrorismOcrResult `json:"TerrorismOcrTask,omitempty" name:"TerrorismOcrTask"`

	// Result for OCR-based recognition of banned content. This parameter is valid when `Type` is `Prohibited.Ocr`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	ProhibitedOcrTask *AiReviewTaskProhibitedOcrResult `json:"ProhibitedOcrTask,omitempty" name:"ProhibitedOcrTask"`

	// Result for ASR-based recognition of banned content. This parameter is valid when `Type` is `Prohibited.Asr`.
	// Note: This field may return `null`, indicating that no valid value can be found.
	ProhibitedAsrTask *AiReviewTaskProhibitedAsrResult `json:"ProhibitedAsrTask,omitempty" name:"ProhibitedAsrTask"`
}

type AiContentReviewTaskInput struct {
	// Intelligent recognition template ID
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	// List of full-text speech recognition segments
	// <font color=red>Note</font>: this list displays up to the first 100 results. You can get all the results from the file whose URL is `SegmentSetFileUrl`.
	SegmentSet []*AiRecognitionTaskAsrFullTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file of the list for full-text speech recognition segments. The file format is JSON, and the data structure is the same as `SegmentSet`. The file will be deleted upon the expiration time `SegmentSetFileUrlExpireTime`, instead of being stored permanently.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file of the list for full-text speech recognition segments, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`

	// Subtitles file URL.
	SubtitleUrl *string `json:"SubtitleUrl,omitempty" name:"SubtitleUrl"`
}

type AiRecognitionTaskAsrFullTextSegmentItem struct {
	// Confidence of recognized segment. Value range: 0-100.
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	SegmentSet []*AiRecognitionTaskAsrWordsSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskAsrWordsResultOutput struct {
	// Speech keyword recognition result set
	// <font color=red>Note</font>: this list displays up to the first 100 results. You can get all the results from the file whose URL is `SegmentSetFileUrl`.
	ResultSet []*AiRecognitionTaskAsrWordsResultItem `json:"ResultSet,omitempty" name:"ResultSet"`

	// URL to the file of the speech keyword recognition result set. The file format is JSON, and the data structure is the same as `SegmentSet`. The file will be deleted upon the expiration time `SegmentSetFileUrlExpireTime`, instead of being stored permanently.
	ResultSetFileUrl *string `json:"ResultSetFileUrl,omitempty" name:"ResultSetFileUrl"`

	// Expiration time of the URL to the file of the speech keyword recognition result set, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732)
	ResultSetFileUrlExpireTime *string `json:"ResultSetFileUrlExpireTime,omitempty" name:"ResultSetFileUrlExpireTime"`
}

type AiRecognitionTaskAsrWordsSegmentItem struct {
	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of recognized segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type AiRecognitionTaskFaceResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	SegmentSet []*AiRecognitionTaskFaceSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskFaceResultOutput struct {
	// Intelligent face recognition result set
	// <font color=red>Note</font>: this list displays up to the first 100 results. You can get all the results from the file whose URL is `SegmentSetFileUrl`.
	ResultSet []*AiRecognitionTaskFaceResultItem `json:"ResultSet,omitempty" name:"ResultSet"`

	// URL to the file of the intelligent face recognition result set. The file format is JSON, and the data structure is the same as `SegmentSet`. The file will be deleted upon the expiration time `SegmentSetFileUrlExpireTime`, instead of being stored permanently.
	ResultSetFileUrl *string `json:"ResultSetFileUrl,omitempty" name:"ResultSetFileUrl"`

	// Expiration time of the URL to the file of the intelligent face recognition result set, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732)
	ResultSetFileUrlExpireTime *string `json:"ResultSetFileUrlExpireTime,omitempty" name:"ResultSetFileUrlExpireTime"`
}

type AiRecognitionTaskFaceSegmentItem struct {
	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of recognized segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskHeadTailResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	// Confidence of recognized opening credits. Value range: 0-100.
	HeadConfidence *float64 `json:"HeadConfidence,omitempty" name:"HeadConfidence"`

	// End time point of video opening credits in seconds.
	HeadTimeOffset *float64 `json:"HeadTimeOffset,omitempty" name:"HeadTimeOffset"`

	// Confidence of recognized closing credits. Value range: 0-100.
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	SegmentSet []*AiRecognitionTaskObjectSeqmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskObjectResultOutput struct {
	// Intelligent object recognition result set
	// <font color=red>Note</font>: this list displays up to the first 100 results. You can get all the results from the file whose URL is `SegmentSetFileUrl`.
	ResultSet []*AiRecognitionTaskObjectResultItem `json:"ResultSet,omitempty" name:"ResultSet"`

	// URL to the file of the object recognition result set. The file format is JSON, and the data structure is the same as `SegmentSet`. The file will be deleted upon the expiration time `SegmentSetFileUrlExpireTime`, instead of being stored permanently.
	ResultSetFileUrl *string `json:"ResultSetFileUrl,omitempty" name:"ResultSetFileUrl"`

	// Expiration time of the URL to the object recognition result set, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732)
	ResultSetFileUrlExpireTime *string `json:"ResultSetFileUrlExpireTime,omitempty" name:"ResultSetFileUrlExpireTime"`
}

type AiRecognitionTaskObjectSeqmentItem struct {
	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of recognized segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskOcrFullTextResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	// Full-text recognition result set
	// <font color=red>Note</font>: this list displays up to the first 100 results. You can get all the results from the file whose URL is `SegmentSetFileUrl`.
	SegmentSet []*AiRecognitionTaskOcrFullTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file of the full-text recognition result set. The file format is JSON, and the data structure is the same as `SegmentSet`. The file will be deleted upon the expiration time `SegmentSetFileUrlExpireTime`, instead of being stored permanently.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file of the full-text recognition result set, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiRecognitionTaskOcrFullTextSegmentItem struct {
	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Recognition segment result set.
	TextSet []*AiRecognitionTaskOcrFullTextSegmentTextItem `json:"TextSet,omitempty" name:"TextSet"`
}

type AiRecognitionTaskOcrFullTextSegmentTextItem struct {
	// Confidence of recognized segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`

	// Recognized text.
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AiRecognitionTaskOcrWordsResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	SegmentSet []*AiRecognitionTaskOcrWordsSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskOcrWordsResultOutput struct {
	// Text keyword recognition result set
	// <font color=red>Note</font>: this list displays up to the first 100 results. You can get all the results from the file whose URL is `SegmentSetFileUrl`.
	ResultSet []*AiRecognitionTaskOcrWordsResultItem `json:"ResultSet,omitempty" name:"ResultSet"`

	// URL to the file of the text keyword recognition result set. The file format is JSON, and the data structure is the same as `SegmentSet`. The file will be deleted upon the expiration time `SegmentSetFileUrlExpireTime`, instead of being stored permanently.
	ResultSetFileUrl *string `json:"ResultSetFileUrl,omitempty" name:"ResultSetFileUrl"`

	// Expiration time of the URL to the file of the text keyword recognition result set, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732)
	ResultSetFileUrlExpireTime *string `json:"ResultSetFileUrlExpireTime,omitempty" name:"ResultSetFileUrlExpireTime"`
}

type AiRecognitionTaskOcrWordsSegmentItem struct {
	// Start time offset of recognized segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence of recognized segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskSegmentResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	// List of split video segments
	// <font color=red>Note</font>: this list displays up to the first 100 results. You can get all the results from the file whose URL is `SegmentSetFileUrl`.
	SegmentSet []*AiRecognitionTaskSegmentSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file of the list for split video segments. The file format is JSON, and the data structure is the same as `SegmentSet`. The file will be deleted upon the expiration time `SegmentSetFileUrlExpireTime`, instead of being stored permanently.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file of the list for split video segments, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiRecognitionTaskSegmentSegmentItem struct {
	// File ID, which is valid only when a VOD file is processed and the subsegments generated through segmentation are also VOD files.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Split video segment URL.
	SegmentUrl *string `json:"SegmentUrl,omitempty" name:"SegmentUrl"`

	// Confidence of split segment. Value range: 0-100.
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
	// ID of the template for recognition of politically sensitive content
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalAsrTaskOutput struct {
	// Confidence score for the ASR-detected politically sensitive content. Value range: 0-100
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Processing suggestion for the ASR-detected politically sensitive content. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain ASR-detected politically sensitive content
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain ASR-detected politically sensitive content. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for video segments that contain ASR-detected politically sensitive content, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPoliticalOcrTaskInput struct {
	// ID of the template for recognition of politically sensitive content
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalOcrTaskOutput struct {
	// Confidence score for the OCR-detected politically sensitive content. Value range: 0-100
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Processing suggestion for the OCR-detected politically sensitive content. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain OCR-detected politically sensitive content
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain OCR-detected politically sensitive content. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for video segments that contain OCR-detected politically sensitive content, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPoliticalTaskInput struct {
	// ID of the template for recognition of politically sensitive content
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalTaskOutput struct {
	// Confidence score for the detected politically sensitive content. Value range: 0-100
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Processing suggestion for the detected politically sensitive content
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Labels for the detected politically sensitive content. The relationship between the values of this parameter and those of the `LabelSet` parameter in [PoliticalImgReviewTemplateInfo](https://intl.cloud.tencent.com/document/api/266/31773?from_cn_redirect=1#PoliticalImgReviewTemplateInfo) is as follows:
	// violation_photo:
	// <li>`violation_photo`: banned images</li>
	// Other values (politician/entertainment/sport/entrepreneur/scholar/celebrity/military):
	// <li>`politician`: politically sensitive people</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// List of video segments that contain detected politically sensitive content
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewPoliticalSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain detected politically sensitive content. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for video segments that contain politically sensitive content, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPornAsrTaskInput struct {
	// ID of the template for recognition of pornographic content
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornAsrTaskOutput struct {
	// Confidence score for the ASR-detected pornographic content
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Processing suggestion for the ASR-detected pornographic content
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain ASR-detected pornographic content
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain ASR-detected pornographic content. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for video segments that contain ASR-detected pornographic content, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPornOcrTaskInput struct {
	// ID of the template for recognition of pornographic content
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornOcrTaskOutput struct {
	// Confidence score for the OCR-detected pornographic content
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Processing suggestion for the OCR-detected pornographic content
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain OCR-detected pornographic content
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain OCR-detected pornographic content. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for video segments that contain OCR-detected pornographic content, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewPornTaskInput struct {
	// ID of the template for recognition of pornographic content
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornTaskOutput struct {
	// Confidence score for the detected pornographic content. Value range: 0-100
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Processing suggestion for the detected pornographic content. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Labels for the detected pornographic content. Valid values:
	// <li>porn</li>
	// <li>sexy</li>
	// <li>vulgar</li>
	// <li>intimacy</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// List of video segments that contain detected pornographic content
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain detected pornographic content. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for video segments that contain detected pornographic content, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
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

	// List of video segments that contain ASR-detected prohibited information
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain ASR-detected prohibited information. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for video segments that contain ASR-detected prohibited information, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
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

	// List of video segments that contain OCR-detected prohibited information
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain OCR-detected prohibited information. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL for video segments that contain OCR-detected prohibited information, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewTaskPoliticalAsrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for ASR-based recognition of politically sensitive content
	Input *AiReviewPoliticalAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for ASR-based recognition of politically sensitive content
	Output *AiReviewPoliticalAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalOcrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for OCR-based recognition of politically sensitive content
	Input *AiReviewPoliticalOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for OCR-based recognition of politically sensitive content
	// Note: This field may return `null`, indicating that no valid value can be found.
	Output *AiReviewPoliticalOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for intelligent recognition of politically sensitive content
	Input *AiReviewPoliticalTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for intelligent recognition of politically sensitive content
	// Note: This field may return `null`, indicating that no valid value can be found.
	Output *AiReviewPoliticalTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornAsrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for ASR-based recognition of pornographic content
	Input *AiReviewPornAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for ASR-based recognition of pornographic content
	// Note: This field may return `null`, indicating that no valid value can be found.
	Output *AiReviewPornAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornOcrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for OCR-based recognition of pornographic content
	Input *AiReviewPornOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for OCR-based recognition of pornographic content
	// Note: This field may return `null`, indicating that no valid value can be found.
	Output *AiReviewPornOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for intelligent recognition of pornographic content
	Input *AiReviewPornTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for intelligent recognition of pornographic content
	// Note: This field may return `null`, indicating that no valid value can be found.
	Output *AiReviewPornTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskProhibitedAsrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for ASR-based recognition of banned content
	Input *AiReviewProhibitedAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for ASR-based recognition of banned content
	// Note: This field may return `null`, indicating that no valid value can be found.
	Output *AiReviewProhibitedAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskProhibitedOcrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for OCR-based recognition of banned content
	Input *AiReviewProhibitedOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for OCR-based recognition of banned content
	// Note: This field may return `null`, indicating that no valid value can be found.
	Output *AiReviewProhibitedOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismOcrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for OCR-based recognition of terrorism content
	Input *AiReviewTerrorismOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for OCR-based recognition of terrorism content
	// Note: This field may return `null`, indicating that no valid value can be found.
	Output *AiReviewTerrorismOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input for intelligent recognition of terrorism content
	Input *AiReviewTerrorismTaskInput `json:"Input,omitempty" name:"Input"`

	// Output for intelligent recognition of terrorism content
	// Note: This field may return `null`, indicating that no valid value can be found.
	Output *AiReviewTerrorismTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTerrorismOcrTaskInput struct {
	// ID of the template for recognition of terrorism content
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewTerrorismOcrTaskOutput struct {
	// Confidence score for the OCR-detected terrorism content. Value range: 0-100
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Processing suggestion for the OCR-detected terrorism content
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of video segments that contain OCR-detected terrorism content
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain OCR-detected terrorism content. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for video segments that contain OCR-detected terrorism content, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
}

type AiReviewTerrorismTaskInput struct {
	// ID of the template for recognition of terrorism content
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

	// Tag of the detected terrorism information in a video. Valid values:
	// <li>`guns`: weapons and guns</li>
	// <li>`crowd`: crowds</li>
	// <li>`police`: police forces</li>
	// <li>`bloody`: bloody images</li>
	// <li>`banners`: terrorism flags</li>
	// <li>`militant`: militants</li>
	// <li>`explosion`: explosions and fires</li>
	// <li>`terrorists`: terrorists</li>
	// <li>`scenario`: terrorism images</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// List of video segments that contain terrorism information
	// <font color=red>Note</font>: This list displays the first 100 results at most. You can get all the results from the file at the URL specified by `SegmentSetFileUrl`.
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`

	// URL to the file for video segments that contain terrorism information. The file is in JSON format and has the same data structure as `SegmentSet`. Instead of being saved permanently, the file is deleted upon the expiration time specified by `SegmentSetFileUrlExpireTime`.
	SegmentSetFileUrl *string `json:"SegmentSetFileUrl,omitempty" name:"SegmentSetFileUrl"`

	// Expiration time of the URL to the file for video segments that contain terrorism information, in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)
	SegmentSetFileUrlExpireTime *string `json:"SegmentSetFileUrlExpireTime,omitempty" name:"SegmentSetFileUrlExpireTime"`
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

	// Face ID set, which is required if `Type` is `delete`.
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds"`

	// String set generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) the face image.
	// <li>This field is required if `Type` is `add` or `reset`;</li>
	// <li>Array length limit: 5 images.</li>
	// Note: the image must be a relatively clear full-face photo of a figure in at least 200 * 200 px.
	FaceContents []*string `json:"FaceContents,omitempty" name:"FaceContents"`
}

type AiSampleFailFaceInfo struct {
	// It corresponds to incorrect image subscripts in the `FaceContents` input parameter, starting from 0.
	Index *uint64 `json:"Index,omitempty" name:"Index"`

	// Error code. Valid values:
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type AiSamplePerson struct {
	// Figure ID.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Figure name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Figure description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Face information.
	FaceInfoSet []*AiSampleFaceInfo `json:"FaceInfoSet,omitempty" name:"FaceInfoSet"`

	// Figure tag.
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet"`

	// Use case.
	UsageSet []*string `json:"UsageSet,omitempty" name:"UsageSet"`

	// Creation time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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

	// Creation time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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
	// Animated image generating template ID
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// Start time offset of an animated image in the video, in seconds.
	// <li>If this parameter is left empty or set to 0, the animated image will start at the same time as the video.</li>
	// <li>If this parameter is set to a positive number (n for example), the animated image will start at the nth second of the video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the animated image will start at the nth second before the end of the video.</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of an animated image in the video, in seconds.
	// <li>If this parameter is left empty or set to 0, the animated image will end at the same time as the video.</li>
	// <li>If this parameter is set to a positive number (n for example), the animated image will end at the nth second of the video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the animated image will end at the nth second before the end of the video.</li>
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

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ApplyUploadRequestParams struct {
	// Media type. For the detailed valid values, please see [Upload Overview](https://intl.cloud.tencent.com/document/product/266/9760?from_cn_redirect=1#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B).
	MediaType *string `json:"MediaType,omitempty" name:"MediaType"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Media name.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// Cover type. For the detailed valid values, please see [Upload Overview](https://intl.cloud.tencent.com/document/product/266/9760?from_cn_redirect=1#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B).
	CoverType *string `json:"CoverType,omitempty" name:"CoverType"`

	// Subsequent task operation on a media file, i.e., after a media file is uploaded, task flow operations will be initiated automatically. This parameter value is a task flow template name. VOD supports [creating task flow templates](https://intl.cloud.tencent.com/document/product/266/33819?from_cn_redirect=1) and naming the templates.
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Expiration time of a media file in ISO 8601 format. For more information, please see [Notes on ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Specifies upload region. This is only applicable to users that have special requirements for the upload region.
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](https://intl.cloud.tencent.com/document/product/266/7812?from_cn_redirect=1) API.
	// <li>Default value: 0, which means "Other".</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Source context, which is used to pass through the user request information. The [upload callback](https://intl.cloud.tencent.com/document/product/266/7830?from_cn_redirect=1) API will return the value of this field. It can contain up to 250 characters.
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`

	// Session context, which is used to pass through the user request information. If the `Procedure` parameter is specified, the [task flow status change callback](https://intl.cloud.tencent.com/document/product/266/9636?from_cn_redirect=1) API will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Reserved parameter for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

type ApplyUploadRequest struct {
	*tchttp.BaseRequest
	
	// Media type. For the detailed valid values, please see [Upload Overview](https://intl.cloud.tencent.com/document/product/266/9760?from_cn_redirect=1#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B).
	MediaType *string `json:"MediaType,omitempty" name:"MediaType"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Media name.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// Cover type. For the detailed valid values, please see [Upload Overview](https://intl.cloud.tencent.com/document/product/266/9760?from_cn_redirect=1#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B).
	CoverType *string `json:"CoverType,omitempty" name:"CoverType"`

	// Subsequent task operation on a media file, i.e., after a media file is uploaded, task flow operations will be initiated automatically. This parameter value is a task flow template name. VOD supports [creating task flow templates](https://intl.cloud.tencent.com/document/product/266/33819?from_cn_redirect=1) and naming the templates.
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Expiration time of a media file in ISO 8601 format. For more information, please see [Notes on ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Specifies upload region. This is only applicable to users that have special requirements for the upload region.
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](https://intl.cloud.tencent.com/document/product/266/7812?from_cn_redirect=1) API.
	// <li>Default value: 0, which means "Other".</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Source context, which is used to pass through the user request information. The [upload callback](https://intl.cloud.tencent.com/document/product/266/7830?from_cn_redirect=1) API will return the value of this field. It can contain up to 250 characters.
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`

	// Session context, which is used to pass through the user request information. If the `Procedure` parameter is specified, the [task flow status change callback](https://intl.cloud.tencent.com/document/product/266/9636?from_cn_redirect=1) API will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Reserved parameter for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

func (r *ApplyUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaType")
	delete(f, "SubAppId")
	delete(f, "MediaName")
	delete(f, "CoverType")
	delete(f, "Procedure")
	delete(f, "ExpireTime")
	delete(f, "StorageRegion")
	delete(f, "ClassId")
	delete(f, "SourceContext")
	delete(f, "SessionContext")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyUploadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyUploadResponseParams struct {
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
}

type ApplyUploadResponse struct {
	*tchttp.BaseResponse
	Response *ApplyUploadResponseParams `json:"Response"`
}

func (r *ApplyUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

type AsrFullTextConfigureInfoForUpdate struct {
	// Switch of full speech recognition task. Valid values:
	// <li>ON: enables intelligent full speech recognition task;</li>
	// <li>OFF: disables intelligent full speech recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Format of generated subtitles file. If an empty string is entered, no subtitles files will be generated. Valid values:
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
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`
}

type AsrWordsConfigureInfoForUpdate struct {
	// Switch of speech keyword recognition task. Valid values:
	// <li>ON: enables speech keyword recognition task;</li>
	// <li>OFF: disables speech keyword recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty or a blank value is entered, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`
}

// Predefined struct for user
type AttachMediaSubtitlesRequestParams struct {
	// Unique ID of the media file
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Operation. Valid values:
	// <li>`Attach`: associates subtitles.</li>
	// <li>`Detach`: disassociates subtitles.</li>
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// [Adaptive bitrate streaming template ID](https://intl.cloud.tencent.com/document/product/266/34071?from_cn_redirect=1#zsy)
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// Unique IDs of the subtitles
	SubtitleIds []*string `json:"SubtitleIds,omitempty" name:"SubtitleIds"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type AttachMediaSubtitlesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the media file
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Operation. Valid values:
	// <li>`Attach`: associates subtitles.</li>
	// <li>`Detach`: disassociates subtitles.</li>
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// [Adaptive bitrate streaming template ID](https://intl.cloud.tencent.com/document/product/266/34071?from_cn_redirect=1#zsy)
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// Unique IDs of the subtitles
	SubtitleIds []*string `json:"SubtitleIds,omitempty" name:"SubtitleIds"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *AttachMediaSubtitlesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachMediaSubtitlesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "Operation")
	delete(f, "AdaptiveDynamicStreamingDefinition")
	delete(f, "SubtitleIds")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachMediaSubtitlesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachMediaSubtitlesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachMediaSubtitlesResponse struct {
	*tchttp.BaseResponse
	Response *AttachMediaSubtitlesResponseParams `json:"Response"`
}

func (r *AttachMediaSubtitlesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachMediaSubtitlesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AudioTemplateInfo struct {
	// The audio codec.
	// If `Container` parameter is `mp3`, the valid value is:
	// <li>libmp3lame</li>
	// If `Container` is `ogg` or `flac`, the valid value is:
	// <li>flac</li>
	// If `Container` is `m4a`, the valid values are:
	// <li>libfdk_aac</li>
	// <li>libmp3lame</li>
	// <li>ac3</li>
	// If `Container` is `mp4` or `flv`, the valid values are:
	// <li>libfdk_aac: more suitable for mp4</li>
	// <li>libmp3lame: More suitable for flv</li>
	// <li>mp2</li>
	// If `Container` is `hls`, the valid values are:
	// <li>libfdk_aac</li>
	// If `Format` is `HLS` or `MPEG-DASH`, the valid values are:
	// <li>libfdk_aac</li>
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
	// <li>1: mono-channel</li>
	// <li>2: dual-channel</li>
	// <li>6: stereo</li>
	// You cannot set the sound channel as stereo for media files in container formats for audios (FLAC, OGG, MP3, M4A).
	// Default value: 2
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type AudioTemplateInfoForUpdate struct {
	// The audio codec.
	// If `Container` parameter is `mp3`, the valid value is:
	// <li>libmp3lame</li>
	// If `Container` is `ogg` or `flac`, the valid value is:
	// <li>flac</li>
	// If `Container` is `m4a`, the valid values are:
	// <li>libfdk_aac</li>
	// <li>libmp3lame</li>
	// <li>ac3</li>
	// If `Container` is `mp4` or `flv`, the valid values are:
	// <li>libfdk_aac: more suitable for mp4</li>
	// <li>libmp3lame: More suitable for flv</li>
	// <li>mp2</li>
	// If `Container` is `hls`, the valid values are:
	// <li>libfdk_aac</li>
	// If `Format` is `HLS` or `MPEG-DASH`, the valid values are:
	// <li>libfdk_aac</li>
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
	// <li>1: mono-channel</li>
	// <li>2: dual-channel</li>
	// <li>6: stereo</li>
	// You cannot set the sound channel as stereo for media files in container formats for audios (FLAC, OGG, MP3, M4A).
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type AudioTrackItem struct {
	// Source of media material for audio segment, which can be:
	// <li>ID of VOD media files</li>
	// <li>Download URL of other media files</li>
	// Note: when a download URL of other media files is used as the material source and access control (such as hotlink protection) is enabled, the URL needs to carry access control parameters (such as hotlink protection signature).
	SourceMedia *string `json:"SourceMedia,omitempty" name:"SourceMedia"`

	// Start time of audio segment in material file in seconds. Default value: 0, which means to start capturing from the beginning position of the material.
	SourceMediaStartTime *float64 `json:"SourceMediaStartTime,omitempty" name:"SourceMediaStartTime"`

	// Audio segment duration in seconds. By default, the length of the material will be used, which means that the entire material will be captured.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// Operation on audio segment, such as volume adjustment.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioOperations []*AudioTransform `json:"AudioOperations,omitempty" name:"AudioOperations"`
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

	// Audio gain. Value range: 0-10.
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

	// Canvas width, which is the width of the output video. Value range: 0-4096 px.
	// Default value: 0, which means that the value is the same as the video width of the first video segment in the first video track.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Canvas height, which is the height (or long side) of the output video. Value range: 0-4096 px.
	// Default value: 0, which means that the value is the same as the video height of the first video segment in the first video track.
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type CdnLogInfo struct {
	// Log date in the format of `yyyy-MM-dd`, such as 2018-03-01.
	Date *string `json:"Date,omitempty" name:"Date"`

	// Log name in the format of date and time-domain name,
	// such as 2018120101-test.vod2.mqcloud.com.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Log download link, which is valid for 24 hours.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Log start time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log end time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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

// Predefined struct for user
type CommitUploadRequestParams struct {
	// VOD session, which takes the returned value (VodSessionKey) of the `ApplyUpload` API.
	VodSessionKey *string `json:"VodSessionKey,omitempty" name:"VodSessionKey"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type CommitUploadRequest struct {
	*tchttp.BaseRequest
	
	// VOD session, which takes the returned value (VodSessionKey) of the `ApplyUpload` API.
	VodSessionKey *string `json:"VodSessionKey,omitempty" name:"VodSessionKey"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CommitUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VodSessionKey")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CommitUploadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommitUploadResponseParams struct {
	// Unique ID of media file.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// The media playback URL.
	MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

	// The thumbnail URL.
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CommitUploadResponse struct {
	*tchttp.BaseResponse
	Response *CommitUploadResponseParams `json:"Response"`
}

func (r *CommitUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComposeMediaOutput struct {
	// Filename of up to 64 characters.
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Description, which can contain up to 128 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](https://intl.cloud.tencent.com/document/product/266/7812?from_cn_redirect=1) API.
	// <li>Default value: 0, which means "Other".</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Expiration time of output media file in ISO 8601 format, after which the file will be deleted. Files will never expire by default. For more information, please see [Notes on ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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

// Predefined struct for user
type ComposeMediaRequestParams struct {
	// List of input media tracks, including video, audio, and image tracks. <li>Input tracks are synced to the output media file.</li><li>Input tracks are synced to each other. Videos and images in higher tracks are superimposed over those in lower tracks. Audio tracks are mixed.</li><li>There can be up to 10 tracks for video, audio, and images each.</li><li>The total number of clips in all tracks cannot exceed 500.</li>
	Tracks []*MediaTrack `json:"Tracks,omitempty" name:"Tracks"`

	// Information of output media file.
	Output *ComposeMediaOutput `json:"Output,omitempty" name:"Output"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Canvas used for composing video file.
	Canvas *Canvas `json:"Canvas,omitempty" name:"Canvas"`

	// Used to pass through user request information. `ComposeMediaComplete` callback will return the value of this parameter. It contains up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type ComposeMediaRequest struct {
	*tchttp.BaseRequest
	
	// List of input media tracks, including video, audio, and image tracks. <li>Input tracks are synced to the output media file.</li><li>Input tracks are synced to each other. Videos and images in higher tracks are superimposed over those in lower tracks. Audio tracks are mixed.</li><li>There can be up to 10 tracks for video, audio, and images each.</li><li>The total number of clips in all tracks cannot exceed 500.</li>
	Tracks []*MediaTrack `json:"Tracks,omitempty" name:"Tracks"`

	// Information of output media file.
	Output *ComposeMediaOutput `json:"Output,omitempty" name:"Output"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Canvas used for composing video file.
	Canvas *Canvas `json:"Canvas,omitempty" name:"Canvas"`

	// Used to pass through user request information. `ComposeMediaComplete` callback will return the value of this parameter. It contains up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

func (r *ComposeMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ComposeMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tracks")
	delete(f, "Output")
	delete(f, "SubAppId")
	delete(f, "Canvas")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ComposeMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ComposeMediaResponseParams struct {
	// Media file composing task ID, which can be used to query the status of composing task (with task type being `MakeMedia`).
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ComposeMediaResponse struct {
	*tchttp.BaseResponse
	Response *ComposeMediaResponseParams `json:"Response"`
}

func (r *ComposeMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Progress of a media file composing task. Value range: [0, 100]
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Input of media file composing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Input *ComposeMediaTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of media file composing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *ComposeMediaTaskOutput `json:"Output,omitempty" name:"Output"`

	// The metadata of the output video.
	// Note: This field may return `null`, indicating that no valid value was found.
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is not carried or is left empty, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this parameter. It can contain up to 1000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

type ComposeMediaTaskInput struct {
	// List of input media tracks, i.e., information of multiple tracks composed of video, audio, image, and other materials.
	Tracks []*MediaTrack `json:"Tracks,omitempty" name:"Tracks"`

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

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](https://intl.cloud.tencent.com/document/product/266/7812?from_cn_redirect=1) API.
	// <li>Default value: 0, which means "Other".</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Expiration time of output media file in ISO 8601 format, after which the file will be deleted. Files will never expire by default. For more information, please see [Notes on ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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
	FileInfoSet []*ConcatFileInfo2017 `json:"FileInfoSet,omitempty" name:"FileInfoSet"`
}

// Predefined struct for user
type ConfirmEventsRequestParams struct {
	// Event handler, i.e., the `EventSet. EventHandle` field in the output parameters of the [event notification pulling](https://intl.cloud.tencent.com/document/product/266/33433?from_cn_redirect=1) API.
	// Array length limit: 16.
	EventHandles []*string `json:"EventHandles,omitempty" name:"EventHandles"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type ConfirmEventsRequest struct {
	*tchttp.BaseRequest
	
	// Event handler, i.e., the `EventSet. EventHandle` field in the output parameters of the [event notification pulling](https://intl.cloud.tencent.com/document/product/266/33433?from_cn_redirect=1) API.
	// Array length limit: 16.
	EventHandles []*string `json:"EventHandles,omitempty" name:"EventHandles"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ConfirmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventHandles")
	delete(f, "ExtInfo")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfirmEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfirmEventsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ConfirmEventsResponse struct {
	*tchttp.BaseResponse
	Response *ConfirmEventsResponseParams `json:"Response"`
}

func (r *ConfirmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContentReviewTemplateItem struct {
	// Unique ID of an intelligent recognition template
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Name of an intelligent recognition template. Max 64 characters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of an intelligent recognition template. Max 256 characters
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Parameters for recognition of pornographic content
	// Note: This field may return `null`, indicating that no valid value can be found.
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// Parameters for recognition of terrorism content
	// Note: This field may return `null`, indicating that no valid value can be found.
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// Parameters for recognition of politically sensitive content
	// Note: This field may return `null`, indicating that no valid value can be found.
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitempty" name:"ProhibitedConfigure"`

	// Custom recognition parameters
	// Note: This field may return `null`, indicating that no valid value can be found.
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// Whether to subject the recognition result to human review
	// <li>ON</li>
	// <li>OFF</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitempty" name:"ReviewWallSwitch"`

	// Frame capturing interval in seconds. If this parameter is left empty, 1 second will be used by default. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet"`
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

// Predefined struct for user
type CreateAIAnalysisTemplateRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
}

type CreateAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ClassificationConfigure")
	delete(f, "TagConfigure")
	delete(f, "CoverConfigure")
	delete(f, "FrameTagConfigure")
	delete(f, "HighlightConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIAnalysisTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIAnalysisTemplateResponseParams struct {
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIAnalysisTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateAIRecognitionTemplateRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Video content recognition template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter of video opening and ending credits recognition.
	HeadTailConfigure *HeadTailConfigureInfo `json:"HeadTailConfigure,omitempty" name:"HeadTailConfigure"`

	// Control parameter of video splitting recognition.
	SegmentConfigure *SegmentConfigureInfo `json:"SegmentConfigure,omitempty" name:"SegmentConfigure"`

	// Control parameter of face recognition.
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// Control parameter of full text recognition.
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// Control parameter of text keyword recognition.
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// Control parameter of full speech recognition.
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// Control parameter of speech keyword recognition.
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`

	// Control parameter of object recognition.
	ObjectConfigure *ObjectConfigureInfo `json:"ObjectConfigure,omitempty" name:"ObjectConfigure"`

	// Frame capturing interval in seconds. If this parameter is left empty, 1 second will be used by default. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
}

type CreateAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Video content recognition template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter of video opening and ending credits recognition.
	HeadTailConfigure *HeadTailConfigureInfo `json:"HeadTailConfigure,omitempty" name:"HeadTailConfigure"`

	// Control parameter of video splitting recognition.
	SegmentConfigure *SegmentConfigureInfo `json:"SegmentConfigure,omitempty" name:"SegmentConfigure"`

	// Control parameter of face recognition.
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// Control parameter of full text recognition.
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// Control parameter of text keyword recognition.
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// Control parameter of full speech recognition.
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// Control parameter of speech keyword recognition.
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`

	// Control parameter of object recognition.
	ObjectConfigure *ObjectConfigureInfo `json:"ObjectConfigure,omitempty" name:"ObjectConfigure"`

	// Frame capturing interval in seconds. If this parameter is left empty, 1 second will be used by default. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "HeadTailConfigure")
	delete(f, "SegmentConfigure")
	delete(f, "FaceConfigure")
	delete(f, "OcrFullTextConfigure")
	delete(f, "OcrWordsConfigure")
	delete(f, "AsrFullTextConfigure")
	delete(f, "AsrWordsConfigure")
	delete(f, "ObjectConfigure")
	delete(f, "ScreenshotInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIRecognitionTemplateResponseParams struct {
	// Unique ID of video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIRecognitionTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateAdaptiveDynamicStreamingTemplateRequestParams struct {
	// The adaptive bitrate streaming format. Valid values:
	// <li>HLS</li>
	// <li>MPEG-DASH</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output.
	// Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitempty" name:"StreamInfos"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The DRM type. Valid values:
	// <li>SimpleAES</li>
	// <li>Widevine</li>
	// <li>FairPlay</li>
	// If this parameter is an empty string, it indicates that the video is not protected by DRM.
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: no.
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: no.
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitempty" name:"DisableHigherVideoResolution"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type CreateAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// The adaptive bitrate streaming format. Valid values:
	// <li>HLS</li>
	// <li>MPEG-DASH</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output.
	// Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitempty" name:"StreamInfos"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The DRM type. Valid values:
	// <li>SimpleAES</li>
	// <li>Widevine</li>
	// <li>FairPlay</li>
	// If this parameter is an empty string, it indicates that the video is not protected by DRM.
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: no.
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: no.
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "DrmType")
	delete(f, "DisableHigherVideoBitrate")
	delete(f, "DisableHigherVideoResolution")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdaptiveDynamicStreamingTemplateResponseParams struct {
	// Unique ID of adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAdaptiveDynamicStreamingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAdaptiveDynamicStreamingTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateAnimatedGraphicsTemplateRequestParams struct {
	// Video frame rate in Hz. Value range: [1, 30].
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
}

type CreateAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Video frame rate in Hz. Value range: [1, 30].
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type CreateAnimatedGraphicsTemplateResponseParams struct {
	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAnimatedGraphicsTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateClassRequestParams struct {
	// Parent category ID. For a first-level category, enter `-1`.
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// Category name. Length limit: 1-64 characters.
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type CreateClassRequest struct {
	*tchttp.BaseRequest
	
	// Parent category ID. For a first-level category, enter `-1`.
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// Category name. Length limit: 1-64 characters.
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ParentId")
	delete(f, "ClassName")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClassResponseParams struct {
	// Category ID
	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClassResponse struct {
	*tchttp.BaseResponse
	Response *CreateClassResponseParams `json:"Response"`
}

func (r *CreateClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContentReviewTemplateRequestParams struct {
	// Whether to allow the recognition result to enter the intelligent recognition platform (for human recognition).
	// <li>ON: yes</li>
	// <li>OFF: no</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitempty" name:"ReviewWallSwitch"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Name of an intelligent content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of an intelligent content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter for porn information.
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// Control parameter for terrorism information.
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// Control parameter for politically sensitive information.
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitempty" name:"ProhibitedConfigure"`

	// Control parameter for custom intelligent content recognition.
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// Frame capturing interval in seconds. If this parameter is left empty, 1 second will be used by default. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
}

type CreateContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Whether to allow the recognition result to enter the intelligent recognition platform (for human recognition).
	// <li>ON: yes</li>
	// <li>OFF: no</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitempty" name:"ReviewWallSwitch"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Name of an intelligent content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of an intelligent content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter for porn information.
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// Control parameter for terrorism information.
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// Control parameter for politically sensitive information.
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitempty" name:"ProhibitedConfigure"`

	// Control parameter for custom intelligent content recognition.
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// Frame capturing interval in seconds. If this parameter is left empty, 1 second will be used by default. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
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
	delete(f, "ReviewWallSwitch")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "PornConfigure")
	delete(f, "TerrorismConfigure")
	delete(f, "PoliticalConfigure")
	delete(f, "ProhibitedConfigure")
	delete(f, "UserDefineConfigure")
	delete(f, "ScreenshotInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContentReviewTemplateResponseParams struct {
	// Unique ID of an intelligent recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateContentReviewTemplateResponseParams `json:"Response"`
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

	// Image sprite specification. For more information, please see [Image Sprite Generating Template](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Total number of subimages in image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Address of output image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageSpriteUrlSet []*string `json:"ImageSpriteUrlSet,omitempty" name:"ImageSpriteUrlSet"`

	// Address of WebVtt file for the position-time relationship among subimages in an image sprite.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WebVttUrl *string `json:"WebVttUrl,omitempty" name:"WebVttUrl"`
}

// Predefined struct for user
type CreateImageSpriteTemplateRequestParams struct {
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

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Name of an image sprite generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`

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

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Name of an image sprite generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`

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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "FillType")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ResolutionAdaptive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageSpriteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageSpriteTemplateResponseParams struct {
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageSpriteTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreatePersonSampleRequestParams struct {
	// Name of a sample. Length limit: 20 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Usage of a sample. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: equivalent to 1+2.
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Description of a sample. Length limit: 1024 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// String generated after the sample image is encoded by [Base64](https://tools.ietf.org/html/rfc4648). Only JPEG and PNG images are supported. Array length limit: 5 images.
	// Note: the image must be a relatively clear full-face photo of a person and has a resolution of no less than 200 x 200.
	FaceContents []*string `json:"FaceContents,omitempty" name:"FaceContents"`

	// Tags of a sample
	// <li>Array length limit: 20 tags</li>
	// <li>Length limit of a tag: 128 characters</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type CreatePersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// Name of a sample. Length limit: 20 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Usage of a sample. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: equivalent to 1+2.
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Description of a sample. Length limit: 1024 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// String generated after the sample image is encoded by [Base64](https://tools.ietf.org/html/rfc4648). Only JPEG and PNG images are supported. Array length limit: 5 images.
	// Note: the image must be a relatively clear full-face photo of a person and has a resolution of no less than 200 x 200.
	FaceContents []*string `json:"FaceContents,omitempty" name:"FaceContents"`

	// Tags of a sample
	// <li>Array length limit: 20 tags</li>
	// <li>Length limit of a tag: 128 characters</li>
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
	delete(f, "SubAppId")
	delete(f, "Description")
	delete(f, "FaceContents")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonSampleResponseParams struct {
	// Information of a sample.
	Person *AiSamplePerson `json:"Person,omitempty" name:"Person"`

	// Information of samples that failed the verification by facial feature positioning.
	FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitempty" name:"FailFaceInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *CreatePersonSampleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateProcedureTemplateRequestParams struct {
	// Task flow name (up to 20 characters).
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Parameter of video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Intelligent recognition task
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Parameter of AI-based content analysis task.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of AI-based content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type CreateProcedureTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Task flow name (up to 20 characters).
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Parameter of video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Intelligent recognition task
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Parameter of AI-based content analysis task.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of AI-based content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcedureTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProcedureTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProcedureTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateProcedureTemplateResponseParams `json:"Response"`
}

func (r *CreateProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcedureTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSampleSnapshotTemplateRequestParams struct {
	// Sampled screencapturing type. Valid values:
	// <li>Percent: by percent.</li>
	// <li>Time: by time interval.</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitempty" name:"SampleInterval"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
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

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type CreateSampleSnapshotTemplateResponseParams struct {
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateSampleSnapshotTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateSnapshotByTimeOffsetTemplateRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type CreateSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type CreateSnapshotByTimeOffsetTemplateResponseParams struct {
	// Unique ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotByTimeOffsetTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateStorageRegionRequestParams struct {
	// The region to enable storage in, which must be a storage region supported by VOD.
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type CreateStorageRegionRequest struct {
	*tchttp.BaseRequest
	
	// The region to enable storage in, which must be a storage region supported by VOD.
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateStorageRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StorageRegion")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStorageRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageRegionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStorageRegionResponse struct {
	*tchttp.BaseResponse
	Response *CreateStorageRegionResponseParams `json:"Response"`
}

func (r *CreateStorageRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubAppIdRequestParams struct {
	// Subapplication name. Length limit: 40 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subapplication overview. Length limit: 300 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateSubAppIdRequest struct {
	*tchttp.BaseRequest
	
	// Subapplication name. Length limit: 40 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subapplication overview. Length limit: 300 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSubAppIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubAppIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubAppIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubAppIdResponseParams struct {
	// ID of created subapplication.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSubAppIdResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubAppIdResponseParams `json:"Response"`
}

func (r *CreateSubAppIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSuperPlayerConfigRequestParams struct {
	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	Name *string `json:"Name,omitempty" name:"Name"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Type of audio/video played. Valid values:
	// <li>AdaptiveDynamicStreaming</li>
	// <li>Transcode</li>
	// <li>Original</li>
	// Default value: `AdaptiveDynamicStream`
	AudioVideoType *string `json:"AudioVideoType,omitempty" name:"AudioVideoType"`

	// Whether to allow only adaptive bitrate streaming playback protected by DRM. Valid values:
	// <li>`ON`: allow only adaptive bitrate streaming playback protected by DRM</li>
	// <li>`OFF`: allow adaptive bitrate streaming playback not protected by DRM</li>
	// Default value: `OFF`
	// This parameter is valid when `AudioVideoType` is `AdaptiveDynamicStream`.
	DrmSwitch *string `json:"DrmSwitch,omitempty" name:"DrmSwitch"`

	// ID of the adaptive bitrate streaming template allowed for playback not protected by DRM.
	// 
	// This parameter is required if `AudioVideoType` is `AdaptiveDynamicStream` and `DrmSwitch` is `OFF`.
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// Content of the adaptive bitrate streaming template allowed for playback protected by DRM.
	// 
	// This parameter is required if `AudioVideoType` is `AdaptiveDynamicStream` and `DrmSwitch` is `ON`.
	DrmStreamingsInfo *DrmStreamingsInfo `json:"DrmStreamingsInfo,omitempty" name:"DrmStreamingsInfo"`

	// ID of the transcoding template allowed for playback
	// 
	// This parameter is required if `AudioVideoType` is `Transcode`.
	TranscodeDefinition *uint64 `json:"TranscodeDefinition,omitempty" name:"TranscodeDefinition"`

	// ID of the image sprite generating template that allows output.
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitempty" name:"ImageSpriteDefinition"`

	// Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used:
	// <li>MinEdgeLength: 240, Name: LD;</li>
	// <li>MinEdgeLength: 480, Name: SD;</li>
	// <li>MinEdgeLength: 720, Name: HD;</li>
	// <li>MinEdgeLength: 1080, Name: FHD;</li>
	// <li>MinEdgeLength: 1440, Name: 2K;</li>
	// <li>MinEdgeLength: 2160, Name: 4K;</li>
	// <li>MinEdgeLength: 4320, Name: 8K.</li>
	ResolutionNames []*ResolutionNameInfo `json:"ResolutionNames,omitempty" name:"ResolutionNames"`

	// Domain name used for playback. If it is left empty or set to `Default`, the domain name configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Scheme used for playback. If it is left empty or set to `Default`, the scheme configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used. Other valid values:
	// <li>HTTP;</li>
	// <li>HTTPS.</li>
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type CreateSuperPlayerConfigRequest struct {
	*tchttp.BaseRequest
	
	// Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	Name *string `json:"Name,omitempty" name:"Name"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Type of audio/video played. Valid values:
	// <li>AdaptiveDynamicStreaming</li>
	// <li>Transcode</li>
	// <li>Original</li>
	// Default value: `AdaptiveDynamicStream`
	AudioVideoType *string `json:"AudioVideoType,omitempty" name:"AudioVideoType"`

	// Whether to allow only adaptive bitrate streaming playback protected by DRM. Valid values:
	// <li>`ON`: allow only adaptive bitrate streaming playback protected by DRM</li>
	// <li>`OFF`: allow adaptive bitrate streaming playback not protected by DRM</li>
	// Default value: `OFF`
	// This parameter is valid when `AudioVideoType` is `AdaptiveDynamicStream`.
	DrmSwitch *string `json:"DrmSwitch,omitempty" name:"DrmSwitch"`

	// ID of the adaptive bitrate streaming template allowed for playback not protected by DRM.
	// 
	// This parameter is required if `AudioVideoType` is `AdaptiveDynamicStream` and `DrmSwitch` is `OFF`.
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// Content of the adaptive bitrate streaming template allowed for playback protected by DRM.
	// 
	// This parameter is required if `AudioVideoType` is `AdaptiveDynamicStream` and `DrmSwitch` is `ON`.
	DrmStreamingsInfo *DrmStreamingsInfo `json:"DrmStreamingsInfo,omitempty" name:"DrmStreamingsInfo"`

	// ID of the transcoding template allowed for playback
	// 
	// This parameter is required if `AudioVideoType` is `Transcode`.
	TranscodeDefinition *uint64 `json:"TranscodeDefinition,omitempty" name:"TranscodeDefinition"`

	// ID of the image sprite generating template that allows output.
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitempty" name:"ImageSpriteDefinition"`

	// Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used:
	// <li>MinEdgeLength: 240, Name: LD;</li>
	// <li>MinEdgeLength: 480, Name: SD;</li>
	// <li>MinEdgeLength: 720, Name: HD;</li>
	// <li>MinEdgeLength: 1080, Name: FHD;</li>
	// <li>MinEdgeLength: 1440, Name: 2K;</li>
	// <li>MinEdgeLength: 2160, Name: 4K;</li>
	// <li>MinEdgeLength: 4320, Name: 8K.</li>
	ResolutionNames []*ResolutionNameInfo `json:"ResolutionNames,omitempty" name:"ResolutionNames"`

	// Domain name used for playback. If it is left empty or set to `Default`, the domain name configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Scheme used for playback. If it is left empty or set to `Default`, the scheme configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used. Other valid values:
	// <li>HTTP;</li>
	// <li>HTTPS.</li>
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateSuperPlayerConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSuperPlayerConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	delete(f, "AudioVideoType")
	delete(f, "DrmSwitch")
	delete(f, "AdaptiveDynamicStreamingDefinition")
	delete(f, "DrmStreamingsInfo")
	delete(f, "TranscodeDefinition")
	delete(f, "ImageSpriteDefinition")
	delete(f, "ResolutionNames")
	delete(f, "Domain")
	delete(f, "Scheme")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSuperPlayerConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSuperPlayerConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSuperPlayerConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateSuperPlayerConfigResponseParams `json:"Response"`
}

func (r *CreateSuperPlayerConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSuperPlayerConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeTemplateRequestParams struct {
	// Container. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
}

type CreateTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Container. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type CreateTranscodeTemplateResponseParams struct {
	// Unique ID of transcoding template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateTranscodeTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateVodDomainRequestParams struct {
	// Domain name to add to VOD. Note: a wildcard domain name is not supported.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Region to enable acceleration. Valid values:
	// <li>`Chinese Mainland`</li>
	// <li>`Outside Chinese Mainland`</li>
	// <li>`Global`</li>
	// If `AccelerateArea` is not specified, VOD will enable acceleration in or outside Chinese mainland based on the regional information a user has configured with Tencent Cloud.
	AccelerateArea *string `json:"AccelerateArea,omitempty" name:"AccelerateArea"`
}

type CreateVodDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain name to add to VOD. Note: a wildcard domain name is not supported.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Region to enable acceleration. Valid values:
	// <li>`Chinese Mainland`</li>
	// <li>`Outside Chinese Mainland`</li>
	// <li>`Global`</li>
	// If `AccelerateArea` is not specified, VOD will enable acceleration in or outside Chinese mainland based on the regional information a user has configured with Tencent Cloud.
	AccelerateArea *string `json:"AccelerateArea,omitempty" name:"AccelerateArea"`
}

func (r *CreateVodDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVodDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SubAppId")
	delete(f, "AccelerateArea")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVodDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVodDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVodDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateVodDomainResponseParams `json:"Response"`
}

func (r *CreateVodDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVodDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWatermarkTemplateRequestParams struct {
	// Watermarking type. Valid values:
	// <li>image: image watermark;</li>
	// <li>text: text watermark;</li>
	// <li>svg: SVG watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
}

type CreateWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Watermarking type. Valid values:
	// <li>image: image watermark;</li>
	// <li>text: text watermark;</li>
	// <li>svg: SVG watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type CreateWatermarkTemplateResponseParams struct {
	// Unique ID of watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Watermark image address. This field is valid only when `Type` is `image`.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateWatermarkTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateWordSamplesRequestParams struct {
	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// 7. All: ASR- and OCR-based content recognition and inappropriate information recognition; equivalent to 1+2+3+4
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Keyword. Array length limit: 100.
	Words []*AiSampleWordInfo `json:"Words,omitempty" name:"Words"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	// 7. All: ASR- and OCR-based content recognition and inappropriate information recognition; equivalent to 1+2+3+4
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Keyword. Array length limit: 100.
	Words []*AiSampleWordInfo `json:"Words,omitempty" name:"Words"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWordSamplesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *CreateWordSamplesResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteAIAnalysisTemplateRequestParams struct {
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIAnalysisTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIAnalysisTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAIAnalysisTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteAIRecognitionTemplateRequestParams struct {
	// Unique ID of video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIRecognitionTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAIRecognitionTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteAdaptiveDynamicStreamingTemplateRequestParams struct {
	// Unique ID of adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAdaptiveDynamicStreamingTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAdaptiveDynamicStreamingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAdaptiveDynamicStreamingTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteAnimatedGraphicsTemplateRequestParams struct {
	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAnimatedGraphicsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAnimatedGraphicsTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAnimatedGraphicsTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteClassRequestParams struct {
	// Category ID
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteClassRequest struct {
	*tchttp.BaseRequest
	
	// Category ID
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClassId")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClassResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClassResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClassResponseParams `json:"Response"`
}

func (r *DeleteClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContentReviewTemplateRequestParams struct {
	// Unique ID of an intelligent content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an intelligent content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContentReviewTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteContentReviewTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteImageSpriteTemplateRequestParams struct {
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageSpriteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageSpriteTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageSpriteTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteMediaRequestParams struct {
	// Unique media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Content to be deleted. The default value is "[]", which indicates to delete the media file and all its corresponding files generated by video processing.
	DeleteParts []*MediaDeleteItem `json:"DeleteParts,omitempty" name:"DeleteParts"`
}

type DeleteMediaRequest struct {
	*tchttp.BaseRequest
	
	// Unique media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Content to be deleted. The default value is "[]", which indicates to delete the media file and all its corresponding files generated by video processing.
	DeleteParts []*MediaDeleteItem `json:"DeleteParts,omitempty" name:"DeleteParts"`
}

func (r *DeleteMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "DeleteParts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMediaResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMediaResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMediaResponseParams `json:"Response"`
}

func (r *DeleteMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonSampleRequestParams struct {
	// ID of a sample.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeletePersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// ID of a sample.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonSampleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *DeletePersonSampleResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteProcedureTemplateRequestParams struct {
	// Task flow name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteProcedureTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Task flow name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProcedureTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProcedureTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProcedureTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProcedureTemplateResponseParams `json:"Response"`
}

func (r *DeleteProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProcedureTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSampleSnapshotTemplateRequestParams struct {
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSampleSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSampleSnapshotTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSampleSnapshotTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteSnapshotByTimeOffsetTemplateRequestParams struct {
	// Unique ID of a specified time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a specified time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotByTimeOffsetTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotByTimeOffsetTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotByTimeOffsetTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteSuperPlayerConfigRequestParams struct {
	// Player configuration name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteSuperPlayerConfigRequest struct {
	*tchttp.BaseRequest
	
	// Player configuration name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteSuperPlayerConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSuperPlayerConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSuperPlayerConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSuperPlayerConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSuperPlayerConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSuperPlayerConfigResponseParams `json:"Response"`
}

func (r *DeleteSuperPlayerConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSuperPlayerConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTranscodeTemplateRequestParams struct {
	// Unique ID of transcoding template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of transcoding template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTranscodeTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTranscodeTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteVodDomainRequestParams struct {
	// Domain name to delete from VOD
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteVodDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain name to delete from VOD
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteVodDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVodDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVodDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVodDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVodDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVodDomainResponseParams `json:"Response"`
}

func (r *DeleteVodDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVodDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWatermarkTemplateRequestParams struct {
	// Unique ID of watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWatermarkTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWatermarkTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWatermarkTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteWordSamplesRequestParams struct {
	// Keyword. Array length limit: 100 words.
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DeleteWordSamplesRequest struct {
	*tchttp.BaseRequest
	
	// Keyword. Array length limit: 100 words.
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWordSamplesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWordSamplesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAIAnalysisTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of video content analysis templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAIAnalysisTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of video content analysis templates. Array length limit: 100.
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIAnalysisTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAnalysisTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of video content analysis template details.
	AIAnalysisTemplateSet []*AIAnalysisTemplateItem `json:"AIAnalysisTemplateSet,omitempty" name:"AIAnalysisTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAIAnalysisTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIAnalysisTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAIRecognitionTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of video content recognition templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAIRecognitionTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of video content recognition templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIRecognitionTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIRecognitionTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of video content recognition template details.
	AIRecognitionTemplateSet []*AIRecognitionTemplateItem `json:"AIRecognitionTemplateSet,omitempty" name:"AIRecognitionTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAIRecognitionTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIRecognitionTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAdaptiveDynamicStreamingTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of transcoding to adaptive bitrate streaming templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeAdaptiveDynamicStreamingTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of transcoding to adaptive bitrate streaming templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdaptiveDynamicStreamingTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdaptiveDynamicStreamingTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of transcoding to adaptive bitrate streaming template details.
	AdaptiveDynamicStreamingTemplateSet []*AdaptiveDynamicStreamingTemplate `json:"AdaptiveDynamicStreamingTemplateSet,omitempty" name:"AdaptiveDynamicStreamingTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAdaptiveDynamicStreamingTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAdaptiveDynamicStreamingTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAllClassRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeAllClassRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAllClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllClassResponseParams struct {
	// Category information set
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassInfoSet []*MediaClassInfo `json:"ClassInfoSet,omitempty" name:"ClassInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllClassResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllClassResponseParams `json:"Response"`
}

func (r *DescribeAllClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnimatedGraphicsTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of animated image generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeAnimatedGraphicsTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of animated image generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAnimatedGraphicsTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnimatedGraphicsTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of animated image generating template details.
	AnimatedGraphicsTemplateSet []*AnimatedGraphicsTemplate `json:"AnimatedGraphicsTemplateSet,omitempty" name:"AnimatedGraphicsTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAnimatedGraphicsTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAnimatedGraphicsTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCDNStatDetailsRequestParams struct {
	// Metrics to query. Valid values:
	// <li>`Traffic`: traffic in bytes</li>
	// <li>`Bandwidth`: bandwidth in bps</li>
	// <li>`Requests`: the number of requests</li>
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Start time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// List of domain names. The usage data of up to 20 domain names can be queried at a time. The usage data of all domain names is returned by default.
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames"`

	// Service region. Valid values:
	// <li>Chinese Mainland</li>
	// <li>Asia Pacific Region 1: Hong Kong (China), Macao (China), Singapore, Vietnam, and Thailand</li>
	// <li>Asia Pacific Region 2: Taiwan (China), Japan, Malaysia, Indonesia, and South Korea</li>
	// <li>Asia Pacific Region 3: Philippines, India, Australia, and other Asia Pacific countries and regions</li>
	// <li>Middle East</li>
	// <li>Europe</li>
	// <li>North America</li>
	// <li>South America</li>
	// <li>Africa</li>
	// Default value: Chinese Mainland
	Area *string `json:"Area,omitempty" name:"Area"`

	// District where users are located. When `Area` is `Chinese Mainland`, valid values for `Districts` are as follows. Otherwise, `Districts` can be ignored.
	// <li>Beijing</li>
	// <li>Inner Mongolia</li>
	// <li>Shanxi</li>
	// <li>Hebei</li>
	// <li>Tianjin</li>
	// <li>Ningxia</li>
	// <li>Shaanxi</li>
	// <li>Gansu</li>
	// <li>Qinghai</li>
	// <li>Xinjiang</li>
	// <li>Heilongjiang</li>
	// <li>Jilin</li>
	// <li>Liaoning</li>
	// <li>Fujian</li>
	// <li>Jiangsu</li>
	// <li>Anhui</li>
	// <li>Shandong</li>
	// <li>Shanghai</li>
	// <li>Zhejiang</li>
	// <li>Henan</li>
	// <li>Hubei</li>
	// <li>Jiangxi</li>
	// <li>Hunan</li>
	// <li>Guizhou</li>
	// <li>Yunnan</li>
	// <li>Chongqing</li>
	// <li>Sichuan</li>
	// <li>Tibet</li>
	// <li>Guangdong</li>
	// <li>Guangxi</li>
	// <li>Hainan</li>
	// <li>Hong Kong, Macao and Taiwan</li>
	// <li>Outside Chinese Mainland</li>
	// <li>Other</li>
	Districts []*string `json:"Districts,omitempty" name:"Districts"`

	// ISP of users. When `Area` is `Chinese Mainland`, valid values for `Isps` are as follows. Otherwise, `Isps` can be ignored.
	// <li>China Telecom</li>
	// <li>China Unicom</li>
	// <li>CERNET</li>
	// <li>Great Wall Broadband Network</li>
	// <li>China Mobile</li>
	// <li>China Mobile Tietong</li>
	// <li>ISPs outside Chinese Mainland</li>
	// <li>Other ISPs</li>
	Isps []*string `json:"Isps,omitempty" name:"Isps"`

	// Time granularity of every piece of data in minutes. Valid values:
	// <li>5: 5-minute granularity. The data at 5-minute granularity in the query period will be returned.</li>
	// <li>1440: 1-day granularity. The data at 1-day granularity in the query period will be returned. If the query period is larger than 24 hours, only data at 1-day granularity can be queried.</li>
	// If the difference between `StartTime` and `EndTime` is larger than 24 hours, the default value of `DataInterval` is 1440.
	DataInterval *uint64 `json:"DataInterval,omitempty" name:"DataInterval"`
}

type DescribeCDNStatDetailsRequest struct {
	*tchttp.BaseRequest
	
	// Metrics to query. Valid values:
	// <li>`Traffic`: traffic in bytes</li>
	// <li>`Bandwidth`: bandwidth in bps</li>
	// <li>`Requests`: the number of requests</li>
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Start time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// List of domain names. The usage data of up to 20 domain names can be queried at a time. The usage data of all domain names is returned by default.
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames"`

	// Service region. Valid values:
	// <li>Chinese Mainland</li>
	// <li>Asia Pacific Region 1: Hong Kong (China), Macao (China), Singapore, Vietnam, and Thailand</li>
	// <li>Asia Pacific Region 2: Taiwan (China), Japan, Malaysia, Indonesia, and South Korea</li>
	// <li>Asia Pacific Region 3: Philippines, India, Australia, and other Asia Pacific countries and regions</li>
	// <li>Middle East</li>
	// <li>Europe</li>
	// <li>North America</li>
	// <li>South America</li>
	// <li>Africa</li>
	// Default value: Chinese Mainland
	Area *string `json:"Area,omitempty" name:"Area"`

	// District where users are located. When `Area` is `Chinese Mainland`, valid values for `Districts` are as follows. Otherwise, `Districts` can be ignored.
	// <li>Beijing</li>
	// <li>Inner Mongolia</li>
	// <li>Shanxi</li>
	// <li>Hebei</li>
	// <li>Tianjin</li>
	// <li>Ningxia</li>
	// <li>Shaanxi</li>
	// <li>Gansu</li>
	// <li>Qinghai</li>
	// <li>Xinjiang</li>
	// <li>Heilongjiang</li>
	// <li>Jilin</li>
	// <li>Liaoning</li>
	// <li>Fujian</li>
	// <li>Jiangsu</li>
	// <li>Anhui</li>
	// <li>Shandong</li>
	// <li>Shanghai</li>
	// <li>Zhejiang</li>
	// <li>Henan</li>
	// <li>Hubei</li>
	// <li>Jiangxi</li>
	// <li>Hunan</li>
	// <li>Guizhou</li>
	// <li>Yunnan</li>
	// <li>Chongqing</li>
	// <li>Sichuan</li>
	// <li>Tibet</li>
	// <li>Guangdong</li>
	// <li>Guangxi</li>
	// <li>Hainan</li>
	// <li>Hong Kong, Macao and Taiwan</li>
	// <li>Outside Chinese Mainland</li>
	// <li>Other</li>
	Districts []*string `json:"Districts,omitempty" name:"Districts"`

	// ISP of users. When `Area` is `Chinese Mainland`, valid values for `Isps` are as follows. Otherwise, `Isps` can be ignored.
	// <li>China Telecom</li>
	// <li>China Unicom</li>
	// <li>CERNET</li>
	// <li>Great Wall Broadband Network</li>
	// <li>China Mobile</li>
	// <li>China Mobile Tietong</li>
	// <li>ISPs outside Chinese Mainland</li>
	// <li>Other ISPs</li>
	Isps []*string `json:"Isps,omitempty" name:"Isps"`

	// Time granularity of every piece of data in minutes. Valid values:
	// <li>5: 5-minute granularity. The data at 5-minute granularity in the query period will be returned.</li>
	// <li>1440: 1-day granularity. The data at 1-day granularity in the query period will be returned. If the query period is larger than 24 hours, only data at 1-day granularity can be queried.</li>
	// If the difference between `StartTime` and `EndTime` is larger than 24 hours, the default value of `DataInterval` is 1440.
	DataInterval *uint64 `json:"DataInterval,omitempty" name:"DataInterval"`
}

func (r *DescribeCDNStatDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDNStatDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Metric")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "DomainNames")
	delete(f, "Area")
	delete(f, "Districts")
	delete(f, "Isps")
	delete(f, "DataInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCDNStatDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCDNStatDetailsResponseParams struct {
	// Time granularity of every piece of data in minutes.
	DataInterval *uint64 `json:"DataInterval,omitempty" name:"DataInterval"`

	// CDN usage statistics.
	Data []*StatDataItem `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCDNStatDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCDNStatDetailsResponseParams `json:"Response"`
}

func (r *DescribeCDNStatDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDNStatDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCDNUsageDataRequestParams struct {
	// Start date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). The end date must be after the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// CDN statistics type. Valid values:
	// <li>Flux: traffic in bytes.</li>
	// <li>Bandwidth: bandwidth in bps.</li>
	DataType *string `json:"DataType,omitempty" name:"DataType"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	// You can set this parameter to 1 to query the total usage of all applications (including the primary application) as an admin (only 1-day granularity is supported).</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Time granularity of usage data in minutes. Valid values:
	// <li>5: 5-minute granularity. The data at 5-minute granularity in the query period will be returned.</li>
	// <li>60: 1-hour granularity. The data at 1-hour granularity in the query period will be returned.</li>
	// <li>1440: 1-day granularity. The data at 1-day granularity in the query period will be returned.</li>
	// Default value: 1440. Data at 1-day granularity will be returned.
	DataInterval *uint64 `json:"DataInterval,omitempty" name:"DataInterval"`

	// List of domain names. The usage data of up to 20 domain names can be queried at a time. You can specify multiple domain names and query their combined usage data. The usage data of all domain names will be returned by default.
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames"`
}

type DescribeCDNUsageDataRequest struct {
	*tchttp.BaseRequest
	
	// Start date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). The end date must be after the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// CDN statistics type. Valid values:
	// <li>Flux: traffic in bytes.</li>
	// <li>Bandwidth: bandwidth in bps.</li>
	DataType *string `json:"DataType,omitempty" name:"DataType"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	// You can set this parameter to 1 to query the total usage of all applications (including the primary application) as an admin (only 1-day granularity is supported).</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Time granularity of usage data in minutes. Valid values:
	// <li>5: 5-minute granularity. The data at 5-minute granularity in the query period will be returned.</li>
	// <li>60: 1-hour granularity. The data at 1-hour granularity in the query period will be returned.</li>
	// <li>1440: 1-day granularity. The data at 1-day granularity in the query period will be returned.</li>
	// Default value: 1440. Data at 1-day granularity will be returned.
	DataInterval *uint64 `json:"DataInterval,omitempty" name:"DataInterval"`

	// List of domain names. The usage data of up to 20 domain names can be queried at a time. You can specify multiple domain names and query their combined usage data. The usage data of all domain names will be returned by default.
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames"`
}

func (r *DescribeCDNUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDNUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DataType")
	delete(f, "SubAppId")
	delete(f, "DataInterval")
	delete(f, "DomainNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCDNUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCDNUsageDataResponseParams struct {
	// Time granularity in minutes.
	DataInterval *int64 `json:"DataInterval,omitempty" name:"DataInterval"`

	// CDN statistics.
	Data []*StatDataItem `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCDNUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCDNUsageDataResponseParams `json:"Response"`
}

func (r *DescribeCDNUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCDNUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnLogsRequestParams struct {
	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Start time for log acquisition in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F), which must be after the start time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Maximum return results of pulling paginated queries. Default value: 100; maximum value: 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number offset from the beginning of paginated queries. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeCdnLogsRequest struct {
	*tchttp.BaseRequest
	
	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Start time for log acquisition in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F), which must be after the start time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Maximum return results of pulling paginated queries. Default value: 100; maximum value: 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number offset from the beginning of paginated queries. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCdnLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnLogsResponseParams struct {
	// Total number of log download links
	// Note: this field may return `null`, indicating that no valid value is obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Log download list for CDN nodes outside Mainland China. If global acceleration is not enabled for the domain name, ignore this parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OverseaCdnLogs []*CdnLogInfo `json:"OverseaCdnLogs,omitempty" name:"OverseaCdnLogs"`

	// Log download list for CDN nodes in Mainland China.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomesticCdnLogs []*CdnLogInfo `json:"DomesticCdnLogs,omitempty" name:"DomesticCdnLogs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCdnLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnLogsResponseParams `json:"Response"`
}

func (r *DescribeCdnLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentReviewTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique IDs for filters of an intelligent content recognition template. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeContentReviewTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique IDs for filters of an intelligent content recognition template. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContentReviewTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentReviewTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of intelligent content recognition template details.
	ContentReviewTemplateSet []*ContentReviewTemplateItem `json:"ContentReviewTemplateSet,omitempty" name:"ContentReviewTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeContentReviewTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContentReviewTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDailyPlayStatFileListRequestParams struct {
	// Start date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeDailyPlayStatFileListRequest struct {
	*tchttp.BaseRequest
	
	// Start date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeDailyPlayStatFileListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDailyPlayStatFileListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDailyPlayStatFileListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDailyPlayStatFileListResponseParams struct {
	// List of playback statistics files.
	PlayStatFileSet []*PlayStatFileInfo `json:"PlayStatFileSet,omitempty" name:"PlayStatFileSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDailyPlayStatFileListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDailyPlayStatFileListResponseParams `json:"Response"`
}

func (r *DescribeDailyPlayStatFileListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDailyPlayStatFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageReviewUsageDataRequestParams struct {
	// The start date for the query in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end date for the query in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format). The end date must be later than the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeImageReviewUsageDataRequest struct {
	*tchttp.BaseRequest
	
	// The start date for the query in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end date for the query in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format). The end date must be later than the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeImageReviewUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageReviewUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageReviewUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageReviewUsageDataResponseParams struct {
	// The image recognition usage statistics (the number of times the image recognition feature is used in the time period specified).
	ImageReviewUsageDataSet []*ImageReviewUsageDataItem `json:"ImageReviewUsageDataSet,omitempty" name:"ImageReviewUsageDataSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageReviewUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageReviewUsageDataResponseParams `json:"Response"`
}

func (r *DescribeImageReviewUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageReviewUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSpriteTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of image sprite generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeImageSpriteTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of image sprite generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageSpriteTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSpriteTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of image sprite generating template details.
	ImageSpriteTemplateSet []*ImageSpriteTemplate `json:"ImageSpriteTemplateSet,omitempty" name:"ImageSpriteTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageSpriteTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageSpriteTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLicenseUsageDataRequestParams struct {
	// The start date for the query in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end date for the query in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format). The end date must be later than the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The license type, which is DRM by default. Valid values:
	// <li> DRM</li>
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeLicenseUsageDataRequest struct {
	*tchttp.BaseRequest
	
	// The start date for the query in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end date for the query in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format). The end date must be later than the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The license type, which is DRM by default. Valid values:
	// <li> DRM</li>
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeLicenseUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "LicenseType")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLicenseUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseUsageDataResponseParams struct {
	// The license request statistics (the number of license requests in the time period specified)
	LicenseUsageDataSet []*LicenseUsageDataItem `json:"LicenseUsageDataSet,omitempty" name:"LicenseUsageDataSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLicenseUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLicenseUsageDataResponseParams `json:"Response"`
}

func (r *DescribeLicenseUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaInfosRequestParams struct {
	// List of media file IDs. N starts from 0 and can be up to 19.
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

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
	Filters []*string `json:"Filters,omitempty" name:"Filters"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeMediaInfosRequest struct {
	*tchttp.BaseRequest
	
	// List of media file IDs. N starts from 0 and can be up to 19.
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

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
	Filters []*string `json:"Filters,omitempty" name:"Filters"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeMediaInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileIds")
	delete(f, "Filters")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaInfosResponseParams struct {
	// Media file information list.
	MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet"`

	// List of IDs of files that do not exist.
	NotExistFileIdSet []*string `json:"NotExistFileIdSet,omitempty" name:"NotExistFileIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMediaInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaInfosResponseParams `json:"Response"`
}

func (r *DescribeMediaInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaPlayStatDetailsRequestParams struct {
	// The ID of the media file.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// The start time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Granularity. Valid values:
	// <li>Hour</li>
	// <li>Day</li>
	// The default value depends on the time period queried. If the time period is shorter than one day, the default value is `Hour`; if the time period is one day or longer, the default value is `Day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type DescribeMediaPlayStatDetailsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the media file.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// The start time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Granularity. Valid values:
	// <li>Hour</li>
	// <li>Day</li>
	// The default value depends on the time period queried. If the time period is shorter than one day, the default value is `Hour`; if the time period is one day or longer, the default value is `Day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeMediaPlayStatDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaPlayStatDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaPlayStatDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaPlayStatDetailsResponseParams struct {
	// The playback statistics.
	PlayStatInfoSet []*PlayStatInfo `json:"PlayStatInfoSet,omitempty" name:"PlayStatInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMediaPlayStatDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaPlayStatDetailsResponseParams `json:"Response"`
}

func (r *DescribeMediaPlayStatDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaPlayStatDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaProcessUsageDataRequestParams struct {
	// Start date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F). The end date must be on or after the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The type of media processing task. Valid values:
	// <li>Transcoding: General transcoding</li>
	// <li>Transcoding-TESHD: TESHD transcoding</li>
	// <li>Editing: Video editing</li>
	// <li>Editing-TESHD: TESHD editing</li>
	// <li>AdaptiveBitrateStreaming: Adaptive bitrate streaming</li>
	// <li>ContentAudit: Content moderation</li>
	// <li>RemoveWatermark: Watermark removal</li>
	// <li>Transcode: Transcoding, including general transcoding, TESHD transcoding, and video editing. This value is not recommended.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeMediaProcessUsageDataRequest struct {
	*tchttp.BaseRequest
	
	// Start date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F). The end date must be on or after the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The type of media processing task. Valid values:
	// <li>Transcoding: General transcoding</li>
	// <li>Transcoding-TESHD: TESHD transcoding</li>
	// <li>Editing: Video editing</li>
	// <li>Editing-TESHD: TESHD editing</li>
	// <li>AdaptiveBitrateStreaming: Adaptive bitrate streaming</li>
	// <li>ContentAudit: Content moderation</li>
	// <li>RemoveWatermark: Watermark removal</li>
	// <li>Transcode: Transcoding, including general transcoding, TESHD transcoding, and video editing. This value is not recommended.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeMediaProcessUsageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaProcessUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaProcessUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaProcessUsageDataResponseParams struct {
	// Overview of video processing statistics, which displays the overview and details of queried tasks.
	MediaProcessDataSet []*TaskStatData `json:"MediaProcessDataSet,omitempty" name:"MediaProcessDataSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMediaProcessUsageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaProcessUsageDataResponseParams `json:"Response"`
}

func (r *DescribeMediaProcessUsageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaProcessUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonSamplesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Type of samples to pull. Valid values:
	// <li>UserDefine: custom sample library</li>
	// <li>Default: default sample library</li>
	// 
	// Default value: UserDefine. Samples in the custom sample library will be pulled.
	// Note: samples from the default library can only be pulled by providing the name or both the ID and name of a sample. Only one face image will be returned.
	Type *string `json:"Type,omitempty" name:"Type"`

	// IDs of samples. Array length limit: 100.
	PersonIds []*string `json:"PersonIds,omitempty" name:"PersonIds"`

	// Names of samples. Array length limit: 20.
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Tags of a sample. Array length limit: 20.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries to be returned. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePersonSamplesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Type of samples to pull. Valid values:
	// <li>UserDefine: custom sample library</li>
	// <li>Default: default sample library</li>
	// 
	// Default value: UserDefine. Samples in the custom sample library will be pulled.
	// Note: samples from the default library can only be pulled by providing the name or both the ID and name of a sample. Only one face image will be returned.
	Type *string `json:"Type,omitempty" name:"Type"`

	// IDs of samples. Array length limit: 100.
	PersonIds []*string `json:"PersonIds,omitempty" name:"PersonIds"`

	// Names of samples. Array length limit: 20.
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Tags of a sample. Array length limit: 20.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries to be returned. Default value: 100. Maximum value: 100.
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
	delete(f, "SubAppId")
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

// Predefined struct for user
type DescribePersonSamplesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Figure information.
	PersonSet []*AiSamplePerson `json:"PersonSet,omitempty" name:"PersonSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePersonSamplesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonSamplesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeProcedureTemplatesRequestParams struct {
	// Name filter of task flow template. Array length limit: 100.
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Filter of task flow template types. Valid values:
	// <li>Preset: preset task flow template;</li>
	// <li>Custom: custom task flow template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeProcedureTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Name filter of task flow template. Array length limit: 100.
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Filter of task flow template types. Valid values:
	// <li>Preset: preset task flow template;</li>
	// <li>Custom: custom task flow template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeProcedureTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcedureTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Names")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcedureTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcedureTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of task flow template details.
	ProcedureTemplateSet []*ProcedureTemplate `json:"ProcedureTemplateSet,omitempty" name:"ProcedureTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProcedureTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProcedureTemplatesResponseParams `json:"Response"`
}

func (r *DescribeProcedureTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcedureTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReviewDetailsRequestParams struct {
	// Start date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). The end date must be after the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeReviewDetailsRequest struct {
	*tchttp.BaseRequest
	
	// Start date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). The end date must be after the start date.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeReviewDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReviewDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReviewDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReviewDetailsResponseParams struct {
	// Times of initiating intelligent content recognition tasks.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Duration of intelligent recognition content.
	TotalDuration *int64 `json:"TotalDuration,omitempty" name:"TotalDuration"`

	// Data of intelligent recognition content duration. One piece of data is collected every day.
	Data []*StatDataItem `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReviewDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReviewDetailsResponseParams `json:"Response"`
}

func (r *DescribeReviewDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReviewDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSampleSnapshotTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of sampled screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeSampleSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of sampled screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSampleSnapshotTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSampleSnapshotTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of sampled screencapturing template details.
	SampleSnapshotTemplateSet []*SampleSnapshotTemplate `json:"SampleSnapshotTemplateSet,omitempty" name:"SampleSnapshotTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSampleSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSampleSnapshotTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSnapshotByTimeOffsetTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of time point screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeSnapshotByTimeOffsetTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of time point screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitempty" name:"Definitions"`

	// Paged offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
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
	delete(f, "SubAppId")
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotByTimeOffsetTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotByTimeOffsetTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of time point screencapturing template details.
	SnapshotByTimeOffsetTemplateSet []*SnapshotByTimeOffsetTemplate `json:"SnapshotByTimeOffsetTemplateSet,omitempty" name:"SnapshotByTimeOffsetTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSnapshotByTimeOffsetTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotByTimeOffsetTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeStorageDataRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeStorageDataRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeStorageDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageDataResponseParams struct {
	// Total number of current media files.
	MediaCount *uint64 `json:"MediaCount,omitempty" name:"MediaCount"`

	// Total current storage capacity in bytes.
	TotalStorage *uint64 `json:"TotalStorage,omitempty" name:"TotalStorage"`

	// Current Standard storage capacity in bytes.
	StandardStorage *uint64 `json:"StandardStorage,omitempty" name:"StandardStorage"`

	// Current Standard_IA storage capacity in bytes.
	InfrequentStorage *uint64 `json:"InfrequentStorage,omitempty" name:"InfrequentStorage"`

	// The current ARCHIVE storage usage in bytes.
	ArchiveStorage *uint64 `json:"ArchiveStorage,omitempty" name:"ArchiveStorage"`

	// The current DEEP ARCHIVE storage usage in bytes.
	DeepArchiveStorage *uint64 `json:"DeepArchiveStorage,omitempty" name:"DeepArchiveStorage"`

	// Storage usage by billing region.
	StorageStat []*StorageStatData `json:"StorageStat,omitempty" name:"StorageStat"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStorageDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageDataResponseParams `json:"Response"`
}

func (r *DescribeStorageDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageDetailsRequestParams struct {
	// Start time in ISO 8601 format. For more information, please see [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in ISO 8601 format, which should be larger than the start time. For more information, please see [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	// You can set this parameter to 1 to query the total usage of all applications (including the primary application) as an admin.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Time granularity. Valid values:
	// <li>Minute: 5-minute granularity</li>
	// <li>Day: 1-day granularity</li>
	// The value is set according to query period length by default. 5-minute granularity is set for periods no longer than 1 day, and 1-day granularity is set for periods longer than 1 day.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Storage class to query. Valid values:
	// <li>`TotalStorage`: total storage usage in classes of STANDARD, STANDARD_IA, ARCHIVE, and DEEP ARCHIVE, excluding the storage usage for data deleted in advance.</li>
	// <li>`StandardStorage`: STANDARD</li>
	// <li>`InfrequentStorage`: STANDARD_IA</li>
	// <li>`ArchiveStorage`: ARCHIVE</li>
	// <li>`DeepArchiveStorage`: DEEP ARCHIVE</li>
	// <li>`DeletedInfrequentStorage`: STANDARD_IA data deleted in advance</li>
	// <li>`DeletedArchiveStorage`: ARCHIVE data deleted in advance</li>
	// <li>`DeletedDeepArchiveStorage`: DEEP ARCHIVE data deleted in advance</li>
	// <li>`ArchiveStandardRetrieval`: ARCHIVE data retrieved using standard retrievals</li>
	// <li>`ArchiveExpeditedRetrieval`: ARCHIVE data retrieved using expedited retrievals</li>
	// <li>`ArchiveBulkRetrieval`: ARCHIVE data retrieved using bulk retrievals</li>
	// <li>`DeepArchiveStandardRetrieval`: DEEP ARCHIVE data retrieved using standard retrievals</li>
	// <li>`DeepArchiveBulkRetrieval`: DEEP ARCHIVE data retrieved using bulk retrievals</li>
	// Default value: `TotalStorage`
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// Storage region to query. Valid values:
	// <li>Chinese Mainland</li>
	// <li>Outside Chinese Mainland</li>
	// Default value: Chinese Mainland
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeStorageDetailsRequest struct {
	*tchttp.BaseRequest
	
	// Start time in ISO 8601 format. For more information, please see [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in ISO 8601 format, which should be larger than the start time. For more information, please see [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	// You can set this parameter to 1 to query the total usage of all applications (including the primary application) as an admin.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Time granularity. Valid values:
	// <li>Minute: 5-minute granularity</li>
	// <li>Day: 1-day granularity</li>
	// The value is set according to query period length by default. 5-minute granularity is set for periods no longer than 1 day, and 1-day granularity is set for periods longer than 1 day.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Storage class to query. Valid values:
	// <li>`TotalStorage`: total storage usage in classes of STANDARD, STANDARD_IA, ARCHIVE, and DEEP ARCHIVE, excluding the storage usage for data deleted in advance.</li>
	// <li>`StandardStorage`: STANDARD</li>
	// <li>`InfrequentStorage`: STANDARD_IA</li>
	// <li>`ArchiveStorage`: ARCHIVE</li>
	// <li>`DeepArchiveStorage`: DEEP ARCHIVE</li>
	// <li>`DeletedInfrequentStorage`: STANDARD_IA data deleted in advance</li>
	// <li>`DeletedArchiveStorage`: ARCHIVE data deleted in advance</li>
	// <li>`DeletedDeepArchiveStorage`: DEEP ARCHIVE data deleted in advance</li>
	// <li>`ArchiveStandardRetrieval`: ARCHIVE data retrieved using standard retrievals</li>
	// <li>`ArchiveExpeditedRetrieval`: ARCHIVE data retrieved using expedited retrievals</li>
	// <li>`ArchiveBulkRetrieval`: ARCHIVE data retrieved using bulk retrievals</li>
	// <li>`DeepArchiveStandardRetrieval`: DEEP ARCHIVE data retrieved using standard retrievals</li>
	// <li>`DeepArchiveBulkRetrieval`: DEEP ARCHIVE data retrieved using bulk retrievals</li>
	// Default value: `TotalStorage`
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// Storage region to query. Valid values:
	// <li>Chinese Mainland</li>
	// <li>Outside Chinese Mainland</li>
	// Default value: Chinese Mainland
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeStorageDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "Interval")
	delete(f, "StorageType")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageDetailsResponseParams struct {
	// Storage statistics with one piece of data for every 5 minutes or 1 day.
	Data []*StatDataItem `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStorageDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageDetailsResponseParams `json:"Response"`
}

func (r *DescribeStorageDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageRegionsRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeStorageRegionsRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeStorageRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageRegionsResponseParams struct {
	// The information of the storage regions.
	StorageRegionInfos []*StorageRegionInfo `json:"StorageRegionInfos,omitempty" name:"StorageRegionInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStorageRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageRegionsResponseParams `json:"Response"`
}

func (r *DescribeStorageRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAppIdsRequestParams struct {
	// Subapplication name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Tag information. You can query the list of subapplications with specified tags.
	Tags []*ResourceTag `json:"Tags,omitempty" name:"Tags"`

	// Page number offset from the beginning of paginated queries. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum return results of pulling paginated queries. Default: 200; maximum: 200.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSubAppIdsRequest struct {
	*tchttp.BaseRequest
	
	// Subapplication name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Tag information. You can query the list of subapplications with specified tags.
	Tags []*ResourceTag `json:"Tags,omitempty" name:"Tags"`

	// Page number offset from the beginning of paginated queries. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum return results of pulling paginated queries. Default: 200; maximum: 200.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubAppIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAppIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Tags")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubAppIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAppIdsResponseParams struct {
	// Subapplication information set.
	SubAppIdInfoSet []*SubAppIdInfo `json:"SubAppIdInfoSet,omitempty" name:"SubAppIdInfoSet"`

	// Total number of subapplications.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubAppIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubAppIdsResponseParams `json:"Response"`
}

func (r *DescribeSubAppIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAppIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSuperPlayerConfigsRequestParams struct {
	// Player configuration name filter. Array length limit: 100.
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Player configuration type filter. Valid values:
	// <li>Preset: preset configuration;</li>
	// <li>Custom: custom configuration.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeSuperPlayerConfigsRequest struct {
	*tchttp.BaseRequest
	
	// Player configuration name filter. Array length limit: 100.
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Player configuration type filter. Valid values:
	// <li>Preset: preset configuration;</li>
	// <li>Custom: custom configuration.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeSuperPlayerConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSuperPlayerConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Names")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSuperPlayerConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSuperPlayerConfigsResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Player configuration array.
	PlayerConfigSet []*PlayerConfig `json:"PlayerConfigSet,omitempty" name:"PlayerConfigSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSuperPlayerConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSuperPlayerConfigsResponseParams `json:"Response"`
}

func (r *DescribeSuperPlayerConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSuperPlayerConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
	// The task type. Valid values:
	// <li>Procedure: Video processing</li>
	// <li>EditMedia: Video editing</li>
	// <li>SplitMedia: Video splitting</li>
	// <li>ComposeMedia: Media file producing</li>
	// <li>WechatPublish: WeChat publishing</li>
	// <li>WechatMiniProgramPublish: Publishing videos on WeChat Mini Program</li>
	// <li>PullUpload: Pulling media files for upload</li>
	// <li>FastClipMedia: Quick clipping</li>
	// 
	// Task types for v2017:
	// <li>Transcode: Transcoding</li>
	// <li>SnapshotByTimeOffset: Screencapturing</li>
	// <li>Concat: Video splicing</li>
	// <li>Clip: Video clipping</li>
	// <li>ImageSprites: Image sprite generating</li>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Task status. Valid values:
	// <li>WAITING: waiting;</li>
	// <li>PROCESSING: processing;</li>
	// <li>FINISH: completed.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Creation time of task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Start time of task execution in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

	// End time of task execution in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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

	// Video splitting task information. This field has a value only when `TaskType` is `EditMedia`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SplitMediaTask *SplitMediaTask `json:"SplitMediaTask,omitempty" name:"SplitMediaTask"`

	// Release on WeChat Mini Program task information. This field has a value only when `TaskType` is `WechatMiniProgramPublish`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WechatMiniProgramPublishTask *WechatMiniProgramPublishTask `json:"WechatMiniProgramPublishTask,omitempty" name:"WechatMiniProgramPublishTask"`

	// Media file pulling for upload task information. This field has a value only when `TaskType` is `PullUpload`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PullUploadTask *PullUploadTask `json:"PullUploadTask,omitempty" name:"PullUploadTask"`

	// Video transcoding task information. This field has a value only when `TaskType` is `Transcode`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TranscodeTask *TranscodeTask2017 `json:"TranscodeTask,omitempty" name:"TranscodeTask"`

	// Video splicing task information. This field has a value only when `TaskType` is `Concat`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConcatTask *ConcatTask2017 `json:"ConcatTask,omitempty" name:"ConcatTask"`

	// Video clipping task information. This field has a value only when `TaskType` is `Clip`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClipTask *ClipTask2017 `json:"ClipTask,omitempty" name:"ClipTask"`

	// Image sprite creating task information. This field has a value only when `TaskType` is `ImageSprite`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateImageSpriteTask *CreateImageSpriteTask2017 `json:"CreateImageSpriteTask,omitempty" name:"CreateImageSpriteTask"`

	// Time point screencapturing task information. This field has a value only when `TaskType` is `SnapshotByTimeOffset`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetTask *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetTask,omitempty" name:"SnapshotByTimeOffsetTask"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Filter: Task status. Valid values: WAITING (waiting), PROCESSING (processing), FINISH (completed).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Filter: file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Filter: task creation time.
	CreateTime *TimeRange `json:"CreateTime,omitempty" name:"CreateTime"`

	// Filter: task end time.
	FinishTime *TimeRange `json:"FinishTime,omitempty" name:"FinishTime"`

	// Sort field. Valid values:
	// <li>`CreateTime`: task creation time</li>
	// <li>`FinishTime`: task end time</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// Number of entries to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Scrolling identifier which is used for pulling in batches. If a single request cannot pull all the data entries, the API will return `ScrollToken`, and if the next request carries it, the next pull will start from the next entry.
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Filter: Task status. Valid values: WAITING (waiting), PROCESSING (processing), FINISH (completed).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Filter: file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Filter: task creation time.
	CreateTime *TimeRange `json:"CreateTime,omitempty" name:"CreateTime"`

	// Filter: task end time.
	FinishTime *TimeRange `json:"FinishTime,omitempty" name:"FinishTime"`

	// Sort field. Valid values:
	// <li>`CreateTime`: task creation time</li>
	// <li>`FinishTime`: task end time</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// Number of entries to be returned. Default value: 10. Maximum value: 100.
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
	delete(f, "SubAppId")
	delete(f, "Status")
	delete(f, "FileId")
	delete(f, "CreateTime")
	delete(f, "FinishTime")
	delete(f, "Sort")
	delete(f, "Limit")
	delete(f, "ScrollToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// Task overview list.
	TaskSet []*TaskSimpleInfo `json:"TaskSet,omitempty" name:"TaskSet"`

	// Scrolling identifier. If a request does not return all the data entries, this field indicates the ID of the next entry. If this field is empty, there is no more data.
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTranscodeTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of transcoding templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

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
}

type DescribeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Unique ID filter of transcoding templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type DescribeTranscodeTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of transcoding template details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TranscodeTemplateSet []*TranscodeTemplate `json:"TranscodeTemplateSet,omitempty" name:"TranscodeTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTranscodeTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeVodDomainsRequestParams struct {
	// List of domain names. If this parameter is left empty, all domain names will be listed.
	// <li>Maximum number of domain names listed: 20</li>
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Maximum results to return for pulling paginated queries. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number offset from the beginning of paginated queries. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DescribeVodDomainsRequest struct {
	*tchttp.BaseRequest
	
	// List of domain names. If this parameter is left empty, all domain names will be listed.
	// <li>Maximum number of domain names listed: 20</li>
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Maximum results to return for pulling paginated queries. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number offset from the beginning of paginated queries. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeVodDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVodDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVodDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVodDomainsResponseParams struct {
	// Total number of domain names
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Domain name information list
	DomainSet []*DomainDetailInfo `json:"DomainSet,omitempty" name:"DomainSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVodDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVodDomainsResponseParams `json:"Response"`
}

func (r *DescribeVodDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVodDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWatermarkTemplatesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Watermark type filter. Valid values:
	// <li>image: image watermark;</li>
	// <li>text: text watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Unique ID filter of watermarking templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

	// Number of returned entries
	// <li>Default value: 10;</li>
	// <li>Maximum value: 100.</li>
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeWatermarkTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Watermark type filter. Valid values:
	// <li>image: image watermark;</li>
	// <li>text: text watermark.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Unique ID filter of watermarking templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions"`

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
	delete(f, "SubAppId")
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Definitions")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWatermarkTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWatermarkTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of watermarking template details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkTemplateSet []*WatermarkTemplate `json:"WatermarkTemplateSet,omitempty" name:"WatermarkTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWatermarkTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWatermarkTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeWordSamplesRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// You can select multiple elements, which are connected by OR logic. If a use case contains any element in this parameter, the keyword sample will be used.
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Keyword filter. Array length limit: 100 words.
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// Tag filter. Array length limit: 20 words.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries to be returned. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeWordSamplesRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// You can select multiple elements, which are connected by OR logic. If a use case contains any element in this parameter, the keyword sample will be used.
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Keyword filter. Array length limit: 100 words.
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// Tag filter. Array length limit: 20 words.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries to be returned. Default value: 100. Maximum value: 100.
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
	delete(f, "SubAppId")
	delete(f, "Usages")
	delete(f, "Keywords")
	delete(f, "Tags")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWordSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWordSamplesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Keyword information.
	WordSet []*AiSampleWord `json:"WordSet,omitempty" name:"WordSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWordSamplesResponseParams `json:"Response"`
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

type DomainDetailInfo struct {
	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Acceleration region information
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AccelerateAreaInfos []*AccelerateAreaInfo `json:"AccelerateAreaInfos,omitempty" name:"AccelerateAreaInfos"`

	// Deployment status. Valid values:
	// <li>Online</li>
	// <li>Deploying</li>
	// <li>Locked: you cannot change the deployment status of locked domain names</li>
	DeployStatus *string `json:"DeployStatus,omitempty" name:"DeployStatus"`

	// HTTPS configuration information
	// Note: this field may return `null`, indicating that no valid value is obtained.
	HTTPSConfig *DomainHTTPSConfig `json:"HTTPSConfig,omitempty" name:"HTTPSConfig"`

	// [Key hotlink protection](https://intl.cloud.tencent.com/document/product/266/33986) configuration
	// Note: this field may return `null`, indicating that no valid value is obtained.
	UrlSignatureAuthPolicy *UrlSignatureAuthPolicy `json:"UrlSignatureAuthPolicy,omitempty" name:"UrlSignatureAuthPolicy"`

	// [Referer hotlink protection](https://intl.cloud.tencent.com/document/product/266/33985) configuration
	// Note: this field may return `null`, indicating that no valid value is obtained.
	RefererAuthPolicy *RefererAuthPolicy `json:"RefererAuthPolicy,omitempty" name:"RefererAuthPolicy"`

	// The time when the domain name was added in the VOD system
	// <li>The time is in [ISO 8601 date format](https://intl.cloud.tencent.com/document/product/266/11732).</li>
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DomainHTTPSConfig struct {
	// Time when the certificate expires
	// <li>The time is in [ISO 8601 date format](https://intl.cloud.tencent.com/document/product/266/11732).</li>
	CertExpireTime *string `json:"CertExpireTime,omitempty" name:"CertExpireTime"`
}

type DrmStreamingsInfo struct {
	// ID of the adaptive bitrate streaming template whose protection type is SimpleAES.
	SimpleAesDefinition *uint64 `json:"SimpleAesDefinition,omitempty" name:"SimpleAesDefinition"`

	// The ID of the adaptive bitrate streaming template that encrypts the streams by Widewine.
	WidevineDefinition *uint64 `json:"WidevineDefinition,omitempty" name:"WidevineDefinition"`

	// The ID of the adaptive bitrate streaming template that encrypts the streams by FairPlay.
	FairPlayDefinition *uint64 `json:"FairPlayDefinition,omitempty" name:"FairPlayDefinition"`
}

type DrmStreamingsInfoForUpdate struct {
	// ID of the adaptive bitrate streaming template whose protection type is SimpleAES.
	SimpleAesDefinition *uint64 `json:"SimpleAesDefinition,omitempty" name:"SimpleAesDefinition"`

	// The ID of the adaptive bitrate streaming template that encrypts the streams by Widewine.
	WidevineDefinition *uint64 `json:"WidevineDefinition,omitempty" name:"WidevineDefinition"`

	// The ID of the adaptive bitrate streaming template that encrypts the streams by FairPlay.
	FairPlayDefinition *uint64 `json:"FairPlayDefinition,omitempty" name:"FairPlayDefinition"`
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

type EditMediaStreamInfo struct {
	// ID of recorded stream
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Start time of stream clipping in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of stream clipping in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Progress of a video editing task. Value range: [0, 100]
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Input of video editing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Input *EditMediaTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of video editing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *EditMediaTaskOutput `json:"Output,omitempty" name:"Output"`

	// The metadata of the output video.
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// If a video processing flow is specified when a video editing task is initiated, this field will be the ID of the task flow.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

type EditMediaTaskInput struct {
	// Input video source type. Valid values: File, Stream.
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// Information of input video file. This field has a value only when `InputType` is `File`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileInfoSet []*EditMediaFileInfo `json:"FileInfoSet,omitempty" name:"FileInfoSet"`

	// Input stream information. This field has a value only when `InputType` is `Stream`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StreamInfoSet []*EditMediaStreamInfo `json:"StreamInfoSet,omitempty" name:"StreamInfoSet"`
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

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [category creating](https://intl.cloud.tencent.com/document/product/266/7812?from_cn_redirect=1) API.
	// <li>Default value: 0, which means "Other".</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Expiration time of output media file in ISO 8601 format, after which the file will be deleted. Files will never expire by default. For more information, please see [Notes on ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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

	// <b>Supported event types:</b>
	// <li>NewFileUpload: finished video upload</li>
	// <li>ProcedureStateChanged: task flow status changed</li>
	// <li>FileDeleted: finished video deletion</li>
	// <li>PullComplete: finished pulling for upload</li>
	// <li>EditMediaComplete: finished video editing</li>
	// <li>SplitMediaComplete: finished video splitting</li>
	// <li>WechatPublishComplete: finished publishing on WeChat</li>
	// <li>ComposeMediaComplete: finished producing the media file</li>
	// <li>WechatMiniProgramPublishComplete: finished publishing on WeChat Mini Program</li>
	// <b>Support v2017 task types:</b>
	// <li>TranscodeComplete: finished video transcoding</li>
	// <li>ConcatComplete: finished video splicing</li>
	// <li>ClipComplete: finished video clipping</li>
	// <li>CreateImageSpriteComplete: finished image sprite generation</li>
	// <li>CreateSnapshotByTimeOffsetComplete: finished point-in-time screencapturing</li>
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

	// Video splitting completion event, which is valid if the event type is `EditMediaComplete`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SplitMediaCompleteEvent *SplitMediaTask `json:"SplitMediaCompleteEvent,omitempty" name:"SplitMediaCompleteEvent"`

	// Media file composing task completion event, which is valid when the event type is `ComposeMediaComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ComposeMediaCompleteEvent *ComposeMediaTask `json:"ComposeMediaCompleteEvent,omitempty" name:"ComposeMediaCompleteEvent"`

	// Video clipping completion event, which is valid if the event type is `ClipComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClipCompleteEvent *ClipTask2017 `json:"ClipCompleteEvent,omitempty" name:"ClipCompleteEvent"`

	// Video transcoding completion event, which is valid if the event type is `TranscodeComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TranscodeCompleteEvent *TranscodeTask2017 `json:"TranscodeCompleteEvent,omitempty" name:"TranscodeCompleteEvent"`

	// Image sprite generating completion event, which is valid if the event type is `CreateImageSpriteComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateImageSpriteCompleteEvent *CreateImageSpriteTask2017 `json:"CreateImageSpriteCompleteEvent,omitempty" name:"CreateImageSpriteCompleteEvent"`

	// Video splicing completion event, which is valid if the event type is `ConcatComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConcatCompleteEvent *ConcatTask2017 `json:"ConcatCompleteEvent,omitempty" name:"ConcatCompleteEvent"`

	// Time point screencapturing completion event, which is valid when the event type is `CreateSnapshotByTimeOffsetComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetCompleteEvent *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetCompleteEvent,omitempty" name:"SnapshotByTimeOffsetCompleteEvent"`

	// Release on WeChat completion event, which is valid if the event type is `WechatPublishComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WechatPublishCompleteEvent *WechatPublishTask `json:"WechatPublishCompleteEvent,omitempty" name:"WechatPublishCompleteEvent"`

	// Release on WeChat Mini Program task completion event, which is valid if the event type is `WechatMiniProgramPublishComplete`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WechatMiniProgramPublishCompleteEvent *WechatMiniProgramPublishTask `json:"WechatMiniProgramPublishCompleteEvent,omitempty" name:"WechatMiniProgramPublishCompleteEvent"`

	// Callback for video retrieval. This parameter is valid when the event type is `RestoreMediaComplete`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RestoreMediaCompleteEvent *RestoreMediaTask `json:"RestoreMediaCompleteEvent,omitempty" name:"RestoreMediaCompleteEvent"`
}

// Predefined struct for user
type ExecuteFunctionRequestParams struct {
	// Name of called backend API.
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// API parameter. For specific parameter format, negotiate with the backend before calling.
	FunctionArg *string `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

type ExecuteFunctionRequest struct {
	*tchttp.BaseRequest
	
	// Name of called backend API.
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// API parameter. For specific parameter format, negotiate with the backend before calling.
	FunctionArg *string `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
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
	delete(f, "SubAppId")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteFunctionResponseParams struct {
	// String generated by packaging processing result. For specifications, negotiate with the backend.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExecuteFunctionResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteFunctionResponseParams `json:"Response"`
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
	// Switch of face recognition task. Valid values:
	// <li>ON: enables intelligent face recognition task;</li>
	// <li>OFF: disables intelligent face recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Face recognition filter score. If this score is reached or exceeded, a recognition result will be returned. Value range: 0–100. Default value: 95.
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// Default face filter labels, which specify the types of faces to return. If this parameter is left empty, the recognition results for all labels are returned. Valid values:
	// <li>`entertainment`: people in the entertainment industry</li>
	// <li>`sport`: sports celebrities</li>
	// <li>`politician`: politically sensitive people</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitempty" name:"DefaultLibraryLabelSet"`

	// Custom face labels for filtering. After you specify a label, callbacks of face images without this label will be returned. If this parameter is not specified or left empty, callbacks of all face images will be returned.
	// You can specify up to 100 labels, with each containing up to 16 characters.
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitempty" name:"UserDefineLibraryLabelSet"`

	// Figure library. Valid values:
	// <li>Default: default figure library;</li>
	// <li>UserDefine: custom figure library.</li>
	// <li>All: both default and custom figure libraries will be used.</li>
	// Default value: All (both default and custom figure libraries will be used.)
	FaceLibrary *string `json:"FaceLibrary,omitempty" name:"FaceLibrary"`
}

type FaceConfigureInfoForUpdate struct {
	// Switch of face recognition task. Valid values:
	// <li>ON: enables intelligent face recognition task;</li>
	// <li>OFF: disables intelligent face recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Face recognition filter score. If this score is reached or exceeded, a recognition result will be returned. Value range: 0–100.
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// Default face filter labels, which specify the types of faces to return. If this parameter is left empty or an empty value is entered, the recognition results for all labels are returned. Valid values:
	// <li>`entertainment`: people in the entertainment industry</li>
	// <li>`sport`: sports celebrities</li>
	// <li>`politician`: politically sensitive people</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitempty" name:"DefaultLibraryLabelSet"`

	// Custom face labels for filtering. After you specify a label, callbacks of face images without this label will be returned. If this parameter is not specified or left empty, callbacks of all face images will be returned.
	// You can specify up to 100 labels, with each containing up to 16 characters.
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitempty" name:"UserDefineLibraryLabelSet"`

	// Figure library. Valid values:
	// <li>Default: default figure library;</li>
	// <li>UserDefine: custom figure library.</li>
	// <li>All: both default and custom figure libraries will be used.</li>
	FaceLibrary *string `json:"FaceLibrary,omitempty" name:"FaceLibrary"`
}

type FileDeleteResultItem struct {
	// The ID of the file deleted.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// The type of the file deleted.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DeleteParts []*MediaDeleteItem `json:"DeleteParts,omitempty" name:"DeleteParts"`
}

type FileDeleteTask struct {
	// List of IDs of deleted files.
	FileIdSet []*string `json:"FileIdSet,omitempty" name:"FileIdSet"`

	// The information of the files deleted.
	FileDeleteResultInfo []*FileDeleteResultItem `json:"FileDeleteResultInfo,omitempty" name:"FileDeleteResultInfo"`
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

// Predefined struct for user
type ForbidMediaDistributionRequestParams struct {
	// List of media files. Up to 20 ones can be submitted at a time.
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// forbid: forbids, recover: unblocks.
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type ForbidMediaDistributionRequest struct {
	*tchttp.BaseRequest
	
	// List of media files. Up to 20 ones can be submitted at a time.
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// forbid: forbids, recover: unblocks.
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ForbidMediaDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForbidMediaDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileIds")
	delete(f, "Operation")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForbidMediaDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ForbidMediaDistributionResponseParams struct {
	// List of IDs of files that do not exist.
	NotExistFileIdSet []*string `json:"NotExistFileIdSet,omitempty" name:"NotExistFileIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ForbidMediaDistributionResponse struct {
	*tchttp.BaseResponse
	Response *ForbidMediaDistributionResponseParams `json:"Response"`
}

func (r *ForbidMediaDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

type HeadTailConfigureInfoForUpdate struct {
	// Switch of video opening and ending credits recognition task. Valid values:
	// <li>ON: enables video opening and ending credits recognition task;</li>
	// <li>OFF: disables video opening and ending credits recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type HeadTailTaskInput struct {
	// Video opening/closing credits configuration template ID
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
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

type ImageReviewUsageDataItem struct {
	// The start time (in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)) of the data returned. For example, if the granularity is a day, `2018-12-01T00:00:00+08:00` indicates that the data is for the whole day of December 1, 2018.
	Time *string `json:"Time,omitempty" name:"Time"`

	// The number of times the image recognition feature is used.
	Count *int64 `json:"Count,omitempty" name:"Count"`
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

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// Template description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type ImageTransform struct {
	// Type. Valid values:
	// <li> Rotate: image rotation.</li>
	// <li> Flip: image flipping.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Rotation angle of image with its center point as origin. Value range: 0-360. This parameter is valid if `Type` is `Rotate`.
	RotateAngle *float64 `json:"RotateAngle,omitempty" name:"RotateAngle"`

	// Image flipping action. Valid values:
	// <li>Horizental: horizontal flipping, i.e., horizontally mirrored.</li>
	// <li>Vertical: vertical flipping, i.e., vertically mirrored.</li>
	// This is valid if `Type` is `Flip`.
	Flip *string `json:"Flip,omitempty" name:"Flip"`
}

type ImageWatermarkInput struct {
	// The [Base64](https://tools.ietf.org/html/rfc4648) encoded string of a watermark image. Only JPEG, PNG, and GIF images are supported.
	ImageContent *string `json:"ImageContent,omitempty" name:"ImageContent"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width. For example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in pixels. For example, `100px` means that `Width` is 100 pixels. Value range: [8, 4096].</li>
	// Default value: 10%.
	Width *string `json:"Width,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px. Valid values: 0 or [8,4096].</li>
	// Default value: 0 px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
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
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px. Valid values: 0 or [8,4096].</li>
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
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px;</li>
	// `0px` means that `Height` will be proportionally scaled according to the video width.
	Height *string `json:"Height,omitempty" name:"Height"`

	// Repeat type of an animated watermark. Valid values:
	// <li>once: no longer appears after watermark playback ends.</li>
	// <li>repeat_last_frame: stays on the last frame after watermark playback ends.</li>
	// <li>repeat (default): repeats the playback until the video ends.</li>
	RepeatType *string `json:"RepeatType,omitempty" name:"RepeatType"`
}

type LicenseUsageDataItem struct {
	// The start time (in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732#iso-date-format)) of the data returned. For example, if the granularity is a day, `2018-12-01T00:00:00+08:00` indicates that the data is for the whole day of December 1, 2018.
	Time *string `json:"Time,omitempty" name:"Time"`

	// The number of license requests.
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

// Predefined struct for user
type LiveRealTimeClipRequestParams struct {
	// The live stream code.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Start time of stream clipping in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of stream clipping in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Whether to clip persistently. 0: no, 1: yes. Default: no.
	IsPersistence *int64 `json:"IsPersistence,omitempty" name:"IsPersistence"`

	// Storage expiration time of video generated by persistent clipping in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). `9999-12-31T23:59:59Z` means `never expire`. After the expiration, the media file and its related resources (such as transcoding results and image sprites) will be permanently deleted. This parameter will be valid only when `IsPersistence` is 1. By default, the video will never expire.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// VOD task flow processing for video generated by persistent clipping. For more information, please see [Specifying Task Flow After Upload](https://intl.cloud.tencent.com/document/product/266/9759?from_cn_redirect=1). This parameter will be valid only when `IsPersistence` is 1.
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Whether the metadata of clipped video needs to be returned. 0: no, 1: yes. Default value: no.
	MetaDataRequired *uint64 `json:"MetaDataRequired,omitempty" name:"MetaDataRequired"`

	// Domain name used for live clipping. Time shifting must be enabled in LVB.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Reserved field. Do not enter a value for it.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

type LiveRealTimeClipRequest struct {
	*tchttp.BaseRequest
	
	// The live stream code.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Start time of stream clipping in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of stream clipping in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Whether to clip persistently. 0: no, 1: yes. Default: no.
	IsPersistence *int64 `json:"IsPersistence,omitempty" name:"IsPersistence"`

	// Storage expiration time of video generated by persistent clipping in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). `9999-12-31T23:59:59Z` means `never expire`. After the expiration, the media file and its related resources (such as transcoding results and image sprites) will be permanently deleted. This parameter will be valid only when `IsPersistence` is 1. By default, the video will never expire.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// VOD task flow processing for video generated by persistent clipping. For more information, please see [Specifying Task Flow After Upload](https://intl.cloud.tencent.com/document/product/266/9759?from_cn_redirect=1). This parameter will be valid only when `IsPersistence` is 1.
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Whether the metadata of clipped video needs to be returned. 0: no, 1: yes. Default value: no.
	MetaDataRequired *uint64 `json:"MetaDataRequired,omitempty" name:"MetaDataRequired"`

	// Domain name used for live clipping. Time shifting must be enabled in LVB.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Reserved field. Do not enter a value for it.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

func (r *LiveRealTimeClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LiveRealTimeClipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SubAppId")
	delete(f, "IsPersistence")
	delete(f, "ExpireTime")
	delete(f, "Procedure")
	delete(f, "MetaDataRequired")
	delete(f, "Host")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LiveRealTimeClipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LiveRealTimeClipResponseParams struct {
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
}

type LiveRealTimeClipResponse struct {
	*tchttp.BaseResponse
	Response *LiveRealTimeClipResponseParams `json:"Response"`
}

func (r *LiveRealTimeClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LiveRealTimeClipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageTaskRequestParams struct {
	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Operation type. Valid value:
	// <li>Abort: terminate a task. You can only terminate initiated tasks in `WAITING` status.</li>
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type ManageTaskRequest struct {
	*tchttp.BaseRequest
	
	// Video processing task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Operation type. Valid value:
	// <li>Abort: terminate a task. You can only terminate initiated tasks in `WAITING` status.</li>
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
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
	delete(f, "TaskId")
	delete(f, "OperationType")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ManageTaskResponse struct {
	*tchttp.BaseResponse
	Response *ManageTaskResponseParams `json:"Response"`
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

type MediaAdaptiveDynamicStreamingInfo struct {
	// Information array of adaptive bitrate streaming.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdaptiveDynamicStreamingSet []*AdaptiveDynamicStreamingInfoItem `json:"AdaptiveDynamicStreamingSet,omitempty" name:"AdaptiveDynamicStreamingSet"`
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

	// Category list of frame-specific tag names. `CategorySet.N` indicates the N+1-level category.
	// For example, if the `Tag` is "tower", and `CategorySet` contains two elements (`CategorySet.0` is "scene", and `CategorySet.1` is "architecture"), then the frame-specific tag is "tower", the first-level category is "scene", and the second-level category is "architecture".
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
	SegmentSet []*HighlightSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet"`
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
	AnimatedGraphicsSet []*MediaAnimatedGraphicsItem `json:"AnimatedGraphicsSet,omitempty" name:"AnimatedGraphicsSet"`
}

type MediaAnimatedGraphicsItem struct {
	// Address of generated animated image.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Animated image generating template ID. For more information, please see [Animated Image Generating Parameter Template](https://intl.cloud.tencent.com/document/product/266/33481?from_cn_redirect=1#.3Cspan-id-.3D-.22zdt.22.3E.3C.2Fspan.3E.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF).
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

	// Creation time of media file in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last update time of media file (by an operation that triggers updating of media file information such as modifying video attributes or initiating video processing) in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Expiration time of media file in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). After the expiration, the media file and its related resources (such as transcoding results and image sprites) will be permanently deleted. `9999-12-31T23:59:59Z` means "never expire".
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

	// Regions where media files are stored, such as `ap-chongqing`. For more regions, see [Storage Region](https://intl.cloud.tencent.com/document/product/266/9760).
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// Tag information of media file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet"`

	// Unique ID of an LVB recording file.
	Vid *string `json:"Vid,omitempty" name:"Vid"`

	// File type:
	// <li>Video: video file</li>
	// <li>Audio: audio file</li>
	// <li>Image: image file</li>
	Category *string `json:"Category,omitempty" name:"Category"`

	// File status. Valid values: Normal, Forbidden.
	// 
	// *Note: this field is not supported yet.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Storage class of a media file:
	// <li>STANDARD</li>
	// <li>STANDARD_IA</li>
	// <li>ARCHIVE</li>
	// <li>DEEP_ARCHIVE</li>
	StorageClass *string `json:"StorageClass,omitempty" name:"StorageClass"`
}

type MediaClassInfo struct {
	// Category ID
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Parent category ID, which is -1 for a first-level category.
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// Category name
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// Category level. 0 for first-level category, up to 3, i.e., up to 4 levels of categories are allowed.
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// Set of IDs of the immediate subcategories in current category
	SubClassIdSet []*int64 `json:"SubClassIdSet,omitempty" name:"SubClassIdSet"`
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

	// Processing suggestion for the detected suspicious content. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of suspected keywords.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet"`
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

	// Processing suggestion for the detected suspicious content. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// List of suspected keywords.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet"`

	// Zone coordinates (at the pixel level) of suspected text: [x1, y1, x2, y2], i.e., the coordinates of the top-left and bottom-right corners.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// Expiration time of suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type MediaContentReviewPoliticalSegmentItem struct {
	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence score for the detected politically sensitive content
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Processing suggestion for the detected politically sensitive content. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Name of the politically sensitive content or banned images
	Name *string `json:"Name,omitempty" name:"Name"`

	// Labels for the detected politically sensitive content. The relationship between the values of this parameter and those of the `LabelSet` parameter in [PoliticalImgReviewTemplateInfo](https://intl.cloud.tencent.com/document/api/266/31773?from_cn_redirect=1#PoliticalImgReviewTemplateInfo) is as follows:
	// violation_photo:
	// <li>`violation_photo`: banned images</li>
	// politician:
	// <li>`nation_politician`: state leader of China</li>
	// <li>`province_politician`: provincial officials</li>
	// <li>`bureau_politician`: bureau-level officials</li>
	// <li>`county_politician`: county-level officials</li>
	// <li>`rural_politician`: township-level officials</li>
	// <li>`sensitive_politician`: politically sensitive people</li>
	// <li>`foreign_politician`: state leaders of other countries</li>
	// entertainment:
	// <li>`sensitive_entertainment`: banned people in the entertainment industry</li>
	// sport:
	// <li>`sensitive_sport`: banned sports celebrities</li>
	// entrepreneur:
	// <li>`sensitive_entrepreneur`: banned businesspeople</li>
	// scholar:
	// <li>sensitive_scholar: banned scholars</li>
	// celebrity:
	// <li>sensitive_celebrity: banned celebrities</li>
	// <li>historical_celebrity: banned historical figures</li>
	// military:
	// <li>sensitive_military: banned people in military</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// URL of a suspected image (which will not be permanently stored
	//  and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// Coordinates (pixel) of the detected politically sensitive content or banned icons. The format is [x1, y1, x2, y2], which indicates the coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet"`

	// This field has been disused. Please use `PicUrlExpireTime`.
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`

	// Expiration time of suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type MediaContentReviewSegmentItem struct {
	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Confidence score for the detected pornographic content
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Label for the detected pornographic content
	Label *string `json:"Label,omitempty" name:"Label"`

	// Processing suggestion for the detected pornographic content. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// URL of a suspected image (which will not be permanently stored
	//  and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitempty" name:"Url"`

	// This field has been disused. Please use `PicUrlExpireTime`.
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`

	// Expiration time of suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitempty" name:"PicUrlExpireTime"`
}

type MediaDeleteItem struct {
	// Type of files to delete. If this parameter is left empty, it will be invalid. Valid values:
	// <li>`OriginalFiles`: original files. You cannot initiate transcoding, publishing on WeChat, or other video processing operations after deleting the original files.</li>
	// <li>`TranscodeFiles`: transcoded files</li>
	// <li>`WechatPublishFiles`: files for publishing on WeChat</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of the template for which to delete the videos of the type specified by the `Type` parameter. For the template definition, please see [Transcoding Template](https://intl.cloud.tencent.com/document/product/266/33478?from_cn_redirect=1#.3Cspan-id-.3D-.22zm.22-.3E.3C.2Fspan.3E.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
	// Default value: 0, which indicates to delete all videos of the type specified by the `Type` parameter.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type MediaImageSpriteInfo struct {
	// Information set of image sprites with specified specifications. Each element represents a set of image sprites with the same specification.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageSpriteSet []*MediaImageSpriteItem `json:"ImageSpriteSet,omitempty" name:"ImageSpriteSet"`
}

type MediaImageSpriteItem struct {
	// Image sprite specification. For more information, please see [Image Sprite Parameter Template](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF).
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
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet"`

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

	// Subtitle information
	// Note: this field may return `null`, indicating that no valid value is obtained.
	SubtitleInfo *MediaSubtitleInfo `json:"SubtitleInfo,omitempty" name:"SubtitleInfo"`

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
	KeyFrameDescSet []*MediaKeyFrameDescItem `json:"KeyFrameDescSet,omitempty" name:"KeyFrameDescSet"`
}

type MediaKeyFrameDescItem struct {
	// Offset time of video timestamp in seconds.
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// Content string of timestamp containing 1-128 characters.
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
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet"`

	// Audio stream information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet"`

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
	MiniProgramReviewList []*MediaMiniProgramReviewInfoItem `json:"MiniProgramReviewList,omitempty" name:"MiniProgramReviewList"`
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
	ReviewSummary []*MediaMiniProgramReviewElem `json:"ReviewSummary,omitempty" name:"ReviewSummary"`
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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
	TranscodeTaskSet []*TranscodeTaskInput `json:"TranscodeTaskSet,omitempty" name:"TranscodeTaskSet"`

	// List of animated image generating tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInput `json:"AnimatedGraphicTaskSet,omitempty" name:"AnimatedGraphicTaskSet"`

	// List of time point screencapturing tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTaskSet,omitempty" name:"SnapshotByTimeOffsetTaskSet"`

	// List of sampled screencapturing tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SampleSnapshotTaskSet []*SampleSnapshotTaskInput `json:"SampleSnapshotTaskSet,omitempty" name:"SampleSnapshotTaskSet"`

	// List of image sprite generating tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageSpriteTaskSet []*ImageSpriteTaskInput `json:"ImageSpriteTaskSet,omitempty" name:"ImageSpriteTaskSet"`

	// List of cover generating tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoverBySnapshotTaskSet []*CoverBySnapshotTaskInput `json:"CoverBySnapshotTaskSet,omitempty" name:"CoverBySnapshotTaskSet"`

	// List of adaptive bitrate streaming tasks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTaskSet,omitempty" name:"AdaptiveDynamicStreamingTaskSet"`
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
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

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You’re not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Input of transcoding task.
	Input *TranscodeTaskInput `json:"Input,omitempty" name:"Input"`

	// Output of transcoding task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *MediaTranscodeItem `json:"Output,omitempty" name:"Output"`

	// Transcoding progress. Value range: 0-100.
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Transcoding task start time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732)
	BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

	// Transcoding task end time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732)
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
}

type MediaSampleSnapshotInfo struct {
	// Information set of sampled screenshots with the specified specifications. Each element represents a set of sampled screenshots with the same specification.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SampleSnapshotSet []*MediaSampleSnapshotItem `json:"SampleSnapshotSet,omitempty" name:"SampleSnapshotSet"`
}

type MediaSampleSnapshotItem struct {
	// Sampled screenshot specification ID. For more information, please see [Sampled Screencapturing Parameter Template](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF).
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
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet"`

	// List of watermarking template IDs if the screenshots are watermarked.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition"`
}

type MediaSnapshotByTimeOffsetInfo struct {
	// Information set of time point screenshots with a specified specification. Currently, there can be only one set of screenshots for each specification.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetSet []*MediaSnapshotByTimeOffsetItem `json:"SnapshotByTimeOffsetSet,omitempty" name:"SnapshotByTimeOffsetSet"`
}

type MediaSnapshotByTimeOffsetItem struct {
	// Specification of a time point screenshot. For more information, please see [Parameter Template for Time Point Screencapturing](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Information set of screenshots of the same specification. Each element represents a screenshot.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet,omitempty" name:"PicInfoSet"`
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
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition"`
}

type MediaSourceData struct {
	// Source of a media file:
	// <li>`Record`: recording, such as live or time-shift recording</li>
	// <li>`Upload`: upload, such as pull for upload, upload from server, and UGC upload from client</li>
	// <li>`VideoProcessing`: video processing, such as video splicing and video clipping</li>
	// <li>`WebPageRecord`: panoramic recording </li>
	// <li>`Unknown`: unknown source</li>
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// Field passed through when a file is created.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`
}

type MediaSubtitleInfo struct {
	// Subtitle information list
	SubtitleSet []*MediaSubtitleItem `json:"SubtitleSet,omitempty" name:"SubtitleSet"`
}

type MediaSubtitleInput struct {
	// Subtitle name. Length limit: 64 characters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subtitle language. Common values:
	// <li>`cn`: Chinese</li>
	// <li>`ja`: Japanese</li>
	// <li>`en-US`: English</li>
	// For other valid values, see [RFC 5646](https://tools.ietf.org/html/rfc5646).
	Language *string `json:"Language,omitempty" name:"Language"`

	// Subtitle format. Valid value:
	// <li>vtt</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Subtitle content, which is [Base64-encoded](https://tools.ietf.org/html/rfc4648) strings
	Content *string `json:"Content,omitempty" name:"Content"`

	// Subtitle ID. Its length cannot exceed 16 characters. Uppercase and lowercase letters, numbers, underscores (_), and hyphens (-) are supported. It cannot be the same as the IDs of the existing subtitles in the media file.
	Id *string `json:"Id,omitempty" name:"Id"`
}

type MediaSubtitleItem struct {
	// Unique subtitle ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Subtitle name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subtitle language. Common values:
	// <li>`cn`: Chinese</li>
	// <li>`ja`: Japanese</li>
	// <li>`en-US`: English</li>
	// For other values, see [RFC 5646](https://tools.ietf.org/html/rfc5646).
	Language *string `json:"Language,omitempty" name:"Language"`

	// Subtitle format. Valid value:
	// <li>vtt</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Subtitle URL
	Url *string `json:"Url,omitempty" name:"Url"`
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
	TrackItems []*MediaTrackItem `json:"TrackItems,omitempty" name:"TrackItems"`
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
	TranscodeSet []*MediaTranscodeItem `json:"TranscodeSet,omitempty" name:"TranscodeSet"`
}

type MediaTranscodeItem struct {
	// Address of output video file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Transcoding specification ID. For more information, please see [Transcoding Parameter Template](https://intl.cloud.tencent.com/document/product/266/33478?from_cn_redirect=1#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
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

	// The file size (bytes).
	// <li>If the file is an HLS file, the value of this parameter is the sum of the size of the M3U8 and TS files.</li>
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Video duration in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// MD5 value of video.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// Container, such as m4a and mp4.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Video stream information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet"`

	// Audio stream information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet"`
}

type MediaTransitionItem struct {
	// Transition duration in seconds. For two media segments that use a transition, the start time of the second segment on the track will be automatically set to the end time of the first segment minus the transition duration.
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// List of transition operations. Up to one video image or audio transition operation is supported.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Transitions []*TransitionOpertion `json:"Transitions,omitempty" name:"Transitions"`
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

// Predefined struct for user
type ModifyAIAnalysisTemplateRequestParams struct {
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
}

type ModifyAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "ClassificationConfigure")
	delete(f, "TagConfigure")
	delete(f, "CoverConfigure")
	delete(f, "FrameTagConfigure")
	delete(f, "HighlightConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAIAnalysisTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAIAnalysisTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAIAnalysisTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAIRecognitionTemplateRequestParams struct {
	// Unique ID of video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Video content recognition template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter of video opening and ending credits recognition.
	HeadTailConfigure *HeadTailConfigureInfoForUpdate `json:"HeadTailConfigure,omitempty" name:"HeadTailConfigure"`

	// Control parameter of video splitting recognition.
	SegmentConfigure *SegmentConfigureInfoForUpdate `json:"SegmentConfigure,omitempty" name:"SegmentConfigure"`

	// Control parameter of face recognition.
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// Control parameter of full text recognition.
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// Control parameter of text keyword recognition.
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// Control parameter of full speech recognition.
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// Control parameter of speech keyword recognition.
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`

	// Control parameter of object recognition.
	ObjectConfigure *ObjectConfigureInfoForUpdate `json:"ObjectConfigure,omitempty" name:"ObjectConfigure"`

	// Frame capturing interval in seconds. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
}

type ModifyAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of video content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Video content recognition template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter of video opening and ending credits recognition.
	HeadTailConfigure *HeadTailConfigureInfoForUpdate `json:"HeadTailConfigure,omitempty" name:"HeadTailConfigure"`

	// Control parameter of video splitting recognition.
	SegmentConfigure *SegmentConfigureInfoForUpdate `json:"SegmentConfigure,omitempty" name:"SegmentConfigure"`

	// Control parameter of face recognition.
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// Control parameter of full text recognition.
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// Control parameter of text keyword recognition.
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// Control parameter of full speech recognition.
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// Control parameter of speech keyword recognition.
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`

	// Control parameter of object recognition.
	ObjectConfigure *ObjectConfigureInfoForUpdate `json:"ObjectConfigure,omitempty" name:"ObjectConfigure"`

	// Frame capturing interval in seconds. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "HeadTailConfigure")
	delete(f, "SegmentConfigure")
	delete(f, "FaceConfigure")
	delete(f, "OcrFullTextConfigure")
	delete(f, "OcrWordsConfigure")
	delete(f, "AsrFullTextConfigure")
	delete(f, "AsrWordsConfigure")
	delete(f, "ObjectConfigure")
	delete(f, "ScreenshotInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAIRecognitionTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAIRecognitionTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAIRecognitionTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAdaptiveDynamicStreamingTemplateRequestParams struct {
	// Unique ID of adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The adaptive bitrate streaming format. Valid values:
	// <li>HLS</li>
	// <li>MPEG-DASH</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitempty" name:"DisableHigherVideoResolution"`

	// Parameter information of input stream for adaptive bitrate streaming. Up to 10 streams can be input.
	// Note: the frame rate of all streams must be the same; otherwise, the frame rate of the first stream will be used as the output frame rate.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitempty" name:"StreamInfos"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type ModifyAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The adaptive bitrate streaming format. Valid values:
	// <li>HLS</li>
	// <li>MPEG-DASH</li>
	Format *string `json:"Format,omitempty" name:"Format"`

	// Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitempty" name:"DisableHigherVideoResolution"`

	// Parameter information of input stream for adaptive bitrate streaming. Up to 10 streams can be input.
	// Note: the frame rate of all streams must be the same; otherwise, the frame rate of the first stream will be used as the output frame rate.
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
	delete(f, "SubAppId")
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

// Predefined struct for user
type ModifyAdaptiveDynamicStreamingTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAdaptiveDynamicStreamingTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAdaptiveDynamicStreamingTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAnimatedGraphicsTemplateRequestParams struct {
	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
}

type ModifyAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type ModifyAnimatedGraphicsTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAnimatedGraphicsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAnimatedGraphicsTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyClassRequestParams struct {
	// Category ID
	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

	// Category name, which can contain 1–64 characters.
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type ModifyClassRequest struct {
	*tchttp.BaseRequest
	
	// Category ID
	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

	// Category name, which can contain 1–64 characters.
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClassId")
	delete(f, "ClassName")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClassResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClassResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClassResponseParams `json:"Response"`
}

func (r *ModifyClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContentReviewTemplateRequestParams struct {
	// Unique ID of an intelligent content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Name of an intelligent content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of an intelligent content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter for terrorism information.
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// Control parameter for porn information.
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// Control parameter for politically sensitive information.
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	ProhibitedConfigure *ProhibitedConfigureInfoForUpdate `json:"ProhibitedConfigure,omitempty" name:"ProhibitedConfigure"`

	// Control parameter for custom intelligent content recognition tasks.
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// Frame capturing interval in seconds. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// Whether to allow the recognition result to enter the intelligent recognition platform (for human recognition).
	// <li>ON: yes</li>
	// <li>OFF: no</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitempty" name:"ReviewWallSwitch"`
}

type ModifyContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an intelligent content recognition template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Name of an intelligent content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description of an intelligent content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Control parameter for terrorism information.
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// Control parameter for porn information.
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// Control parameter for politically sensitive information.
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	ProhibitedConfigure *ProhibitedConfigureInfoForUpdate `json:"ProhibitedConfigure,omitempty" name:"ProhibitedConfigure"`

	// Control parameter for custom intelligent content recognition tasks.
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// Frame capturing interval in seconds. Minimum value: 0.5 seconds.
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// Whether to allow the recognition result to enter the intelligent recognition platform (for human recognition).
	// <li>ON: yes</li>
	// <li>OFF: no</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitempty" name:"ReviewWallSwitch"`
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
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "TerrorismConfigure")
	delete(f, "PornConfigure")
	delete(f, "PoliticalConfigure")
	delete(f, "ProhibitedConfigure")
	delete(f, "UserDefineConfigure")
	delete(f, "ScreenshotInterval")
	delete(f, "ReviewWallSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyContentReviewTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContentReviewTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyContentReviewTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDefaultStorageRegionRequestParams struct {
	// The default storage region, which must be a region you have storage access to. You can use the `DescribeStorageRegions` API to query such regions.
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type ModifyDefaultStorageRegionRequest struct {
	*tchttp.BaseRequest
	
	// The default storage region, which must be a region you have storage access to. You can use the `DescribeStorageRegions` API to query such regions.
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyDefaultStorageRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultStorageRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StorageRegion")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDefaultStorageRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultStorageRegionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDefaultStorageRegionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDefaultStorageRegionResponseParams `json:"Response"`
}

func (r *ModifyDefaultStorageRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultStorageRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSpriteTemplateRequestParams struct {
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type ModifyImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type ModifyImageSpriteTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyImageSpriteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageSpriteTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyMediaInfoRequestParams struct {
	// Unique media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Media filename, which can contain up to 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Media file description, which can contain up to 128 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Media file category ID.
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Media file expiration time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). The value `9999-12-31T23:59:59Z` indicates that the media file never expires. After the expiration, the media file and its related resources (such as transcoding results and image sprites) will be permanently deleted.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// String generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) the video cover image file (such as .jpeg or .png file). Only .gif, .jpeg, and .png image formats are supported.
	CoverData *string `json:"CoverData,omitempty" name:"CoverData"`

	// Set of video timestamps to be added. If a timestamp already exists at an offset time point, it will be overwritten. Up to 100 timestamps can be added to one media file. In the same request, the time offset parameters of `AddKeyFrameDescs` must be different from those of `DeleteKeyFrameDescs`.
	AddKeyFrameDescs []*MediaKeyFrameDescItem `json:"AddKeyFrameDescs,omitempty" name:"AddKeyFrameDescs"`

	// Time offset of the set of video timestamps to be deleted in seconds. In the same request, the time offset parameters of `AddKeyFrameDescs` must be different from those of `DeleteKeyFrameDescs`.
	DeleteKeyFrameDescs []*float64 `json:"DeleteKeyFrameDescs,omitempty" name:"DeleteKeyFrameDescs"`

	// The value `1` indicates to delete all timestamps in the video. Other values are meaningless.
	// In the same request, `ClearKeyFrameDescs` and `AddKeyFrameDescs` cannot be present at the same time.
	ClearKeyFrameDescs *int64 `json:"ClearKeyFrameDescs,omitempty" name:"ClearKeyFrameDescs"`

	// Set of tags to be added. Up to 16 tags can be added to one media file, and one tag can contain up to 16 characters. In the same request, the parameters of `AddTags` must be different from those of `DeleteTags`.
	AddTags []*string `json:"AddTags,omitempty" name:"AddTags"`

	// Set of tags to be deleted. In the same request, the parameters of `AddTags` must be different from those of `DeleteTags`.
	DeleteTags []*string `json:"DeleteTags,omitempty" name:"DeleteTags"`

	// The value `1` indicates to delete all tags of the media file. Other values are meaningless.
	// In the same request, `ClearTags` and `AddTags` cannot be present at the same time.
	ClearTags *int64 `json:"ClearTags,omitempty" name:"ClearTags"`

	// Information of multiple subtitles to be added. A single media file can have up to 16 subtitles. In the same request, the subtitle IDs specified in `AddSubtitles` must be different from those in `DeleteSubtitleIds`.
	AddSubtitles []*MediaSubtitleInput `json:"AddSubtitles,omitempty" name:"AddSubtitles"`

	// Unique IDs of the subtitles to be deleted. In the same request, the subtitle IDs specified in `AddSubtitles` must be different from those in `DeleteSubtitleIds`.
	DeleteSubtitleIds []*string `json:"DeleteSubtitleIds,omitempty" name:"DeleteSubtitleIds"`

	// The value `1` indicates to delete all subtitle information of the media file. Other values are meaningless.
	// `ClearSubtitles` and `AddSubtitles` cannot co-exist in the same request.
	ClearSubtitles *int64 `json:"ClearSubtitles,omitempty" name:"ClearSubtitles"`
}

type ModifyMediaInfoRequest struct {
	*tchttp.BaseRequest
	
	// Unique media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Media filename, which can contain up to 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Media file description, which can contain up to 128 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Media file category ID.
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Media file expiration time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). The value `9999-12-31T23:59:59Z` indicates that the media file never expires. After the expiration, the media file and its related resources (such as transcoding results and image sprites) will be permanently deleted.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// String generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) the video cover image file (such as .jpeg or .png file). Only .gif, .jpeg, and .png image formats are supported.
	CoverData *string `json:"CoverData,omitempty" name:"CoverData"`

	// Set of video timestamps to be added. If a timestamp already exists at an offset time point, it will be overwritten. Up to 100 timestamps can be added to one media file. In the same request, the time offset parameters of `AddKeyFrameDescs` must be different from those of `DeleteKeyFrameDescs`.
	AddKeyFrameDescs []*MediaKeyFrameDescItem `json:"AddKeyFrameDescs,omitempty" name:"AddKeyFrameDescs"`

	// Time offset of the set of video timestamps to be deleted in seconds. In the same request, the time offset parameters of `AddKeyFrameDescs` must be different from those of `DeleteKeyFrameDescs`.
	DeleteKeyFrameDescs []*float64 `json:"DeleteKeyFrameDescs,omitempty" name:"DeleteKeyFrameDescs"`

	// The value `1` indicates to delete all timestamps in the video. Other values are meaningless.
	// In the same request, `ClearKeyFrameDescs` and `AddKeyFrameDescs` cannot be present at the same time.
	ClearKeyFrameDescs *int64 `json:"ClearKeyFrameDescs,omitempty" name:"ClearKeyFrameDescs"`

	// Set of tags to be added. Up to 16 tags can be added to one media file, and one tag can contain up to 16 characters. In the same request, the parameters of `AddTags` must be different from those of `DeleteTags`.
	AddTags []*string `json:"AddTags,omitempty" name:"AddTags"`

	// Set of tags to be deleted. In the same request, the parameters of `AddTags` must be different from those of `DeleteTags`.
	DeleteTags []*string `json:"DeleteTags,omitempty" name:"DeleteTags"`

	// The value `1` indicates to delete all tags of the media file. Other values are meaningless.
	// In the same request, `ClearTags` and `AddTags` cannot be present at the same time.
	ClearTags *int64 `json:"ClearTags,omitempty" name:"ClearTags"`

	// Information of multiple subtitles to be added. A single media file can have up to 16 subtitles. In the same request, the subtitle IDs specified in `AddSubtitles` must be different from those in `DeleteSubtitleIds`.
	AddSubtitles []*MediaSubtitleInput `json:"AddSubtitles,omitempty" name:"AddSubtitles"`

	// Unique IDs of the subtitles to be deleted. In the same request, the subtitle IDs specified in `AddSubtitles` must be different from those in `DeleteSubtitleIds`.
	DeleteSubtitleIds []*string `json:"DeleteSubtitleIds,omitempty" name:"DeleteSubtitleIds"`

	// The value `1` indicates to delete all subtitle information of the media file. Other values are meaningless.
	// `ClearSubtitles` and `AddSubtitles` cannot co-exist in the same request.
	ClearSubtitles *int64 `json:"ClearSubtitles,omitempty" name:"ClearSubtitles"`
}

func (r *ModifyMediaInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ClassId")
	delete(f, "ExpireTime")
	delete(f, "CoverData")
	delete(f, "AddKeyFrameDescs")
	delete(f, "DeleteKeyFrameDescs")
	delete(f, "ClearKeyFrameDescs")
	delete(f, "AddTags")
	delete(f, "DeleteTags")
	delete(f, "ClearTags")
	delete(f, "AddSubtitles")
	delete(f, "DeleteSubtitleIds")
	delete(f, "ClearSubtitles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMediaInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMediaInfoResponseParams struct {
	// URL of new video cover.
	// * Note: this returned value is valid only if the request carries `CoverData`.*
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// Added subtitle information
	AddedSubtitleSet []*MediaSubtitleItem `json:"AddedSubtitleSet,omitempty" name:"AddedSubtitleSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMediaInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMediaInfoResponseParams `json:"Response"`
}

func (r *ModifyMediaInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMediaStorageClassRequestParams struct {
	// The unique IDs of media files
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// The target storage class. Valid values:
	// <li>STANDARD</li>
	// <li>STANDARD_IA</li>
	// <li>ARCHIVE</li>
	// <li>DEEP_ARCHIVE</li>
	StorageClass *string `json:"StorageClass,omitempty" name:"StorageClass"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// The retrieval mode. When switching files from DEEP ARCHIVE or ARCHIVE to STANDARD, you need to specify the retrieval mode. For details, see [Data retrieval and retrieval mode](https://intl.cloud.tencent.com/document/product/266/43051#data-retrieval-and-retrieval-mode.3Ca-id.3D.22retake.22.3E.3C.2Fa.3E).
	// If the current storage class is ARCHIVE, the valid values for this parameter are as follows:
	// <li>Expedited</li>
	// <li>Standard</li>
	// <li>Bulk</li>
	// If the current storage class is DEEP ARCHIVE, the valid values for this parameter are as follows:
	// <li>Standard</li>
	// <li>Bulk</li>
	RestoreTier *string `json:"RestoreTier,omitempty" name:"RestoreTier"`
}

type ModifyMediaStorageClassRequest struct {
	*tchttp.BaseRequest
	
	// The unique IDs of media files
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// The target storage class. Valid values:
	// <li>STANDARD</li>
	// <li>STANDARD_IA</li>
	// <li>ARCHIVE</li>
	// <li>DEEP_ARCHIVE</li>
	StorageClass *string `json:"StorageClass,omitempty" name:"StorageClass"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// The retrieval mode. When switching files from DEEP ARCHIVE or ARCHIVE to STANDARD, you need to specify the retrieval mode. For details, see [Data retrieval and retrieval mode](https://intl.cloud.tencent.com/document/product/266/43051#data-retrieval-and-retrieval-mode.3Ca-id.3D.22retake.22.3E.3C.2Fa.3E).
	// If the current storage class is ARCHIVE, the valid values for this parameter are as follows:
	// <li>Expedited</li>
	// <li>Standard</li>
	// <li>Bulk</li>
	// If the current storage class is DEEP ARCHIVE, the valid values for this parameter are as follows:
	// <li>Standard</li>
	// <li>Bulk</li>
	RestoreTier *string `json:"RestoreTier,omitempty" name:"RestoreTier"`
}

func (r *ModifyMediaStorageClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaStorageClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileIds")
	delete(f, "StorageClass")
	delete(f, "SubAppId")
	delete(f, "RestoreTier")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMediaStorageClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMediaStorageClassResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMediaStorageClassResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMediaStorageClassResponseParams `json:"Response"`
}

func (r *ModifyMediaStorageClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaStorageClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonSampleRequestParams struct {
	// ID of a sample.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Name. Length limit: 128 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description. Length limit: 1,024 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Sample usage. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: used for content recognition and inappropriate information recognition; equivalent to 1+2
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Information of operations on facial features.
	FaceOperationInfo *AiSampleFaceOperation `json:"FaceOperationInfo,omitempty" name:"FaceOperationInfo"`

	// Tag operation information.
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitempty" name:"TagOperationInfo"`
}

type ModifyPersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// ID of a sample.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Name. Length limit: 128 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description. Length limit: 1,024 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Sample usage. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: used for content recognition and inappropriate information recognition; equivalent to 1+2
	Usages []*string `json:"Usages,omitempty" name:"Usages"`

	// Information of operations on facial features.
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
	delete(f, "SubAppId")
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

// Predefined struct for user
type ModifyPersonSampleResponseParams struct {
	// Information of a sample.
	Person *AiSamplePerson `json:"Person,omitempty" name:"Person"`

	// Information of samples that failed the verification by facial feature positioning.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitempty" name:"FailFaceInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPersonSampleResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifySampleSnapshotTemplateRequestParams struct {
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type ModifySampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type ModifySampleSnapshotTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySampleSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifySampleSnapshotTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifySnapshotByTimeOffsetTemplateRequestParams struct {
	// Unique ID of a specified time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type ModifySnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a specified time point screencapturing template.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type ModifySnapshotByTimeOffsetTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySnapshotByTimeOffsetTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotByTimeOffsetTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifySubAppIdInfoRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Subapplication name. Length limit: 40 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subapplication overview. Length limit: 300 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifySubAppIdInfoRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubAppIdInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubAppIdInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubAppIdInfoResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubAppIdInfoResponseParams `json:"Response"`
}

func (r *ModifySubAppIdInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubAppIdInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubAppIdStatusRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Subapplication status. Valid values:
	// <li>On: enabled</li>
	// <li>Off: disabled</li>
	// <li>Destroyed: terminated</li>
	// You cannot enable a subapplication whose status is “Destroying”. You can enable it after it was terminated.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifySubAppIdStatusRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Subapplication status. Valid values:
	// <li>On: enabled</li>
	// <li>Off: disabled</li>
	// <li>Destroyed: terminated</li>
	// You cannot enable a subapplication whose status is “Destroying”. You can enable it after it was terminated.
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifySubAppIdStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubAppIdStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubAppIdStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubAppIdStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubAppIdStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubAppIdStatusResponseParams `json:"Response"`
}

func (r *ModifySubAppIdStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubAppIdStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySuperPlayerConfigRequestParams struct {
	// Player configuration name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type of audio/video played. Valid values:
	// <li>AdaptiveDynamicStreaming</li>
	// <li>Transcode</li>
	// <li>Original</li>
	AudioVideoType *string `json:"AudioVideoType,omitempty" name:"AudioVideoType"`

	// Switch of DRM-protected adaptive bitstream playback:
	// <li>ON: enabled, indicating to play back only output adaptive bitstreams protected by DRM;</li>
	// <li>OFF: disabled, indicating to play back unencrypted output adaptive bitstreams.</li>
	DrmSwitch *string `json:"DrmSwitch,omitempty" name:"DrmSwitch"`

	// ID of the unencrypted adaptive bitrate streaming template that allows output.
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// Content of the DRM-protected adaptive bitrate streaming template that allows output.
	DrmStreamingsInfo *DrmStreamingsInfoForUpdate `json:"DrmStreamingsInfo,omitempty" name:"DrmStreamingsInfo"`

	// ID of the transcoding template allowed for playback
	TranscodeDefinition *uint64 `json:"TranscodeDefinition,omitempty" name:"TranscodeDefinition"`

	// ID of the image sprite generating template that allows output.
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitempty" name:"ImageSpriteDefinition"`

	// Display name of player for substreams with different resolutions.
	ResolutionNames []*ResolutionNameInfo `json:"ResolutionNames,omitempty" name:"ResolutionNames"`

	// Domain name used for playback. If its value is `Default`, the domain name configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Scheme used for playback. Valid values:
	// <li>Default: the scheme configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used;</li>
	// <li>HTTP;</li>
	// <li>HTTPS.</li>
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type ModifySuperPlayerConfigRequest struct {
	*tchttp.BaseRequest
	
	// Player configuration name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type of audio/video played. Valid values:
	// <li>AdaptiveDynamicStreaming</li>
	// <li>Transcode</li>
	// <li>Original</li>
	AudioVideoType *string `json:"AudioVideoType,omitempty" name:"AudioVideoType"`

	// Switch of DRM-protected adaptive bitstream playback:
	// <li>ON: enabled, indicating to play back only output adaptive bitstreams protected by DRM;</li>
	// <li>OFF: disabled, indicating to play back unencrypted output adaptive bitstreams.</li>
	DrmSwitch *string `json:"DrmSwitch,omitempty" name:"DrmSwitch"`

	// ID of the unencrypted adaptive bitrate streaming template that allows output.
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// Content of the DRM-protected adaptive bitrate streaming template that allows output.
	DrmStreamingsInfo *DrmStreamingsInfoForUpdate `json:"DrmStreamingsInfo,omitempty" name:"DrmStreamingsInfo"`

	// ID of the transcoding template allowed for playback
	TranscodeDefinition *uint64 `json:"TranscodeDefinition,omitempty" name:"TranscodeDefinition"`

	// ID of the image sprite generating template that allows output.
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitempty" name:"ImageSpriteDefinition"`

	// Display name of player for substreams with different resolutions.
	ResolutionNames []*ResolutionNameInfo `json:"ResolutionNames,omitempty" name:"ResolutionNames"`

	// Domain name used for playback. If its value is `Default`, the domain name configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Scheme used for playback. Valid values:
	// <li>Default: the scheme configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used;</li>
	// <li>HTTP;</li>
	// <li>HTTPS.</li>
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifySuperPlayerConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySuperPlayerConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AudioVideoType")
	delete(f, "DrmSwitch")
	delete(f, "AdaptiveDynamicStreamingDefinition")
	delete(f, "DrmStreamingsInfo")
	delete(f, "TranscodeDefinition")
	delete(f, "ImageSpriteDefinition")
	delete(f, "ResolutionNames")
	delete(f, "Domain")
	delete(f, "Scheme")
	delete(f, "Comment")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySuperPlayerConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySuperPlayerConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySuperPlayerConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifySuperPlayerConfigResponseParams `json:"Response"`
}

func (r *ModifySuperPlayerConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySuperPlayerConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTranscodeTemplateRequestParams struct {
	// Unique ID of transcoding template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Container. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Transcoding template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
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
}

type ModifyTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of transcoding template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Container. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitempty" name:"Container"`

	// Transcoding template name. Length limit: 64 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
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
	delete(f, "SubAppId")
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

// Predefined struct for user
type ModifyTranscodeTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTranscodeTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyVodDomainAccelerateConfigRequestParams struct {
	// Domain name for acceleration setting
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Region. Valid values:
	// <li>`Chinese Mainland`</li>
	// <li>`Outside Chinese Mainland`</li>
	// <li>`Global`</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether to enable or disable domain name acceleration for the selected region. Valid values:
	// <li>`Enabled`: enable</li>
	// <li>`Disabled`: disable</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type ModifyVodDomainAccelerateConfigRequest struct {
	*tchttp.BaseRequest
	
	// Domain name for acceleration setting
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Region. Valid values:
	// <li>`Chinese Mainland`</li>
	// <li>`Outside Chinese Mainland`</li>
	// <li>`Global`</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether to enable or disable domain name acceleration for the selected region. Valid values:
	// <li>`Enabled`: enable</li>
	// <li>`Disabled`: disable</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyVodDomainAccelerateConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVodDomainAccelerateConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Area")
	delete(f, "Status")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVodDomainAccelerateConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVodDomainAccelerateConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVodDomainAccelerateConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVodDomainAccelerateConfigResponseParams `json:"Response"`
}

func (r *ModifyVodDomainAccelerateConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVodDomainAccelerateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVodDomainConfigRequestParams struct {
	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// [Referer hotlink protection](https://intl.cloud.tencent.com/document/product/266/14046?from_cn_redirect=1) policy
	RefererAuthPolicy *RefererAuthPolicy `json:"RefererAuthPolicy,omitempty" name:"RefererAuthPolicy"`

	// [Key hotlink protection](https://intl.cloud.tencent.com/document/product/266/14047?from_cn_redirect=1) policy
	UrlSignatureAuthPolicy *UrlSignatureAuthPolicy `json:"UrlSignatureAuthPolicy,omitempty" name:"UrlSignatureAuthPolicy"`
}

type ModifyVodDomainConfigRequest struct {
	*tchttp.BaseRequest
	
	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// [Referer hotlink protection](https://intl.cloud.tencent.com/document/product/266/14046?from_cn_redirect=1) policy
	RefererAuthPolicy *RefererAuthPolicy `json:"RefererAuthPolicy,omitempty" name:"RefererAuthPolicy"`

	// [Key hotlink protection](https://intl.cloud.tencent.com/document/product/266/14047?from_cn_redirect=1) policy
	UrlSignatureAuthPolicy *UrlSignatureAuthPolicy `json:"UrlSignatureAuthPolicy,omitempty" name:"UrlSignatureAuthPolicy"`
}

func (r *ModifyVodDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVodDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SubAppId")
	delete(f, "RefererAuthPolicy")
	delete(f, "UrlSignatureAuthPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVodDomainConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVodDomainConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVodDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVodDomainConfigResponseParams `json:"Response"`
}

func (r *ModifyVodDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVodDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWatermarkTemplateRequestParams struct {
	// Unique ID of watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
}

type ModifyWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of watermarking template.
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
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

// Predefined struct for user
type ModifyWatermarkTemplateResponseParams struct {
	// Image watermark address. This field has a value only when `ImageTemplate.ImageContent` is not empty.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWatermarkTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyWordSampleRequestParams struct {
	// Keyword. Length limit: 128 characters.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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

type ModifyWordSampleRequest struct {
	*tchttp.BaseRequest
	
	// Keyword. Length limit: 128 characters.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

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
	delete(f, "SubAppId")
	delete(f, "Usages")
	delete(f, "TagOperationInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWordSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWordSampleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWordSampleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWordSampleResponseParams `json:"Response"`
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

type ObjectConfigureInfoForUpdate struct {
	// Switch of object recognition task. Valid values:
	// <li>ON: enables intelligent object recognition task;</li>
	// <li>OFF: disables intelligent object recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Object library. Valid values:
	// <li>Default: default object library;</li>
	// <li>UserDefine: custom object library.</li>
	// <li>All: both default and custom object libraries will be used.</li>
	ObjectLibrary *string `json:"ObjectLibrary,omitempty" name:"ObjectLibrary"`
}

type OcrFullTextConfigureInfo struct {
	// Switch of full text recognition task. Valid values:
	// <li>ON: enables intelligent full text recognition task;</li>
	// <li>OFF: disables intelligent full text recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type OcrFullTextConfigureInfoForUpdate struct {
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
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`
}

type OcrWordsConfigureInfoForUpdate struct {
	// Switch of text keyword recognition task. Valid values:
	// <li>ON: enables text keyword recognition task;</li>
	// <li>OFF: disables text keyword recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty or a blank value is entered, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`
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

// Predefined struct for user
type ParseStreamingManifestRequestParams struct {
	// Index file content to be parsed.
	MediaManifestContent *string `json:"MediaManifestContent,omitempty" name:"MediaManifestContent"`

	// Video index file format, which is `m3u8` by default.
	// <li>m3u8</li>
	// <li>mpd</li>
	ManifestType *string `json:"ManifestType,omitempty" name:"ManifestType"`
}

type ParseStreamingManifestRequest struct {
	*tchttp.BaseRequest
	
	// Index file content to be parsed.
	MediaManifestContent *string `json:"MediaManifestContent,omitempty" name:"MediaManifestContent"`

	// Video index file format, which is `m3u8` by default.
	// <li>m3u8</li>
	// <li>mpd</li>
	ManifestType *string `json:"ManifestType,omitempty" name:"ManifestType"`
}

func (r *ParseStreamingManifestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseStreamingManifestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaManifestContent")
	delete(f, "ManifestType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseStreamingManifestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseStreamingManifestResponseParams struct {
	// Segment file list.
	MediaSegmentSet []*string `json:"MediaSegmentSet,omitempty" name:"MediaSegmentSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ParseStreamingManifestResponse struct {
	*tchttp.BaseResponse
	Response *ParseStreamingManifestResponseParams `json:"Response"`
}

func (r *ParseStreamingManifestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseStreamingManifestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlayStatFileInfo struct {
	// Date of playback statistics in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	Date *string `json:"Date,omitempty" name:"Date"`

	// URL of a playback statistics file, including the following contents:
	// <li> date: playback date</li>
	// <li> file_id: video file ID</li>
	// <li> ip_count: number of client IPs after deduplication</li>
	// <li> flux: playback traffic in bytes</li>
	// <li> play_times: total playback times</li>
	// <li> pc_play_times: playback times on PC clients</li>
	// <li> mobile_play_times: playback times on mobile clients</li>
	// <li> iphone_play_times: playback times on iPhone</li>
	// <li> android_play_times: playback times on Android</li>
	// <li> host_name: domain name</li>
	Url *string `json:"Url,omitempty" name:"Url"`
}

type PlayStatInfo struct {
	// The start time (in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I)) of the data returned. For example, if the granularity is a day, `2018-12-01T00:00:00+08:00` indicates that the data is for the period between December 1, 2018 (inclusive) and December 2, 2018 (exclusive).
	// <li>If the granularity is an hour, `2019-08-22T00:00:00+08:00` indicates the data is for the period between 00:00 and 01:00 AM on August 22, 2019.</li>
	// <li>If the granularity is a day, `2019-08-22T00:00:00+08:00` indicates the data is for August 22, 2019.</li>
	Time *string `json:"Time,omitempty" name:"Time"`

	// The ID of the media file.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// The playback times.
	PlayTimes *uint64 `json:"PlayTimes,omitempty" name:"PlayTimes"`

	// The traffic (in bytes) consumed for playback.
	Traffic *uint64 `json:"Traffic,omitempty" name:"Traffic"`
}

type PlayerConfig struct {
	// Player configuration name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Player configuration type. Valid values:
	// <li>Preset: preset configuration;</li>
	// <li>Custom: custom configuration.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Switch of DRM-protected adaptive bitstream playback:
	// <li>ON: enabled, indicating to play back only output adaptive bitstreams protected by DRM;</li>
	// <li>OFF: disabled, indicating to play back unencrypted output adaptive bitstreams.</li>
	DrmSwitch *string `json:"DrmSwitch,omitempty" name:"DrmSwitch"`

	// ID of the unencrypted adaptive bitrate streaming template that allows output.
	AdaptiveDynamicStreamingDefinition *uint64 `json:"AdaptiveDynamicStreamingDefinition,omitempty" name:"AdaptiveDynamicStreamingDefinition"`

	// Content of the DRM-protected adaptive bitrate streaming template that allows output.
	DrmStreamingsInfo *DrmStreamingsInfo `json:"DrmStreamingsInfo,omitempty" name:"DrmStreamingsInfo"`

	// ID of the image sprite generating template that allows output.
	ImageSpriteDefinition *uint64 `json:"ImageSpriteDefinition,omitempty" name:"ImageSpriteDefinition"`

	// Display name of player for substreams with different resolutions.
	ResolutionNameSet []*ResolutionNameInfo `json:"ResolutionNameSet,omitempty" name:"ResolutionNameSet"`

	// Creation time of player configuration in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of player configuration in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Domain name used for playback. If its value is `Default`, the domain name configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Scheme used for playback. Valid values:
	// <li>Default: the scheme configured in [Default Distribution Configuration](https://intl.cloud.tencent.com/document/product/266/33373?from_cn_redirect=1) will be used;</li>
	// <li>HTTP;</li>
	// <li>HTTPS.</li>
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// Template description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type PoliticalAsrReviewTemplateInfo struct {
	// Whether to enable ASR-based recognition of politically sensitive content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`
}

type PoliticalAsrReviewTemplateInfoForUpdate struct {
	// Whether to enable ASR-based recognition of politically sensitive content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalConfigureInfo struct {
	// Parameters for recognition of politically sensitive content in images
	// Note: This field may return `null`, indicating that no valid value can be found.
	ImgReviewInfo *PoliticalImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Parameters for ASR-based recognition of politically sensitive content
	// Note: This field may return `null`, indicating that no valid value can be found.
	AsrReviewInfo *PoliticalAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Parameters for OCR-based recognition of politically sensitive content
	// Note: This field may return `null`, indicating that no valid value can be found.
	OcrReviewInfo *PoliticalOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PoliticalConfigureInfoForUpdate struct {
	// Parameters for recognition of politically sensitive content in images
	ImgReviewInfo *PoliticalImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Parameters for ASR-based recognition of politically sensitive content
	AsrReviewInfo *PoliticalAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Parameters for OCR-based recognition of politically sensitive content
	OcrReviewInfo *PoliticalOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PoliticalImgReviewTemplateInfo struct {
	// Whether to enable recognition of politically sensitive content in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for recognition of politically sensitive content in images. Results containing the specified labels are returned. If no labels are specified, all results are returned. Valid values:
	// <li>`violation_photo`: banned images</li>
	// <li>`politician`: politically sensitive people</li>
	// <li>`entertainment`: people in the entertainment industry</li>
	// <li>`sport`: sportspeople</li>
	// <li>`entrepreneur`: businesspeople</li>
	// <li>`scholar`: scholars</li>
	// <li>`celebrity`: celebrities</li>
	// <li>`military`: people in military</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `97` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `95` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalImgReviewTemplateInfoForUpdate struct {
	// Whether to enable recognition of politically sensitive content in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for recognition of politically sensitive content in images. Results containing the specified labels are returned. If no labels are specified, all results are returned. Valid values:
	// <li>`violation_photo`: banned images</li>
	// <li>`politician`: politically sensitive people</li>
	// <li>`entertainment`: people in the entertainment industry</li>
	// <li>`sport`: sportspeople</li>
	// <li>`entrepreneur`: businesspeople</li>
	// <li>`scholar`: scholars</li>
	// <li>`celebrity`: celebrities</li>
	// <li>`military`: people in military</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfo struct {
	// Whether to enable OCR-based recognition of politically sensitive content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfoForUpdate struct {
	// Whether to enable OCR-based recognition of politically sensitive content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfo struct {
	// Whether to enable ASR-based recognition of pornographic content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfoForUpdate struct {
	// Whether to enable ASR-based recognition of pornographic content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornConfigureInfo struct {
	// Parameters for recognition of pornographic content in images
	// Note: This field may return `null`, indicating that no valid value can be found.
	ImgReviewInfo *PornImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Parameters for ASR-based recognition of pornographic content
	// Note: This field may return `null`, indicating that no valid value can be found.
	AsrReviewInfo *PornAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Parameters for OCR-based recognition of pornographic content
	// Note: This field may return `null`, indicating that no valid value can be found.
	OcrReviewInfo *PornOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PornConfigureInfoForUpdate struct {
	// Parameters for recognition of pornographic content in images
	ImgReviewInfo *PornImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Parameters for ASR-based recognition of pornographic content
	AsrReviewInfo *PornAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Parameters for OCR-based recognition of pornographic content
	OcrReviewInfo *PornOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PornImgReviewTemplateInfo struct {
	// Whether to enable recognition of pornographic content in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for recognition of pornographic content in images. Results containing the specified labels are returned. If no labels are specified, all results are returned. Valid values:
	// <li>porn</li>
	// <li>vulgar</li>
	// <li>intimacy</li>
	// <li>sexy</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `90` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `0` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornImgReviewTemplateInfoForUpdate struct {
	// Whether to enable recognition of pornographic content in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for recognition of pornographic content in images. Results containing the specified labels are returned. If no labels are specified, all results are returned. Valid values:
	// <li>porn</li>
	// <li>vulgar</li>
	// <li>intimacy</li>
	// <li>sexy</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfo struct {
	// Whether to enable OCR-based recognition of pornographic content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfoForUpdate struct {
	// Whether to enable OCR-based recognition of pornographic content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0–100.
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
	MediaProcessResultSet []*MediaProcessTaskResult `json:"MediaProcessResultSet,omitempty" name:"MediaProcessResultSet"`

	// Status and result of an intelligent recognition task
	AiContentReviewResultSet []*AiContentReviewResult `json:"AiContentReviewResultSet,omitempty" name:"AiContentReviewResultSet"`

	// Execution status and result of video content analysis task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AiAnalysisResultSet []*AiAnalysisResult `json:"AiAnalysisResultSet,omitempty" name:"AiAnalysisResultSet"`

	// Execution status and result of video content recognition task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AiRecognitionResultSet []*AiRecognitionResult `json:"AiRecognitionResultSet,omitempty" name:"AiRecognitionResultSet"`

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

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Parameter of video processing task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Intelligent recognition task
	// Note: This field may return `null`, indicating that no valid value can be found.
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

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ProcessMediaByProcedureRequestParams struct {
	// Media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// [Task flow template](https://intl.cloud.tencent.com/document/product/266/11700?from_cn_redirect=1#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF) name.
	ProcedureName *string `json:"ProcedureName,omitempty" name:"ProcedureName"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Task flow priority. The higher the value, the higher the priority. Value range: -10-10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for task flow status change. Valid values: Finish, Change, None. If this parameter is left empty, `Finish` will be used.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

type ProcessMediaByProcedureRequest struct {
	*tchttp.BaseRequest
	
	// Media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// [Task flow template](https://intl.cloud.tencent.com/document/product/266/11700?from_cn_redirect=1#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF) name.
	ProcedureName *string `json:"ProcedureName,omitempty" name:"ProcedureName"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Task flow priority. The higher the value, the higher the priority. Value range: -10-10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for task flow status change. Valid values: Finish, Change, None. If this parameter is left empty, `Finish` will be used.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

func (r *ProcessMediaByProcedureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaByProcedureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileId")
	delete(f, "ProcedureName")
	delete(f, "SubAppId")
	delete(f, "TasksPriority")
	delete(f, "TasksNotifyMode")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaByProcedureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaByProcedureResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ProcessMediaByProcedureResponse struct {
	*tchttp.BaseResponse
	Response *ProcessMediaByProcedureResponseParams `json:"Response"`
}

func (r *ProcessMediaByProcedureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaByProcedureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaByUrlRequestParams struct {
	// This API is<font color='red'>disused</font>. You are advised to use an alternative API. For more information, see API overview.
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// Information of COS path to output file.
	OutputInfo *MediaOutputInfo `json:"OutputInfo,omitempty" name:"OutputInfo"`

	// Type parameter of video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Task flow priority. The higher the value, the higher the priority. Value range: -10-10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for task flow status change. Valid values: Finish, Change, None. If this parameter is left empty, `Finish` will be used.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type ProcessMediaByUrlRequest struct {
	*tchttp.BaseRequest
	
	// This API is<font color='red'>disused</font>. You are advised to use an alternative API. For more information, see API overview.
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// Information of COS path to output file.
	OutputInfo *MediaOutputInfo `json:"OutputInfo,omitempty" name:"OutputInfo"`

	// Type parameter of video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Task flow priority. The higher the value, the higher the priority. Value range: -10-10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for task flow status change. Valid values: Finish, Change, None. If this parameter is left empty, `Finish` will be used.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ProcessMediaByUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaByUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	delete(f, "OutputInfo")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "TasksPriority")
	delete(f, "TasksNotifyMode")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaByUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaByUrlResponseParams struct {
	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ProcessMediaByUrlResponse struct {
	*tchttp.BaseResponse
	Response *ProcessMediaByUrlResponseParams `json:"Response"`
}

func (r *ProcessMediaByUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaByUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaRequestParams struct {
	// Media file ID, i.e., the globally unique ID of a file in VOD assigned by the VOD backend after successful upload. This field can be obtained through the [video upload completion event notification](https://intl.cloud.tencent.com/document/product/266/7830?from_cn_redirect=1) or [VOD Console](https://console.cloud.tencent.com/vod/media).
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Parameter of video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Parameters for intelligent recognition
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Task flow priority. The higher the value, the higher the priority. Value range: -10-10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for task flow status change. Valid values: Finish, Change, None. If this parameter is left empty, `Finish` will be used.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

type ProcessMediaRequest struct {
	*tchttp.BaseRequest
	
	// Media file ID, i.e., the globally unique ID of a file in VOD assigned by the VOD backend after successful upload. This field can be obtained through the [video upload completion event notification](https://intl.cloud.tencent.com/document/product/266/7830?from_cn_redirect=1) or [VOD Console](https://console.cloud.tencent.com/vod/media).
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Parameter of video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Parameters for intelligent recognition
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// Task flow priority. The higher the value, the higher the priority. Value range: -10-10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// Notification mode for task flow status change. Valid values: Finish, Change, None. If this parameter is left empty, `Finish` will be used.
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
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
	delete(f, "FileId")
	delete(f, "SubAppId")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "TasksPriority")
	delete(f, "TasksNotifyMode")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaResponseParams struct {
	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ProcessMediaResponse struct {
	*tchttp.BaseResponse
	Response *ProcessMediaResponseParams `json:"Response"`
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

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type ProhibitedAsrReviewTemplateInfoForUpdate struct {
	// Switch of prohibited information detection in speech task. Valid values:
	// <li>ON: enables prohibited information detection in speech task;</li>
	// <li>OFF: disables prohibited information detection in speech task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
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

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type ProhibitedOcrReviewTemplateInfoForUpdate struct {
	// Switch of prohibited information detection in text task. Valid values:
	// <li>ON: enables prohibited information detection in text task;</li>
	// <li>OFF: disables prohibited information detection in text task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

// Predefined struct for user
type PullEventsRequestParams struct {
	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type PullEventsRequest struct {
	*tchttp.BaseRequest
	
	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *PullEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExtInfo")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullEventsResponseParams struct {
	// List of events.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EventSet []*EventContent `json:"EventSet,omitempty" name:"EventSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PullEventsResponse struct {
	*tchttp.BaseResponse
	Response *PullEventsResponseParams `json:"Response"`
}

func (r *PullEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullUploadRequestParams struct {
	// The URL of the media to pull, which can be in HLS format, but not DASH format.
	// For more information about supported extensions, see [Media types](https://intl.cloud.tencent.com/document/product/266/9760#media-types). Please make sure the URL is accessible.
	MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Media name.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// URL of video cover to be pulled. Only gif, jpeg, and png formats are supported.
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// Subsequent task for media. For more information, please see [Specifying Task Flow After Upload](https://intl.cloud.tencent.com/document/product/266/9759?from_cn_redirect=1).
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Expiration time of media file in ISO 8601 format. For more information, please see [Notes on ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Specifies upload region. This is only applicable to users that have special requirements for the upload region:
	// <li>If it is left empty, the upload region is your [default region](https://intl.cloud.tencent.com/document/product/266/14059?from=11329?from_cn_redirect=1#.E5.AD.98.E5.82.A8.E5.9C.B0.E5.9F.9F.E6.AD.A5.E9.AA.A4);</li>
	// <li>If it is specified, please make sure that the upload region has been [enabled for storage](https://intl.cloud.tencent.com/document/product/266/14059?from=11329?from_cn_redirect=1#.E5.AD.98.E5.82.A8.E5.9C.B0.E5.9F.9F.E6.AD.A5.E9.AA.A4).</li>
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [CreateClass](https://intl.cloud.tencent.com/document/product/266/7812?from_cn_redirect=1) API.
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// The source context which is used to pass through the user request information. After `Procedure` is specified, the task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// Source context, which is used to pass through the user request information. The [upload callback](https://intl.cloud.tencent.com/document/product/266/7830?from_cn_redirect=1) API will return the value of this field. It can contain up to 250 characters.
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`
}

type PullUploadRequest struct {
	*tchttp.BaseRequest
	
	// The URL of the media to pull, which can be in HLS format, but not DASH format.
	// For more information about supported extensions, see [Media types](https://intl.cloud.tencent.com/document/product/266/9760#media-types). Please make sure the URL is accessible.
	MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Media name.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// URL of video cover to be pulled. Only gif, jpeg, and png formats are supported.
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// Subsequent task for media. For more information, please see [Specifying Task Flow After Upload](https://intl.cloud.tencent.com/document/product/266/9759?from_cn_redirect=1).
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Expiration time of media file in ISO 8601 format. For more information, please see [Notes on ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Specifies upload region. This is only applicable to users that have special requirements for the upload region:
	// <li>If it is left empty, the upload region is your [default region](https://intl.cloud.tencent.com/document/product/266/14059?from=11329?from_cn_redirect=1#.E5.AD.98.E5.82.A8.E5.9C.B0.E5.9F.9F.E6.AD.A5.E9.AA.A4);</li>
	// <li>If it is specified, please make sure that the upload region has been [enabled for storage](https://intl.cloud.tencent.com/document/product/266/14059?from=11329?from_cn_redirect=1#.E5.AD.98.E5.82.A8.E5.9C.B0.E5.9F.9F.E6.AD.A5.E9.AA.A4).</li>
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// Category ID, which is used to categorize the media for management. A category can be created and its ID can be obtained by using the [CreateClass](https://intl.cloud.tencent.com/document/product/266/7812?from_cn_redirect=1) API.
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// The source context which is used to pass through the user request information. After `Procedure` is specified, the task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// Used to identify duplicate requests. After you send a request, if any request with the same `SessionId` has already been sent in the last three days (72 hours), an error message will be returned. `SessionId` contains up to 50 characters. If this parameter is not carried or is an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Reserved field for special purposes.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// Source context, which is used to pass through the user request information. The [upload callback](https://intl.cloud.tencent.com/document/product/266/7830?from_cn_redirect=1) API will return the value of this field. It can contain up to 250 characters.
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`
}

func (r *PullUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PullUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaUrl")
	delete(f, "SubAppId")
	delete(f, "MediaName")
	delete(f, "CoverUrl")
	delete(f, "Procedure")
	delete(f, "ExpireTime")
	delete(f, "StorageRegion")
	delete(f, "ClassId")
	delete(f, "SessionContext")
	delete(f, "SessionId")
	delete(f, "ExtInfo")
	delete(f, "SourceContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PullUploadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PullUploadResponseParams struct {
	// Video pull for upload task ID, which can be used to query the status of pull for upload task.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PullUploadResponse struct {
	*tchttp.BaseResponse
	Response *PullUploadResponseParams `json:"Response"`
}

func (r *PullUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// The metadata of the output video.
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// Playback address generated after pull for upload is completed.
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// If a video processing flow is specified when a video is pulled for upload, this parameter will be the ID of the task flow.
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or a blank string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

// Predefined struct for user
type PushUrlCacheRequestParams struct {
	// List of prefetched URLs. Up to 20 ones can be specified at a time.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type PushUrlCacheRequest struct {
	*tchttp.BaseRequest
	
	// List of prefetched URLs. Up to 20 ones can be specified at a time.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *PushUrlCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PushUrlCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PushUrlCacheResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PushUrlCacheResponse struct {
	*tchttp.BaseResponse
	Response *PushUrlCacheResponseParams `json:"Response"`
}

func (r *PushUrlCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefererAuthPolicy struct {
	// [Referer hotlink protection](https://intl.cloud.tencent.com/document/product/266/33985) status. Valid values:
	// <li>Enabled</li>
	// <li>Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Referer authentication method. Valid values:
	// <li>`Black`: blocklist. Any HTTP request carrying a referer in the `Referers` list will be rejected. </li>
	// <li>`White`: allowlist. Only HTTP requests carrying referers in the `Referers` list will be accepted.</li>
	// When `Status` is set to `Enabled`, `AuthType` must be specified.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// The list of referers (up to 20). When `Status` is set to `Enabled`, `Referers` cannot be empty. Enter domain names as referers.
	Referers []*string `json:"Referers,omitempty" name:"Referers"`

	// Whether to allow requests with empty referer to access this domain name. Valid values:
	// <li>`Yes`</li>
	// <li>`No`</li>
	// When `Status` is set to `Enabled`, `BlankRefererAllowed` must be specified.
	BlankRefererAllowed *string `json:"BlankRefererAllowed,omitempty" name:"BlankRefererAllowed"`
}

// Predefined struct for user
type RefreshUrlCacheRequestParams struct {
	// The URLs to purge. You can specify up to 20 URLs per request.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type RefreshUrlCacheRequest struct {
	*tchttp.BaseRequest
	
	// The URLs to purge. You can specify up to 20 URLs per request.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *RefreshUrlCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshUrlCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshUrlCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshUrlCacheResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefreshUrlCacheResponse struct {
	*tchttp.BaseResponse
	Response *RefreshUrlCacheResponseParams `json:"Response"`
}

func (r *RefreshUrlCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshUrlCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetProcedureTemplateRequestParams struct {
	// Task flow name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Parameter of video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Parameter of AI-based content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Parameter of AI-based content analysis task.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of AI-based content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type ResetProcedureTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Task flow name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Parameter of video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// Parameter of AI-based content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// Parameter of AI-based content analysis task.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// Type parameter of AI-based content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// [Subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ResetProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetProcedureTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetProcedureTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetProcedureTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ResetProcedureTemplateResponseParams `json:"Response"`
}

func (r *ResetProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetProcedureTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResolutionNameInfo struct {
	// Length of video short side in px.
	MinEdgeLength *uint64 `json:"MinEdgeLength,omitempty" name:"MinEdgeLength"`

	// Display name.
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ResourceTag struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

// Predefined struct for user
type RestoreMediaRequestParams struct {
	// The IDs of media files.
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// The number of days during which the restored files will remain available.
	RestoreDay *uint64 `json:"RestoreDay,omitempty" name:"RestoreDay"`

	// The retrieval mode. If the current storage class is ARCHIVE, the valid values for this parameter are as follows:
	// <li>Expedited: The files are made available in five minutes.</li>
	// <li>Standard: The files are made available in five hours.</li>
	// <li>Bulk: The files are made available in 12 hours.</li>
	// If the current storage class is DEEP ARCHIVE, the valid values for this parameter are as follows:
	// <li>Standard: The files are made available in 24 hours.</li>
	// <li>Bulk: The files are made available in 48 hours.</li>
	RestoreTier *string `json:"RestoreTier,omitempty" name:"RestoreTier"`

	// The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

type RestoreMediaRequest struct {
	*tchttp.BaseRequest
	
	// The IDs of media files.
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// The number of days during which the restored files will remain available.
	RestoreDay *uint64 `json:"RestoreDay,omitempty" name:"RestoreDay"`

	// The retrieval mode. If the current storage class is ARCHIVE, the valid values for this parameter are as follows:
	// <li>Expedited: The files are made available in five minutes.</li>
	// <li>Standard: The files are made available in five hours.</li>
	// <li>Bulk: The files are made available in 12 hours.</li>
	// If the current storage class is DEEP ARCHIVE, the valid values for this parameter are as follows:
	// <li>Standard: The files are made available in 24 hours.</li>
	// <li>Bulk: The files are made available in 48 hours.</li>
	RestoreTier *string `json:"RestoreTier,omitempty" name:"RestoreTier"`

	// The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *RestoreMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileIds")
	delete(f, "RestoreDay")
	delete(f, "RestoreTier")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreMediaResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestoreMediaResponse struct {
	*tchttp.BaseResponse
	Response *RestoreMediaResponseParams `json:"Response"`
}

func (r *RestoreMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreMediaTask struct {
	// File ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Original storage class
	OriginalStorageClass *string `json:"OriginalStorageClass,omitempty" name:"OriginalStorageClass"`

	// Target storage class. For temporary retrieval, the target storage class is the same as the original.
	TargetStorageClass *string `json:"TargetStorageClass,omitempty" name:"TargetStorageClass"`

	// Retrieval mode. Valid values:
	// <li>Expedited</li>
	// <li>Standard</li>
	// <li>Bulk</li>
	RestoreTier *string `json:"RestoreTier,omitempty" name:"RestoreTier"`

	// Validity period (days) for a temporary copy. `0` indicates permanent retrieval.
	RestoreDay *int64 `json:"RestoreDay,omitempty" name:"RestoreDay"`

	// This field has been disused.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// This field has been disused.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SampleSnapshotTaskInput struct {
	// Sampled screencapturing template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of up to 10 image or text watermarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet"`
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

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The fill mode, or the way of processing a screenshot when the configured aspect ratio is different from that of the source video. Valid values:
	// <li>stretch: Stretch the image frame by frame to fill the entire screen. The video image may become "squashed" or "stretched" after transcoding.</li>
	// <li>black: Keep the image's original aspect ratio and fill the blank space with black bars.</li>
	// <li>white: Keep the image’s original aspect ratio and fill the blank space with white bars.</li>
	// <li>gauss: Keep the image’s original aspect ratio and apply Gaussian blur to the blank space.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

// Predefined struct for user
type SearchMediaRequestParams struct {
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// File ID set. Any element in the set can be matched.
	// <li>Array length limit: 10.</li>
	// <li>ID length limit: 40 characters.</li>
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// Filename set. Filenames of media files are fuzzily matched. The higher the match rate, the higher-ranked the result.
	// <li>Filename length limit: 40 characters.</li>
	// <li>Array length limit: 10.</li>
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Filename prefix, which matches the filenames of media files.
	// <li>Filename prefix length limit: 20 characters.</li>
	// <li>Array length limit: 10.</li>
	NamePrefixes []*string `json:"NamePrefixes,omitempty" name:"NamePrefixes"`

	// File description set. Media file descriptions are fuzzily matched. The higher the match rate, the higher-ranked the result.
	// <li>Length limit for a single description: 100 characters</li>
	// <li>Array length limit: 10</li>
	Descriptions []*string `json:"Descriptions,omitempty" name:"Descriptions"`

	// Category ID set. The categories of the specified IDs and all subcategories in the set are matched.
	// <li>Array length limit: 10.</li>
	ClassIds []*int64 `json:"ClassIds,omitempty" name:"ClassIds"`

	// The tag set. A file is considered a match if it has any of the tags in the tag set.
	// <li>Tag length limit: 16 characters.</li>
	// <li>Array length limit: 10.</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// File type. Any element in the set can be matched.
	// <li>Video: video file</li>
	// <li>Audio: audio file</li>
	// <li>Image: image file</li>
	Categories []*string `json:"Categories,omitempty" name:"Categories"`

	// Media file source set. For valid values, please see [SourceType](https://intl.cloud.tencent.com/document/product/266/31773?from_cn_redirect=1#MediaSourceData).
	// <li>Array length limit: 10.</li>
	SourceTypes []*string `json:"SourceTypes,omitempty" name:"SourceTypes"`

	// The live stream code array. A media file will be returned if it matches any element in the array.
	// <li>Array length limit: 10</li>
	StreamIds []*string `json:"StreamIds,omitempty" name:"StreamIds"`

	// Unique ID of LVB recording file. Any element in the set can be matched.
	// <li>Array length limit: 10.</li>
	Vids []*string `json:"Vids,omitempty" name:"Vids"`

	// Matches files created within the time period.
	// <li>Includes specified start and end points in time.</li>
	CreateTime *TimeRange `json:"CreateTime,omitempty" name:"CreateTime"`

	// Files whose expiration time points are within the specified time range will be returned. Expired files will not be returned.
	// <li>The files whose expiration time points are on the start or end time of the specified range will also be returned.</li>
	ExpireTime *TimeRange `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Sorting order.
	// <li>Valid value of `Sort.Field`: CreateTime.</li>
	// <li>If `Text`, `Names`, or `Descriptions` is not empty, the `Sort.Field` field will not take effect, and the search results will be sorted by match rate.</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// <div id="p_offset">Start offset of a paged return. Default value: 0. Entries from No. "Offset" to No. "Offset + Limit - 1" will be returned.
	// <li>Value range: "Offset + Limit" cannot be more than 5,000. (For more information, please see <a href="#maxResultsDesc">Limit on the Number of Results Returned by API</a>)</li></div>
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// <div id="p_limit">Number of entries returned by a paged query. Default value: 10. Entries from No. "Offset" to No. "Offset + Limit - 1" will be returned.
	// <li>Value range: "Offset + Limit" cannot be more than 5,000. (For more information, please see <a href="#maxResultsDesc">Limit on the Number of Results Returned by API</a>)</li></div>
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies information entry that needs to be returned for all media files. Multiple entries can be specified simultaneously. N starts from 0. If this field is left empty, all information entries will be returned by default. Valid values:
	// <li>basicInfo (basic video information).</li>
	// <li>metaData (video metadata).</li>
	// <li>transcodeInfo (result information of video transcoding).</li>
	// <li>animatedGraphicsInfo (result information of animated image generating task).</li>
	// <li>imageSpriteInfo (image sprite information).</li>
	// <li>snapshotByTimeOffsetInfo (point-in-time screenshot information).</li>
	// <li>sampleSnapshotInfo (sampled screenshot information).</li>
	// <li>keyFrameDescInfo (timestamp information).</li>
	// <li>adaptiveDynamicStreamingInfo (information of adaptive bitrate streaming).</li>
	// <li>miniProgramReviewInfo (WeChat Mini Program audit information).</li>
	Filters []*string `json:"Filters,omitempty" name:"Filters"`

	// Regions where media files are stored, such as `ap-chongqing`. For more regions, see [Storage Regions](https://intl.cloud.tencent.com/document/product/266/9760?from_cn_redirect=1#.E5.B7.B2.E6.94.AF.E6.8C.81.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8).
	// <li>Length limit for a single region: 20 characters</li>
	// <li>Array length limit: 20</li>
	StorageRegions []*string `json:"StorageRegions,omitempty" name:"StorageRegions"`

	// An array of storage classes. Valid values:
	// <li>STANDARD</li>
	// <li>STANDARD_IA</li>
	// <li>ARCHIVE</li>
	// <li>DEEP_ARCHIVE</li>
	StorageClasses []*string `json:"StorageClasses,omitempty" name:"StorageClasses"`

	// (This is not recommended. `Names`, `NamePrefixes`, or `Descriptions` should be used instead)
	// Search text, which fuzzily matches the media file name or description. The more matching items and the higher the match rate, the higher-ranked the result. It can contain up to 64 characters.
	Text *string `json:"Text,omitempty" name:"Text"`

	// (This is not recommended. `SourceTypes` should be used instead)
	// Media file source. For valid values, please see [SourceType](https://intl.cloud.tencent.com/document/product/266/31773?from_cn_redirect=1#MediaSourceData).
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// (Not recommended. Consider using `StreamIds` instead.)
	// The live stream code.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// (This is not recommended. `Vids` should be used instead)
	// Unique ID of LVB recording file.
	Vid *string `json:"Vid,omitempty" name:"Vid"`

	// (This is not recommended. `CreateTime` should be used instead)
	// Start time in the creation time range.
	// <li>After or at the start time.</li>
	// <li>If `CreateTime.After` also exists, it will be used first.</li>
	// <li>In ISO 8601 format. For more information, please see [ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).</li>
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// (This is not recommended. `CreateTime` should be used instead)
	// End time in the creation time range.
	// <li>Before the end time.</li>
	// <li>If `CreateTime.Before` also exists, it will be used first.</li>
	// <li>In ISO 8601 format. For more information, please see [ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).</li>
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type SearchMediaRequest struct {
	*tchttp.BaseRequest
	
	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// File ID set. Any element in the set can be matched.
	// <li>Array length limit: 10.</li>
	// <li>ID length limit: 40 characters.</li>
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// Filename set. Filenames of media files are fuzzily matched. The higher the match rate, the higher-ranked the result.
	// <li>Filename length limit: 40 characters.</li>
	// <li>Array length limit: 10.</li>
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Filename prefix, which matches the filenames of media files.
	// <li>Filename prefix length limit: 20 characters.</li>
	// <li>Array length limit: 10.</li>
	NamePrefixes []*string `json:"NamePrefixes,omitempty" name:"NamePrefixes"`

	// File description set. Media file descriptions are fuzzily matched. The higher the match rate, the higher-ranked the result.
	// <li>Length limit for a single description: 100 characters</li>
	// <li>Array length limit: 10</li>
	Descriptions []*string `json:"Descriptions,omitempty" name:"Descriptions"`

	// Category ID set. The categories of the specified IDs and all subcategories in the set are matched.
	// <li>Array length limit: 10.</li>
	ClassIds []*int64 `json:"ClassIds,omitempty" name:"ClassIds"`

	// The tag set. A file is considered a match if it has any of the tags in the tag set.
	// <li>Tag length limit: 16 characters.</li>
	// <li>Array length limit: 10.</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// File type. Any element in the set can be matched.
	// <li>Video: video file</li>
	// <li>Audio: audio file</li>
	// <li>Image: image file</li>
	Categories []*string `json:"Categories,omitempty" name:"Categories"`

	// Media file source set. For valid values, please see [SourceType](https://intl.cloud.tencent.com/document/product/266/31773?from_cn_redirect=1#MediaSourceData).
	// <li>Array length limit: 10.</li>
	SourceTypes []*string `json:"SourceTypes,omitempty" name:"SourceTypes"`

	// The live stream code array. A media file will be returned if it matches any element in the array.
	// <li>Array length limit: 10</li>
	StreamIds []*string `json:"StreamIds,omitempty" name:"StreamIds"`

	// Unique ID of LVB recording file. Any element in the set can be matched.
	// <li>Array length limit: 10.</li>
	Vids []*string `json:"Vids,omitempty" name:"Vids"`

	// Matches files created within the time period.
	// <li>Includes specified start and end points in time.</li>
	CreateTime *TimeRange `json:"CreateTime,omitempty" name:"CreateTime"`

	// Files whose expiration time points are within the specified time range will be returned. Expired files will not be returned.
	// <li>The files whose expiration time points are on the start or end time of the specified range will also be returned.</li>
	ExpireTime *TimeRange `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Sorting order.
	// <li>Valid value of `Sort.Field`: CreateTime.</li>
	// <li>If `Text`, `Names`, or `Descriptions` is not empty, the `Sort.Field` field will not take effect, and the search results will be sorted by match rate.</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// <div id="p_offset">Start offset of a paged return. Default value: 0. Entries from No. "Offset" to No. "Offset + Limit - 1" will be returned.
	// <li>Value range: "Offset + Limit" cannot be more than 5,000. (For more information, please see <a href="#maxResultsDesc">Limit on the Number of Results Returned by API</a>)</li></div>
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// <div id="p_limit">Number of entries returned by a paged query. Default value: 10. Entries from No. "Offset" to No. "Offset + Limit - 1" will be returned.
	// <li>Value range: "Offset + Limit" cannot be more than 5,000. (For more information, please see <a href="#maxResultsDesc">Limit on the Number of Results Returned by API</a>)</li></div>
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies information entry that needs to be returned for all media files. Multiple entries can be specified simultaneously. N starts from 0. If this field is left empty, all information entries will be returned by default. Valid values:
	// <li>basicInfo (basic video information).</li>
	// <li>metaData (video metadata).</li>
	// <li>transcodeInfo (result information of video transcoding).</li>
	// <li>animatedGraphicsInfo (result information of animated image generating task).</li>
	// <li>imageSpriteInfo (image sprite information).</li>
	// <li>snapshotByTimeOffsetInfo (point-in-time screenshot information).</li>
	// <li>sampleSnapshotInfo (sampled screenshot information).</li>
	// <li>keyFrameDescInfo (timestamp information).</li>
	// <li>adaptiveDynamicStreamingInfo (information of adaptive bitrate streaming).</li>
	// <li>miniProgramReviewInfo (WeChat Mini Program audit information).</li>
	Filters []*string `json:"Filters,omitempty" name:"Filters"`

	// Regions where media files are stored, such as `ap-chongqing`. For more regions, see [Storage Regions](https://intl.cloud.tencent.com/document/product/266/9760?from_cn_redirect=1#.E5.B7.B2.E6.94.AF.E6.8C.81.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8).
	// <li>Length limit for a single region: 20 characters</li>
	// <li>Array length limit: 20</li>
	StorageRegions []*string `json:"StorageRegions,omitempty" name:"StorageRegions"`

	// An array of storage classes. Valid values:
	// <li>STANDARD</li>
	// <li>STANDARD_IA</li>
	// <li>ARCHIVE</li>
	// <li>DEEP_ARCHIVE</li>
	StorageClasses []*string `json:"StorageClasses,omitempty" name:"StorageClasses"`

	// (This is not recommended. `Names`, `NamePrefixes`, or `Descriptions` should be used instead)
	// Search text, which fuzzily matches the media file name or description. The more matching items and the higher the match rate, the higher-ranked the result. It can contain up to 64 characters.
	Text *string `json:"Text,omitempty" name:"Text"`

	// (This is not recommended. `SourceTypes` should be used instead)
	// Media file source. For valid values, please see [SourceType](https://intl.cloud.tencent.com/document/product/266/31773?from_cn_redirect=1#MediaSourceData).
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// (Not recommended. Consider using `StreamIds` instead.)
	// The live stream code.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// (This is not recommended. `Vids` should be used instead)
	// Unique ID of LVB recording file.
	Vid *string `json:"Vid,omitempty" name:"Vid"`

	// (This is not recommended. `CreateTime` should be used instead)
	// Start time in the creation time range.
	// <li>After or at the start time.</li>
	// <li>If `CreateTime.After` also exists, it will be used first.</li>
	// <li>In ISO 8601 format. For more information, please see [ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).</li>
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// (This is not recommended. `CreateTime` should be used instead)
	// End time in the creation time range.
	// <li>Before the end time.</li>
	// <li>If `CreateTime.Before` also exists, it will be used first.</li>
	// <li>In ISO 8601 format. For more information, please see [ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).</li>
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *SearchMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "FileIds")
	delete(f, "Names")
	delete(f, "NamePrefixes")
	delete(f, "Descriptions")
	delete(f, "ClassIds")
	delete(f, "Tags")
	delete(f, "Categories")
	delete(f, "SourceTypes")
	delete(f, "StreamIds")
	delete(f, "Vids")
	delete(f, "CreateTime")
	delete(f, "ExpireTime")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "StorageRegions")
	delete(f, "StorageClasses")
	delete(f, "Text")
	delete(f, "SourceType")
	delete(f, "StreamId")
	delete(f, "Vid")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchMediaResponseParams struct {
	// Number of eligible entries.
	// <li>Maximum value: 5000. If the number of eligible entries is greater than 5,000, this field will return 5,000 instead of the actual number.</li>
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Media file information list
	MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchMediaResponse struct {
	*tchttp.BaseResponse
	Response *SearchMediaResponseParams `json:"Response"`
}

func (r *SearchMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SegmentConfigureInfo struct {
	// Switch of video splitting recognition task. Valid values:
	// <li>ON: enables intelligent video splitting recognition task;</li>
	// <li>OFF: disables intelligent video splitting recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type SegmentConfigureInfoForUpdate struct {
	// Switch of video splitting recognition task. Valid values:
	// <li>ON: enables intelligent video splitting recognition task;</li>
	// <li>OFF: disables intelligent video splitting recognition task.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

// Predefined struct for user
type SimpleHlsClipRequestParams struct {
	// URL of the HLS video in VOD that needs to be clipped.
	Url *string `json:"Url,omitempty" name:"Url"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Start offset time of clipping in seconds. Default value: 0, which means to clip from the beginning of the video. A negative number indicates how many seconds from the end of the video clipping will start at. For example, -10 means that clipping will start at the 10th second from the end.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End offset time of clipping in seconds. Default value: 0, which means to clip till the end of the video. A negative number indicates how many seconds from the end of the video clipping will end. For example, -10 means that clipping will end at the 10th second from the end.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Whether to store the video clip persistently. 0: no (default), 1: yes.
	IsPersistence *int64 `json:"IsPersistence,omitempty" name:"IsPersistence"`
}

type SimpleHlsClipRequest struct {
	*tchttp.BaseRequest
	
	// URL of the HLS video in VOD that needs to be clipped.
	Url *string `json:"Url,omitempty" name:"Url"`

	// <b>The VOD [subapplication](https://intl.cloud.tencent.com/document/product/266/14574?from_cn_redirect=1) ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.</b>
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Start offset time of clipping in seconds. Default value: 0, which means to clip from the beginning of the video. A negative number indicates how many seconds from the end of the video clipping will start at. For example, -10 means that clipping will start at the 10th second from the end.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// End offset time of clipping in seconds. Default value: 0, which means to clip till the end of the video. A negative number indicates how many seconds from the end of the video clipping will end. For example, -10 means that clipping will end at the 10th second from the end.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Whether to store the video clip persistently. 0: no (default), 1: yes.
	IsPersistence *int64 `json:"IsPersistence,omitempty" name:"IsPersistence"`
}

func (r *SimpleHlsClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SimpleHlsClipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "SubAppId")
	delete(f, "StartTimeOffset")
	delete(f, "EndTimeOffset")
	delete(f, "IsPersistence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SimpleHlsClipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SimpleHlsClipResponseParams struct {
	// Address of clipped video.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Metadata of clipped video. Currently, `Size`, `Rotate`, `VideoDuration`, and `AudioDuration` fields use default value with no actual data.
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// Unique ID of a video clip for persistent storage.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SimpleHlsClipResponse struct {
	*tchttp.BaseResponse
	Response *SimpleHlsClipResponseParams `json:"Response"`
}

func (r *SimpleHlsClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// screenshot specification. For more information, please see [Parameter Template for Time Point Screencapturing](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// Result information of screencapturing.
	SnapshotInfoSet []*SnapshotByTimeOffset2017 `json:"SnapshotInfoSet,omitempty" name:"SnapshotInfoSet"`
}

type SnapshotByTimeOffsetTaskInput struct {
	// Time point screencapturing template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// The list of screenshot time points. “s” and “%” formats are supported:
	// <li>When a time point string ends with “s”, its unit is second. For example, “3.5s” means the 3.5th second of the video.</li>
	// <li>When a time point string ends with “%”, it represents the percentage of the video’s duration. For example, “10%” means that the time point is at the 10% of the video’s entire duration.</li>
	ExtTimeOffsetSet []*string `json:"ExtTimeOffsetSet,omitempty" name:"ExtTimeOffsetSet"`

	// List of time points for screencapturing in <font color=red>milliseconds</font>.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitempty" name:"TimeOffsetSet"`

	// List of up to 10 image or text watermarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet"`
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

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The fill mode, or the way of processing a screenshot when the configured aspect ratio is different from that of the source video. Valid values:
	// <li>stretch: Stretch the image frame by frame to fill the entire screen. The video image may become "squashed" or "stretched" after transcoding.</li>
	// <li>black: Keep the image's original aspect ratio and fill the blank space with black bars.</li>
	// <li>white: Keep the image’s original aspect ratio and fill the blank space with white bars.</li>
	// <li>gauss: Keep the image’s original aspect ratio and apply Gaussian blur to the blank space.</li>
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
	Data []*TaskStatDataItem `json:"Data,omitempty" name:"Data"`
}

type SplitMediaOutputConfig struct {
	// Name of an output file. This parameter can contain up to 64 characters, and will be generated by the system if it is left empty.
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// Output file format. Valid values: mp4 (default), hls.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Category ID, which is used to categorize the media file for management. You can use [CreateClass](https://intl.cloud.tencent.com/document/product/266/7812?from_cn_redirect=1) API to create a category and get the category ID.
	// <li>Default value: 0, which means other categories.</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// Expiration time of an output file. After passing the expiration time, the file will be deleted. There is no expiration time set for a file by default. The time is in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?lang=en&pg=).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type SplitMediaTask struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task flow status. Valid values:
	// <li>PROCESSING: processing</li>
	// <li>FINISH: finished</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error code. An empty string indicates the task is successful; other values indicate failure. For details, see [Video Processing Error Codes](https://intl.cloud.tencent.com/zh/document/product/266/39145).
	ErrCodeExt *string `json:"ErrCodeExt,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; other values indicate failure. You're not recommended to use this parameter, but to use the new parameter `ErrCodeExt`.
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Error information.
	Message *string `json:"Message,omitempty" name:"Message"`

	// List of video splitting task details.
	FileInfoSet []*SplitMediaTaskSegmentInfo `json:"FileInfoSet,omitempty" name:"FileInfoSet"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1000 characters.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or set to an empty string, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type SplitMediaTaskInput struct {
	// Video ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Offset of the video splitting start time in seconds.
	// <li>If this parameter is left empty or set to 0, the transcoded video will start at the same time as the original video.</li>
	// <li>If this parameter is set to a positive number (n for example), the transcoded video will start at the nth second of the original video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the transcoded video will start at the nth second before the end of the original video.</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// Offset of the video splitting end time in seconds.
	// <li>If this parameter is left empty or set to 0, the transcoded video will end at the same time as the original video.</li>
	// <li>If this parameter is set to a positive number (n for example), the transcoded video will end at the nth second of the original video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the transcoded video will end at the nth second before the end of the original video.</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// [Task flow template](https://intl.cloud.tencent.com/document/product/266/33931?lang=en&pg=) name, which should be entered if you want to perform a task flow on the generated new video.
	ProcedureName *string `json:"ProcedureName,omitempty" name:"ProcedureName"`

	// Output information of a video splitting task.
	OutputConfig *SplitMediaOutputConfig `json:"OutputConfig,omitempty" name:"OutputConfig"`
}

type SplitMediaTaskSegmentInfo struct {
	// Input information of a video splitting task.
	Input *SplitMediaTaskInput `json:"Input,omitempty" name:"Input"`

	// Output information of a video splitting task.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Output *TaskOutputMediaInfo `json:"Output,omitempty" name:"Output"`

	// If a video processing flow is specified when a video splitting task is initiated, this field will be the task flow ID.
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`
}

type StatDataItem struct {
	// Start time of data time range in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). For example, if the time granularity is 1-day, `2018-12-01T00:00:00+08:00` represents the time range between December 1, 2018 (inclusive) and December 2, 2018 (not inclusive).
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
	// Source of media material for sticker segment, which can be:
	// <li>ID of VOD media files</li>
	// <li>Download URL of other media files</li>
	// Note: when a download URL of other media files is used as the material source and access control (such as hotlink protection) is enabled, the URL needs to carry access control parameters (such as hotlink protection signature).
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
	ImageOperations []*ImageTransform `json:"ImageOperations,omitempty" name:"ImageOperations"`
}

type StorageRegionInfo struct {
	// Storage region.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Description of the storage region.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether storage is enabled in the region. Valid values:
	// <li>opened: Enabled</li>
	// <li>unopened: Not enabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Whether the region is the default storage region. Valid values: true, false.
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type StorageStatData struct {
	// VOD storage billing region. Valid values:
	// <li>Chinese Mainland</li>
	// <li>Outside Chinese Mainland</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// Current total storage capacity in bytes.
	TotalStorage *uint64 `json:"TotalStorage,omitempty" name:"TotalStorage"`

	// Current STANDARD_IA storage capacity in bytes.
	InfrequentStorage *uint64 `json:"InfrequentStorage,omitempty" name:"InfrequentStorage"`

	// Current STANDARD storage capacity in bytes.
	StandardStorage *uint64 `json:"StandardStorage,omitempty" name:"StandardStorage"`

	// Current ARCHIVE storage usage in bytes
	ArchiveStorage *uint64 `json:"ArchiveStorage,omitempty" name:"ArchiveStorage"`

	// Current DEEP ARCHIVE storage usage in bytes
	DeepArchiveStorage *uint64 `json:"DeepArchiveStorage,omitempty" name:"DeepArchiveStorage"`
}

type SubAppIdInfo struct {
	// Subapplication ID.
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// Subapplication name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subapplication overview.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Subapplication creation time of task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Subapplication status. Valid values:
	// <li>On: enabled</li>
	// <li>Off: disabled</li>
	// <li>Destroying: terminating</li>
	// <li>Destroyed: terminated</li>
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
	// TESHD transcoding type. Valid values: <li>TEHD-100</li> <li>OFF (default)</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Maximum bitrate, which is valid when `Type` is `TESHD`.
	// If this parameter is left blank or 0 is entered, there will be no upper limit for bitrate.
	MaxVideoBitrate *uint64 `json:"MaxVideoBitrate,omitempty" name:"MaxVideoBitrate"`
}

type TEHDConfigForUpdate struct {
	// TESHD transcoding type. Valid values: <li>TEHD-100</li> <li>OFF (default)</li>
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

type TaskOutputMediaInfo struct {
	// Media file ID.
	FileId *string `json:"FileId,omitempty" name:"FileId"`


	MediaBasicInfo *MediaBasicInfo `json:"MediaBasicInfo,omitempty" name:"MediaBasicInfo"`
}

type TaskSimpleInfo struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task status. Valid values: `WAITING` (waiting), `PROCESSING` (processing), `FINISH` (completed)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Video ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Task type. Valid values:
	// <li>Procedure: video processing task;</li>
	// <li>EditMedia: video editing task</li>
	// <li>WechatDistribute: release on WeChat task.</li>
	// Task types compatible with v2017:
	// <li>Transcode: transcoding task;</li>
	// <li>SnapshotByTimeOffset: video screencapturing task</li>
	// <li>Concat: video splicing task;</li>
	// <li>Clip: video clipping task;</li>
	// <li>ImageSprites: image sprite generating task.</li>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Creation time of task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Start time of task execution in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). If the task has not been started yet, this field will be empty.
	BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

	// End time of task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I). If the task has not been completed yet, this field will be empty.
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// ID used for deduplication if there was a request with the same ID in the last seven days.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Source context, which is used to pass through the user request information.
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
}

type TaskStatData struct {
	// The task type.
	// <li>Transcoding: General transcoding</li>
	// <li>Transcoding-TESHD: TESHD transcoding</li>
	// <li>Editing: Video editing</li>
	// <li>Editing-TESHD: TESHD editing</li>
	// <li>AdaptiveBitrateStreaming: Adaptive bitrate streaming</li>
	// <li>ContentAudit: Content moderation</li>
	// <li>RemoveWatermark: Watermark removal</li>
	// <li>Transcode: Transcoding, including general transcoding, TESHD transcoding, and video editing. This value is not recommended.</li>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Task statistics overview (usage unit: second).
	Summary []*TaskStatDataItem `json:"Summary,omitempty" name:"Summary"`

	// The detailed statistics of different tasks.
	// Transcoding:
	// <li>Remuxing</li>
	// <li>Audio</li>
	// <li>Standard.H264.SD</li>
	// <li>Standard.H264.HD</li>
	// <li>Standard.H264.FHD</li>
	// <li>Standard.H264.2K</li>
	// <li>Standard.H264.4K</li>
	// <li>Standard.H265.SD</li>
	// <li>Standard.H265.HD</li>
	// <li>Standard.H265.FHD</li>
	// <li>Standard.H265.2K</li>
	// <li>Standard.H265.4K</li>
	// <li>TESHD-10.H264.SD</li>
	// <li>TESHD-10.H264.HD</li>
	// <li>TESHD-10.H264.FHD</li>
	// <li>TESHD-10.H264.2K</li>
	// <li>TESHD-10.H264.4K</li>
	// <li>TESHD-10.H265.SD</li>
	// <li>TESHD-10.H265.HD</li>
	// <li>TESHD-10.H265.FHD</li>
	// <li>TESHD-10.H265.2K</li>
	// <li>TESHD-10.H265.4K</li>
	// <li>Edit.Audio</li>
	// <li>Edit.H264.SD</li>
	// <li>Edit.H264.HD</li>
	// <li>Edit.H264.FHD</li>
	// <li>Edit.H264.2K</li>
	// <li>Edit.H264.4K</li>
	// <li>Edit.H265.SD</li>
	// <li>Edit.H265.HD</li>
	// <li>Edit.H265.FHD</li>
	// <li>Edit.H265.2K</li>
	// <li>Edit.H265.4K</li>
	// <li>Edit.TESHD-10.H264.SD</li>
	// <li>Edit.TESHD-10.H264.HD</li>
	// <li>Edit.TESHD-10.H264.FHD</li>
	// <li>Edit.TESHD-10.H264.2K</li>
	// <li>Edit.TESHD-10.H264.4K</li>
	// <li>Edit.TESHD-10.H265.SD</li>
	// <li>Edit.TESHD-10.H265.HD</li>
	// <li>Edit.TESHD-10.H265.FHD</li>
	// <li>Edit.TESHD-10.H265.2K</li>
	// <li>Edit.TESHD-10.H265.4K</li>
	// Watermark removal:
	// <li>480P: 640 x 480 and below</li>
	// <li>720P: 1280 x 720 and below</li>
	// <li>1080P: 1920 x 1080 and below</li>
	// <li>2K: 2560 x 1440 and below</li>
	// <li>4K: 3840 x 2160 and below</li>
	// <li>8K: 7680 x 4320 and below</li>
	Details []*SpecificationDataItem `json:"Details,omitempty" name:"Details"`
}

type TaskStatDataItem struct {
	// Start time of data time range in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F). For example, if the time granularity is 1-day, `2018-12-01T00:00:00+08:00` represents the time range between December 1, 2018 (inclusive) and December 2, 2018 (not inclusive).
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
	// Parameters for recognition of terrorism content in images
	// Note: This field may return `null`, indicating that no valid value can be found.
	ImgReviewInfo *TerrorismImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Parameters for OCR-based recognition of terrorism content
	// Note: This field may return `null`, indicating that no valid value can be found.
	OcrReviewInfo *TerrorismOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type TerrorismConfigureInfoForUpdate struct {
	// Parameters for recognition of terrorism content in images
	ImgReviewInfo *TerrorismImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// Parameters for OCR-based recognition of terrorism content
	OcrReviewInfo *TerrorismOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type TerrorismImgReviewTemplateInfo struct {
	// Whether to enable recognition of terrorism content in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for recognition of terrorism content in images. Results containing the specified labels are returned. If no labels are specified, all results are returned. Valid values:
	// <li>`guns`: weapons and guns</li>
	// <li>`crowd`: crowd</li>
	// <li>`bloody`: bloody scenes</li>
	// <li>`police`: police force</li>
	// <li>`banners`: terrorism flags</li>
	// <li>`militant`: militants</li>
	// <li>`explosion`: explosions and fires</li>
	// <li>`terrorists`: terrorists</li>
	// <li>`scenario`: terrorism images</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `90` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `80` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type TerrorismImgReviewTemplateInfoForUpdate struct {
	// Whether to enable recognition of terrorism content in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for recognition of terrorism content in images. Results containing the specified labels are returned. If no labels are specified, all results are returned. Valid values:
	// <li>`guns`: weapons and guns</li>
	// <li>`crowd`: crowd</li>
	// <li>`bloody`: bloody scenes</li>
	// <li>`police`: police force</li>
	// <li>`banners`: terrorism flags</li>
	// <li>`militant`: militants</li>
	// <li>`explosion`: explosions and fires</li>
	// <li>`terrorists`: terrorists</li>
	// <li>`scenario`: terrorism images</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type TerrorismOcrReviewTemplateInfo struct {
	// Whether to enable OCR-based recognition of terrorism content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type TerrorismOcrReviewTemplateInfoForUpdate struct {
	// Whether to enable OCR-based recognition of terrorism content. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
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

type TimeRange struct {
	// <li>After or at this time (start time).</li>
	// <li>In ISO 8601 format. For more information, please see [ISO Date Format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).</li>
	After *string `json:"After,omitempty" name:"After"`

	// <li>Earlier than this time (end time).</li>
	// <li>In ISO 8601 format. For more information, please see [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).</li>
	Before *string `json:"Before,omitempty" name:"Before"`
}

type TraceWatermarkInput struct {
	// The watermark template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type TranscodePlayInfo2017 struct {
	// Playback address.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Transcoding specification ID. For more information, please see [Transcoding Parameter Template](https://intl.cloud.tencent.com/document/product/266/33478?from_cn_redirect=1#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
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
	PlayInfoSet []*TranscodePlayInfo2017 `json:"PlayInfoSet,omitempty" name:"PlayInfoSet"`
}

type TranscodeTaskInput struct {
	// Video transcoding template ID.
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// List of up to 10 image or text watermarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// Digital watermark.
	TraceWatermark *TraceWatermarkInput `json:"TraceWatermark,omitempty" name:"TraceWatermark"`

	// List of video opening/closing credits configuration template IDs. You can enter up to 10 IDs.
	HeadTailSet []*HeadTailTaskInput `json:"HeadTailSet,omitempty" name:"HeadTailSet"`

	// List of blurs. Up to 10 ones can be supported.
	MosaicSet []*MosaicInput `json:"MosaicSet,omitempty" name:"MosaicSet"`

	// End time offset of a transcoded video, in seconds.
	// <li>If this parameter is left empty or set to 0, the transcoded video will end at the same time as the original video.</li>
	// <li>If this parameter is set to a positive number (n for example), the transcoded video will end at the nth second of the original video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the transcoded video will end at the nth second before the end of the original video.</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Start time offset of a transcoded video, in seconds.
	// <li>If this parameter is left empty or set to 0, the transcoded video will start at the same time as the original video.</li>
	// <li>If this parameter is set to a positive number (n for example), the transcoded video will start at the nth second of the original video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the transcoded video will start at the nth second before the end of the original video.</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`
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

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
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

type UrlSignatureAuthPolicy struct {
	// Whether to enable or disable [key hotlink protection](https://intl.cloud.tencent.com/document/product/266/33986). Valid values:
	// <li>`Enabled`: enable</li>
	// <li>`Disabled`: disable</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The key for generating the signature of [key hotlink protection](https://intl.cloud.tencent.com/document/product/266/33986).
	// `EncryptedKey` can contain 8-40 bytes, and cannot contain non-printable characters.
	EncryptedKey *string `json:"EncryptedKey,omitempty" name:"EncryptedKey"`
}

type UserDefineAsrTextReviewTemplateInfo struct {
	// Whether to enable custom ASR-based recognition. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for custom ASR-based recognition. Results containing the specified labels are returned. If no labels are specified, all results are returned. To filter by labels, specify the labels when adding keywords for custom ASR-based recognition.
	// Up to 10 labels are allowed, each containing no more than 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineAsrTextReviewTemplateInfoForUpdate struct {
	// Whether to enable custom ASR-based recognition. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for custom ASR-based recognition. Results containing the specified labels are returned. If no labels are specified, all results are returned. To filter by labels, specify the labels when adding keywords for custom ASR-based recognition.
	// Up to 10 labels are allowed, each containing no more than 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineConfigureInfo struct {
	// Parameters for custom facial recognition
	// Note: This field may return `null`, indicating that no valid value can be found.
	FaceReviewInfo *UserDefineFaceReviewTemplateInfo `json:"FaceReviewInfo,omitempty" name:"FaceReviewInfo"`

	// Parameters for custom ASR-based recognition
	// Note: This field may return `null`, indicating that no valid value can be found.
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Parameters for custom OCR-based recognition
	// Note: This field may return `null`, indicating that no valid value can be found.
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type UserDefineConfigureInfoForUpdate struct {
	// Parameters for custom facial recognition
	FaceReviewInfo *UserDefineFaceReviewTemplateInfoForUpdate `json:"FaceReviewInfo,omitempty" name:"FaceReviewInfo"`

	// Parameters for custom ASR-based recognition
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// Parameters for custom OCR-based recognition
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type UserDefineFaceReviewTemplateInfo struct {
	// Whether to enable custom facial recognition. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for custom facial recognition. Results containing the specified labels are returned. If no labels are specified, all results are returned. To filter by labels, specify the labels when adding custom facial libraries.
	// Up to 10 labels are allowed, each containing no more than 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `97` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `95` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineFaceReviewTemplateInfoForUpdate struct {
	// Whether to enable custom facial recognition. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for custom facial recognition. Results containing the specified labels are returned. If no labels are specified, all results are returned. To filter by labels, specify the labels when adding custom facial libraries.
	// Up to 10 labels are allowed, each containing no more than 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfo struct {
	// Whether to enable custom OCR-based recognition. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for custom OCR-based recognition. Results containing the specified labels are returned. If no labels are specified, all results are returned. To filter by labels, specify the labels when adding keywords for custom OCR-based recognition.
	// Up to 10 labels are allowed, each containing no more than 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. If this parameter is left empty, `100` will be used by default. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. If this parameter is left empty, `75` will be used by default. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfoForUpdate struct {
	// Whether to enable custom OCR-based recognition. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Filter labels for custom OCR-based recognition. Results containing the specified labels are returned. If no labels are specified, all results are returned. To filter by labels, specify the labels when adding keywords for custom OCR-based recognition.
	// Up to 10 labels are allowed, each containing no more than 16 characters.
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet"`

	// Confidence score threshold for determining that something should be blocked. If this threshold is reached, VOD will suggest that the content be blocked. Value range: 0-100
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// Confidence score threshold for human review. If this threshold is reached, human review is needed. Value range: 0-100
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type VideoTemplateInfo struct {
	// The video codec. Valid values:
	// <li>libx264: H.264</li>
	// <li>libx265: H.265</li>
	// <li>av1: AOMedia Video 1</li>
	// <li>H.266: H.266</li>
	// <font color=red>Notes:</font>
	// <li>The AOMedia Video 1 and H.266 codecs can only be used for MP4 files.</li>
	// <li> Only CRF is supported for H.266 currently.</li>
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Video frame rate in Hz. Value range: [0,100].
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

	// Fill type, the way of processing a screenshot when the configured aspect ratio is different from that of the source video. Valid values:
	// <li>stretch: stretches the video image frame by frame to fill the screen. The video image may become "squashed" or "stretched" after transcoding.</li>
	// <li>black: fills the uncovered area with black color, without changing the image's aspect ratio.</li>
	// <li>white: fills the uncovered area with white color, without changing the image's aspect ratio.</li>
	// <li>gauss: applies Gaussian blur to the uncovered area, without changing the image's aspect ratio.</li>
	// Default value: black
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// The video constant rate factor (CRF). Value range: 1-51.
	// 
	// <font color=red>Notes:</font>
	// <li>If this parameter is specified, CRF encoding will be used and the bitrate parameter will be ignored.</li>
	// <li>If `Codec` is `H.266`, this parameter is required (`28` is recommended).</li>
	// <li>We don’t recommend using this parameter unless you have special requirements.</li>
	Vcrf *uint64 `json:"Vcrf,omitempty" name:"Vcrf"`

	// I-frame interval in frames. Valid values: 0 and 1-100000.
	// When this parameter is set to 0 or left empty, `Gop` will be automatically set.
	Gop *uint64 `json:"Gop,omitempty" name:"Gop"`
}

type VideoTemplateInfoForUpdate struct {
	// The video codec. Valid values:
	// <li>libx264: H.264</li>
	// <li>libx265: H.265</li>
	// <li>av1: AOMedia Video 1</li>
	// <li>H.266: H.266</li>
	// <font color=red>Notes:</font>
	// <li>The AOMedia Video 1 and H.266 codecs can only be used for MP4 files.</li>
	// <li> Only CRF is supported for H.266 currently.</li>
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// Video frame rate in Hz. Value range: [0,100].
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

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. Valid values:
	// <li>stretch: stretches video image frame by frame to fill the screen. The video image may become "squashed" or "stretched" after transcoding.</li>
	// <li>black: fills the uncovered area with black color, without changing the image's aspect ratio.</li>
	// <li>white: fills the uncovered area with white color, without changing the image's aspect ratio.</li>
	// <li>gauss: applies Gaussian blur to the uncovered area, without changing the image's aspect ratio.</li>
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// The video constant rate factor (CRF). Value range: 1-51. `0` means to disable this parameter.
	// 
	// <font color=red>Notes:</font>
	// <li>If this parameter is specified, CRF encoding will be used and the bitrate parameter will be ignored.</li>
	// <li>If `Codec` is `H.266`, this parameter is required (`28` is recommended).</li>
	// <li>We don’t recommend using this parameter unless you have special requirements.</li>
	Vcrf *uint64 `json:"Vcrf,omitempty" name:"Vcrf"`

	// I-frame interval in frames. Valid values: 0 and 1-100000.
	// When this parameter is set to 0 or left empty, `Gop` will be automatically set.
	Gop *uint64 `json:"Gop,omitempty" name:"Gop"`
}

type VideoTrackItem struct {
	// Source of media material for video segment, which can be:
	// <li>ID of VOD media files</li>
	// <li>Download URL of other media files.</li>
	// Note: when a download URL of other media files is used as the material source and access control (such as hotlink protection) is enabled, the URL needs to carry access control parameters (such as hotlink protection signature).
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
	ImageOperations []*ImageTransform `json:"ImageOperations,omitempty" name:"ImageOperations"`

	// Operation on audio such as muting.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AudioOperations []*AudioTransform `json:"AudioOperations,omitempty" name:"AudioOperations"`
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

	// Text content, which contains up to 100 characters. Set this parameter only when the watermark type is text.
	// VOD does not support adding text watermarks on screenshots.
	TextContent *string `json:"TextContent,omitempty" name:"TextContent"`

	// SVG content, which contains up to 2,000,000 characters. Set this parameter only when the watermark type is SVG.
	// VOD does not support adding SVG watermarks on screenshots.
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

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Origin position. Valid values:
	// <li>topLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>topRight: the origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>bottomLeft: the origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>bottomRight: the origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`
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