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

package v20240523

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2024-05-23"

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


func NewDescribeImageAnimateJobRequest() (request *DescribeImageAnimateJobRequest) {
    request = &DescribeImageAnimateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeImageAnimateJob")
    
    
    return
}

func NewDescribeImageAnimateJobResponse() (response *DescribeImageAnimateJobResponse) {
    response = &DescribeImageAnimateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageAnimateJob
// This API is used to query image animation tasks. The image animation feature supports generating videos based on dance movements and images to meet the needs of scenarios such as social entertainment and interactive marketing.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageAnimateJob(request *DescribeImageAnimateJobRequest) (response *DescribeImageAnimateJobResponse, err error) {
    return c.DescribeImageAnimateJobWithContext(context.Background(), request)
}

// DescribeImageAnimateJob
// This API is used to query image animation tasks. The image animation feature supports generating videos based on dance movements and images to meet the needs of scenarios such as social entertainment and interactive marketing.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageAnimateJobWithContext(ctx context.Context, request *DescribeImageAnimateJobRequest) (response *DescribeImageAnimateJobResponse, err error) {
    if request == nil {
        request = NewDescribeImageAnimateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAnimateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAnimateJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitImageAnimateJobRequest() (request *SubmitImageAnimateJobRequest) {
    request = &SubmitImageAnimateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitImageAnimateJob")
    
    
    return
}

func NewSubmitImageAnimateJobResponse() (response *SubmitImageAnimateJobResponse) {
    response = &SubmitImageAnimateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitImageAnimateJob
// This API is used to submit image animation tasks. The image animation feature supports generating videos based on dance movements and images to meet the needs of scenarios such as social entertainment and interactive marketing.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageAnimateJob(request *SubmitImageAnimateJobRequest) (response *SubmitImageAnimateJobResponse, err error) {
    return c.SubmitImageAnimateJobWithContext(context.Background(), request)
}

// SubmitImageAnimateJob
// This API is used to submit image animation tasks. The image animation feature supports generating videos based on dance movements and images to meet the needs of scenarios such as social entertainment and interactive marketing.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageAnimateJobWithContext(ctx context.Context, request *SubmitImageAnimateJobRequest) (response *SubmitImageAnimateJobResponse, err error) {
    if request == nil {
        request = NewSubmitImageAnimateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitImageAnimateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitImageAnimateJobResponse()
    err = c.Send(request, response)
    return
}
