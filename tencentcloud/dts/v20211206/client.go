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

package v20211206

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-12-06"

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


func NewCompleteMigrateJobRequest() (request *CompleteMigrateJobRequest) {
    request = &CompleteMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CompleteMigrateJob")
    
    
    return
}

func NewCompleteMigrateJobResponse() (response *CompleteMigrateJobResponse) {
    response = &CompleteMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CompleteMigrateJob
// This API is used to complete a data migration task.
//
// For tasks in incremental migration mode, you need to call this API before migration gets ready for completion to stop migrating incremental data.
//
// If the task status queried through the `DescribeMigrationJobs` API is ready (`Status` = `readyComplete`), you can call this API to complete the migration task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CompleteMigrateJob(request *CompleteMigrateJobRequest) (response *CompleteMigrateJobResponse, err error) {
    return c.CompleteMigrateJobWithContext(context.Background(), request)
}

// CompleteMigrateJob
// This API is used to complete a data migration task.
//
// For tasks in incremental migration mode, you need to call this API before migration gets ready for completion to stop migrating incremental data.
//
// If the task status queried through the `DescribeMigrationJobs` API is ready (`Status` = `readyComplete`), you can call this API to complete the migration task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CompleteMigrateJobWithContext(ctx context.Context, request *CompleteMigrateJobRequest) (response *CompleteMigrateJobResponse, err error) {
    if request == nil {
        request = NewCompleteMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompleteMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompleteMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewConfigureSubscribeJobRequest() (request *ConfigureSubscribeJobRequest) {
    request = &ConfigureSubscribeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ConfigureSubscribeJob")
    
    
    return
}

func NewConfigureSubscribeJobResponse() (response *ConfigureSubscribeJobResponse) {
    response = &ConfigureSubscribeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfigureSubscribeJob
// This API is used to configure data subscription instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ConfigureSubscribeJob(request *ConfigureSubscribeJobRequest) (response *ConfigureSubscribeJobResponse, err error) {
    return c.ConfigureSubscribeJobWithContext(context.Background(), request)
}

// ConfigureSubscribeJob
// This API is used to configure data subscription instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ConfigureSubscribeJobWithContext(ctx context.Context, request *ConfigureSubscribeJobRequest) (response *ConfigureSubscribeJobResponse, err error) {
    if request == nil {
        request = NewConfigureSubscribeJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfigureSubscribeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfigureSubscribeJobResponse()
    err = c.Send(request, response)
    return
}

func NewConfigureSyncJobRequest() (request *ConfigureSyncJobRequest) {
    request = &ConfigureSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ConfigureSyncJob")
    
    
    return
}

func NewConfigureSyncJobResponse() (response *ConfigureSyncJobResponse) {
    response = &ConfigureSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfigureSyncJob
// This API is used to configure a sync task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ConfigureSyncJob(request *ConfigureSyncJobRequest) (response *ConfigureSyncJobResponse, err error) {
    return c.ConfigureSyncJobWithContext(context.Background(), request)
}

// ConfigureSyncJob
// This API is used to configure a sync task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ConfigureSyncJobWithContext(ctx context.Context, request *ConfigureSyncJobRequest) (response *ConfigureSyncJobResponse, err error) {
    if request == nil {
        request = NewConfigureSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfigureSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfigureSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewContinueMigrateJobRequest() (request *ContinueMigrateJobRequest) {
    request = &ContinueMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ContinueMigrateJob")
    
    
    return
}

func NewContinueMigrateJobResponse() (response *ContinueMigrateJobResponse) {
    response = &ContinueMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ContinueMigrateJob
// This API is used to resume a paused migration task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ContinueMigrateJob(request *ContinueMigrateJobRequest) (response *ContinueMigrateJobResponse, err error) {
    return c.ContinueMigrateJobWithContext(context.Background(), request)
}

// ContinueMigrateJob
// This API is used to resume a paused migration task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ContinueMigrateJobWithContext(ctx context.Context, request *ContinueMigrateJobRequest) (response *ContinueMigrateJobResponse, err error) {
    if request == nil {
        request = NewContinueMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ContinueMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewContinueMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewContinueSyncJobRequest() (request *ContinueSyncJobRequest) {
    request = &ContinueSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ContinueSyncJob")
    
    
    return
}

func NewContinueSyncJobResponse() (response *ContinueSyncJobResponse) {
    response = &ContinueSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ContinueSyncJob
// This API is used to resume a paused data sync task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ContinueSyncJob(request *ContinueSyncJobRequest) (response *ContinueSyncJobResponse, err error) {
    return c.ContinueSyncJobWithContext(context.Background(), request)
}

// ContinueSyncJob
// This API is used to resume a paused data sync task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ContinueSyncJobWithContext(ctx context.Context, request *ContinueSyncJobRequest) (response *ContinueSyncJobResponse, err error) {
    if request == nil {
        request = NewContinueSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ContinueSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewContinueSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCheckSyncJobRequest() (request *CreateCheckSyncJobRequest) {
    request = &CreateCheckSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateCheckSyncJob")
    
    
    return
}

func NewCreateCheckSyncJobResponse() (response *CreateCheckSyncJobResponse) {
    response = &CreateCheckSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCheckSyncJob
// This API is used to verify a sync task by checking required parameters and peripheral configuration.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDFORSYNCJOBERROR = "UnsupportedOperation.IntraNetUserNotTaggedForSyncJobError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateCheckSyncJob(request *CreateCheckSyncJobRequest) (response *CreateCheckSyncJobResponse, err error) {
    return c.CreateCheckSyncJobWithContext(context.Background(), request)
}

// CreateCheckSyncJob
// This API is used to verify a sync task by checking required parameters and peripheral configuration.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDFORSYNCJOBERROR = "UnsupportedOperation.IntraNetUserNotTaggedForSyncJobError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateCheckSyncJobWithContext(ctx context.Context, request *CreateCheckSyncJobRequest) (response *CreateCheckSyncJobResponse, err error) {
    if request == nil {
        request = NewCreateCheckSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCheckSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCheckSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCompareTaskRequest() (request *CreateCompareTaskRequest) {
    request = &CreateCompareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateCompareTask")
    
    
    return
}

func NewCreateCompareTaskResponse() (response *CreateCompareTaskResponse) {
    response = &CreateCompareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCompareTask
// This API is used to create a data consistency check task. After the task is successfully created, its ID will be returned in the format of `dts-8yv4w2i1-cmp-37skmii9`, and you can call `StartCompare` to start it.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateCompareTask(request *CreateCompareTaskRequest) (response *CreateCompareTaskResponse, err error) {
    return c.CreateCompareTaskWithContext(context.Background(), request)
}

// CreateCompareTask
// This API is used to create a data consistency check task. After the task is successfully created, its ID will be returned in the format of `dts-8yv4w2i1-cmp-37skmii9`, and you can call `StartCompare` to start it.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateCompareTaskWithContext(ctx context.Context, request *CreateCompareTaskRequest) (response *CreateCompareTaskResponse, err error) {
    if request == nil {
        request = NewCreateCompareTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCompareTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCompareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsumerGroupRequest() (request *CreateConsumerGroupRequest) {
    request = &CreateConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateConsumerGroup")
    
    
    return
}

func NewCreateConsumerGroupResponse() (response *CreateConsumerGroupResponse) {
    response = &CreateConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsumerGroup
// This API is used to creat a consumer group for the subscription instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    return c.CreateConsumerGroupWithContext(context.Background(), request)
}

// CreateConsumerGroup
// This API is used to creat a consumer group for the subscription instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateConsumerGroupWithContext(ctx context.Context, request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    if request == nil {
        request = NewCreateConsumerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrateCheckJobRequest() (request *CreateMigrateCheckJobRequest) {
    request = &CreateMigrateCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateCheckJob")
    
    
    return
}

func NewCreateMigrateCheckJobResponse() (response *CreateMigrateCheckJobResponse) {
    response = &CreateMigrateCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMigrateCheckJob
// This API is used to verify a migration task.
//
// Before migration, you should call this API to create a check task. Migration will start only if the check succeeds. You can view the check result through the `DescribeMigrationCheckJob` API.
//
// After successful check, if the migration task needs to be modified, a new check task should be created, and migration will start only after the new check succeeds.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_CGWSYSTEMERROR = "InternalError.CgwSystemError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateMigrateCheckJob(request *CreateMigrateCheckJobRequest) (response *CreateMigrateCheckJobResponse, err error) {
    return c.CreateMigrateCheckJobWithContext(context.Background(), request)
}

// CreateMigrateCheckJob
// This API is used to verify a migration task.
//
// Before migration, you should call this API to create a check task. Migration will start only if the check succeeds. You can view the check result through the `DescribeMigrationCheckJob` API.
//
// After successful check, if the migration task needs to be modified, a new check task should be created, and migration will start only after the new check succeeds.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_CGWSYSTEMERROR = "InternalError.CgwSystemError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateMigrateCheckJobWithContext(ctx context.Context, request *CreateMigrateCheckJobRequest) (response *CreateMigrateCheckJobResponse, err error) {
    if request == nil {
        request = NewCreateMigrateCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigrateCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMigrateCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrationServiceRequest() (request *CreateMigrationServiceRequest) {
    request = &CreateMigrationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrationService")
    
    
    return
}

func NewCreateMigrationServiceResponse() (response *CreateMigrationServiceResponse) {
    response = &CreateMigrationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMigrationService
// This API is used to purchase migration tasks. After the tasks are purchased successfully, a randomly generated list of task IDs will be returned. You can also call the `DescribeMigrationJobs` API to query the IDs of the successfully purchased tasks. Note that once a task is purchased successfully, the types and regions of the source and target databases cannot be changed.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateMigrationService(request *CreateMigrationServiceRequest) (response *CreateMigrationServiceResponse, err error) {
    return c.CreateMigrationServiceWithContext(context.Background(), request)
}

// CreateMigrationService
// This API is used to purchase migration tasks. After the tasks are purchased successfully, a randomly generated list of task IDs will be returned. You can also call the `DescribeMigrationJobs` API to query the IDs of the successfully purchased tasks. Note that once a task is purchased successfully, the types and regions of the source and target databases cannot be changed.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateMigrationServiceWithContext(ctx context.Context, request *CreateMigrationServiceRequest) (response *CreateMigrationServiceResponse, err error) {
    if request == nil {
        request = NewCreateMigrationServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigrationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMigrationServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModifyCheckSyncJobRequest() (request *CreateModifyCheckSyncJobRequest) {
    request = &CreateModifyCheckSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateModifyCheckSyncJob")
    
    
    return
}

func NewCreateModifyCheckSyncJobResponse() (response *CreateModifyCheckSyncJobResponse) {
    response = &CreateModifyCheckSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateModifyCheckSyncJob
// This API is used to check whether the current data sync task supports object modification after the task configuration is modified.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateModifyCheckSyncJob(request *CreateModifyCheckSyncJobRequest) (response *CreateModifyCheckSyncJobResponse, err error) {
    return c.CreateModifyCheckSyncJobWithContext(context.Background(), request)
}

// CreateModifyCheckSyncJob
// This API is used to check whether the current data sync task supports object modification after the task configuration is modified.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateModifyCheckSyncJobWithContext(ctx context.Context, request *CreateModifyCheckSyncJobRequest) (response *CreateModifyCheckSyncJobResponse, err error) {
    if request == nil {
        request = NewCreateModifyCheckSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModifyCheckSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModifyCheckSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubscribeRequest() (request *CreateSubscribeRequest) {
    request = &CreateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateSubscribe")
    
    
    return
}

func NewCreateSubscribeResponse() (response *CreateSubscribeResponse) {
    response = &CreateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSubscribe
// This API is used to create a data subscription task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_USERNOTAUTHED = "FailedOperation.UserNotAuthed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    return c.CreateSubscribeWithContext(context.Background(), request)
}

// CreateSubscribe
// This API is used to create a data subscription task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_USERNOTAUTHED = "FailedOperation.UserNotAuthed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateSubscribeWithContext(ctx context.Context, request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubscribeCheckJobRequest() (request *CreateSubscribeCheckJobRequest) {
    request = &CreateSubscribeCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateSubscribeCheckJob")
    
    
    return
}

func NewCreateSubscribeCheckJobResponse() (response *CreateSubscribeCheckJobResponse) {
    response = &CreateSubscribeCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSubscribeCheckJob
// This API is used to create a subscription check task. The task must have successfully called the ConfigureSubscribeJob interface to configure all necessary information before starting the check.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateSubscribeCheckJob(request *CreateSubscribeCheckJobRequest) (response *CreateSubscribeCheckJobResponse, err error) {
    return c.CreateSubscribeCheckJobWithContext(context.Background(), request)
}

// CreateSubscribeCheckJob
// This API is used to create a subscription check task. The task must have successfully called the ConfigureSubscribeJob interface to configure all necessary information before starting the check.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateSubscribeCheckJobWithContext(ctx context.Context, request *CreateSubscribeCheckJobRequest) (response *CreateSubscribeCheckJobResponse, err error) {
    if request == nil {
        request = NewCreateSubscribeCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubscribeCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSubscribeCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSyncJobRequest() (request *CreateSyncJobRequest) {
    request = &CreateSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateSyncJob")
    
    
    return
}

func NewCreateSyncJobResponse() (response *CreateSyncJobResponse) {
    response = &CreateSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSyncJob
// This API is used to create a sync task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_NOTENOUGHMONEYERROR = "InternalError.NotEnoughMoneyError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) CreateSyncJob(request *CreateSyncJobRequest) (response *CreateSyncJobResponse, err error) {
    return c.CreateSyncJobWithContext(context.Background(), request)
}

// CreateSyncJob
// This API is used to create a sync task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_NOTENOUGHMONEYERROR = "InternalError.NotEnoughMoneyError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) CreateSyncJobWithContext(ctx context.Context, request *CreateSyncJobRequest) (response *CreateSyncJobResponse, err error) {
    if request == nil {
        request = NewCreateSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCompareTaskRequest() (request *DeleteCompareTaskRequest) {
    request = &DeleteCompareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DeleteCompareTask")
    
    
    return
}

func NewDeleteCompareTaskResponse() (response *DeleteCompareTaskResponse) {
    response = &DeleteCompareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCompareTask
// This API is used to delete a data consistency check task, which can be called when the task status is `success`, `failed`, or `canceled`.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DeleteCompareTask(request *DeleteCompareTaskRequest) (response *DeleteCompareTaskResponse, err error) {
    return c.DeleteCompareTaskWithContext(context.Background(), request)
}

// DeleteCompareTask
// This API is used to delete a data consistency check task, which can be called when the task status is `success`, `failed`, or `canceled`.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DeleteCompareTaskWithContext(ctx context.Context, request *DeleteCompareTaskRequest) (response *DeleteCompareTaskResponse, err error) {
    if request == nil {
        request = NewDeleteCompareTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCompareTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCompareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsumerGroupRequest() (request *DeleteConsumerGroupRequest) {
    request = &DeleteConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DeleteConsumerGroup")
    
    
    return
}

func NewDeleteConsumerGroupResponse() (response *DeleteConsumerGroupResponse) {
    response = &DeleteConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsumerGroup
// This API is used to delete a consumer group of a subscription task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    return c.DeleteConsumerGroupWithContext(context.Background(), request)
}

// DeleteConsumerGroup
// This API is used to delete a consumer group of a subscription task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DeleteConsumerGroupWithContext(ctx context.Context, request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteConsumerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckSyncJobResultRequest() (request *DescribeCheckSyncJobResultRequest) {
    request = &DescribeCheckSyncJobResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeCheckSyncJobResult")
    
    
    return
}

func NewDescribeCheckSyncJobResultResponse() (response *DescribeCheckSyncJobResultResponse) {
    response = &DescribeCheckSyncJobResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCheckSyncJobResult
// This API is used to query the result of the sync check task and check the required parameters and peripheral configurations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) DescribeCheckSyncJobResult(request *DescribeCheckSyncJobResultRequest) (response *DescribeCheckSyncJobResultResponse, err error) {
    return c.DescribeCheckSyncJobResultWithContext(context.Background(), request)
}

// DescribeCheckSyncJobResult
// This API is used to query the result of the sync check task and check the required parameters and peripheral configurations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) DescribeCheckSyncJobResultWithContext(ctx context.Context, request *DescribeCheckSyncJobResultRequest) (response *DescribeCheckSyncJobResultResponse, err error) {
    if request == nil {
        request = NewDescribeCheckSyncJobResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckSyncJobResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckSyncJobResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompareReportRequest() (request *DescribeCompareReportRequest) {
    request = &DescribeCompareReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeCompareReport")
    
    
    return
}

func NewDescribeCompareReportResponse() (response *DescribeCompareReportResponse) {
    response = &DescribeCompareReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCompareReport
// This API is used to query the details of a data consistency check task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeCompareReport(request *DescribeCompareReportRequest) (response *DescribeCompareReportResponse, err error) {
    return c.DescribeCompareReportWithContext(context.Background(), request)
}

// DescribeCompareReport
// This API is used to query the details of a data consistency check task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeCompareReportWithContext(ctx context.Context, request *DescribeCompareReportRequest) (response *DescribeCompareReportResponse, err error) {
    if request == nil {
        request = NewDescribeCompareReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompareReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompareReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompareTasksRequest() (request *DescribeCompareTasksRequest) {
    request = &DescribeCompareTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeCompareTasks")
    
    
    return
}

func NewDescribeCompareTasksResponse() (response *DescribeCompareTasksResponse) {
    response = &DescribeCompareTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCompareTasks
// This API is used to query the list of data consistency check tasks under the current task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeCompareTasks(request *DescribeCompareTasksRequest) (response *DescribeCompareTasksResponse, err error) {
    return c.DescribeCompareTasksWithContext(context.Background(), request)
}

// DescribeCompareTasks
// This API is used to query the list of data consistency check tasks under the current task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeCompareTasksWithContext(ctx context.Context, request *DescribeCompareTasksRequest) (response *DescribeCompareTasksResponse, err error) {
    if request == nil {
        request = NewDescribeCompareTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompareTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompareTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerGroupsRequest() (request *DescribeConsumerGroupsRequest) {
    request = &DescribeConsumerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeConsumerGroups")
    
    
    return
}

func NewDescribeConsumerGroupsResponse() (response *DescribeConsumerGroupsResponse) {
    response = &DescribeConsumerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerGroups
// This API is used to get consumer group details for the subscription instance configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeConsumerGroups(request *DescribeConsumerGroupsRequest) (response *DescribeConsumerGroupsResponse, err error) {
    return c.DescribeConsumerGroupsWithContext(context.Background(), request)
}

// DescribeConsumerGroups
// This API is used to get consumer group details for the subscription instance configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeConsumerGroupsWithContext(ctx context.Context, request *DescribeConsumerGroupsRequest) (response *DescribeConsumerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateDBInstancesRequest() (request *DescribeMigrateDBInstancesRequest) {
    request = &DescribeMigrateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateDBInstances")
    
    
    return
}

func NewDescribeMigrateDBInstancesResponse() (response *DescribeMigrateDBInstancesResponse) {
    response = &DescribeMigrateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrateDBInstances
// This API is used to query migratable database instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrateDBInstances(request *DescribeMigrateDBInstancesRequest) (response *DescribeMigrateDBInstancesResponse, err error) {
    return c.DescribeMigrateDBInstancesWithContext(context.Background(), request)
}

// DescribeMigrateDBInstances
// This API is used to query migratable database instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrateDBInstancesWithContext(ctx context.Context, request *DescribeMigrateDBInstancesRequest) (response *DescribeMigrateDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationCheckJobRequest() (request *DescribeMigrationCheckJobRequest) {
    request = &DescribeMigrationCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrationCheckJob")
    
    
    return
}

func NewDescribeMigrationCheckJobResponse() (response *DescribeMigrationCheckJobResponse) {
    response = &DescribeMigrationCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrationCheckJob
// This API is used to get the check result and query the check status and progress after a check is created. 
//
// If the check succeeds, you can call the `StartMigrateJob` API to start migration.
//
// If the check fails, the cause can be queried. Modify the migration configuration or adjust relevant parameters of the source/target instances through the `ModifyMigrationJob` API based on the error message.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationCheckJob(request *DescribeMigrationCheckJobRequest) (response *DescribeMigrationCheckJobResponse, err error) {
    return c.DescribeMigrationCheckJobWithContext(context.Background(), request)
}

// DescribeMigrationCheckJob
// This API is used to get the check result and query the check status and progress after a check is created. 
//
// If the check succeeds, you can call the `StartMigrateJob` API to start migration.
//
// If the check fails, the cause can be queried. Modify the migration configuration or adjust relevant parameters of the source/target instances through the `ModifyMigrationJob` API based on the error message.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationCheckJobWithContext(ctx context.Context, request *DescribeMigrationCheckJobRequest) (response *DescribeMigrationCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationDetailRequest() (request *DescribeMigrationDetailRequest) {
    request = &DescribeMigrationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrationDetail")
    
    
    return
}

func NewDescribeMigrationDetailResponse() (response *DescribeMigrationDetailResponse) {
    response = &DescribeMigrationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrationDetail
// This API is used to query the details of a migration task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationDetail(request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    return c.DescribeMigrationDetailWithContext(context.Background(), request)
}

// DescribeMigrationDetail
// This API is used to query the details of a migration task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationDetailWithContext(ctx context.Context, request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationJobsRequest() (request *DescribeMigrationJobsRequest) {
    request = &DescribeMigrationJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrationJobs")
    
    
    return
}

func NewDescribeMigrationJobsResponse() (response *DescribeMigrationJobsResponse) {
    response = &DescribeMigrationJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrationJobs
// This API is used to query the list of data migration tasks.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationJobs(request *DescribeMigrationJobsRequest) (response *DescribeMigrationJobsResponse, err error) {
    return c.DescribeMigrationJobsWithContext(context.Background(), request)
}

// DescribeMigrationJobs
// This API is used to query the list of data migration tasks.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationJobsWithContext(ctx context.Context, request *DescribeMigrationJobsRequest) (response *DescribeMigrationJobsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModifyCheckSyncJobResultRequest() (request *DescribeModifyCheckSyncJobResultRequest) {
    request = &DescribeModifyCheckSyncJobResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeModifyCheckSyncJobResult")
    
    
    return
}

func NewDescribeModifyCheckSyncJobResultResponse() (response *DescribeModifyCheckSyncJobResultResponse) {
    response = &DescribeModifyCheckSyncJobResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModifyCheckSyncJobResult
// This API is used to query the result of the created check task for object modification.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) DescribeModifyCheckSyncJobResult(request *DescribeModifyCheckSyncJobResultRequest) (response *DescribeModifyCheckSyncJobResultResponse, err error) {
    return c.DescribeModifyCheckSyncJobResultWithContext(context.Background(), request)
}

// DescribeModifyCheckSyncJobResult
// This API is used to query the result of the created check task for object modification.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) DescribeModifyCheckSyncJobResultWithContext(ctx context.Context, request *DescribeModifyCheckSyncJobResultRequest) (response *DescribeModifyCheckSyncJobResultResponse, err error) {
    if request == nil {
        request = NewDescribeModifyCheckSyncJobResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModifyCheckSyncJobResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModifyCheckSyncJobResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOffsetByTimeRequest() (request *DescribeOffsetByTimeRequest) {
    request = &DescribeOffsetByTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeOffsetByTime")
    
    
    return
}

func NewDescribeOffsetByTimeResponse() (response *DescribeOffsetByTimeResponse) {
    response = &DescribeOffsetByTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOffsetByTime
// This API is used to query the latest offset before the specified time in KafkaTopic.The offset output by the interface is the closest offset to this time.If the input time is much later than the current time, the output is equivalent to the latest offset;If the input time is much earlier than the current time, the output is equivalent to the oldest offset;If the input is empty, the default time is 0, which is the oldest offset to be queried.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeOffsetByTime(request *DescribeOffsetByTimeRequest) (response *DescribeOffsetByTimeResponse, err error) {
    return c.DescribeOffsetByTimeWithContext(context.Background(), request)
}

// DescribeOffsetByTime
// This API is used to query the latest offset before the specified time in KafkaTopic.The offset output by the interface is the closest offset to this time.If the input time is much later than the current time, the output is equivalent to the latest offset;If the input time is much earlier than the current time, the output is equivalent to the oldest offset;If the input is empty, the default time is 0, which is the oldest offset to be queried.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeOffsetByTimeWithContext(ctx context.Context, request *DescribeOffsetByTimeRequest) (response *DescribeOffsetByTimeResponse, err error) {
    if request == nil {
        request = NewDescribeOffsetByTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOffsetByTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOffsetByTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeCheckJobRequest() (request *DescribeSubscribeCheckJobRequest) {
    request = &DescribeSubscribeCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeCheckJob")
    
    
    return
}

func NewDescribeSubscribeCheckJobResponse() (response *DescribeSubscribeCheckJobResponse) {
    response = &DescribeSubscribeCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubscribeCheckJob
// This API is used to query the results of a subscription check task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeCheckJob(request *DescribeSubscribeCheckJobRequest) (response *DescribeSubscribeCheckJobResponse, err error) {
    return c.DescribeSubscribeCheckJobWithContext(context.Background(), request)
}

// DescribeSubscribeCheckJob
// This API is used to query the results of a subscription check task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeCheckJobWithContext(ctx context.Context, request *DescribeSubscribeCheckJobRequest) (response *DescribeSubscribeCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeDetailRequest() (request *DescribeSubscribeDetailRequest) {
    request = &DescribeSubscribeDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeDetail")
    
    
    return
}

func NewDescribeSubscribeDetailResponse() (response *DescribeSubscribeDetailResponse) {
    response = &DescribeSubscribeDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubscribeDetail
// This API is used to get the configuration information of the data subscription instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeDetail(request *DescribeSubscribeDetailRequest) (response *DescribeSubscribeDetailResponse, err error) {
    return c.DescribeSubscribeDetailWithContext(context.Background(), request)
}

// DescribeSubscribeDetail
// This API is used to get the configuration information of the data subscription instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeDetailWithContext(ctx context.Context, request *DescribeSubscribeDetailRequest) (response *DescribeSubscribeDetailResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeJobsRequest() (request *DescribeSubscribeJobsRequest) {
    request = &DescribeSubscribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeJobs")
    
    
    return
}

func NewDescribeSubscribeJobsResponse() (response *DescribeSubscribeJobsResponse) {
    response = &DescribeSubscribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubscribeJobs
// This API is used to get the information list of data subscription instances. Pagination is enabled by default with 20 results returned each time.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeJobs(request *DescribeSubscribeJobsRequest) (response *DescribeSubscribeJobsResponse, err error) {
    return c.DescribeSubscribeJobsWithContext(context.Background(), request)
}

// DescribeSubscribeJobs
// This API is used to get the information list of data subscription instances. Pagination is enabled by default with 20 results returned each time.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeJobsWithContext(ctx context.Context, request *DescribeSubscribeJobsRequest) (response *DescribeSubscribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeReturnableRequest() (request *DescribeSubscribeReturnableRequest) {
    request = &DescribeSubscribeReturnableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeReturnable")
    
    
    return
}

func NewDescribeSubscribeReturnableResponse() (response *DescribeSubscribeReturnableResponse) {
    response = &DescribeSubscribeReturnableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubscribeReturnable
// This API is used to query whether a subscription task can be terminated and returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeReturnable(request *DescribeSubscribeReturnableRequest) (response *DescribeSubscribeReturnableResponse, err error) {
    return c.DescribeSubscribeReturnableWithContext(context.Background(), request)
}

// DescribeSubscribeReturnable
// This API is used to query whether a subscription task can be terminated and returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeReturnableWithContext(ctx context.Context, request *DescribeSubscribeReturnableRequest) (response *DescribeSubscribeReturnableResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeReturnableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeReturnable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeReturnableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSyncJobsRequest() (request *DescribeSyncJobsRequest) {
    request = &DescribeSyncJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSyncJobs")
    
    
    return
}

func NewDescribeSyncJobsResponse() (response *DescribeSyncJobsResponse) {
    response = &DescribeSyncJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSyncJobs
// This API is used to query the information of a sync task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeSyncJobs(request *DescribeSyncJobsRequest) (response *DescribeSyncJobsResponse, err error) {
    return c.DescribeSyncJobsWithContext(context.Background(), request)
}

// DescribeSyncJobs
// This API is used to query the information of a sync task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeSyncJobsWithContext(ctx context.Context, request *DescribeSyncJobsRequest) (response *DescribeSyncJobsResponse, err error) {
    if request == nil {
        request = NewDescribeSyncJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSyncJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSyncJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyIsolatedSubscribeRequest() (request *DestroyIsolatedSubscribeRequest) {
    request = &DestroyIsolatedSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DestroyIsolatedSubscribe")
    
    
    return
}

func NewDestroyIsolatedSubscribeResponse() (response *DestroyIsolatedSubscribeResponse) {
    response = &DestroyIsolatedSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyIsolatedSubscribe
// This API is used to deactivate an isolated data subscription instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DestroyIsolatedSubscribe(request *DestroyIsolatedSubscribeRequest) (response *DestroyIsolatedSubscribeResponse, err error) {
    return c.DestroyIsolatedSubscribeWithContext(context.Background(), request)
}

// DestroyIsolatedSubscribe
// This API is used to deactivate an isolated data subscription instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DestroyIsolatedSubscribeWithContext(ctx context.Context, request *DestroyIsolatedSubscribeRequest) (response *DestroyIsolatedSubscribeResponse, err error) {
    if request == nil {
        request = NewDestroyIsolatedSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyIsolatedSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyIsolatedSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyMigrateJobRequest() (request *DestroyMigrateJobRequest) {
    request = &DestroyMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DestroyMigrateJob")
    
    
    return
}

func NewDestroyMigrateJobResponse() (response *DestroyMigrateJobResponse) {
    response = &DestroyMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyMigrateJob
// This API is used to delete a data migration task. For a billed task, you must first call the `IsolateMigrateJob` API to isolate it and make sure that it is in **Isolated** status before calling this API to delete it. For a free task, you can directly call the `IsolateMigrateJob` API to delete it.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DestroyMigrateJob(request *DestroyMigrateJobRequest) (response *DestroyMigrateJobResponse, err error) {
    return c.DestroyMigrateJobWithContext(context.Background(), request)
}

// DestroyMigrateJob
// This API is used to delete a data migration task. For a billed task, you must first call the `IsolateMigrateJob` API to isolate it and make sure that it is in **Isolated** status before calling this API to delete it. For a free task, you can directly call the `IsolateMigrateJob` API to delete it.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DestroyMigrateJobWithContext(ctx context.Context, request *DestroyMigrateJobRequest) (response *DestroyMigrateJobResponse, err error) {
    if request == nil {
        request = NewDestroyMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewDestroySyncJobRequest() (request *DestroySyncJobRequest) {
    request = &DestroySyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DestroySyncJob")
    
    
    return
}

func NewDestroySyncJobResponse() (response *DestroySyncJobResponse) {
    response = &DestroySyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroySyncJob
// This API is used to delete a sync task. Only tasks in **Isolated** status can be completely deleted. After deletion, you can call the `DescribeSyncJobs` API to get the task list. If the deleted task is not in the list, it is deleted successfully.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DestroySyncJob(request *DestroySyncJobRequest) (response *DestroySyncJobResponse, err error) {
    return c.DestroySyncJobWithContext(context.Background(), request)
}

// DestroySyncJob
// This API is used to delete a sync task. Only tasks in **Isolated** status can be completely deleted. After deletion, you can call the `DescribeSyncJobs` API to get the task list. If the deleted task is not in the list, it is deleted successfully.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DestroySyncJobWithContext(ctx context.Context, request *DestroySyncJobRequest) (response *DestroySyncJobResponse, err error) {
    if request == nil {
        request = NewDestroySyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroySyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroySyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateMigrateJobRequest() (request *IsolateMigrateJobRequest) {
    request = &IsolateMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "IsolateMigrateJob")
    
    
    return
}

func NewIsolateMigrateJobResponse() (response *IsolateMigrateJobResponse) {
    response = &IsolateMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateMigrateJob
//  This API is used to isolate and return a data migration task. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status. For a billed task, after isolating it, you can call `RecoverMigrationJob` to recover it or call `DestroyMigrateJob` to delete it. For a free task, calling this API will directly delete it permanently.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateMigrateJob(request *IsolateMigrateJobRequest) (response *IsolateMigrateJobResponse, err error) {
    return c.IsolateMigrateJobWithContext(context.Background(), request)
}

// IsolateMigrateJob
//  This API is used to isolate and return a data migration task. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status. For a billed task, after isolating it, you can call `RecoverMigrationJob` to recover it or call `DestroyMigrateJob` to delete it. For a free task, calling this API will directly delete it permanently.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateMigrateJobWithContext(ctx context.Context, request *IsolateMigrateJobRequest) (response *IsolateMigrateJobResponse, err error) {
    if request == nil {
        request = NewIsolateMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateSubscribeRequest() (request *IsolateSubscribeRequest) {
    request = &IsolateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "IsolateSubscribe")
    
    
    return
}

func NewIsolateSubscribeResponse() (response *IsolateSubscribeResponse) {
    response = &IsolateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateSubscribe
// This API is used to isolate the subscription task. After calling, the subscription task will not be available. Pay-as-you-go tasks will stop billing, and monthly subscription tasks will refund automatically.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateSubscribe(request *IsolateSubscribeRequest) (response *IsolateSubscribeResponse, err error) {
    return c.IsolateSubscribeWithContext(context.Background(), request)
}

// IsolateSubscribe
// This API is used to isolate the subscription task. After calling, the subscription task will not be available. Pay-as-you-go tasks will stop billing, and monthly subscription tasks will refund automatically.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateSubscribeWithContext(ctx context.Context, request *IsolateSubscribeRequest) (response *IsolateSubscribeResponse, err error) {
    if request == nil {
        request = NewIsolateSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateSyncJobRequest() (request *IsolateSyncJobRequest) {
    request = &IsolateSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "IsolateSyncJob")
    
    
    return
}

func NewIsolateSyncJobResponse() (response *IsolateSyncJobResponse) {
    response = &IsolateSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateSyncJob
// This API is used to isolate a sync task. After the task is isolated, you can call the `DescribeSyncJobs` API to query its status, call `RecoverSyncJob` to recover it, or directly delete it. For a free task, calling this API will directly delete it permanently.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateSyncJob(request *IsolateSyncJobRequest) (response *IsolateSyncJobResponse, err error) {
    return c.IsolateSyncJobWithContext(context.Background(), request)
}

// IsolateSyncJob
// This API is used to isolate a sync task. After the task is isolated, you can call the `DescribeSyncJobs` API to query its status, call `RecoverSyncJob` to recover it, or directly delete it. For a free task, calling this API will directly delete it permanently.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateSyncJobWithContext(ctx context.Context, request *IsolateSyncJobRequest) (response *IsolateSyncJobResponse, err error) {
    if request == nil {
        request = NewIsolateSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCompareTaskRequest() (request *ModifyCompareTaskRequest) {
    request = &ModifyCompareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyCompareTask")
    
    
    return
}

func NewModifyCompareTaskResponse() (response *ModifyCompareTaskResponse) {
    response = &ModifyCompareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCompareTask
// This API is used to modify the parameters of a data consistency check task after it is created and before it starts.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyCompareTask(request *ModifyCompareTaskRequest) (response *ModifyCompareTaskResponse, err error) {
    return c.ModifyCompareTaskWithContext(context.Background(), request)
}

// ModifyCompareTask
// This API is used to modify the parameters of a data consistency check task after it is created and before it starts.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyCompareTaskWithContext(ctx context.Context, request *ModifyCompareTaskRequest) (response *ModifyCompareTaskResponse, err error) {
    if request == nil {
        request = NewModifyCompareTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCompareTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCompareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCompareTaskNameRequest() (request *ModifyCompareTaskNameRequest) {
    request = &ModifyCompareTaskNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyCompareTaskName")
    
    
    return
}

func NewModifyCompareTaskNameResponse() (response *ModifyCompareTaskNameResponse) {
    response = &ModifyCompareTaskNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCompareTaskName
// This API is used to rename a data consistency check task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyCompareTaskName(request *ModifyCompareTaskNameRequest) (response *ModifyCompareTaskNameResponse, err error) {
    return c.ModifyCompareTaskNameWithContext(context.Background(), request)
}

// ModifyCompareTaskName
// This API is used to rename a data consistency check task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyCompareTaskNameWithContext(ctx context.Context, request *ModifyCompareTaskNameRequest) (response *ModifyCompareTaskNameResponse, err error) {
    if request == nil {
        request = NewModifyCompareTaskNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCompareTaskName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCompareTaskNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsumerGroupDescriptionRequest() (request *ModifyConsumerGroupDescriptionRequest) {
    request = &ModifyConsumerGroupDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyConsumerGroupDescription")
    
    
    return
}

func NewModifyConsumerGroupDescriptionResponse() (response *ModifyConsumerGroupDescriptionResponse) {
    response = &ModifyConsumerGroupDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsumerGroupDescription
// This API is used to modify the description of the specified subscription consumption group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyConsumerGroupDescription(request *ModifyConsumerGroupDescriptionRequest) (response *ModifyConsumerGroupDescriptionResponse, err error) {
    return c.ModifyConsumerGroupDescriptionWithContext(context.Background(), request)
}

// ModifyConsumerGroupDescription
// This API is used to modify the description of the specified subscription consumption group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyConsumerGroupDescriptionWithContext(ctx context.Context, request *ModifyConsumerGroupDescriptionRequest) (response *ModifyConsumerGroupDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyConsumerGroupDescriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumerGroupDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerGroupDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsumerGroupPasswordRequest() (request *ModifyConsumerGroupPasswordRequest) {
    request = &ModifyConsumerGroupPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyConsumerGroupPassword")
    
    
    return
}

func NewModifyConsumerGroupPasswordResponse() (response *ModifyConsumerGroupPasswordResponse) {
    response = &ModifyConsumerGroupPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsumerGroupPassword
// This API is used to modify the password of the specified subscription consumer group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyConsumerGroupPassword(request *ModifyConsumerGroupPasswordRequest) (response *ModifyConsumerGroupPasswordResponse, err error) {
    return c.ModifyConsumerGroupPasswordWithContext(context.Background(), request)
}

// ModifyConsumerGroupPassword
// This API is used to modify the password of the specified subscription consumer group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyConsumerGroupPasswordWithContext(ctx context.Context, request *ModifyConsumerGroupPasswordRequest) (response *ModifyConsumerGroupPasswordResponse, err error) {
    if request == nil {
        request = NewModifyConsumerGroupPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumerGroupPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerGroupPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateJobSpecRequest() (request *ModifyMigrateJobSpecRequest) {
    request = &ModifyMigrateJobSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateJobSpec")
    
    
    return
}

func NewModifyMigrateJobSpecResponse() (response *ModifyMigrateJobSpecResponse) {
    response = &ModifyMigrateJobSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrateJobSpec
// This API is used to adjust the specification of a pay-as-you-go task. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateJobSpec(request *ModifyMigrateJobSpecRequest) (response *ModifyMigrateJobSpecResponse, err error) {
    return c.ModifyMigrateJobSpecWithContext(context.Background(), request)
}

// ModifyMigrateJobSpec
// This API is used to adjust the specification of a pay-as-you-go task. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateJobSpecWithContext(ctx context.Context, request *ModifyMigrateJobSpecRequest) (response *ModifyMigrateJobSpecResponse, err error) {
    if request == nil {
        request = NewModifyMigrateJobSpecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateJobSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateJobSpecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateNameRequest() (request *ModifyMigrateNameRequest) {
    request = &ModifyMigrateNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateName")
    
    
    return
}

func NewModifyMigrateNameResponse() (response *ModifyMigrateNameResponse) {
    response = &ModifyMigrateNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrateName
// This API is used to rename a migration task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateName(request *ModifyMigrateNameRequest) (response *ModifyMigrateNameResponse, err error) {
    return c.ModifyMigrateNameWithContext(context.Background(), request)
}

// ModifyMigrateName
// This API is used to rename a migration task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateNameWithContext(ctx context.Context, request *ModifyMigrateNameRequest) (response *ModifyMigrateNameResponse, err error) {
    if request == nil {
        request = NewModifyMigrateNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateRateLimitRequest() (request *ModifyMigrateRateLimitRequest) {
    request = &ModifyMigrateRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateRateLimit")
    
    
    return
}

func NewModifyMigrateRateLimitResponse() (response *ModifyMigrateRateLimitResponse) {
    response = &ModifyMigrateRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrateRateLimit
// This API is used to restrict the rate limit of the task, when a user finds that migration task has a large impact on the load of user's database.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateRateLimit(request *ModifyMigrateRateLimitRequest) (response *ModifyMigrateRateLimitResponse, err error) {
    return c.ModifyMigrateRateLimitWithContext(context.Background(), request)
}

// ModifyMigrateRateLimit
// This API is used to restrict the rate limit of the task, when a user finds that migration task has a large impact on the load of user's database.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateRateLimitWithContext(ctx context.Context, request *ModifyMigrateRateLimitRequest) (response *ModifyMigrateRateLimitResponse, err error) {
    if request == nil {
        request = NewModifyMigrateRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateRuntimeAttributeRequest() (request *ModifyMigrateRuntimeAttributeRequest) {
    request = &ModifyMigrateRuntimeAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateRuntimeAttribute")
    
    
    return
}

func NewModifyMigrateRuntimeAttributeResponse() (response *ModifyMigrateRuntimeAttributeResponse) {
    response = &ModifyMigrateRuntimeAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrateRuntimeAttribute
// This API is used to modify task runtime attributes. This interface is different from the configuration class interface and does not perform state machine evaluation.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
func (c *Client) ModifyMigrateRuntimeAttribute(request *ModifyMigrateRuntimeAttributeRequest) (response *ModifyMigrateRuntimeAttributeResponse, err error) {
    return c.ModifyMigrateRuntimeAttributeWithContext(context.Background(), request)
}

// ModifyMigrateRuntimeAttribute
// This API is used to modify task runtime attributes. This interface is different from the configuration class interface and does not perform state machine evaluation.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
func (c *Client) ModifyMigrateRuntimeAttributeWithContext(ctx context.Context, request *ModifyMigrateRuntimeAttributeRequest) (response *ModifyMigrateRuntimeAttributeResponse, err error) {
    if request == nil {
        request = NewModifyMigrateRuntimeAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateRuntimeAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateRuntimeAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrationJobRequest() (request *ModifyMigrationJobRequest) {
    request = &ModifyMigrationJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrationJob")
    
    
    return
}

func NewModifyMigrationJobResponse() (response *ModifyMigrationJobResponse) {
    response = &ModifyMigrationJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrationJob
// This API is used to configure a migration task. After it is configured successfully, you can call the `CreateMigrationCheckJob` API to create a check task, and only after it passes the check can it be started.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_AUTHORIZEDOPERATIONDENYERROR = "AuthFailure.AuthorizedOperationDenyError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrationJob(request *ModifyMigrationJobRequest) (response *ModifyMigrationJobResponse, err error) {
    return c.ModifyMigrationJobWithContext(context.Background(), request)
}

// ModifyMigrationJob
// This API is used to configure a migration task. After it is configured successfully, you can call the `CreateMigrationCheckJob` API to create a check task, and only after it passes the check can it be started.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_AUTHORIZEDOPERATIONDENYERROR = "AuthFailure.AuthorizedOperationDenyError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrationJobWithContext(ctx context.Context, request *ModifyMigrationJobRequest) (response *ModifyMigrationJobResponse, err error) {
    if request == nil {
        request = NewModifyMigrationJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrationJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrationJobResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeAutoRenewFlagRequest() (request *ModifySubscribeAutoRenewFlagRequest) {
    request = &ModifySubscribeAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeAutoRenewFlag")
    
    
    return
}

func NewModifySubscribeAutoRenewFlagResponse() (response *ModifySubscribeAutoRenewFlagResponse) {
    response = &ModifySubscribeAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySubscribeAutoRenewFlag
// This API is used to modify the auto-renewal flag of your subscription instance. Only the monthly subscription modification task makes sense. After modifying, the pay-as-you-go task has no impact.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifySubscribeAutoRenewFlag(request *ModifySubscribeAutoRenewFlagRequest) (response *ModifySubscribeAutoRenewFlagResponse, err error) {
    return c.ModifySubscribeAutoRenewFlagWithContext(context.Background(), request)
}

// ModifySubscribeAutoRenewFlag
// This API is used to modify the auto-renewal flag of your subscription instance. Only the monthly subscription modification task makes sense. After modifying, the pay-as-you-go task has no impact.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifySubscribeAutoRenewFlagWithContext(ctx context.Context, request *ModifySubscribeAutoRenewFlagRequest) (response *ModifySubscribeAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifySubscribeAutoRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeNameRequest() (request *ModifySubscribeNameRequest) {
    request = &ModifySubscribeNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeName")
    
    
    return
}

func NewModifySubscribeNameResponse() (response *ModifySubscribeNameResponse) {
    response = &ModifySubscribeNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySubscribeName
// This API is used to modify the name of the data subscription instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifySubscribeName(request *ModifySubscribeNameRequest) (response *ModifySubscribeNameResponse, err error) {
    return c.ModifySubscribeNameWithContext(context.Background(), request)
}

// ModifySubscribeName
// This API is used to modify the name of the data subscription instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifySubscribeNameWithContext(ctx context.Context, request *ModifySubscribeNameRequest) (response *ModifySubscribeNameResponse, err error) {
    if request == nil {
        request = NewModifySubscribeNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeObjectsRequest() (request *ModifySubscribeObjectsRequest) {
    request = &ModifySubscribeObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeObjects")
    
    
    return
}

func NewModifySubscribeObjectsResponse() (response *ModifySubscribeObjectsResponse) {
    response = &ModifySubscribeObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySubscribeObjects
// This API is used to modify the data subscription object and Kafka partition rule. For MongoDB subscription, you can also modify the output aggregation rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifySubscribeObjects(request *ModifySubscribeObjectsRequest) (response *ModifySubscribeObjectsResponse, err error) {
    return c.ModifySubscribeObjectsWithContext(context.Background(), request)
}

// ModifySubscribeObjects
// This API is used to modify the data subscription object and Kafka partition rule. For MongoDB subscription, you can also modify the output aggregation rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifySubscribeObjectsWithContext(ctx context.Context, request *ModifySubscribeObjectsRequest) (response *ModifySubscribeObjectsResponse, err error) {
    if request == nil {
        request = NewModifySubscribeObjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySyncJobConfigRequest() (request *ModifySyncJobConfigRequest) {
    request = &ModifySyncJobConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySyncJobConfig")
    
    
    return
}

func NewModifySyncJobConfigResponse() (response *ModifySyncJobConfigResponse) {
    response = &ModifySyncJobConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySyncJobConfig
// This API is used to modify the configuration of a data sync task after the task is started.\n Configuration modification steps:  Modify sync task configuration -> Create a modification check task -> Query the check task result -> Start the configuration modification check task
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) ModifySyncJobConfig(request *ModifySyncJobConfigRequest) (response *ModifySyncJobConfigResponse, err error) {
    return c.ModifySyncJobConfigWithContext(context.Background(), request)
}

// ModifySyncJobConfig
// This API is used to modify the configuration of a data sync task after the task is started.\n Configuration modification steps:  Modify sync task configuration -> Create a modification check task -> Query the check task result -> Start the configuration modification check task
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) ModifySyncJobConfigWithContext(ctx context.Context, request *ModifySyncJobConfigRequest) (response *ModifySyncJobConfigResponse, err error) {
    if request == nil {
        request = NewModifySyncJobConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySyncJobConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySyncJobConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifySyncRateLimitRequest() (request *ModifySyncRateLimitRequest) {
    request = &ModifySyncRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySyncRateLimit")
    
    
    return
}

func NewModifySyncRateLimitResponse() (response *ModifySyncRateLimitResponse) {
    response = &ModifySyncRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySyncRateLimit
// This API is used to restrict the rate limit of the task, when a user finds that the sync task has a large impact on the load of user's database.
//
// error code that may be returned:
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ModifySyncRateLimit(request *ModifySyncRateLimitRequest) (response *ModifySyncRateLimitResponse, err error) {
    return c.ModifySyncRateLimitWithContext(context.Background(), request)
}

// ModifySyncRateLimit
// This API is used to restrict the rate limit of the task, when a user finds that the sync task has a large impact on the load of user's database.
//
// error code that may be returned:
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ModifySyncRateLimitWithContext(ctx context.Context, request *ModifySyncRateLimitRequest) (response *ModifySyncRateLimitResponse, err error) {
    if request == nil {
        request = NewModifySyncRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySyncRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySyncRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewPauseMigrateJobRequest() (request *PauseMigrateJobRequest) {
    request = &PauseMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "PauseMigrateJob")
    
    
    return
}

func NewPauseMigrateJobResponse() (response *PauseMigrateJobResponse) {
    response = &PauseMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseMigrateJob
// This API is used to pause a migration task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) PauseMigrateJob(request *PauseMigrateJobRequest) (response *PauseMigrateJobResponse, err error) {
    return c.PauseMigrateJobWithContext(context.Background(), request)
}

// PauseMigrateJob
// This API is used to pause a migration task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) PauseMigrateJobWithContext(ctx context.Context, request *PauseMigrateJobRequest) (response *PauseMigrateJobResponse, err error) {
    if request == nil {
        request = NewPauseMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewPauseSyncJobRequest() (request *PauseSyncJobRequest) {
    request = &PauseSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "PauseSyncJob")
    
    
    return
}

func NewPauseSyncJobResponse() (response *PauseSyncJobResponse) {
    response = &PauseSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseSyncJob
// This API is used to pause a data sync task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) PauseSyncJob(request *PauseSyncJobRequest) (response *PauseSyncJobResponse, err error) {
    return c.PauseSyncJobWithContext(context.Background(), request)
}

// PauseSyncJob
// This API is used to pause a data sync task.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) PauseSyncJobWithContext(ctx context.Context, request *PauseSyncJobRequest) (response *PauseSyncJobResponse, err error) {
    if request == nil {
        request = NewPauseSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverMigrateJobRequest() (request *RecoverMigrateJobRequest) {
    request = &RecoverMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "RecoverMigrateJob")
    
    
    return
}

func NewRecoverMigrateJobResponse() (response *RecoverMigrateJobResponse) {
    response = &RecoverMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecoverMigrateJob
// This API is used to recover a data migration task in **Isolated** status. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) RecoverMigrateJob(request *RecoverMigrateJobRequest) (response *RecoverMigrateJobResponse, err error) {
    return c.RecoverMigrateJobWithContext(context.Background(), request)
}

// RecoverMigrateJob
// This API is used to recover a data migration task in **Isolated** status. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) RecoverMigrateJobWithContext(ctx context.Context, request *RecoverMigrateJobRequest) (response *RecoverMigrateJobResponse, err error) {
    if request == nil {
        request = NewRecoverMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverSyncJobRequest() (request *RecoverSyncJobRequest) {
    request = &RecoverSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "RecoverSyncJob")
    
    
    return
}

func NewRecoverSyncJobResponse() (response *RecoverSyncJobResponse) {
    response = &RecoverSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecoverSyncJob
// This API is used to recover a sync task in **Isolated** status. After calling this API, you can call the `DescribeSyncJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) RecoverSyncJob(request *RecoverSyncJobRequest) (response *RecoverSyncJobResponse, err error) {
    return c.RecoverSyncJobWithContext(context.Background(), request)
}

// RecoverSyncJob
// This API is used to recover a sync task in **Isolated** status. After calling this API, you can call the `DescribeSyncJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) RecoverSyncJobWithContext(ctx context.Context, request *RecoverSyncJobRequest) (response *RecoverSyncJobResponse, err error) {
    if request == nil {
        request = NewRecoverSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewResetConsumerGroupOffsetRequest() (request *ResetConsumerGroupOffsetRequest) {
    request = &ResetConsumerGroupOffsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResetConsumerGroupOffset")
    
    
    return
}

func NewResetConsumerGroupOffsetResponse() (response *ResetConsumerGroupOffsetResponse) {
    response = &ResetConsumerGroupOffsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetConsumerGroupOffset
// This API is used to reset the offset of the subscription consumer group. Call the DescribeConsumerGroups API to query the status of the consumer group, only when the status is Dead or Empty can this operation be executed. Otherwise, the reset will not be effective and the API will not return any error.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResetConsumerGroupOffset(request *ResetConsumerGroupOffsetRequest) (response *ResetConsumerGroupOffsetResponse, err error) {
    return c.ResetConsumerGroupOffsetWithContext(context.Background(), request)
}

// ResetConsumerGroupOffset
// This API is used to reset the offset of the subscription consumer group. Call the DescribeConsumerGroups API to query the status of the consumer group, only when the status is Dead or Empty can this operation be executed. Otherwise, the reset will not be effective and the API will not return any error.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResetConsumerGroupOffsetWithContext(ctx context.Context, request *ResetConsumerGroupOffsetRequest) (response *ResetConsumerGroupOffsetResponse, err error) {
    if request == nil {
        request = NewResetConsumerGroupOffsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetConsumerGroupOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetConsumerGroupOffsetResponse()
    err = c.Send(request, response)
    return
}

func NewResetSubscribeRequest() (request *ResetSubscribeRequest) {
    request = &ResetSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResetSubscribe")
    
    
    return
}

func NewResetSubscribeResponse() (response *ResetSubscribeResponse) {
    response = &ResetSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetSubscribe
// This API is used to reset the subscription instance. After resetting, the subscription task can be reconfigured.You can call DescribeSubscribeDetail to query the subscription information to determine whether the subscription is successful. When SubsStatus changes to notStarted, it means the reset is successful.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResetSubscribe(request *ResetSubscribeRequest) (response *ResetSubscribeResponse, err error) {
    return c.ResetSubscribeWithContext(context.Background(), request)
}

// ResetSubscribe
// This API is used to reset the subscription instance. After resetting, the subscription task can be reconfigured.You can call DescribeSubscribeDetail to query the subscription information to determine whether the subscription is successful. When SubsStatus changes to notStarted, it means the reset is successful.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResetSubscribeWithContext(ctx context.Context, request *ResetSubscribeRequest) (response *ResetSubscribeResponse, err error) {
    if request == nil {
        request = NewResetSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewResizeSyncJobRequest() (request *ResizeSyncJobRequest) {
    request = &ResizeSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResizeSyncJob")
    
    
    return
}

func NewResizeSyncJobResponse() (response *ResizeSyncJobResponse) {
    response = &ResizeSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResizeSyncJob
// This API is used to adjust the specification of a pay-as-you-go sync task. After this API is called, the backend needs to take about 3-5 minutes to implement the adjustment. You can call the `DescribeSyncJobs` API to query the latest task status.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_NOTENOUGHMONEYERROR = "InternalError.NotEnoughMoneyError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResizeSyncJob(request *ResizeSyncJobRequest) (response *ResizeSyncJobResponse, err error) {
    return c.ResizeSyncJobWithContext(context.Background(), request)
}

// ResizeSyncJob
// This API is used to adjust the specification of a pay-as-you-go sync task. After this API is called, the backend needs to take about 3-5 minutes to implement the adjustment. You can call the `DescribeSyncJobs` API to query the latest task status.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_NOTENOUGHMONEYERROR = "InternalError.NotEnoughMoneyError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResizeSyncJobWithContext(ctx context.Context, request *ResizeSyncJobRequest) (response *ResizeSyncJobResponse, err error) {
    if request == nil {
        request = NewResizeSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewResumeMigrateJobRequest() (request *ResumeMigrateJobRequest) {
    request = &ResumeMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResumeMigrateJob")
    
    
    return
}

func NewResumeMigrateJobResponse() (response *ResumeMigrateJobResponse) {
    response = &ResumeMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeMigrateJob
// This API is used to retry an abnormal or failed Redis data migration task. Note that this operation will skip the check stage and directly start the task, just like by calling `StartMigrationJob`. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResumeMigrateJob(request *ResumeMigrateJobRequest) (response *ResumeMigrateJobResponse, err error) {
    return c.ResumeMigrateJobWithContext(context.Background(), request)
}

// ResumeMigrateJob
// This API is used to retry an abnormal or failed Redis data migration task. Note that this operation will skip the check stage and directly start the task, just like by calling `StartMigrationJob`. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResumeMigrateJobWithContext(ctx context.Context, request *ResumeMigrateJobRequest) (response *ResumeMigrateJobResponse, err error) {
    if request == nil {
        request = NewResumeMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewResumeSubscribeRequest() (request *ResumeSubscribeRequest) {
    request = &ResumeSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResumeSubscribe")
    
    
    return
}

func NewResumeSubscribeResponse() (response *ResumeSubscribeResponse) {
    response = &ResumeSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeSubscribe
// This API is used to resume faulty subscription tasks. When the status of the subscription task is set to error, you can resume task via this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResumeSubscribe(request *ResumeSubscribeRequest) (response *ResumeSubscribeResponse, err error) {
    return c.ResumeSubscribeWithContext(context.Background(), request)
}

// ResumeSubscribe
// This API is used to resume faulty subscription tasks. When the status of the subscription task is set to error, you can resume task via this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResumeSubscribeWithContext(ctx context.Context, request *ResumeSubscribeRequest) (response *ResumeSubscribeResponse, err error) {
    if request == nil {
        request = NewResumeSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewResumeSyncJobRequest() (request *ResumeSyncJobRequest) {
    request = &ResumeSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResumeSyncJob")
    
    
    return
}

func NewResumeSyncJobResponse() (response *ResumeSyncJobResponse) {
    response = &ResumeSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeSyncJob
// This API is used to retry a sync task after certain resolvable errors are reported. After calling this API, you can call the `DescribeSyncJobs` API to query the latest task status.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ResumeSyncJob(request *ResumeSyncJobRequest) (response *ResumeSyncJobResponse, err error) {
    return c.ResumeSyncJobWithContext(context.Background(), request)
}

// ResumeSyncJob
// This API is used to retry a sync task after certain resolvable errors are reported. After calling this API, you can call the `DescribeSyncJobs` API to query the latest task status.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ResumeSyncJobWithContext(ctx context.Context, request *ResumeSyncJobRequest) (response *ResumeSyncJobResponse, err error) {
    if request == nil {
        request = NewResumeSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewSkipCheckItemRequest() (request *SkipCheckItemRequest) {
    request = &SkipCheckItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "SkipCheckItem")
    
    
    return
}

func NewSkipCheckItemResponse() (response *SkipCheckItemResponse) {
    response = &SkipCheckItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SkipCheckItem
// This API is used for the backend to skip a failed check item. Theoretically, to execute a migration task normally, any check step cannot be skipped, and the check must be passed. For products or links that support check item skipping, see [Check Item Overview](https://www.tencentcloud.com/document/product/571/42551).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) SkipCheckItem(request *SkipCheckItemRequest) (response *SkipCheckItemResponse, err error) {
    return c.SkipCheckItemWithContext(context.Background(), request)
}

// SkipCheckItem
// This API is used for the backend to skip a failed check item. Theoretically, to execute a migration task normally, any check step cannot be skipped, and the check must be passed. For products or links that support check item skipping, see [Check Item Overview](https://www.tencentcloud.com/document/product/571/42551).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) SkipCheckItemWithContext(ctx context.Context, request *SkipCheckItemRequest) (response *SkipCheckItemResponse, err error) {
    if request == nil {
        request = NewSkipCheckItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SkipCheckItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewSkipCheckItemResponse()
    err = c.Send(request, response)
    return
}

func NewSkipSyncCheckItemRequest() (request *SkipSyncCheckItemRequest) {
    request = &SkipSyncCheckItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "SkipSyncCheckItem")
    
    
    return
}

func NewSkipSyncCheckItemResponse() (response *SkipSyncCheckItemResponse) {
    response = &SkipSyncCheckItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SkipSyncCheckItem
// This API is used for the backend to skip a failed check item. Theoretically, to execute a sync task normally, any check step cannot be skipped, and the check must be passed. For products or links that support check item skipping, see [Check Item Overview](https://www.tencentcloud.com/document/product/571/42551).
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) SkipSyncCheckItem(request *SkipSyncCheckItemRequest) (response *SkipSyncCheckItemResponse, err error) {
    return c.SkipSyncCheckItemWithContext(context.Background(), request)
}

// SkipSyncCheckItem
// This API is used for the backend to skip a failed check item. Theoretically, to execute a sync task normally, any check step cannot be skipped, and the check must be passed. For products or links that support check item skipping, see [Check Item Overview](https://www.tencentcloud.com/document/product/571/42551).
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) SkipSyncCheckItemWithContext(ctx context.Context, request *SkipSyncCheckItemRequest) (response *SkipSyncCheckItemResponse, err error) {
    if request == nil {
        request = NewSkipSyncCheckItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SkipSyncCheckItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewSkipSyncCheckItemResponse()
    err = c.Send(request, response)
    return
}

func NewStartCompareRequest() (request *StartCompareRequest) {
    request = &StartCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartCompare")
    
    
    return
}

func NewStartCompareResponse() (response *StartCompareResponse) {
    response = &StartCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartCompare
// This API is used to start a data consistency check task after creating it by calling the `CreateCompareTask` API. After calling this API, you can call the `DescribeCompareTasks` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartCompare(request *StartCompareRequest) (response *StartCompareResponse, err error) {
    return c.StartCompareWithContext(context.Background(), request)
}

// StartCompare
// This API is used to start a data consistency check task after creating it by calling the `CreateCompareTask` API. After calling this API, you can call the `DescribeCompareTasks` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartCompareWithContext(ctx context.Context, request *StartCompareRequest) (response *StartCompareResponse, err error) {
    if request == nil {
        request = NewStartCompareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCompare require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartCompareResponse()
    err = c.Send(request, response)
    return
}

func NewStartMigrateJobRequest() (request *StartMigrateJobRequest) {
    request = &StartMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartMigrateJob")
    
    
    return
}

func NewStartMigrateJobResponse() (response *StartMigrateJobResponse) {
    response = &StartMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartMigrateJob
// This API (`StartMigrationJob`) is used to start a migration task. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STARTJOBFAILED = "FailedOperation.StartJobFailed"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartMigrateJob(request *StartMigrateJobRequest) (response *StartMigrateJobResponse, err error) {
    return c.StartMigrateJobWithContext(context.Background(), request)
}

// StartMigrateJob
// This API (`StartMigrationJob`) is used to start a migration task. After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STARTJOBFAILED = "FailedOperation.StartJobFailed"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartMigrateJobWithContext(ctx context.Context, request *StartMigrateJobRequest) (response *StartMigrateJobResponse, err error) {
    if request == nil {
        request = NewStartMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewStartModifySyncJobRequest() (request *StartModifySyncJobRequest) {
    request = &StartModifySyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartModifySyncJob")
    
    
    return
}

func NewStartModifySyncJobResponse() (response *StartModifySyncJobResponse) {
    response = &StartModifySyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartModifySyncJob
// This API is used to start the configuration modification process when the modification check task status becomes “success”.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StartModifySyncJob(request *StartModifySyncJobRequest) (response *StartModifySyncJobResponse, err error) {
    return c.StartModifySyncJobWithContext(context.Background(), request)
}

// StartModifySyncJob
// This API is used to start the configuration modification process when the modification check task status becomes “success”.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StartModifySyncJobWithContext(ctx context.Context, request *StartModifySyncJobRequest) (response *StartModifySyncJobResponse, err error) {
    if request == nil {
        request = NewStartModifySyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartModifySyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartModifySyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewStartSubscribeRequest() (request *StartSubscribeRequest) {
    request = &StartSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartSubscribe")
    
    
    return
}

func NewStartSubscribeResponse() (response *StartSubscribeResponse) {
    response = &StartSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartSubscribe
// This API is used to start a Kafka version of the data subscription instance. This interface can be called only when the status of the subscription task is checkPass.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartSubscribe(request *StartSubscribeRequest) (response *StartSubscribeResponse, err error) {
    return c.StartSubscribeWithContext(context.Background(), request)
}

// StartSubscribe
// This API is used to start a Kafka version of the data subscription instance. This interface can be called only when the status of the subscription task is checkPass.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartSubscribeWithContext(ctx context.Context, request *StartSubscribeRequest) (response *StartSubscribeResponse, err error) {
    if request == nil {
        request = NewStartSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewStartSyncJobRequest() (request *StartSyncJobRequest) {
    request = &StartSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartSyncJob")
    
    
    return
}

func NewStartSyncJobResponse() (response *StartSyncJobResponse) {
    response = &StartSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartSyncJob
// This API is used to start a sync task.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StartSyncJob(request *StartSyncJobRequest) (response *StartSyncJobResponse, err error) {
    return c.StartSyncJobWithContext(context.Background(), request)
}

// StartSyncJob
// This API is used to start a sync task.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StartSyncJobWithContext(ctx context.Context, request *StartSyncJobRequest) (response *StartSyncJobResponse, err error) {
    if request == nil {
        request = NewStartSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopCompareRequest() (request *StopCompareRequest) {
    request = &StopCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StopCompare")
    
    
    return
}

func NewStopCompareResponse() (response *StopCompareResponse) {
    response = &StopCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCompare
// This API is used to stop a data consistency check task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StopCompare(request *StopCompareRequest) (response *StopCompareResponse, err error) {
    return c.StopCompareWithContext(context.Background(), request)
}

// StopCompare
// This API is used to stop a data consistency check task.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StopCompareWithContext(ctx context.Context, request *StopCompareRequest) (response *StopCompareResponse, err error) {
    if request == nil {
        request = NewStopCompareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCompare require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCompareResponse()
    err = c.Send(request, response)
    return
}

func NewStopMigrateJobRequest() (request *StopMigrateJobRequest) {
    request = &StopMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StopMigrateJob")
    
    
    return
}

func NewStopMigrateJobResponse() (response *StopMigrateJobResponse) {
    response = &StopMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopMigrateJob
// This API is used to stop a data migration task.
//
// After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StopMigrateJob(request *StopMigrateJobRequest) (response *StopMigrateJobResponse, err error) {
    return c.StopMigrateJobWithContext(context.Background(), request)
}

// StopMigrateJob
// This API is used to stop a data migration task.
//
// After calling this API, you can call the `DescribeMigrationJobs` API to query the latest task status.
//
// error code that may be returned:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StopMigrateJobWithContext(ctx context.Context, request *StopMigrateJobRequest) (response *StopMigrateJobResponse, err error) {
    if request == nil {
        request = NewStopMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopSyncJobRequest() (request *StopSyncJobRequest) {
    request = &StopSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StopSyncJob")
    
    
    return
}

func NewStopSyncJobResponse() (response *StopSyncJobResponse) {
    response = &StopSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopSyncJob
// This API is used to stop a sync task. After calling this API, you can call the `DescribeSyncJobs` API to query the latest task status.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StopSyncJob(request *StopSyncJobRequest) (response *StopSyncJobResponse, err error) {
    return c.StopSyncJobWithContext(context.Background(), request)
}

// StopSyncJob
// This API is used to stop a sync task. After calling this API, you can call the `DescribeSyncJobs` API to query the latest task status.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StopSyncJobWithContext(ctx context.Context, request *StopSyncJobRequest) (response *StopSyncJobResponse, err error) {
    if request == nil {
        request = NewStopSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopSyncJobResponse()
    err = c.Send(request, response)
    return
}
