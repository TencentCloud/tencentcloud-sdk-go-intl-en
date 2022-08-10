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

package v20180330

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-30"

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


func NewActivateSubscribeRequest() (request *ActivateSubscribeRequest) {
    request = &ActivateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ActivateSubscribe")
    
    
    return
}

func NewActivateSubscribeResponse() (response *ActivateSubscribeResponse) {
    response = &ActivateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ActivateSubscribe
// This API is used to configure a data subscription, which can be called only for subscription instances in unconfigured status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ActivateSubscribe(request *ActivateSubscribeRequest) (response *ActivateSubscribeResponse, err error) {
    return c.ActivateSubscribeWithContext(context.Background(), request)
}

// ActivateSubscribe
// This API is used to configure a data subscription, which can be called only for subscription instances in unconfigured status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ActivateSubscribeWithContext(ctx context.Context, request *ActivateSubscribeRequest) (response *ActivateSubscribeResponse, err error) {
    if request == nil {
        request = NewActivateSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateSubscribeResponse()
    err = c.Send(request, response)
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
// This API (CompleteMigrateJob) is used to complete a data migration task.
//
// For tasks in incremental migration mode, you need to call this API before migration gets ready, so as to stop migrating incremental data.
//
// If the task status queried through the (DescribeMigrateJobs) API is ready (status=8), you can call this API to complete the migration task.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) CompleteMigrateJob(request *CompleteMigrateJobRequest) (response *CompleteMigrateJobResponse, err error) {
    return c.CompleteMigrateJobWithContext(context.Background(), request)
}

// CompleteMigrateJob
// This API (CompleteMigrateJob) is used to complete a data migration task.
//
// For tasks in incremental migration mode, you need to call this API before migration gets ready, so as to stop migrating incremental data.
//
// If the task status queried through the (DescribeMigrateJobs) API is ready (status=8), you can call this API to complete the migration task.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
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
// This API is used to create a migration check task.
//
// Before migration, you should call this API to create a check. Migration will start only if the check succeeds. You can view the check result through the `DescribeMigrateCheckJob` API.
//
// After successful check, if the migration task needs to be modified, a new check task should be created and migration will begin only after the new check succeeds.
//
// 
//
// For a finance zone link, use the domain name https://dts.ap-shenzhen-fsi.tencentcloudapi.com.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_CGWSYSTEMERROR = "InternalError.CgwSystemError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) CreateMigrateCheckJob(request *CreateMigrateCheckJobRequest) (response *CreateMigrateCheckJobResponse, err error) {
    return c.CreateMigrateCheckJobWithContext(context.Background(), request)
}

// CreateMigrateCheckJob
// This API is used to create a migration check task.
//
// Before migration, you should call this API to create a check. Migration will start only if the check succeeds. You can view the check result through the `DescribeMigrateCheckJob` API.
//
// After successful check, if the migration task needs to be modified, a new check task should be created and migration will begin only after the new check succeeds.
//
// 
//
// For a finance zone link, use the domain name https://dts.ap-shenzhen-fsi.tencentcloudapi.com.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_CGWSYSTEMERROR = "InternalError.CgwSystemError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
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

func NewCreateMigrateJobRequest() (request *CreateMigrateJobRequest) {
    request = &CreateMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateJob")
    
    
    return
}

func NewCreateMigrateJobResponse() (response *CreateMigrateJobResponse) {
    response = &CreateMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMigrateJob
// This API (CreateMigrateJob) is used to create a data migration task.
//
// 
//
// For a finance zone linkage, please use the domain name dts.ap-shenzhen-fsi.tencentcloudapi.com.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameter.BizInvalidParameterValueError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) CreateMigrateJob(request *CreateMigrateJobRequest) (response *CreateMigrateJobResponse, err error) {
    return c.CreateMigrateJobWithContext(context.Background(), request)
}

// CreateMigrateJob
// This API (CreateMigrateJob) is used to create a data migration task.
//
// 
//
// For a finance zone linkage, please use the domain name dts.ap-shenzhen-fsi.tencentcloudapi.com.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameter.BizInvalidParameterValueError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) CreateMigrateJobWithContext(ctx context.Context, request *CreateMigrateJobRequest) (response *CreateMigrateJobResponse, err error) {
    if request == nil {
        request = NewCreateMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMigrateJobResponse()
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
// This API is used to create a data subscription instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    return c.CreateSubscribeWithContext(context.Background(), request)
}

// CreateSubscribe
// This API is used to create a data subscription instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
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

func NewDeleteMigrateJobRequest() (request *DeleteMigrateJobRequest) {
    request = &DeleteMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DeleteMigrateJob")
    
    
    return
}

func NewDeleteMigrateJobResponse() (response *DeleteMigrateJobResponse) {
    response = &DeleteMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMigrateJob
// This API (DeleteMigrationJob) is used to delete a data migration task. If the task status queried through the DescribeMigrateJobs API is checking (status=3), running (status=7), ready (status=8), canceling (status=11), or completing (status=12), the task cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_MIGRATESERVICESUPPORTERROR = "OperationDenied.MigrateServiceSupportError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DeleteMigrateJob(request *DeleteMigrateJobRequest) (response *DeleteMigrateJobResponse, err error) {
    return c.DeleteMigrateJobWithContext(context.Background(), request)
}

// DeleteMigrateJob
// This API (DeleteMigrationJob) is used to delete a data migration task. If the task status queried through the DescribeMigrateJobs API is checking (status=3), running (status=7), ready (status=8), canceling (status=11), or completing (status=12), the task cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_MIGRATESERVICESUPPORTERROR = "OperationDenied.MigrateServiceSupportError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DeleteMigrateJobWithContext(ctx context.Context, request *DeleteMigrateJobRequest) (response *DeleteMigrateJobResponse, err error) {
    if request == nil {
        request = NewDeleteMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
    request = &DescribeAsyncRequestInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeAsyncRequestInfo")
    
    
    return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
    response = &DescribeAsyncRequestInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAsyncRequestInfo
// This API is used to query the execution result of a task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    return c.DescribeAsyncRequestInfoWithContext(context.Background(), request)
}

// DescribeAsyncRequestInfo
// This API is used to query the execution result of a task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncRequestInfoWithContext(ctx context.Context, request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncRequestInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateCheckJobRequest() (request *DescribeMigrateCheckJobRequest) {
    request = &DescribeMigrateCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateCheckJob")
    
    
    return
}

func NewDescribeMigrateCheckJobResponse() (response *DescribeMigrateCheckJobResponse) {
    response = &DescribeMigrateCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMigrateCheckJob
// This API is used to get the check result and query check status and progress after a check is created. 
//
// If the check succeeds, you can call the StartMigrateJob API to start migration.
//
// If the check fails, the reason can be queried. Please modify the migration configuration or adjust relevant parameters of the source/target instances through the ModifyMigrateJob API based on the error message.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DescribeMigrateCheckJob(request *DescribeMigrateCheckJobRequest) (response *DescribeMigrateCheckJobResponse, err error) {
    return c.DescribeMigrateCheckJobWithContext(context.Background(), request)
}

// DescribeMigrateCheckJob
// This API is used to get the check result and query check status and progress after a check is created. 
//
// If the check succeeds, you can call the StartMigrateJob API to start migration.
//
// If the check fails, the reason can be queried. Please modify the migration configuration or adjust relevant parameters of the source/target instances through the ModifyMigrateJob API based on the error message.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DescribeMigrateCheckJobWithContext(ctx context.Context, request *DescribeMigrateCheckJobRequest) (response *DescribeMigrateCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrateCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrateCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateJobsRequest() (request *DescribeMigrateJobsRequest) {
    request = &DescribeMigrateJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateJobs")
    
    
    return
}

func NewDescribeMigrateJobsResponse() (response *DescribeMigrateJobsResponse) {
    response = &DescribeMigrateJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMigrateJobs
// This API is used to query data migration tasks.
//
// For a finance zone linkage, please use the domain name https://dts.ap-shenzhen-fsi.tencentcloudapi.com.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DescribeMigrateJobs(request *DescribeMigrateJobsRequest) (response *DescribeMigrateJobsResponse, err error) {
    return c.DescribeMigrateJobsWithContext(context.Background(), request)
}

// DescribeMigrateJobs
// This API is used to query data migration tasks.
//
// For a finance zone linkage, please use the domain name https://dts.ap-shenzhen-fsi.tencentcloudapi.com.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DescribeMigrateJobsWithContext(ctx context.Context, request *DescribeMigrateJobsRequest) (response *DescribeMigrateJobsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrateJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrateJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionConfRequest() (request *DescribeRegionConfRequest) {
    request = &DescribeRegionConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeRegionConf")
    
    
    return
}

func NewDescribeRegionConfResponse() (response *DescribeRegionConfResponse) {
    response = &DescribeRegionConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegionConf
// This API is used to query the purchasable subscription instance regions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegionConf(request *DescribeRegionConfRequest) (response *DescribeRegionConfResponse, err error) {
    return c.DescribeRegionConfWithContext(context.Background(), request)
}

// DescribeRegionConf
// This API is used to query the purchasable subscription instance regions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegionConfWithContext(ctx context.Context, request *DescribeRegionConfRequest) (response *DescribeRegionConfResponse, err error) {
    if request == nil {
        request = NewDescribeRegionConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegionConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeConfRequest() (request *DescribeSubscribeConfRequest) {
    request = &DescribeSubscribeConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeConf")
    
    
    return
}

func NewDescribeSubscribeConfResponse() (response *DescribeSubscribeConfResponse) {
    response = &DescribeSubscribeConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubscribeConf
// This API is used to query the subscription instance configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeSubscribeConf(request *DescribeSubscribeConfRequest) (response *DescribeSubscribeConfResponse, err error) {
    return c.DescribeSubscribeConfWithContext(context.Background(), request)
}

// DescribeSubscribeConf
// This API is used to query the subscription instance configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeSubscribeConfWithContext(ctx context.Context, request *DescribeSubscribeConfRequest) (response *DescribeSubscribeConfResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribesRequest() (request *DescribeSubscribesRequest) {
    request = &DescribeSubscribesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribes")
    
    
    return
}

func NewDescribeSubscribesResponse() (response *DescribeSubscribesResponse) {
    response = &DescribeSubscribesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubscribes
// This API is used to get the information list of data subscription instances. Pagination is enabled by default with 20 results returned each time.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubscribes(request *DescribeSubscribesRequest) (response *DescribeSubscribesResponse, err error) {
    return c.DescribeSubscribesWithContext(context.Background(), request)
}

// DescribeSubscribes
// This API is used to get the information list of data subscription instances. Pagination is enabled by default with 20 results returned each time.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubscribesWithContext(ctx context.Context, request *DescribeSubscribesRequest) (response *DescribeSubscribesResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribesResponse()
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
// This API is used to isolate an hourly billed subscription instance. After this API is called, the instance will become unavailable and billing will stop for it.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameter.BizInvalidParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) IsolateSubscribe(request *IsolateSubscribeRequest) (response *IsolateSubscribeResponse, err error) {
    return c.IsolateSubscribeWithContext(context.Background(), request)
}

// IsolateSubscribe
// This API is used to isolate an hourly billed subscription instance. After this API is called, the instance will become unavailable and billing will stop for it.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameter.BizInvalidParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewModifyMigrateJobRequest() (request *ModifyMigrateJobRequest) {
    request = &ModifyMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateJob")
    
    
    return
}

func NewModifyMigrateJobResponse() (response *ModifyMigrateJobResponse) {
    response = &ModifyMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMigrateJob
// This API (ModifyMigrateJob) is used to modify a data migration task.
//
// If the status of a migration task is creating (status=1), check succeeded (status=4), check failed (status=5), or migration failed (status=10), this API can be called to modify the task, but the type of the source and target instances and the region of the target instance cannot be modified.
//
// 
//
// For a finance zone linkage, please use the domain name dts.ap-shenzhen-fsi.tencentcloudapi.com.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) ModifyMigrateJob(request *ModifyMigrateJobRequest) (response *ModifyMigrateJobResponse, err error) {
    return c.ModifyMigrateJobWithContext(context.Background(), request)
}

// ModifyMigrateJob
// This API (ModifyMigrateJob) is used to modify a data migration task.
//
// If the status of a migration task is creating (status=1), check succeeded (status=4), check failed (status=5), or migration failed (status=10), this API can be called to modify the task, but the type of the source and target instances and the region of the target instance cannot be modified.
//
// 
//
// For a finance zone linkage, please use the domain name dts.ap-shenzhen-fsi.tencentcloudapi.com.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) ModifyMigrateJobWithContext(ctx context.Context, request *ModifyMigrateJobRequest) (response *ModifyMigrateJobResponse, err error) {
    if request == nil {
        request = NewModifyMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeConsumeTimeRequest() (request *ModifySubscribeConsumeTimeRequest) {
    request = &ModifySubscribeConsumeTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeConsumeTime")
    
    
    return
}

func NewModifySubscribeConsumeTimeResponse() (response *ModifySubscribeConsumeTimeResponse) {
    response = &ModifySubscribeConsumeTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubscribeConsumeTime
// This API is used to modify the consumption time point of a data subscription channel.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeConsumeTime(request *ModifySubscribeConsumeTimeRequest) (response *ModifySubscribeConsumeTimeResponse, err error) {
    return c.ModifySubscribeConsumeTimeWithContext(context.Background(), request)
}

// ModifySubscribeConsumeTime
// This API is used to modify the consumption time point of a data subscription channel.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeConsumeTimeWithContext(ctx context.Context, request *ModifySubscribeConsumeTimeRequest) (response *ModifySubscribeConsumeTimeResponse, err error) {
    if request == nil {
        request = NewModifySubscribeConsumeTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeConsumeTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeConsumeTimeResponse()
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
// This API is used to rename a data subscription instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeName(request *ModifySubscribeNameRequest) (response *ModifySubscribeNameResponse, err error) {
    return c.ModifySubscribeNameWithContext(context.Background(), request)
}

// ModifySubscribeName
// This API is used to rename a data subscription instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
// This API is used to modify the subscription rule of a data subscription channel.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) ModifySubscribeObjects(request *ModifySubscribeObjectsRequest) (response *ModifySubscribeObjectsResponse, err error) {
    return c.ModifySubscribeObjectsWithContext(context.Background(), request)
}

// ModifySubscribeObjects
// This API is used to modify the subscription rule of a data subscription channel.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
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

func NewModifySubscribeVipVportRequest() (request *ModifySubscribeVipVportRequest) {
    request = &ModifySubscribeVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeVipVport")
    
    
    return
}

func NewModifySubscribeVipVportResponse() (response *ModifySubscribeVipVportResponse) {
    response = &ModifySubscribeVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubscribeVipVport
// This API is used to modify the IP and port number of a data subscription instance.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeVipVport(request *ModifySubscribeVipVportRequest) (response *ModifySubscribeVipVportResponse, err error) {
    return c.ModifySubscribeVipVportWithContext(context.Background(), request)
}

// ModifySubscribeVipVport
// This API is used to modify the IP and port number of a data subscription instance.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeVipVportWithContext(ctx context.Context, request *ModifySubscribeVipVportRequest) (response *ModifySubscribeVipVportResponse, err error) {
    if request == nil {
        request = NewModifySubscribeVipVportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeVipVport require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeVipVportResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineIsolatedSubscribeRequest() (request *OfflineIsolatedSubscribeRequest) {
    request = &OfflineIsolatedSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "OfflineIsolatedSubscribe")
    
    
    return
}

func NewOfflineIsolatedSubscribeResponse() (response *OfflineIsolatedSubscribeResponse) {
    response = &OfflineIsolatedSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OfflineIsolatedSubscribe
// This API is used to deactivate an isolated data subscription instance.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) OfflineIsolatedSubscribe(request *OfflineIsolatedSubscribeRequest) (response *OfflineIsolatedSubscribeResponse, err error) {
    return c.OfflineIsolatedSubscribeWithContext(context.Background(), request)
}

// OfflineIsolatedSubscribe
// This API is used to deactivate an isolated data subscription instance.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) OfflineIsolatedSubscribeWithContext(ctx context.Context, request *OfflineIsolatedSubscribeRequest) (response *OfflineIsolatedSubscribeResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineIsolatedSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineIsolatedSubscribeResponse()
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
// This API is used to reset a data subscription instance. Once reset, an activated instance can be bound to other database instances through the `ActivateSubscribe` API.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetSubscribe(request *ResetSubscribeRequest) (response *ResetSubscribeResponse, err error) {
    return c.ResetSubscribeWithContext(context.Background(), request)
}

// ResetSubscribe
// This API is used to reset a data subscription instance. Once reset, an activated instance can be bound to other database instances through the `ActivateSubscribe` API.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
// This API (StartMigrationJob) is used to start a migration task. After the API is called, non-scheduled migration tasks will start the migration immediately, while scheduled tasks will start the countdown.
//
// Before calling this API, be sure to use the CreateMigrateCheckJob API to check the data migration task, which can be started only if its status queried through the DescribeMigrateJobs API is check succeeded (status=4).
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STARTJOBFAILED = "FailedOperation.StartJobFailed"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) StartMigrateJob(request *StartMigrateJobRequest) (response *StartMigrateJobResponse, err error) {
    return c.StartMigrateJobWithContext(context.Background(), request)
}

// StartMigrateJob
// This API (StartMigrationJob) is used to start a migration task. After the API is called, non-scheduled migration tasks will start the migration immediately, while scheduled tasks will start the countdown.
//
// Before calling this API, be sure to use the CreateMigrateCheckJob API to check the data migration task, which can be started only if its status queried through the DescribeMigrateJobs API is check succeeded (status=4).
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STARTJOBFAILED = "FailedOperation.StartJobFailed"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
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
// This API (StopMigrateJob) is used to cancel a data migration task.
//
// During migration, this API can be used to cancel migration if the task status queried through the DescribeMigrateJobs API is running (status=7) or ready (status=8), and the migration task will fail.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) StopMigrateJob(request *StopMigrateJobRequest) (response *StopMigrateJobResponse, err error) {
    return c.StopMigrateJobWithContext(context.Background(), request)
}

// StopMigrateJob
// This API (StopMigrateJob) is used to cancel a data migration task.
//
// During migration, this API can be used to cancel migration if the task status queried through the DescribeMigrateJobs API is running (status=7) or ready (status=8), and the migration task will fail.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
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
