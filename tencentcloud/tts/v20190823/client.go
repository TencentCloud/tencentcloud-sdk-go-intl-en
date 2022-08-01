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

package v20190823

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-08-23"

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


func NewTextToVoiceRequest() (request *TextToVoiceRequest) {
    request = &TextToVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tts", APIVersion, "TextToVoice")
    
    
    return
}

func NewTextToVoiceResponse() (response *TextToVoiceResponse) {
    response = &TextToVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextToVoice
// This API is used to convert any text to speech, allowing your devices and applications to talk to users.
//
// ​Tencent Cloud Text To Speech (TTS) can synthesize speech from text in real time for many use cases, such as audiobook and news apps, voice reminders on smart devices, quick synthesis of a celebrity's voice based on existing programs or certain voice records available on the Internet, and personalized vehicle navigation systems.
//
// It is free for use in beta.
//
// It supports SSML. For syntax details, see [SSML](https://intl.cloud.tencent.com/document/product/1073/49575?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_EXCEEDMAXLIMIT = "InternalError.ExceedMaxLimit"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDTEXT = "InvalidParameter.InvalidText"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_ERRORCARDINALFORMAT = "InvalidParameterValue.ErrorCardinalFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEXT = "InvalidParameterValue.InvalidText"
//  INVALIDPARAMETERVALUE_MISSPARAMETERS = "InvalidParameterValue.MissParameters"
//  INVALIDPARAMETERVALUE_PRIMARYLANGUAGE = "InvalidParameterValue.PrimaryLanguage"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SPEED = "InvalidParameterValue.Speed"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  INVALIDPARAMETERVALUE_TEXTEMPTY = "InvalidParameterValue.TextEmpty"
//  INVALIDPARAMETERVALUE_TEXTNOTUTF8 = "InvalidParameterValue.TextNotUtf8"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VOICETYPE = "InvalidParameterValue.VoiceType"
//  INVALIDPARAMETERVALUE_VOLUME = "InvalidParameterValue.Volume"
//  LIMITEXCEEDED_ACCESSLIMIT = "LimitExceeded.AccessLimit"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONEXPIRED = "UnsupportedOperation.AuthorizationExpired"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONFAILED = "UnsupportedOperation.AuthorizationFailed"
//  UNSUPPORTEDOPERATION_FORBIDDENUSE = "UnsupportedOperation.ForbiddenUse"
//  UNSUPPORTEDOPERATION_NOFREEACCOUNT = "UnsupportedOperation.NoFreeAccount"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) TextToVoice(request *TextToVoiceRequest) (response *TextToVoiceResponse, err error) {
    return c.TextToVoiceWithContext(context.Background(), request)
}

// TextToVoice
// This API is used to convert any text to speech, allowing your devices and applications to talk to users.
//
// ​Tencent Cloud Text To Speech (TTS) can synthesize speech from text in real time for many use cases, such as audiobook and news apps, voice reminders on smart devices, quick synthesis of a celebrity's voice based on existing programs or certain voice records available on the Internet, and personalized vehicle navigation systems.
//
// It is free for use in beta.
//
// It supports SSML. For syntax details, see [SSML](https://intl.cloud.tencent.com/document/product/1073/49575?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_EXCEEDMAXLIMIT = "InternalError.ExceedMaxLimit"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDTEXT = "InvalidParameter.InvalidText"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"
//  INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"
//  INVALIDPARAMETERVALUE_ERRORCARDINALFORMAT = "InvalidParameterValue.ErrorCardinalFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEXT = "InvalidParameterValue.InvalidText"
//  INVALIDPARAMETERVALUE_MISSPARAMETERS = "InvalidParameterValue.MissParameters"
//  INVALIDPARAMETERVALUE_PRIMARYLANGUAGE = "InvalidParameterValue.PrimaryLanguage"
//  INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SPEED = "InvalidParameterValue.Speed"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  INVALIDPARAMETERVALUE_TEXTEMPTY = "InvalidParameterValue.TextEmpty"
//  INVALIDPARAMETERVALUE_TEXTNOTUTF8 = "InvalidParameterValue.TextNotUtf8"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VOICETYPE = "InvalidParameterValue.VoiceType"
//  INVALIDPARAMETERVALUE_VOLUME = "InvalidParameterValue.Volume"
//  LIMITEXCEEDED_ACCESSLIMIT = "LimitExceeded.AccessLimit"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONEXPIRED = "UnsupportedOperation.AuthorizationExpired"
//  UNSUPPORTEDOPERATION_AUTHORIZATIONFAILED = "UnsupportedOperation.AuthorizationFailed"
//  UNSUPPORTEDOPERATION_FORBIDDENUSE = "UnsupportedOperation.ForbiddenUse"
//  UNSUPPORTEDOPERATION_NOFREEACCOUNT = "UnsupportedOperation.NoFreeAccount"
//  UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"
//  UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"
//  UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
func (c *Client) TextToVoiceWithContext(ctx context.Context, request *TextToVoiceRequest) (response *TextToVoiceResponse, err error) {
    if request == nil {
        request = NewTextToVoiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToVoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToVoiceResponse()
    err = c.Send(request, response)
    return
}
