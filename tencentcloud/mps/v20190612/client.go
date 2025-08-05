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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewBatchProcessMediaRequest() (request *BatchProcessMediaRequest) {
    request = &BatchProcessMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "BatchProcessMedia")
    
    
    return
}

func NewBatchProcessMediaResponse() (response *BatchProcessMediaResponse) {
    response = &BatchProcessMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchProcessMedia
// This API is used to initiate batch processing tasks for URL video links, with features including:
//
// Smart subtitle (full speech, speech hotword, and speech translation)
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) BatchProcessMedia(request *BatchProcessMediaRequest) (response *BatchProcessMediaResponse, err error) {
    return c.BatchProcessMediaWithContext(context.Background(), request)
}

// BatchProcessMedia
// This API is used to initiate batch processing tasks for URL video links, with features including:
//
// Smart subtitle (full speech, speech hotword, and speech translation)
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) BatchProcessMediaWithContext(ctx context.Context, request *BatchProcessMediaRequest) (response *BatchProcessMediaResponse, err error) {
    if request == nil {
        request = NewBatchProcessMediaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "BatchProcessMedia")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchProcessMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchProcessMediaResponse()
    err = c.Send(request, response)
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
    return c.CreateAIAnalysisTemplateWithContext(context.Background(), request)
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
func (c *Client) CreateAIAnalysisTemplateWithContext(ctx context.Context, request *CreateAIAnalysisTemplateRequest) (response *CreateAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIAnalysisTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAIAnalysisTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAIAnalysisTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_DESTINATIONLANGUAGE = "InvalidParameterValue.DestinationLanguage"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAIRecognitionTemplate(request *CreateAIRecognitionTemplateRequest) (response *CreateAIRecognitionTemplateResponse, err error) {
    return c.CreateAIRecognitionTemplateWithContext(context.Background(), request)
}

// CreateAIRecognitionTemplate
// This API is used to create a custom content recognition template. Up to 50 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_DESTINATIONLANGUAGE = "InvalidParameterValue.DestinationLanguage"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAIRecognitionTemplateWithContext(ctx context.Context, request *CreateAIRecognitionTemplateRequest) (response *CreateAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIRecognitionTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAIRecognitionTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAIRecognitionTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateAdaptiveDynamicStreamingTemplateWithContext(context.Background(), request)
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
func (c *Client) CreateAdaptiveDynamicStreamingTemplateWithContext(ctx context.Context, request *CreateAdaptiveDynamicStreamingTemplateRequest) (response *CreateAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAdaptiveDynamicStreamingTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAdaptiveDynamicStreamingTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateAnimatedGraphicsTemplateWithContext(context.Background(), request)
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
func (c *Client) CreateAnimatedGraphicsTemplateWithContext(ctx context.Context, request *CreateAnimatedGraphicsTemplateRequest) (response *CreateAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAnimatedGraphicsTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAnimatedGraphicsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAsrHotwordsRequest() (request *CreateAsrHotwordsRequest) {
    request = &CreateAsrHotwordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateAsrHotwords")
    
    
    return
}

func NewCreateAsrHotwordsResponse() (response *CreateAsrHotwordsResponse) {
    response = &CreateAsrHotwordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAsrHotwords
// This API is used to create a smart subtitle hotword lexicon.
//
// error code that may be returned:
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSFORMATERROR = "InvalidParameterValue.HotwordsFormatError"
//  INVALIDPARAMETERVALUE_INPUTINFO = "InvalidParameterValue.InputInfo"
//  LIMITEXCEEDED_TOOMUCHHOTWORDS = "LimitExceeded.TooMuchHotWords"
//  LIMITEXCEEDED_TOOMUCHLARGEHOTWORDS = "LimitExceeded.TooMuchLargeHotWords"
func (c *Client) CreateAsrHotwords(request *CreateAsrHotwordsRequest) (response *CreateAsrHotwordsResponse, err error) {
    return c.CreateAsrHotwordsWithContext(context.Background(), request)
}

// CreateAsrHotwords
// This API is used to create a smart subtitle hotword lexicon.
//
// error code that may be returned:
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSFORMATERROR = "InvalidParameterValue.HotwordsFormatError"
//  INVALIDPARAMETERVALUE_INPUTINFO = "InvalidParameterValue.InputInfo"
//  LIMITEXCEEDED_TOOMUCHHOTWORDS = "LimitExceeded.TooMuchHotWords"
//  LIMITEXCEEDED_TOOMUCHLARGEHOTWORDS = "LimitExceeded.TooMuchLargeHotWords"
func (c *Client) CreateAsrHotwordsWithContext(ctx context.Context, request *CreateAsrHotwordsRequest) (response *CreateAsrHotwordsResponse, err error) {
    if request == nil {
        request = NewCreateAsrHotwordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateAsrHotwords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAsrHotwords require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAsrHotwordsResponse()
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
// This API is used to create a custom content moderation template. Up to 50 templates can be created in total.
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
    return c.CreateContentReviewTemplateWithContext(context.Background(), request)
}

// CreateContentReviewTemplate
// This API is used to create a custom content moderation template. Up to 50 templates can be created in total.
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
func (c *Client) CreateContentReviewTemplateWithContext(ctx context.Context, request *CreateContentReviewTemplateRequest) (response *CreateContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewCreateContentReviewTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateContentReviewTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateImageSpriteTemplate(request *CreateImageSpriteTemplateRequest) (response *CreateImageSpriteTemplateResponse, err error) {
    return c.CreateImageSpriteTemplateWithContext(context.Background(), request)
}

// CreateImageSpriteTemplate
// This API is used to create a custom image sprite generating template. Up to 16 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateImageSpriteTemplateWithContext(ctx context.Context, request *CreateImageSpriteTemplateRequest) (response *CreateImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewCreateImageSpriteTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateImageSpriteTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordTemplateRequest() (request *CreateLiveRecordTemplateRequest) {
    request = &CreateLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateLiveRecordTemplate")
    
    
    return
}

func NewCreateLiveRecordTemplateResponse() (response *CreateLiveRecordTemplateResponse) {
    response = &CreateLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveRecordTemplate
// This API is used to create a live recording template.
//
// error code that may be returned:
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateLiveRecordTemplate(request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    return c.CreateLiveRecordTemplateWithContext(context.Background(), request)
}

// CreateLiveRecordTemplate
// This API is used to create a live recording template.
//
// error code that may be returned:
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateLiveRecordTemplateWithContext(ctx context.Context, request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveRecordTemplateResponse()
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
    return c.CreatePersonSampleWithContext(context.Background(), request)
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
func (c *Client) CreatePersonSampleWithContext(ctx context.Context, request *CreatePersonSampleRequest) (response *CreatePersonSampleResponse, err error) {
    if request == nil {
        request = NewCreatePersonSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreatePersonSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQualityControlTemplateRequest() (request *CreateQualityControlTemplateRequest) {
    request = &CreateQualityControlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateQualityControlTemplate")
    
    
    return
}

func NewCreateQualityControlTemplateResponse() (response *CreateQualityControlTemplateResponse) {
    response = &CreateQualityControlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQualityControlTemplate
// This API is used to create a media quality inspection template. Up to 50 templates can be created.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYDETECTITEM = "InvalidParameterValue.EmptyDetectItem"
//  INVALIDPARAMETERVALUE_UNKNOWNCATEGORY = "InvalidParameterValue.UnknownCategory"
func (c *Client) CreateQualityControlTemplate(request *CreateQualityControlTemplateRequest) (response *CreateQualityControlTemplateResponse, err error) {
    return c.CreateQualityControlTemplateWithContext(context.Background(), request)
}

// CreateQualityControlTemplate
// This API is used to create a media quality inspection template. Up to 50 templates can be created.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYDETECTITEM = "InvalidParameterValue.EmptyDetectItem"
//  INVALIDPARAMETERVALUE_UNKNOWNCATEGORY = "InvalidParameterValue.UnknownCategory"
func (c *Client) CreateQualityControlTemplateWithContext(ctx context.Context, request *CreateQualityControlTemplateRequest) (response *CreateQualityControlTemplateResponse, err error) {
    if request == nil {
        request = NewCreateQualityControlTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateQualityControlTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQualityControlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQualityControlTemplateResponse()
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
    return c.CreateSampleSnapshotTemplateWithContext(context.Background(), request)
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
func (c *Client) CreateSampleSnapshotTemplateWithContext(ctx context.Context, request *CreateSampleSnapshotTemplateRequest) (response *CreateSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSampleSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateSampleSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScheduleRequest() (request *CreateScheduleRequest) {
    request = &CreateScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateSchedule")
    
    
    return
}

func NewCreateScheduleResponse() (response *CreateScheduleResponse) {
    response = &CreateScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSchedule
// This API is used to create a scheme for media files uploaded to a specified COS bucket. A scheme may include the following tasks:
//
// 1. Video transcoding (with watermark)
//
// 2. Animated screenshot generating
//
// 3. Time point screencapturing
//
// 4. Sampled screencapturing
//
// 5. Image sprite generating
//
// 6. Adaptive bitrate streaming
//
// 7. Intelligent content moderation (detection of pornographic and sensitive content)
//
// 8. Intelligent content analysis (labeling, categorization, thumbnail generation, labeling by frame)
//
// 9. Intelligent content recognition (face, full text, text keyword, full speech, and speech keyword)
//
// 
//
// Note: A scheme is disabled upon creation. You need to manually enable it.
//
// error code that may be returned:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
func (c *Client) CreateSchedule(request *CreateScheduleRequest) (response *CreateScheduleResponse, err error) {
    return c.CreateScheduleWithContext(context.Background(), request)
}

// CreateSchedule
// This API is used to create a scheme for media files uploaded to a specified COS bucket. A scheme may include the following tasks:
//
// 1. Video transcoding (with watermark)
//
// 2. Animated screenshot generating
//
// 3. Time point screencapturing
//
// 4. Sampled screencapturing
//
// 5. Image sprite generating
//
// 6. Adaptive bitrate streaming
//
// 7. Intelligent content moderation (detection of pornographic and sensitive content)
//
// 8. Intelligent content analysis (labeling, categorization, thumbnail generation, labeling by frame)
//
// 9. Intelligent content recognition (face, full text, text keyword, full speech, and speech keyword)
//
// 
//
// Note: A scheme is disabled upon creation. You need to manually enable it.
//
// error code that may be returned:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
func (c *Client) CreateScheduleWithContext(ctx context.Context, request *CreateScheduleRequest) (response *CreateScheduleResponse, err error) {
    if request == nil {
        request = NewCreateScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateSchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSmartSubtitleTemplateRequest() (request *CreateSmartSubtitleTemplateRequest) {
    request = &CreateSmartSubtitleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "CreateSmartSubtitleTemplate")
    
    
    return
}

func NewCreateSmartSubtitleTemplateResponse() (response *CreateSmartSubtitleTemplateResponse) {
    response = &CreateSmartSubtitleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSmartSubtitleTemplate
// This API is used to create a custom smart subtitle template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ASRHOTWORDSCONFIGURE = "InvalidParameterValue.AsrHotWordsConfigure"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSLIBRARYID = "InvalidParameterValue.AsrHotWordsLibraryId"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSSWITCH = "InvalidParameterValue.AsrHotWordsSwitch"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLETYPE = "InvalidParameterValue.SubtitleType"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  INVALIDPARAMETERVALUE_TRANSLATESWITCH = "InvalidParameterValue.TranslateSwitch"
//  INVALIDPARAMETERVALUE_VIDEOSRCLANGUAGE = "InvalidParameterValue.VideoSrcLanguage"
func (c *Client) CreateSmartSubtitleTemplate(request *CreateSmartSubtitleTemplateRequest) (response *CreateSmartSubtitleTemplateResponse, err error) {
    return c.CreateSmartSubtitleTemplateWithContext(context.Background(), request)
}

// CreateSmartSubtitleTemplate
// This API is used to create a custom smart subtitle template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ASRHOTWORDSCONFIGURE = "InvalidParameterValue.AsrHotWordsConfigure"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSLIBRARYID = "InvalidParameterValue.AsrHotWordsLibraryId"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSSWITCH = "InvalidParameterValue.AsrHotWordsSwitch"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLETYPE = "InvalidParameterValue.SubtitleType"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  INVALIDPARAMETERVALUE_TRANSLATESWITCH = "InvalidParameterValue.TranslateSwitch"
//  INVALIDPARAMETERVALUE_VIDEOSRCLANGUAGE = "InvalidParameterValue.VideoSrcLanguage"
func (c *Client) CreateSmartSubtitleTemplateWithContext(ctx context.Context, request *CreateSmartSubtitleTemplateRequest) (response *CreateSmartSubtitleTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSmartSubtitleTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateSmartSubtitleTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSmartSubtitleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSmartSubtitleTemplateResponse()
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
    return c.CreateSnapshotByTimeOffsetTemplateWithContext(context.Background(), request)
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
func (c *Client) CreateSnapshotByTimeOffsetTemplateWithContext(ctx context.Context, request *CreateSnapshotByTimeOffsetTemplateRequest) (response *CreateSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotByTimeOffsetTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateSnapshotByTimeOffsetTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
    return c.CreateTranscodeTemplateWithContext(context.Background(), request)
}

// CreateTranscodeTemplate
// This API is used to create a custom transcoding template. Up to 1,000 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
func (c *Client) CreateTranscodeTemplateWithContext(ctx context.Context, request *CreateTranscodeTemplateRequest) (response *CreateTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateWatermarkTemplateWithContext(context.Background(), request)
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
func (c *Client) CreateWatermarkTemplateWithContext(ctx context.Context, request *CreateWatermarkTemplateRequest) (response *CreateWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewCreateWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateWordSamplesWithContext(context.Background(), request)
}

// CreateWordSamples
// This API is used to create keyword samples in batches for video processing operations such as content recognition and inappropriate information detection with the help of the OCR and ASR technologies.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateWordSamplesWithContext(ctx context.Context, request *CreateWordSamplesRequest) (response *CreateWordSamplesResponse, err error) {
    if request == nil {
        request = NewCreateWordSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateWordSamples")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWordSamples require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a workflow for media files uploaded to a specified COS bucket. A workflow may include the following tasks:
//
// 1. Video transcoding (with watermark)
//
// 2. Animated image generating
//
// 3. Time point screencapturing
//
// 4. Sampled screencapturing
//
// 5. Image sprite generating
//
// 6. Adaptive bitrate streaming
//
// 7. Intelligent content moderation (detection of pornographic and sensitive content)
//
// 8. Intelligent content analysis (labeling, categorization, thumbnail generation, frame-specific labeling)
//
// 9. Intelligent content recognition (face, full text, text keyword, full speech, and speech keyword)
//
// 
//
// Note: A workflow is disabled upon creation. You need to manually enable it.
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
    return c.CreateWorkflowWithContext(context.Background(), request)
}

// CreateWorkflow
// This API is used to create a workflow for media files uploaded to a specified COS bucket. A workflow may include the following tasks:
//
// 1. Video transcoding (with watermark)
//
// 2. Animated image generating
//
// 3. Time point screencapturing
//
// 4. Sampled screencapturing
//
// 5. Image sprite generating
//
// 6. Adaptive bitrate streaming
//
// 7. Intelligent content moderation (detection of pornographic and sensitive content)
//
// 8. Intelligent content analysis (labeling, categorization, thumbnail generation, frame-specific labeling)
//
// 9. Intelligent content recognition (face, full text, text keyword, full speech, and speech keyword)
//
// 
//
// Note: A workflow is disabled upon creation. You need to manually enable it.
//
// error code that may be returned:
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
func (c *Client) CreateWorkflowWithContext(ctx context.Context, request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "CreateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflow require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteAIAnalysisTemplateWithContext(context.Background(), request)
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
func (c *Client) DeleteAIAnalysisTemplateWithContext(ctx context.Context, request *DeleteAIAnalysisTemplateRequest) (response *DeleteAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIAnalysisTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAIAnalysisTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAIAnalysisTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteAIRecognitionTemplateWithContext(context.Background(), request)
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
func (c *Client) DeleteAIRecognitionTemplateWithContext(ctx context.Context, request *DeleteAIRecognitionTemplateRequest) (response *DeleteAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIRecognitionTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAIRecognitionTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAIRecognitionTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAdaptiveDynamicStreamingTemplate(request *DeleteAdaptiveDynamicStreamingTemplateRequest) (response *DeleteAdaptiveDynamicStreamingTemplateResponse, err error) {
    return c.DeleteAdaptiveDynamicStreamingTemplateWithContext(context.Background(), request)
}

// DeleteAdaptiveDynamicStreamingTemplate
// This API is used to delete an adaptive bitrate streaming template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAdaptiveDynamicStreamingTemplateWithContext(ctx context.Context, request *DeleteAdaptiveDynamicStreamingTemplateRequest) (response *DeleteAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAdaptiveDynamicStreamingTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAdaptiveDynamicStreamingTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteAnimatedGraphicsTemplateWithContext(context.Background(), request)
}

// DeleteAnimatedGraphicsTemplate
// This API is used to delete a custom animated image generating template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAnimatedGraphicsTemplateWithContext(ctx context.Context, request *DeleteAnimatedGraphicsTemplateRequest) (response *DeleteAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAnimatedGraphicsTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAnimatedGraphicsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAsrHotwordsRequest() (request *DeleteAsrHotwordsRequest) {
    request = &DeleteAsrHotwordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteAsrHotwords")
    
    
    return
}

func NewDeleteAsrHotwordsResponse() (response *DeleteAsrHotwordsResponse) {
    response = &DeleteAsrHotwordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAsrHotwords
// This API is used to delete a smart subtitle hotword lexicon.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSNOTEXIST = "InvalidParameterValue.HotWordsNotExist"
func (c *Client) DeleteAsrHotwords(request *DeleteAsrHotwordsRequest) (response *DeleteAsrHotwordsResponse, err error) {
    return c.DeleteAsrHotwordsWithContext(context.Background(), request)
}

// DeleteAsrHotwords
// This API is used to delete a smart subtitle hotword lexicon.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSNOTEXIST = "InvalidParameterValue.HotWordsNotExist"
func (c *Client) DeleteAsrHotwordsWithContext(ctx context.Context, request *DeleteAsrHotwordsRequest) (response *DeleteAsrHotwordsResponse, err error) {
    if request == nil {
        request = NewDeleteAsrHotwordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteAsrHotwords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAsrHotwords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAsrHotwordsResponse()
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
// This API is used to delete a custom content moderation template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteContentReviewTemplate(request *DeleteContentReviewTemplateRequest) (response *DeleteContentReviewTemplateResponse, err error) {
    return c.DeleteContentReviewTemplateWithContext(context.Background(), request)
}

// DeleteContentReviewTemplate
// This API is used to delete a custom content moderation template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteContentReviewTemplateWithContext(ctx context.Context, request *DeleteContentReviewTemplateRequest) (response *DeleteContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteContentReviewTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteContentReviewTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteImageSpriteTemplateWithContext(context.Background(), request)
}

// DeleteImageSpriteTemplate
// This API is used to delete an image sprite generating template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteImageSpriteTemplateWithContext(ctx context.Context, request *DeleteImageSpriteTemplateRequest) (response *DeleteImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteImageSpriteTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteImageSpriteTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordTemplateRequest() (request *DeleteLiveRecordTemplateRequest) {
    request = &DeleteLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteLiveRecordTemplate")
    
    
    return
}

func NewDeleteLiveRecordTemplateResponse() (response *DeleteLiveRecordTemplateResponse) {
    response = &DeleteLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveRecordTemplate
// This API is used to delete a live recording template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteLiveRecordTemplate(request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    return c.DeleteLiveRecordTemplateWithContext(context.Background(), request)
}

// DeleteLiveRecordTemplate
// This API is used to delete a live recording template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteLiveRecordTemplateWithContext(ctx context.Context, request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveRecordTemplateResponse()
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
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) DeletePersonSample(request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    return c.DeletePersonSampleWithContext(context.Background(), request)
}

// DeletePersonSample
// This API is used to delete image samples by image ID.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) DeletePersonSampleWithContext(ctx context.Context, request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    if request == nil {
        request = NewDeletePersonSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeletePersonSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQualityControlTemplateRequest() (request *DeleteQualityControlTemplateRequest) {
    request = &DeleteQualityControlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteQualityControlTemplate")
    
    
    return
}

func NewDeleteQualityControlTemplateResponse() (response *DeleteQualityControlTemplateResponse) {
    response = &DeleteQualityControlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQualityControlTemplate
// This API is used to delete a media quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteQualityControlTemplate(request *DeleteQualityControlTemplateRequest) (response *DeleteQualityControlTemplateResponse, err error) {
    return c.DeleteQualityControlTemplateWithContext(context.Background(), request)
}

// DeleteQualityControlTemplate
// This API is used to delete a media quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteQualityControlTemplateWithContext(ctx context.Context, request *DeleteQualityControlTemplateRequest) (response *DeleteQualityControlTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteQualityControlTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteQualityControlTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQualityControlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQualityControlTemplateResponse()
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
    return c.DeleteSampleSnapshotTemplateWithContext(context.Background(), request)
}

// DeleteSampleSnapshotTemplate
// This API is used to delete a custom sampled screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSampleSnapshotTemplateWithContext(ctx context.Context, request *DeleteSampleSnapshotTemplateRequest) (response *DeleteSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSampleSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteSampleSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScheduleRequest() (request *DeleteScheduleRequest) {
    request = &DeleteScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteSchedule")
    
    
    return
}

func NewDeleteScheduleResponse() (response *DeleteScheduleResponse) {
    response = &DeleteScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSchedule
// This API is used to delete a scheme.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSchedule(request *DeleteScheduleRequest) (response *DeleteScheduleResponse, err error) {
    return c.DeleteScheduleWithContext(context.Background(), request)
}

// DeleteSchedule
// This API is used to delete a scheme.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteScheduleWithContext(ctx context.Context, request *DeleteScheduleRequest) (response *DeleteScheduleResponse, err error) {
    if request == nil {
        request = NewDeleteScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteSchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmartSubtitleTemplateRequest() (request *DeleteSmartSubtitleTemplateRequest) {
    request = &DeleteSmartSubtitleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DeleteSmartSubtitleTemplate")
    
    
    return
}

func NewDeleteSmartSubtitleTemplateResponse() (response *DeleteSmartSubtitleTemplateResponse) {
    response = &DeleteSmartSubtitleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSmartSubtitleTemplate
// This API is used to delete a user-defined smart subtitle template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSmartSubtitleTemplate(request *DeleteSmartSubtitleTemplateRequest) (response *DeleteSmartSubtitleTemplateResponse, err error) {
    return c.DeleteSmartSubtitleTemplateWithContext(context.Background(), request)
}

// DeleteSmartSubtitleTemplate
// This API is used to delete a user-defined smart subtitle template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSmartSubtitleTemplateWithContext(ctx context.Context, request *DeleteSmartSubtitleTemplateRequest) (response *DeleteSmartSubtitleTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSmartSubtitleTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteSmartSubtitleTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSmartSubtitleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSmartSubtitleTemplateResponse()
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
    return c.DeleteSnapshotByTimeOffsetTemplateWithContext(context.Background(), request)
}

// DeleteSnapshotByTimeOffsetTemplate
// This API is used to delete a custom time point screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteSnapshotByTimeOffsetTemplateWithContext(ctx context.Context, request *DeleteSnapshotByTimeOffsetTemplateRequest) (response *DeleteSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotByTimeOffsetTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteSnapshotByTimeOffsetTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteTranscodeTemplate(request *DeleteTranscodeTemplateRequest) (response *DeleteTranscodeTemplateResponse, err error) {
    return c.DeleteTranscodeTemplateWithContext(context.Background(), request)
}

// DeleteTranscodeTemplate
// This API is used to delete a custom transcoding template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteTranscodeTemplateWithContext(ctx context.Context, request *DeleteTranscodeTemplateRequest) (response *DeleteTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteWatermarkTemplateWithContext(context.Background(), request)
}

// DeleteWatermarkTemplate
// This API is used to delete a custom watermarking template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWatermarkTemplateWithContext(ctx context.Context, request *DeleteWatermarkTemplateRequest) (response *DeleteWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteWordSamplesWithContext(context.Background(), request)
}

// DeleteWordSamples
// This API is used to delete keyword samples in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteWordSamplesWithContext(ctx context.Context, request *DeleteWordSamplesRequest) (response *DeleteWordSamplesResponse, err error) {
    if request == nil {
        request = NewDeleteWordSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteWordSamples")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWordSamples require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteWorkflowWithContext(context.Background(), request)
}

// DeleteWorkflow
// This API is used to delete a workflow. An enabled workflow must be disabled before it can be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteWorkflowWithContext(ctx context.Context, request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DeleteWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflow require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeAIAnalysisTemplatesWithContext(context.Background(), request)
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
func (c *Client) DescribeAIAnalysisTemplatesWithContext(ctx context.Context, request *DescribeAIAnalysisTemplatesRequest) (response *DescribeAIAnalysisTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIAnalysisTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAIAnalysisTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIAnalysisTemplates require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeAIRecognitionTemplatesWithContext(context.Background(), request)
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
func (c *Client) DescribeAIRecognitionTemplatesWithContext(ctx context.Context, request *DescribeAIRecognitionTemplatesRequest) (response *DescribeAIRecognitionTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIRecognitionTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAIRecognitionTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIRecognitionTemplates require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeAdaptiveDynamicStreamingTemplatesWithContext(context.Background(), request)
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
func (c *Client) DescribeAdaptiveDynamicStreamingTemplatesWithContext(ctx context.Context, request *DescribeAdaptiveDynamicStreamingTemplatesRequest) (response *DescribeAdaptiveDynamicStreamingTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAdaptiveDynamicStreamingTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAdaptiveDynamicStreamingTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAdaptiveDynamicStreamingTemplates require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeAnimatedGraphicsTemplates(request *DescribeAnimatedGraphicsTemplatesRequest) (response *DescribeAnimatedGraphicsTemplatesResponse, err error) {
    return c.DescribeAnimatedGraphicsTemplatesWithContext(context.Background(), request)
}

// DescribeAnimatedGraphicsTemplates
// This API is used to query the list of animated image generating templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeAnimatedGraphicsTemplatesWithContext(ctx context.Context, request *DescribeAnimatedGraphicsTemplatesRequest) (response *DescribeAnimatedGraphicsTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAnimatedGraphicsTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAnimatedGraphicsTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAnimatedGraphicsTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAnimatedGraphicsTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsrHotwordsRequest() (request *DescribeAsrHotwordsRequest) {
    request = &DescribeAsrHotwordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAsrHotwords")
    
    
    return
}

func NewDescribeAsrHotwordsResponse() (response *DescribeAsrHotwordsResponse) {
    response = &DescribeAsrHotwordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAsrHotwords
// This API is used to query a smart subtitle hotword lexicon.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAsrHotwords(request *DescribeAsrHotwordsRequest) (response *DescribeAsrHotwordsResponse, err error) {
    return c.DescribeAsrHotwordsWithContext(context.Background(), request)
}

// DescribeAsrHotwords
// This API is used to query a smart subtitle hotword lexicon.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAsrHotwordsWithContext(ctx context.Context, request *DescribeAsrHotwordsRequest) (response *DescribeAsrHotwordsResponse, err error) {
    if request == nil {
        request = NewDescribeAsrHotwordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAsrHotwords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsrHotwords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsrHotwordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsrHotwordsListRequest() (request *DescribeAsrHotwordsListRequest) {
    request = &DescribeAsrHotwordsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeAsrHotwordsList")
    
    
    return
}

func NewDescribeAsrHotwordsListResponse() (response *DescribeAsrHotwordsListResponse) {
    response = &DescribeAsrHotwordsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAsrHotwordsList
// This API is used to obtain the hotword lexicon list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAsrHotwordsList(request *DescribeAsrHotwordsListRequest) (response *DescribeAsrHotwordsListResponse, err error) {
    return c.DescribeAsrHotwordsListWithContext(context.Background(), request)
}

// DescribeAsrHotwordsList
// This API is used to obtain the hotword lexicon list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAsrHotwordsListWithContext(ctx context.Context, request *DescribeAsrHotwordsListRequest) (response *DescribeAsrHotwordsListResponse, err error) {
    if request == nil {
        request = NewDescribeAsrHotwordsListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeAsrHotwordsList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsrHotwordsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsrHotwordsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchTaskDetailRequest() (request *DescribeBatchTaskDetailRequest) {
    request = &DescribeBatchTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeBatchTaskDetail")
    
    
    return
}

func NewDescribeBatchTaskDetailResponse() (response *DescribeBatchTaskDetailResponse) {
    response = &DescribeBatchTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBatchTaskDetail
// This API is used to query the details of the task execution status and results by task ID (tasks submitted within the last 7 days can be queried).
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBatchTaskDetail(request *DescribeBatchTaskDetailRequest) (response *DescribeBatchTaskDetailResponse, err error) {
    return c.DescribeBatchTaskDetailWithContext(context.Background(), request)
}

// DescribeBatchTaskDetail
// This API is used to query the details of the task execution status and results by task ID (tasks submitted within the last 7 days can be queried).
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBatchTaskDetailWithContext(ctx context.Context, request *DescribeBatchTaskDetailRequest) (response *DescribeBatchTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBatchTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeBatchTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchTaskDetailResponse()
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
// This API is used to query content moderation templates by template ID. Both custom and preset templates that match the template IDs passed in will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeContentReviewTemplates(request *DescribeContentReviewTemplatesRequest) (response *DescribeContentReviewTemplatesResponse, err error) {
    return c.DescribeContentReviewTemplatesWithContext(context.Background(), request)
}

// DescribeContentReviewTemplates
// This API is used to query content moderation templates by template ID. Both custom and preset templates that match the template IDs passed in will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeContentReviewTemplatesWithContext(ctx context.Context, request *DescribeContentReviewTemplatesRequest) (response *DescribeContentReviewTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeContentReviewTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeContentReviewTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContentReviewTemplates require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeImageSpriteTemplates(request *DescribeImageSpriteTemplatesRequest) (response *DescribeImageSpriteTemplatesResponse, err error) {
    return c.DescribeImageSpriteTemplatesWithContext(context.Background(), request)
}

// DescribeImageSpriteTemplates
// This API is used to query the list of image sprite generating templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeImageSpriteTemplatesWithContext(ctx context.Context, request *DescribeImageSpriteTemplatesRequest) (response *DescribeImageSpriteTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeImageSpriteTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeImageSpriteTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageSpriteTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageSpriteTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageTaskDetailRequest() (request *DescribeImageTaskDetailRequest) {
    request = &DescribeImageTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeImageTaskDetail")
    
    
    return
}

func NewDescribeImageTaskDetailResponse() (response *DescribeImageTaskDetailResponse) {
    response = &DescribeImageTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageTaskDetail
// This API is used to query the details of the task execution status and results by task ID (tasks submitted within the last 7 days can be queried).
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageTaskDetail(request *DescribeImageTaskDetailRequest) (response *DescribeImageTaskDetailResponse, err error) {
    return c.DescribeImageTaskDetailWithContext(context.Background(), request)
}

// DescribeImageTaskDetail
// This API is used to query the details of the task execution status and results by task ID (tasks submitted within the last 7 days can be queried).
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageTaskDetailWithContext(ctx context.Context, request *DescribeImageTaskDetailRequest) (response *DescribeImageTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeImageTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeImageTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordTemplatesRequest() (request *DescribeLiveRecordTemplatesRequest) {
    request = &DescribeLiveRecordTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeLiveRecordTemplates")
    
    
    return
}

func NewDescribeLiveRecordTemplatesResponse() (response *DescribeLiveRecordTemplatesResponse) {
    response = &DescribeLiveRecordTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveRecordTemplates
// This API is used to get a live recording template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeLiveRecordTemplates(request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    return c.DescribeLiveRecordTemplatesWithContext(context.Background(), request)
}

// DescribeLiveRecordTemplates
// This API is used to get a live recording template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeLiveRecordTemplatesWithContext(ctx context.Context, request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeLiveRecordTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveRecordTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveRecordTemplatesResponse()
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
    return c.DescribeMediaMetaDataWithContext(context.Background(), request)
}

// DescribeMediaMetaData
// This API is used to get the metadata of media, such as video image width/height, codec, length, and frame rate.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) DescribeMediaMetaDataWithContext(ctx context.Context, request *DescribeMediaMetaDataRequest) (response *DescribeMediaMetaDataResponse, err error) {
    if request == nil {
        request = NewDescribeMediaMetaDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeMediaMetaData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMediaMetaData require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribePersonSamplesWithContext(context.Background(), request)
}

// DescribePersonSamples
// This API is used to query the information of image samples. It supports paginated queries by image ID, name, and tag.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePersonSamplesWithContext(ctx context.Context, request *DescribePersonSamplesRequest) (response *DescribePersonSamplesResponse, err error) {
    if request == nil {
        request = NewDescribePersonSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribePersonSamples")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityControlTemplatesRequest() (request *DescribeQualityControlTemplatesRequest) {
    request = &DescribeQualityControlTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeQualityControlTemplates")
    
    
    return
}

func NewDescribeQualityControlTemplatesResponse() (response *DescribeQualityControlTemplatesResponse) {
    response = &DescribeQualityControlTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQualityControlTemplates
// This API is used to query custom media quality inspection templates, supporting paged queries by conditions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeQualityControlTemplates(request *DescribeQualityControlTemplatesRequest) (response *DescribeQualityControlTemplatesResponse, err error) {
    return c.DescribeQualityControlTemplatesWithContext(context.Background(), request)
}

// DescribeQualityControlTemplates
// This API is used to query custom media quality inspection templates, supporting paged queries by conditions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeQualityControlTemplatesWithContext(ctx context.Context, request *DescribeQualityControlTemplatesRequest) (response *DescribeQualityControlTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeQualityControlTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeQualityControlTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQualityControlTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQualityControlTemplatesResponse()
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
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeSampleSnapshotTemplates(request *DescribeSampleSnapshotTemplatesRequest) (response *DescribeSampleSnapshotTemplatesResponse, err error) {
    return c.DescribeSampleSnapshotTemplatesWithContext(context.Background(), request)
}

// DescribeSampleSnapshotTemplates
// This API is used to query the list of sampled screencapturing templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeSampleSnapshotTemplatesWithContext(ctx context.Context, request *DescribeSampleSnapshotTemplatesRequest) (response *DescribeSampleSnapshotTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSampleSnapshotTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeSampleSnapshotTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleSnapshotTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleSnapshotTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSchedulesRequest() (request *DescribeSchedulesRequest) {
    request = &DescribeSchedulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeSchedules")
    
    
    return
}

func NewDescribeSchedulesResponse() (response *DescribeSchedulesResponse) {
    response = &DescribeSchedulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSchedules
// This API is used to query a scheme.
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSchedules(request *DescribeSchedulesRequest) (response *DescribeSchedulesResponse, err error) {
    return c.DescribeSchedulesWithContext(context.Background(), request)
}

// DescribeSchedules
// This API is used to query a scheme.
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSchedulesWithContext(ctx context.Context, request *DescribeSchedulesRequest) (response *DescribeSchedulesResponse, err error) {
    if request == nil {
        request = NewDescribeSchedulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeSchedules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSchedules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSchedulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmartSubtitleTemplatesRequest() (request *DescribeSmartSubtitleTemplatesRequest) {
    request = &DescribeSmartSubtitleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeSmartSubtitleTemplates")
    
    
    return
}

func NewDescribeSmartSubtitleTemplatesResponse() (response *DescribeSmartSubtitleTemplatesResponse) {
    response = &DescribeSmartSubtitleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSmartSubtitleTemplates
// This API is used to obtain the list of smart subtitle templates based on template unique identifier. The returned result includes all matching user-defined smart subtitle templates and system preset smart subtitle templates.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeSmartSubtitleTemplates(request *DescribeSmartSubtitleTemplatesRequest) (response *DescribeSmartSubtitleTemplatesResponse, err error) {
    return c.DescribeSmartSubtitleTemplatesWithContext(context.Background(), request)
}

// DescribeSmartSubtitleTemplates
// This API is used to obtain the list of smart subtitle templates based on template unique identifier. The returned result includes all matching user-defined smart subtitle templates and system preset smart subtitle templates.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeSmartSubtitleTemplatesWithContext(ctx context.Context, request *DescribeSmartSubtitleTemplatesRequest) (response *DescribeSmartSubtitleTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSmartSubtitleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeSmartSubtitleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmartSubtitleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmartSubtitleTemplatesResponse()
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
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeSnapshotByTimeOffsetTemplates(request *DescribeSnapshotByTimeOffsetTemplatesRequest) (response *DescribeSnapshotByTimeOffsetTemplatesResponse, err error) {
    return c.DescribeSnapshotByTimeOffsetTemplatesWithContext(context.Background(), request)
}

// DescribeSnapshotByTimeOffsetTemplates
// This API is used to query the list of time point screencapturing templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) DescribeSnapshotByTimeOffsetTemplatesWithContext(ctx context.Context, request *DescribeSnapshotByTimeOffsetTemplatesRequest) (response *DescribeSnapshotByTimeOffsetTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotByTimeOffsetTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeSnapshotByTimeOffsetTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotByTimeOffsetTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotByTimeOffsetTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkSecurityGroupRequest() (request *DescribeStreamLinkSecurityGroupRequest) {
    request = &DescribeStreamLinkSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DescribeStreamLinkSecurityGroup")
    
    
    return
}

func NewDescribeStreamLinkSecurityGroupResponse() (response *DescribeStreamLinkSecurityGroupResponse) {
    response = &DescribeStreamLinkSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamLinkSecurityGroup
// This API is used to query a security group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLinkSecurityGroup(request *DescribeStreamLinkSecurityGroupRequest) (response *DescribeStreamLinkSecurityGroupResponse, err error) {
    return c.DescribeStreamLinkSecurityGroupWithContext(context.Background(), request)
}

// DescribeStreamLinkSecurityGroup
// This API is used to query a security group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLinkSecurityGroupWithContext(ctx context.Context, request *DescribeStreamLinkSecurityGroupRequest) (response *DescribeStreamLinkSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeStreamLinkSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkSecurityGroupResponse()
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
    return c.DescribeTaskDetailWithContext(context.Background(), request)
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
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskDetail require credential")
    }

    request.SetContext(ctx)
    
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
// * This API is used to query tasks.
//
// * If the data is large, one API call may not be able to obtain all the tasks in the query. You can use the `ScrollToken` parameter to query tasks with multiple calls.
//
// * Only tasks in the last seven days (168 hours) can be queried.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// * This API is used to query tasks.
//
// * If the data is large, one API call may not be able to obtain all the tasks in the query. You can use the `ScrollToken` parameter to query tasks with multiple calls.
//
// * Only tasks in the last seven days (168 hours) can be queried.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeTranscodeTemplates(request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    return c.DescribeTranscodeTemplatesWithContext(context.Background(), request)
}

// DescribeTranscodeTemplates
// This API is used to get the list of transcoding templates based on unique template ID. The return result includes all eligible custom and [preset transcoding templates](https://intl.cloud.tencent.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeTranscodeTemplatesWithContext(ctx context.Context, request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeTranscodeTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeWatermarkTemplatesWithContext(context.Background(), request)
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
func (c *Client) DescribeWatermarkTemplatesWithContext(ctx context.Context, request *DescribeWatermarkTemplatesRequest) (response *DescribeWatermarkTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeWatermarkTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeWatermarkTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWatermarkTemplates require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeWordSamplesWithContext(context.Background(), request)
}

// DescribeWordSamples
// This API is used to perform paged queries of keyword sample information by use case, keyword, and tag.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWordSamplesWithContext(ctx context.Context, request *DescribeWordSamplesRequest) (response *DescribeWordSamplesResponse, err error) {
    if request == nil {
        request = NewDescribeWordSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeWordSamples")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWordSamples require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeWorkflowsWithContext(context.Background(), request)
}

// DescribeWorkflows
// This API is used to get the list of workflow details by workflow ID.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeWorkflowsWithContext(ctx context.Context, request *DescribeWorkflowsRequest) (response *DescribeWorkflowsResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DescribeWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableScheduleRequest() (request *DisableScheduleRequest) {
    request = &DisableScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "DisableSchedule")
    
    
    return
}

func NewDisableScheduleResponse() (response *DisableScheduleResponse) {
    response = &DisableScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableSchedule
// This API is used to disable a scheme.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GETSOURCENOTIFY = "FailedOperation.GetSourceNotify"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  FAILEDOPERATION_SETSOURCENOTIFY = "FailedOperation.SetSourceNotify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DisableSchedule(request *DisableScheduleRequest) (response *DisableScheduleResponse, err error) {
    return c.DisableScheduleWithContext(context.Background(), request)
}

// DisableSchedule
// This API is used to disable a scheme.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GETSOURCENOTIFY = "FailedOperation.GetSourceNotify"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  FAILEDOPERATION_SETSOURCENOTIFY = "FailedOperation.SetSourceNotify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DisableScheduleWithContext(ctx context.Context, request *DisableScheduleRequest) (response *DisableScheduleResponse, err error) {
    if request == nil {
        request = NewDisableScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DisableSchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableScheduleResponse()
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
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DisableWorkflow(request *DisableWorkflowRequest) (response *DisableWorkflowResponse, err error) {
    return c.DisableWorkflowWithContext(context.Background(), request)
}

// DisableWorkflow
// This API is used to disable a workflow.
//
// error code that may be returned:
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DisableWorkflowWithContext(ctx context.Context, request *DisableWorkflowRequest) (response *DisableWorkflowResponse, err error) {
    if request == nil {
        request = NewDisableWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "DisableWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableWorkflow require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to edit a video to generate a new one. Editing features include:
//
//  
//
// 
//
// 1. **Editing task**: simple video editing, such as clipping and splicing.
//
// 1) Edit a file to generate a new video.
//
// 2) Splice multiple files to generate a new video.
//
// 3) Edit multiple files and then splice them to generate a new video.
//
// 
//
// 2. **Compositing task**: Generate a new video by describing information through APIs.
//
// 1) Multi-track (video, audio, and subtitles) and multi-type elements (video, image, audio, text, and empty).
//
// 2) Image level: mapping, zoom in/out, arbitrary rotation, mirroring, and more.
//
// 3) Audio level: volume control, fade in/out, mixing, and more.
//
// 4) Video level: transition, playback speed adjustment, splicing, clipping, subtitles, picture-in-picture, audio-video separation, entrance and exit animations, and more.
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditMedia(request *EditMediaRequest) (response *EditMediaResponse, err error) {
    return c.EditMediaWithContext(context.Background(), request)
}

// EditMedia
// This API is used to edit a video to generate a new one. Editing features include:
//
//  
//
// 
//
// 1. **Editing task**: simple video editing, such as clipping and splicing.
//
// 1) Edit a file to generate a new video.
//
// 2) Splice multiple files to generate a new video.
//
// 3) Edit multiple files and then splice them to generate a new video.
//
// 
//
// 2. **Compositing task**: Generate a new video by describing information through APIs.
//
// 1) Multi-track (video, audio, and subtitles) and multi-type elements (video, image, audio, text, and empty).
//
// 2) Image level: mapping, zoom in/out, arbitrary rotation, mirroring, and more.
//
// 3) Audio level: volume control, fade in/out, mixing, and more.
//
// 4) Video level: transition, playback speed adjustment, splicing, clipping, subtitles, picture-in-picture, audio-video separation, entrance and exit animations, and more.
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditMediaWithContext(ctx context.Context, request *EditMediaRequest) (response *EditMediaResponse, err error) {
    if request == nil {
        request = NewEditMediaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "EditMedia")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditMediaResponse()
    err = c.Send(request, response)
    return
}

func NewEnableScheduleRequest() (request *EnableScheduleRequest) {
    request = &EnableScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "EnableSchedule")
    
    
    return
}

func NewEnableScheduleResponse() (response *EnableScheduleResponse) {
    response = &EnableScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableSchedule
// This API is used to enable a scheme.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_GETSOURCENOTIFY = "FailedOperation.GetSourceNotify"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  FAILEDOPERATION_SETSOURCENOTIFY = "FailedOperation.SetSourceNotify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) EnableSchedule(request *EnableScheduleRequest) (response *EnableScheduleResponse, err error) {
    return c.EnableScheduleWithContext(context.Background(), request)
}

// EnableSchedule
// This API is used to enable a scheme.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"
//  FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_GETSOURCENOTIFY = "FailedOperation.GetSourceNotify"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  FAILEDOPERATION_SETSOURCENOTIFY = "FailedOperation.SetSourceNotify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) EnableScheduleWithContext(ctx context.Context, request *EnableScheduleRequest) (response *EnableScheduleResponse, err error) {
    if request == nil {
        request = NewEnableScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "EnableSchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableScheduleResponse()
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
    return c.EnableWorkflowWithContext(context.Background(), request)
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
func (c *Client) EnableWorkflowWithContext(ctx context.Context, request *EnableWorkflowRequest) (response *EnableWorkflowResponse, err error) {
    if request == nil {
        request = NewEnableWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "EnableWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableWorkflow require credential")
    }

    request.SetContext(ctx)
    
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
// This API is reserved for special circumstances. Do not use it unless you are directed to use it by technical support.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExecuteFunction(request *ExecuteFunctionRequest) (response *ExecuteFunctionResponse, err error) {
    return c.ExecuteFunctionWithContext(context.Background(), request)
}

// ExecuteFunction
// This API is reserved for special circumstances. Do not use it unless you are directed to use it by technical support.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExecuteFunctionWithContext(ctx context.Context, request *ExecuteFunctionRequest) (response *ExecuteFunctionResponse, err error) {
    if request == nil {
        request = NewExecuteFunctionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ExecuteFunction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteFunction require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ManageTaskWithContext(context.Background(), request)
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
func (c *Client) ManageTaskWithContext(ctx context.Context, request *ManageTaskRequest) (response *ManageTaskResponse, err error) {
    if request == nil {
        request = NewManageTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ManageTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageTask require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyAIAnalysisTemplateWithContext(context.Background(), request)
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
func (c *Client) ModifyAIAnalysisTemplateWithContext(ctx context.Context, request *ModifyAIAnalysisTemplateRequest) (response *ModifyAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIAnalysisTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAIAnalysisTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAIAnalysisTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAIRecognitionTemplate(request *ModifyAIRecognitionTemplateRequest) (response *ModifyAIRecognitionTemplateResponse, err error) {
    return c.ModifyAIRecognitionTemplateWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAIRecognitionTemplateWithContext(ctx context.Context, request *ModifyAIRecognitionTemplateRequest) (response *ModifyAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIRecognitionTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAIRecognitionTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAIRecognitionTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyAdaptiveDynamicStreamingTemplateWithContext(context.Background(), request)
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
func (c *Client) ModifyAdaptiveDynamicStreamingTemplateWithContext(ctx context.Context, request *ModifyAdaptiveDynamicStreamingTemplateRequest) (response *ModifyAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAdaptiveDynamicStreamingTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAdaptiveDynamicStreamingTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyAnimatedGraphicsTemplateWithContext(context.Background(), request)
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
func (c *Client) ModifyAnimatedGraphicsTemplateWithContext(ctx context.Context, request *ModifyAnimatedGraphicsTemplateRequest) (response *ModifyAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAnimatedGraphicsTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAnimatedGraphicsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAsrHotwordsRequest() (request *ModifyAsrHotwordsRequest) {
    request = &ModifyAsrHotwordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyAsrHotwords")
    
    
    return
}

func NewModifyAsrHotwordsResponse() (response *ModifyAsrHotwordsResponse) {
    response = &ModifyAsrHotwordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAsrHotwords
// This API is used to update a smart subtitle hotword lexicon.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSNOTEXIST = "InvalidParameterValue.HotWordsNotExist"
func (c *Client) ModifyAsrHotwords(request *ModifyAsrHotwordsRequest) (response *ModifyAsrHotwordsResponse, err error) {
    return c.ModifyAsrHotwordsWithContext(context.Background(), request)
}

// ModifyAsrHotwords
// This API is used to update a smart subtitle hotword lexicon.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HOTWORDSNOTEXIST = "InvalidParameterValue.HotWordsNotExist"
func (c *Client) ModifyAsrHotwordsWithContext(ctx context.Context, request *ModifyAsrHotwordsRequest) (response *ModifyAsrHotwordsResponse, err error) {
    if request == nil {
        request = NewModifyAsrHotwordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyAsrHotwords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAsrHotwords require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAsrHotwordsResponse()
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
// This API is used to modify a custom content moderation template.
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
    return c.ModifyContentReviewTemplateWithContext(context.Background(), request)
}

// ModifyContentReviewTemplate
// This API is used to modify a custom content moderation template.
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
func (c *Client) ModifyContentReviewTemplateWithContext(ctx context.Context, request *ModifyContentReviewTemplateRequest) (response *ModifyContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewModifyContentReviewTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyContentReviewTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyImageSpriteTemplate(request *ModifyImageSpriteTemplateRequest) (response *ModifyImageSpriteTemplateResponse, err error) {
    return c.ModifyImageSpriteTemplateWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyImageSpriteTemplateWithContext(ctx context.Context, request *ModifyImageSpriteTemplateRequest) (response *ModifyImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewModifyImageSpriteTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyImageSpriteTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveRecordTemplateRequest() (request *ModifyLiveRecordTemplateRequest) {
    request = &ModifyLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyLiveRecordTemplate")
    
    
    return
}

func NewModifyLiveRecordTemplateResponse() (response *ModifyLiveRecordTemplateResponse) {
    response = &ModifyLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveRecordTemplate
// This API is used to modify a live recording template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyLiveRecordTemplate(request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    return c.ModifyLiveRecordTemplateWithContext(context.Background(), request)
}

// ModifyLiveRecordTemplate
// This API is used to modify a live recording template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyLiveRecordTemplateWithContext(ctx context.Context, request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveRecordTemplateResponse()
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
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) ModifyPersonSample(request *ModifyPersonSampleRequest) (response *ModifyPersonSampleResponse, err error) {
    return c.ModifyPersonSampleWithContext(context.Background(), request)
}

// ModifyPersonSample
// This API is used to modify image samples by image ID. You can use it to modify the name and description of an image sample and add/delete/reset facial features or tags. There must be at least one image left after the deletion of facial features; otherwise, please reset instead of delete the facial features.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
func (c *Client) ModifyPersonSampleWithContext(ctx context.Context, request *ModifyPersonSampleRequest) (response *ModifyPersonSampleResponse, err error) {
    if request == nil {
        request = NewModifyPersonSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyPersonSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQualityControlTemplateRequest() (request *ModifyQualityControlTemplateRequest) {
    request = &ModifyQualityControlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifyQualityControlTemplate")
    
    
    return
}

func NewModifyQualityControlTemplateResponse() (response *ModifyQualityControlTemplateResponse) {
    response = &ModifyQualityControlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQualityControlTemplate
// This API is used to modify a media quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYDETECTITEM = "InvalidParameterValue.EmptyDetectItem"
//  INVALIDPARAMETERVALUE_UNKNOWNCATEGORY = "InvalidParameterValue.UnknownCategory"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyQualityControlTemplate(request *ModifyQualityControlTemplateRequest) (response *ModifyQualityControlTemplateResponse, err error) {
    return c.ModifyQualityControlTemplateWithContext(context.Background(), request)
}

// ModifyQualityControlTemplate
// This API is used to modify a media quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYDETECTITEM = "InvalidParameterValue.EmptyDetectItem"
//  INVALIDPARAMETERVALUE_UNKNOWNCATEGORY = "InvalidParameterValue.UnknownCategory"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyQualityControlTemplateWithContext(ctx context.Context, request *ModifyQualityControlTemplateRequest) (response *ModifyQualityControlTemplateResponse, err error) {
    if request == nil {
        request = NewModifyQualityControlTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyQualityControlTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQualityControlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQualityControlTemplateResponse()
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
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySampleSnapshotTemplate(request *ModifySampleSnapshotTemplateRequest) (response *ModifySampleSnapshotTemplateResponse, err error) {
    return c.ModifySampleSnapshotTemplateWithContext(context.Background(), request)
}

// ModifySampleSnapshotTemplate
// This API is used to modify a custom sampled screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySampleSnapshotTemplateWithContext(ctx context.Context, request *ModifySampleSnapshotTemplateRequest) (response *ModifySampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewModifySampleSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifySampleSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScheduleRequest() (request *ModifyScheduleRequest) {
    request = &ModifyScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifySchedule")
    
    
    return
}

func NewModifyScheduleResponse() (response *ModifyScheduleResponse) {
    response = &ModifyScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySchedule
// This API is used to modify a scheme.
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySchedule(request *ModifyScheduleRequest) (response *ModifyScheduleResponse, err error) {
    return c.ModifyScheduleWithContext(context.Background(), request)
}

// ModifySchedule
// This API is used to modify a scheme.
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyScheduleWithContext(ctx context.Context, request *ModifyScheduleRequest) (response *ModifyScheduleResponse, err error) {
    if request == nil {
        request = NewModifyScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifySchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmartSubtitleTemplateRequest() (request *ModifySmartSubtitleTemplateRequest) {
    request = &ModifySmartSubtitleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ModifySmartSubtitleTemplate")
    
    
    return
}

func NewModifySmartSubtitleTemplateResponse() (response *ModifySmartSubtitleTemplateResponse) {
    response = &ModifySmartSubtitleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySmartSubtitleTemplate
// This API is used to modify a user-defined smart subtitle template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSLIBRARYID = "InvalidParameterValue.AsrHotWordsLibraryId"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSSWITCH = "InvalidParameterValue.AsrHotWordsSwitch"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLETYPE = "InvalidParameterValue.SubtitleType"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  INVALIDPARAMETERVALUE_TRANSLATESWITCH = "InvalidParameterValue.TranslateSwitch"
//  INVALIDPARAMETERVALUE_VIDEOSRCLANGUAGE = "InvalidParameterValue.VideoSrcLanguage"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySmartSubtitleTemplate(request *ModifySmartSubtitleTemplateRequest) (response *ModifySmartSubtitleTemplateResponse, err error) {
    return c.ModifySmartSubtitleTemplateWithContext(context.Background(), request)
}

// ModifySmartSubtitleTemplate
// This API is used to modify a user-defined smart subtitle template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSLIBRARYID = "InvalidParameterValue.AsrHotWordsLibraryId"
//  INVALIDPARAMETERVALUE_ASRHOTWORDSSWITCH = "InvalidParameterValue.AsrHotWordsSwitch"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SUBTITLETYPE = "InvalidParameterValue.SubtitleType"
//  INVALIDPARAMETERVALUE_TRANSLATEDSTLANGUAGE = "InvalidParameterValue.TranslateDstLanguage"
//  INVALIDPARAMETERVALUE_TRANSLATESWITCH = "InvalidParameterValue.TranslateSwitch"
//  INVALIDPARAMETERVALUE_VIDEOSRCLANGUAGE = "InvalidParameterValue.VideoSrcLanguage"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySmartSubtitleTemplateWithContext(ctx context.Context, request *ModifySmartSubtitleTemplateRequest) (response *ModifySmartSubtitleTemplateResponse, err error) {
    if request == nil {
        request = NewModifySmartSubtitleTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifySmartSubtitleTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySmartSubtitleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySmartSubtitleTemplateResponse()
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
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySnapshotByTimeOffsetTemplate(request *ModifySnapshotByTimeOffsetTemplateRequest) (response *ModifySnapshotByTimeOffsetTemplateResponse, err error) {
    return c.ModifySnapshotByTimeOffsetTemplateWithContext(context.Background(), request)
}

// ModifySnapshotByTimeOffsetTemplate
// This API is used to modify a custom time point screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifySnapshotByTimeOffsetTemplateWithContext(ctx context.Context, request *ModifySnapshotByTimeOffsetTemplateRequest) (response *ModifySnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewModifySnapshotByTimeOffsetTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifySnapshotByTimeOffsetTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyTranscodeTemplate(request *ModifyTranscodeTemplateRequest) (response *ModifyTranscodeTemplateResponse, err error) {
    return c.ModifyTranscodeTemplateWithContext(context.Background(), request)
}

// ModifyTranscodeTemplate
// This API is used to modify a custom transcoding template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyTranscodeTemplateWithContext(ctx context.Context, request *ModifyTranscodeTemplateRequest) (response *ModifyTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyWatermarkTemplateWithContext(context.Background(), request)
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
func (c *Client) ModifyWatermarkTemplateWithContext(ctx context.Context, request *ModifyWatermarkTemplateRequest) (response *ModifyWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewModifyWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyWordSampleWithContext(context.Background(), request)
}

// ModifyWordSample
// This API is used to modify the use case and tag of a keyword. The keyword itself cannot be modified, but you can delete it and create another one if needed.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_WORD = "ResourceNotFound.Word"
func (c *Client) ModifyWordSampleWithContext(ctx context.Context, request *ModifyWordSampleRequest) (response *ModifyWordSampleResponse, err error) {
    if request == nil {
        request = NewModifyWordSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ModifyWordSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWordSample require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ParseLiveStreamProcessNotificationWithContext(context.Background(), request)
}

// ParseLiveStreamProcessNotification
// This API is used to parse the content of an MPS live stream processing event notification from the `msgBody` field in the message received from CMQ.
//
// Instead of initiating a video processing task, this API is used to help generate SDKs for various programming languages. You can parse the event notification based on the analytic function of the SDKs.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
func (c *Client) ParseLiveStreamProcessNotificationWithContext(ctx context.Context, request *ParseLiveStreamProcessNotificationRequest) (response *ParseLiveStreamProcessNotificationResponse, err error) {
    if request == nil {
        request = NewParseLiveStreamProcessNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ParseLiveStreamProcessNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseLiveStreamProcessNotification require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ParseNotificationWithContext(context.Background(), request)
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
func (c *Client) ParseNotificationWithContext(ctx context.Context, request *ParseNotificationRequest) (response *ParseNotificationResponse, err error) {
    if request == nil {
        request = NewParseNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ParseNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewProcessImageRequest() (request *ProcessImageRequest) {
    request = &ProcessImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mps", APIVersion, "ProcessImage")
    
    
    return
}

func NewProcessImageResponse() (response *ProcessImageResponse) {
    response = &ProcessImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessImage
// This API is used to initiate image processing, with features including:
//
// 1. Format conversion.
//
// 2. Image enhancement.
//
// 3. Image erasure.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessImage(request *ProcessImageRequest) (response *ProcessImageResponse, err error) {
    return c.ProcessImageWithContext(context.Background(), request)
}

// ProcessImage
// This API is used to initiate image processing, with features including:
//
// 1. Format conversion.
//
// 2. Image enhancement.
//
// 3. Image erasure.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessImageWithContext(ctx context.Context, request *ProcessImageRequest) (response *ProcessImageResponse, err error) {
    if request == nil {
        request = NewProcessImageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ProcessImage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessImageResponse()
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
// This API is used to initiate live stream processing tasks. Such tasks may include the following:
//
// 
//
// * Intelligent content moderation (detection of pornographic content in images and audio, detection of sensitive information)
//
// * Intelligent content recognition (face, full text, text keyword, full speech, speech keyword, real-time speech translation, object recognition, game event tracking)
//
// * Intelligent content analysis (real-time news splitting)
//
// * Quality control, including recognizing live stream format, checking audio/video content for flickering, blur, low light, overexposure, black bars, white bars, black screen, white screen, noise, pixelation, QR code, etc., and no-reference scoring.
//
// * Recording
//
// 
//
// HTTP callbacks are supported for live stream processing events. Notifications can also be written in real time to and read from a CMQ queue. The output files of processing tasks are saved to the storage you specify.
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) ProcessLiveStream(request *ProcessLiveStreamRequest) (response *ProcessLiveStreamResponse, err error) {
    return c.ProcessLiveStreamWithContext(context.Background(), request)
}

// ProcessLiveStream
// This API is used to initiate live stream processing tasks. Such tasks may include the following:
//
// 
//
// * Intelligent content moderation (detection of pornographic content in images and audio, detection of sensitive information)
//
// * Intelligent content recognition (face, full text, text keyword, full speech, speech keyword, real-time speech translation, object recognition, game event tracking)
//
// * Intelligent content analysis (real-time news splitting)
//
// * Quality control, including recognizing live stream format, checking audio/video content for flickering, blur, low light, overexposure, black bars, white bars, black screen, white screen, noise, pixelation, QR code, etc., and no-reference scoring.
//
// * Recording
//
// 
//
// HTTP callbacks are supported for live stream processing events. Notifications can also be written in real time to and read from a CMQ queue. The output files of processing tasks are saved to the storage you specify.
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) ProcessLiveStreamWithContext(ctx context.Context, request *ProcessLiveStreamRequest) (response *ProcessLiveStreamResponse, err error) {
    if request == nil {
        request = NewProcessLiveStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ProcessLiveStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessLiveStream require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to initiate a processing task for video URLs or media files in COS. Features include:
//
// 1. Video transcoding (standard transcoding, TSC transcoding, and audio/video enhancement).
//
// 2. Video-to-GIF conversion.
//
// 3. Time point screenshot of videos.
//
// 4. Sampled screenshot of videos.
//
// 5. Image sprite of videos.
//
// 6. Adaptive bitrate stream conversion for videos.
//
// 7. Intelligent content moderation (such as pornography detection and sensitive information detection).
//
// 8. Intelligent content analysis (such as tag, category, cover, frame tag, video splitting, highlight, opening and ending clips, and game tracking).
//
// 9. Intelligent content recognition (such as human face, full text, text keyword, full speech, speech keyword, speech translation, and object).
//
// 10. Media quality inspection, such as media format diagnosis, audio/visual content moderation (such as jitter, blur, low lighting, overexposure, black edge, white edge, black screen, white screen, screen glitch, noise, mosaic, and QR code), and no-reference quality scoring.
//
// 11. Smart subtitle (such as full speech, speech hotword, and speech translation).
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessMedia(request *ProcessMediaRequest) (response *ProcessMediaResponse, err error) {
    return c.ProcessMediaWithContext(context.Background(), request)
}

// ProcessMedia
// This API is used to initiate a processing task for video URLs or media files in COS. Features include:
//
// 1. Video transcoding (standard transcoding, TSC transcoding, and audio/video enhancement).
//
// 2. Video-to-GIF conversion.
//
// 3. Time point screenshot of videos.
//
// 4. Sampled screenshot of videos.
//
// 5. Image sprite of videos.
//
// 6. Adaptive bitrate stream conversion for videos.
//
// 7. Intelligent content moderation (such as pornography detection and sensitive information detection).
//
// 8. Intelligent content analysis (such as tag, category, cover, frame tag, video splitting, highlight, opening and ending clips, and game tracking).
//
// 9. Intelligent content recognition (such as human face, full text, text keyword, full speech, speech keyword, speech translation, and object).
//
// 10. Media quality inspection, such as media format diagnosis, audio/visual content moderation (such as jitter, blur, low lighting, overexposure, black edge, white edge, black screen, white screen, screen glitch, noise, mosaic, and QR code), and no-reference quality scoring.
//
// 11. Smart subtitle (such as full speech, speech hotword, and speech translation).
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessMediaWithContext(ctx context.Context, request *ProcessMediaRequest) (response *ProcessMediaResponse, err error) {
    if request == nil {
        request = NewProcessMediaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ProcessMedia")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMedia require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ResetWorkflowWithContext(context.Background(), request)
}

// ResetWorkflow
// This API is used to reset an existing workflow that is disabled.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ResetWorkflowWithContext(ctx context.Context, request *ResetWorkflowRequest) (response *ResetWorkflowResponse, err error) {
    if request == nil {
        request = NewResetWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mps", APIVersion, "ResetWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetWorkflowResponse()
    err = c.Send(request, response)
    return
}
