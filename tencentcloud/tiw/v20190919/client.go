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

package v20190919

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-09-19"

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


func NewApplyTiwTrialRequest() (request *ApplyTiwTrialRequest) {
    request = &ApplyTiwTrialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ApplyTiwTrial")
    
    
    return
}

func NewApplyTiwTrialResponse() (response *ApplyTiwTrialResponse) {
    response = &ApplyTiwTrialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyTiwTrial
// This API is used to apply for a Tencent Interactive Whiteboard trial (15-day by default).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ALREADYENABLED = "FailedOperation.AlreadyEnabled"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ApplyTiwTrial(request *ApplyTiwTrialRequest) (response *ApplyTiwTrialResponse, err error) {
    return c.ApplyTiwTrialWithContext(context.Background(), request)
}

// ApplyTiwTrial
// This API is used to apply for a Tencent Interactive Whiteboard trial (15-day by default).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ALREADYENABLED = "FailedOperation.AlreadyEnabled"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ApplyTiwTrialWithContext(ctx context.Context, request *ApplyTiwTrialRequest) (response *ApplyTiwTrialResponse, err error) {
    if request == nil {
        request = NewApplyTiwTrialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyTiwTrial require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyTiwTrialResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateApplication")
    
    
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplication
// This API is used to create a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATIONALREADYEXISTS = "InvalidParameter.ApplicationAlreadyExists"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    return c.CreateApplicationWithContext(context.Background(), request)
}

// CreateApplication
// This API is used to create a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATIONALREADYEXISTS = "InvalidParameter.ApplicationAlreadyExists"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOfflineRecordRequest() (request *CreateOfflineRecordRequest) {
    request = &CreateOfflineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateOfflineRecord")
    
    
    return
}

func NewCreateOfflineRecordResponse() (response *CreateOfflineRecordResponse) {
    response = &CreateOfflineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOfflineRecord
// 课后录制服务已下线
//
// 
//
// This API is used to create an offline recording task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateOfflineRecord(request *CreateOfflineRecordRequest) (response *CreateOfflineRecordResponse, err error) {
    return c.CreateOfflineRecordWithContext(context.Background(), request)
}

// CreateOfflineRecord
// 课后录制服务已下线
//
// 
//
// This API is used to create an offline recording task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateOfflineRecordWithContext(ctx context.Context, request *CreateOfflineRecordRequest) (response *CreateOfflineRecordResponse, err error) {
    if request == nil {
        request = NewCreateOfflineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOfflineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOfflineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotTaskRequest() (request *CreateSnapshotTaskRequest) {
    request = &CreateSnapshotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateSnapshotTask")
    
    
    return
}

func NewCreateSnapshotTaskResponse() (response *CreateSnapshotTaskResponse) {
    response = &CreateSnapshotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSnapshotTask
// This API is used to create a whiteboard snapshot task. If a callback URL is provided, the whiteboard snapshot result is sent to the callback URL after the task is complete.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateSnapshotTask(request *CreateSnapshotTaskRequest) (response *CreateSnapshotTaskResponse, err error) {
    return c.CreateSnapshotTaskWithContext(context.Background(), request)
}

// CreateSnapshotTask
// This API is used to create a whiteboard snapshot task. If a callback URL is provided, the whiteboard snapshot result is sent to the callback URL after the task is complete.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateSnapshotTaskWithContext(ctx context.Context, request *CreateSnapshotTaskRequest) (response *CreateSnapshotTaskResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTranscodeRequest() (request *CreateTranscodeRequest) {
    request = &CreateTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateTranscode")
    
    
    return
}

func NewCreateTranscodeResponse() (response *CreateTranscodeResponse) {
    response = &CreateTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTranscode
// This API is used to create a document transcoding task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateTranscode(request *CreateTranscodeRequest) (response *CreateTranscodeResponse, err error) {
    return c.CreateTranscodeWithContext(context.Background(), request)
}

// CreateTranscode
// This API is used to create a document transcoding task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateTranscodeWithContext(ctx context.Context, request *CreateTranscodeRequest) (response *CreateTranscodeResponse, err error) {
    if request == nil {
        request = NewCreateTranscodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoGenerationTaskRequest() (request *CreateVideoGenerationTaskRequest) {
    request = &CreateVideoGenerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateVideoGenerationTask")
    
    
    return
}

func NewCreateVideoGenerationTaskResponse() (response *CreateVideoGenerationTaskResponse) {
    response = &CreateVideoGenerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVideoGenerationTask
// This API is used to create a recording video generation task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateVideoGenerationTask(request *CreateVideoGenerationTaskRequest) (response *CreateVideoGenerationTaskResponse, err error) {
    return c.CreateVideoGenerationTaskWithContext(context.Background(), request)
}

// CreateVideoGenerationTask
// This API is used to create a recording video generation task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateVideoGenerationTaskWithContext(ctx context.Context, request *CreateVideoGenerationTaskRequest) (response *CreateVideoGenerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateVideoGenerationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoGenerationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoGenerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAPIServiceRequest() (request *DescribeAPIServiceRequest) {
    request = &DescribeAPIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeAPIService")
    
    
    return
}

func NewDescribeAPIServiceResponse() (response *DescribeAPIServiceResponse) {
    response = &DescribeAPIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAPIService
// This API is used to query the information about other cloud products by using the service role.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAPIService(request *DescribeAPIServiceRequest) (response *DescribeAPIServiceResponse, err error) {
    return c.DescribeAPIServiceWithContext(context.Background(), request)
}

// DescribeAPIService
// This API is used to query the information about other cloud products by using the service role.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAPIServiceWithContext(ctx context.Context, request *DescribeAPIServiceRequest) (response *DescribeAPIServiceResponse, err error) {
    if request == nil {
        request = NewDescribeAPIServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAPIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAPIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationInfosRequest() (request *DescribeApplicationInfosRequest) {
    request = &DescribeApplicationInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeApplicationInfos")
    
    
    return
}

func NewDescribeApplicationInfosResponse() (response *DescribeApplicationInfosResponse) {
    response = &DescribeApplicationInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationInfos
// This API is used to query the details of a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeApplicationInfos(request *DescribeApplicationInfosRequest) (response *DescribeApplicationInfosResponse, err error) {
    return c.DescribeApplicationInfosWithContext(context.Background(), request)
}

// DescribeApplicationInfos
// This API is used to query the details of a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeApplicationInfosWithContext(ctx context.Context, request *DescribeApplicationInfosRequest) (response *DescribeApplicationInfosResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationUsageRequest() (request *DescribeApplicationUsageRequest) {
    request = &DescribeApplicationUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeApplicationUsage")
    
    
    return
}

func NewDescribeApplicationUsageResponse() (response *DescribeApplicationUsageResponse) {
    response = &DescribeApplicationUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationUsage
// This API is used to query the subproduct usage of Tencent Interactive Whiteboard.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeApplicationUsage(request *DescribeApplicationUsageRequest) (response *DescribeApplicationUsageResponse, err error) {
    return c.DescribeApplicationUsageWithContext(context.Background(), request)
}

// DescribeApplicationUsage
// This API is used to query the subproduct usage of Tencent Interactive Whiteboard.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeApplicationUsageWithContext(ctx context.Context, request *DescribeApplicationUsageRequest) (response *DescribeApplicationUsageResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBoardSDKLogRequest() (request *DescribeBoardSDKLogRequest) {
    request = &DescribeBoardSDKLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeBoardSDKLog")
    
    
    return
}

func NewDescribeBoardSDKLogResponse() (response *DescribeBoardSDKLogResponse) {
    response = &DescribeBoardSDKLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBoardSDKLog
// This API is used to query the logs of a whiteboard application on a client.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeBoardSDKLog(request *DescribeBoardSDKLogRequest) (response *DescribeBoardSDKLogResponse, err error) {
    return c.DescribeBoardSDKLogWithContext(context.Background(), request)
}

// DescribeBoardSDKLog
// This API is used to query the logs of a whiteboard application on a client.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeBoardSDKLogWithContext(ctx context.Context, request *DescribeBoardSDKLogRequest) (response *DescribeBoardSDKLogResponse, err error) {
    if request == nil {
        request = NewDescribeBoardSDKLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBoardSDKLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBoardSDKLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIMApplicationsRequest() (request *DescribeIMApplicationsRequest) {
    request = &DescribeIMApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeIMApplications")
    
    
    return
}

func NewDescribeIMApplicationsResponse() (response *DescribeIMApplicationsResponse) {
    response = &DescribeIMApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIMApplications
// This API is used to query the instant messaging (IM) applications that are available for creating a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeIMApplications(request *DescribeIMApplicationsRequest) (response *DescribeIMApplicationsResponse, err error) {
    return c.DescribeIMApplicationsWithContext(context.Background(), request)
}

// DescribeIMApplications
// This API is used to query the instant messaging (IM) applications that are available for creating a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeIMApplicationsWithContext(ctx context.Context, request *DescribeIMApplicationsRequest) (response *DescribeIMApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeIMApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIMApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIMApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineRecordRequest() (request *DescribeOfflineRecordRequest) {
    request = &DescribeOfflineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOfflineRecord")
    
    
    return
}

func NewDescribeOfflineRecordResponse() (response *DescribeOfflineRecordResponse) {
    response = &DescribeOfflineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOfflineRecord
// 课后录制服务已下线
//
// 
//
// This API is used to query the information about an offline recording task, including the recording progress and recording result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOfflineRecord(request *DescribeOfflineRecordRequest) (response *DescribeOfflineRecordResponse, err error) {
    return c.DescribeOfflineRecordWithContext(context.Background(), request)
}

// DescribeOfflineRecord
// 课后录制服务已下线
//
// 
//
// This API is used to query the information about an offline recording task, including the recording progress and recording result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOfflineRecordWithContext(ctx context.Context, request *DescribeOfflineRecordRequest) (response *DescribeOfflineRecordResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfflineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfflineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineRecordCallbackRequest() (request *DescribeOfflineRecordCallbackRequest) {
    request = &DescribeOfflineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOfflineRecordCallback")
    
    
    return
}

func NewDescribeOfflineRecordCallbackResponse() (response *DescribeOfflineRecordCallbackResponse) {
    response = &DescribeOfflineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOfflineRecordCallback
// 课后录制服务已下线
//
// 
//
// This API is used to query the offline recording callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOfflineRecordCallback(request *DescribeOfflineRecordCallbackRequest) (response *DescribeOfflineRecordCallbackResponse, err error) {
    return c.DescribeOfflineRecordCallbackWithContext(context.Background(), request)
}

// DescribeOfflineRecordCallback
// 课后录制服务已下线
//
// 
//
// This API is used to query the offline recording callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOfflineRecordCallbackWithContext(ctx context.Context, request *DescribeOfflineRecordCallbackRequest) (response *DescribeOfflineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfflineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfflineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOnlineRecordRequest() (request *DescribeOnlineRecordRequest) {
    request = &DescribeOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOnlineRecord")
    
    
    return
}

func NewDescribeOnlineRecordResponse() (response *DescribeOnlineRecordResponse) {
    response = &DescribeOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOnlineRecord
// This API is used to query the status and result of a real-time recording task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecord(request *DescribeOnlineRecordRequest) (response *DescribeOnlineRecordResponse, err error) {
    return c.DescribeOnlineRecordWithContext(context.Background(), request)
}

// DescribeOnlineRecord
// This API is used to query the status and result of a real-time recording task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecordWithContext(ctx context.Context, request *DescribeOnlineRecordRequest) (response *DescribeOnlineRecordResponse, err error) {
    if request == nil {
        request = NewDescribeOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOnlineRecordCallbackRequest() (request *DescribeOnlineRecordCallbackRequest) {
    request = &DescribeOnlineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOnlineRecordCallback")
    
    
    return
}

func NewDescribeOnlineRecordCallbackResponse() (response *DescribeOnlineRecordCallbackResponse) {
    response = &DescribeOnlineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOnlineRecordCallback
// This API is used to query the real-time recording callback address.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecordCallback(request *DescribeOnlineRecordCallbackRequest) (response *DescribeOnlineRecordCallbackResponse, err error) {
    return c.DescribeOnlineRecordCallbackWithContext(context.Background(), request)
}

// DescribeOnlineRecordCallback
// This API is used to query the real-time recording callback address.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecordCallbackWithContext(ctx context.Context, request *DescribeOnlineRecordCallbackRequest) (response *DescribeOnlineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeOnlineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOnlineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOnlineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePostpaidUsageRequest() (request *DescribePostpaidUsageRequest) {
    request = &DescribePostpaidUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribePostpaidUsage")
    
    
    return
}

func NewDescribePostpaidUsageResponse() (response *DescribePostpaidUsageResponse) {
    response = &DescribePostpaidUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePostpaidUsage
// This API is used to query the pay-as-you-go usage of a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribePostpaidUsage(request *DescribePostpaidUsageRequest) (response *DescribePostpaidUsageResponse, err error) {
    return c.DescribePostpaidUsageWithContext(context.Background(), request)
}

// DescribePostpaidUsage
// This API is used to query the pay-as-you-go usage of a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribePostpaidUsageWithContext(ctx context.Context, request *DescribePostpaidUsageRequest) (response *DescribePostpaidUsageResponse, err error) {
    if request == nil {
        request = NewDescribePostpaidUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePostpaidUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePostpaidUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityMetricsRequest() (request *DescribeQualityMetricsRequest) {
    request = &DescribeQualityMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeQualityMetrics")
    
    
    return
}

func NewDescribeQualityMetricsResponse() (response *DescribeQualityMetricsResponse) {
    response = &DescribeQualityMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQualityMetrics
// This API is used to query the quality data of a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeQualityMetrics(request *DescribeQualityMetricsRequest) (response *DescribeQualityMetricsResponse, err error) {
    return c.DescribeQualityMetricsWithContext(context.Background(), request)
}

// DescribeQualityMetrics
// This API is used to query the quality data of a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeQualityMetricsWithContext(ctx context.Context, request *DescribeQualityMetricsRequest) (response *DescribeQualityMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeQualityMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQualityMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQualityMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordSearchRequest() (request *DescribeRecordSearchRequest) {
    request = &DescribeRecordSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeRecordSearch")
    
    
    return
}

func NewDescribeRecordSearchResponse() (response *DescribeRecordSearchResponse) {
    response = &DescribeRecordSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordSearch
// This API is used to query real-time recording tasks by room ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeRecordSearch(request *DescribeRecordSearchRequest) (response *DescribeRecordSearchResponse, err error) {
    return c.DescribeRecordSearchWithContext(context.Background(), request)
}

// DescribeRecordSearch
// This API is used to query real-time recording tasks by room ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeRecordSearchWithContext(ctx context.Context, request *DescribeRecordSearchRequest) (response *DescribeRecordSearchResponse, err error) {
    if request == nil {
        request = NewDescribeRecordSearchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordSearchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomListRequest() (request *DescribeRoomListRequest) {
    request = &DescribeRoomListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeRoomList")
    
    
    return
}

func NewDescribeRoomListResponse() (response *DescribeRoomListResponse) {
    response = &DescribeRoomListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoomList
// This API is used to query the rooms of a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeRoomList(request *DescribeRoomListRequest) (response *DescribeRoomListResponse, err error) {
    return c.DescribeRoomListWithContext(context.Background(), request)
}

// DescribeRoomList
// This API is used to query the rooms of a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeRoomListWithContext(ctx context.Context, request *DescribeRoomListRequest) (response *DescribeRoomListResponse, err error) {
    if request == nil {
        request = NewDescribeRoomListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotTaskRequest() (request *DescribeSnapshotTaskRequest) {
    request = &DescribeSnapshotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeSnapshotTask")
    
    
    return
}

func NewDescribeSnapshotTaskResponse() (response *DescribeSnapshotTaskResponse) {
    response = &DescribeSnapshotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotTask
// This API is used to query the information about a whiteboard snapshot task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSnapshotTask(request *DescribeSnapshotTaskRequest) (response *DescribeSnapshotTaskResponse, err error) {
    return c.DescribeSnapshotTaskWithContext(context.Background(), request)
}

// DescribeSnapshotTask
// This API is used to query the information about a whiteboard snapshot task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSnapshotTaskWithContext(ctx context.Context, request *DescribeSnapshotTaskRequest) (response *DescribeSnapshotTaskResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTIWDailyUsageRequest() (request *DescribeTIWDailyUsageRequest) {
    request = &DescribeTIWDailyUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTIWDailyUsage")
    
    
    return
}

func NewDescribeTIWDailyUsageResponse() (response *DescribeTIWDailyUsageResponse) {
    response = &DescribeTIWDailyUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTIWDailyUsage
// This API is used to query the daily billable usage of a whiteboard application.
//
// 1. The period queried per request cannot be longer than 31 days.
//
// 2. Due to statistics delays and other reasons, you cannot query the usage data of the current day. You can query today's usage after 7:00 tomorrow.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTIWDailyUsage(request *DescribeTIWDailyUsageRequest) (response *DescribeTIWDailyUsageResponse, err error) {
    return c.DescribeTIWDailyUsageWithContext(context.Background(), request)
}

// DescribeTIWDailyUsage
// This API is used to query the daily billable usage of a whiteboard application.
//
// 1. The period queried per request cannot be longer than 31 days.
//
// 2. Due to statistics delays and other reasons, you cannot query the usage data of the current day. You can query today's usage after 7:00 tomorrow.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTIWDailyUsageWithContext(ctx context.Context, request *DescribeTIWDailyUsageRequest) (response *DescribeTIWDailyUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTIWDailyUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTIWDailyUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTIWDailyUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTIWRoomDailyUsageRequest() (request *DescribeTIWRoomDailyUsageRequest) {
    request = &DescribeTIWRoomDailyUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTIWRoomDailyUsage")
    
    
    return
}

func NewDescribeTIWRoomDailyUsageResponse() (response *DescribeTIWRoomDailyUsageResponse) {
    response = &DescribeTIWRoomDailyUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTIWRoomDailyUsage
// This API is used to query the daily billable usage by each room of a whiteboard application.
//
// 1. The period queried per request cannot be longer than 31 days.
//
// 2. Due to statistics delays and other reasons, you cannot query the usage data of the current day. You can query today's usage after 7:00 tomorrow.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTIWRoomDailyUsage(request *DescribeTIWRoomDailyUsageRequest) (response *DescribeTIWRoomDailyUsageResponse, err error) {
    return c.DescribeTIWRoomDailyUsageWithContext(context.Background(), request)
}

// DescribeTIWRoomDailyUsage
// This API is used to query the daily billable usage by each room of a whiteboard application.
//
// 1. The period queried per request cannot be longer than 31 days.
//
// 2. Due to statistics delays and other reasons, you cannot query the usage data of the current day. You can query today's usage after 7:00 tomorrow.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTIWRoomDailyUsageWithContext(ctx context.Context, request *DescribeTIWRoomDailyUsageRequest) (response *DescribeTIWRoomDailyUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTIWRoomDailyUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTIWRoomDailyUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTIWRoomDailyUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeRequest() (request *DescribeTranscodeRequest) {
    request = &DescribeTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscode")
    
    
    return
}

func NewDescribeTranscodeResponse() (response *DescribeTranscodeResponse) {
    response = &DescribeTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTranscode
// This API is used to query the progress and result of a document transcoding task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEDOWNLOADFAIL = "FailedOperation.FileDownloadFail"
//  FAILEDOPERATION_FILEFORMATERROR = "FailedOperation.FileFormatError"
//  FAILEDOPERATION_FILEOPENFAIL = "FailedOperation.FileOpenFail"
//  FAILEDOPERATION_FILEUPLOADFAIL = "FailedOperation.FileUploadFail"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  FAILEDOPERATION_TRANSCODESERVERERROR = "FailedOperation.TranscodeServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  INVALIDPARAMETER_URLFORMATERROR = "InvalidParameter.UrlFormatError"
//  LIMITEXCEEDED_TRANSCODEPAGESLIMITATION = "LimitExceeded.TranscodePagesLimitation"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscode(request *DescribeTranscodeRequest) (response *DescribeTranscodeResponse, err error) {
    return c.DescribeTranscodeWithContext(context.Background(), request)
}

// DescribeTranscode
// This API is used to query the progress and result of a document transcoding task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEDOWNLOADFAIL = "FailedOperation.FileDownloadFail"
//  FAILEDOPERATION_FILEFORMATERROR = "FailedOperation.FileFormatError"
//  FAILEDOPERATION_FILEOPENFAIL = "FailedOperation.FileOpenFail"
//  FAILEDOPERATION_FILEUPLOADFAIL = "FailedOperation.FileUploadFail"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  FAILEDOPERATION_TRANSCODESERVERERROR = "FailedOperation.TranscodeServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  INVALIDPARAMETER_URLFORMATERROR = "InvalidParameter.UrlFormatError"
//  LIMITEXCEEDED_TRANSCODEPAGESLIMITATION = "LimitExceeded.TranscodePagesLimitation"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeWithContext(ctx context.Context, request *DescribeTranscodeRequest) (response *DescribeTranscodeResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeCallbackRequest() (request *DescribeTranscodeCallbackRequest) {
    request = &DescribeTranscodeCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscodeCallback")
    
    
    return
}

func NewDescribeTranscodeCallbackResponse() (response *DescribeTranscodeCallbackResponse) {
    response = &DescribeTranscodeCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTranscodeCallback
// This API is used to query the document transcoding callback address.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeCallback(request *DescribeTranscodeCallbackRequest) (response *DescribeTranscodeCallbackResponse, err error) {
    return c.DescribeTranscodeCallbackWithContext(context.Background(), request)
}

// DescribeTranscodeCallback
// This API is used to query the document transcoding callback address.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeCallbackWithContext(ctx context.Context, request *DescribeTranscodeCallbackRequest) (response *DescribeTranscodeCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeSearchRequest() (request *DescribeTranscodeSearchRequest) {
    request = &DescribeTranscodeSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscodeSearch")
    
    
    return
}

func NewDescribeTranscodeSearchResponse() (response *DescribeTranscodeSearchResponse) {
    response = &DescribeTranscodeSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTranscodeSearch
// This API is used to query transcoding tasks by document name.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeSearch(request *DescribeTranscodeSearchRequest) (response *DescribeTranscodeSearchResponse, err error) {
    return c.DescribeTranscodeSearchWithContext(context.Background(), request)
}

// DescribeTranscodeSearch
// This API is used to query transcoding tasks by document name.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeSearchWithContext(ctx context.Context, request *DescribeTranscodeSearchRequest) (response *DescribeTranscodeSearchResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeSearchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeSearchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageSummaryRequest() (request *DescribeUsageSummaryRequest) {
    request = &DescribeUsageSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeUsageSummary")
    
    
    return
}

func NewDescribeUsageSummaryResponse() (response *DescribeUsageSummaryResponse) {
    response = &DescribeUsageSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsageSummary
// This API is used to query the summary of subproduct usage within a specified period of time.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUsageSummary(request *DescribeUsageSummaryRequest) (response *DescribeUsageSummaryResponse, err error) {
    return c.DescribeUsageSummaryWithContext(context.Background(), request)
}

// DescribeUsageSummary
// This API is used to query the summary of subproduct usage within a specified period of time.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUsageSummaryWithContext(ctx context.Context, request *DescribeUsageSummaryRequest) (response *DescribeUsageSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeUsageSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsageSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserListRequest() (request *DescribeUserListRequest) {
    request = &DescribeUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeUserList")
    
    
    return
}

func NewDescribeUserListResponse() (response *DescribeUserListResponse) {
    response = &DescribeUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserList
// This API is used to query the users of a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeUserList(request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    return c.DescribeUserListWithContext(context.Background(), request)
}

// DescribeUserList
// This API is used to query the users of a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeUserListWithContext(ctx context.Context, request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    if request == nil {
        request = NewDescribeUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserResourcesRequest() (request *DescribeUserResourcesRequest) {
    request = &DescribeUserResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeUserResources")
    
    
    return
}

func NewDescribeUserResourcesResponse() (response *DescribeUserResourcesResponse) {
    response = &DescribeUserResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserResources
// This API is used to query user resources.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUserResources(request *DescribeUserResourcesRequest) (response *DescribeUserResourcesResponse, err error) {
    return c.DescribeUserResourcesWithContext(context.Background(), request)
}

// DescribeUserResources
// This API is used to query user resources.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUserResourcesWithContext(ctx context.Context, request *DescribeUserResourcesRequest) (response *DescribeUserResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeUserResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserStatusRequest() (request *DescribeUserStatusRequest) {
    request = &DescribeUserStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeUserStatus")
    
    
    return
}

func NewDescribeUserStatusResponse() (response *DescribeUserStatusResponse) {
    response = &DescribeUserStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserStatus
// This API is used to query the Tencent Interactive Whiteboard service status of a user, including the activation status and validity period.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ALREADYENABLED = "FailedOperation.AlreadyEnabled"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUserStatus(request *DescribeUserStatusRequest) (response *DescribeUserStatusResponse, err error) {
    return c.DescribeUserStatusWithContext(context.Background(), request)
}

// DescribeUserStatus
// This API is used to query the Tencent Interactive Whiteboard service status of a user, including the activation status and validity period.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ALREADYENABLED = "FailedOperation.AlreadyEnabled"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUserStatusWithContext(ctx context.Context, request *DescribeUserStatusRequest) (response *DescribeUserStatusResponse, err error) {
    if request == nil {
        request = NewDescribeUserStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoGenerationTaskRequest() (request *DescribeVideoGenerationTaskRequest) {
    request = &DescribeVideoGenerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeVideoGenerationTask")
    
    
    return
}

func NewDescribeVideoGenerationTaskResponse() (response *DescribeVideoGenerationTaskResponse) {
    response = &DescribeVideoGenerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoGenerationTask
// This API is used to query the status and result of a recording video generation task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeVideoGenerationTask(request *DescribeVideoGenerationTaskRequest) (response *DescribeVideoGenerationTaskResponse, err error) {
    return c.DescribeVideoGenerationTaskWithContext(context.Background(), request)
}

// DescribeVideoGenerationTask
// This API is used to query the status and result of a recording video generation task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeVideoGenerationTaskWithContext(ctx context.Context, request *DescribeVideoGenerationTaskRequest) (response *DescribeVideoGenerationTaskResponse, err error) {
    if request == nil {
        request = NewDescribeVideoGenerationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoGenerationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoGenerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoGenerationTaskCallbackRequest() (request *DescribeVideoGenerationTaskCallbackRequest) {
    request = &DescribeVideoGenerationTaskCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeVideoGenerationTaskCallback")
    
    
    return
}

func NewDescribeVideoGenerationTaskCallbackResponse() (response *DescribeVideoGenerationTaskCallbackResponse) {
    response = &DescribeVideoGenerationTaskCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoGenerationTaskCallback
// This API is used to query the callback URL for recording video generation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeVideoGenerationTaskCallback(request *DescribeVideoGenerationTaskCallbackRequest) (response *DescribeVideoGenerationTaskCallbackResponse, err error) {
    return c.DescribeVideoGenerationTaskCallbackWithContext(context.Background(), request)
}

// DescribeVideoGenerationTaskCallback
// This API is used to query the callback URL for recording video generation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeVideoGenerationTaskCallbackWithContext(ctx context.Context, request *DescribeVideoGenerationTaskCallbackRequest) (response *DescribeVideoGenerationTaskCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeVideoGenerationTaskCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoGenerationTaskCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoGenerationTaskCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardApplicationConfigRequest() (request *DescribeWhiteboardApplicationConfigRequest) {
    request = &DescribeWhiteboardApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardApplicationConfig")
    
    
    return
}

func NewDescribeWhiteboardApplicationConfigResponse() (response *DescribeWhiteboardApplicationConfigResponse) {
    response = &DescribeWhiteboardApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWhiteboardApplicationConfig
// This API is used to query the task-related configurations of a whiteboard application, including the bucket and callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardApplicationConfig(request *DescribeWhiteboardApplicationConfigRequest) (response *DescribeWhiteboardApplicationConfigResponse, err error) {
    return c.DescribeWhiteboardApplicationConfigWithContext(context.Background(), request)
}

// DescribeWhiteboardApplicationConfig
// This API is used to query the task-related configurations of a whiteboard application, including the bucket and callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardApplicationConfigWithContext(ctx context.Context, request *DescribeWhiteboardApplicationConfigRequest) (response *DescribeWhiteboardApplicationConfigResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardApplicationConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardBucketConfigRequest() (request *DescribeWhiteboardBucketConfigRequest) {
    request = &DescribeWhiteboardBucketConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardBucketConfig")
    
    
    return
}

func NewDescribeWhiteboardBucketConfigResponse() (response *DescribeWhiteboardBucketConfigResponse) {
    response = &DescribeWhiteboardBucketConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWhiteboardBucketConfig
// This API is used to query the bucket configurations for document transcoding and real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardBucketConfig(request *DescribeWhiteboardBucketConfigRequest) (response *DescribeWhiteboardBucketConfigResponse, err error) {
    return c.DescribeWhiteboardBucketConfigWithContext(context.Background(), request)
}

// DescribeWhiteboardBucketConfig
// This API is used to query the bucket configurations for document transcoding and real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardBucketConfigWithContext(ctx context.Context, request *DescribeWhiteboardBucketConfigRequest) (response *DescribeWhiteboardBucketConfigResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardBucketConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardBucketConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardBucketConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardPushRequest() (request *DescribeWhiteboardPushRequest) {
    request = &DescribeWhiteboardPushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardPush")
    
    
    return
}

func NewDescribeWhiteboardPushResponse() (response *DescribeWhiteboardPushResponse) {
    response = &DescribeWhiteboardPushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWhiteboardPush
// This API is used to query the status and result of a whiteboard push task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_WHITEBOARDPUSH = "FailedOperation.WhiteboardPush"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPush(request *DescribeWhiteboardPushRequest) (response *DescribeWhiteboardPushResponse, err error) {
    return c.DescribeWhiteboardPushWithContext(context.Background(), request)
}

// DescribeWhiteboardPush
// This API is used to query the status and result of a whiteboard push task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_WHITEBOARDPUSH = "FailedOperation.WhiteboardPush"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushWithContext(ctx context.Context, request *DescribeWhiteboardPushRequest) (response *DescribeWhiteboardPushResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardPushRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardPush require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardPushResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardPushCallbackRequest() (request *DescribeWhiteboardPushCallbackRequest) {
    request = &DescribeWhiteboardPushCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardPushCallback")
    
    
    return
}

func NewDescribeWhiteboardPushCallbackResponse() (response *DescribeWhiteboardPushCallbackResponse) {
    response = &DescribeWhiteboardPushCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWhiteboardPushCallback
// This API is used to query the whiteboard push callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushCallback(request *DescribeWhiteboardPushCallbackRequest) (response *DescribeWhiteboardPushCallbackResponse, err error) {
    return c.DescribeWhiteboardPushCallbackWithContext(context.Background(), request)
}

// DescribeWhiteboardPushCallback
// This API is used to query the whiteboard push callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushCallbackWithContext(ctx context.Context, request *DescribeWhiteboardPushCallbackRequest) (response *DescribeWhiteboardPushCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardPushCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardPushCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardPushCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardPushSearchRequest() (request *DescribeWhiteboardPushSearchRequest) {
    request = &DescribeWhiteboardPushSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardPushSearch")
    
    
    return
}

func NewDescribeWhiteboardPushSearchResponse() (response *DescribeWhiteboardPushSearchResponse) {
    response = &DescribeWhiteboardPushSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWhiteboardPushSearch
// This API is used to query whiteboard push tasks by room ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushSearch(request *DescribeWhiteboardPushSearchRequest) (response *DescribeWhiteboardPushSearchResponse, err error) {
    return c.DescribeWhiteboardPushSearchWithContext(context.Background(), request)
}

// DescribeWhiteboardPushSearch
// This API is used to query whiteboard push tasks by room ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushSearchWithContext(ctx context.Context, request *DescribeWhiteboardPushSearchRequest) (response *DescribeWhiteboardPushSearchResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardPushSearchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardPushSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardPushSearchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
    request = &ModifyApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ModifyApplication")
    
    
    return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
    response = &ModifyApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplication
// This API is used to modify a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    return c.ModifyApplicationWithContext(context.Background(), request)
}

// ModifyApplication
// This API is used to modify a whiteboard application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ModifyApplicationWithContext(ctx context.Context, request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoRenewFlagRequest() (request *ModifyAutoRenewFlagRequest) {
    request = &ModifyAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ModifyAutoRenewFlag")
    
    
    return
}

func NewModifyAutoRenewFlagResponse() (response *ModifyAutoRenewFlagResponse) {
    response = &ModifyAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoRenewFlag
// This API is used to set auto-renewal for a Tencent Interactive Whiteboard monthly feature package.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    return c.ModifyAutoRenewFlagWithContext(context.Background(), request)
}

// ModifyAutoRenewFlag
// This API is used to set auto-renewal for a Tencent Interactive Whiteboard monthly feature package.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ModifyAutoRenewFlagWithContext(ctx context.Context, request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWhiteboardApplicationConfigRequest() (request *ModifyWhiteboardApplicationConfigRequest) {
    request = &ModifyWhiteboardApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ModifyWhiteboardApplicationConfig")
    
    
    return
}

func NewModifyWhiteboardApplicationConfigResponse() (response *ModifyWhiteboardApplicationConfigResponse) {
    response = &ModifyWhiteboardApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWhiteboardApplicationConfig
// This API is used to modify the task-related configurations of a whiteboard application, including the bucket and callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COSBUCKETINVALID = "FailedOperation.CosBucketInvalid"
//  FAILEDOPERATION_GETCREDENTIALFAIL = "FailedOperation.GetCredentialFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  INVALIDPARAMETER_CDNDOMAINNOTFOUND = "InvalidParameter.CdnDomainNotFound"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) ModifyWhiteboardApplicationConfig(request *ModifyWhiteboardApplicationConfigRequest) (response *ModifyWhiteboardApplicationConfigResponse, err error) {
    return c.ModifyWhiteboardApplicationConfigWithContext(context.Background(), request)
}

// ModifyWhiteboardApplicationConfig
// This API is used to modify the task-related configurations of a whiteboard application, including the bucket and callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COSBUCKETINVALID = "FailedOperation.CosBucketInvalid"
//  FAILEDOPERATION_GETCREDENTIALFAIL = "FailedOperation.GetCredentialFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  INVALIDPARAMETER_CDNDOMAINNOTFOUND = "InvalidParameter.CdnDomainNotFound"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) ModifyWhiteboardApplicationConfigWithContext(ctx context.Context, request *ModifyWhiteboardApplicationConfigRequest) (response *ModifyWhiteboardApplicationConfigResponse, err error) {
    if request == nil {
        request = NewModifyWhiteboardApplicationConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWhiteboardApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWhiteboardApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWhiteboardBucketConfigRequest() (request *ModifyWhiteboardBucketConfigRequest) {
    request = &ModifyWhiteboardBucketConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ModifyWhiteboardBucketConfig")
    
    
    return
}

func NewModifyWhiteboardBucketConfigResponse() (response *ModifyWhiteboardBucketConfigResponse) {
    response = &ModifyWhiteboardBucketConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWhiteboardBucketConfig
// This API is used to modify the bucket configurations for document transcoding and real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_COSBUCKETINVALID = "FailedOperation.CosBucketInvalid"
//  FAILEDOPERATION_GETCREDENTIALFAIL = "FailedOperation.GetCredentialFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CDNDOMAINNOTFOUND = "InvalidParameter.CdnDomainNotFound"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) ModifyWhiteboardBucketConfig(request *ModifyWhiteboardBucketConfigRequest) (response *ModifyWhiteboardBucketConfigResponse, err error) {
    return c.ModifyWhiteboardBucketConfigWithContext(context.Background(), request)
}

// ModifyWhiteboardBucketConfig
// This API is used to modify the bucket configurations for document transcoding and real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_COSBUCKETINVALID = "FailedOperation.CosBucketInvalid"
//  FAILEDOPERATION_GETCREDENTIALFAIL = "FailedOperation.GetCredentialFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CDNDOMAINNOTFOUND = "InvalidParameter.CdnDomainNotFound"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) ModifyWhiteboardBucketConfigWithContext(ctx context.Context, request *ModifyWhiteboardBucketConfigRequest) (response *ModifyWhiteboardBucketConfigResponse, err error) {
    if request == nil {
        request = NewModifyWhiteboardBucketConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWhiteboardBucketConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWhiteboardBucketConfigResponse()
    err = c.Send(request, response)
    return
}

func NewPauseOnlineRecordRequest() (request *PauseOnlineRecordRequest) {
    request = &PauseOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "PauseOnlineRecord")
    
    
    return
}

func NewPauseOnlineRecordResponse() (response *PauseOnlineRecordResponse) {
    response = &PauseOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseOnlineRecord
// This API is used to pause real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) PauseOnlineRecord(request *PauseOnlineRecordRequest) (response *PauseOnlineRecordResponse, err error) {
    return c.PauseOnlineRecordWithContext(context.Background(), request)
}

// PauseOnlineRecord
// This API is used to pause real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) PauseOnlineRecordWithContext(ctx context.Context, request *PauseOnlineRecordRequest) (response *PauseOnlineRecordResponse, err error) {
    if request == nil {
        request = NewPauseOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewResumeOnlineRecordRequest() (request *ResumeOnlineRecordRequest) {
    request = &ResumeOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ResumeOnlineRecord")
    
    
    return
}

func NewResumeOnlineRecordResponse() (response *ResumeOnlineRecordResponse) {
    response = &ResumeOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeOnlineRecord
// This API is used to resume real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) ResumeOnlineRecord(request *ResumeOnlineRecordRequest) (response *ResumeOnlineRecordResponse, err error) {
    return c.ResumeOnlineRecordWithContext(context.Background(), request)
}

// ResumeOnlineRecord
// This API is used to resume real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) ResumeOnlineRecordWithContext(ctx context.Context, request *ResumeOnlineRecordRequest) (response *ResumeOnlineRecordResponse, err error) {
    if request == nil {
        request = NewResumeOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewSetOfflineRecordCallbackRequest() (request *SetOfflineRecordCallbackRequest) {
    request = &SetOfflineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetOfflineRecordCallback")
    
    
    return
}

func NewSetOfflineRecordCallbackResponse() (response *SetOfflineRecordCallbackResponse) {
    response = &SetOfflineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetOfflineRecordCallback
// 课后录制服务已下线
//
// 
//
// This API is used to set the offline recording callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOfflineRecordCallback(request *SetOfflineRecordCallbackRequest) (response *SetOfflineRecordCallbackResponse, err error) {
    return c.SetOfflineRecordCallbackWithContext(context.Background(), request)
}

// SetOfflineRecordCallback
// 课后录制服务已下线
//
// 
//
// This API is used to set the offline recording callback URL.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOfflineRecordCallbackWithContext(ctx context.Context, request *SetOfflineRecordCallbackRequest) (response *SetOfflineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewSetOfflineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetOfflineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetOfflineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetOnlineRecordCallbackRequest() (request *SetOnlineRecordCallbackRequest) {
    request = &SetOnlineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetOnlineRecordCallback")
    
    
    return
}

func NewSetOnlineRecordCallbackResponse() (response *SetOnlineRecordCallbackResponse) {
    response = &SetOnlineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetOnlineRecordCallback
// This API is used to set the real-time recording callback address. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40258?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallback(request *SetOnlineRecordCallbackRequest) (response *SetOnlineRecordCallbackResponse, err error) {
    return c.SetOnlineRecordCallbackWithContext(context.Background(), request)
}

// SetOnlineRecordCallback
// This API is used to set the real-time recording callback address. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40258?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallbackWithContext(ctx context.Context, request *SetOnlineRecordCallbackRequest) (response *SetOnlineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewSetOnlineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetOnlineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetOnlineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetOnlineRecordCallbackKeyRequest() (request *SetOnlineRecordCallbackKeyRequest) {
    request = &SetOnlineRecordCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetOnlineRecordCallbackKey")
    
    
    return
}

func NewSetOnlineRecordCallbackKeyResponse() (response *SetOnlineRecordCallbackKeyResponse) {
    response = &SetOnlineRecordCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetOnlineRecordCallbackKey
// This API is used to set the callback authentication key for real-time recording. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallbackKey(request *SetOnlineRecordCallbackKeyRequest) (response *SetOnlineRecordCallbackKeyResponse, err error) {
    return c.SetOnlineRecordCallbackKeyWithContext(context.Background(), request)
}

// SetOnlineRecordCallbackKey
// This API is used to set the callback authentication key for real-time recording. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallbackKeyWithContext(ctx context.Context, request *SetOnlineRecordCallbackKeyRequest) (response *SetOnlineRecordCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetOnlineRecordCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetOnlineRecordCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetOnlineRecordCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetTranscodeCallbackRequest() (request *SetTranscodeCallbackRequest) {
    request = &SetTranscodeCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetTranscodeCallback")
    
    
    return
}

func NewSetTranscodeCallbackResponse() (response *SetTranscodeCallbackResponse) {
    response = &SetTranscodeCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetTranscodeCallback
// This API is used to set the callback address for document transcoding. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40260?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallback(request *SetTranscodeCallbackRequest) (response *SetTranscodeCallbackResponse, err error) {
    return c.SetTranscodeCallbackWithContext(context.Background(), request)
}

// SetTranscodeCallback
// This API is used to set the callback address for document transcoding. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40260?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallbackWithContext(ctx context.Context, request *SetTranscodeCallbackRequest) (response *SetTranscodeCallbackResponse, err error) {
    if request == nil {
        request = NewSetTranscodeCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTranscodeCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTranscodeCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetTranscodeCallbackKeyRequest() (request *SetTranscodeCallbackKeyRequest) {
    request = &SetTranscodeCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetTranscodeCallbackKey")
    
    
    return
}

func NewSetTranscodeCallbackKeyResponse() (response *SetTranscodeCallbackKeyResponse) {
    response = &SetTranscodeCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetTranscodeCallbackKey
// This API is used to set the callback authentication key for document transcoding. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallbackKey(request *SetTranscodeCallbackKeyRequest) (response *SetTranscodeCallbackKeyResponse, err error) {
    return c.SetTranscodeCallbackKeyWithContext(context.Background(), request)
}

// SetTranscodeCallbackKey
// This API is used to set the callback authentication key for document transcoding. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallbackKeyWithContext(ctx context.Context, request *SetTranscodeCallbackKeyRequest) (response *SetTranscodeCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetTranscodeCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTranscodeCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTranscodeCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetVideoGenerationTaskCallbackRequest() (request *SetVideoGenerationTaskCallbackRequest) {
    request = &SetVideoGenerationTaskCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetVideoGenerationTaskCallback")
    
    
    return
}

func NewSetVideoGenerationTaskCallbackResponse() (response *SetVideoGenerationTaskCallbackResponse) {
    response = &SetVideoGenerationTaskCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetVideoGenerationTaskCallback
// This API is used to set the callback URL for recording video generation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetVideoGenerationTaskCallback(request *SetVideoGenerationTaskCallbackRequest) (response *SetVideoGenerationTaskCallbackResponse, err error) {
    return c.SetVideoGenerationTaskCallbackWithContext(context.Background(), request)
}

// SetVideoGenerationTaskCallback
// This API is used to set the callback URL for recording video generation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetVideoGenerationTaskCallbackWithContext(ctx context.Context, request *SetVideoGenerationTaskCallbackRequest) (response *SetVideoGenerationTaskCallbackResponse, err error) {
    if request == nil {
        request = NewSetVideoGenerationTaskCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetVideoGenerationTaskCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetVideoGenerationTaskCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetVideoGenerationTaskCallbackKeyRequest() (request *SetVideoGenerationTaskCallbackKeyRequest) {
    request = &SetVideoGenerationTaskCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetVideoGenerationTaskCallbackKey")
    
    
    return
}

func NewSetVideoGenerationTaskCallbackKeyResponse() (response *SetVideoGenerationTaskCallbackKeyResponse) {
    response = &SetVideoGenerationTaskCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetVideoGenerationTaskCallbackKey
// This API is used to set the callback authentication key for recording video generation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetVideoGenerationTaskCallbackKey(request *SetVideoGenerationTaskCallbackKeyRequest) (response *SetVideoGenerationTaskCallbackKeyResponse, err error) {
    return c.SetVideoGenerationTaskCallbackKeyWithContext(context.Background(), request)
}

// SetVideoGenerationTaskCallbackKey
// This API is used to set the callback authentication key for recording video generation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetVideoGenerationTaskCallbackKeyWithContext(ctx context.Context, request *SetVideoGenerationTaskCallbackKeyRequest) (response *SetVideoGenerationTaskCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetVideoGenerationTaskCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetVideoGenerationTaskCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetVideoGenerationTaskCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetWhiteboardPushCallbackRequest() (request *SetWhiteboardPushCallbackRequest) {
    request = &SetWhiteboardPushCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetWhiteboardPushCallback")
    
    
    return
}

func NewSetWhiteboardPushCallbackResponse() (response *SetWhiteboardPushCallbackResponse) {
    response = &SetWhiteboardPushCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetWhiteboardPushCallback
// This API is used to set the whiteboard push callback URL. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWhiteboardPushCallback(request *SetWhiteboardPushCallbackRequest) (response *SetWhiteboardPushCallbackResponse, err error) {
    return c.SetWhiteboardPushCallbackWithContext(context.Background(), request)
}

// SetWhiteboardPushCallback
// This API is used to set the whiteboard push callback URL. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWhiteboardPushCallbackWithContext(ctx context.Context, request *SetWhiteboardPushCallbackRequest) (response *SetWhiteboardPushCallbackResponse, err error) {
    if request == nil {
        request = NewSetWhiteboardPushCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetWhiteboardPushCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetWhiteboardPushCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetWhiteboardPushCallbackKeyRequest() (request *SetWhiteboardPushCallbackKeyRequest) {
    request = &SetWhiteboardPushCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetWhiteboardPushCallbackKey")
    
    
    return
}

func NewSetWhiteboardPushCallbackKeyResponse() (response *SetWhiteboardPushCallbackKeyResponse) {
    response = &SetWhiteboardPushCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetWhiteboardPushCallbackKey
// This API is used to set the callback authentication key for whiteboard push. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWhiteboardPushCallbackKey(request *SetWhiteboardPushCallbackKeyRequest) (response *SetWhiteboardPushCallbackKeyResponse, err error) {
    return c.SetWhiteboardPushCallbackKeyWithContext(context.Background(), request)
}

// SetWhiteboardPushCallbackKey
// This API is used to set the callback authentication key for whiteboard push. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWhiteboardPushCallbackKeyWithContext(ctx context.Context, request *SetWhiteboardPushCallbackKeyRequest) (response *SetWhiteboardPushCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetWhiteboardPushCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetWhiteboardPushCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetWhiteboardPushCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewStartOnlineRecordRequest() (request *StartOnlineRecordRequest) {
    request = &StartOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "StartOnlineRecord")
    
    
    return
}

func NewStartOnlineRecordResponse() (response *StartOnlineRecordResponse) {
    response = &StartOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartOnlineRecord
// This API is used to start a real-time recording task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartOnlineRecord(request *StartOnlineRecordRequest) (response *StartOnlineRecordResponse, err error) {
    return c.StartOnlineRecordWithContext(context.Background(), request)
}

// StartOnlineRecord
// This API is used to start a real-time recording task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartOnlineRecordWithContext(ctx context.Context, request *StartOnlineRecordRequest) (response *StartOnlineRecordResponse, err error) {
    if request == nil {
        request = NewStartOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStartWhiteboardPushRequest() (request *StartWhiteboardPushRequest) {
    request = &StartWhiteboardPushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "StartWhiteboardPush")
    
    
    return
}

func NewStartWhiteboardPushResponse() (response *StartWhiteboardPushResponse) {
    response = &StartWhiteboardPushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartWhiteboardPush
// This API is used to start a whiteboard push task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartWhiteboardPush(request *StartWhiteboardPushRequest) (response *StartWhiteboardPushResponse, err error) {
    return c.StartWhiteboardPushWithContext(context.Background(), request)
}

// StartWhiteboardPush
// This API is used to start a whiteboard push task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartWhiteboardPushWithContext(ctx context.Context, request *StartWhiteboardPushRequest) (response *StartWhiteboardPushResponse, err error) {
    if request == nil {
        request = NewStartWhiteboardPushRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartWhiteboardPush require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartWhiteboardPushResponse()
    err = c.Send(request, response)
    return
}

func NewStopOnlineRecordRequest() (request *StopOnlineRecordRequest) {
    request = &StopOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "StopOnlineRecord")
    
    
    return
}

func NewStopOnlineRecordResponse() (response *StopOnlineRecordResponse) {
    response = &StopOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopOnlineRecord
// This API is used to stop real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopOnlineRecord(request *StopOnlineRecordRequest) (response *StopOnlineRecordResponse, err error) {
    return c.StopOnlineRecordWithContext(context.Background(), request)
}

// StopOnlineRecord
// This API is used to stop real-time recording.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopOnlineRecordWithContext(ctx context.Context, request *StopOnlineRecordRequest) (response *StopOnlineRecordResponse, err error) {
    if request == nil {
        request = NewStopOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopWhiteboardPushRequest() (request *StopWhiteboardPushRequest) {
    request = &StopWhiteboardPushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "StopWhiteboardPush")
    
    
    return
}

func NewStopWhiteboardPushResponse() (response *StopWhiteboardPushResponse) {
    response = &StopWhiteboardPushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopWhiteboardPush
// This API is used to stop a whiteboard push task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopWhiteboardPush(request *StopWhiteboardPushRequest) (response *StopWhiteboardPushResponse, err error) {
    return c.StopWhiteboardPushWithContext(context.Background(), request)
}

// StopWhiteboardPush
// This API is used to stop a whiteboard push task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopWhiteboardPushWithContext(ctx context.Context, request *StopWhiteboardPushRequest) (response *StopWhiteboardPushResponse, err error) {
    if request == nil {
        request = NewStopWhiteboardPushRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopWhiteboardPush require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopWhiteboardPushResponse()
    err = c.Send(request, response)
    return
}
