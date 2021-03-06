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
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-06-12"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateAIAnalysisTemplateRequest() (request *CreateAIAnalysisTemplateRequest) {
    request = &CreateAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateAIAnalysisTemplate")
    return
}

func NewCreateAIAnalysisTemplateResponse() (response *CreateAIAnalysisTemplateResponse) {
    response = &CreateAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAIAnalysisTemplate
// This API is used to create a custom content analysis template. Up to 50 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLASSIFCATIONCONFIGURE = "InvalidParameterValue.ClassifcationConfigure"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COVERCONFIGURE = "InvalidParameterValue.CoverConfigure"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_FRAMETAGCONFIGURE = "InvalidParameterValue.FrameTagConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAIAnalysisTemplate(request *CreateAIAnalysisTemplateRequest) (response *CreateAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIAnalysisTemplateRequest()
    }
    response = NewCreateAIAnalysisTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIRecognitionTemplateRequest() (request *CreateAIRecognitionTemplateRequest) {
    request = &CreateAIRecognitionTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateAIRecognitionTemplate")
    return
}

func NewCreateAIRecognitionTemplateResponse() (response *CreateAIRecognitionTemplateResponse) {
    response = &CreateAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAIRecognitionTemplate
// This API is used to create a custom content recognition template. Up to 50 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAIRecognitionTemplate(request *CreateAIRecognitionTemplateRequest) (response *CreateAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIRecognitionTemplateRequest()
    }
    response = NewCreateAIRecognitionTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAdaptiveDynamicStreamingTemplateRequest() (request *CreateAdaptiveDynamicStreamingTemplateRequest) {
    request = &CreateAdaptiveDynamicStreamingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateAdaptiveDynamicStreamingTemplate")
    return
}

func NewCreateAdaptiveDynamicStreamingTemplateResponse() (response *CreateAdaptiveDynamicStreamingTemplateResponse) {
    response = &CreateAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAdaptiveDynamicStreamingTemplate
// This API is used to create an adaptive bitrate streaming template. Up up to 100 such templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAdaptiveDynamicStreamingTemplate(request *CreateAdaptiveDynamicStreamingTemplateRequest) (response *CreateAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAdaptiveDynamicStreamingTemplateRequest()
    }
    response = NewCreateAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAnimatedGraphicsTemplateRequest() (request *CreateAnimatedGraphicsTemplateRequest) {
    request = &CreateAnimatedGraphicsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateAnimatedGraphicsTemplate")
    return
}

func NewCreateAnimatedGraphicsTemplateResponse() (response *CreateAnimatedGraphicsTemplateResponse) {
    response = &CreateAnimatedGraphicsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAnimatedGraphicsTemplate
// This API is used to create a custom animated image generating template. Up to 16 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FORMATWEBPLACKWIDTHANDHEIGHT = "InvalidParameterValue.FormatWebpLackWidthAndHeight"
//  INVALIDPARAMETERVALUE_FORMATWEBPWIDTHANDHEIGHTBOTHZERO = "InvalidParameterValue.FormatWebpWidthAndHeightBothZero"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_QUALITY = "InvalidParameterValue.Quality"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAnimatedGraphicsTemplate(request *CreateAnimatedGraphicsTemplateRequest) (response *CreateAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAnimatedGraphicsTemplateRequest()
    }
    response = NewCreateAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContentReviewTemplateRequest() (request *CreateContentReviewTemplateRequest) {
    request = &CreateContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateContentReviewTemplate")
    return
}

func NewCreateContentReviewTemplateResponse() (response *CreateContentReviewTemplateResponse) {
    response = &CreateContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateContentReviewTemplate
// This API is used to create a custom template for intelligent content recognition. Up to 50 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateContentReviewTemplate(request *CreateContentReviewTemplateRequest) (response *CreateContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewCreateContentReviewTemplateRequest()
    }
    response = NewCreateContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageSpriteTemplateRequest() (request *CreateImageSpriteTemplateRequest) {
    request = &CreateImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateImageSpriteTemplate")
    return
}

func NewCreateImageSpriteTemplateResponse() (response *CreateImageSpriteTemplateResponse) {
    response = &CreateImageSpriteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImageSpriteTemplate
// This API is used to create a custom image sprite generating template. Up to 16 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateImageSpriteTemplate(request *CreateImageSpriteTemplateRequest) (response *CreateImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewCreateImageSpriteTemplateRequest()
    }
    response = NewCreateImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonSampleRequest() (request *CreatePersonSampleRequest) {
    request = &CreatePersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreatePersonSample")
    return
}

func NewCreatePersonSampleResponse() (response *CreatePersonSampleResponse) {
    response = &CreatePersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePersonSample
// This API is used to create image samples for video processing operations such as content recognition and inappropriate information detection with the help of technologies such as facial feature positioning.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
func (c *Client) CreatePersonSample(request *CreatePersonSampleRequest) (response *CreatePersonSampleResponse, err error) {
    if request == nil {
        request = NewCreatePersonSampleRequest()
    }
    response = NewCreatePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSampleSnapshotTemplateRequest() (request *CreateSampleSnapshotTemplateRequest) {
    request = &CreateSampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateSampleSnapshotTemplate")
    return
}

func NewCreateSampleSnapshotTemplateResponse() (response *CreateSampleSnapshotTemplateResponse) {
    response = &CreateSampleSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSampleSnapshotTemplate
// This API is used to create a custom sampled screencapturing template. Up to 16 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateSampleSnapshotTemplate(request *CreateSampleSnapshotTemplateRequest) (response *CreateSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSampleSnapshotTemplateRequest()
    }
    response = NewCreateSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotByTimeOffsetTemplateRequest() (request *CreateSnapshotByTimeOffsetTemplateRequest) {
    request = &CreateSnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateSnapshotByTimeOffsetTemplate")
    return
}

func NewCreateSnapshotByTimeOffsetTemplateResponse() (response *CreateSnapshotByTimeOffsetTemplateResponse) {
    response = &CreateSnapshotByTimeOffsetTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSnapshotByTimeOffsetTemplate
// This API is used to create a custom time point screencapturing template. Up to 16 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateSnapshotByTimeOffsetTemplate(request *CreateSnapshotByTimeOffsetTemplateRequest) (response *CreateSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotByTimeOffsetTemplateRequest()
    }
    response = NewCreateSnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTranscodeTemplateRequest() (request *CreateTranscodeTemplateRequest) {
    request = &CreateTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateTranscodeTemplate")
    return
}

func NewCreateTranscodeTemplateResponse() (response *CreateTranscodeTemplateResponse) {
    response = &CreateTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTranscodeTemplate
// This API is used to create a custom transcoding template. Up to 1,000 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateTranscodeTemplate(request *CreateTranscodeTemplateRequest) (response *CreateTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTranscodeTemplateRequest()
    }
    response = NewCreateTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWatermarkTemplateRequest() (request *CreateWatermarkTemplateRequest) {
    request = &CreateWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateWatermarkTemplate")
    return
}

func NewCreateWatermarkTemplateResponse() (response *CreateWatermarkTemplateResponse) {
    response = &CreateWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWatermarkTemplate
// This API is used to create a custom watermarking template. Up to 1,000 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INTERNALERROR_UPLOADWATERMARKERROR = "InternalError.UploadWatermarkError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COORDINATEORIGIN = "InvalidParameterValue.CoordinateOrigin"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_IMAGECONTENT = "InvalidParameterValue.ImageContent"
//  INVALIDPARAMETERVALUE_IMAGETEMPLATE = "InvalidParameterValue.ImageTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REPEATTYPE = "InvalidParameterValue.RepeatType"
//  INVALIDPARAMETERVALUE_SVGTEMPLATE = "InvalidParameterValue.SvgTemplate"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEHEIGHT = "InvalidParameterValue.SvgTemplateHeight"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEWIDTH = "InvalidParameterValue.SvgTemplateWidth"
//  INVALIDPARAMETERVALUE_TEXTALPHA = "InvalidParameterValue.TextAlpha"
//  INVALIDPARAMETERVALUE_TEXTTEMPLATE = "InvalidParameterValue.TextTemplate"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  INVALIDPARAMETERVALUE_XPOS = "InvalidParameterValue.XPos"
//  INVALIDPARAMETERVALUE_YPOS = "InvalidParameterValue.YPos"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateWatermarkTemplate(request *CreateWatermarkTemplateRequest) (response *CreateWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewCreateWatermarkTemplateRequest()
    }
    response = NewCreateWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWordSamplesRequest() (request *CreateWordSamplesRequest) {
    request = &CreateWordSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateWordSamples")
    return
}

func NewCreateWordSamplesResponse() (response *CreateWordSamplesResponse) {
    response = &CreateWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWordSamples
// This API is used to create keyword samples in batches for video processing operations such as content recognition and inappropriate information detection with the help of the OCR and ASR technologies.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateWordSamples(request *CreateWordSamplesRequest) (response *CreateWordSamplesResponse, err error) {
    if request == nil {
        request = NewCreateWordSamplesRequest()
    }
    response = NewCreateWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRequest() (request *CreateWorkflowRequest) {
    request = &CreateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "CreateWorkflow")
    return
}

func NewCreateWorkflowResponse() (response *CreateWorkflowResponse) {
    response = &CreateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWorkflow
// This API is used to set a processing rule for media files uploaded to the specified directory of a COS bucket, including:
//
// 1. Video transcoding (with watermark);
//
// 2. Animated image generating;
//
// 3. Time point screencapturing;
//
// 4. Sampled screencapturing;
//
// 5. Image sprite generating;
//
// 6. Video conversion to adaptive bitrate streaming;
//
// 7. Intelligent content audit (detection of porn, terrorism, and politically sensitive information);
//
// 8. Intelligent content recognition (face recognition, full text recognition, text keyword recognition, full speech recognition, and speech keyword recognition).
//
// 
//
// Note: Once successfully created, a workflow is disabled by default and needs to be enabled manually.
//
// error code that may be returned:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
func (c *Client) CreateWorkflow(request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRequest()
    }
    response = NewCreateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIAnalysisTemplateRequest() (request *DeleteAIAnalysisTemplateRequest) {
    request = &DeleteAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAIAnalysisTemplate")
    return
}

func NewDeleteAIAnalysisTemplateResponse() (response *DeleteAIAnalysisTemplateResponse) {
    response = &DeleteAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAIAnalysisTemplate
// This API is used to delete a custom content analysis template.
//
// 
//
// Note: templates with an ID below 10000 are preset and cannot be deleted.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAIAnalysisTemplate(request *DeleteAIAnalysisTemplateRequest) (response *DeleteAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIAnalysisTemplateRequest()
    }
    response = NewDeleteAIAnalysisTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIRecognitionTemplateRequest() (request *DeleteAIRecognitionTemplateRequest) {
    request = &DeleteAIRecognitionTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAIRecognitionTemplate")
    return
}

func NewDeleteAIRecognitionTemplateResponse() (response *DeleteAIRecognitionTemplateResponse) {
    response = &DeleteAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAIRecognitionTemplate
// This API is used to delete a custom content recognition template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAIRecognitionTemplate(request *DeleteAIRecognitionTemplateRequest) (response *DeleteAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIRecognitionTemplateRequest()
    }
    response = NewDeleteAIRecognitionTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAdaptiveDynamicStreamingTemplateRequest() (request *DeleteAdaptiveDynamicStreamingTemplateRequest) {
    request = &DeleteAdaptiveDynamicStreamingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAdaptiveDynamicStreamingTemplate")
    return
}

func NewDeleteAdaptiveDynamicStreamingTemplateResponse() (response *DeleteAdaptiveDynamicStreamingTemplateResponse) {
    response = &DeleteAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAdaptiveDynamicStreamingTemplate
// This API is used to delete an adaptive bitrate streaming template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAdaptiveDynamicStreamingTemplate(request *DeleteAdaptiveDynamicStreamingTemplateRequest) (response *DeleteAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAdaptiveDynamicStreamingTemplateRequest()
    }
    response = NewDeleteAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAnimatedGraphicsTemplateRequest() (request *DeleteAnimatedGraphicsTemplateRequest) {
    request = &DeleteAnimatedGraphicsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAnimatedGraphicsTemplate")
    return
}

func NewDeleteAnimatedGraphicsTemplateResponse() (response *DeleteAnimatedGraphicsTemplateResponse) {
    response = &DeleteAnimatedGraphicsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAnimatedGraphicsTemplate
// This API is used to delete a custom animated image generating template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAnimatedGraphicsTemplate(request *DeleteAnimatedGraphicsTemplateRequest) (response *DeleteAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAnimatedGraphicsTemplateRequest()
    }
    response = NewDeleteAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContentReviewTemplateRequest() (request *DeleteContentReviewTemplateRequest) {
    request = &DeleteContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteContentReviewTemplate")
    return
}

func NewDeleteContentReviewTemplateResponse() (response *DeleteContentReviewTemplateResponse) {
    response = &DeleteContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteContentReviewTemplate
// This API is used to delete a custom intelligent content recognition template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteContentReviewTemplate(request *DeleteContentReviewTemplateRequest) (response *DeleteContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteContentReviewTemplateRequest()
    }
    response = NewDeleteContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageSpriteTemplateRequest() (request *DeleteImageSpriteTemplateRequest) {
    request = &DeleteImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteImageSpriteTemplate")
    return
}

func NewDeleteImageSpriteTemplateResponse() (response *DeleteImageSpriteTemplateResponse) {
    response = &DeleteImageSpriteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImageSpriteTemplate
// This API is used to delete an image sprite generating template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteImageSpriteTemplate(request *DeleteImageSpriteTemplateRequest) (response *DeleteImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteImageSpriteTemplateRequest()
    }
    response = NewDeleteImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonSampleRequest() (request *DeletePersonSampleRequest) {
    request = &DeletePersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeletePersonSample")
    return
}

func NewDeletePersonSampleResponse() (response *DeletePersonSampleResponse) {
    response = &DeletePersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePersonSample
// This API is used to delete image samples by image ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) DeletePersonSample(request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    if request == nil {
        request = NewDeletePersonSampleRequest()
    }
    response = NewDeletePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSampleSnapshotTemplateRequest() (request *DeleteSampleSnapshotTemplateRequest) {
    request = &DeleteSampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteSampleSnapshotTemplate")
    return
}

func NewDeleteSampleSnapshotTemplateResponse() (response *DeleteSampleSnapshotTemplateResponse) {
    response = &DeleteSampleSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSampleSnapshotTemplate
// This API is used to delete a custom sampled screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSampleSnapshotTemplate(request *DeleteSampleSnapshotTemplateRequest) (response *DeleteSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSampleSnapshotTemplateRequest()
    }
    response = NewDeleteSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotByTimeOffsetTemplateRequest() (request *DeleteSnapshotByTimeOffsetTemplateRequest) {
    request = &DeleteSnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteSnapshotByTimeOffsetTemplate")
    return
}

func NewDeleteSnapshotByTimeOffsetTemplateResponse() (response *DeleteSnapshotByTimeOffsetTemplateResponse) {
    response = &DeleteSnapshotByTimeOffsetTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSnapshotByTimeOffsetTemplate
// This API is used to delete a custom time point screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSnapshotByTimeOffsetTemplate(request *DeleteSnapshotByTimeOffsetTemplateRequest) (response *DeleteSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotByTimeOffsetTemplateRequest()
    }
    response = NewDeleteSnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTranscodeTemplateRequest() (request *DeleteTranscodeTemplateRequest) {
    request = &DeleteTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteTranscodeTemplate")
    return
}

func NewDeleteTranscodeTemplateResponse() (response *DeleteTranscodeTemplateResponse) {
    response = &DeleteTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTranscodeTemplate
// This API is used to delete a custom transcoding template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteTranscodeTemplate(request *DeleteTranscodeTemplateRequest) (response *DeleteTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTranscodeTemplateRequest()
    }
    response = NewDeleteTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWatermarkTemplateRequest() (request *DeleteWatermarkTemplateRequest) {
    request = &DeleteWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteWatermarkTemplate")
    return
}

func NewDeleteWatermarkTemplateResponse() (response *DeleteWatermarkTemplateResponse) {
    response = &DeleteWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWatermarkTemplate
// This API is used to delete a custom watermarking template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWatermarkTemplate(request *DeleteWatermarkTemplateRequest) (response *DeleteWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteWatermarkTemplateRequest()
    }
    response = NewDeleteWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWordSamplesRequest() (request *DeleteWordSamplesRequest) {
    request = &DeleteWordSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteWordSamples")
    return
}

func NewDeleteWordSamplesResponse() (response *DeleteWordSamplesResponse) {
    response = &DeleteWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWordSamples
// This API is used to delete keyword samples in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteWordSamples(request *DeleteWordSamplesRequest) (response *DeleteWordSamplesResponse, err error) {
    if request == nil {
        request = NewDeleteWordSamplesRequest()
    }
    response = NewDeleteWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowRequest() (request *DeleteWorkflowRequest) {
    request = &DeleteWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DeleteWorkflow")
    return
}

func NewDeleteWorkflowResponse() (response *DeleteWorkflowResponse) {
    response = &DeleteWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWorkflow
// This API is used to delete a workflow. An enabled workflow must be disabled before it can be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowRequest()
    }
    response = NewDeleteWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIAnalysisTemplatesRequest() (request *DescribeAIAnalysisTemplatesRequest) {
    request = &DescribeAIAnalysisTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAIAnalysisTemplates")
    return
}

func NewDescribeAIAnalysisTemplatesResponse() (response *DescribeAIAnalysisTemplatesResponse) {
    response = &DescribeAIAnalysisTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAIAnalysisTemplates
// This API is used to get the list of content analysis templates based on unique template ID. The returned result includes all eligible custom and preset video content analysis templates.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeAIAnalysisTemplates(request *DescribeAIAnalysisTemplatesRequest) (response *DescribeAIAnalysisTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIAnalysisTemplatesRequest()
    }
    response = NewDescribeAIAnalysisTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIRecognitionTemplatesRequest() (request *DescribeAIRecognitionTemplatesRequest) {
    request = &DescribeAIRecognitionTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAIRecognitionTemplates")
    return
}

func NewDescribeAIRecognitionTemplatesResponse() (response *DescribeAIRecognitionTemplatesResponse) {
    response = &DescribeAIRecognitionTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAIRecognitionTemplates
// This API is used to get the list of content recognition templates based on unique template ID. The return result includes all eligible custom and preset content recognition templates.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeAIRecognitionTemplates(request *DescribeAIRecognitionTemplatesRequest) (response *DescribeAIRecognitionTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIRecognitionTemplatesRequest()
    }
    response = NewDescribeAIRecognitionTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAdaptiveDynamicStreamingTemplatesRequest() (request *DescribeAdaptiveDynamicStreamingTemplatesRequest) {
    request = &DescribeAdaptiveDynamicStreamingTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAdaptiveDynamicStreamingTemplates")
    return
}

func NewDescribeAdaptiveDynamicStreamingTemplatesResponse() (response *DescribeAdaptiveDynamicStreamingTemplatesResponse) {
    response = &DescribeAdaptiveDynamicStreamingTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAdaptiveDynamicStreamingTemplates
// This API is used to query the list of adaptive bitrate streaming templates and supports paginated queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAdaptiveDynamicStreamingTemplates(request *DescribeAdaptiveDynamicStreamingTemplatesRequest) (response *DescribeAdaptiveDynamicStreamingTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAdaptiveDynamicStreamingTemplatesRequest()
    }
    response = NewDescribeAdaptiveDynamicStreamingTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAnimatedGraphicsTemplatesRequest() (request *DescribeAnimatedGraphicsTemplatesRequest) {
    request = &DescribeAnimatedGraphicsTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAnimatedGraphicsTemplates")
    return
}

func NewDescribeAnimatedGraphicsTemplatesResponse() (response *DescribeAnimatedGraphicsTemplatesResponse) {
    response = &DescribeAnimatedGraphicsTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAnimatedGraphicsTemplates
// This API is used to query the list of animated image generating templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeAnimatedGraphicsTemplates(request *DescribeAnimatedGraphicsTemplatesRequest) (response *DescribeAnimatedGraphicsTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAnimatedGraphicsTemplatesRequest()
    }
    response = NewDescribeAnimatedGraphicsTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentReviewTemplatesRequest() (request *DescribeContentReviewTemplatesRequest) {
    request = &DescribeContentReviewTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeContentReviewTemplates")
    return
}

func NewDescribeContentReviewTemplatesResponse() (response *DescribeContentReviewTemplatesResponse) {
    response = &DescribeContentReviewTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContentReviewTemplates
// This API is used to get the list of intelligent content recognition templates using their unique ID. The returned result includes all custom templates that meet the conditions as well as preset intelligent content recognition templates.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeContentReviewTemplates(request *DescribeContentReviewTemplatesRequest) (response *DescribeContentReviewTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeContentReviewTemplatesRequest()
    }
    response = NewDescribeContentReviewTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageSpriteTemplatesRequest() (request *DescribeImageSpriteTemplatesRequest) {
    request = &DescribeImageSpriteTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeImageSpriteTemplates")
    return
}

func NewDescribeImageSpriteTemplatesResponse() (response *DescribeImageSpriteTemplatesResponse) {
    response = &DescribeImageSpriteTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageSpriteTemplates
// This API is used to query the list of image sprite generating templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeImageSpriteTemplates(request *DescribeImageSpriteTemplatesRequest) (response *DescribeImageSpriteTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeImageSpriteTemplatesRequest()
    }
    response = NewDescribeImageSpriteTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaMetaDataRequest() (request *DescribeMediaMetaDataRequest) {
    request = &DescribeMediaMetaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeMediaMetaData")
    return
}

func NewDescribeMediaMetaDataResponse() (response *DescribeMediaMetaDataResponse) {
    response = &DescribeMediaMetaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMediaMetaData
// This API is used to get the metadata of media, such as video image width/height, codec, length, and frame rate.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) DescribeMediaMetaData(request *DescribeMediaMetaDataRequest) (response *DescribeMediaMetaDataResponse, err error) {
    if request == nil {
        request = NewDescribeMediaMetaDataRequest()
    }
    response = NewDescribeMediaMetaDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonSamplesRequest() (request *DescribePersonSamplesRequest) {
    request = &DescribePersonSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribePersonSamples")
    return
}

func NewDescribePersonSamplesResponse() (response *DescribePersonSamplesResponse) {
    response = &DescribePersonSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePersonSamples
// This API is used to query the information of image samples. It supports paginated queries by image ID, name, and tag.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePersonSamples(request *DescribePersonSamplesRequest) (response *DescribePersonSamplesResponse, err error) {
    if request == nil {
        request = NewDescribePersonSamplesRequest()
    }
    response = NewDescribePersonSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleSnapshotTemplatesRequest() (request *DescribeSampleSnapshotTemplatesRequest) {
    request = &DescribeSampleSnapshotTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeSampleSnapshotTemplates")
    return
}

func NewDescribeSampleSnapshotTemplatesResponse() (response *DescribeSampleSnapshotTemplatesResponse) {
    response = &DescribeSampleSnapshotTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSampleSnapshotTemplates
// This API is used to query the list of sampled screencapturing templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeSampleSnapshotTemplates(request *DescribeSampleSnapshotTemplatesRequest) (response *DescribeSampleSnapshotTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSampleSnapshotTemplatesRequest()
    }
    response = NewDescribeSampleSnapshotTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotByTimeOffsetTemplatesRequest() (request *DescribeSnapshotByTimeOffsetTemplatesRequest) {
    request = &DescribeSnapshotByTimeOffsetTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeSnapshotByTimeOffsetTemplates")
    return
}

func NewDescribeSnapshotByTimeOffsetTemplatesResponse() (response *DescribeSnapshotByTimeOffsetTemplatesResponse) {
    response = &DescribeSnapshotByTimeOffsetTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotByTimeOffsetTemplates
// This API is used to query the list of time point screencapturing templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeSnapshotByTimeOffsetTemplates(request *DescribeSnapshotByTimeOffsetTemplatesRequest) (response *DescribeSnapshotByTimeOffsetTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotByTimeOffsetTemplatesRequest()
    }
    response = NewDescribeSnapshotByTimeOffsetTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeTaskDetail")
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetail
// This API is used to query the details of execution status and result of a task submitted in the last 3 days by task ID.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeTasks")
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasks
// * This API is used to query the task list;
//
// * If there are many data entries in the list, one single call of the API may not be able to pull the entire list. The `ScrollToken` parameter can be used to pull the list in batches;
//
// * Only tasks in the last three days (72 hours) can be queried.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeTemplatesRequest() (request *DescribeTranscodeTemplatesRequest) {
    request = &DescribeTranscodeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeTranscodeTemplates")
    return
}

func NewDescribeTranscodeTemplatesResponse() (response *DescribeTranscodeTemplatesResponse) {
    response = &DescribeTranscodeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTranscodeTemplates
// This API is used to get the list of transcoding templates based on unique template ID. The return result includes all eligible custom and [preset transcoding templates](https://intl.cloud.tencent.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeTranscodeTemplates(request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeTemplatesRequest()
    }
    response = NewDescribeTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWatermarkTemplatesRequest() (request *DescribeWatermarkTemplatesRequest) {
    request = &DescribeWatermarkTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeWatermarkTemplates")
    return
}

func NewDescribeWatermarkTemplatesResponse() (response *DescribeWatermarkTemplatesResponse) {
    response = &DescribeWatermarkTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWatermarkTemplates
// This API is used to query custom watermarking templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeWatermarkTemplates(request *DescribeWatermarkTemplatesRequest) (response *DescribeWatermarkTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeWatermarkTemplatesRequest()
    }
    response = NewDescribeWatermarkTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWordSamplesRequest() (request *DescribeWordSamplesRequest) {
    request = &DescribeWordSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeWordSamples")
    return
}

func NewDescribeWordSamplesResponse() (response *DescribeWordSamplesResponse) {
    response = &DescribeWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWordSamples
// This API is used to perform paged queries of keyword sample information by use case, keyword, and tag.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWordSamples(request *DescribeWordSamplesRequest) (response *DescribeWordSamplesResponse, err error) {
    if request == nil {
        request = NewDescribeWordSamplesRequest()
    }
    response = NewDescribeWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowsRequest() (request *DescribeWorkflowsRequest) {
    request = &DescribeWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DescribeWorkflows")
    return
}

func NewDescribeWorkflowsResponse() (response *DescribeWorkflowsResponse) {
    response = &DescribeWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWorkflows
// This API is used to get the list of workflow details by workflow ID.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWorkflows(request *DescribeWorkflowsRequest) (response *DescribeWorkflowsResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowsRequest()
    }
    response = NewDescribeWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableWorkflowRequest() (request *DisableWorkflowRequest) {
    request = &DisableWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "DisableWorkflow")
    return
}

func NewDisableWorkflowResponse() (response *DisableWorkflowResponse) {
    response = &DisableWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableWorkflow
// This API is used to disable a workflow.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DisableWorkflow(request *DisableWorkflowRequest) (response *DisableWorkflowResponse, err error) {
    if request == nil {
        request = NewDisableWorkflowRequest()
    }
    response = NewDisableWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewEditMediaRequest() (request *EditMediaRequest) {
    request = &EditMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "EditMedia")
    return
}

func NewEditMediaResponse() (response *EditMediaResponse) {
    response = &EditMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditMedia
// This API is used to edit a video (by clipping, splicing, etc.) to generate a new VOD video. Editing features include:
//
// 
//
// 1. Clipping a file to generate a new video;
//
// 2. Splicing multiple files to generate a new video;
//
// 3. Clipping multiple files and then splicing the clips to generate a new video.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditMedia(request *EditMediaRequest) (response *EditMediaResponse, err error) {
    if request == nil {
        request = NewEditMediaRequest()
    }
    response = NewEditMediaResponse()
    err = c.Send(request, response)
    return
}

func NewEnableWorkflowRequest() (request *EnableWorkflowRequest) {
    request = &EnableWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "EnableWorkflow")
    return
}

func NewEnableWorkflowResponse() (response *EnableWorkflowResponse) {
    response = &EnableWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableWorkflow
// This API is used to enable a workflow.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) EnableWorkflow(request *EnableWorkflowRequest) (response *EnableWorkflowResponse, err error) {
    if request == nil {
        request = NewEnableWorkflowRequest()
    }
    response = NewEnableWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteFunctionRequest() (request *ExecuteFunctionRequest) {
    request = &ExecuteFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ExecuteFunction")
    return
}

func NewExecuteFunctionResponse() (response *ExecuteFunctionResponse) {
    response = &ExecuteFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExecuteFunction
// This API is only used in unique custom development scenarios. Unless requested by Media Processing Service customer service, please do not call it.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExecuteFunction(request *ExecuteFunctionRequest) (response *ExecuteFunctionResponse, err error) {
    if request == nil {
        request = NewExecuteFunctionRequest()
    }
    response = NewExecuteFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewManageTaskRequest() (request *ManageTaskRequest) {
    request = &ManageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ManageTask")
    return
}

func NewManageTaskResponse() (response *ManageTaskResponse) {
    response = &ManageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManageTask
// This API is used to manage initiated tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDOPERATIONTYPE = "InvalidParameterValue.InvalidOperationType"
//  INVALIDPARAMETERVALUE_NOTPROCESSINGTASK = "InvalidParameterValue.NotProcessingTask"
//  INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"
func (c *Client) ManageTask(request *ManageTaskRequest) (response *ManageTaskResponse, err error) {
    if request == nil {
        request = NewManageTaskRequest()
    }
    response = NewManageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAIAnalysisTemplateRequest() (request *ModifyAIAnalysisTemplateRequest) {
    request = &ModifyAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAIAnalysisTemplate")
    return
}

func NewModifyAIAnalysisTemplateResponse() (response *ModifyAIAnalysisTemplateResponse) {
    response = &ModifyAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAIAnalysisTemplate
// This API is used to modify a custom content analysis template.
//
// 
//
// Note: templates with an ID below 10000 are preset and cannot be modified.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLASSIFCATIONCONFIGURE = "InvalidParameterValue.ClassifcationConfigure"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COVERCONFIGURE = "InvalidParameterValue.CoverConfigure"
//  INVALIDPARAMETERVALUE_FRAMETAGCONFIGURE = "InvalidParameterValue.FrameTagConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAIAnalysisTemplate(request *ModifyAIAnalysisTemplateRequest) (response *ModifyAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIAnalysisTemplateRequest()
    }
    response = NewModifyAIAnalysisTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAIRecognitionTemplateRequest() (request *ModifyAIRecognitionTemplateRequest) {
    request = &ModifyAIRecognitionTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAIRecognitionTemplate")
    return
}

func NewModifyAIRecognitionTemplateResponse() (response *ModifyAIRecognitionTemplateResponse) {
    response = &ModifyAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAIRecognitionTemplate
// This API is used to modify a custom content recognition template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAIRecognitionTemplate(request *ModifyAIRecognitionTemplateRequest) (response *ModifyAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIRecognitionTemplateRequest()
    }
    response = NewModifyAIRecognitionTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAdaptiveDynamicStreamingTemplateRequest() (request *ModifyAdaptiveDynamicStreamingTemplateRequest) {
    request = &ModifyAdaptiveDynamicStreamingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAdaptiveDynamicStreamingTemplate")
    return
}

func NewModifyAdaptiveDynamicStreamingTemplateResponse() (response *ModifyAdaptiveDynamicStreamingTemplateResponse) {
    response = &ModifyAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAdaptiveDynamicStreamingTemplate
// This API is used to modify an adaptive bitrate streaming template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
func (c *Client) ModifyAdaptiveDynamicStreamingTemplate(request *ModifyAdaptiveDynamicStreamingTemplateRequest) (response *ModifyAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAdaptiveDynamicStreamingTemplateRequest()
    }
    response = NewModifyAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAnimatedGraphicsTemplateRequest() (request *ModifyAnimatedGraphicsTemplateRequest) {
    request = &ModifyAnimatedGraphicsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAnimatedGraphicsTemplate")
    return
}

func NewModifyAnimatedGraphicsTemplateResponse() (response *ModifyAnimatedGraphicsTemplateResponse) {
    response = &ModifyAnimatedGraphicsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAnimatedGraphicsTemplate
// This API is used to modify a custom animated image generating template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FORMATWEBPLACKWIDTHANDHEIGHT = "InvalidParameterValue.FormatWebpLackWidthAndHeight"
//  INVALIDPARAMETERVALUE_FORMATWEBPWIDTHANDHEIGHTBOTHZERO = "InvalidParameterValue.FormatWebpWidthAndHeightBothZero"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_QUALITY = "InvalidParameterValue.Quality"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAnimatedGraphicsTemplate(request *ModifyAnimatedGraphicsTemplateRequest) (response *ModifyAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAnimatedGraphicsTemplateRequest()
    }
    response = NewModifyAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContentReviewTemplateRequest() (request *ModifyContentReviewTemplateRequest) {
    request = &ModifyContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyContentReviewTemplate")
    return
}

func NewModifyContentReviewTemplateResponse() (response *ModifyContentReviewTemplateResponse) {
    response = &ModifyContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyContentReviewTemplate
// This API is used to modify a custom intelligent content recognition template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyContentReviewTemplate(request *ModifyContentReviewTemplateRequest) (response *ModifyContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewModifyContentReviewTemplateRequest()
    }
    response = NewModifyContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageSpriteTemplateRequest() (request *ModifyImageSpriteTemplateRequest) {
    request = &ModifyImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyImageSpriteTemplate")
    return
}

func NewModifyImageSpriteTemplateResponse() (response *ModifyImageSpriteTemplateResponse) {
    response = &ModifyImageSpriteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyImageSpriteTemplate
// This API is used to modify a custom image sprite generating template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyImageSpriteTemplate(request *ModifyImageSpriteTemplateRequest) (response *ModifyImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewModifyImageSpriteTemplateRequest()
    }
    response = NewModifyImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonSampleRequest() (request *ModifyPersonSampleRequest) {
    request = &ModifyPersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyPersonSample")
    return
}

func NewModifyPersonSampleResponse() (response *ModifyPersonSampleResponse) {
    response = &ModifyPersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPersonSample
// This API is used to modify image samples by image ID. You can use it to modify the name and description of an image sample and add/delete/reset facial features or tags. There must be at least one image left after the deletion of facial features; otherwise, please reset instead of delete the facial features.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) ModifyPersonSample(request *ModifyPersonSampleRequest) (response *ModifyPersonSampleResponse, err error) {
    if request == nil {
        request = NewModifyPersonSampleRequest()
    }
    response = NewModifyPersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySampleSnapshotTemplateRequest() (request *ModifySampleSnapshotTemplateRequest) {
    request = &ModifySampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifySampleSnapshotTemplate")
    return
}

func NewModifySampleSnapshotTemplateResponse() (response *ModifySampleSnapshotTemplateResponse) {
    response = &ModifySampleSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySampleSnapshotTemplate
// This API is used to modify a custom sampled screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySampleSnapshotTemplate(request *ModifySampleSnapshotTemplateRequest) (response *ModifySampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewModifySampleSnapshotTemplateRequest()
    }
    response = NewModifySampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotByTimeOffsetTemplateRequest() (request *ModifySnapshotByTimeOffsetTemplateRequest) {
    request = &ModifySnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifySnapshotByTimeOffsetTemplate")
    return
}

func NewModifySnapshotByTimeOffsetTemplateResponse() (response *ModifySnapshotByTimeOffsetTemplateResponse) {
    response = &ModifySnapshotByTimeOffsetTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySnapshotByTimeOffsetTemplate
// This API is used to modify a custom time point screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySnapshotByTimeOffsetTemplate(request *ModifySnapshotByTimeOffsetTemplateRequest) (response *ModifySnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewModifySnapshotByTimeOffsetTemplateRequest()
    }
    response = NewModifySnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTranscodeTemplateRequest() (request *ModifyTranscodeTemplateRequest) {
    request = &ModifyTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyTranscodeTemplate")
    return
}

func NewModifyTranscodeTemplateResponse() (response *ModifyTranscodeTemplateResponse) {
    response = &ModifyTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTranscodeTemplate
// This API is used to modify a custom transcoding template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyTranscodeTemplate(request *ModifyTranscodeTemplateRequest) (response *ModifyTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyTranscodeTemplateRequest()
    }
    response = NewModifyTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWatermarkTemplateRequest() (request *ModifyWatermarkTemplateRequest) {
    request = &ModifyWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyWatermarkTemplate")
    return
}

func NewModifyWatermarkTemplateResponse() (response *ModifyWatermarkTemplateResponse) {
    response = &ModifyWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWatermarkTemplate
// This API is used to modify a custom watermarking template. The watermark type cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UPLOADWATERMARKERROR = "InternalError.UploadWatermarkError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COORDINATEORIGIN = "InvalidParameterValue.CoordinateOrigin"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_IMAGECONTENT = "InvalidParameterValue.ImageContent"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REPEATTYPE = "InvalidParameterValue.RepeatType"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEHEIGHT = "InvalidParameterValue.SvgTemplateHeight"
//  INVALIDPARAMETERVALUE_SVGTEMPLATEWIDTH = "InvalidParameterValue.SvgTemplateWidth"
//  INVALIDPARAMETERVALUE_TEXTALPHA = "InvalidParameterValue.TextAlpha"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  INVALIDPARAMETERVALUE_XPOS = "InvalidParameterValue.XPos"
//  INVALIDPARAMETERVALUE_YPOS = "InvalidParameterValue.YPos"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyWatermarkTemplate(request *ModifyWatermarkTemplateRequest) (response *ModifyWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewModifyWatermarkTemplateRequest()
    }
    response = NewModifyWatermarkTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWordSampleRequest() (request *ModifyWordSampleRequest) {
    request = &ModifyWordSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ModifyWordSample")
    return
}

func NewModifyWordSampleResponse() (response *ModifyWordSampleResponse) {
    response = &ModifyWordSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWordSample
// This API is used to modify the use case and tag of a keyword. The keyword itself cannot be modified, but you can delete it and create another one if needed.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_WORD = "ResourceNotFound.Word"
func (c *Client) ModifyWordSample(request *ModifyWordSampleRequest) (response *ModifyWordSampleResponse, err error) {
    if request == nil {
        request = NewModifyWordSampleRequest()
    }
    response = NewModifyWordSampleResponse()
    err = c.Send(request, response)
    return
}

func NewParseLiveStreamProcessNotificationRequest() (request *ParseLiveStreamProcessNotificationRequest) {
    request = &ParseLiveStreamProcessNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ParseLiveStreamProcessNotification")
    return
}

func NewParseLiveStreamProcessNotificationResponse() (response *ParseLiveStreamProcessNotificationResponse) {
    response = &ParseLiveStreamProcessNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ParseLiveStreamProcessNotification
// This API is used to parse the content of an MPS live stream processing event notification from the `msgBody` field in the message received from CMQ.
//
// Instead of initiating a video processing task, this API is used to help generate SDKs for various programming languages. You can parse the event notification based on the analytic function of the SDKs.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
func (c *Client) ParseLiveStreamProcessNotification(request *ParseLiveStreamProcessNotificationRequest) (response *ParseLiveStreamProcessNotificationResponse, err error) {
    if request == nil {
        request = NewParseLiveStreamProcessNotificationRequest()
    }
    response = NewParseLiveStreamProcessNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewParseNotificationRequest() (request *ParseNotificationRequest) {
    request = &ParseNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ParseNotification")
    return
}

func NewParseNotificationResponse() (response *ParseNotificationResponse) {
    response = &ParseNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ParseNotification
// This API is used to parse the content of an MPS event notification from the `msgBody` field in the message received from CMQ.
//
// Instead of initiating a video processing task, this API is used to help generate SDKs for various programming languages. You can parse the event notification based on the analytic function of the SDKs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
func (c *Client) ParseNotification(request *ParseNotificationRequest) (response *ParseNotificationResponse, err error) {
    if request == nil {
        request = NewParseNotificationRequest()
    }
    response = NewParseNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewProcessLiveStreamRequest() (request *ProcessLiveStreamRequest) {
    request = &ProcessLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ProcessLiveStream")
    return
}

func NewProcessLiveStreamResponse() (response *ProcessLiveStreamResponse) {
    response = &ProcessLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ProcessLiveStream
// This API is used to initiate tasks to process live streaming media files. Such tasks include:
//
// 
//
// * Intelligent content moderation (detecting porn, terrorism, and politically sensitive information in image and porn information in speech)
//
// * Intelligent content recognition (face, full text, text keyword, full speech, and speech keyword)
//
// 
//
// Live stream processing event notifications are written into specified CMQ queues in real time. Users need to obtain event notification results from such CMQ queues. Output files of the processing tasks are written into destination buckets specified by users.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) ProcessLiveStream(request *ProcessLiveStreamRequest) (response *ProcessLiveStreamResponse, err error) {
    if request == nil {
        request = NewProcessLiveStreamRequest()
    }
    response = NewProcessLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaRequest() (request *ProcessMediaRequest) {
    request = &ProcessMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ProcessMedia")
    return
}

func NewProcessMediaResponse() (response *ProcessMediaResponse) {
    response = &ProcessMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ProcessMedia
// This API is used to initiate a processing task for media files in COS, including:
//
// 1. Video transcoding (with watermark);
//
// 2. Animated image generating;
//
// 3. Time point screencapturing;
//
// 4. Sampled screencapturing;
//
// 5. Image sprite generating;
//
// 6. Video conversion to adaptive bitrate streaming;
//
// 7. Intelligent content audit (detection of porn, terrorism, and politically sensitive information);
//
// 8. Intelligent content recognition (face recognition, full text recognition, text keyword recognition, full speech recognition, and speech keyword recognition).
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessMedia(request *ProcessMediaRequest) (response *ProcessMediaResponse, err error) {
    if request == nil {
        request = NewProcessMediaRequest()
    }
    response = NewProcessMediaResponse()
    err = c.Send(request, response)
    return
}

func NewResetWorkflowRequest() (request *ResetWorkflowRequest) {
    request = &ResetWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mps", APIVersion, "ResetWorkflow")
    return
}

func NewResetWorkflowResponse() (response *ResetWorkflowResponse) {
    response = &ResetWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetWorkflow
// This API is used to reset an existing workflow that is disabled.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ResetWorkflow(request *ResetWorkflowRequest) (response *ResetWorkflowResponse, err error) {
    if request == nil {
        request = NewResetWorkflowRequest()
    }
    response = NewResetWorkflowResponse()
    err = c.Send(request, response)
    return
}
