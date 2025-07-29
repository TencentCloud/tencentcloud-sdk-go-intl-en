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

package v20220927

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-09-27"

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


func NewFuseFaceRequest() (request *FuseFaceRequest) {
    request = &FuseFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("facefusion", APIVersion, "FuseFace")
    
    
    return
}

func NewFuseFaceResponse() (response *FuseFaceResponse) {
    response = &FuseFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FuseFace
// This API is used to perform the fusion of a single face, multiple faces, and specified faces with the material template by uploading face images. Users can add logos to generated images. See <a href="https://intl.cloud.tencent.com/document/product/670/38247?from_cn_redirect=1" target="_blank">Fusion Access Guide</a>.
//
// 
//
// 
//
// 
//
// - The signature method in the public parameters must be specified as the V3 version. In other words, set the SignatureMethod parameter to TC3-HMAC-SHA256.
//
// error code that may be returned:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) FuseFace(request *FuseFaceRequest) (response *FuseFaceResponse, err error) {
    return c.FuseFaceWithContext(context.Background(), request)
}

// FuseFace
// This API is used to perform the fusion of a single face, multiple faces, and specified faces with the material template by uploading face images. Users can add logos to generated images. See <a href="https://intl.cloud.tencent.com/document/product/670/38247?from_cn_redirect=1" target="_blank">Fusion Access Guide</a>.
//
// 
//
// 
//
// 
//
// - The signature method in the public parameters must be specified as the V3 version. In other words, set the SignatureMethod parameter to TC3-HMAC-SHA256.
//
// error code that may be returned:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) FuseFaceWithContext(ctx context.Context, request *FuseFaceRequest) (response *FuseFaceResponse, err error) {
    if request == nil {
        request = NewFuseFaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "facefusion", APIVersion, "FuseFace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FuseFace require credential")
    }

    request.SetContext(ctx)
    
    response = NewFuseFaceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVideoFaceFusionJobRequest() (request *QueryVideoFaceFusionJobRequest) {
    request = &QueryVideoFaceFusionJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("facefusion", APIVersion, "QueryVideoFaceFusionJob")
    
    
    return
}

func NewQueryVideoFaceFusionJobResponse() (response *QueryVideoFaceFusionJobResponse) {
    response = &QueryVideoFaceFusionJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryVideoFaceFusionJob
// This API is used to query the progress and status of video face fusion tasks by Job ID.
//
// error code that may be returned:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBHASBEENCANCELED = "FailedOperation.JobHasBeenCanceled"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
func (c *Client) QueryVideoFaceFusionJob(request *QueryVideoFaceFusionJobRequest) (response *QueryVideoFaceFusionJobResponse, err error) {
    return c.QueryVideoFaceFusionJobWithContext(context.Background(), request)
}

// QueryVideoFaceFusionJob
// This API is used to query the progress and status of video face fusion tasks by Job ID.
//
// error code that may be returned:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBHASBEENCANCELED = "FailedOperation.JobHasBeenCanceled"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
func (c *Client) QueryVideoFaceFusionJobWithContext(ctx context.Context, request *QueryVideoFaceFusionJobRequest) (response *QueryVideoFaceFusionJobResponse, err error) {
    if request == nil {
        request = NewQueryVideoFaceFusionJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "facefusion", APIVersion, "QueryVideoFaceFusionJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVideoFaceFusionJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVideoFaceFusionJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitVideoFaceFusionJobRequest() (request *SubmitVideoFaceFusionJobRequest) {
    request = &SubmitVideoFaceFusionJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("facefusion", APIVersion, "SubmitVideoFaceFusionJob")
    
    
    return
}

func NewSubmitVideoFaceFusionJobResponse() (response *SubmitVideoFaceFusionJobResponse) {
    response = &SubmitVideoFaceFusionJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoFaceFusionJob
// This API is used to submit asynchronous processing tasks of video face fusion. After a task is submitted, the Job ID, estimated completion time, and current queue length will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_FACEIDNOTINVIDEO = "FailedOperation.FaceIdNotInVideo"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//  FAILEDOPERATION_FUSEMATERIALNOTAVAILABLE = "FailedOperation.FuseMaterialNotAvailable"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) SubmitVideoFaceFusionJob(request *SubmitVideoFaceFusionJobRequest) (response *SubmitVideoFaceFusionJobResponse, err error) {
    return c.SubmitVideoFaceFusionJobWithContext(context.Background(), request)
}

// SubmitVideoFaceFusionJob
// This API is used to submit asynchronous processing tasks of video face fusion. After a task is submitted, the Job ID, estimated completion time, and current queue length will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_FACEIDNOTINVIDEO = "FailedOperation.FaceIdNotInVideo"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//  FAILEDOPERATION_FUSEMATERIALNOTAVAILABLE = "FailedOperation.FuseMaterialNotAvailable"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) SubmitVideoFaceFusionJobWithContext(ctx context.Context, request *SubmitVideoFaceFusionJobRequest) (response *SubmitVideoFaceFusionJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoFaceFusionJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "facefusion", APIVersion, "SubmitVideoFaceFusionJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoFaceFusionJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoFaceFusionJobResponse()
    err = c.Send(request, response)
    return
}
