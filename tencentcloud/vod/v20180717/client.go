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
// * We strongly recommend that you use the [server-side upload SDK](https://www.tencentcloud.com/document/product/266/33912#1.-initiate-upload) provided by VOD to upload files. Directly invoking the API for upload is significantly more difficult and involves a much larger workload than using the SDK.
//
// This API is used to apply for uploading media files (and cover files), obtain the meta information for uploading files to VOD (including upload path and upload signature), and is used for subsequent upload APIs.
//
// For the upload process, see [Server-Side Upload Overview](https://www.tencentcloud.com/document/product/266/9759?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXPIRETIME = "InvalidParameter.ExpireTime"
//  INVALIDPARAMETERVALUE_COVERTYPE = "InvalidParameterValue.CoverType"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_MEDIASTORAGEPATH = "InvalidParameterValue.MediaStoragePath"
//  INVALIDPARAMETERVALUE_MEDIATYPE = "InvalidParameterValue.MediaType"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ApplyUpload(request *ApplyUploadRequest) (response *ApplyUploadResponse, err error) {
    return c.ApplyUploadWithContext(context.Background(), request)
}

// ApplyUpload
// * We strongly recommend that you use the [server-side upload SDK](https://www.tencentcloud.com/document/product/266/33912#1.-initiate-upload) provided by VOD to upload files. Directly invoking the API for upload is significantly more difficult and involves a much larger workload than using the SDK.
//
// This API is used to apply for uploading media files (and cover files), obtain the meta information for uploading files to VOD (including upload path and upload signature), and is used for subsequent upload APIs.
//
// For the upload process, see [Server-Side Upload Overview](https://www.tencentcloud.com/document/product/266/9759?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXPIRETIME = "InvalidParameter.ExpireTime"
//  INVALIDPARAMETERVALUE_COVERTYPE = "InvalidParameterValue.CoverType"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_MEDIASTORAGEPATH = "InvalidParameterValue.MediaStoragePath"
//  INVALIDPARAMETERVALUE_MEDIATYPE = "InvalidParameterValue.MediaType"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ApplyUploadWithContext(ctx context.Context, request *ApplyUploadRequest) (response *ApplyUploadResponse, err error) {
    if request == nil {
        request = NewApplyUploadRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ApplyUpload")
    
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
// Associate media asset subtitles with the media output file corresponding to the designated adaptive bitrate streaming template ID (or disassociate them).
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
// Associate media asset subtitles with the media output file corresponding to the designated adaptive bitrate streaming template ID (or disassociate them).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "AttachMediaSubtitles")
    
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
// Clone a CDN domain.
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
// Clone a CDN domain.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CloneCDNDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneCDNDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneCDNDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCloneVoiceAsyncRequest() (request *CloneVoiceAsyncRequest) {
    request = &CloneVoiceAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CloneVoiceAsync")
    
    
    return
}

func NewCloneVoiceAsyncResponse() (response *CloneVoiceAsyncResponse) {
    response = &CloneVoiceAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloneVoiceAsync
// This API is used to initiate a voice cloning task. It clones an exclusive voice based on reference audio. The generated voice can be used for subsequent text to speech. Voice cloning is an asynchronous task. The voice ID and audio audition are generated after task completion.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CloneVoiceAsync(request *CloneVoiceAsyncRequest) (response *CloneVoiceAsyncResponse, err error) {
    return c.CloneVoiceAsyncWithContext(context.Background(), request)
}

// CloneVoiceAsync
// This API is used to initiate a voice cloning task. It clones an exclusive voice based on reference audio. The generated voice can be used for subsequent text to speech. Voice cloning is an asynchronous task. The voice ID and audio audition are generated after task completion.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CloneVoiceAsyncWithContext(ctx context.Context, request *CloneVoiceAsyncRequest) (response *CloneVoiceAsyncResponse, err error) {
    if request == nil {
        request = NewCloneVoiceAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CloneVoiceAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneVoiceAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneVoiceAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewCloneVoiceSyncRequest() (request *CloneVoiceSyncRequest) {
    request = &CloneVoiceSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CloneVoiceSync")
    
    
    return
}

func NewCloneVoiceSyncResponse() (response *CloneVoiceSyncResponse) {
    response = &CloneVoiceSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloneVoiceSync
// This API is used to initiate a voice cloning task. It clones an exclusive voice based on reference audio. The generated voice can be used for subsequent text to speech.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CloneVoiceSync(request *CloneVoiceSyncRequest) (response *CloneVoiceSyncResponse, err error) {
    return c.CloneVoiceSyncWithContext(context.Background(), request)
}

// CloneVoiceSync
// This API is used to initiate a voice cloning task. It clones an exclusive voice based on reference audio. The generated voice can be used for subsequent text to speech.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDACCOUNT = "FailedOperation.InvalidAccount"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CloneVoiceSyncWithContext(ctx context.Context, request *CloneVoiceSyncRequest) (response *CloneVoiceSyncResponse, err error) {
    if request == nil {
        request = NewCloneVoiceSyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CloneVoiceSync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneVoiceSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneVoiceSyncResponse()
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
// This API is used to confirm the upload result of media files and cover files to Tencent Cloud VOD, store media information, and return the playback address and file ID.
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
// This API is used to confirm the upload result of media files and cover files to Tencent Cloud VOD, store media information, and return the playback address and file ID.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CommitUpload")
    
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
// This API is used to compose media files to achieve the following effects:
//
// 
//
// 1. **Image rotation**: Rotate the video or image by a certain degree, or flip it in a certain direction.
//
// 2. **Audio control**: Increase or reduce the volume of video and audio, or mute the video.
//
// 3. **Screen overlay**: Overlay frames from videos and images in sequence, for example, to achieve a Picture-in-Picture effect.
//
// 4. **Audio mixing**: Mix the sound in video and audio together.
//
// 5. **Audio extraction**: Extract the audio from the video (visuals are not retained).
//
// 6. **Crop**: Crop a specified time period from video or audio.
//
// 7. **Splicing**: Splice videos, audio, and images in chronological order.
//
// 8. **Transitions**: When stitching multiple videos or images, you can add transition effects between paragraphs.
//
// 
//
// The muxing format of the composed media can be MP4 (video) or MP3 (audio). If event notification is used, the event notification type is [Video Synthesis Completed](https://www.tencentcloud.com/document/product/266/43000?from_cn_redirect=1).
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
// This API is used to compose media files to achieve the following effects:
//
// 
//
// 1. **Image rotation**: Rotate the video or image by a certain degree, or flip it in a certain direction.
//
// 2. **Audio control**: Increase or reduce the volume of video and audio, or mute the video.
//
// 3. **Screen overlay**: Overlay frames from videos and images in sequence, for example, to achieve a Picture-in-Picture effect.
//
// 4. **Audio mixing**: Mix the sound in video and audio together.
//
// 5. **Audio extraction**: Extract the audio from the video (visuals are not retained).
//
// 6. **Crop**: Crop a specified time period from video or audio.
//
// 7. **Splicing**: Splice videos, audio, and images in chronological order.
//
// 8. **Transitions**: When stitching multiple videos or images, you can add transition effects between paragraphs.
//
// 
//
// The muxing format of the composed media can be MP4 (video) or MP3 (audio). If event notification is used, the event notification type is [Video Synthesis Completed](https://www.tencentcloud.com/document/product/266/43000?from_cn_redirect=1).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ComposeMedia")
    
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
// * Developers call the event notification pull API. After obtaining an event, they must call this API to acknowledge that the message has been received.
//
// * After the developer obtains the event handler, the validity time for pending confirmation is 30 seconds. If it exceeds 30 seconds, a parameter error (4000) will be reported.
//
// * For more references on reliable callback for event notification, see [Reliable Callback](https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1#.E5.8F.AF.E9.9D.A0.E5.9B.9E.E8.B0.83).
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
// * Developers call the event notification pull API. After obtaining an event, they must call this API to acknowledge that the message has been received.
//
// * After the developer obtains the event handler, the validity time for pending confirmation is 30 seconds. If it exceeds 30 seconds, a parameter error (4000) will be reported.
//
// * For more references on reliable callback for event notification, see [Reliable Callback](https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1#.E5.8F.AF.E9.9D.A0.E5.9B.9E.E8.B0.83).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ConfirmEvents")
    
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
// This API is used to create a user-defined audio and video content analysis template. Maximum quantity: 50. HLS format is not supported currently.
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
// This API is used to create a user-defined audio and video content analysis template. Maximum quantity: 50. HLS format is not supported currently.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAIAnalysisTemplate")
    
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
// This API is used to create a user-defined audio and video content recognition template. Maximum quantity: 50.
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
// This API is used to create a user-defined audio and video content recognition template. Maximum quantity: 50.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAIRecognitionTemplate")
    
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
// Create adaptive bitrate streaming templates. Maximum quantity: 100.
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
// Create adaptive bitrate streaming templates. Maximum quantity: 100.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAdaptiveDynamicStreamingTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcAdvancedCustomElementRequest() (request *CreateAigcAdvancedCustomElementRequest) {
    request = &CreateAigcAdvancedCustomElementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcAdvancedCustomElement")
    
    
    return
}

func NewCreateAigcAdvancedCustomElementResponse() (response *CreateAigcAdvancedCustomElementResponse) {
    response = &CreateAigcAdvancedCustomElementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcAdvancedCustomElement
// This API is used to create an AIGC advanced custom subject.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcAdvancedCustomElement(request *CreateAigcAdvancedCustomElementRequest) (response *CreateAigcAdvancedCustomElementResponse, err error) {
    return c.CreateAigcAdvancedCustomElementWithContext(context.Background(), request)
}

// CreateAigcAdvancedCustomElement
// This API is used to create an AIGC advanced custom subject.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcAdvancedCustomElementWithContext(ctx context.Context, request *CreateAigcAdvancedCustomElementRequest) (response *CreateAigcAdvancedCustomElementResponse, err error) {
    if request == nil {
        request = NewCreateAigcAdvancedCustomElementRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcAdvancedCustomElement")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcAdvancedCustomElement require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcAdvancedCustomElementResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcApiTokenRequest() (request *CreateAigcApiTokenRequest) {
    request = &CreateAigcApiTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcApiToken")
    
    
    return
}

func NewCreateAigcApiTokenResponse() (response *CreateAigcApiTokenResponse) {
    response = &CreateAigcApiTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcApiToken
// This API is used to create a Token for AIGC API calls. Data sync may be delayed once created. It can be queried or deleted after about 30 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateAigcApiToken(request *CreateAigcApiTokenRequest) (response *CreateAigcApiTokenResponse, err error) {
    return c.CreateAigcApiTokenWithContext(context.Background(), request)
}

// CreateAigcApiToken
// This API is used to create a Token for AIGC API calls. Data sync may be delayed once created. It can be queried or deleted after about 30 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateAigcApiTokenWithContext(ctx context.Context, request *CreateAigcApiTokenRequest) (response *CreateAigcApiTokenResponse, err error) {
    if request == nil {
        request = NewCreateAigcApiTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcApiToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcApiToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcApiTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcAudioCloneRequest() (request *CreateAigcAudioCloneRequest) {
    request = &CreateAigcAudioCloneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcAudioClone")
    
    
    return
}

func NewCreateAigcAudioCloneResponse() (response *CreateAigcAudioCloneResponse) {
    response = &CreateAigcAudioCloneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcAudioClone
// This API is used to create AIGC voice replication. Note that calling this API incurs fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#96b3b59a-f9e1-49e9-966a-bedb70a4bf12).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcAudioClone(request *CreateAigcAudioCloneRequest) (response *CreateAigcAudioCloneResponse, err error) {
    return c.CreateAigcAudioCloneWithContext(context.Background(), request)
}

// CreateAigcAudioClone
// This API is used to create AIGC voice replication. Note that calling this API incurs fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#96b3b59a-f9e1-49e9-966a-bedb70a4bf12).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcAudioCloneWithContext(ctx context.Context, request *CreateAigcAudioCloneRequest) (response *CreateAigcAudioCloneResponse, err error) {
    if request == nil {
        request = NewCreateAigcAudioCloneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcAudioClone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcAudioClone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcAudioCloneResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcAudioTaskRequest() (request *CreateAigcAudioTaskRequest) {
    request = &CreateAigcAudioTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcAudioTask")
    
    
    return
}

func NewCreateAigcAudioTaskResponse() (response *CreateAigcAudioTaskResponse) {
    response = &CreateAigcAudioTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcAudioTask
// This API is used to create AI audio generation tasks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIOLATIONCONTENT = "InvalidParameter.ViolationContent"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) CreateAigcAudioTask(request *CreateAigcAudioTaskRequest) (response *CreateAigcAudioTaskResponse, err error) {
    return c.CreateAigcAudioTaskWithContext(context.Background(), request)
}

// CreateAigcAudioTask
// This API is used to create AI audio generation tasks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIOLATIONCONTENT = "InvalidParameter.ViolationContent"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) CreateAigcAudioTaskWithContext(ctx context.Context, request *CreateAigcAudioTaskRequest) (response *CreateAigcAudioTaskResponse, err error) {
    if request == nil {
        request = NewCreateAigcAudioTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcAudioTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcAudioTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcAudioTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcCustomElementRequest() (request *CreateAigcCustomElementRequest) {
    request = &CreateAigcCustomElementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcCustomElement")
    
    
    return
}

func NewCreateAigcCustomElementResponse() (response *CreateAigcCustomElementResponse) {
    response = &CreateAigcCustomElementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcCustomElement
// Call this API to create a subject for a specified model.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAigcCustomElement(request *CreateAigcCustomElementRequest) (response *CreateAigcCustomElementResponse, err error) {
    return c.CreateAigcCustomElementWithContext(context.Background(), request)
}

// CreateAigcCustomElement
// Call this API to create a subject for a specified model.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAigcCustomElementWithContext(ctx context.Context, request *CreateAigcCustomElementRequest) (response *CreateAigcCustomElementResponse, err error) {
    if request == nil {
        request = NewCreateAigcCustomElementRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcCustomElement")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcCustomElement require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcCustomElementResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcCustomVoiceRequest() (request *CreateAigcCustomVoiceRequest) {
    request = &CreateAigcCustomVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcCustomVoice")
    
    
    return
}

func NewCreateAigcCustomVoiceResponse() (response *CreateAigcCustomVoiceResponse) {
    response = &CreateAigcCustomVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcCustomVoice
// This API is used to create AIGC custom voice types. Note that calling this API incurs custom voice type creation fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#5e5217e8-29fc-467e-ac2d-853648f988b7).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcCustomVoice(request *CreateAigcCustomVoiceRequest) (response *CreateAigcCustomVoiceResponse, err error) {
    return c.CreateAigcCustomVoiceWithContext(context.Background(), request)
}

// CreateAigcCustomVoice
// This API is used to create AIGC custom voice types. Note that calling this API incurs custom voice type creation fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#5e5217e8-29fc-467e-ac2d-853648f988b7).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcCustomVoiceWithContext(ctx context.Context, request *CreateAigcCustomVoiceRequest) (response *CreateAigcCustomVoiceResponse, err error) {
    if request == nil {
        request = NewCreateAigcCustomVoiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcCustomVoice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcCustomVoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcCustomVoiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcHunyuan3DTaskRequest() (request *CreateAigcHunyuan3DTaskRequest) {
    request = &CreateAigcHunyuan3DTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcHunyuan3DTask")
    
    
    return
}

func NewCreateAigcHunyuan3DTaskResponse() (response *CreateAigcHunyuan3DTaskResponse) {
    response = &CreateAigcHunyuan3DTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcHunyuan3DTask
// This API is used to create AIGC Hunyuan 3D tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcHunyuan3DTask(request *CreateAigcHunyuan3DTaskRequest) (response *CreateAigcHunyuan3DTaskResponse, err error) {
    return c.CreateAigcHunyuan3DTaskWithContext(context.Background(), request)
}

// CreateAigcHunyuan3DTask
// This API is used to create AIGC Hunyuan 3D tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcHunyuan3DTaskWithContext(ctx context.Context, request *CreateAigcHunyuan3DTaskRequest) (response *CreateAigcHunyuan3DTaskResponse, err error) {
    if request == nil {
        request = NewCreateAigcHunyuan3DTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcHunyuan3DTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcHunyuan3DTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcHunyuan3DTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcImageTaskRequest() (request *CreateAigcImageTaskRequest) {
    request = &CreateAigcImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcImageTask")
    
    
    return
}

func NewCreateAigcImageTaskResponse() (response *CreateAigcImageTaskResponse) {
    response = &CreateAigcImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcImageTask
// This API is used to generate AIGC images. The default limit is 1 concurrent processing. API calls will incur actual fees. Refer to the VOD AIGC image generation billing documentation. The settlement mode for the feature is pay-as-you-go. For daily billing customers, usage on the day is billed on the second day. For monthly settlement customers, the usage fees of the previous month are billed on the 1st of the next month.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcImageTask(request *CreateAigcImageTaskRequest) (response *CreateAigcImageTaskResponse, err error) {
    return c.CreateAigcImageTaskWithContext(context.Background(), request)
}

// CreateAigcImageTask
// This API is used to generate AIGC images. The default limit is 1 concurrent processing. API calls will incur actual fees. Refer to the VOD AIGC image generation billing documentation. The settlement mode for the feature is pay-as-you-go. For daily billing customers, usage on the day is billed on the second day. For monthly settlement customers, the usage fees of the previous month are billed on the 1st of the next month.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcImageTaskWithContext(ctx context.Context, request *CreateAigcImageTaskRequest) (response *CreateAigcImageTaskResponse, err error) {
    if request == nil {
        request = NewCreateAigcImageTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcImageTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcImageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcQuotaRequest() (request *CreateAigcQuotaRequest) {
    request = &CreateAigcQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcQuota")
    
    
    return
}

func NewCreateAigcQuotaResponse() (response *CreateAigcQuotaResponse) {
    response = &CreateAigcQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcQuota
// This API is used to create and enable AIGC quota configuration. Quota usage starts accumulating when the quota feature is enabled. Once the quota is reached, AIGC features will no longer be usable.
//
// 
//
// If the quota is deleted and re-enabled, the amount will be cleared and recalculated.
//
// 
//
// Since AGC content generation is an async task, real-time usage data cannot be obtained. Therefore, quota limits result in some errors, and precise control over the set limit cannot be achieved.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateAigcQuota(request *CreateAigcQuotaRequest) (response *CreateAigcQuotaResponse, err error) {
    return c.CreateAigcQuotaWithContext(context.Background(), request)
}

// CreateAigcQuota
// This API is used to create and enable AIGC quota configuration. Quota usage starts accumulating when the quota feature is enabled. Once the quota is reached, AIGC features will no longer be usable.
//
// 
//
// If the quota is deleted and re-enabled, the amount will be cleared and recalculated.
//
// 
//
// Since AGC content generation is an async task, real-time usage data cannot be obtained. Therefore, quota limits result in some errors, and precise control over the set limit cannot be achieved.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateAigcQuotaWithContext(ctx context.Context, request *CreateAigcQuotaRequest) (response *CreateAigcQuotaResponse, err error) {
    if request == nil {
        request = NewCreateAigcQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcSubjectRequest() (request *CreateAigcSubjectRequest) {
    request = &CreateAigcSubjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcSubject")
    
    
    return
}

func NewCreateAigcSubjectResponse() (response *CreateAigcSubjectResponse) {
    response = &CreateAigcSubjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcSubject
// This API is used to create AIGC custom subjects (Vidu). Note that calling this API incurs fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#96b3b59a-f9e1-49e9-966a-bedb70a4bf12).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcSubject(request *CreateAigcSubjectRequest) (response *CreateAigcSubjectResponse, err error) {
    return c.CreateAigcSubjectWithContext(context.Background(), request)
}

// CreateAigcSubject
// This API is used to create AIGC custom subjects (Vidu). Note that calling this API incurs fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#96b3b59a-f9e1-49e9-966a-bedb70a4bf12).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcSubjectWithContext(ctx context.Context, request *CreateAigcSubjectRequest) (response *CreateAigcSubjectResponse, err error) {
    if request == nil {
        request = NewCreateAigcSubjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcSubject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcSubject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcSubjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcVideoRedrawTaskRequest() (request *CreateAigcVideoRedrawTaskRequest) {
    request = &CreateAigcVideoRedrawTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcVideoRedrawTask")
    
    
    return
}

func NewCreateAigcVideoRedrawTaskResponse() (response *CreateAigcVideoRedrawTaskResponse) {
    response = &CreateAigcVideoRedrawTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcVideoRedrawTask
// This API is used to generate AIGC videos. API calls incur actual fees. Refer to the VOD AIGC video generation billing documentation. The settlement mode of this feature is pay-as-you-go. For daily billing customers, usage on the day is billed on the second day. For monthly billing customers, usage fees of the previous month are billed on the 1st of the next month.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcVideoRedrawTask(request *CreateAigcVideoRedrawTaskRequest) (response *CreateAigcVideoRedrawTaskResponse, err error) {
    return c.CreateAigcVideoRedrawTaskWithContext(context.Background(), request)
}

// CreateAigcVideoRedrawTask
// This API is used to generate AIGC videos. API calls incur actual fees. Refer to the VOD AIGC video generation billing documentation. The settlement mode of this feature is pay-as-you-go. For daily billing customers, usage on the day is billed on the second day. For monthly billing customers, usage fees of the previous month are billed on the 1st of the next month.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcVideoRedrawTaskWithContext(ctx context.Context, request *CreateAigcVideoRedrawTaskRequest) (response *CreateAigcVideoRedrawTaskResponse, err error) {
    if request == nil {
        request = NewCreateAigcVideoRedrawTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcVideoRedrawTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcVideoRedrawTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcVideoRedrawTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcVideoTaskRequest() (request *CreateAigcVideoTaskRequest) {
    request = &CreateAigcVideoTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateAigcVideoTask")
    
    
    return
}

func NewCreateAigcVideoTaskResponse() (response *CreateAigcVideoTaskResponse) {
    response = &CreateAigcVideoTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcVideoTask
// This API is used to generate AIGC videos. The default limit is 1 concurrent processing. API calls incur actual fees. Refer to the VOD AIGC video generation billing documentation. The feature uses postpaid settlement mode. For daily billing customers, usage on the day is billed on the second day. For monthly settlement customers, the previous month's usage fees are billed on the 1st of the next month.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcVideoTask(request *CreateAigcVideoTaskRequest) (response *CreateAigcVideoTaskResponse, err error) {
    return c.CreateAigcVideoTaskWithContext(context.Background(), request)
}

// CreateAigcVideoTask
// This API is used to generate AIGC videos. The default limit is 1 concurrent processing. API calls incur actual fees. Refer to the VOD AIGC video generation billing documentation. The feature uses postpaid settlement mode. For daily billing customers, usage on the day is billed on the second day. For monthly settlement customers, the previous month's usage fees are billed on the 1st of the next month.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAigcVideoTaskWithContext(ctx context.Context, request *CreateAigcVideoTaskRequest) (response *CreateAigcVideoTaskResponse, err error) {
    if request == nil {
        request = NewCreateAigcVideoTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAigcVideoTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcVideoTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcVideoTaskResponse()
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
// This API is used to create a custom animated image generating template. Maximum number: 16.
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
// This API is used to create a custom animated image generating template. Maximum number: 16.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateAnimatedGraphicsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBlindWatermarkTemplateRequest() (request *CreateBlindWatermarkTemplateRequest) {
    request = &CreateBlindWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateBlindWatermarkTemplate")
    
    
    return
}

func NewCreateBlindWatermarkTemplateResponse() (response *CreateBlindWatermarkTemplateResponse) {
    response = &CreateBlindWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBlindWatermarkTemplate
// This API is used to create a user-defined digital watermark template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateBlindWatermarkTemplate(request *CreateBlindWatermarkTemplateRequest) (response *CreateBlindWatermarkTemplateResponse, err error) {
    return c.CreateBlindWatermarkTemplateWithContext(context.Background(), request)
}

// CreateBlindWatermarkTemplate
// This API is used to create a user-defined digital watermark template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateBlindWatermarkTemplateWithContext(ctx context.Context, request *CreateBlindWatermarkTemplateRequest) (response *CreateBlindWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewCreateBlindWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateBlindWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBlindWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBlindWatermarkTemplateResponse()
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
// This API is used to add domain names to VOD. A user can add up to 20 domain names. 1. After the domain name is successfully added, VOD will deploy the domain name. It takes about 2 minutes for the domain name to change from the deployment state to the online status.
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
// This API is used to add domain names to VOD. A user can add up to 20 domain names. 1. After the domain name is successfully added, VOD will deploy the domain name. It takes about 2 minutes for the domain name to change from the deployment state to the online status.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateCDNDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCDNDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCDNDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCLSLogsetRequest() (request *CreateCLSLogsetRequest) {
    request = &CreateCLSLogsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateCLSLogset")
    
    
    return
}

func NewCreateCLSLogsetResponse() (response *CreateCLSLogsetResponse) {
    response = &CreateCLSLogsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCLSLogset
// Create a logset via VOD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCLSLogset(request *CreateCLSLogsetRequest) (response *CreateCLSLogsetResponse, err error) {
    return c.CreateCLSLogsetWithContext(context.Background(), request)
}

// CreateCLSLogset
// Create a logset via VOD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCLSLogsetWithContext(ctx context.Context, request *CreateCLSLogsetRequest) (response *CreateCLSLogsetResponse, err error) {
    if request == nil {
        request = NewCreateCLSLogsetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateCLSLogset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCLSLogset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCLSLogsetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCLSTopicRequest() (request *CreateCLSTopicRequest) {
    request = &CreateCLSTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateCLSTopic")
    
    
    return
}

func NewCreateCLSTopicResponse() (response *CreateCLSTopicResponse) {
    response = &CreateCLSTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCLSTopic
// This API is used to create a CLS log topic for VOD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCLSTopic(request *CreateCLSTopicRequest) (response *CreateCLSTopicResponse, err error) {
    return c.CreateCLSTopicWithContext(context.Background(), request)
}

// CreateCLSTopic
// This API is used to create a CLS log topic for VOD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCLSTopicWithContext(ctx context.Context, request *CreateCLSTopicRequest) (response *CreateCLSTopicResponse, err error) {
    if request == nil {
        request = NewCreateCLSTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateCLSTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCLSTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCLSTopicResponse()
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
// * Used to categorize and manage media;
//
// * This API does not affect the category of existing media. To classify media, call the [ModifyMediaInfo](https://www.tencentcloud.com/document/product/266/31762?from_cn_redirect=1) API.
//
// * The classification hierarchy cannot exceed 4 levels.
//
// * The number of subcategories in each category cannot exceed 500.
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
// * Used to categorize and manage media;
//
// * This API does not affect the category of existing media. To classify media, call the [ModifyMediaInfo](https://www.tencentcloud.com/document/product/266/31762?from_cn_redirect=1) API.
//
// * The classification hierarchy cannot exceed 4 levels.
//
// * The number of subcategories in each category cannot exceed 500.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateClass")
    
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
// Initiate a complex adaptive bitstream processing task. Features include:
//
// 1. Output HLS and DASH adaptive bitrate streams based on the designated adaptive bitrate template;
//
// 2. Content protection solutions for adaptive bitrate streams are available in unencrypted, Widevine, or FairPlay.
//
// 3. Support adding opening and ending segments;
//
// 4. The output adaptive bitrate stream can contain multilingual audio streams, with each language coming from a different media file;
//
// 5. The output adaptive bitrate stream can include multilingual subtitle streams.
//
// 
//
// Notes:
//
// 1. When using an opening scene, the video stream in the opening scene media needs to align with the audio stream; otherwise, it will cause audio and video synchronization issues in the output.
//
// 2. If the output adaptive bitrate stream needs to include the audio of the main media, specify the FileId of the main media in the AudioSet parameter.
//
// 3. To use subtitles, add them to the main media first via the ModifyMediaInfo API or the audio and video details page in the console.
//
// 4. Top speed Codec and watermark are not currently supported.
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
// Initiate a complex adaptive bitstream processing task. Features include:
//
// 1. Output HLS and DASH adaptive bitrate streams based on the designated adaptive bitrate template;
//
// 2. Content protection solutions for adaptive bitrate streams are available in unencrypted, Widevine, or FairPlay.
//
// 3. Support adding opening and ending segments;
//
// 4. The output adaptive bitrate stream can contain multilingual audio streams, with each language coming from a different media file;
//
// 5. The output adaptive bitrate stream can include multilingual subtitle streams.
//
// 
//
// Notes:
//
// 1. When using an opening scene, the video stream in the opening scene media needs to align with the audio stream; otherwise, it will cause audio and video synchronization issues in the output.
//
// 2. If the output adaptive bitrate stream needs to include the audio of the main media, specify the FileId of the main media in the AudioSet parameter.
//
// 3. To use subtitles, add them to the main media first via the ModifyMediaInfo API or the audio and video details page in the console.
//
// 4. Top speed Codec and watermark are not currently supported.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateComplexAdaptiveDynamicStreamingTask")
    
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
// This API is <font color=red>no longer maintained</font>. The new version of the moderation template supports video moderation and image moderation. For details, please see [Create Moderation Template](https://www.tencentcloud.com/document/api/266/84391?from_cn_redirect=1).
//
// This API is used to create a user-customized audio/video moderation template. Up to 50 templates can be created.
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
// This API is <font color=red>no longer maintained</font>. The new version of the moderation template supports video moderation and image moderation. For details, please see [Create Moderation Template](https://www.tencentcloud.com/document/api/266/84391?from_cn_redirect=1).
//
// This API is used to create a user-customized audio/video moderation template. Up to 50 templates can be created.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateContentReviewTemplate")
    
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
// This API is used to generate a subdomain name resolution record and prompt the customer to add it to the domain name resolution for wildcard domain name and domain name retrieval ownership verification.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDomainVerifyRecord(request *CreateDomainVerifyRecordRequest) (response *CreateDomainVerifyRecordResponse, err error) {
    return c.CreateDomainVerifyRecordWithContext(context.Background(), request)
}

// CreateDomainVerifyRecord
// This API is used to generate a subdomain name resolution record and prompt the customer to add it to the domain name resolution for wildcard domain name and domain name retrieval ownership verification.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDomainVerifyRecordWithContext(ctx context.Context, request *CreateDomainVerifyRecordRequest) (response *CreateDomainVerifyRecordResponse, err error) {
    if request == nil {
        request = NewCreateDomainVerifyRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateDomainVerifyRecord")
    
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
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Creates an Audio and Video Quality Rebirth Template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEnhanceMediaTemplate(request *CreateEnhanceMediaTemplateRequest) (response *CreateEnhanceMediaTemplateResponse, err error) {
    return c.CreateEnhanceMediaTemplateWithContext(context.Background(), request)
}

// CreateEnhanceMediaTemplate
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Creates an Audio and Video Quality Rebirth Template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEnhanceMediaTemplateWithContext(ctx context.Context, request *CreateEnhanceMediaTemplateRequest) (response *CreateEnhanceMediaTemplateResponse, err error) {
    if request == nil {
        request = NewCreateEnhanceMediaTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateEnhanceMediaTemplate")
    
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
// Creates a title and trailer template.
//
// -Maximum supported template quantity: 100.
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
// Creates a title and trailer template.
//
// -Maximum supported template quantity: 100.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateHeadTailTemplate")
    
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
// Create a custom image processing template. Maximum quantity: 16. Supports up to ten operations, for example: crop - thumbnail - crop - blurry - thumbnail - crop - thumbnail - crop - blurry - thumbnail.
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
// Create a custom image processing template. Maximum quantity: 16. Supports up to ten operations, for example: crop - thumbnail - crop - blurry - thumbnail - crop - thumbnail - crop - blurry - thumbnail.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateImageProcessingTemplate")
    
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
// This API is used to create a user-customized image sprite template. Maximum number: 16.
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
// This API is used to create a user-customized image sprite template. Maximum number: 16.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateImageSpriteTemplate")
    
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
// This API is used to create a just in time transcoding template.
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
// This API is used to create a just in time transcoding template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateJustInTimeTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJustInTimeTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateJustInTimeTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKnowledgeBaseRequest() (request *CreateKnowledgeBaseRequest) {
    request = &CreateKnowledgeBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateKnowledgeBase")
    
    
    return
}

func NewCreateKnowledgeBaseResponse() (response *CreateKnowledgeBaseResponse) {
    response = &CreateKnowledgeBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKnowledgeBase
// Create a knowledge base. This API is used to create a new knowledge base for Intelligent Media Assets. Each user can create up to 20 knowledge bases.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateKnowledgeBase(request *CreateKnowledgeBaseRequest) (response *CreateKnowledgeBaseResponse, err error) {
    return c.CreateKnowledgeBaseWithContext(context.Background(), request)
}

// CreateKnowledgeBase
// Create a knowledge base. This API is used to create a new knowledge base for Intelligent Media Assets. Each user can create up to 20 knowledge bases.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateKnowledgeBaseWithContext(ctx context.Context, request *CreateKnowledgeBaseRequest) (response *CreateKnowledgeBaseResponse, err error) {
    if request == nil {
        request = NewCreateKnowledgeBaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateKnowledgeBase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKnowledgeBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKnowledgeBaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLLMComprehendTemplateRequest() (request *CreateLLMComprehendTemplateRequest) {
    request = &CreateLLMComprehendTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateLLMComprehendTemplate")
    
    
    return
}

func NewCreateLLMComprehendTemplateResponse() (response *CreateLLMComprehendTemplateResponse) {
    response = &CreateLLMComprehendTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLLMComprehendTemplate
// This API is used to create a large model parsing template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_EXTENDEDPARAMETER = "InvalidParameterValue.ExtendedParameter"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateLLMComprehendTemplate(request *CreateLLMComprehendTemplateRequest) (response *CreateLLMComprehendTemplateResponse, err error) {
    return c.CreateLLMComprehendTemplateWithContext(context.Background(), request)
}

// CreateLLMComprehendTemplate
// This API is used to create a large model parsing template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_EXTENDEDPARAMETER = "InvalidParameterValue.ExtendedParameter"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateLLMComprehendTemplateWithContext(ctx context.Context, request *CreateLLMComprehendTemplateRequest) (response *CreateLLMComprehendTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLLMComprehendTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateLLMComprehendTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLLMComprehendTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLLMComprehendTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMPSTemplateRequest() (request *CreateMPSTemplateRequest) {
    request = &CreateMPSTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateMPSTemplate")
    
    
    return
}

func NewCreateMPSTemplateResponse() (response *CreateMPSTemplateResponse) {
    response = &CreateMPSTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMPSTemplate
// This API is used to create a custom template for partial features of the ProcessMediaByMPS API.
//
// When creating a template, fill in MPS related parameters in JSON format into the MPSCreateTemplateParams parameter. For specific task parameter configuration methods, refer to the MPS task template related documentation.
//
// Currently supported MPS features for creating custom templates:
//
// 1. [Audio and video enhancement](https://www.tencentcloud.com/document/product/862/118703?from_cn_redirect=1).
//
// 2. [Media AI](https://www.tencentcloud.com/document/product/862/113756?from_cn_redirect=1)
//
// 
//
// > Template for tasks created using this method:
//
// > Template management is still completed in the VOD platform.
//
// > 2. The feature is currently in beta test. If needed, you can contact us for support.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateMPSTemplate(request *CreateMPSTemplateRequest) (response *CreateMPSTemplateResponse, err error) {
    return c.CreateMPSTemplateWithContext(context.Background(), request)
}

// CreateMPSTemplate
// This API is used to create a custom template for partial features of the ProcessMediaByMPS API.
//
// When creating a template, fill in MPS related parameters in JSON format into the MPSCreateTemplateParams parameter. For specific task parameter configuration methods, refer to the MPS task template related documentation.
//
// Currently supported MPS features for creating custom templates:
//
// 1. [Audio and video enhancement](https://www.tencentcloud.com/document/product/862/118703?from_cn_redirect=1).
//
// 2. [Media AI](https://www.tencentcloud.com/document/product/862/113756?from_cn_redirect=1)
//
// 
//
// > Template for tasks created using this method:
//
// > Template management is still completed in the VOD platform.
//
// > 2. The feature is currently in beta test. If needed, you can contact us for support.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) CreateMPSTemplateWithContext(ctx context.Context, request *CreateMPSTemplateRequest) (response *CreateMPSTemplateResponse, err error) {
    if request == nil {
        request = NewCreateMPSTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateMPSTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMPSTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMPSTemplateResponse()
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
// This API is used to create material samples for video processing such as content recognition and inappropriate video recognition through technologies like facial feature positioning.
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
// This API is used to create material samples for video processing such as content recognition and inappropriate video recognition through technologies like facial feature positioning.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreatePersonSample")
    
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
// This API is used to create user-defined task flow templates. Template capacity limit: 50.
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
// This API is used to create user-defined task flow templates. Template capacity limit: 50.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateProcedureTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProcedureTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProcedureTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcessImageAsyncTemplateRequest() (request *CreateProcessImageAsyncTemplateRequest) {
    request = &CreateProcessImageAsyncTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateProcessImageAsyncTemplate")
    
    
    return
}

func NewCreateProcessImageAsyncTemplateResponse() (response *CreateProcessImageAsyncTemplateResponse) {
    response = &CreateProcessImageAsyncTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProcessImageAsyncTemplate
// This API is used to create a user-customized image async processing template. Maximum quantity: 50. HLS format is not supported currently.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProcessImageAsyncTemplate(request *CreateProcessImageAsyncTemplateRequest) (response *CreateProcessImageAsyncTemplateResponse, err error) {
    return c.CreateProcessImageAsyncTemplateWithContext(context.Background(), request)
}

// CreateProcessImageAsyncTemplate
// This API is used to create a user-customized image async processing template. Maximum quantity: 50. HLS format is not supported currently.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProcessImageAsyncTemplateWithContext(ctx context.Context, request *CreateProcessImageAsyncTemplateRequest) (response *CreateProcessImageAsyncTemplateResponse, err error) {
    if request == nil {
        request = NewCreateProcessImageAsyncTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateProcessImageAsyncTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProcessImageAsyncTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProcessImageAsyncTemplateResponse()
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
// This API is used to create an audio-visual quality inspection template.
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
// This API is used to create an audio-visual quality inspection template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateQualityInspectTemplate")
    
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
// This API is no longer maintained. The new version of the [Audio and Video Quality Revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Revival Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// This API is used to create a video rebirth template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
func (c *Client) CreateRebuildMediaTemplate(request *CreateRebuildMediaTemplateRequest) (response *CreateRebuildMediaTemplateResponse, err error) {
    return c.CreateRebuildMediaTemplateWithContext(context.Background(), request)
}

// CreateRebuildMediaTemplate
// This API is no longer maintained. The new version of the [Audio and Video Quality Revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Revival Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// This API is used to create a video rebirth template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"
func (c *Client) CreateRebuildMediaTemplateWithContext(ctx context.Context, request *CreateRebuildMediaTemplateRequest) (response *CreateRebuildMediaTemplateResponse, err error) {
    if request == nil {
        request = NewCreateRebuildMediaTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateRebuildMediaTemplate")
    
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
// This API is used to create a user-defined moderation template. Maximum quantity: 50.
//
// >Template is applicable only to the [audio/video moderation (ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) and [image moderation (ReviewImage)](https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) APIs.
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
// This API is used to create a user-defined moderation template. Maximum quantity: 50.
//
// >Template is applicable only to the [audio/video moderation (ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) and [image moderation (ReviewImage)](https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) APIs.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateReviewTemplate")
    
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
// This API is used to create a carousel playlist. Maximum quantity: 100.
//
// Each file in the Carousel Playlist can specify a source file or a transcoded file.
//
// The designated file must be in hls format. All playlist files should preferably maintain the same bitrate and resolution.
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
// This API is used to create a carousel playlist. Maximum quantity: 100.
//
// Each file in the Carousel Playlist can specify a source file or a transcoded file.
//
// The designated file must be in hls format. All playlist files should preferably maintain the same bitrate and resolution.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateRoundPlay")
    
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
// This API is used to create a custom sampled screenshot template. Maximum number: 16.
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
// This API is used to create a custom sampled screenshot template. Maximum number: 16.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateSampleSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSampleSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSampleSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSceneAigcImageTaskRequest() (request *CreateSceneAigcImageTaskRequest) {
    request = &CreateSceneAigcImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateSceneAigcImageTask")
    
    
    return
}

func NewCreateSceneAigcImageTaskResponse() (response *CreateSceneAigcImageTaskResponse) {
    response = &CreateSceneAigcImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSceneAigcImageTask
// This API is used to generate scenario-based AIGC images. API calls incur actual fees. Refer to the VOD [AIGC image generation billing document](https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#9c4dc6ff-4b3f-4b25-bf2d-393889dfb9ac). The feature uses the [postpaid](https://www.tencentcloud.com/document/product/266/2838?from_cn_redirect=1) settlement mode. For daily billing customers, usage on the day is billed on the second day. For monthly billing customers, the previous month's usage fees are billed on the 1st of the next month.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSceneAigcImageTask(request *CreateSceneAigcImageTaskRequest) (response *CreateSceneAigcImageTaskResponse, err error) {
    return c.CreateSceneAigcImageTaskWithContext(context.Background(), request)
}

// CreateSceneAigcImageTask
// This API is used to generate scenario-based AIGC images. API calls incur actual fees. Refer to the VOD [AIGC image generation billing document](https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#9c4dc6ff-4b3f-4b25-bf2d-393889dfb9ac). The feature uses the [postpaid](https://www.tencentcloud.com/document/product/266/2838?from_cn_redirect=1) settlement mode. For daily billing customers, usage on the day is billed on the second day. For monthly billing customers, the previous month's usage fees are billed on the 1st of the next month.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSceneAigcImageTaskWithContext(ctx context.Context, request *CreateSceneAigcImageTaskRequest) (response *CreateSceneAigcImageTaskResponse, err error) {
    if request == nil {
        request = NewCreateSceneAigcImageTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateSceneAigcImageTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSceneAigcImageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSceneAigcImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSceneAigcVideoTaskRequest() (request *CreateSceneAigcVideoTaskRequest) {
    request = &CreateSceneAigcVideoTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateSceneAigcVideoTask")
    
    
    return
}

func NewCreateSceneAigcVideoTaskResponse() (response *CreateSceneAigcVideoTaskResponse) {
    response = &CreateSceneAigcVideoTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSceneAigcVideoTask
// This API is used to generate scenario-based AIGC images.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSceneAigcVideoTask(request *CreateSceneAigcVideoTaskRequest) (response *CreateSceneAigcVideoTaskResponse, err error) {
    return c.CreateSceneAigcVideoTaskWithContext(context.Background(), request)
}

// CreateSceneAigcVideoTask
// This API is used to generate scenario-based AIGC images.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSceneAigcVideoTaskWithContext(ctx context.Context, request *CreateSceneAigcVideoTaskRequest) (response *CreateSceneAigcVideoTaskResponse, err error) {
    if request == nil {
        request = NewCreateSceneAigcVideoTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateSceneAigcVideoTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSceneAigcVideoTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSceneAigcVideoTaskResponse()
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
// This API is used to create a user-customized specified time point screenshot template. Maximum quantity: 16.
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
// This API is used to create a user-customized specified time point screenshot template. Maximum quantity: 16.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateSnapshotByTimeOffsetTemplate")
    
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
// 1. When a user enables on-demand services, storage in partial regions is enabled by default. To enable storage in other regions, use this API.
//
// 2. The DescribeStorageRegions API can query all storage regions and opened regions.
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
// 1. When a user enables on-demand services, storage in partial regions is enabled by default. To enable storage in other regions, use this API.
//
// 2. The DescribeStorageRegions API can query all storage regions and opened regions.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateStorageRegion")
    
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
// This API is used to create a VOD application.
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
// This API is used to create a VOD application.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateSubAppId")
    
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
// This API is <font color='red'>no longer maintained</font>. The new version of player signature no longer uses player configuration templates. For details, please see [Player Signature](https://www.tencentcloud.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to create player configurations. Maximum quantity: 100.
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
// This API is <font color='red'>no longer maintained</font>. The new version of player signature no longer uses player configuration templates. For details, please see [Player Signature](https://www.tencentcloud.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to create player configurations. Maximum quantity: 100.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateSuperPlayerConfig")
    
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
// Create custom transcoding templates. Maximum quantity: 100.
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
// Create custom transcoding templates. Maximum quantity: 100.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateTranscodeTemplate")
    
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
// This API is used to add an acceleration domain name to VOD. A user can add up to 20 acceleration domain names.
//
// 1. After the domain name is successfully added, VOD will deploy the domain. It takes about 2 minutes for the domain to change from deployment status to online status.
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
// This API is used to add an acceleration domain name to VOD. A user can add up to 20 acceleration domain names.
//
// 1. After the domain name is successfully added, VOD will deploy the domain. It takes about 2 minutes for the domain to change from deployment status to online status.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateVodDomain")
    
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
// This API is used to create a user-defined watermark template with an upper limit of 1000.
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
// This API is used to create a user-defined watermark template with an upper limit of 1000.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateWatermarkTemplate")
    
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
// This API is used to create keyword samples in batches. The samples are used for video processing such as inappropriate content recognition and content recognition through OCR and ASR technologies.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateWordSamples(request *CreateWordSamplesRequest) (response *CreateWordSamplesResponse, err error) {
    return c.CreateWordSamplesWithContext(context.Background(), request)
}

// CreateWordSamples
// This API is used to create keyword samples in batches. The samples are used for video processing such as inappropriate content recognition and content recognition through OCR and ASR technologies.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateWordSamplesWithContext(ctx context.Context, request *CreateWordSamplesRequest) (response *CreateWordSamplesResponse, err error) {
    if request == nil {
        request = NewCreateWordSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "CreateWordSamples")
    
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
// This API is used to delete a user-defined audio and video content analysis template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and cannot be deleted.
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
// This API is used to delete a user-defined audio and video content analysis template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and cannot be deleted.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteAIAnalysisTemplate")
    
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
// This API is used to delete a user-defined audio and video content recognition template.
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
// This API is used to delete a user-defined audio and video content recognition template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteAIRecognitionTemplate")
    
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
// Delete an adaptive bitrate streaming template
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
// Delete an adaptive bitrate streaming template
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteAdaptiveDynamicStreamingTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAigcAdvancedCustomElementRequest() (request *DeleteAigcAdvancedCustomElementRequest) {
    request = &DeleteAigcAdvancedCustomElementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAigcAdvancedCustomElement")
    
    
    return
}

func NewDeleteAigcAdvancedCustomElementResponse() (response *DeleteAigcAdvancedCustomElementResponse) {
    response = &DeleteAigcAdvancedCustomElementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAigcAdvancedCustomElement
// This API is used to delete AIGC advanced custom subjects.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
func (c *Client) DeleteAigcAdvancedCustomElement(request *DeleteAigcAdvancedCustomElementRequest) (response *DeleteAigcAdvancedCustomElementResponse, err error) {
    return c.DeleteAigcAdvancedCustomElementWithContext(context.Background(), request)
}

// DeleteAigcAdvancedCustomElement
// This API is used to delete AIGC advanced custom subjects.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
func (c *Client) DeleteAigcAdvancedCustomElementWithContext(ctx context.Context, request *DeleteAigcAdvancedCustomElementRequest) (response *DeleteAigcAdvancedCustomElementResponse, err error) {
    if request == nil {
        request = NewDeleteAigcAdvancedCustomElementRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteAigcAdvancedCustomElement")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAigcAdvancedCustomElement require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAigcAdvancedCustomElementResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAigcApiTokenRequest() (request *DeleteAigcApiTokenRequest) {
    request = &DeleteAigcApiTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAigcApiToken")
    
    
    return
}

func NewDeleteAigcApiTokenResponse() (response *DeleteAigcApiTokenResponse) {
    response = &DeleteAigcApiTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAigcApiToken
// Deletes an AIGC API Token. The associated AIGC quota will also be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteAigcApiToken(request *DeleteAigcApiTokenRequest) (response *DeleteAigcApiTokenResponse, err error) {
    return c.DeleteAigcApiTokenWithContext(context.Background(), request)
}

// DeleteAigcApiToken
// Deletes an AIGC API Token. The associated AIGC quota will also be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteAigcApiTokenWithContext(ctx context.Context, request *DeleteAigcApiTokenRequest) (response *DeleteAigcApiTokenResponse, err error) {
    if request == nil {
        request = NewDeleteAigcApiTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteAigcApiToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAigcApiToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAigcApiTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAigcQuotaRequest() (request *DeleteAigcQuotaRequest) {
    request = &DeleteAigcQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteAigcQuota")
    
    
    return
}

func NewDeleteAigcQuotaResponse() (response *DeleteAigcQuotaResponse) {
    response = &DeleteAigcQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAigcQuota
// This API is used to delete AIGC quota configurations. Once deleted, AIGC task initiation will no longer be limited.
//
// 
//
// If the quota is re-enabled after deletion, the amount will be cleared and recalculated.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteAigcQuota(request *DeleteAigcQuotaRequest) (response *DeleteAigcQuotaResponse, err error) {
    return c.DeleteAigcQuotaWithContext(context.Background(), request)
}

// DeleteAigcQuota
// This API is used to delete AIGC quota configurations. Once deleted, AIGC task initiation will no longer be limited.
//
// 
//
// If the quota is re-enabled after deletion, the amount will be cleared and recalculated.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteAigcQuotaWithContext(ctx context.Context, request *DeleteAigcQuotaRequest) (response *DeleteAigcQuotaResponse, err error) {
    if request == nil {
        request = NewDeleteAigcQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteAigcQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAigcQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAigcQuotaResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteAnimatedGraphicsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlindWatermarkTemplateRequest() (request *DeleteBlindWatermarkTemplateRequest) {
    request = &DeleteBlindWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteBlindWatermarkTemplate")
    
    
    return
}

func NewDeleteBlindWatermarkTemplateResponse() (response *DeleteBlindWatermarkTemplateResponse) {
    response = &DeleteBlindWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBlindWatermarkTemplate
// This API is used to delete a user-defined digital watermark template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBlindWatermarkTemplate(request *DeleteBlindWatermarkTemplateRequest) (response *DeleteBlindWatermarkTemplateResponse, err error) {
    return c.DeleteBlindWatermarkTemplateWithContext(context.Background(), request)
}

// DeleteBlindWatermarkTemplate
// This API is used to delete a user-defined digital watermark template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBlindWatermarkTemplateWithContext(ctx context.Context, request *DeleteBlindWatermarkTemplateRequest) (response *DeleteBlindWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteBlindWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteBlindWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBlindWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBlindWatermarkTemplateResponse()
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
// Delete a CDN domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCDNDomain(request *DeleteCDNDomainRequest) (response *DeleteCDNDomainResponse, err error) {
    return c.DeleteCDNDomainWithContext(context.Background(), request)
}

// DeleteCDNDomain
// Delete a CDN domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCDNDomainWithContext(ctx context.Context, request *DeleteCDNDomainRequest) (response *DeleteCDNDomainResponse, err error) {
    if request == nil {
        request = NewDeleteCDNDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteCDNDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCDNDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCDNDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCLSTopicRequest() (request *DeleteCLSTopicRequest) {
    request = &DeleteCLSTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteCLSTopic")
    
    
    return
}

func NewDeleteCLSTopicResponse() (response *DeleteCLSTopicResponse) {
    response = &DeleteCLSTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCLSTopic
// Delete the log topic enabled for VOD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCLSTopic(request *DeleteCLSTopicRequest) (response *DeleteCLSTopicResponse, err error) {
    return c.DeleteCLSTopicWithContext(context.Background(), request)
}

// DeleteCLSTopic
// Delete the log topic enabled for VOD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCLSTopicWithContext(ctx context.Context, request *DeleteCLSTopicRequest) (response *DeleteCLSTopicResponse, err error) {
    if request == nil {
        request = NewDeleteCLSTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteCLSTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCLSTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCLSTopicResponse()
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
// * A category can be deleted only when it has no subcategories and no associated media.
//
// * Otherwise, execute [delete media](https://www.tencentcloud.com/document/product/266/31764?from_cn_redirect=1) and subcategories first, then delete the category;
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
// * A category can be deleted only when it has no subcategories and no associated media.
//
// * Otherwise, execute [delete media](https://www.tencentcloud.com/document/product/266/31764?from_cn_redirect=1) and subcategories first, then delete the category;
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteClass")
    
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
// This API is <font color=red>no longer maintained</font>. The new version moderation template supports video moderation and image moderation. For details, please see [Delete Moderation Template](https://www.tencentcloud.com/document/api/266/84390?from_cn_redirect=1).
//
// Delete a user-customized audio/video moderation template.
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
// This API is <font color=red>no longer maintained</font>. The new version moderation template supports video moderation and image moderation. For details, please see [Delete Moderation Template](https://www.tencentcloud.com/document/api/266/84390?from_cn_redirect=1).
//
// Delete a user-customized audio/video moderation template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteContentReviewTemplate")
    
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
// This API is <font color=red>no longer maintained</font>. The new version of [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Delete an audio and video quality rebirth template.
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
// This API is <font color=red>no longer maintained</font>. The new version of [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Delete an audio and video quality rebirth template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteEnhanceMediaTemplate")
    
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
// Delete a title and trailer template.
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
// Delete a title and trailer template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteHeadTailTemplate")
    
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
// This API is used to delete a user-defined image processing template.
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
// This API is used to delete a user-defined image processing template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteImageProcessingTemplate")
    
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
// Delete an image sprite template.
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
// Delete an image sprite template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteImageSpriteTemplate")
    
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
// This API is used to delete a just in time transcoding template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteJustInTimeTranscodeTemplate(request *DeleteJustInTimeTranscodeTemplateRequest) (response *DeleteJustInTimeTranscodeTemplateResponse, err error) {
    return c.DeleteJustInTimeTranscodeTemplateWithContext(context.Background(), request)
}

// DeleteJustInTimeTranscodeTemplate
// This API is used to delete a just in time transcoding template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteJustInTimeTranscodeTemplateWithContext(ctx context.Context, request *DeleteJustInTimeTranscodeTemplateRequest) (response *DeleteJustInTimeTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteJustInTimeTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteJustInTimeTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJustInTimeTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJustInTimeTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKnowledgeBaseRequest() (request *DeleteKnowledgeBaseRequest) {
    request = &DeleteKnowledgeBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteKnowledgeBase")
    
    
    return
}

func NewDeleteKnowledgeBaseResponse() (response *DeleteKnowledgeBaseResponse) {
    response = &DeleteKnowledgeBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteKnowledgeBase
// Delete a knowledge base.
//
// After the API is called, the knowledge base will be in the "Deleting" status and the deletion operation will be performed in the backend.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteKnowledgeBase(request *DeleteKnowledgeBaseRequest) (response *DeleteKnowledgeBaseResponse, err error) {
    return c.DeleteKnowledgeBaseWithContext(context.Background(), request)
}

// DeleteKnowledgeBase
// Delete a knowledge base.
//
// After the API is called, the knowledge base will be in the "Deleting" status and the deletion operation will be performed in the backend.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteKnowledgeBaseWithContext(ctx context.Context, request *DeleteKnowledgeBaseRequest) (response *DeleteKnowledgeBaseResponse, err error) {
    if request == nil {
        request = NewDeleteKnowledgeBaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteKnowledgeBase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKnowledgeBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKnowledgeBaseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLLMComprehendTemplateRequest() (request *DeleteLLMComprehendTemplateRequest) {
    request = &DeleteLLMComprehendTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteLLMComprehendTemplate")
    
    
    return
}

func NewDeleteLLMComprehendTemplateResponse() (response *DeleteLLMComprehendTemplateResponse) {
    response = &DeleteLLMComprehendTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLLMComprehendTemplate
// This API is used to delete a user-defined customized large model parsing template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLLMComprehendTemplate(request *DeleteLLMComprehendTemplateRequest) (response *DeleteLLMComprehendTemplateResponse, err error) {
    return c.DeleteLLMComprehendTemplateWithContext(context.Background(), request)
}

// DeleteLLMComprehendTemplate
// This API is used to delete a user-defined customized large model parsing template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLLMComprehendTemplateWithContext(ctx context.Context, request *DeleteLLMComprehendTemplateRequest) (response *DeleteLLMComprehendTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLLMComprehendTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteLLMComprehendTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLLMComprehendTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLLMComprehendTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMPSTemplateRequest() (request *DeleteMPSTemplateRequest) {
    request = &DeleteMPSTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteMPSTemplate")
    
    
    return
}

func NewDeleteMPSTemplateResponse() (response *DeleteMPSTemplateResponse) {
    response = &DeleteMPSTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMPSTemplate
// This API is used to delete a user-defined MPS task template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteMPSTemplate(request *DeleteMPSTemplateRequest) (response *DeleteMPSTemplateResponse, err error) {
    return c.DeleteMPSTemplateWithContext(context.Background(), request)
}

// DeleteMPSTemplate
// This API is used to delete a user-defined MPS task template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteMPSTemplateWithContext(ctx context.Context, request *DeleteMPSTemplateRequest) (response *DeleteMPSTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteMPSTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteMPSTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMPSTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMPSTemplateResponse()
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
// * Delete media and its corresponding video processing files (raw files, such as transcoded videos, sprite sheets, screenshots, WeChat video releases, etc.);
//
// * You can separately delete the source file, transcoded video, and WeChat-published video under a specified video file ID.
//
// * Note: After the original file is deleted, you cannot initiate any video processing operation such as transcoding or publishing on WeChat.
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
// * Delete media and its corresponding video processing files (raw files, such as transcoded videos, sprite sheets, screenshots, WeChat video releases, etc.);
//
// * You can separately delete the source file, transcoded video, and WeChat-published video under a specified video file ID.
//
// * Note: After the original file is deleted, you cannot initiate any video processing operation such as transcoding or publishing on WeChat.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteMedia")
    
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
// This API is used to delete material samples based on character ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePersonSample(request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    return c.DeletePersonSampleWithContext(context.Background(), request)
}

// DeletePersonSample
// This API is used to delete material samples based on character ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePersonSampleWithContext(ctx context.Context, request *DeletePersonSampleRequest) (response *DeletePersonSampleResponse, err error) {
    if request == nil {
        request = NewDeletePersonSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeletePersonSample")
    
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
// This API is used to delete a user-defined task flow template.
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
// This API is used to delete a user-defined task flow template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteProcedureTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProcedureTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProcedureTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProcessImageAsyncTemplateRequest() (request *DeleteProcessImageAsyncTemplateRequest) {
    request = &DeleteProcessImageAsyncTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteProcessImageAsyncTemplate")
    
    
    return
}

func NewDeleteProcessImageAsyncTemplateResponse() (response *DeleteProcessImageAsyncTemplateResponse) {
    response = &DeleteProcessImageAsyncTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProcessImageAsyncTemplate
// This API is used to delete a user-customized asynchronous image processing template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteProcessImageAsyncTemplate(request *DeleteProcessImageAsyncTemplateRequest) (response *DeleteProcessImageAsyncTemplateResponse, err error) {
    return c.DeleteProcessImageAsyncTemplateWithContext(context.Background(), request)
}

// DeleteProcessImageAsyncTemplate
// This API is used to delete a user-customized asynchronous image processing template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteProcessImageAsyncTemplateWithContext(ctx context.Context, request *DeleteProcessImageAsyncTemplateRequest) (response *DeleteProcessImageAsyncTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteProcessImageAsyncTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteProcessImageAsyncTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProcessImageAsyncTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProcessImageAsyncTemplateResponse()
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
// This API is used to delete an audio-visual quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteQualityInspectTemplate(request *DeleteQualityInspectTemplateRequest) (response *DeleteQualityInspectTemplateResponse, err error) {
    return c.DeleteQualityInspectTemplateWithContext(context.Background(), request)
}

// DeleteQualityInspectTemplate
// This API is used to delete an audio-visual quality inspection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteQualityInspectTemplateWithContext(ctx context.Context, request *DeleteQualityInspectTemplateRequest) (response *DeleteQualityInspectTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteQualityInspectTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteQualityInspectTemplate")
    
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
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Delete a video rebirth template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteRebuildMediaTemplate(request *DeleteRebuildMediaTemplateRequest) (response *DeleteRebuildMediaTemplateResponse, err error) {
    return c.DeleteRebuildMediaTemplateWithContext(context.Background(), request)
}

// DeleteRebuildMediaTemplate
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Delete a video rebirth template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteRebuildMediaTemplateWithContext(ctx context.Context, request *DeleteRebuildMediaTemplateRequest) (response *DeleteRebuildMediaTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteRebuildMediaTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteRebuildMediaTemplate")
    
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
// This API is used to delete a user-defined moderation template.
//
// >Template is applicable only to the [audio/video moderation (ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) and [image moderation (ReviewImage)](https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) APIs.
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
// This API is used to delete a user-defined moderation template.
//
// >Template is applicable only to the [audio/video moderation (ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) and [image moderation (ReviewImage)](https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) APIs.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteReviewTemplate")
    
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
// This API is used to delete a carousel playlist.
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
// This API is used to delete a carousel playlist.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteRoundPlay")
    
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
// This API is used to delete a user-customized sampled screenshot template.
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
// This API is used to delete a user-customized sampled screenshot template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteSampleSnapshotTemplate")
    
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
// This API is used to delete a user-customized specified time point screenshot template.
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
// This API is used to delete a user-customized specified time point screenshot template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteSnapshotByTimeOffsetTemplate")
    
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
// This API is <font color='red'>no longer maintained</font>. The new version of player signature no longer uses player configuration templates. For details, please see [Player Signature](https://www.tencentcloud.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to delete player configurations.  
//
// *Note: The system preset player configuration cannot be deleted.*
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
// This API is <font color='red'>no longer maintained</font>. The new version of player signature no longer uses player configuration templates. For details, please see [Player Signature](https://www.tencentcloud.com/document/product/266/45554?from_cn_redirect=1).
//
// This API is used to delete player configurations.  
//
// *Note: The system preset player configuration cannot be deleted.*
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteSuperPlayerConfig")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteTranscodeTemplate")
    
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
// This API is used to delete VOD acceleration domains.
//
// 1. Acceleration in all regions must be disabled before domain deletion.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteVodDomain(request *DeleteVodDomainRequest) (response *DeleteVodDomainResponse, err error) {
    return c.DeleteVodDomainWithContext(context.Background(), request)
}

// DeleteVodDomain
// This API is used to delete VOD acceleration domains.
//
// 1. Acceleration in all regions must be disabled before domain deletion.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteVodDomainWithContext(ctx context.Context, request *DeleteVodDomainRequest) (response *DeleteVodDomainResponse, err error) {
    if request == nil {
        request = NewDeleteVodDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteVodDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVodDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVodDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVoiceRequest() (request *DeleteVoiceRequest) {
    request = &DeleteVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteVoice")
    
    
    return
}

func NewDeleteVoiceResponse() (response *DeleteVoiceResponse) {
    response = &DeleteVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVoice
// Deletes a specified voice type by voice ID. The deletion cannot be undone, and the voice type cannot be used for subsequent APIs. Only voice types for this account can be deleted. System preset voice types cannot be deleted.
//
// 
//
// Note: Newly designed or cloned voice types cannot be deleted before activation. They are activated only after the new voice type is used for TTS once.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteVoice(request *DeleteVoiceRequest) (response *DeleteVoiceResponse, err error) {
    return c.DeleteVoiceWithContext(context.Background(), request)
}

// DeleteVoice
// Deletes a specified voice type by voice ID. The deletion cannot be undone, and the voice type cannot be used for subsequent APIs. Only voice types for this account can be deleted. System preset voice types cannot be deleted.
//
// 
//
// Note: Newly designed or cloned voice types cannot be deleted before activation. They are activated only after the new voice type is used for TTS once.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteVoiceWithContext(ctx context.Context, request *DeleteVoiceRequest) (response *DeleteVoiceResponse, err error) {
    if request == nil {
        request = NewDeleteVoiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteVoice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVoiceResponse()
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
// This API is used to delete a user-defined watermark template.
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
// This API is used to delete a user-defined watermark template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteWatermarkTemplate")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DeleteWordSamples")
    
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
// This API is used to retrieve the audio/video content analysis template detail list based on the unique identifier of an audio/video content analysis template. The returned results include all eligible user-defined audio/video content analysis templates and [system preset audio/video content analysis templates](https://www.tencentcloud.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.88.86.E6.9E.90.E6.A8.A1.E6.9D.BF).
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
// This API is used to retrieve the audio/video content analysis template detail list based on the unique identifier of an audio/video content analysis template. The returned results include all eligible user-defined audio/video content analysis templates and [system preset audio/video content analysis templates](https://www.tencentcloud.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.88.86.E6.9E.90.E6.A8.A1.E6.9D.BF).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAIAnalysisTemplates")
    
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
// This API is used to get the list of audio/video content recognition templates by unique ID. The returned results include all eligible user-defined audio/video content recognition templates and [system preset audio/video content recognition templates](https://www.tencentcloud.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E8.AF.86.E5.88.AB.E6.A8.A1.E6.9D.BF).
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
// This API is used to get the list of audio/video content recognition templates by unique ID. The returned results include all eligible user-defined audio/video content recognition templates and [system preset audio/video content recognition templates](https://www.tencentcloud.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E8.AF.86.E5.88.AB.E6.A8.A1.E6.9D.BF).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAIRecognitionTemplates")
    
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
// This API is used to query adaptive bitrate streaming templates, and the pagination query is supported based on conditions.
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
// This API is used to query adaptive bitrate streaming templates, and the pagination query is supported based on conditions.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAdaptiveDynamicStreamingTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAdaptiveDynamicStreamingTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAdaptiveDynamicStreamingTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcAdvancedCustomElementsRequest() (request *DescribeAigcAdvancedCustomElementsRequest) {
    request = &DescribeAigcAdvancedCustomElementsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAigcAdvancedCustomElements")
    
    
    return
}

func NewDescribeAigcAdvancedCustomElementsResponse() (response *DescribeAigcAdvancedCustomElementsResponse) {
    response = &DescribeAigcAdvancedCustomElementsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcAdvancedCustomElements
// This API is used to retrieve advanced custom AIGC subjects.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
func (c *Client) DescribeAigcAdvancedCustomElements(request *DescribeAigcAdvancedCustomElementsRequest) (response *DescribeAigcAdvancedCustomElementsResponse, err error) {
    return c.DescribeAigcAdvancedCustomElementsWithContext(context.Background(), request)
}

// DescribeAigcAdvancedCustomElements
// This API is used to retrieve advanced custom AIGC subjects.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
func (c *Client) DescribeAigcAdvancedCustomElementsWithContext(ctx context.Context, request *DescribeAigcAdvancedCustomElementsRequest) (response *DescribeAigcAdvancedCustomElementsResponse, err error) {
    if request == nil {
        request = NewDescribeAigcAdvancedCustomElementsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAigcAdvancedCustomElements")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcAdvancedCustomElements require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcAdvancedCustomElementsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcApiTokensRequest() (request *DescribeAigcApiTokensRequest) {
    request = &DescribeAigcApiTokensRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAigcApiTokens")
    
    
    return
}

func NewDescribeAigcApiTokensResponse() (response *DescribeAigcApiTokensResponse) {
    response = &DescribeAigcApiTokensResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcApiTokens
// Query the list of AIGC API tokens. Data sync has a delay after creation or deletion. You can query the latest data after about 30 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeAigcApiTokens(request *DescribeAigcApiTokensRequest) (response *DescribeAigcApiTokensResponse, err error) {
    return c.DescribeAigcApiTokensWithContext(context.Background(), request)
}

// DescribeAigcApiTokens
// Query the list of AIGC API tokens. Data sync has a delay after creation or deletion. You can query the latest data after about 30 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeAigcApiTokensWithContext(ctx context.Context, request *DescribeAigcApiTokensRequest) (response *DescribeAigcApiTokensResponse, err error) {
    if request == nil {
        request = NewDescribeAigcApiTokensRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAigcApiTokens")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcApiTokens require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcApiTokensResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcFaceInfoRequest() (request *DescribeAigcFaceInfoRequest) {
    request = &DescribeAigcFaceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAigcFaceInfo")
    
    
    return
}

func NewDescribeAigcFaceInfoResponse() (response *DescribeAigcFaceInfoResponse) {
    response = &DescribeAigcFaceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcFaceInfo
// This API is used to retrieve AIGC face information. Note that calling this API incurs face recognition fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#96b3b59a-f9e1-49e9-966a-bedb70a4bf12).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAigcFaceInfo(request *DescribeAigcFaceInfoRequest) (response *DescribeAigcFaceInfoResponse, err error) {
    return c.DescribeAigcFaceInfoWithContext(context.Background(), request)
}

// DescribeAigcFaceInfo
// This API is used to retrieve AIGC face information. Note that calling this API incurs face recognition fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#96b3b59a-f9e1-49e9-966a-bedb70a4bf12).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAigcFaceInfoWithContext(ctx context.Context, request *DescribeAigcFaceInfoRequest) (response *DescribeAigcFaceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAigcFaceInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAigcFaceInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcFaceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcFaceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcFaceInfoAsyncRequest() (request *DescribeAigcFaceInfoAsyncRequest) {
    request = &DescribeAigcFaceInfoAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAigcFaceInfoAsync")
    
    
    return
}

func NewDescribeAigcFaceInfoAsyncResponse() (response *DescribeAigcFaceInfoAsyncResponse) {
    response = &DescribeAigcFaceInfoAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcFaceInfoAsync
// This API is used to asynchronously fetch AIGC face information. Note that calling this API incurs face recognition fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#96b3b59a-f9e1-49e9-966a-bedb70a4bf12).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAigcFaceInfoAsync(request *DescribeAigcFaceInfoAsyncRequest) (response *DescribeAigcFaceInfoAsyncResponse, err error) {
    return c.DescribeAigcFaceInfoAsyncWithContext(context.Background(), request)
}

// DescribeAigcFaceInfoAsync
// This API is used to asynchronously fetch AIGC face information. Note that calling this API incurs face recognition fees. Refer to the billing documentation (https://www.tencentcloud.com/document/product/266/95125?from_cn_redirect=1#96b3b59a-f9e1-49e9-966a-bedb70a4bf12).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAigcFaceInfoAsyncWithContext(ctx context.Context, request *DescribeAigcFaceInfoAsyncRequest) (response *DescribeAigcFaceInfoAsyncResponse, err error) {
    if request == nil {
        request = NewDescribeAigcFaceInfoAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAigcFaceInfoAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcFaceInfoAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcFaceInfoAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcQuotasRequest() (request *DescribeAigcQuotasRequest) {
    request = &DescribeAigcQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAigcQuotas")
    
    
    return
}

func NewDescribeAigcQuotasResponse() (response *DescribeAigcQuotasResponse) {
    response = &DescribeAigcQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcQuotas
// This API is used to query AIGC quota configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeAigcQuotas(request *DescribeAigcQuotasRequest) (response *DescribeAigcQuotasResponse, err error) {
    return c.DescribeAigcQuotasWithContext(context.Background(), request)
}

// DescribeAigcQuotas
// This API is used to query AIGC quota configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeAigcQuotasWithContext(ctx context.Context, request *DescribeAigcQuotasRequest) (response *DescribeAigcQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeAigcQuotasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAigcQuotas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcQuotas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcUsageDataRequest() (request *DescribeAigcUsageDataRequest) {
    request = &DescribeAigcUsageDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeAigcUsageData")
    
    
    return
}

func NewDescribeAigcUsageDataResponse() (response *DescribeAigcUsageDataResponse) {
    response = &DescribeAigcUsageDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcUsageData
// This API returns AIGC statistical information within a specified time range.
//
// 1. AIGC statistical data from the last 365 days can be queried.
//
//    2. The query time span should not exceed 90 days.
//
// 3. If the query time span exceeds 1 day, the data of day granularity is returned. Otherwise, the data of 5-minute granularity is returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAigcUsageData(request *DescribeAigcUsageDataRequest) (response *DescribeAigcUsageDataResponse, err error) {
    return c.DescribeAigcUsageDataWithContext(context.Background(), request)
}

// DescribeAigcUsageData
// This API returns AIGC statistical information within a specified time range.
//
// 1. AIGC statistical data from the last 365 days can be queried.
//
//    2. The query time span should not exceed 90 days.
//
// 3. If the query time span exceeds 1 day, the data of day granularity is returned. Otherwise, the data of 5-minute granularity is returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAigcUsageDataWithContext(ctx context.Context, request *DescribeAigcUsageDataRequest) (response *DescribeAigcUsageDataResponse, err error) {
    if request == nil {
        request = NewDescribeAigcUsageDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAigcUsageData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcUsageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcUsageDataResponse()
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
// * Obtain all classification information of the user.
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
// * Obtain all classification information of the user.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAllClass")
    
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
// This API is used to query a list of rotating image templates, and the pagination query is supported based on conditions.
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
// This API is used to query a list of rotating image templates, and the pagination query is supported based on conditions.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeAnimatedGraphicsTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAnimatedGraphicsTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAnimatedGraphicsTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlindWatermarkTemplatesRequest() (request *DescribeBlindWatermarkTemplatesRequest) {
    request = &DescribeBlindWatermarkTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeBlindWatermarkTemplates")
    
    
    return
}

func NewDescribeBlindWatermarkTemplatesResponse() (response *DescribeBlindWatermarkTemplatesResponse) {
    response = &DescribeBlindWatermarkTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlindWatermarkTemplates
// This API is used to query user-customized digital watermark templates.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeBlindWatermarkTemplates(request *DescribeBlindWatermarkTemplatesRequest) (response *DescribeBlindWatermarkTemplatesResponse, err error) {
    return c.DescribeBlindWatermarkTemplatesWithContext(context.Background(), request)
}

// DescribeBlindWatermarkTemplates
// This API is used to query user-customized digital watermark templates.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeBlindWatermarkTemplatesWithContext(ctx context.Context, request *DescribeBlindWatermarkTemplatesRequest) (response *DescribeBlindWatermarkTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeBlindWatermarkTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeBlindWatermarkTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlindWatermarkTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlindWatermarkTemplatesResponse()
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
// This API is used to query on-demand domain names.
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
// This API is used to query on-demand domain names.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeCDNDomains")
    
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
// This API is used to query CDN bandwidth, traffic, and other stats for on-demand domain names.
//
// * The time span between the query start time and end time should not exceed 90 days.
//
// * Data in different service regions can be queried.
//
// * Statistical data within the Chinese mainland supports querying stats by designated region and carrier.
//
// * Playback statistics only target VOD domains, excluding distribution from EdgeOne domain names.
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
// This API is used to query CDN bandwidth, traffic, and other stats for on-demand domain names.
//
// * The time span between the query start time and end time should not exceed 90 days.
//
// * Data in different service regions can be queried.
//
// * Statistical data within the Chinese mainland supports querying stats by designated region and carrier.
//
// * Playback statistics only target VOD domains, excluding distribution from EdgeOne domain names.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeCDNStatDetails")
    
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
// This API is used to query VOD CDN stats such as traffic and bandwidth.
//
// 1. CDN usage data is retained on the system side for 13 months. You can only query usage data from the most recent 365 days through the API. If you need to retrieve historical usage data beyond 365 days, contact us.
//
//    2. The query time span should not exceed 90 days.
//
// 3. You can specify the time granularity of usage data, supporting 5-minute, 1-hour, and 1-day granularities.
//
// 4. Traffic is the total traffic within the query time granularity, and bandwidth is the peak bandwidth within the query time granularity.
//
// 5. Playback statistics only target VOD domains, excluding distribution from EdgeOne domain names.
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
// This API is used to query VOD CDN stats such as traffic and bandwidth.
//
// 1. CDN usage data is retained on the system side for 13 months. You can only query usage data from the most recent 365 days through the API. If you need to retrieve historical usage data beyond 365 days, contact us.
//
//    2. The query time span should not exceed 90 days.
//
// 3. You can specify the time granularity of usage data, supporting 5-minute, 1-hour, and 1-day granularities.
//
// 4. Traffic is the total traffic within the query time granularity, and bandwidth is the peak bandwidth within the query time granularity.
//
// 5. Playback statistics only target VOD domains, excluding distribution from EdgeOne domain names.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeCDNUsageData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCDNUsageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCDNUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCLSLogsetsRequest() (request *DescribeCLSLogsetsRequest) {
    request = &DescribeCLSLogsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCLSLogsets")
    
    
    return
}

func NewDescribeCLSLogsetsResponse() (response *DescribeCLSLogsetsResponse) {
    response = &DescribeCLSLogsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCLSLogsets
// Query the CLS log set created by VOD.
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
func (c *Client) DescribeCLSLogsets(request *DescribeCLSLogsetsRequest) (response *DescribeCLSLogsetsResponse, err error) {
    return c.DescribeCLSLogsetsWithContext(context.Background(), request)
}

// DescribeCLSLogsets
// Query the CLS log set created by VOD.
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
func (c *Client) DescribeCLSLogsetsWithContext(ctx context.Context, request *DescribeCLSLogsetsRequest) (response *DescribeCLSLogsetsResponse, err error) {
    if request == nil {
        request = NewDescribeCLSLogsetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeCLSLogsets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCLSLogsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCLSLogsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCLSPushTargetsRequest() (request *DescribeCLSPushTargetsRequest) {
    request = &DescribeCLSPushTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCLSPushTargets")
    
    
    return
}

func NewDescribeCLSPushTargetsResponse() (response *DescribeCLSPushTargetsResponse) {
    response = &DescribeCLSPushTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCLSPushTargets
// Queries the destination topic for log delivery under an on-demand domain name.
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
func (c *Client) DescribeCLSPushTargets(request *DescribeCLSPushTargetsRequest) (response *DescribeCLSPushTargetsResponse, err error) {
    return c.DescribeCLSPushTargetsWithContext(context.Background(), request)
}

// DescribeCLSPushTargets
// Queries the destination topic for log delivery under an on-demand domain name.
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
func (c *Client) DescribeCLSPushTargetsWithContext(ctx context.Context, request *DescribeCLSPushTargetsRequest) (response *DescribeCLSPushTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeCLSPushTargetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeCLSPushTargets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCLSPushTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCLSPushTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCLSTopicsRequest() (request *DescribeCLSTopicsRequest) {
    request = &DescribeCLSTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeCLSTopics")
    
    
    return
}

func NewDescribeCLSTopicsResponse() (response *DescribeCLSTopicsResponse) {
    response = &DescribeCLSTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCLSTopics
// Queries the list of CLS log topics created by VOD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCLSTopics(request *DescribeCLSTopicsRequest) (response *DescribeCLSTopicsResponse, err error) {
    return c.DescribeCLSTopicsWithContext(context.Background(), request)
}

// DescribeCLSTopics
// Queries the list of CLS log topics created by VOD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCLSTopicsWithContext(ctx context.Context, request *DescribeCLSTopicsRequest) (response *DescribeCLSTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeCLSTopicsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeCLSTopics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCLSTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCLSTopicsResponse()
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
// This API is used to query the download URL of CDN access logs for an on-demand domain name, excluding EdgeOne origin-pull to VOD domains.
//
// 1. Can query CDN log download links from the most recent 30 days.
//
// 2. By default, CDN generates a log file per hour. If there is no CDN access in an hour, no log file is generated.    
//
// 3. The CDN log download link has a validity of 24 hours.
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
// This API is used to query the download URL of CDN access logs for an on-demand domain name, excluding EdgeOne origin-pull to VOD domains.
//
// 1. Can query CDN log download links from the most recent 30 days.
//
// 2. By default, CDN generates a log file per hour. If there is no CDN access in an hour, no log file is generated.    
//
// 3. The CDN log download link has a validity of 24 hours.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeCdnLogs")
    
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
// This API is used to return client upload acceleration statistics within a specified time range.
//
// 1. Can query client upload acceleration statistics data for the most recent 365 days.
//
//    2. The query time span should not exceed 90 days.
//
// 3. If the query time span exceeds 1 day, the data returned is at a daily granularity. Otherwise, the data returned is at a 5-minute granularity.
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
// This API is used to return client upload acceleration statistics within a specified time range.
//
// 1. Can query client upload acceleration statistics data for the most recent 365 days.
//
//    2. The query time span should not exceed 90 days.
//
// 3. If the query time span exceeds 1 day, the data returned is at a daily granularity. Otherwise, the data returned is at a 5-minute granularity.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeClientUploadAccelerationUsageData")
    
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
// This API is <font color=red>no longer maintained</font>. The new version of the moderation template supports video and image moderation. For details, please see [Query the moderation template list](https://www.tencentcloud.com/document/api/266/84389?from_cn_redirect=1).
//
// This API is used to obtain the list of audio/video moderation template details based on the unique identifier of the template. The returned results include all eligible custom templates and system preset content review templates (https://www.tencentcloud.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.AE.A1.E6.A0.B8.E6.A8.A1.E6.9D.BF).
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
// This API is <font color=red>no longer maintained</font>. The new version of the moderation template supports video and image moderation. For details, please see [Query the moderation template list](https://www.tencentcloud.com/document/api/266/84389?from_cn_redirect=1).
//
// This API is used to obtain the list of audio/video moderation template details based on the unique identifier of the template. The returned results include all eligible custom templates and system preset content review templates (https://www.tencentcloud.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.A7.86.E9.A2.91.E5.86.85.E5.AE.B9.E5.AE.A1.E6.A0.B8.E6.A8.A1.E6.9D.BF).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeContentReviewTemplates")
    
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
// Query the carousel current playlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ROUNDPLAYID = "InvalidParameterValue.RoundPlayId"
func (c *Client) DescribeCurrentPlaylist(request *DescribeCurrentPlaylistRequest) (response *DescribeCurrentPlaylistResponse, err error) {
    return c.DescribeCurrentPlaylistWithContext(context.Background(), request)
}

// DescribeCurrentPlaylist
// Query the carousel current playlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ROUNDPLAYID = "InvalidParameterValue.RoundPlayId"
func (c *Client) DescribeCurrentPlaylistWithContext(ctx context.Context, request *DescribeCurrentPlaylistRequest) (response *DescribeCurrentPlaylistResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentPlaylistRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeCurrentPlaylist")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCurrentPlaylist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCurrentPlaylistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDailyMediaPlayStatRequest() (request *DescribeDailyMediaPlayStatRequest) {
    request = &DescribeDailyMediaPlayStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeDailyMediaPlayStat")
    
    
    return
}

func NewDescribeDailyMediaPlayStatResponse() (response *DescribeDailyMediaPlayStatResponse) {
    response = &DescribeDailyMediaPlayStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDailyMediaPlayStat
// This API is used to query daily playback statistics within a specified date range.
//
// Playback statistics from the past one year can be queried.
//
// * The time span between the start date and end date can be up to 90 days.
//
// * Playback statistics only target VOD domains, excluding distribution from EdgeOne domain names.
//
// * Due to data delay, you are advised to query the usage data of the previous day after 12:00 noon the next day.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_ENDDATE = "InvalidParameterValue.EndDate"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_STARTDATE = "InvalidParameterValue.StartDate"
func (c *Client) DescribeDailyMediaPlayStat(request *DescribeDailyMediaPlayStatRequest) (response *DescribeDailyMediaPlayStatResponse, err error) {
    return c.DescribeDailyMediaPlayStatWithContext(context.Background(), request)
}

// DescribeDailyMediaPlayStat
// This API is used to query daily playback statistics within a specified date range.
//
// Playback statistics from the past one year can be queried.
//
// * The time span between the start date and end date can be up to 90 days.
//
// * Playback statistics only target VOD domains, excluding distribution from EdgeOne domain names.
//
// * Due to data delay, you are advised to query the usage data of the previous day after 12:00 noon the next day.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_ENDDATE = "InvalidParameterValue.EndDate"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_STARTDATE = "InvalidParameterValue.StartDate"
func (c *Client) DescribeDailyMediaPlayStatWithContext(ctx context.Context, request *DescribeDailyMediaPlayStatRequest) (response *DescribeDailyMediaPlayStatResponse, err error) {
    if request == nil {
        request = NewDescribeDailyMediaPlayStatRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeDailyMediaPlayStat")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDailyMediaPlayStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDailyMediaPlayStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDailyMostPlayedStatRequest() (request *DescribeDailyMostPlayedStatRequest) {
    request = &DescribeDailyMostPlayedStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeDailyMostPlayedStat")
    
    
    return
}

func NewDescribeDailyMostPlayedStatResponse() (response *DescribeDailyMostPlayedStatResponse) {
    response = &DescribeDailyMostPlayedStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDailyMostPlayedStat
// This API is used to query daily playback statistics for the Top 100 media files.
//
// * Playback statistics from the past one year can be queried.
//
// * You can query by number of plays or playback traffic.
//
// * Playback count statistics description:
//
// 1. HLS file: The number of plays is counted when accessing M3U8 files, but not when accessing TS files.
//
// 2. Other files (for example, MP4 files): If a playback request includes the range parameter and the start parameter of range is not equal to 0, the number of plays is not counted. In other cases, the number of plays is counted.
//
// * Playback statistics only target VOD domains (i.e., EdgeOne domain distribution is not included in playback statistics).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetWorkError"
//  INVALIDPARAMETERVALUE_DATE = "InvalidParameterValue.Date"
//  INVALIDPARAMETERVALUE_DOMAINNAME = "InvalidParameterValue.DomainName"
//  INVALIDPARAMETERVALUE_METRIC = "InvalidParameterValue.Metric"
func (c *Client) DescribeDailyMostPlayedStat(request *DescribeDailyMostPlayedStatRequest) (response *DescribeDailyMostPlayedStatResponse, err error) {
    return c.DescribeDailyMostPlayedStatWithContext(context.Background(), request)
}

// DescribeDailyMostPlayedStat
// This API is used to query daily playback statistics for the Top 100 media files.
//
// * Playback statistics from the past one year can be queried.
//
// * You can query by number of plays or playback traffic.
//
// * Playback count statistics description:
//
// 1. HLS file: The number of plays is counted when accessing M3U8 files, but not when accessing TS files.
//
// 2. Other files (for example, MP4 files): If a playback request includes the range parameter and the start parameter of range is not equal to 0, the number of plays is not counted. In other cases, the number of plays is counted.
//
// * Playback statistics only target VOD domains (i.e., EdgeOne domain distribution is not included in playback statistics).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetWorkError"
//  INVALIDPARAMETERVALUE_DATE = "InvalidParameterValue.Date"
//  INVALIDPARAMETERVALUE_DOMAINNAME = "InvalidParameterValue.DomainName"
//  INVALIDPARAMETERVALUE_METRIC = "InvalidParameterValue.Metric"
func (c *Client) DescribeDailyMostPlayedStatWithContext(ctx context.Context, request *DescribeDailyMostPlayedStatRequest) (response *DescribeDailyMostPlayedStatResponse, err error) {
    if request == nil {
        request = NewDescribeDailyMostPlayedStatRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeDailyMostPlayedStat")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDailyMostPlayedStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDailyMostPlayedStatResponse()
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
// This API is used to query the download address of playback statistics files.
//
// * You can query the download links for playback statistics files from the past one year. The time span between the start date and end date cannot exceed 90 days.
//
// VOD analyzes and processes the CDN request logs of the previous day to generate playback statistics files.
//
// * The playback statistics file contains statistical information such as the number of plays and total traffic of media files.
//
// * Playback count statistics description:
//
// 1. HLS file: The number of plays is counted when an M3U8 file is accessed, but not when a TS file is accessed.
//
// 2. Other files (for example, MP4 files): If the playback request includes the range parameter and the start parameter of range is not equal to 0, the number of plays is not counted. In other cases, the number of plays is counted.
//
// * Statistics of playback devices: If a playback request includes the UserAgent parameter and the UserAgent contains identifiers such as Android or iPhone, it is counted as a mobile playback count. Otherwise, it is counted as a PC playback count.
//
// * Playback statistics only target VOD domains, excluding EdgeOne domain name distribution.
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
// This API is used to query the download address of playback statistics files.
//
// * You can query the download links for playback statistics files from the past one year. The time span between the start date and end date cannot exceed 90 days.
//
// VOD analyzes and processes the CDN request logs of the previous day to generate playback statistics files.
//
// * The playback statistics file contains statistical information such as the number of plays and total traffic of media files.
//
// * Playback count statistics description:
//
// 1. HLS file: The number of plays is counted when an M3U8 file is accessed, but not when a TS file is accessed.
//
// 2. Other files (for example, MP4 files): If the playback request includes the range parameter and the start parameter of range is not equal to 0, the number of plays is not counted. In other cases, the number of plays is counted.
//
// * Statistics of playback devices: If a playback request includes the UserAgent parameter and the UserAgent contains identifiers such as Android or iPhone, it is counted as a mobile playback count. Otherwise, it is counted as a PC playback count.
//
// * Playback statistics only target VOD domains, excluding EdgeOne domain name distribution.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeDailyPlayStatFileList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDailyPlayStatFileList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDailyPlayStatFileListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultDistributionConfigRequest() (request *DescribeDefaultDistributionConfigRequest) {
    request = &DescribeDefaultDistributionConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeDefaultDistributionConfig")
    
    
    return
}

func NewDescribeDefaultDistributionConfigResponse() (response *DescribeDefaultDistributionConfigResponse) {
    response = &DescribeDefaultDistributionConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDefaultDistributionConfig
// This API is used to query the default distribution configuration.
//
// * Distribution domain name and distribution protocol, i.e., the domain name and protocol in the media file distribution URL. Media files are distributed based on the default distribution configuration.
//
// * Playback key, used to calculate player signature.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDefaultDistributionConfig(request *DescribeDefaultDistributionConfigRequest) (response *DescribeDefaultDistributionConfigResponse, err error) {
    return c.DescribeDefaultDistributionConfigWithContext(context.Background(), request)
}

// DescribeDefaultDistributionConfig
// This API is used to query the default distribution configuration.
//
// * Distribution domain name and distribution protocol, i.e., the domain name and protocol in the media file distribution URL. Media files are distributed based on the default distribution configuration.
//
// * Playback key, used to calculate player signature.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDefaultDistributionConfigWithContext(ctx context.Context, request *DescribeDefaultDistributionConfigRequest) (response *DescribeDefaultDistributionConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultDistributionConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeDefaultDistributionConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultDistributionConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultDistributionConfigResponse()
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
// This API is used to query DRM Key Provider Information.
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
// This API is used to query DRM Key Provider Information.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeDrmKeyProviderInfo")
    
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
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// This API is used to get the audio and video quality rebirth template list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEnhanceMediaTemplates(request *DescribeEnhanceMediaTemplatesRequest) (response *DescribeEnhanceMediaTemplatesResponse, err error) {
    return c.DescribeEnhanceMediaTemplatesWithContext(context.Background(), request)
}

// DescribeEnhanceMediaTemplates
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// This API is used to get the audio and video quality rebirth template list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEnhanceMediaTemplatesWithContext(ctx context.Context, request *DescribeEnhanceMediaTemplatesRequest) (response *DescribeEnhanceMediaTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeEnhanceMediaTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeEnhanceMediaTemplates")
    
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
// Tencent Cloud VOD provides media upload, media management, media processing, and other services. During or after the execution of these services, Tencent Cloud VOD also provides various event notifications, allowing developers to detect service processing status and perform next business operations.
//
// 
//
// Developers can use this API to query the current configuration of event notification receiving methods, recipient addresses, and which events have callback notification enabled.
//
// 
//
// Default API request rate limit: 100 requests/second.
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
// Tencent Cloud VOD provides media upload, media management, media processing, and other services. During or after the execution of these services, Tencent Cloud VOD also provides various event notifications, allowing developers to detect service processing status and perform next business operations.
//
// 
//
// Developers can use this API to query the current configuration of event notification receiving methods, recipient addresses, and which events have callback notification enabled.
//
// 
//
// Default API request rate limit: 100 requests/second.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeEventConfig")
    
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
// Used to asynchronously get file attributes.
//
// -Currently only support getting the Md5 and Sha1 of the source file.
//
// -For HLS or DASH input files, only get the attributes of the index file.
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
// Used to asynchronously get file attributes.
//
// -Currently only support getting the Md5 and Sha1 of the source file.
//
// -For HLS or DASH input files, only get the attributes of the index file.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeFileAttributes")
    
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
// This API is used to search for a list of title and trailer templates.
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
// This API is used to search for a list of title and trailer templates.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeHeadTailTemplates")
    
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
// This API is used to query the list of image processing templates, and the pagination query is supported based on conditions.
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
// This API is used to query the list of image processing templates, and the pagination query is supported based on conditions.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeImageProcessingTemplates")
    
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
// This API is used to return the daily image moderation usage information within the specified query time range.
//
// 1. Image moderation statistics data from the last 365 days can be queried.
//
//    2. The query time span should not exceed 90 days.
//
// 3. If the query time span exceeds 1 day, the data returned is at a daily granularity. Otherwise, the data returned is at a 5-minute granularity.
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
// This API is used to return the daily image moderation usage information within the specified query time range.
//
// 1. Image moderation statistics data from the last 365 days can be queried.
//
//    2. The query time span should not exceed 90 days.
//
// 3. If the query time span exceeds 1 day, the data returned is at a daily granularity. Otherwise, the data returned is at a 5-minute granularity.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeImageReviewUsageData")
    
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
// This API is used to query image sprite templates based on conditions with paging.
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
// This API is used to query image sprite templates based on conditions with paging.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeImageSpriteTemplates")
    
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
// This API is used to search the instant transcoding template list.
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
// This API is used to search the instant transcoding template list.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeJustInTimeTranscodeTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJustInTimeTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJustInTimeTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKnowledgeBasesRequest() (request *DescribeKnowledgeBasesRequest) {
    request = &DescribeKnowledgeBasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeKnowledgeBases")
    
    
    return
}

func NewDescribeKnowledgeBasesResponse() (response *DescribeKnowledgeBasesResponse) {
    response = &DescribeKnowledgeBasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeBases
// Queries the knowledge base list. Returns all knowledge base info under the designated user.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeKnowledgeBases(request *DescribeKnowledgeBasesRequest) (response *DescribeKnowledgeBasesResponse, err error) {
    return c.DescribeKnowledgeBasesWithContext(context.Background(), request)
}

// DescribeKnowledgeBases
// Queries the knowledge base list. Returns all knowledge base info under the designated user.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeKnowledgeBasesWithContext(ctx context.Context, request *DescribeKnowledgeBasesRequest) (response *DescribeKnowledgeBasesResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeBasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeKnowledgeBases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeBases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeBasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLLMComprehendTemplatesRequest() (request *DescribeLLMComprehendTemplatesRequest) {
    request = &DescribeLLMComprehendTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeLLMComprehendTemplates")
    
    
    return
}

func NewDescribeLLMComprehendTemplatesResponse() (response *DescribeLLMComprehendTemplatesResponse) {
    response = &DescribeLLMComprehendTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLLMComprehendTemplates
// This API is used to obtain the parsing template detail list of a large model based on the template unique identifier. The returned results include all user-customized large model parsing templates that meet the conditions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLLMComprehendTemplates(request *DescribeLLMComprehendTemplatesRequest) (response *DescribeLLMComprehendTemplatesResponse, err error) {
    return c.DescribeLLMComprehendTemplatesWithContext(context.Background(), request)
}

// DescribeLLMComprehendTemplates
// This API is used to obtain the parsing template detail list of a large model based on the template unique identifier. The returned results include all user-customized large model parsing templates that meet the conditions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLLMComprehendTemplatesWithContext(ctx context.Context, request *DescribeLLMComprehendTemplatesRequest) (response *DescribeLLMComprehendTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLLMComprehendTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeLLMComprehendTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLLMComprehendTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLLMComprehendTemplatesResponse()
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
// This API is used to return the daily License request count within the specified query time range.
//
// 1. License request count stats from the last 365 days can be queried.
//
//    2. The query time span should not exceed 90 days.
//
// 3. If the query time span exceeds 1 day, the data returned is at a daily granularity. Otherwise, the data returned is at a 5-minute granularity.
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
// This API is used to return the daily License request count within the specified query time range.
//
// 1. License request count stats from the last 365 days can be queried.
//
//    2. The query time span should not exceed 90 days.
//
// 3. If the query time span exceeds 1 day, the data returned is at a daily granularity. Otherwise, the data returned is at a 5-minute granularity.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeLicenseUsageData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicenseUsageData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseUsageDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMPSTemplatesRequest() (request *DescribeMPSTemplatesRequest) {
    request = &DescribeMPSTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeMPSTemplates")
    
    
    return
}

func NewDescribeMPSTemplatesResponse() (response *DescribeMPSTemplatesResponse) {
    response = &DescribeMPSTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMPSTemplates
// Query a user-customized media processing service task template.
//
// To query the template list, fill in the MPS related parameters in MPSDescribeTemplateParams in JSON format. For task parameter configuration, see the MPS task template documentation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeMPSTemplates(request *DescribeMPSTemplatesRequest) (response *DescribeMPSTemplatesResponse, err error) {
    return c.DescribeMPSTemplatesWithContext(context.Background(), request)
}

// DescribeMPSTemplates
// Query a user-customized media processing service task template.
//
// To query the template list, fill in the MPS related parameters in MPSDescribeTemplateParams in JSON format. For task parameter configuration, see the MPS task template documentation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeMPSTemplatesWithContext(ctx context.Context, request *DescribeMPSTemplatesRequest) (response *DescribeMPSTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeMPSTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeMPSTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMPSTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMPSTemplatesResponse()
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
// 1. This API can obtain multiple types of info of multiple media files, including:
//
// 1. Basic information (basicInfo): including media name, category, playback address, cover image, etc.
//
// 2. Meta information (metaData): including size, duration, video stream information, audio stream information, etc.
//
// 3. transcodeInfo: includes media addresses, video stream parameters, and audio stream parameters of various specifications generated for the media.
//
// 4. Animated graphics info (animatedGraphicsInfo): the animated graphics info after converting a video to GIF (for example, gif).
//
// 5. sampleSnapshotInfo: Screenshot information after sampling screenshots of a video.
//
// 6. Sprite image information (imageSpriteInfo): sprite image information after capturing sprite image files from a video.
//
// 7. snapshotByTimeOffsetInfo: screenshot information after taking screenshots of a video at specified time points.
//
// 8. Video timestamp information (keyFrameDescInfo): Dotting information set for the video.
//
// 9. Adaptive Bitrate Streaming information (adaptiveDynamicStreamingInfo): information including specification, encryption type, packaging format, and other related details.
//
// 10. Review information (reviewInfo): includes media moderation and media cover review information.
//
// 2. You can specify to only return partial info in the response.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETMEDIALISTERROR = "InternalError.GetMediaListError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_FILEIDSEMPTY = "InvalidParameterValue.FileIdsEmpty"
//  INVALIDPARAMETERVALUE_FILEIDSOVERLIMIT = "InvalidParameterValue.FileIdsOverLimit"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMediaInfos(request *DescribeMediaInfosRequest) (response *DescribeMediaInfosResponse, err error) {
    return c.DescribeMediaInfosWithContext(context.Background(), request)
}

// DescribeMediaInfos
// 1. This API can obtain multiple types of info of multiple media files, including:
//
// 1. Basic information (basicInfo): including media name, category, playback address, cover image, etc.
//
// 2. Meta information (metaData): including size, duration, video stream information, audio stream information, etc.
//
// 3. transcodeInfo: includes media addresses, video stream parameters, and audio stream parameters of various specifications generated for the media.
//
// 4. Animated graphics info (animatedGraphicsInfo): the animated graphics info after converting a video to GIF (for example, gif).
//
// 5. sampleSnapshotInfo: Screenshot information after sampling screenshots of a video.
//
// 6. Sprite image information (imageSpriteInfo): sprite image information after capturing sprite image files from a video.
//
// 7. snapshotByTimeOffsetInfo: screenshot information after taking screenshots of a video at specified time points.
//
// 8. Video timestamp information (keyFrameDescInfo): Dotting information set for the video.
//
// 9. Adaptive Bitrate Streaming information (adaptiveDynamicStreamingInfo): information including specification, encryption type, packaging format, and other related details.
//
// 10. Review information (reviewInfo): includes media moderation and media cover review information.
//
// 2. You can specify to only return partial info in the response.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETMEDIALISTERROR = "InternalError.GetMediaListError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_FILEIDS = "InvalidParameterValue.FileIds"
//  INVALIDPARAMETERVALUE_FILEIDSEMPTY = "InvalidParameterValue.FileIdsEmpty"
//  INVALIDPARAMETERVALUE_FILEIDSOVERLIMIT = "InvalidParameterValue.FileIdsOverLimit"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMediaInfosWithContext(ctx context.Context, request *DescribeMediaInfosRequest) (response *DescribeMediaInfosResponse, err error) {
    if request == nil {
        request = NewDescribeMediaInfosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeMediaInfos")
    
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
// This API is used to query playback data of media files by specified time granularity.
//
// * Playback statistics from the past one year can be queried.
//
// Time granularity: hr. The max span between the end time and start time is 7 days.
//
// Time granularity is day, and the maximum span between the end time and start time is 90 days.
//
// * Playback statistics only target VOD domains, excluding distribution from EdgeOne domain names.
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
// This API is used to query playback data of media files by specified time granularity.
//
// * Playback statistics from the past one year can be queried.
//
// Time granularity: hr. The max span between the end time and start time is 7 days.
//
// Time granularity is day, and the maximum span between the end time and start time is 90 days.
//
// * Playback statistics only target VOD domains, excluding distribution from EdgeOne domain names.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeMediaPlayStatDetails")
    
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
// This API is used to return the daily video processing usage information within the specified query time range.
//
// 1. Video processing usage data is retained on the system side for 13 months. You can only query usage data from the most recent 365 days through the API. If you need to retrieve historical usage data beyond 365 days, contact us.
//
//    2. The query time span should not exceed 90 days.
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
// This API is used to return the daily video processing usage information within the specified query time range.
//
// 1. Video processing usage data is retained on the system side for 13 months. You can only query usage data from the most recent 365 days through the API. If you need to retrieve historical usage data beyond 365 days, contact us.
//
//    2. The query time span should not exceed 90 days.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeMediaProcessUsageData")
    
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
// This API is used to query material sample information by material ID, name, or tag with pagination.
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
// This API is used to query material sample information by material ID, name, or tag with pagination.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribePersonSamples")
    
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
// This API is used to search the task flow template detail list based on the task flow template name.
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
// This API is used to search the task flow template detail list based on the task flow template name.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeProcedureTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcedureTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcedureTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessImageAsyncTemplatesRequest() (request *DescribeProcessImageAsyncTemplatesRequest) {
    request = &DescribeProcessImageAsyncTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeProcessImageAsyncTemplates")
    
    
    return
}

func NewDescribeProcessImageAsyncTemplatesResponse() (response *DescribeProcessImageAsyncTemplatesResponse) {
    response = &DescribeProcessImageAsyncTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProcessImageAsyncTemplates
// This API is used to obtain the template details list based on the Template Unique Identifier. The returned results include all eligible user-customized image asynchronous processing templates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProcessImageAsyncTemplates(request *DescribeProcessImageAsyncTemplatesRequest) (response *DescribeProcessImageAsyncTemplatesResponse, err error) {
    return c.DescribeProcessImageAsyncTemplatesWithContext(context.Background(), request)
}

// DescribeProcessImageAsyncTemplates
// This API is used to obtain the template details list based on the Template Unique Identifier. The returned results include all eligible user-customized image asynchronous processing templates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProcessImageAsyncTemplatesWithContext(ctx context.Context, request *DescribeProcessImageAsyncTemplatesRequest) (response *DescribeProcessImageAsyncTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeProcessImageAsyncTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeProcessImageAsyncTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcessImageAsyncTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcessImageAsyncTemplatesResponse()
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
// This API is used to query the audio and video quality detection template list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeQualityInspectTemplates(request *DescribeQualityInspectTemplatesRequest) (response *DescribeQualityInspectTemplatesResponse, err error) {
    return c.DescribeQualityInspectTemplatesWithContext(context.Background(), request)
}

// DescribeQualityInspectTemplates
// This API is used to query the audio and video quality detection template list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeQualityInspectTemplatesWithContext(ctx context.Context, request *DescribeQualityInspectTemplatesRequest) (response *DescribeQualityInspectTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeQualityInspectTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeQualityInspectTemplates")
    
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
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Queries the video rebirth template list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRebuildMediaTemplates(request *DescribeRebuildMediaTemplatesRequest) (response *DescribeRebuildMediaTemplatesResponse, err error) {
    return c.DescribeRebuildMediaTemplatesWithContext(context.Background(), request)
}

// DescribeRebuildMediaTemplates
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Queries the video rebirth template list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRebuildMediaTemplatesWithContext(ctx context.Context, request *DescribeRebuildMediaTemplatesRequest) (response *DescribeRebuildMediaTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeRebuildMediaTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeRebuildMediaTemplates")
    
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
// <b>This API is not recommended. Use [DescribeMediaProcessUsageData](https://www.tencentcloud.com/document/product/266/41464?from_cn_redirect=1) as an alternative.</b>
//
// 
//
// This API is used to return the daily video content intelligent identification duration data within the specified query time range. Measurement unit: second.
//
// 
//
// 1. Video content intelligent identification duration stats from the last 365 days can be queried.
//
// 2. The query time span should not exceed 90 days.
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
// <b>This API is not recommended. Use [DescribeMediaProcessUsageData](https://www.tencentcloud.com/document/product/266/41464?from_cn_redirect=1) as an alternative.</b>
//
// 
//
// This API is used to return the daily video content intelligent identification duration data within the specified query time range. Measurement unit: second.
//
// 
//
// 1. Video content intelligent identification duration stats from the last 365 days can be queried.
//
// 2. The query time span should not exceed 90 days.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeReviewDetails")
    
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
// Retrieves the moderation template list.
//
// >Template is applicable only to the ReviewAudioVideo (https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) and ReviewImage (https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) APIs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeReviewTemplates(request *DescribeReviewTemplatesRequest) (response *DescribeReviewTemplatesResponse, err error) {
    return c.DescribeReviewTemplatesWithContext(context.Background(), request)
}

// DescribeReviewTemplates
// Retrieves the moderation template list.
//
// >Template is applicable only to the ReviewAudioVideo (https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) and ReviewImage (https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) APIs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeReviewTemplatesWithContext(ctx context.Context, request *DescribeReviewTemplatesRequest) (response *DescribeReviewTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeReviewTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeReviewTemplates")
    
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
// This API is used to get the carousel playlist list.
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
// This API is used to get the carousel playlist list.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeRoundPlays")
    
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
// This API is used to query sampled screenshot templates based on conditions with paging.
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
// This API is used to query sampled screenshot templates based on conditions with paging.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeSampleSnapshotTemplates")
    
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
// This API is used to query specified time point screenshot templates based on conditions with paging.
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
// This API is used to query specified time point screenshot templates based on conditions with paging.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeSnapshotByTimeOffsetTemplates")
    
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
// Queries storage space usage and number of files.
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
// Queries storage space usage and number of files.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeStorageData")
    
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
// This API is used to return the VOD storage space used within a specified time range, in bytes.
//
// 1. Storage usage data is reserved on the system side for 13 months. You can only query usage data from the most recent 365 days through the API. If you need to call historical usage data beyond 365 days, contact us;
//
// 2. The query time span should not exceed 90 days.
//
// 3. The query span at a minute granularity should not exceed 7 days.
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
// This API is used to return the VOD storage space used within a specified time range, in bytes.
//
// 1. Storage usage data is reserved on the system side for 13 months. You can only query usage data from the most recent 365 days through the API. If you need to call historical usage data beyond 365 days, contact us;
//
// 2. The query time span should not exceed 90 days.
//
// 3. The query span at a minute granularity should not exceed 7 days.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeStorageDetails")
    
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
// This API is used to:
//
// 1. Query the list of all storage campuses available for on-demand activation.
//
// 2. Query the opened park list.
//
// 3. Query the storage campus used by default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStorageRegions(request *DescribeStorageRegionsRequest) (response *DescribeStorageRegionsResponse, err error) {
    return c.DescribeStorageRegionsWithContext(context.Background(), request)
}

// DescribeStorageRegions
// This API is used to:
//
// 1. Query the list of all storage campuses available for on-demand activation.
//
// 2. Query the opened park list.
//
// 3. Query the storage campus used by default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStorageRegionsWithContext(ctx context.Context, request *DescribeStorageRegionsRequest) (response *DescribeStorageRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeStorageRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeStorageRegions")
    
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
// This API is used to get the application list of the current account.
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
// This API is used to get the application list of the current account.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeSubAppIds")
    
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
// This API is <font color='red'>no longer maintained</font>. The new version of player signature no longer uses the player configuration template. For details, please see [Player Signature](https://www.tencentcloud.com/document/product/266/45554?from_cn_redirect=1).
//
// Queries player configurations based on conditions with paging.
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
// This API is <font color='red'>no longer maintained</font>. The new version of player signature no longer uses the player configuration template. For details, please see [Player Signature](https://www.tencentcloud.com/document/product/266/45554?from_cn_redirect=1).
//
// Queries player configurations based on conditions with paging.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeSuperPlayerConfigs")
    
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
// This API is used to query the details of the task execution status and results by task ID (tasks submitted within the last 3 days can be queried).
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
// This API is used to query the details of the task execution status and results by task ID (tasks submitted within the last 3 days can be queried).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeTaskDetail")
    
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
// * This API is used to query the task list.
//
// * When the list contains a large amount of data, a single API call cannot pull the entire list. You can use the ScrollToken parameter to pull in batches.
//
// * Only query tasks from the last three days (72 hr).
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
// * This API is used to query the task list.
//
// * When the list contains a large amount of data, a single API call cannot pull the entire list. You can use the ScrollToken parameter to pull in batches.
//
// * Only query tasks from the last three days (72 hr).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeTasks")
    
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
// This API is used to search the transcoding template detail list by transcoding template unique identifier. The returned results include all eligible custom templates and [system preset transcoding templates](https://www.tencentcloud.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_SCENARIOTYPE = "InvalidParameterValue.ScenarioType"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTranscodeTemplates(request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    return c.DescribeTranscodeTemplatesWithContext(context.Background(), request)
}

// DescribeTranscodeTemplates
// This API is used to search the transcoding template detail list by transcoding template unique identifier. The returned results include all eligible custom templates and [system preset transcoding templates](https://www.tencentcloud.com/document/product/266/33476?from_cn_redirect=1#.E9.A2.84.E7.BD.AE.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"
//  INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_SCENARIOTYPE = "InvalidParameterValue.ScenarioType"
//  INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTranscodeTemplatesWithContext(ctx context.Context, request *DescribeTranscodeTemplatesRequest) (response *DescribeTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeTranscodeTemplates")
    
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
// This API is used to query the information list of on-demand domain names.
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
// This API is used to query the information list of on-demand domain names.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeVodDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVodDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVodDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoicesRequest() (request *DescribeVoicesRequest) {
    request = &DescribeVoicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeVoices")
    
    
    return
}

func NewDescribeVoicesResponse() (response *DescribeVoicesResponse) {
    response = &DescribeVoicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVoices
// Queries the list of available timbres under the current account, supporting filtering by optional conditions such as voice ID, type, name, gender, age, language, tag, and scenario.
//
// 
//
// Note: Newly designed or cloned voice types cannot be queried before activation. They are activated only after the new voice type is used for TTS once.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_LIMITTOOLARGE = "InvalidParameterValue.LimitTooLarge"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_OFFSETTOOLARGE = "InvalidParameterValue.OffsetTooLarge"
func (c *Client) DescribeVoices(request *DescribeVoicesRequest) (response *DescribeVoicesResponse, err error) {
    return c.DescribeVoicesWithContext(context.Background(), request)
}

// DescribeVoices
// Queries the list of available timbres under the current account, supporting filtering by optional conditions such as voice ID, type, name, gender, age, language, tag, and scenario.
//
// 
//
// Note: Newly designed or cloned voice types cannot be queried before activation. They are activated only after the new voice type is used for TTS once.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_LIMITTOOLARGE = "InvalidParameterValue.LimitTooLarge"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_OFFSETTOOLARGE = "InvalidParameterValue.OffsetTooLarge"
func (c *Client) DescribeVoicesWithContext(ctx context.Context, request *DescribeVoicesRequest) (response *DescribeVoicesResponse, err error) {
    if request == nil {
        request = NewDescribeVoicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeVoices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoicesResponse()
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
// This API is used to query user-defined watermark templates, and the pagination query is supported based on conditions.
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
// This API is used to query user-defined watermark templates, and the pagination query is supported based on conditions.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeWatermarkTemplates")
    
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
// This API is used to paginate keyword sample information by scenario, keyword, and tag.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWordSamples(request *DescribeWordSamplesRequest) (response *DescribeWordSamplesResponse, err error) {
    return c.DescribeWordSamplesWithContext(context.Background(), request)
}

// DescribeWordSamples
// This API is used to paginate keyword sample information by scenario, keyword, and tag.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWordSamplesWithContext(ctx context.Context, request *DescribeWordSamplesRequest) (response *DescribeWordSamplesResponse, err error) {
    if request == nil {
        request = NewDescribeWordSamplesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DescribeWordSamples")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWordSamples require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWordSamplesResponse()
    err = c.Send(request, response)
    return
}

func NewDesignVoiceAsyncRequest() (request *DesignVoiceAsyncRequest) {
    request = &DesignVoiceAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DesignVoiceAsync")
    
    
    return
}

func NewDesignVoiceAsyncResponse() (response *DesignVoiceAsyncResponse) {
    response = &DesignVoiceAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DesignVoiceAsync
// This API is used to initiate a voice design task. It generates a custom voice based on a natural language description. You can specify a voice profile at the same time, including name, gender, age, language, tag, and scenario. If trial text is attached during submission, audio audition is generated after task completion. Voice design is an asynchronous task. The voice ID is generated after task completion.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DesignVoiceAsync(request *DesignVoiceAsyncRequest) (response *DesignVoiceAsyncResponse, err error) {
    return c.DesignVoiceAsyncWithContext(context.Background(), request)
}

// DesignVoiceAsync
// This API is used to initiate a voice design task. It generates a custom voice based on a natural language description. You can specify a voice profile at the same time, including name, gender, age, language, tag, and scenario. If trial text is attached during submission, audio audition is generated after task completion. Voice design is an asynchronous task. The voice ID is generated after task completion.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DesignVoiceAsyncWithContext(ctx context.Context, request *DesignVoiceAsyncRequest) (response *DesignVoiceAsyncResponse, err error) {
    if request == nil {
        request = NewDesignVoiceAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "DesignVoiceAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DesignVoiceAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDesignVoiceAsyncResponse()
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
// This API is used to edit a video, such as clipping and concatenation, to generate a new video on demand. Editing features include:
//
// 
//
// 1) Edit a file in on-demand video to generate a new video.
//
// 2) Splice multiple on-demand files to generate a new video.
//
// 3) Edit multiple on-demand video files and then splice them to generate a new video;
//
// 4. Directly generate a new video for one of the streams in VOD;
//
// 5. Edit one of the VOD streams to generate a new video;
//
// 6) Splice multiple on-demand streams to generate a new video.
//
// 7) Edit multiple streams in VOD and then splice them to generate a new video.
//
// 
//
// For the generated new video, you can also specify whether to execute task flow for the generated video.
//
// 
//
// >When editing or splicing a live stream, please ensure the stream ended before you operate. Otherwise, the generated video may be incomplete.
//
// 
//
// If event notification is used, its type is video editing completed (https://www.tencentcloud.com/document/product/266/33794?from_cn_redirect=1).
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
// This API is used to edit a video, such as clipping and concatenation, to generate a new video on demand. Editing features include:
//
// 
//
// 1) Edit a file in on-demand video to generate a new video.
//
// 2) Splice multiple on-demand files to generate a new video.
//
// 3) Edit multiple on-demand video files and then splice them to generate a new video;
//
// 4. Directly generate a new video for one of the streams in VOD;
//
// 5. Edit one of the VOD streams to generate a new video;
//
// 6) Splice multiple on-demand streams to generate a new video.
//
// 7) Edit multiple streams in VOD and then splice them to generate a new video.
//
// 
//
// For the generated new video, you can also specify whether to execute task flow for the generated video.
//
// 
//
// >When editing or splicing a live stream, please ensure the stream ended before you operate. Otherwise, the generated video may be incomplete.
//
// 
//
// If event notification is used, its type is video editing completed (https://www.tencentcloud.com/document/product/266/33794?from_cn_redirect=1).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "EditMedia")
    
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
// This API is <font color=red>no longer maintained</font>. Please use the new version of APIs for [audio and video quality revival](https://www.tencentcloud.com/document/api/266/102571?from_cn_redirect=1).
//
// Use a template to initiate audio and video quality revival.
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
// This API is <font color=red>no longer maintained</font>. Please use the new version of APIs for [audio and video quality revival](https://www.tencentcloud.com/document/api/266/102571?from_cn_redirect=1).
//
// Use a template to initiate audio and video quality revival.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "EnhanceMediaByTemplate")
    
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
// This API is used to initiate an audio and video quality regeneration task for on-demand audio-video media.
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
// This API is used to initiate an audio and video quality regeneration task for on-demand audio-video media.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "EnhanceMediaQuality")
    
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
// This API is only used for special customized development scenarios. Do not call this API unless VOD customer service proactively informs you to do so.
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
// This API is only used for special customized development scenarios. Do not call this API unless VOD customer service proactively informs you to do so.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ExecuteFunction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewExtractBlindWatermarkRequest() (request *ExtractBlindWatermarkRequest) {
    request = &ExtractBlindWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ExtractBlindWatermark")
    
    
    return
}

func NewExtractBlindWatermarkResponse() (response *ExtractBlindWatermarkResponse) {
    response = &ExtractBlindWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExtractBlindWatermark
// This API is used to initiate a digital watermark extraction task for a video. The extraction result can be queried through DescribeTaskDetail.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) ExtractBlindWatermark(request *ExtractBlindWatermarkRequest) (response *ExtractBlindWatermarkResponse, err error) {
    return c.ExtractBlindWatermarkWithContext(context.Background(), request)
}

// ExtractBlindWatermark
// This API is used to initiate a digital watermark extraction task for a video. The extraction result can be queried through DescribeTaskDetail.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) ExtractBlindWatermarkWithContext(ctx context.Context, request *ExtractBlindWatermarkRequest) (response *ExtractBlindWatermarkResponse, err error) {
    if request == nil {
        request = NewExtractBlindWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ExtractBlindWatermark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExtractBlindWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewExtractBlindWatermarkResponse()
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
// If you need source tracing for piracy, see Ghost Watermark (https://www.tencentcloud.com/document/product/266/94228?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) ExtractCopyRightWatermark(request *ExtractCopyRightWatermarkRequest) (response *ExtractCopyRightWatermarkResponse, err error) {
    return c.ExtractCopyRightWatermarkWithContext(context.Background(), request)
}

// ExtractCopyRightWatermark
// If you need source tracing for piracy, see Ghost Watermark (https://www.tencentcloud.com/document/product/266/94228?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) ExtractCopyRightWatermarkWithContext(ctx context.Context, request *ExtractCopyRightWatermarkRequest) (response *ExtractCopyRightWatermarkResponse, err error) {
    if request == nil {
        request = NewExtractCopyRightWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ExtractCopyRightWatermark")
    
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
// If source tracing for piracy is required, ghost watermark is recommended for use (https://www.tencentcloud.com/document/product/266/94228?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) ExtractTraceWatermark(request *ExtractTraceWatermarkRequest) (response *ExtractTraceWatermarkResponse, err error) {
    return c.ExtractTraceWatermarkWithContext(context.Background(), request)
}

// ExtractTraceWatermark
// If source tracing for piracy is required, ghost watermark is recommended for use (https://www.tencentcloud.com/document/product/266/94228?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) ExtractTraceWatermarkWithContext(ctx context.Context, request *ExtractTraceWatermarkRequest) (response *ExtractTraceWatermarkResponse, err error) {
    if request == nil {
        request = NewExtractTraceWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ExtractTraceWatermark")
    
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
// Quickly splice and edit HLS videos in VOD to generate new media in HLS format.
//
// 
//
// Quickly splice or edit the generated video to generate a new FileId and solidify it. After successful solidification, the new video file exists independent of the original input video and is not affected by deletion of the original video.
//
// 
//
// <font color='red'>Note:</font> Enable reception of editing solidification event notifications through the ModifyEventConfig API. After successful solidification, you will receive a PersistenceComplete event notification. Before receiving this event notification, you should not delete or reduce the storage class of the original input video. Otherwise, playback of the video generated by splicing and clipping may be abnormal.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) FastEditMedia(request *FastEditMediaRequest) (response *FastEditMediaResponse, err error) {
    return c.FastEditMediaWithContext(context.Background(), request)
}

// FastEditMedia
// Quickly splice and edit HLS videos in VOD to generate new media in HLS format.
//
// 
//
// Quickly splice or edit the generated video to generate a new FileId and solidify it. After successful solidification, the new video file exists independent of the original input video and is not affected by deletion of the original video.
//
// 
//
// <font color='red'>Note:</font> Enable reception of editing solidification event notifications through the ModifyEventConfig API. After successful solidification, you will receive a PersistenceComplete event notification. Before receiving this event notification, you should not delete or reduce the storage class of the original input video. Otherwise, playback of the video generated by splicing and clipping may be abnormal.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) FastEditMediaWithContext(ctx context.Context, request *FastEditMediaRequest) (response *FastEditMediaResponse, err error) {
    if request == nil {
        request = NewFastEditMediaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "FastEditMedia")
    
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
// After media blocking, except for VOD console preview, accessing video resource URLs (raw files, transcoding output files, screenshots, etc.) for other scenarios will return 403.
//
// It takes about 5 to 10 minutes for the block/unblock operation to take effect across the entire network.
//
// * Note: Blocking media can only operate on media in standard storage and infrequent storage. Media in infrequent storage must be stored for at least 30 days. If deleted early or changed to another storage class, it is still billed for 30 days. If you block media in infrequent storage and its infrequent storage duration is less than 30 days, early deletion billing occurs. At the same time, after blocking, the infrequent storage duration of the media restarts from the current time. If the media is deleted or changed to another storage class before reaching 30 days, early deletion billing also occurs. For example, media 001 has been in infrequent storage for 10 days. If you block 001 at this point, infrequent storage is still billed for 30 days (early deletion billing duration: 30 - 10 = 20 days). After blocking, the infrequent storage duration of 001 restarts. If 001 is deleted on day 5 after blocking, infrequent storage is still billed for 30 days (early deletion billing duration: 30 - 5 = 25 days). The actual infrequent storage duration of 001 is 10 + 5 = 15 days, and the infrequent storage billing duration is 10 + 20 (early deletion billing) + 5 + 25 (early deletion billing) = 60 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_MEDIAFORBIDEDBYSYSTEM = "FailedOperation.MediaForbidedBySystem"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEIDSTOOMANY = "InvalidParameterValue.FileIdsTooMany"
//  INVALIDPARAMETERVALUE_OPERATION = "InvalidParameterValue.Operation"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ForbidMediaDistribution(request *ForbidMediaDistributionRequest) (response *ForbidMediaDistributionResponse, err error) {
    return c.ForbidMediaDistributionWithContext(context.Background(), request)
}

// ForbidMediaDistribution
// After media blocking, except for VOD console preview, accessing video resource URLs (raw files, transcoding output files, screenshots, etc.) for other scenarios will return 403.
//
// It takes about 5 to 10 minutes for the block/unblock operation to take effect across the entire network.
//
// * Note: Blocking media can only operate on media in standard storage and infrequent storage. Media in infrequent storage must be stored for at least 30 days. If deleted early or changed to another storage class, it is still billed for 30 days. If you block media in infrequent storage and its infrequent storage duration is less than 30 days, early deletion billing occurs. At the same time, after blocking, the infrequent storage duration of the media restarts from the current time. If the media is deleted or changed to another storage class before reaching 30 days, early deletion billing also occurs. For example, media 001 has been in infrequent storage for 10 days. If you block 001 at this point, infrequent storage is still billed for 30 days (early deletion billing duration: 30 - 10 = 20 days). After blocking, the infrequent storage duration of 001 restarts. If 001 is deleted on day 5 after blocking, infrequent storage is still billed for 30 days (early deletion billing duration: 30 - 5 = 25 days). The actual infrequent storage duration of 001 is 10 + 5 = 15 days, and the infrequent storage billing duration is 10 + 20 (early deletion billing) + 5 + 25 (early deletion billing) = 60 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  FAILEDOPERATION_MEDIAFORBIDEDBYSYSTEM = "FailedOperation.MediaForbidedBySystem"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEIDSTOOMANY = "InvalidParameterValue.FileIdsTooMany"
//  INVALIDPARAMETERVALUE_OPERATION = "InvalidParameterValue.Operation"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ForbidMediaDistributionWithContext(ctx context.Context, request *ForbidMediaDistributionRequest) (response *ForbidMediaDistributionResponse, err error) {
    if request == nil {
        request = NewForbidMediaDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ForbidMediaDistribution")
    
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
// Perform operations on the carousel current playlist. Supported operations: <li> Insert: insert a play program into the current playlist.</li><li> Delete: delete a play program from the playlist.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_INDEX = "InvalidParameterValue.Index"
//  INVALIDPARAMETERVALUE_ITEMID = "InvalidParameterValue.ItemId"
//  INVALIDPARAMETERVALUE_ROUNDPLAYID = "InvalidParameterValue.RoundPlayId"
//  INVALIDPARAMETERVALUE_ROUNDPLAYLIST = "InvalidParameterValue.RoundPlaylist"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_PLAYLIST = "LimitExceeded.PlayList"
//  LIMITEXCEEDED_ROUNDPLAYLIST = "LimitExceeded.RoundPlaylist"
func (c *Client) HandleCurrentPlaylist(request *HandleCurrentPlaylistRequest) (response *HandleCurrentPlaylistResponse, err error) {
    return c.HandleCurrentPlaylistWithContext(context.Background(), request)
}

// HandleCurrentPlaylist
// Perform operations on the carousel current playlist. Supported operations: <li> Insert: insert a play program into the current playlist.</li><li> Delete: delete a play program from the playlist.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_INDEX = "InvalidParameterValue.Index"
//  INVALIDPARAMETERVALUE_ITEMID = "InvalidParameterValue.ItemId"
//  INVALIDPARAMETERVALUE_ROUNDPLAYID = "InvalidParameterValue.RoundPlayId"
//  INVALIDPARAMETERVALUE_ROUNDPLAYLIST = "InvalidParameterValue.RoundPlaylist"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  LIMITEXCEEDED_PLAYLIST = "LimitExceeded.PlayList"
//  LIMITEXCEEDED_ROUNDPLAYLIST = "LimitExceeded.RoundPlaylist"
func (c *Client) HandleCurrentPlaylistWithContext(ctx context.Context, request *HandleCurrentPlaylistRequest) (response *HandleCurrentPlaylistResponse, err error) {
    if request == nil {
        request = NewHandleCurrentPlaylistRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "HandleCurrentPlaylist")
    
    if c.GetCredential() == nil {
        return nil, errors.New("HandleCurrentPlaylist require credential")
    }

    request.SetContext(ctx)
    
    response = NewHandleCurrentPlaylistResponse()
    err = c.Send(request, response)
    return
}

func NewImportMediaKnowledgeRequest() (request *ImportMediaKnowledgeRequest) {
    request = &ImportMediaKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ImportMediaKnowledge")
    
    
    return
}

func NewImportMediaKnowledgeResponse() (response *ImportMediaKnowledgeResponse) {
    response = &ImportMediaKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportMediaKnowledge
// Used to import AI analysis results into the knowledge base.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ImportMediaKnowledge(request *ImportMediaKnowledgeRequest) (response *ImportMediaKnowledgeResponse, err error) {
    return c.ImportMediaKnowledgeWithContext(context.Background(), request)
}

// ImportMediaKnowledge
// Used to import AI analysis results into the knowledge base.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ImportMediaKnowledgeWithContext(ctx context.Context, request *ImportMediaKnowledgeRequest) (response *ImportMediaKnowledgeResponse, err error) {
    if request == nil {
        request = NewImportMediaKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ImportMediaKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportMediaKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportMediaKnowledgeResponse()
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
// This API is used to trigger an audio and video quality inspection task for on-demand audio-video media.
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
// This API is used to trigger an audio and video quality inspection task for on-demand audio-video media.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "InspectMediaQuality")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InspectMediaQuality require credential")
    }

    request.SetContext(ctx)
    
    response = NewInspectMediaQualityResponse()
    err = c.Send(request, response)
    return
}

func NewListFilesRequest() (request *ListFilesRequest) {
    request = &ListFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ListFiles")
    
    
    return
}

func NewListFilesResponse() (response *ListFilesResponse) {
    response = &ListFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListFiles
// This API is used to list stored file entries under a sub-app.
//
// 
//
// **This API is only available in FileID+Path mode**
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
func (c *Client) ListFiles(request *ListFilesRequest) (response *ListFilesResponse, err error) {
    return c.ListFilesWithContext(context.Background(), request)
}

// ListFiles
// This API is used to list stored file entries under a sub-app.
//
// 
//
// **This API is only available in FileID+Path mode**
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
func (c *Client) ListFilesWithContext(ctx context.Context, request *ListFilesRequest) (response *ListFilesResponse, err error) {
    if request == nil {
        request = NewListFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ListFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewListFilesResponse()
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
// Live stream clipping refers to the ability for customers to select a segment from past live stream content during live streaming (when the live stream has not yet ended), generate a new video in real time (HLS format), and developers can share it instantly or store it for long-term preservation.
//
// 
//
// Tencent Cloud VOD supports two real-time clipping modes:
//
// - Clip solidification: Save the edited video as an independent video with its own FileId. This is suitable for long-term preservation of highlights.
//
// - Editing is not solidified: The edited video is attached to the live streaming recording file and has no standalone FileId. This is suitable for scenarios where highlights are shared temporarily.
//
// 
//
// Note:
//
// - The premise for using the live stream clipping feature is that the target live stream has the time shifting and playback (https://www.tencentcloud.com/document/product/267/32742?from_cn_redirect=1) feature enabled.
//
// -Live streaming Instant Editing is based on the m3u8 file generated by live recording, so its minimum editing precision is one ts slice. Second-level or more precise editing precision cannot be achieved.
//
// -Since stream disconnection may occur during live streaming, the actual video duration generated by editing might differ from the expected duration. For example, if you edit a live stream from 2018-09-20T10:30:00Z to 2018-09-20T10:40:00Z, and a stream disconnection occurred during this time interval, the returned media asset file duration will be less than 10 minutes. In such cases, you can perceive it through the output parameter <a href="#p_segmentset">SegmentSet</a>.
//
// 
//
// ### Edit solidification
//
// Editing solidification means saving the edited video as an independent video (with an independent FileId). Its lifecycle is not subject to any impact from the original live recorded video (even if the original recorded video is deleted, the clipping result will not be affected). It can also be transcoded, published on WeChat, or undergo other secondary processing.
//
// 
//
// For example, a complete football match may last for more than 2 hours. The customer can store the original video for 2 months for cost savings, but can specify a longer storage period for the highlight reel from live stream clipping. You can also perform additional on-demand operations on the highlight reel, such as transcoding and publishing on WeChat. In this case, you can choose the live stream clipping and persistence solution.
//
// 
//
// The advantage of solidified editing is that its lifecycle is independent of the original recorded video, allowing for separate management and long-term preservation.
//
// 
//
// <font color='red'>Note:</font> If solidification is specified when editing, enable reception of editing solidification event notifications through the ModifyEventConfig API. After successful solidification, you will receive a PersistenceComplete event notification. Before receiving this event notification, you should not delete or transition the live video recording to colder storage. Otherwise, playback of the video generated by editing may be abnormal.
//
// 
//
// ### Editing is not solidified
//
// Editing is not solidified, meaning the result of editing (m3u8 file) shares the same TS segments with the live video recording. The newly generated video is not an independent and complete video (no standalone FileId, only a playback URL), and its valid period is consistent with that of the full live recording video. Once the live recording video is deleted, the clip will also become unplayable.
//
// 
//
// Editing is not solidified. Since the clipping result is not an independent video, it is not included in video management of on-demand media assets. For example, the total number of videos in the console does not count this clip. You also cannot separately transcode, publish on WeChat, or perform any other video processing operation on this clip.
//
// 
//
// The advantage of editing not being solidified is that the editing operation is relatively "lightweight" and will not generate additional storage overhead. However, its shortcoming is that the lifecycle is identical to the original recorded video, and it is unable to further transcode or perform other video processing.
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
// Live stream clipping refers to the ability for customers to select a segment from past live stream content during live streaming (when the live stream has not yet ended), generate a new video in real time (HLS format), and developers can share it instantly or store it for long-term preservation.
//
// 
//
// Tencent Cloud VOD supports two real-time clipping modes:
//
// - Clip solidification: Save the edited video as an independent video with its own FileId. This is suitable for long-term preservation of highlights.
//
// - Editing is not solidified: The edited video is attached to the live streaming recording file and has no standalone FileId. This is suitable for scenarios where highlights are shared temporarily.
//
// 
//
// Note:
//
// - The premise for using the live stream clipping feature is that the target live stream has the time shifting and playback (https://www.tencentcloud.com/document/product/267/32742?from_cn_redirect=1) feature enabled.
//
// -Live streaming Instant Editing is based on the m3u8 file generated by live recording, so its minimum editing precision is one ts slice. Second-level or more precise editing precision cannot be achieved.
//
// -Since stream disconnection may occur during live streaming, the actual video duration generated by editing might differ from the expected duration. For example, if you edit a live stream from 2018-09-20T10:30:00Z to 2018-09-20T10:40:00Z, and a stream disconnection occurred during this time interval, the returned media asset file duration will be less than 10 minutes. In such cases, you can perceive it through the output parameter <a href="#p_segmentset">SegmentSet</a>.
//
// 
//
// ### Edit solidification
//
// Editing solidification means saving the edited video as an independent video (with an independent FileId). Its lifecycle is not subject to any impact from the original live recorded video (even if the original recorded video is deleted, the clipping result will not be affected). It can also be transcoded, published on WeChat, or undergo other secondary processing.
//
// 
//
// For example, a complete football match may last for more than 2 hours. The customer can store the original video for 2 months for cost savings, but can specify a longer storage period for the highlight reel from live stream clipping. You can also perform additional on-demand operations on the highlight reel, such as transcoding and publishing on WeChat. In this case, you can choose the live stream clipping and persistence solution.
//
// 
//
// The advantage of solidified editing is that its lifecycle is independent of the original recorded video, allowing for separate management and long-term preservation.
//
// 
//
// <font color='red'>Note:</font> If solidification is specified when editing, enable reception of editing solidification event notifications through the ModifyEventConfig API. After successful solidification, you will receive a PersistenceComplete event notification. Before receiving this event notification, you should not delete or transition the live video recording to colder storage. Otherwise, playback of the video generated by editing may be abnormal.
//
// 
//
// ### Editing is not solidified
//
// Editing is not solidified, meaning the result of editing (m3u8 file) shares the same TS segments with the live video recording. The newly generated video is not an independent and complete video (no standalone FileId, only a playback URL), and its valid period is consistent with that of the full live recording video. Once the live recording video is deleted, the clip will also become unplayable.
//
// 
//
// Editing is not solidified. Since the clipping result is not an independent video, it is not included in video management of on-demand media assets. For example, the total number of videos in the console does not count this clip. You also cannot separately transcode, publish on WeChat, or perform any other video processing operation on this clip.
//
// 
//
// The advantage of editing not being solidified is that the editing operation is relatively "lightweight" and will not generate additional storage overhead. However, its shortcoming is that the lifecycle is identical to the original recorded video, and it is unable to further transcode or perform other video processing.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "LiveRealTimeClip")
    
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
// Manage initiated tasks.
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
// Manage initiated tasks.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ManageTask")
    
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
// Modify a user-defined audio and video content analysis template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and are not allowed to be modified.
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
// Modify a user-defined audio and video content analysis template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and are not allowed to be modified.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyAIAnalysisTemplate")
    
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
// This API is used to modify a user-defined audio and video content recognition template.
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
// This API is used to modify a user-defined audio and video content recognition template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyAIRecognitionTemplate")
    
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
// Modifying an Adaptive Bitrate Streaming Template
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
// Modifying an Adaptive Bitrate Streaming Template
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyAdaptiveDynamicStreamingTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAdaptiveDynamicStreamingTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAdaptiveDynamicStreamingTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAigcQuotaRequest() (request *ModifyAigcQuotaRequest) {
    request = &ModifyAigcQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyAigcQuota")
    
    
    return
}

func NewModifyAigcQuotaResponse() (response *ModifyAigcQuotaResponse) {
    response = &ModifyAigcQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAigcQuota
// Used to edit AIGC quota configuration. Quota usage is accumulated from the start of the quota feature. Once the quota is reached, the AIGC feature will no longer be usable.
//
// 
//
// Since AGC content generation is an async task, real-time usage data cannot be obtained. Therefore, quota limits result in some errors, and precise control over the set limit cannot be achieved.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyAigcQuota(request *ModifyAigcQuotaRequest) (response *ModifyAigcQuotaResponse, err error) {
    return c.ModifyAigcQuotaWithContext(context.Background(), request)
}

// ModifyAigcQuota
// Used to edit AIGC quota configuration. Quota usage is accumulated from the start of the quota feature. Once the quota is reached, the AIGC feature will no longer be usable.
//
// 
//
// Since AGC content generation is an async task, real-time usage data cannot be obtained. Therefore, quota limits result in some errors, and precise control over the set limit cannot be achieved.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyAigcQuotaWithContext(ctx context.Context, request *ModifyAigcQuotaRequest) (response *ModifyAigcQuotaResponse, err error) {
    if request == nil {
        request = NewModifyAigcQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyAigcQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAigcQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAigcQuotaResponse()
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
// Modify a custom animated image generating template.
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
// Modify a custom animated image generating template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyAnimatedGraphicsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAnimatedGraphicsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAnimatedGraphicsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlindWatermarkTemplateRequest() (request *ModifyBlindWatermarkTemplateRequest) {
    request = &ModifyBlindWatermarkTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyBlindWatermarkTemplate")
    
    
    return
}

func NewModifyBlindWatermarkTemplateResponse() (response *ModifyBlindWatermarkTemplateResponse) {
    response = &ModifyBlindWatermarkTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBlindWatermarkTemplate
// This API is used to modify a user-defined digital watermark template. The digital watermark type cannot be modified.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyBlindWatermarkTemplate(request *ModifyBlindWatermarkTemplateRequest) (response *ModifyBlindWatermarkTemplateResponse, err error) {
    return c.ModifyBlindWatermarkTemplateWithContext(context.Background(), request)
}

// ModifyBlindWatermarkTemplate
// This API is used to modify a user-defined digital watermark template. The digital watermark type cannot be modified.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyBlindWatermarkTemplateWithContext(ctx context.Context, request *ModifyBlindWatermarkTemplateRequest) (response *ModifyBlindWatermarkTemplateResponse, err error) {
    if request == nil {
        request = NewModifyBlindWatermarkTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyBlindWatermarkTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlindWatermarkTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlindWatermarkTemplateResponse()
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
// This API is used to modify a CDN domain name configuration.
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
// This API is used to modify a CDN domain name configuration.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyCDNDomainConfig")
    
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
// Modify media classification attributes.
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
// Modify media classification attributes.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyClass")
    
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
// This API is <font color=red>no longer maintained</font>. The new version of the moderation template supports video moderation and image moderation. For details, please see [Modify Moderation Template](https://www.tencentcloud.com/document/api/266/84388?from_cn_redirect=1).
//
// Modify a user-customized audio/video moderation template.
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
// This API is <font color=red>no longer maintained</font>. The new version of the moderation template supports video moderation and image moderation. For details, please see [Modify Moderation Template](https://www.tencentcloud.com/document/api/266/84388?from_cn_redirect=1).
//
// Modify a user-customized audio/video moderation template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyContentReviewTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyContentReviewTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyContentReviewTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDefaultDistributionConfigRequest() (request *ModifyDefaultDistributionConfigRequest) {
    request = &ModifyDefaultDistributionConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyDefaultDistributionConfig")
    
    
    return
}

func NewModifyDefaultDistributionConfigResponse() (response *ModifyDefaultDistributionConfigResponse) {
    response = &ModifyDefaultDistributionConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDefaultDistributionConfig
// This API is used to modify the default distribution configuration.
//
// * Distribution domain name and distribution protocol, i.e., the domain name and protocol in the media file distribution URL. Media files are distributed based on the default distribution configuration.
//
// Playback key, used to calculate player signature.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETERVALUE_DOMAINNAME = "InvalidParameterValue.DomainName"
//  INVALIDPARAMETERVALUE_SCHEME = "InvalidParameterValue.Scheme"
//  INVALIDPARAMETERVALUE_SCHEMECONFLICT = "InvalidParameterValue.SchemeConflict"
//  INVALIDPARAMETERVALUE_UNSUPPORTDOMAIN = "InvalidParameterValue.UnsupportDomain"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDefaultDistributionConfig(request *ModifyDefaultDistributionConfigRequest) (response *ModifyDefaultDistributionConfigResponse, err error) {
    return c.ModifyDefaultDistributionConfigWithContext(context.Background(), request)
}

// ModifyDefaultDistributionConfig
// This API is used to modify the default distribution configuration.
//
// * Distribution domain name and distribution protocol, i.e., the domain name and protocol in the media file distribution URL. Media files are distributed based on the default distribution configuration.
//
// Playback key, used to calculate player signature.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INVALIDPARAMETERVALUE_DOMAINNAME = "InvalidParameterValue.DomainName"
//  INVALIDPARAMETERVALUE_SCHEME = "InvalidParameterValue.Scheme"
//  INVALIDPARAMETERVALUE_SCHEMECONFLICT = "InvalidParameterValue.SchemeConflict"
//  INVALIDPARAMETERVALUE_UNSUPPORTDOMAIN = "InvalidParameterValue.UnsupportDomain"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDefaultDistributionConfigWithContext(ctx context.Context, request *ModifyDefaultDistributionConfigRequest) (response *ModifyDefaultDistributionConfigResponse, err error) {
    if request == nil {
        request = NewModifyDefaultDistributionConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyDefaultDistributionConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDefaultDistributionConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDefaultDistributionConfigResponse()
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
// This API is used to set the default storage region. If no region is specified during file upload, files are uploaded to the default region.
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
// This API is used to set the default storage region. If no region is specified during file upload, files are uploaded to the default region.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyDefaultStorageRegion")
    
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
// This API is no longer maintained. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) interface uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Modifies an Audio and Video Quality Rebirth Template.
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
// This API is no longer maintained. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) interface uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Modifies an Audio and Video Quality Rebirth Template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyEnhanceMediaTemplate")
    
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
// Tencent Cloud Video on Demand (VOD) provides customers with media upload, media management, media processing, and other services. During or after these services are executed, VOD also offers various event notifications, helping developers detect service processing status and perform the next business operation.
//
// 
//
// Developers can call this interface to achieve the following:
//
// - Set the type of callback notification to receive. Currently, there are two types: [HTTP callback notification](https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1) and [Reliable Notification Based on Message Queue](https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1).
//
// - For [HTTP callback notification](https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1), you can set the address for 3.0 format callback. For the description of 3.0 format callback, see [Historical format callback](https://www.tencentcloud.com/document/product/266/33796?from_cn_redirect=1).
//
// -Select to set receipt or ignore for notification events of a specific event service.
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
// Tencent Cloud Video on Demand (VOD) provides customers with media upload, media management, media processing, and other services. During or after these services are executed, VOD also offers various event notifications, helping developers detect service processing status and perform the next business operation.
//
// 
//
// Developers can call this interface to achieve the following:
//
// - Set the type of callback notification to receive. Currently, there are two types: [HTTP callback notification](https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1) and [Reliable Notification Based on Message Queue](https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1).
//
// - For [HTTP callback notification](https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1), you can set the address for 3.0 format callback. For the description of 3.0 format callback, see [Historical format callback](https://www.tencentcloud.com/document/product/266/33796?from_cn_redirect=1).
//
// -Select to set receipt or ignore for notification events of a specific event service.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyEventConfig")
    
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
// Modifies a title and trailer template.
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
// Modifies a title and trailer template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyHeadTailTemplate")
    
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
// Modify a custom image sprite template.
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
// Modify a custom image sprite template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyImageSpriteTemplate")
    
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
// This API is used to modify a just in time transcoding template.
//
// -Note: After a just in time transcoding template is created, modification is not recommended. If parameter modification is needed, add a template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyJustInTimeTranscodeTemplate(request *ModifyJustInTimeTranscodeTemplateRequest) (response *ModifyJustInTimeTranscodeTemplateResponse, err error) {
    return c.ModifyJustInTimeTranscodeTemplateWithContext(context.Background(), request)
}

// ModifyJustInTimeTranscodeTemplate
// This API is used to modify a just in time transcoding template.
//
// -Note: After a just in time transcoding template is created, modification is not recommended. If parameter modification is needed, add a template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyJustInTimeTranscodeTemplateWithContext(ctx context.Context, request *ModifyJustInTimeTranscodeTemplateRequest) (response *ModifyJustInTimeTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyJustInTimeTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyJustInTimeTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyJustInTimeTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyJustInTimeTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKnowledgeBaseRequest() (request *ModifyKnowledgeBaseRequest) {
    request = &ModifyKnowledgeBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyKnowledgeBase")
    
    
    return
}

func NewModifyKnowledgeBaseResponse() (response *ModifyKnowledgeBaseResponse) {
    response = &ModifyKnowledgeBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyKnowledgeBase
// This API is used to modify a knowledge base. The name and/or description of the knowledge base can be modified. A minimum of one field, Name or Description, is required.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyKnowledgeBase(request *ModifyKnowledgeBaseRequest) (response *ModifyKnowledgeBaseResponse, err error) {
    return c.ModifyKnowledgeBaseWithContext(context.Background(), request)
}

// ModifyKnowledgeBase
// This API is used to modify a knowledge base. The name and/or description of the knowledge base can be modified. A minimum of one field, Name or Description, is required.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyKnowledgeBaseWithContext(ctx context.Context, request *ModifyKnowledgeBaseRequest) (response *ModifyKnowledgeBaseResponse, err error) {
    if request == nil {
        request = NewModifyKnowledgeBaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyKnowledgeBase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyKnowledgeBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyKnowledgeBaseResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLLMComprehendTemplateRequest() (request *ModifyLLMComprehendTemplateRequest) {
    request = &ModifyLLMComprehendTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyLLMComprehendTemplate")
    
    
    return
}

func NewModifyLLMComprehendTemplateResponse() (response *ModifyLLMComprehendTemplateResponse) {
    response = &ModifyLLMComprehendTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLLMComprehendTemplate
// Modify a large model parsing template
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_EXTENDEDPARAMETER = "InvalidParameterValue.ExtendedParameter"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) ModifyLLMComprehendTemplate(request *ModifyLLMComprehendTemplateRequest) (response *ModifyLLMComprehendTemplateResponse, err error) {
    return c.ModifyLLMComprehendTemplateWithContext(context.Background(), request)
}

// ModifyLLMComprehendTemplate
// Modify a large model parsing template
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"
//  INVALIDPARAMETERVALUE_EXTENDEDPARAMETER = "InvalidParameterValue.ExtendedParameter"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"
//  LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"
func (c *Client) ModifyLLMComprehendTemplateWithContext(ctx context.Context, request *ModifyLLMComprehendTemplateRequest) (response *ModifyLLMComprehendTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLLMComprehendTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyLLMComprehendTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLLMComprehendTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLLMComprehendTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMPSTemplateRequest() (request *ModifyMPSTemplateRequest) {
    request = &ModifyMPSTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyMPSTemplate")
    
    
    return
}

func NewModifyMPSTemplateResponse() (response *ModifyMPSTemplateResponse) {
    response = &ModifyMPSTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMPSTemplate
// Modify a user-customized MPS task template.
//
// When modifying a template, fill in MPS related parameters in MPSModifyTemplateParams in JSON format. For task parameter configuration methods, refer to the MPS task template documentation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyMPSTemplate(request *ModifyMPSTemplateRequest) (response *ModifyMPSTemplateResponse, err error) {
    return c.ModifyMPSTemplateWithContext(context.Background(), request)
}

// ModifyMPSTemplate
// Modify a user-customized MPS task template.
//
// When modifying a template, fill in MPS related parameters in MPSModifyTemplateParams in JSON format. For task parameter configuration methods, refer to the MPS task template documentation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyMPSTemplateWithContext(ctx context.Context, request *ModifyMPSTemplateRequest) (response *ModifyMPSTemplateResponse, err error) {
    if request == nil {
        request = NewModifyMPSTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyMPSTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMPSTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMPSTemplateResponse()
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
// This API is used to modify media file attributes, including category, name, description, tag, expiration time, dotting information, video cover, and subtitle information.
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
// This API is used to modify media file attributes, including category, name, description, tag, expiration time, dotting information, video cover, and subtitle information.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyMediaInfo")
    
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
// Modify the storage type of media files.
//
// When the storage type of a media file is standard storage, it can be modified to the following types:
//
// <li>Infrequent storage</li>
//
// <li>Archive storage</li>
//
// <li>DEEP_ARCHIVE</li>
//
// When the current storage type of a media file is infrequent storage, it can be modified to the following types:
//
// <li>Standard storage</li>
//
// <li>Archive storage</li>
//
// <li>DEEP_ARCHIVE</li>
//
// When the current storage type of a media file is archive storage, it can be modified to the following types:
//
// <li>Standard storage</li>
//
// When the current storage type of a media file is DEEP_ARCHIVE, it can be modified to the following types:
//
// <li>Standard storage</li>
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
// Modify the storage type of media files.
//
// When the storage type of a media file is standard storage, it can be modified to the following types:
//
// <li>Infrequent storage</li>
//
// <li>Archive storage</li>
//
// <li>DEEP_ARCHIVE</li>
//
// When the current storage type of a media file is infrequent storage, it can be modified to the following types:
//
// <li>Standard storage</li>
//
// <li>Archive storage</li>
//
// <li>DEEP_ARCHIVE</li>
//
// When the current storage type of a media file is archive storage, it can be modified to the following types:
//
// <li>Standard storage</li>
//
// When the current storage type of a media file is DEEP_ARCHIVE, it can be modified to the following types:
//
// <li>Standard storage</li>
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyMediaStorageClass")
    
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
// This API is used to modify material sample info based on the material ID, including modification of the name and description, as well as addition, deletion, and reset of facial features and tags. Ensure at least 1 image remains after facial feature deletion. Otherwise, use reset.
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
// This API is used to modify material sample info based on the material ID, including modification of the name and description, as well as addition, deletion, and reset of facial features and tags. Ensure at least 1 image remains after facial feature deletion. Otherwise, use reset.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyPersonSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPersonSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPersonSampleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProcessImageAsyncTemplateRequest() (request *ModifyProcessImageAsyncTemplateRequest) {
    request = &ModifyProcessImageAsyncTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyProcessImageAsyncTemplate")
    
    
    return
}

func NewModifyProcessImageAsyncTemplateResponse() (response *ModifyProcessImageAsyncTemplateResponse) {
    response = &ModifyProcessImageAsyncTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProcessImageAsyncTemplate
// This API is used to modify a user-customized asynchronous image processing template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and are not allowed to be modified.
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
func (c *Client) ModifyProcessImageAsyncTemplate(request *ModifyProcessImageAsyncTemplateRequest) (response *ModifyProcessImageAsyncTemplateResponse, err error) {
    return c.ModifyProcessImageAsyncTemplateWithContext(context.Background(), request)
}

// ModifyProcessImageAsyncTemplate
// This API is used to modify a user-customized asynchronous image processing template.
//
// 
//
// Note: Templates with IDs below 10000 are preset templates and are not allowed to be modified.
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
func (c *Client) ModifyProcessImageAsyncTemplateWithContext(ctx context.Context, request *ModifyProcessImageAsyncTemplateRequest) (response *ModifyProcessImageAsyncTemplateResponse, err error) {
    if request == nil {
        request = NewModifyProcessImageAsyncTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyProcessImageAsyncTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProcessImageAsyncTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProcessImageAsyncTemplateResponse()
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
// This API is used to modify an audio and video quality detection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyQualityInspectTemplate(request *ModifyQualityInspectTemplateRequest) (response *ModifyQualityInspectTemplateResponse, err error) {
    return c.ModifyQualityInspectTemplateWithContext(context.Background(), request)
}

// ModifyQualityInspectTemplate
// This API is used to modify an audio and video quality detection template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyQualityInspectTemplateWithContext(ctx context.Context, request *ModifyQualityInspectTemplateRequest) (response *ModifyQualityInspectTemplateResponse, err error) {
    if request == nil {
        request = NewModifyQualityInspectTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyQualityInspectTemplate")
    
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
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Modify a video rebirth template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyRebuildMediaTemplate(request *ModifyRebuildMediaTemplateRequest) (response *ModifyRebuildMediaTemplateResponse, err error) {
    return c.ModifyRebuildMediaTemplateWithContext(context.Background(), request)
}

// ModifyRebuildMediaTemplate
// This API is <font color=red>no longer maintained</font>. The new version of the [audio and video quality revival](https://www.tencentcloud.com/document/product/266/102571?from_cn_redirect=1) API uses preset templates. For details, see [Audio and Video Quality Rebirth Template](https://www.tencentcloud.com/document/product/266/102586?from_cn_redirect=1#50604b3f-0286-4a10-a3f7-18218116aff7).
//
// Modify a video rebirth template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) ModifyRebuildMediaTemplateWithContext(ctx context.Context, request *ModifyRebuildMediaTemplateRequest) (response *ModifyRebuildMediaTemplateResponse, err error) {
    if request == nil {
        request = NewModifyRebuildMediaTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyRebuildMediaTemplate")
    
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
// Modify a user-customized moderation template.
//
// >Template is applicable only to the [audio/video moderation (ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) and [image moderation (ReviewImage)](https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) APIs.
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
// Modify a user-customized moderation template.
//
// >Template is applicable only to the [audio/video moderation (ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) and [image moderation (ReviewImage)](https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) APIs.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyReviewTemplate")
    
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
// This API is used to modify a carousel playlist.
//
// After modification, only new playback requests will take effect. Users already playing can still play the playlist before modification within seven days.
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
// This API is used to modify a carousel playlist.
//
// After modification, only new playback requests will take effect. Users already playing can still play the playlist before modification within seven days.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyRoundPlay")
    
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
// Modify a custom sampled screenshot template.
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
// Modify a custom sampled screenshot template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifySampleSnapshotTemplate")
    
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
// Modify a user-customized specified time point screenshot template.
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
// Modify a user-customized specified time point screenshot template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifySnapshotByTimeOffsetTemplate")
    
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
// This API is used to change application information, but the default application information is not allowed to be modified.
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
// This API is used to change application information, but the default application information is not allowed to be modified.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifySubAppIdInfo")
    
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
// This API is used to enable or deactivate applications. Deactivated applications will have their corresponding domains blocked and console access restricted.
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
// This API is used to enable or deactivate applications. Deactivated applications will have their corresponding domains blocked and console access restricted.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifySubAppIdStatus")
    
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
// This API is <font color='red'>no longer maintained</font>. The new version of player signature no longer uses player configuration templates. For details, please see [Player Signature](https://www.tencentcloud.com/document/product/266/45554?from_cn_redirect=1).
//
// Modifies a player configuration.
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
// This API is <font color='red'>no longer maintained</font>. The new version of player signature no longer uses player configuration templates. For details, please see [Player Signature](https://www.tencentcloud.com/document/product/266/45554?from_cn_redirect=1).
//
// Modifies a player configuration.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifySuperPlayerConfig")
    
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
// Modify the information of a custom transcoding template.
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
// Modify the information of a custom transcoding template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyTranscodeTemplate")
    
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
// This API is used to modify the acceleration region of a VOD domain.
//
// 1. The acceleration region can be modified only when the domain name deployment status is Online.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVodDomainAccelerateConfig(request *ModifyVodDomainAccelerateConfigRequest) (response *ModifyVodDomainAccelerateConfigResponse, err error) {
    return c.ModifyVodDomainAccelerateConfigWithContext(context.Background(), request)
}

// ModifyVodDomainAccelerateConfig
// This API is used to modify the acceleration region of a VOD domain.
//
// 1. The acceleration region can be modified only when the domain name deployment status is Online.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVodDomainAccelerateConfigWithContext(ctx context.Context, request *ModifyVodDomainAccelerateConfigRequest) (response *ModifyVodDomainAccelerateConfigResponse, err error) {
    if request == nil {
        request = NewModifyVodDomainAccelerateConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyVodDomainAccelerateConfig")
    
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
// This API is used to modify domain name configuration, including hotlink protection configuration.
//
// 1. The domain name configuration can be modified only when the deployment state is Online.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVodDomainConfig(request *ModifyVodDomainConfigRequest) (response *ModifyVodDomainConfigResponse, err error) {
    return c.ModifyVodDomainConfigWithContext(context.Background(), request)
}

// ModifyVodDomainConfig
// This API is used to modify domain name configuration, including hotlink protection configuration.
//
// 1. The domain name configuration can be modified only when the deployment state is Online.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINDEPLOYING = "FailedOperation.DomainDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVodDomainConfigWithContext(ctx context.Context, request *ModifyVodDomainConfigRequest) (response *ModifyVodDomainConfigResponse, err error) {
    if request == nil {
        request = NewModifyVodDomainConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyVodDomainConfig")
    
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
// This API is used to modify a user-defined watermark template. The watermark type cannot be modified.
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
// This API is used to modify a user-defined watermark template. The watermark type cannot be modified.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyWatermarkTemplate")
    
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
// This API is used to modify the scenario and tags of a keyword. The keyword itself cannot be modified. If modification is needed, delete and rebuild it.
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
// This API is used to modify the scenario and tags of a keyword. The keyword itself cannot be modified. If modification is needed, delete and rebuild it.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ModifyWordSample")
    
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
// When uploading HLS videos, this API parses the index file content and returns the list of shard files to be uploaded. The sharded file path must be a relative path in the current directory or subdirectory. It cannot be a URL or an absolute path.
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
// When uploading HLS videos, this API parses the index file content and returns the list of shard files to be uploaded. The sharded file path must be a relative path in the current directory or subdirectory. It cannot be a URL or an absolute path.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ParseStreamingManifest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseStreamingManifest require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseStreamingManifestResponse()
    err = c.Send(request, response)
    return
}

func NewProcessImageAsyncRequest() (request *ProcessImageAsyncRequest) {
    request = &ProcessImageAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ProcessImageAsync")
    
    
    return
}

func NewProcessImageAsyncResponse() (response *ProcessImageAsyncResponse) {
    response = &ProcessImageAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessImageAsync
// This API is used to process image tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessImageAsync(request *ProcessImageAsyncRequest) (response *ProcessImageAsyncResponse, err error) {
    return c.ProcessImageAsyncWithContext(context.Background(), request)
}

// ProcessImageAsync
// This API is used to process image tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessImageAsyncWithContext(ctx context.Context, request *ProcessImageAsyncRequest) (response *ProcessImageAsyncResponse, err error) {
    if request == nil {
        request = NewProcessImageAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ProcessImageAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessImageAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessImageAsyncResponse()
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
// This API is used to initiate processing tasks for audio-video media in VOD, with features including:
//
// 1. Watermarked video transcoding;
//
// 2. Animated image generating;
//
// 3. Screenshot taking at specified time points;
//
// 4. Sampled screenshot taking;
//
// 5. Capture CSS sprites for videos;
//
// 6. Capture a frame from a video as the cover.
//
// 7. Transcoding to adaptive bitrate streaming with encryption;
//
// 8. Content review (offensive content, unsafe information, inappropriate information). It is <font color=red>not recommended</font> to use this API to initiate it. [Audio/Video Moderation (ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) or [Image Moderation (ReviewImage)](https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) is recommended;
//
// 9. Content analysis (tag, category, cover, frame tagging) is not supported for HLS format currently.
//
// 10. Content recognition (video intro and outro, human face, full text, text keyword, full speech, speech keyword, object).
//
// 
//
// If event notification is used, its type is task flow status change (https://www.tencentcloud.com/document/product/266/9636?from_cn_redirect=1).
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
// This API is used to initiate processing tasks for audio-video media in VOD, with features including:
//
// 1. Watermarked video transcoding;
//
// 2. Animated image generating;
//
// 3. Screenshot taking at specified time points;
//
// 4. Sampled screenshot taking;
//
// 5. Capture CSS sprites for videos;
//
// 6. Capture a frame from a video as the cover.
//
// 7. Transcoding to adaptive bitrate streaming with encryption;
//
// 8. Content review (offensive content, unsafe information, inappropriate information). It is <font color=red>not recommended</font> to use this API to initiate it. [Audio/Video Moderation (ReviewAudioVideo)](https://www.tencentcloud.com/document/api/266/80283?from_cn_redirect=1) or [Image Moderation (ReviewImage)](https://www.tencentcloud.com/document/api/266/73217?from_cn_redirect=1) is recommended;
//
// 9. Content analysis (tag, category, cover, frame tagging) is not supported for HLS format currently.
//
// 10. Content recognition (video intro and outro, human face, full text, text keyword, full speech, speech keyword, object).
//
// 
//
// If event notification is used, its type is task flow status change (https://www.tencentcloud.com/document/product/266/9636?from_cn_redirect=1).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ProcessMedia")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessMediaResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaByMPSRequest() (request *ProcessMediaByMPSRequest) {
    request = &ProcessMediaByMPSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ProcessMediaByMPS")
    
    
    return
}

func NewProcessMediaByMPSResponse() (response *ProcessMediaByMPSResponse) {
    response = &ProcessMediaByMPSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessMediaByMPS
// Use the media processing capacity of the media processing service (MPS) to initiate media processing for videos in video-on-demand.
//
// Currently supported MPS features:
//
// 1. Smart subtitling: This feature supports processing offline audio files, video files, and live streams. It can extract subtitles in the video source language through ASR speech recognition or OCR text recognition, and implement multilingual translation. View details in the integration guide (https://www.tencentcloud.com/document/product/266/131210?from_cn_redirect=1).
//
// 2. Intelligent erasure: It can blur, mosaic, or seamlessly process elements such as logos, subtitles, human faces, and license plates in video footage, making it easy to spread and share content. The new video generated by this task will be assigned a new FileId and stored in a sub-application of the VOD platform. View details in the [Access Guide](https://www.tencentcloud.com/document/product/266/131211?from_cn_redirect=1).
//
// 3. AI analysis: This feature supports all-in-one translation (https://www.tencentcloud.com/document/product/266/131212?from_cn_redirect=1), highlights (https://www.tencentcloud.com/document/product/266/131213?from_cn_redirect=1), LLM video summary (https://www.tencentcloud.com/document/product/266/131214?from_cn_redirect=1), LLM audio/video understanding (https://www.tencentcloud.com/document/product/266/131215?from_cn_redirect=1), intelligent splitting (https://www.tencentcloud.com/document/product/266/131216?from_cn_redirect=1), intelligent landscape-to-portrait (https://www.tencentcloud.com/document/product/266/131217?from_cn_redirect=1), video deduplication (https://www.tencentcloud.com/document/product/266/131218?from_cn_redirect=1), and other features.
//
// 
//
// 
//
// > Video processing task initiated this method:
//
// > 1. Query of task status and results is still completed in the VOD platform. Use [DescribeTaskDetail](https://www.tencentcloud.com/document/product/266/33431?from_cn_redirect=1) or [DescribeTasks](https://www.tencentcloud.com/document/product/266/33430?from_cn_redirect=1) to query tasks.
//
// > 2. The amount and bills of related features will be provided on the PS platform. Before using this feature, start by enabling Media Processing Service (MPS) in the console. For the activation method, see the preliminary operations in the integration guide.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessMediaByMPS(request *ProcessMediaByMPSRequest) (response *ProcessMediaByMPSResponse, err error) {
    return c.ProcessMediaByMPSWithContext(context.Background(), request)
}

// ProcessMediaByMPS
// Use the media processing capacity of the media processing service (MPS) to initiate media processing for videos in video-on-demand.
//
// Currently supported MPS features:
//
// 1. Smart subtitling: This feature supports processing offline audio files, video files, and live streams. It can extract subtitles in the video source language through ASR speech recognition or OCR text recognition, and implement multilingual translation. View details in the integration guide (https://www.tencentcloud.com/document/product/266/131210?from_cn_redirect=1).
//
// 2. Intelligent erasure: It can blur, mosaic, or seamlessly process elements such as logos, subtitles, human faces, and license plates in video footage, making it easy to spread and share content. The new video generated by this task will be assigned a new FileId and stored in a sub-application of the VOD platform. View details in the [Access Guide](https://www.tencentcloud.com/document/product/266/131211?from_cn_redirect=1).
//
// 3. AI analysis: This feature supports all-in-one translation (https://www.tencentcloud.com/document/product/266/131212?from_cn_redirect=1), highlights (https://www.tencentcloud.com/document/product/266/131213?from_cn_redirect=1), LLM video summary (https://www.tencentcloud.com/document/product/266/131214?from_cn_redirect=1), LLM audio/video understanding (https://www.tencentcloud.com/document/product/266/131215?from_cn_redirect=1), intelligent splitting (https://www.tencentcloud.com/document/product/266/131216?from_cn_redirect=1), intelligent landscape-to-portrait (https://www.tencentcloud.com/document/product/266/131217?from_cn_redirect=1), video deduplication (https://www.tencentcloud.com/document/product/266/131218?from_cn_redirect=1), and other features.
//
// 
//
// 
//
// > Video processing task initiated this method:
//
// > 1. Query of task status and results is still completed in the VOD platform. Use [DescribeTaskDetail](https://www.tencentcloud.com/document/product/266/33431?from_cn_redirect=1) or [DescribeTasks](https://www.tencentcloud.com/document/product/266/33430?from_cn_redirect=1) to query tasks.
//
// > 2. The amount and bills of related features will be provided on the PS platform. Before using this feature, start by enabling Media Processing Service (MPS) in the console. For the activation method, see the preliminary operations in the integration guide.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FILEID = "InvalidParameterValue.FileId"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
//  INVALIDPARAMETERVALUE_SUBAPPID = "InvalidParameterValue.SubAppId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ProcessMediaByMPSWithContext(ctx context.Context, request *ProcessMediaByMPSRequest) (response *ProcessMediaByMPSResponse, err error) {
    if request == nil {
        request = NewProcessMediaByMPSRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ProcessMediaByMPS")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMediaByMPS require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessMediaByMPSResponse()
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
// Use a task flow template to initiate processing tasks for videos in VOD.
//
// There are two ways to create a task flow template:
//
// 1. Create and modify a task flow template in the console;
//
// 2. Create a task flow template through the task flow template API.
//
// 
//
// For event notification, the type of event notifications other than audio/video moderation tasks is [task flow status change](https://www.tencentcloud.com/document/product/266/9636?from_cn_redirect=1); the type of audio/video moderation task event notification is [audio/video moderation completed](https://www.tencentcloud.com/document/product/266/81258?from_cn_redirect=1).
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
// Use a task flow template to initiate processing tasks for videos in VOD.
//
// There are two ways to create a task flow template:
//
// 1. Create and modify a task flow template in the console;
//
// 2. Create a task flow template through the task flow template API.
//
// 
//
// For event notification, the type of event notifications other than audio/video moderation tasks is [task flow status change](https://www.tencentcloud.com/document/product/266/9636?from_cn_redirect=1); the type of audio/video moderation task event notification is [audio/video moderation completed](https://www.tencentcloud.com/document/product/266/81258?from_cn_redirect=1).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ProcessMediaByProcedure")
    
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
// This API is <font color='red'>no longer maintained</font>. Please use the [ProcessMedia](https://www.tencentcloud.com/document/product/862/37578?from_cn_redirect=1) API of MPS and specify the video URL in the input parameter InputInfo.UrlInputInfo.Url.
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
// This API is <font color='red'>no longer maintained</font>. Please use the [ProcessMedia](https://www.tencentcloud.com/document/product/862/37578?from_cn_redirect=1) API of MPS and specify the video URL in the input parameter InputInfo.UrlInputInfo.Url.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ProcessMediaByUrl")
    
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
// * This API is used for business servers to get event notifications via reliable callback (https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1#.E5.8F.AF.E9.9D.A0.E5.9B.9E.E8.B0.83).
//
// * The API uses long polling mode. If there are unconsumed events on the server, they will be returned to the requester immediately. If there are no unconsumed events, the request will be suspended in the backend until a new event occurs.
//
// * The request can be suspended for up to 5 seconds. It is advisable to set the timeout to 10 seconds for the requester.
//
// * Event notifications that are not pulled are retained for up to 4 days. Notifications exceeding this time limit may be purged.
//
// * If this API returns an event, the caller must call the [Confirm Event Notification](https://www.tencentcloud.com/document/product/266/33434?from_cn_redirect=1) API within <font color="red">30 seconds</font> to confirm that the event notification has been processed. Otherwise, the event notification will be pulled again after <font color="red">30 seconds</font>.
//
// * Currently, a maximum of 16 event notifications can be obtained per API call.
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
// * This API is used for business servers to get event notifications via reliable callback (https://www.tencentcloud.com/document/product/266/33779?from_cn_redirect=1#.E5.8F.AF.E9.9D.A0.E5.9B.9E.E8.B0.83).
//
// * The API uses long polling mode. If there are unconsumed events on the server, they will be returned to the requester immediately. If there are no unconsumed events, the request will be suspended in the backend until a new event occurs.
//
// * The request can be suspended for up to 5 seconds. It is advisable to set the timeout to 10 seconds for the requester.
//
// * Event notifications that are not pulled are retained for up to 4 days. Notifications exceeding this time limit may be purged.
//
// * If this API returns an event, the caller must call the [Confirm Event Notification](https://www.tencentcloud.com/document/product/266/33434?from_cn_redirect=1) API within <font color="red">30 seconds</font> to confirm that the event notification has been processed. Otherwise, the event notification will be pulled again after <font color="red">30 seconds</font>.
//
// * Currently, a maximum of 16 event notifications can be obtained per API call.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "PullEvents")
    
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
// This API is used to pull a video from the network to the VOD platform.
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
//  INVALIDPARAMETERVALUE_MEDIASTORAGEPATH = "InvalidParameterValue.MediaStoragePath"
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
// This API is used to pull a video from the network to the VOD platform.
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
//  INVALIDPARAMETERVALUE_MEDIASTORAGEPATH = "InvalidParameterValue.MediaStoragePath"
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "PullUpload")
    
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
// 1. Preheats a specified URL list.
//
// 2. The domain name of the URL must be registered in VOD.
//
// 3. You can specify up to 20 URLs per request.
//
// 4. The default prefetch quota is 10,000 URLs per day.
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
// 1. Preheats a specified URL list.
//
// 2. The domain name of the URL must be registered in VOD.
//
// 3. You can specify up to 20 URLs per request.
//
// 4. The default prefetch quota is 10,000 URLs per day.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "PushUrlCache")
    
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
// This API is <font color=red>no longer maintained</font>. Please use the new version of APIs [audio and video quality revival](https://www.tencentcloud.com/document/api/266/102571?from_cn_redirect=1).
//
// This API is used to initiate audio and video quality revival.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) RebuildMedia(request *RebuildMediaRequest) (response *RebuildMediaResponse, err error) {
    return c.RebuildMediaWithContext(context.Background(), request)
}

// RebuildMedia
// This API is <font color=red>no longer maintained</font>. Please use the new version of APIs [audio and video quality revival](https://www.tencentcloud.com/document/api/266/102571?from_cn_redirect=1).
//
// This API is used to initiate audio and video quality revival.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) RebuildMediaWithContext(ctx context.Context, request *RebuildMediaRequest) (response *RebuildMediaResponse, err error) {
    if request == nil {
        request = NewRebuildMediaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "RebuildMedia")
    
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
// This API is <font color=red>no longer maintained</font>. Please use the new version of APIs for [audio and video quality revival](https://www.tencentcloud.com/document/api/266/102571?from_cn_redirect=1).
//
// Use a template to initiate video rebirth.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) RebuildMediaByTemplate(request *RebuildMediaByTemplateRequest) (response *RebuildMediaByTemplateResponse, err error) {
    return c.RebuildMediaByTemplateWithContext(context.Background(), request)
}

// RebuildMediaByTemplate
// This API is <font color=red>no longer maintained</font>. Please use the new version of APIs for [audio and video quality revival](https://www.tencentcloud.com/document/api/266/102571?from_cn_redirect=1).
//
// Use a template to initiate video rebirth.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
func (c *Client) RebuildMediaByTemplateWithContext(ctx context.Context, request *RebuildMediaByTemplateRequest) (response *RebuildMediaByTemplateResponse, err error) {
    if request == nil {
        request = NewRebuildMediaByTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "RebuildMediaByTemplate")
    
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
// 1. Refresh a specified URL list.
//
// 2. The domain name of the URL must be registered in VOD.
//
// 3. You can specify up to 20 URLs per request.
//
// 4. The default refresh quota is 100,000 URLs per day.
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
// 1. Refresh a specified URL list.
//
// 2. The domain name of the URL must be registered in VOD.
//
// 3. You can specify up to 20 URLs per request.
//
// 4. The default refresh quota is 100,000 URLs per day.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "RefreshUrlCache")
    
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
// Watermark removal
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
// Watermark removal
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "RemoveWatermark")
    
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
// Reset the content of the user-defined task flow template.
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
// Reset the content of the user-defined task flow template.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ResetProcedureTemplate")
    
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
// When the storage type of a media file is archive storage or deep archive storage, it is unreachable. If you need access, call this API to unfreeze it. The unfrozen media file is temporarily accessible and becomes unreachable after the validity period expires.
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
// When the storage type of a media file is archive storage or deep archive storage, it is unreachable. If you need access, call this API to unfreeze it. The unfrozen media file is temporarily accessible and becomes unreachable after the validity period expires.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "RestoreMedia")
    
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
// This API is used to initiate moderation tasks for on-demand audio-video media, intelligently detecting violative content in video footage, text in visuals, text in speech, and sound.
//
// 
//
// If event notification is used, its type is audio/video moderation completed (https://www.tencentcloud.com/document/product/266/81258?from_cn_redirect=1).
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
// This API is used to initiate moderation tasks for on-demand audio-video media, intelligently detecting violative content in video footage, text in visuals, text in speech, and sound.
//
// 
//
// If event notification is used, its type is audio/video moderation completed (https://www.tencentcloud.com/document/product/266/81258?from_cn_redirect=1).
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ReviewAudioVideo")
    
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
// This API is used to initiate review tasks for image files in on-demand video, including offensive, unsafe, and inappropriate content.
//
// 
//
// ><<li>Supported image file size: file < 5M;</li>
//
// ><<li>Supported image file resolution: recommended resolution above 256x256, otherwise review effectiveness may be affected;</li>
//
// ><<li>Supported image file formats: PNG, JPG, JPEG, BMP, GIF, WEBP.</li>
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
// This API is used to initiate review tasks for image files in on-demand video, including offensive, unsafe, and inappropriate content.
//
// 
//
// ><<li>Supported image file size: file < 5M;</li>
//
// ><<li>Supported image file resolution: recommended resolution above 256x256, otherwise review effectiveness may be affected;</li>
//
// ><<li>Supported image file formats: PNG, JPG, JPEG, BMP, GIF, WEBP.</li>
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "ReviewImage")
    
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
// This API is used to search media information with conditional filtering, sort and filter returned results, and other features. This includes:
//
// -Specify the file ID collection FileIds to return media matching any ID in the collection.
//
// -Perform fuzzy search by multiple media file names (Names) or descriptions (Descriptions).
//
// -Search by multiple filename prefixes NamePrefixes.
//
// -Specify the category collection ClassIds (see input parameters), and media that meet any category in the collection will be returned. For example, media categories include movie, TV series, and variety show. The movie category has subcategories such as historical film, action film, and romance film. If ClassIds specifies movie and TV series, all subcategories under movie and TV series will be returned. If ClassIds specifies historical film and action film, only media under these two subcategories will be returned.
//
// - Specify tag collection Tags (see input parameter) to return media that match any tag in the collection. For example, if media tags include ACG, palace intrigue, and parody remix, and Tags specifies ACG and parody remix, then media matching any one of these two tags will be retrieved.
//
// -Specify the file type set Categories (see input parameters), and return media that meet any type in the collection. For example, file types include Video, Audio, and Image. If Categories specifies Video and Audio, media that meet these types will be retrieved.
//
// -Specify the source collection SourceTypes (see input parameters) to return media that meets any source in the collection. For example, media sources include Record (live recording), Upload, and so on. If SourceTypes specifies Record and Upload, media that meets these sources will be retrieved.
//
// -Specify the file packaging format set MediaTypes (see input parameters) to return media that meets any packaging format in the collection. For example, packaging formats include MP4, AVI, MP3, and so on. If MediaTypes specifies MP4 and MP3, media that complies with these packaging formats will be retrieved.
//
// -Specify the file status collection Status (see input parameters), and return media that meets any status in the collection. For example, file statuses include Normal, SystemForbidden (Platform Ban), and Forbidden (proactive ban). If Status specifies Normal and Forbidden, media that meets these statuses will be retrieved.
//
// -Specify the file review result set ReviewResults (see input parameters) to return media that meets any status in the collection. For example, file review results include pass and block. If ReviewResults specifies both pass and block, media that meets these review results will be retrieved.
//
// -Filter media for live streaming recording by specifying the collection of live streaming codes StreamIds (see input parameters).
//
// -Filter media by the creation time range of the specified media.
//
// -Specify a TRTC application ID collection to filter media.
//
// -Specify a TRTC room ID collection to filter media.
//
// 
//
// -The above parameters can be combined in any way for search. For example, filter media with a creation time between 2018-12-01 12:00:00 and 2018-12-08 12:00:00, categorized as movie or TV series, and tagged with palace intrigue and suspense. Note that for any parameter that supports array input, the search logic between its elements is OR. The logical relationship between all parameters is AND.
//
// 
//
// - Allow controlling the type of media information returned through Filters (return all information by default). Options include:
//
// 1. basicInfo: including media name, category, playback address, cover image, and more.
//
// 2. Meta information (metaData): including size, duration, video stream information, and audio stream information.
//
// 3. transcodeInfo: includes media addresses, video stream parameters, audio stream parameters, and more for various specifications generated by transcoding this media.
//
// 4. Animated graphics info (animatedGraphicsInfo): the animated graphics info after converting a video to gif (for example, gif).
//
// 5. sampleSnapshotInfo: screenshot information after sampling screenshot taking from a video.
//
// 6. Sprite image information (imageSpriteInfo): sprite image information of the captured sprite image file from the video.
//
// 7. snapshotByTimeOffsetInfo: screenshot information after taking screenshots at specified time points.
//
// 8. Video timestamp information (keyFrameDescInfo): dotting information set for the video.
//
// 9. Adaptive Bitstreaming information (adaptiveDynamicStreamingInfo): includes specification, encryption type, packaging format and other related information.
//
// 
//
// -Allow sorting results by creation time and return in pages. Use Offset and Limit (see input parameters) to control pagination.
//
// 
//
// <div id="maxResultsDesc">API return result count limit:</div>
//
// 
//
// -<b><a href="#p_offset">Offset</a> and <a href="#p_limit">Limit</a> impact the number of results per pagination query. Special attention: when both values are omitted, this interface returns up to 10 query results only.</b>
//
// -<b>Supports up to 5000 search results. Excess results are no longer queryable. If the search result volume is too large, use more granular criteria to reduce the results.</b>
//
// 
//
// <br>Not recommended conditional filtering:
//
// -(Not recommended: use Names, NamePrefixes, or Descriptions instead) Specify a single Text to do fuzzy search on media file Names or Descriptions.
//
// -(Not recommended: use SourceTypes instead) Specify a single media file source SourceType for search.
//
// -(Not recommended: use StreamIds instead) Specify a single push stream live code StreamId to search.
//
// -(Not recommended: use CreateTime instead) Specify a single start creation time StartTime to search.
//
// -(Not recommended: use CreateTime instead) Specify a single end time EndTime for search.
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
// This API is used to search media information with conditional filtering, sort and filter returned results, and other features. This includes:
//
// -Specify the file ID collection FileIds to return media matching any ID in the collection.
//
// -Perform fuzzy search by multiple media file names (Names) or descriptions (Descriptions).
//
// -Search by multiple filename prefixes NamePrefixes.
//
// -Specify the category collection ClassIds (see input parameters), and media that meet any category in the collection will be returned. For example, media categories include movie, TV series, and variety show. The movie category has subcategories such as historical film, action film, and romance film. If ClassIds specifies movie and TV series, all subcategories under movie and TV series will be returned. If ClassIds specifies historical film and action film, only media under these two subcategories will be returned.
//
// - Specify tag collection Tags (see input parameter) to return media that match any tag in the collection. For example, if media tags include ACG, palace intrigue, and parody remix, and Tags specifies ACG and parody remix, then media matching any one of these two tags will be retrieved.
//
// -Specify the file type set Categories (see input parameters), and return media that meet any type in the collection. For example, file types include Video, Audio, and Image. If Categories specifies Video and Audio, media that meet these types will be retrieved.
//
// -Specify the source collection SourceTypes (see input parameters) to return media that meets any source in the collection. For example, media sources include Record (live recording), Upload, and so on. If SourceTypes specifies Record and Upload, media that meets these sources will be retrieved.
//
// -Specify the file packaging format set MediaTypes (see input parameters) to return media that meets any packaging format in the collection. For example, packaging formats include MP4, AVI, MP3, and so on. If MediaTypes specifies MP4 and MP3, media that complies with these packaging formats will be retrieved.
//
// -Specify the file status collection Status (see input parameters), and return media that meets any status in the collection. For example, file statuses include Normal, SystemForbidden (Platform Ban), and Forbidden (proactive ban). If Status specifies Normal and Forbidden, media that meets these statuses will be retrieved.
//
// -Specify the file review result set ReviewResults (see input parameters) to return media that meets any status in the collection. For example, file review results include pass and block. If ReviewResults specifies both pass and block, media that meets these review results will be retrieved.
//
// -Filter media for live streaming recording by specifying the collection of live streaming codes StreamIds (see input parameters).
//
// -Filter media by the creation time range of the specified media.
//
// -Specify a TRTC application ID collection to filter media.
//
// -Specify a TRTC room ID collection to filter media.
//
// 
//
// -The above parameters can be combined in any way for search. For example, filter media with a creation time between 2018-12-01 12:00:00 and 2018-12-08 12:00:00, categorized as movie or TV series, and tagged with palace intrigue and suspense. Note that for any parameter that supports array input, the search logic between its elements is OR. The logical relationship between all parameters is AND.
//
// 
//
// - Allow controlling the type of media information returned through Filters (return all information by default). Options include:
//
// 1. basicInfo: including media name, category, playback address, cover image, and more.
//
// 2. Meta information (metaData): including size, duration, video stream information, and audio stream information.
//
// 3. transcodeInfo: includes media addresses, video stream parameters, audio stream parameters, and more for various specifications generated by transcoding this media.
//
// 4. Animated graphics info (animatedGraphicsInfo): the animated graphics info after converting a video to gif (for example, gif).
//
// 5. sampleSnapshotInfo: screenshot information after sampling screenshot taking from a video.
//
// 6. Sprite image information (imageSpriteInfo): sprite image information of the captured sprite image file from the video.
//
// 7. snapshotByTimeOffsetInfo: screenshot information after taking screenshots at specified time points.
//
// 8. Video timestamp information (keyFrameDescInfo): dotting information set for the video.
//
// 9. Adaptive Bitstreaming information (adaptiveDynamicStreamingInfo): includes specification, encryption type, packaging format and other related information.
//
// 
//
// -Allow sorting results by creation time and return in pages. Use Offset and Limit (see input parameters) to control pagination.
//
// 
//
// <div id="maxResultsDesc">API return result count limit:</div>
//
// 
//
// -<b><a href="#p_offset">Offset</a> and <a href="#p_limit">Limit</a> impact the number of results per pagination query. Special attention: when both values are omitted, this interface returns up to 10 query results only.</b>
//
// -<b>Supports up to 5000 search results. Excess results are no longer queryable. If the search result volume is too large, use more granular criteria to reduce the results.</b>
//
// 
//
// <br>Not recommended conditional filtering:
//
// -(Not recommended: use Names, NamePrefixes, or Descriptions instead) Specify a single Text to do fuzzy search on media file Names or Descriptions.
//
// -(Not recommended: use SourceTypes instead) Specify a single media file source SourceType for search.
//
// -(Not recommended: use StreamIds instead) Specify a single push stream live code StreamId to search.
//
// -(Not recommended: use CreateTime instead) Specify a single start creation time StartTime to search.
//
// -(Not recommended: use CreateTime instead) Specify a single end time EndTime for search.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "SearchMedia")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchMedia require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchMediaResponse()
    err = c.Send(request, response)
    return
}

func NewSearchMediaBySemanticsRequest() (request *SearchMediaBySemanticsRequest) {
    request = &SearchMediaBySemanticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "SearchMediaBySemantics")
    
    
    return
}

func NewSearchMediaBySemanticsResponse() (response *SearchMediaBySemanticsResponse) {
    response = &SearchMediaBySemanticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchMediaBySemantics
// Use natural language to conduct semantic search on media.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SearchMediaBySemantics(request *SearchMediaBySemanticsRequest) (response *SearchMediaBySemanticsResponse, err error) {
    return c.SearchMediaBySemanticsWithContext(context.Background(), request)
}

// SearchMediaBySemantics
// Use natural language to conduct semantic search on media.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SearchMediaBySemanticsWithContext(ctx context.Context, request *SearchMediaBySemanticsRequest) (response *SearchMediaBySemanticsResponse, err error) {
    if request == nil {
        request = NewSearchMediaBySemanticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "SearchMediaBySemantics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchMediaBySemantics require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchMediaBySemanticsResponse()
    err = c.Send(request, response)
    return
}

func NewSetCLSPushTargetRequest() (request *SetCLSPushTargetRequest) {
    request = &SetCLSPushTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "SetCLSPushTarget")
    
    
    return
}

func NewSetCLSPushTargetResponse() (response *SetCLSPushTargetResponse) {
    response = &SetCLSPushTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetCLSPushTarget
// Set a delivery destination for CLS for the vod domain.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SetCLSPushTarget(request *SetCLSPushTargetRequest) (response *SetCLSPushTargetResponse, err error) {
    return c.SetCLSPushTargetWithContext(context.Background(), request)
}

// SetCLSPushTarget
// Set a delivery destination for CLS for the vod domain.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SetCLSPushTargetWithContext(ctx context.Context, request *SetCLSPushTargetRequest) (response *SetCLSPushTargetResponse, err error) {
    if request == nil {
        request = NewSetCLSPushTargetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "SetCLSPushTarget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetCLSPushTarget require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetCLSPushTargetResponse()
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
// Sets DRM key provider information.
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
// Sets DRM key provider information.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "SetDrmKeyProviderInfo")
    
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
// Set the HTTPS certificate for the vod domain.
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
// Set the HTTPS certificate for the vod domain.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "SetVodDomainCertificate")
    
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
// Crop an HLS video by time period to generate a new HLS video in real time. Developers can share it immediately or save it for long-term preservation.
//
// 
//
// Tencent Cloud VOD supports two editing modes:
//
// -Clip and save: Save the edited video as an independent video with a separate FileId. This is suitable for long-term preservation of highlights.
//
// -Editing is not solidified: The edited video is attached to the input file with no standalone FileId, suitable for temporary sharing of highlight clips.
//
// 
//
// This API is used to crop an m3u8 file based on input. The minimum editing precision is one ts slice, so second-level or more precise editing precision cannot be achieved.
//
// 
//
// ### Edit solidification
//
// Editing and solidification refers to saving an edited video as an independent video with an independent FileId. Its lifecycle is not subject to any impact from the original input video. Even if the original input video is deleted, the clipping result is not affected. You can also transcode it or publish it on WeChat.
//
// 
//
// For example, a complete football match raw video may last for over 2 hours. For cost savings, a customer can store this video for 2 months, but can specify longer storage for the edited "highlights" video. You can also separately transcode, publish on WeChat, and perform other additional on-demand operations on the "highlights" video. In this case, you can choose the edit and solidify solution.
//
// 
//
// The advantage of solidified edits is that their lifecycle is independent of the original input video, allowing separate management and long-term preservation.
//
// 
//
// <font color='red'>Note:</font> If solidification is specified when editing, enable reception of editing solidification event notifications through the ModifyEventConfig API. After successful solidification, you will receive a PersistenceComplete event notification. Before receiving this event notification, you should not delete or reduce the storage class of the original input video. Otherwise, playback of the generated video may be abnormal.
//
// 
//
// ### Editing is not solidified
//
// Editing is not solidified, meaning the result of editing (m3u8 file) shares the same TS segments with the original input video. The generated video is not an independent and complete video (no standalone FileId, only a playback URL), and its valid period is consistent with that of the original input full video. Once the original input video is deleted, the clip will also become unplayable.
//
// 
//
// Editing is not solidified. Since the clipping result is not an independent video, it is not included in video management of on-demand media assets. For example, the total number of videos in the console does not include this video clip. You cannot separately perform any video processing operation on this clip, such as transcoding or publishing on WeChat.
//
// 
//
// The advantage of non-solidified editing is that the editing operation is relatively "lightweight" and will not generate additional storage overhead. However, its shortcoming is that the lifecycle is the same as the original recorded video, and it is unable to further transcode or perform other video processing.
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
// Crop an HLS video by time period to generate a new HLS video in real time. Developers can share it immediately or save it for long-term preservation.
//
// 
//
// Tencent Cloud VOD supports two editing modes:
//
// -Clip and save: Save the edited video as an independent video with a separate FileId. This is suitable for long-term preservation of highlights.
//
// -Editing is not solidified: The edited video is attached to the input file with no standalone FileId, suitable for temporary sharing of highlight clips.
//
// 
//
// This API is used to crop an m3u8 file based on input. The minimum editing precision is one ts slice, so second-level or more precise editing precision cannot be achieved.
//
// 
//
// ### Edit solidification
//
// Editing and solidification refers to saving an edited video as an independent video with an independent FileId. Its lifecycle is not subject to any impact from the original input video. Even if the original input video is deleted, the clipping result is not affected. You can also transcode it or publish it on WeChat.
//
// 
//
// For example, a complete football match raw video may last for over 2 hours. For cost savings, a customer can store this video for 2 months, but can specify longer storage for the edited "highlights" video. You can also separately transcode, publish on WeChat, and perform other additional on-demand operations on the "highlights" video. In this case, you can choose the edit and solidify solution.
//
// 
//
// The advantage of solidified edits is that their lifecycle is independent of the original input video, allowing separate management and long-term preservation.
//
// 
//
// <font color='red'>Note:</font> If solidification is specified when editing, enable reception of editing solidification event notifications through the ModifyEventConfig API. After successful solidification, you will receive a PersistenceComplete event notification. Before receiving this event notification, you should not delete or reduce the storage class of the original input video. Otherwise, playback of the generated video may be abnormal.
//
// 
//
// ### Editing is not solidified
//
// Editing is not solidified, meaning the result of editing (m3u8 file) shares the same TS segments with the original input video. The generated video is not an independent and complete video (no standalone FileId, only a playback URL), and its valid period is consistent with that of the original input full video. Once the original input video is deleted, the clip will also become unplayable.
//
// 
//
// Editing is not solidified. Since the clipping result is not an independent video, it is not included in video management of on-demand media assets. For example, the total number of videos in the console does not include this video clip. You cannot separately perform any video processing operation on this clip, such as transcoding or publishing on WeChat.
//
// 
//
// The advantage of non-solidified editing is that the editing operation is relatively "lightweight" and will not generate additional storage overhead. However, its shortcoming is that the lifecycle is the same as the original recorded video, and it is unable to further transcode or perform other video processing.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "SimpleHlsClip")
    
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
// This API is used to split an on-demand video into multiple new on-demand videos.
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
// This API is used to split an on-demand video into multiple new on-demand videos.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "SplitMedia")
    
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
// This API is used to enable or disable a CDN acceleration domain name.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StartCDNDomain(request *StartCDNDomainRequest) (response *StartCDNDomainResponse, err error) {
    return c.StartCDNDomainWithContext(context.Background(), request)
}

// StartCDNDomain
// This API is used to enable or disable a CDN acceleration domain name.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StartCDNDomainWithContext(ctx context.Context, request *StartCDNDomainRequest) (response *StartCDNDomainResponse, err error) {
    if request == nil {
        request = NewStartCDNDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "StartCDNDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCDNDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartCDNDomainResponse()
    err = c.Send(request, response)
    return
}

func NewTextToSpeechAsyncRequest() (request *TextToSpeechAsyncRequest) {
    request = &TextToSpeechAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "TextToSpeechAsync")
    
    
    return
}

func NewTextToSpeechAsyncResponse() (response *TextToSpeechAsyncResponse) {
    response = &TextToSpeechAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextToSpeechAsync
// Initiate a speech synthesis task to convert text to speech for long text scenarios (up to 200,000 characters). It supports specifying voice tone, speaking rate, volume, pitch, sampling rate, output format, and other synthesis parameters. Speech synthesis is an asynchronous task, and audio results are generated upon completion.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TextToSpeechAsync(request *TextToSpeechAsyncRequest) (response *TextToSpeechAsyncResponse, err error) {
    return c.TextToSpeechAsyncWithContext(context.Background(), request)
}

// TextToSpeechAsync
// Initiate a speech synthesis task to convert text to speech for long text scenarios (up to 200,000 characters). It supports specifying voice tone, speaking rate, volume, pitch, sampling rate, output format, and other synthesis parameters. Speech synthesis is an asynchronous task, and audio results are generated upon completion.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TextToSpeechAsyncWithContext(ctx context.Context, request *TextToSpeechAsyncRequest) (response *TextToSpeechAsyncResponse, err error) {
    if request == nil {
        request = NewTextToSpeechAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "TextToSpeechAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToSpeechAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToSpeechAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewTextToSpeechSyncRequest() (request *TextToSpeechSyncRequest) {
    request = &TextToSpeechSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "TextToSpeechSync")
    
    
    return
}

func NewTextToSpeechSyncResponse() (response *TextToSpeechSyncResponse) {
    response = &TextToSpeechSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextToSpeechSync
// This API is used to initiate a text to speech task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TextToSpeechSync(request *TextToSpeechSyncRequest) (response *TextToSpeechSyncResponse, err error) {
    return c.TextToSpeechSyncWithContext(context.Background(), request)
}

// TextToSpeechSync
// This API is used to initiate a text to speech task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TextToSpeechSyncWithContext(ctx context.Context, request *TextToSpeechSyncRequest) (response *TextToSpeechSyncResponse, err error) {
    if request == nil {
        request = NewTextToSpeechSyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "TextToSpeechSync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToSpeechSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToSpeechSyncResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAigcApiTokenRequest() (request *UpdateAigcApiTokenRequest) {
    request = &UpdateAigcApiTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "UpdateAigcApiToken")
    
    
    return
}

func NewUpdateAigcApiTokenResponse() (response *UpdateAigcApiTokenResponse) {
    response = &UpdateAigcApiTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAigcApiToken
// This API is used to create a Token for AIGC API calls. Data sync may delay once created. It can be queried or deleted after about 30 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateAigcApiToken(request *UpdateAigcApiTokenRequest) (response *UpdateAigcApiTokenResponse, err error) {
    return c.UpdateAigcApiTokenWithContext(context.Background(), request)
}

// UpdateAigcApiToken
// This API is used to create a Token for AIGC API calls. Data sync may delay once created. It can be queried or deleted after about 30 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateAigcApiTokenWithContext(ctx context.Context, request *UpdateAigcApiTokenRequest) (response *UpdateAigcApiTokenResponse, err error) {
    if request == nil {
        request = NewUpdateAigcApiTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "UpdateAigcApiToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAigcApiToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAigcApiTokenResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateVoiceRequest() (request *UpdateVoiceRequest) {
    request = &UpdateVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "UpdateVoice")
    
    
    return
}

func NewUpdateVoiceResponse() (response *UpdateVoiceResponse) {
    response = &UpdateVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateVoice
// This API is used to update the profile information of a voice type by voice ID, including name, description, gender, age, language, tag, and scenario. It returns the complete voice type information after the update. Only voice types for this account can be updated. System preset voice types do not support update.
//
// 
//
// Note: Newly designed or cloned voice types cannot be updated before activation. They are activated only after the new voice type is used for TTS once.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateVoice(request *UpdateVoiceRequest) (response *UpdateVoiceResponse, err error) {
    return c.UpdateVoiceWithContext(context.Background(), request)
}

// UpdateVoice
// This API is used to update the profile information of a voice type by voice ID, including name, description, gender, age, language, tag, and scenario. It returns the complete voice type information after the update. Only voice types for this account can be updated. System preset voice types do not support update.
//
// 
//
// Note: Newly designed or cloned voice types cannot be updated before activation. They are activated only after the new voice type is used for TTS once.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateVoiceWithContext(ctx context.Context, request *UpdateVoiceRequest) (response *UpdateVoiceResponse, err error) {
    if request == nil {
        request = NewUpdateVoiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "UpdateVoice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateVoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateVoiceResponse()
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
// This API is used to verify domain name resolution values.
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
// This API is used to verify domain name resolution values.
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
    c.InitBaseRequest(&request.BaseRequest, "vod", APIVersion, "VerifyDomainRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyDomainRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyDomainRecordResponse()
    err = c.Send(request, response)
    return
}
