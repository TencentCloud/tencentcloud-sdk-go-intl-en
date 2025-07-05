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

package v20190612

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AIAnalysisTemplateItem struct {
	// Unique ID of intelligent analysis template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Intelligent analysis template name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Intelligent analysis template description.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Control parameter of intelligent categorization task.
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// Control parameter of intelligent tagging task.
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// Control parameter of intelligent cover generating task.
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// Control parameter of intelligent frame-specific tagging task.
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The template type. Valid values:
	// * Preset
	// * Custom
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type AIRecognitionTemplateItem struct {
	// Unique ID of a video content recognition template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of a video content recognition template.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description of a video content recognition template.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Face recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// Full text recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// Text keyword recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// Full speech recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// Speech keyword recognition control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`

	// Voice translation control parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranslateConfigure *TranslateConfigureInfo `json:"TranslateConfigure,omitnil,omitempty" name:"TranslateConfigure"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The template type. Valid values:
	// * Preset
	// * Custom
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type Activity struct {
	// Atomic task type.
	// <li>input: start node</li>
	// <li>output: end node</li>
	// <li>action-trans: transcoding</li>
	// <li>action-samplesnapshot: sampled screenshot</li>
	// <li>action-AIAnalysis: analysis</li>
	// <li>action-AIRecognition: recognition</li>
	// <li>action-aiReview: review</li>
	// <li>action-animated-graphics: conversion to GIF</li>
	// <li>action-image-sprite: image sprite</li>
	// <li>action-snapshotByTimeOffset: time point screenshot</li>
	// <li>action-adaptive-substream: adaptive bitrate stream</li>
	// <li>action-AIQualityControl: media quality inspection</li>
	// <li>action-SmartSubtitles: smart subtitle</li>
	// 
	// 
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	ActivityType *string `json:"ActivityType,omitnil,omitempty" name:"ActivityType"`

	// The indexes of the subsequent actions.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReardriveIndex []*int64 `json:"ReardriveIndex,omitnil,omitempty" name:"ReardriveIndex"`

	// The parameters of a subtask.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActivityPara *ActivityPara `json:"ActivityPara,omitnil,omitempty" name:"ActivityPara"`
}

type ActivityPara struct {
	// A transcoding task.
	TranscodeTask *TranscodeTaskInput `json:"TranscodeTask,omitnil,omitempty" name:"TranscodeTask"`

	// An animated screenshot generation task.
	AnimatedGraphicTask *AnimatedGraphicTaskInput `json:"AnimatedGraphicTask,omitnil,omitempty" name:"AnimatedGraphicTask"`

	// A time point screenshot task.
	SnapshotByTimeOffsetTask *SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTask,omitnil,omitempty" name:"SnapshotByTimeOffsetTask"`

	// A sampled screenshot task.
	SampleSnapshotTask *SampleSnapshotTaskInput `json:"SampleSnapshotTask,omitnil,omitempty" name:"SampleSnapshotTask"`

	// An image sprite screenshot task.
	ImageSpriteTask *ImageSpriteTaskInput `json:"ImageSpriteTask,omitnil,omitempty" name:"ImageSpriteTask"`

	// An adaptive bitrate streaming task.
	AdaptiveDynamicStreamingTask *AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTask,omitnil,omitempty" name:"AdaptiveDynamicStreamingTask"`

	// A content moderation task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// A content analysis task.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// A content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// Media quality inspection task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	QualityControlTask *AiQualityControlTaskInput `json:"QualityControlTask,omitnil,omitempty" name:"QualityControlTask"`

	// Smart subtitle task.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SmartSubtitlesTask *SmartSubtitlesTaskInput `json:"SmartSubtitlesTask,omitnil,omitempty" name:"SmartSubtitlesTask"`
}

type ActivityResItem struct {
	// The result of a transcoding task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitnil,omitempty" name:"TranscodeTask"`

	// The result of an animated image generating task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitnil,omitempty" name:"AnimatedGraphicTask"`

	// The result of a time point screenshot task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitnil,omitempty" name:"SnapshotByTimeOffsetTask"`

	// The result of a sampled screenshot task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitnil,omitempty" name:"SampleSnapshotTask"`

	// The result of an image sprite task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitnil,omitempty" name:"ImageSpriteTask"`

	// The result of an adaptive bitrate streaming task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AdaptiveDynamicStreamingTask *MediaProcessTaskAdaptiveDynamicStreamingResult `json:"AdaptiveDynamicStreamingTask,omitnil,omitempty" name:"AdaptiveDynamicStreamingTask"`

	// The result of a content recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecognitionTask *ScheduleRecognitionTaskResult `json:"RecognitionTask,omitnil,omitempty" name:"RecognitionTask"`

	// The result of a content moderation task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReviewTask *ScheduleReviewTaskResult `json:"ReviewTask,omitnil,omitempty" name:"ReviewTask"`

	// The result of a content analysis task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnalysisTask *ScheduleAnalysisTaskResult `json:"AnalysisTask,omitnil,omitempty" name:"AnalysisTask"`

	// Media quality inspection task output.
	// Note: This field may return null, indicating that no valid values can be obtained.
	QualityControlTask *ScheduleQualityControlTaskResult `json:"QualityControlTask,omitnil,omitempty" name:"QualityControlTask"`

	// Smart subtitle task output.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SmartSubtitlesTask *ScheduleSmartSubtitleTaskResult `json:"SmartSubtitlesTask,omitnil,omitempty" name:"SmartSubtitlesTask"`
}

type ActivityResult struct {
	// Atomic task type.
	// <Li>Transcode: transcoding</li>
	// <Li>SampleSnapshot: sampled screenshot</li>
	// <Li>AnimatedGraphics: conversion to GIF</li>
	// <Li>SnapshotByTimeOffset: time point screenshot</li>
	// <Li>ImageSprites: image sprite</li>
	// <Li>AdaptiveDynamicStreaming: adaptive bitrate stream</li>
	// <Li>AiContentReview: content review</li>
	// <Li>AIRecognition: intelligent recognition</li>
	// <Li>AIAnalysis: intelligent analysis</li>
	// <li>AiQualityControl: media quality inspection.</li>
	// 
	// <Li>SmartSubtitles: smart subtitle</li>
	ActivityType *string `json:"ActivityType,omitnil,omitempty" name:"ActivityType"`

	// The execution results of the subtasks of the scheme.
	ActivityResItem *ActivityResItem `json:"ActivityResItem,omitnil,omitempty" name:"ActivityResItem"`
}

type AdaptiveDynamicStreamingInfoItem struct {
	// Adaptive bitrate streaming specification.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Container format. Valid values: HLS, MPEG-DASH.
	Package *string `json:"Package,omitnil,omitempty" name:"Package"`

	// Playback address.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Storage location of adaptive bitrate streaming files.
	Storage *TaskOutputStorage `json:"Storage,omitnil,omitempty" name:"Storage"`
}

type AdaptiveDynamicStreamingTaskInput struct {
	// Adaptive dynamic streaming template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Watermark list. Multiple image or text watermarks up to a maximum of 10 are supported.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// Target storage for files after adaptive dynamic streaming. If left blank, it inherits the upper-level OutputStorage value.
	// Note: This field may return null, indicating that no valid value can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Output path for the manifest file after adaptive dynamic streaming. It can be either a relative path or an absolute path.
	// If you need to define an output path, the path must end with `.{format}`. Refer to [Filename Variable Description](https://intl.cloud.tencent.com/document/product/862/37039?from_cn_redirect=1) for variable names.
	// Example of relative path:
	// <li>filename_{variable name}.{format}</li>
	// <li>filename.{format}</li>
	// Example of absolute path:
	// <li>/custom path/filename_{variable name}.{format}</li>
	// If not filled in, it is a relative path by default: {inputName}_adaptiveDynamicStreaming_{definition}.{format}.
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// After adaptive dynamic streaming, the output path of substream files can only be a relative path. If not filled in, it is a relative path by default: `{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}.{format}`.
	SubStreamObjectName *string `json:"SubStreamObjectName,omitnil,omitempty" name:"SubStreamObjectName"`

	// After adaptive dynamic streaming (for HLS only), the output path of segment files can only be a relative path. If not filled in, it is a relative path by default: `{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}_{segmentNumber}.{format}`.
	SegmentObjectName *string `json:"SegmentObjectName,omitnil,omitempty" name:"SegmentObjectName"`

	// Subtitle file to be inserted.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AddOnSubtitles []*AddOnSubtitle `json:"AddOnSubtitles,omitnil,omitempty" name:"AddOnSubtitles"`

	// Drm information.
	// Note: This field may return null, indicating that no valid value can be obtained.
	DrmInfo *DrmInfo `json:"DrmInfo,omitnil,omitempty" name:"DrmInfo"`

	// Adaptive transcoding template type.
	// Common: audio-video.
	// PureAudio: audio-only.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DefinitionType *string `json:"DefinitionType,omitnil,omitempty" name:"DefinitionType"`
}

type AdaptiveDynamicStreamingTemplate struct {
	// Unique ID of an adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Name of an adaptive bitrate streaming template.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description of an adaptive bitrate streaming template.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Adaptive bitrate streaming format. Valid values:
	// <li>HLS;</li>
	// <li>MPEG-DASH.</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Parameter information of input streams for transcoding to adaptive bitrate streaming. Up to 10 streams can be input.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// Whether to prohibit transcoding from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// Creation time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#I).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Whether it is an audio-only template. 0: video template. 1: audio-only template.Note: This field may return null, indicating that no valid values can be obtained.
	PureAudio *uint64 `json:"PureAudio,omitnil,omitempty" name:"PureAudio"`

	// HLS segment type. Valid values:
	// <li>ts-segment: HLS+TS segment.</li>
	// <li>ts-byterange: HLS+TS byte range.</li>
	// <li>mp4-segment: HLS+MP4 segment.</li>
	// <li>mp4-byterange: HLS+MP4 byte range.</li>
	// <li>ts-packed-audio: TS+Packed audio.</li>
	// <li>mp4-packed-audio: MP4+Packed audio.</li>
	// Default value: ts-segment.
	// 
	// Note: The HLS segment format for adaptive bitrate streaming is based on this field.Note: This field may return null, indicating that no valid values can be obtained.
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
}

type AdaptiveStreamTemplate struct {
	// Audio parameter information.
	Audio *AudioTemplateInfo `json:"Audio,omitnil,omitempty" name:"Audio"`

	// Video parameter information.
	Video *VideoTemplateInfo `json:"Video,omitnil,omitempty" name:"Video"`

	// Whether to remove audio stream. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	RemoveAudio *uint64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// Whether to remove video stream. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	RemoveVideo *uint64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// Audio parameter information list.
	// The parameter is only used when merging multiple audio tracks in adaptive bitrate transcoding. the maximum length of the parameter array is 64.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	AudioList []*AudioTemplateInfo `json:"AudioList,omitnil,omitempty" name:"AudioList"`
}

type AddOnSubtitle struct {
	// The mode. Valid values:
	// <li>`subtitle-stream`: Add a subtitle track.</li>
	// <li>`close-caption-708`: Embed CEA-708 subtitles in SEI frames.</li>
	// <li>`close-caption-608`: Embed CEA-608 subtitles in SEI frames.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The subtitle file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Subtitle *MediaInputInfo `json:"Subtitle,omitnil,omitempty" name:"Subtitle"`

	// Subtitle name.
	// Note: supports Chinese characters, letters, digits, spaces, underscores (_), hyphens (-), periods (.), and parentheses. Max 64 characters.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SubtitleName *string `json:"SubtitleName,omitnil,omitempty" name:"SubtitleName"`
}

type AiAnalysisResult struct {
	// Task type. Valid values:
	// <li>Classification: intelligent classification.</li>
	// <li>Cover: intelligent thumbnail generating.</li>
	// <li>Tag: intelligent tagging.</li>
	// <li>FrameTag: intelligent frame-by-frame tagging.</li>
	// <li>Highlight: intelligent highlights generating.</li>
	// 
	// <li>DeLogo: intelligent removal.</li>
	// <li>Description: large model summarization.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Query result of intelligent categorization task in video content analysis, which is valid if task type is `Classification`.
	ClassificationTask *AiAnalysisTaskClassificationResult `json:"ClassificationTask,omitnil,omitempty" name:"ClassificationTask"`

	// Query result of intelligent cover generating task in video content analysis, which is valid if task type is `Cover`.
	CoverTask *AiAnalysisTaskCoverResult `json:"CoverTask,omitnil,omitempty" name:"CoverTask"`

	// Query result of intelligent tagging task in video content analysis, which is valid if task type is `Tag`.
	TagTask *AiAnalysisTaskTagResult `json:"TagTask,omitnil,omitempty" name:"TagTask"`

	// Query result of intelligent frame-specific tagging task in video content analysis, which is valid if task type is `FrameTag`.
	FrameTagTask *AiAnalysisTaskFrameTagResult `json:"FrameTagTask,omitnil,omitempty" name:"FrameTagTask"`

	// The result of a highlight generation task. This parameter is valid if `Type` is `Highlight`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HighlightTask *AiAnalysisTaskHighlightResult `json:"HighlightTask,omitnil,omitempty" name:"HighlightTask"`

	// The query result of an intelligent removal task for video analysis, which is valid when the task type is DeLogo.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeLogoTask *AiAnalysisTaskDelLogoResult `json:"DeLogoTask,omitnil,omitempty" name:"DeLogoTask"`

	// The query result of a splitting task for video analysis, which is valid when the task type is SegmentRecognition.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SegmentTask *AiAnalysisTaskSegmentResult `json:"SegmentTask,omitnil,omitempty" name:"SegmentTask"`

	// The query result of an opening and closing segments recognition task for video analysis, which is valid when the task type is HeadTailRecognition.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HeadTailTask *AiAnalysisTaskHeadTailResult `json:"HeadTailTask,omitnil,omitempty" name:"HeadTailTask"`

	// The query result of a video analysis summarization task, which is valid when the task type is Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DescriptionTask *AiAnalysisTaskDescriptionResult `json:"DescriptionTask,omitnil,omitempty" name:"DescriptionTask"`

	// The query result of a landscape-to-portrait task for video analysis, which is valid when the task type is HorizontalToVertical.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HorizontalToVerticalTask *AiAnalysisTaskHorizontalToVerticalResult `json:"HorizontalToVerticalTask,omitnil,omitempty" name:"HorizontalToVerticalTask"`
}

type AiAnalysisTaskClassificationInput struct {
	// Intelligent video categorization template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskClassificationOutput struct {
	// List of intelligently generated video categories.
	ClassificationSet []*MediaAiAnalysisClassificationItem `json:"ClassificationSet,omitnil,omitempty" name:"ClassificationSet"`
}

type AiAnalysisTaskClassificationResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input of intelligent categorization task.
	Input *AiAnalysisTaskClassificationInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of intelligent categorization task.
	Output *AiAnalysisTaskClassificationOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiAnalysisTaskCoverInput struct {
	// Intelligent video cover generating template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskCoverOutput struct {
	// List of intelligently generated covers.
	CoverSet []*MediaAiAnalysisCoverItem `json:"CoverSet,omitnil,omitempty" name:"CoverSet"`

	// Storage location of intelligently generated cover.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`
}

type AiAnalysisTaskCoverResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input of intelligent cover generating task.
	Input *AiAnalysisTaskCoverInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of intelligent cover generating task.
	Output *AiAnalysisTaskCoverOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiAnalysisTaskDelLogoInput struct {
	// Intelligent removal template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskDelLogoOutput struct {
	// Path of a file after removal.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Storage location of a file after removal.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Path of a subtitle file extracted from a video.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginSubtitlePath *string `json:"OriginSubtitlePath,omitnil,omitempty" name:"OriginSubtitlePath"`

	// Path of a subtitle translation file extracted from a video.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranslateSubtitlePath *string `json:"TranslateSubtitlePath,omitnil,omitempty" name:"TranslateSubtitlePath"`
}

type AiAnalysisTaskDelLogoResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. `0`: Task successful. Other values: Task failed.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Intelligent removal task input.
	Input *AiAnalysisTaskDelLogoInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Intelligent removal task output.Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskDelLogoOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiAnalysisTaskDescriptionInput struct {
	// Intelligent description template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskDescriptionOutput struct {
	// Intelligent video description list.
	DescriptionSet []*MediaAiAnalysisDescriptionItem `json:"DescriptionSet,omitnil,omitempty" name:"DescriptionSet"`
}

type AiAnalysisTaskDescriptionResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. `0`: Task successful. Other values: Task failed.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Intelligent description task input.
	Input *AiAnalysisTaskDescriptionInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Intelligent description task output.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskDescriptionOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiAnalysisTaskFrameTagInput struct {
	// Intelligent frame-specific video tagging template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskFrameTagOutput struct {
	// List of frame-specific video tags.
	SegmentSet []*MediaAiAnalysisFrameTagSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiAnalysisTaskFrameTagResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input of intelligent frame-specific tagging task.
	Input *AiAnalysisTaskFrameTagInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of intelligent frame-specific tagging task.
	Output *AiAnalysisTaskFrameTagOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiAnalysisTaskHeadTailInput struct {
	// Opening and closing segments recognition template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskHeadTailOutput struct {
	// Opening segment PTS.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HeadTimeOffset *float64 `json:"HeadTimeOffset,omitnil,omitempty" name:"HeadTimeOffset"`

	// Closing segment PTS.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TailTimeOffset *float64 `json:"TailTimeOffset,omitnil,omitempty" name:"TailTimeOffset"`
}

type AiAnalysisTaskHeadTailResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. `0`: Task successful. Other values: Task failed.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Opening and closing segments recognition task input.
	Input *AiAnalysisTaskHeadTailInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Opening and closing segments recognition task output.Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskHeadTailOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiAnalysisTaskHighlightInput struct {
	// The ID of the intelligent highlight generation template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskHighlightOutput struct {
	// A list of the highlight segments generated.
	HighlightSet []*MediaAiAnalysisHighlightItem `json:"HighlightSet,omitnil,omitempty" name:"HighlightSet"`

	// The storage location of the highlight segments.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`
}

type AiAnalysisTaskHighlightResult struct {
	// The task status. Valid values: `PROCESSING`, `SUCCESS`, `FAIL`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. `0`: The task succeeded; other values: The task failed.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// The error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input of the intelligent highlight generation task.
	Input *AiAnalysisTaskHighlightInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of the intelligent highlight generation task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskHighlightOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiAnalysisTaskHorizontalToVerticalInput struct {
	// Intelligent landscape-to-portrait template ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskHorizontalToVerticalOutput struct {
	// Intelligent landscape-to-portrait video list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Storage location of intelligent landscape-to-portrait videos.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Confidence.	
	// 	
	// Note: This field may return null, indicating that no valid values can be obtained.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type AiAnalysisTaskHorizontalToVerticalResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. 0: Task successful. Other values: Task failed.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Intelligent landscape-to-portrait task input.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Input *AiAnalysisTaskHorizontalToVerticalInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Intelligent landscape-to-portrait task output.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskHorizontalToVerticalOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiAnalysisTaskInput struct {
	// Video content analysis template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Additional parameter. Its value is a serialized JSON string.
	// Note: This parameter is used to meet customization requirements. References:
	// [Smart Erase Tutorial]: https://intl.cloud.tencent.com/document/product/862/101530?from_cn_redirect=1
	// [Video Splitting (Long Videos to Short Videos) Tutorial](https://intl.cloud.tencent.com/document/product/862/112098?from_cn_redirect=1)
	// [Intelligent Highlights Tutorial](https://intl.cloud.tencent.com/document/product/862/107280?from_cn_redirect=1)
	// [Horizontal-to-Vertical Video Transformation Tutorial](https://intl.cloud.tencent.com/document/product/862/112112?from_cn_redirect=1)
	// Note: This field may return null, indicating that no valid value can be obtained.
	ExtendedParameter *string `json:"ExtendedParameter,omitnil,omitempty" name:"ExtendedParameter"`
}

type AiAnalysisTaskSegmentInput struct {
	// Splitting task template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskSegmentOutput struct {
	// Intelligent splitting sub-segment list.
	SegmentSet []*SegmentRecognitionItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Video abstract, used for offline scenarios.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Abstract *string `json:"Abstract,omitnil,omitempty" name:"Abstract"`
}

type AiAnalysisTaskSegmentResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. `0`: Task successful. Other values: Task failed.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Splitting task input.
	Input *AiAnalysisTaskSegmentInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Splitting task output.Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiAnalysisTaskSegmentOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiAnalysisTaskTagInput struct {
	// Intelligent video tagging template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiAnalysisTaskTagOutput struct {
	// List of intelligently generated video tags.
	TagSet []*MediaAiAnalysisTagItem `json:"TagSet,omitnil,omitempty" name:"TagSet"`
}

type AiAnalysisTaskTagResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input of intelligent tagging task.
	Input *AiAnalysisTaskTagInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of intelligent tagging task.
	Output *AiAnalysisTaskTagOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiContentReviewResult struct {
	// Task type. Valid values:
	// <li>Porn (in images)</li>
	// <li>Terrorism (in images)</li>
	// <li>Political (in images)</li>
	// <li>Porn.Asr</li>
	// <li>Porn.Ocr</li>
	// <li>Political.Asr</li>
	// <li>Political.Ocr</li>
	// <li>Terrorism.Ocr</li>
	// <li>Prohibited.Asr</li>
	// <li>Prohibited.Ocr</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Sample rate, which indicates the number of video frames captured per second for audit
	SampleRate *float64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Audited video duration in seconds.
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Query result of an intelligent porn information detection in image task in video content audit, which is valid when task type is `Porn`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PornTask *AiReviewTaskPornResult `json:"PornTask,omitnil,omitempty" name:"PornTask"`

	// The result of detecting terrorism content in images, which is valid when the task type is `Terrorism`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TerrorismTask *AiReviewTaskTerrorismResult `json:"TerrorismTask,omitnil,omitempty" name:"TerrorismTask"`

	// The result of detecting politically sensitive information in images, which is valid when the task type is `Political`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PoliticalTask *AiReviewTaskPoliticalResult `json:"PoliticalTask,omitnil,omitempty" name:"PoliticalTask"`

	// Query result of an ASR-based porn information detection in text task in video content audit, which is valid when task type is `Porn.Asr`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PornAsrTask *AiReviewTaskPornAsrResult `json:"PornAsrTask,omitnil,omitempty" name:"PornAsrTask"`

	// Query result of an OCR-based porn information detection in text task in video content audit, which is valid when task type is `Porn.Ocr`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PornOcrTask *AiReviewTaskPornOcrResult `json:"PornOcrTask,omitnil,omitempty" name:"PornOcrTask"`

	// The result of detecting politically sensitive information based on ASR, which is valid when the task type is `Political.Asr`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PoliticalAsrTask *AiReviewTaskPoliticalAsrResult `json:"PoliticalAsrTask,omitnil,omitempty" name:"PoliticalAsrTask"`

	// The result of detecting politically sensitive information based on OCR, which is valid when the task type is `Political.Ocr`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PoliticalOcrTask *AiReviewTaskPoliticalOcrResult `json:"PoliticalOcrTask,omitnil,omitempty" name:"PoliticalOcrTask"`

	// The result of detecting terrorism content based on OCR, which is valid when task type is `Terrorism.Ocr`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TerrorismOcrTask *AiReviewTaskTerrorismOcrResult `json:"TerrorismOcrTask,omitnil,omitempty" name:"TerrorismOcrTask"`

	// Query result of ASR-based prohibited information detection in speech task in video content audit, which is valid if task type is `Prohibited.Asr`.
	ProhibitedAsrTask *AiReviewTaskProhibitedAsrResult `json:"ProhibitedAsrTask,omitnil,omitempty" name:"ProhibitedAsrTask"`

	// Query result of OCR-based prohibited information detection in text task in video content audit, which is valid if task type is `Prohibited.Ocr`.
	ProhibitedOcrTask *AiReviewTaskProhibitedOcrResult `json:"ProhibitedOcrTask,omitnil,omitempty" name:"ProhibitedOcrTask"`
}

type AiContentReviewTaskInput struct {
	// Video content audit template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiParagraphInfo struct {
	// Segment summary.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// Segment title.
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Segment keywords.
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// Segmentation start time point, in seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// Segmentation end time point, in seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type AiQualityControlTaskInput struct {
	// Media quality inspection template ID.
	// You can directly use a preset template or customize a template in the console. The preset templates are as follows:
	// - 10: Enable all quality inspection items.
	// - 20: Only enable quality inspection items corresponding to format diagnosis.
	// - 30: Only enable quality inspection items corresponding to no-reference scoring.
	// - 40: Only enable quality inspection items corresponding to screen quality.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The channel extension parameter, which is a serialized JSON string.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChannelExtPara *string `json:"ChannelExtPara,omitnil,omitempty" name:"ChannelExtPara"`
}

type AiRecognitionResult struct {
	// The task type. Valid values:
	// <li>FaceRecognition: Face recognition</li>
	// <li>AsrWordsRecognition: Speech keyword recognition</li>
	// <li>OcrWordsRecognition: Text keyword recognition</li>
	// <li>AsrFullTextRecognition: Full speech recognition</li>
	// <li>OcrFullTextRecognition: Full text recognition</li>
	// <li>TransTextRecognition: Speech translation</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Face recognition result, which is valid when `Type` is 
	//  `FaceRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FaceTask *AiRecognitionTaskFaceResult `json:"FaceTask,omitnil,omitempty" name:"FaceTask"`

	// Speech keyword recognition result, which is valid when `Type` is
	//  `AsrWordsRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrWordsTask *AiRecognitionTaskAsrWordsResult `json:"AsrWordsTask,omitnil,omitempty" name:"AsrWordsTask"`

	// Full speech recognition result, which is valid when `Type` is
	//  `AsrFullTextRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrFullTextTask *AiRecognitionTaskAsrFullTextResult `json:"AsrFullTextTask,omitnil,omitempty" name:"AsrFullTextTask"`

	// Text keyword recognition result, which is valid when `Type` is
	//  `OcrWordsRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrWordsTask *AiRecognitionTaskOcrWordsResult `json:"OcrWordsTask,omitnil,omitempty" name:"OcrWordsTask"`

	// Full text recognition result, which is valid when `Type` is
	//  `OcrFullTextRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrFullTextTask *AiRecognitionTaskOcrFullTextResult `json:"OcrFullTextTask,omitnil,omitempty" name:"OcrFullTextTask"`

	// The translation result. This parameter is valid only if `Type` is
	//  `TransTextRecognition`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransTextTask *AiRecognitionTaskTransTextResult `json:"TransTextTask,omitnil,omitempty" name:"TransTextTask"`

	// Object recognition result, which is valid when Type is
	// 
	// ObjectRecognition.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectTask *AiRecognitionTaskObjectResult `json:"ObjectTask,omitnil,omitempty" name:"ObjectTask"`
}

type AiRecognitionTaskAsrFullTextResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input information of a full speech recognition task.
	Input *AiRecognitionTaskAsrFullTextResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output information of a full speech recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskAsrFullTextResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// Task progress.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiRecognitionTaskAsrFullTextResultInput struct {
	// Full speech recognition template ID.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrFullTextResultOutput struct {
	// List of full speech recognition segments.
	SegmentSet []*AiRecognitionTaskAsrFullTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Subtitles file address.
	SubtitlePath *string `json:"SubtitlePath,omitnil,omitempty" name:"SubtitlePath"`

	// Subtitles file storage location.
	//
	// Deprecated: OutputStorage is deprecated.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`
}

type AiRecognitionTaskAsrFullTextSegmentItem struct {
	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Recognized text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Word timestamp information.
	Wordlist []*WordResult `json:"Wordlist,omitnil,omitempty" name:"Wordlist"`
}

type AiRecognitionTaskAsrWordsResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input information of a speech keyword recognition task.
	Input *AiRecognitionTaskAsrWordsResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output information of a speech keyword recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskAsrWordsResultOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiRecognitionTaskAsrWordsResultInput struct {
	// Speech keyword recognition template ID.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrWordsResultItem struct {
	// Speech keyword.
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// List of time segments that contain the speech keyword.
	SegmentSet []*AiRecognitionTaskAsrWordsSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskAsrWordsResultOutput struct {
	// Speech keyword recognition result set.
	ResultSet []*AiRecognitionTaskAsrWordsResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`
}

type AiRecognitionTaskAsrWordsSegmentItem struct {
	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type AiRecognitionTaskFaceResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input information of a face recognition task.
	Input *AiRecognitionTaskFaceResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output information of a face recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskFaceResultOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiRecognitionTaskFaceResultInput struct {
	// Face recognition template ID.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskFaceResultItem struct {
	// Unique ID of a figure.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Figure library type, indicating to which figure library the recognized figure belongs:
	// <li>Default: Default figure library;</li>
	// <li>UserDefine: Custom figure library.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Name of a figure.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Result set of segments that contain a figure.
	SegmentSet []*AiRecognitionTaskFaceSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// The person’s gender.
	// <li>Male</li>
	// <li>Female</li>
	// Note: This field may return null, indicating that no valid value can be obtained.
	Gender *string `json:"Gender,omitnil,omitempty" name:"Gender"`

	// The person’s birth date.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// The person’s job or job title.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Profession *string `json:"Profession,omitnil,omitempty" name:"Profession"`

	// The college the person graduated from.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SchoolOfGraduation *string `json:"SchoolOfGraduation,omitnil,omitempty" name:"SchoolOfGraduation"`

	// The person’s profile.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Abstract *string `json:"Abstract,omitnil,omitempty" name:"Abstract"`

	// The person’s place of birth.
	// Note: This field may return null, indicating that no valid value can be obtained.
	PlaceOfBirth *string `json:"PlaceOfBirth,omitnil,omitempty" name:"PlaceOfBirth"`

	// Whether the person is a politician or artist.
	// <li>Politician</li>
	// <li>Artist</li>
	// Note: This field may return null, indicating that no valid value can be obtained.
	PersonType *string `json:"PersonType,omitnil,omitempty" name:"PersonType"`

	// Sensitivity
	// <li>Normal</li>
	// <li>Sensitive</li>
	// Note: This field may return null, indicating that no valid value can be obtained.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// The screenshot URL.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type AiRecognitionTaskFaceResultOutput struct {
	// Intelligent face recognition result set.
	ResultSet []*AiRecognitionTaskFaceResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`
}

type AiRecognitionTaskFaceSegmentItem struct {
	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Zone coordinates of a recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskInput struct {
	// Intelligent video recognition template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// User extension field, which does not need to be filled in for general scenarios.
	UserExtPara *string `json:"UserExtPara,omitnil,omitempty" name:"UserExtPara"`
}

type AiRecognitionTaskObjectResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. `0`: Task successful. Other values: Task failed.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Object recognition task input.
	Input *AiRecognitionTaskObjectResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Object recognition task output.Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskObjectResultOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiRecognitionTaskObjectResultInput struct {
	// Object recognition template ID.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskObjectResultItem struct {
	// Name of a recognized object.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// List of segments that contain the object.
	SegmentSet []*AiRecognitionTaskObjectSeqmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskObjectResultOutput struct {
	// Intelligent object recognition result set.
	ResultSet []*AiRecognitionTaskObjectResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`
}

type AiRecognitionTaskObjectSeqmentItem struct {
	// Start time offset of a recognized segment, in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognized segment, in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Confidence of a recognized segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Zone coordinates of the recognition result. An array contains four elements: [x1, y1, x2, y2], representing the horizontal and vertical coordinates of the top-left and bottom-right corners, respectively.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskOcrFullTextResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input information of a full text recognition task.
	Input *AiRecognitionTaskOcrFullTextResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output information of a full text recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskOcrFullTextResultOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiRecognitionTaskOcrFullTextResultInput struct {
	// Full text recognition template ID.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrFullTextResultOutput struct {
	// Full text recognition result set.
	SegmentSet []*AiRecognitionTaskOcrFullTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskOcrFullTextSegmentItem struct {
	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Recognition segment result set.
	TextSet []*AiRecognitionTaskOcrFullTextSegmentTextItem `json:"TextSet,omitnil,omitempty" name:"TextSet"`
}

type AiRecognitionTaskOcrFullTextSegmentTextItem struct {
	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Zone coordinates of a recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// Recognized text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type AiRecognitionTaskOcrWordsResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input information of a text keyword recognition task.
	Input *AiRecognitionTaskOcrWordsResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output information of a text keyword recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskOcrWordsResultOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiRecognitionTaskOcrWordsResultInput struct {
	// Text keyword recognition template ID.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrWordsResultItem struct {
	// Text keyword.
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// List of segments that contain a text keyword.
	SegmentSet []*AiRecognitionTaskOcrWordsSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiRecognitionTaskOcrWordsResultOutput struct {
	// Text keyword recognition result set.
	ResultSet []*AiRecognitionTaskOcrWordsResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`
}

type AiRecognitionTaskOcrWordsSegmentItem struct {
	// Start time offset of a recognition segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognition segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Confidence of a recognition segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Zone coordinates of a recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type AiRecognitionTaskTransTextResult struct {
	// The task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value indicates the task has failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// The error code. `0` indicates the task is successful; other values indicate the task has failed. This parameter is not recommended. Please use `ErrCodeExt` instead.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// The error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input of the translation task.
	Input *AiRecognitionTaskTransTextResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of the translation task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiRecognitionTaskTransTextResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// Task progress.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type AiRecognitionTaskTransTextResultInput struct {
	// The translation template ID.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiRecognitionTaskTransTextResultOutput struct {
	// The translated segments.
	SegmentSet []*AiRecognitionTaskTransTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// The subtitle URL.
	SubtitlePath *string `json:"SubtitlePath,omitnil,omitempty" name:"SubtitlePath"`
}

type AiRecognitionTaskTransTextSegmentItem struct {
	// The confidence score for a segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The start time offset (seconds) of a segment.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// The end time offset (seconds) of a segment.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// The text transcript.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// The translation.
	Trans *string `json:"Trans,omitnil,omitempty" name:"Trans"`

	// Word timestamp information.
	Wordlist []*WordResult `json:"Wordlist,omitnil,omitempty" name:"Wordlist"`
}

type AiReviewPoliticalAsrTaskInput struct {
	// The template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPoliticalAsrTaskOutput struct {
	// The confidence score for the ASR-based detection of sensitive information. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The suggestion for handling the sensitive information detected based on ASR. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// The video segments that contain sensitive information detected based on ASR.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiReviewPoliticalOcrTaskInput struct {
	// The template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPoliticalOcrTaskOutput struct {
	// The confidence score for the OCR-based detection of sensitive information. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The suggestion for handling the sensitive information detected based on OCR. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// The video segments that contain sensitive information detected based on OCR.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiReviewPoliticalTaskInput struct {
	// The template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPoliticalTaskOutput struct {
	// The confidence score for the detection of sensitive information. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The suggestion for handling the sensitive information detected. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// The labels for the detected sensitive content. The relationship between the values of this parameter and those of the `LabelSet` parameter in [PoliticalImgReviewTemplateInfo](https://intl.cloud.tencent.com/document/api/862/37615?from_cn_redirect=1#AiReviewPoliticalTaskOutput) is as follows:
	// violation_photo:
	// <li>violation_photo (banned icons)</li>
	// Other values (politician/entertainment/sport/entrepreneur/scholar/celebrity/military):
	// <li>politician</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// The video segments that contain sensitive information.
	SegmentSet []*MediaContentReviewPoliticalSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiReviewPornAsrTaskInput struct {
	// ID of a porn information detection template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPornAsrTaskOutput struct {
	// Score of the ASR-detected porn information in text from 0 to 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for the ASR-detected porn information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// List of video segments that contain the ASR-detected porn information in text.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiReviewPornOcrTaskInput struct {
	// ID of a porn information detection template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPornOcrTaskOutput struct {
	// Score of the OCR-detected porn information in text from 0 to 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for the OCR-detected porn information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// List of video segments that contain the OCR-detected porn information in text.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiReviewPornTaskInput struct {
	// The ID of a porn detection template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewPornTaskOutput struct {
	// Score of the detected porn information in video from 0 to 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for the detected porn information. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Tag of the detected porn information in video. Valid values:
	// <li>porn: Porn.</li>
	// <li>sexy: Sexiness.</li>
	// <li>vulgar: Vulgarity.</li>
	// <li>intimacy: Intimacy.</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// List of video segments that contain the detected porn information.
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiReviewProhibitedAsrTaskInput struct {
	// Prohibited information detection template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewProhibitedAsrTaskOutput struct {
	// Score of ASR-detected prohibited information in speech between 0 and 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for ASR-detected prohibited information in speech. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// List of video segments that contain the ASR-detected prohibited information in speech.
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiReviewProhibitedOcrTaskInput struct {
	// Prohibited information detection template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewProhibitedOcrTaskOutput struct {
	// Score of OCR-detected prohibited information in text between 0 and 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for OCR-detected prohibited information in text. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// List of video segments that contain the OCR-detected prohibited information in text.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiReviewTaskPoliticalAsrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input parameter for ASR-based detection of politically sensitive information.
	Input *AiReviewPoliticalAsrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of ASR-based detection of politically sensitive information.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Output *AiReviewPoliticalAsrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalOcrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input parameter for OCR-based detection of politically sensitive information.
	Input *AiReviewPoliticalOcrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of OCR-based detection of politically sensitive information.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Output *AiReviewPoliticalOcrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input parameter for sensitive information detection.
	Input *AiReviewPoliticalTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of sensitive information detection.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Output *AiReviewPoliticalTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTaskPornAsrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input for an ASR-based porn information detection in text task during content audit.
	Input *AiReviewPornAsrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of an ASR-based porn information detection in text task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPornAsrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTaskPornOcrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input for an OCR-based porn information detection in text task during content audit.
	Input *AiReviewPornOcrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of an OCR-based porn information detection in text task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPornOcrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTaskPornResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input for a porn information detection task during content audit.
	Input *AiReviewPornTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of a porn information detection task during content audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *AiReviewPornTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTaskProhibitedAsrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input of ASR-based prohibited information detection in speech task in content audit
	Input *AiReviewProhibitedAsrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of ASR-based prohibited information detection in speech task in content audit
	Output *AiReviewProhibitedAsrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTaskProhibitedOcrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input of OCR-based prohibited information detection in text task in content audit
	Input *AiReviewProhibitedOcrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of OCR-based prohibited information detection in text task in content audit
	Output *AiReviewProhibitedOcrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismOcrResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0: success; other values: failure.
	// <li>40000: invalid input parameter. Please check it;</li>
	// <li>60000: invalid source file (e.g., video data is corrupted). Please check whether the source file is normal;</li>
	// <li>70000: internal service error. Please try again.</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input parameter for OCR-based detection of terrorism content.
	Input *AiReviewTerrorismOcrTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of OCR-based detection of terrorism content.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Output *AiReviewTerrorismOcrTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input parameter for sensitive information detection.
	Input *AiReviewTerrorismTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of sensitive information detection.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Output *AiReviewTerrorismTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type AiReviewTerrorismOcrTaskInput struct {
	// The template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewTerrorismOcrTaskOutput struct {
	// The confidence score for the OCR-based detection of sensitive information. Value range: 1-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The suggestion for handling the sensitive information detected based on OCR. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// The video segments that contain sensitive information detected based on OCR.
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiReviewTerrorismTaskInput struct {
	// The template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type AiReviewTerrorismTaskOutput struct {
	// The confidence score for the detection of sensitive information. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The suggestion for handling the sensitive information detected. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// The labels for the detected sensitive content. Valid values:
	// <li>guns</li>
	// <li>crowd</li>
	// <li>police</li>
	// <li>bloody</li>
	// <li>banners (sensitive flags)</li>
	// <li>militant</li>
	// <li>explosion</li>
	// <li>terrorists</li>
	// <li>scenario (sensitive scenes) </li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// The video segments that contain sensitive information.
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type AiSampleFaceInfo struct {
	// Face image ID.
	FaceId *string `json:"FaceId,omitnil,omitempty" name:"FaceId"`

	// Face image address.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type AiSampleFaceOperation struct {
	// Operation type. Valid values: add, delete, reset. The `reset` operation will clear the existing face data of a figure and add `FaceContents` as the specified face data.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Face ID set. This field is required when `Type` is `delete`.
	FaceIds []*string `json:"FaceIds,omitnil,omitempty" name:"FaceIds"`

	// String set generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) the face image.
	// <li>This field is required when `Type` is `add` or `reset`;</li>
	// <li>Array length limit: 5 images.</li>
	// Note: The image must be a relatively clear full-face photo of a figure in at least 200 * 200 px.
	FaceContents []*string `json:"FaceContents,omitnil,omitempty" name:"FaceContents"`
}

type AiSampleFailFaceInfo struct {
	// Corresponds to incorrect image subscripts in the `FaceContents` input parameter, starting from 0.
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`

	// Error code. Valid values:
	// <li>0: Succeeded;</li>
	// <li>Other values: Failed.</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error description.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type AiSamplePerson struct {
	// Figure ID.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Name of a figure.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Figure description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Face information.
	FaceInfoSet []*AiSampleFaceInfo `json:"FaceInfoSet,omitnil,omitempty" name:"FaceInfoSet"`

	// Figure tag.
	TagSet []*string `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// Use case.
	UsageSet []*string `json:"UsageSet,omitnil,omitempty" name:"UsageSet"`

	// Creation time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AiSampleTagOperation struct {
	// Operation type. Valid values: add, delete, reset.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Tag. Length limit: 128 characters.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type AiSampleWord struct {
	// Keyword.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Keyword tag.
	TagSet []*string `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// Keyword use case.
	UsageSet []*string `json:"UsageSet,omitnil,omitempty" name:"UsageSet"`

	// Creation time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AiSampleWordInfo struct {
	// Keyword. Length limit: 20 characters.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Keyword tag
	// <li>Array length limit: 20 tags;</li>
	// <li>Tag length limit: 128 characters.</li>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type AnimatedGraphicTaskInput struct {
	// Animated image generating template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Start time of an animated image in a video in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time of an animated image in a video in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Target bucket of a generated animated image file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Output path of a file after animated image generating, which can be a relative or absolute path.
	// If you need to define an output path, the path must end with `.{format}`. For variable names, refer to [Filename Variable](https://intl.cloud.tencent.com/document/product/862/37039?from_cn_redirect=1).
	// Relative path example:
	// <li>Filename_{Variable name}.{format}.</li>
	// <li>Filename.{format}.</li>
	// Absolute path example:
	// <li>/Custom path/Filename_{Variable name}.{format}.</li>
	// If left empty, a relative path is used by default: `{inputName}_animatedGraphic_{definition}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`
}

type AnimatedGraphicsTemplate struct {
	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Name of an animated image generating template.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description of an animated image generating template.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Maximum value of the width (or long side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: Enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: Disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Animated image format.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Frame rate.
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Image quality.
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type ArtifactRepairConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The strength. Valid values:
	// <li>weak</li>
	// <li>strong</li>
	// Default value: weak.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type AsrFullTextConfigureInfo struct {
	// Switch of a full speech recognition task. Valid values:
	// <li>ON: Enables an intelligent full speech recognition task;</li>
	// <li>OFF: Disables an intelligent full speech recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Format of the generated subtitles file. If this parameter is left empty or an empty string is entered, no subtitles files will be generated. Valid value:
	// <li>vtt: Generates a WebVTT subtitles file.</li>
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`
}

type AsrFullTextConfigureInfoForUpdate struct {
	// Switch of a full speech recognition task. Valid values:
	// <li>ON: Enables an intelligent full speech recognition task;</li>
	// <li>OFF: Disables an intelligent full speech recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Format of the generated subtitles file. If an empty string is entered, no subtitles files will be generated. Valid value:
	// <li>vtt: Generates a WebVTT subtitles file.</li>
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`
}

type AsrHotWordsConfigure struct {
	// Hotword switch.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Hotword lexicon ID.
	// Note: This field may return null, indicating that no valid value can be obtained.
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`
}

type AsrHotwordsSet struct {
	// Hotword lexicon ID.
	// Note: This field may return null, indicating that no valid value can be obtained.
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`

	// Current hotword lexicon status. The value indicates the number of smart subtitle templates bound to this hotword lexicon.
	// If the value of Status is 0, it indicates that the hotword lexicon is not referenced by any smart subtitle template and that it can be deleted.
	// If the value of Status is not 0, it indicates that the hotword lexicon cannot be deleted.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Hotword lexicon name.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Number of hotwords in the hotword lexicon.
	// Note: This field may return null, indicating that no valid value can be obtained.
	WordCount *uint64 `json:"WordCount,omitnil,omitempty" name:"WordCount"`

	// Name of the uploaded hotword file.
	// Note: This field may return null, indicating that no valid value can be obtained.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Creation time of the hotword lexicon in ISO datetime format (UTC time). For example, 2006-01-02T15:04:05Z.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Creation time of the hotword lexicon in ISO datetime format (UTC time). For example, 2006-01-02T15:04:05Z.
	// Note: This field may return null, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 0: temporary hotword lexicon
	// 1: file-based hotword lexicon
	// Note: This field may return null, indicating that no valid value can be obtained.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type AsrHotwordsSetItem struct {
	// Hotword ID.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Hotword text.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Hotword weight. The value can be 11 or 100 or be in the range of 1 to 10.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type AsrWordsConfigureInfo struct {
	// Switch of a speech keyword recognition task. Valid values:
	// <li>ON: Enables a speech keyword recognition task;</li>
	// <li>OFF: Disables a speech keyword recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`
}

type AsrWordsConfigureInfoForUpdate struct {
	// Switch of a speech keyword recognition task. Valid values:
	// <li>ON: Enables a speech keyword recognition task;</li>
	// <li>OFF: Disables a speech keyword recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`
}

type AudioBeautifyConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>`ON`</li>
	// <li>`OFF` </li>
	// Default value: `ON`.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The audio improvement options. You can specify multiple options. Valid values:
	// <li>`declick`: Noise removal.</li>
	// <li>`deesser`: De-essing.</li>
	// Default: `declick`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`
}

type AudioDenoiseConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>`ON`</li>
	// <li>`OFF` </li>
	// Default value: `ON`.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type AudioEnhanceConfig struct {
	// The audio noise reduction configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Denoise *AudioDenoiseConfig `json:"Denoise,omitnil,omitempty" name:"Denoise"`

	// The audio separation configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Separate *AudioSeparateConfig `json:"Separate,omitnil,omitempty" name:"Separate"`

	// The volume equalization configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VolumeBalance *VolumeBalanceConfig `json:"VolumeBalance,omitnil,omitempty" name:"VolumeBalance"`

	// The audio improvement configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Beautify *AudioBeautifyConfig `json:"Beautify,omitnil,omitempty" name:"Beautify"`
}

type AudioSeparateConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>`ON`</li>
	// <li>`OFF` </li>
	// Default value: `ON`.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The scenario. Valid values:
	// <li>`normal`: Separate voice and background audio.</li>
	// <li>`music`: Separate vocals and instrumentals.</li>
	// Default value: `normal`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The output audio track. Valid values:
	// <li>`vocal`: Voice.</li>
	// <li>`background`: Output background audio if the scenario is `normal`, and output instrumentals if the scenario is `music`.</li>
	// Default value: `vocal`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Track *string `json:"Track,omitnil,omitempty" name:"Track"`
}

type AudioTemplateInfo struct {
	// Specifies the encoding format of the audio stream.
	// When audio transcoding is not needed, the optional values are:.
	// <li>copy.</li>
	// When the outer parameter Container is mp3, the valid values are:.
	// <li>mp3.</li>
	// When the outer parameter Container is ogg or flac, the valid values are:.
	// <li>flac.</li>
	// When the outer parameter Container is m4a, valid values are:.
	// <li>aac;</li>
	// <li>ac3.</li>
	// When the outer parameter Container is mp4 or flv, valid values are:.
	// <li>aac: more suitable for mp4;</li>.
	// <li>mp3: more suitable for flv;</li>.
	// <li>mp2.</li>
	// When the outer parameter Container is hls, valid values are:.
	// <li>aac;</li>
	// <li>mp3;</li>
	// <li>eac3: used when merging adaptive transcoding audio tracks.</li>.
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// The bitrate of the audio stream. value ranges from 0 to 26 and in the range of [26, 256]. measurement unit: kbps.
	// If the value is 0, the audio bitrate will be the same as that of the original audio.
	// Specifies that when using the TrackChannelInfo parameter for adaptive transcoding audio track merging, the valid values are:.
	// Cannot be set to 0.
	// 2). when Codec is aac, valid values: [26, 256].
	// 3). when Codec is ac3, valid values: [26, 640].
	// 4) when Codec is eac3, value range: [26, 6144]. remark: when SampleRate is 44100HZ, maximum value: 5644. when SampleRate is 48000HZ, maximum value: 6144.
	// 
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// Audio stream sampling rate. Different sampling rate options are provided for different encoding standards. For details, see [Audio/Video Transcoding Template](https://intl.cloud.tencent.com/document/product/862/77166?from_cn_redirect=1#f3b039f1-d817-4a96-b4e4-90132d31cd53).
	// Unit: Hz.
	// Note: Make sure that the sampling rate of the source audio stream is among the above options. Otherwise, transcoding may fail.
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Audio channel mode. Valid values:
	// <li>1: mono-channel.</li>
	// <li>2: dual-channel.</li>
	// <li>6: 5.1 surround sound.
	// <li>Default value: 2.
	// When the container format is audio (flac, ogg, mp3, and m4a), the audio channel cannot be set to 5.1 surround sound.
	AudioChannel *int64 `json:"AudioChannel,omitnil,omitempty" name:"AudioChannel"`

	// Merge audio track information.
	// This field only takes effect in adaptive bitrate transcoding.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	TrackChannelInfo *AudioTrackChannelInfo `json:"TrackChannelInfo,omitnil,omitempty" name:"TrackChannelInfo"`
}

type AudioTemplateInfoForUpdate struct {
	// Audio stream encoding format.
	// When audio transcoding is not needed, the value is:
	// <li>copy.</li>
	// When the outer parameter Container is mp3, the value is:
	// <li>mp3.</li>
	// When the outer parameter Container is ogg or flac, the value is:
	// <li>flac.</li>
	// When the outer parameter Container is m4a, valid values are:
	// <li>aac;</li>
	// <li>ac3.</li>
	// When the outer parameter Container is mp4 or flv, valid values are:
	// <li>aac: more suitable for mp4;</li>
	// <li>mp3: more suitable for flv;</li>
	// <li>mp2.</li>
	// When the outer parameter Container is hls, valid values are:
	// <li>aac;</li>
	// <li>mp3.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// Audio stream bitrate in Kbps. Value range: 0 and [26, 256]. If the value is 0, the bitrate of the audio stream will be the same as that of the original audio.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// The sampling rate of the audio stream. the sampling rate options supported by different encoding standards are different. for details, see the audio sample rate support scope document (https://intl.cloud.tencent.com/document/product/862/77166?from_cn_redirect=1#f3b039f1-d817-4a96-b4e4-90132d31cd53).
	// Unit: Hz.
	// Please ensure that the sampling rate of the source audio stream is within the scope of the above options. otherwise, transcoding failure may occur.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Audio channel mode. Valid values:
	// <li>1: mono-channel.</li>
	// <li>2: dual-channel.</li>
	// <li>6: 5.1 surround sound.
	// When the container format is audio (flac, ogg, mp3, and m4a), the audio channel cannot be set to 5.1 surround sound.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioChannel *int64 `json:"AudioChannel,omitnil,omitempty" name:"AudioChannel"`

	// The audio tracks to retain. All audio tracks are retained by default.
	StreamSelects []*int64 `json:"StreamSelects,omitnil,omitempty" name:"StreamSelects"`
}

type AudioTrackChannelInfo struct {
	// Whether to enable the feature of multi-audio track mixing. Valid values:
	// <li>0: To disable the multi-audio track mixing feature.
	// <li>1: To enable the multi-audio track mixing feature. 
	// <li>Default value: 0.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	ChannelsRemix *int64 `json:"ChannelsRemix,omitnil,omitempty" name:"ChannelsRemix"`

	// Set the selector type for the input audio track. Valid values:
	// <li>track: indicates the usage of audio track id to identify the track to be used.
	// <li>track_channel: indicates the usage of both the audio track id and sound channel id to identify the track and channel to be used.
	// <li>Default value: track.
	// If the original audio track has multiple sound channels, please use track_channel.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	SelectType *string `json:"SelectType,omitnil,omitempty" name:"SelectType"`

	// Audio track information.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	InputTrackInfo []*TrackInfo `json:"InputTrackInfo,omitnil,omitempty" name:"InputTrackInfo"`
}

type AwsS3FileUploadTrigger struct {
	// The AWS S3 bucket bound to the scheme.
	S3Bucket *string `json:"S3Bucket,omitnil,omitempty" name:"S3Bucket"`

	// The region of the AWS S3 bucket.
	S3Region *string `json:"S3Region,omitnil,omitempty" name:"S3Region"`

	// The bucket directory bound. It must be an absolute path that starts and ends with `/`, such as `/movie/201907/`. If you do not specify this, the root directory will be bound.	
	Dir *string `json:"Dir,omitnil,omitempty" name:"Dir"`

	// The file formats that will trigger the scheme, such as ["mp4", "flv", "mov"]. If you do not specify this, the upload of files in any format will trigger the scheme.	
	Formats []*string `json:"Formats,omitnil,omitempty" name:"Formats"`

	// The key ID of the AWS S3 bucket.
	// Note: This field may return null, indicating that no valid values can be obtained.
	S3SecretId *string `json:"S3SecretId,omitnil,omitempty" name:"S3SecretId"`

	// The key of the AWS S3 bucket.
	// Note: This field may return null, indicating that no valid values can be obtained.
	S3SecretKey *string `json:"S3SecretKey,omitnil,omitempty" name:"S3SecretKey"`

	// The SQS queue of the AWS S3 bucket.
	// Note: The queue must be in the same region as the bucket.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AwsSQS *AwsSQS `json:"AwsSQS,omitnil,omitempty" name:"AwsSQS"`
}

type AwsSQS struct {
	// The region of the SQS queue.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SQSRegion *string `json:"SQSRegion,omitnil,omitempty" name:"SQSRegion"`

	// The name of the SQS queue.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SQSQueueName *string `json:"SQSQueueName,omitnil,omitempty" name:"SQSQueueName"`

	// The key ID required to read from/write to the SQS queue.
	// Note: This field may return null, indicating that no valid values can be obtained.
	S3SecretId *string `json:"S3SecretId,omitnil,omitempty" name:"S3SecretId"`

	// The key required to read from/write to the SQS queue.
	// Note: This field may return null, indicating that no valid values can be obtained.
	S3SecretKey *string `json:"S3SecretKey,omitnil,omitempty" name:"S3SecretKey"`
}

// Predefined struct for user
type BatchProcessMediaRequestParams struct {
	// Path of the input file.
	InputInfo []*MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// Storage bucket for the output file. If it is left blank, the storage location in InputInfo will be inherited.
	// Note: When InputInfo.Type is URL, this parameter is required.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Storage directory for the output file. It should start and end with a slash (/), such as `/movie/201907/`.
	// If left blank, it indicates that the directory is the same as the one where the file is located in InputInfo.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Smart subtitle.
	SmartSubtitlesTask *SmartSubtitlesTaskInput `json:"SmartSubtitlesTask,omitnil,omitempty" name:"SmartSubtitlesTask"`

	// Event notification information of the task. If left blank, it indicates that no event notification will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Priority of the task flow. The higher the value, the higher the priority. The value range is from -10 to 10. If left blank, the default value is 0.
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// Source context, which is used to pass through the user request information. The callback for task flow status changes will return the value of this field. The maximum length is 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// Resource ID. Ensure the corresponding resource is in the enabled state. The default value is an account's primary resource ID.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Whether to skip metadata acquisition. Valid values:
	// 0: do not skip.
	// 1: skip.
	// Default value: 0		
	SkipMateData *int64 `json:"SkipMateData,omitnil,omitempty" name:"SkipMateData"`
}

type BatchProcessMediaRequest struct {
	*tchttp.BaseRequest
	
	// Path of the input file.
	InputInfo []*MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// Storage bucket for the output file. If it is left blank, the storage location in InputInfo will be inherited.
	// Note: When InputInfo.Type is URL, this parameter is required.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Storage directory for the output file. It should start and end with a slash (/), such as `/movie/201907/`.
	// If left blank, it indicates that the directory is the same as the one where the file is located in InputInfo.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Smart subtitle.
	SmartSubtitlesTask *SmartSubtitlesTaskInput `json:"SmartSubtitlesTask,omitnil,omitempty" name:"SmartSubtitlesTask"`

	// Event notification information of the task. If left blank, it indicates that no event notification will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Priority of the task flow. The higher the value, the higher the priority. The value range is from -10 to 10. If left blank, the default value is 0.
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// Source context, which is used to pass through the user request information. The callback for task flow status changes will return the value of this field. The maximum length is 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// Resource ID. Ensure the corresponding resource is in the enabled state. The default value is an account's primary resource ID.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Whether to skip metadata acquisition. Valid values:
	// 0: do not skip.
	// 1: skip.
	// Default value: 0		
	SkipMateData *int64 `json:"SkipMateData,omitnil,omitempty" name:"SkipMateData"`
}

func (r *BatchProcessMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchProcessMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "SmartSubtitlesTask")
	delete(f, "TaskNotifyConfig")
	delete(f, "TasksPriority")
	delete(f, "SessionContext")
	delete(f, "ResourceId")
	delete(f, "SkipMateData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchProcessMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchProcessMediaResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchProcessMediaResponse struct {
	*tchttp.BaseResponse
	Response *BatchProcessMediaResponseParams `json:"Response"`
}

func (r *BatchProcessMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchProcessMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchSmartSubtitlesResult struct {
	// Input information for smart subtitle tasks.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Input *SmartSubtitleTaskResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output information for smart subtitle tasks.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Outputs []*SmartSubtitleTaskBatchOutput `json:"Outputs,omitnil,omitempty" name:"Outputs"`
}

type BatchSubTaskResult struct {
	// Input information for a batch task.
	// Note: This field may return null, indicating that no valid value can be obtained.
	InputInfos []*MediaInputInfo `json:"InputInfos,omitnil,omitempty" name:"InputInfos"`

	// Metadata of the original video.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Metadatas []*MediaMetaData `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`

	// Execution result of the smart subtitle task.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SmartSubtitlesTaskResult *BatchSmartSubtitlesResult `json:"SmartSubtitlesTaskResult,omitnil,omitempty" name:"SmartSubtitlesTaskResult"`
}

type ClassificationConfigureInfo struct {
	// Switch of intelligent categorization task. Valid values:
	// <li>ON: enables intelligent categorization task;</li>
	// <li>OFF: disables intelligent categorization task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ClassificationConfigureInfoForUpdate struct {
	// Switch of intelligent categorization task. Valid values:
	// <li>ON: enables intelligent categorization task;</li>
	// <li>OFF: disables intelligent categorization task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ColorEnhanceConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The strength. Valid values:
	// <li>weak</li>
	// <li>normal</li>
	// <li>strong</li>
	// Default value: weak.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ComposeAudioItem struct {
	// The media information of the element.
	SourceMedia *ComposeSourceMedia `json:"SourceMedia,omitnil,omitempty" name:"SourceMedia"`

	// The time of the element in the timeline. If this is not specified, the element will follow the previous element.
	TrackTime *ComposeTrackTime `json:"TrackTime,omitnil,omitempty" name:"TrackTime"`

	// The operations performed, such as muting.
	AudioOperations []*ComposeAudioOperation `json:"AudioOperations,omitnil,omitempty" name:"AudioOperations"`
}

type ComposeAudioOperation struct {
	// The operation type. Valid values:
	// <li>`Volume`: Volume adjustment. </li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	//  The volume level. This parameter is valid if `Type` is `Volume`. Value range: 0–5. 
	// <li>If the parameter value is `0`, the video will be muted. </li>
	// <li>If the parameter value is smaller than 1, the volume will be reduced. </li>
	// <li>If the parameter value is `1`, the original volume will be kept. </li>
	// <li>If the parameter value is greater than 1, the volume will be increased. </li>
	Volume *float64 `json:"Volume,omitnil,omitempty" name:"Volume"`
}

type ComposeAudioStream struct {
	// The codec of the audio stream. Valid values:
	// <li>`AAC`: AAC (default), which is used for the MP4 container. </li>
	// <li>`MP3`: MP3 codec, which is used for the MP3 container. </li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// The sample rate (Hz) of the audio stream.
	// <li>16000 (default)</li>
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// The number of sound channels. Valid values:
	// u200c<li>`1`: Mono. </li>
	// <li>`2`: Dual (default). </li>
	AudioChannel *int64 `json:"AudioChannel,omitnil,omitempty" name:"AudioChannel"`

	// Reference bitrate, in kbps. Value range: 26-10000.
	// If set, the encoder will try to encode at this bitrate.
	// If not set, the service will automatically adopt a suitable bitrate based on audio parameters.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`
}

type ComposeCanvas struct {
	// The RGB value of the background color. The format is #RRGGBB, such as `#F0F0F0`. 
	// Default value: `#000000` (black).
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`

	// The canvas width (px), which is the width of the output video. Value range: 0–3840.  
	// The default value is `0`, which means that the canvas is as wide as the first video.
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// The canvas height (px), which is the height of the output video. Value range: 0–3840.  
	// The default value is `0`, which means that the canvas is as high as the first video.
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type ComposeEmptyItem struct {
	// The element duration.
	// <li>The value of this parameter ends with `s`, which means seconds. For example, `3.5s` indicates 3.5 seconds. </li>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type ComposeImageItem struct {
	// The media information of the element.
	SourceMedia *ComposeSourceMedia `json:"SourceMedia,omitnil,omitempty" name:"SourceMedia"`

	// The time of the element in the timeline. If this is not specified, the element will follow the previous element.
	TrackTime *ComposeTrackTime `json:"TrackTime,omitnil,omitempty" name:"TrackTime"`

	// The horizontal distance of the element's center from the canvas origin. Two formats are supported:
	// <li>If the value ends with %, it specifies the distance as a percentage of the canvas width. For example, `10%` means that the distance is 10% of the canvas width. </li>
	// u200c<li>If the value ends with px, it specifies the distance in pixels. For example, `100px` means that the distance is 100 pixels. </li>
	// Default value: `50%`.
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// The vertical distance of the element's center from the canvas origin. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the distance as a percentage of the canvas height. For example, `10%` means that the distance is 10% of the canvas height. </li>
	// u200c<li>If the value ends with px, it specifies the distance in pixels. For example, `100px` means that the distance is 100 pixels. </li>
	// Default value: `50%`.
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// The width of the video segment. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the width as a percentage of the canvas width. For example, `10%` means that the video width is 10% of the canvas width. </li>
	// u200c<li>If the value ends with px, it specifies the width in pixels. For example, `100px` means that the video width is 100 pixels. </li>
	// If one or both parameters are empty or set to `0`:
	// <li>If both `Width` and `Height` are empty, the original width and height of the element will be kept. </li>
	// <li>If `Width` is empty and `Height` is not, the width will be auto scaled. </li>
	// <li>If `Width` is not empty and `Height` is, the height will be auto scaled. </li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// The height of the element. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the height as a percentage of the canvas height. For example, `10%` means that the height is 10% of the canvas height. </li>
	// u200c<li>If the value ends with px, it specifies the height in pixels. For example, `100px` means that the height is 100 pixels. </li>
	// If one or both parameters are empty or set to `0`:
	// <li>If both `Width` and `Height` are empty, the original width and height of the video will be kept. </li>
	// <li>If `Width` is empty and `Height` is not, the width will be auto scaled. </li>
	// <li>If `Width` is not empty and `Height` is, the height will be auto scaled. </li>
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// The image operations, such as image rotation.
	ImageOperations []*ComposeImageOperation `json:"ImageOperations,omitnil,omitempty" name:"ImageOperations"`
}

type ComposeImageOperation struct {
	// The type. Valid values:
	// u200c<li>`Rotate`: Image rotation. </li>
	// <li>`Flip`: Image flipping. </li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// This is valid if `Type` is `Rotate`. The angle of rotation around the image center. Value range: 0–360.
	RotateAngle *float64 `json:"RotateAngle,omitnil,omitempty" name:"RotateAngle"`

	// This is valid if `Type` is `Flip`. How to flip the image. Valid values:xa0
	// u200c<li>`Horizental`: Flip horizontally. </li>
	// <li>`Vertical`: Flip vertically. </li>
	FlipType *string `json:"FlipType,omitnil,omitempty" name:"FlipType"`
}

type ComposeMediaConfig struct {
	// The information of the output video.
	TargetInfo *ComposeTargetInfo `json:"TargetInfo,omitnil,omitempty" name:"TargetInfo"`

	// The canvas information of the output video.
	Canvas *ComposeCanvas `json:"Canvas,omitnil,omitempty" name:"Canvas"`

	// The global styles. This parameter is used together with `Tracks` to specify styles, such as the subtitle style.
	Styles []*ComposeStyles `json:"Styles,omitnil,omitempty" name:"Styles"`

	// The information of media tracks (consisting of video, audio, image, and text elements) used to composite the video. About tracks and the timeline:
	// <ul><li>The timeline of a track is the same as the timeline of the output video. </li><li>The elements of different tracks are overlaid at the same time point in the timeline.</li><ul><li>Video, image, and text elements are overlaid according to their track number, with the first track on top. </li><li>Audio elements are mixed. </li></ul></ul>Note: The different elements of the same track cannot be overlaid (except subtitles).
	Tracks []*ComposeMediaTrack `json:"Tracks,omitnil,omitempty" name:"Tracks"`
}

type ComposeMediaItem struct {
	// The element type. Valid values:
	// <li>`Video` </li>
	// <li>`Audio` </li>
	// <li>`Image` </li>
	// <li>`Transition` </li>
	// <li>`Subtitle` </li>
	// <li>`Empty` </li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The information of the video element, which is valid if `Type` is `Video`.
	Video *ComposeVideoItem `json:"Video,omitnil,omitempty" name:"Video"`

	// The information of the audio element, which is valid if `Type` is `Audio`.
	Audio *ComposeAudioItem `json:"Audio,omitnil,omitempty" name:"Audio"`

	// The information of the image element, which is valid if `Type` is `Image`.
	Image *ComposeImageItem `json:"Image,omitnil,omitempty" name:"Image"`

	// The information of the transition element, which is valid if `Type` is `Transition`.
	Transition *ComposeTransitionItem `json:"Transition,omitnil,omitempty" name:"Transition"`

	// The information of the subtitle element, which is valid if `Type` is `Subtitle`.
	Subtitle *ComposeSubtitleItem `json:"Subtitle,omitnil,omitempty" name:"Subtitle"`

	// The information of the empty element, which is valid if `Type` is `Empty`. An empty element is used as a placeholder in the timeline.
	Empty *ComposeEmptyItem `json:"Empty,omitnil,omitempty" name:"Empty"`
}

type ComposeMediaTrack struct {
	// Track type. Valid values: <ul><li>Video: video track. It can consist of the following elements:</li><ul><li>Video elements</li><li>Image elements</li><li>Transition elements</li><li>Empty elements</li></ul><li>Audio: audio track. It can consist of the following elements:</li><ul><li>Audio elements</li><li>Transition elements</li><li>Empty elements</li></ul><li>Title: text track. It can consist of the following elements:</li><ul><li>Subtitle elements</li></ul></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The elements of a track.
	Items []*ComposeMediaItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type ComposeSourceMedia struct {
	// The material ID, which can be found in `FileInfos`.
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// The start time of the material. The following two formats are supported.
	// <li>If the value of this parameter ends with `s`, it specifies the time in seconds. For example, `3.5s` indicates the time when 3.5 seconds of the material elapses.</li>
	// u200c<li>If the value of this parameter ends with `%`, it specifies the time as a percentage of the material's duration. For example, `10%` indicates the time when 10% of the material's duration elapses. </li>
	// Default value: `0s`.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time of the material. This parameter and `StartTime` determine which time segment of the material is used. The following two formats are supported:
	// <li>If the value of this parameter ends with `s`, it specifies the time in seconds. For example, `3.5s` indicates the time when 3.5 seconds of the material elapses.</li>
	// u200c<li>If the value of this parameter ends with `%`, it specifies the time as a percentage of the material's duration. For example, `10%` indicates the time when 10% of the material's duration elapses. </li>
	// If the track duration is set, the default value is `StartTime` plus the track duration. If not, the default value is `StartTime` plus 1 second.
	// Note: `EndTime` must be at least 0.02 seconds later than `StartTime`.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ComposeStyles struct {
	// The style ID, which identifies an element style.
	// Note: The style ID can be up to 32 characters long and can contain letters, digits, and special characters -_
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// The type. Valid values:
	// <li>`Subtitle`: The subtitle style. </li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The subtitle style details. This parameter is valid if `Type` is `Subtitle`.
	Subtitle *ComposeSubtitleStyle `json:"Subtitle,omitnil,omitempty" name:"Subtitle"`
}

type ComposeSubtitleItem struct {
	// The subtitle style ID, which corresponds to the `Id` field of `ComposeStyles`.
	StyleId *string `json:"StyleId,omitnil,omitempty" name:"StyleId"`

	// The subtitle text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// The time of the element in the timeline. If this is not specified, the element will follow the previous element.	
	TrackTime *ComposeTrackTime `json:"TrackTime,omitnil,omitempty" name:"TrackTime"`
}

type ComposeSubtitleStyle struct {
	// The subtitle height. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the height as a percentage of the canvas height. For example, `10%` means that the height is 10% of the canvas height. </li>
	// u200c<li>If the value ends with px, it specifies the height in pixels. For example, `100px` means that the height is 100 pixels. </li>
	// The default value is the value of `FontSize`.
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// The bottom margin. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the margin as a percentage of the canvas height. For example, `10%` means that the margin is 10% of the canvas height. </li>
	// u200c<li>If the value ends with px, it specifies the margin in pixels. For example, `100px` means that the margin is 100 pixels. </li>
	// Default value: `0px`.
	MarginBottom *string `json:"MarginBottom,omitnil,omitempty" name:"MarginBottom"`

	// The font type. Valid values:
	// <li>`SimHei`(default): Chinese font Heiti. </li>
	// <Li>`SimSun`: Chinese font Songti. </li>
	FontType *string `json:"FontType,omitnil,omitempty" name:"FontType"`

	// The font size. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the size as a percentage of the canvas height. For example, `10%` means that the size is 10% of the canvas height. </li>
	// u200c<li>If the value ends with px, it specifies the size in pixels. For example, `100px` means that the size is 100 pixels. </li>
	// Default value: `2%`.
	FontSize *string `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// Whether to bold the text (some fonts may not support bold). Valid values:
	// <li>`0` (default): No. </li>
	// <li>`1`: Yes. </li>
	FontBold *int64 `json:"FontBold,omitnil,omitempty" name:"FontBold"`

	// Whether to italicize the text (some fonts may not support italics). Valid values:
	// <li>`0` (default): No. </li>
	// <li>`1`: Yes. </li>
	FontItalic *int64 `json:"FontItalic,omitnil,omitempty" name:"FontItalic"`

	// The font color (#RRGGBBAA).  
	// Default value: `0x000000FF` (black).  
	// Note: `AA` in the color notation defines the opacity of the color. It's optional.
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// The text alignment. Valid values:
	// <li>`Center`(default) </li>
	// <li>`Left` </li>
	// <li>`Right` </li>
	FontAlign *string `json:"FontAlign,omitnil,omitempty" name:"FontAlign"`

	// The margin for left/right align.
	// <li>If `FontAlign` is `Left`, this parameter specifies the left margin of the subtitles. </li>
	// <li>If `FontAlign` is `Right`, this parameter specifies the right margin of the subtitles. </li>
	// Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the margin as a percentage of the canvas width. For example, `10%` means that the margin is 10% of the canvas width. </li>
	// u200c<li>If the value ends with px, it specifies the margin in pixels. For example, `100px` means that the margin is 100 pixels. </li>
	FontAlignMargin *string `json:"FontAlignMargin,omitnil,omitempty" name:"FontAlignMargin"`

	// The subtitle border width. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the width as a percentage of the canvas height. For example, `10%` means that the width is 10% of the canvas height. </li>
	// u200c<li>If the value ends with px, it specifies the width in pixels. For example, `100px` means that the width is 100 pixels. </li>
	// The default value is `0`, which means the subtitles will have no borders.
	BorderWidth *string `json:"BorderWidth,omitnil,omitempty" name:"BorderWidth"`

	// The border color, whose format is the same as that for `FontColor`. This parameter is valid if `BorderWidth` is not `0`.
	BorderColor *string `json:"BorderColor,omitnil,omitempty" name:"BorderColor"`

	// The text background color, whose format is the same as that for `FontColor`.  
	// The default value is an empty string, which means the subtitles will not have a background color.
	BottomColor *string `json:"BottomColor,omitnil,omitempty" name:"BottomColor"`
}

type ComposeTargetInfo struct {
	// The container. Valid values:
	// <li>`mp4` (default), for video files. </li>
	// <li>`mp3`, for audio files. </li>
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Whether to remove video data. Valid values:
	// <li>`0` (default): No. </li>
	// <li>`1`: Yes. </li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>`0` (default): No. </li>
	// <li>`1`: Yes. </li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// The information of the output video stream.
	VideoStream *ComposeVideoStream `json:"VideoStream,omitnil,omitempty" name:"VideoStream"`

	// The information of the output audio stream.
	AudioStream *ComposeAudioStream `json:"AudioStream,omitnil,omitempty" name:"AudioStream"`
}

type ComposeTrackTime struct {
	// The time when the element starts on the track.
	// <li>The value of this parameter ends with `s`, which means seconds. For example, `3.5s` indicates the time when 3.5 seconds of the video elapses.</li>
	// Note: If this parameter is not specified, the start time will be the end time of the previous element. Therefore, you can also use the placeholder parameter `ComposeEmptyItem` to configure the start time.
	Start *string `json:"Start,omitnil,omitempty" name:"Start"`

	// The element duration.
	// <li>The value of this parameter ends with `s`, which means seconds. For example, `3.5s` means 3.5 seconds.</li>
	// The default value is the material duration, which is determined by `EndTime` and `StartTime` of `ComposeSourceMedia`. If `ComposeSourceMedia` is not specified, the duration will be 1 second.
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type ComposeTransitionItem struct {
	// The element duration. <li>The value of this parameter ends with `s`, which means seconds. For example, `3s` indicates 3 seconds. </li>
	// Default value: `1s`.
	// Note
	// <li>The number before `s` must be an integer. Non-integers will be rounded down to the nearest integer. </li>
	// <li>The transition element must be between two non-empty elements. </li>
	// <li>The duration of the transition element must be shorter than that of the preceding element and the following element. </li>
	// u200c<li>The start time of the following element on the track will be automatically changed to the end time of the preceding element minus the duration of the transition element. </li>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// The transition effects.
	// The default transition effect is fade.
	// Note: You can add at most one image transition and one audio transition.
	Transitions []*ComposeTransitionOperation `json:"Transitions,omitnil,omitempty" name:"Transitions"`
}

type ComposeTransitionOperation struct {
	// The transition type.
	// 
	// The image transition, which connects two video segments.
	// <li>`ImageFadeInFadeOut` </li>
	// u200c<li>`BowTieHorizontal` </li>
	// u200c<li>`BowTieVertical` </li>
	// u200c<li>`ButterflyWaveScrawler` </li>
	// <li>`Cannabisleaf` </li>
	// <li>`Circle` </li>
	// <li>`CircleCrop` </li>
	// u200c<li>`Circleopen` </li>
	// <li>`Crosswarp` </li>
	// <li>`Cube` </li>
	// <li>`DoomScreenTransition` </li>
	// <li>`Doorway` </li>
	// <li>`Dreamy` </li>
	// <li>`DreamyZoom` </li>
	// <li>`FilmBurn` </li>
	// <li>`GlitchMemories` </li>
	// <li>`Heart` </li>
	// <li>`InvertedPageCurl` </li>
	// <li>`Luma` </li>
	// <li>`Mosaic` </li>
	// <li>`Pinwheel` </li>
	// <li>`PolarFunction` </li>
	// <li>`PolkaDotsCurtain` </li>
	// <li>`Radial` </li>
	// <li>`RotateScaleFade` </li>
	// <li>`Squeeze` </li>
	// <li>`Swap` </li>
	// <li>`Swirl` </li>
	// <li>`UndulatingBurnOutSwirl` </li>
	// <li>`Windowblinds` </li>
	// <li>`WipeDown` </li>
	// <li>`WipeLeft` </li>
	// <li>`WipeRight` </li>
	// <li>`WipeUp` </li>
	// <li>`ZoomInCircles` </li> 
	// The audio transition, which connects two audio segments.
	// <li>`AudioFadeInFadeOut` </li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ComposeVideoItem struct {
	// The media information of the element.
	SourceMedia *ComposeSourceMedia `json:"SourceMedia,omitnil,omitempty" name:"SourceMedia"`

	// The time of the element in the timeline. If this is not specified, the element will follow the previous element.
	TrackTime *ComposeTrackTime `json:"TrackTime,omitnil,omitempty" name:"TrackTime"`

	// The horizontal distance of the element's center from the canvas origin. Two formats are supported:
	// <li>If the value ends with %, it specifies the distance as a percentage of the canvas width. For example, `10%` means that the distance is 10% of the canvas width. </li>
	// u200c<li>If the value ends with px, it specifies the distance in pixels. For example, `100px` means that the distance is 100 pixels. </li>
	// Default value: `50%`.
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// The vertical distance of the element's center from the canvas origin. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the distance as a percentage of the canvas height. For example, `10%` means that the distance is 10% of the canvas height. </li>
	// u200c<li>If the value ends with px, it specifies the distance in pixels. For example, `100px` means that the distance is 100 pixels. </li>
	// Default value: `50%`.
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// The width of the video segment. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the width as a percentage of the canvas width. For example, `10%` means that the video width is 10% of the canvas width. </li>
	// u200c<li>If the value ends with px, it specifies the width in pixels. For example, `100px` means that the video width is 100 pixels. </li>
	// If one or both parameters are empty or set to `0`:
	// <li>If both `Width` and `Height` are empty, the original width and height of the element will be kept. </li>
	// <li>If `Width` is empty and `Height` is not, the width will be auto scaled. </li>
	// <li>If `Width` is not empty and `Height` is, the height will be auto scaled. </li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// The height of the element. Two formats are supported:
	// u200c<li>If the value ends with %, it specifies the height as a percentage of the canvas height. For example, `10%` means that the height is 10% of the canvas height. </li>
	// u200c<li>If the value ends with px, it specifies the height in pixels. For example, `100px` means that the height is 100 pixels. </li>
	// If one or both parameters are empty or set to `0`:
	// <li>If both `Width` and `Height` are empty, the original width and height of the element will be kept. </li>
	// <li>If `Width` is empty and `Height` is not, the width will be auto scaled. </li>
	// <li>If `Width` is not empty and `Height` is, the height will be auto scaled. </li>
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// The image operations, such as image rotation.
	ImageOperations []*ComposeImageOperation `json:"ImageOperations,omitnil,omitempty" name:"ImageOperations"`

	// The audio operations, such as muting.
	AudioOperations []*ComposeAudioOperation `json:"AudioOperations,omitnil,omitempty" name:"AudioOperations"`
}

type ComposeVideoStream struct {
	// The codec. Valid values:
	// <li>`H.264` (default) </li>
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// The video frame rate (Hz). Value range: 0–60.  
	// The default value is `0`, which means that the frame rate will be the same as that of the first video.
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Reference bitrate, in kbps. Value range: 50-35000.
	// If set, the encoder will try to encode at this bitrate.
	// If not set, the service will automatically adopt a suitable bitrate based on the complexity of an image.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`
}

type ContainerDiagnoseResultItem struct {
	// Diagnosed exception category. Valid values:
	// DecodeParamException: decoding parameter exception.
	// TimeStampException: timestamp exception.
	// FrameException: frame rate exception.
	// StreamStatusException: stream status exception.
	// StreamInfo: stream information exception.
	// StreamAbnormalCharacteristics: stream characteristics exception.
	// DecodeException: decoding exception.
	// HLSRequirements: HLS format exception.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// Diagnosed specific exception type. Valid values:
	// 
	// VideoResolutionChanged: video resolution change.
	// AudioSampleRateChanged: audio sample rate change.
	// AudioChannelsChanged: audio channel quantity change.
	// ParameterSetsChanged: stream parameter set information change.
	// DarOrSarInvalid: video aspect ratio exception.
	// TimestampFallback: DTS timestamp rollback.
	// DtsJitter: DTS jitter too high.
	// PtsJitter: PTS jitter too high.
	// AACDurationDeviation: improper AAC frame timestamp interval.
	// AudioDroppingFrames: audio frame dropping.
	// VideoDroppingFrames: video frame dropping.
	// AVTimestampInterleave: improper audio-video interleaving.
	// PtsLessThanDts: PTS less than DTS for media streams.
	// ReceiveFpsJitter: significant jitter in the network receive frame rate.
	// ReceiveFpsTooSmall: network receive video frame rate too low.
	// FpsJitter: significant jitter in the stream frame rate calculated via PTS.
	// StreamOpenFailed: stream open failure.
	// StreamEnd: stream end.
	// StreamParseFailed: stream parsing failure.
	// VideoFirstFrameNotIdr: first frame not an IDR frame.
	// StreamNALUError: NALU start code error.
	// TsStreamNoAud: no AUD NALU in the H26x stream of MPEG-TS.
	// AudioStreamLack: no audio stream.
	// VideoStreamLack: no video stream.
	// LackAudioRecover: missing audio stream recovery.
	// LackVideoRecover: missing video stream recovery.
	// VideoBitrateOutofRange: video stream bitrate (kbps) out of range.
	// AudioBitrateOutofRange: audio stream bitrate (kbps) out of range.
	// VideoDecodeFailed: video decoding error.
	// AudioDecodeFailed: audio decoding error.
	// AudioOutOfPhase: opposite phase in dual-channel audio.
	// VideoDuplicatedFrame: duplicate frames in video streams.
	// AudioDuplicatedFrame: duplicate frames in audio streams.
	// VideoRotation: video rotation.
	// TsMultiPrograms: multiple programs in MPEG2-TS streams
	// Mp4InvalidCodecFourcc: codec FourCC in MP4 not meeting Apple HLS requirements.
	// HLSBadM3u8Format: invalid M3U8 file.
	// HLSInvalidMasterM3u8: invalid main M3U8 file.
	// HLSInvalidMediaM3u8: invalid media M3U8 file.
	// HLSMasterM3u8Recommended: parameters recommended by standards missing in main M3U8.
	// HLSMediaM3u8Recommended: parameters recommended by standards missing in media M3U8.
	// HLSMediaM3u8DiscontinuityExist: EXT-X-DISCONTINUITY in media M3U8.
	// HLSMediaSegmentsStreamNumChange: changed number of streams in segments.
	// HLSMediaSegmentsPTSJitterDeviation: PTS jumps between segments without EXT-X-DISCONTINUITY.
	// HLSMediaSegmentsDTSJitterDeviation: DTS jumps between segments without EXT-X-DISCONTINUITY.
	// TimecodeTrackExist: TMCD track in MP4.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Diagnosed exception level. Valid values:
	// Fatal: affecting subsequent playback and parsing.
	// Error: may affect playback.
	// Warning: potential risk, which may not necessarily affect playback.
	// Notice: important stream information.
	// Info: general stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SeverityLevel *string `json:"SeverityLevel,omitnil,omitempty" name:"SeverityLevel"`

	// Timestamp of warning, in the format of 2022-12-25T13:14:16Z.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DateTimeSet []*string `json:"DateTimeSet,omitnil,omitempty" name:"DateTimeSet"`

	// Timestamp.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimestampSet []*float64 `json:"TimestampSet,omitnil,omitempty" name:"TimestampSet"`
}

type ContentReviewTemplateItem struct {
	// Unique ID of a content audit template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of a content audit template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description of a content audit template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Porn information detection control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// The parameters for detecting sensitive information.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// The parameters for detecting sensitive information.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// Custom content audit control parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The template type. Valid values:
	// * Preset
	// * Custom
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CosFileUploadTrigger struct {
	// Name of the COS bucket bound to a workflow, such as `TopRankVideo-125xxx88`.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Region of the COS bucket bound to a workflow, such as `ap-chongiqng`.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Input path directory bound to a workflow, such as `/movie/201907/`. If this parameter is left empty, the `/` root directory will be used.
	Dir *string `json:"Dir,omitnil,omitempty" name:"Dir"`

	// Format list of files that can trigger a workflow, such as ["mp4", "flv", "mov"]. If this parameter is left empty, files in all formats can trigger the workflow.
	Formats []*string `json:"Formats,omitnil,omitempty" name:"Formats"`
}

type CosInputInfo struct {
	// The COS bucket of the object to process, such as `TopRankVideo-125xxx88`.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// The region of the COS bucket, such as `ap-chongqing`.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The path of the object to process, such as `/movie/201907/WildAnimal.mov`.
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`
}

type CosOutputStorage struct {
	// The bucket to which the output file of media processing is saved, such as `TopRankVideo-125xxx88`. If this parameter is left empty, the value of the upper layer will be inherited.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// The region of the output bucket, such as `ap-chongqing`. If this parameter is left empty, the value of the upper layer will be inherited.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type CoverConfigureInfo struct {
	// Switch of intelligent cover generating task. Valid values:
	// <li>ON: enables intelligent cover generating task;</li>
	// <li>OFF: disables intelligent cover generating task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type CoverConfigureInfoForUpdate struct {
	// Switch of intelligent cover generating task. Valid values:
	// <li>ON: enables intelligent cover generating task;</li>
	// <li>OFF: disables intelligent cover generating task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type CreateAIAnalysisTemplateRequestParams struct {
	// Video content analysis template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Video content analysis template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Control parameter of intelligent categorization task.
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// Control parameter of intelligent tagging task.
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// Control parameter of intelligent cover generating task.
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// Control parameter of intelligent frame-specific tagging task.
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`
}

type CreateAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Video content analysis template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Video content analysis template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Control parameter of intelligent categorization task.
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// Control parameter of intelligent tagging task.
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// Control parameter of intelligent cover generating task.
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// Control parameter of intelligent frame-specific tagging task.
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`
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

// Predefined struct for user
type CreateAIAnalysisTemplateResponseParams struct {
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Name of a video content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description of a video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Face recognition control parameter.
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// Full text recognition control parameter.
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// Text keyword recognition control parameter.
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// Full speech recognition control parameter.
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// Speech keyword recognition control parameter.
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`
}

type CreateAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Name of a video content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description of a video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Face recognition control parameter.
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// Full text recognition control parameter.
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// Text keyword recognition control parameter.
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// Full speech recognition control parameter.
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// Speech keyword recognition control parameter.
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`
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

// Predefined struct for user
type CreateAIRecognitionTemplateResponseParams struct {
	// Unique ID of a video content recognition template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Adaptive bitrate streaming format. Valid values:
	// <li>HLS,</li>
	// <li>MPEG-DASH.</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Parameter information of output substreams for transcoding to adaptive bitrate streaming. Up to 10 substreams can be output.
	// Note: the frame rate of each substream must be consistent; otherwise, the frame rate of the first substream is used as the output frame rate.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether to prohibit transcoding from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: 0.
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: 0.
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Indicates whether it is audio-only. 0 means video template, 1 means audio-only template.
	// When the value is 1.
	// 1. StreamInfos.N.RemoveVideo=1
	// 2. StreamInfos.N.RemoveAudio=0
	// 3. StreamInfos.N.Video.Codec=copy
	// When the value is 0.
	// 1. StreamInfos.N.Video.Codec cannot be copy.
	// 2. StreamInfos.N.Video.Fps cannot be null.
	// 
	// Note:
	// 
	// This value only distinguishes template types. The task uses the values of RemoveAudio and RemoveVideo.
	PureAudio *uint64 `json:"PureAudio,omitnil,omitempty" name:"PureAudio"`

	// HLS segment type. Valid values: <li>ts-segment: HLS+TS segment.</li> <li>ts-byterange: HLS+TS byte range.</li> <li>mp4-segment: HLS+MP4 segment.</li> <li>mp4-byterange: HLS+MP4 byte range.</li> <li>ts-packed-audio: TS+Packed audio.</li> <li>mp4-packed-audio: MP4+Packed audio.</li> Default value: ts-segment.
	// Note: The HLS segment format for adaptive bitrate streaming is based on this field.
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
}

type CreateAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Adaptive bitrate streaming format. Valid values:
	// <li>HLS,</li>
	// <li>MPEG-DASH.</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Parameter information of output substreams for transcoding to adaptive bitrate streaming. Up to 10 substreams can be output.
	// Note: the frame rate of each substream must be consistent; otherwise, the frame rate of the first substream is used as the output frame rate.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether to prohibit transcoding from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: 0.
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	// Default value: 0.
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Indicates whether it is audio-only. 0 means video template, 1 means audio-only template.
	// When the value is 1.
	// 1. StreamInfos.N.RemoveVideo=1
	// 2. StreamInfos.N.RemoveAudio=0
	// 3. StreamInfos.N.Video.Codec=copy
	// When the value is 0.
	// 1. StreamInfos.N.Video.Codec cannot be copy.
	// 2. StreamInfos.N.Video.Fps cannot be null.
	// 
	// Note:
	// 
	// This value only distinguishes template types. The task uses the values of RemoveAudio and RemoveVideo.
	PureAudio *uint64 `json:"PureAudio,omitnil,omitempty" name:"PureAudio"`

	// HLS segment type. Valid values: <li>ts-segment: HLS+TS segment.</li> <li>ts-byterange: HLS+TS byte range.</li> <li>mp4-segment: HLS+MP4 segment.</li> <li>mp4-byterange: HLS+MP4 byte range.</li> <li>ts-packed-audio: TS+Packed audio.</li> <li>mp4-packed-audio: MP4+Packed audio.</li> Default value: ts-segment.
	// Note: The HLS segment format for adaptive bitrate streaming is based on this field.
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
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
	delete(f, "PureAudio")
	delete(f, "SegmentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdaptiveDynamicStreamingTemplateResponseParams struct {
	// Unique ID of an adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Maximum value of the width (or long side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Animated image format. Valid values: gif; webp. Default value: gif.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Image quality. Value range: [1, 100]. Default value: 75.
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// Name of an animated image generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Video frame rate in Hz. Value range: [1, 30].
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Maximum value of the width (or long side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Animated image format. Valid values: gif; webp. Default value: gif.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Image quality. Value range: [1, 100]. Default value: 75.
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// Name of an animated image generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
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

// Predefined struct for user
type CreateAnimatedGraphicsTemplateResponseParams struct {
	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateAsrHotwordsRequestParams struct {
	// 0: temporary hotword; 1 file-based hotword.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Hotword lexicon name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Hotword lexicon text. This field is required if Type is set to 0.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Base64-encoded content of the hotword file. This field is required if Type is set to 1.
	// 
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// Name of the uploaded file.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type CreateAsrHotwordsRequest struct {
	*tchttp.BaseRequest
	
	// 0: temporary hotword; 1 file-based hotword.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Hotword lexicon name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Hotword lexicon text. This field is required if Type is set to 0.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Base64-encoded content of the hotword file. This field is required if Type is set to 1.
	// 
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// Name of the uploaded file.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

func (r *CreateAsrHotwordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsrHotwordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "FileContent")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAsrHotwordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAsrHotwordsResponseParams struct {
	// Hotword lexicon ID.
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAsrHotwordsResponse struct {
	*tchttp.BaseResponse
	Response *CreateAsrHotwordsResponseParams `json:"Response"`
}

func (r *CreateAsrHotwordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsrHotwordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContentReviewTemplateRequestParams struct {
	// The name of the content moderation template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Control parameter for a pornography detection task.
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// Control parameter for a violence detection task.
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// Control parameter for a sensitive content detection task.
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this parameter is not supported yet.
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// Custom content moderation parameters.
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`
}

type CreateContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// The name of the content moderation template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Control parameter for a pornography detection task.
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// Control parameter for a violence detection task.
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// Control parameter for a sensitive content detection task.
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this parameter is not supported yet.
	ProhibitedConfigure *ProhibitedConfigureInfo `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// Custom content moderation parameters.
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`
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

// Predefined struct for user
type CreateContentReviewTemplateResponseParams struct {
	// The unique ID of the content moderation template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateImageSpriteTemplateRequestParams struct {
	// Sampling type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// Subimage row count of an image sprite.
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// Subimage column count of an image sprite.
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// Name of an image sprite generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Subimage width of an image sprite in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Subimage height of an image sprite in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// The image format. Valid values: jpg (default), png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type CreateImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Sampling type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// Subimage row count of an image sprite.
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// Subimage column count of an image sprite.
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// Name of an image sprite generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Subimage width of an image sprite in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Subimage height of an image sprite in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// The image format. Valid values: jpg (default), png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
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
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageSpriteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageSpriteTemplateResponseParams struct {
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateLiveRecordTemplateRequestParams struct {
	// HLS configuration parameter. Either this parameter or MP4Configure should be specified.
	HLSConfigure *HLSConfigureInfo `json:"HLSConfigure,omitnil,omitempty" name:"HLSConfigure"`

	// MP4 configuration parameter. Either this parameter or HLSConfigure should be specified.
	MP4Configure *MP4ConfigureInfo `json:"MP4Configure,omitnil,omitempty" name:"MP4Configure"`

	// Recording template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description, with a length limit of 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// HLS configuration parameter. Either this parameter or MP4Configure should be specified.
	HLSConfigure *HLSConfigureInfo `json:"HLSConfigure,omitnil,omitempty" name:"HLSConfigure"`

	// MP4 configuration parameter. Either this parameter or HLSConfigure should be specified.
	MP4Configure *MP4ConfigureInfo `json:"MP4Configure,omitnil,omitempty" name:"MP4Configure"`

	// Recording template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description, with a length limit of 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HLSConfigure")
	delete(f, "MP4Configure")
	delete(f, "Name")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveRecordTemplateResponseParams struct {
	// Unique identifier of the recording template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveRecordTemplateResponseParams `json:"Response"`
}

func (r *CreateLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonSampleRequestParams struct {
	// Name of an image. Length limit: 20 characters
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image usage. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: equivalent to 1+2
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Image description. Length limit: 1,024 characters
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// [Base64](https://tools.ietf.org/html/rfc4648) string converted from an image. Only JPEG and PNG images are supported. Array length limit: 5 images
	// Note: the image must be a relatively clear facial feature photo of one person with a size of at least 200 x 200 pixels.
	FaceContents []*string `json:"FaceContents,omitnil,omitempty" name:"FaceContents"`

	// Image tag
	// <li>Array length limit: 20 tags</li>
	// <li>Tag length limit: 128 characters</li>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreatePersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// Name of an image. Length limit: 20 characters
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image usage. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: equivalent to 1+2
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Image description. Length limit: 1,024 characters
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// [Base64](https://tools.ietf.org/html/rfc4648) string converted from an image. Only JPEG and PNG images are supported. Array length limit: 5 images
	// Note: the image must be a relatively clear facial feature photo of one person with a size of at least 200 x 200 pixels.
	FaceContents []*string `json:"FaceContents,omitnil,omitempty" name:"FaceContents"`

	// Image tag
	// <li>Array length limit: 20 tags</li>
	// <li>Tag length limit: 128 characters</li>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
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

// Predefined struct for user
type CreatePersonSampleResponseParams struct {
	// Image information
	Person *AiSamplePerson `json:"Person,omitnil,omitempty" name:"Person"`

	// Information of images that failed the verification by facial feature positioning
	FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitnil,omitempty" name:"FailFaceInfoSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateQualityControlTemplateRequestParams struct {
	// Media quality inspection template name, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Media quality inspection control parameters.
	QualityControlItemSet []*QualityControlItemConfig `json:"QualityControlItemSet,omitnil,omitempty" name:"QualityControlItemSet"`

	// Media quality inspection template description, with a length limit of 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateQualityControlTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Media quality inspection template name, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Media quality inspection control parameters.
	QualityControlItemSet []*QualityControlItemConfig `json:"QualityControlItemSet,omitnil,omitempty" name:"QualityControlItemSet"`

	// Media quality inspection template description, with a length limit of 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateQualityControlTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQualityControlTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "QualityControlItemSet")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQualityControlTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQualityControlTemplateResponseParams struct {
	// Unique identifier of a media quality inspection template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQualityControlTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateQualityControlTemplateResponseParams `json:"Response"`
}

func (r *CreateQualityControlTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQualityControlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSampleSnapshotTemplateRequestParams struct {
	// Sampled screencapturing type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// The image format. Valid values: jpg (default), png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type CreateSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Sampled screencapturing type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// The image format. Valid values: jpg (default), png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
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

// Predefined struct for user
type CreateSampleSnapshotTemplateResponseParams struct {
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateScheduleRequestParams struct {
	// The scheme name (max 128 characters). This name should be unique across your account.
	ScheduleName *string `json:"ScheduleName,omitnil,omitempty" name:"ScheduleName"`

	// The trigger of the scheme. If a file is uploaded to the specified bucket, the scheme will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The subtasks of the scheme.
	Activities []*Activity `json:"Activities,omitnil,omitempty" name:"Activities"`

	// The bucket to save the output file. If you do not specify this parameter, the bucket in `Trigger` will be used.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The directory to save the media processing output file, which must start and end with `/`, such as `/movie/201907/`.
	// If you do not specify this, the file will be saved to the trigger directory.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// The notification configuration. If you do not specify this parameter, notifications will not be sent.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Resource ID. Ensure the corresponding resource is in the enabled state. The default value is an account's primary resource ID.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type CreateScheduleRequest struct {
	*tchttp.BaseRequest
	
	// The scheme name (max 128 characters). This name should be unique across your account.
	ScheduleName *string `json:"ScheduleName,omitnil,omitempty" name:"ScheduleName"`

	// The trigger of the scheme. If a file is uploaded to the specified bucket, the scheme will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The subtasks of the scheme.
	Activities []*Activity `json:"Activities,omitnil,omitempty" name:"Activities"`

	// The bucket to save the output file. If you do not specify this parameter, the bucket in `Trigger` will be used.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The directory to save the media processing output file, which must start and end with `/`, such as `/movie/201907/`.
	// If you do not specify this, the file will be saved to the trigger directory.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// The notification configuration. If you do not specify this parameter, notifications will not be sent.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Resource ID. Ensure the corresponding resource is in the enabled state. The default value is an account's primary resource ID.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *CreateScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleName")
	delete(f, "Trigger")
	delete(f, "Activities")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "TaskNotifyConfig")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScheduleResponseParams struct {
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateScheduleResponse struct {
	*tchttp.BaseResponse
	Response *CreateScheduleResponseParams `json:"Response"`
}

func (r *CreateScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSmartSubtitleTemplateRequestParams struct {
	// Smart subtitle template name.
	// Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Source language of the video with smart subtitles.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// zh-PY: Chinese-English-Cantonese
	// zh-medical: Medical Chinese
	// yue: Cantonese
	// vi: Vietnamese
	// ms: Malay
	// id: Indonesian
	// fil: Filipino
	// th: Thai
	// pt: Portuguese
	// tr: Turkish
	// ar: Arabic
	// es: Spanish
	// hi: Hindi
	// fr: French
	// de: German
	// zh-dialect: Chinese dialect
	VideoSrcLanguage *string `json:"VideoSrcLanguage,omitnil,omitempty" name:"VideoSrcLanguage"`

	// Smart subtitle language type.
	// 0: source language1: target language
	// 2: source language + target language
	// The value can only be 0 when TranslateSwitch is set to OFF.The value can only be 1 or 2 when TranslateSwitch is set to ON.
	SubtitleType *int64 `json:"SubtitleType,omitnil,omitempty" name:"SubtitleType"`

	// Smart subtitle template description.
	// Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Smart subtitle file format.
	// vtt: WebVTT format
	// If this field is left blank, no subtitle file will be generated.
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`

	// ASR hotword lexicon parameter.
	AsrHotWordsConfigure *AsrHotWordsConfigure `json:"AsrHotWordsConfigure,omitnil,omitempty" name:"AsrHotWordsConfigure"`

	// Subtitle translation switch.
	// ON: enable translation
	// OFF: disable translation
	TranslateSwitch *string `json:"TranslateSwitch,omitnil,omitempty" name:"TranslateSwitch"`

	// Target language for subtitle translation.
	// This field takes effect when TranslateSwitch is set to ON.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// fr: French
	// es: Spanish
	// it: Italian
	// de: German
	// tr: Turkish
	// ru: Russian
	// pt: Portuguese
	// vi: Vietnamese
	// id: Indonesian
	// ms: Malay
	// th: Thai
	// ar: Arabic
	// hi: Hindi
	TranslateDstLanguage *string `json:"TranslateDstLanguage,omitnil,omitempty" name:"TranslateDstLanguage"`
}

type CreateSmartSubtitleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Smart subtitle template name.
	// Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Source language of the video with smart subtitles.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// zh-PY: Chinese-English-Cantonese
	// zh-medical: Medical Chinese
	// yue: Cantonese
	// vi: Vietnamese
	// ms: Malay
	// id: Indonesian
	// fil: Filipino
	// th: Thai
	// pt: Portuguese
	// tr: Turkish
	// ar: Arabic
	// es: Spanish
	// hi: Hindi
	// fr: French
	// de: German
	// zh-dialect: Chinese dialect
	VideoSrcLanguage *string `json:"VideoSrcLanguage,omitnil,omitempty" name:"VideoSrcLanguage"`

	// Smart subtitle language type.
	// 0: source language1: target language
	// 2: source language + target language
	// The value can only be 0 when TranslateSwitch is set to OFF.The value can only be 1 or 2 when TranslateSwitch is set to ON.
	SubtitleType *int64 `json:"SubtitleType,omitnil,omitempty" name:"SubtitleType"`

	// Smart subtitle template description.
	// Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Smart subtitle file format.
	// vtt: WebVTT format
	// If this field is left blank, no subtitle file will be generated.
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`

	// ASR hotword lexicon parameter.
	AsrHotWordsConfigure *AsrHotWordsConfigure `json:"AsrHotWordsConfigure,omitnil,omitempty" name:"AsrHotWordsConfigure"`

	// Subtitle translation switch.
	// ON: enable translation
	// OFF: disable translation
	TranslateSwitch *string `json:"TranslateSwitch,omitnil,omitempty" name:"TranslateSwitch"`

	// Target language for subtitle translation.
	// This field takes effect when TranslateSwitch is set to ON.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// fr: French
	// es: Spanish
	// it: Italian
	// de: German
	// tr: Turkish
	// ru: Russian
	// pt: Portuguese
	// vi: Vietnamese
	// id: Indonesian
	// ms: Malay
	// th: Thai
	// ar: Arabic
	// hi: Hindi
	TranslateDstLanguage *string `json:"TranslateDstLanguage,omitnil,omitempty" name:"TranslateDstLanguage"`
}

func (r *CreateSmartSubtitleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSmartSubtitleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "VideoSrcLanguage")
	delete(f, "SubtitleType")
	delete(f, "Comment")
	delete(f, "SubtitleFormat")
	delete(f, "AsrHotWordsConfigure")
	delete(f, "TranslateSwitch")
	delete(f, "TranslateDstLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSmartSubtitleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSmartSubtitleTemplateResponseParams struct {
	// Unique identifier of the smart subtitle template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSmartSubtitleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateSmartSubtitleTemplateResponseParams `json:"Response"`
}

func (r *CreateSmartSubtitleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSmartSubtitleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotByTimeOffsetTemplateRequestParams struct {
	// Name of a time point screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// The image format. Valid values: jpg (default), png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type CreateSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Name of a time point screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// The image format. Valid values: jpg (default), png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
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

// Predefined struct for user
type CreateSnapshotByTimeOffsetTemplateResponseParams struct {
	// Unique ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateTranscodeTemplateRequestParams struct {
	// Container format. Valid values: mp4, flv, hls, ts, webm, mkv, mxf, mov, mp3, flac, ogg, and m4a. Among them, mp3, flac, ogg, and m4a are for audio-only files.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Name of a transcoding template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Whether to remove video data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	// Default value: 0.
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	// Default value: 0.
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter. This field is required when `RemoveVideo` is 0.
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter. This field is required when `RemoveAudio` is 0.
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter. To enable it, please contact your Tencent Cloud sales rep.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// Audio/Video enhancement configuration.
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil,omitempty" name:"EnhanceConfig"`

	// Additional parameter, which is a serialized JSON string.
	StdExtInfo *string `json:"StdExtInfo,omitnil,omitempty" name:"StdExtInfo"`
}

type CreateTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Container format. Valid values: mp4, flv, hls, ts, webm, mkv, mxf, mov, mp3, flac, ogg, and m4a. Among them, mp3, flac, ogg, and m4a are for audio-only files.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Name of a transcoding template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Whether to remove video data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	// Default value: 0.
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	// Default value: 0.
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter. This field is required when `RemoveVideo` is 0.
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter. This field is required when `RemoveAudio` is 0.
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter. To enable it, please contact your Tencent Cloud sales rep.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// Audio/Video enhancement configuration.
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil,omitempty" name:"EnhanceConfig"`

	// Additional parameter, which is a serialized JSON string.
	StdExtInfo *string `json:"StdExtInfo,omitnil,omitempty" name:"StdExtInfo"`
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
	delete(f, "EnhanceConfig")
	delete(f, "StdExtInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeTemplateResponseParams struct {
	// Unique ID of a transcoding template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateWatermarkTemplateRequestParams struct {
	// Watermarking type. Valid values:
	// <li>image: Image watermark;</li>
	// <li>text: Text watermark;</li>
	// <li>svg: SVG watermark.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Watermarking template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Origin position. Valid values:
	// <li>TopLeft: The origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>TopRight: The origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>BottomLeft: The origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>BottomRight: The origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	// Default value: TopLeft.
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the watermark will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width;</li>
	// <li>If the string ends in px, the `XPos` of the watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.</li>
	// Default value: 0 px.
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the watermark will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height;</li>
	// <li>If the string ends in px, the `YPos` of the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.</li>
	// Default value: 0 px.
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// Image watermarking template. This field is required and valid only when `Type` is `image`.
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is required and valid only when `Type` is `text`.
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is required and valid only when `Type` is `svg`.
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`
}

type CreateWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Watermarking type. Valid values:
	// <li>image: Image watermark;</li>
	// <li>text: Text watermark;</li>
	// <li>svg: SVG watermark.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Watermarking template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Origin position. Valid values:
	// <li>TopLeft: The origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>TopRight: The origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>BottomLeft: The origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>BottomRight: The origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	// Default value: TopLeft.
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the watermark will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width;</li>
	// <li>If the string ends in px, the `XPos` of the watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.</li>
	// Default value: 0 px.
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the watermark will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height;</li>
	// <li>If the string ends in px, the `YPos` of the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.</li>
	// Default value: 0 px.
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// Image watermarking template. This field is required and valid only when `Type` is `image`.
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is required and valid only when `Type` is `text`.
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is required and valid only when `Type` is `svg`.
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`
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

// Predefined struct for user
type CreateWatermarkTemplateResponseParams struct {
	// Unique ID of a watermarking template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Watermark image address. This field is valid only when `Type` is `image`.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 7. All: ASR- and OCR-based content recognition and inappropriate information detection; equivalent to 1+2+3+4
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Keyword. Array length limit: 100.
	Words []*AiSampleWordInfo `json:"Words,omitnil,omitempty" name:"Words"`
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
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Keyword. Array length limit: 100.
	Words []*AiSampleWordInfo `json:"Words,omitnil,omitempty" name:"Words"`
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

// Predefined struct for user
type CreateWordSamplesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateWorkflowRequestParams struct {
	// Workflow name of up to 128 characters, which must be unique for the same user.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Triggering rule bound to a workflow. If an uploaded video hits the rule for the object, the workflow will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The location to save the output file of media processing. If this parameter is left empty, the storage location in `Trigger` will be inherited.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The directory to save the media processing output file, which must start and end with `/`, such as `/movie/201907/`.
	// If you do not specify this, the file will be saved to the trigger directory.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// The media processing parameters to use.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// Event notification configuration for a task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Workflow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TaskPriority *int64 `json:"TaskPriority,omitnil,omitempty" name:"TaskPriority"`
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Workflow name of up to 128 characters, which must be unique for the same user.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Triggering rule bound to a workflow. If an uploaded video hits the rule for the object, the workflow will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The location to save the output file of media processing. If this parameter is left empty, the storage location in `Trigger` will be inherited.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The directory to save the media processing output file, which must start and end with `/`, such as `/movie/201907/`.
	// If you do not specify this, the file will be saved to the trigger directory.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// The media processing parameters to use.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// Event notification configuration for a task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Workflow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TaskPriority *int64 `json:"TaskPriority,omitnil,omitempty" name:"TaskPriority"`
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

// Predefined struct for user
type CreateWorkflowResponseParams struct {
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteAIAnalysisTemplateRequestParams struct {
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteAIAnalysisTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Unique ID of a video content recognition template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a video content recognition template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteAIRecognitionTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Unique ID of an adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteAdaptiveDynamicStreamingTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteAnimatedGraphicsTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteAsrHotwordsRequestParams struct {
	// ID of the hotword lexicon to be deleted.
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`
}

type DeleteAsrHotwordsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the hotword lexicon to be deleted.
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`
}

func (r *DeleteAsrHotwordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsrHotwordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HotwordsId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAsrHotwordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAsrHotwordsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAsrHotwordsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAsrHotwordsResponseParams `json:"Response"`
}

func (r *DeleteAsrHotwordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsrHotwordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContentReviewTemplateRequestParams struct {
	// The unique ID of the content moderation template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the content moderation template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteContentReviewTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteImageSpriteTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteLiveRecordTemplateRequestParams struct {
	// Unique identifier of the recording template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique identifier of the recording template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

func (r *DeleteLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLiveRecordTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveRecordTemplateResponseParams `json:"Response"`
}

func (r *DeleteLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonSampleRequestParams struct {
	// Image ID
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

type DeletePersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// Image ID
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
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

// Predefined struct for user
type DeletePersonSampleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteQualityControlTemplateRequestParams struct {
	// Unique identifier of a media quality inspection template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteQualityControlTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique identifier of a media quality inspection template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

func (r *DeleteQualityControlTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQualityControlTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQualityControlTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQualityControlTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQualityControlTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQualityControlTemplateResponseParams `json:"Response"`
}

func (r *DeleteQualityControlTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQualityControlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSampleSnapshotTemplateRequestParams struct {
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteSampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteSampleSnapshotTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteScheduleRequestParams struct {
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`
}

type DeleteScheduleRequest struct {
	*tchttp.BaseRequest
	
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`
}

func (r *DeleteScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScheduleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScheduleResponseParams `json:"Response"`
}

func (r *DeleteScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmartSubtitleTemplateRequestParams struct {
	// Unique identifier of the smart subtitle template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteSmartSubtitleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique identifier of the smart subtitle template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

func (r *DeleteSmartSubtitleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmartSubtitleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSmartSubtitleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmartSubtitleTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSmartSubtitleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSmartSubtitleTemplateResponseParams `json:"Response"`
}

func (r *DeleteSmartSubtitleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmartSubtitleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotByTimeOffsetTemplateRequestParams struct {
	// Unique ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteSnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteSnapshotByTimeOffsetTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteTranscodeTemplateRequestParams struct {
	// Unique ID of a transcoding template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a transcoding template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteTranscodeTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteWatermarkTemplateRequestParams struct {
	// Unique ID of a watermarking template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type DeleteWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a watermarking template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`
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

// Predefined struct for user
type DeleteWatermarkTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`
}

type DeleteWordSamplesRequest struct {
	*tchttp.BaseRequest
	
	// Keyword. Array length limit: 100 words.
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`
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

// Predefined struct for user
type DeleteWordSamplesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteWorkflowRequestParams struct {
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DeleteWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
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

// Predefined struct for user
type DeleteWorkflowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAIAnalysisTemplatesRequestParams struct {
	// Unique ID filter of video content analysis templates. Array length limit: 10.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The filter for querying templates. If this parameter is left empty, both preset and custom templates are returned. Valid values:
	// * Preset
	// * Custom
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for video analysis template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeAIAnalysisTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID filter of video content analysis templates. Array length limit: 10.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The filter for querying templates. If this parameter is left empty, both preset and custom templates are returned. Valid values:
	// * Preset
	// * Custom
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for video analysis template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Type")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIAnalysisTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAnalysisTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of video content analysis template details.
	AIAnalysisTemplateSet []*AIAnalysisTemplateItem `json:"AIAnalysisTemplateSet,omitnil,omitempty" name:"AIAnalysisTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Unique ID filter of video content recognition templates. Array length limit: 10.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The filter for querying templates. If this parameter is left empty, both preset and custom templates are returned. Valid values:
	// * Preset
	// * Custom
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for video recognition template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeAIRecognitionTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID filter of video content recognition templates. Array length limit: 10.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The filter for querying templates. If this parameter is left empty, both preset and custom templates are returned. Valid values:
	// * Preset
	// * Custom
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for video recognition template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Type")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIRecognitionTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIRecognitionTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of video content recognition template details.
	AIRecognitionTemplateSet []*AIRecognitionTemplateItem `json:"AIRecognitionTemplateSet,omitnil,omitempty" name:"AIRecognitionTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Unique ID filter of adaptive bitrate streaming templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether it is an audio-only template. 0: video template. 1: audio-only template.
	// 
	// Default value: 0
	PureAudio *uint64 `json:"PureAudio,omitnil,omitempty" name:"PureAudio"`

	// Filter condition for adaptive transcoding template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeAdaptiveDynamicStreamingTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID filter of adaptive bitrate streaming templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: preset template;</li>
	// <li>Custom: custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether it is an audio-only template. 0: video template. 1: audio-only template.
	// 
	// Default value: 0
	PureAudio *uint64 `json:"PureAudio,omitnil,omitempty" name:"PureAudio"`

	// Filter condition for adaptive transcoding template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "PureAudio")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdaptiveDynamicStreamingTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdaptiveDynamicStreamingTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of adaptive bitrate streaming template details.
	AdaptiveDynamicStreamingTemplateSet []*AdaptiveDynamicStreamingTemplate `json:"AdaptiveDynamicStreamingTemplateSet,omitnil,omitempty" name:"AdaptiveDynamicStreamingTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAnimatedGraphicsTemplatesRequestParams struct {
	// Unique ID filter of animated image generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for animated image generating template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeAnimatedGraphicsTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID filter of animated image generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for animated image generating template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAnimatedGraphicsTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnimatedGraphicsTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of animated image generating template details.
	AnimatedGraphicsTemplateSet []*AnimatedGraphicsTemplate `json:"AnimatedGraphicsTemplateSet,omitnil,omitempty" name:"AnimatedGraphicsTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAsrHotwordsListRequestParams struct {
	// Parameter for querying by hotword lexicon ID.
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`

	// Parameter for querying by hotword lexicon name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. All hotword lexicons are returned by default.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Hotword lexicon sorting order.
	// 
	// 0: ascending (default)
	// 1: descending
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Sorts hotword lexicons by a specific field. By default, hotword lexicons are sorted by creation time. If an invalid field is used for sorting, the default sorting field applies.
	// 
	//  - CreateTime: sort by creation time
	//  - UpdateTime: sort by update time
	//  - Name: sort by hotword lexicon name
	//  - WordCount: sort by the number of hotwords
	//  - HotwordsId: sort by hotword lexicon ID
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 0: temporary hotword; 1 file-based hotword.
	Types []*uint64 `json:"Types,omitnil,omitempty" name:"Types"`
}

type DescribeAsrHotwordsListRequest struct {
	*tchttp.BaseRequest
	
	// Parameter for querying by hotword lexicon ID.
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`

	// Parameter for querying by hotword lexicon name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. All hotword lexicons are returned by default.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Hotword lexicon sorting order.
	// 
	// 0: ascending (default)
	// 1: descending
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Sorts hotword lexicons by a specific field. By default, hotword lexicons are sorted by creation time. If an invalid field is used for sorting, the default sorting field applies.
	// 
	//  - CreateTime: sort by creation time
	//  - UpdateTime: sort by update time
	//  - Name: sort by hotword lexicon name
	//  - WordCount: sort by the number of hotwords
	//  - HotwordsId: sort by hotword lexicon ID
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 0: temporary hotword; 1 file-based hotword.
	Types []*uint64 `json:"Types,omitnil,omitempty" name:"Types"`
}

func (r *DescribeAsrHotwordsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsrHotwordsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HotwordsId")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderType")
	delete(f, "OrderBy")
	delete(f, "Types")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsrHotwordsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsrHotwordsListResponseParams struct {
	// Total number of hotword lexicons.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. All hotword lexicons are returned by default.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Hotword lexicon list.
	AsrHotwordsSet []*AsrHotwordsSet `json:"AsrHotwordsSet,omitnil,omitempty" name:"AsrHotwordsSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAsrHotwordsListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsrHotwordsListResponseParams `json:"Response"`
}

func (r *DescribeAsrHotwordsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsrHotwordsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsrHotwordsRequestParams struct {
	// ID of the hotword lexicon to be queried.
	// **Note: Either HotwordsId or Name should be specified. If both are specified, HotwordsId has a higher priority than Name.**
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`

	// Hotword lexicon name.
	// **Note: Either HotwordsId or Name should be specified. If both are specified, HotwordsId has a higher priority than Name.**
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Paging offset. Default value: 0.
	// 
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Hotword sorting field. Valid values:
	// 
	//  - Default: Sort by the hotword upload sequence.
	//  - Weight: Sort by the weight.
	//  - Lexical: Sort by the first letter of hotwords.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Hotword sorting order. 0: ascending (default); 1: descending.
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type DescribeAsrHotwordsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the hotword lexicon to be queried.
	// **Note: Either HotwordsId or Name should be specified. If both are specified, HotwordsId has a higher priority than Name.**
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`

	// Hotword lexicon name.
	// **Note: Either HotwordsId or Name should be specified. If both are specified, HotwordsId has a higher priority than Name.**
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Paging offset. Default value: 0.
	// 
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Hotword sorting field. Valid values:
	// 
	//  - Default: Sort by the hotword upload sequence.
	//  - Weight: Sort by the weight.
	//  - Lexical: Sort by the first letter of hotwords.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Hotword sorting order. 0: ascending (default); 1: descending.
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *DescribeAsrHotwordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsrHotwordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HotwordsId")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsrHotwordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsrHotwordsResponseParams struct {
	// ID of the hotword lexicon to be queried.
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`

	// Current status of the hotword lexicon corresponding to the ID. The value 0 indicates that no template is bound to this hotword lexicon when the query is performed and that the hotword lexicon can be deleted.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Hotword lexicon name.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The value is 0 for a temporary hotword lexicon, and the string provided during creation is returned.
	// The value is 1 for a file-based hotword lexicon, and the content of the file uploaded during creation is returned.
	// 
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Name of the uploaded hotword file.
	// Note: This field may return null, indicating that no valid value can be obtained.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// List of hotwords returned for the query.
	HotWords []*AsrHotwordsSetItem `json:"HotWords,omitnil,omitempty" name:"HotWords"`

	// Hotword text, which depends on the value of Type.
	// If the value of Type is 0, the hotword string is returned.
	// If the value of Type is 1, the base64-encoded content of the hotword file is returned.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Number of words contained in the hotword lexicon.
	// Note: This field may return null, indicating that no valid value can be obtained.
	WordCount *uint64 `json:"WordCount,omitnil,omitempty" name:"WordCount"`

	// Paging offset. Default value: 0.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Creation time of the hotword lexicon in ISO datetime format (UTC time). For example, "2006-01-02T15:04:05Z".Note: This field may return null, indicating that no valid value can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Modification time of the hotword lexicon in ISO datetime format (UTC time). For example, "2006-01-02T15:04:05Z".
	// Note: This field may return null, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAsrHotwordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsrHotwordsResponseParams `json:"Response"`
}

func (r *DescribeAsrHotwordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsrHotwordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchTaskDetailRequestParams struct {
	// Video processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeBatchTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Video processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeBatchTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchTaskDetailResponseParams struct {
	// Task type. Currently, the valid values include:
	// <Li>BatchTask: batch processing task for video workflows.</li>.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task status. Valid values:
	// <Li>WAITING: waiting.</li>
	// <Li>PROCESSING: processing.</li>
	// <li>FINISH: completed.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task creation time in [ISO datetime format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Task execution start time in [ISO datetime format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// Task execution completion time in [ISO datetime format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Media processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Video processing task information. This field has a value only when TaskType is BatchTask.
	// Note: This field may return null, indicating that no valid value can be obtained.
	BatchTaskResult *BatchSubTaskResult `json:"BatchTaskResult,omitnil,omitempty" name:"BatchTaskResult"`

	// Event notification information of the task.
	// Note: This field may return null, indicating that no valid value can be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Priority of the task flow, with a value range of [-10, 10].
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// An identifier for deduplication. If there has been a request with the same identifier within the past seven days, an error will be returned for the current request. The maximum length is 50 characters. Leaving it blank or using a null string indicates no deduplication is required.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Source context, which is used to pass through user request information. The callback for task flow status changes will return the value of this field. The maximum length is 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// Additional information field, only used in specific scenarios.
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeBatchTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentReviewTemplatesRequestParams struct {
	// The IDs of the content moderation templates to query. Array length limit: 50.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The filter for querying templates. If this parameter is left empty, both preset and custom templates are returned. Valid values:
	// * Preset
	// * Custom
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for intelligent auditing template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeContentReviewTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// The IDs of the content moderation templates to query. Array length limit: 50.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The filter for querying templates. If this parameter is left empty, both preset and custom templates are returned. Valid values:
	// * Preset
	// * Custom
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for intelligent auditing template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Type")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContentReviewTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentReviewTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of content audit template details.
	ContentReviewTemplateSet []*ContentReviewTemplateItem `json:"ContentReviewTemplateSet,omitnil,omitempty" name:"ContentReviewTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeImageSpriteTemplatesRequestParams struct {
	// Unique ID filter of image sprite generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for sprite template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeImageSpriteTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID filter of image sprite generating templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for sprite template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageSpriteTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSpriteTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of image sprite generating template details.
	ImageSpriteTemplateSet []*ImageSpriteTemplate `json:"ImageSpriteTemplateSet,omitnil,omitempty" name:"ImageSpriteTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeImageTaskDetailRequestParams struct {
	// Image processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeImageTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Image processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeImageTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageTaskDetailResponseParams struct {
	// Task type. Currently, the valid values include:
	// <Li>WorkflowTask: workflow processing task.</li>
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task status. Valid values:
	// <Li>WAITING: waiting.</li>
	// <Li>PROCESSING: processing.</li>
	// <li>FINISH: completed.</li>
	// Note: This field may return null, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Execution status and results of the image processing task.
	// Note: This field may return null, indicating that no valid value can be obtained.
	ImageProcessTaskResultSet []*ImageProcessTaskResult `json:"ImageProcessTaskResultSet,omitnil,omitempty" name:"ImageProcessTaskResultSet"`

	// Task creation time in [ISO datetime format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid value can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Task execution completion time in [ISO datetime format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid value can be obtained.
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeImageTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveRecordTemplatesRequestParams struct {
	// Specifies the recording template unique identifier filter condition, with an array length limit of 100.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Indicates the template type filter condition. If left empty, all templates are returned. Valid values:
	// * Preset: System preset template;
	// * Custom
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Specifies the recording template identifier filter condition, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeLiveRecordTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the recording template unique identifier filter condition, with an array length limit of 100.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Indicates the template type filter condition. If left empty, all templates are returned. Valid values:
	// * Preset: System preset template;
	// * Custom
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Specifies the recording template identifier filter condition, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeLiveRecordTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveRecordTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveRecordTemplatesResponseParams struct {
	// Total number of records that meet filter conditions.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Recording template details list.
	LiveRecordTemplateSet []*LiveRecordTemplate `json:"LiveRecordTemplateSet,omitnil,omitempty" name:"LiveRecordTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveRecordTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveRecordTemplatesResponseParams `json:"Response"`
}

func (r *DescribeLiveRecordTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaMetaDataRequestParams struct {
	// Input information of file for metadata getting.
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`
}

type DescribeMediaMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// Input information of file for metadata getting.
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`
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

// Predefined struct for user
type DescribeMediaMetaDataResponseParams struct {
	// Media metadata.
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMediaMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaMetaDataResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePersonSamplesRequestParams struct {
	// Type of images to pull. Valid values:
	// <li>UserDefine: custom image library</li>
	// <li>Default: default image library</li>
	// 
	// Default value: UserDefine. Samples in the custom image library will be pulled.
	// Note: you can pull the default image library only using the image name or a combination of the image name and ID, and only one face image is returned.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Image ID. Array length limit: 100
	PersonIds []*string `json:"PersonIds,omitnil,omitempty" name:"PersonIds"`

	// Image name. Array length limit: 20
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// Image tag. Array length limit: 20
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePersonSamplesRequest struct {
	*tchttp.BaseRequest
	
	// Type of images to pull. Valid values:
	// <li>UserDefine: custom image library</li>
	// <li>Default: default image library</li>
	// 
	// Default value: UserDefine. Samples in the custom image library will be pulled.
	// Note: you can pull the default image library only using the image name or a combination of the image name and ID, and only one face image is returned.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Image ID. Array length limit: 100
	PersonIds []*string `json:"PersonIds,omitnil,omitempty" name:"PersonIds"`

	// Image name. Array length limit: 20
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// Image tag. Array length limit: 20
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribePersonSamplesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Image information
	PersonSet []*AiSamplePerson `json:"PersonSet,omitnil,omitempty" name:"PersonSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeQualityControlTemplatesRequestParams struct {
	// Filter condition for media quality inspection template unique identifiers, with an array length limit of 100.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries.
	// 
	// <li>Default value: 10.</li>
	// <li>Maximum value: 100.</li>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// "Preset": preset template, "Custom": custom template
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for media quality inspection template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeQualityControlTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Filter condition for media quality inspection template unique identifiers, with an array length limit of 100.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries.
	// 
	// <li>Default value: 10.</li>
	// <li>Maximum value: 100.</li>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// "Preset": preset template, "Custom": custom template
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for media quality inspection template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeQualityControlTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityControlTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQualityControlTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityControlTemplatesResponseParams struct {
	// Total number of records that meet filter conditions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Media quality inspection template details list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	QualityControlTemplateSet []*QualityControlTemplate `json:"QualityControlTemplateSet,omitnil,omitempty" name:"QualityControlTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQualityControlTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQualityControlTemplatesResponseParams `json:"Response"`
}

func (r *DescribeQualityControlTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityControlTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSampleSnapshotTemplatesRequestParams struct {
	// Unique ID filter of sampled screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for sampled screenshot template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeSampleSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID filter of sampled screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for sampled screenshot template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSampleSnapshotTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSampleSnapshotTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of sampled screencapturing template details.
	SampleSnapshotTemplateSet []*SampleSnapshotTemplate `json:"SampleSnapshotTemplateSet,omitnil,omitempty" name:"SampleSnapshotTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSchedulesRequestParams struct {
	// The IDs of the schemes to query. Array length limit: 100.
	ScheduleIds []*int64 `json:"ScheduleIds,omitnil,omitempty" name:"ScheduleIds"`

	// The trigger type. Valid values:
	// <li>`CosFileUpload`: The scheme is triggered when a file is uploaded to Tencent Cloud Object Storage (COS).</li>
	// <li>`AwsS3FileUpload`: The scheme is triggered when a file is uploaded to AWS S3.</li>
	// If you do not specify this parameter or leave it empty, all schemes will be returned regardless of the trigger type.
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// The scheme status. Valid values:
	// <li>`Enabled`</li>
	// <li>`Disabled`</li>
	// If you do not specify this parameter, all schemes will be returned regardless of the status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The maximum number of records to return. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSchedulesRequest struct {
	*tchttp.BaseRequest
	
	// The IDs of the schemes to query. Array length limit: 100.
	ScheduleIds []*int64 `json:"ScheduleIds,omitnil,omitempty" name:"ScheduleIds"`

	// The trigger type. Valid values:
	// <li>`CosFileUpload`: The scheme is triggered when a file is uploaded to Tencent Cloud Object Storage (COS).</li>
	// <li>`AwsS3FileUpload`: The scheme is triggered when a file is uploaded to AWS S3.</li>
	// If you do not specify this parameter or leave it empty, all schemes will be returned regardless of the trigger type.
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// The scheme status. Valid values:
	// <li>`Enabled`</li>
	// <li>`Disabled`</li>
	// If you do not specify this parameter, all schemes will be returned regardless of the status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The pagination offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The maximum number of records to return. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSchedulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleIds")
	delete(f, "TriggerType")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSchedulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulesResponseParams struct {
	// The total number of records that meet the conditions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The information of the schemes.
	ScheduleInfoSet []*SchedulesInfo `json:"ScheduleInfoSet,omitnil,omitempty" name:"ScheduleInfoSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSchedulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSchedulesResponseParams `json:"Response"`
}

func (r *DescribeSchedulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmartSubtitleTemplatesRequestParams struct {
	// Unique identifiers of smart subtitle templates for filtering. The array can contain up to 100 unique identifiers.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Condition for filtering templates by type. If this field is not specified, all templates are returned. Valid values:
	// * Preset: system preset template
	// * Custom: user-defined template
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Condition for filtering smart subtitle templates by ID. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeSmartSubtitleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique identifiers of smart subtitle templates for filtering. The array can contain up to 100 unique identifiers.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Condition for filtering templates by type. If this field is not specified, all templates are returned. Valid values:
	// * Preset: system preset template
	// * Custom: user-defined template
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Condition for filtering smart subtitle templates by ID. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeSmartSubtitleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmartSubtitleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definitions")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmartSubtitleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmartSubtitleTemplatesResponseParams struct {
	// Total number of records that meet filter conditions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of smart subtitle template details.
	SmartSubtitleTemplateSet []*SmartSubtitleTemplateItem `json:"SmartSubtitleTemplateSet,omitnil,omitempty" name:"SmartSubtitleTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSmartSubtitleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmartSubtitleTemplatesResponseParams `json:"Response"`
}

func (r *DescribeSmartSubtitleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmartSubtitleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotByTimeOffsetTemplatesRequestParams struct {
	// Unique ID filter of time point screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for time point screenshot template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeSnapshotByTimeOffsetTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID filter of time point screencapturing templates. Array length limit: 100.
	Definitions []*uint64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition for time point screenshot template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotByTimeOffsetTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotByTimeOffsetTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of time point screencapturing template details.
	SnapshotByTimeOffsetTemplateSet []*SnapshotByTimeOffsetTemplate `json:"SnapshotByTimeOffsetTemplateSet,omitnil,omitempty" name:"SnapshotByTimeOffsetTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeStreamLinkSecurityGroupRequestParams struct {
	// Security group ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeStreamLinkSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Security group ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeStreamLinkSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkSecurityGroupResponseParams struct {
	// Information on the input security group.
	Info *SecurityGroupInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamLinkSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// Video processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Video processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
	// Task type. Valid values:<li>WorkflowTask: video workflow processing task.</li><li>EditMediaTask: video editing task.</li><li>LiveStreamProcessTask: live stream processing task.</li><li>ScheduleTask: orchestration processing task.</li><li>EvaluationTask: evaluation task.</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task status. Valid values:
	// <li>WAITING: Waiting;</li>
	// <li>PROCESSING: Processing;</li>
	// <li>FINISH: Completed.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Creation time of a task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Start time of task execution in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// End time of task execution in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Video editing task information. This field has a value only when `TaskType` is `EditMediaTask`.
	EditMediaTask *EditMediaTask `json:"EditMediaTask,omitnil,omitempty" name:"EditMediaTask"`

	// Information of a video processing task. This field has a value only when `TaskType` is `WorkflowTask`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WorkflowTask *WorkflowTask `json:"WorkflowTask,omitnil,omitempty" name:"WorkflowTask"`

	// Information of a live stream processing task. This field has a value only when `TaskType` is `LiveStreamProcessTask`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LiveStreamProcessTask *LiveStreamProcessTask `json:"LiveStreamProcessTask,omitnil,omitempty" name:"LiveStreamProcessTask"`

	// Event notification information of a task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Task flow priority. Value range: [-10, 10].
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// Extended information field, used in specific scenarios.
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// The information of a scheme. This parameter is valid only if `TaskType` is `ScheduleTask`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScheduleTask *ScheduleTask `json:"ScheduleTask,omitnil,omitempty" name:"ScheduleTask"`

	// The information of a live scheme. This parameter is valid only if `TaskType` is `LiveScheduleTask`.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LiveScheduleTask *LiveScheduleTask `json:"LiveScheduleTask,omitnil,omitempty" name:"LiveScheduleTask"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Filter: Task status. Valid values: WAITING (waiting), PROCESSING (processing), FINISH (completed).
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Scrolling identifier which is used for pulling in batches. If a single request cannot pull all the data entries, the API will return `ScrollToken`, and if the next request carries it, the next pull will start from the next entry.
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// Filter: Task status. Valid values: WAITING (waiting), PROCESSING (processing), FINISH (completed).
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Scrolling identifier which is used for pulling in batches. If a single request cannot pull all the data entries, the API will return `ScrollToken`, and if the next request carries it, the next pull will start from the next entry.
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`
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

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// Task overview list.
	TaskSet []*TaskSimpleInfo `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// Scrolling identifier. If a request does not return all the data entries, this field indicates the ID of the next entry. If this field is an empty string, there is no more data.
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// The total number of records that meet the conditions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Unique ID filter of transcoding templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Container format filter. Valid values:
	// <li>Video: Video container format that can contain both video stream and audio stream;</li>
	// <li>PureAudio: Audio container format that can contain only audio stream.</li>
	ContainerType *string `json:"ContainerType,omitnil,omitempty" name:"ContainerType"`

	// TESHD filter, which is used to filter common transcoding or ultra-fast HD transcoding templates. Valid values:
	// <li>Common: Common transcoding template;</li>
	// <li>TEHD: TESHD template.</li>
	TEHDType *string `json:"TEHDType,omitnil,omitempty" name:"TEHDType"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The template type (replacing `TEHDType`). Valid values:
	// <li>Common: Common transcoding template</li>
	// <li>TEHD: TESHD template</li>
	// <li>Enhance: Audio/Video enhancement template.</li>
	// This parameter is left empty by default, which indicates to return all types of templates.
	TranscodeType *string `json:"TranscodeType,omitnil,omitempty" name:"TranscodeType"`

	// Filter condition for transcoding template identifiers, with a length limit of 64 characters.	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Video scenario. Optional values: 
	// normal: General transcoding scenario: General transcoding and compression scenario. 
	// pgc: PGC HD TV shows and movies: At the time of compression, focus is placed on the viewing experience of TV shows and movies and ROI encoding is performed according to their characteristics, while high-quality contents of videos and audio are retained. 
	// materials_video: HD materials: Scenario involving material resources, where requirements for image quality are extremely high and there are many transparent images, with almost no visual loss during compression. 
	// ugc: UGC content: It is suitable for a wide range of UGC/short video scenarios, with an optimized encoding bitrate for short video characteristics, improved image quality, and enhanced business QOS/QOE metrics. 
	// e-commerce_video: Fashion show/e-commerce: At the time of compression, emphasis is placed on detail clarity and ROI enhancement, with a particular focus on maintaining the image quality of the face region. 
	// educational_video: Education: At the time of compression, emphasis is placed on the clarity and readability of text and images to help students better understand the content, ensuring that the teaching content is clearly conveyed. 
	// no_config: Not configured.
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Transcoding policy. Optional values: 
	// ultra_compress: Extreme compression: Compared to standard compression, this policy can maximize bitrate compression while ensuring a certain level of image quality, thus greatly saving bandwidth and storage costs. 
	// standard_compress: Comprehensively optimal: The compression ratio and image quality are balanced, and files are compressed as much as possible without a noticeable reduction in subjective image quality. Only audio and video TSC transcoding fees are charged for this policy. 
	// high_compress: Bitrate priority: Priority is given to reducing file size, which may result in certain image quality loss. Only audio and video TSC transcoding fees are charged for this policy. 
	// low_compress: Image quality priority: Priority is given to ensuring image quality, and the size of compressed files may be relatively large. Only audio and video TSC transcoding fees are charged for this policy. 
	// no_config: Not configured.
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
}

type DescribeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID filter of transcoding templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Template type filter. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Container format filter. Valid values:
	// <li>Video: Video container format that can contain both video stream and audio stream;</li>
	// <li>PureAudio: Audio container format that can contain only audio stream.</li>
	ContainerType *string `json:"ContainerType,omitnil,omitempty" name:"ContainerType"`

	// TESHD filter, which is used to filter common transcoding or ultra-fast HD transcoding templates. Valid values:
	// <li>Common: Common transcoding template;</li>
	// <li>TEHD: TESHD template.</li>
	TEHDType *string `json:"TEHDType,omitnil,omitempty" name:"TEHDType"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The template type (replacing `TEHDType`). Valid values:
	// <li>Common: Common transcoding template</li>
	// <li>TEHD: TESHD template</li>
	// <li>Enhance: Audio/Video enhancement template.</li>
	// This parameter is left empty by default, which indicates to return all types of templates.
	TranscodeType *string `json:"TranscodeType,omitnil,omitempty" name:"TranscodeType"`

	// Filter condition for transcoding template identifiers, with a length limit of 64 characters.	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Video scenario. Optional values: 
	// normal: General transcoding scenario: General transcoding and compression scenario. 
	// pgc: PGC HD TV shows and movies: At the time of compression, focus is placed on the viewing experience of TV shows and movies and ROI encoding is performed according to their characteristics, while high-quality contents of videos and audio are retained. 
	// materials_video: HD materials: Scenario involving material resources, where requirements for image quality are extremely high and there are many transparent images, with almost no visual loss during compression. 
	// ugc: UGC content: It is suitable for a wide range of UGC/short video scenarios, with an optimized encoding bitrate for short video characteristics, improved image quality, and enhanced business QOS/QOE metrics. 
	// e-commerce_video: Fashion show/e-commerce: At the time of compression, emphasis is placed on detail clarity and ROI enhancement, with a particular focus on maintaining the image quality of the face region. 
	// educational_video: Education: At the time of compression, emphasis is placed on the clarity and readability of text and images to help students better understand the content, ensuring that the teaching content is clearly conveyed. 
	// no_config: Not configured.
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Transcoding policy. Optional values: 
	// ultra_compress: Extreme compression: Compared to standard compression, this policy can maximize bitrate compression while ensuring a certain level of image quality, thus greatly saving bandwidth and storage costs. 
	// standard_compress: Comprehensively optimal: The compression ratio and image quality are balanced, and files are compressed as much as possible without a noticeable reduction in subjective image quality. Only audio and video TSC transcoding fees are charged for this policy. 
	// high_compress: Bitrate priority: Priority is given to reducing file size, which may result in certain image quality loss. Only audio and video TSC transcoding fees are charged for this policy. 
	// low_compress: Image quality priority: Priority is given to ensuring image quality, and the size of compressed files may be relatively large. Only audio and video TSC transcoding fees are charged for this policy. 
	// no_config: Not configured.
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
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
	delete(f, "TranscodeType")
	delete(f, "Name")
	delete(f, "SceneType")
	delete(f, "CompressType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of transcoding template details.
	TranscodeTemplateSet []*TranscodeTemplate `json:"TranscodeTemplateSet,omitnil,omitempty" name:"TranscodeTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeWatermarkTemplatesRequestParams struct {
	// Unique ID filter of watermarking templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Watermark type filter. Valid values:
	// <li>image: Image watermark;</li>
	// <li>text: Text watermark.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries
	// <li>Default value: 10;</li>
	// <li>Maximum value: 100.</li>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter condition for watermark template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeWatermarkTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID filter of watermarking templates. Array length limit: 100.
	Definitions []*int64 `json:"Definitions,omitnil,omitempty" name:"Definitions"`

	// Watermark type filter. Valid values:
	// <li>image: Image watermark;</li>
	// <li>text: Text watermark.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries
	// <li>Default value: 10;</li>
	// <li>Maximum value: 100.</li>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter condition for watermark template identifiers, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWatermarkTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWatermarkTemplatesResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of watermarking template details.
	WatermarkTemplateSet []*WatermarkTemplate `json:"WatermarkTemplateSet,omitnil,omitempty" name:"WatermarkTemplateSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Keyword filter. Array length limit: 100 words.
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// You can select multiple elements, which are connected by OR logic. If a usage contains any element in this parameter, the keyword sample will be used.
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Tag filter. Array length limit: 20 words.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeWordSamplesRequest struct {
	*tchttp.BaseRequest
	
	// Keyword filter. Array length limit: 100 words.
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// You can select multiple elements, which are connected by OR logic. If a usage contains any element in this parameter, the keyword sample will be used.
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Tag filter. Array length limit: 20 words.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Paging offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeWordSamplesResponseParams struct {
	// Number of eligible entries.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Keyword information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WordSet []*AiSampleWord `json:"WordSet,omitnil,omitempty" name:"WordSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeWorkflowsRequestParams struct {
	// Workflow ID filter. Array length limit: 100.
	WorkflowIds []*int64 `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// Workflow status. Valid values:
	// <li>Enabled: Enabled,</li>
	// <li>Disabled: Disabled.</li>
	// If this parameter is left empty, the workflow status will not be distinguished.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Paging offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeWorkflowsRequest struct {
	*tchttp.BaseRequest
	
	// Workflow ID filter. Array length limit: 100.
	WorkflowIds []*int64 `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// Workflow status. Valid values:
	// <li>Enabled: Enabled,</li>
	// <li>Disabled: Disabled.</li>
	// If this parameter is left empty, the workflow status will not be distinguished.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Paging offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned entries. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeWorkflowsResponseParams struct {
	// Number of eligible entries.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Workflow information array.
	WorkflowInfoSet []*WorkflowInfo `json:"WorkflowInfoSet,omitnil,omitempty" name:"WorkflowInfoSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowsResponseParams `json:"Response"`
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

type DiagnoseResult struct {
	// Diagnosed exception category. Valid values:
	// DecodeParamException: decoding parameter exception.
	// TimeStampException: timestamp exception.
	// FrameException: frame rate exception.
	// StreamStatusException: stream status exception.
	// StreamInfo: stream information exception.
	// StreamAbnormalCharacteristics: stream characteristics exception.
	// DecodeException: decoding exception.
	// HLSRequirements: HLS format exception.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// Diagnosed specific exception type. Valid values:
	// 
	// VideoResolutionChanged: video resolution change.
	// AudioSampleRateChanged: audio sample rate change.
	// AudioChannelsChanged: audio channel quantity change.ParameterSetsChanged: stream parameter set information change.
	// DarOrSarInvalid: video aspect ratio exception.
	// TimestampFallback: DTS timestamp rollback.DtsJitter: DTS jitter too high.
	// PtsJitter: PTS jitter too high.
	// AACDurationDeviation: improper AAC frame timestamp interval.
	// AudioDroppingFrames: audio frame dropping.
	// VideoDroppingFrames: video frame dropping.
	// AVTimestampInterleave: improper audio-video interleaving.
	// PtsLessThanDts: PTS less than DTS for media streams.
	// ReceiveFpsJitter: significant jitter in the network receive frame rate.ReceiveFpsTooSmall: network receive video frame rate too low.FpsJitter: significant jitter in the stream frame rate calculated via PTS.StreamOpenFailed: stream open failure.
	// StreamEnd: stream end.
	// StreamParseFailed: stream parsing failure.
	// VideoFirstFrameNotIdr: first frame not an IDR frame.
	// StreamNALUError: NALU start code error.
	// TsStreamNoAud: no AUD NALU in the H26x stream of MPEG-TS.AudioStreamLack: no audio stream.
	// VideoStreamLack: no video stream.
	// LackAudioRecover: missing audio stream recovery.
	// LackVideoRecover: missing video stream recovery.
	// VideoBitrateOutofRange: video stream bitrate (kbps) out of range.
	// AudioBitrateOutofRange: audio stream bitrate (kbps) out of range.
	// VideoDecodeFailed: video decoding error.
	// AudioDecodeFailed: audio decoding error.
	// AudioOutOfPhase: opposite phase in dual-channel audio.
	// VideoDuplicatedFrame: duplicate frames in video streams.
	// AudioDuplicatedFrame: duplicate frames in audio streams.
	// VideoRotation: video rotation.
	// TsMultiPrograms: multiple programs in MPEG2-TS streams.Mp4InvalidCodecFourcc: codec FourCC in MP4 not meeting Apple HLS requirements.
	// HLSBadM3u8Format: invalid M3U8 file.
	// HLSInvalidMasterM3u8: invalid main M3U8 file.
	// HLSInvalidMediaM3u8: invalid media M3U8 file.
	// HLSMasterM3u8Recommended: parameters recommended by standards missing in main M3U8.
	// HLSMediaM3u8Recommended: parameters recommended by standards missing in media M3U8.
	// HLSMediaM3u8DiscontinuityExist: EXT-X-DISCONTINUITY in media M3U8.
	// HLSMediaSegmentsStreamNumChange: changed number of streams in segments.
	// HLSMediaSegmentsPTSJitterDeviation: PTS jumps between segments without EXT-X-DISCONTINUITY.
	// HLSMediaSegmentsDTSJitterDeviation: DTS jumps between segments without EXT-X-DISCONTINUITY.
	// TimecodeTrackExist: TMCD track in MP4.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`


	Timestamp *float64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`


	Description *string `json:"Description,omitnil,omitempty" name:"Description"`


	DateTime *string `json:"DateTime,omitnil,omitempty" name:"DateTime"`

	// Diagnosed exception level. Valid values:
	// Fatal: affecting subsequent playback and parsing.
	// Error: may affect playback.
	// Warning: potential risk, which may not necessarily affect playback.
	// Notice: important stream information.
	// Info: general stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SeverityLevel *string `json:"SeverityLevel,omitnil,omitempty" name:"SeverityLevel"`
}

// Predefined struct for user
type DisableScheduleRequestParams struct {
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`
}

type DisableScheduleRequest struct {
	*tchttp.BaseRequest
	
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`
}

func (r *DisableScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableScheduleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DisableScheduleResponseParams `json:"Response"`
}

func (r *DisableScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWorkflowRequestParams struct {
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DisableWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
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

// Predefined struct for user
type DisableWorkflowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *DisableWorkflowResponseParams `json:"Response"`
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

type DrmInfo struct {
	// Encryption type.
	// <li>simpleaes: AES-128 encryption</li>
	// <li> widevine</li>
	// <li>fairplay: not supported for DASH streams</li>
	// <li> playready</li>
	// Note: This field may return null, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The AES-128 encryption details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SimpleAesDrm *SimpleAesDrm `json:"SimpleAesDrm,omitnil,omitempty" name:"SimpleAesDrm"`

	// Information about FairPlay, WideVine, and PlayReady encryption.
	SpekeDrm *SpekeDrm `json:"SpekeDrm,omitnil,omitempty" name:"SpekeDrm"`
}

type EditMediaFileInfo struct {
	// Video input information.
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// The start offset (seconds) for video clipping. This parameter is valid for video clipping tasks.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// The end offset (seconds) for video clipping. This parameter is valid for video clipping tasks.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// The ID of the material associated with an element. This parameter is required for video compositing tasks.
	// 
	// Note: The ID can be up to 32 characters long and can contain letters, digits, and special characters -_
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type EditMediaOutputConfig struct {
	// The container. Valid values: `mp4` (default), `hls`, `mov`, `flv`, `avi`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Editing mode. Optional values:
	// normal (default): Precise editing
	// fast: Fast editing, with faster processing speed but lower precision to some extent
	// Note: fast only supports individual files, and the default output transcoding format of normal is h264.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type EditMediaRequestParams struct {
	// Information of input video file.
	FileInfos []*EditMediaFileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// The storage location of the media processing output file.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The path to save the media processing output file.
	// 
	// Note: For complex compositing tasks, the filename can be up to 64 characters long and can only contain digits, letters, and special characters -_
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// The output settings for a video clipping task.
	OutputConfig *EditMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// The settings for a video compositing task.
	// 
	// Note: If this parameter is not empty, the task is a video compositing task. Otherwise, the task is a video clipping task.
	ComposeConfig *ComposeMediaConfig `json:"ComposeConfig,omitnil,omitempty" name:"ComposeConfig"`

	// Event notification information of task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Task priority. The higher the value, the higher the priority. Value range: -10 - 10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// The ID used for deduplication. If there was a request with the same ID in the last three days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type EditMediaRequest struct {
	*tchttp.BaseRequest
	
	// Information of input video file.
	FileInfos []*EditMediaFileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// The storage location of the media processing output file.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The path to save the media processing output file.
	// 
	// Note: For complex compositing tasks, the filename can be up to 64 characters long and can only contain digits, letters, and special characters -_
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// The output settings for a video clipping task.
	OutputConfig *EditMediaOutputConfig `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// The settings for a video compositing task.
	// 
	// Note: If this parameter is not empty, the task is a video compositing task. Otherwise, the task is a video clipping task.
	ComposeConfig *ComposeMediaConfig `json:"ComposeConfig,omitnil,omitempty" name:"ComposeConfig"`

	// Event notification information of task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Task priority. The higher the value, the higher the priority. Value range: -10 - 10. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// The ID used for deduplication. If there was a request with the same ID in the last three days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
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
	delete(f, "OutputConfig")
	delete(f, "ComposeConfig")
	delete(f, "TaskNotifyConfig")
	delete(f, "TasksPriority")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditMediaResponseParams struct {
	// Video editing task ID, which can be used to query the status of an editing task.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EditMediaResponse struct {
	*tchttp.BaseResponse
	Response *EditMediaResponseParams `json:"Response"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task status. Valid values:
	// <li>PROCESSING: processing;</li>
	// <li>FINISH: completed.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code
	// <li>0: success;</li>
	// <li>Other values: failure.</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input of video editing task.
	Input *EditMediaTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of video editing task.
	Output *EditMediaTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type EditMediaTaskInput struct {
	// Information of input video file.
	FileInfoSet []*EditMediaFileInfo `json:"FileInfoSet,omitnil,omitempty" name:"FileInfoSet"`
}

type EditMediaTaskOutput struct {
	// Target storage of edited file.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Path of edited video file.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

// Predefined struct for user
type EnableScheduleRequestParams struct {
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`
}

type EnableScheduleRequest struct {
	*tchttp.BaseRequest
	
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`
}

func (r *EnableScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableScheduleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableScheduleResponse struct {
	*tchttp.BaseResponse
	Response *EnableScheduleResponseParams `json:"Response"`
}

func (r *EnableScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWorkflowRequestParams struct {
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type EnableWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
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

// Predefined struct for user
type EnableWorkflowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *EnableWorkflowResponseParams `json:"Response"`
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

type EnhanceConfig struct {
	// Video enhancement configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoEnhance *VideoEnhanceConfig `json:"VideoEnhance,omitnil,omitempty" name:"VideoEnhance"`

	// The audio enhancement configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioEnhance *AudioEnhanceConfig `json:"AudioEnhance,omitnil,omitempty" name:"AudioEnhance"`
}

// Predefined struct for user
type ExecuteFunctionRequestParams struct {
	// Name of called backend API.
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// API parameter. Parameter format will depend on the actual function definition.
	FunctionArg *string `json:"FunctionArg,omitnil,omitempty" name:"FunctionArg"`
}

type ExecuteFunctionRequest struct {
	*tchttp.BaseRequest
	
	// Name of called backend API.
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// API parameter. Parameter format will depend on the actual function definition.
	FunctionArg *string `json:"FunctionArg,omitnil,omitempty" name:"FunctionArg"`
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

// Predefined struct for user
type ExecuteFunctionResponseParams struct {
	// Packed string, which will vary according to the custom API.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Switch of a face recognition task. Valid values:
	// <li>ON: Enables an intelligent face recognition task;</li>
	// <li>OFF: Disables an intelligent face recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Face recognition filter score. If this score is reached or exceeded, a recognition result will be returned. Value range: 0-100. Default value: 95.
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// The default face filter labels, which specify the types of faces to return. If this parameter is left empty, the detection results for all labels are returned. Valid values:
	// <li>entertainment (people in the entertainment industry)</li>
	// <li>sport (sports celebrities)</li>
	// <li>politician</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitnil,omitempty" name:"DefaultLibraryLabelSet"`

	// Custom face tags for filter, which specify the face recognition results to return. If this parameter is not specified or left empty, the recognition results for all custom face tags are returned.
	// Up to 100 tags are allowed, each containing no more than 16 characters.
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitnil,omitempty" name:"UserDefineLibraryLabelSet"`

	// Figure library. Valid values:
	// <li>Default: Default figure library;</li>
	// <li>UserDefine: Custom figure library.</li>
	// <li>All: Both default and custom figure libraries will be used.</li>
	// Default value: All (both default and custom figure libraries will be used.)
	FaceLibrary *string `json:"FaceLibrary,omitnil,omitempty" name:"FaceLibrary"`
}

type FaceConfigureInfoForUpdate struct {
	// Switch of a face recognition task. Valid values:
	// <li>ON: Enables an intelligent face recognition task;</li>
	// <li>OFF: Disables an intelligent face recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Face recognition filter score. If this score is reached or exceeded, a recognition result will be returned. Value range: 0-100.
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// The default face filter labels, which specify the types of faces to return. If this parameter is left empty, the detection results for all labels are returned. Valid values:
	// <li>entertainment (people in the entertainment industry)</li>
	// <li>sport (sports celebrities)</li>
	// <li>politician</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitnil,omitempty" name:"DefaultLibraryLabelSet"`

	// Custom face tags for filter, which specify the face recognition results to return. If this parameter is not specified or left empty, the recognition results for all custom face tags are returned.
	// Up to 100 tags are allowed, each containing no more than 16 characters.
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitnil,omitempty" name:"UserDefineLibraryLabelSet"`

	// Figure library. Valid values:
	// <li>Default: Default figure library;</li>
	// <li>UserDefine: Custom figure library.</li>
	// <li>All: Both default and custom figure libraries will be used.</li>
	FaceLibrary *string `json:"FaceLibrary,omitnil,omitempty" name:"FaceLibrary"`
}

type FaceEnhanceConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The strength. Value range: 0.0-1.0
	// Default value: 0.0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Intensity *float64 `json:"Intensity,omitnil,omitempty" name:"Intensity"`
}

type FrameRateConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The frame rate (Hz). Value range: [0, 100].
	// Default value: 0.
	// Note: For transcoding, this parameter will overwrite `Fps` of `VideoTemplate`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`
}

type FrameTagConfigureInfo struct {
	// Switch of intelligent frame-specific tagging task. Valid values:
	// <li>ON: enables intelligent frame-specific tagging task;</li>
	// <li>OFF: disables intelligent frame-specific tagging task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type FrameTagConfigureInfoForUpdate struct {
	// Switch of intelligent frame-specific tagging task. Valid values:
	// <li>ON: enables intelligent frame-specific tagging task;</li>
	// <li>OFF: disables intelligent frame-specific tagging task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type HLSConfigureInfo struct {
	// Duration of a single TS file in seconds. Value range: 5-30 seconds.
	// 
	// If this parameter is left empty, 30 seconds will be used by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	ItemDuration *int64 `json:"ItemDuration,omitnil,omitempty" name:"ItemDuration"`

	// Recording cycle in seconds. Value range: 10 minutes to 12 hours.
	// 
	// If this parameter is left empty, 10 minutes (3600 seconds) will be used by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Resume recording waiting time, unit: seconds. Value range: 60-1800 seconds.
	// If this parameter is left empty, 0 (resume recording not enabled) will be used by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	ContinueTimeout *int64 `json:"ContinueTimeout,omitnil,omitempty" name:"ContinueTimeout"`
}

type HdrConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Type. Valid values:
	// <li>HDR10</li>
	// <li>HLG</li>
	// Default Value: HDR10.
	// Note: The video encoding method should be H.265.
	// Note: The video encoding bit depth is 10.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type HeadTailParameter struct {
	// The opening segments.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HeadSet []*MediaInputInfo `json:"HeadSet,omitnil,omitempty" name:"HeadSet"`

	// The closing segments.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TailSet []*MediaInputInfo `json:"TailSet,omitnil,omitempty" name:"TailSet"`
}

type HighlightSegmentItem struct {
	// The confidence score.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The start time offset of the segment.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// The end time offset of the segment.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Segment tag.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SegmentTags []*string `json:"SegmentTags,omitnil,omitempty" name:"SegmentTags"`

	// The live streaming segment corresponds to the live start time point, in the ISO date format.	
	// Note: This field may return null, indicating that no valid value can be obtained.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The live streaming segment corresponds to the live streaming end time, in the ISO date format.	
	// Note: This field may return null, indicating that no valid value can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ImageAreaBoxInfo struct {
	// Type of the box selection area in the image. Valid values:
	// <li>logo: icon.</li>
	// <li>Text: text.</li>
	// Default value: logo.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Coordinates (pixel-level) of the box selection area in the image. Format: [x1, y1, x2, y2], which indicates the coordinates of the top left corner and the bottom right corner.
	// For example, [101, 85, 111, 95].
	// Note: This field may return null, indicating that no valid value can be obtained.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// Coordinates of the box selection area in the image. Format: [x1, y1, x2, y2], which indicates the coordinates of the top left corner and the bottom right corner. This parameter takes effect when AreaCoordSet is not specified.
	//  - [0.1, 0.1, 0.3, 0.3]: Indicates the ratio (values are less than 1).
	//  -[50, 50, 350, 280]: Indicates the pixel (values are greater than or equal to 1).
	// Note: This field may return null, indicating that no valid value can be obtained.
	BoundingBox []*float64 `json:"BoundingBox,omitnil,omitempty" name:"BoundingBox"`
}

type ImageDenoiseConfig struct {
	// Capability configuration enabling status. Valid values:
	// <li>ON: enabled.</li>
	// <li>OFF: disabled.</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Type, with valid values including:
	// <li>weak</li>
	// <li>strong</li>
	// Default value: weak.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ImageEncodeConfig struct {
	// Image format. Valid values: JPEG, PNG, BMP, and WebP. If it is not specified, the original image format is used. Animations are not supported.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Relative image quality. Valid range: 1 - 100. The value is based on the original image quality, and the default is the original image quality.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Quality *int64 `json:"Quality,omitnil,omitempty" name:"Quality"`
}

type ImageEnhanceConfig struct {
	// Super-resolution configuration.
	SuperResolution *SuperResolutionConfig `json:"SuperResolution,omitnil,omitempty" name:"SuperResolution"`

	// Denoising configuration.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Denoise *ImageDenoiseConfig `json:"Denoise,omitnil,omitempty" name:"Denoise"`

	// Comprehensive enhancement configuration.
	// Note: This field may return null, indicating that no valid value can be obtained.
	ImageQualityEnhance *ImageQualityEnhanceConfig `json:"ImageQualityEnhance,omitnil,omitempty" name:"ImageQualityEnhance"`

	// Color enhancement configuration.
	ColorEnhance *ColorEnhanceConfig `json:"ColorEnhance,omitnil,omitempty" name:"ColorEnhance"`

	// Detail enhancement configuration.
	SharpEnhance *SharpEnhanceConfig `json:"SharpEnhance,omitnil,omitempty" name:"SharpEnhance"`

	// Face enhancement configuration.
	FaceEnhance *FaceEnhanceConfig `json:"FaceEnhance,omitnil,omitempty" name:"FaceEnhance"`

	// Low-light enhancement configuration.
	// Note: This field may return null, indicating that no valid value can be obtained.
	LowLightEnhance *LowLightEnhanceConfig `json:"LowLightEnhance,omitnil,omitempty" name:"LowLightEnhance"`
}

type ImageEraseConfig struct {
	// Icon erasing configuration.
	// Note: This field may return null, indicating that no valid value can be obtained.
	ImageEraseLogo *ImageEraseLogoConfig `json:"ImageEraseLogo,omitnil,omitempty" name:"ImageEraseLogo"`
}

type ImageEraseLogoConfig struct {
	// Capability configuration enabling status. Valid values:
	// <li>ON: enabled</li>
	// <li>OFF: disabled</li>
	// Default value: ON.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Multiple box selection areas that need to be erased, with a maximum of 16 areas available.
	// Note: This field may return null, indicating that no valid value can be obtained.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	ImageAreaBoxes []*ImageAreaBoxInfo `json:"ImageAreaBoxes,omitnil,omitempty" name:"ImageAreaBoxes"`
}

type ImageProcessTaskOutput struct {
	// Path of the output file.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Storage location of the output file.
	// Note: This field may return null, indicating that no valid value can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`
}

type ImageProcessTaskResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error message.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Transcoding task output.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Output *ImageProcessTaskOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// Transcoding progress, with a value range of [0-100].
	// Note: This field may return null, indicating that no valid value can be obtained.
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type ImageQualityEnhanceConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The strength. Valid values:
	// <li>weak</li>
	// <li>normal</li>
	// <li>strong</li>
	// Default value: weak.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ImageSpriteTaskInput struct {
	// ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Target bucket of a generated image sprite. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Output path of a captured sprite image file, which can be a relative or absolute path.
	// If you need to define an output path, the path must end with `.{format}`. For variable names, refer to [Filename Variable](https://intl.cloud.tencent.com/document/product/862/37039?from_cn_redirect=1).Relative path example:
	// <li>Filename_{Variable name}.{format}.</li>
	// <li>Filename.{format}.</li>
	// Absolute path example:
	// <li>/Custom path/Filename_{Variable name}.{format}.</li>
	// If left empty, a relative path is used by default: `{inputName}_imageSprite_{definition}_{number}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// Output path to the WebVTT file after an image sprite is generated, which can only be a relative path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_imageSprite_{definition}.{format}`.
	WebVttObjectName *string `json:"WebVttObjectName,omitnil,omitempty" name:"WebVttObjectName"`

	// Rule of the `{number}` variable in the image sprite output path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitnil,omitempty" name:"ObjectNumberFormat"`
}

type ImageSpriteTemplate struct {
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Name of an image sprite generating template.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Subimage width of an image sprite.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Subimage height of an image sprite.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Sampling type.
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// Subimage row count of an image sprite.
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// Subimage column count of an image sprite.
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: Stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: Fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// Template description.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// The image format.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type ImageTaskInput struct {
	// Image encoding configuration.
	// Note: This field may return null, indicating that no valid value can be obtained.
	EncodeConfig *ImageEncodeConfig `json:"EncodeConfig,omitnil,omitempty" name:"EncodeConfig"`

	// Image enhancement configuration.
	// Note: This field may return null, indicating that no valid value can be obtained.
	EnhanceConfig *ImageEnhanceConfig `json:"EnhanceConfig,omitnil,omitempty" name:"EnhanceConfig"`

	// Image erasing configuration.
	// Note: This field may return null, indicating that no valid value can be obtained.
	EraseConfig *ImageEraseConfig `json:"EraseConfig,omitnil,omitempty" name:"EraseConfig"`
}

type ImageWatermarkInput struct {
	// String generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) a watermark image. JPEG and PNG images are supported.
	ImageContent *string `json:"ImageContent,omitnil,omitempty" name:"ImageContent"`

	// Width of a watermark, supporting two formats: % and px.
	// <li>If a string ends with %, it indicates that the `Width` of a watermark is a percentage of a video's width. For example, `10%` means that `Width` is 10% of a video's width.</li>
	// <li>If a string ends with px, the `Width` of a watermark will be in pixels. For example, `100px` means that `Width` is 100 pixels. Value range: [8, 4096].</li>
	// 
	// When width and height are not specified or set to 0, the default value is 10%.
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height. For example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in pixels. For example, `100px` means that `Height` is 100 pixels. Value range: 0 or [8, 4096].</li>
	// Default value: 0px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// Repeat type of an animated watermark. Valid values:
	// <li>once: no longer appears after watermark playback ends.</li>
	// <li>repeat_last_frame: stays on the last frame after watermark playback ends.</li>
	// <li>repeat (default): repeats the playback until the video ends.</li>
	RepeatType *string `json:"RepeatType,omitnil,omitempty" name:"RepeatType"`
}

type ImageWatermarkInputForUpdate struct {
	// String generated by [Base64-encoding](https://tools.ietf.org/html/rfc4648) a watermark image. JPEG and PNG images are supported.
	ImageContent *string `json:"ImageContent,omitnil,omitempty" name:"ImageContent"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width. For example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in pixels. For example, `100px` means that `Width` is 100 pixels. Value range: [8, 4096].</li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// Height of a watermark, supporting two formats: % and px.
	// <li>If a string ends with %, it indicates that the `Height` of a watermark is a percentage of a video's height. For example, `10%` means that `Height` is 10% of a video's height.</li>
	// <li>If a string ends with px, the `Height` of a watermark will be in pixels. For example, `100px` means that `Height` is 100 pixels. Value range: 0 or [8, 4096].</li>
	// 
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// Repeat type of an animated watermark. Valid values:
	// <li>once: no longer appears after watermark playback ends.</li>
	// <li>repeat_last_frame: stays on the last frame after watermark playback ends.</li>
	// <li>repeat (default): repeats the playback until the video ends.</li>
	RepeatType *string `json:"RepeatType,omitnil,omitempty" name:"RepeatType"`
}

type ImageWatermarkTemplate struct {
	// Watermark image address.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.</li>
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height. For example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in pixels. For example, `100px` means that `Height` is 100 pixels.</li>
	// `0px` means that `Height` will be proportionally scaled according to the video width.
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// Repeat type of an animated watermark. Valid values:
	// <li>once: no longer appears after watermark playback ends.</li>
	// <li>repeat_last_frame: stays on the last frame after watermark playback ends.</li>
	// <li>repeat (default): repeats the playback until the video ends.</li>
	RepeatType *string `json:"RepeatType,omitnil,omitempty" name:"RepeatType"`
}

type LiveActivityResItem struct {
	// The output of a live recording task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LiveRecordTask *LiveScheduleLiveRecordTaskResult `json:"LiveRecordTask,omitnil,omitempty" name:"LiveRecordTask"`

	// Media quality inspection task output.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LiveQualityControlTask *ScheduleQualityControlTaskResult `json:"LiveQualityControlTask,omitnil,omitempty" name:"LiveQualityControlTask"`
}

type LiveActivityResult struct {
	// Atomic task type.
	// <li>LiveRecord: live recording.</li>
	// <li>AiQualityControl: media quality inspection.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActivityType *string `json:"ActivityType,omitnil,omitempty" name:"ActivityType"`

	// The task output.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LiveActivityResItem *LiveActivityResItem `json:"LiveActivityResItem,omitnil,omitempty" name:"LiveActivityResItem"`
}

type LiveRecordFile struct {
	// The URL of the recording file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// The size of the recording file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// The duration of the recording file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// The recording start time in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The recording end time in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type LiveRecordResult struct {
	// The storage of the recording file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The recording segments.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileList []*LiveRecordFile `json:"FileList,omitnil,omitempty" name:"FileList"`
}

type LiveRecordTaskInput struct {
	// The live recording template ID.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The storage of the recording file. If this parameter is left empty, the `OutputStorage` value of the parent folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The output path of the recording file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`
}

type LiveRecordTemplate struct {
	// Specifies the recording template unique identifier.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// HLS configuration parameters
	HLSConfigure *HLSConfigureInfo `json:"HLSConfigure,omitnil,omitempty" name:"HLSConfigure"`

	// MP4 configuration parameter.
	MP4Configure *MP4ConfigureInfo `json:"MP4Configure,omitnil,omitempty" name:"MP4Configure"`

	// Recording template name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Template type. Valid values:
	// <li>Preset: system-preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type LiveScheduleLiveRecordTaskResult struct {
	// The task status. Valid values: `PROCESSING`, `SUCCESS`, `FAIL`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value indicates the task has failed. For details, see [Error Codes](https://www.tencentcloud.com/document/product/1041/40249).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// The error code. `0` indicates the task is successful; other values indicate the task has failed. This parameter is not recommended. Please use `ErrCodeExt` instead.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// The error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input of a live recording task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Input *LiveRecordTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of a live recording task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *LiveRecordResult `json:"Output,omitnil,omitempty" name:"Output"`

	// The time when the task was started, in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid values can be obtained.
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// The time when the task was completed, in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid values can be obtained.
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type LiveScheduleTask struct {
	// The ID of a live scheme subtask.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The task status. Valid values:
	// <li>`PROCESSING`</li>
	// <li>`FINISH` </li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// If the value returned is not `0`, there was a source error. If `0` is returned, refer to the error codes of the corresponding task type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// If there was a source error, this parameter is the error message. For other errors, refer to the error messages of the corresponding task type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The URL of the live stream.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// The task output.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LiveActivityResultSet []*LiveActivityResult `json:"LiveActivityResultSet,omitnil,omitempty" name:"LiveActivityResultSet"`
}

type LiveStreamAiAnalysisResultInfo struct {

	ResultSet []*LiveStreamAiAnalysisResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`
}

type LiveStreamAiAnalysisResultItem struct {

	Type *string `json:"Type,omitnil,omitempty" name:"Type"`


	SegmentResultSet []*SegmentRecognitionItem `json:"SegmentResultSet,omitnil,omitempty" name:"SegmentResultSet"`
}

type LiveStreamAiQualityControlResultInfo struct {
	// Content quality inspection result list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: QualityControlResults is deprecated.
	QualityControlResults []*QualityControlResult `json:"QualityControlResults,omitnil,omitempty" name:"QualityControlResults"`


	//
	// Deprecated: DiagnoseResults is deprecated.
	DiagnoseResults []*DiagnoseResult `json:"DiagnoseResults,omitnil,omitempty" name:"DiagnoseResults"`

	// Content quality inspection result list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	QualityControlResultSet []*QualityControlResult `json:"QualityControlResultSet,omitnil,omitempty" name:"QualityControlResultSet"`

	// Format diagnostic result list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiagnoseResultSet []*DiagnoseResult `json:"DiagnoseResultSet,omitnil,omitempty" name:"DiagnoseResultSet"`
}

type LiveStreamAiRecognitionResultInfo struct {
	// Content recognition result list.
	ResultSet []*LiveStreamAiRecognitionResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`
}

type LiveStreamAiRecognitionResultItem struct {
	// Result type. Valid values:
	// <li>FaceRecognition: face recognition.</li>
	// <li>AsrWordsRecognition: speech keyword recognition.</li>
	// <li>OcrWordsRecognition: text keyword recognition.</li>
	// <li>AsrFullTextRecognition: full speech recognition.</li>
	// <li>OcrFullTextRecognition: full text recognition.</li>
	// <li>TransTextRecognition: speech translation.</li>
	// 
	// <li>ObjectRecognition: object recognition.</li>
	// <li>TagRecognition: highlights marking.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Face recognition result, which is valid when `Type` is
	// `FaceRecognition`.
	FaceRecognitionResultSet []*LiveStreamFaceRecognitionResult `json:"FaceRecognitionResultSet,omitnil,omitempty" name:"FaceRecognitionResultSet"`

	// Speech keyword recognition result, which is valid when `Type` is
	// `AsrWordsRecognition`.
	AsrWordsRecognitionResultSet []*LiveStreamAsrWordsRecognitionResult `json:"AsrWordsRecognitionResultSet,omitnil,omitempty" name:"AsrWordsRecognitionResultSet"`

	// Text keyword recognition result, which is valid when `Type` is
	// `OcrWordsRecognition`.
	OcrWordsRecognitionResultSet []*LiveStreamOcrWordsRecognitionResult `json:"OcrWordsRecognitionResultSet,omitnil,omitempty" name:"OcrWordsRecognitionResultSet"`

	// Full speech recognition result, which is valid when `Type` is
	// `AsrFullTextRecognition`.
	AsrFullTextRecognitionResultSet []*LiveStreamAsrFullTextRecognitionResult `json:"AsrFullTextRecognitionResultSet,omitnil,omitempty" name:"AsrFullTextRecognitionResultSet"`

	// Full text recognition result, which is valid when `Type` is
	// `OcrFullTextRecognition`.
	OcrFullTextRecognitionResultSet []*LiveStreamOcrFullTextRecognitionResult `json:"OcrFullTextRecognitionResultSet,omitnil,omitempty" name:"OcrFullTextRecognitionResultSet"`

	// The translation result. This parameter is valid only if `Type` is `TransTextRecognition`.
	TransTextRecognitionResultSet []*LiveStreamTransTextRecognitionResult `json:"TransTextRecognitionResultSet,omitnil,omitempty" name:"TransTextRecognitionResultSet"`

	// Object recognition result, which is valid when Type is ObjectRecognition.
	ObjectRecognitionResultSet []*LiveStreamObjectRecognitionResult `json:"ObjectRecognitionResultSet,omitnil,omitempty" name:"ObjectRecognitionResultSet"`


	TagRecognitionResultSet []*LiveStreamTagRecognitionResult `json:"TagRecognitionResultSet,omitnil,omitempty" name:"TagRecognitionResultSet"`
}

type LiveStreamAiReviewImagePoliticalResult struct {
	// Start PTS time of a suspected segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// End PTS time of a suspected segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// The confidence score for the detected sensitive segments.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for porn information detection of a suspected segment. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// The labels for the detected sensitive information. Valid values:
	// <li>politician</li>
	// <li>violation_photo (banned icons)</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// The name of a sensitive person or banned icon.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The pixel coordinates of the detected sensitive people or banned icons. The format is [x1, y1, x2, y2], which indicates the coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type LiveStreamAiReviewImagePornResult struct {
	// Start PTS time of a suspected segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// End PTS time of a suspected segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// Score of a suspected porn segment.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for porn information detection of a suspected segment. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Tag of the detected porn information in video. Valid values:
	// <li>porn: Porn.</li>
	// <li>sexy: Sexiness.</li>
	// <li>vulgar: Vulgarity.</li>
	// <li>intimacy: Intimacy.</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type LiveStreamAiReviewImageTerrorismResult struct {
	// Start PTS time of a suspected segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// End PTS time of a suspected segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// The confidence score for the detected sensitive segments.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The suggestion for handling the sensitive segments. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// The labels for the detected sensitive content. Valid values:
	// <li>guns</li>
	// <li>crowd</li>
	// <li>police</li>
	// <li>bloody</li>
	// <li>banners (sensitive flags)</li>
	// <li>militant</li>
	// <li>explosion</li>
	// <li>terrorists</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type LiveStreamAiReviewResultInfo struct {
	// List of content audit results.
	ResultSet []*LiveStreamAiReviewResultItem `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`
}

type LiveStreamAiReviewResultItem struct {
	// The type of moderation result. Valid values:
	// <li>ImagePorn</li>
	// <li>ImageTerrorism</li>
	// <li>ImagePolitical</li>
	// <li>VoicePorn</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Result of porn information detection in image, which is valid when `Type` is `ImagePorn`.
	ImagePornResultSet []*LiveStreamAiReviewImagePornResult `json:"ImagePornResultSet,omitnil,omitempty" name:"ImagePornResultSet"`

	// The result of detecting sensitive information in images, which is valid if `Type` is `ImageTerrorism`.
	ImageTerrorismResultSet []*LiveStreamAiReviewImageTerrorismResult `json:"ImageTerrorismResultSet,omitnil,omitempty" name:"ImageTerrorismResultSet"`

	// The result of detecting sensitive information in images, which is valid if `Type` is `ImagePolitical`.
	ImagePoliticalResultSet []*LiveStreamAiReviewImagePoliticalResult `json:"ImagePoliticalResultSet,omitnil,omitempty" name:"ImagePoliticalResultSet"`

	// The result for moderation of pornographic content in audio. This parameter is valid if `Type` is `VoicePorn`.
	VoicePornResultSet []*LiveStreamAiReviewVoicePornResult `json:"VoicePornResultSet,omitnil,omitempty" name:"VoicePornResultSet"`
}

type LiveStreamAiReviewVoicePornResult struct {
	// Start PTS time of a suspected segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// End PTS time of a suspected segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// Score of a suspected porn segment.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for porn information detection of a suspected segment. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Tag of the detected porn information in video. Valid values:
	// <li>sexual_moan: Sexual moans.</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`
}

type LiveStreamAsrFullTextRecognitionResult struct {
	// Recognized text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type LiveStreamAsrWordsRecognitionResult struct {
	// Speech keyword.
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type LiveStreamFaceRecognitionResult struct {
	// Unique ID of figure.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Figure name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Figure library type, indicating to which figure library the recognized figure belongs:
	// <li>Default: default figure library</li><li>UserDefine: custom figure library</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type LiveStreamObjectRecognitionResult struct {
	// Name of a recognized object.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Start PTS time of a recognized segment, in seconds.
	StartPtsOffset *float64 `json:"StartPtsOffset,omitnil,omitempty" name:"StartPtsOffset"`

	// End PTS time of a recognized segment, in seconds.
	EndPtsOffset *float64 `json:"EndPtsOffset,omitnil,omitempty" name:"EndPtsOffset"`

	// Confidence of a recognized segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Zone coordinates of the recognition result. An array contains four elements: [x1, y1, x2, y2], representing the horizontal and vertical coordinates of the top-left and bottom-right corners, respectively.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// Screenshot link.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type LiveStreamOcrFullTextRecognitionResult struct {
	// Speech text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type LiveStreamOcrWordsRecognitionResult struct {
	// Text keyword.
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// Start PTS time of recognized segment in seconds.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// End PTS time of recognized segment in seconds.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// Confidence of recognized segment. Value range: 0–100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Zone coordinates of recognition result. The array contains four elements: [x1,y1,x2,y2], i.e., the horizontal and vertical coordinates of the top-left and bottom-right corners.
	AreaCoords []*int64 `json:"AreaCoords,omitnil,omitempty" name:"AreaCoords"`
}

type LiveStreamProcessErrorInfo struct {
	// Error code:
	// <li>0: No error;</li>
	// <li>If this parameter is not 0, an error has occurred. Please see the error message (`Message`).</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type LiveStreamProcessTask struct {
	// The media processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task flow status. Valid values:
	// <li>PROCESSING: Processing;</li>
	// <li>FINISH: Completed.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. 0: success; other values: failure.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Live stream URL.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type LiveStreamRecordResultInfo struct {
	// Whether recording ends.
	// 0: Recording does not end, returning a single file.
	// 1: Recording ends, returning all recording files.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecordOver *uint64 `json:"RecordOver,omitnil,omitempty" name:"RecordOver"`

	// File list.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileResults []*LiveRecordFile `json:"FileResults,omitnil,omitempty" name:"FileResults"`
}

type LiveStreamTagRecognitionResult struct {

	Id *string `json:"Id,omitnil,omitempty" name:"Id"`


	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`


	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`


	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type LiveStreamTaskNotifyConfig struct {
	// Notification type:
	// 
	// "CMQ": Callback messages are written to the CMQ queue; 
	// "URL": When a URL is specified, the HTTP callback is pushed to the address specified by NotifyUrl. The callback protocol is http+json. The content of the packet body is the same as the output parameters of the [ParseLiveStreamProcessNotification API](https://intl.cloud.tencent.com/document/product/862/39229?from_cn_redirect=1).
	// 
	// <font color="red">Note: If left blank, it is CMQ by default. To use the other type, you need to fill in the corresponding type value.</font>
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// HTTP callback URL, required if `NotifyType` is set to `URL`
	NotifyUrl *string `json:"NotifyUrl,omitnil,omitempty" name:"NotifyUrl"`

	// CMQ model. There are two types: `Queue` and `Topic`. Currently, only `Queue` is supported.
	CmqModel *string `json:"CmqModel,omitnil,omitempty" name:"CmqModel"`

	// CMQ region, such as `sh` and `bj`.
	CmqRegion *string `json:"CmqRegion,omitnil,omitempty" name:"CmqRegion"`

	// This parameter is valid when the model is `Queue`, indicating the name of the CMQ queue for receiving event notifications.
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// This parameter is valid when the model is `Topic`, indicating the name of the CMQ topic for receiving event notifications.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Key used to generate a callback signature.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NotifyKey *string `json:"NotifyKey,omitnil,omitempty" name:"NotifyKey"`
}

type LiveStreamTransTextRecognitionResult struct {
	// The text transcript.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// The PTS (seconds) of the start of a segment.
	StartPtsTime *float64 `json:"StartPtsTime,omitnil,omitempty" name:"StartPtsTime"`

	// The PTS (seconds) of the end of a segment.
	EndPtsTime *float64 `json:"EndPtsTime,omitnil,omitempty" name:"EndPtsTime"`

	// The confidence score for a segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The translation.
	Trans *string `json:"Trans,omitnil,omitempty" name:"Trans"`
}

type LowLightEnhanceConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The strength. Valid values:
	// <li>normal</li>
	// Default value: normal.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type MP4ConfigureInfo struct {
	// Recording duration, in seconds. The interval can range from 10 minutes to 720 minutes. It is 60 minutes (3,600 seconds) by default.
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`
}

// Predefined struct for user
type ManageTaskRequestParams struct {
	// Operation type. Valid values:
	// <ul>
	// <li>Abort: task termination. Description:
	// <ul><li>If the [task type](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) is live stream processing (`LiveStreamProcessTask`), tasks whose [task status](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) is `WAITING` or `PROCESSING` can be terminated.</li>
	// <li>For other [task types](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0), only tasks whose [task status](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) is `WAITING` can be terminated.</li></ul>
	// </li></ul>
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Video processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type ManageTaskRequest struct {
	*tchttp.BaseRequest
	
	// Operation type. Valid values:
	// <ul>
	// <li>Abort: task termination. Description:
	// <ul><li>If the [task type](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) is live stream processing (`LiveStreamProcessTask`), tasks whose [task status](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) is `WAITING` or `PROCESSING` can be terminated.</li>
	// <li>For other [task types](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0), only tasks whose [task status](https://intl.cloud.tencent.com/document/product/862/37614?from_cn_redirect=1#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0) is `WAITING` can be terminated.</li></ul>
	// </li></ul>
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Video processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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

// Predefined struct for user
type ManageTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type MediaAiAnalysisClassificationItem struct {
	// Name of intelligently generated category.
	Classification *string `json:"Classification,omitnil,omitempty" name:"Classification"`

	// Confidence of intelligently generated category between 0 and 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type MediaAiAnalysisCoverItem struct {
	// Storage path of intelligently generated cover.
	CoverPath *string `json:"CoverPath,omitnil,omitempty" name:"CoverPath"`

	// Confidence of intelligently generated cover between 0 and 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type MediaAiAnalysisDescriptionItem struct {
	// Intelligent description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Confidence of the intelligent description, with a value range from 0 to 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Intelligent description title.
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Intelligent description keywords.
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// Segmentation result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Paragraphs []*AiParagraphInfo `json:"Paragraphs,omitnil,omitempty" name:"Paragraphs"`
}

type MediaAiAnalysisFrameTagItem struct {
	// Frame-specific tag name.
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`


	CategorySet []*string `json:"CategorySet,omitnil,omitempty" name:"CategorySet"`

	// Confidence of intelligently generated frame-specific tag between 0 and 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type MediaAiAnalysisFrameTagSegmentItem struct {
	// Start time offset of frame-specific tag.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of frame-specific tag.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// List of tags in time period.
	TagSet []*MediaAiAnalysisFrameTagItem `json:"TagSet,omitnil,omitempty" name:"TagSet"`
}

type MediaAiAnalysisHighlightItem struct {
	// The URL of the highlight segments.
	HighlightPath *string `json:"HighlightPath,omitnil,omitempty" name:"HighlightPath"`

	// The URL of the thumbnail.
	CovImgPath *string `json:"CovImgPath,omitnil,omitempty" name:"CovImgPath"`

	// The confidence score. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The duration of the highlights.
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// A list of the highlight segments.
	SegmentSet []*HighlightSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`
}

type MediaAiAnalysisTagItem struct {
	// Tag name.
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Confidence of tag between 0 and 100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`
}

type MediaAnimatedGraphicsItem struct {
	// Storage location of a generated animated image file.
	Storage *TaskOutputStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Path to a generated animated image file.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// ID of an animated image generating template. For more information, please see [Animated Image Generating Parameter Template](https://intl.cloud.tencent.com/document/product/266/33481?from_cn_redirect=1#.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Animated image format, such as gif.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Height of an animated image in px.
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Width of an animated image in px.
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Bitrate of an animated image in bps.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// Size of an animated image in bytes.
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// MD5 value of an animated image.
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// Start time offset of an animated image in the video in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of an animated image in the video in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type MediaAudioStreamItem struct {
	// Bitrate of an audio stream in bps.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// Sample rate of an audio stream in Hz.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SamplingRate *int64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// Audio stream codec, such as aac.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// Number of sound channels, e.g., 2
	// Note: this field may return `null`, indicating that no valid value was found.
	Channel *int64 `json:"Channel,omitnil,omitempty" name:"Channel"`
}

type MediaContentReviewAsrTextSegmentItem struct {
	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Confidence of a suspected segment.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for suspected segment audit. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// List of suspected keywords.
	KeywordSet []*string `json:"KeywordSet,omitnil,omitempty" name:"KeywordSet"`
}

type MediaContentReviewOcrTextSegmentItem struct {
	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Confidence of a suspected segment.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Suggestion for suspected segment audit. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// List of suspected keywords.
	KeywordSet []*string `json:"KeywordSet,omitnil,omitempty" name:"KeywordSet"`

	// Zone coordinates (at the pixel level) of suspected text: [x1, y1, x2, y2], i.e., the coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// URL of a suspected image (which will not be permanently stored
	// and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type MediaContentReviewPoliticalSegmentItem struct {
	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// The confidence score for the detected sensitive segments.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The suggestion for handling the sensitive segments. Valid values:
	// <li>pass</li>
	// <li>review</li>
	// <li>block</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// The name of a sensitive person or banned icon.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The labels for the detected sensitive segments. The relationship between the values of this parameter and those of the `LabelSet` parameter in [PoliticalImgReviewTemplateInfo](https://intl.cloud.tencent.com/document/api/862/37615?from_cn_redirect=1#PoliticalImgReviewTemplateInfo) is as follows:
	// violation_photo:
	// <li>violation_photo (banned icons)</li>
	// politician:
	// <li>nation_politician (state leader)</li>
	// <li>province_politician (provincial officials)</li>
	// <li>bureau_politician (bureau-level officials)</li>
	// <li>county_politician (county-level officials)</li>
	// <li>rural_politician (township-level officials)</li>
	// <li>sensitive_politician (sensitive people)</li>
	// <li>foreign_politician (state leaders of other countries)</li>
	// entertainment:
	// <li>sensitive_entertainment (sensitive people in the entertainment industry</li>
	// sport:
	// <li>sensitive_sport (sensitive sports celebrities)</li>
	// entrepreneur:
	// <li>sensitive_entrepreneur</li>
	// scholar:
	// <li>sensitive_scholar</li>
	// celebrity:
	// <li>sensitive_celebrity</li>
	// <li>historical_celebrity (sensitive historical figures)</li>
	// military:
	// <li>sensitive_military (sensitive people in military)</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// URL of a suspected image (which will not be permanently stored
	//  and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// The pixel coordinates of the detected sensitive people or banned icons. The format is [x1, y1, x2, y2], which indicates the coordinates of the top-left and bottom-right corners.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type MediaContentReviewSegmentItem struct {
	// Start time offset of a suspected segment in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a suspected segment in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Score of a suspected porn segment.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Tag of porn information detection result of a suspected segment.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// Suggestion for porn information detection of a suspected segment. Valid values:
	// <li>pass.</li>
	// <li>review.</li>
	// <li>block.</li>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// URL of a suspected image (which will not be permanently stored
	//  and will be deleted after `PicUrlExpireTime`).
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Expiration time of a suspected image URL in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	PicUrlExpireTime *string `json:"PicUrlExpireTime,omitnil,omitempty" name:"PicUrlExpireTime"`
}

type MediaImageSpriteItem struct {
	// Image sprite specification. For more information, please see [Image Sprite Parameter Template](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Subimage height of an image sprite.
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Subimage width of an image sprite.
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Total number of subimages in each image sprite.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Path to each image sprite.
	ImagePathSet []*string `json:"ImagePathSet,omitnil,omitempty" name:"ImagePathSet"`

	// Path to a WebVtt file for the position-time relationship among subimages in an image sprite. The WebVtt file indicates the corresponding time points of each subimage and their coordinates in the image sprite, which is typically used by the player for implementing preview.
	WebVttPath *string `json:"WebVttPath,omitnil,omitempty" name:"WebVttPath"`

	// Storage location of an image sprite file.
	Storage *TaskOutputStorage `json:"Storage,omitnil,omitempty" name:"Storage"`
}

type MediaInputInfo struct {
	// The input type. Valid values:
	// <li>`COS`: A COS bucket address.</li>
	// <li> `URL`: A URL.</li>
	// <li> `AWS-S3`: An AWS S3 bucket address. Currently, this type is only supported for transcoding tasks.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The information of the COS object to process. This parameter is valid and required when `Type` is `COS`.
	CosInputInfo *CosInputInfo `json:"CosInputInfo,omitnil,omitempty" name:"CosInputInfo"`

	// The URL of the object to process. This parameter is valid and required when `Type` is `URL`.
	// Note: This field may return null, indicating that no valid value can be obtained.
	UrlInputInfo *UrlInputInfo `json:"UrlInputInfo,omitnil,omitempty" name:"UrlInputInfo"`

	// The information of the AWS S3 object processed. This parameter is required if `Type` is `AWS-S3`.
	// Note: This field may return null, indicating that no valid value can be obtained.
	S3InputInfo *S3InputInfo `json:"S3InputInfo,omitnil,omitempty" name:"S3InputInfo"`
}

type MediaMetaData struct {
	// Size of an uploaded media file in bytes (which is the sum of size of m3u8 and ts files if the video is in HLS format).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Container, such as m4a and mp4.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Sum of the average bitrate of a video stream and that of an audio stream in bps.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// Maximum value of the height of a video stream in px.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Maximum value of the width of a video stream in px.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Video duration in seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Selected angle during video recording in degrees.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rotate *int64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`

	// Video stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitnil,omitempty" name:"VideoStreamSet"`

	// Audio stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitnil,omitempty" name:"AudioStreamSet"`

	// Video duration in seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoDuration *float64 `json:"VideoDuration,omitnil,omitempty" name:"VideoDuration"`

	// Audio duration in seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioDuration *float64 `json:"AudioDuration,omitnil,omitempty" name:"AudioDuration"`
}

type MediaProcessTaskAdaptiveDynamicStreamingResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input of an adaptive bitrate streaming task.
	Input *AdaptiveDynamicStreamingTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of an adaptive bitrate streaming task.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Output *AdaptiveDynamicStreamingInfoItem `json:"Output,omitnil,omitempty" name:"Output"`
}

type MediaProcessTaskAnimatedGraphicResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input for an animated image generating task.
	Input *AnimatedGraphicTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of an animated image generating task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaAnimatedGraphicsItem `json:"Output,omitnil,omitempty" name:"Output"`
}

type MediaProcessTaskImageSpriteResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input for an image sprite generating task.
	Input *ImageSpriteTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of an image sprite generating task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaImageSpriteItem `json:"Output,omitnil,omitempty" name:"Output"`
}

type MediaProcessTaskInput struct {
	// List of transcoding tasks.
	TranscodeTaskSet []*TranscodeTaskInput `json:"TranscodeTaskSet,omitnil,omitempty" name:"TranscodeTaskSet"`

	// List of animated image screenshot tasks.
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInput `json:"AnimatedGraphicTaskSet,omitnil,omitempty" name:"AnimatedGraphicTaskSet"`

	// List of time point screenshot tasks.
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTaskSet,omitnil,omitempty" name:"SnapshotByTimeOffsetTaskSet"`

	// List of sampled screenshot tasks.
	SampleSnapshotTaskSet []*SampleSnapshotTaskInput `json:"SampleSnapshotTaskSet,omitnil,omitempty" name:"SampleSnapshotTaskSet"`

	// List of image sprite screenshot tasks.
	ImageSpriteTaskSet []*ImageSpriteTaskInput `json:"ImageSpriteTaskSet,omitnil,omitempty" name:"ImageSpriteTaskSet"`

	// List of adaptive bitrate streaming tasks.
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTaskSet,omitnil,omitempty" name:"AdaptiveDynamicStreamingTaskSet"`
}

type MediaProcessTaskResult struct {
	// Task type. Valid values:
	// <li>Transcode: Transcoding</li>
	// <li>AnimatedGraphics: Animated image generating</li>
	// <li>SnapshotByTimeOffset: Time point screenshot</li>
	// <li>SampleSnapshot: Sampled screenshot</li>
	// <li>ImageSprites: Image sprite screenshot</li>
	// <li>CoverBySnapshot: Screenshot for cover image</li>
	// <li>AdaptiveDynamicStreaming: Adaptive bitrate streaming</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Query result of a transcoding task, which is valid when task type is `Transcode`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitnil,omitempty" name:"TranscodeTask"`

	// Query result of an animated image generating task, which is valid when task type is `AnimatedGraphics`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitnil,omitempty" name:"AnimatedGraphicTask"`

	// Query result of a time point screenshot task, which is valid when task type is `SnapshotByTimeOffset`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitnil,omitempty" name:"SnapshotByTimeOffsetTask"`

	// Query result of a sampled screenshot task, which is valid when task type is `SampleSnapshot`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitnil,omitempty" name:"SampleSnapshotTask"`

	// Query result of an image sprite screenshot task, which is valid when task type is `ImageSprite`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitnil,omitempty" name:"ImageSpriteTask"`

	// Query result of an adaptive bitrate streaming task, which is valid if the task type is `AdaptiveDynamicStreaming`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdaptiveDynamicStreamingTask *MediaProcessTaskAdaptiveDynamicStreamingResult `json:"AdaptiveDynamicStreamingTask,omitnil,omitempty" name:"AdaptiveDynamicStreamingTask"`
}

type MediaProcessTaskSampleSnapshotResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input for a sampled screenshot task.
	Input *SampleSnapshotTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of a sampled screenshot task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaSampleSnapshotItem `json:"Output,omitnil,omitempty" name:"Output"`


	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`


	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type MediaProcessTaskSnapshotByTimeOffsetResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input for a time point screenshot task.
	Input *SnapshotByTimeOffsetTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of a time point screenshot task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaSnapshotByTimeOffsetItem `json:"Output,omitnil,omitempty" name:"Output"`

	// The time when the task started executing, in ISO date format.
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// The time when the task finished, in ISO date format.
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type MediaProcessTaskTranscodeResult struct {
	// Task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates the task is successful; otherwise it is failed. This parameter is no longer recommended. Consider using the new error code parameter ErrCodeExt.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input for a transcoding task.
	Input *TranscodeTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output of a transcoding task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *MediaTranscodeItem `json:"Output,omitnil,omitempty" name:"Output"`

	// Transcoding progress. Value range: 0-100
	// Note: This field may return `null`, indicating that no valid value was found.
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type MediaSampleSnapshotItem struct {
	// Sampled screenshot specification ID. For more information, please see [Sampled Screenshot Parameter Template](https://intl.cloud.tencent.com/document/product/266/33480?from_cn_redirect=1#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Sample type. Valid values:
	// <li>Percent: Samples at the specified percentage interval.</li>
	// <li>Time: Samples at the specified time interval.</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval
	// <li>If `SampleType` is `Percent`, this value means taking a screenshot at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, this value means taking a screenshot at an interval of the specified time (in seconds). The first screenshot is always the first video frame.</li>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Storage location of a generated screenshot file.
	Storage *TaskOutputStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// List of paths to generated screenshots.
	ImagePathSet []*string `json:"ImagePathSet,omitnil,omitempty" name:"ImagePathSet"`

	// List of watermarking template IDs if the screenshots are watermarked.
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitnil,omitempty" name:"WaterMarkDefinition"`
}

type MediaSnapshotByTimeOffsetItem struct {
	// Specification of a time point screenshot template. 
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Information set of screenshots of the same specification. Each element represents a screenshot.
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet,omitnil,omitempty" name:"PicInfoSet"`

	// Location of a time point screenshot file.
	Storage *TaskOutputStorage `json:"Storage,omitnil,omitempty" name:"Storage"`
}

type MediaSnapshotByTimePicInfoItem struct {
	// The timestamp (seconds) of the screenshot.
	TimeOffset *float64 `json:"TimeOffset,omitnil,omitempty" name:"TimeOffset"`

	// Path to the screenshot.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// List of watermarking template IDs if the screenshots are watermarked.
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitnil,omitempty" name:"WaterMarkDefinition"`
}

type MediaTranscodeItem struct {
	// Target bucket of an output file.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Path to an output video file.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Transcoding specification ID. For more information, please see [Transcoding Parameter Template](https://intl.cloud.tencent.com/document/product/266/33478?from_cn_redirect=1#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Sum of the average bitrate of a video stream and that of an audio stream in bps.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// Maximum value of the height of a video stream in px.
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Maximum value of the width of a video stream in px.
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Total size of a media file in bytes (which is the sum of size of m3u8 and ts files if the video is in HLS format).
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Video duration in seconds.
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Container, such as m4a and mp4.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// MD5 value of a video.
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// Audio stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitnil,omitempty" name:"AudioStreamSet"`

	// Video stream information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitnil,omitempty" name:"VideoStreamSet"`

	// Enhancement items used for video transcoding. Descriptions of enhancement items:
	// <li>hdr: HDR configuration</li>
	// <li>wd_fps: configuration of frame interpolation for higher frame rate</li>
	// <li>video_super_resolution: 	super-resolution configuration</li>
	// <li>repair: comprehensive enhancement configuration</li>
	// <li>denoise: video denoising configuration</li>
	// <Li>color_enhance: color enhancement configuration</li>
	// <Li>scratch: scratch removal configuration</li>
	// <li>artifact: artifact (glitch) removal configuration</li>
	// <li>sharp: detail enhancement configuration</li>
	// <Li>low_light: low-light enhancement configuration</li>
	// <Li>face_enhance: face enhancement configuration</li>
	// Note: This field may return null, indicating that no valid value can be obtained.
	CallBackExtInfo *string `json:"CallBackExtInfo,omitnil,omitempty" name:"CallBackExtInfo"`
}

type MediaVideoStreamItem struct {
	// Bitrate of a video stream in bps.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// Height of a video stream in px.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Width of a video stream in px.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Video stream codec, such as h264.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// Frame rate in Hz.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Color primaries
	// Note: this field may return `null`, indicating that no valid value was found.
	ColorPrimaries *string `json:"ColorPrimaries,omitnil,omitempty" name:"ColorPrimaries"`

	// Color space
	// Note: this field may return `null`, indicating that no valid value was found.
	ColorSpace *string `json:"ColorSpace,omitnil,omitempty" name:"ColorSpace"`

	// Color transfer
	// Note: this field may return `null`, indicating that no valid value was found.
	ColorTransfer *string `json:"ColorTransfer,omitnil,omitempty" name:"ColorTransfer"`

	// HDR type
	// Note: This field may return `null`, indicating that no valid value was found.
	HdrType *string `json:"HdrType,omitnil,omitempty" name:"HdrType"`


	Codecs *string `json:"Codecs,omitnil,omitempty" name:"Codecs"`

	// Numerator of the frame rate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FpsNumerator *int64 `json:"FpsNumerator,omitnil,omitempty" name:"FpsNumerator"`

	// Denominator of the frame rate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FpsDenominator *int64 `json:"FpsDenominator,omitnil,omitempty" name:"FpsDenominator"`
}

// Predefined struct for user
type ModifyAIAnalysisTemplateRequestParams struct {
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Video content analysis template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Video content analysis template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Control parameter of intelligent categorization task.
	ClassificationConfigure *ClassificationConfigureInfoForUpdate `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// Control parameter of intelligent tagging task.
	TagConfigure *TagConfigureInfoForUpdate `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// Control parameter of intelligent cover generating task.
	CoverConfigure *CoverConfigureInfoForUpdate `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// Control parameter of intelligent frame-specific tagging task.
	FrameTagConfigure *FrameTagConfigureInfoForUpdate `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`
}

type ModifyAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of video content analysis template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Video content analysis template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Video content analysis template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Control parameter of intelligent categorization task.
	ClassificationConfigure *ClassificationConfigureInfoForUpdate `json:"ClassificationConfigure,omitnil,omitempty" name:"ClassificationConfigure"`

	// Control parameter of intelligent tagging task.
	TagConfigure *TagConfigureInfoForUpdate `json:"TagConfigure,omitnil,omitempty" name:"TagConfigure"`

	// Control parameter of intelligent cover generating task.
	CoverConfigure *CoverConfigureInfoForUpdate `json:"CoverConfigure,omitnil,omitempty" name:"CoverConfigure"`

	// Control parameter of intelligent frame-specific tagging task.
	FrameTagConfigure *FrameTagConfigureInfoForUpdate `json:"FrameTagConfigure,omitnil,omitempty" name:"FrameTagConfigure"`
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

// Predefined struct for user
type ModifyAIAnalysisTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Unique ID of a video content recognition template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of a video content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description of a video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Face recognition control parameter.
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// Full text recognition control parameter.
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// Text keyword recognition control parameter.
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// Full speech recognition control parameter.
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// Speech keyword recognition control parameter.
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`
}

type ModifyAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a video content recognition template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of a video content recognition template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description of a video content recognition template. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Face recognition control parameter.
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitnil,omitempty" name:"FaceConfigure"`

	// Full text recognition control parameter.
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitnil,omitempty" name:"OcrFullTextConfigure"`

	// Text keyword recognition control parameter.
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitnil,omitempty" name:"OcrWordsConfigure"`

	// Full speech recognition control parameter.
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitnil,omitempty" name:"AsrFullTextConfigure"`

	// Speech keyword recognition control parameter.
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitnil,omitempty" name:"AsrWordsConfigure"`
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

// Predefined struct for user
type ModifyAIRecognitionTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Unique ID of an adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Adaptive bitrate streaming format. Valid values:
	// <li>HLS,</li>
	// <li>MPEG-DASH.</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Whether to prohibit transcoding from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// Parameter information of input streams for transcoding to adaptive bitrate streaming. Up to 10 streams can be input.
	// 
	// Note:
	// 
	// 1. The frame rate of each stream must be consistent; otherwise, the frame rate of the first stream is used as the output frame rate.
	// 2. When modifying substream information, all field values must be fully modified and added; otherwise, fields not filled in will use default values.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Indicates whether it is audio-only. 0 means video template, 1 means audio-only template.
	// When the value is 1.
	// 1. StreamInfos.N.RemoveVideo=1
	// 2. StreamInfos.N.RemoveAudio=0
	// 3. StreamInfos.N.Video.Codec=copy
	// When the value is 0.
	// 1. StreamInfos.N.Video.Codec cannot be copy.
	// 2. StreamInfos.N.Video.Fps cannot be null.
	// 
	// Note:
	// 
	// This value only distinguishes template types. The task uses the values of RemoveAudio and RemoveVideo.
	PureAudio *uint64 `json:"PureAudio,omitnil,omitempty" name:"PureAudio"`

	// HLS segment type. Valid values: <li>ts-segment: HLS+TS segment.</li> <li>ts-byterange: HLS+TS byte range.</li> <li>mp4-segment: HLS+MP4 segment.</li> <li>mp4-byterange: HLS+MP4 byte range.</li> <li>ts-packed-audio: TS+Packed audio.</li> <li>mp4-packed-audio: MP4+Packed audio.</li> Default value: ts-segment.
	// Note: The HLS segment format for adaptive bitrate streaming is based on this field.
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
}

type ModifyAdaptiveDynamicStreamingTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an adaptive bitrate streaming template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Adaptive bitrate streaming format. Valid values:
	// <li>HLS,</li>
	// <li>MPEG-DASH.</li>
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Whether to prohibit transcoding from low bitrate to high bitrate. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoBitrate *uint64 `json:"DisableHigherVideoBitrate,omitnil,omitempty" name:"DisableHigherVideoBitrate"`

	// Whether to prohibit transcoding from low resolution to high resolution. Valid values:
	// <li>0: no,</li>
	// <li>1: yes.</li>
	DisableHigherVideoResolution *uint64 `json:"DisableHigherVideoResolution,omitnil,omitempty" name:"DisableHigherVideoResolution"`

	// Parameter information of input streams for transcoding to adaptive bitrate streaming. Up to 10 streams can be input.
	// 
	// Note:
	// 
	// 1. The frame rate of each stream must be consistent; otherwise, the frame rate of the first stream is used as the output frame rate.
	// 2. When modifying substream information, all field values must be fully modified and added; otherwise, fields not filled in will use default values.
	StreamInfos []*AdaptiveStreamTemplate `json:"StreamInfos,omitnil,omitempty" name:"StreamInfos"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Indicates whether it is audio-only. 0 means video template, 1 means audio-only template.
	// When the value is 1.
	// 1. StreamInfos.N.RemoveVideo=1
	// 2. StreamInfos.N.RemoveAudio=0
	// 3. StreamInfos.N.Video.Codec=copy
	// When the value is 0.
	// 1. StreamInfos.N.Video.Codec cannot be copy.
	// 2. StreamInfos.N.Video.Fps cannot be null.
	// 
	// Note:
	// 
	// This value only distinguishes template types. The task uses the values of RemoveAudio and RemoveVideo.
	PureAudio *uint64 `json:"PureAudio,omitnil,omitempty" name:"PureAudio"`

	// HLS segment type. Valid values: <li>ts-segment: HLS+TS segment.</li> <li>ts-byterange: HLS+TS byte range.</li> <li>mp4-segment: HLS+MP4 segment.</li> <li>mp4-byterange: HLS+MP4 byte range.</li> <li>ts-packed-audio: TS+Packed audio.</li> <li>mp4-packed-audio: MP4+Packed audio.</li> Default value: ts-segment.
	// Note: The HLS segment format for adaptive bitrate streaming is based on this field.
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`
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
	delete(f, "PureAudio")
	delete(f, "SegmentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAdaptiveDynamicStreamingTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAdaptiveDynamicStreamingTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of an animated image generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Maximum value of the width (or long side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Animated image format. Valid values: gif, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Video frame rate in Hz. Value range: [1, 30].
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Image quality. Value range: [1, 100]. Default value: 75.
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ModifyAnimatedGraphicsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an animated image generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of an animated image generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Maximum value of the width (or long side) of an animated image in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Animated image format. Valid values: gif, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Video frame rate in Hz. Value range: [1, 30].
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Image quality. Value range: [1, 100]. Default value: 75.
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
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

// Predefined struct for user
type ModifyAnimatedGraphicsTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyAsrHotwordsRequestParams struct {
	// Hotword lexicon ID. 
	//  
	// Either Name or Content should be specified if the hotword lexicon is a temporary hotword lexicon.
	// Either Name, FileContent, or FileName should be specified if the hotword lexicon is a file-based hotword lexicon.
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`

	// Hotword lexicon name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Hotword lexicon text.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Base64-encoded content of the hotword file. This field is required if Type is set to 1.
	// 
	// 
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// Name of the uploaded hotword file.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type ModifyAsrHotwordsRequest struct {
	*tchttp.BaseRequest
	
	// Hotword lexicon ID. 
	//  
	// Either Name or Content should be specified if the hotword lexicon is a temporary hotword lexicon.
	// Either Name, FileContent, or FileName should be specified if the hotword lexicon is a file-based hotword lexicon.
	HotwordsId *string `json:"HotwordsId,omitnil,omitempty" name:"HotwordsId"`

	// Hotword lexicon name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Hotword lexicon text.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Base64-encoded content of the hotword file. This field is required if Type is set to 1.
	// 
	// 
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// Name of the uploaded hotword file.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

func (r *ModifyAsrHotwordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAsrHotwordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HotwordsId")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "FileContent")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAsrHotwordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAsrHotwordsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAsrHotwordsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAsrHotwordsResponseParams `json:"Response"`
}

func (r *ModifyAsrHotwordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAsrHotwordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContentReviewTemplateRequestParams struct {
	// The unique ID of the content moderation template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The name of the content moderation template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Control parameter for porn information
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// Control parameter for terrorism information
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// Control parameter for politically sensitive information
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this parameter is not supported yet.
	ProhibitedConfigure *ProhibitedConfigureInfoForUpdate `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// Custom content moderation parameters.
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`
}

type ModifyContentReviewTemplateRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the content moderation template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// The name of the content moderation template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Control parameter for porn information
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitnil,omitempty" name:"PornConfigure"`

	// Control parameter for terrorism information
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitnil,omitempty" name:"TerrorismConfigure"`

	// Control parameter for politically sensitive information
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitnil,omitempty" name:"PoliticalConfigure"`

	// Control parameter of prohibited information detection. Prohibited information includes:
	// <li>Abusive;</li>
	// <li>Drug-related.</li>
	// Note: this parameter is not supported yet.
	ProhibitedConfigure *ProhibitedConfigureInfoForUpdate `json:"ProhibitedConfigure,omitnil,omitempty" name:"ProhibitedConfigure"`

	// Custom content moderation parameters.
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitnil,omitempty" name:"UserDefineConfigure"`
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

// Predefined struct for user
type ModifyContentReviewTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyImageSpriteTemplateRequestParams struct {
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of an image sprite generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Subimage width of an image sprite in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Subimage height of an image sprite in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Sampling type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// Subimage row count of an image sprite.
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// Subimage column count of an image sprite.
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// The image format. Valid values: jpg, png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type ModifyImageSpriteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an image sprite generating template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of an image sprite generating template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Subimage width of an image sprite in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Subimage height of an image sprite in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Sampling type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// Subimage row count of an image sprite.
	RowCount *uint64 `json:"RowCount,omitnil,omitempty" name:"RowCount"`

	// Subimage column count of an image sprite.
	ColumnCount *uint64 `json:"ColumnCount,omitnil,omitempty" name:"ColumnCount"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// The image format. Valid values: jpg, png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
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
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageSpriteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSpriteTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyLiveRecordTemplateRequestParams struct {
	// Specifies the recording template unique identifier.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// HLS configuration parameter. Either this parameter or MP4Configure should be specified.
	HLSConfigure *HLSConfigureInfo `json:"HLSConfigure,omitnil,omitempty" name:"HLSConfigure"`

	// MP4 configuration parameter. Either this parameter or HLSConfigure should be specified.
	MP4Configure *MP4ConfigureInfo `json:"MP4Configure,omitnil,omitempty" name:"MP4Configure"`

	// Recording template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description, with a length limit of 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ModifyLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the recording template unique identifier.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// HLS configuration parameter. Either this parameter or MP4Configure should be specified.
	HLSConfigure *HLSConfigureInfo `json:"HLSConfigure,omitnil,omitempty" name:"HLSConfigure"`

	// MP4 configuration parameter. Either this parameter or HLSConfigure should be specified.
	MP4Configure *MP4ConfigureInfo `json:"MP4Configure,omitnil,omitempty" name:"MP4Configure"`

	// Recording template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description, with a length limit of 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *ModifyLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "HLSConfigure")
	delete(f, "MP4Configure")
	delete(f, "Name")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveRecordTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveRecordTemplateResponseParams `json:"Response"`
}

func (r *ModifyLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonSampleRequestParams struct {
	// Image ID
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Name. Length limit: 128 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description. Length limit: 1,024 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Image usage. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: used for content recognition and inappropriate information recognition; equivalent to 1+2
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Information of operations on facial features
	FaceOperationInfo *AiSampleFaceOperation `json:"FaceOperationInfo,omitnil,omitempty" name:"FaceOperationInfo"`

	// Tag operation information.
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil,omitempty" name:"TagOperationInfo"`
}

type ModifyPersonSampleRequest struct {
	*tchttp.BaseRequest
	
	// Image ID
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// Name. Length limit: 128 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description. Length limit: 1,024 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Image usage. Valid values:
	// 1. Recognition: used for content recognition; equivalent to `Recognition.Face`
	// 2. Review: used for inappropriate information recognition; equivalent to `Review.Face`
	// 3. All: used for content recognition and inappropriate information recognition; equivalent to 1+2
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Information of operations on facial features
	FaceOperationInfo *AiSampleFaceOperation `json:"FaceOperationInfo,omitnil,omitempty" name:"FaceOperationInfo"`

	// Tag operation information.
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil,omitempty" name:"TagOperationInfo"`
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

// Predefined struct for user
type ModifyPersonSampleResponseParams struct {
	// Image information
	Person *AiSamplePerson `json:"Person,omitnil,omitempty" name:"Person"`

	// Information of images that failed the verification by facial feature positioning.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitnil,omitempty" name:"FailFaceInfoSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyQualityControlTemplateRequestParams struct {
	// Unique identifier of a media quality inspection template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Media quality inspection template name, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description, with a length limit of 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Media quality inspection configuration parameters.
	QualityControlItemSet []*QualityControlItemConfig `json:"QualityControlItemSet,omitnil,omitempty" name:"QualityControlItemSet"`
}

type ModifyQualityControlTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique identifier of a media quality inspection template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Media quality inspection template name, with a length limit of 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description, with a length limit of 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Media quality inspection configuration parameters.
	QualityControlItemSet []*QualityControlItemConfig `json:"QualityControlItemSet,omitnil,omitempty" name:"QualityControlItemSet"`
}

func (r *ModifyQualityControlTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQualityControlTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "QualityControlItemSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQualityControlTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQualityControlTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQualityControlTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQualityControlTemplateResponseParams `json:"Response"`
}

func (r *ModifyQualityControlTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQualityControlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySampleSnapshotTemplateRequestParams struct {
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Sampled screencapturing type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// The image format. Valid values: jpg, png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type ModifySampleSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a sampled screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Sampled screencapturing type. Valid values:
	// <li>Percent: By percent.</li>
	// <li>Time: By time interval.</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	// <li>If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.</li>
	// <li>If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.</li>
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// The image format. Valid values: jpg, png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
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

// Predefined struct for user
type ModifySampleSnapshotTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyScheduleRequestParams struct {
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// The scheme name.
	ScheduleName *string `json:"ScheduleName,omitnil,omitempty" name:"ScheduleName"`

	// The trigger of the scheme.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The subtasks of the scheme.
	// Note: You need to pass in the full list of subtasks even if you want to change only some of the subtasks.
	Activities []*Activity `json:"Activities,omitnil,omitempty" name:"Activities"`

	// The bucket to save the output file.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The directory to save the media processing output file, which must start and end with `/`.
	// Note: If this parameter is left empty, the current `OutputDir` value will be invalidated.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// The notification configuration.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Resource ID. Ensure the corresponding resource is in the enabled state.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type ModifyScheduleRequest struct {
	*tchttp.BaseRequest
	
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// The scheme name.
	ScheduleName *string `json:"ScheduleName,omitnil,omitempty" name:"ScheduleName"`

	// The trigger of the scheme.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The subtasks of the scheme.
	// Note: You need to pass in the full list of subtasks even if you want to change only some of the subtasks.
	Activities []*Activity `json:"Activities,omitnil,omitempty" name:"Activities"`

	// The bucket to save the output file.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The directory to save the media processing output file, which must start and end with `/`.
	// Note: If this parameter is left empty, the current `OutputDir` value will be invalidated.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// The notification configuration.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Resource ID. Ensure the corresponding resource is in the enabled state.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *ModifyScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	delete(f, "ScheduleName")
	delete(f, "Trigger")
	delete(f, "Activities")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "TaskNotifyConfig")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScheduleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyScheduleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyScheduleResponseParams `json:"Response"`
}

func (r *ModifyScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySmartSubtitleTemplateRequestParams struct {
	// Unique identifier of the smart subtitle template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Subtitle translation switch.
	// ON: enable translation
	// OFF: disable translation
	TranslateSwitch *string `json:"TranslateSwitch,omitnil,omitempty" name:"TranslateSwitch"`

	// Smart subtitle template name.
	// Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Smart subtitle template description.
	// Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Source language of the video with smart subtitles.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// zh-PY: Chinese-English-Cantonese
	// zh-medical: Medical Chinese
	// yue: Cantonese
	// vi: Vietnamese
	// ms: Malay
	// id: Indonesian
	// fil: Filipino
	// th: Thai
	// pt: Portuguese
	// tr: Turkish
	// ar: Arabic
	// es: Spanish
	// hi: Hindi
	// fr: French
	// de: German
	// zh-dialect: Chinese dialect
	VideoSrcLanguage *string `json:"VideoSrcLanguage,omitnil,omitempty" name:"VideoSrcLanguage"`

	// Smart subtitle file format.
	// vtt: WebVTT format
	// If this field is left blank, no subtitle file will be generated.
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`

	// Smart subtitle language type.
	// 0: source language1: target language
	// 2: source language + target language
	// The value can only be 0 when TranslateSwitch is set to OFF.The value can only be 1 or 2 when TranslateSwitch is set to ON.
	SubtitleType *int64 `json:"SubtitleType,omitnil,omitempty" name:"SubtitleType"`

	// ASR hotword lexicon parameter.
	AsrHotWordsConfigure *AsrHotWordsConfigure `json:"AsrHotWordsConfigure,omitnil,omitempty" name:"AsrHotWordsConfigure"`

	// Target language for subtitle translation.
	// This field takes effect when TranslateSwitch is set to ON.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// fr: French
	// es: Spanish
	// it: Italian
	// de: German
	// tr: Turkish
	// ru: Russian
	// pt: Portuguese
	// vi: Vietnamese
	// id: Indonesian
	// ms: Malay
	// th: Thai
	// ar: Arabic
	// hi: Hindi
	TranslateDstLanguage *string `json:"TranslateDstLanguage,omitnil,omitempty" name:"TranslateDstLanguage"`
}

type ModifySmartSubtitleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique identifier of the smart subtitle template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Subtitle translation switch.
	// ON: enable translation
	// OFF: disable translation
	TranslateSwitch *string `json:"TranslateSwitch,omitnil,omitempty" name:"TranslateSwitch"`

	// Smart subtitle template name.
	// Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Smart subtitle template description.
	// Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Source language of the video with smart subtitles.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// zh-PY: Chinese-English-Cantonese
	// zh-medical: Medical Chinese
	// yue: Cantonese
	// vi: Vietnamese
	// ms: Malay
	// id: Indonesian
	// fil: Filipino
	// th: Thai
	// pt: Portuguese
	// tr: Turkish
	// ar: Arabic
	// es: Spanish
	// hi: Hindi
	// fr: French
	// de: German
	// zh-dialect: Chinese dialect
	VideoSrcLanguage *string `json:"VideoSrcLanguage,omitnil,omitempty" name:"VideoSrcLanguage"`

	// Smart subtitle file format.
	// vtt: WebVTT format
	// If this field is left blank, no subtitle file will be generated.
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`

	// Smart subtitle language type.
	// 0: source language1: target language
	// 2: source language + target language
	// The value can only be 0 when TranslateSwitch is set to OFF.The value can only be 1 or 2 when TranslateSwitch is set to ON.
	SubtitleType *int64 `json:"SubtitleType,omitnil,omitempty" name:"SubtitleType"`

	// ASR hotword lexicon parameter.
	AsrHotWordsConfigure *AsrHotWordsConfigure `json:"AsrHotWordsConfigure,omitnil,omitempty" name:"AsrHotWordsConfigure"`

	// Target language for subtitle translation.
	// This field takes effect when TranslateSwitch is set to ON.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// fr: French
	// es: Spanish
	// it: Italian
	// de: German
	// tr: Turkish
	// ru: Russian
	// pt: Portuguese
	// vi: Vietnamese
	// id: Indonesian
	// ms: Malay
	// th: Thai
	// ar: Arabic
	// hi: Hindi
	TranslateDstLanguage *string `json:"TranslateDstLanguage,omitnil,omitempty" name:"TranslateDstLanguage"`
}

func (r *ModifySmartSubtitleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmartSubtitleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "TranslateSwitch")
	delete(f, "Name")
	delete(f, "Comment")
	delete(f, "VideoSrcLanguage")
	delete(f, "SubtitleFormat")
	delete(f, "SubtitleType")
	delete(f, "AsrHotWordsConfigure")
	delete(f, "TranslateDstLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySmartSubtitleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySmartSubtitleTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySmartSubtitleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifySmartSubtitleTemplateResponseParams `json:"Response"`
}

func (r *ModifySmartSubtitleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySmartSubtitleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotByTimeOffsetTemplateRequestParams struct {
	// Unique ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of a time point screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// The image format. Valid values: jpg, png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type ModifySnapshotByTimeOffsetTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a time point screencapturing template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Name of a time point screencapturing template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Image width in px. Value range: [128, 4,096].
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Image height in px. Value range: [128, 4,096].
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// The image format. Valid values: jpg, png, webp.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
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

// Predefined struct for user
type ModifySnapshotByTimeOffsetTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyTranscodeTemplateRequestParams struct {
	// Unique ID of a transcoding template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Container format. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Name of a transcoding template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Whether to remove video data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter.
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter.
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter. To enable it, please contact your Tencent Cloud sales rep.
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// Audio/Video enhancement settings.
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil,omitempty" name:"EnhanceConfig"`
}

type ModifyTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a transcoding template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Container format. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Name of a transcoding template. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Whether to remove video data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: Retain</li>
	// <li>1: Remove</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter.
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter.
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter. To enable it, please contact your Tencent Cloud sales rep.
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// Audio/Video enhancement settings.
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil,omitempty" name:"EnhanceConfig"`
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
	delete(f, "EnhanceConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTranscodeTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyWatermarkTemplateRequestParams struct {
	// Unique ID of a watermarking template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Watermarking template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Origin position. Valid values:
	// <li>TopLeft: The origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>TopRight: The origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>BottomLeft: The origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>BottomRight: The origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the watermark will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width;</li>
	// <li>If the string ends in px, the `XPos` of the watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.</li>
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the watermark will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height;</li>
	// <li>If the string ends in px, the `YPos` of the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.</li>
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// Image watermarking template. This field is valid only for image watermarking templates.
	ImageTemplate *ImageWatermarkInputForUpdate `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is valid only for text watermarking templates.
	TextTemplate *TextWatermarkTemplateInputForUpdate `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is required when `Type` is `svg` and is invalid when `Type` is `image` or `text`.
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`
}

type ModifyWatermarkTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a watermarking template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Watermarking template name. Length limit: 64 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description. Length limit: 256 characters.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Origin position. Valid values:
	// <li>TopLeft: The origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>TopRight: The origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>BottomLeft: The origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>BottomRight: The origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the watermark will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width;</li>
	// <li>If the string ends in px, the `XPos` of the watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.</li>
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the watermark will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height;</li>
	// <li>If the string ends in px, the `YPos` of the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.</li>
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// Image watermarking template. This field is valid only for image watermarking templates.
	ImageTemplate *ImageWatermarkInputForUpdate `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is valid only for text watermarking templates.
	TextTemplate *TextWatermarkTemplateInputForUpdate `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is required when `Type` is `svg` and is invalid when `Type` is `image` or `text`.
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`
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

// Predefined struct for user
type ModifyWatermarkTemplateResponseParams struct {
	// Image watermark address. This field is valid only when `ImageTemplate.ImageContent` is non-empty.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// 7. All: equivalent to 1+2+3+4
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Tag operation information.
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil,omitempty" name:"TagOperationInfo"`
}

type ModifyWordSampleRequest struct {
	*tchttp.BaseRequest
	
	// Keyword. Length limit: 128 characters.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <b>Keyword usage. Valid values:</b>
	// 1. Recognition.Ocr: OCR-based content recognition
	// 2. Recognition.Asr: ASR-based content recognition
	// 3. Review.Ocr: OCR-based inappropriate information recognition
	// 4. Review.Asr: ASR-based inappropriate information recognition
	// <b>Valid values can also be:</b>
	// 5. Recognition: ASR- and OCR-based content recognition; equivalent to 1+2
	// 6. Review: ASR- and OCR-based inappropriate information recognition; equivalent to 3+4
	// 7. All: equivalent to 1+2+3+4
	Usages []*string `json:"Usages,omitnil,omitempty" name:"Usages"`

	// Tag operation information.
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitnil,omitempty" name:"TagOperationInfo"`
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

// Predefined struct for user
type ModifyWordSampleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the blur relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the blur will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width;</li>
	// <li>If the string ends in px, the `XPos` of the blur will be the specified px; for example, `100px` means that `XPos` is 100 px.</li>
	// Default value: 0 px.
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// Vertical position of the origin of blur relative to the origin of coordinates of video. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the blur will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height;</li>
	// <li>If the string ends in px, the `YPos` of the blur will be the specified px; for example, `100px` means that `YPos` is 100 px.</li>
	// Default value: 0 px.
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// Blur width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the blur will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the blur will be in px; for example, `100px` means that `Width` is 100 px.</li>
	// Default value: 10%.
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// Blur height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the blur will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the blur will be in px; for example, `100px` means that `Height` is 100 px.</li>
	// Default value: 10%.
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// Start time offset of blur in seconds. If this parameter is left empty or 0 is entered, the blur will appear upon the first video frame.
	// <li>If this parameter is left empty or 0 is entered, the blur will appear upon the first video frame;</li>
	// <li>If this value is greater than 0 (e.g., n), the blur will appear at second n after the first video frame;</li>
	// <li>If this value is smaller than 0 (e.g., -n), the blur will appear at second n before the last video frame.</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of blur in seconds.
	// <li>If this parameter is left empty or 0 is entered, the blur will exist till the last video frame;</li>
	// <li>If this value is greater than 0 (e.g., n), the blur will exist till second n;</li>
	// <li>If this value is smaller than 0 (e.g., -n), the blur will exist till second n before the last video frame.</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type NumberFormat struct {
	// Start value of the `{number}` variable. Default value: 0.
	InitialValue *uint64 `json:"InitialValue,omitnil,omitempty" name:"InitialValue"`

	// Increment of the `{number}` variable. Default value: 1.
	Increment *uint64 `json:"Increment,omitnil,omitempty" name:"Increment"`

	// Minimum length of the `{number}` variable. A placeholder will be used if the variable length is below the minimum requirement. Default value: 1.
	MinLength *uint64 `json:"MinLength,omitnil,omitempty" name:"MinLength"`

	// Placeholder used when the `{number}` variable length is below the minimum requirement. Default value: 0.
	PlaceHolder *string `json:"PlaceHolder,omitnil,omitempty" name:"PlaceHolder"`
}

type OcrFullTextConfigureInfo struct {
	// Switch of a full text recognition task. Valid values:
	// <li>ON: Enables an intelligent full text recognition task;</li>
	// <li>OFF: Disables an intelligent full text recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type OcrFullTextConfigureInfoForUpdate struct {
	// Switch of a full text recognition task. Valid values:
	// <li>ON: Enables an intelligent full text recognition task;</li>
	// <li>OFF: Disables an intelligent full text recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type OcrWordsConfigureInfo struct {
	// Switch of a text keyword recognition task. Valid values:
	// <li>ON: Enables a text keyword recognition task;</li>
	// <li>OFF: Disables a text keyword recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`
}

type OcrWordsConfigureInfoForUpdate struct {
	// Switch of a text keyword recognition task. Valid values:
	// <li>ON: Enables a text keyword recognition task;</li>
	// <li>OFF: Disables a text keyword recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Keyword filter tag, which specifies the keyword tag that needs to be returned. If this parameter is left empty, all results will be returned.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`
}

type OverrideTranscodeParameter struct {
	// Container format. Valid values: mp4, flv, hls, mp3, flac, ogg, and m4a; mp3, flac, ogg, and m4a are formats of audio files.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Whether to remove video data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	RemoveVideo *uint64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: retain</li>
	// <li>1: remove</li>
	RemoveAudio *uint64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter.
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter.
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// The TSC transcoding parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TEHDConfig *TEHDConfigForUpdate `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// The subtitle settings.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubtitleTemplate *SubtitleTemplate `json:"SubtitleTemplate,omitnil,omitempty" name:"SubtitleTemplate"`

	// The information of the external audio track to add.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddonAudioStream []*MediaInputInfo `json:"AddonAudioStream,omitnil,omitempty" name:"AddonAudioStream"`

	// An extended field for transcoding.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StdExtInfo *string `json:"StdExtInfo,omitnil,omitempty" name:"StdExtInfo"`

	// The subtitle file to add.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddOnSubtitles []*AddOnSubtitle `json:"AddOnSubtitles,omitnil,omitempty" name:"AddOnSubtitles"`
}

// Predefined struct for user
type ParseLiveStreamProcessNotificationRequestParams struct {
	// Live stream event notification obtained from CMQ.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ParseLiveStreamProcessNotificationRequest struct {
	*tchttp.BaseRequest
	
	// Live stream event notification obtained from CMQ.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
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

// Predefined struct for user
type ParseLiveStreamProcessNotificationResponseParams struct {
	// Live stream processing result type, including:
	// <li>AiReviewResult: content auditing result.</li>
	// <li>AiRecognitionResult: content recognition result.</li>
	// <li>LiveRecordResult: live recording result.</li>
	// <li>AiQualityControlResult: media quality inspection result.</li>
	// <li>ProcessEof: live stream processing result.</li>
	NotificationType *string `json:"NotificationType,omitnil,omitempty" name:"NotificationType"`

	// Video processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Information of a live stream processing error, which is valid when `NotificationType` is `ProcessEof`.
	// Note: when this field return null, means no valid values can be obtained.
	ProcessEofInfo *LiveStreamProcessErrorInfo `json:"ProcessEofInfo,omitnil,omitempty" name:"ProcessEofInfo"`

	// Content audit result, which is valid when `NotificationType` is `AiReviewResult`.
	// Note: when this field return null, means no valid values can be obtained.
	AiReviewResultInfo *LiveStreamAiReviewResultInfo `json:"AiReviewResultInfo,omitnil,omitempty" name:"AiReviewResultInfo"`

	// Content recognition result, which is valid if `NotificationType` is `AiRecognitionResult`.
	AiRecognitionResultInfo *LiveStreamAiRecognitionResultInfo `json:"AiRecognitionResultInfo,omitnil,omitempty" name:"AiRecognitionResultInfo"`

	// Content analysis result, which is valid if `NotificationType` is `AiAnalysisResult`.
	AiAnalysisResultInfo *LiveStreamAiAnalysisResultInfo `json:"AiAnalysisResultInfo,omitnil,omitempty" name:"AiAnalysisResultInfo"`

	// Media quality inspection result, which is valid if `NotificationType` is `AiQualityControlResult`.
	AiQualityControlResultInfo *LiveStreamAiQualityControlResultInfo `json:"AiQualityControlResultInfo,omitnil,omitempty" name:"AiQualityControlResultInfo"`

	// Live recording result is valid when NotificationType is LiveRecordResult.
	// Note: when this field return null, means no valid values can be obtained.
	LiveRecordResultInfo *LiveStreamRecordResultInfo `json:"LiveRecordResultInfo,omitnil,omitempty" name:"LiveRecordResultInfo"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// - Expiration time, event notification signature expiration UNIX timestamp. - By default, notifications sent by MPS expire after 10 minutes. If the expiration time specified has elapsed, a notification will be considered invalid. This can prevent replay attacks. - The format of Timestamp is a decimal UNIX timestamp, which is the number of seconds that have elapsed since January 1, 1970 (midnight UTC/GMT).
	// Note: This field may return null, indicating that no valid value can be obtained.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Event notification security signature. Sign = MD5 (Timestamp + NotifyKey). Note: Media Processing Service concatenates Timestamp and NotifyKey from TaskNotifyConfig as a string and calculates the Sign value through MD5. This value is included in the notification message. Your backend server can verify whether the Sign is correct using the same algorithm, to confirm whether the message is indeed from the Media Processing Service backend.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	Sign *string `json:"Sign,omitnil,omitempty" name:"Sign"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ParseLiveStreamProcessNotificationResponse struct {
	*tchttp.BaseResponse
	Response *ParseLiveStreamProcessNotificationResponseParams `json:"Response"`
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

// Predefined struct for user
type ParseNotificationRequestParams struct {
	// Event notification obtained from CMQ.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ParseNotificationRequest struct {
	*tchttp.BaseRequest
	
	// Event notification obtained from CMQ.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
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

// Predefined struct for user
type ParseNotificationResponseParams struct {
	// The event type. Valid values:
	// <li>WorkflowTask</li>
	// <li>EditMediaTask</li>
	// <li>ScheduleTask (scheme)</li>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// The information of a video processing task. Information will be returned only if `EventType` is `WorkflowTask`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WorkflowTaskEvent *WorkflowTask `json:"WorkflowTaskEvent,omitnil,omitempty" name:"WorkflowTaskEvent"`

	// The information of a video editing task. Information will be returned only if `EventType` is `EditMediaTask`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EditMediaTaskEvent *EditMediaTask `json:"EditMediaTaskEvent,omitnil,omitempty" name:"EditMediaTaskEvent"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// The information of a scheme. Information will be returned only if `EventType` is `ScheduleTask`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScheduleTaskEvent *ScheduleTask `json:"ScheduleTaskEvent,omitnil,omitempty" name:"ScheduleTaskEvent"`

	// - The expiration time (Unix timestamp) of the notification's signature.
	// - By default, notifications sent by MPS expire after 10 minutes. If the expiration time specified has elapsed, a notification will be considered invalid. This can prevent replay attacks.
	// - The format of this parameter is a decimal Unix timestamp, i.e., the number of seconds that have elapsed since 00:00 (UTC/GMT time) on January 1, 1970.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Event notification security signature. Sign = MD5 (Timestamp + NotifyKey). Note: Media Processing Service concatenates Timestamp and NotifyKey from TaskNotifyConfig as a string and calculates the Sign value through MD5. This value is included in the notification message. Your backend server can verify whether the Sign is correct using the same algorithm, to confirm whether the message is indeed from the Media Processing Service backend.
	Sign *string `json:"Sign,omitnil,omitempty" name:"Sign"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ParseNotificationResponse struct {
	*tchttp.BaseResponse
	Response *ParseNotificationResponseParams `json:"Response"`
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
	// Whether to detect sensitive information based on ASR. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PoliticalAsrReviewTemplateInfoForUpdate struct {
	// Whether to detect sensitive information based on ASR. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PoliticalConfigureInfo struct {
	// The parameters for detecting sensitive information in images.
	ImgReviewInfo *PoliticalImgReviewTemplateInfo `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// The parameters for detecting sensitive information based on ASR.
	AsrReviewInfo *PoliticalAsrReviewTemplateInfo `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// The parameters for detecting sensitive information based on OCR.
	OcrReviewInfo *PoliticalOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type PoliticalConfigureInfoForUpdate struct {
	// The parameters for detecting sensitive information in images.
	ImgReviewInfo *PoliticalImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// The parameters for detecting sensitive information based on ASR.
	AsrReviewInfo *PoliticalAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// The parameters for detecting sensitive information based on OCR.
	OcrReviewInfo *PoliticalOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type PoliticalImgReviewTemplateInfo struct {
	// Whether to detect sensitive information in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The filter labels for sensitive information detection in images, which specify the types of sensitive information to return. If this parameter is left empty, the detection results for all labels are returned. Valid values:
	// <li>violation_photo (banned icons)</li>
	// <li>politician</li>
	// <li>entertainment (people in the entertainment industry)</li>
	// <li>sport (people in the sports industry)</li>
	// <li>entrepreneur</li>
	// <li>scholar</li>
	// <li>celebrity</li>
	// <li>military (people in military)</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 97 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 95 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PoliticalImgReviewTemplateInfoForUpdate struct {
	// Whether to detect sensitive information in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The filter labels for sensitive information detection in images, which specify the types of sensitive information to return. If this parameter is left empty, the detection results for all labels are returned. Valid values:
	// <li>violation_photo (banned icons)</li>
	// <li>politician</li>
	// <li>entertainment (people in the entertainment industry)</li>
	// <li>sport (people in the sports industry)</li>
	// <li>entrepreneur</li>
	// <li>scholar</li>
	// <li>celebrity</li>
	// <li>military (people in military)</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfo struct {
	// Whether to detect sensitive information based on OCR. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfoForUpdate struct {
	// Whether to detect sensitive information based on OCR. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfo struct {
	// Switch of a porn information detection in speech task. Valid values:
	// <li>ON: Enables a porn information detection in speech task;</li>
	// <li>OFF: Disables a porn information detection in speech task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfoForUpdate struct {
	// Switch of a porn information detection in speech task. Valid values:
	// <li>ON: Enables a porn information detection in speech task;</li>
	// <li>OFF: Disables a porn information detection in speech task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornConfigureInfo struct {
	// Control parameter of porn information detection in image.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImgReviewInfo *PornImgReviewTemplateInfo `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// Control parameter of porn information detection in speech.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrReviewInfo *PornAsrReviewTemplateInfo `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// Control parameter of porn information detection in text.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *PornOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type PornConfigureInfoForUpdate struct {
	// Control parameter of porn information detection in image.
	ImgReviewInfo *PornImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// Control parameter of porn information detection in speech.
	AsrReviewInfo *PornAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// Control parameter of porn information detection in text.
	OcrReviewInfo *PornOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type PornImgReviewTemplateInfo struct {
	// Switch of a porn information detection in image task. Valid values:
	// <li>ON: Enables a porn information detection in image task;</li>
	// <li>OFF: Disables a porn information detection in image task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Filter tag for porn information detection in image. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. Valid values:
	// <li>porn: Porn;</li>
	// <li>vulgar: Vulgarity;</li>
	// <li>intimacy: Intimacy;</li>
	// <li>sexy: Sexiness.</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 90 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 0 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornImgReviewTemplateInfoForUpdate struct {
	// Switch of a porn information detection in image task. Valid values:
	// <li>ON: Enables a porn information detection in image task;</li>
	// <li>OFF: Disables a porn information detection in image task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Filter tag for porn information detection in image. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. Valid values:
	// <li>porn: Porn;</li>
	// <li>vulgar: Vulgarity;</li>
	// <li>intimacy: Intimacy;</li>
	// <li>sexy: Sexiness.</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfo struct {
	// Switch of a porn information detection in text task. Valid values:
	// <li>ON: Enables a porn information detection in text task;</li>
	// <li>OFF: Disables a porn information detection in text task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfoForUpdate struct {
	// Switch of a porn information detection in text task. Valid values:
	// <li>ON: Enables a porn information detection in text task;</li>
	// <li>OFF: Disables a porn information detection in text task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

// Predefined struct for user
type ProcessImageRequestParams struct {
	// File input information for image processing.
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// Target storage for image processing output files. If left blank, it inherits the storage location in InputInfo.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Output file path for image processing. If left blank, it is the directory of the file in InputInfo. If it is a directory, such as `/image/201907/`, it means inheriting the original filename and outputting to this directory.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Image processing parameter.
	ImageTask *ImageTaskInput `json:"ImageTask,omitnil,omitempty" name:"ImageTask"`
}

type ProcessImageRequest struct {
	*tchttp.BaseRequest
	
	// File input information for image processing.
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// Target storage for image processing output files. If left blank, it inherits the storage location in InputInfo.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Output file path for image processing. If left blank, it is the directory of the file in InputInfo. If it is a directory, such as `/image/201907/`, it means inheriting the original filename and outputting to this directory.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Image processing parameter.
	ImageTask *ImageTaskInput `json:"ImageTask,omitnil,omitempty" name:"ImageTask"`
}

func (r *ProcessImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "ImageTask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessImageResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProcessImageResponse struct {
	*tchttp.BaseResponse
	Response *ProcessImageResponseParams `json:"Response"`
}

func (r *ProcessImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessLiveStreamRequestParams struct {
	// Live stream URL, which must be a live stream file address. RTMP, HLS, and FLV are supported.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Event notification information of a task, which is used to specify the live stream processing result.
	TaskNotifyConfig *LiveStreamTaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Target bucket of a live stream processing output file. This parameter is required if a file will be output.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Target directory of a live stream processing output file, such as `/movie/201909/`. If this parameter is left empty, the `/` directory will be used.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// Type parameter of video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`


	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// Media quality inspection type task parameters.
	AiQualityControlTask *AiQualityControlTaskInput `json:"AiQualityControlTask,omitnil,omitempty" name:"AiQualityControlTask"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// The live scheme ID.
	// Note 1:
	// <li>If an output storage (`OutputStorage`) and directory (`OutputDir`) are specified for a subtask of the scheme, those output settings will be applied. </li>
	// u200c<li>If an output storage (`OutputStorage`) and directory (`OutputDir`) are not specified for a subtask of the scheme, the output parameters specified for `ProcessLiveStream` (if any) will be applied. </li>
	// Note 2: If `TaskNotifyConfig` is specified when `ProcessLiveStream` is called, the specified settings will be applied instead of the default callback settings of the scheme.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`
}

type ProcessLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// Live stream URL, which must be a live stream file address. RTMP, HLS, and FLV are supported.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Event notification information of a task, which is used to specify the live stream processing result.
	TaskNotifyConfig *LiveStreamTaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Target bucket of a live stream processing output file. This parameter is required if a file will be output.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Target directory of a live stream processing output file, such as `/movie/201909/`. If this parameter is left empty, the `/` directory will be used.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// Type parameter of video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// Media quality inspection type task parameters.
	AiQualityControlTask *AiQualityControlTaskInput `json:"AiQualityControlTask,omitnil,omitempty" name:"AiQualityControlTask"`

	// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// The live scheme ID.
	// Note 1:
	// <li>If an output storage (`OutputStorage`) and directory (`OutputDir`) are specified for a subtask of the scheme, those output settings will be applied. </li>
	// u200c<li>If an output storage (`OutputStorage`) and directory (`OutputDir`) are not specified for a subtask of the scheme, the output parameters specified for `ProcessLiveStream` (if any) will be applied. </li>
	// Note 2: If `TaskNotifyConfig` is specified when `ProcessLiveStream` is called, the specified settings will be applied instead of the default callback settings of the scheme.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`
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
	delete(f, "AiAnalysisTask")
	delete(f, "AiQualityControlTask")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "ScheduleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessLiveStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessLiveStreamResponseParams struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProcessLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *ProcessLiveStreamResponseParams `json:"Response"`
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

// Predefined struct for user
type ProcessMediaRequestParams struct {
	// The information of the file to process.
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// Target storage for Media Processing Service output files. If left blank, it inherits the storage location in InputInfo.
	// 
	// Note: When InputInfo.Type is URL, this parameter is required.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The directory to save the media processing output file, which must start and end with `/`, such as `/movie/201907/`.
	// If you do not specify this parameter, the file will be saved to the directory specified in `InputInfo`.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Orchestration ID.
	// Note 1: For parameters OutputStorage and OutputDir:
	// <li>When a sub-task node in service orchestration has OutputStorage and OutputDir configured, the output configured in this sub-task node is used as the output of the sub-task.</li>
	// <li>When a sub-task node in service orchestration does not have OutputStorage and OutputDir configured, if the task creation API (ProcessMedia) has specified an output, it will override the default output of the original orchestration.</li>
	// <li>The priority of output settings is: Orchestration sub-task node > Output specified by the task API > Corresponding configuration within an orchestration.</li>
	// Note 2: For the TaskNotifyConfig parameter, if the task creation API (ProcessMedia) has set this parameter, it will override the default callback of the original orchestration.
	// 
	// Note 3: The trigger configured for an orchestration is for automatically starting the orchestration. It stops working when you manually call this API to start an orchestration.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// The media processing parameters to use.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// Media quality inspection type task parameters.
	AiQualityControlTask *AiQualityControlTaskInput `json:"AiQualityControlTask,omitnil,omitempty" name:"AiQualityControlTask"`

	// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Task flow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// Identification code for deduplication, up to 50 characters. If a request with the same identification code was made within the past 3 days, an error will be returned for the current request. If this parameter is not provided or is an empty string, deduplication will not be performed for this request.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// The task type.
	// <li> `Online` (default): A task that is executed immediately.</li>
	// <li> `Offline`: A task that is executed when the system is idle (within three days by default).</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Resource ID. Ensure the corresponding resource is in the enabled state. The default value is an account's primary resource ID.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Smart subtitle task.
	SmartSubtitlesTask *SmartSubtitlesTaskInput `json:"SmartSubtitlesTask,omitnil,omitempty" name:"SmartSubtitlesTask"`

	// Whether to skip metadata acquisition. Valid values:
	// 0: do not skip
	// 1: skip
	// Default value: 0		
	SkipMateData *int64 `json:"SkipMateData,omitnil,omitempty" name:"SkipMateData"`
}

type ProcessMediaRequest struct {
	*tchttp.BaseRequest
	
	// The information of the file to process.
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// Target storage for Media Processing Service output files. If left blank, it inherits the storage location in InputInfo.
	// 
	// Note: When InputInfo.Type is URL, this parameter is required.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The directory to save the media processing output file, which must start and end with `/`, such as `/movie/201907/`.
	// If you do not specify this parameter, the file will be saved to the directory specified in `InputInfo`.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Orchestration ID.
	// Note 1: For parameters OutputStorage and OutputDir:
	// <li>When a sub-task node in service orchestration has OutputStorage and OutputDir configured, the output configured in this sub-task node is used as the output of the sub-task.</li>
	// <li>When a sub-task node in service orchestration does not have OutputStorage and OutputDir configured, if the task creation API (ProcessMedia) has specified an output, it will override the default output of the original orchestration.</li>
	// <li>The priority of output settings is: Orchestration sub-task node > Output specified by the task API > Corresponding configuration within an orchestration.</li>
	// Note 2: For the TaskNotifyConfig parameter, if the task creation API (ProcessMedia) has set this parameter, it will override the default callback of the original orchestration.
	// 
	// Note 3: The trigger configured for an orchestration is for automatically starting the orchestration. It stops working when you manually call this API to start an orchestration.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// The media processing parameters to use.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// Media quality inspection type task parameters.
	AiQualityControlTask *AiQualityControlTaskInput `json:"AiQualityControlTask,omitnil,omitempty" name:"AiQualityControlTask"`

	// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Task flow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// Identification code for deduplication, up to 50 characters. If a request with the same identification code was made within the past 3 days, an error will be returned for the current request. If this parameter is not provided or is an empty string, deduplication will not be performed for this request.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// The task type.
	// <li> `Online` (default): A task that is executed immediately.</li>
	// <li> `Offline`: A task that is executed when the system is idle (within three days by default).</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Resource ID. Ensure the corresponding resource is in the enabled state. The default value is an account's primary resource ID.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Smart subtitle task.
	SmartSubtitlesTask *SmartSubtitlesTaskInput `json:"SmartSubtitlesTask,omitnil,omitempty" name:"SmartSubtitlesTask"`

	// Whether to skip metadata acquisition. Valid values:
	// 0: do not skip
	// 1: skip
	// Default value: 0		
	SkipMateData *int64 `json:"SkipMateData,omitnil,omitempty" name:"SkipMateData"`
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
	delete(f, "ScheduleId")
	delete(f, "MediaProcessTask")
	delete(f, "AiContentReviewTask")
	delete(f, "AiAnalysisTask")
	delete(f, "AiRecognitionTask")
	delete(f, "AiQualityControlTask")
	delete(f, "TaskNotifyConfig")
	delete(f, "TasksPriority")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	delete(f, "TaskType")
	delete(f, "ResourceId")
	delete(f, "SmartSubtitlesTask")
	delete(f, "SkipMateData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type ProhibitedAsrReviewTemplateInfoForUpdate struct {
	// Switch of prohibited information detection in speech task. Valid values:
	// <li>ON: enables prohibited information detection in speech task;</li>
	// <li>OFF: disables prohibited information detection in speech task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type ProhibitedConfigureInfo struct {
	// Control parameter of prohibited information detection in speech.
	AsrReviewInfo *ProhibitedAsrReviewTemplateInfo `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// Control parameter of prohibited information detection in text.
	OcrReviewInfo *ProhibitedOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type ProhibitedConfigureInfoForUpdate struct {
	// Control parameter of prohibited information detection in speech.
	AsrReviewInfo *ProhibitedAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// Control parameter of prohibited information detection in text.
	OcrReviewInfo *ProhibitedOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type ProhibitedOcrReviewTemplateInfo struct {
	// Switch of prohibited information detection in text task. Valid values:
	// <li>ON: enables prohibited information detection in text task;</li>
	// <li>OFF: disables prohibited information detection in text task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type ProhibitedOcrReviewTemplateInfoForUpdate struct {
	// Switch of prohibited information detection in text task. Valid values:
	// <li>ON: enables prohibited information detection in text task;</li>
	// <li>OFF: disables prohibited information detection in text task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type QualityControlData struct {
	// Whether there is an audio track. `true` indicates that there isn't.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NoAudio *bool `json:"NoAudio,omitnil,omitempty" name:"NoAudio"`

	// Whether there is a video track. `true` indicates that there isn't.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NoVideo *bool `json:"NoVideo,omitnil,omitempty" name:"NoVideo"`

	// No-reference quality score of the video (100 points in total).
	// Note: This field may return null, indicating that no valid value can be obtained.
	QualityEvaluationScore *int64 `json:"QualityEvaluationScore,omitnil,omitempty" name:"QualityEvaluationScore"`

	// No-reference quality score of the video (MOS).
	// Note: This field may return null, indicating that no valid value can be obtained.
	QualityEvaluationMeanOpinionScore *float64 `json:"QualityEvaluationMeanOpinionScore,omitnil,omitempty" name:"QualityEvaluationMeanOpinionScore"`

	// Exception items detected in content quality inspection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	QualityControlResultSet []*QualityControlResult `json:"QualityControlResultSet,omitnil,omitempty" name:"QualityControlResultSet"`

	// Exception items detected in format diagnosis.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ContainerDiagnoseResultSet []*ContainerDiagnoseResultItem `json:"ContainerDiagnoseResultSet,omitnil,omitempty" name:"ContainerDiagnoseResultSet"`
}

type QualityControlItem struct {
	// The confidence score. Value range: 0-100.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// The start timestamp (second) of the segment.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// The end timestamp (second) of the segment.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// The coordinates (px) of the top left and bottom right corner.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitnil,omitempty" name:"AreaCoordSet"`
}

type QualityControlItemConfig struct {
	// Quality control item name. The quality control item values are as follows:
	// <li>LowEvaluation: No reference score.</li>
	// <li>Mosaic: Mosaic detection.</li>
	// <li>CrashScreen: Screen crash detection.</li>
	// <li>Blur: Blur detection.</li>
	// <li>BlackWhiteEdge: Black and white edge detection.</li>
	// <li>SolidColorScreen: Solid color screen detection.</li>
	// <li>LowLighting: Low lighting.</li>
	// <li>HighLighting: Overexposure.</li>
	// <li>NoVoice: Silence detection.</li>
	// <li>LowVoice: Low voice detection.</li>
	// <li>HighVoice: High voice detection.</li>
	// <li>Jitter: Jitter detection.</li>
	// <li>Noise: Noise detection.</li>
	// <li>QRCode: QR code detection.</li>
	// <li>BarCode: Barcode detection.</li>
	// <li>AppletCode: Applet code detection.</li>
	// <li>VideoResolutionChanged: The video resolution changed.</li>
	// <li>AudioSampleRateChanged: The audio sampling rate changed.</li>
	// <li>AudioChannelsChanged: The audio channel count changed.</li>
	// <li>ParameterSetsChanged: The stream parameter set information changed.</li>
	// <li>DarOrSarInvalid: Abnormal video aspect ratio.</li>
	// <li>TimestampFallback: DTS timestamp fallback.</li>
	// <li>DtsJitter: Excessive DTS jitter.</li>
	// <li>PtsJitter: Excessive PTS jitter.</li>
	// <li>AACDurationDeviation: Unreasonable AAC frame timestamp interval.</li>
	// <li>AudioDroppingFrames: Audio frame loss.</li>
	// <li>VideoDroppingFrames: Video frame loss.</li>
	// <li>AVTimestampInterleave: Unreasonable audio and video interleaving.</li>
	// <li>PtsLessThanDts: The PTS of media streams is less than DTS.</li>
	// <li>ReceiveFpsJitter: Excessive jitter of the frame rate received by the network.</li>
	// <li>ReceiveFpsTooSmall: Too low video frame rate received by the network.</li>
	// <li>FpsJitter: Excessive stream frame rate jitter calculated through PTS.</li>
	// <li>StreamOpenFailed: Stream opening failed.</li>
	// <li>StreamEnd: The stream ended.</li>
	// <li>StreamParseFailed: Stream parsing failed.</li>
	// <li>VideoFirstFrameNotIdr: The first frame is not an IDR frame.</li>
	// <li>StreamNALUError: NALU start code error.</li>
	// <li>TsStreamNoAud: The H26x stream of MPEGTS lacks AUD NALU.</li>
	// <li>AudioStreamLack: No audio stream.</li>
	// <li>VideoStreamLack: No video stream.</li>
	// <li>LackAudioRecover: Lack of audio stream recovery.</li>
	// <li>LackVideoRecover: Lack of video stream recovery.</li>
	// <li>VideoBitrateOutofRange: Out-of-range video stream bitrate (kbps).</li>
	// <li>AudioBitrateOutofRange: Out-of-range audio stream bitrate (kbps).</li>
	// <li>VideoDecodeFailed: Video decoding error.</li>
	// <li>AudioDecodeFailed: Audio decoding error.</li>
	// <li>AudioOutOfPhase: Opposite phase in Dual-channel audio.</li>
	// <li>VideoDuplicatedFrame: Duplicate frames in the video stream.</li>
	// <li>AudioDuplicatedFrame: Duplicate frames in the audio stream.</li>
	// <li>VideoRotation: Video image rotation.</li>
	// <li>TsMultiPrograms: The MPEG2-TS stream has multiple programs.</li>
	// <li>Mp4InvalidCodecFourcc: The codec fourcc in MP4 does not meet Apple HLS requirements.</li>
	// <li>HLSBadM3u8Format: Invalid m3u8 file.</li>
	// <li>HLSInvalidMasterM3u8: Invalid main m3u8 file.</li>
	// <li>HLSInvalidMediaM3u8: Invalid media m3u8 file.</li>
	// <li>HLSMasterM3u8Recommended: The main m3u8 file lacks parameters recommended by the standard.</li>
	// <li>HLSMediaM3u8Recommended: The media m3u8 file lacks parameters recommended by the standard.</li>
	// <li>HLSMediaM3u8DiscontinuityExist: EXT-X-DISCONTINUITY exists in the media m3u8 file.</li>
	// <li>HLSMediaSegmentsStreamNumChange: The number of streams in the segment has changed.</li>
	// <li>HLSMediaSegmentsPTSJitterDeviation: PTS jitter between segments without EXT-X-DISCONTINUITY.</li>
	// <li>HLSMediaSegmentsDTSJitterDeviation: DTS jitter between segments without EXT-X-DISCONTINUITY.</li>
	// <li>TimecodeTrackExist: TMCD track in MP4.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Capability configuration switch. Valid values:
	// <li>ON: enabled;</li>
	// <li>OFF: disabled.</li>
	// 
	// Default value: ON.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Sampling method, Valid value:
	// - Time: sampling based on time interval.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sampling *string `json:"Sampling,omitnil,omitempty" name:"Sampling"`

	// Sampling interval time, in ms.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntervalTime *uint64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// Duration of abnormality, in ms.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Threshold of a detection item. Different detection items have different thresholds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Threshold *string `json:"Threshold,omitnil,omitempty" name:"Threshold"`
}

type QualityControlResult struct {
	// The issue type. Valid values:
	// `Jitter`
	// `Blur`
	// `LowLighting`
	// `HighLighting` (overexposure)
	// `CrashScreen` (video corruption)
	// `BlackWhiteEdge`
	// `SolidColorScreen` (blank screen)
	// `Noise`
	// `Mosaic` (pixelation)
	// `QRCode`
	// `AppletCode` (Weixin Mini Program code)
	// `BarCode`
	// `LowVoice`
	// `HighVoice`
	// `NoVoice`
	// `LowEvaluation` (low no-reference video quality score)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The information of a checked segment in quality control.
	QualityControlItems []*QualityControlItem `json:"QualityControlItems,omitnil,omitempty" name:"QualityControlItems"`
}

type QualityControlTemplate struct {
	// Unique identifier of a media quality inspection template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Media quality inspection template name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description.
	// 
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Template type. Valid values:
	// <li>Preset: system preset template;</li>
	// <li>Custom: custom template.</li>
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Media quality inspection configuration parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	QualityControlItemSet []*QualityControlItemConfig `json:"QualityControlItemSet,omitnil,omitempty" name:"QualityControlItemSet"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type RawImageWatermarkInput struct {
	// Input content of watermark image. JPEG and PNG images are supported.
	ImageContent *MediaInputInfo `json:"ImageContent,omitnil,omitempty" name:"ImageContent"`

	// Watermark width. % and px formats are supported:
	// <li>If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width;</li>
	// <li>If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.</li>
	// Default value: 10%.
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// Watermark height. % and px formats are supported:
	// <li>If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px.</li>
	// Default value: 0 px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// Repeat type of an animated watermark. Valid values:
	// <li>`once`: no longer appears after watermark playback ends.</li>
	// <li>`repeat_last_frame`: stays on the last frame after watermark playback ends.</li>
	// <li>`repeat` (default): repeats the playback until the video ends.</li>
	RepeatType *string `json:"RepeatType,omitnil,omitempty" name:"RepeatType"`
}

type RawSmartSubtitleParameter struct {
	// Smart subtitle language type.
	// 0: source language
	// 1: target language
	// 2: source language + target language
	// The value can only be 0 when TranslateSwitch is set to OFF. The value can only be 1 or 2 when TranslateSwitch is set to ON.
	SubtitleType *int64 `json:"SubtitleType,omitnil,omitempty" name:"SubtitleType"`

	// Source language of the video with smart subtitles.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// zh-PY: Chinese-English-Cantonese
	// zh-medical: Medical Chinese
	// yue: Cantonese
	// vi: Vietnamese
	// ms: Malay
	// id: Indonesian
	// fil: Filipino
	// th: Thai
	// pt: Portuguese
	// tr: Turkish
	// ar: Arabic
	// es: Spanish
	// hi: Hindi
	// fr: French
	// de: German
	// zh-dialect: Chinese dialect
	VideoSrcLanguage *string `json:"VideoSrcLanguage,omitnil,omitempty" name:"VideoSrcLanguage"`

	// Smart subtitle file format.
	// vtt: WebVTT format
	// If this field is left blank, no subtitle file will be generated.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`

	// Subtitle translation switch.
	// ON: enable translation
	// OFF: disable translation
	// Note: This field may return null, indicating that no valid value can be obtained.
	TranslateSwitch *string `json:"TranslateSwitch,omitnil,omitempty" name:"TranslateSwitch"`

	// Target language for subtitle translation.
	// This field takes effect when TranslateSwitch is set to ON.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// fr: French
	// es: Spanish
	// it: Italian
	// de: German
	// tr: Turkish
	// ru: Russian
	// pt: Portuguese
	// vi: Vietnamese
	// id: Indonesian
	// ms: Malay
	// th: Thai
	// ar: Arabic
	// hi: Hindi
	// Note: This field may return null, indicating that no valid value can be obtained.
	TranslateDstLanguage *string `json:"TranslateDstLanguage,omitnil,omitempty" name:"TranslateDstLanguage"`

	// ASR hotword lexicon parameter.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AsrHotWordsConfigure *AsrHotWordsConfigure `json:"AsrHotWordsConfigure,omitnil,omitempty" name:"AsrHotWordsConfigure"`

	// Custom parameter.
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`
}

type RawTranscodeParameter struct {
	// Container. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Whether to remove video data. Valid values:
	// <li>0: retain;</li>
	// <li>1: remove.</li>
	// Default value: 0.
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: retain;</li>
	// <li>1: remove.</li>
	// Default value: 0.
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter. This field is required when `RemoveVideo` is 0.
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter. This field is required when `RemoveAudio` is 0.
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// Additional parameter, which is a serialized JSON string.
	// Note: This field may return null, indicating that no valid value can be obtained.
	StdExtInfo *string `json:"StdExtInfo,omitnil,omitempty" name:"StdExtInfo"`

	// Audio/Video enhancement configuration.
	// Note: This field may return null, indicating that no valid value can be obtained.
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil,omitempty" name:"EnhanceConfig"`
}

type RawWatermarkParameter struct {
	// Watermark type. Valid values:
	// <li>image: image watermark.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Origin position, which currently can only be:
	// <li>TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text.</li>
	// Default value: TopLeft.
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `XPos` of the watermark will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width;</li>
	// <li>If the string ends in px, the `XPos` of the watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.</li>
	// Default value: 0 px.
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported:
	// <li>If the string ends in %, the `YPos` of the watermark will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height;</li>
	// <li>If the string ends in px, the `YPos` of the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.</li>
	// Default value: 0 px.
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// Image watermark template. This field is required when `Type` is `image` and is invalid when `Type` is `text`.
	ImageTemplate *RawImageWatermarkInput `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`
}

// Predefined struct for user
type ResetWorkflowRequestParams struct {
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name of up to 128 characters, which must be unique for the same user.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Triggering rule bound to a workflow. If an uploaded video hits the rule for the object, the workflow will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// Output configuration of a video processing output file. If this parameter is left empty, the storage location in `Trigger` will be inherited.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The target directory for the output files generated by video processing. It must start and end with a slash (/), such as `/movie/201907/`.
	// If left empty, it is the same as the directory of the trigger file, that is, `{inputDir}`.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Parameter of a video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// Workflow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TaskPriority *int64 `json:"TaskPriority,omitnil,omitempty" name:"TaskPriority"`

	// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`
}

type ResetWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name of up to 128 characters, which must be unique for the same user.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Triggering rule bound to a workflow. If an uploaded video hits the rule for the object, the workflow will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// Output configuration of a video processing output file. If this parameter is left empty, the storage location in `Trigger` will be inherited.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The target directory for the output files generated by video processing. It must start and end with a slash (/), such as `/movie/201907/`.
	// If left empty, it is the same as the directory of the trigger file, that is, `{inputDir}`.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Parameter of a video processing task.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// Workflow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TaskPriority *int64 `json:"TaskPriority,omitnil,omitempty" name:"TaskPriority"`

	// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`
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

// Predefined struct for user
type ResetWorkflowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *ResetWorkflowResponseParams `json:"Response"`
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

type S3InputInfo struct {
	// The AWS S3 bucket.
	S3Bucket *string `json:"S3Bucket,omitnil,omitempty" name:"S3Bucket"`

	// The region of the AWS S3 bucket.
	S3Region *string `json:"S3Region,omitnil,omitempty" name:"S3Region"`

	// The path of the AWS S3 object.
	S3Object *string `json:"S3Object,omitnil,omitempty" name:"S3Object"`

	// The key ID required to access the AWS S3 object.
	S3SecretId *string `json:"S3SecretId,omitnil,omitempty" name:"S3SecretId"`

	// The key required to access the AWS S3 object.
	S3SecretKey *string `json:"S3SecretKey,omitnil,omitempty" name:"S3SecretKey"`
}

type S3OutputStorage struct {
	// The AWS S3 bucket.
	S3Bucket *string `json:"S3Bucket,omitnil,omitempty" name:"S3Bucket"`

	// The region of the AWS S3 bucket.
	S3Region *string `json:"S3Region,omitnil,omitempty" name:"S3Region"`

	// The key ID required to upload files to the AWS S3 object.
	S3SecretId *string `json:"S3SecretId,omitnil,omitempty" name:"S3SecretId"`

	// The key required to upload files to the AWS S3 object.
	S3SecretKey *string `json:"S3SecretKey,omitnil,omitempty" name:"S3SecretKey"`
}

type SampleSnapshotTaskInput struct {
	// Sampled screenshot template ID.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// List of up to 10 image or text watermarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// Target bucket of a sampled screenshot. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Output path of an image file after sampled screenshot taking, which can be a relative or absolute path.
	// If you need to define an output path, the path must end with `.{format}`. For variable names, refer to [Filename Variable](https://intl.cloud.tencent.com/document/product/862/37039?from_cn_redirect=1).Relative path example:
	// <li>Filename_{Variable name}.{format}.</li>
	// <li>Filename.{format}.</li>
	// Absolute path example:
	// <li>/Custom path/Filename_{Variable name}.{format}.</li>
	// If left empty, a relative path is used by default: `{inputName}_sampleSnapshot_{definition}_{number}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// Rule of the `{number}` variable in the sampled screenshot output path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitnil,omitempty" name:"ObjectNumberFormat"`
}

type SampleSnapshotTemplate struct {
	// Unique ID of a sampled screenshot template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Name of a sampled screenshot template.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: Enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: Disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Image format.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Sampled screenshot type.
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// Sampling interval.
	SampleInterval *uint64 `json:"SampleInterval,omitnil,omitempty" name:"SampleInterval"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: Stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: Fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: Fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: Fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type ScheduleAnalysisTaskResult struct {
	// The task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task has failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// The error code. 0 indicates the task is successful; other values indicate the task has failed. This parameter is not recommended. Please use `ErrCodeExt` instead.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// The error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input of the content analysis task.
	Input *AiAnalysisTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of the content analysis task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output []*AiAnalysisResult `json:"Output,omitnil,omitempty" name:"Output"`
}

type ScheduleQualityControlTaskResult struct {
	// The task status. Valid values: `PROCESSING`, `SUCCESS`, `FAIL`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value indicates the task has failed. For details, see [Error Codes](https://www.tencentcloud.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// The error code. `0` indicates the task is successful; other values indicate the task has failed. This parameter is not recommended. Please use `ErrCodeExt` instead.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// The error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Media quality inspection task input.
	Input *AiQualityControlTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Media quality inspection task output.Note: This field may return null, indicating that no valid values can be obtained.
	Output *QualityControlData `json:"Output,omitnil,omitempty" name:"Output"`
}

type ScheduleRecognitionTaskResult struct {
	// The task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task has failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// The error code. 0 indicates the task is successful; other values indicate the task has failed. This parameter is not recommended. Please use `ErrCodeExt` instead.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// The error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input of the content recognition task.
	Input *AiRecognitionTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of the content recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output []*AiRecognitionResult `json:"Output,omitnil,omitempty" name:"Output"`
}

type ScheduleReviewTaskResult struct {
	// The task status. Valid values: PROCESSING, SUCCESS, FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The error code. An empty string indicates the task is successful; any other value returned indicates the task has failed. For details, see [Error Codes](https://intl.cloud.tencent.com/document/product/1041/40249).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// The error code. 0 indicates the task is successful; other values indicate the task has failed. This parameter is not recommended. Please use `ErrCodeExt` instead.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// The error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The input of the content moderation task.
	Input *AiContentReviewTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// The output of the content moderation task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output []*AiContentReviewResult `json:"Output,omitnil,omitempty" name:"Output"`
}

type ScheduleSmartSubtitleTaskResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. An empty string indicates that the task is successful, and other values indicate that the task has failed. For specific values, see [Error Codes] (https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates that the task is successful, and other values indicate that the task has failed. (This field is not recommended. Use the new error code field ErrCodeExt instead.)
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Recognition task input.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Input *SmartSubtitlesTaskInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Recognition task output.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Output []*SmartSubtitlesResult `json:"Output,omitnil,omitempty" name:"Output"`

	// Task execution start time in [ISO datetime format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid value can be obtained.
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// Task execution completion time in [ISO datetime format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid value can be obtained.
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type ScheduleTask struct {
	// The scheme ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The scheme status. Valid values:
	// <li>PROCESSING</li>
	// <li>FINISH</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// If the value returned is not 0, there was a source error. If 0 is returned, refer to the error codes of the corresponding task type.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// If there was a source error, this parameter is the error message. For other errors, refer to the error messages of the corresponding task type.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The information of the file processed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// The metadata of the source video.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// The output of the scheme.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActivityResultSet []*ActivityResult `json:"ActivityResultSet,omitnil,omitempty" name:"ActivityResultSet"`
}

type SchedulesInfo struct {
	// The scheme ID.
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// The scheme name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScheduleName *string `json:"ScheduleName,omitnil,omitempty" name:"ScheduleName"`

	// The scheme type. Valid values:
	//  <li>`Preset`</li>
	// <li>`Custom` </li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The scheme status. Valid values:
	// `Enabled`
	// `Disabled`
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The trigger of the scheme.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The subtasks of the scheme.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Activities []*Activity `json:"Activities,omitnil,omitempty" name:"Activities"`

	// The bucket to save the output file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The directory to save the output file.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// The notification configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// The creation time in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The last updated time in [ISO date format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Resource ID. For those without an associated resource ID, fill in with an account's primary resource ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type ScratchRepairConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The strength. Value range: 0.0-1.0
	// Default value: 0.0
	// Note: This field may return null, indicating that no valid values can be obtained.
	Intensity *float64 `json:"Intensity,omitnil,omitempty" name:"Intensity"`
}

type SecurityGroupInfo struct {
	// Security group ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Security group name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Allowlist list.
	Whitelist []*string `json:"Whitelist,omitnil,omitempty" name:"Whitelist"`

	// List of bound input streams.
	// Note: This field may return null, indicating that no valid value can be obtained.
	OccupiedInputs []*string `json:"OccupiedInputs,omitnil,omitempty" name:"OccupiedInputs"`

	// Security group address.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// List of bound output streams.
	// Note: This field may return null, indicating that no valid value can be obtained.
	OccupiedOutputs []*string `json:"OccupiedOutputs,omitnil,omitempty" name:"OccupiedOutputs"`
}

type SegmentRecognitionItem struct {

	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`


	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`


	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`


	SegmentUrl *string `json:"SegmentUrl,omitnil,omitempty" name:"SegmentUrl"`

	// Segment cover.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CovImgUrl *string `json:"CovImgUrl,omitnil,omitempty" name:"CovImgUrl"`

	// Segment title.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Segment summary.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// Segmentation keywords.
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// The start time of a live streaming segment, in the ISO date format.
	// Note: This field may return null, indicating that no valid value can be obtained.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of a live streaming segment, in the ISO date format.
	// Note: This field may return null, indicating that no valid value can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Specifies the character ID.
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

type SegmentSpecificInfo struct {
	// Switch for segment duration at startup. Optional values:
	// on: Turn on the switch
	// off: Turn off the switch
	// Default value: off
	// Note: This field may return null, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Segment duration at startup. Unit: second
	// Note: This field may return null, indicating that no valid value can be obtained.
	FragmentTime *int64 `json:"FragmentTime,omitnil,omitempty" name:"FragmentTime"`

	// Number of effective segments, indicating the first FragmentEndNum segments with FragmentTime. Value range: >=1
	// Note: This field may return null, indicating that no valid value can be obtained.
	FragmentEndNum *int64 `json:"FragmentEndNum,omitnil,omitempty" name:"FragmentEndNum"`
}

type SharpEnhanceConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The strength. Value range: 0.0-1.0
	// Default value: 0.0
	// Note: This field may return null, indicating that no valid values can be obtained.
	Intensity *float64 `json:"Intensity,omitnil,omitempty" name:"Intensity"`
}

type SimpleAesDrm struct {
	// The URI of decryption key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`

	// The encryption key (a 32-byte string).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// The initialization vector for encryption (a 32-byte string).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vector *string `json:"Vector,omitnil,omitempty" name:"Vector"`
}

type SmartSubtitleTaskAsrFullTextResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. An empty string indicates that the task is successful, and other values indicate that the task has failed. For specific values, see [Error Codes] (https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates that the task is successful, and other values indicate that the task has failed. (This field is not recommended. Use the new error code field ErrCodeExt instead.)
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Input information on the full speech recognition task.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Input *SmartSubtitleTaskResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Output information on the full speech recognition task.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Output *SmartSubtitleTaskAsrFullTextResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// Task progress.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type SmartSubtitleTaskAsrFullTextResultOutput struct {
	// List of segments for full speech recognition.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SegmentSet []*SmartSubtitleTaskAsrFullTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Subtitle file path.
	SubtitlePath *string `json:"SubtitlePath,omitnil,omitempty" name:"SubtitlePath"`
}

type SmartSubtitleTaskAsrFullTextSegmentItem struct {
	// Confidence of a recognized segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Start time offset of a recognized segment, in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognized segment, in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Recognized text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Word timestamp information.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	Wordlist []*WordResult `json:"Wordlist,omitnil,omitempty" name:"Wordlist"`
}

type SmartSubtitleTaskBatchOutput struct {
	// Task progress.
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Task status, including PROCESSING, SUCCESS, and FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. An empty string indicates that the task is successful, and other values indicate that the task has failed. For specific values, see [Error Codes] (https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Translation task output information.
	// Note: This field may return null, indicating that no valid value can be obtained.
	TransTextTask *SmartSubtitleTaskTransTextResultOutput `json:"TransTextTask,omitnil,omitempty" name:"TransTextTask"`

	// Output information on the full speech recognition task.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AsrFullTextTask *SmartSubtitleTaskAsrFullTextResultOutput `json:"AsrFullTextTask,omitnil,omitempty" name:"AsrFullTextTask"`
}

type SmartSubtitleTaskResultInput struct {
	// Smart subtitle template ID.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Custom smart subtitle parameter. It takes effect when Definition is set to 0.
	// This parameter is used in high customization scenarios. It is recommended that you preferentially use Definition to specify smart subtitle parameters.
	// Note: This field may return null, indicating that no valid value can be obtained.
	RawParameter *RawSmartSubtitleParameter `json:"RawParameter,omitnil,omitempty" name:"RawParameter"`
}

type SmartSubtitleTaskTransTextResult struct {
	// Task status, including PROCESSING, SUCCESS, and FAIL.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code. An empty string indicates that the task is successful, and other values indicate that the task has failed. For specific values, see [Error Codes] (https://intl.cloud.tencent.com/document/product/862/50369?from_cn_redirect=1#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81).
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// Error code. 0 indicates that the task is successful, and other values indicate that the task has failed. (This field is not recommended. Use the new error code field ErrCodeExt instead.)
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Translation task input information.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Input *SmartSubtitleTaskResultInput `json:"Input,omitnil,omitempty" name:"Input"`

	// Translation task output information.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Output *SmartSubtitleTaskTransTextResultOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// Task progress.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type SmartSubtitleTaskTransTextResultOutput struct {
	// List of segments for translation.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SegmentSet []*SmartSubtitleTaskTransTextSegmentItem `json:"SegmentSet,omitnil,omitempty" name:"SegmentSet"`

	// Subtitle file path.
	SubtitlePath *string `json:"SubtitlePath,omitnil,omitempty" name:"SubtitlePath"`
}

type SmartSubtitleTaskTransTextSegmentItem struct {
	// Confidence of a recognized segment. Value range: 0-100.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Start time offset of a recognized segment, in seconds.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a recognized segment, in seconds.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Recognized text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Translated text.
	Trans *string `json:"Trans,omitnil,omitempty" name:"Trans"`

	// Word timestamp information.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	Wordlist []*WordResult `json:"Wordlist,omitnil,omitempty" name:"Wordlist"`
}

type SmartSubtitleTemplateItem struct {
	// Unique identifier of the smart subtitle template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Smart subtitle template name.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Smart subtitle template description.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Template type. Valid values:
	// * Preset: system preset template
	// * Custom: user-defined template
	// Note: This field may return null, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// ASR hotword lexicon parameter.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AsrHotWordsConfigure *AsrHotWordsConfigure `json:"AsrHotWordsConfigure,omitnil,omitempty" name:"AsrHotWordsConfigure"`

	// Name of the hotword lexicon associated with the template.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AsrHotWordsLibraryName *string `json:"AsrHotWordsLibraryName,omitnil,omitempty" name:"AsrHotWordsLibraryName"`

	// Source language of the video with smart subtitles.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// zh-PY: Chinese-English-Cantonese
	// zh-medical: Medical Chinese
	// yue: Cantonese
	// vi: Vietnamese
	// ms: Malay
	// id: Indonesian
	// fli: Filipino
	// th: Thai
	// pt: Portuguese
	// tr: Turkish
	// ar: Arabic
	// es: Spanish
	// hi: Hindi
	// fr: French
	// de: German
	// zh-dialect: Chinese dialect
	VideoSrcLanguage *string `json:"VideoSrcLanguage,omitnil,omitempty" name:"VideoSrcLanguage"`

	// Smart subtitle file format.
	// vtt: WebVTT format
	// If this field is left blank, no subtitle file will be generated.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`

	// Smart subtitle language type.
	// 0: source language1: target language
	// 2: source language + target language
	// The value can only be 0 when TranslateSwitch is set to OFF.The value can only be 1 or 2 when TranslateSwitch is set to ON.
	SubtitleType *int64 `json:"SubtitleType,omitnil,omitempty" name:"SubtitleType"`

	// Subtitle translation switch.
	// ON: enable translation
	// OFF: disable translation
	// Note: This field may return null, indicating that no valid value can be obtained.
	TranslateSwitch *string `json:"TranslateSwitch,omitnil,omitempty" name:"TranslateSwitch"`

	// Target language for subtitle translation.
	// This field takes effect when TranslateSwitch is set to ON.
	// Supported languages:
	// zh: Simplified Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// fr: French
	// es: Spanish
	// it: Italian
	// de: German
	// tr: Turkish
	// ru: Russian
	// pt: Portuguese
	// vi: Vietnamese
	// id: Indonesian
	// ms: Malay
	// th: Thai
	// ar: Arabic
	// hi: Hindi
	// Note: This field may return null, indicating that no valid value can be obtained.
	TranslateDstLanguage *string `json:"TranslateDstLanguage,omitnil,omitempty" name:"TranslateDstLanguage"`

	// Template creation time in [ISO datetime format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modification time of the template in [ISO datetime format](https://intl.cloud.tencent.com/document/product/862/37710?from_cn_redirect=1#52).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Alias of the preset smart subtitle template.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`
}

type SmartSubtitlesResult struct {
	// Task type. Valid values:
	// <Li>AsrFullTextRecognition: full speech recognition</li>
	// <Li>TransTextRecognition: speech translation</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Full speech recognition result. When Type is
	//  set to AsrFullTextRecognition, this parameter takes effect.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AsrFullTextTask *SmartSubtitleTaskAsrFullTextResult `json:"AsrFullTextTask,omitnil,omitempty" name:"AsrFullTextTask"`

	// Translation result. When Type is
	// 
	//  set to TransTextRecognition, this parameter takes effect.
	// Note: This field may return null, indicating that no valid value can be obtained.
	TransTextTask *SmartSubtitleTaskTransTextResult `json:"TransTextTask,omitnil,omitempty" name:"TransTextTask"`
}

type SmartSubtitlesTaskInput struct {
	// Smart subtitle template ID.	
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// User extension field, which does not need to be filled in for general scenarios.
	UserExtPara *string `json:"UserExtPara,omitnil,omitempty" name:"UserExtPara"`

	// Custom smart subtitle parameter. It takes effect when Definition is set to 0. This parameter is used in high customization scenarios. It is recommended that you preferentially use Definition to specify smart subtitle parameters.	
	// Note: This field may return null, indicating that no valid value can be obtained.
	RawParameter *RawSmartSubtitleParameter `json:"RawParameter,omitnil,omitempty" name:"RawParameter"`
}

type SnapshotByTimeOffsetTaskInput struct {
	// ID of a time point screenshot template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// List of screenshot time points in the format of `s` or `%`:
	// <li>If the string ends in `s`, it means that the time point is in seconds; for example, `3.5s` means that the time point is the 3.5th second;</li>
	// <li>If the string ends in `%`, it means that the time point is the specified percentage of the video duration; for example, `10%` means that the time point is 10% of the video duration.</li>
	ExtTimeOffsetSet []*string `json:"ExtTimeOffsetSet,omitnil,omitempty" name:"ExtTimeOffsetSet"`

	// List of time points of screenshots in <font color=red>seconds</font>.
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitnil,omitempty" name:"TimeOffsetSet"`

	// List of up to 10 image or text watermarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// Target bucket of a generated time point screenshot file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Output path for an image file of screenshots taken at specific time points, which can be a relative or absolute path.
	// If you need to define an output path, the path must end with `.{format}`. For variable names, refer to [Filename Variable](https://intl.cloud.tencent.com/document/product/862/37039?from_cn_redirect=1).
	// Relative path example:
	// <li>Filename_{Variable name}.{format}.</li>
	// <li>Filename.{format}.</li>
	// Absolute path example:
	// <li>/Custom path/Filename_{Variable name}.{format}.</li>
	// If left empty, a relative path is used by default: `{inputName}_snapshotByTimeOffset_{definition}_{number}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// Rule of the `{number}` variable in the time point screenshot output path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitnil,omitempty" name:"ObjectNumberFormat"`
}

type SnapshotByTimeOffsetTemplate struct {
	// Unique ID of a time point screenshot template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Name of a time point screenshot template.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096].
	// <li>If both `Width` and `Height` are 0, the resolution will be the same as that of the source video;</li>
	// <li>If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled;</li>
	// <li>If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;</li>
	// <li>If both `Width` and `Height` are not 0, the custom resolution will be used.</li>
	// Default value: 0.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Resolution adaption. Valid values:
	// <li>open: Enabled. In this case, `Width` represents the long side of a video, while `Height` the short side;</li>
	// <li>close: Disabled. In this case, `Width` represents the width of a video, while `Height` the height.</li>
	// Default value: open.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Image format.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:
	// <li> stretch: Stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: Fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: Fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: Fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type SpekeDrm struct {
	// Resource tagging. the field content is user-customized.
	// It supports 1 to 128 characters consisting of digits, letters, underscores (_), and hyphens (-).
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// DRM manufacturer access address. the field content is obtained from the drm manufacturer.
	// 
	// Note: different DRM manufacturers have different limitations on the number of substreams. for example, PallyCon limits the number of substreams to no more than 5, and DRMtoday only supports encryption of up to 9 substreams.
	KeyServerUrl *string `json:"KeyServerUrl,omitnil,omitempty" name:"KeyServerUrl"`

	// Encryption initialization vector (32-byte string). the field content is user-customized.
	Vector *string `json:"Vector,omitnil,omitempty" name:"Vector"`

	// Encryption method. cbcs: default method of FairPlay; cenc: default method of PlayReady and Widevine.
	// 
	// cbcs: supported by PlayReady, Widevine, and FairPlay
	// cenc: supported by PlayReady and Widevine
	EncryptionMethod *string `json:"EncryptionMethod,omitnil,omitempty" name:"EncryptionMethod"`

	// Substream encryption rule. Default value: preset0.
	// preset 0: use the same key to encrypt all substreams
	// preset1: use different keys for each substream
	EncryptionPreset *string `json:"EncryptionPreset,omitnil,omitempty" name:"EncryptionPreset"`
}

type SubtitleTemplate struct {
	// The URL of the subtitles to add to the video.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// The subtitle track to add to the video. If both `Path` and `StreamIndex` are specified, `Path` will be used. You need to specify at least one of the two parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StreamIndex *int64 `json:"StreamIndex,omitnil,omitempty" name:"StreamIndex"`

	// The font. Valid values:
	// <li>`hei.ttf`: Heiti.</li>
	// <li>`song.ttf`: Songti.</li>
	// <li>`simkai.ttf`: Kaiti.</li>
	// <li>`arial.ttf`: Arial.</li>
	// The default is `hei.ttf`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FontType *string `json:"FontType,omitnil,omitempty" name:"FontType"`

	// The font size (pixels). If this is not specified, the font size in the subtitle file will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FontSize *string `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// The font color in 0xRRGGBB format. Default value: 0xFFFFFF (white).
	// Note: This field may return null, indicating that no valid values can be obtained.
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// The text transparency. Value range: 0-1.
	// <li>`0`: Fully transparent.</li>
	// <li>`1`: Fully opaque.</li>
	// Default value: 1.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FontAlpha *float64 `json:"FontAlpha,omitnil,omitempty" name:"FontAlpha"`
}

type SuperResolutionConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The strength. Valid values:
	// <li>lq: For low-resolution videos with obvious noise</li>
	// <li>hq: For high-resolution videos</li>
	// Default value: lq.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The ratio of the target resolution to the original resolution. Valid values:
	// <li>2</li>
	// Default value: 2.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
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
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// Watermark height, which supports six formats of px, %, W%, H%, S%, and L%:
	// <li>If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px; if `0px` is entered
	//  and `Width` is not `0px`, the watermark height will be proportionally scaled based on the source SVG image; if `0px` is entered for both `Width` and `Height`, the watermark height will be the height of the source SVG image;</li>
	// <li>If the string ends in `W%`, the `Height` of the watermark will be the specified percentage of the video width; for example, `10W%` means that `Height` is 10% of the video width;</li>
	// <li>If the string ends in `H%`, the `Height` of the watermark will be the specified percentage of the video height; for example, `10H%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in `S%`, the `Height` of the watermark will be the specified percentage of the short side of the video; for example, `10S%` means that `Height` is 10% of the short side of the video;</li>
	// <li>If the string ends in `L%`, the `Height` of the watermark will be the specified percentage of the long side of the video; for example, `10L%` means that `Height` is 10% of the long side of the video;</li>
	// <li>If the string ends in %, the meaning is the same as `H%`.</li>
	// Default value: 0 px.
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`
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
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// Watermark Height, which supports six formats of px, %, W%, H%, S%, and L%:
	// <li>If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px; if `0px` is entered
	//  and `Width` is not `0px`, the watermark height will be proportionally scaled based on the source SVG image; if `0px` is entered for both `Width` and `Height`, the watermark size will be the size of the source SVG image;</li>
	// <li>If the string ends in `W%`, the `Height` of the watermark will be the specified percentage of the video width; for example, `10W%` means that `Height` is 10% of the video width;</li>
	// <li>If the string ends in `H%`, the `Height` of the watermark will be the specified percentage of the video height; for example, `10H%` means that `Height` is 10% of the video height;</li>
	// <li>If the string ends in `S%`, the `Height` of the watermark will be the specified percentage of the short side of the video; for example, `10S%` means that `Height` is 10% of the short side of the video;</li>
	// <li>If the string ends in `L%`, the `Height` of the watermark will be the specified percentage of the long side of the video; for example, `10L%` means that `Height` is 10% of the long side of the video;</li>
	// <li>If the string ends in %, the meaning is the same as `W%`.</li>
	// Default value: 0px.
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`
}

type TEHDConfig struct {
	// TESHD type. Valid values:
	// <li>TEHD-100: TESHD-100.</li>
	// If this parameter is left empty, TESHD will not be enabled.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Maximum bitrate, which is valid when `Type` is `TESHD`.
	// If this parameter is left empty or 0 is entered, there will be no upper limit for bitrate.
	MaxVideoBitrate *int64 `json:"MaxVideoBitrate,omitnil,omitempty" name:"MaxVideoBitrate"`
}

type TEHDConfigForUpdate struct {
	// The TSC type. Valid values:
	// <li>`TEHD-100`: TSC-100 (video TSC). </li>
	// <li>`TEHD-200`: TSC-200 (audio TSC). </li>
	// If this parameter is left blank, no modification will be made.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The maximum video bitrate. If this parameter is not specified, no modifications will be made.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxVideoBitrate *int64 `json:"MaxVideoBitrate,omitnil,omitempty" name:"MaxVideoBitrate"`
}

type TagConfigureInfo struct {
	// Switch of intelligent tagging task. Valid values:
	// <li>ON: enables intelligent tagging task;</li>
	// <li>OFF: disables intelligent tagging task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type TagConfigureInfoForUpdate struct {
	// Switch of intelligent tagging task. Valid values:
	// <li>ON: enables intelligent tagging task;</li>
	// <li>OFF: disables intelligent tagging task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type TaskNotifyConfig struct {
	// The notification type. Valid values:
	// <li>`CMQ`: This value is no longer used. Please use `TDMQ-CMQ` instead.</li>
	// <li>`TDMQ-CMQ`: Message queue</li>
	// <li>`URL`: If `NotifyType` is set to `URL`, HTTP callbacks are sent to the URL specified by `NotifyUrl`. HTTP and JSON are used for the callbacks. The packet contains the response parameters of the `ParseNotification` API.</li>
	// <li>`SCF`: This notification type is not recommended. You need to configure it in the SCF console.</li>
	// <li>`AWS-SQS`: AWS queue. This type is only supported for AWS tasks, and the queue must be in the same region as the AWS bucket.</li>
	// <font color="red">Note: If you do not pass this parameter or pass in an empty string, `CMQ` will be used. To use a different notification type, specify this parameter accordingly.</font>
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// Workflow notification method. Valid values: Finish, Change. If this parameter is left empty, `Finish` will be used.
	NotifyMode *string `json:"NotifyMode,omitnil,omitempty" name:"NotifyMode"`

	// HTTP callback URL, required if `NotifyType` is set to `URL`
	NotifyUrl *string `json:"NotifyUrl,omitnil,omitempty" name:"NotifyUrl"`

	// The CMQ or TDMQ-CMQ model. Valid values: Queue, Topic.
	CmqModel *string `json:"CmqModel,omitnil,omitempty" name:"CmqModel"`

	// The CMQ or TDMQ-CMQ region, such as `sh` (Shanghai) or `bj` (Beijing).
	CmqRegion *string `json:"CmqRegion,omitnil,omitempty" name:"CmqRegion"`

	// The CMQ or TDMQ-CMQ topic to receive notifications. This parameter is valid when `CmqModel` is `Topic`.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// The CMQ or TDMQ-CMQ queue to receive notifications. This parameter is valid when `CmqModel` is `Queue`.
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// The AWS SQS queue. This parameter is required if `NotifyType` is `AWS-SQS`.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	AwsSQS *AwsSQS `json:"AwsSQS,omitnil,omitempty" name:"AwsSQS"`

	// The key used to generate the callback signature.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NotifyKey *string `json:"NotifyKey,omitnil,omitempty" name:"NotifyKey"`
}

type TaskOutputStorage struct {
	// The storage type for a media processing output file. Valid values:
	// <li>`COS`: Tencent Cloud COS</li>
	// <li>`AWS-S3`: AWS S3. This type is only supported for AWS tasks, and the output bucket must be in the same region as the bucket of the source file.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The location to save the output object in COS. This parameter is valid and required when `Type` is COS.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CosOutputStorage *CosOutputStorage `json:"CosOutputStorage,omitnil,omitempty" name:"CosOutputStorage"`

	// The AWS S3 bucket to save the output file. This parameter is required if `Type` is `AWS-S3`.
	// Note: This field may return null, indicating that no valid value can be obtained.
	S3OutputStorage *S3OutputStorage `json:"S3OutputStorage,omitnil,omitempty" name:"S3OutputStorage"`
}

type TaskSimpleInfo struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task type. Valid values:
	// <li> WorkflowTask: Workflow processing task;</li>
	// <li> LiveProcessTask: Live stream processing task.</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Creation time of a task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Start time of task execution in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F). If the task has not been started yet, this field will be `0000-00-00T00:00:00Z`.
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// End time of a task in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F). If the task has not been completed yet, this field will be `0000-00-00T00:00:00Z`.
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// The subtask type.
	SubTaskTypes []*string `json:"SubTaskTypes,omitnil,omitempty" name:"SubTaskTypes"`
}

type TerrorismConfigureInfo struct {
	// The parameters for detecting sensitive information in images.
	ImgReviewInfo *TerrorismImgReviewTemplateInfo `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// The parameters for detecting sensitive information based on OCR.
	OcrReviewInfo *TerrorismOcrReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type TerrorismConfigureInfoForUpdate struct {
	// The parameters for detecting sensitive information in images.
	ImgReviewInfo *TerrorismImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitnil,omitempty" name:"ImgReviewInfo"`

	// The parameters for detecting sensitive information based on OCR.
	OcrReviewInfo *TerrorismOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type TerrorismImgReviewTemplateInfo struct {
	// Whether to detect sensitive information in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Sensitive content filter tags. The auditing results including the selected tags are returned. If the filter tag is empty, all auditing results will be returned. Valid values:
	// <li>guns: weapons and guns;</li>
	// <li>crowd: crowd gathering;</li>
	// <li>bloody: bloodiness;</li>
	// <li>police: police force;</li>
	// <li>banners: sensitive flags;</li>
	// <li>militant: militants;</li>
	// <li>explosion: explosions and fires;</li>
	// <li>terrorists: sensitive persons.</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 90 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 80 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type TerrorismImgReviewTemplateInfoForUpdate struct {
	// Whether to detect sensitive information in images. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Sensitive content filter tags. The auditing results including the selected tags are returned. If the filter tag is empty, all auditing results will be returned. Valid values:
	// <li>guns: weapons and guns;</li>
	// <li>crowd: crowd gathering;</li>
	// <li>bloody: bloodiness;</li>
	// <li>police: police force;</li>
	// <li>banners: sensitive flags;</li>
	// <li>militant: militants;</li>
	// <li>explosion: explosions and fires;</li>
	// <li>terrorists: sensitive persons.</li>
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type TerrorismOcrReviewTemplateInfo struct {
	// Whether to detect sensitive information based on OCR. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type TerrorismOcrReviewTemplateInfoForUpdate struct {
	// Whether to detect sensitive information based on OCR. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0–100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0–100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type TextWatermarkTemplateInput struct {
	// Font type. Currently, two types are supported:
	// <li>simkai.ttf: Both Chinese and English are supported;</li>
	// <li>arial.ttf: Only English is supported.</li>
	FontType *string `json:"FontType,omitnil,omitempty" name:"FontType"`

	// Font size, in the format of Npx. N is a numerical value with a value range of [0, 1] or [8, 4096].
	FontSize *string `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// Font color in 0xRRGGBB format. Default value: 0xFFFFFF (white).
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// Text transparency. Value range: (0, 1]
	// <li>0: Completely transparent</li>
	// <li>1: Completely opaque</li>
	// Default value: 1.
	FontAlpha *float64 `json:"FontAlpha,omitnil,omitempty" name:"FontAlpha"`

	// Text content, up to 100 characters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TextContent *string `json:"TextContent,omitnil,omitempty" name:"TextContent"`
}

type TextWatermarkTemplateInputForUpdate struct {
	// Font type. Currently, two types are supported:
	// <li>simkai.ttf: Both Chinese and English are supported;</li>
	// <li>arial.ttf: Only English is supported.</li>
	FontType *string `json:"FontType,omitnil,omitempty" name:"FontType"`

	// Font size, in the format of Npx. N is a numerical value with a value range of [0, 1] or [8, 4096].
	FontSize *string `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// Font color in 0xRRGGBB format. Default value: 0xFFFFFF (white).
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// Text transparency. Value range: (0, 1]
	// <li>0: Completely transparent</li>
	// <li>1: Completely opaque</li>
	FontAlpha *float64 `json:"FontAlpha,omitnil,omitempty" name:"FontAlpha"`

	// Text content, up to 100 characters.
	TextContent *string `json:"TextContent,omitnil,omitempty" name:"TextContent"`
}

type TrackInfo struct {
	// The serial number of the audio track and sound channel.
	// <li>When the value of SelectType is track, this value is an integer, for example: 1.
	// <li>When the value of SelectType is track_channel, this value is a decimal, for example: 1.0.
	// <li>Default value: 1.0.
	// The integer part represents the audio track serial number, and the decimal part represents the sound channel. The audio track serial number is the stream index value of the audio track, which can be 0 or a positive integer. The decimal part supports up to 2 decimal places, and only 0 - 63 is supported. However, when the Codec is aac/eac3/ac3, only 0 - 15 is supported for the decimal part. For example: for an audio track with a stream index value of 1, 1.0 represents the first sound channel of this audio track, and 1.1 represents the second sound channel of this audio track.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	TrackNum *string `json:"TrackNum,omitnil,omitempty" name:"TrackNum"`

	// The volume of the sound channel.
	// <li>When the value of AudioChannel is 1, the length of this array is 1. For example: [6].
	// <li>When the value of AudioChannel is 2, the length of this array is 2. For example: [0,6].
	// <li>When the value of AudioChannel is 6, the length of this array is greater than 2 and less than 16. For example: [-60,0,0,6].
	// 
	// Please specify the value array for this parameter. The value range is between -60 and 6, where -60 indicates mute, 0 maintains the original volume, and 6 doubles the original volume. The default value is -60. Please note: This field supports up to 3 decimal places.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	ChannelVolume []*float64 `json:"ChannelVolume,omitnil,omitempty" name:"ChannelVolume"`
}

type TranscodeTaskInput struct {
	// ID of a video transcoding template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Custom video transcoding parameter, which is valid if `Definition` is 0.
	// This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the transcoding parameter preferably.
	RawParameter *RawTranscodeParameter `json:"RawParameter,omitnil,omitempty" name:"RawParameter"`

	// Video transcoding custom parameter, which is valid when `Definition` is not 0.
	// When any parameters in this structure are entered, they will be used to override corresponding parameters in templates.
	// This parameter is used in highly customized scenarios. We recommend you only use `Definition` to specify the transcoding parameter.
	// Note: this field may return `null`, indicating that no valid value was found.
	OverrideParameter *OverrideTranscodeParameter `json:"OverrideParameter,omitnil,omitempty" name:"OverrideParameter"`

	// List of up to 10 image or text watermarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// List of blurs. Up to 10 ones can be supported.
	MosaicSet []*MosaicInput `json:"MosaicSet,omitnil,omitempty" name:"MosaicSet"`

	// Start time offset of a transcoded video, in seconds.
	// <li>If this parameter is left empty or set to 0, the transcoded video will start at the same time as the original video.</li>
	// <li>If this parameter is set to a positive number (n for example), the transcoded video will start at the nth second of the original video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the transcoded video will start at the nth second before the end of the original video.</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a transcoded video, in seconds.
	// <li>If this parameter is left empty or set to 0, the transcoded video will end at the same time as the original video.</li>
	// <li>If this parameter is set to a positive number (n for example), the transcoded video will end at the nth second of the original video.</li>
	// <li>If this parameter is set to a negative number (-n for example), the transcoded video will end at the nth second before the end of the original video.</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Target bucket of an output file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// Output path of the main file after transcoding, which can be a relative or absolute path.
	// If you need to define an output path, the path must end with `.{format}`. For variable names, refer to [Filename Variable](https://intl.cloud.tencent.com/document/product/862/37039?from_cn_redirect=1).Relative path example:
	// <li>Filename_{Variable name}.{format}.</li>
	// <li>Filename.{format}.</li>
	// Absolute path example:
	// <li>/Custom path/Filename_{Variable name}.{format}.</li>
	// If left empty, a relative path is used by default: `{inputName}_transcode_{definition}.{format}`.
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// Path to an output file part (the path to ts during transcoding to HLS), which can only be a relative path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_transcode_{definition}_{number}.{format}`.
	SegmentObjectName *string `json:"SegmentObjectName,omitnil,omitempty" name:"SegmentObjectName"`

	// Rule of the `{number}` variable in the output path after transcoding.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectNumberFormat *NumberFormat `json:"ObjectNumberFormat,omitnil,omitempty" name:"ObjectNumberFormat"`

	// Opening and closing credits parameters
	// Note: this field may return `null`, indicating that no valid value was found.
	HeadTailParameter *HeadTailParameter `json:"HeadTailParameter,omitnil,omitempty" name:"HeadTailParameter"`
}

type TranscodeTemplate struct {
	// Unique ID of a transcoding template.
	Definition *string `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Container format. Valid values: mp4, flv, hls, mp3, flac, ogg.
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// Name of a transcoding template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Template type. Valid values:
	// <li>Preset: Preset template;</li>
	// <li>Custom: Custom template.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether to remove video data. Valid values:
	// <li>0: Retain;</li>
	// <li>1: Remove.</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitnil,omitempty" name:"RemoveVideo"`

	// Whether to remove audio data. Valid values:
	// <li>0: Retain;</li>
	// <li>1: Remove.</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitnil,omitempty" name:"RemoveAudio"`

	// Video stream configuration parameter. This field is valid only when `RemoveVideo` is 0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitnil,omitempty" name:"VideoTemplate"`

	// Audio stream configuration parameter. This field is valid only when `RemoveAudio` is 0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitnil,omitempty" name:"AudioTemplate"`

	// TESHD transcoding parameter. To enable it, please contact your Tencent Cloud sales rep.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitnil,omitempty" name:"TEHDConfig"`

	// Container format filter. Valid values:
	// <li>Video: Video container format that can contain both video stream and audio stream;</li>
	// <li>PureAudio: Audio container format that can contain only audio stream.</li>
	ContainerType *string `json:"ContainerType,omitnil,omitempty" name:"ContainerType"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Audio/Video enhancement configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnhanceConfig *EnhanceConfig `json:"EnhanceConfig,omitnil,omitempty" name:"EnhanceConfig"`

	// Transcoding template alias.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`
}

type TranslateConfigureInfo struct {
	// Switch of a full speech recognition task. Valid values:
	// <li>ON: Enables an intelligent full speech recognition task;</li>
	// <li>OFF: Disables an intelligent full speech recognition task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`


	SourceLanguage *string `json:"SourceLanguage,omitnil,omitempty" name:"SourceLanguage"`


	DestinationLanguage *string `json:"DestinationLanguage,omitnil,omitempty" name:"DestinationLanguage"`

	// Generated subtitle file format. Leaving it as an empty string means no subtitle file will be generated. Valid value:
	// <li>vtt: Generate a WebVTT subtitle file.</li>
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubtitleFormat *string `json:"SubtitleFormat,omitnil,omitempty" name:"SubtitleFormat"`
}

type UrlInputInfo struct {
	// URL of a video.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type UserDefineAsrTextReviewTemplateInfo struct {
	// Switch of a custom speech audit task. Valid values:
	// <li>ON: Enables a custom speech audit task;</li>
	// <li>OFF: Disables a custom speech audit task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Custom speech filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom speech keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineAsrTextReviewTemplateInfoForUpdate struct {
	// Switch of a custom speech audit task. Valid values:
	// <li>ON: Enables a custom speech audit task;</li>
	// <li>OFF: Disables a custom speech audit task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Custom speech filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom speech keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineConfigureInfo struct {
	// Control parameter of custom figure audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FaceReviewInfo *UserDefineFaceReviewTemplateInfo `json:"FaceReviewInfo,omitnil,omitempty" name:"FaceReviewInfo"`

	// Control parameter of custom speech audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfo `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// Control parameter of custom text audit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfo `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type UserDefineConfigureInfoForUpdate struct {
	// Control parameter of custom figure audit.
	FaceReviewInfo *UserDefineFaceReviewTemplateInfoForUpdate `json:"FaceReviewInfo,omitnil,omitempty" name:"FaceReviewInfo"`

	// Control parameter of custom speech audit.
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitnil,omitempty" name:"AsrReviewInfo"`

	// Control parameter of custom text audit.
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitnil,omitempty" name:"OcrReviewInfo"`
}

type UserDefineFaceReviewTemplateInfo struct {
	// Switch of a custom figure audit task. Valid values:
	// <li>ON: Enables a custom figure audit task;</li>
	// <li>OFF: Disables a custom figure audit task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Custom figure filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for the custom figure library.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 97 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 95 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineFaceReviewTemplateInfoForUpdate struct {
	// Switch of a custom figure audit task. Valid values:
	// <li>ON: Enables a custom figure audit task;</li>
	// <li>OFF: Disables a custom figure audit task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Custom figure filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for the custom figure library.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfo struct {
	// Switch of a custom text audit task. Valid values:
	// <li>ON: Enables a custom text audit task;</li>
	// <li>OFF: Disables a custom text audit task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Custom text filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom text keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. If this parameter is left empty, 100 will be used by default. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. If this parameter is left empty, 75 will be used by default. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfoForUpdate struct {
	// Switch of a custom text audit task. Valid values:
	// <li>ON: Enables a custom text audit task;</li>
	// <li>OFF: Disables a custom text audit task.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Custom text filter tag. If an audit result contains the selected tag, it will be returned; if the filter tag is empty, all audit results will be returned. To use the tag filtering feature, you need to add the corresponding tag when adding materials for custom text keywords.
	// There can be up to 10 tags, each with a length limit of 16 characters.
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// Threshold score for violation. If this score is reached or exceeded during intelligent audit, it will be deemed that a suspected violation has occurred. Value range: 0-100.
	BlockConfidence *int64 `json:"BlockConfidence,omitnil,omitempty" name:"BlockConfidence"`

	// Threshold score for human audit. If this score is reached or exceeded during intelligent audit, human audit will be considered necessary. Value range: 0-100.
	ReviewConfidence *int64 `json:"ReviewConfidence,omitnil,omitempty" name:"ReviewConfidence"`
}

type VideoDenoiseConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>ON</li>
	// <li>OFF</li>
	// Default value: ON.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The strength. Valid values:
	// <li>weak</li>
	// <li>strong</li>
	// Default value: weak.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type VideoEnhanceConfig struct {
	// Frame interpolation configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FrameRate *FrameRateConfig `json:"FrameRate,omitnil,omitempty" name:"FrameRate"`

	// Super resolution configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SuperResolution *SuperResolutionConfig `json:"SuperResolution,omitnil,omitempty" name:"SuperResolution"`

	// HDR configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Hdr *HdrConfig `json:"Hdr,omitnil,omitempty" name:"Hdr"`

	// Image noise removal configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Denoise *VideoDenoiseConfig `json:"Denoise,omitnil,omitempty" name:"Denoise"`

	// Overall enhancement configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImageQualityEnhance *ImageQualityEnhanceConfig `json:"ImageQualityEnhance,omitnil,omitempty" name:"ImageQualityEnhance"`

	// Color enhancement configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ColorEnhance *ColorEnhanceConfig `json:"ColorEnhance,omitnil,omitempty" name:"ColorEnhance"`

	// Detail enhancement configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SharpEnhance *SharpEnhanceConfig `json:"SharpEnhance,omitnil,omitempty" name:"SharpEnhance"`

	// Face enhancement configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FaceEnhance *FaceEnhanceConfig `json:"FaceEnhance,omitnil,omitempty" name:"FaceEnhance"`

	// Low-light enhancement configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LowLightEnhance *LowLightEnhanceConfig `json:"LowLightEnhance,omitnil,omitempty" name:"LowLightEnhance"`

	// Banding removal configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScratchRepair *ScratchRepairConfig `json:"ScratchRepair,omitnil,omitempty" name:"ScratchRepair"`

	// Artifact removal (smoothing) configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ArtifactRepair *ArtifactRepairConfig `json:"ArtifactRepair,omitnil,omitempty" name:"ArtifactRepair"`
}

type VideoTemplateInfo struct {
	// Encoding format for video streams. Optional values:
	// <li>h264: H.264 encoding</li>
	// <li>h265: H.265 encoding</li>
	// <li>h266: H.266 encoding</li>
	// <li>av1: AOMedia Video 1 encoding</li>
	// <li>vp8: VP8 encoding</li>
	// <li>vp9: VP9 encoding</li>
	// <li>mpeg2: MPEG2 encoding</li>
	// <li>dnxhd: DNxHD encoding</li>
	// <li>mv-hevc: MV-HEVC encoding</li>
	// 
	// Note: AV1 encoding containers currently only support mp4, webm, and mkv.
	// Note: H.266 encoding containers currently only support mp4, hls, ts, and mov.
	// Note: VP8 and VP9 encoding containers currently only support webm and mkv.
	// Note: MPEG2 and DNxHD encoding containers currently only support mxf.
	// Note: MV-HEVC encoding containers only support mp4, hls, and mov. Among them, the hls format only supports mp4 segmentation format.
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// Video frame rate. Value range:
	// When FpsDenominator is empty, the range is [0, 120], in Hz.
	// When FpsDenominator is not empty, the Fps/FpsDenominator range is [0, 120].
	// If the value is 0, the frame rate will be the same as that of the source video.
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Bitrate of a video stream, in kbps. Value range: 0 and [128, 100000].If the value is 0, the bitrate of the video will be the same as that of the source video.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// Resolution adaption. Valid values:
	// <li>open: Enabled. When resolution adaption is enabled, `Width` indicates the long side of a video, while `Height` indicates the short side.</li>
	// <li>close: Disabled. When resolution adaption is disabled, `Width` indicates the width of a video, while `Height` indicates the height.</li>
	// Default value: open.
	// Note: When resolution adaption is enabled, `Width` cannot be smaller than `Height`.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Maximum value of the video stream width (or long edge) in px. Value range: 0 and [128, 4096].
	// <li>If both Width and Height are 0, the resolution is the same as the source.</li>
	// <li>If Width is 0 but Height is not 0, the width will be proportionally scaled.</li>
	// <li>If Width is not 0 but Height is 0, the height will be proportionally scaled.</li>
	// <li>If both Width and Height are not 0, the resolution is as specified by the user.</li>
	// Default value: 0.
	// Note: If Codec is set to MV-HEVC, the maximum value can be 7680.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Maximum value of the video stream height (or short edge) in px. Value range: 0 and [128, 4,096].
	// <li>If both Width and Height are 0, the resolution is the same as the source.</li>
	// <li>If Width is 0 but Height is not 0, the width will be proportionally scaled.</li>
	// <li>If Width is not 0 but Height is 0, the height will be proportionally scaled.</li>
	// <li>If both Width and Height are not 0, the resolution is as specified by the user.</li>
	// Default value: 0.
	// Note: If Codec is set to MV-HEVC, the maximum value can be 7680.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Interval between I-frames (keyframes), which can be customized in frames or seconds. GOP value range: 0 and [1, 100000].
	// If this parameter is 0 or left blank, the system will automatically set the GOP length.
	Gop *uint64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// GOP value unit. Optional values:
	// frame: indicates frame
	// second: indicates second
	// Default value: frame
	// Note: This field may return null, indicating that no valid value can be obtained.
	GopUnit *string `json:"GopUnit,omitnil,omitempty" name:"GopUnit"`

	// Padding method. When the video stream configuration width and height parameters are inconsistent with the aspect ratio of the original video, the transcoding processing method is "padding". Optional filling method:
	// <li> stretch: Stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot "shorter" or "longer";</li>
	// <li>black: Fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.</li>
	// <li>white: Fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks.</li>
	// <li>gauss: applies Gaussian blur to the uncovered area, without changing the image's aspect ratio.</li>
	// 
	// <li>smarttailor: Video images are smartly selected to ensure proportional image cropping.</li>
	// Default value: black.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// Specifies the constant bitrate control factor for the video. Value range: [0, 51]. Leaving this parameter blank sets it to "Automatic". It is recommended not to specify this parameter unless necessary.
	// If the Mode parameter is set to VBR and the Vcrf value is also configured, MPS will process the video in VBR mode, considering both Vcrf and Bitrate parameters to balance video quality, bitrate, transcoding efficiency, and file size.
	// If the Mode parameter is set to CRF, the Bitrate setting will be invalid, and encoding will be based on the Vcrf value.
	// If the Mode parameter is set to ABR or CBR, the Vcrf value does not need to be configured.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Vcrf *uint64 `json:"Vcrf,omitnil,omitempty" name:"Vcrf"`

	// Average segment duration. Value range: (0-10], unit: second.
	// This parameter will be set to automatic if not specified. The segment duration will be automatically selected based on the GOP and other characteristics of the video.
	// Note: It can be used only in the container format of hls.
	// Note: This field may return null, indicating that no valid value can be obtained.
	HlsTime *uint64 `json:"HlsTime,omitnil,omitempty" name:"HlsTime"`

	// HLS segment type. Valid values:
	// <li>0: HLS+TS segment</li>
	// <li>2: HLS+TS byte range</li>
	// <li>7: HLS+MP4 segment</li>
	// <li>5: HLS+MP4 byte range</li>
	// Default value: 0
	// 
	// Note: This field is used for normal/TSC transcoding settings and does not apply to adaptive bitrate streaming. To configure the segment type for adaptive bitrate streaming, use the outer field.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SegmentType *int64 `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`

	// Denominator of the frame rate.
	// Note: The value must be greater than 0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FpsDenominator *int64 `json:"FpsDenominator,omitnil,omitempty" name:"FpsDenominator"`

	// 3D video splicing mode, applicable only to mv-hevc and effective for 3d videos. valid values:
	// <Li>side_by_side: the original video content is arranged in a left-right layout.</li>
	// <li>top_bottom: vertical layout arrangement of original video content.</li>
	// Submit the amount and cost based on the segmented resolution size.
	// Default value: side_by_side.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Stereo3dType *string `json:"Stereo3dType,omitnil,omitempty" name:"Stereo3dType"`

	// Profile, suitable for different scenarios.
	// baseline: It only supports I/P-frames and non-interlaced scenarios, and is suitable for scenarios such as video calls and mobile videos.
	// main: It offers I-frames, P-frames, and B-frames, and supports both interlaced and non-interlaced modes. It is mainly used in mainstream audio and video consumption products such as video players and streaming media transmission devices.
	// high: the highest encoding level, with 8x8 prediction added to the main profile and support for custom quantification. It is widely used in scenarios such as Blu-ray storage and HDTV.
	// default: automatic filling along with the original video.    
	// 
	// This configuration appears only when the encoding standard is set to H264. baseline/main/high is supported. Default value: default
	// Note: This field may return null, indicating that no valid value can be obtained.
	VideoProfile *string `json:"VideoProfile,omitnil,omitempty" name:"VideoProfile"`

	// Encoder level. Default value: auto ("")
	// If the encoding standard is set to H264, the following options are supported: "", 1, 1.1, 1.2, 1.3, 2, 2.1, 2.2, 3, 3.1, 3.2, 4, 4.1, 4.2, 5, and 5.1.
	// If the encoding standard is set to H265, the following options are supported: "", 1, 2, 2.1, 3, 3.1, 4, 4.1, 5, 5.1, 5.2, 6, 6.1, 6.2, and 8.5.
	// Note: This field may return null, indicating that no valid value can be obtained.
	VideoLevel *string `json:"VideoLevel,omitnil,omitempty" name:"VideoLevel"`

	// Number of B-frames between reference frames. The default is auto, and a range of 0 - 16 is supported.
	// Note: Leaving it blank means using the auto option.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Bframes *int64 `json:"Bframes,omitnil,omitempty" name:"Bframes"`

	// Bitrate control mode. Optional values:
	// VBR: variable bitrate. The output bitrate is adjusted based on the complexity of the video image, ensuring higher image quality. This mode is suitable for storage scenarios as well as applications with high image quality requirements.
	// ABR: average bitrate. The average bitrate of the output video is kept stable to the greatest extent, but short-term bitrate fluctuations are allowed. This mode is suitable for scenarios where it is necessary to minimize the overall bitrate while a certain quality is maintained.
	// CBR: constant bitrate. The output bitrate remains constant during the video encoding process, regardless of changes in image complexity. This mode is suitable for scenarios with strict network bandwidth requirements, such as live streaming.
	// VCRF: constant rate factor. The video quality is controlled by setting a quality factor, achieving constant quality encoding of videos. The bitrate is automatically adjusted based on the complexity of the content. This mode is suitable for scenarios where maintaining a certain quality is desired.
	// VBR is selected by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Display aspect ratio. Optional values: [1:1, 2:1, default]
	// Default value: default
	// Note: This field may return null, indicating that no valid value can be obtained.
	Sar *string `json:"Sar,omitnil,omitempty" name:"Sar"`

	// Adaptive I-frame decision. When it is enabled, Media Processing Service will automatically identify transition points between different scenarios in the video (usually they are visually distinct frames, such as those of switching from one shot to another) and adaptively insert keyframes (I-frames) at these points to improve the random accessibility and encoding efficiency of the video. Optional values:
	// 0: Disable the adaptive I-frame decision 
	// 1: Enable the adaptive I-frame decision
	// Default value: 0
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	NoScenecut *int64 `json:"NoScenecut,omitnil,omitempty" name:"NoScenecut"`

	// Bit: 8/10 is supported. Default value: 8
	// Note: This field may return null, indicating that no valid value can be obtained.
	BitDepth *int64 `json:"BitDepth,omitnil,omitempty" name:"BitDepth"`

	// Preservation of original timestamp. Optional values:
	// 0: Disabled
	// 1: Enabled
	// Default value: Disabled
	// Note: This field may return null, indicating that no valid value can be obtained.
	RawPts *int64 `json:"RawPts,omitnil,omitempty" name:"RawPts"`

	// Proportional compression bitrate. When it is enabled, the bitrate of the output video will be adjusted according to the proportion. After the compression ratio is entered, the system will automatically calculate the target output bitrate based on the source video bitrate. Compression ratio range: 0-100
	// Leaving this value blank means it is not enabled by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Compress *int64 `json:"Compress,omitnil,omitempty" name:"Compress"`

	// Segment duration at startup.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SegmentSpecificInfo *SegmentSpecificInfo `json:"SegmentSpecificInfo,omitnil,omitempty" name:"SegmentSpecificInfo"`

	// Whether the template enables scenario-based settings. 
	// 0: disable. 
	// 1: enable 
	//  
	// Default value: 0	
	// 	
	// Note: The values of SceneType and CompressType fields only take effect when this field value is 1.
	// Note: This field may return null, indicating that no valid value can be obtained.
	ScenarioBased *int64 `json:"ScenarioBased,omitnil,omitempty" name:"ScenarioBased"`

	// Video scenario. Valid values: 
	// normal: General transcoding scenario: General transcoding and compression scenario.
	// pgc: PGC HD TV shows and movies: At the time of compression, focus is placed on the viewing experience of TV shows and movies and ROI encoding is performed according to their characteristics, while high-quality contents of videos and audio are retained. 
	// materials_video: HD materials: Scenario involving material resources, where requirements for image quality are extremely high and there are many transparent images, with almost no visual loss during compression. 
	// ugc: UGC content: It is suitable for a wide range of UGC/short video scenarios, with an optimized encoding bitrate for short video characteristics, improved image quality, and enhanced business QOS/QOE metrics. 
	// e-commerce_video: Fashion show/e-commerce: At the time of compression, emphasis is placed on detail clarity and ROI enhancement, with a particular focus on maintaining the image quality of the face region. 
	// educational_video: Education: At the time of compression, emphasis is placed on the clarity and readability of text and images to help students better understand the content, ensuring that the teaching content is clearly conveyed. 
	// Default value: normal.
	// Note: To use this value, the value of ScenarioBased must be 1; otherwise, this value will not take effect.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Transcoding policy. Valid values: 
	// ultra_compress: Extreme compression: Compared to standard compression, this policy can maximize bitrate compression while ensuring a certain level of image quality, thus greatly saving bandwidth and storage costs. 
	// standard_compress: Comprehensively optimal: Balances compression ratio and image quality, compressing files as much as possible without a noticeable reduction in subjective image quality. Only audio and video TSC transcoding fees are charged for this policy. 
	// high_compress: Bitrate priority: Prioritizes reducing file size, which may result in certain image quality loss. Only audio and video TSC transcoding fees are charged for this policy. 
	// low_compress: Image quality priority: Prioritizes ensuring image quality, and the size of compressed files may be relatively large. Only audio and video TSC transcoding fees are charged for this policy. 
	// Default value: standard_compress. 
	// Note: If you need to watch videos on TV, it is recommended not to use the ultra_compress policy. The billing standard for the ultra_compress policy is TSC transcoding + audio and video enhancement - artifacts removal.
	// Note: To use this value, the value of ScenarioBased must be 1; otherwise, this value will not take effect.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
}

type VideoTemplateInfoForUpdate struct {
	// Encoding format for video streams. Optional values:
	// <li>h264: H.264 encoding</li>
	// <li>h265: H.265 encoding</li>
	// <li>h266: H.266 encoding</li>
	// <li>av1: AOMedia Video 1 encoding</li>
	// <li>vp8: VP8 encoding</li>
	// <li>vp9: VP9 encoding</li>
	// <li>mpeg2: MPEG2 encoding</li>
	// <li>dnxhd: DNxHD encoding</li>
	// <li>mv-hevc: MV-HEVC encoding</li>
	// 
	// Note: 
	// AV1 encoding containers currently only support mp4, webm, and mkv.
	// H.266 encoding containers currently only support mp4, hls, ts, and mov. 
	// VP8 and VP9 encoding containers currently only support webm and mkv.
	// MPEG2 and DNxHD encoding containers currently only support mxf.
	// MV-HEVC encoding containers only support mp4, hls, and mov. Also, the hls format only supports mp4 segmentation format.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// Video frame rate. Value range:
	// When FpsDenominator is empty, the range is [0, 120], in Hz.
	// When FpsDenominator is not empty, the Fps/FpsDenominator range is [0, 120].
	// If the value is 0, the frame rate will be the same as that of the source video.Note: This field may return null, indicating that no valid values can be obtained.
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Bitrate of a video stream, in kbps. Value range: 0 and [128, 100000].If the value is 0, the bitrate of the video will be the same as that of the source video.Note: This field may return null, indicating that no valid values can be obtained.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// Resolution adaption. Valid values:
	// <li>open: Enabled. When resolution adaption is enabled, `Width` indicates the long side of a video, while `Height` indicates the short side.</li>
	// <li>close: Disabled. When resolution adaption is disabled, `Width` indicates the width of a video, while `Height` indicates the height.</li>
	// Note: When resolution adaption is enabled, `Width` cannot be smaller than `Height`.
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitnil,omitempty" name:"ResolutionAdaptive"`

	// Maximum value of the video stream width (or long edge) in px. Value range: 0 and [128, 4096].
	// <li>If both Width and Height are 0, the resolution is the same as the source.</li>
	// <li>If Width is 0 but Height is not 0, the width will be proportionally scaled.</li>
	// <li>If Width is not 0 but Height is 0, the height will be proportionally scaled.</li>
	// <li>If both Width and Height are not 0, the resolution is as specified by the user.</li>
	// Note: If Codec is set to MV-HEVC, the maximum value can be 7680.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Maximum value of the video stream height (or short edge) in px. Value range: 0 and [128, 4,096].
	// Note: If Codec is set to MV-HEVC, the maximum value can be 7680.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// Interval between I-frames (keyframes), which can be customized in frames or seconds. GOP value range: 0 and [1, 100000].
	// If this parameter is 0, the system will automatically set the GOP length.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Gop *uint64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// GOP value unit. Optional values: 
	// frame: indicates frame 
	// second: indicates second
	// Default value: frame
	// Note: This field may return null, indicating that no valid value can be obtained.
	GopUnit *string `json:"GopUnit,omitnil,omitempty" name:"GopUnit"`

	// Fill type. "Fill" refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. Valid values:
	//  <li>stretch: Each frame is stretched to fill the entire screen, which may cause the transcoded video to be "flattened" or "stretched".</li>
	// <li>black: Keep the image's original aspect ratio and fill the blank space with black bars.</li>
	// <li>white: The aspect ratio of the video is kept unchanged, and the rest of the edges is filled with white.</li>
	// <li>gauss: applies Gaussian blur to the uncovered area, without changing the image's aspect ratio.</li>
	// 
	// <li>smarttailor: Video images are smartly selected to ensure proportional image cropping.</li>
	// Default value: black.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	FillType *string `json:"FillType,omitnil,omitempty" name:"FillType"`

	// The control factor of video constant bitrate. Value range: [0, 51]. If not specified, it means "auto". It is recommended not to specify this parameter unless necessary.
	// When the Mode parameter is set to VBR, if the Vcrf value is also configured, MPS will process the video in VBR mode, considering both Vcrf and Bitrate parameters to balance video quality, bitrate, transcoding efficiency, and file size.
	// When the Mode parameter is set to CRF, the Bitrate setting will be invalid, and the encoding will be based on the Vcrf value.
	// When the Mode parameter is set to ABR or CBR, the Vcrf value does not need to be configured.
	// Note: When you need to set it to auto, fill in 100.
	// 
	// Note: This field may return null, indicating that no valid value can be obtained.
	Vcrf *uint64 `json:"Vcrf,omitnil,omitempty" name:"Vcrf"`

	// Whether to enable adaptive encoding. Valid values:
	// <li>0: Disable</li>
	// <li>1: Enable</li>
	// Default value: 0. If this parameter is set to `1`, multiple streams with different resolutions and bitrates will be generated automatically. The highest resolution, bitrate, and quality of the streams are determined by the values of `width` and `height`, `Bitrate`, and `Vcrf` in `VideoTemplate` respectively. If these parameters are not set in `VideoTemplate`, the highest resolution generated will be the same as that of the source video, and the highest video quality will be close to VMAF 95. To use this parameter or learn about the billing details of adaptive encoding, please contact your sales rep.
	ContentAdaptStream *uint64 `json:"ContentAdaptStream,omitnil,omitempty" name:"ContentAdaptStream"`

	// Average segment duration. Value range: (0-10], unit: second
	// Default value: 10
	// Note: It is used only in the format of HLS.
	// Note: This field may return null, indicating that no valid value can be obtained.
	HlsTime *uint64 `json:"HlsTime,omitnil,omitempty" name:"HlsTime"`

	// HLS segment type. Valid values:
	// <li>0: HLS+TS segment</li>
	// <li>2: HLS+TS byte range</li>
	// <li>7: HLS+MP4 segment</li>
	// <li>5: HLS+MP4 byte range</li>
	// Default value: 0
	// 
	// Note: This field is used for normal/Top Speed Codec transcoding settings and does not apply to adaptive bitrate streaming. To configure the segment type for adaptive bitrate streaming, use the outer field.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SegmentType *int64 `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`

	// Denominator of the frame rate.
	// Note: The value must be greater than 0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FpsDenominator *int64 `json:"FpsDenominator,omitnil,omitempty" name:"FpsDenominator"`

	// 3D video splicing mode, applicable only to mv-hevc and effective for 3d videos. valid values:
	// <Li>side_by_side: the original video content is arranged in a left-right layout.</li>
	// <Li>top_bottom: layout arrangement of the original video content from top to bottom.</li>
	// The usage and charges will be reported based on the segmented resolution dimensions.
	// Default value: side_by_side.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Stereo3dType *string `json:"Stereo3dType,omitnil,omitempty" name:"Stereo3dType"`

	// Profile, suitable for different scenarios. 
	// baseline: It only supports I/P-frames and non-interlaced scenarios, and is suitable for scenarios such as video calls and mobile videos. 
	// main: It offers I-frames, P-frames, and B-frames, and supports both interlaced and non-interlaced modes. It is mainly used in mainstream audio and video consumption products such as video players and streaming media transmission devices. 
	// high: the highest encoding level, with 8x8 prediction added to the main profile and support for custom quantification. It is widely used in scenarios such as Blu-ray storage and HDTV.
	// default: automatic filling along with the original video
	// 
	// This configuration appears only when the encoding standard is set to H264. Default value: default
	// Note: This field may return null, indicating that no valid value can be obtained.
	VideoProfile *string `json:"VideoProfile,omitnil,omitempty" name:"VideoProfile"`

	// Encoder level. Default value: auto ("")
	// If the encoding standard is set to H264, the following options are supported: "", 1, 1.1, 1.2, 1.3, 2, 2.1, 2.2, 3, 3.1, 3.2, 4, 4.1, 4.2, 5, and 5.1. 
	// If the encoding standard is set to H265, the following options are supported: "", 1, 2, 2.1, 3, 3.1, 4, 4.1, 5, 5.1, 5.2, 6, 6.1, 6.2, and 8.5.
	// Note: This field may return null, indicating that no valid value can be obtained.
	VideoLevel *string `json:"VideoLevel,omitnil,omitempty" name:"VideoLevel"`

	// Maximum number of consecutive B-frames. The default is auto, and 0 - 16 and -1 are supported.
	// Note:
	// 
	// -1 indicates auto.	
	// Note: This field may return null, indicating that no valid value can be obtained.
	Bframes *int64 `json:"Bframes,omitnil,omitempty" name:"Bframes"`

	// Bitrate control mode. Optional values: 
	// VBR: variable bitrate. The output bitrate is adjusted based on the complexity of the video image, ensuring higher image quality. This mode is suitable for storage scenarios as well as applications with high image quality requirements. 
	// ABR: average bitrate. The average bitrate of the output video is kept stable to the greatest extent, but short-term bitrate fluctuations are allowed. This mode is suitable for scenarios where it is necessary to minimize the overall bitrate while a certain quality is maintained. 
	// CBR: constant bitrate. The output bitrate remains constant during the video encoding process, regardless of changes in image complexity. This mode is suitable for scenarios with strict network bandwidth requirements, such as live streaming. 
	// VCRF: constant rate factor. The video quality is controlled by setting a quality factor, achieving constant quality encoding of videos. The bitrate is automatically adjusted based on the complexity of the content. This mode is suitable for scenarios where maintaining a certain quality is desired. 
	// VBR is selected by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Display aspect ratio. Optional values: [1:1, 2:1, default]
	// Default value: default
	// Note: This field may return null, indicating that no valid value can be obtained.
	Sar *string `json:"Sar,omitnil,omitempty" name:"Sar"`

	// Adaptive I-frame decision. When it is enabled, Media Processing Service will automatically identify transition points between different scenarios in the video (usually they are visually distinct frames, such as those of switching from one shot to another) and adaptively insert keyframes (I-frames) at these points to improve the random accessibility and encoding efficiency of the video. Optional values: 
	// 0: Disable the adaptive I-frame decision 
	// 1: Enable the adaptive I-frame decision 
	// Default value: 0	
	// 	
	// Note: This field may return null, indicating that no valid value can be obtained.
	NoScenecut *int64 `json:"NoScenecut,omitnil,omitempty" name:"NoScenecut"`

	// Bit: 8/10 is supported. Default value: 8	
	// Note: This field may return null, indicating that no valid value can be obtained.
	BitDepth *int64 `json:"BitDepth,omitnil,omitempty" name:"BitDepth"`

	// Preservation of original timestamp. Optional values: 
	// 0: Disabled 
	// 1: Enabled 
	// Default value: Disabled	
	// Note: This field may return null, indicating that no valid value can be obtained.
	RawPts *int64 `json:"RawPts,omitnil,omitempty" name:"RawPts"`

	// Proportional compression bitrate. When it is enabled, the bitrate of the output video will be adjusted according to the proportion. After the compression ratio is entered, the system will automatically calculate the target output bitrate based on the source video bitrate. Compression ratio range: 0-100, optional values: [0-100] and -1 
	// Note: -1 indicates auto.	
	// Note: This field may return null, indicating that no valid value can be obtained.
	Compress *int64 `json:"Compress,omitnil,omitempty" name:"Compress"`

	// Segment duration at startup.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SegmentSpecificInfo *SegmentSpecificInfo `json:"SegmentSpecificInfo,omitnil,omitempty" name:"SegmentSpecificInfo"`

	// Indicates whether to enable scenario-based settings for the template. 
	// 0: Disable. 
	// 1: enable 
	//  
	// Default value: 0	
	// 	
	// Note: This value takes effect only when the value of this field is 1.
	// Note: This field may return null, indicating that no valid value can be obtained.
	ScenarioBased *int64 `json:"ScenarioBased,omitnil,omitempty" name:"ScenarioBased"`

	// Video scenario. Valid values: 
	// normal: General transcoding scenario: General transcoding and compression scenario. pgc: PGC HD film and television: Emphasis is placed on the viewing experience of films and TV shows during compression, with ROI encoding based on the characteristics of films and TV shows, while maintaining high-quality video and audio content. 
	// materials_video: HD materials: Scenario involving material resources, where requirements for image quality are extremely high and there are many transparent images, with almost no visual loss during compression. 
	// ugc: UGC content: It is suitable for a wide range of UGC/short video scenarios, with an optimized encoding bitrate for short video characteristics, improved image quality, and enhanced business QOS/QOE metrics. 
	// e-commerce_video: Fashion show/e-commerce: At the time of compression, emphasis is placed on detail clarity and ROI enhancement, with a particular focus on maintaining the image quality of the face region. 
	// educational_video: Education: At the time of compression, emphasis is placed on the clarity and readability of text and images to help students better understand the content, ensuring that the teaching content is clearly conveyed.
	// Default value: normal.
	// Note: To use this value, the value of ScenarioBased must be 1; otherwise, this value will not take effect.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Transcoding policy. Valid values: 
	// ultra_compress: Extreme compression: Compared to standard compression, this policy can maximize bitrate compression while ensuring a certain level of image quality, thus greatly saving bandwidth and storage costs. 
	// standard_compress: Comprehensively optimal: Balances compression ratio and image quality, compressing files as much as possible without a noticeable reduction in subjective image quality. This policy only charges audio and video TSC transcoding fees. 
	// high_compress: Bitrate priority: Prioritizes reducing file size, which may result in some image quality loss. This policy only charges audio and video TSC transcoding fees. 
	// low_compress: Image quality priority: Prioritizes ensuring image quality, and the size of compressed files may be relatively large. This policy only charges audio and video TSC transcoding fees. 
	// Default value: standard_compress. 
	// Note: If you need to watch videos on TV, it is recommended not to use the ultra_compress policy. The billing standard for the ultra_compress policy is TSC transcoding + audio and video enhancement - artifacts removal.
	// Note: To use this value, the value of ScenarioBased must be 1; otherwise, this value will not take effect.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
}

type VolumeBalanceConfig struct {
	// Whether to enable the feature. Valid values:
	// <li>`ON`</li>
	// <li>`OFF` </li>
	// Default value: `ON`.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The type. Valid values:
	// <li>`loudNorm`: Loudness normalization.</li>
	// <li>`gainControl`: Volume leveling.</li>
	// Default value: `loudNorm`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type WatermarkInput struct {
	// ID of a watermarking template.
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Custom watermark parameter, which is valid if `Definition` is 0.
	// This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the watermark parameter preferably.
	// Custom watermark parameter is not available for screenshot.
	RawParameter *RawWatermarkParameter `json:"RawParameter,omitnil,omitempty" name:"RawParameter"`

	// Text content of up to 100 characters. This field is required only when the watermark type is text.
	// Text watermark is not available for screenshot.
	TextContent *string `json:"TextContent,omitnil,omitempty" name:"TextContent"`

	// SVG content of up to 2,000,000 characters. This field is required only when the watermark type is `SVG`.
	// SVG watermark is not available for screenshot.
	SvgContent *string `json:"SvgContent,omitnil,omitempty" name:"SvgContent"`

	// Start time offset of a watermark, in seconds. If not set or set to 0, a watermark starts appearing when a video starts.
	// <li>If not set or set to 0, a watermark starts appearing when a video starts.</li>
	// <li>If the value is greater than 0 (for example, n), a watermark will appear at second n of a video.</li>
	// <li>If the value is less than 0 (for example, -n), a watermark will appear n seconds before the end of a video.</li>
	// 
	// Note: It is only used for video scenarios. Screenshots are not supported.
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// End time offset of a watermark, in seconds.
	// <li>If not set or set to 0, a watermark will last until the end of a video.</li>
	// <li>If the value is greater than 0 (for example, n), a watermark will disappear at second n.</li>
	// <li>If the value is less than 0 (for example, -n), a watermark will disappear n seconds before the end of a video.</li>
	// 
	// Note: It is only used for video scenarios. Screenshots are not supported.
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type WatermarkTemplate struct {
	// Unique ID of a watermarking template.
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// Watermark type. Valid values:
	// <li>image: Image watermark;</li>
	// <li>text: Text watermark.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Name of a watermarking template.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// Horizontal position of the origin of the watermark image relative to the origin of the video.
	// <li>If the string ends in %, the `Left` edge of the watermark will be at the position of the specified percentage of the video width; for example, `10%` means that the `Left` edge is at 10% of the video width;</li>
	// <li>If the string ends in px, the `Left` edge of the watermark will be at the position of the specified px of the video width; for example, `100px` means that the `Left` edge is at the position of 100 px.</li>
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// Vertical position of the origin of the watermark image relative to the origin of the video.
	// <li>If the string ends in %, the `Top` edge of the watermark will beat the position of the specified percentage of the video height; for example, `10%` means that the `Top` edge is at 10% of the video height;</li>
	// <li>If the string ends in px, the `Top` edge of the watermark will be at the position of the specified px of the video height; for example, `100px` means that the `Top` edge is at the position of 100 px.</li>
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// Image watermarking template. This field is valid only when `Type` is `image`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImageTemplate *ImageWatermarkTemplate `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`

	// Text watermarking template. This field is valid only when `Type` is `text`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitnil,omitempty" name:"TextTemplate"`

	// SVG watermarking template. This field is valid when `Type` is `svg`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitnil,omitempty" name:"SvgTemplate"`

	// Creation time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a template in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Origin position. Valid values:
	// <li>topLeft: The origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text;</li>
	// <li>topRight: The origin of coordinates is in the top-right corner of the video, and the origin of the watermark is in the top-right corner of the image or text;</li>
	// <li>bottomLeft: The origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left corner of the image or text;</li>
	// <li>bottomRight: The origin of coordinates is in the bottom-right corner of the video, and the origin of the watermark is in the bottom-right corner of the image or text.</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`
}

type WordResult struct {
	// Word text.
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// Word start timestamp, in seconds.
	Start *float64 `json:"Start,omitnil,omitempty" name:"Start"`

	// Word end timestamp, in seconds.
	End *float64 `json:"End,omitnil,omitempty" name:"End"`

	// Text after translation.
	Trans *string `json:"Trans,omitnil,omitempty" name:"Trans"`
}

type WorkflowInfo struct {
	// Workflow ID.
	WorkflowId *int64 `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Workflow status. Valid values:
	// <li>Enabled: Enabled,</li>
	// <li>Disabled: Disabled.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Input rule bound to a workflow. If an uploaded video hits the rule for the object, the workflow will be triggered.
	Trigger *WorkflowTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The location to save the media processing output file.
	// Note: This field may return null, indicating that no valid value can be obtained.
	OutputStorage *TaskOutputStorage `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// The media processing parameters to use.
	// Note: This field may return null, indicating that no valid value can be obtained.
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// Type parameter of a video content audit task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitnil,omitempty" name:"AiContentReviewTask"`

	// Video content analysis task parameter.
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitnil,omitempty" name:"AiAnalysisTask"`

	// Type parameter of a video content recognition task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitnil,omitempty" name:"AiRecognitionTask"`

	// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskNotifyConfig *TaskNotifyConfig `json:"TaskNotifyConfig,omitnil,omitempty" name:"TaskNotifyConfig"`

	// Task flow priority. The higher the value, the higher the priority. Value range: [-10, 10]. If this parameter is left empty, 0 will be used.
	TaskPriority *int64 `json:"TaskPriority,omitnil,omitempty" name:"TaskPriority"`

	// The directory to save the media processing output file, such as `/movie/201907/`.
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// Creation time of a workflow in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of a workflow in [ISO date format](https://intl.cloud.tencent.com/document/product/266/11732?from_cn_redirect=1#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type WorkflowTask struct {
	// The media processing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task flow status. Valid values:
	// <li>PROCESSING: Processing;</li>
	// <li>FINISH: Completed.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// If the value returned is not 0, there was a source error. If 0 is returned, refer to the error codes of the corresponding task type.
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Except those for source errors, error messages vary with task type.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The information of the file processed.
	// Note: This field may return null, indicating that no valid value can be obtained.
	InputInfo *MediaInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// Metadata of a source video.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetaData *MediaMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// The execution status and result of the media processing task.
	MediaProcessResultSet []*MediaProcessTaskResult `json:"MediaProcessResultSet,omitnil,omitempty" name:"MediaProcessResultSet"`

	// Execution status and result of a video content audit task.
	AiContentReviewResultSet []*AiContentReviewResult `json:"AiContentReviewResultSet,omitnil,omitempty" name:"AiContentReviewResultSet"`

	// Execution status and result of video content analysis task.
	AiAnalysisResultSet []*AiAnalysisResult `json:"AiAnalysisResultSet,omitnil,omitempty" name:"AiAnalysisResultSet"`

	// Execution status and result of a video content recognition task.
	AiRecognitionResultSet []*AiRecognitionResult `json:"AiRecognitionResultSet,omitnil,omitempty" name:"AiRecognitionResultSet"`

	// Execution status and results of a media quality inspection task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AiQualityControlTaskResult *ScheduleQualityControlTaskResult `json:"AiQualityControlTaskResult,omitnil,omitempty" name:"AiQualityControlTaskResult"`

	// Execution result of the smart subtitle task.
	// Note: This field may return null, indicating that no valid value can be obtained.
	SmartSubtitlesTaskResult []*SmartSubtitlesResult `json:"SmartSubtitlesTaskResult,omitnil,omitempty" name:"SmartSubtitlesTaskResult"`
}

type WorkflowTrigger struct {
	// The trigger type. Valid values:
	// <li>`CosFileUpload`: Tencent Cloud COS trigger.</li>
	// <li>`AwsS3FileUpload`: AWS S3 trigger. Currently, this type is only supported for transcoding tasks and schemes (not supported for workflows).</li>
	// 
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// This parameter is required and valid when `Type` is `CosFileUpload`, indicating the COS trigger rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosFileUploadTrigger *CosFileUploadTrigger `json:"CosFileUploadTrigger,omitnil,omitempty" name:"CosFileUploadTrigger"`

	// The AWS S3 trigger. This parameter is valid and required if `Type` is `AwsS3FileUpload`.
	// 
	// Note: Currently, the key for the AWS S3 bucket, the trigger SQS queue, and the callback SQS queue must be the same.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AwsS3FileUploadTrigger *AwsS3FileUploadTrigger `json:"AwsS3FileUploadTrigger,omitnil,omitempty" name:"AwsS3FileUploadTrigger"`
}