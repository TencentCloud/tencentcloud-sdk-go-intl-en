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

package v20210622

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-06-22"

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


func NewCreateLogExportRequest() (request *CreateLogExportRequest) {
    request = &CreateLogExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "CreateLogExport")
    
    
    return
}

func NewCreateLogExportResponse() (response *CreateLogExportResponse) {
    response = &CreateLogExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLogExport
// API domain name: `rum.tencentcloudapi.com`.
//
// 
//
// This API is used to create a log download task.
//
// 
//
// Default API request rate limit: 20 requests/sec.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateLogExport(request *CreateLogExportRequest) (response *CreateLogExportResponse, err error) {
    if request == nil {
        request = NewCreateLogExportRequest()
    }
    
    response = NewCreateLogExportResponse()
    err = c.Send(request, response)
    return
}

// CreateLogExport
// API domain name: `rum.tencentcloudapi.com`.
//
// 
//
// This API is used to create a log download task.
//
// 
//
// Default API request rate limit: 20 requests/sec.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateLogExportWithContext(ctx context.Context, request *CreateLogExportRequest) (response *CreateLogExportResponse, err error) {
    if request == nil {
        request = NewCreateLogExportRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateLogExportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOfflineLogConfigRequest() (request *CreateOfflineLogConfigRequest) {
    request = &CreateOfflineLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "CreateOfflineLogConfig")
    
    
    return
}

func NewCreateOfflineLogConfigResponse() (response *CreateOfflineLogConfigResponse) {
    response = &CreateOfflineLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOfflineLogConfig
// This API is used to create an offline log listener to report offline logs of particular users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateOfflineLogConfig(request *CreateOfflineLogConfigRequest) (response *CreateOfflineLogConfigResponse, err error) {
    if request == nil {
        request = NewCreateOfflineLogConfigRequest()
    }
    
    response = NewCreateOfflineLogConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateOfflineLogConfig
// This API is used to create an offline log listener to report offline logs of particular users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateOfflineLogConfigWithContext(ctx context.Context, request *CreateOfflineLogConfigRequest) (response *CreateOfflineLogConfigResponse, err error) {
    if request == nil {
        request = NewCreateOfflineLogConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateOfflineLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProject
// This API is used to create a project (owned by the specified team).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

// CreateProject
// This API is used to create a project (owned by the specified team).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReleaseFileRequest() (request *CreateReleaseFileRequest) {
    request = &CreateReleaseFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "CreateReleaseFile")
    
    
    return
}

func NewCreateReleaseFileResponse() (response *CreateReleaseFileResponse) {
    response = &CreateReleaseFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReleaseFile
// This API is used to create a file record for the specified project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReleaseFile(request *CreateReleaseFileRequest) (response *CreateReleaseFileResponse, err error) {
    if request == nil {
        request = NewCreateReleaseFileRequest()
    }
    
    response = NewCreateReleaseFileResponse()
    err = c.Send(request, response)
    return
}

// CreateReleaseFile
// This API is used to create a file record for the specified project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReleaseFileWithContext(ctx context.Context, request *CreateReleaseFileRequest) (response *CreateReleaseFileResponse, err error) {
    if request == nil {
        request = NewCreateReleaseFileRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateReleaseFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStarProjectRequest() (request *CreateStarProjectRequest) {
    request = &CreateStarProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "CreateStarProject")
    
    
    return
}

func NewCreateStarProjectResponse() (response *CreateStarProjectResponse) {
    response = &CreateStarProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStarProject
// This API is used to add a starred project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateStarProject(request *CreateStarProjectRequest) (response *CreateStarProjectResponse, err error) {
    if request == nil {
        request = NewCreateStarProjectRequest()
    }
    
    response = NewCreateStarProjectResponse()
    err = c.Send(request, response)
    return
}

// CreateStarProject
// This API is used to add a starred project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateStarProjectWithContext(ctx context.Context, request *CreateStarProjectRequest) (response *CreateStarProjectResponse, err error) {
    if request == nil {
        request = NewCreateStarProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateStarProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTawInstanceRequest() (request *CreateTawInstanceRequest) {
    request = &CreateTawInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "CreateTawInstance")
    
    
    return
}

func NewCreateTawInstanceResponse() (response *CreateTawInstanceResponse) {
    response = &CreateTawInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTawInstance
// This API is used to create a RUM instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTawInstance(request *CreateTawInstanceRequest) (response *CreateTawInstanceResponse, err error) {
    if request == nil {
        request = NewCreateTawInstanceRequest()
    }
    
    response = NewCreateTawInstanceResponse()
    err = c.Send(request, response)
    return
}

// CreateTawInstance
// This API is used to create a RUM instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTawInstanceWithContext(ctx context.Context, request *CreateTawInstanceRequest) (response *CreateTawInstanceResponse, err error) {
    if request == nil {
        request = NewCreateTawInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTawInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWhitelistRequest() (request *CreateWhitelistRequest) {
    request = &CreateWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "CreateWhitelist")
    
    
    return
}

func NewCreateWhitelistResponse() (response *CreateWhitelistResponse) {
    response = &CreateWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWhitelist
// This API is used to create an allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWhitelist(request *CreateWhitelistRequest) (response *CreateWhitelistResponse, err error) {
    if request == nil {
        request = NewCreateWhitelistRequest()
    }
    
    response = NewCreateWhitelistResponse()
    err = c.Send(request, response)
    return
}

// CreateWhitelist
// This API is used to create an allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWhitelistWithContext(ctx context.Context, request *CreateWhitelistRequest) (response *CreateWhitelistResponse, err error) {
    if request == nil {
        request = NewCreateWhitelistRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstance
// This API is used to delete an instance. The deleted instance cannot be recovered.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

// DeleteInstance
// This API is used to delete an instance. The deleted instance cannot be recovered.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogExportRequest() (request *DeleteLogExportRequest) {
    request = &DeleteLogExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DeleteLogExport")
    
    
    return
}

func NewDeleteLogExportResponse() (response *DeleteLogExportResponse) {
    response = &DeleteLogExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLogExport
// API domain name: `rum.tencentcloudapi.com`.
//
// 
//
// This API is used to delete a log download task.
//
// 
//
// Default API request rate limit: 20 requests/sec.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteLogExport(request *DeleteLogExportRequest) (response *DeleteLogExportResponse, err error) {
    if request == nil {
        request = NewDeleteLogExportRequest()
    }
    
    response = NewDeleteLogExportResponse()
    err = c.Send(request, response)
    return
}

// DeleteLogExport
// API domain name: `rum.tencentcloudapi.com`.
//
// 
//
// This API is used to delete a log download task.
//
// 
//
// Default API request rate limit: 20 requests/sec.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteLogExportWithContext(ctx context.Context, request *DeleteLogExportRequest) (response *DeleteLogExportResponse, err error) {
    if request == nil {
        request = NewDeleteLogExportRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteLogExportResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOfflineLogConfigRequest() (request *DeleteOfflineLogConfigRequest) {
    request = &DeleteOfflineLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DeleteOfflineLogConfig")
    
    
    return
}

func NewDeleteOfflineLogConfigResponse() (response *DeleteOfflineLogConfigResponse) {
    response = &DeleteOfflineLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteOfflineLogConfig
// This API is used to delete an offline RUM log listener. Then, offline logs of particular users will not be reported.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteOfflineLogConfig(request *DeleteOfflineLogConfigRequest) (response *DeleteOfflineLogConfigResponse, err error) {
    if request == nil {
        request = NewDeleteOfflineLogConfigRequest()
    }
    
    response = NewDeleteOfflineLogConfigResponse()
    err = c.Send(request, response)
    return
}

// DeleteOfflineLogConfig
// This API is used to delete an offline RUM log listener. Then, offline logs of particular users will not be reported.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteOfflineLogConfigWithContext(ctx context.Context, request *DeleteOfflineLogConfigRequest) (response *DeleteOfflineLogConfigResponse, err error) {
    if request == nil {
        request = NewDeleteOfflineLogConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteOfflineLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOfflineLogRecordRequest() (request *DeleteOfflineLogRecordRequest) {
    request = &DeleteOfflineLogRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DeleteOfflineLogRecord")
    
    
    return
}

func NewDeleteOfflineLogRecordResponse() (response *DeleteOfflineLogRecordResponse) {
    response = &DeleteOfflineLogRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteOfflineLogRecord
// This API is used to delete an offline log record.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteOfflineLogRecord(request *DeleteOfflineLogRecordRequest) (response *DeleteOfflineLogRecordResponse, err error) {
    if request == nil {
        request = NewDeleteOfflineLogRecordRequest()
    }
    
    response = NewDeleteOfflineLogRecordResponse()
    err = c.Send(request, response)
    return
}

// DeleteOfflineLogRecord
// This API is used to delete an offline log record.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteOfflineLogRecordWithContext(ctx context.Context, request *DeleteOfflineLogRecordRequest) (response *DeleteOfflineLogRecordResponse, err error) {
    if request == nil {
        request = NewDeleteOfflineLogRecordRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteOfflineLogRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DeleteProject")
    
    
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProject
// This API is used to delete the specified RUM project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

// DeleteProject
// This API is used to delete the specified RUM project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReleaseFileRequest() (request *DeleteReleaseFileRequest) {
    request = &DeleteReleaseFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DeleteReleaseFile")
    
    
    return
}

func NewDeleteReleaseFileResponse() (response *DeleteReleaseFileResponse) {
    response = &DeleteReleaseFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteReleaseFile
// This API is used to delete the specified sourcemap file.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteReleaseFile(request *DeleteReleaseFileRequest) (response *DeleteReleaseFileResponse, err error) {
    if request == nil {
        request = NewDeleteReleaseFileRequest()
    }
    
    response = NewDeleteReleaseFileResponse()
    err = c.Send(request, response)
    return
}

// DeleteReleaseFile
// This API is used to delete the specified sourcemap file.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteReleaseFileWithContext(ctx context.Context, request *DeleteReleaseFileRequest) (response *DeleteReleaseFileResponse, err error) {
    if request == nil {
        request = NewDeleteReleaseFileRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteReleaseFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStarProjectRequest() (request *DeleteStarProjectRequest) {
    request = &DeleteStarProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DeleteStarProject")
    
    
    return
}

func NewDeleteStarProjectResponse() (response *DeleteStarProjectResponse) {
    response = &DeleteStarProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStarProject
// This API is used to delete a starred project for the specified user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteStarProject(request *DeleteStarProjectRequest) (response *DeleteStarProjectResponse, err error) {
    if request == nil {
        request = NewDeleteStarProjectRequest()
    }
    
    response = NewDeleteStarProjectResponse()
    err = c.Send(request, response)
    return
}

// DeleteStarProject
// This API is used to delete a starred project for the specified user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteStarProjectWithContext(ctx context.Context, request *DeleteStarProjectRequest) (response *DeleteStarProjectResponse, err error) {
    if request == nil {
        request = NewDeleteStarProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteStarProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWhitelistRequest() (request *DeleteWhitelistRequest) {
    request = &DeleteWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DeleteWhitelist")
    
    
    return
}

func NewDeleteWhitelistResponse() (response *DeleteWhitelistResponse) {
    response = &DeleteWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWhitelist
// This API is used to delete an allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteWhitelist(request *DeleteWhitelistRequest) (response *DeleteWhitelistResponse, err error) {
    if request == nil {
        request = NewDeleteWhitelistRequest()
    }
    
    response = NewDeleteWhitelistResponse()
    err = c.Send(request, response)
    return
}

// DeleteWhitelist
// This API is used to delete an allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteWhitelistWithContext(ctx context.Context, request *DeleteWhitelistRequest) (response *DeleteWhitelistResponse, err error) {
    if request == nil {
        request = NewDeleteWhitelistRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataRequest() (request *DescribeDataRequest) {
    request = &DescribeDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeData")
    
    
    return
}

func NewDescribeDataResponse() (response *DescribeDataResponse) {
    response = &DescribeDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeData
// This API is used to query the forwarding monitor.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeData(request *DescribeDataRequest) (response *DescribeDataResponse, err error) {
    if request == nil {
        request = NewDescribeDataRequest()
    }
    
    response = NewDescribeDataResponse()
    err = c.Send(request, response)
    return
}

// DescribeData
// This API is used to query the forwarding monitor.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataWithContext(ctx context.Context, request *DescribeDataRequest) (response *DescribeDataResponse, err error) {
    if request == nil {
        request = NewDescribeDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataCustomUrlRequest() (request *DescribeDataCustomUrlRequest) {
    request = &DescribeDataCustomUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataCustomUrl")
    
    
    return
}

func NewDescribeDataCustomUrlResponse() (response *DescribeDataCustomUrlResponse) {
    response = &DescribeDataCustomUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataCustomUrl
// This API is used to get the DescribeDataCustomUrl information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataCustomUrl(request *DescribeDataCustomUrlRequest) (response *DescribeDataCustomUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataCustomUrlRequest()
    }
    
    response = NewDescribeDataCustomUrlResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataCustomUrl
// This API is used to get the DescribeDataCustomUrl information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataCustomUrlWithContext(ctx context.Context, request *DescribeDataCustomUrlRequest) (response *DescribeDataCustomUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataCustomUrlRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataCustomUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataEventUrlRequest() (request *DescribeDataEventUrlRequest) {
    request = &DescribeDataEventUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataEventUrl")
    
    
    return
}

func NewDescribeDataEventUrlResponse() (response *DescribeDataEventUrlResponse) {
    response = &DescribeDataEventUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataEventUrl
// This API is used to get the DescribeDataEventUrl information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataEventUrl(request *DescribeDataEventUrlRequest) (response *DescribeDataEventUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataEventUrlRequest()
    }
    
    response = NewDescribeDataEventUrlResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataEventUrl
// This API is used to get the DescribeDataEventUrl information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataEventUrlWithContext(ctx context.Context, request *DescribeDataEventUrlRequest) (response *DescribeDataEventUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataEventUrlRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataEventUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataFetchProjectRequest() (request *DescribeDataFetchProjectRequest) {
    request = &DescribeDataFetchProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataFetchProject")
    
    
    return
}

func NewDescribeDataFetchProjectResponse() (response *DescribeDataFetchProjectResponse) {
    response = &DescribeDataFetchProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataFetchProject
// This API is used to get the DescribeDataFetchProject information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchProject(request *DescribeDataFetchProjectRequest) (response *DescribeDataFetchProjectResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchProjectRequest()
    }
    
    response = NewDescribeDataFetchProjectResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataFetchProject
// This API is used to get the DescribeDataFetchProject information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchProjectWithContext(ctx context.Context, request *DescribeDataFetchProjectRequest) (response *DescribeDataFetchProjectResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataFetchProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataFetchUrlRequest() (request *DescribeDataFetchUrlRequest) {
    request = &DescribeDataFetchUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataFetchUrl")
    
    
    return
}

func NewDescribeDataFetchUrlResponse() (response *DescribeDataFetchUrlResponse) {
    response = &DescribeDataFetchUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataFetchUrl
// This API is used to get the DescribeDataFetchUrl information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchUrl(request *DescribeDataFetchUrlRequest) (response *DescribeDataFetchUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchUrlRequest()
    }
    
    response = NewDescribeDataFetchUrlResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataFetchUrl
// This API is used to get the DescribeDataFetchUrl information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchUrlWithContext(ctx context.Context, request *DescribeDataFetchUrlRequest) (response *DescribeDataFetchUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchUrlRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataFetchUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataFetchUrlInfoRequest() (request *DescribeDataFetchUrlInfoRequest) {
    request = &DescribeDataFetchUrlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataFetchUrlInfo")
    
    
    return
}

func NewDescribeDataFetchUrlInfoResponse() (response *DescribeDataFetchUrlInfoResponse) {
    response = &DescribeDataFetchUrlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataFetchUrlInfo
// This API is used to get the DescribeDataFetchUrlInfo information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchUrlInfo(request *DescribeDataFetchUrlInfoRequest) (response *DescribeDataFetchUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchUrlInfoRequest()
    }
    
    response = NewDescribeDataFetchUrlInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataFetchUrlInfo
// This API is used to get the DescribeDataFetchUrlInfo information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataFetchUrlInfoWithContext(ctx context.Context, request *DescribeDataFetchUrlInfoRequest) (response *DescribeDataFetchUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataFetchUrlInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataFetchUrlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataLogUrlInfoRequest() (request *DescribeDataLogUrlInfoRequest) {
    request = &DescribeDataLogUrlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataLogUrlInfo")
    
    
    return
}

func NewDescribeDataLogUrlInfoResponse() (response *DescribeDataLogUrlInfoResponse) {
    response = &DescribeDataLogUrlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataLogUrlInfo
// This API is used to get the loginfo information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataLogUrlInfo(request *DescribeDataLogUrlInfoRequest) (response *DescribeDataLogUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataLogUrlInfoRequest()
    }
    
    response = NewDescribeDataLogUrlInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataLogUrlInfo
// This API is used to get the loginfo information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataLogUrlInfoWithContext(ctx context.Context, request *DescribeDataLogUrlInfoRequest) (response *DescribeDataLogUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataLogUrlInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataLogUrlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataLogUrlStatisticsRequest() (request *DescribeDataLogUrlStatisticsRequest) {
    request = &DescribeDataLogUrlStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataLogUrlStatistics")
    
    
    return
}

func NewDescribeDataLogUrlStatisticsResponse() (response *DescribeDataLogUrlStatisticsResponse) {
    response = &DescribeDataLogUrlStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataLogUrlStatistics
// This API is used to get the LogUrlStatistics information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataLogUrlStatistics(request *DescribeDataLogUrlStatisticsRequest) (response *DescribeDataLogUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataLogUrlStatisticsRequest()
    }
    
    response = NewDescribeDataLogUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataLogUrlStatistics
// This API is used to get the LogUrlStatistics information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataLogUrlStatisticsWithContext(ctx context.Context, request *DescribeDataLogUrlStatisticsRequest) (response *DescribeDataLogUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataLogUrlStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataLogUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPerformancePageRequest() (request *DescribeDataPerformancePageRequest) {
    request = &DescribeDataPerformancePageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPerformancePage")
    
    
    return
}

func NewDescribeDataPerformancePageResponse() (response *DescribeDataPerformancePageResponse) {
    response = &DescribeDataPerformancePageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataPerformancePage
// This API is used to get the PerformancePage information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPerformancePage(request *DescribeDataPerformancePageRequest) (response *DescribeDataPerformancePageResponse, err error) {
    if request == nil {
        request = NewDescribeDataPerformancePageRequest()
    }
    
    response = NewDescribeDataPerformancePageResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataPerformancePage
// This API is used to get the PerformancePage information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPerformancePageWithContext(ctx context.Context, request *DescribeDataPerformancePageRequest) (response *DescribeDataPerformancePageResponse, err error) {
    if request == nil {
        request = NewDescribeDataPerformancePageRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataPerformancePageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPerformanceProjectRequest() (request *DescribeDataPerformanceProjectRequest) {
    request = &DescribeDataPerformanceProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPerformanceProject")
    
    
    return
}

func NewDescribeDataPerformanceProjectResponse() (response *DescribeDataPerformanceProjectResponse) {
    response = &DescribeDataPerformanceProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataPerformanceProject
// This API is used to get the PerformanceProject information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPerformanceProject(request *DescribeDataPerformanceProjectRequest) (response *DescribeDataPerformanceProjectResponse, err error) {
    if request == nil {
        request = NewDescribeDataPerformanceProjectRequest()
    }
    
    response = NewDescribeDataPerformanceProjectResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataPerformanceProject
// This API is used to get the PerformanceProject information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPerformanceProjectWithContext(ctx context.Context, request *DescribeDataPerformanceProjectRequest) (response *DescribeDataPerformanceProjectResponse, err error) {
    if request == nil {
        request = NewDescribeDataPerformanceProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataPerformanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPvUrlInfoRequest() (request *DescribeDataPvUrlInfoRequest) {
    request = &DescribeDataPvUrlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPvUrlInfo")
    
    
    return
}

func NewDescribeDataPvUrlInfoResponse() (response *DescribeDataPvUrlInfoResponse) {
    response = &DescribeDataPvUrlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataPvUrlInfo
// This API is used to get the PvUrlInfo information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPvUrlInfo(request *DescribeDataPvUrlInfoRequest) (response *DescribeDataPvUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataPvUrlInfoRequest()
    }
    
    response = NewDescribeDataPvUrlInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataPvUrlInfo
// This API is used to get the PvUrlInfo information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPvUrlInfoWithContext(ctx context.Context, request *DescribeDataPvUrlInfoRequest) (response *DescribeDataPvUrlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataPvUrlInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataPvUrlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataPvUrlStatisticsRequest() (request *DescribeDataPvUrlStatisticsRequest) {
    request = &DescribeDataPvUrlStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataPvUrlStatistics")
    
    
    return
}

func NewDescribeDataPvUrlStatisticsResponse() (response *DescribeDataPvUrlStatisticsResponse) {
    response = &DescribeDataPvUrlStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataPvUrlStatistics
// This API is used to get the DescribeDataPvUrlStatistics information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPvUrlStatistics(request *DescribeDataPvUrlStatisticsRequest) (response *DescribeDataPvUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataPvUrlStatisticsRequest()
    }
    
    response = NewDescribeDataPvUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataPvUrlStatistics
// This API is used to get the DescribeDataPvUrlStatistics information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataPvUrlStatisticsWithContext(ctx context.Context, request *DescribeDataPvUrlStatisticsRequest) (response *DescribeDataPvUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataPvUrlStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataPvUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataReportCountRequest() (request *DescribeDataReportCountRequest) {
    request = &DescribeDataReportCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataReportCount")
    
    
    return
}

func NewDescribeDataReportCountResponse() (response *DescribeDataReportCountResponse) {
    response = &DescribeDataReportCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataReportCount
// This API is used to get the number of reported data entries for a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataReportCount(request *DescribeDataReportCountRequest) (response *DescribeDataReportCountResponse, err error) {
    if request == nil {
        request = NewDescribeDataReportCountRequest()
    }
    
    response = NewDescribeDataReportCountResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataReportCount
// This API is used to get the number of reported data entries for a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataReportCountWithContext(ctx context.Context, request *DescribeDataReportCountRequest) (response *DescribeDataReportCountResponse, err error) {
    if request == nil {
        request = NewDescribeDataReportCountRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataReportCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSetUrlStatisticsRequest() (request *DescribeDataSetUrlStatisticsRequest) {
    request = &DescribeDataSetUrlStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataSetUrlStatistics")
    
    
    return
}

func NewDescribeDataSetUrlStatisticsResponse() (response *DescribeDataSetUrlStatisticsResponse) {
    response = &DescribeDataSetUrlStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataSetUrlStatistics
// This API is used to get the DescribeDataSetUrlStatistics information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataSetUrlStatistics(request *DescribeDataSetUrlStatisticsRequest) (response *DescribeDataSetUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataSetUrlStatisticsRequest()
    }
    
    response = NewDescribeDataSetUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataSetUrlStatistics
// This API is used to get the DescribeDataSetUrlStatistics information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataSetUrlStatisticsWithContext(ctx context.Context, request *DescribeDataSetUrlStatisticsRequest) (response *DescribeDataSetUrlStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDataSetUrlStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataSetUrlStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataStaticProjectRequest() (request *DescribeDataStaticProjectRequest) {
    request = &DescribeDataStaticProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataStaticProject")
    
    
    return
}

func NewDescribeDataStaticProjectResponse() (response *DescribeDataStaticProjectResponse) {
    response = &DescribeDataStaticProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataStaticProject
// This API is used to get the DescribeDataStaticProject information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataStaticProject(request *DescribeDataStaticProjectRequest) (response *DescribeDataStaticProjectResponse, err error) {
    if request == nil {
        request = NewDescribeDataStaticProjectRequest()
    }
    
    response = NewDescribeDataStaticProjectResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataStaticProject
// This API is used to get the DescribeDataStaticProject information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataStaticProjectWithContext(ctx context.Context, request *DescribeDataStaticProjectRequest) (response *DescribeDataStaticProjectResponse, err error) {
    if request == nil {
        request = NewDescribeDataStaticProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataStaticProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataStaticResourceRequest() (request *DescribeDataStaticResourceRequest) {
    request = &DescribeDataStaticResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataStaticResource")
    
    
    return
}

func NewDescribeDataStaticResourceResponse() (response *DescribeDataStaticResourceResponse) {
    response = &DescribeDataStaticResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataStaticResource
// This API is used to get the DescribeDataStaticResource information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataStaticResource(request *DescribeDataStaticResourceRequest) (response *DescribeDataStaticResourceResponse, err error) {
    if request == nil {
        request = NewDescribeDataStaticResourceRequest()
    }
    
    response = NewDescribeDataStaticResourceResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataStaticResource
// This API is used to get the DescribeDataStaticResource information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataStaticResourceWithContext(ctx context.Context, request *DescribeDataStaticResourceRequest) (response *DescribeDataStaticResourceResponse, err error) {
    if request == nil {
        request = NewDescribeDataStaticResourceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataStaticResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataStaticUrlRequest() (request *DescribeDataStaticUrlRequest) {
    request = &DescribeDataStaticUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataStaticUrl")
    
    
    return
}

func NewDescribeDataStaticUrlResponse() (response *DescribeDataStaticUrlResponse) {
    response = &DescribeDataStaticUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataStaticUrl
// This API is used to get the DescribeDataStaticUrl information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataStaticUrl(request *DescribeDataStaticUrlRequest) (response *DescribeDataStaticUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataStaticUrlRequest()
    }
    
    response = NewDescribeDataStaticUrlResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataStaticUrl
// This API is used to get the DescribeDataStaticUrl information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataStaticUrlWithContext(ctx context.Context, request *DescribeDataStaticUrlRequest) (response *DescribeDataStaticUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDataStaticUrlRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataStaticUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataWebVitalsPageRequest() (request *DescribeDataWebVitalsPageRequest) {
    request = &DescribeDataWebVitalsPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeDataWebVitalsPage")
    
    
    return
}

func NewDescribeDataWebVitalsPageResponse() (response *DescribeDataWebVitalsPageResponse) {
    response = &DescribeDataWebVitalsPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataWebVitalsPage
// This API is used to get the DescribeDataWebVitalsPage information, which is about core user activities.
//
// It includes the Web Vitals metric for the page loading performance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataWebVitalsPage(request *DescribeDataWebVitalsPageRequest) (response *DescribeDataWebVitalsPageResponse, err error) {
    if request == nil {
        request = NewDescribeDataWebVitalsPageRequest()
    }
    
    response = NewDescribeDataWebVitalsPageResponse()
    err = c.Send(request, response)
    return
}

// DescribeDataWebVitalsPage
// This API is used to get the DescribeDataWebVitalsPage information, which is about core user activities.
//
// It includes the Web Vitals metric for the page loading performance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataWebVitalsPageWithContext(ctx context.Context, request *DescribeDataWebVitalsPageRequest) (response *DescribeDataWebVitalsPageResponse, err error) {
    if request == nil {
        request = NewDescribeDataWebVitalsPageRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDataWebVitalsPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeErrorRequest() (request *DescribeErrorRequest) {
    request = &DescribeErrorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeError")
    
    
    return
}

func NewDescribeErrorResponse() (response *DescribeErrorResponse) {
    response = &DescribeErrorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeError
// This API is used to get the homepage error information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeError(request *DescribeErrorRequest) (response *DescribeErrorResponse, err error) {
    if request == nil {
        request = NewDescribeErrorRequest()
    }
    
    response = NewDescribeErrorResponse()
    err = c.Send(request, response)
    return
}

// DescribeError
// This API is used to get the homepage error information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeErrorWithContext(ctx context.Context, request *DescribeErrorRequest) (response *DescribeErrorResponse, err error) {
    if request == nil {
        request = NewDescribeErrorRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeErrorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogExportsRequest() (request *DescribeLogExportsRequest) {
    request = &DescribeLogExportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeLogExports")
    
    
    return
}

func NewDescribeLogExportsResponse() (response *DescribeLogExportsResponse) {
    response = &DescribeLogExportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogExports
// API domain name: `rum.tencentcloudapi.com`.
//
// 
//
// This API is used to get the list of log download tasks.
//
// 
//
// Default API request rate limit: 20 requests/sec.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLogExports(request *DescribeLogExportsRequest) (response *DescribeLogExportsResponse, err error) {
    if request == nil {
        request = NewDescribeLogExportsRequest()
    }
    
    response = NewDescribeLogExportsResponse()
    err = c.Send(request, response)
    return
}

// DescribeLogExports
// API domain name: `rum.tencentcloudapi.com`.
//
// 
//
// This API is used to get the list of log download tasks.
//
// 
//
// Default API request rate limit: 20 requests/sec.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLogExportsWithContext(ctx context.Context, request *DescribeLogExportsRequest) (response *DescribeLogExportsResponse, err error) {
    if request == nil {
        request = NewDescribeLogExportsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLogExportsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogListRequest() (request *DescribeLogListRequest) {
    request = &DescribeLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeLogList")
    
    
    return
}

func NewDescribeLogListResponse() (response *DescribeLogListResponse) {
    response = &DescribeLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogList
// This API is used to get the list of logs in a project (created by an instance).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLogList(request *DescribeLogListRequest) (response *DescribeLogListResponse, err error) {
    if request == nil {
        request = NewDescribeLogListRequest()
    }
    
    response = NewDescribeLogListResponse()
    err = c.Send(request, response)
    return
}

// DescribeLogList
// This API is used to get the list of logs in a project (created by an instance).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLogListWithContext(ctx context.Context, request *DescribeLogListRequest) (response *DescribeLogListResponse, err error) {
    if request == nil {
        request = NewDescribeLogListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineLogConfigsRequest() (request *DescribeOfflineLogConfigsRequest) {
    request = &DescribeOfflineLogConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeOfflineLogConfigs")
    
    
    return
}

func NewDescribeOfflineLogConfigsResponse() (response *DescribeOfflineLogConfigsResponse) {
    response = &DescribeOfflineLogConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOfflineLogConfigs
// This API is used to get the configuration of the set offline log listener and return the unique user ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeOfflineLogConfigs(request *DescribeOfflineLogConfigsRequest) (response *DescribeOfflineLogConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineLogConfigsRequest()
    }
    
    response = NewDescribeOfflineLogConfigsResponse()
    err = c.Send(request, response)
    return
}

// DescribeOfflineLogConfigs
// This API is used to get the configuration of the set offline log listener and return the unique user ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeOfflineLogConfigsWithContext(ctx context.Context, request *DescribeOfflineLogConfigsRequest) (response *DescribeOfflineLogConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineLogConfigsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeOfflineLogConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineLogRecordsRequest() (request *DescribeOfflineLogRecordsRequest) {
    request = &DescribeOfflineLogRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeOfflineLogRecords")
    
    
    return
}

func NewDescribeOfflineLogRecordsResponse() (response *DescribeOfflineLogRecordsResponse) {
    response = &DescribeOfflineLogRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOfflineLogRecords
// This API is used to get all (up to 100) offline log records.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogRecords(request *DescribeOfflineLogRecordsRequest) (response *DescribeOfflineLogRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineLogRecordsRequest()
    }
    
    response = NewDescribeOfflineLogRecordsResponse()
    err = c.Send(request, response)
    return
}

// DescribeOfflineLogRecords
// This API is used to get all (up to 100) offline log records.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogRecordsWithContext(ctx context.Context, request *DescribeOfflineLogRecordsRequest) (response *DescribeOfflineLogRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineLogRecordsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeOfflineLogRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineLogsRequest() (request *DescribeOfflineLogsRequest) {
    request = &DescribeOfflineLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeOfflineLogs")
    
    
    return
}

func NewDescribeOfflineLogsResponse() (response *DescribeOfflineLogsResponse) {
    response = &DescribeOfflineLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOfflineLogs
// This API is used to get the specified offline log.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogs(request *DescribeOfflineLogsRequest) (response *DescribeOfflineLogsResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineLogsRequest()
    }
    
    response = NewDescribeOfflineLogsResponse()
    err = c.Send(request, response)
    return
}

// DescribeOfflineLogs
// This API is used to get the specified offline log.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOfflineLogsWithContext(ctx context.Context, request *DescribeOfflineLogsRequest) (response *DescribeOfflineLogsResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineLogsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeOfflineLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectLimitsRequest() (request *DescribeProjectLimitsRequest) {
    request = &DescribeProjectLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeProjectLimits")
    
    
    return
}

func NewDescribeProjectLimitsResponse() (response *DescribeProjectLimitsResponse) {
    response = &DescribeProjectLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjectLimits
// This API is used to get the list of project reporting rates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProjectLimits(request *DescribeProjectLimitsRequest) (response *DescribeProjectLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectLimitsRequest()
    }
    
    response = NewDescribeProjectLimitsResponse()
    err = c.Send(request, response)
    return
}

// DescribeProjectLimits
// This API is used to get the list of project reporting rates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProjectLimitsWithContext(ctx context.Context, request *DescribeProjectLimitsRequest) (response *DescribeProjectLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectLimitsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProjectLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjects
// This API is used to get the list of projects (under teams created by an instance).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

// DescribeProjects
// This API is used to get the list of projects (under teams created by an instance).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePvListRequest() (request *DescribePvListRequest) {
    request = &DescribePvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribePvList")
    
    
    return
}

func NewDescribePvListResponse() (response *DescribePvListResponse) {
    response = &DescribePvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePvList
// This API is used to get the list of PVs under a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePvList(request *DescribePvListRequest) (response *DescribePvListResponse, err error) {
    if request == nil {
        request = NewDescribePvListRequest()
    }
    
    response = NewDescribePvListResponse()
    err = c.Send(request, response)
    return
}

// DescribePvList
// This API is used to get the list of PVs under a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePvListWithContext(ctx context.Context, request *DescribePvListRequest) (response *DescribePvListResponse, err error) {
    if request == nil {
        request = NewDescribePvListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseFileSignRequest() (request *DescribeReleaseFileSignRequest) {
    request = &DescribeReleaseFileSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeReleaseFileSign")
    
    
    return
}

func NewDescribeReleaseFileSignResponse() (response *DescribeReleaseFileSignResponse) {
    response = &DescribeReleaseFileSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReleaseFileSign
// This API is used to get the temporary key for uploaded file storage.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReleaseFileSign(request *DescribeReleaseFileSignRequest) (response *DescribeReleaseFileSignResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseFileSignRequest()
    }
    
    response = NewDescribeReleaseFileSignResponse()
    err = c.Send(request, response)
    return
}

// DescribeReleaseFileSign
// This API is used to get the temporary key for uploaded file storage.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReleaseFileSignWithContext(ctx context.Context, request *DescribeReleaseFileSignRequest) (response *DescribeReleaseFileSignResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseFileSignRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeReleaseFileSignResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseFilesRequest() (request *DescribeReleaseFilesRequest) {
    request = &DescribeReleaseFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeReleaseFiles")
    
    
    return
}

func NewDescribeReleaseFilesResponse() (response *DescribeReleaseFilesResponse) {
    response = &DescribeReleaseFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReleaseFiles
// This API is used to get the list of sourcemap files of a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReleaseFiles(request *DescribeReleaseFilesRequest) (response *DescribeReleaseFilesResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseFilesRequest()
    }
    
    response = NewDescribeReleaseFilesResponse()
    err = c.Send(request, response)
    return
}

// DescribeReleaseFiles
// This API is used to get the list of sourcemap files of a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReleaseFilesWithContext(ctx context.Context, request *DescribeReleaseFilesRequest) (response *DescribeReleaseFilesResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseFilesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeReleaseFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScoresRequest() (request *DescribeScoresRequest) {
    request = &DescribeScoresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeScores")
    
    
    return
}

func NewDescribeScoresResponse() (response *DescribeScoresResponse) {
    response = &DescribeScoresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScores
// This API is used to get the list of homepage scores.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScores(request *DescribeScoresRequest) (response *DescribeScoresResponse, err error) {
    if request == nil {
        request = NewDescribeScoresRequest()
    }
    
    response = NewDescribeScoresResponse()
    err = c.Send(request, response)
    return
}

// DescribeScores
// This API is used to get the list of homepage scores.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScoresWithContext(ctx context.Context, request *DescribeScoresRequest) (response *DescribeScoresResponse, err error) {
    if request == nil {
        request = NewDescribeScoresRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeScoresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTawAreasRequest() (request *DescribeTawAreasRequest) {
    request = &DescribeTawAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTawAreas")
    
    
    return
}

func NewDescribeTawAreasResponse() (response *DescribeTawAreasResponse) {
    response = &DescribeTawAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTawAreas
// This API is used to query region information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTawAreas(request *DescribeTawAreasRequest) (response *DescribeTawAreasResponse, err error) {
    if request == nil {
        request = NewDescribeTawAreasRequest()
    }
    
    response = NewDescribeTawAreasResponse()
    err = c.Send(request, response)
    return
}

// DescribeTawAreas
// This API is used to query region information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTawAreasWithContext(ctx context.Context, request *DescribeTawAreasRequest) (response *DescribeTawAreasResponse, err error) {
    if request == nil {
        request = NewDescribeTawAreasRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTawAreasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTawInstancesRequest() (request *DescribeTawInstancesRequest) {
    request = &DescribeTawInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeTawInstances")
    
    
    return
}

func NewDescribeTawInstancesResponse() (response *DescribeTawInstancesResponse) {
    response = &DescribeTawInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTawInstances
// This API is used to query instance information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTawInstances(request *DescribeTawInstancesRequest) (response *DescribeTawInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTawInstancesRequest()
    }
    
    response = NewDescribeTawInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeTawInstances
// This API is used to query instance information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTawInstancesWithContext(ctx context.Context, request *DescribeTawInstancesRequest) (response *DescribeTawInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTawInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTawInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUvListRequest() (request *DescribeUvListRequest) {
    request = &DescribeUvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeUvList")
    
    
    return
}

func NewDescribeUvListResponse() (response *DescribeUvListResponse) {
    response = &DescribeUvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUvList
// This API is used to get the list of UVs under a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUvList(request *DescribeUvListRequest) (response *DescribeUvListResponse, err error) {
    if request == nil {
        request = NewDescribeUvListRequest()
    }
    
    response = NewDescribeUvListResponse()
    err = c.Send(request, response)
    return
}

// DescribeUvList
// This API is used to get the list of UVs under a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUvListWithContext(ctx context.Context, request *DescribeUvListRequest) (response *DescribeUvListResponse, err error) {
    if request == nil {
        request = NewDescribeUvListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhitelistsRequest() (request *DescribeWhitelistsRequest) {
    request = &DescribeWhitelistsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "DescribeWhitelists")
    
    
    return
}

func NewDescribeWhitelistsResponse() (response *DescribeWhitelistsResponse) {
    response = &DescribeWhitelistsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWhitelists
// This API is used to get the list of allowlists.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeWhitelists(request *DescribeWhitelistsRequest) (response *DescribeWhitelistsResponse, err error) {
    if request == nil {
        request = NewDescribeWhitelistsRequest()
    }
    
    response = NewDescribeWhitelistsResponse()
    err = c.Send(request, response)
    return
}

// DescribeWhitelists
// This API is used to get the list of allowlists.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeWhitelistsWithContext(ctx context.Context, request *DescribeWhitelistsRequest) (response *DescribeWhitelistsResponse, err error) {
    if request == nil {
        request = NewDescribeWhitelistsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeWhitelistsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstance
// This API is used to modify an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

// ModifyInstance
// This API is used to modify an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
    request = &ModifyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "ModifyProject")
    
    
    return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
    response = &ModifyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProject
// This API is used to modify a RUM project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}

// ModifyProject
// This API is used to modify a RUM project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyProjectWithContext(ctx context.Context, request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectLimitRequest() (request *ModifyProjectLimitRequest) {
    request = &ModifyProjectLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "ModifyProjectLimit")
    
    
    return
}

func NewModifyProjectLimitResponse() (response *ModifyProjectLimitResponse) {
    response = &ModifyProjectLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProjectLimit
// This API is used to add or modify a project limit.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProjectLimit(request *ModifyProjectLimitRequest) (response *ModifyProjectLimitResponse, err error) {
    if request == nil {
        request = NewModifyProjectLimitRequest()
    }
    
    response = NewModifyProjectLimitResponse()
    err = c.Send(request, response)
    return
}

// ModifyProjectLimit
// This API is used to add or modify a project limit.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSCALLFAIL = "FailedOperation.ClsCallFail"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DataBaseException"
//  FAILEDOPERATION_INFRASTRUCTUREERROR = "FailedOperation.InfrastructureError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOINSTANCE = "ResourceNotFound.NoInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProjectLimitWithContext(ctx context.Context, request *ModifyProjectLimitRequest) (response *ModifyProjectLimitResponse, err error) {
    if request == nil {
        request = NewModifyProjectLimitRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyProjectLimitResponse()
    err = c.Send(request, response)
    return
}

func NewResumeInstanceRequest() (request *ResumeInstanceRequest) {
    request = &ResumeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "ResumeInstance")
    
    
    return
}

func NewResumeInstanceResponse() (response *ResumeInstanceResponse) {
    response = &ResumeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeInstance
// This API is used to resume an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeInstance(request *ResumeInstanceRequest) (response *ResumeInstanceResponse, err error) {
    if request == nil {
        request = NewResumeInstanceRequest()
    }
    
    response = NewResumeInstanceResponse()
    err = c.Send(request, response)
    return
}

// ResumeInstance
// This API is used to resume an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeInstanceWithContext(ctx context.Context, request *ResumeInstanceRequest) (response *ResumeInstanceResponse, err error) {
    if request == nil {
        request = NewResumeInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewResumeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstanceRequest() (request *StopInstanceRequest) {
    request = &StopInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("rum", APIVersion, "StopInstance")
    
    
    return
}

func NewStopInstanceResponse() (response *StopInstanceResponse) {
    response = &StopInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopInstance
// This API is used to stop an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopInstance(request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
    if request == nil {
        request = NewStopInstanceRequest()
    }
    
    response = NewStopInstanceResponse()
    err = c.Send(request, response)
    return
}

// StopInstance
// This API is used to stop an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopInstanceWithContext(ctx context.Context, request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
    if request == nil {
        request = NewStopInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopInstanceResponse()
    err = c.Send(request, response)
    return
}
