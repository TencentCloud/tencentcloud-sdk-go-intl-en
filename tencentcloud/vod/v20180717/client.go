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

package v20180717

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-07-17"

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


func NewApplyUploadRequest() (request *ApplyUploadRequest) {
    request = &ApplyUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ApplyUpload")
    
    
    return
}

func NewApplyUploadResponse() (response *ApplyUploadResponse) {
    response = &ApplyUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyUpload
// * This API is used to apply for uploading a media file (and cover file) to VOD and obtain the metadata of file storage (including upload path and upload signature) for subsequent use by the uploading API.
//
// * For the detailed upload process, please see [Overview of Upload from Client](https://intl.cloud.tencent.com/document/product/266/9759?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXPIRETIME = "InvalidParameter.ExpireTime"
//  INVALIDPARAMETERVALUE_COVERTYPE = "InvalidParameterValue.CoverType"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_MEDIATYPE = "InvalidParameterValue.MediaType"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ApplyUpload(request *ApplyUploadRequest) (response *ApplyUploadResponse, err error) {
    return c.ApplyUploadWithContext(context.Background(), request)
}

// ApplyUpload
// * This API is used to apply for uploading a media file (and cover file) to VOD and obtain the metadata of file storage (including upload path and upload signature) for subsequent use by the uploading API.
//
// * For the detailed upload process, please see [Overview of Upload from Client](https://intl.cloud.tencent.com/document/product/266/9759?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXPIRETIME = "InvalidParameter.ExpireTime"
//  INVALIDPARAMETERVALUE_COVERTYPE = "InvalidParameterValue.CoverType"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_MEDIATYPE = "InvalidParameterValue.MediaType"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ApplyUploadWithContext(ctx context.Context, request *ApplyUploadRequest) (response *ApplyUploadResponse, err error) {
    if request == nil {
        request = NewApplyUploadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyUpload require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyUploadResponse()
    err = c.Send(request, response)
    return
}

func NewAttachMediaSubtitlesRequest() (request *AttachMediaSubtitlesRequest) {
    request = &AttachMediaSubtitlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "AttachMediaSubtitles")
    
    
    return
}

func NewAttachMediaSubtitlesResponse() (response *AttachMediaSubtitlesResponse) {
    response = &AttachMediaSubtitlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachMediaSubtitles
// This API is used to associate/disassociate subtitles with/from a media file of a specific adaptive bitrate streaming template ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AttachMediaSubtitles(request *AttachMediaSubtitlesRequest) (response *AttachMediaSubtitlesResponse, err error) {
    return c.AttachMediaSubtitlesWithContext(context.Background(), request)
}

// AttachMediaSubtitles
// This API is used to associate/disassociate subtitles with/from a media file of a specific adaptive bitrate streaming template ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AttachMediaSubtitlesWithContext(ctx context.Context, request *AttachMediaSubtitlesRequest) (response *AttachMediaSubtitlesResponse, err error) {
    if request == nil {
        request = NewAttachMediaSubtitlesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachMediaSubtitles require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachMediaSubtitlesResponse()
    err = c.Send(request, response)
    return
}

func NewCloneCDNDomainRequest() (request *CloneCDNDomainRequest) {
    request = &CloneCDNDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CloneCDNDomain")
    
    
    return
}

func NewCloneCDNDomainResponse() (response *CloneCDNDomainResponse) {
    response = &CloneCDNDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloneCDNDomain
// Clone CDN Domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CloneCDNDomain(request *CloneCDNDomainRequest) (response *CloneCDNDomainResponse, err error) {
    return c.CloneCDNDomainWithContext(context.Background(), request)
}

// CloneCDNDomain
// Clone CDN Domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CloneCDNDomainWithContext(ctx context.Context, request *CloneCDNDomainRequest) (response *CloneCDNDomainResponse, err error) {
    if request == nil {
        request = NewCloneCDNDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneCDNDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneCDNDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCommitUploadRequest() (request *CommitUploadRequest) {
    request = &CommitUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CommitUpload")
    
    
    return
}

func NewCommitUploadResponse() (response *CommitUploadResponse) {
    response = &CommitUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CommitUpload
// This API is used to confirm the result of uploading a media file (and cover file) to VOD, store the media information, and return the playback address and ID of the file.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  INVALIDPARAMETERVALUE_VODSESSIONKEY = "InvalidParameterValue.VodSessionKey"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CommitUpload(request *CommitUploadRequest) (response *CommitUploadResponse, err error) {
    return c.CommitUploadWithContext(context.Background(), request)
}

// CommitUpload
// This API is used to confirm the result of uploading a media file (and cover file) to VOD, store the media information, and return the playback address and ID of the file.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  INVALIDPARAMETERVALUE_VODSESSIONKEY = "InvalidParameterValue.VodSessionKey"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CommitUploadWithContext(ctx context.Context, request *CommitUploadRequest) (response *CommitUploadResponse, err error) {
    if request == nil {
        request = NewCommitUploadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CommitUpload require credential")
    }

    request.SetContext(ctx)
    
    response = NewCommitUploadResponse()
    err = c.Send(request, response)
    return
}

func NewComposeMediaRequest() (request *ComposeMediaRequest) {
    request = &ComposeMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ComposeMedia")
    
    
    return
}

func NewComposeMediaResponse() (response *ComposeMediaResponse) {
    response = &ComposeMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ComposeMedia
// This API is used to compose a media file. You can use it to do the following:
//
// 
//
// 1. **Rotation/Flipping**: Rotate a video or image by a specific angle or flip a video or image.
//
// 2. **Audio control**: Increase/Lower the volume of an audio/video file or mute an audio/video file.
//
// 3. **Overlaying**: Overlay videos/images in a specified sequence to achieve the picture-in-picture effect.
//
// 4. **Audio mixing**: Mix the audios of audio/video files.
//
// 5 **Audio extraction**: Extract audio from a video.
//
// 6. **Clipping**: Clip segments from audio/video files according to a specified start and end time.
//
// 7. **Splicing**: Splice videos/audios/images in a specified sequence.
//
// 8. **Transition**: Add transition effects between video segments or images that are spliced together.
//
// 
//
// The output file is in MP4 or MP3 format. In the callback for media composition, the event type is [ComposeMediaComplete](https://intl.cloud.tencent.com/document/product/266/43000?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_UPLOADCOSFAIL = "FailedOperation.UploadCosFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETFILEINFOERROR = "InternalError.GetFileInfoError"
//  INTERNALERROR_GETMEDIALISTERROR = "InternalError.GetMediaListError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLIPDURATION = "InvalidParameterValue.ClipDuration"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ComposeMedia(request *ComposeMediaRequest) (response *ComposeMediaResponse, err error) {
    return c.ComposeMediaWithContext(context.Background(), request)
}

// ComposeMedia
// This API is used to compose a media file. You can use it to do the following:
//
// 
//
// 1. **Rotation/Flipping**: Rotate a video or image by a specific angle or flip a video or image.
//
// 2. **Audio control**: Increase/Lower the volume of an audio/video file or mute an audio/video file.
//
// 3. **Overlaying**: Overlay videos/images in a specified sequence to achieve the picture-in-picture effect.
//
// 4. **Audio mixing**: Mix the audios of audio/video files.
//
// 5 **Audio extraction**: Extract audio from a video.
//
// 6. **Clipping**: Clip segments from audio/video files according to a specified start and end time.
//
// 7. **Splicing**: Splice videos/audios/images in a specified sequence.
//
// 8. **Transition**: Add transition effects between video segments or images that are spliced together.
//
// 
//
// The output file is in MP4 or MP3 format. In the callback for media composition, the event type is [ComposeMediaComplete](https://intl.cloud.tencent.com/document/product/266/43000?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_UPLOADCOSFAIL = "FailedOperation.UploadCosFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETFILEINFOERROR = "InternalError.GetFileInfoError"
//  INTERNALERROR_GETMEDIALISTERROR = "InternalError.GetMediaListError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLIPDURATION = "InvalidParameterValue.ClipDuration"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ComposeMediaWithContext(ctx context.Context, request *ComposeMediaRequest) (response *ComposeMediaResponse, err error) {
    if request == nil {
        request = NewComposeMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ComposeMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewComposeMediaResponse()
    err = c.Send(request, response)
    return
}

func NewConfirmEventsRequest() (request *ConfirmEventsRequest) {
    request = &ConfirmEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ConfirmEvents")
    
    
    return
}

func NewConfirmEventsResponse() (response *ConfirmEventsResponse) {
    response = &ConfirmEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfirmEvents
// * After the `PullEvents` API is called to get an event, this API must be called to confirm that the message has been received;
//
// * After the event handler is obtained, the validity period of waiting for confirmation is 30 seconds. If the wait exceeds 30 seconds, a parameter error will be reported (4000);
//
// * For more information, please see the [reliable callback](https://intl.cloud.tencent.com/document/product/266/33779?from_cn_redirect=1#.E5.8F.AF.E9.9D.A0.E5.9B.9E.E8.B0.83) of event notification.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ConfirmEvents(request *ConfirmEventsRequest) (response *ConfirmEventsResponse, err error) {
    return c.ConfirmEventsWithContext(context.Background(), request)
}

// ConfirmEvents
// * After the `PullEvents` API is called to get an event, this API must be called to confirm that the message has been received;
//
// * After the event handler is obtained, the validity period of waiting for confirmation is 30 seconds. If the wait exceeds 30 seconds, a parameter error will be reported (4000);
//
// * For more information, please see the [reliable callback](https://intl.cloud.tencent.com/document/product/266/33779?from_cn_redirect=1#.E5.8F.AF.E9.9D.A0.E5.9B.9E.E8.B0.83) of event notification.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ConfirmEventsWithContext(ctx context.Context, request *ConfirmEventsRequest) (response *ConfirmEventsResponse, err error) {
    if request == nil {
        request = NewConfirmEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfirmEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfirmEventsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIAnalysisTemplateRequest() (request *CreateAIAnalysisTemplateRequest) {
    request = &CreateAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAIAnalysisTemplate")
    
    
    return
}

func NewCreateAIAnalysisTemplateResponse() (response *CreateAIAnalysisTemplateResponse) {
    response = &CreateAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAIAnalysisTemplate
// This API is used to create a custom video content analysis template. Up to 50 templates can be created.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLASSIFCATIONCONFIGURE = "InvalidParameterValue.ClassifcationConfigure"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COVERCONFIGURE = "InvalidParameterValue.CoverConfigure"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_FRAMETAGCONFIGURE = "InvalidParameterValue.FrameTagConfigure"
//  INVALIDPARAMETERVALUE_HIGHLIGHTCONFIGURE = "InvalidParameterValue.HighlightConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAIAnalysisTemplate(request *CreateAIAnalysisTemplateRequest) (response *CreateAIAnalysisTemplateResponse, err error) {
    return c.CreateAIAnalysisTemplateWithContext(context.Background(), request)
}

// CreateAIAnalysisTemplate
// This API is used to create a custom video content analysis template. Up to 50 templates can be created.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLASSIFCATIONCONFIGURE = "InvalidParameterValue.ClassifcationConfigure"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_COVERCONFIGURE = "InvalidParameterValue.CoverConfigure"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_FRAMETAGCONFIGURE = "InvalidParameterValue.FrameTagConfigure"
//  INVALIDPARAMETERVALUE_HIGHLIGHTCONFIGURE = "InvalidParameterValue.HighlightConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAIAnalysisTemplateWithContext(ctx context.Context, request *CreateAIAnalysisTemplateRequest) (response *CreateAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIAnalysisTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAIRecognitionTemplate")
    
    
    return
}

func NewCreateAIRecognitionTemplateResponse() (response *CreateAIRecognitionTemplateResponse) {
    response = &CreateAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAIRecognitionTemplate
// This API is used to create a custom video content recognition template. Up to 50 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SCREENSHOTINTERVAL = "InvalidParameterValue.ScreenshotInterval"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAIRecognitionTemplate(request *CreateAIRecognitionTemplateRequest) (response *CreateAIRecognitionTemplateResponse, err error) {
    return c.CreateAIRecognitionTemplateWithContext(context.Background(), request)
}

// CreateAIRecognitionTemplate
// This API is used to create a custom video content recognition template. Up to 50 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"
//  INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"
//  INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"
//  INVALIDPARAMETERVALUE_SCREENSHOTINTERVAL = "InvalidParameterValue.ScreenshotInterval"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAIRecognitionTemplateWithContext(ctx context.Context, request *CreateAIRecognitionTemplateRequest) (response *CreateAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAIRecognitionTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAdaptiveDynamicStreamingTemplate")
    
    
    return
}

func NewCreateAdaptiveDynamicStreamingTemplateResponse() (response *CreateAdaptiveDynamicStreamingTemplateResponse) {
    response = &CreateAdaptiveDynamicStreamingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAdaptiveDynamicStreamingTemplate
// This API is used to create an adaptive bitrate streaming template. Up to 100 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_DRMTYPE = "InvalidParameterValue.DrmType"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAdaptiveDynamicStreamingTemplate(request *CreateAdaptiveDynamicStreamingTemplateRequest) (response *CreateAdaptiveDynamicStreamingTemplateResponse, err error) {
    return c.CreateAdaptiveDynamicStreamingTemplateWithContext(context.Background(), request)
}

// CreateAdaptiveDynamicStreamingTemplate
// This API is used to create an adaptive bitrate streaming template. Up to 100 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_DRMTYPE = "InvalidParameterValue.DrmType"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateAdaptiveDynamicStreamingTemplateWithContext(ctx context.Context, request *CreateAdaptiveDynamicStreamingTemplateRequest) (response *CreateAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAdaptiveDynamicStreamingTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAnimatedGraphicsTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAnimatedGraphicsTemplate(request *CreateAnimatedGraphicsTemplateRequest) (response *CreateAnimatedGraphicsTemplateResponse, err error) {
    return c.CreateAnimatedGraphicsTemplateWithContext(context.Background(), request)
}

// CreateAnimatedGraphicsTemplate
// This API is used to create a custom animated image generating template. Up to 16 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAnimatedGraphicsTemplateWithContext(ctx context.Context, request *CreateAnimatedGraphicsTemplateRequest) (response *CreateAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAnimatedGraphicsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCDNDomainRequest() (request *CreateCDNDomainRequest) {
    request = &CreateCDNDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateCDNDomain")
    
    
    return
}

func NewCreateCDNDomainResponse() (response *CreateCDNDomainResponse) {
    response = &CreateCDNDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCDNDomain
// This API is used for adding domain names to VOD. A user can add up to 20 domain names. 1. After the domain name is added successfully, VOD will carry out the deployment of the domain name. It takes approximately 2 minutes for the domain name to change from the deployment status to the online status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DOMAINNAMEINBLACKLIST = "InvalidParameterValue.DomainNameInBlackList"
func (c *Client) CreateCDNDomain(request *CreateCDNDomainRequest) (response *CreateCDNDomainResponse, err error) {
    return c.CreateCDNDomainWithContext(context.Background(), request)
}

// CreateCDNDomain
// This API is used for adding domain names to VOD. A user can add up to 20 domain names. 1. After the domain name is added successfully, VOD will carry out the deployment of the domain name. It takes approximately 2 minutes for the domain name to change from the deployment status to the online status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DOMAINNAMEINBLACKLIST = "InvalidParameterValue.DomainNameInBlackList"
func (c *Client) CreateCDNDomainWithContext(ctx context.Context, request *CreateCDNDomainRequest) (response *CreateCDNDomainResponse, err error) {
    if request == nil {
        request = NewCreateCDNDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCDNDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCDNDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClassRequest() (request *CreateClassRequest) {
    request = &CreateClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateClass")
    
    
    return
}

func NewCreateClassResponse() (response *CreateClassResponse) {
    response = &CreateClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClass
// * This API is used to categorize media assets for management;
//
// * It does not affect the categories of existing media assets. If you want to modify the category of a media asset, call the [ModifyMediaInfo](https://intl.cloud.tencent.com/document/product/266/31762?from_cn_redirect=1) API.
//
// * There can be up to 4 levels of categories.
//
// * One category can have up to 500 subcategories under it.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSLEVELLIMITEXCEEDED = "FailedOperation.ClassLevelLimitExceeded"
//  FAILEDOPERATION_CLASSNAMEDUPLICATE = "FailedOperation.ClassNameDuplicate"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_PARENTIDNOFOUND = "FailedOperation.ParentIdNoFound"
//  FAILEDOPERATION_SUBCLASSLIMITEXCEEDED = "FailedOperation.SubclassLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSNAME = "InvalidParameterValue.ClassName"
//  INVALIDPARAMETERVALUE_PARENTID = "InvalidParameterValue.ParentId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClass(request *CreateClassRequest) (response *CreateClassResponse, err error) {
    return c.CreateClassWithContext(context.Background(), request)
}

// CreateClass
// * This API is used to categorize media assets for management;
//
// * It does not affect the categories of existing media assets. If you want to modify the category of a media asset, call the [ModifyMediaInfo](https://intl.cloud.tencent.com/document/product/266/31762?from_cn_redirect=1) API.
//
// * There can be up to 4 levels of categories.
//
// * One category can have up to 500 subcategories under it.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSLEVELLIMITEXCEEDED = "FailedOperation.ClassLevelLimitExceeded"
//  FAILEDOPERATION_CLASSNAMEDUPLICATE = "FailedOperation.ClassNameDuplicate"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_PARENTIDNOFOUND = "FailedOperation.ParentIdNoFound"
//  FAILEDOPERATION_SUBCLASSLIMITEXCEEDED = "FailedOperation.SubclassLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSNAME = "InvalidParameterValue.ClassName"
//  INVALIDPARAMETERVALUE_PARENTID = "InvalidParameterValue.ParentId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClassWithContext(ctx context.Context, request *CreateClassRequest) (response *CreateClassResponse, err error) {
    if request == nil {
        request = NewCreateClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClass require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClassResponse()
    err = c.Send(request, response)
    return
}

func NewCreateComplexAdaptiveDynamicStreamingTaskRequest() (request *CreateComplexAdaptiveDynamicStreamingTaskRequest) {
    request = &CreateComplexAdaptiveDynamicStreamingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateComplexAdaptiveDynamicStreamingTask")
    
    
    return
}

func NewCreateComplexAdaptiveDynamicStreamingTaskResponse() (response *CreateComplexAdaptiveDynamicStreamingTaskResponse) {
    response = &CreateComplexAdaptiveDynamicStreamingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateComplexAdaptiveDynamicStreamingTask
// Initiate an adaptive bitrate streaming processing task, with the following features:
//
// 1. Output HLS and MPEG-DASH adaptive streams based on the specified templates;
//
// 2. DRM options for the adaptive bitrate streaming include no encryption, Widevine, or FairPlay.
//
// 3. The output adaptive bitrate streaming can include multiple audio streams in different languages, with each language sourced from different media files.
//
// 4. The output adaptive stream can also include multiple subtitle streams in different languages.
//
// 
//
// **Notes:**
//
// 1. When using opening/closing credits, the video stream in the opening credit must be synchronized with the audio stream; otherwise, it will result in audio and video desynchronization in the output content.
//
// 2. If the output adaptive bitrate streaming needs to include audio from the main media, the FileId of the main media must be specified in the AudioSet parameter.
//
// 3. Subtitles must be added to the main media beforehand, which can be done through the ModifyMediaInfo API interface or the Video Management page in the console.
//
// 4. Support for TESHD transcoding and watermarking is currently not available.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateComplexAdaptiveDynamicStreamingTask(request *CreateComplexAdaptiveDynamicStreamingTaskRequest) (response *CreateComplexAdaptiveDynamicStreamingTaskResponse, err error) {
    return c.CreateComplexAdaptiveDynamicStreamingTaskWithContext(context.Background(), request)
}

// CreateComplexAdaptiveDynamicStreamingTask
// Initiate an adaptive bitrate streaming processing task, with the following features:
//
// 1. Output HLS and MPEG-DASH adaptive streams based on the specified templates;
//
// 2. DRM options for the adaptive bitrate streaming include no encryption, Widevine, or FairPlay.
//
// 3. The output adaptive bitrate streaming can include multiple audio streams in different languages, with each language sourced from different media files.
//
// 4. The output adaptive stream can also include multiple subtitle streams in different languages.
//
// 
//
// **Notes:**
//
// 1. When using opening/closing credits, the video stream in the opening credit must be synchronized with the audio stream; otherwise, it will result in audio and video desynchronization in the output content.
//
// 2. If the output adaptive bitrate streaming needs to include audio from the main media, the FileId of the main media must be specified in the AudioSet parameter.
//
// 3. Subtitles must be added to the main media beforehand, which can be done through the ModifyMediaInfo API interface or the Video Management page in the console.
//
// 4. Support for TESHD transcoding and watermarking is currently not available.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateComplexAdaptiveDynamicStreamingTaskWithContext(ctx context.Context, request *CreateComplexAdaptiveDynamicStreamingTaskRequest) (response *CreateComplexAdaptiveDynamicStreamingTaskResponse, err error) {
    if request == nil {
        request = NewCreateComplexAdaptiveDynamicStreamingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateComplexAdaptiveDynamicStreamingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateComplexAdaptiveDynamicStreamingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContentReviewTemplateRequest() (request *CreateContentReviewTemplateRequest) {
    request = &CreateContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateContentReviewTemplate")
    
    
    return
}

func NewCreateContentReviewTemplateResponse() (response *CreateContentReviewTemplateResponse) {
    response = &CreateContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateContentReviewTemplate
// We have <font color=red>stopped updating</font> this API. Our new moderation templates can moderate audio/video as well as images. For details, see [CreateReviewTemplate](https://intl.cloud.tencent.com/document/api/266/84391?from_cn_redirect=1).
//
// This API is used to create a custom audio/video moderation template. Up to 50 templates can be created in total.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_REVIEWWALLSWITCH = "InvalidParameterValue.ReviewWallSwitch"
//  INVALIDPARAMETERVALUE_SCREENSHOTINTERVAL = "InvalidParameterValue.ScreenshotInterval"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateContentReviewTemplate(request *CreateContentReviewTemplateRequest) (response *CreateContentReviewTemplateResponse, err error) {
    return c.CreateContentReviewTemplateWithContext(context.Background(), request)
}

// CreateContentReviewTemplate
// We have <font color=red>stopped updating</font> this API. Our new moderation templates can moderate audio/video as well as images. For details, see [CreateReviewTemplate](https://intl.cloud.tencent.com/document/api/266/84391?from_cn_redirect=1).
//
// This API is used to create a custom audio/video moderation template. Up to 50 templates can be created in total.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_REVIEWWALLSWITCH = "InvalidParameterValue.ReviewWallSwitch"
//  INVALIDPARAMETERVALUE_SCREENSHOTINTERVAL = "InvalidParameterValue.ScreenshotInterval"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateContentReviewTemplateWithContext(ctx context.Context, request *CreateContentReviewTemplateRequest) (response *CreateContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewCreateContentReviewTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainVerifyRecordRequest() (request *CreateDomainVerifyRecordRequest) {
    request = &CreateDomainVerifyRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateDomainVerifyRecord")
    
    
    return
}

func NewCreateDomainVerifyRecordResponse() (response *CreateDomainVerifyRecordResponse) {
    response = &CreateDomainVerifyRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDomainVerifyRecord
// This API is used for generating a subdomain resolution, prompting the customer to add it to the domain name resolution, used for wildcard domain and domain name retrieval verification of ownership.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDomainVerifyRecord(request *CreateDomainVerifyRecordRequest) (response *CreateDomainVerifyRecordResponse, err error) {
    return c.CreateDomainVerifyRecordWithContext(context.Background(), request)
}

// CreateDomainVerifyRecord
// This API is used for generating a subdomain resolution, prompting the customer to add it to the domain name resolution, used for wildcard domain and domain name retrieval verification of ownership.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDomainVerifyRecordWithContext(ctx context.Context, request *CreateDomainVerifyRecordRequest) (response *CreateDomainVerifyRecordResponse, err error) {
    if request == nil {
        request = NewCreateDomainVerifyRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainVerifyRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainVerifyRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnhanceMediaTemplateRequest() (request *CreateEnhanceMediaTemplateRequest) {
    request = &CreateEnhanceMediaTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateEnhanceMediaTemplate")
    
    
    return
}

func NewCreateEnhanceMediaTemplateResponse() (response *CreateEnhanceMediaTemplateResponse) {
    response = &CreateEnhanceMediaTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnhanceMediaTemplate
// Create enhance media template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEnhanceMediaTemplate(request *CreateEnhanceMediaTemplateRequest) (response *CreateEnhanceMediaTemplateResponse, err error) {
    return c.CreateEnhanceMediaTemplateWithContext(context.Background(), request)
}

// CreateEnhanceMediaTemplate
// Create enhance media template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEnhanceMediaTemplateWithContext(ctx context.Context, request *CreateEnhanceMediaTemplateRequest) (response *CreateEnhanceMediaTemplateResponse, err error) {
    if request == nil {
        request = NewCreateEnhanceMediaTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnhanceMediaTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnhanceMediaTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHeadTailTemplateRequest() (request *CreateHeadTailTemplateRequest) {
    request = &CreateHeadTailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateHeadTailTemplate")
    
    
    return
}

func NewCreateHeadTailTemplateResponse() (response *CreateHeadTailTemplateResponse) {
    response = &CreateHeadTailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHeadTailTemplate
// Create HeadTail Template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateHeadTailTemplate(request *CreateHeadTailTemplateRequest) (response *CreateHeadTailTemplateResponse, err error) {
    return c.CreateHeadTailTemplateWithContext(context.Background(), request)
}

// CreateHeadTailTemplate
// Create HeadTail Template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateHeadTailTemplateWithContext(ctx context.Context, request *CreateHeadTailTemplateRequest) (response *CreateHeadTailTemplateResponse, err error) {
    if request == nil {
        request = NewCreateHeadTailTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHeadTailTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHeadTailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageProcessingTemplateRequest() (request *CreateImageProcessingTemplateRequest) {
    request = &CreateImageProcessingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateImageProcessingTemplate")
    
    
    return
}

func NewCreateImageProcessingTemplateResponse() (response *CreateImageProcessingTemplateResponse) {
    response = &CreateImageProcessingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateImageProcessingTemplate
// This API is used to create a custom image processing template. A template can include at most 10 operations, for example, crop-scale-crop-blur-scale-crop-scale-crop-blur-scale. You can have up to 16 image processing templates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_CUTANDCROPS = "InvalidParameterValue.CutAndCrops"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_THUMBNAILS = "InvalidParameterValue.Thumbnails"
//  INVALIDPARAMETERVALUE_WATERMARKS = "InvalidParameterValue.Watermarks"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateImageProcessingTemplate(request *CreateImageProcessingTemplateRequest) (response *CreateImageProcessingTemplateResponse, err error) {
    return c.CreateImageProcessingTemplateWithContext(context.Background(), request)
}

// CreateImageProcessingTemplate
// This API is used to create a custom image processing template. A template can include at most 10 operations, for example, crop-scale-crop-blur-scale-crop-scale-crop-blur-scale. You can have up to 16 image processing templates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_CUTANDCROPS = "InvalidParameterValue.CutAndCrops"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_THUMBNAILS = "InvalidParameterValue.Thumbnails"
//  INVALIDPARAMETERVALUE_WATERMARKS = "InvalidParameterValue.Watermarks"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateImageProcessingTemplateWithContext(ctx context.Context, request *CreateImageProcessingTemplateRequest) (response *CreateImageProcessingTemplateResponse, err error) {
    if request == nil {
        request = NewCreateImageProcessingTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageProcessingTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageProcessingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageSpriteTemplateRequest() (request *CreateImageSpriteTemplateRequest) {
    request = &CreateImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateImageSpriteTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateImageSpriteTemplate(request *CreateImageSpriteTemplateRequest) (response *CreateImageSpriteTemplateResponse, err error) {
    return c.CreateImageSpriteTemplateWithContext(context.Background(), request)
}

// CreateImageSpriteTemplate
// This API is used to create a custom image sprite generating template. Up to 16 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateImageSpriteTemplateWithContext(ctx context.Context, request *CreateImageSpriteTemplateRequest) (response *CreateImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewCreateImageSpriteTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateJustInTimeTranscodeTemplateRequest() (request *CreateJustInTimeTranscodeTemplateRequest) {
    request = &CreateJustInTimeTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateJustInTimeTranscodeTemplate")
    
    
    return
}

func NewCreateJustInTimeTranscodeTemplateResponse() (response *CreateJustInTimeTranscodeTemplateResponse) {
    response = &CreateJustInTimeTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateJustInTimeTranscodeTemplate
// Create Just In Time Transcode Template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTEDNAME = "InvalidParameterValue.ExistedName"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateJustInTimeTranscodeTemplate(request *CreateJustInTimeTranscodeTemplateRequest) (response *CreateJustInTimeTranscodeTemplateResponse, err error) {
    return c.CreateJustInTimeTranscodeTemplateWithContext(context.Background(), request)
}

// CreateJustInTimeTranscodeTemplate
// Create Just In Time Transcode Template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTEDNAME = "InvalidParameterValue.ExistedName"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateJustInTimeTranscodeTemplateWithContext(ctx context.Context, request *CreateJustInTimeTranscodeTemplateRequest) (response *CreateJustInTimeTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateJustInTimeTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJustInTimeTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateJustInTimeTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonSampleRequest() (request *CreatePersonSampleRequest) {
    request = &CreatePersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreatePersonSample")
    
    
    return
}

func NewCreatePersonSampleResponse() (response *CreatePersonSampleResponse) {
    response = &CreatePersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePersonSample
// This API is used to create samples for using facial features positioning and other technologies to perform video processing operations such as content recognition and inappropriate information recognition.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreatePersonSample(request *CreatePersonSampleRequest) (response *CreatePersonSampleResponse, err error) {
    return c.CreatePersonSampleWithContext(context.Background(), request)
}

// CreatePersonSample
// This API is used to create samples for using facial features positioning and other technologies to perform video processing operations such as content recognition and inappropriate information recognition.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreatePersonSampleWithContext(ctx context.Context, request *CreatePersonSampleRequest) (response *CreatePersonSampleResponse, err error) {
    if request == nil {
        request = NewCreatePersonSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcedureTemplateRequest() (request *CreateProcedureTemplateRequest) {
    request = &CreateProcedureTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateProcedureTemplate")
    
    
    return
}

func NewCreateProcedureTemplateResponse() (response *CreateProcedureTemplateResponse) {
    response = &CreateProcedureTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProcedureTemplate
// This API is used to create a custom task flow template. Up to 50 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXISTEDPROCEDURENAME = "InvalidParameter.ExistedProcedureName"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProcedureTemplate(request *CreateProcedureTemplateRequest) (response *CreateProcedureTemplateResponse, err error) {
    return c.CreateProcedureTemplateWithContext(context.Background(), request)
}

// CreateProcedureTemplate
// This API is used to create a custom task flow template. Up to 50 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXISTEDPROCEDURENAME = "InvalidParameter.ExistedProcedureName"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProcedureTemplateWithContext(ctx context.Context, request *CreateProcedureTemplateRequest) (response *CreateProcedureTemplateResponse, err error) {
    if request == nil {
        request = NewCreateProcedureTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProcedureTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProcedureTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQualityInspectTemplateRequest() (request *CreateQualityInspectTemplateRequest) {
    request = &CreateQualityInspectTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateQualityInspectTemplate")
    
    
    return
}

func NewCreateQualityInspectTemplateResponse() (response *CreateQualityInspectTemplateResponse) {
    response = &CreateQualityInspectTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQualityInspectTemplate
// Creates media quality inspection template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateQualityInspectTemplate(request *CreateQualityInspectTemplateRequest) (response *CreateQualityInspectTemplateResponse, err error) {
    return c.CreateQualityInspectTemplateWithContext(context.Background(), request)
}

// CreateQualityInspectTemplate
// Creates media quality inspection template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateQualityInspectTemplateWithContext(ctx context.Context, request *CreateQualityInspectTemplateRequest) (response *CreateQualityInspectTemplateResponse, err error) {
    if request == nil {
        request = NewCreateQualityInspectTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQualityInspectTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQualityInspectTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRebuildMediaTemplateRequest() (request *CreateRebuildMediaTemplateRequest) {
    request = &CreateRebuildMediaTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateRebuildMediaTemplate")
    
    
    return
}

func NewCreateRebuildMediaTemplateResponse() (response *CreateRebuildMediaTemplateResponse) {
    response = &CreateRebuildMediaTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRebuildMediaTemplate
// Create rebuild media template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
func (c *Client) CreateRebuildMediaTemplate(request *CreateRebuildMediaTemplateRequest) (response *CreateRebuildMediaTemplateResponse, err error) {
    return c.CreateRebuildMediaTemplateWithContext(context.Background(), request)
}

// CreateRebuildMediaTemplate
// Create rebuild media template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
func (c *Client) CreateRebuildMediaTemplateWithContext(ctx context.Context, request *CreateRebuildMediaTemplateRequest) (response *CreateRebuildMediaTemplateResponse, err error) {
    if request == nil {
        request = NewCreateRebuildMediaTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRebuildMediaTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRebuildMediaTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReviewTemplateRequest() (request *CreateReviewTemplateRequest) {
    request = &CreateReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateReviewTemplate")
    
    
    return
}

func NewCreateReviewTemplateResponse() (response *CreateReviewTemplateResponse) {
    response = &CreateReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReviewTemplate
// This API is used to create a custom moderation template. Up to 50 templates can be created in total.
//
// > The templates can only be used by the APIs [ReviewAudioVideo](https://intl.cloud.tencent.com/document/api/266/80283?from_cn_redirect=1) and [ReviewImage](https://intl.cloud.tencent.com/document/api/266/73217?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER_LABELS = "InvalidParameter.Labels"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELS = "InvalidParameterValue.Labels"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateReviewTemplate(request *CreateReviewTemplateRequest) (response *CreateReviewTemplateResponse, err error) {
    return c.CreateReviewTemplateWithContext(context.Background(), request)
}

// CreateReviewTemplate
// This API is used to create a custom moderation template. Up to 50 templates can be created in total.
//
// > The templates can only be used by the APIs [ReviewAudioVideo](https://intl.cloud.tencent.com/document/api/266/80283?from_cn_redirect=1) and [ReviewImage](https://intl.cloud.tencent.com/document/api/266/73217?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETER_LABELS = "InvalidParameter.Labels"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELS = "InvalidParameterValue.Labels"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateReviewTemplateWithContext(ctx context.Context, request *CreateReviewTemplateRequest) (response *CreateReviewTemplateResponse, err error) {
    if request == nil {
        request = NewCreateReviewTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoundPlayRequest() (request *CreateRoundPlayRequest) {
    request = &CreateRoundPlayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateRoundPlay")
    
    
    return
}

func NewCreateRoundPlayResponse() (response *CreateRoundPlayResponse) {
    response = &CreateRoundPlayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoundPlay
// This API is used to create a playlist. You can create at most 100 playlists.
//
// For each video on the list, you can either use the original file or a transcoding file.
//
// The files must be in HLS format. Preferably, they should have the same bitrate and resolution.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EXPIREDTIME = "InvalidParameterValue.ExpiredTime"
//  INVALIDPARAMETERVALUE_ROUNDPLAYALREADYEXISTS = "InvalidParameterValue.RoundPlayAlreadyExists"
//  INVALIDPARAMETERVALUE_ROUNDPLAYLIST = "InvalidParameterValue.RoundPlaylist"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  LIMITEXCEEDED_PLAYLIST = "LimitExceeded.PlayList"
//  LIMITEXCEEDED_ROUNDPLAYS = "LimitExceeded.RoundPlays"
//  LIMITEXCEEDED_RUNNINGROUNDPLAYS = "LimitExceeded.RunningRoundPlays"
func (c *Client) CreateRoundPlay(request *CreateRoundPlayRequest) (response *CreateRoundPlayResponse, err error) {
    return c.CreateRoundPlayWithContext(context.Background(), request)
}

// CreateRoundPlay
// This API is used to create a playlist. You can create at most 100 playlists.
//
// For each video on the list, you can either use the original file or a transcoding file.
//
// The files must be in HLS format. Preferably, they should have the same bitrate and resolution.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EXPIREDTIME = "InvalidParameterValue.ExpiredTime"
//  INVALIDPARAMETERVALUE_ROUNDPLAYALREADYEXISTS = "InvalidParameterValue.RoundPlayAlreadyExists"
//  INVALIDPARAMETERVALUE_ROUNDPLAYLIST = "InvalidParameterValue.RoundPlaylist"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  LIMITEXCEEDED_PLAYLIST = "LimitExceeded.PlayList"
//  LIMITEXCEEDED_ROUNDPLAYS = "LimitExceeded.RoundPlays"
//  LIMITEXCEEDED_RUNNINGROUNDPLAYS = "LimitExceeded.RunningRoundPlays"
func (c *Client) CreateRoundPlayWithContext(ctx context.Context, request *CreateRoundPlayRequest) (response *CreateRoundPlayResponse, err error) {
    if request == nil {
        request = NewCreateRoundPlayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoundPlay require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoundPlayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSampleSnapshotTemplateRequest() (request *CreateSampleSnapshotTemplateRequest) {
    request = &CreateSampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateSampleSnapshotTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSampleSnapshotTemplate(request *CreateSampleSnapshotTemplateRequest) (response *CreateSampleSnapshotTemplateResponse, err error) {
    return c.CreateSampleSnapshotTemplateWithContext(context.Background(), request)
}

// CreateSampleSnapshotTemplate
// This API is used to create a custom sampled screencapturing template. Up to 16 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSampleSnapshotTemplateWithContext(ctx context.Context, request *CreateSampleSnapshotTemplateRequest) (response *CreateSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSampleSnapshotTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotByTimeOffsetTemplateRequest() (request *CreateSnapshotByTimeOffsetTemplateRequest) {
    request = &CreateSnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateSnapshotByTimeOffsetTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSnapshotByTimeOffsetTemplate(request *CreateSnapshotByTimeOffsetTemplateRequest) (response *CreateSnapshotByTimeOffsetTemplateResponse, err error) {
    return c.CreateSnapshotByTimeOffsetTemplateWithContext(context.Background(), request)
}

// CreateSnapshotByTimeOffsetTemplate
// This API is used to create a custom time point screencapturing template. Up to 16 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSnapshotByTimeOffsetTemplateWithContext(ctx context.Context, request *CreateSnapshotByTimeOffsetTemplateRequest) (response *CreateSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotByTimeOffsetTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStorageRegionRequest() (request *CreateStorageRegionRequest) {
    request = &CreateStorageRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateStorageRegion")
    
    
    return
}

func NewCreateStorageRegionResponse() (response *CreateStorageRegionResponse) {
    response = &CreateStorageRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStorageRegion
// This API is used to enable storage in a region.
//
//   1. When you activate VOD, the system will enable storage for you in certain regions. If you need to store data in another region, you can use this API to enable storage in that region.
//
//   2. You can use the `DescribeStorageRegions` API to query all supported storage regions and the regions you have storage access to currently.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_USERSTATUSINAVLID = "FailedOperation.UserStatusInavlid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_STORAGEREGION = "InvalidParameterValue.StorageRegion"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateStorageRegion(request *CreateStorageRegionRequest) (response *CreateStorageRegionResponse, err error) {
    return c.CreateStorageRegionWithContext(context.Background(), request)
}

// CreateStorageRegion
// This API is used to enable storage in a region.
//
//   1. When you activate VOD, the system will enable storage for you in certain regions. If you need to store data in another region, you can use this API to enable storage in that region.
//
//   2. You can use the `DescribeStorageRegions` API to query all supported storage regions and the regions you have storage access to currently.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_USERSTATUSINAVLID = "FailedOperation.UserStatusInavlid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_STORAGEREGION = "InvalidParameterValue.StorageRegion"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateStorageRegionWithContext(ctx context.Context, request *CreateStorageRegionRequest) (response *CreateStorageRegionResponse, err error) {
    if request == nil {
        request = NewCreateStorageRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStorageRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStorageRegionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubAppIdRequest() (request *CreateSubAppIdRequest) {
    request = &CreateSubAppIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateSubAppId")
    
    
    return
}

func NewCreateSubAppIdResponse() (response *CreateSubAppIdResponse) {
    response = &CreateSubAppIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSubAppId
// This API is used to create a VOD subapplication.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSubAppId(request *CreateSubAppIdRequest) (response *CreateSubAppIdResponse, err error) {
    return c.CreateSubAppIdWithContext(context.Background(), request)
}

// CreateSubAppId
// This API is used to create a VOD subapplication.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSubAppIdWithContext(ctx context.Context, request *CreateSubAppIdRequest) (response *CreateSubAppIdResponse, err error) {
    if request == nil {
        request = NewCreateSubAppIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubAppId require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSubAppIdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSuperPlayerConfigRequest() (request *CreateSuperPlayerConfigRequest) {
    request = &CreateSuperPlayerConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateSuperPlayerConfig")
    
    
    return
}

func NewCreateSuperPlayerConfigResponse() (response *CreateSuperPlayerConfigResponse) {
    response = &CreateSuperPlayerConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSuperPlayerConfig
// We have <font color='red'>stopped updating</font> this API. Currently, you no longer need a player configuration to use player signatures. For details, see [Player Signature](https://intl.cloud.tencent.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to create a player configuration. Up to 100 configurations can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSuperPlayerConfig(request *CreateSuperPlayerConfigRequest) (response *CreateSuperPlayerConfigResponse, err error) {
    return c.CreateSuperPlayerConfigWithContext(context.Background(), request)
}

// CreateSuperPlayerConfig
// We have <font color='red'>stopped updating</font> this API. Currently, you no longer need a player configuration to use player signatures. For details, see [Player Signature](https://intl.cloud.tencent.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to create a player configuration. Up to 100 configurations can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSuperPlayerConfigWithContext(ctx context.Context, request *CreateSuperPlayerConfigRequest) (response *CreateSuperPlayerConfigResponse, err error) {
    if request == nil {
        request = NewCreateSuperPlayerConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSuperPlayerConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSuperPlayerConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTranscodeTemplateRequest() (request *CreateTranscodeTemplateRequest) {
    request = &CreateTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateTranscodeTemplate")
    
    
    return
}

func NewCreateTranscodeTemplateResponse() (response *CreateTranscodeTemplateResponse) {
    response = &CreateTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTranscodeTemplate
// This API is used to create a custom transcoding template. Up to 100 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VCRF = "InvalidParameterValue.Vcrf"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTranscodeTemplate(request *CreateTranscodeTemplateRequest) (response *CreateTranscodeTemplateResponse, err error) {
    return c.CreateTranscodeTemplateWithContext(context.Background(), request)
}

// CreateTranscodeTemplate
// This API is used to create a custom transcoding template. Up to 100 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VCRF = "InvalidParameterValue.Vcrf"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTranscodeTemplateWithContext(ctx context.Context, request *CreateTranscodeTemplateRequest) (response *CreateTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVodDomainRequest() (request *CreateVodDomainRequest) {
    request = &CreateVodDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateVodDomain")
    
    
    return
}

func NewCreateVodDomainResponse() (response *CreateVodDomainResponse) {
    response = &CreateVodDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVodDomain
// This API is used to add an acceleration domain name to VOD. One user can add up to 20 domain names.
//
// 1. After a domain name is added, VOD will deploy it, and it takes about 2 minutes for the domain name status to change from `Deploying` to `Online`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DOMAINNAMEINBLACKLIST = "InvalidParameterValue.DomainNameInBlackList"
//  UNAUTHORIZEDOPERATION_DOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.DomainRecordNotVerified"
func (c *Client) CreateVodDomain(request *CreateVodDomainRequest) (response *CreateVodDomainResponse, err error) {
    return c.CreateVodDomainWithContext(context.Background(), request)
}

// CreateVodDomain
// This API is used to add an acceleration domain name to VOD. One user can add up to 20 domain names.
//
// 1. After a domain name is added, VOD will deploy it, and it takes about 2 minutes for the domain name status to change from `Deploying` to `Online`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DOMAINNAMEINBLACKLIST = "InvalidParameterValue.DomainNameInBlackList"
//  UNAUTHORIZEDOPERATION_DOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.DomainRecordNotVerified"
func (c *Client) CreateVodDomainWithContext(ctx context.Context, request *CreateVodDomainRequest) (response *CreateVodDomainResponse, err error) {
    if request == nil {
        request = NewCreateVodDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVodDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVodDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWatermarkTemplateRequest() (request *CreateWatermarkTemplateRequest) {
    request = &CreateWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateWatermarkTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateWatermarkTemplate(request *CreateWatermarkTemplateRequest) (response *CreateWatermarkTemplateResponse, err error) {
    return c.CreateWatermarkTemplateWithContext(context.Background(), request)
}

// CreateWatermarkTemplate
// This API is used to create a custom watermarking template. Up to 1,000 templates can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateWatermarkTemplateWithContext(ctx context.Context, request *CreateWatermarkTemplateRequest) (response *CreateWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewCreateWatermarkTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateWordSamples")
    
    
    return
}

func NewCreateWordSamplesResponse() (response *CreateWordSamplesResponse) {
    response = &CreateWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWordSamples
// This API is used to create keyword samples in batches for using OCR and ASR technologies to perform video processing operations such as content recognition and inappropriate information recognition.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateWordSamples(request *CreateWordSamplesRequest) (response *CreateWordSamplesResponse, err error) {
    return c.CreateWordSamplesWithContext(context.Background(), request)
}

// CreateWordSamples
// This API is used to create keyword samples in batches for using OCR and ASR technologies to perform video processing operations such as content recognition and inappropriate information recognition.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateWordSamplesWithContext(ctx context.Context, request *CreateWordSamplesRequest) (response *CreateWordSamplesResponse, err error) {
    if request == nil {
        request = NewCreateWordSamplesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWordSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIAnalysisTemplateRequest() (request *DeleteAIAnalysisTemplateRequest) {
    request = &DeleteAIAnalysisTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAIAnalysisTemplate")
    
    
    return
}

func NewDeleteAIAnalysisTemplateResponse() (response *DeleteAIAnalysisTemplateResponse) {
    response = &DeleteAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAIAnalysisTemplate
// This API is used to delete a custom video content analysis template.
//
// 
//
// Note: templates with an ID below 10000 are preset and cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAIAnalysisTemplate(request *DeleteAIAnalysisTemplateRequest) (response *DeleteAIAnalysisTemplateResponse, err error) {
    return c.DeleteAIAnalysisTemplateWithContext(context.Background(), request)
}

// DeleteAIAnalysisTemplate
// This API is used to delete a custom video content analysis template.
//
// 
//
// Note: templates with an ID below 10000 are preset and cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAIAnalysisTemplateWithContext(ctx context.Context, request *DeleteAIAnalysisTemplateRequest) (response *DeleteAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIAnalysisTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAIRecognitionTemplate")
    
    
    return
}

func NewDeleteAIRecognitionTemplateResponse() (response *DeleteAIRecognitionTemplateResponse) {
    response = &DeleteAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAIRecognitionTemplate
// This API is used to delete a custom video content recognition template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAIRecognitionTemplate(request *DeleteAIRecognitionTemplateRequest) (response *DeleteAIRecognitionTemplateResponse, err error) {
    return c.DeleteAIRecognitionTemplateWithContext(context.Background(), request)
}

// DeleteAIRecognitionTemplate
// This API is used to delete a custom video content recognition template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAIRecognitionTemplateWithContext(ctx context.Context, request *DeleteAIRecognitionTemplateRequest) (response *DeleteAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAIRecognitionTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAdaptiveDynamicStreamingTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteAdaptiveDynamicStreamingTemplateWithContext(ctx context.Context, request *DeleteAdaptiveDynamicStreamingTemplateRequest) (response *DeleteAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAdaptiveDynamicStreamingTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAnimatedGraphicsTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAnimatedGraphicsTemplate(request *DeleteAnimatedGraphicsTemplateRequest) (response *DeleteAnimatedGraphicsTemplateResponse, err error) {
    return c.DeleteAnimatedGraphicsTemplateWithContext(context.Background(), request)
}

// DeleteAnimatedGraphicsTemplate
// This API is used to delete a custom animated image generating template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAnimatedGraphicsTemplateWithContext(ctx context.Context, request *DeleteAnimatedGraphicsTemplateRequest) (response *DeleteAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAnimatedGraphicsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCDNDomainRequest() (request *DeleteCDNDomainRequest) {
    request = &DeleteCDNDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteCDNDomain")
    
    
    return
}

func NewDeleteCDNDomainResponse() (response *DeleteCDNDomainResponse) {
    response = &DeleteCDNDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCDNDomain
// Delete CDN Domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCDNDomain(request *DeleteCDNDomainRequest) (response *DeleteCDNDomainResponse, err error) {
    return c.DeleteCDNDomainWithContext(context.Background(), request)
}

// DeleteCDNDomain
// Delete CDN Domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCDNDomainWithContext(ctx context.Context, request *DeleteCDNDomainRequest) (response *DeleteCDNDomainResponse, err error) {
    if request == nil {
        request = NewDeleteCDNDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCDNDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCDNDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClassRequest() (request *DeleteClassRequest) {
    request = &DeleteClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteClass")
    
    
    return
}

func NewDeleteClassResponse() (response *DeleteClassResponse) {
    response = &DeleteClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClass
// * A category can be deleted only if it has no subcategories and associated media files;
//
// * Otherwise, [delete the media files](https://intl.cloud.tencent.com/document/product/266/31764?from_cn_redirect=1) and subcategories first before deleting the category.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSID = "InvalidParameterValue.ClassId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_CLASSNOTEMPTY = "UnsupportedOperation.ClassNotEmpty"
func (c *Client) DeleteClass(request *DeleteClassRequest) (response *DeleteClassResponse, err error) {
    return c.DeleteClassWithContext(context.Background(), request)
}

// DeleteClass
// * A category can be deleted only if it has no subcategories and associated media files;
//
// * Otherwise, [delete the media files](https://intl.cloud.tencent.com/document/product/266/31764?from_cn_redirect=1) and subcategories first before deleting the category.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSID = "InvalidParameterValue.ClassId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_CLASSNOTEMPTY = "UnsupportedOperation.ClassNotEmpty"
func (c *Client) DeleteClassWithContext(ctx context.Context, request *DeleteClassRequest) (response *DeleteClassResponse, err error) {
    if request == nil {
        request = NewDeleteClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClass require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClassResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContentReviewTemplateRequest() (request *DeleteContentReviewTemplateRequest) {
    request = &DeleteContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteContentReviewTemplate")
    
    
    return
}

func NewDeleteContentReviewTemplateResponse() (response *DeleteContentReviewTemplateResponse) {
    response = &DeleteContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteContentReviewTemplate
// We have <font color=red>stopped updating</font> this API. Our new moderation templates can moderate audio/video as well as images. For details, see [DeleteReviewTemplate](https://intl.cloud.tencent.com/document/api/266/84390?from_cn_redirect=1).
//
// This API is used to delete a custom audio/video moderation template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteContentReviewTemplate(request *DeleteContentReviewTemplateRequest) (response *DeleteContentReviewTemplateResponse, err error) {
    return c.DeleteContentReviewTemplateWithContext(context.Background(), request)
}

// DeleteContentReviewTemplate
// We have <font color=red>stopped updating</font> this API. Our new moderation templates can moderate audio/video as well as images. For details, see [DeleteReviewTemplate](https://intl.cloud.tencent.com/document/api/266/84390?from_cn_redirect=1).
//
// This API is used to delete a custom audio/video moderation template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteContentReviewTemplateWithContext(ctx context.Context, request *DeleteContentReviewTemplateRequest) (response *DeleteContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteContentReviewTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEnhanceMediaTemplateRequest() (request *DeleteEnhanceMediaTemplateRequest) {
    request = &DeleteEnhanceMediaTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteEnhanceMediaTemplate")
    
    
    return
}

func NewDeleteEnhanceMediaTemplateResponse() (response *DeleteEnhanceMediaTemplateResponse) {
    response = &DeleteEnhanceMediaTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEnhanceMediaTemplate
// Delete Enhance Media template
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteEnhanceMediaTemplate(request *DeleteEnhanceMediaTemplateRequest) (response *DeleteEnhanceMediaTemplateResponse, err error) {
    return c.DeleteEnhanceMediaTemplateWithContext(context.Background(), request)
}

// DeleteEnhanceMediaTemplate
// Delete Enhance Media template
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteEnhanceMediaTemplateWithContext(ctx context.Context, request *DeleteEnhanceMediaTemplateRequest) (response *DeleteEnhanceMediaTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteEnhanceMediaTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEnhanceMediaTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEnhanceMediaTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHeadTailTemplateRequest() (request *DeleteHeadTailTemplateRequest) {
    request = &DeleteHeadTailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteHeadTailTemplate")
    
    
    return
}

func NewDeleteHeadTailTemplateResponse() (response *DeleteHeadTailTemplateResponse) {
    response = &DeleteHeadTailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHeadTailTemplate
// Delete HeadTail Template
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteHeadTailTemplate(request *DeleteHeadTailTemplateRequest) (response *DeleteHeadTailTemplateResponse, err error) {
    return c.DeleteHeadTailTemplateWithContext(context.Background(), request)
}

// DeleteHeadTailTemplate
// Delete HeadTail Template
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteHeadTailTemplateWithContext(ctx context.Context, request *DeleteHeadTailTemplateRequest) (response *DeleteHeadTailTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteHeadTailTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHeadTailTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHeadTailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageProcessingTemplateRequest() (request *DeleteImageProcessingTemplateRequest) {
    request = &DeleteImageProcessingTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteImageProcessingTemplate")
    
    
    return
}

func NewDeleteImageProcessingTemplateResponse() (response *DeleteImageProcessingTemplateResponse) {
    response = &DeleteImageProcessingTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteImageProcessingTemplate
// This API is used to delete an image processing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteImageProcessingTemplate(request *DeleteImageProcessingTemplateRequest) (response *DeleteImageProcessingTemplateResponse, err error) {
    return c.DeleteImageProcessingTemplateWithContext(context.Background(), request)
}

// DeleteImageProcessingTemplate
// This API is used to delete an image processing template.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteImageProcessingTemplateWithContext(ctx context.Context, request *DeleteImageProcessingTemplateRequest) (response *DeleteImageProcessingTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteImageProcessingTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageProcessingTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageProcessingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageSpriteTemplateRequest() (request *DeleteImageSpriteTemplateRequest) {
    request = &DeleteImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteImageSpriteTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteImageSpriteTemplate(request *DeleteImageSpriteTemplateRequest) (response *DeleteImageSpriteTemplateResponse, err error) {
    return c.DeleteImageSpriteTemplateWithContext(context.Background(), request)
}

// DeleteImageSpriteTemplate
// This API is used to delete an image sprite generating template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteImageSpriteTemplateWithContext(ctx context.Context, request *DeleteImageSpriteTemplateRequest) (response *DeleteImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteImageSpriteTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJustInTimeTranscodeTemplateRequest() (request *DeleteJustInTimeTranscodeTemplateRequest) {
    request = &DeleteJustInTimeTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteJustInTimeTranscodeTemplate")
    
    
    return
}

func NewDeleteJustInTimeTranscodeTemplateResponse() (response *DeleteJustInTimeTranscodeTemplateResponse) {
    response = &DeleteJustInTimeTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteJustInTimeTranscodeTemplate
// Delete Just In Time Transcode Template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteJustInTimeTranscodeTemplate(request *DeleteJustInTimeTranscodeTemplateRequest) (response *DeleteJustInTimeTranscodeTemplateResponse, err error) {
    return c.DeleteJustInTimeTranscodeTemplateWithContext(context.Background(), request)
}

// DeleteJustInTimeTranscodeTemplate
// Delete Just In Time Transcode Template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteJustInTimeTranscodeTemplateWithContext(ctx context.Context, request *DeleteJustInTimeTranscodeTemplateRequest) (response *DeleteJustInTimeTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteJustInTimeTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJustInTimeTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJustInTimeTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaRequest() (request *DeleteMediaRequest) {
    request = &DeleteMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteMedia")
    
    
    return
}

func NewDeleteMediaResponse() (response *DeleteMediaResponse) {
    response = &DeleteMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMedia
// * This API is used to delete a media file and its processed files, such as the transcoded video files, image sprites, screenshots, and videos for publishing on WeChat.
//
// * You can delete the original files, transcoded video files, and videos for publishing on WeChat, etc. of videos with specified IDs.
//
// * Note: after the original file of a video is deleted, you cannot transcode the video, publish it on WeChat, or perform other operations on it.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMedia(request *DeleteMediaRequest) (response *DeleteMediaResponse, err error) {
    return c.DeleteMediaWithContext(context.Background(), request)
}

// DeleteMedia
// * This API is used to delete a media file and its processed files, such as the transcoded video files, image sprites, screenshots, and videos for publishing on WeChat.
//
// * You can delete the original files, transcoded video files, and videos for publishing on WeChat, etc. of videos with specified IDs.
//
// * Note: after the original file of a video is deleted, you cannot transcode the video, publish it on WeChat, or perform other operations on it.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMediaWithContext(ctx context.Context, request *DeleteMediaRequest) (response *DeleteMediaResponse, err error) {
    if request == nil {
        request = NewDeleteMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMediaResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonSampleRequest() (request *DeletePersonSampleRequest) {
    request = &DeletePersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeletePersonSample")
    
    
    return
}

func NewDeletePersonSampleResponse() (response *DeletePersonSampleResponse) {
    response = &DeletePersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePersonSample
// This API is used to delete samples according to sample IDs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePersonSample(request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    return c.DeletePersonSampleWithContext(context.Background(), request)
}

// DeletePersonSample
// This API is used to delete samples according to sample IDs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePersonSampleWithContext(ctx context.Context, request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    if request == nil {
        request = NewDeletePersonSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProcedureTemplateRequest() (request *DeleteProcedureTemplateRequest) {
    request = &DeleteProcedureTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteProcedureTemplate")
    
    
    return
}

func NewDeleteProcedureTemplateResponse() (response *DeleteProcedureTemplateResponse) {
    response = &DeleteProcedureTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProcedureTemplate
// Delete user-created task flow templates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteProcedureTemplate(request *DeleteProcedureTemplateRequest) (response *DeleteProcedureTemplateResponse, err error) {
    return c.DeleteProcedureTemplateWithContext(context.Background(), request)
}

// DeleteProcedureTemplate
// Delete user-created task flow templates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteProcedureTemplateWithContext(ctx context.Context, request *DeleteProcedureTemplateRequest) (response *DeleteProcedureTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteProcedureTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProcedureTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProcedureTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQualityInspectTemplateRequest() (request *DeleteQualityInspectTemplateRequest) {
    request = &DeleteQualityInspectTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteQualityInspectTemplate")
    
    
    return
}

func NewDeleteQualityInspectTemplateResponse() (response *DeleteQualityInspectTemplateResponse) {
    response = &DeleteQualityInspectTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQualityInspectTemplate
// Deletes media quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteQualityInspectTemplate(request *DeleteQualityInspectTemplateRequest) (response *DeleteQualityInspectTemplateResponse, err error) {
    return c.DeleteQualityInspectTemplateWithContext(context.Background(), request)
}

// DeleteQualityInspectTemplate
// Deletes media quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteQualityInspectTemplateWithContext(ctx context.Context, request *DeleteQualityInspectTemplateRequest) (response *DeleteQualityInspectTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteQualityInspectTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQualityInspectTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQualityInspectTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRebuildMediaTemplateRequest() (request *DeleteRebuildMediaTemplateRequest) {
    request = &DeleteRebuildMediaTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteRebuildMediaTemplate")
    
    
    return
}

func NewDeleteRebuildMediaTemplateResponse() (response *DeleteRebuildMediaTemplateResponse) {
    response = &DeleteRebuildMediaTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRebuildMediaTemplate
// Delete rebuild media template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteRebuildMediaTemplate(request *DeleteRebuildMediaTemplateRequest) (response *DeleteRebuildMediaTemplateResponse, err error) {
    return c.DeleteRebuildMediaTemplateWithContext(context.Background(), request)
}

// DeleteRebuildMediaTemplate
// Delete rebuild media template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteRebuildMediaTemplateWithContext(ctx context.Context, request *DeleteRebuildMediaTemplateRequest) (response *DeleteRebuildMediaTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteRebuildMediaTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRebuildMediaTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRebuildMediaTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReviewTemplateRequest() (request *DeleteReviewTemplateRequest) {
    request = &DeleteReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteReviewTemplate")
    
    
    return
}

func NewDeleteReviewTemplateResponse() (response *DeleteReviewTemplateResponse) {
    response = &DeleteReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReviewTemplate
// This API is used to delete a custom moderation template.
//
// > The templates can only be used by the APIs [ReviewAudioVideo](https://intl.cloud.tencent.com/document/api/266/80283?from_cn_redirect=1) and [ReviewImage](https://intl.cloud.tencent.com/document/api/266/73217?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteReviewTemplate(request *DeleteReviewTemplateRequest) (response *DeleteReviewTemplateResponse, err error) {
    return c.DeleteReviewTemplateWithContext(context.Background(), request)
}

// DeleteReviewTemplate
// This API is used to delete a custom moderation template.
//
// > The templates can only be used by the APIs [ReviewAudioVideo](https://intl.cloud.tencent.com/document/api/266/80283?from_cn_redirect=1) and [ReviewImage](https://intl.cloud.tencent.com/document/api/266/73217?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteReviewTemplateWithContext(ctx context.Context, request *DeleteReviewTemplateRequest) (response *DeleteReviewTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteReviewTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoundPlayRequest() (request *DeleteRoundPlayRequest) {
    request = &DeleteRoundPlayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteRoundPlay")
    
    
    return
}

func NewDeleteRoundPlayResponse() (response *DeleteRoundPlayResponse) {
    response = &DeleteRoundPlayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoundPlay
// This API is used to delete a playlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRoundPlay(request *DeleteRoundPlayRequest) (response *DeleteRoundPlayResponse, err error) {
    return c.DeleteRoundPlayWithContext(context.Background(), request)
}

// DeleteRoundPlay
// This API is used to delete a playlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRoundPlayWithContext(ctx context.Context, request *DeleteRoundPlayRequest) (response *DeleteRoundPlayResponse, err error) {
    if request == nil {
        request = NewDeleteRoundPlayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoundPlay require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoundPlayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSampleSnapshotTemplateRequest() (request *DeleteSampleSnapshotTemplateRequest) {
    request = &DeleteSampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteSampleSnapshotTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSampleSnapshotTemplate(request *DeleteSampleSnapshotTemplateRequest) (response *DeleteSampleSnapshotTemplateResponse, err error) {
    return c.DeleteSampleSnapshotTemplateWithContext(context.Background(), request)
}

// DeleteSampleSnapshotTemplate
// This API is used to delete a custom sampled screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSampleSnapshotTemplateWithContext(ctx context.Context, request *DeleteSampleSnapshotTemplateRequest) (response *DeleteSampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSampleSnapshotTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotByTimeOffsetTemplateRequest() (request *DeleteSnapshotByTimeOffsetTemplateRequest) {
    request = &DeleteSnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteSnapshotByTimeOffsetTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSnapshotByTimeOffsetTemplate(request *DeleteSnapshotByTimeOffsetTemplateRequest) (response *DeleteSnapshotByTimeOffsetTemplateResponse, err error) {
    return c.DeleteSnapshotByTimeOffsetTemplateWithContext(context.Background(), request)
}

// DeleteSnapshotByTimeOffsetTemplate
// This API is used to delete a custom time point screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSnapshotByTimeOffsetTemplateWithContext(ctx context.Context, request *DeleteSnapshotByTimeOffsetTemplateRequest) (response *DeleteSnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotByTimeOffsetTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSuperPlayerConfigRequest() (request *DeleteSuperPlayerConfigRequest) {
    request = &DeleteSuperPlayerConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteSuperPlayerConfig")
    
    
    return
}

func NewDeleteSuperPlayerConfigResponse() (response *DeleteSuperPlayerConfigResponse) {
    response = &DeleteSuperPlayerConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSuperPlayerConfig
// We have <font color='red'>stopped updating</font> this API. Currently, you no longer need a player configuration to use player signatures. For details, see [Player Signature](https://intl.cloud.tencent.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to delete a player configuration.  
//
// *Note: Preset player configurations cannot be deleted.*
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSuperPlayerConfig(request *DeleteSuperPlayerConfigRequest) (response *DeleteSuperPlayerConfigResponse, err error) {
    return c.DeleteSuperPlayerConfigWithContext(context.Background(), request)
}

// DeleteSuperPlayerConfig
// We have <font color='red'>stopped updating</font> this API. Currently, you no longer need a player configuration to use player signatures. For details, see [Player Signature](https://intl.cloud.tencent.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to delete a player configuration.  
//
// *Note: Preset player configurations cannot be deleted.*
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSuperPlayerConfigWithContext(ctx context.Context, request *DeleteSuperPlayerConfigRequest) (response *DeleteSuperPlayerConfigResponse, err error) {
    if request == nil {
        request = NewDeleteSuperPlayerConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSuperPlayerConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSuperPlayerConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTranscodeTemplateRequest() (request *DeleteTranscodeTemplateRequest) {
    request = &DeleteTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteTranscodeTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteTranscodeTemplate(request *DeleteTranscodeTemplateRequest) (response *DeleteTranscodeTemplateResponse, err error) {
    return c.DeleteTranscodeTemplateWithContext(context.Background(), request)
}

// DeleteTranscodeTemplate
// This API is used to delete a custom transcoding template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteTranscodeTemplateWithContext(ctx context.Context, request *DeleteTranscodeTemplateRequest) (response *DeleteTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVodDomainRequest() (request *DeleteVodDomainRequest) {
    request = &DeleteVodDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteVodDomain")
    
    
    return
}

func NewDeleteVodDomainResponse() (response *DeleteVodDomainResponse) {
    response = &DeleteVodDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVodDomain
// This API is used to delete an acceleration domain name from VOD.
//
// 1. Before deleting a domain name, disable its acceleration in all regions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteVodDomain(request *DeleteVodDomainRequest) (response *DeleteVodDomainResponse, err error) {
    return c.DeleteVodDomainWithContext(context.Background(), request)
}

// DeleteVodDomain
// This API is used to delete an acceleration domain name from VOD.
//
// 1. Before deleting a domain name, disable its acceleration in all regions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteVodDomainWithContext(ctx context.Context, request *DeleteVodDomainRequest) (response *DeleteVodDomainResponse, err error) {
    if request == nil {
        request = NewDeleteVodDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVodDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVodDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWatermarkTemplateRequest() (request *DeleteWatermarkTemplateRequest) {
    request = &DeleteWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteWatermarkTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteWatermarkTemplate(request *DeleteWatermarkTemplateRequest) (response *DeleteWatermarkTemplateResponse, err error) {
    return c.DeleteWatermarkTemplateWithContext(context.Background(), request)
}

// DeleteWatermarkTemplate
// This API is used to delete a custom watermarking template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteWatermarkTemplateWithContext(ctx context.Context, request *DeleteWatermarkTemplateRequest) (response *DeleteWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteWatermarkTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteWordSamples")
    
    
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteWordSamples(request *DeleteWordSamplesRequest) (response *DeleteWordSamplesResponse, err error) {
    return c.DeleteWordSamplesWithContext(context.Background(), request)
}

// DeleteWordSamples
// This API is used to delete keyword samples in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteWordSamplesWithContext(ctx context.Context, request *DeleteWordSamplesRequest) (response *DeleteWordSamplesResponse, err error) {
    if request == nil {
        request = NewDeleteWordSamplesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWordSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIAnalysisTemplatesRequest() (request *DescribeAIAnalysisTemplatesRequest) {
    request = &DescribeAIAnalysisTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAIAnalysisTemplates")
    
    
    return
}

func NewDescribeAIAnalysisTemplatesResponse() (response *DescribeAIAnalysisTemplatesResponse) {
    response = &DescribeAIAnalysisTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIAnalysisTemplates
// This API is used to get the list of video content analysis templates based on unique template ID. The returned result includes all eligible custom and [preset video content analysis templates](https://intl.cloud.tencent.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.88.86.E6.9E.90.E6.A8.A1.E6.9D.BF).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIAnalysisTemplates(request *DescribeAIAnalysisTemplatesRequest) (response *DescribeAIAnalysisTemplatesResponse, err error) {
    return c.DescribeAIAnalysisTemplatesWithContext(context.Background(), request)
}

// DescribeAIAnalysisTemplates
// This API is used to get the list of video content analysis templates based on unique template ID. The returned result includes all eligible custom and [preset video content analysis templates](https://intl.cloud.tencent.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.88.86.E6.9E.90.E6.A8.A1.E6.9D.BF).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIAnalysisTemplatesWithContext(ctx context.Context, request *DescribeAIAnalysisTemplatesRequest) (response *DescribeAIAnalysisTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIAnalysisTemplatesRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAIRecognitionTemplates")
    
    
    return
}

func NewDescribeAIRecognitionTemplatesResponse() (response *DescribeAIRecognitionTemplatesResponse) {
    response = &DescribeAIRecognitionTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIRecognitionTemplates
// This API is used to get the list of video content recognition templates based on unique template ID. The return result includes all eligible custom and [preset video content recognition templates](https://intl.cloud.tencent.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E8.AF.86.E5.88.AB.E6.A8.A1.E6.9D.BF).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIRecognitionTemplates(request *DescribeAIRecognitionTemplatesRequest) (response *DescribeAIRecognitionTemplatesResponse, err error) {
    return c.DescribeAIRecognitionTemplatesWithContext(context.Background(), request)
}

// DescribeAIRecognitionTemplates
// This API is used to get the list of video content recognition templates based on unique template ID. The return result includes all eligible custom and [preset video content recognition templates](https://intl.cloud.tencent.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E8.AF.86.E5.88.AB.E6.A8.A1.E6.9D.BF).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIRecognitionTemplatesWithContext(ctx context.Context, request *DescribeAIRecognitionTemplatesRequest) (response *DescribeAIRecognitionTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAIRecognitionTemplatesRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAdaptiveDynamicStreamingTemplates")
    
    
    return
}

func NewDescribeAdaptiveDynamicStreamingTemplatesResponse() (response *DescribeAdaptiveDynamicStreamingTemplatesResponse) {
    response = &DescribeAdaptiveDynamicStreamingTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAdaptiveDynamicStreamingTemplates
// This API is used to query the list of transcoding to adaptive bitrate streaming templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAdaptiveDynamicStreamingTemplates(request *DescribeAdaptiveDynamicStreamingTemplatesRequest) (response *DescribeAdaptiveDynamicStreamingTemplatesResponse, err error) {
    return c.DescribeAdaptiveDynamicStreamingTemplatesWithContext(context.Background(), request)
}

// DescribeAdaptiveDynamicStreamingTemplates
// This API is used to query the list of transcoding to adaptive bitrate streaming templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAdaptiveDynamicStreamingTemplatesWithContext(ctx context.Context, request *DescribeAdaptiveDynamicStreamingTemplatesRequest) (response *DescribeAdaptiveDynamicStreamingTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAdaptiveDynamicStreamingTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAdaptiveDynamicStreamingTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAdaptiveDynamicStreamingTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllClassRequest() (request *DescribeAllClassRequest) {
    request = &DescribeAllClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAllClass")
    
    
    return
}

func NewDescribeAllClassResponse() (response *DescribeAllClassResponse) {
    response = &DescribeAllClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllClass
// * This API is used to get the information of all categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAllClass(request *DescribeAllClassRequest) (response *DescribeAllClassResponse, err error) {
    return c.DescribeAllClassWithContext(context.Background(), request)
}

// DescribeAllClass
// * This API is used to get the information of all categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAllClassWithContext(ctx context.Context, request *DescribeAllClassRequest) (response *DescribeAllClassResponse, err error) {
    if request == nil {
        request = NewDescribeAllClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllClass require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllClassResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAnimatedGraphicsTemplatesRequest() (request *DescribeAnimatedGraphicsTemplatesRequest) {
    request = &DescribeAnimatedGraphicsTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAnimatedGraphicsTemplates")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAnimatedGraphicsTemplates(request *DescribeAnimatedGraphicsTemplatesRequest) (response *DescribeAnimatedGraphicsTemplatesResponse, err error) {
    return c.DescribeAnimatedGraphicsTemplatesWithContext(context.Background(), request)
}

// DescribeAnimatedGraphicsTemplates
// This API is used to query the list of animated image generating templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAnimatedGraphicsTemplatesWithContext(ctx context.Context, request *DescribeAnimatedGraphicsTemplatesRequest) (response *DescribeAnimatedGraphicsTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAnimatedGraphicsTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAnimatedGraphicsTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAnimatedGraphicsTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCDNDomainsRequest() (request *DescribeCDNDomainsRequest) {
    request = &DescribeCDNDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCDNDomains")
    
    
    return
}

func NewDescribeCDNDomainsResponse() (response *DescribeCDNDomainsResponse) {
    response = &DescribeCDNDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCDNDomains
// Describe CDN Domains
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCDNDomains(request *DescribeCDNDomainsRequest) (response *DescribeCDNDomainsResponse, err error) {
    return c.DescribeCDNDomainsWithContext(context.Background(), request)
}

// DescribeCDNDomains
// Describe CDN Domains
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCDNDomainsWithContext(ctx context.Context, request *DescribeCDNDomainsRequest) (response *DescribeCDNDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeCDNDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCDNDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCDNDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCDNStatDetailsRequest() (request *DescribeCDNStatDetailsRequest) {
    request = &DescribeCDNStatDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCDNStatDetails")
    
    
    return
}

func NewDescribeCDNStatDetailsResponse() (response *DescribeCDNStatDetailsResponse) {
    response = &DescribeCDNStatDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCDNStatDetails
// This API is used to query CDN bandwidth, traffic, and other data of VOD domain names.
//
// * The query period is up to 90 days.
//
// * You can query data of different service regions.
//
// * You can query data of Chinese mainland by region and ISP.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_AREA = "InvalidParameterValue.Area"
//  INVALIDPARAMETERVALUE_DISTRICTS = "InvalidParameterValue.Districts"
//  INVALIDPARAMETERVALUE_DOMAINNAMES = "InvalidParameterValue.DomainNames"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_ISPS = "InvalidParameterValue.Isps"
//  INVALIDPARAMETERVALUE_METRIC = "InvalidParameterValue.Metric"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
func (c *Client) DescribeCDNStatDetails(request *DescribeCDNStatDetailsRequest) (response *DescribeCDNStatDetailsResponse, err error) {
    return c.DescribeCDNStatDetailsWithContext(context.Background(), request)
}

// DescribeCDNStatDetails
// This API is used to query CDN bandwidth, traffic, and other data of VOD domain names.
//
// * The query period is up to 90 days.
//
// * You can query data of different service regions.
//
// * You can query data of Chinese mainland by region and ISP.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_AREA = "InvalidParameterValue.Area"
//  INVALIDPARAMETERVALUE_DISTRICTS = "InvalidParameterValue.Districts"
//  INVALIDPARAMETERVALUE_DOMAINNAMES = "InvalidParameterValue.DomainNames"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_ISPS = "InvalidParameterValue.Isps"
//  INVALIDPARAMETERVALUE_METRIC = "InvalidParameterValue.Metric"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
func (c *Client) DescribeCDNStatDetailsWithContext(ctx context.Context, request *DescribeCDNStatDetailsRequest) (response *DescribeCDNStatDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeCDNStatDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCDNStatDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCDNStatDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCDNUsageDataRequest() (request *DescribeCDNUsageDataRequest) {
    request = &DescribeCDNUsageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCDNUsageData")
    
    
    return
}

func NewDescribeCDNUsageDataResponse() (response *DescribeCDNUsageDataResponse) {
    response = &DescribeCDNUsageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCDNUsageData
// This API is used to query the CDN statistics of VOD such as traffic and bandwidth.
//
//    1. Only CDN usage data for the last 365 days can be queried.
//
//    2. The query time range cannot be more than 90 days.
//
//    3. The time granularity of usage data can be specified, including 5-minute, 1-hour, and 1-day.
//
//    4. Traffic refers to the total traffic within the query time granularity, while bandwidth the peak bandwidth.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetWorkError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATAINTERVAL = "InvalidParameterValue.DataInterval"
//  INVALIDPARAMETERVALUE_DATATYPE = "InvalidParameterValue.DataType"
//  INVALIDPARAMETERVALUE_DOMAINNAME = "InvalidParameterValue.DomainName"
//  INVALIDPARAMETERVALUE_DOMAINNAMES = "InvalidParameterValue.DomainNames"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCDNUsageData(request *DescribeCDNUsageDataRequest) (response *DescribeCDNUsageDataResponse, err error) {
    return c.DescribeCDNUsageDataWithContext(context.Background(), request)
}

// DescribeCDNUsageData
// This API is used to query the CDN statistics of VOD such as traffic and bandwidth.
//
//    1. Only CDN usage data for the last 365 days can be queried.
//
//    2. The query time range cannot be more than 90 days.
//
//    3. The time granularity of usage data can be specified, including 5-minute, 1-hour, and 1-day.
//
//    4. Traffic refers to the total traffic within the query time granularity, while bandwidth the peak bandwidth.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetWorkError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATAINTERVAL = "InvalidParameterValue.DataInterval"
//  INVALIDPARAMETERVALUE_DATATYPE = "InvalidParameterValue.DataType"
//  INVALIDPARAMETERVALUE_DOMAINNAME = "InvalidParameterValue.DomainName"
//  INVALIDPARAMETERVALUE_DOMAINNAMES = "InvalidParameterValue.DomainNames"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCDNUsageDataWithContext(ctx context.Context, request *DescribeCDNUsageDataRequest) (response *DescribeCDNUsageDataResponse, err error) {
    if request == nil {
        request = NewDescribeCDNUsageDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCDNUsageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCDNUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnLogsRequest() (request *DescribeCdnLogsRequest) {
    request = &DescribeCdnLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCdnLogs")
    
    
    return
}

func NewDescribeCdnLogsResponse() (response *DescribeCdnLogsResponse) {
    response = &DescribeCdnLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCdnLogs
// This API is used to query the download links of CDN access logs of a VOD domain name.
//
//     1. Only download links of CDN logs for the last 30 days can be queried.
//
//     2. By default, CDN generates a log file every hour. If there is no CDN access for a certain hour, no log file will be generated for the hour.    
//
//     3. A CDN log download link is valid for 24 hours.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetWorkError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DOMAINNAME = "InvalidParameterValue.DomainName"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCdnLogs(request *DescribeCdnLogsRequest) (response *DescribeCdnLogsResponse, err error) {
    return c.DescribeCdnLogsWithContext(context.Background(), request)
}

// DescribeCdnLogs
// This API is used to query the download links of CDN access logs of a VOD domain name.
//
//     1. Only download links of CDN logs for the last 30 days can be queried.
//
//     2. By default, CDN generates a log file every hour. If there is no CDN access for a certain hour, no log file will be generated for the hour.    
//
//     3. A CDN log download link is valid for 24 hours.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetWorkError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DOMAINNAME = "InvalidParameterValue.DomainName"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCdnLogsWithContext(ctx context.Context, request *DescribeCdnLogsRequest) (response *DescribeCdnLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCdnLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdnLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdnLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientUploadAccelerationUsageDataRequest() (request *DescribeClientUploadAccelerationUsageDataRequest) {
    request = &DescribeClientUploadAccelerationUsageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeClientUploadAccelerationUsageData")
    
    
    return
}

func NewDescribeClientUploadAccelerationUsageDataResponse() (response *DescribeClientUploadAccelerationUsageDataResponse) {
    response = &DescribeClientUploadAccelerationUsageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClientUploadAccelerationUsageData
// This interface returns client upload acceleration statistics within the query time range. 
//
//  1. You can query the client upload acceleration statistics in the last 365 days. 
//
//  2. The query time span does not exceed 90 days. 
//
//  3. If the query time span exceeds 1 day, data with day granularity will be returned. Otherwise, data with 5-minute granularity will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeClientUploadAccelerationUsageData(request *DescribeClientUploadAccelerationUsageDataRequest) (response *DescribeClientUploadAccelerationUsageDataResponse, err error) {
    return c.DescribeClientUploadAccelerationUsageDataWithContext(context.Background(), request)
}

// DescribeClientUploadAccelerationUsageData
// This interface returns client upload acceleration statistics within the query time range. 
//
//  1. You can query the client upload acceleration statistics in the last 365 days. 
//
//  2. The query time span does not exceed 90 days. 
//
//  3. If the query time span exceeds 1 day, data with day granularity will be returned. Otherwise, data with 5-minute granularity will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeClientUploadAccelerationUsageDataWithContext(ctx context.Context, request *DescribeClientUploadAccelerationUsageDataRequest) (response *DescribeClientUploadAccelerationUsageDataResponse, err error) {
    if request == nil {
        request = NewDescribeClientUploadAccelerationUsageDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClientUploadAccelerationUsageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientUploadAccelerationUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentReviewTemplatesRequest() (request *DescribeContentReviewTemplatesRequest) {
    request = &DescribeContentReviewTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeContentReviewTemplates")
    
    
    return
}

func NewDescribeContentReviewTemplatesResponse() (response *DescribeContentReviewTemplatesResponse) {
    response = &DescribeContentReviewTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContentReviewTemplates
// We have <font color=red>stopped updating</font> this API. Our new moderation templates can moderate audio/video as well as images. For details, see [DescribeReviewTemplates](https://intl.cloud.tencent.com/document/api/266/84389?from_cn_redirect=1).
//
// This API is used to get the information of custom and [preset](https://intl.cloud.tencent.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.AE.A1.E6.A0.B8.E6.A8.A1.E6.9D.BF) audio/video moderation templates based on template IDs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeContentReviewTemplates(request *DescribeContentReviewTemplatesRequest) (response *DescribeContentReviewTemplatesResponse, err error) {
    return c.DescribeContentReviewTemplatesWithContext(context.Background(), request)
}

// DescribeContentReviewTemplates
// We have <font color=red>stopped updating</font> this API. Our new moderation templates can moderate audio/video as well as images. For details, see [DescribeReviewTemplates](https://intl.cloud.tencent.com/document/api/266/84389?from_cn_redirect=1).
//
// This API is used to get the information of custom and [preset](https://intl.cloud.tencent.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.AE.A1.E6.A0.B8.E6.A8.A1.E6.9D.BF) audio/video moderation templates based on template IDs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeContentReviewTemplatesWithContext(ctx context.Context, request *DescribeContentReviewTemplatesRequest) (response *DescribeContentReviewTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeContentReviewTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContentReviewTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContentReviewTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurrentPlaylistRequest() (request *DescribeCurrentPlaylistRequest) {
    request = &DescribeCurrentPlaylistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCurrentPlaylist")
    
    
    return
}

func NewDescribeCurrentPlaylistResponse() (response *DescribeCurrentPlaylistResponse) {
    response = &DescribeCurrentPlaylistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCurrentPlaylist
// Query current playlist of the round play.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ROUNDPLAYID = "InvalidParameterValue.RoundPlayId"
func (c *Client) DescribeCurrentPlaylist(request *DescribeCurrentPlaylistRequest) (response *DescribeCurrentPlaylistResponse, err error) {
    return c.DescribeCurrentPlaylistWithContext(context.Background(), request)
}

// DescribeCurrentPlaylist
// Query current playlist of the round play.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ROUNDPLAYID = "InvalidParameterValue.RoundPlayId"
func (c *Client) DescribeCurrentPlaylistWithContext(ctx context.Context, request *DescribeCurrentPlaylistRequest) (response *DescribeCurrentPlaylistResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentPlaylistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCurrentPlaylist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCurrentPlaylistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDailyPlayStatFileListRequest() (request *DescribeDailyPlayStatFileListRequest) {
    request = &DescribeDailyPlayStatFileListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeDailyPlayStatFileList")
    
    
    return
}

func NewDescribeDailyPlayStatFileListResponse() (response *DescribeDailyPlayStatFileListResponse) {
    response = &DescribeDailyPlayStatFileListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDailyPlayStatFileList
// This API is used to query the download links of playback statistics files.
//
// * You can query the download links of playback statistics files in the past year. The start and end dates for query cannot be more than 90 days apart.
//
// * Every day, VOD will analyze CDN request logs of the previous day and then generate a playback statistics file.
//
// * A playback statistics file includes playback times and traffic of media files.
//
// * Notes on playback times:
//
//     1. HLS file: VOD counts playback times when M3U8 files are accessed, but not when TS files are accessed.
//
//     2. Other files (MP4 files for example): VOD does not count playback times when the playback request carries the `range` parameter and the `start` parameter in `range` is not `0`. In other cases, VOD counts playback times.
//
// * Statistics on playback devices: VOD counts playback times on mobile clients when the playback request carries the `UserAgent` parameter which includes an identifier such as `Android` or `iPhone`. In other cases, VOD counts playback times on PC clients.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
func (c *Client) DescribeDailyPlayStatFileList(request *DescribeDailyPlayStatFileListRequest) (response *DescribeDailyPlayStatFileListResponse, err error) {
    return c.DescribeDailyPlayStatFileListWithContext(context.Background(), request)
}

// DescribeDailyPlayStatFileList
// This API is used to query the download links of playback statistics files.
//
// * You can query the download links of playback statistics files in the past year. The start and end dates for query cannot be more than 90 days apart.
//
// * Every day, VOD will analyze CDN request logs of the previous day and then generate a playback statistics file.
//
// * A playback statistics file includes playback times and traffic of media files.
//
// * Notes on playback times:
//
//     1. HLS file: VOD counts playback times when M3U8 files are accessed, but not when TS files are accessed.
//
//     2. Other files (MP4 files for example): VOD does not count playback times when the playback request carries the `range` parameter and the `start` parameter in `range` is not `0`. In other cases, VOD counts playback times.
//
// * Statistics on playback devices: VOD counts playback times on mobile clients when the playback request carries the `UserAgent` parameter which includes an identifier such as `Android` or `iPhone`. In other cases, VOD counts playback times on PC clients.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
func (c *Client) DescribeDailyPlayStatFileListWithContext(ctx context.Context, request *DescribeDailyPlayStatFileListRequest) (response *DescribeDailyPlayStatFileListResponse, err error) {
    if request == nil {
        request = NewDescribeDailyPlayStatFileListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDailyPlayStatFileList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDailyPlayStatFileListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDrmKeyProviderInfoRequest() (request *DescribeDrmKeyProviderInfoRequest) {
    request = &DescribeDrmKeyProviderInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeDrmKeyProviderInfo")
    
    
    return
}

func NewDescribeDrmKeyProviderInfoResponse() (response *DescribeDrmKeyProviderInfoResponse) {
    response = &DescribeDrmKeyProviderInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDrmKeyProviderInfo
// This API is used to query DRM key information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDrmKeyProviderInfo(request *DescribeDrmKeyProviderInfoRequest) (response *DescribeDrmKeyProviderInfoResponse, err error) {
    return c.DescribeDrmKeyProviderInfoWithContext(context.Background(), request)
}

// DescribeDrmKeyProviderInfo
// This API is used to query DRM key information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDrmKeyProviderInfoWithContext(ctx context.Context, request *DescribeDrmKeyProviderInfoRequest) (response *DescribeDrmKeyProviderInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDrmKeyProviderInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDrmKeyProviderInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDrmKeyProviderInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnhanceMediaTemplatesRequest() (request *DescribeEnhanceMediaTemplatesRequest) {
    request = &DescribeEnhanceMediaTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeEnhanceMediaTemplates")
    
    
    return
}

func NewDescribeEnhanceMediaTemplatesResponse() (response *DescribeEnhanceMediaTemplatesResponse) {
    response = &DescribeEnhanceMediaTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnhanceMediaTemplates
// Describe Enhance Media Templates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEnhanceMediaTemplates(request *DescribeEnhanceMediaTemplatesRequest) (response *DescribeEnhanceMediaTemplatesResponse, err error) {
    return c.DescribeEnhanceMediaTemplatesWithContext(context.Background(), request)
}

// DescribeEnhanceMediaTemplates
// Describe Enhance Media Templates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEnhanceMediaTemplatesWithContext(ctx context.Context, request *DescribeEnhanceMediaTemplatesRequest) (response *DescribeEnhanceMediaTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeEnhanceMediaTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnhanceMediaTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnhanceMediaTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventConfigRequest() (request *DescribeEventConfigRequest) {
    request = &DescribeEventConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeEventConfig")
    
    
    return
}

func NewDescribeEventConfigResponse() (response *DescribeEventConfigResponse) {
    response = &DescribeEventConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEventConfig
// Tencent Cloud VOD provides customers with services such as media upload, media management, and media processing. During or after the execution of these services, Tencent Cloud VOD also offers various corresponding event notifications to facilitate developers' awareness of the service processing status and to perform subsequent business operations. Developers can use this interface to query the current configuration of event notification reception methods, reception addresses, and which events have callback notification reception enabled.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeEventConfig(request *DescribeEventConfigRequest) (response *DescribeEventConfigResponse, err error) {
    return c.DescribeEventConfigWithContext(context.Background(), request)
}

// DescribeEventConfig
// Tencent Cloud VOD provides customers with services such as media upload, media management, and media processing. During or after the execution of these services, Tencent Cloud VOD also offers various corresponding event notifications to facilitate developers' awareness of the service processing status and to perform subsequent business operations. Developers can use this interface to query the current configuration of event notification reception methods, reception addresses, and which events have callback notification reception enabled.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeEventConfigWithContext(ctx context.Context, request *DescribeEventConfigRequest) (response *DescribeEventConfigResponse, err error) {
    if request == nil {
        request = NewDescribeEventConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileAttributesRequest() (request *DescribeFileAttributesRequest) {
    request = &DescribeFileAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeFileAttributes")
    
    
    return
}

func NewDescribeFileAttributesResponse() (response *DescribeFileAttributesResponse) {
    response = &DescribeFileAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileAttributes
// This API is used to get file attributes asynchronously.
//
// - Currently, this API can only get the MD5 hash of a file.
//
// - If the file queried is in HLS or DASH format, the attributes of the index file will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_TASKDUPLICATE = "FailedOperation.TaskDuplicate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETFILEINFOERROR = "InternalError.GetFileInfoError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileAttributes(request *DescribeFileAttributesRequest) (response *DescribeFileAttributesResponse, err error) {
    return c.DescribeFileAttributesWithContext(context.Background(), request)
}

// DescribeFileAttributes
// This API is used to get file attributes asynchronously.
//
// - Currently, this API can only get the MD5 hash of a file.
//
// - If the file queried is in HLS or DASH format, the attributes of the index file will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_TASKDUPLICATE = "FailedOperation.TaskDuplicate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETFILEINFOERROR = "InternalError.GetFileInfoError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileAttributesWithContext(ctx context.Context, request *DescribeFileAttributesRequest) (response *DescribeFileAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeFileAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHeadTailTemplatesRequest() (request *DescribeHeadTailTemplatesRequest) {
    request = &DescribeHeadTailTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeHeadTailTemplates")
    
    
    return
}

func NewDescribeHeadTailTemplatesResponse() (response *DescribeHeadTailTemplatesResponse) {
    response = &DescribeHeadTailTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHeadTailTemplates
// Describe HeadTail Templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHeadTailTemplates(request *DescribeHeadTailTemplatesRequest) (response *DescribeHeadTailTemplatesResponse, err error) {
    return c.DescribeHeadTailTemplatesWithContext(context.Background(), request)
}

// DescribeHeadTailTemplates
// Describe HeadTail Templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHeadTailTemplatesWithContext(ctx context.Context, request *DescribeHeadTailTemplatesRequest) (response *DescribeHeadTailTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeHeadTailTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHeadTailTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHeadTailTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageProcessingTemplatesRequest() (request *DescribeImageProcessingTemplatesRequest) {
    request = &DescribeImageProcessingTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeImageProcessingTemplates")
    
    
    return
}

func NewDescribeImageProcessingTemplatesResponse() (response *DescribeImageProcessingTemplatesResponse) {
    response = &DescribeImageProcessingTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageProcessingTemplates
// This API is used to query image processing templates. You can specify the filters as well as the offset to start returning records from.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageProcessingTemplates(request *DescribeImageProcessingTemplatesRequest) (response *DescribeImageProcessingTemplatesResponse, err error) {
    return c.DescribeImageProcessingTemplatesWithContext(context.Background(), request)
}

// DescribeImageProcessingTemplates
// This API is used to query image processing templates. You can specify the filters as well as the offset to start returning records from.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageProcessingTemplatesWithContext(ctx context.Context, request *DescribeImageProcessingTemplatesRequest) (response *DescribeImageProcessingTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeImageProcessingTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageProcessingTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageProcessingTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageReviewUsageDataRequest() (request *DescribeImageReviewUsageDataRequest) {
    request = &DescribeImageReviewUsageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeImageReviewUsageData")
    
    
    return
}

func NewDescribeImageReviewUsageDataResponse() (response *DescribeImageReviewUsageDataResponse) {
    response = &DescribeImageReviewUsageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageReviewUsageData
// This interface returns the image review usage information used every day within the query time range.
//
//    1. You can query the image review statistics for the last 365 days.
//
//    2. The query time span does not exceed 90 days.
//
//    3. If the query time span exceeds 1 day, data with a granularity of days will be returned. Otherwise, data with a granularity of 5 minutes will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeImageReviewUsageData(request *DescribeImageReviewUsageDataRequest) (response *DescribeImageReviewUsageDataResponse, err error) {
    return c.DescribeImageReviewUsageDataWithContext(context.Background(), request)
}

// DescribeImageReviewUsageData
// This interface returns the image review usage information used every day within the query time range.
//
//    1. You can query the image review statistics for the last 365 days.
//
//    2. The query time span does not exceed 90 days.
//
//    3. If the query time span exceeds 1 day, data with a granularity of days will be returned. Otherwise, data with a granularity of 5 minutes will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeImageReviewUsageDataWithContext(ctx context.Context, request *DescribeImageReviewUsageDataRequest) (response *DescribeImageReviewUsageDataResponse, err error) {
    if request == nil {
        request = NewDescribeImageReviewUsageDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageReviewUsageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageReviewUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageSpriteTemplatesRequest() (request *DescribeImageSpriteTemplatesRequest) {
    request = &DescribeImageSpriteTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeImageSpriteTemplates")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageSpriteTemplates(request *DescribeImageSpriteTemplatesRequest) (response *DescribeImageSpriteTemplatesResponse, err error) {
    return c.DescribeImageSpriteTemplatesWithContext(context.Background(), request)
}

// DescribeImageSpriteTemplates
// This API is used to query the list of image sprite generating templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageSpriteTemplatesWithContext(ctx context.Context, request *DescribeImageSpriteTemplatesRequest) (response *DescribeImageSpriteTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeImageSpriteTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageSpriteTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageSpriteTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJustInTimeTranscodeTemplatesRequest() (request *DescribeJustInTimeTranscodeTemplatesRequest) {
    request = &DescribeJustInTimeTranscodeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeJustInTimeTranscodeTemplates")
    
    
    return
}

func NewDescribeJustInTimeTranscodeTemplatesResponse() (response *DescribeJustInTimeTranscodeTemplatesResponse) {
    response = &DescribeJustInTimeTranscodeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJustInTimeTranscodeTemplates
// Describe Just In Time Transcode Templates.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeJustInTimeTranscodeTemplates(request *DescribeJustInTimeTranscodeTemplatesRequest) (response *DescribeJustInTimeTranscodeTemplatesResponse, err error) {
    return c.DescribeJustInTimeTranscodeTemplatesWithContext(context.Background(), request)
}

// DescribeJustInTimeTranscodeTemplates
// Describe Just In Time Transcode Templates.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeJustInTimeTranscodeTemplatesWithContext(ctx context.Context, request *DescribeJustInTimeTranscodeTemplatesRequest) (response *DescribeJustInTimeTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeJustInTimeTranscodeTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJustInTimeTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJustInTimeTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLicenseUsageDataRequest() (request *DescribeLicenseUsageDataRequest) {
    request = &DescribeLicenseUsageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeLicenseUsageData")
    
    
    return
}

func NewDescribeLicenseUsageDataResponse() (response *DescribeLicenseUsageDataResponse) {
    response = &DescribeLicenseUsageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLicenseUsageData
// This interface returns information about the number of license requests per day within the query time range.
//
//    1. You can query the license request statistics in the last 365 days.
//
//    2. The query time span does not exceed 90 days.
//
//    3. If the query time span exceeds 1 day, data with a granularity of days will be returned. Otherwise, data with a granularity of 5 minutes will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLicenseUsageData(request *DescribeLicenseUsageDataRequest) (response *DescribeLicenseUsageDataResponse, err error) {
    return c.DescribeLicenseUsageDataWithContext(context.Background(), request)
}

// DescribeLicenseUsageData
// This interface returns information about the number of license requests per day within the query time range.
//
//    1. You can query the license request statistics in the last 365 days.
//
//    2. The query time span does not exceed 90 days.
//
//    3. If the query time span exceeds 1 day, data with a granularity of days will be returned. Otherwise, data with a granularity of 5 minutes will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLicenseUsageDataWithContext(ctx context.Context, request *DescribeLicenseUsageDataRequest) (response *DescribeLicenseUsageDataResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseUsageDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicenseUsageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaInfosRequest() (request *DescribeMediaInfosRequest) {
    request = &DescribeMediaInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeMediaInfos")
    
    
    return
}

func NewDescribeMediaInfosResponse() (response *DescribeMediaInfosResponse) {
    response = &DescribeMediaInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMediaInfos
// 1. This API is used to get the information of multiple media files. Specifically, the information returned is as follows:
//
//     1. `basicInfo`: Basic information including the file name, category, playback URL, and thumbnail.
//
//     2. `metaData`: Metadata including the file size, duration, video stream information, and audio stream information.
//
//     3. `transcodeInfo`: Transcoding information including the URLs, video stream parameters, and audio stream parameters of transcoding outputs.
//
//     4. `animatedGraphicsInfo`: The information of the animated images (such as GIF images) generated.
//
//     5. `sampleSnapshotInfo`: The information of the sampled screenshots generated.
//
//     6. `imageSpriteInfo`: The information of the image sprites generated.
//
//     7. `snapshotByTimeOffsetInfo`: The information of the time point screenshots generated.
//
//     8. `keyFrameDescInfo`: The video timestamp information.
//
//     9. `adaptiveDynamicStreamingInfo`: Adaptive bitrate information including the specifications, encryption type, and formats of the streams.
//
//     10. `reviewInfo`: Moderation details for audio/video content and thumbnails.
//
// 2. You can specify what information to return.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETMEDIALISTERROR = "InternalError.GetMediaListError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_FILEIDSEMPTY = "InvalidParameterValue.FileIdsEmpty"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMediaInfos(request *DescribeMediaInfosRequest) (response *DescribeMediaInfosResponse, err error) {
    return c.DescribeMediaInfosWithContext(context.Background(), request)
}

// DescribeMediaInfos
// 1. This API is used to get the information of multiple media files. Specifically, the information returned is as follows:
//
//     1. `basicInfo`: Basic information including the file name, category, playback URL, and thumbnail.
//
//     2. `metaData`: Metadata including the file size, duration, video stream information, and audio stream information.
//
//     3. `transcodeInfo`: Transcoding information including the URLs, video stream parameters, and audio stream parameters of transcoding outputs.
//
//     4. `animatedGraphicsInfo`: The information of the animated images (such as GIF images) generated.
//
//     5. `sampleSnapshotInfo`: The information of the sampled screenshots generated.
//
//     6. `imageSpriteInfo`: The information of the image sprites generated.
//
//     7. `snapshotByTimeOffsetInfo`: The information of the time point screenshots generated.
//
//     8. `keyFrameDescInfo`: The video timestamp information.
//
//     9. `adaptiveDynamicStreamingInfo`: Adaptive bitrate information including the specifications, encryption type, and formats of the streams.
//
//     10. `reviewInfo`: Moderation details for audio/video content and thumbnails.
//
// 2. You can specify what information to return.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETMEDIALISTERROR = "InternalError.GetMediaListError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_FILEIDSEMPTY = "InvalidParameterValue.FileIdsEmpty"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMediaInfosWithContext(ctx context.Context, request *DescribeMediaInfosRequest) (response *DescribeMediaInfosResponse, err error) {
    if request == nil {
        request = NewDescribeMediaInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMediaInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMediaInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaPlayStatDetailsRequest() (request *DescribeMediaPlayStatDetailsRequest) {
    request = &DescribeMediaPlayStatDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeMediaPlayStatDetails")
    
    
    return
}

func NewDescribeMediaPlayStatDetailsResponse() (response *DescribeMediaPlayStatDetailsResponse) {
    response = &DescribeMediaPlayStatDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMediaPlayStatDetails
// This API is used to query the playback statistics of a media file at the specified granularity.
//
// * You can query playback statistics in the past year.
//
// * If the granularity is an hour, the start and end time cannot be more than seven days apart.
//
// * If the granularity is a day, the start and end time cannot be more than 90 days apart.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_INTERVAL = "InvalidParameterValue.Interval"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
func (c *Client) DescribeMediaPlayStatDetails(request *DescribeMediaPlayStatDetailsRequest) (response *DescribeMediaPlayStatDetailsResponse, err error) {
    return c.DescribeMediaPlayStatDetailsWithContext(context.Background(), request)
}

// DescribeMediaPlayStatDetails
// This API is used to query the playback statistics of a media file at the specified granularity.
//
// * You can query playback statistics in the past year.
//
// * If the granularity is an hour, the start and end time cannot be more than seven days apart.
//
// * If the granularity is a day, the start and end time cannot be more than 90 days apart.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_INTERVAL = "InvalidParameterValue.Interval"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
func (c *Client) DescribeMediaPlayStatDetailsWithContext(ctx context.Context, request *DescribeMediaPlayStatDetailsRequest) (response *DescribeMediaPlayStatDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeMediaPlayStatDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMediaPlayStatDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMediaPlayStatDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaProcessUsageDataRequest() (request *DescribeMediaProcessUsageDataRequest) {
    request = &DescribeMediaProcessUsageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeMediaProcessUsageData")
    
    
    return
}

func NewDescribeMediaProcessUsageDataResponse() (response *DescribeMediaProcessUsageDataResponse) {
    response = &DescribeMediaProcessUsageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMediaProcessUsageData
// This API is used to query the information of video processing usage within the specified time range.
//
//    1. Statistics on video processing for the last 365 days can be queried.
//
//    2. The query time range cannot be more than 90 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMediaProcessUsageData(request *DescribeMediaProcessUsageDataRequest) (response *DescribeMediaProcessUsageDataResponse, err error) {
    return c.DescribeMediaProcessUsageDataWithContext(context.Background(), request)
}

// DescribeMediaProcessUsageData
// This API is used to query the information of video processing usage within the specified time range.
//
//    1. Statistics on video processing for the last 365 days can be queried.
//
//    2. The query time range cannot be more than 90 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMediaProcessUsageDataWithContext(ctx context.Context, request *DescribeMediaProcessUsageDataRequest) (response *DescribeMediaProcessUsageDataResponse, err error) {
    if request == nil {
        request = NewDescribeMediaProcessUsageDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMediaProcessUsageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMediaProcessUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonSamplesRequest() (request *DescribePersonSamplesRequest) {
    request = &DescribePersonSamplesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribePersonSamples")
    
    
    return
}

func NewDescribePersonSamplesResponse() (response *DescribePersonSamplesResponse) {
    response = &DescribePersonSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePersonSamples
// This API is used to query the information of samples and supports paginated queries by sample ID, name, and tag.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonSamples(request *DescribePersonSamplesRequest) (response *DescribePersonSamplesResponse, err error) {
    return c.DescribePersonSamplesWithContext(context.Background(), request)
}

// DescribePersonSamples
// This API is used to query the information of samples and supports paginated queries by sample ID, name, and tag.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonSamplesWithContext(ctx context.Context, request *DescribePersonSamplesRequest) (response *DescribePersonSamplesResponse, err error) {
    if request == nil {
        request = NewDescribePersonSamplesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcedureTemplatesRequest() (request *DescribeProcedureTemplatesRequest) {
    request = &DescribeProcedureTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeProcedureTemplates")
    
    
    return
}

func NewDescribeProcedureTemplatesResponse() (response *DescribeProcedureTemplatesResponse) {
    response = &DescribeProcedureTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProcedureTemplates
// This API is used to get the list of task flow template details by task flow template name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_NAMES = "InvalidParameterValue.Names"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProcedureTemplates(request *DescribeProcedureTemplatesRequest) (response *DescribeProcedureTemplatesResponse, err error) {
    return c.DescribeProcedureTemplatesWithContext(context.Background(), request)
}

// DescribeProcedureTemplates
// This API is used to get the list of task flow template details by task flow template name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_NAMES = "InvalidParameterValue.Names"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProcedureTemplatesWithContext(ctx context.Context, request *DescribeProcedureTemplatesRequest) (response *DescribeProcedureTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeProcedureTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcedureTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcedureTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityInspectTemplatesRequest() (request *DescribeQualityInspectTemplatesRequest) {
    request = &DescribeQualityInspectTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeQualityInspectTemplates")
    
    
    return
}

func NewDescribeQualityInspectTemplatesResponse() (response *DescribeQualityInspectTemplatesResponse) {
    response = &DescribeQualityInspectTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQualityInspectTemplates
// Get media quality inspection Template List.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeQualityInspectTemplates(request *DescribeQualityInspectTemplatesRequest) (response *DescribeQualityInspectTemplatesResponse, err error) {
    return c.DescribeQualityInspectTemplatesWithContext(context.Background(), request)
}

// DescribeQualityInspectTemplates
// Get media quality inspection Template List.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeQualityInspectTemplatesWithContext(ctx context.Context, request *DescribeQualityInspectTemplatesRequest) (response *DescribeQualityInspectTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeQualityInspectTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQualityInspectTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQualityInspectTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRebuildMediaTemplatesRequest() (request *DescribeRebuildMediaTemplatesRequest) {
    request = &DescribeRebuildMediaTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeRebuildMediaTemplates")
    
    
    return
}

func NewDescribeRebuildMediaTemplatesResponse() (response *DescribeRebuildMediaTemplatesResponse) {
    response = &DescribeRebuildMediaTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRebuildMediaTemplates
// Describe Rebuild Media Templates
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRebuildMediaTemplates(request *DescribeRebuildMediaTemplatesRequest) (response *DescribeRebuildMediaTemplatesResponse, err error) {
    return c.DescribeRebuildMediaTemplatesWithContext(context.Background(), request)
}

// DescribeRebuildMediaTemplates
// Describe Rebuild Media Templates
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRebuildMediaTemplatesWithContext(ctx context.Context, request *DescribeRebuildMediaTemplatesRequest) (response *DescribeRebuildMediaTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeRebuildMediaTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRebuildMediaTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRebuildMediaTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReviewDetailsRequest() (request *DescribeReviewDetailsRequest) {
    request = &DescribeReviewDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeReviewDetails")
    
    
    return
}

func NewDescribeReviewDetailsResponse() (response *DescribeReviewDetailsResponse) {
    response = &DescribeReviewDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReviewDetails
// <b>This API is disused and replaced by [DescribeMediaProcessUsageData](https://intl.cloud.tencent.com/document/product/266/41464?from_cn_redirect=1).</b>
//
// 
//
// This API returns the video content duration for intelligent recognition in seconds per day within the queried period.
//
// 
//
// 1. The API is used to query statistics on the video content duration for intelligent recognition in the last 365 days.
//
// 2. The query period is up to 90 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeReviewDetails(request *DescribeReviewDetailsRequest) (response *DescribeReviewDetailsResponse, err error) {
    return c.DescribeReviewDetailsWithContext(context.Background(), request)
}

// DescribeReviewDetails
// <b>This API is disused and replaced by [DescribeMediaProcessUsageData](https://intl.cloud.tencent.com/document/product/266/41464?from_cn_redirect=1).</b>
//
// 
//
// This API returns the video content duration for intelligent recognition in seconds per day within the queried period.
//
// 
//
// 1. The API is used to query statistics on the video content duration for intelligent recognition in the last 365 days.
//
// 2. The query period is up to 90 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeReviewDetailsWithContext(ctx context.Context, request *DescribeReviewDetailsRequest) (response *DescribeReviewDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeReviewDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReviewDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReviewDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReviewTemplatesRequest() (request *DescribeReviewTemplatesRequest) {
    request = &DescribeReviewTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeReviewTemplates")
    
    
    return
}

func NewDescribeReviewTemplatesResponse() (response *DescribeReviewTemplatesResponse) {
    response = &DescribeReviewTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReviewTemplates
// This API is used to get the information of moderation templates.
//
// > The templates can only be used by the APIs [ReviewAudioVideo](https://intl.cloud.tencent.com/document/api/266/80283?from_cn_redirect=1) and [ReviewImage](https://intl.cloud.tencent.com/document/api/266/73217?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeReviewTemplates(request *DescribeReviewTemplatesRequest) (response *DescribeReviewTemplatesResponse, err error) {
    return c.DescribeReviewTemplatesWithContext(context.Background(), request)
}

// DescribeReviewTemplates
// This API is used to get the information of moderation templates.
//
// > The templates can only be used by the APIs [ReviewAudioVideo](https://intl.cloud.tencent.com/document/api/266/80283?from_cn_redirect=1) and [ReviewImage](https://intl.cloud.tencent.com/document/api/266/73217?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeReviewTemplatesWithContext(ctx context.Context, request *DescribeReviewTemplatesRequest) (response *DescribeReviewTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeReviewTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReviewTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReviewTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoundPlaysRequest() (request *DescribeRoundPlaysRequest) {
    request = &DescribeRoundPlaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeRoundPlays")
    
    
    return
}

func NewDescribeRoundPlaysResponse() (response *DescribeRoundPlaysResponse) {
    response = &DescribeRoundPlaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoundPlays
// This API is used to query playlists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRoundPlays(request *DescribeRoundPlaysRequest) (response *DescribeRoundPlaysResponse, err error) {
    return c.DescribeRoundPlaysWithContext(context.Background(), request)
}

// DescribeRoundPlays
// This API is used to query playlists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRoundPlaysWithContext(ctx context.Context, request *DescribeRoundPlaysRequest) (response *DescribeRoundPlaysResponse, err error) {
    if request == nil {
        request = NewDescribeRoundPlaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoundPlays require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoundPlaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleSnapshotTemplatesRequest() (request *DescribeSampleSnapshotTemplatesRequest) {
    request = &DescribeSampleSnapshotTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeSampleSnapshotTemplates")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSampleSnapshotTemplates(request *DescribeSampleSnapshotTemplatesRequest) (response *DescribeSampleSnapshotTemplatesResponse, err error) {
    return c.DescribeSampleSnapshotTemplatesWithContext(context.Background(), request)
}

// DescribeSampleSnapshotTemplates
// This API is used to query the list of sampled screencapturing templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSampleSnapshotTemplatesWithContext(ctx context.Context, request *DescribeSampleSnapshotTemplatesRequest) (response *DescribeSampleSnapshotTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSampleSnapshotTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleSnapshotTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleSnapshotTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotByTimeOffsetTemplatesRequest() (request *DescribeSnapshotByTimeOffsetTemplatesRequest) {
    request = &DescribeSnapshotByTimeOffsetTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeSnapshotByTimeOffsetTemplates")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSnapshotByTimeOffsetTemplates(request *DescribeSnapshotByTimeOffsetTemplatesRequest) (response *DescribeSnapshotByTimeOffsetTemplatesResponse, err error) {
    return c.DescribeSnapshotByTimeOffsetTemplatesWithContext(context.Background(), request)
}

// DescribeSnapshotByTimeOffsetTemplates
// This API is used to query the list of time point screencapturing templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSnapshotByTimeOffsetTemplatesWithContext(ctx context.Context, request *DescribeSnapshotByTimeOffsetTemplatesRequest) (response *DescribeSnapshotByTimeOffsetTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotByTimeOffsetTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotByTimeOffsetTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotByTimeOffsetTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageDataRequest() (request *DescribeStorageDataRequest) {
    request = &DescribeStorageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeStorageData")
    
    
    return
}

func NewDescribeStorageDataResponse() (response *DescribeStorageDataResponse) {
    response = &DescribeStorageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStorageData
// This API is used to query the storage capacity usage and number of files.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStorageData(request *DescribeStorageDataRequest) (response *DescribeStorageDataResponse, err error) {
    return c.DescribeStorageDataWithContext(context.Background(), request)
}

// DescribeStorageData
// This API is used to query the storage capacity usage and number of files.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStorageDataWithContext(ctx context.Context, request *DescribeStorageDataRequest) (response *DescribeStorageDataResponse, err error) {
    if request == nil {
        request = NewDescribeStorageDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStorageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStorageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageDetailsRequest() (request *DescribeStorageDetailsRequest) {
    request = &DescribeStorageDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeStorageDetails")
    
    
    return
}

func NewDescribeStorageDetailsResponse() (response *DescribeStorageDetailsResponse) {
    response = &DescribeStorageDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStorageDetails
// This API is used to query VOD storage usage in bytes within the query period.
//
//     1. You can only query storage usage for the last 365 days.
//
//     2. The query period is up to 90 days.
//
//     3. The query period at minute-level granularity is up to 7 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEPARSEERROR = "InternalError.TimeParseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_AREA = "InvalidParameterValue.Area"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  INVALIDPARAMETERVALUE_STORAGETYPE = "InvalidParameterValue.StorageType"
//  INVALIDPARAMETERVALUE_TIMETYPE = "InvalidParameterValue.TimeType"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStorageDetails(request *DescribeStorageDetailsRequest) (response *DescribeStorageDetailsResponse, err error) {
    return c.DescribeStorageDetailsWithContext(context.Background(), request)
}

// DescribeStorageDetails
// This API is used to query VOD storage usage in bytes within the query period.
//
//     1. You can only query storage usage for the last 365 days.
//
//     2. The query period is up to 90 days.
//
//     3. The query period at minute-level granularity is up to 7 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEPARSEERROR = "InternalError.TimeParseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_AREA = "InvalidParameterValue.Area"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  INVALIDPARAMETERVALUE_STORAGETYPE = "InvalidParameterValue.StorageType"
//  INVALIDPARAMETERVALUE_TIMETYPE = "InvalidParameterValue.TimeType"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStorageDetailsWithContext(ctx context.Context, request *DescribeStorageDetailsRequest) (response *DescribeStorageDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeStorageDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStorageDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStorageDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageRegionsRequest() (request *DescribeStorageRegionsRequest) {
    request = &DescribeStorageRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeStorageRegions")
    
    
    return
}

func NewDescribeStorageRegionsResponse() (response *DescribeStorageRegionsResponse) {
    response = &DescribeStorageRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStorageRegions
// This API is used to query the following information:
//
//   1. All supported storage regions.
//
//   2. The regions you have storage access to currently.
//
//   3. The default storage region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStorageRegions(request *DescribeStorageRegionsRequest) (response *DescribeStorageRegionsResponse, err error) {
    return c.DescribeStorageRegionsWithContext(context.Background(), request)
}

// DescribeStorageRegions
// This API is used to query the following information:
//
//   1. All supported storage regions.
//
//   2. The regions you have storage access to currently.
//
//   3. The default storage region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStorageRegionsWithContext(ctx context.Context, request *DescribeStorageRegionsRequest) (response *DescribeStorageRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeStorageRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStorageRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStorageRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubAppIdsRequest() (request *DescribeSubAppIdsRequest) {
    request = &DescribeSubAppIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeSubAppIds")
    
    
    return
}

func NewDescribeSubAppIdsResponse() (response *DescribeSubAppIdsResponse) {
    response = &DescribeSubAppIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubAppIds
// This API is used to query the list of the primary application and subapplications of the current account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSubAppIds(request *DescribeSubAppIdsRequest) (response *DescribeSubAppIdsResponse, err error) {
    return c.DescribeSubAppIdsWithContext(context.Background(), request)
}

// DescribeSubAppIds
// This API is used to query the list of the primary application and subapplications of the current account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSubAppIdsWithContext(ctx context.Context, request *DescribeSubAppIdsRequest) (response *DescribeSubAppIdsResponse, err error) {
    if request == nil {
        request = NewDescribeSubAppIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubAppIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubAppIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSuperPlayerConfigsRequest() (request *DescribeSuperPlayerConfigsRequest) {
    request = &DescribeSuperPlayerConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeSuperPlayerConfigs")
    
    
    return
}

func NewDescribeSuperPlayerConfigsResponse() (response *DescribeSuperPlayerConfigsResponse) {
    response = &DescribeSuperPlayerConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSuperPlayerConfigs
// We have <font color='red'>stopped updating</font> this API. Currently, you no longer need a player configuration to use player signatures. For details, see [Player Signature](https://intl.cloud.tencent.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to query player configurations. It supports pagination.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSuperPlayerConfigs(request *DescribeSuperPlayerConfigsRequest) (response *DescribeSuperPlayerConfigsResponse, err error) {
    return c.DescribeSuperPlayerConfigsWithContext(context.Background(), request)
}

// DescribeSuperPlayerConfigs
// We have <font color='red'>stopped updating</font> this API. Currently, you no longer need a player configuration to use player signatures. For details, see [Player Signature](https://intl.cloud.tencent.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to query player configurations. It supports pagination.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSuperPlayerConfigsWithContext(ctx context.Context, request *DescribeSuperPlayerConfigsRequest) (response *DescribeSuperPlayerConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeSuperPlayerConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSuperPlayerConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSuperPlayerConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeTaskDetail")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    return c.DescribeTaskDetailWithContext(context.Background(), request)
}

// DescribeTaskDetail
// This API is used to query the details of execution status and result of a task submitted in the last 3 days by task ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeTasks")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_STATUS = "InvalidParameterValue.Status"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// * This API is used to query the task list;
//
// * If there are many data entries in the list, one single call of the API may not be able to pull the entire list. The `ScrollToken` parameter can be used to pull the list in batches;
//
// * Only tasks in the last three days (72 hours) can be queried.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_STATUS = "InvalidParameterValue.Status"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeTranscodeTemplates")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTranscodeTemplates(request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    return c.DescribeTranscodeTemplatesWithContext(context.Background(), request)
}

// DescribeTranscodeTemplates
// This API is used to get the list of transcoding templates based on unique template ID. The return result includes all eligible custom and [preset transcoding templates](https://intl.cloud.tencent.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTranscodeTemplatesWithContext(ctx context.Context, request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVodDomainsRequest() (request *DescribeVodDomainsRequest) {
    request = &DescribeVodDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeVodDomains")
    
    
    return
}

func NewDescribeVodDomainsResponse() (response *DescribeVodDomainsResponse) {
    response = &DescribeVodDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVodDomains
// This API is used to query the list of VOD domain names.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_LIMITTOOLARGE = "InvalidParameterValue.LimitTooLarge"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_OFFSETTOOLARGE = "InvalidParameterValue.OffsetTooLarge"
func (c *Client) DescribeVodDomains(request *DescribeVodDomainsRequest) (response *DescribeVodDomainsResponse, err error) {
    return c.DescribeVodDomainsWithContext(context.Background(), request)
}

// DescribeVodDomains
// This API is used to query the list of VOD domain names.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_LIMITTOOLARGE = "InvalidParameterValue.LimitTooLarge"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_OFFSETTOOLARGE = "InvalidParameterValue.OffsetTooLarge"
func (c *Client) DescribeVodDomainsWithContext(ctx context.Context, request *DescribeVodDomainsRequest) (response *DescribeVodDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeVodDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVodDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVodDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWatermarkTemplatesRequest() (request *DescribeWatermarkTemplatesRequest) {
    request = &DescribeWatermarkTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeWatermarkTemplates")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWatermarkTemplates(request *DescribeWatermarkTemplatesRequest) (response *DescribeWatermarkTemplatesResponse, err error) {
    return c.DescribeWatermarkTemplatesWithContext(context.Background(), request)
}

// DescribeWatermarkTemplates
// This API is used to query custom watermarking templates and supports paged queries by filters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWatermarkTemplatesWithContext(ctx context.Context, request *DescribeWatermarkTemplatesRequest) (response *DescribeWatermarkTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeWatermarkTemplatesRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeWordSamples")
    
    
    return
}

func NewDescribeWordSamplesResponse() (response *DescribeWordSamplesResponse) {
    response = &DescribeWordSamplesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWordSamples
// This API is used to perform paginated queries of keyword sample information by use case, keyword, and tag.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWordSamples(request *DescribeWordSamplesRequest) (response *DescribeWordSamplesResponse, err error) {
    return c.DescribeWordSamplesWithContext(context.Background(), request)
}

// DescribeWordSamples
// This API is used to perform paginated queries of keyword sample information by use case, keyword, and tag.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWordSamplesWithContext(ctx context.Context, request *DescribeWordSamplesRequest) (response *DescribeWordSamplesResponse, err error) {
    if request == nil {
        request = NewDescribeWordSamplesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWordSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewEditMediaRequest() (request *EditMediaRequest) {
    request = &EditMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "EditMedia")
    
    
    return
}

func NewEditMediaResponse() (response *EditMediaResponse) {
    response = &EditMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EditMedia
// Edit the video (cut, splice, etc.) to generate a new video. The editing functions include:
//
// 
//
// 1. Edit a file in the VOD video to generate a new video;
//
// 2. Splice multiple files to generate a new video;
//
// 3. Edit multiple files and then splice them to generate a new video;
//
// 4. Directly generate a new video for a stream;
//
// 5. Edit a stream to generate a new video. Video;
//
// 6. Splice multiple streams to generate a new video;
//
// 7. Clip and then splice multiple streams to generate a new video. 
//
// 
//
// For the generated new video, you can also specify whether the generated video needs to execute the task flow. 
//
// 
//
// >When editing or splicing live streams, please make sure to do so after the stream is over. Otherwise the resulting video may be incomplete. 
//
// 
//
// If event notification is used, the type of event notification is [Video editing completed](https://intl.cloud.tencent.com/document/product/266/33794?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EditMedia(request *EditMediaRequest) (response *EditMediaResponse, err error) {
    return c.EditMediaWithContext(context.Background(), request)
}

// EditMedia
// Edit the video (cut, splice, etc.) to generate a new video. The editing functions include:
//
// 
//
// 1. Edit a file in the VOD video to generate a new video;
//
// 2. Splice multiple files to generate a new video;
//
// 3. Edit multiple files and then splice them to generate a new video;
//
// 4. Directly generate a new video for a stream;
//
// 5. Edit a stream to generate a new video. Video;
//
// 6. Splice multiple streams to generate a new video;
//
// 7. Clip and then splice multiple streams to generate a new video. 
//
// 
//
// For the generated new video, you can also specify whether the generated video needs to execute the task flow. 
//
// 
//
// >When editing or splicing live streams, please make sure to do so after the stream is over. Otherwise the resulting video may be incomplete. 
//
// 
//
// If event notification is used, the type of event notification is [Video editing completed](https://intl.cloud.tencent.com/document/product/266/33794?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EditMediaWithContext(ctx context.Context, request *EditMediaRequest) (response *EditMediaResponse, err error) {
    if request == nil {
        request = NewEditMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditMediaResponse()
    err = c.Send(request, response)
    return
}

func NewEnhanceMediaByTemplateRequest() (request *EnhanceMediaByTemplateRequest) {
    request = &EnhanceMediaByTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "EnhanceMediaByTemplate")
    
    
    return
}

func NewEnhanceMediaByTemplateResponse() (response *EnhanceMediaByTemplateResponse) {
    response = &EnhanceMediaByTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnhanceMediaByTemplate
// Enhance Media By Template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EnhanceMediaByTemplate(request *EnhanceMediaByTemplateRequest) (response *EnhanceMediaByTemplateResponse, err error) {
    return c.EnhanceMediaByTemplateWithContext(context.Background(), request)
}

// EnhanceMediaByTemplate
// Enhance Media By Template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EnhanceMediaByTemplateWithContext(ctx context.Context, request *EnhanceMediaByTemplateRequest) (response *EnhanceMediaByTemplateResponse, err error) {
    if request == nil {
        request = NewEnhanceMediaByTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnhanceMediaByTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnhanceMediaByTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewEnhanceMediaQualityRequest() (request *EnhanceMediaQualityRequest) {
    request = &EnhanceMediaQualityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "EnhanceMediaQuality")
    
    
    return
}

func NewEnhanceMediaQualityResponse() (response *EnhanceMediaQualityResponse) {
    response = &EnhanceMediaQualityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnhanceMediaQuality
// Initiate a Remaster task for audio and video media in VOD
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) EnhanceMediaQuality(request *EnhanceMediaQualityRequest) (response *EnhanceMediaQualityResponse, err error) {
    return c.EnhanceMediaQualityWithContext(context.Background(), request)
}

// EnhanceMediaQuality
// Initiate a Remaster task for audio and video media in VOD
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) EnhanceMediaQualityWithContext(ctx context.Context, request *EnhanceMediaQualityRequest) (response *EnhanceMediaQualityResponse, err error) {
    if request == nil {
        request = NewEnhanceMediaQualityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnhanceMediaQuality require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnhanceMediaQualityResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteFunctionRequest() (request *ExecuteFunctionRequest) {
    request = &ExecuteFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ExecuteFunction")
    
    
    return
}

func NewExecuteFunctionResponse() (response *ExecuteFunctionResponse) {
    response = &ExecuteFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExecuteFunction
// This API is only used in special scenarios of custom development. Unless requested by VOD customer service, please do not call it.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExecuteFunction(request *ExecuteFunctionRequest) (response *ExecuteFunctionResponse, err error) {
    return c.ExecuteFunctionWithContext(context.Background(), request)
}

// ExecuteFunction
// This API is only used in special scenarios of custom development. Unless requested by VOD customer service, please do not call it.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExecuteFunctionWithContext(ctx context.Context, request *ExecuteFunctionRequest) (response *ExecuteFunctionResponse, err error) {
    if request == nil {
        request = NewExecuteFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewExtractCopyRightWatermarkRequest() (request *ExtractCopyRightWatermarkRequest) {
    request = &ExtractCopyRightWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ExtractCopyRightWatermark")
    
    
    return
}

func NewExtractCopyRightWatermarkResponse() (response *ExtractCopyRightWatermarkResponse) {
    response = &ExtractCopyRightWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExtractCopyRightWatermark
// Extract CopyRight Watermark.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExtractCopyRightWatermark(request *ExtractCopyRightWatermarkRequest) (response *ExtractCopyRightWatermarkResponse, err error) {
    return c.ExtractCopyRightWatermarkWithContext(context.Background(), request)
}

// ExtractCopyRightWatermark
// Extract CopyRight Watermark.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExtractCopyRightWatermarkWithContext(ctx context.Context, request *ExtractCopyRightWatermarkRequest) (response *ExtractCopyRightWatermarkResponse, err error) {
    if request == nil {
        request = NewExtractCopyRightWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExtractCopyRightWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewExtractCopyRightWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewExtractTraceWatermarkRequest() (request *ExtractTraceWatermarkRequest) {
    request = &ExtractTraceWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ExtractTraceWatermark")
    
    
    return
}

func NewExtractTraceWatermarkResponse() (response *ExtractTraceWatermarkResponse) {
    response = &ExtractTraceWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExtractTraceWatermark
// This API is used to extract the user ID of a user that distributed a video containing a digital watermark.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExtractTraceWatermark(request *ExtractTraceWatermarkRequest) (response *ExtractTraceWatermarkResponse, err error) {
    return c.ExtractTraceWatermarkWithContext(context.Background(), request)
}

// ExtractTraceWatermark
// This API is used to extract the user ID of a user that distributed a video containing a digital watermark.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ExtractTraceWatermarkWithContext(ctx context.Context, request *ExtractTraceWatermarkRequest) (response *ExtractTraceWatermarkResponse, err error) {
    if request == nil {
        request = NewExtractTraceWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExtractTraceWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewExtractTraceWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewFastEditMediaRequest() (request *FastEditMediaRequest) {
    request = &FastEditMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "FastEditMedia")
    
    
    return
}

func NewFastEditMediaResponse() (response *FastEditMediaResponse) {
    response = &FastEditMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FastEditMedia
// Implement quick concatenation and quick clipping for Tencent Cloud VOD's HLS videos to generate new media in HLS format. The video generated by quick concatenation or clipping will produce a new FileId and undergo solidification. After successful solidification, the new video file exists independently of the original input video and is not affected by deletions or other actions on the original video. <font color='red'>Note:</font> Enable the reception of persistence completed event notifications through the ModifyEvent config interface, and a PersistenceComplete type event notification will be received after successful solidification. Before receiving this event notification, operations such as deletion or cooling down of the original input video should not be performed, otherwise, abnormal playback may occur in the video generated by the concatenation and clipping.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) FastEditMedia(request *FastEditMediaRequest) (response *FastEditMediaResponse, err error) {
    return c.FastEditMediaWithContext(context.Background(), request)
}

// FastEditMedia
// Implement quick concatenation and quick clipping for Tencent Cloud VOD's HLS videos to generate new media in HLS format. The video generated by quick concatenation or clipping will produce a new FileId and undergo solidification. After successful solidification, the new video file exists independently of the original input video and is not affected by deletions or other actions on the original video. <font color='red'>Note:</font> Enable the reception of persistence completed event notifications through the ModifyEvent config interface, and a PersistenceComplete type event notification will be received after successful solidification. Before receiving this event notification, operations such as deletion or cooling down of the original input video should not be performed, otherwise, abnormal playback may occur in the video generated by the concatenation and clipping.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) FastEditMediaWithContext(ctx context.Context, request *FastEditMediaRequest) (response *FastEditMediaResponse, err error) {
    if request == nil {
        request = NewFastEditMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FastEditMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewFastEditMediaResponse()
    err = c.Send(request, response)
    return
}

func NewForbidMediaDistributionRequest() (request *ForbidMediaDistributionRequest) {
    request = &ForbidMediaDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ForbidMediaDistribution")
    
    
    return
}

func NewForbidMediaDistributionResponse() (response *ForbidMediaDistributionResponse) {
    response = &ForbidMediaDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ForbidMediaDistribution
// * After a media file is forbidden, except previewing it in the VOD Console, accessing the URLs of its various resources (such as source file, output files, and screenshots) in other scenarios will return error 403.
//
//   It takes about 5-10 minutes for a forbidding/unblocking operation to take effect across the entire network.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_MEDIAFORBIDEDBYSYSTEM = "FailedOperation.MediaForbidedBySystem"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEIDSTOOMANY = "InvalidParameterValue.FileIdsTooMany"
//  INVALIDPARAMETERVALUE_OPERATION = "InvalidParameterValue.Operation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ForbidMediaDistribution(request *ForbidMediaDistributionRequest) (response *ForbidMediaDistributionResponse, err error) {
    return c.ForbidMediaDistributionWithContext(context.Background(), request)
}

// ForbidMediaDistribution
// * After a media file is forbidden, except previewing it in the VOD Console, accessing the URLs of its various resources (such as source file, output files, and screenshots) in other scenarios will return error 403.
//
//   It takes about 5-10 minutes for a forbidding/unblocking operation to take effect across the entire network.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_MEDIAFORBIDEDBYSYSTEM = "FailedOperation.MediaForbidedBySystem"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEIDSTOOMANY = "InvalidParameterValue.FileIdsTooMany"
//  INVALIDPARAMETERVALUE_OPERATION = "InvalidParameterValue.Operation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ForbidMediaDistributionWithContext(ctx context.Context, request *ForbidMediaDistributionRequest) (response *ForbidMediaDistributionResponse, err error) {
    if request == nil {
        request = NewForbidMediaDistributionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForbidMediaDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewForbidMediaDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewHandleCurrentPlaylistRequest() (request *HandleCurrentPlaylistRequest) {
    request = &HandleCurrentPlaylistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "HandleCurrentPlaylist")
    
    
    return
}

func NewHandleCurrentPlaylistResponse() (response *HandleCurrentPlaylistResponse) {
    response = &HandleCurrentPlaylistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// HandleCurrentPlaylist
// Operate the current play list . Supported operations include:<li> Insert: Insert a playing program into the current playlist.</li><li> Delete: Remove a playing program from the playlist.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_INDEX = "InvalidParameterValue.Index"
//  INVALIDPARAMETERVALUE_ITEMID = "InvalidParameterValue.ItemId"
//  INVALIDPARAMETERVALUE_ROUNDPLAYID = "InvalidParameterValue.RoundPlayId"
//  INVALIDPARAMETERVALUE_ROUNDPLAYLIST = "InvalidParameterValue.RoundPlaylist"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_PLAYLIST = "LimitExceeded.PlayList"
func (c *Client) HandleCurrentPlaylist(request *HandleCurrentPlaylistRequest) (response *HandleCurrentPlaylistResponse, err error) {
    return c.HandleCurrentPlaylistWithContext(context.Background(), request)
}

// HandleCurrentPlaylist
// Operate the current play list . Supported operations include:<li> Insert: Insert a playing program into the current playlist.</li><li> Delete: Remove a playing program from the playlist.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_INDEX = "InvalidParameterValue.Index"
//  INVALIDPARAMETERVALUE_ITEMID = "InvalidParameterValue.ItemId"
//  INVALIDPARAMETERVALUE_ROUNDPLAYID = "InvalidParameterValue.RoundPlayId"
//  INVALIDPARAMETERVALUE_ROUNDPLAYLIST = "InvalidParameterValue.RoundPlaylist"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_PLAYLIST = "LimitExceeded.PlayList"
func (c *Client) HandleCurrentPlaylistWithContext(ctx context.Context, request *HandleCurrentPlaylistRequest) (response *HandleCurrentPlaylistResponse, err error) {
    if request == nil {
        request = NewHandleCurrentPlaylistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HandleCurrentPlaylist require credential")
    }

    request.SetContext(ctx)
    
    response = NewHandleCurrentPlaylistResponse()
    err = c.Send(request, response)
    return
}

func NewInspectMediaQualityRequest() (request *InspectMediaQualityRequest) {
    request = &InspectMediaQualityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "InspectMediaQuality")
    
    
    return
}

func NewInspectMediaQualityResponse() (response *InspectMediaQualityResponse) {
    response = &InspectMediaQualityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InspectMediaQuality
// Initiate media quality inspection task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) InspectMediaQuality(request *InspectMediaQualityRequest) (response *InspectMediaQualityResponse, err error) {
    return c.InspectMediaQualityWithContext(context.Background(), request)
}

// InspectMediaQuality
// Initiate media quality inspection task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) InspectMediaQualityWithContext(ctx context.Context, request *InspectMediaQualityRequest) (response *InspectMediaQualityResponse, err error) {
    if request == nil {
        request = NewInspectMediaQualityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InspectMediaQuality require credential")
    }

    request.SetContext(ctx)
    
    response = NewInspectMediaQualityResponse()
    err = c.Send(request, response)
    return
}

func NewLiveRealTimeClipRequest() (request *LiveRealTimeClipRequest) {
    request = &LiveRealTimeClipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "LiveRealTimeClip")
    
    
    return
}

func NewLiveRealTimeClipResponse() (response *LiveRealTimeClipResponse) {
    response = &LiveRealTimeClipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LiveRealTimeClip
// Live broadcast real-time editing means that during the live broadcast process (that is, before the live broadcast has ended), customers can select a section of past live broadcast content to generate a new video (HLS format) in real time. Developers can share it immediately, or permanently Save it.
//
// 
//
// Tencent Cloud VOD supports two real-time editing modes:
//
// - Editing and solidification: Save the edited video as an independent video with an independent FileId; suitable for scenes where the highlight clips are saved for a long time;
//
// - Editing is not solidified: the edited video is attached to the live broadcast recording file and does not have an independent FileId; it is suitable for scenarios where highlight clips are **temporarily shared**.
//
// 
//
// Notice:
//
// - The prerequisite for using the live broadcast real-time editing function is that the target live stream has the [Time Shift Replay](https://intl.cloud.tencent.com/document/product/267/32742?from_cn_redirect=1) function enabled.
//
// - Live broadcast real-time editing is based on the m3u8 file generated by live broadcast recording, so its minimum editing accuracy is one ts slice, and it is impossible to achieve second-level or more precise editing accuracy.
//
// - Since the stream may be interrupted during the live broadcast, the actual video duration generated by editing may be inconsistent with expectations. For example, the time interval for editing a live stream is from 2018-09-20T10:30:00Z to 2018-09-20T10:40:00Z. If the stream is interrupted during this time interval, the duration of the returned media asset file will be Less than 10 minutes, in which case it can be sensed via the output parameter <a href="#p_segmentset">SegmentSet</a>.
//
// 
//
// ### Clip solidification
//
// Clip solidification means saving the clipped video into an independent video (with an independent FileId). Its life cycle is not affected by the original live broadcast recording video (even if the original recording video is deleted, the editing results will not be affected in any way); it can also be subjected to secondary processing such as transcoding.
//
// 
//
// For example: for a complete football match, the original video recorded live may be as long as 2 hours. To save costs, the customer can store this video for 2 months, but for the "highlight moment" video that is edited in real time during the live broadcast However, you can specify a longer storage period, and at the same time, you can separately transcode the "highlight moments" video and other on-demand operations. At this time, you can choose the solution of real-time editing and solidification of the live broadcast.
//
// 
//
// The advantage of editing and curing is that its life cycle is independent of the original recorded video, and can be managed independently and stored for a long time.
//
// 
//
// <font color='red'>Notice:</font> If solidification is specified during clipping, enable the reception of persistence completed event notifications through the ModifyEventConfig interface. After successful solidification, an event notification of type PersistenceComplete will be received. Before receiving this event notification, operations such as deletion or cooling down of live recording videos should not be performed, otherwise, abnormal playback may occur in the video generated by the clipping.
//
// 
//
// ### Clips are not solidified
//
// The editing is not solidified, which means that the result of editing (m3u8 file) and the live recording video share the same ts fragment. The newly generated video is not an independent and complete video (no independent FileId, only playback URL), and its validity period is the same as the live broadcast. The validity period of the complete recorded video is the same. Once the video recorded during the live broadcast is deleted, the clip will also become unplayable.
//
// 
//
// The clip is not solidified. Because the clip result is not an independent video, it will not be included in on-demand media asset video management (for example, the total number of videos in the console will not count this clip), and this clip cannot be transcoded separately. Video processing operations.
//
// 
//
// The advantage of not solidifying editing is that its editing operation is very "lightweight" and does not generate additional storage overhead. However, its disadvantage is that the life cycle is the same as the original recorded video, and further video processing such as transcoding cannot be performed.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLIPDURATION = "InvalidParameterValue.ClipDuration"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  INVALIDPARAMETERVALUE_STREAMIDINVALID = "InvalidParameterValue.StreamIdInvalid"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LiveRealTimeClip(request *LiveRealTimeClipRequest) (response *LiveRealTimeClipResponse, err error) {
    return c.LiveRealTimeClipWithContext(context.Background(), request)
}

// LiveRealTimeClip
// Live broadcast real-time editing means that during the live broadcast process (that is, before the live broadcast has ended), customers can select a section of past live broadcast content to generate a new video (HLS format) in real time. Developers can share it immediately, or permanently Save it.
//
// 
//
// Tencent Cloud VOD supports two real-time editing modes:
//
// - Editing and solidification: Save the edited video as an independent video with an independent FileId; suitable for scenes where the highlight clips are saved for a long time;
//
// - Editing is not solidified: the edited video is attached to the live broadcast recording file and does not have an independent FileId; it is suitable for scenarios where highlight clips are **temporarily shared**.
//
// 
//
// Notice:
//
// - The prerequisite for using the live broadcast real-time editing function is that the target live stream has the [Time Shift Replay](https://intl.cloud.tencent.com/document/product/267/32742?from_cn_redirect=1) function enabled.
//
// - Live broadcast real-time editing is based on the m3u8 file generated by live broadcast recording, so its minimum editing accuracy is one ts slice, and it is impossible to achieve second-level or more precise editing accuracy.
//
// - Since the stream may be interrupted during the live broadcast, the actual video duration generated by editing may be inconsistent with expectations. For example, the time interval for editing a live stream is from 2018-09-20T10:30:00Z to 2018-09-20T10:40:00Z. If the stream is interrupted during this time interval, the duration of the returned media asset file will be Less than 10 minutes, in which case it can be sensed via the output parameter <a href="#p_segmentset">SegmentSet</a>.
//
// 
//
// ### Clip solidification
//
// Clip solidification means saving the clipped video into an independent video (with an independent FileId). Its life cycle is not affected by the original live broadcast recording video (even if the original recording video is deleted, the editing results will not be affected in any way); it can also be subjected to secondary processing such as transcoding.
//
// 
//
// For example: for a complete football match, the original video recorded live may be as long as 2 hours. To save costs, the customer can store this video for 2 months, but for the "highlight moment" video that is edited in real time during the live broadcast However, you can specify a longer storage period, and at the same time, you can separately transcode the "highlight moments" video and other on-demand operations. At this time, you can choose the solution of real-time editing and solidification of the live broadcast.
//
// 
//
// The advantage of editing and curing is that its life cycle is independent of the original recorded video, and can be managed independently and stored for a long time.
//
// 
//
// <font color='red'>Notice:</font> If solidification is specified during clipping, enable the reception of persistence completed event notifications through the ModifyEventConfig interface. After successful solidification, an event notification of type PersistenceComplete will be received. Before receiving this event notification, operations such as deletion or cooling down of live recording videos should not be performed, otherwise, abnormal playback may occur in the video generated by the clipping.
//
// 
//
// ### Clips are not solidified
//
// The editing is not solidified, which means that the result of editing (m3u8 file) and the live recording video share the same ts fragment. The newly generated video is not an independent and complete video (no independent FileId, only playback URL), and its validity period is the same as the live broadcast. The validity period of the complete recorded video is the same. Once the video recorded during the live broadcast is deleted, the clip will also become unplayable.
//
// 
//
// The clip is not solidified. Because the clip result is not an independent video, it will not be included in on-demand media asset video management (for example, the total number of videos in the console will not count this clip), and this clip cannot be transcoded separately. Video processing operations.
//
// 
//
// The advantage of not solidifying editing is that its editing operation is very "lightweight" and does not generate additional storage overhead. However, its disadvantage is that the life cycle is the same as the original recorded video, and further video processing such as transcoding cannot be performed.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLIPDURATION = "InvalidParameterValue.ClipDuration"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  INVALIDPARAMETERVALUE_STREAMIDINVALID = "InvalidParameterValue.StreamIdInvalid"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LiveRealTimeClipWithContext(ctx context.Context, request *LiveRealTimeClipRequest) (response *LiveRealTimeClipResponse, err error) {
    if request == nil {
        request = NewLiveRealTimeClipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LiveRealTimeClip require credential")
    }

    request.SetContext(ctx)
    
    response = NewLiveRealTimeClipResponse()
    err = c.Send(request, response)
    return
}

func NewManageTaskRequest() (request *ManageTaskRequest) {
    request = &ManageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ManageTask")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDOPERATIONTYPE = "InvalidParameterValue.InvalidOperationType"
//  INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ManageTask(request *ManageTaskRequest) (response *ManageTaskResponse, err error) {
    return c.ManageTaskWithContext(context.Background(), request)
}

// ManageTask
// This API is used to manage initiated tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDOPERATIONTYPE = "InvalidParameterValue.InvalidOperationType"
//  INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ManageTaskWithContext(ctx context.Context, request *ManageTaskRequest) (response *ManageTaskResponse, err error) {
    if request == nil {
        request = NewManageTaskRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyAIAnalysisTemplate")
    
    
    return
}

func NewModifyAIAnalysisTemplateResponse() (response *ModifyAIAnalysisTemplateResponse) {
    response = &ModifyAIAnalysisTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAIAnalysisTemplate
// This API is used to modify a custom video content analysis template.
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
//  INVALIDPARAMETERVALUE_HIGHLIGHTCONFIGURE = "InvalidParameterValue.HighlightConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAIAnalysisTemplate(request *ModifyAIAnalysisTemplateRequest) (response *ModifyAIAnalysisTemplateResponse, err error) {
    return c.ModifyAIAnalysisTemplateWithContext(context.Background(), request)
}

// ModifyAIAnalysisTemplate
// This API is used to modify a custom video content analysis template.
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
//  INVALIDPARAMETERVALUE_HIGHLIGHTCONFIGURE = "InvalidParameterValue.HighlightConfigure"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAIAnalysisTemplateWithContext(ctx context.Context, request *ModifyAIAnalysisTemplateRequest) (response *ModifyAIAnalysisTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIAnalysisTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyAIRecognitionTemplate")
    
    
    return
}

func NewModifyAIRecognitionTemplateResponse() (response *ModifyAIRecognitionTemplateResponse) {
    response = &ModifyAIRecognitionTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAIRecognitionTemplate
// This API is used to modify a custom video content recognition template.
//
// error code that may be returned:
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
//  INVALIDPARAMETERVALUE_SCREENSHOTINTERVAL = "InvalidParameterValue.ScreenshotInterval"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAIRecognitionTemplate(request *ModifyAIRecognitionTemplateRequest) (response *ModifyAIRecognitionTemplateResponse, err error) {
    return c.ModifyAIRecognitionTemplateWithContext(context.Background(), request)
}

// ModifyAIRecognitionTemplate
// This API is used to modify a custom video content recognition template.
//
// error code that may be returned:
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
//  INVALIDPARAMETERVALUE_SCREENSHOTINTERVAL = "InvalidParameterValue.ScreenshotInterval"
//  INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAIRecognitionTemplateWithContext(ctx context.Context, request *ModifyAIRecognitionTemplateRequest) (response *ModifyAIRecognitionTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAIRecognitionTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyAdaptiveDynamicStreamingTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_DRMTYPE = "InvalidParameterValue.DrmType"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAdaptiveDynamicStreamingTemplate(request *ModifyAdaptiveDynamicStreamingTemplateRequest) (response *ModifyAdaptiveDynamicStreamingTemplateResponse, err error) {
    return c.ModifyAdaptiveDynamicStreamingTemplateWithContext(context.Background(), request)
}

// ModifyAdaptiveDynamicStreamingTemplate
// This API is used to modify an adaptive bitrate streaming template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"
//  INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"
//  INVALIDPARAMETERVALUE_DRMTYPE = "InvalidParameterValue.DrmType"
//  INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyAdaptiveDynamicStreamingTemplateWithContext(ctx context.Context, request *ModifyAdaptiveDynamicStreamingTemplateRequest) (response *ModifyAdaptiveDynamicStreamingTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAdaptiveDynamicStreamingTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyAnimatedGraphicsTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAnimatedGraphicsTemplate(request *ModifyAnimatedGraphicsTemplateRequest) (response *ModifyAnimatedGraphicsTemplateResponse, err error) {
    return c.ModifyAnimatedGraphicsTemplateWithContext(context.Background(), request)
}

// ModifyAnimatedGraphicsTemplate
// This API is used to modify a custom animated image generating template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAnimatedGraphicsTemplateWithContext(ctx context.Context, request *ModifyAnimatedGraphicsTemplateRequest) (response *ModifyAnimatedGraphicsTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAnimatedGraphicsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCDNDomainConfigRequest() (request *ModifyCDNDomainConfigRequest) {
    request = &ModifyCDNDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyCDNDomainConfig")
    
    
    return
}

func NewModifyCDNDomainConfigResponse() (response *ModifyCDNDomainConfigResponse) {
    response = &ModifyCDNDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCDNDomainConfig
// Modify CDN Domain Config.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCDNDomainConfig(request *ModifyCDNDomainConfigRequest) (response *ModifyCDNDomainConfigResponse, err error) {
    return c.ModifyCDNDomainConfigWithContext(context.Background(), request)
}

// ModifyCDNDomainConfig
// Modify CDN Domain Config.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCDNDomainConfigWithContext(ctx context.Context, request *ModifyCDNDomainConfigRequest) (response *ModifyCDNDomainConfigResponse, err error) {
    if request == nil {
        request = NewModifyCDNDomainConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCDNDomainConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCDNDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClassRequest() (request *ModifyClassRequest) {
    request = &ModifyClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyClass")
    
    
    return
}

func NewModifyClassResponse() (response *ModifyClassResponse) {
    response = &ModifyClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClass
// This API is used to modify the category of a media file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSNAMEDUPLICATE = "FailedOperation.ClassNameDuplicate"
//  FAILEDOPERATION_CLASSNOFOUND = "FailedOperation.ClassNoFound"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSID = "InvalidParameterValue.ClassId"
//  INVALIDPARAMETERVALUE_CLASSNAME = "InvalidParameterValue.ClassName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyClass(request *ModifyClassRequest) (response *ModifyClassResponse, err error) {
    return c.ModifyClassWithContext(context.Background(), request)
}

// ModifyClass
// This API is used to modify the category of a media file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSNAMEDUPLICATE = "FailedOperation.ClassNameDuplicate"
//  FAILEDOPERATION_CLASSNOFOUND = "FailedOperation.ClassNoFound"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLASSID = "InvalidParameterValue.ClassId"
//  INVALIDPARAMETERVALUE_CLASSNAME = "InvalidParameterValue.ClassName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyClassWithContext(ctx context.Context, request *ModifyClassRequest) (response *ModifyClassResponse, err error) {
    if request == nil {
        request = NewModifyClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClass require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClassResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContentReviewTemplateRequest() (request *ModifyContentReviewTemplateRequest) {
    request = &ModifyContentReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyContentReviewTemplate")
    
    
    return
}

func NewModifyContentReviewTemplateResponse() (response *ModifyContentReviewTemplateResponse) {
    response = &ModifyContentReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyContentReviewTemplate
// We have <font color=red>stopped updating</font> this API. Our new moderation templates can moderate audio/video as well as images. For details, see [ModifyReviewTemplate](https://intl.cloud.tencent.com/document/api/266/84388?from_cn_redirect=1).
//
// This API is used to modify a custom audio/video moderation template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_REVIEWWALLSWITCH = "InvalidParameterValue.ReviewWallSwitch"
//  INVALIDPARAMETERVALUE_SCREENSHOTINTERVAL = "InvalidParameterValue.ScreenshotInterval"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyContentReviewTemplate(request *ModifyContentReviewTemplateRequest) (response *ModifyContentReviewTemplateResponse, err error) {
    return c.ModifyContentReviewTemplateWithContext(context.Background(), request)
}

// ModifyContentReviewTemplate
// We have <font color=red>stopped updating</font> this API. Our new moderation templates can moderate audio/video as well as images. For details, see [ModifyReviewTemplate](https://intl.cloud.tencent.com/document/api/266/84388?from_cn_redirect=1).
//
// This API is used to modify a custom audio/video moderation template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"
//  INVALIDPARAMETERVALUE_REVIEWWALLSWITCH = "InvalidParameterValue.ReviewWallSwitch"
//  INVALIDPARAMETERVALUE_SCREENSHOTINTERVAL = "InvalidParameterValue.ScreenshotInterval"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyContentReviewTemplateWithContext(ctx context.Context, request *ModifyContentReviewTemplateRequest) (response *ModifyContentReviewTemplateResponse, err error) {
    if request == nil {
        request = NewModifyContentReviewTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDefaultStorageRegionRequest() (request *ModifyDefaultStorageRegionRequest) {
    request = &ModifyDefaultStorageRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyDefaultStorageRegion")
    
    
    return
}

func NewModifyDefaultStorageRegionResponse() (response *ModifyDefaultStorageRegionResponse) {
    response = &ModifyDefaultStorageRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDefaultStorageRegion
// This API is used to set the default storage region. A file will be stored in the default region if no region is specified for file upload.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_STORAGEREGION = "InvalidParameterValue.StorageRegion"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDefaultStorageRegion(request *ModifyDefaultStorageRegionRequest) (response *ModifyDefaultStorageRegionResponse, err error) {
    return c.ModifyDefaultStorageRegionWithContext(context.Background(), request)
}

// ModifyDefaultStorageRegion
// This API is used to set the default storage region. A file will be stored in the default region if no region is specified for file upload.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_STORAGEREGION = "InvalidParameterValue.StorageRegion"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDefaultStorageRegionWithContext(ctx context.Context, request *ModifyDefaultStorageRegionRequest) (response *ModifyDefaultStorageRegionResponse, err error) {
    if request == nil {
        request = NewModifyDefaultStorageRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDefaultStorageRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDefaultStorageRegionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnhanceMediaTemplateRequest() (request *ModifyEnhanceMediaTemplateRequest) {
    request = &ModifyEnhanceMediaTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyEnhanceMediaTemplate")
    
    
    return
}

func NewModifyEnhanceMediaTemplateResponse() (response *ModifyEnhanceMediaTemplateResponse) {
    response = &ModifyEnhanceMediaTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnhanceMediaTemplate
// Modify enhance media template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_STORAGEREGION = "InvalidParameterValue.StorageRegion"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyEnhanceMediaTemplate(request *ModifyEnhanceMediaTemplateRequest) (response *ModifyEnhanceMediaTemplateResponse, err error) {
    return c.ModifyEnhanceMediaTemplateWithContext(context.Background(), request)
}

// ModifyEnhanceMediaTemplate
// Modify enhance media template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_STORAGEREGION = "InvalidParameterValue.StorageRegion"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyEnhanceMediaTemplateWithContext(ctx context.Context, request *ModifyEnhanceMediaTemplateRequest) (response *ModifyEnhanceMediaTemplateResponse, err error) {
    if request == nil {
        request = NewModifyEnhanceMediaTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnhanceMediaTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnhanceMediaTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEventConfigRequest() (request *ModifyEventConfigRequest) {
    request = &ModifyEventConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyEventConfig")
    
    
    return
}

func NewModifyEventConfigResponse() (response *ModifyEventConfigResponse) {
    response = &ModifyEventConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEventConfig
// Tencent Cloud VOD provides customers with media upload, media management, media processing and other services. During or at the end of the execution of these services, Tencent Cloud On-Demand also provides various corresponding event notifications to facilitate developers to perceive the service processing status and Do the next business operation. 
//
// 
//
// Developers can achieve this by calling this interface:
//
// - Set the type of callback notification received. Currently, there is [HTTP callback notification](https://www.tencentcloud.com/document/product/266/33948) and [reliable notification based on message queue](https://www.tencentcloud.com/document/product/266/33948) two types. 
//
// - For [HTTP callback notification](https://www.tencentcloud.com/document/product/266/33948), you can set the address of the 3.0 format callback. For the description of 3.0 format callback, see [Historical Format Callback](https://intl.cloud.tencent.com/document/product/266/33796?from_cn_redirect=1). 
//
// - Select settings to receive or ignore notification events for specific event services.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyEventConfig(request *ModifyEventConfigRequest) (response *ModifyEventConfigResponse, err error) {
    return c.ModifyEventConfigWithContext(context.Background(), request)
}

// ModifyEventConfig
// Tencent Cloud VOD provides customers with media upload, media management, media processing and other services. During or at the end of the execution of these services, Tencent Cloud On-Demand also provides various corresponding event notifications to facilitate developers to perceive the service processing status and Do the next business operation. 
//
// 
//
// Developers can achieve this by calling this interface:
//
// - Set the type of callback notification received. Currently, there is [HTTP callback notification](https://www.tencentcloud.com/document/product/266/33948) and [reliable notification based on message queue](https://www.tencentcloud.com/document/product/266/33948) two types. 
//
// - For [HTTP callback notification](https://www.tencentcloud.com/document/product/266/33948), you can set the address of the 3.0 format callback. For the description of 3.0 format callback, see [Historical Format Callback](https://intl.cloud.tencent.com/document/product/266/33796?from_cn_redirect=1). 
//
// - Select settings to receive or ignore notification events for specific event services.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyEventConfigWithContext(ctx context.Context, request *ModifyEventConfigRequest) (response *ModifyEventConfigResponse, err error) {
    if request == nil {
        request = NewModifyEventConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEventConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEventConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHeadTailTemplateRequest() (request *ModifyHeadTailTemplateRequest) {
    request = &ModifyHeadTailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyHeadTailTemplate")
    
    
    return
}

func NewModifyHeadTailTemplateResponse() (response *ModifyHeadTailTemplateResponse) {
    response = &ModifyHeadTailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHeadTailTemplate
// Modify HeadTail Template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_STORAGEREGIONS = "InvalidParameterValue.StorageRegions"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyHeadTailTemplate(request *ModifyHeadTailTemplateRequest) (response *ModifyHeadTailTemplateResponse, err error) {
    return c.ModifyHeadTailTemplateWithContext(context.Background(), request)
}

// ModifyHeadTailTemplate
// Modify HeadTail Template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_STORAGEREGIONS = "InvalidParameterValue.StorageRegions"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyHeadTailTemplateWithContext(ctx context.Context, request *ModifyHeadTailTemplateRequest) (response *ModifyHeadTailTemplateResponse, err error) {
    if request == nil {
        request = NewModifyHeadTailTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHeadTailTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHeadTailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageSpriteTemplateRequest() (request *ModifyImageSpriteTemplateRequest) {
    request = &ModifyImageSpriteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyImageSpriteTemplate")
    
    
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyImageSpriteTemplate(request *ModifyImageSpriteTemplateRequest) (response *ModifyImageSpriteTemplateResponse, err error) {
    return c.ModifyImageSpriteTemplateWithContext(context.Background(), request)
}

// ModifyImageSpriteTemplate
// This API is used to modify a custom image sprite generating template.
//
// error code that may be returned:
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyImageSpriteTemplateWithContext(ctx context.Context, request *ModifyImageSpriteTemplateRequest) (response *ModifyImageSpriteTemplateResponse, err error) {
    if request == nil {
        request = NewModifyImageSpriteTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyImageSpriteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyImageSpriteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyJustInTimeTranscodeTemplateRequest() (request *ModifyJustInTimeTranscodeTemplateRequest) {
    request = &ModifyJustInTimeTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyJustInTimeTranscodeTemplate")
    
    
    return
}

func NewModifyJustInTimeTranscodeTemplateResponse() (response *ModifyJustInTimeTranscodeTemplateResponse) {
    response = &ModifyJustInTimeTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyJustInTimeTranscodeTemplate
// Modify Just In Time Transcode Template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyJustInTimeTranscodeTemplate(request *ModifyJustInTimeTranscodeTemplateRequest) (response *ModifyJustInTimeTranscodeTemplateResponse, err error) {
    return c.ModifyJustInTimeTranscodeTemplateWithContext(context.Background(), request)
}

// ModifyJustInTimeTranscodeTemplate
// Modify Just In Time Transcode Template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyJustInTimeTranscodeTemplateWithContext(ctx context.Context, request *ModifyJustInTimeTranscodeTemplateRequest) (response *ModifyJustInTimeTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyJustInTimeTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyJustInTimeTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyJustInTimeTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaInfoRequest() (request *ModifyMediaInfoRequest) {
    request = &ModifyMediaInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyMediaInfo")
    
    
    return
}

func NewModifyMediaInfoResponse() (response *ModifyMediaInfoResponse) {
    response = &ModifyMediaInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMediaInfo
// This API is used to modify the attributes of a media file, including category, name, description, tag, expiration time, timestamp information, video thumbnail, and subtitle information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR_GETFILEINFOERROR = "InternalError.GetFileInfoError"
//  INTERNALERROR_UPDATEMEDIAERROR = "InternalError.UpdateMediaError"
//  INTERNALERROR_UPLOADCOVERIMAGEERROR = "InternalError.UploadCoverImageError"
//  INVALIDPARAMETERVALUE_ADDKEYFRAMEDESCSANDCLEARKEYFRAMEDESCSCONFLICT = "InvalidParameterValue.AddKeyFrameDescsAndClearKeyFrameDescsConflict"
//  INVALIDPARAMETERVALUE_ADDKEYFRAMEDESCSANDDELETEKEYFRAMEDESCSCONFLICT = "InvalidParameterValue.AddKeyFrameDescsAndDeleteKeyFrameDescsConflict"
//  INVALIDPARAMETERVALUE_ADDTAGSANDCLEARTAGSCONFLICT = "InvalidParameterValue.AddTagsAndClearTagsConflict"
//  INVALIDPARAMETERVALUE_ADDTAGSANDDELETETAGSCONFLICT = "InvalidParameterValue.AddTagsAndDeleteTagsConflict"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_IMAGEDECODEERROR = "InvalidParameterValue.ImageDecodeError"
//  INVALIDPARAMETERVALUE_KEYFRAMEDESCCONTENTTOOLONG = "InvalidParameterValue.KeyFrameDescContentTooLong"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGTOOLONG = "InvalidParameterValue.TagTooLong"
//  LIMITEXCEEDED_KEYFRAMEDESCCOUNTREACHMAX = "LimitExceeded.KeyFrameDescCountReachMax"
//  LIMITEXCEEDED_TAGCOUNTREACHMAX = "LimitExceeded.TagCountReachMax"
//  RESOURCENOTFOUND_FILENOTEXIST = "ResourceNotFound.FileNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyMediaInfo(request *ModifyMediaInfoRequest) (response *ModifyMediaInfoResponse, err error) {
    return c.ModifyMediaInfoWithContext(context.Background(), request)
}

// ModifyMediaInfo
// This API is used to modify the attributes of a media file, including category, name, description, tag, expiration time, timestamp information, video thumbnail, and subtitle information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR_GETFILEINFOERROR = "InternalError.GetFileInfoError"
//  INTERNALERROR_UPDATEMEDIAERROR = "InternalError.UpdateMediaError"
//  INTERNALERROR_UPLOADCOVERIMAGEERROR = "InternalError.UploadCoverImageError"
//  INVALIDPARAMETERVALUE_ADDKEYFRAMEDESCSANDCLEARKEYFRAMEDESCSCONFLICT = "InvalidParameterValue.AddKeyFrameDescsAndClearKeyFrameDescsConflict"
//  INVALIDPARAMETERVALUE_ADDKEYFRAMEDESCSANDDELETEKEYFRAMEDESCSCONFLICT = "InvalidParameterValue.AddKeyFrameDescsAndDeleteKeyFrameDescsConflict"
//  INVALIDPARAMETERVALUE_ADDTAGSANDCLEARTAGSCONFLICT = "InvalidParameterValue.AddTagsAndClearTagsConflict"
//  INVALIDPARAMETERVALUE_ADDTAGSANDDELETETAGSCONFLICT = "InvalidParameterValue.AddTagsAndDeleteTagsConflict"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_IMAGEDECODEERROR = "InvalidParameterValue.ImageDecodeError"
//  INVALIDPARAMETERVALUE_KEYFRAMEDESCCONTENTTOOLONG = "InvalidParameterValue.KeyFrameDescContentTooLong"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TAGTOOLONG = "InvalidParameterValue.TagTooLong"
//  LIMITEXCEEDED_KEYFRAMEDESCCOUNTREACHMAX = "LimitExceeded.KeyFrameDescCountReachMax"
//  LIMITEXCEEDED_TAGCOUNTREACHMAX = "LimitExceeded.TagCountReachMax"
//  RESOURCENOTFOUND_FILENOTEXIST = "ResourceNotFound.FileNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyMediaInfoWithContext(ctx context.Context, request *ModifyMediaInfoRequest) (response *ModifyMediaInfoResponse, err error) {
    if request == nil {
        request = NewModifyMediaInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMediaInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMediaInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaStorageClassRequest() (request *ModifyMediaStorageClassRequest) {
    request = &ModifyMediaStorageClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyMediaStorageClass")
    
    
    return
}

func NewModifyMediaStorageClassResponse() (response *ModifyMediaStorageClassResponse) {
    response = &ModifyMediaStorageClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMediaStorageClass
// This API is used to modify the storage class of media files.
//
// If the current storage class is STANDARD, it can be changed to one of the following classes:
//
// <li>STANDARD_IA</li>
//
// <li>ARCHIVE</li>
//
// <li>DEEP ARCHIVE</li>
//
// If the current storage class is STANDARD_IA, it can be changed to one of the following classes:
//
// <li>STANDARD</li>
//
// <li>ARCHIVE</li>
//
// <li>DEEP ARCHIVE</li>
//
// If the current storage class is ARCHIVE, it can be changed to the following class:
//
// <li>STANDARD</li>
//
// If the current storage class is DEEP ARCHIVE, it can be changed to the following class:
//
// <li>STANDARD</li>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_ORIGINALSTORAGECLASS = "InvalidParameterValue.OriginalStorageClass"
//  INVALIDPARAMETERVALUE_STORAGECLASS = "InvalidParameterValue.StorageClass"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDRESTORETIER = "InvalidParameterValue.UnsupportedRestoreTier"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTORAGECLASS = "InvalidParameterValue.UnsupportedStorageClass"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTRANSITION = "InvalidParameterValue.UnsupportedTransition"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILENOTEXIST = "ResourceNotFound.FileNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMediaStorageClass(request *ModifyMediaStorageClassRequest) (response *ModifyMediaStorageClassResponse, err error) {
    return c.ModifyMediaStorageClassWithContext(context.Background(), request)
}

// ModifyMediaStorageClass
// This API is used to modify the storage class of media files.
//
// If the current storage class is STANDARD, it can be changed to one of the following classes:
//
// <li>STANDARD_IA</li>
//
// <li>ARCHIVE</li>
//
// <li>DEEP ARCHIVE</li>
//
// If the current storage class is STANDARD_IA, it can be changed to one of the following classes:
//
// <li>STANDARD</li>
//
// <li>ARCHIVE</li>
//
// <li>DEEP ARCHIVE</li>
//
// If the current storage class is ARCHIVE, it can be changed to the following class:
//
// <li>STANDARD</li>
//
// If the current storage class is DEEP ARCHIVE, it can be changed to the following class:
//
// <li>STANDARD</li>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_ORIGINALSTORAGECLASS = "InvalidParameterValue.OriginalStorageClass"
//  INVALIDPARAMETERVALUE_STORAGECLASS = "InvalidParameterValue.StorageClass"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDRESTORETIER = "InvalidParameterValue.UnsupportedRestoreTier"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTORAGECLASS = "InvalidParameterValue.UnsupportedStorageClass"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTRANSITION = "InvalidParameterValue.UnsupportedTransition"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILENOTEXIST = "ResourceNotFound.FileNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMediaStorageClassWithContext(ctx context.Context, request *ModifyMediaStorageClassRequest) (response *ModifyMediaStorageClassResponse, err error) {
    if request == nil {
        request = NewModifyMediaStorageClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMediaStorageClass require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMediaStorageClassResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonSampleRequest() (request *ModifyPersonSampleRequest) {
    request = &ModifyPersonSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyPersonSample")
    
    
    return
}

func NewModifyPersonSampleResponse() (response *ModifyPersonSampleResponse) {
    response = &ModifyPersonSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPersonSample
// This API is used to modify sample information according to the sample ID. You can modify the name and description, add, delete, and reset facial features or tags. Leave at least one image after deleting facial features. To leave no image, please use the reset operation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyPersonSample(request *ModifyPersonSampleRequest) (response *ModifyPersonSampleResponse, err error) {
    return c.ModifyPersonSampleWithContext(context.Background(), request)
}

// ModifyPersonSample
// This API is used to modify sample information according to the sample ID. You can modify the name and description, add, delete, and reset facial features or tags. Leave at least one image after deleting facial features. To leave no image, please use the reset operation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"
//  INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyPersonSampleWithContext(ctx context.Context, request *ModifyPersonSampleRequest) (response *ModifyPersonSampleResponse, err error) {
    if request == nil {
        request = NewModifyPersonSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQualityInspectTemplateRequest() (request *ModifyQualityInspectTemplateRequest) {
    request = &ModifyQualityInspectTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyQualityInspectTemplate")
    
    
    return
}

func NewModifyQualityInspectTemplateResponse() (response *ModifyQualityInspectTemplateResponse) {
    response = &ModifyQualityInspectTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQualityInspectTemplate
// Modifies media quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyQualityInspectTemplate(request *ModifyQualityInspectTemplateRequest) (response *ModifyQualityInspectTemplateResponse, err error) {
    return c.ModifyQualityInspectTemplateWithContext(context.Background(), request)
}

// ModifyQualityInspectTemplate
// Modifies media quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyQualityInspectTemplateWithContext(ctx context.Context, request *ModifyQualityInspectTemplateRequest) (response *ModifyQualityInspectTemplateResponse, err error) {
    if request == nil {
        request = NewModifyQualityInspectTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQualityInspectTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQualityInspectTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRebuildMediaTemplateRequest() (request *ModifyRebuildMediaTemplateRequest) {
    request = &ModifyRebuildMediaTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyRebuildMediaTemplate")
    
    
    return
}

func NewModifyRebuildMediaTemplateResponse() (response *ModifyRebuildMediaTemplateResponse) {
    response = &ModifyRebuildMediaTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRebuildMediaTemplate
// Modify Rebuild Media Template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyRebuildMediaTemplate(request *ModifyRebuildMediaTemplateRequest) (response *ModifyRebuildMediaTemplateResponse, err error) {
    return c.ModifyRebuildMediaTemplateWithContext(context.Background(), request)
}

// ModifyRebuildMediaTemplate
// Modify Rebuild Media Template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyRebuildMediaTemplateWithContext(ctx context.Context, request *ModifyRebuildMediaTemplateRequest) (response *ModifyRebuildMediaTemplateResponse, err error) {
    if request == nil {
        request = NewModifyRebuildMediaTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRebuildMediaTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRebuildMediaTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReviewTemplateRequest() (request *ModifyReviewTemplateRequest) {
    request = &ModifyReviewTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyReviewTemplate")
    
    
    return
}

func NewModifyReviewTemplateResponse() (response *ModifyReviewTemplateResponse) {
    response = &ModifyReviewTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyReviewTemplate
// This API is used to modify a custom moderation template.
//
// > The templates can only be used by the APIs [ReviewAudioVideo](https://intl.cloud.tencent.com/document/api/266/80283?from_cn_redirect=1) and [ReviewImage](https://intl.cloud.tencent.com/document/api/266/73217?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELS = "InvalidParameterValue.Labels"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyReviewTemplate(request *ModifyReviewTemplateRequest) (response *ModifyReviewTemplateResponse, err error) {
    return c.ModifyReviewTemplateWithContext(context.Background(), request)
}

// ModifyReviewTemplate
// This API is used to modify a custom moderation template.
//
// > The templates can only be used by the APIs [ReviewAudioVideo](https://intl.cloud.tencent.com/document/api/266/80283?from_cn_redirect=1) and [ReviewImage](https://intl.cloud.tencent.com/document/api/266/73217?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_LABELS = "InvalidParameterValue.Labels"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyReviewTemplateWithContext(ctx context.Context, request *ModifyReviewTemplateRequest) (response *ModifyReviewTemplateResponse, err error) {
    if request == nil {
        request = NewModifyReviewTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoundPlayRequest() (request *ModifyRoundPlayRequest) {
    request = &ModifyRoundPlayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyRoundPlay")
    
    
    return
}

func NewModifyRoundPlayResponse() (response *ModifyRoundPlayResponse) {
    response = &ModifyRoundPlayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRoundPlay
// This API is used to modify a playlist.
//
// The modification will only take effect for new playback requests. For ongoing playback, the old playlist will be playable for seven days after the modification.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_PLAYLIST = "LimitExceeded.PlayList"
//  LIMITEXCEEDED_ROUNDPLAYS = "LimitExceeded.RoundPlays"
//  LIMITEXCEEDED_RUNNINGROUNDPLAYS = "LimitExceeded.RunningRoundPlays"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRoundPlay(request *ModifyRoundPlayRequest) (response *ModifyRoundPlayResponse, err error) {
    return c.ModifyRoundPlayWithContext(context.Background(), request)
}

// ModifyRoundPlay
// This API is used to modify a playlist.
//
// The modification will only take effect for new playback requests. For ongoing playback, the old playlist will be playable for seven days after the modification.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_PLAYLIST = "LimitExceeded.PlayList"
//  LIMITEXCEEDED_ROUNDPLAYS = "LimitExceeded.RoundPlays"
//  LIMITEXCEEDED_RUNNINGROUNDPLAYS = "LimitExceeded.RunningRoundPlays"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRoundPlayWithContext(ctx context.Context, request *ModifyRoundPlayRequest) (response *ModifyRoundPlayResponse, err error) {
    if request == nil {
        request = NewModifyRoundPlayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoundPlay require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoundPlayResponse()
    err = c.Send(request, response)
    return
}

func NewModifySampleSnapshotTemplateRequest() (request *ModifySampleSnapshotTemplateRequest) {
    request = &ModifySampleSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifySampleSnapshotTemplate")
    
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySampleSnapshotTemplate(request *ModifySampleSnapshotTemplateRequest) (response *ModifySampleSnapshotTemplateResponse, err error) {
    return c.ModifySampleSnapshotTemplateWithContext(context.Background(), request)
}

// ModifySampleSnapshotTemplate
// This API is used to modify a custom sampled screencapturing template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"
//  INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySampleSnapshotTemplateWithContext(ctx context.Context, request *ModifySampleSnapshotTemplateRequest) (response *ModifySampleSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewModifySampleSnapshotTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotByTimeOffsetTemplateRequest() (request *ModifySnapshotByTimeOffsetTemplateRequest) {
    request = &ModifySnapshotByTimeOffsetTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifySnapshotByTimeOffsetTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySnapshotByTimeOffsetTemplate(request *ModifySnapshotByTimeOffsetTemplateRequest) (response *ModifySnapshotByTimeOffsetTemplateResponse, err error) {
    return c.ModifySnapshotByTimeOffsetTemplateWithContext(context.Background(), request)
}

// ModifySnapshotByTimeOffsetTemplate
// This API is used to modify a custom time point screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"
//  INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySnapshotByTimeOffsetTemplateWithContext(ctx context.Context, request *ModifySnapshotByTimeOffsetTemplateRequest) (response *ModifySnapshotByTimeOffsetTemplateResponse, err error) {
    if request == nil {
        request = NewModifySnapshotByTimeOffsetTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotByTimeOffsetTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotByTimeOffsetTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubAppIdInfoRequest() (request *ModifySubAppIdInfoRequest) {
    request = &ModifySubAppIdInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifySubAppIdInfo")
    
    
    return
}

func NewModifySubAppIdInfoResponse() (response *ModifySubAppIdInfoResponse) {
    response = &ModifySubAppIdInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySubAppIdInfo
// This API is used to modify subapplication information, but it is not allowed to modify primary application information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubAppIdInfo(request *ModifySubAppIdInfoRequest) (response *ModifySubAppIdInfoResponse, err error) {
    return c.ModifySubAppIdInfoWithContext(context.Background(), request)
}

// ModifySubAppIdInfo
// This API is used to modify subapplication information, but it is not allowed to modify primary application information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubAppIdInfoWithContext(ctx context.Context, request *ModifySubAppIdInfoRequest) (response *ModifySubAppIdInfoResponse, err error) {
    if request == nil {
        request = NewModifySubAppIdInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubAppIdInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubAppIdInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubAppIdStatusRequest() (request *ModifySubAppIdStatusRequest) {
    request = &ModifySubAppIdStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifySubAppIdStatus")
    
    
    return
}

func NewModifySubAppIdStatusResponse() (response *ModifySubAppIdStatusResponse) {
    response = &ModifySubAppIdStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySubAppIdStatus
// This API is used to enable/disable a subapplication. After a subapplication is disabled, its corresponding domain name will be blocked and its access to the console will be restricted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubAppIdStatus(request *ModifySubAppIdStatusRequest) (response *ModifySubAppIdStatusResponse, err error) {
    return c.ModifySubAppIdStatusWithContext(context.Background(), request)
}

// ModifySubAppIdStatus
// This API is used to enable/disable a subapplication. After a subapplication is disabled, its corresponding domain name will be blocked and its access to the console will be restricted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySubAppIdStatusWithContext(ctx context.Context, request *ModifySubAppIdStatusRequest) (response *ModifySubAppIdStatusResponse, err error) {
    if request == nil {
        request = NewModifySubAppIdStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubAppIdStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubAppIdStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySuperPlayerConfigRequest() (request *ModifySuperPlayerConfigRequest) {
    request = &ModifySuperPlayerConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifySuperPlayerConfig")
    
    
    return
}

func NewModifySuperPlayerConfigResponse() (response *ModifySuperPlayerConfigResponse) {
    response = &ModifySuperPlayerConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySuperPlayerConfig
// We have <font color='red'>stopped updating</font> this API. Currently, you no longer need a player configuration to use player signatures. For details, see [Player Signature](https://intl.cloud.tencent.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to modify a player configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySuperPlayerConfig(request *ModifySuperPlayerConfigRequest) (response *ModifySuperPlayerConfigResponse, err error) {
    return c.ModifySuperPlayerConfigWithContext(context.Background(), request)
}

// ModifySuperPlayerConfig
// We have <font color='red'>stopped updating</font> this API. Currently, you no longer need a player configuration to use player signatures. For details, see [Player Signature](https://intl.cloud.tencent.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to modify a player configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySuperPlayerConfigWithContext(ctx context.Context, request *ModifySuperPlayerConfigRequest) (response *ModifySuperPlayerConfigResponse, err error) {
    if request == nil {
        request = NewModifySuperPlayerConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySuperPlayerConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySuperPlayerConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTranscodeTemplateRequest() (request *ModifyTranscodeTemplateRequest) {
    request = &ModifyTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyTranscodeTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FILTRATEAUDIO = "InvalidParameterValue.FiltrateAudio"
//  INVALIDPARAMETERVALUE_FILTRATEVIDEO = "InvalidParameterValue.FiltrateVideo"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VCRF = "InvalidParameterValue.Vcrf"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTranscodeTemplate(request *ModifyTranscodeTemplateRequest) (response *ModifyTranscodeTemplateResponse, err error) {
    return c.ModifyTranscodeTemplateWithContext(context.Background(), request)
}

// ModifyTranscodeTemplate
// This API is used to modify a custom transcoding template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"
//  INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"
//  INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"
//  INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"
//  INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"
//  INVALIDPARAMETERVALUE_FILTRATEAUDIO = "InvalidParameterValue.FiltrateAudio"
//  INVALIDPARAMETERVALUE_FILTRATEVIDEO = "InvalidParameterValue.FiltrateVideo"
//  INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"
//  INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"
//  INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"
//  INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"
//  INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VCRF = "InvalidParameterValue.Vcrf"
//  INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTranscodeTemplateWithContext(ctx context.Context, request *ModifyTranscodeTemplateRequest) (response *ModifyTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVodDomainAccelerateConfigRequest() (request *ModifyVodDomainAccelerateConfigRequest) {
    request = &ModifyVodDomainAccelerateConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyVodDomainAccelerateConfig")
    
    
    return
}

func NewModifyVodDomainAccelerateConfigResponse() (response *ModifyVodDomainAccelerateConfigResponse) {
    response = &ModifyVodDomainAccelerateConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVodDomainAccelerateConfig
// This API is used to modify the acceleration region of a domain name on VOD.
//
// 1. You can modify acceleration regions of only domain names whose status is `Online`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVodDomainAccelerateConfig(request *ModifyVodDomainAccelerateConfigRequest) (response *ModifyVodDomainAccelerateConfigResponse, err error) {
    return c.ModifyVodDomainAccelerateConfigWithContext(context.Background(), request)
}

// ModifyVodDomainAccelerateConfig
// This API is used to modify the acceleration region of a domain name on VOD.
//
// 1. You can modify acceleration regions of only domain names whose status is `Online`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVodDomainAccelerateConfigWithContext(ctx context.Context, request *ModifyVodDomainAccelerateConfigRequest) (response *ModifyVodDomainAccelerateConfigResponse, err error) {
    if request == nil {
        request = NewModifyVodDomainAccelerateConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVodDomainAccelerateConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVodDomainAccelerateConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVodDomainConfigRequest() (request *ModifyVodDomainConfigRequest) {
    request = &ModifyVodDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyVodDomainConfig")
    
    
    return
}

func NewModifyVodDomainConfigResponse() (response *ModifyVodDomainConfigResponse) {
    response = &ModifyVodDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVodDomainConfig
// This API is used to modify domain name settings, such as the hotlink protection configuration.
//
// 1. You can modify settings of only domain names whose status is `Online`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVodDomainConfig(request *ModifyVodDomainConfigRequest) (response *ModifyVodDomainConfigResponse, err error) {
    return c.ModifyVodDomainConfigWithContext(context.Background(), request)
}

// ModifyVodDomainConfig
// This API is used to modify domain name settings, such as the hotlink protection configuration.
//
// 1. You can modify settings of only domain names whose status is `Online`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVodDomainConfigWithContext(ctx context.Context, request *ModifyVodDomainConfigRequest) (response *ModifyVodDomainConfigResponse, err error) {
    if request == nil {
        request = NewModifyVodDomainConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVodDomainConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVodDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWatermarkTemplateRequest() (request *ModifyWatermarkTemplateRequest) {
    request = &ModifyWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyWatermarkTemplate")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyWatermarkTemplate(request *ModifyWatermarkTemplateRequest) (response *ModifyWatermarkTemplateResponse, err error) {
    return c.ModifyWatermarkTemplateWithContext(context.Background(), request)
}

// ModifyWatermarkTemplate
// This API is used to modify a custom watermarking template. The watermark type cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyWatermarkTemplateWithContext(ctx context.Context, request *ModifyWatermarkTemplateRequest) (response *ModifyWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewModifyWatermarkTemplateRequest()
    }
    
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
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyWordSample")
    
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_WORD = "ResourceNotFound.Word"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyWordSample(request *ModifyWordSampleRequest) (response *ModifyWordSampleResponse, err error) {
    return c.ModifyWordSampleWithContext(context.Background(), request)
}

// ModifyWordSample
// This API is used to modify the use case and tag of a keyword. The keyword itself cannot be modified, but you can delete it and create another one if needed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_WORD = "ResourceNotFound.Word"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyWordSampleWithContext(ctx context.Context, request *ModifyWordSampleRequest) (response *ModifyWordSampleResponse, err error) {
    if request == nil {
        request = NewModifyWordSampleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWordSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWordSampleResponse()
    err = c.Send(request, response)
    return
}

func NewParseStreamingManifestRequest() (request *ParseStreamingManifestRequest) {
    request = &ParseStreamingManifestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ParseStreamingManifest")
    
    
    return
}

func NewParseStreamingManifestResponse() (response *ParseStreamingManifestResponse) {
    response = &ParseStreamingManifestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ParseStreamingManifest
// This API is used to parse the index file content and return the list of segment files to be uploaded when an HLS video is uploaded. A segment file path must be a relative path of the current directory or subdirectory instead of a URL or absolute path.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MEDIAMANIFESTCONTENT = "InvalidParameterValue.MediaManifestContent"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ParseStreamingManifest(request *ParseStreamingManifestRequest) (response *ParseStreamingManifestResponse, err error) {
    return c.ParseStreamingManifestWithContext(context.Background(), request)
}

// ParseStreamingManifest
// This API is used to parse the index file content and return the list of segment files to be uploaded when an HLS video is uploaded. A segment file path must be a relative path of the current directory or subdirectory instead of a URL or absolute path.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MEDIAMANIFESTCONTENT = "InvalidParameterValue.MediaManifestContent"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ParseStreamingManifestWithContext(ctx context.Context, request *ParseStreamingManifestRequest) (response *ParseStreamingManifestResponse, err error) {
    if request == nil {
        request = NewParseStreamingManifestRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseStreamingManifest require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseStreamingManifestResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaRequest() (request *ProcessMediaRequest) {
    request = &ProcessMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ProcessMedia")
    
    
    return
}

func NewProcessMediaResponse() (response *ProcessMediaResponse) {
    response = &ProcessMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessMedia
// Initiate processing tasks for media file in VOD, including: 
//
// 1. Video transcoding (with Watermark); 
//
// 2. Video to animated image; 
//
// 3. Screencapturing the video at specified Time point;
//
// 4. Sampled screenshot of the video;
//
// 5. Generated Image sprite template from the video;
//
// 6. Generated a cover image from the video;
//
// 7. Adaptive Bitrate Streaming for the video (with encryption); 
//
// 8. Content Moderation (offensive Information, unsafe Information, inappropriate Information), it is <font color=red>not recommended</font> to use this API, recommend using [Video moderation(ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/50634) or [Image moderation(ReviewImage)](https://www.tencentcloud.com/document/api/266/47138);
//
// 9. Content analysis (tag, category, cover, frame-by-frame tag);
//
// 10. Content recognition (video intro and outro, face, Text, keyword, voice, Key object).
//
// 
//
// If using event notification, the event notification type is [task flow status change](https://www.tencentcloud.com/document/product/266/33953).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_NONEEDTOREDUCEMEDIABITRATE = "FailedOperation.NoNeedToReduceMediaBitrate"
//  FAILEDOPERATION_TASKDUPLICATE = "FailedOperation.TaskDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AIANALYSISTASKDEFINITION = "InvalidParameterValue.AiAnalysisTaskDefinition"
//  INVALIDPARAMETERVALUE_AICONTENTREVIEWTASKDEFINITION = "InvalidParameterValue.AiContentReviewTaskDefinition"
//  INVALIDPARAMETERVALUE_AIRECOGNITIONTASKDEFINITION = "InvalidParameterValue.AiRecognitionTaskDefinition"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessMedia(request *ProcessMediaRequest) (response *ProcessMediaResponse, err error) {
    return c.ProcessMediaWithContext(context.Background(), request)
}

// ProcessMedia
// Initiate processing tasks for media file in VOD, including: 
//
// 1. Video transcoding (with Watermark); 
//
// 2. Video to animated image; 
//
// 3. Screencapturing the video at specified Time point;
//
// 4. Sampled screenshot of the video;
//
// 5. Generated Image sprite template from the video;
//
// 6. Generated a cover image from the video;
//
// 7. Adaptive Bitrate Streaming for the video (with encryption); 
//
// 8. Content Moderation (offensive Information, unsafe Information, inappropriate Information), it is <font color=red>not recommended</font> to use this API, recommend using [Video moderation(ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/50634) or [Image moderation(ReviewImage)](https://www.tencentcloud.com/document/api/266/47138);
//
// 9. Content analysis (tag, category, cover, frame-by-frame tag);
//
// 10. Content recognition (video intro and outro, face, Text, keyword, voice, Key object).
//
// 
//
// If using event notification, the event notification type is [task flow status change](https://www.tencentcloud.com/document/product/266/33953).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_NONEEDTOREDUCEMEDIABITRATE = "FailedOperation.NoNeedToReduceMediaBitrate"
//  FAILEDOPERATION_TASKDUPLICATE = "FailedOperation.TaskDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AIANALYSISTASKDEFINITION = "InvalidParameterValue.AiAnalysisTaskDefinition"
//  INVALIDPARAMETERVALUE_AICONTENTREVIEWTASKDEFINITION = "InvalidParameterValue.AiContentReviewTaskDefinition"
//  INVALIDPARAMETERVALUE_AIRECOGNITIONTASKDEFINITION = "InvalidParameterValue.AiRecognitionTaskDefinition"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessMediaWithContext(ctx context.Context, request *ProcessMediaRequest) (response *ProcessMediaResponse, err error) {
    if request == nil {
        request = NewProcessMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessMediaResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaByProcedureRequest() (request *ProcessMediaByProcedureRequest) {
    request = &ProcessMediaByProcedureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ProcessMediaByProcedure")
    
    
    return
}

func NewProcessMediaByProcedureResponse() (response *ProcessMediaByProcedureResponse) {
    response = &ProcessMediaByProcedureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessMediaByProcedure
// This API is used to start a task flow on a video.
//
// There are two ways to create a task flow template:
//
// 1. Create and modify a task flow template in the console;
//
// 2. Create a task flow template using the `CreateProcedureTemplate` API.
//
// 
//
// If event notifications are used, the event type for moderation tasks is [ReviewAudioVideoComplete](https://intl.cloud.tencent.com/document/product/266/81258?from_cn_redirect=1), and that for other tasks is [ProcedureStateChanged](https://intl.cloud.tencent.com/document/product/266/9636?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_PROCEDURENAME = "InvalidParameterValue.ProcedureName"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  INVALIDPARAMETERVALUE_UNIQUEIDENTIFIER = "InvalidParameterValue.UniqueIdentifier"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessMediaByProcedure(request *ProcessMediaByProcedureRequest) (response *ProcessMediaByProcedureResponse, err error) {
    return c.ProcessMediaByProcedureWithContext(context.Background(), request)
}

// ProcessMediaByProcedure
// This API is used to start a task flow on a video.
//
// There are two ways to create a task flow template:
//
// 1. Create and modify a task flow template in the console;
//
// 2. Create a task flow template using the `CreateProcedureTemplate` API.
//
// 
//
// If event notifications are used, the event type for moderation tasks is [ReviewAudioVideoComplete](https://intl.cloud.tencent.com/document/product/266/81258?from_cn_redirect=1), and that for other tasks is [ProcedureStateChanged](https://intl.cloud.tencent.com/document/product/266/9636?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_PROCEDURENAME = "InvalidParameterValue.ProcedureName"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  INVALIDPARAMETERVALUE_UNIQUEIDENTIFIER = "InvalidParameterValue.UniqueIdentifier"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessMediaByProcedureWithContext(ctx context.Context, request *ProcessMediaByProcedureRequest) (response *ProcessMediaByProcedureResponse, err error) {
    if request == nil {
        request = NewProcessMediaByProcedureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMediaByProcedure require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessMediaByProcedureResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaByUrlRequest() (request *ProcessMediaByUrlRequest) {
    request = &ProcessMediaByUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ProcessMediaByUrl")
    
    
    return
}

func NewProcessMediaByUrlResponse() (response *ProcessMediaByUrlResponse) {
    response = &ProcessMediaByUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessMediaByUrl
// This API is <font color='red'>disused</font>, please use [ProcessMedia](https://intl.cloud.tencent.com/document/product/862/37578?from_cn_redirect=1) API of MPS, with the input parameter `InputInfo.UrlInputInfo.Url` set to a video URL.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_TASKDUPLICATE = "FailedOperation.TaskDuplicate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETFILEINFOERROR = "InternalError.GetFileInfoError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_AIANALYSISTASKDEFINITION = "InvalidParameterValue.AiAnalysisTaskDefinition"
//  INVALIDPARAMETERVALUE_AICONTENTREVIEWTASKDEFINITION = "InvalidParameterValue.AiContentReviewTaskDefinition"
//  INVALIDPARAMETERVALUE_AIRECOGNITIONTASKDEFINITION = "InvalidParameterValue.AiRecognitionTaskDefinition"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessMediaByUrl(request *ProcessMediaByUrlRequest) (response *ProcessMediaByUrlResponse, err error) {
    return c.ProcessMediaByUrlWithContext(context.Background(), request)
}

// ProcessMediaByUrl
// This API is <font color='red'>disused</font>, please use [ProcessMedia](https://intl.cloud.tencent.com/document/product/862/37578?from_cn_redirect=1) API of MPS, with the input parameter `InputInfo.UrlInputInfo.Url` set to a video URL.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_TASKDUPLICATE = "FailedOperation.TaskDuplicate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETFILEINFOERROR = "InternalError.GetFileInfoError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_AIANALYSISTASKDEFINITION = "InvalidParameterValue.AiAnalysisTaskDefinition"
//  INVALIDPARAMETERVALUE_AICONTENTREVIEWTASKDEFINITION = "InvalidParameterValue.AiContentReviewTaskDefinition"
//  INVALIDPARAMETERVALUE_AIRECOGNITIONTASKDEFINITION = "InvalidParameterValue.AiRecognitionTaskDefinition"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessMediaByUrlWithContext(ctx context.Context, request *ProcessMediaByUrlRequest) (response *ProcessMediaByUrlResponse, err error) {
    if request == nil {
        request = NewProcessMediaByUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMediaByUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessMediaByUrlResponse()
    err = c.Send(request, response)
    return
}

func NewPullEventsRequest() (request *PullEventsRequest) {
    request = &PullEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "PullEvents")
    
    
    return
}

func NewPullEventsResponse() (response *PullEventsResponse) {
    response = &PullEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullEvents
// * This API is used to get event notifications from the business server through [reliable callback](https://intl.cloud.tencent.com/document/product/266/33948).
//
// * The API gets event data through long polling. That is, if there is any unconsumed event on the server, the event notification will be returned to the requester immediately. If there is no unconsumed event on the server, the request will be suspended in the backend until a new event is generated.
//
// * The request can be suspended for up to 5 seconds. It's recommended to set the request timeout period to 10 seconds.
//
// * Event notifications not pulled will be retained for up to 4 days and may be cleared after this period.
//
// * After the API returns an event, the caller must call the [ConfirmEvents](https://intl.cloud.tencent.com/document/product/266/34184) API within <font color="red">30 seconds</font> to confirm that the event notification has been processed. Otherwise, the event notification will be pulled again after <font color="red">30 seconds</font>.
//
// * This API can get up to 16 event notifications at a time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PullEvents(request *PullEventsRequest) (response *PullEventsResponse, err error) {
    return c.PullEventsWithContext(context.Background(), request)
}

// PullEvents
// * This API is used to get event notifications from the business server through [reliable callback](https://intl.cloud.tencent.com/document/product/266/33948).
//
// * The API gets event data through long polling. That is, if there is any unconsumed event on the server, the event notification will be returned to the requester immediately. If there is no unconsumed event on the server, the request will be suspended in the backend until a new event is generated.
//
// * The request can be suspended for up to 5 seconds. It's recommended to set the request timeout period to 10 seconds.
//
// * Event notifications not pulled will be retained for up to 4 days and may be cleared after this period.
//
// * After the API returns an event, the caller must call the [ConfirmEvents](https://intl.cloud.tencent.com/document/product/266/34184) API within <font color="red">30 seconds</font> to confirm that the event notification has been processed. Otherwise, the event notification will be pulled again after <font color="red">30 seconds</font>.
//
// * This API can get up to 16 event notifications at a time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PullEventsWithContext(ctx context.Context, request *PullEventsRequest) (response *PullEventsResponse, err error) {
    if request == nil {
        request = NewPullEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullEventsResponse()
    err = c.Send(request, response)
    return
}

func NewPullUploadRequest() (request *PullUploadRequest) {
    request = &PullUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "PullUpload")
    
    
    return
}

func NewPullUploadResponse() (response *PullUploadResponse) {
    response = &PullUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullUpload
// This API is used to pull a video on the internet to the VOD platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COVERTYPE = "FailedOperation.CoverType"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_MEDIATYPE = "FailedOperation.MediaType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXPIRETIME = "InvalidParameter.ExpireTime"
//  INVALIDPARAMETER_STORAGEREGION = "InvalidParameter.StorageRegion"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COVERURL = "InvalidParameterValue.CoverUrl"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_MEDIAURL = "InvalidParameterValue.MediaUrl"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_STORAGEREGION = "InvalidParameterValue.StorageRegion"
//  RESOURCENOTFOUND_COVERURL = "ResourceNotFound.CoverUrl"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PullUpload(request *PullUploadRequest) (response *PullUploadResponse, err error) {
    return c.PullUploadWithContext(context.Background(), request)
}

// PullUpload
// This API is used to pull a video on the internet to the VOD platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COVERTYPE = "FailedOperation.CoverType"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_MEDIATYPE = "FailedOperation.MediaType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXPIRETIME = "InvalidParameter.ExpireTime"
//  INVALIDPARAMETER_STORAGEREGION = "InvalidParameter.StorageRegion"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COVERURL = "InvalidParameterValue.CoverUrl"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_MEDIAURL = "InvalidParameterValue.MediaUrl"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_STORAGEREGION = "InvalidParameterValue.StorageRegion"
//  RESOURCENOTFOUND_COVERURL = "ResourceNotFound.CoverUrl"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PullUploadWithContext(ctx context.Context, request *PullUploadRequest) (response *PullUploadResponse, err error) {
    if request == nil {
        request = NewPullUploadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullUpload require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullUploadResponse()
    err = c.Send(request, response)
    return
}

func NewPushUrlCacheRequest() (request *PushUrlCacheRequest) {
    request = &PushUrlCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "PushUrlCache")
    
    
    return
}

func NewPushUrlCacheResponse() (response *PushUrlCacheResponse) {
    response = &PushUrlCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PushUrlCache
// 1. This API is used to prefetch a list of specified URLs.
//
// 2. The URL domain names must have already been registered with VOD.
//
// 3. Up to 20 URLs can be specified in one request.
//
// 4. By default, the maximum number of URLs that can be refreshed per day is 10,000.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PushUrlCache(request *PushUrlCacheRequest) (response *PushUrlCacheResponse, err error) {
    return c.PushUrlCacheWithContext(context.Background(), request)
}

// PushUrlCache
// 1. This API is used to prefetch a list of specified URLs.
//
// 2. The URL domain names must have already been registered with VOD.
//
// 3. Up to 20 URLs can be specified in one request.
//
// 4. By default, the maximum number of URLs that can be refreshed per day is 10,000.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PushUrlCacheWithContext(ctx context.Context, request *PushUrlCacheRequest) (response *PushUrlCacheResponse, err error) {
    if request == nil {
        request = NewPushUrlCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PushUrlCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewPushUrlCacheResponse()
    err = c.Send(request, response)
    return
}

func NewRebuildMediaRequest() (request *RebuildMediaRequest) {
    request = &RebuildMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "RebuildMedia")
    
    
    return
}

func NewRebuildMediaResponse() (response *RebuildMediaResponse) {
    response = &RebuildMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RebuildMedia
// Initiate rebuild media
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) RebuildMedia(request *RebuildMediaRequest) (response *RebuildMediaResponse, err error) {
    return c.RebuildMediaWithContext(context.Background(), request)
}

// RebuildMedia
// Initiate rebuild media
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) RebuildMediaWithContext(ctx context.Context, request *RebuildMediaRequest) (response *RebuildMediaResponse, err error) {
    if request == nil {
        request = NewRebuildMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebuildMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebuildMediaResponse()
    err = c.Send(request, response)
    return
}

func NewRebuildMediaByTemplateRequest() (request *RebuildMediaByTemplateRequest) {
    request = &RebuildMediaByTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "RebuildMediaByTemplate")
    
    
    return
}

func NewRebuildMediaByTemplateResponse() (response *RebuildMediaByTemplateResponse) {
    response = &RebuildMediaByTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RebuildMediaByTemplate
// Rebuild media by template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) RebuildMediaByTemplate(request *RebuildMediaByTemplateRequest) (response *RebuildMediaByTemplateResponse, err error) {
    return c.RebuildMediaByTemplateWithContext(context.Background(), request)
}

// RebuildMediaByTemplate
// Rebuild media by template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) RebuildMediaByTemplateWithContext(ctx context.Context, request *RebuildMediaByTemplateRequest) (response *RebuildMediaByTemplateResponse, err error) {
    if request == nil {
        request = NewRebuildMediaByTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebuildMediaByTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebuildMediaByTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshUrlCacheRequest() (request *RefreshUrlCacheRequest) {
    request = &RefreshUrlCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "RefreshUrlCache")
    
    
    return
}

func NewRefreshUrlCacheResponse() (response *RefreshUrlCacheResponse) {
    response = &RefreshUrlCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RefreshUrlCache
// 1. This API is used to purge URLs.
//
// 2. The URL domain names must have already been registered with VOD.
//
// 3. Up to 20 URLs can be specified in one request.
//
// 4. By default, the maximum number of URLs allowed for purge per day is 100,000.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RefreshUrlCache(request *RefreshUrlCacheRequest) (response *RefreshUrlCacheResponse, err error) {
    return c.RefreshUrlCacheWithContext(context.Background(), request)
}

// RefreshUrlCache
// 1. This API is used to purge URLs.
//
// 2. The URL domain names must have already been registered with VOD.
//
// 3. Up to 20 URLs can be specified in one request.
//
// 4. By default, the maximum number of URLs allowed for purge per day is 100,000.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RefreshUrlCacheWithContext(ctx context.Context, request *RefreshUrlCacheRequest) (response *RefreshUrlCacheResponse, err error) {
    if request == nil {
        request = NewRefreshUrlCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefreshUrlCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefreshUrlCacheResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveWatermarkRequest() (request *RemoveWatermarkRequest) {
    request = &RemoveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "RemoveWatermark")
    
    
    return
}

func NewRemoveWatermarkResponse() (response *RemoveWatermarkResponse) {
    response = &RemoveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveWatermark
// This API is used to remove watermarks from a video.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_TASKDUPLICATE = "FailedOperation.TaskDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RemoveWatermark(request *RemoveWatermarkRequest) (response *RemoveWatermarkResponse, err error) {
    return c.RemoveWatermarkWithContext(context.Background(), request)
}

// RemoveWatermark
// This API is used to remove watermarks from a video.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_TASKDUPLICATE = "FailedOperation.TaskDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RemoveWatermarkWithContext(ctx context.Context, request *RemoveWatermarkRequest) (response *RemoveWatermarkResponse, err error) {
    if request == nil {
        request = NewRemoveWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewResetProcedureTemplateRequest() (request *ResetProcedureTemplateRequest) {
    request = &ResetProcedureTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ResetProcedureTemplate")
    
    
    return
}

func NewResetProcedureTemplateResponse() (response *ResetProcedureTemplateResponse) {
    response = &ResetProcedureTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetProcedureTemplate
// This API is used to modify a custom task flow template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROCEDURENAMENOTEXIST = "InvalidParameter.ProcedureNameNotExist"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ResetProcedureTemplate(request *ResetProcedureTemplateRequest) (response *ResetProcedureTemplateResponse, err error) {
    return c.ResetProcedureTemplateWithContext(context.Background(), request)
}

// ResetProcedureTemplate
// This API is used to modify a custom task flow template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROCEDURENAMENOTEXIST = "InvalidParameter.ProcedureNameNotExist"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ResetProcedureTemplateWithContext(ctx context.Context, request *ResetProcedureTemplateRequest) (response *ResetProcedureTemplateResponse, err error) {
    if request == nil {
        request = NewResetProcedureTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetProcedureTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetProcedureTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreMediaRequest() (request *RestoreMediaRequest) {
    request = &RestoreMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "RestoreMedia")
    
    
    return
}

func NewRestoreMediaResponse() (response *RestoreMediaResponse) {
    response = &RestoreMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestoreMedia
// This API is used to restore files from ARCHIVE or DEEP ARCHIVE. Files stored in ARCHIVE or DEEP ARCHIVE must be restored before they can be accessed. Restored files are available for a limited period of time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_NOTRESTORABLE = "InvalidParameterValue.NotRestorable"
//  INVALIDPARAMETERVALUE_ORIGINALSTORAGECLASS = "InvalidParameterValue.OriginalStorageClass"
//  INVALIDPARAMETERVALUE_RESTOREDAY = "InvalidParameterValue.RestoreDay"
//  INVALIDPARAMETERVALUE_RESTORETIER = "InvalidParameterValue.RestoreTier"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDRESTORETIER = "InvalidParameterValue.UnsupportedRestoreTier"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILENOTEXIST = "ResourceNotFound.FileNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestoreMedia(request *RestoreMediaRequest) (response *RestoreMediaResponse, err error) {
    return c.RestoreMediaWithContext(context.Background(), request)
}

// RestoreMedia
// This API is used to restore files from ARCHIVE or DEEP ARCHIVE. Files stored in ARCHIVE or DEEP ARCHIVE must be restored before they can be accessed. Restored files are available for a limited period of time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_NOTRESTORABLE = "InvalidParameterValue.NotRestorable"
//  INVALIDPARAMETERVALUE_ORIGINALSTORAGECLASS = "InvalidParameterValue.OriginalStorageClass"
//  INVALIDPARAMETERVALUE_RESTOREDAY = "InvalidParameterValue.RestoreDay"
//  INVALIDPARAMETERVALUE_RESTORETIER = "InvalidParameterValue.RestoreTier"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDRESTORETIER = "InvalidParameterValue.UnsupportedRestoreTier"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILENOTEXIST = "ResourceNotFound.FileNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestoreMediaWithContext(ctx context.Context, request *RestoreMediaRequest) (response *RestoreMediaResponse, err error) {
    if request == nil {
        request = NewRestoreMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestoreMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestoreMediaResponse()
    err = c.Send(request, response)
    return
}

func NewReviewAudioVideoRequest() (request *ReviewAudioVideoRequest) {
    request = &ReviewAudioVideoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ReviewAudioVideo")
    
    
    return
}

func NewReviewAudioVideoResponse() (response *ReviewAudioVideoResponse) {
    response = &ReviewAudioVideoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReviewAudioVideo
// This API is used to start a moderation task on a file stored in VOD to detect non-compliant content in images, text, speech, and voice.
//
// 
//
// If event notifications are used, the event type is [ReviewAudioVideoComplete](https://intl.cloud.tencent.com/document/product/266/81258?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReviewAudioVideo(request *ReviewAudioVideoRequest) (response *ReviewAudioVideoResponse, err error) {
    return c.ReviewAudioVideoWithContext(context.Background(), request)
}

// ReviewAudioVideo
// This API is used to start a moderation task on a file stored in VOD to detect non-compliant content in images, text, speech, and voice.
//
// 
//
// If event notifications are used, the event type is [ReviewAudioVideoComplete](https://intl.cloud.tencent.com/document/product/266/81258?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReviewAudioVideoWithContext(ctx context.Context, request *ReviewAudioVideoRequest) (response *ReviewAudioVideoResponse, err error) {
    if request == nil {
        request = NewReviewAudioVideoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReviewAudioVideo require credential")
    }

    request.SetContext(ctx)
    
    response = NewReviewAudioVideoResponse()
    err = c.Send(request, response)
    return
}

func NewReviewImageRequest() (request *ReviewImageRequest) {
    request = &ReviewImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ReviewImage")
    
    
    return
}

func NewReviewImageResponse() (response *ReviewImageResponse) {
    response = &ReviewImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReviewImage
// This API is used to moderate an image stored in VOD (detect pornographic and terrorist content).><li>The image file must be smaller than 5 MB.</li> ><li>To ensure the accuracy of moderation results, the image resolution must be higher than 256 x 256 px.</li> ><li>The format must be PNG, JPG, JPEG, BMP, GIF, or WEBP.</li>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_MEDIATYPE = "FailedOperation.MediaType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReviewImage(request *ReviewImageRequest) (response *ReviewImageResponse, err error) {
    return c.ReviewImageWithContext(context.Background(), request)
}

// ReviewImage
// This API is used to moderate an image stored in VOD (detect pornographic and terrorist content).><li>The image file must be smaller than 5 MB.</li> ><li>To ensure the accuracy of moderation results, the image resolution must be higher than 256 x 256 px.</li> ><li>The format must be PNG, JPG, JPEG, BMP, GIF, or WEBP.</li>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_MEDIATYPE = "FailedOperation.MediaType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReviewImageWithContext(ctx context.Context, request *ReviewImageRequest) (response *ReviewImageResponse, err error) {
    if request == nil {
        request = NewReviewImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReviewImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewReviewImageResponse()
    err = c.Send(request, response)
    return
}

func NewSearchMediaRequest() (request *SearchMediaRequest) {
    request = &SearchMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "SearchMedia")
    
    
    return
}

func NewSearchMediaResponse() (response *SearchMediaResponse) {
    response = &SearchMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchMedia
// Search for media information, supporting multiple condition filtering, as well as sorting and filtering of returned results. This includes:
//
// - Specify a collection of file IDs (FileIds) to return media with any ID in the collection.
//
// - Fuzzy search based on multiple media file names (Names) or description information (Descriptions).
//
// - Search based on multiple file name prefixes (NamePrefixes).
//
// - Specify a collection of categories (ClassIds, see input parameters) to return media that meets any category in the collection. For example, media categories include movies, TV shows, variety shows, etc., and there are subcategories such as historical films, action films, and romance films under the movie category. If ClassIds specifies movies and TV shows, all subcategories under movies and TV shows will be returned; if ClassIds specifies historical films and action films, only media under these two subcategories will be returned.
//
// - Specify a collection of tags (Tags, see input parameters) to return media that meets any tag in the collection. For example, media tags include ACG, palace fighting, and ghost animal. If Tags specifies ACG and ghost animal, media that meets any of these two tags will be retrieved.
//
// - Specify a collection of file types (Categories, see input parameters) to return media that meets any type in the collection. For example, file types include Video (video), Audio (audio), and Image (image). If Categories specifies Video and Audio, media that meets these types will be retrieved.
//
// - Specify a collection of sources (SourceTypes, see input parameters) to return media that meets any source in the collection. For example, media sources include Record (live recording) and Upload (upload). If SourceTypes specifies Record and Upload, media that meets these sources will be retrieved.
//
// - Specify a collection of file container formats (MediaTypes, see input parameters) to return media that meets any container format in the collection. For example, container formats include MP4, AVI, MP3, etc. If MediaTypes specifies MP4 and MP3, media that meets these container formats will be retrieved.
//
// - Specify a collection of file statuses (Status, see input parameters) to return media that meets any status in the collection. For example, file statuses include Normal (normal), SystemForbidden (platform ban), and Forbidden (active ban). If Status specifies Normal and Forbidden, media that meets these statuses will be retrieved.
//
// - Specify a collection of file moderation results (ReviewResults, see input parameters) to return media that meets any status in the collection. For example, file moderation results include pass (passed) and block (not compliant). If ReviewResults specifies pass and block, media that meets these moderation results will be retrieved.
//
// - Filter live recorded media by specifying a collection of live streaming codes (StreamIds, see input parameters).
//
// - Filter media by specifying the creation time range of the media.
//
// - Filter media by specifying a collection of TRTC application IDs.
//
// - Filter media by specifying a collection of TRTC room IDs.
//
// 
//
// - The above parameters can be combined in any way to search. For example, filter media created between 12:00:00 on December 1, 2018, and 12:00:00 on December 8, 2018, categorized as movies or TV shows, and tagged with palace fighting and suspense. Note that the search logic for elements of any parameter that supports array input is 'or'. The logical relationship between all parameters is 'and'
//
// 
//
// - Allow to control the type of media information returned through Filters (default to return all information). Optional inputs include:
//
//     1. Basic information (basicInfo): including media name, category, playback URL, cover image, etc.
//
//     2. Metadata (metaData): including size, duration, video stream information, audio stream information, etc.
//
//     3. Transcoding result information (transcodeInfo): including the media addresses, video stream parameters, audio stream parameters, etc., generated by transcoding the media into various specifications.
//
//     4.  Animated image result information (animatedGraphicsInfo): information on the animated image (such as gif) generated after converting the video.
//
//     Sampled screenshot information (sampleSnapshotInfo): screenshot information after sampling the video.
//
//     6. Image sprite information (imageSpriteInfo): image sprite information after generating the sprite from the video.
//
//     7. Specified time point screenshot information (snapshotByTimeOffsetInfo): screenshot information after capturing the video at specified time points.
//
//     8. Video timestamp info (keyFrameDescInfo): timestamp information set for the video.
//
//     9. Adaptive Bitrate Streaming information (adaptiveDynamicStreamingInfo): including specifications, encryption types, muxing formats, and other relevant information.
//
// 
//
// - Allow sorting the results by creation time and returning them in pages. Pagination is controlled by Offset and Limit (see input parameters).
//
// 
//
// <div id="maxResultsDesc">API result count limitation:</div>
//
// - <b><a href="#p_offset">Offset</a> and <a href="#p_limit">Limit</a> both affect the number of results returned in a single page query. Please pay special attention: when both of these values are missing, this API will return a maximum of 10 query results.</b>
//
// - <b>Supports up to 5,000 search results, and queries beyond this limit are not supported. If the search result volume is too large, it is recommended to use more refined filtering conditions to reduce search results.</b>
//
// 
//
// <br>Not recommended for conditional filtering:
//
// - (Not recommended: use Names, NamePrefixes, or Descriptions instead) Fuzzy search for media file names or description information with a single text (Text).
//
// - (Not recommended: use SourceTypes instead) Search for media files with a single source (SourceType).
//
// - (Not recommended: use StreamIds instead) Search for media files with a single live streaming code (StreamId).
//
// - (Not recommended: use CreateTime instead) Search for media files with a single start creation time (StartTime).
//
// - (Not recommended: use CreateTime instead) Search for media files with a single end creation time (EndTime).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETMEDIALISTERROR = "InternalError.GetMediaListError"
//  INVALIDPARAMETERVALUE_CATEGORIES = "InvalidParameterValue.Categories"
//  INVALIDPARAMETERVALUE_CLASSIDS = "InvalidParameterValue.ClassIds"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_NAMEPREFIXES = "InvalidParameterValue.NamePrefixes"
//  INVALIDPARAMETERVALUE_NAMES = "InvalidParameterValue.Names"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_SORT = "InvalidParameterValue.Sort"
//  INVALIDPARAMETERVALUE_SOURCETYPE = "InvalidParameterValue.SourceType"
//  INVALIDPARAMETERVALUE_SOURCETYPES = "InvalidParameterValue.SourceTypes"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  INVALIDPARAMETERVALUE_STORAGEREGIONS = "InvalidParameterValue.StorageRegions"
//  INVALIDPARAMETERVALUE_STREAMIDS = "InvalidParameterValue.StreamIds"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  INVALIDPARAMETERVALUE_TAGS = "InvalidParameterValue.Tags"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  INVALIDPARAMETERVALUE_TYPES = "InvalidParameterValue.Types"
//  INVALIDPARAMETERVALUE_VIDS = "InvalidParameterValue.Vids"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchMedia(request *SearchMediaRequest) (response *SearchMediaResponse, err error) {
    return c.SearchMediaWithContext(context.Background(), request)
}

// SearchMedia
// Search for media information, supporting multiple condition filtering, as well as sorting and filtering of returned results. This includes:
//
// - Specify a collection of file IDs (FileIds) to return media with any ID in the collection.
//
// - Fuzzy search based on multiple media file names (Names) or description information (Descriptions).
//
// - Search based on multiple file name prefixes (NamePrefixes).
//
// - Specify a collection of categories (ClassIds, see input parameters) to return media that meets any category in the collection. For example, media categories include movies, TV shows, variety shows, etc., and there are subcategories such as historical films, action films, and romance films under the movie category. If ClassIds specifies movies and TV shows, all subcategories under movies and TV shows will be returned; if ClassIds specifies historical films and action films, only media under these two subcategories will be returned.
//
// - Specify a collection of tags (Tags, see input parameters) to return media that meets any tag in the collection. For example, media tags include ACG, palace fighting, and ghost animal. If Tags specifies ACG and ghost animal, media that meets any of these two tags will be retrieved.
//
// - Specify a collection of file types (Categories, see input parameters) to return media that meets any type in the collection. For example, file types include Video (video), Audio (audio), and Image (image). If Categories specifies Video and Audio, media that meets these types will be retrieved.
//
// - Specify a collection of sources (SourceTypes, see input parameters) to return media that meets any source in the collection. For example, media sources include Record (live recording) and Upload (upload). If SourceTypes specifies Record and Upload, media that meets these sources will be retrieved.
//
// - Specify a collection of file container formats (MediaTypes, see input parameters) to return media that meets any container format in the collection. For example, container formats include MP4, AVI, MP3, etc. If MediaTypes specifies MP4 and MP3, media that meets these container formats will be retrieved.
//
// - Specify a collection of file statuses (Status, see input parameters) to return media that meets any status in the collection. For example, file statuses include Normal (normal), SystemForbidden (platform ban), and Forbidden (active ban). If Status specifies Normal and Forbidden, media that meets these statuses will be retrieved.
//
// - Specify a collection of file moderation results (ReviewResults, see input parameters) to return media that meets any status in the collection. For example, file moderation results include pass (passed) and block (not compliant). If ReviewResults specifies pass and block, media that meets these moderation results will be retrieved.
//
// - Filter live recorded media by specifying a collection of live streaming codes (StreamIds, see input parameters).
//
// - Filter media by specifying the creation time range of the media.
//
// - Filter media by specifying a collection of TRTC application IDs.
//
// - Filter media by specifying a collection of TRTC room IDs.
//
// 
//
// - The above parameters can be combined in any way to search. For example, filter media created between 12:00:00 on December 1, 2018, and 12:00:00 on December 8, 2018, categorized as movies or TV shows, and tagged with palace fighting and suspense. Note that the search logic for elements of any parameter that supports array input is 'or'. The logical relationship between all parameters is 'and'
//
// 
//
// - Allow to control the type of media information returned through Filters (default to return all information). Optional inputs include:
//
//     1. Basic information (basicInfo): including media name, category, playback URL, cover image, etc.
//
//     2. Metadata (metaData): including size, duration, video stream information, audio stream information, etc.
//
//     3. Transcoding result information (transcodeInfo): including the media addresses, video stream parameters, audio stream parameters, etc., generated by transcoding the media into various specifications.
//
//     4.  Animated image result information (animatedGraphicsInfo): information on the animated image (such as gif) generated after converting the video.
//
//     Sampled screenshot information (sampleSnapshotInfo): screenshot information after sampling the video.
//
//     6. Image sprite information (imageSpriteInfo): image sprite information after generating the sprite from the video.
//
//     7. Specified time point screenshot information (snapshotByTimeOffsetInfo): screenshot information after capturing the video at specified time points.
//
//     8. Video timestamp info (keyFrameDescInfo): timestamp information set for the video.
//
//     9. Adaptive Bitrate Streaming information (adaptiveDynamicStreamingInfo): including specifications, encryption types, muxing formats, and other relevant information.
//
// 
//
// - Allow sorting the results by creation time and returning them in pages. Pagination is controlled by Offset and Limit (see input parameters).
//
// 
//
// <div id="maxResultsDesc">API result count limitation:</div>
//
// - <b><a href="#p_offset">Offset</a> and <a href="#p_limit">Limit</a> both affect the number of results returned in a single page query. Please pay special attention: when both of these values are missing, this API will return a maximum of 10 query results.</b>
//
// - <b>Supports up to 5,000 search results, and queries beyond this limit are not supported. If the search result volume is too large, it is recommended to use more refined filtering conditions to reduce search results.</b>
//
// 
//
// <br>Not recommended for conditional filtering:
//
// - (Not recommended: use Names, NamePrefixes, or Descriptions instead) Fuzzy search for media file names or description information with a single text (Text).
//
// - (Not recommended: use SourceTypes instead) Search for media files with a single source (SourceType).
//
// - (Not recommended: use StreamIds instead) Search for media files with a single live streaming code (StreamId).
//
// - (Not recommended: use CreateTime instead) Search for media files with a single start creation time (StartTime).
//
// - (Not recommended: use CreateTime instead) Search for media files with a single end creation time (EndTime).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETMEDIALISTERROR = "InternalError.GetMediaListError"
//  INVALIDPARAMETERVALUE_CATEGORIES = "InvalidParameterValue.Categories"
//  INVALIDPARAMETERVALUE_CLASSIDS = "InvalidParameterValue.ClassIds"
//  INVALIDPARAMETERVALUE_ENDTIME = "InvalidParameterValue.EndTime"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_NAMEPREFIXES = "InvalidParameterValue.NamePrefixes"
//  INVALIDPARAMETERVALUE_NAMES = "InvalidParameterValue.Names"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_SORT = "InvalidParameterValue.Sort"
//  INVALIDPARAMETERVALUE_SOURCETYPE = "InvalidParameterValue.SourceType"
//  INVALIDPARAMETERVALUE_SOURCETYPES = "InvalidParameterValue.SourceTypes"
//  INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"
//  INVALIDPARAMETERVALUE_STORAGEREGIONS = "InvalidParameterValue.StorageRegions"
//  INVALIDPARAMETERVALUE_STREAMIDS = "InvalidParameterValue.StreamIds"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  INVALIDPARAMETERVALUE_TAGS = "InvalidParameterValue.Tags"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  INVALIDPARAMETERVALUE_TYPES = "InvalidParameterValue.Types"
//  INVALIDPARAMETERVALUE_VIDS = "InvalidParameterValue.Vids"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchMediaWithContext(ctx context.Context, request *SearchMediaRequest) (response *SearchMediaResponse, err error) {
    if request == nil {
        request = NewSearchMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchMediaResponse()
    err = c.Send(request, response)
    return
}

func NewSetDrmKeyProviderInfoRequest() (request *SetDrmKeyProviderInfoRequest) {
    request = &SetDrmKeyProviderInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "SetDrmKeyProviderInfo")
    
    
    return
}

func NewSetDrmKeyProviderInfoResponse() (response *SetDrmKeyProviderInfoResponse) {
    response = &SetDrmKeyProviderInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetDrmKeyProviderInfo
// This API is used to configure DRM key information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SetDrmKeyProviderInfo(request *SetDrmKeyProviderInfoRequest) (response *SetDrmKeyProviderInfoResponse, err error) {
    return c.SetDrmKeyProviderInfoWithContext(context.Background(), request)
}

// SetDrmKeyProviderInfo
// This API is used to configure DRM key information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SetDrmKeyProviderInfoWithContext(ctx context.Context, request *SetDrmKeyProviderInfoRequest) (response *SetDrmKeyProviderInfoResponse, err error) {
    if request == nil {
        request = NewSetDrmKeyProviderInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetDrmKeyProviderInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetDrmKeyProviderInfoResponse()
    err = c.Send(request, response)
    return
}

func NewSetVodDomainCertificateRequest() (request *SetVodDomainCertificateRequest) {
    request = &SetVodDomainCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "SetVodDomainCertificate")
    
    
    return
}

func NewSetVodDomainCertificateResponse() (response *SetVodDomainCertificateResponse) {
    response = &SetVodDomainCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetVodDomainCertificate
// Set Vod Domain Certificate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SetVodDomainCertificate(request *SetVodDomainCertificateRequest) (response *SetVodDomainCertificateResponse, err error) {
    return c.SetVodDomainCertificateWithContext(context.Background(), request)
}

// SetVodDomainCertificate
// Set Vod Domain Certificate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SetVodDomainCertificateWithContext(ctx context.Context, request *SetVodDomainCertificateRequest) (response *SetVodDomainCertificateResponse, err error) {
    if request == nil {
        request = NewSetVodDomainCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetVodDomainCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetVodDomainCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewSimpleHlsClipRequest() (request *SimpleHlsClipRequest) {
    request = &SimpleHlsClipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "SimpleHlsClip")
    
    
    return
}

func NewSimpleHlsClipResponse() (response *SimpleHlsClipResponse) {
    response = &SimpleHlsClipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SimpleHlsClip
// This API is used to cut a clip from an HLS video to generate a new video (in HLS format). You can either share the new video or save it.
//
// 
//
// VOD supports two types of clipping:
//
// - Clipping for persistent storage: The video clip is saved as an independent video file with its own `FileId`.
//
// - Clipping for temporary sharing: The video clip shares `FileId` with the input file.
//
// 
//
// Notes:
//
// - Clipping is based on the M3U8 file that contains a list of TS segments, so the smallest clipping unit is one TS segment instead of a second or less.
//
// 
//
// 
//
// ### Clipping for persistent storage
//
// In this mode, a video clip is saved as an independent video file with a `FileId`, and its lifecycle is not subject to the input video. Even if the source video is deleted, the video clip still exists. Moreover, the video clip can be transcoded, published on WeChat, and processed in other ways.
//
// 
//
// Suppose you recorded a two-hour football match. You want to save the full video for only two months to save costs, but want to save the highlights for a longer time and perhaps transcode and publish the highlight clip to WeChat. In this case, you can choose clipping for persistent storage.
//
// 
//
// The advantage of clipping for persistent storage is that the video clip has a lifecycle independent of the input video and can be managed independently and stored persistently.
//
// 
//
// <font color='red'>Notice:</font> If solidification is specified during clipping, enable the reception of persistence completed event notifications through the ModifyEventConfig interface. After successful solidification, an event notification of type PersistenceComplete will be received. Before receiving this event notification, operations such as deletion or cooling down of live recording videos should not be performed, otherwise, abnormal playback may occur in the video generated by the clipping.
//
// 
//
// ### Clipping for temporary sharing
//
// The video clip (an M3U8 file) shares the same TS segments with the input video instead of being an independent video. It only has a playback URL but has no `FileId`, and its validity period is the same as that of the input video. Once the input video is deleted, the video clip cannot be played back.
//
// 
//
// Because the video clip is not an independent video, it's not displayed as a media asset in the VOD console, and cannot be transcoded or published to WeChat.
//
// 
//
// Clipping for temporary sharing is lightweight and incurs no additional storage fees. However, the video clip has the same lifecycle as the source recording video and cannot be transcoded or processed in other ways.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEOFFSET = "InvalidParameterValue.EndTimeOffset"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_ISPERSISTENCE = "InvalidParameterValue.IsPersistence"
//  INVALIDPARAMETERVALUE_OUTPUTMEDIATYPE = "InvalidParameterValue.OutputMediaType"
//  INVALIDPARAMETERVALUE_PRECISION = "InvalidParameterValue.Precision"
//  INVALIDPARAMETERVALUE_PROCEDURE = "InvalidParameterValue.Procedure"
//  INVALIDPARAMETERVALUE_STARTTIMEOFFSET = "InvalidParameterValue.StartTimeOffset"
//  INVALIDPARAMETERVALUE_URL = "InvalidParameterValue.Url"
//  RESOURCEUNAVAILABLE_MASTERPLAYLIST = "ResourceUnavailable.MasterPlaylist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SimpleHlsClip(request *SimpleHlsClipRequest) (response *SimpleHlsClipResponse, err error) {
    return c.SimpleHlsClipWithContext(context.Background(), request)
}

// SimpleHlsClip
// This API is used to cut a clip from an HLS video to generate a new video (in HLS format). You can either share the new video or save it.
//
// 
//
// VOD supports two types of clipping:
//
// - Clipping for persistent storage: The video clip is saved as an independent video file with its own `FileId`.
//
// - Clipping for temporary sharing: The video clip shares `FileId` with the input file.
//
// 
//
// Notes:
//
// - Clipping is based on the M3U8 file that contains a list of TS segments, so the smallest clipping unit is one TS segment instead of a second or less.
//
// 
//
// 
//
// ### Clipping for persistent storage
//
// In this mode, a video clip is saved as an independent video file with a `FileId`, and its lifecycle is not subject to the input video. Even if the source video is deleted, the video clip still exists. Moreover, the video clip can be transcoded, published on WeChat, and processed in other ways.
//
// 
//
// Suppose you recorded a two-hour football match. You want to save the full video for only two months to save costs, but want to save the highlights for a longer time and perhaps transcode and publish the highlight clip to WeChat. In this case, you can choose clipping for persistent storage.
//
// 
//
// The advantage of clipping for persistent storage is that the video clip has a lifecycle independent of the input video and can be managed independently and stored persistently.
//
// 
//
// <font color='red'>Notice:</font> If solidification is specified during clipping, enable the reception of persistence completed event notifications through the ModifyEventConfig interface. After successful solidification, an event notification of type PersistenceComplete will be received. Before receiving this event notification, operations such as deletion or cooling down of live recording videos should not be performed, otherwise, abnormal playback may occur in the video generated by the clipping.
//
// 
//
// ### Clipping for temporary sharing
//
// The video clip (an M3U8 file) shares the same TS segments with the input video instead of being an independent video. It only has a playback URL but has no `FileId`, and its validity period is the same as that of the input video. Once the input video is deleted, the video clip cannot be played back.
//
// 
//
// Because the video clip is not an independent video, it's not displayed as a media asset in the VOD console, and cannot be transcoded or published to WeChat.
//
// 
//
// Clipping for temporary sharing is lightweight and incurs no additional storage fees. However, the video clip has the same lifecycle as the source recording video and cannot be transcoded or processed in other ways.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEOFFSET = "InvalidParameterValue.EndTimeOffset"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_ISPERSISTENCE = "InvalidParameterValue.IsPersistence"
//  INVALIDPARAMETERVALUE_OUTPUTMEDIATYPE = "InvalidParameterValue.OutputMediaType"
//  INVALIDPARAMETERVALUE_PRECISION = "InvalidParameterValue.Precision"
//  INVALIDPARAMETERVALUE_PROCEDURE = "InvalidParameterValue.Procedure"
//  INVALIDPARAMETERVALUE_STARTTIMEOFFSET = "InvalidParameterValue.StartTimeOffset"
//  INVALIDPARAMETERVALUE_URL = "InvalidParameterValue.Url"
//  RESOURCEUNAVAILABLE_MASTERPLAYLIST = "ResourceUnavailable.MasterPlaylist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SimpleHlsClipWithContext(ctx context.Context, request *SimpleHlsClipRequest) (response *SimpleHlsClipResponse, err error) {
    if request == nil {
        request = NewSimpleHlsClipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SimpleHlsClip require credential")
    }

    request.SetContext(ctx)
    
    response = NewSimpleHlsClipResponse()
    err = c.Send(request, response)
    return
}

func NewSplitMediaRequest() (request *SplitMediaRequest) {
    request = &SplitMediaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "SplitMedia")
    
    
    return
}

func NewSplitMediaResponse() (response *SplitMediaResponse) {
    response = &SplitMediaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SplitMedia
// Split the video into strips to generate multiple new videos.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SplitMedia(request *SplitMediaRequest) (response *SplitMediaResponse, err error) {
    return c.SplitMediaWithContext(context.Background(), request)
}

// SplitMedia
// Split the video into strips to generate multiple new videos.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SplitMediaWithContext(ctx context.Context, request *SplitMediaRequest) (response *SplitMediaResponse, err error) {
    if request == nil {
        request = NewSplitMediaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SplitMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewSplitMediaResponse()
    err = c.Send(request, response)
    return
}

func NewStartCDNDomainRequest() (request *StartCDNDomainRequest) {
    request = &StartCDNDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "StartCDNDomain")
    
    
    return
}

func NewStartCDNDomainResponse() (response *StartCDNDomainResponse) {
    response = &StartCDNDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartCDNDomain
// This API is used for enabling/disabling the CDN acceleration domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StartCDNDomain(request *StartCDNDomainRequest) (response *StartCDNDomainResponse, err error) {
    return c.StartCDNDomainWithContext(context.Background(), request)
}

// StartCDNDomain
// This API is used for enabling/disabling the CDN acceleration domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StartCDNDomainWithContext(ctx context.Context, request *StartCDNDomainRequest) (response *StartCDNDomainResponse, err error) {
    if request == nil {
        request = NewStartCDNDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCDNDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartCDNDomainResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyDomainRecordRequest() (request *VerifyDomainRecordRequest) {
    request = &VerifyDomainRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "VerifyDomainRecord")
    
    
    return
}

func NewVerifyDomainRecordResponse() (response *VerifyDomainRecordResponse) {
    response = &VerifyDomainRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyDomainRecord
// This API is used to verify the domain name resolution value.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) VerifyDomainRecord(request *VerifyDomainRecordRequest) (response *VerifyDomainRecordResponse, err error) {
    return c.VerifyDomainRecordWithContext(context.Background(), request)
}

// VerifyDomainRecord
// This API is used to verify the domain name resolution value.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) VerifyDomainRecordWithContext(ctx context.Context, request *VerifyDomainRecordRequest) (response *VerifyDomainRecordResponse, err error) {
    if request == nil {
        request = NewVerifyDomainRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyDomainRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyDomainRecordResponse()
    err = c.Send(request, response)
    return
}
