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

package v20210728

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-07-28"

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


func NewAbortCronJobsRequest() (request *AbortCronJobsRequest) {
    request = &AbortCronJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "AbortCronJobs")
    
    
    return
}

func NewAbortCronJobsResponse() (response *AbortCronJobsResponse) {
    response = &AbortCronJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AbortCronJobs
// This API is used to stop cron jobs.
func (c *Client) AbortCronJobs(request *AbortCronJobsRequest) (response *AbortCronJobsResponse, err error) {
    return c.AbortCronJobsWithContext(context.Background(), request)
}

// AbortCronJobs
// This API is used to stop cron jobs.
func (c *Client) AbortCronJobsWithContext(ctx context.Context, request *AbortCronJobsRequest) (response *AbortCronJobsResponse, err error) {
    if request == nil {
        request = NewAbortCronJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AbortCronJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewAbortCronJobsResponse()
    err = c.Send(request, response)
    return
}

func NewAbortJobRequest() (request *AbortJobRequest) {
    request = &AbortJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "AbortJob")
    
    
    return
}

func NewAbortJobResponse() (response *AbortJobResponse) {
    response = &AbortJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AbortJob
// This API is used to stop test job.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_JOBSTATUSNOTRUNNING = "FailedOperation.JobStatusNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AbortJob(request *AbortJobRequest) (response *AbortJobResponse, err error) {
    return c.AbortJobWithContext(context.Background(), request)
}

// AbortJob
// This API is used to stop test job.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_JOBSTATUSNOTRUNNING = "FailedOperation.JobStatusNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AbortJobWithContext(ctx context.Context, request *AbortJobRequest) (response *AbortJobResponse, err error) {
    if request == nil {
        request = NewAbortJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AbortJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewAbortJobResponse()
    err = c.Send(request, response)
    return
}

func NewAdjustJobSpeedRequest() (request *AdjustJobSpeedRequest) {
    request = &AdjustJobSpeedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "AdjustJobSpeed")
    
    
    return
}

func NewAdjustJobSpeedResponse() (response *AdjustJobSpeedResponse) {
    response = &AdjustJobSpeedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AdjustJobSpeed
// This API is used to adjust the target RPS of a job.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AdjustJobSpeed(request *AdjustJobSpeedRequest) (response *AdjustJobSpeedResponse, err error) {
    return c.AdjustJobSpeedWithContext(context.Background(), request)
}

// AdjustJobSpeed
// This API is used to adjust the target RPS of a job.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AdjustJobSpeedWithContext(ctx context.Context, request *AdjustJobSpeedRequest) (response *AdjustJobSpeedResponse, err error) {
    if request == nil {
        request = NewAdjustJobSpeedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AdjustJobSpeed require credential")
    }

    request.SetContext(ctx)
    
    response = NewAdjustJobSpeedResponse()
    err = c.Send(request, response)
    return
}

func NewCopyScenarioRequest() (request *CopyScenarioRequest) {
    request = &CopyScenarioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CopyScenario")
    
    
    return
}

func NewCopyScenarioResponse() (response *CopyScenarioResponse) {
    response = &CopyScenarioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyScenario
// This API is used to copy a scenario.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) CopyScenario(request *CopyScenarioRequest) (response *CopyScenarioResponse, err error) {
    return c.CopyScenarioWithContext(context.Background(), request)
}

// CopyScenario
// This API is used to copy a scenario.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) CopyScenarioWithContext(ctx context.Context, request *CopyScenarioRequest) (response *CopyScenarioResponse, err error) {
    if request == nil {
        request = NewCopyScenarioRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyScenario require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyScenarioResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlertChannelRequest() (request *CreateAlertChannelRequest) {
    request = &CreateAlertChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateAlertChannel")
    
    
    return
}

func NewCreateAlertChannelResponse() (response *CreateAlertChannelResponse) {
    response = &CreateAlertChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlertChannel
// This API is used to create an alert recipient group.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateAlertChannel(request *CreateAlertChannelRequest) (response *CreateAlertChannelResponse, err error) {
    return c.CreateAlertChannelWithContext(context.Background(), request)
}

// CreateAlertChannel
// This API is used to create an alert recipient group.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateAlertChannelWithContext(ctx context.Context, request *CreateAlertChannelRequest) (response *CreateAlertChannelResponse, err error) {
    if request == nil {
        request = NewCreateAlertChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlertChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlertChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCronJobRequest() (request *CreateCronJobRequest) {
    request = &CreateCronJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateCronJob")
    
    
    return
}

func NewCreateCronJobResponse() (response *CreateCronJobResponse) {
    response = &CreateCronJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCronJob
// This API is used to create a cron job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateCronJob(request *CreateCronJobRequest) (response *CreateCronJobResponse, err error) {
    return c.CreateCronJobWithContext(context.Background(), request)
}

// CreateCronJob
// This API is used to create a cron job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateCronJobWithContext(ctx context.Context, request *CreateCronJobRequest) (response *CreateCronJobResponse, err error) {
    if request == nil {
        request = NewCreateCronJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCronJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCronJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvironmentRequest() (request *CreateEnvironmentRequest) {
    request = &CreateEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateEnvironment")
    
    
    return
}

func NewCreateEnvironmentResponse() (response *CreateEnvironmentResponse) {
    response = &CreateEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnvironment
// This API is used to create an environment.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    return c.CreateEnvironmentWithContext(context.Background(), request)
}

// CreateEnvironment
// This API is used to create an environment.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateEnvironmentWithContext(ctx context.Context, request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileRequest() (request *CreateFileRequest) {
    request = &CreateFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateFile")
    
    
    return
}

func NewCreateFileResponse() (response *CreateFileResponse) {
    response = &CreateFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFile
// This API is used to create a file.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateFile(request *CreateFileRequest) (response *CreateFileResponse, err error) {
    return c.CreateFileWithContext(context.Background(), request)
}

// CreateFile
// This API is used to create a file.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateFileWithContext(ctx context.Context, request *CreateFileRequest) (response *CreateFileResponse, err error) {
    if request == nil {
        request = NewCreateFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProject
// This API is used to create a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// This API is used to create a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScenarioRequest() (request *CreateScenarioRequest) {
    request = &CreateScenarioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateScenario")
    
    
    return
}

func NewCreateScenarioResponse() (response *CreateScenarioResponse) {
    response = &CreateScenarioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateScenario
// This API is used to create a scenario.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateScenario(request *CreateScenarioRequest) (response *CreateScenarioResponse, err error) {
    return c.CreateScenarioWithContext(context.Background(), request)
}

// CreateScenario
// This API is used to create a scenario.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateScenarioWithContext(ctx context.Context, request *CreateScenarioRequest) (response *CreateScenarioResponse, err error) {
    if request == nil {
        request = NewCreateScenarioRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScenario require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScenarioResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlertChannelRequest() (request *DeleteAlertChannelRequest) {
    request = &DeleteAlertChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteAlertChannel")
    
    
    return
}

func NewDeleteAlertChannelResponse() (response *DeleteAlertChannelResponse) {
    response = &DeleteAlertChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlertChannel
// This API is used to delete an alert recipient group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteAlertChannel(request *DeleteAlertChannelRequest) (response *DeleteAlertChannelResponse, err error) {
    return c.DeleteAlertChannelWithContext(context.Background(), request)
}

// DeleteAlertChannel
// This API is used to delete an alert recipient group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteAlertChannelWithContext(ctx context.Context, request *DeleteAlertChannelRequest) (response *DeleteAlertChannelResponse, err error) {
    if request == nil {
        request = NewDeleteAlertChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlertChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlertChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCronJobsRequest() (request *DeleteCronJobsRequest) {
    request = &DeleteCronJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteCronJobs")
    
    
    return
}

func NewDeleteCronJobsResponse() (response *DeleteCronJobsResponse) {
    response = &DeleteCronJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCronJobs
// This API is used to delete cron jobs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteCronJobs(request *DeleteCronJobsRequest) (response *DeleteCronJobsResponse, err error) {
    return c.DeleteCronJobsWithContext(context.Background(), request)
}

// DeleteCronJobs
// This API is used to delete cron jobs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteCronJobsWithContext(ctx context.Context, request *DeleteCronJobsRequest) (response *DeleteCronJobsResponse, err error) {
    if request == nil {
        request = NewDeleteCronJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCronJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCronJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEnvironmentsRequest() (request *DeleteEnvironmentsRequest) {
    request = &DeleteEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteEnvironments")
    
    
    return
}

func NewDeleteEnvironmentsResponse() (response *DeleteEnvironmentsResponse) {
    response = &DeleteEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEnvironments
// This API is used to delete environments.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteEnvironments(request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
    return c.DeleteEnvironmentsWithContext(context.Background(), request)
}

// DeleteEnvironments
// This API is used to delete environments.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteEnvironmentsWithContext(ctx context.Context, request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDeleteEnvironmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFilesRequest() (request *DeleteFilesRequest) {
    request = &DeleteFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteFiles")
    
    
    return
}

func NewDeleteFilesResponse() (response *DeleteFilesResponse) {
    response = &DeleteFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFiles
// This API is used to delete files.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteFiles(request *DeleteFilesRequest) (response *DeleteFilesResponse, err error) {
    return c.DeleteFilesWithContext(context.Background(), request)
}

// DeleteFiles
// This API is used to delete files.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteFilesWithContext(ctx context.Context, request *DeleteFilesRequest) (response *DeleteFilesResponse, err error) {
    if request == nil {
        request = NewDeleteFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobsRequest() (request *DeleteJobsRequest) {
    request = &DeleteJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteJobs")
    
    
    return
}

func NewDeleteJobsResponse() (response *DeleteJobsResponse) {
    response = &DeleteJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteJobs
// This API is used to delete jobs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteJobs(request *DeleteJobsRequest) (response *DeleteJobsResponse, err error) {
    return c.DeleteJobsWithContext(context.Background(), request)
}

// DeleteJobs
// This API is used to delete jobs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteJobsWithContext(ctx context.Context, request *DeleteJobsRequest) (response *DeleteJobsResponse, err error) {
    if request == nil {
        request = NewDeleteJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectsRequest() (request *DeleteProjectsRequest) {
    request = &DeleteProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteProjects")
    
    
    return
}

func NewDeleteProjectsResponse() (response *DeleteProjectsResponse) {
    response = &DeleteProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProjects
// This API is used to delete projects.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteProjects(request *DeleteProjectsRequest) (response *DeleteProjectsResponse, err error) {
    return c.DeleteProjectsWithContext(context.Background(), request)
}

// DeleteProjects
// This API is used to delete projects.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteProjectsWithContext(ctx context.Context, request *DeleteProjectsRequest) (response *DeleteProjectsResponse, err error) {
    if request == nil {
        request = NewDeleteProjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScenariosRequest() (request *DeleteScenariosRequest) {
    request = &DeleteScenariosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteScenarios")
    
    
    return
}

func NewDeleteScenariosResponse() (response *DeleteScenariosResponse) {
    response = &DeleteScenariosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteScenarios
// This API is used to delete scenarios.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteScenarios(request *DeleteScenariosRequest) (response *DeleteScenariosResponse, err error) {
    return c.DeleteScenariosWithContext(context.Background(), request)
}

// DeleteScenarios
// This API is used to delete scenarios.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteScenariosWithContext(ctx context.Context, request *DeleteScenariosRequest) (response *DeleteScenariosResponse, err error) {
    if request == nil {
        request = NewDeleteScenariosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScenarios require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScenariosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertChannelsRequest() (request *DescribeAlertChannelsRequest) {
    request = &DescribeAlertChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeAlertChannels")
    
    
    return
}

func NewDescribeAlertChannelsResponse() (response *DescribeAlertChannelsResponse) {
    response = &DescribeAlertChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertChannels
// This API is used to query alert recipient groups.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAlertChannels(request *DescribeAlertChannelsRequest) (response *DescribeAlertChannelsResponse, err error) {
    return c.DescribeAlertChannelsWithContext(context.Background(), request)
}

// DescribeAlertChannels
// This API is used to query alert recipient groups.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAlertChannelsWithContext(ctx context.Context, request *DescribeAlertChannelsRequest) (response *DescribeAlertChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeAlertChannelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertRecordsRequest() (request *DescribeAlertRecordsRequest) {
    request = &DescribeAlertRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeAlertRecords")
    
    
    return
}

func NewDescribeAlertRecordsResponse() (response *DescribeAlertRecordsResponse) {
    response = &DescribeAlertRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertRecords
// This API is used to query alert records.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAlertRecords(request *DescribeAlertRecordsRequest) (response *DescribeAlertRecordsResponse, err error) {
    return c.DescribeAlertRecordsWithContext(context.Background(), request)
}

// DescribeAlertRecords
// This API is used to query alert records.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAlertRecordsWithContext(ctx context.Context, request *DescribeAlertRecordsRequest) (response *DescribeAlertRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeAlertRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableMetricsRequest() (request *DescribeAvailableMetricsRequest) {
    request = &DescribeAvailableMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeAvailableMetrics")
    
    
    return
}

func NewDescribeAvailableMetricsResponse() (response *DescribeAvailableMetricsResponse) {
    response = &DescribeAvailableMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailableMetrics
// This API is used to query all supported metrics.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAvailableMetrics(request *DescribeAvailableMetricsRequest) (response *DescribeAvailableMetricsResponse, err error) {
    return c.DescribeAvailableMetricsWithContext(context.Background(), request)
}

// DescribeAvailableMetrics
// This API is used to query all supported metrics.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAvailableMetricsWithContext(ctx context.Context, request *DescribeAvailableMetricsRequest) (response *DescribeAvailableMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckSummaryRequest() (request *DescribeCheckSummaryRequest) {
    request = &DescribeCheckSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeCheckSummary")
    
    
    return
}

func NewDescribeCheckSummaryResponse() (response *DescribeCheckSummaryResponse) {
    response = &DescribeCheckSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCheckSummary
// This API is used to query checkpoint summary information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeCheckSummary(request *DescribeCheckSummaryRequest) (response *DescribeCheckSummaryResponse, err error) {
    return c.DescribeCheckSummaryWithContext(context.Background(), request)
}

// DescribeCheckSummary
// This API is used to query checkpoint summary information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeCheckSummaryWithContext(ctx context.Context, request *DescribeCheckSummaryRequest) (response *DescribeCheckSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeCheckSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCronJobsRequest() (request *DescribeCronJobsRequest) {
    request = &DescribeCronJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeCronJobs")
    
    
    return
}

func NewDescribeCronJobsResponse() (response *DescribeCronJobsResponse) {
    response = &DescribeCronJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCronJobs
// This API is used to list cron jobs, selecting all by default if a non-mandatory array parameter is empty.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeCronJobs(request *DescribeCronJobsRequest) (response *DescribeCronJobsResponse, err error) {
    return c.DescribeCronJobsWithContext(context.Background(), request)
}

// DescribeCronJobs
// This API is used to list cron jobs, selecting all by default if a non-mandatory array parameter is empty.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeCronJobsWithContext(ctx context.Context, request *DescribeCronJobsRequest) (response *DescribeCronJobsResponse, err error) {
    if request == nil {
        request = NewDescribeCronJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCronJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCronJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeEnvironments")
    
    
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironments
// This API is used to query the environment list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// This API is used to query the environment list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeErrorSummaryRequest() (request *DescribeErrorSummaryRequest) {
    request = &DescribeErrorSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeErrorSummary")
    
    
    return
}

func NewDescribeErrorSummaryResponse() (response *DescribeErrorSummaryResponse) {
    response = &DescribeErrorSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeErrorSummary
// This API is used to query error summary information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeErrorSummary(request *DescribeErrorSummaryRequest) (response *DescribeErrorSummaryResponse, err error) {
    return c.DescribeErrorSummaryWithContext(context.Background(), request)
}

// DescribeErrorSummary
// This API is used to query error summary information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeErrorSummaryWithContext(ctx context.Context, request *DescribeErrorSummaryRequest) (response *DescribeErrorSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeErrorSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeErrorSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeErrorSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFilesRequest() (request *DescribeFilesRequest) {
    request = &DescribeFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeFiles")
    
    
    return
}

func NewDescribeFilesResponse() (response *DescribeFilesResponse) {
    response = &DescribeFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFiles
// This API is used to query file list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFiles(request *DescribeFilesRequest) (response *DescribeFilesResponse, err error) {
    return c.DescribeFilesWithContext(context.Background(), request)
}

// DescribeFiles
// This API is used to query file list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFilesWithContext(ctx context.Context, request *DescribeFilesRequest) (response *DescribeFilesResponse, err error) {
    if request == nil {
        request = NewDescribeFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeJobs")
    
    
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobs
// This API is used to query job list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    return c.DescribeJobsWithContext(context.Background(), request)
}

// DescribeJobs
// This API is used to query job list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeJobsWithContext(ctx context.Context, request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLabelValuesRequest() (request *DescribeLabelValuesRequest) {
    request = &DescribeLabelValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeLabelValues")
    
    
    return
}

func NewDescribeLabelValuesResponse() (response *DescribeLabelValuesResponse) {
    response = &DescribeLabelValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLabelValues
// This API is used to query label values.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLabelValues(request *DescribeLabelValuesRequest) (response *DescribeLabelValuesResponse, err error) {
    return c.DescribeLabelValuesWithContext(context.Background(), request)
}

// DescribeLabelValues
// This API is used to query label values.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLabelValuesWithContext(ctx context.Context, request *DescribeLabelValuesRequest) (response *DescribeLabelValuesResponse, err error) {
    if request == nil {
        request = NewDescribeLabelValuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLabelValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLabelValuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricLabelWithValuesRequest() (request *DescribeMetricLabelWithValuesRequest) {
    request = &DescribeMetricLabelWithValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeMetricLabelWithValues")
    
    
    return
}

func NewDescribeMetricLabelWithValuesResponse() (response *DescribeMetricLabelWithValuesResponse) {
    response = &DescribeMetricLabelWithValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricLabelWithValues
// This API is used to query all labels and values of metrics
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricLabelWithValues(request *DescribeMetricLabelWithValuesRequest) (response *DescribeMetricLabelWithValuesResponse, err error) {
    return c.DescribeMetricLabelWithValuesWithContext(context.Background(), request)
}

// DescribeMetricLabelWithValues
// This API is used to query all labels and values of metrics
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricLabelWithValuesWithContext(ctx context.Context, request *DescribeMetricLabelWithValuesRequest) (response *DescribeMetricLabelWithValuesResponse, err error) {
    if request == nil {
        request = NewDescribeMetricLabelWithValuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricLabelWithValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricLabelWithValuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNormalLogsRequest() (request *DescribeNormalLogsRequest) {
    request = &DescribeNormalLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeNormalLogs")
    
    
    return
}

func NewDescribeNormalLogsResponse() (response *DescribeNormalLogsResponse) {
    response = &DescribeNormalLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNormalLogs
// This API is used to query logs in performance testing, including engine logs and console logs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNormalLogs(request *DescribeNormalLogsRequest) (response *DescribeNormalLogsResponse, err error) {
    return c.DescribeNormalLogsWithContext(context.Background(), request)
}

// DescribeNormalLogs
// This API is used to query logs in performance testing, including engine logs and console logs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNormalLogsWithContext(ctx context.Context, request *DescribeNormalLogsRequest) (response *DescribeNormalLogsResponse, err error) {
    if request == nil {
        request = NewDescribeNormalLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNormalLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNormalLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjects
// The API is used to query project list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    return c.DescribeProjectsWithContext(context.Background(), request)
}

// DescribeProjects
// The API is used to query project list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegions
// This API is used to query region list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// This API is used to query region list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRequestSummaryRequest() (request *DescribeRequestSummaryRequest) {
    request = &DescribeRequestSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeRequestSummary")
    
    
    return
}

func NewDescribeRequestSummaryResponse() (response *DescribeRequestSummaryResponse) {
    response = &DescribeRequestSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRequestSummary
// This API is used to query request summary information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRequestSummary(request *DescribeRequestSummaryRequest) (response *DescribeRequestSummaryResponse, err error) {
    return c.DescribeRequestSummaryWithContext(context.Background(), request)
}

// DescribeRequestSummary
// This API is used to query request summary information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRequestSummaryWithContext(ctx context.Context, request *DescribeRequestSummaryRequest) (response *DescribeRequestSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeRequestSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRequestSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRequestSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleBatchQueryRequest() (request *DescribeSampleBatchQueryRequest) {
    request = &DescribeSampleBatchQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleBatchQuery")
    
    
    return
}

func NewDescribeSampleBatchQueryResponse() (response *DescribeSampleBatchQueryResponse) {
    response = &DescribeSampleBatchQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleBatchQuery
// This API is used to query metrics in batch and return metric content at a specific time point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSampleBatchQuery(request *DescribeSampleBatchQueryRequest) (response *DescribeSampleBatchQueryResponse, err error) {
    return c.DescribeSampleBatchQueryWithContext(context.Background(), request)
}

// DescribeSampleBatchQuery
// This API is used to query metrics in batch and return metric content at a specific time point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSampleBatchQueryWithContext(ctx context.Context, request *DescribeSampleBatchQueryRequest) (response *DescribeSampleBatchQueryResponse, err error) {
    if request == nil {
        request = NewDescribeSampleBatchQueryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleBatchQuery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleBatchQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleLogsRequest() (request *DescribeSampleLogsRequest) {
    request = &DescribeSampleLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleLogs")
    
    
    return
}

func NewDescribeSampleLogsResponse() (response *DescribeSampleLogsResponse) {
    response = &DescribeSampleLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleLogs
// This API is used to query sampled request logs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSampleLogs(request *DescribeSampleLogsRequest) (response *DescribeSampleLogsResponse, err error) {
    return c.DescribeSampleLogsWithContext(context.Background(), request)
}

// DescribeSampleLogs
// This API is used to query sampled request logs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSampleLogsWithContext(ctx context.Context, request *DescribeSampleLogsRequest) (response *DescribeSampleLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSampleLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleMatrixBatchQueryRequest() (request *DescribeSampleMatrixBatchQueryRequest) {
    request = &DescribeSampleMatrixBatchQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleMatrixBatchQuery")
    
    
    return
}

func NewDescribeSampleMatrixBatchQueryResponse() (response *DescribeSampleMatrixBatchQueryResponse) {
    response = &DescribeSampleMatrixBatchQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleMatrixBatchQuery
// This API is used to batch query metric matrices.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSampleMatrixBatchQuery(request *DescribeSampleMatrixBatchQueryRequest) (response *DescribeSampleMatrixBatchQueryResponse, err error) {
    return c.DescribeSampleMatrixBatchQueryWithContext(context.Background(), request)
}

// DescribeSampleMatrixBatchQuery
// This API is used to batch query metric matrices.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSampleMatrixBatchQueryWithContext(ctx context.Context, request *DescribeSampleMatrixBatchQueryRequest) (response *DescribeSampleMatrixBatchQueryResponse, err error) {
    if request == nil {
        request = NewDescribeSampleMatrixBatchQueryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleMatrixBatchQuery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleMatrixBatchQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleMatrixQueryRequest() (request *DescribeSampleMatrixQueryRequest) {
    request = &DescribeSampleMatrixQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleMatrixQuery")
    
    
    return
}

func NewDescribeSampleMatrixQueryResponse() (response *DescribeSampleMatrixQueryResponse) {
    response = &DescribeSampleMatrixQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleMatrixQuery
// This API is used to query metric matrix.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSampleMatrixQuery(request *DescribeSampleMatrixQueryRequest) (response *DescribeSampleMatrixQueryResponse, err error) {
    return c.DescribeSampleMatrixQueryWithContext(context.Background(), request)
}

// DescribeSampleMatrixQuery
// This API is used to query metric matrix.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSampleMatrixQueryWithContext(ctx context.Context, request *DescribeSampleMatrixQueryRequest) (response *DescribeSampleMatrixQueryResponse, err error) {
    if request == nil {
        request = NewDescribeSampleMatrixQueryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleMatrixQuery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleMatrixQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleQueryRequest() (request *DescribeSampleQueryRequest) {
    request = &DescribeSampleQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleQuery")
    
    
    return
}

func NewDescribeSampleQueryResponse() (response *DescribeSampleQueryResponse) {
    response = &DescribeSampleQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleQuery
// This API is used to query metrics and return metric content at a specific time point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSampleQuery(request *DescribeSampleQueryRequest) (response *DescribeSampleQueryResponse, err error) {
    return c.DescribeSampleQueryWithContext(context.Background(), request)
}

// DescribeSampleQuery
// This API is used to query metrics and return metric content at a specific time point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSampleQueryWithContext(ctx context.Context, request *DescribeSampleQueryRequest) (response *DescribeSampleQueryResponse, err error) {
    if request == nil {
        request = NewDescribeSampleQueryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleQuery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScenarioWithJobsRequest() (request *DescribeScenarioWithJobsRequest) {
    request = &DescribeScenarioWithJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeScenarioWithJobs")
    
    
    return
}

func NewDescribeScenarioWithJobsResponse() (response *DescribeScenarioWithJobsResponse) {
    response = &DescribeScenarioWithJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScenarioWithJobs
// This API is used to query scenarios with executed jobs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeScenarioWithJobs(request *DescribeScenarioWithJobsRequest) (response *DescribeScenarioWithJobsResponse, err error) {
    return c.DescribeScenarioWithJobsWithContext(context.Background(), request)
}

// DescribeScenarioWithJobs
// This API is used to query scenarios with executed jobs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeScenarioWithJobsWithContext(ctx context.Context, request *DescribeScenarioWithJobsRequest) (response *DescribeScenarioWithJobsResponse, err error) {
    if request == nil {
        request = NewDescribeScenarioWithJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScenarioWithJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScenarioWithJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScenariosRequest() (request *DescribeScenariosRequest) {
    request = &DescribeScenariosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeScenarios")
    
    
    return
}

func NewDescribeScenariosResponse() (response *DescribeScenariosResponse) {
    response = &DescribeScenariosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScenarios
// This API is used to query scenario list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeScenarios(request *DescribeScenariosRequest) (response *DescribeScenariosResponse, err error) {
    return c.DescribeScenariosWithContext(context.Background(), request)
}

// DescribeScenarios
// This API is used to query scenario list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeScenariosWithContext(ctx context.Context, request *DescribeScenariosRequest) (response *DescribeScenariosResponse, err error) {
    if request == nil {
        request = NewDescribeScenariosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScenarios require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScenariosResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateTmpKeyRequest() (request *GenerateTmpKeyRequest) {
    request = &GenerateTmpKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "GenerateTmpKey")
    
    
    return
}

func NewGenerateTmpKeyResponse() (response *GenerateTmpKeyResponse) {
    response = &GenerateTmpKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateTmpKey
// This API is used to generate temporary COS credentials.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GenerateTmpKey(request *GenerateTmpKeyRequest) (response *GenerateTmpKeyResponse, err error) {
    return c.GenerateTmpKeyWithContext(context.Background(), request)
}

// GenerateTmpKey
// This API is used to generate temporary COS credentials.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GenerateTmpKeyWithContext(ctx context.Context, request *GenerateTmpKeyRequest) (response *GenerateTmpKeyResponse, err error) {
    if request == nil {
        request = NewGenerateTmpKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateTmpKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateTmpKeyResponse()
    err = c.Send(request, response)
    return
}

func NewRestartCronJobsRequest() (request *RestartCronJobsRequest) {
    request = &RestartCronJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "RestartCronJobs")
    
    
    return
}

func NewRestartCronJobsResponse() (response *RestartCronJobsResponse) {
    response = &RestartCronJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartCronJobs
// This API is used to restart cron jobs that have been aborted.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RestartCronJobs(request *RestartCronJobsRequest) (response *RestartCronJobsResponse, err error) {
    return c.RestartCronJobsWithContext(context.Background(), request)
}

// RestartCronJobs
// This API is used to restart cron jobs that have been aborted.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RestartCronJobsWithContext(ctx context.Context, request *RestartCronJobsRequest) (response *RestartCronJobsResponse, err error) {
    if request == nil {
        request = NewRestartCronJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartCronJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartCronJobsResponse()
    err = c.Send(request, response)
    return
}

func NewStartJobRequest() (request *StartJobRequest) {
    request = &StartJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "StartJob")
    
    
    return
}

func NewStartJobResponse() (response *StartJobResponse) {
    response = &StartJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartJob
// This API is used to create and start test job.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOTASKSINJOB = "FailedOperation.NoTasksInJob"
//  FAILEDOPERATION_NOTSUPPORTEDINENV = "FailedOperation.NotSupportedInEnv"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) StartJob(request *StartJobRequest) (response *StartJobResponse, err error) {
    return c.StartJobWithContext(context.Background(), request)
}

// StartJob
// This API is used to create and start test job.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOTASKSINJOB = "FailedOperation.NoTasksInJob"
//  FAILEDOPERATION_NOTSUPPORTEDINENV = "FailedOperation.NotSupportedInEnv"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) StartJobWithContext(ctx context.Context, request *StartJobRequest) (response *StartJobResponse, err error) {
    if request == nil {
        request = NewStartJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCronJobRequest() (request *UpdateCronJobRequest) {
    request = &UpdateCronJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateCronJob")
    
    
    return
}

func NewUpdateCronJobResponse() (response *UpdateCronJobResponse) {
    response = &UpdateCronJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCronJob
// This API is used to update a cron job.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpdateCronJob(request *UpdateCronJobRequest) (response *UpdateCronJobResponse, err error) {
    return c.UpdateCronJobWithContext(context.Background(), request)
}

// UpdateCronJob
// This API is used to update a cron job.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpdateCronJobWithContext(ctx context.Context, request *UpdateCronJobRequest) (response *UpdateCronJobResponse, err error) {
    if request == nil {
        request = NewUpdateCronJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCronJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCronJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEnvironmentRequest() (request *UpdateEnvironmentRequest) {
    request = &UpdateEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateEnvironment")
    
    
    return
}

func NewUpdateEnvironmentResponse() (response *UpdateEnvironmentResponse) {
    response = &UpdateEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateEnvironment
// This API is used to update environments.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateEnvironment(request *UpdateEnvironmentRequest) (response *UpdateEnvironmentResponse, err error) {
    return c.UpdateEnvironmentWithContext(context.Background(), request)
}

// UpdateEnvironment
// This API is used to update environments.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateEnvironmentWithContext(ctx context.Context, request *UpdateEnvironmentRequest) (response *UpdateEnvironmentResponse, err error) {
    if request == nil {
        request = NewUpdateEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFileScenarioRelationRequest() (request *UpdateFileScenarioRelationRequest) {
    request = &UpdateFileScenarioRelationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateFileScenarioRelation")
    
    
    return
}

func NewUpdateFileScenarioRelationResponse() (response *UpdateFileScenarioRelationResponse) {
    response = &UpdateFileScenarioRelationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateFileScenarioRelation
// This API is used to update relation between files and scenarios.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) UpdateFileScenarioRelation(request *UpdateFileScenarioRelationRequest) (response *UpdateFileScenarioRelationResponse, err error) {
    return c.UpdateFileScenarioRelationWithContext(context.Background(), request)
}

// UpdateFileScenarioRelation
// This API is used to update relation between files and scenarios.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) UpdateFileScenarioRelationWithContext(ctx context.Context, request *UpdateFileScenarioRelationRequest) (response *UpdateFileScenarioRelationResponse, err error) {
    if request == nil {
        request = NewUpdateFileScenarioRelationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFileScenarioRelation require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateFileScenarioRelationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateJobRequest() (request *UpdateJobRequest) {
    request = &UpdateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateJob")
    
    
    return
}

func NewUpdateJobResponse() (response *UpdateJobResponse) {
    response = &UpdateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateJob
// This API is used to update a job.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateJob(request *UpdateJobRequest) (response *UpdateJobResponse, err error) {
    return c.UpdateJobWithContext(context.Background(), request)
}

// UpdateJob
// This API is used to update a job.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateJobWithContext(ctx context.Context, request *UpdateJobRequest) (response *UpdateJobResponse, err error) {
    if request == nil {
        request = NewUpdateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProjectRequest() (request *UpdateProjectRequest) {
    request = &UpdateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateProject")
    
    
    return
}

func NewUpdateProjectResponse() (response *UpdateProjectResponse) {
    response = &UpdateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateProject
// This API is used to update a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateProject(request *UpdateProjectRequest) (response *UpdateProjectResponse, err error) {
    return c.UpdateProjectWithContext(context.Background(), request)
}

// UpdateProject
// This API is used to update a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateProjectWithContext(ctx context.Context, request *UpdateProjectRequest) (response *UpdateProjectResponse, err error) {
    if request == nil {
        request = NewUpdateProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateScenarioRequest() (request *UpdateScenarioRequest) {
    request = &UpdateScenarioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateScenario")
    
    
    return
}

func NewUpdateScenarioResponse() (response *UpdateScenarioResponse) {
    response = &UpdateScenarioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateScenario
// This API is used to update a scenario.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateScenario(request *UpdateScenarioRequest) (response *UpdateScenarioResponse, err error) {
    return c.UpdateScenarioWithContext(context.Background(), request)
}

// UpdateScenario
// This API is used to update a scenario.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateScenarioWithContext(ctx context.Context, request *UpdateScenarioRequest) (response *UpdateScenarioResponse, err error) {
    if request == nil {
        request = NewUpdateScenarioRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateScenario require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateScenarioResponse()
    err = c.Send(request, response)
    return
}
