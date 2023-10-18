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

package v20180301

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-03-01"

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


func NewApplyLivenessTokenRequest() (request *ApplyLivenessTokenRequest) {
    request = &ApplyLivenessTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "ApplyLivenessToken")
    
    
    return
}

func NewApplyLivenessTokenResponse() (response *ApplyLivenessTokenResponse) {
    response = &ApplyLivenessTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyLivenessToken
// This API is used to apply for a token before calling the liveness detection service each time. This token is required for initiating the verification process and getting the result after the verification is completed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ApplyLivenessToken(request *ApplyLivenessTokenRequest) (response *ApplyLivenessTokenResponse, err error) {
    return c.ApplyLivenessTokenWithContext(context.Background(), request)
}

// ApplyLivenessToken
// This API is used to apply for a token before calling the liveness detection service each time. This token is required for initiating the verification process and getting the result after the verification is completed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ApplyLivenessTokenWithContext(ctx context.Context, request *ApplyLivenessTokenRequest) (response *ApplyLivenessTokenResponse, err error) {
    if request == nil {
        request = NewApplyLivenessTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyLivenessToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyLivenessTokenResponse()
    err = c.Send(request, response)
    return
}

func NewApplySdkVerificationTokenRequest() (request *ApplySdkVerificationTokenRequest) {
    request = &ApplySdkVerificationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "ApplySdkVerificationToken")
    
    
    return
}

func NewApplySdkVerificationTokenResponse() (response *ApplySdkVerificationTokenResponse) {
    response = &ApplySdkVerificationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplySdkVerificationToken
// This API is used to apply for a token before calling the Identity Verification SDK service each time. This token is required for initiating the verification process and getting the result after the verification is completed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ApplySdkVerificationToken(request *ApplySdkVerificationTokenRequest) (response *ApplySdkVerificationTokenResponse, err error) {
    return c.ApplySdkVerificationTokenWithContext(context.Background(), request)
}

// ApplySdkVerificationToken
// This API is used to apply for a token before calling the Identity Verification SDK service each time. This token is required for initiating the verification process and getting the result after the verification is completed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ApplySdkVerificationTokenWithContext(ctx context.Context, request *ApplySdkVerificationTokenRequest) (response *ApplySdkVerificationTokenResponse, err error) {
    if request == nil {
        request = NewApplySdkVerificationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplySdkVerificationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplySdkVerificationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewApplyWebVerificationBizTokenIntlRequest() (request *ApplyWebVerificationBizTokenIntlRequest) {
    request = &ApplyWebVerificationBizTokenIntlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "ApplyWebVerificationBizTokenIntl")
    
    
    return
}

func NewApplyWebVerificationBizTokenIntlResponse() (response *ApplyWebVerificationBizTokenIntlResponse) {
    response = &ApplyWebVerificationBizTokenIntlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyWebVerificationBizTokenIntl
// This API is used to apply for a BizToken before calling the web-based verification service each time. This token is required for initiating a verification process and getting the result after the verification is completed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) ApplyWebVerificationBizTokenIntl(request *ApplyWebVerificationBizTokenIntlRequest) (response *ApplyWebVerificationBizTokenIntlResponse, err error) {
    return c.ApplyWebVerificationBizTokenIntlWithContext(context.Background(), request)
}

// ApplyWebVerificationBizTokenIntl
// This API is used to apply for a BizToken before calling the web-based verification service each time. This token is required for initiating a verification process and getting the result after the verification is completed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) ApplyWebVerificationBizTokenIntlWithContext(ctx context.Context, request *ApplyWebVerificationBizTokenIntlRequest) (response *ApplyWebVerificationBizTokenIntlResponse, err error) {
    if request == nil {
        request = NewApplyWebVerificationBizTokenIntlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyWebVerificationBizTokenIntl require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyWebVerificationBizTokenIntlResponse()
    err = c.Send(request, response)
    return
}

func NewApplyWebVerificationTokenRequest() (request *ApplyWebVerificationTokenRequest) {
    request = &ApplyWebVerificationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "ApplyWebVerificationToken")
    
    
    return
}

func NewApplyWebVerificationTokenResponse() (response *ApplyWebVerificationTokenResponse) {
    response = &ApplyWebVerificationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyWebVerificationToken
// This API is used to apply for a token before calling the web-based verification service each time. This token is required for initiating the verification process and getting the result after the verification is completed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) ApplyWebVerificationToken(request *ApplyWebVerificationTokenRequest) (response *ApplyWebVerificationTokenResponse, err error) {
    return c.ApplyWebVerificationTokenWithContext(context.Background(), request)
}

// ApplyWebVerificationToken
// This API is used to apply for a token before calling the web-based verification service each time. This token is required for initiating the verification process and getting the result after the verification is completed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) ApplyWebVerificationTokenWithContext(ctx context.Context, request *ApplyWebVerificationTokenRequest) (response *ApplyWebVerificationTokenResponse, err error) {
    if request == nil {
        request = NewApplyWebVerificationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyWebVerificationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyWebVerificationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUploadUrlRequest() (request *CreateUploadUrlRequest) {
    request = &CreateUploadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "CreateUploadUrl")
    
    
    return
}

func NewCreateUploadUrlResponse() (response *CreateUploadUrlResponse) {
    response = &CreateUploadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUploadUrl
// This API is used to generate a temporary `UploadUrl` for uploading resource files (with the `HTTP PUT` method). After resource upload, `ResourceUrl` will be passed to the `TargetAction` API to complete the resource passing (specific fields vary by case). 
//
// The data will be stored in a COS bucket in the region specified by the parameter `Region` for two hours.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateUploadUrl(request *CreateUploadUrlRequest) (response *CreateUploadUrlResponse, err error) {
    return c.CreateUploadUrlWithContext(context.Background(), request)
}

// CreateUploadUrl
// This API is used to generate a temporary `UploadUrl` for uploading resource files (with the `HTTP PUT` method). After resource upload, `ResourceUrl` will be passed to the `TargetAction` API to complete the resource passing (specific fields vary by case). 
//
// The data will be stored in a COS bucket in the region specified by the parameter `Region` for two hours.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateUploadUrlWithContext(ctx context.Context, request *CreateUploadUrlRequest) (response *CreateUploadUrlResponse, err error) {
    if request == nil {
        request = NewCreateUploadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUploadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUploadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDetectReflectLivenessAndCompareRequest() (request *DetectReflectLivenessAndCompareRequest) {
    request = &DetectReflectLivenessAndCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "DetectReflectLivenessAndCompare")
    
    
    return
}

func NewDetectReflectLivenessAndCompareResponse() (response *DetectReflectLivenessAndCompareResponse) {
    response = &DetectReflectLivenessAndCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetectReflectLivenessAndCompare
// This API is used to detect liveness with the package generated by the liveness comparison (reflection-based) SDK, and to compare the person detected with that in the image passed in.
//
// The image and the data generated with the SDK must be stored in COS, and the region of the COS bucket must be same as that of requests made with this API. We recommend that you pass resources with upload link APIs.
//
// error code that may be returned:
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTFACEDETECTFAIL = "FailedOperation.SilentFaceDetectFail"
//  FAILEDOPERATION_SILENTFACEQUALITYFAIL = "FailedOperation.SilentFaceQualityFail"
//  FAILEDOPERATION_SILENTMULTIFACEFAIL = "FailedOperation.SilentMultiFaceFail"
//  FAILEDOPERATION_SILENTPICTURELIVEFAIL = "FailedOperation.SilentPictureLiveFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  INTERNALERROR_ACTIONLIGHTDARK = "InternalError.ActionLightDark"
//  INTERNALERROR_ACTIONLIGHTSTRONG = "InternalError.ActionLightStrong"
//  INTERNALERROR_ACTIONNODETECTFACE = "InternalError.ActionNodetectFace"
//  INTERNALERROR_COMPARELOWSIMILARITY = "InternalError.CompareLowSimilarity"
//  INTERNALERROR_LIFEPHOTOPOORQUALITY = "InternalError.LifePhotoPoorQuality"
//  INTERNALERROR_LIFEPHOTOSIZEERROR = "InternalError.LifePhotoSizeError"
func (c *Client) DetectReflectLivenessAndCompare(request *DetectReflectLivenessAndCompareRequest) (response *DetectReflectLivenessAndCompareResponse, err error) {
    return c.DetectReflectLivenessAndCompareWithContext(context.Background(), request)
}

// DetectReflectLivenessAndCompare
// This API is used to detect liveness with the package generated by the liveness comparison (reflection-based) SDK, and to compare the person detected with that in the image passed in.
//
// The image and the data generated with the SDK must be stored in COS, and the region of the COS bucket must be same as that of requests made with this API. We recommend that you pass resources with upload link APIs.
//
// error code that may be returned:
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTFACEDETECTFAIL = "FailedOperation.SilentFaceDetectFail"
//  FAILEDOPERATION_SILENTFACEQUALITYFAIL = "FailedOperation.SilentFaceQualityFail"
//  FAILEDOPERATION_SILENTMULTIFACEFAIL = "FailedOperation.SilentMultiFaceFail"
//  FAILEDOPERATION_SILENTPICTURELIVEFAIL = "FailedOperation.SilentPictureLiveFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  INTERNALERROR_ACTIONLIGHTDARK = "InternalError.ActionLightDark"
//  INTERNALERROR_ACTIONLIGHTSTRONG = "InternalError.ActionLightStrong"
//  INTERNALERROR_ACTIONNODETECTFACE = "InternalError.ActionNodetectFace"
//  INTERNALERROR_COMPARELOWSIMILARITY = "InternalError.CompareLowSimilarity"
//  INTERNALERROR_LIFEPHOTOPOORQUALITY = "InternalError.LifePhotoPoorQuality"
//  INTERNALERROR_LIFEPHOTOSIZEERROR = "InternalError.LifePhotoSizeError"
func (c *Client) DetectReflectLivenessAndCompareWithContext(ctx context.Context, request *DetectReflectLivenessAndCompareRequest) (response *DetectReflectLivenessAndCompareResponse, err error) {
    if request == nil {
        request = NewDetectReflectLivenessAndCompareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectReflectLivenessAndCompare require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectReflectLivenessAndCompareResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateReflectSequenceRequest() (request *GenerateReflectSequenceRequest) {
    request = &GenerateReflectSequenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GenerateReflectSequence")
    
    
    return
}

func NewGenerateReflectSequenceResponse() (response *GenerateReflectSequenceResponse) {
    response = &GenerateReflectSequenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateReflectSequence
// This API is used to generate an appropriate light sequence based on the information collected by the liveness comparison (reflection-based) SDK and pass the light sequence into the SDK to start the identity verification process.
//
// The data generated with the SDK must be stored in COS, and the region of the COS bucket must be same as that of requests made with this API. We recommend that you pass resources with upload link APIs.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
func (c *Client) GenerateReflectSequence(request *GenerateReflectSequenceRequest) (response *GenerateReflectSequenceResponse, err error) {
    return c.GenerateReflectSequenceWithContext(context.Background(), request)
}

// GenerateReflectSequence
// This API is used to generate an appropriate light sequence based on the information collected by the liveness comparison (reflection-based) SDK and pass the light sequence into the SDK to start the identity verification process.
//
// The data generated with the SDK must be stored in COS, and the region of the COS bucket must be same as that of requests made with this API. We recommend that you pass resources with upload link APIs.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
func (c *Client) GenerateReflectSequenceWithContext(ctx context.Context, request *GenerateReflectSequenceRequest) (response *GenerateReflectSequenceResponse, err error) {
    if request == nil {
        request = NewGenerateReflectSequenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateReflectSequence require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateReflectSequenceResponse()
    err = c.Send(request, response)
    return
}

func NewGetFaceIdResultIntlRequest() (request *GetFaceIdResultIntlRequest) {
    request = &GetFaceIdResultIntlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetFaceIdResultIntl")
    
    
    return
}

func NewGetFaceIdResultIntlResponse() (response *GetFaceIdResultIntlResponse) {
    response = &GetFaceIdResultIntlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFaceIdResultIntl
// This API is used to get the verification result with the corresponding SDK token after the identity verification process is completed. The SDK token is valid for two hours (2*3,600s) after generation and can be called multiple times.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetFaceIdResultIntl(request *GetFaceIdResultIntlRequest) (response *GetFaceIdResultIntlResponse, err error) {
    return c.GetFaceIdResultIntlWithContext(context.Background(), request)
}

// GetFaceIdResultIntl
// This API is used to get the verification result with the corresponding SDK token after the identity verification process is completed. The SDK token is valid for two hours (2*3,600s) after generation and can be called multiple times.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetFaceIdResultIntlWithContext(ctx context.Context, request *GetFaceIdResultIntlRequest) (response *GetFaceIdResultIntlResponse, err error) {
    if request == nil {
        request = NewGetFaceIdResultIntlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFaceIdResultIntl require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFaceIdResultIntlResponse()
    err = c.Send(request, response)
    return
}

func NewGetFaceIdTokenIntlRequest() (request *GetFaceIdTokenIntlRequest) {
    request = &GetFaceIdTokenIntlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetFaceIdTokenIntl")
    
    
    return
}

func NewGetFaceIdTokenIntlResponse() (response *GetFaceIdTokenIntlResponse) {
    response = &GetFaceIdTokenIntlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFaceIdTokenIntl
// This API is used to apply for an SDK token before calling the liveness detection and face comparison SDK each time. The SDK token is used throughout the identity verification process and to get the verification result after the verification is completed. A token is valid for one identity verification process only.
//
// error code that may be returned:
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetFaceIdTokenIntl(request *GetFaceIdTokenIntlRequest) (response *GetFaceIdTokenIntlResponse, err error) {
    return c.GetFaceIdTokenIntlWithContext(context.Background(), request)
}

// GetFaceIdTokenIntl
// This API is used to apply for an SDK token before calling the liveness detection and face comparison SDK each time. The SDK token is used throughout the identity verification process and to get the verification result after the verification is completed. A token is valid for one identity verification process only.
//
// error code that may be returned:
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetFaceIdTokenIntlWithContext(ctx context.Context, request *GetFaceIdTokenIntlRequest) (response *GetFaceIdTokenIntlResponse, err error) {
    if request == nil {
        request = NewGetFaceIdTokenIntlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFaceIdTokenIntl require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFaceIdTokenIntlResponse()
    err = c.Send(request, response)
    return
}

func NewGetLivenessResultRequest() (request *GetLivenessResultRequest) {
    request = &GetLivenessResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetLivenessResult")
    
    
    return
}

func NewGetLivenessResultResponse() (response *GetLivenessResultResponse) {
    response = &GetLivenessResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetLivenessResult
// This API is used to get the verification result with the corresponding token (SdkToken) after the liveness detection is completed. The token is valid for two hours after issuance and can be called multiple times.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetLivenessResult(request *GetLivenessResultRequest) (response *GetLivenessResultResponse, err error) {
    return c.GetLivenessResultWithContext(context.Background(), request)
}

// GetLivenessResult
// This API is used to get the verification result with the corresponding token (SdkToken) after the liveness detection is completed. The token is valid for two hours after issuance and can be called multiple times.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetLivenessResultWithContext(ctx context.Context, request *GetLivenessResultRequest) (response *GetLivenessResultResponse, err error) {
    if request == nil {
        request = NewGetLivenessResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLivenessResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLivenessResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetSdkVerificationResultRequest() (request *GetSdkVerificationResultRequest) {
    request = &GetSdkVerificationResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetSdkVerificationResult")
    
    
    return
}

func NewGetSdkVerificationResultResponse() (response *GetSdkVerificationResultResponse) {
    response = &GetSdkVerificationResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSdkVerificationResult
// This API is used to get the verification result with the corresponding token after the SDK-based verification is completed. The token is valid for three days after issuance and can be called multiple times.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetSdkVerificationResult(request *GetSdkVerificationResultRequest) (response *GetSdkVerificationResultResponse, err error) {
    return c.GetSdkVerificationResultWithContext(context.Background(), request)
}

// GetSdkVerificationResult
// This API is used to get the verification result with the corresponding token after the SDK-based verification is completed. The token is valid for three days after issuance and can be called multiple times.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetSdkVerificationResultWithContext(ctx context.Context, request *GetSdkVerificationResultRequest) (response *GetSdkVerificationResultResponse, err error) {
    if request == nil {
        request = NewGetSdkVerificationResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSdkVerificationResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSdkVerificationResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetWebVerificationResultRequest() (request *GetWebVerificationResultRequest) {
    request = &GetWebVerificationResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetWebVerificationResult")
    
    
    return
}

func NewGetWebVerificationResultResponse() (response *GetWebVerificationResultResponse) {
    response = &GetWebVerificationResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWebVerificationResult
// This API is used to get the verification result with the corresponding token (BizToken) after the web-based verification is completed. The BizToken is valid for three days (3*24*3,600s) after issuance and can be called multiple times.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) GetWebVerificationResult(request *GetWebVerificationResultRequest) (response *GetWebVerificationResultResponse, err error) {
    return c.GetWebVerificationResultWithContext(context.Background(), request)
}

// GetWebVerificationResult
// This API is used to get the verification result with the corresponding token (BizToken) after the web-based verification is completed. The BizToken is valid for three days (3*24*3,600s) after issuance and can be called multiple times.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) GetWebVerificationResultWithContext(ctx context.Context, request *GetWebVerificationResultRequest) (response *GetWebVerificationResultResponse, err error) {
    if request == nil {
        request = NewGetWebVerificationResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWebVerificationResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWebVerificationResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetWebVerificationResultIntlRequest() (request *GetWebVerificationResultIntlRequest) {
    request = &GetWebVerificationResultIntlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetWebVerificationResultIntl")
    
    
    return
}

func NewGetWebVerificationResultIntlResponse() (response *GetWebVerificationResultIntlResponse) {
    response = &GetWebVerificationResultIntlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWebVerificationResultIntl
// This API is used to get the verification result with the corresponding BizToken after the u200dweb-based verification is completed. The token is valid for three days (259,200s) after issuance and can be called multiple times.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetWebVerificationResultIntl(request *GetWebVerificationResultIntlRequest) (response *GetWebVerificationResultIntlResponse, err error) {
    return c.GetWebVerificationResultIntlWithContext(context.Background(), request)
}

// GetWebVerificationResultIntl
// This API is used to get the verification result with the corresponding BizToken after the u200dweb-based verification is completed. The token is valid for three days (259,200s) after issuance and can be called multiple times.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetWebVerificationResultIntlWithContext(ctx context.Context, request *GetWebVerificationResultIntlRequest) (response *GetWebVerificationResultIntlResponse, err error) {
    if request == nil {
        request = NewGetWebVerificationResultIntlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWebVerificationResultIntl require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWebVerificationResultIntlResponse()
    err = c.Send(request, response)
    return
}

func NewLivenessCompareRequest() (request *LivenessCompareRequest) {
    request = &LivenessCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "LivenessCompare")
    
    
    return
}

func NewLivenessCompareResponse() (response *LivenessCompareResponse) {
    response = &LivenessCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LivenessCompare
// This API is used to pass in a video and a photo, determine whether the person in the video is real, and if yes, then determine whether the person in the video is the same as that in the photo.
//
// This API on the legacy version will continue to serve existing users but will be unavailable to new users. We recommend you use `VideoLivenessCompare` for better service quality.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTEYELIVEFAIL = "FailedOperation.SilentEyeLiveFail"
//  FAILEDOPERATION_SILENTFACEDETECTFAIL = "FailedOperation.SilentFaceDetectFail"
//  FAILEDOPERATION_SILENTFACEQUALITYFAIL = "FailedOperation.SilentFaceQualityFail"
//  FAILEDOPERATION_SILENTFACEWITHMASKFAIL = "FailedOperation.SilentFaceWithMaskFail"
//  FAILEDOPERATION_SILENTMOUTHLIVEFAIL = "FailedOperation.SilentMouthLiveFail"
//  FAILEDOPERATION_SILENTMULTIFACEFAIL = "FailedOperation.SilentMultiFaceFail"
//  FAILEDOPERATION_SILENTPICTURELIVEFAIL = "FailedOperation.SilentPictureLiveFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LivenessCompare(request *LivenessCompareRequest) (response *LivenessCompareResponse, err error) {
    return c.LivenessCompareWithContext(context.Background(), request)
}

// LivenessCompare
// This API is used to pass in a video and a photo, determine whether the person in the video is real, and if yes, then determine whether the person in the video is the same as that in the photo.
//
// This API on the legacy version will continue to serve existing users but will be unavailable to new users. We recommend you use `VideoLivenessCompare` for better service quality.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTEYELIVEFAIL = "FailedOperation.SilentEyeLiveFail"
//  FAILEDOPERATION_SILENTFACEDETECTFAIL = "FailedOperation.SilentFaceDetectFail"
//  FAILEDOPERATION_SILENTFACEQUALITYFAIL = "FailedOperation.SilentFaceQualityFail"
//  FAILEDOPERATION_SILENTFACEWITHMASKFAIL = "FailedOperation.SilentFaceWithMaskFail"
//  FAILEDOPERATION_SILENTMOUTHLIVEFAIL = "FailedOperation.SilentMouthLiveFail"
//  FAILEDOPERATION_SILENTMULTIFACEFAIL = "FailedOperation.SilentMultiFaceFail"
//  FAILEDOPERATION_SILENTPICTURELIVEFAIL = "FailedOperation.SilentPictureLiveFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LivenessCompareWithContext(ctx context.Context, request *LivenessCompareRequest) (response *LivenessCompareResponse, err error) {
    if request == nil {
        request = NewLivenessCompareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LivenessCompare require credential")
    }

    request.SetContext(ctx)
    
    response = NewLivenessCompareResponse()
    err = c.Send(request, response)
    return
}

func NewVideoLivenessCompareRequest() (request *VideoLivenessCompareRequest) {
    request = &VideoLivenessCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "VideoLivenessCompare")
    
    
    return
}

func NewVideoLivenessCompareResponse() (response *VideoLivenessCompareResponse) {
    response = &VideoLivenessCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VideoLivenessCompare
// This API is used to pass in URLs of a video and a photo, determine whether the person in the video is real, and if yes, then determine whether the person in the video is the same as that in the photo.
//
// error code that may be returned:
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTEYELIVEFAIL = "FailedOperation.SilentEyeLiveFail"
//  FAILEDOPERATION_SILENTFACEDETECTFAIL = "FailedOperation.SilentFaceDetectFail"
//  FAILEDOPERATION_SILENTFACEQUALITYFAIL = "FailedOperation.SilentFaceQualityFail"
//  FAILEDOPERATION_SILENTFACEWITHMASKFAIL = "FailedOperation.SilentFaceWithMaskFail"
//  FAILEDOPERATION_SILENTMOUTHLIVEFAIL = "FailedOperation.SilentMouthLiveFail"
//  FAILEDOPERATION_SILENTMULTIFACEFAIL = "FailedOperation.SilentMultiFaceFail"
//  FAILEDOPERATION_SILENTPICTURELIVEFAIL = "FailedOperation.SilentPictureLiveFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VideoLivenessCompare(request *VideoLivenessCompareRequest) (response *VideoLivenessCompareResponse, err error) {
    return c.VideoLivenessCompareWithContext(context.Background(), request)
}

// VideoLivenessCompare
// This API is used to pass in URLs of a video and a photo, determine whether the person in the video is real, and if yes, then determine whether the person in the video is the same as that in the photo.
//
// error code that may be returned:
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTEYELIVEFAIL = "FailedOperation.SilentEyeLiveFail"
//  FAILEDOPERATION_SILENTFACEDETECTFAIL = "FailedOperation.SilentFaceDetectFail"
//  FAILEDOPERATION_SILENTFACEQUALITYFAIL = "FailedOperation.SilentFaceQualityFail"
//  FAILEDOPERATION_SILENTFACEWITHMASKFAIL = "FailedOperation.SilentFaceWithMaskFail"
//  FAILEDOPERATION_SILENTMOUTHLIVEFAIL = "FailedOperation.SilentMouthLiveFail"
//  FAILEDOPERATION_SILENTMULTIFACEFAIL = "FailedOperation.SilentMultiFaceFail"
//  FAILEDOPERATION_SILENTPICTURELIVEFAIL = "FailedOperation.SilentPictureLiveFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VideoLivenessCompareWithContext(ctx context.Context, request *VideoLivenessCompareRequest) (response *VideoLivenessCompareResponse, err error) {
    if request == nil {
        request = NewVideoLivenessCompareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VideoLivenessCompare require credential")
    }

    request.SetContext(ctx)
    
    response = NewVideoLivenessCompareResponse()
    err = c.Send(request, response)
    return
}
