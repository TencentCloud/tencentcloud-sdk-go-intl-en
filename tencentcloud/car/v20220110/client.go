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

package v20220110

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-01-10"

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


func NewApplyConcurrentRequest() (request *ApplyConcurrentRequest) {
    request = &ApplyConcurrentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "ApplyConcurrent")
    
    
    return
}

func NewApplyConcurrentResponse() (response *ApplyConcurrentResponse) {
    response = &ApplyConcurrentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyConcurrent
// This API is used to request concurrency quota. The timeout period of the API is 20 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) ApplyConcurrent(request *ApplyConcurrentRequest) (response *ApplyConcurrentResponse, err error) {
    return c.ApplyConcurrentWithContext(context.Background(), request)
}

// ApplyConcurrent
// This API is used to request concurrency quota. The timeout period of the API is 20 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) ApplyConcurrentWithContext(ctx context.Context, request *ApplyConcurrentRequest) (response *ApplyConcurrentResponse, err error) {
    if request == nil {
        request = NewApplyConcurrentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "ApplyConcurrent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyConcurrent require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyConcurrentResponse()
    err = c.Send(request, response)
    return
}

func NewBindConcurrentPackagesToProjectRequest() (request *BindConcurrentPackagesToProjectRequest) {
    request = &BindConcurrentPackagesToProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "BindConcurrentPackagesToProject")
    
    
    return
}

func NewBindConcurrentPackagesToProjectResponse() (response *BindConcurrentPackagesToProjectResponse) {
    response = &BindConcurrentPackagesToProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindConcurrentPackagesToProject
// This API is used to bind a concurrency pack to a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) BindConcurrentPackagesToProject(request *BindConcurrentPackagesToProjectRequest) (response *BindConcurrentPackagesToProjectResponse, err error) {
    return c.BindConcurrentPackagesToProjectWithContext(context.Background(), request)
}

// BindConcurrentPackagesToProject
// This API is used to bind a concurrency pack to a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) BindConcurrentPackagesToProjectWithContext(ctx context.Context, request *BindConcurrentPackagesToProjectRequest) (response *BindConcurrentPackagesToProjectResponse, err error) {
    if request == nil {
        request = NewBindConcurrentPackagesToProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "BindConcurrentPackagesToProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindConcurrentPackagesToProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindConcurrentPackagesToProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "CreateApplication")
    
    
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplication
// This API is used to create an application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONCREATEFAIL = "FailedOperation.ApplicationCreateFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_APPLICATIONLIMITEXCEEDED = "OperationDenied.ApplicationLimitExceeded"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    return c.CreateApplicationWithContext(context.Background(), request)
}

// CreateApplication
// This API is used to create an application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONCREATEFAIL = "FailedOperation.ApplicationCreateFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_APPLICATIONLIMITEXCEEDED = "OperationDenied.ApplicationLimitExceeded"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "CreateApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProjectRequest() (request *CreateApplicationProjectRequest) {
    request = &CreateApplicationProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "CreateApplicationProject")
    
    
    return
}

func NewCreateApplicationProjectResponse() (response *CreateApplicationProjectResponse) {
    response = &CreateApplicationProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationProject
// This API is used to create a cloud application project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateApplicationProject(request *CreateApplicationProjectRequest) (response *CreateApplicationProjectResponse, err error) {
    return c.CreateApplicationProjectWithContext(context.Background(), request)
}

// CreateApplicationProject
// This API is used to create a cloud application project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateApplicationProjectWithContext(ctx context.Context, request *CreateApplicationProjectRequest) (response *CreateApplicationProjectResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "CreateApplicationProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationSnapshotRequest() (request *CreateApplicationSnapshotRequest) {
    request = &CreateApplicationSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "CreateApplicationSnapshot")
    
    
    return
}

func NewCreateApplicationSnapshotResponse() (response *CreateApplicationSnapshotResponse) {
    response = &CreateApplicationSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationSnapshot
// This API is used to create a cloud application version snapshot.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateApplicationSnapshot(request *CreateApplicationSnapshotRequest) (response *CreateApplicationSnapshotResponse, err error) {
    return c.CreateApplicationSnapshotWithContext(context.Background(), request)
}

// CreateApplicationSnapshot
// This API is used to create a cloud application version snapshot.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateApplicationSnapshotWithContext(ctx context.Context, request *CreateApplicationSnapshotRequest) (response *CreateApplicationSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateApplicationSnapshotRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "CreateApplicationSnapshot")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationVersionRequest() (request *CreateApplicationVersionRequest) {
    request = &CreateApplicationVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "CreateApplicationVersion")
    
    
    return
}

func NewCreateApplicationVersionResponse() (response *CreateApplicationVersionResponse) {
    response = &CreateApplicationVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationVersion
// This API is used to create a cloud application version.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_VERSIONCREATING = "OperationDenied.VersionCreating"
//  OPERATIONDENIED_VERSIONLIMITEXCEEDED = "OperationDenied.VersionLimitExceeded"
func (c *Client) CreateApplicationVersion(request *CreateApplicationVersionRequest) (response *CreateApplicationVersionResponse, err error) {
    return c.CreateApplicationVersionWithContext(context.Background(), request)
}

// CreateApplicationVersion
// This API is used to create a cloud application version.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_VERSIONCREATING = "OperationDenied.VersionCreating"
//  OPERATIONDENIED_VERSIONLIMITEXCEEDED = "OperationDenied.VersionLimitExceeded"
func (c *Client) CreateApplicationVersionWithContext(ctx context.Context, request *CreateApplicationVersionRequest) (response *CreateApplicationVersionResponse, err error) {
    if request == nil {
        request = NewCreateApplicationVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "CreateApplicationVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSessionRequest() (request *CreateSessionRequest) {
    request = &CreateSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "CreateSession")
    
    
    return
}

func NewCreateSessionResponse() (response *CreateSessionResponse) {
    response = &CreateSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSession
// This API is used to create a session. The timeout period of the API is 5 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOCKTIMEOUT = "FailedOperation.LockTimeout"
//  FAILEDOPERATION_PATHNOTFOUND = "FailedOperation.PathNotFound"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_ROLE = "LimitExceeded.Role"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) CreateSession(request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    return c.CreateSessionWithContext(context.Background(), request)
}

// CreateSession
// This API is used to create a session. The timeout period of the API is 5 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOCKTIMEOUT = "FailedOperation.LockTimeout"
//  FAILEDOPERATION_PATHNOTFOUND = "FailedOperation.PathNotFound"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_ROLE = "LimitExceeded.Role"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) CreateSessionWithContext(ctx context.Context, request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    if request == nil {
        request = NewCreateSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "CreateSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
    request = &DeleteApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DeleteApplication")
    
    
    return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
    response = &DeleteApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplication
// This API is used to delete a cloud application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  FAILEDOPERATION_VERSIONLOCKFAIL = "FailedOperation.VersionLockFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    return c.DeleteApplicationWithContext(context.Background(), request)
}

// DeleteApplication
// This API is used to delete a cloud application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  FAILEDOPERATION_VERSIONLOCKFAIL = "FailedOperation.VersionLockFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DeleteApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProjectsRequest() (request *DeleteApplicationProjectsRequest) {
    request = &DeleteApplicationProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DeleteApplicationProjects")
    
    
    return
}

func NewDeleteApplicationProjectsResponse() (response *DeleteApplicationProjectsResponse) {
    response = &DeleteApplicationProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplicationProjects
// This API is used to delete cloud application projects in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  FAILEDOPERATION_VERSIONLOCKFAIL = "FailedOperation.VersionLockFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteApplicationProjects(request *DeleteApplicationProjectsRequest) (response *DeleteApplicationProjectsResponse, err error) {
    return c.DeleteApplicationProjectsWithContext(context.Background(), request)
}

// DeleteApplicationProjects
// This API is used to delete cloud application projects in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  FAILEDOPERATION_VERSIONLOCKFAIL = "FailedOperation.VersionLockFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteApplicationProjectsWithContext(ctx context.Context, request *DeleteApplicationProjectsRequest) (response *DeleteApplicationProjectsResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DeleteApplicationProjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationVersionRequest() (request *DeleteApplicationVersionRequest) {
    request = &DeleteApplicationVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DeleteApplicationVersion")
    
    
    return
}

func NewDeleteApplicationVersionResponse() (response *DeleteApplicationVersionResponse) {
    response = &DeleteApplicationVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplicationVersion
// This API is used to delete a cloud application version.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  FAILEDOPERATION_VERSIONLOCKFAIL = "FailedOperation.VersionLockFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteApplicationVersion(request *DeleteApplicationVersionRequest) (response *DeleteApplicationVersionResponse, err error) {
    return c.DeleteApplicationVersionWithContext(context.Background(), request)
}

// DeleteApplicationVersion
// This API is used to delete a cloud application version.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  FAILEDOPERATION_VERSIONLOCKFAIL = "FailedOperation.VersionLockFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteApplicationVersionWithContext(ctx context.Context, request *DeleteApplicationVersionRequest) (response *DeleteApplicationVersionResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DeleteApplicationVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationFileInfoRequest() (request *DescribeApplicationFileInfoRequest) {
    request = &DescribeApplicationFileInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeApplicationFileInfo")
    
    
    return
}

func NewDescribeApplicationFileInfoResponse() (response *DescribeApplicationFileInfoResponse) {
    response = &DescribeApplicationFileInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationFileInfo
// This API is used to query application file information.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApplicationFileInfo(request *DescribeApplicationFileInfoRequest) (response *DescribeApplicationFileInfoResponse, err error) {
    return c.DescribeApplicationFileInfoWithContext(context.Background(), request)
}

// DescribeApplicationFileInfo
// This API is used to query application file information.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApplicationFileInfoWithContext(ctx context.Context, request *DescribeApplicationFileInfoRequest) (response *DescribeApplicationFileInfoResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationFileInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeApplicationFileInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationFileInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationFileInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationListRequest() (request *DescribeApplicationListRequest) {
    request = &DescribeApplicationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeApplicationList")
    
    
    return
}

func NewDescribeApplicationListResponse() (response *DescribeApplicationListResponse) {
    response = &DescribeApplicationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationList
// This API is used to query the cloud application list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeApplicationList(request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    return c.DescribeApplicationListWithContext(context.Background(), request)
}

// DescribeApplicationList
// This API is used to query the cloud application list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeApplicationListWithContext(ctx context.Context, request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeApplicationList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationPathListRequest() (request *DescribeApplicationPathListRequest) {
    request = &DescribeApplicationPathListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeApplicationPathList")
    
    
    return
}

func NewDescribeApplicationPathListResponse() (response *DescribeApplicationPathListResponse) {
    response = &DescribeApplicationPathListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationPathList
// This API is used to query the cloud application startup path list.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApplicationPathList(request *DescribeApplicationPathListRequest) (response *DescribeApplicationPathListResponse, err error) {
    return c.DescribeApplicationPathListWithContext(context.Background(), request)
}

// DescribeApplicationPathList
// This API is used to query the cloud application startup path list.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApplicationPathListWithContext(ctx context.Context, request *DescribeApplicationPathListRequest) (response *DescribeApplicationPathListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationPathListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeApplicationPathList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationPathList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationPathListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationProjectAdvancedConfigRequest() (request *DescribeApplicationProjectAdvancedConfigRequest) {
    request = &DescribeApplicationProjectAdvancedConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeApplicationProjectAdvancedConfig")
    
    
    return
}

func NewDescribeApplicationProjectAdvancedConfigResponse() (response *DescribeApplicationProjectAdvancedConfigResponse) {
    response = &DescribeApplicationProjectAdvancedConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationProjectAdvancedConfig
// This API is used to obtain the advanced configuration information of a cloud application project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeApplicationProjectAdvancedConfig(request *DescribeApplicationProjectAdvancedConfigRequest) (response *DescribeApplicationProjectAdvancedConfigResponse, err error) {
    return c.DescribeApplicationProjectAdvancedConfigWithContext(context.Background(), request)
}

// DescribeApplicationProjectAdvancedConfig
// This API is used to obtain the advanced configuration information of a cloud application project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeApplicationProjectAdvancedConfigWithContext(ctx context.Context, request *DescribeApplicationProjectAdvancedConfigRequest) (response *DescribeApplicationProjectAdvancedConfigResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationProjectAdvancedConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeApplicationProjectAdvancedConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationProjectAdvancedConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationProjectAdvancedConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationProjectsRequest() (request *DescribeApplicationProjectsRequest) {
    request = &DescribeApplicationProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeApplicationProjects")
    
    
    return
}

func NewDescribeApplicationProjectsResponse() (response *DescribeApplicationProjectsResponse) {
    response = &DescribeApplicationProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationProjects
// This API is used to obtain the cloud application project list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeApplicationProjects(request *DescribeApplicationProjectsRequest) (response *DescribeApplicationProjectsResponse, err error) {
    return c.DescribeApplicationProjectsWithContext(context.Background(), request)
}

// DescribeApplicationProjects
// This API is used to obtain the cloud application project list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeApplicationProjectsWithContext(ctx context.Context, request *DescribeApplicationProjectsRequest) (response *DescribeApplicationProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationProjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeApplicationProjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationStatusRequest() (request *DescribeApplicationStatusRequest) {
    request = &DescribeApplicationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeApplicationStatus")
    
    
    return
}

func NewDescribeApplicationStatusResponse() (response *DescribeApplicationStatusResponse) {
    response = &DescribeApplicationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationStatus
// This API is used to query the running status of a cloud application and update status information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeApplicationStatus(request *DescribeApplicationStatusRequest) (response *DescribeApplicationStatusResponse, err error) {
    return c.DescribeApplicationStatusWithContext(context.Background(), request)
}

// DescribeApplicationStatus
// This API is used to query the running status of a cloud application and update status information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeApplicationStatusWithContext(ctx context.Context, request *DescribeApplicationStatusRequest) (response *DescribeApplicationStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeApplicationStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationVersionRequest() (request *DescribeApplicationVersionRequest) {
    request = &DescribeApplicationVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeApplicationVersion")
    
    
    return
}

func NewDescribeApplicationVersionResponse() (response *DescribeApplicationVersionResponse) {
    response = &DescribeApplicationVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationVersion
// This API is used to query the version information of a cloud application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeApplicationVersion(request *DescribeApplicationVersionRequest) (response *DescribeApplicationVersionResponse, err error) {
    return c.DescribeApplicationVersionWithContext(context.Background(), request)
}

// DescribeApplicationVersion
// This API is used to query the version information of a cloud application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeApplicationVersionWithContext(ctx context.Context, request *DescribeApplicationVersionRequest) (response *DescribeApplicationVersionResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeApplicationVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrentCountRequest() (request *DescribeConcurrentCountRequest) {
    request = &DescribeConcurrentCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeConcurrentCount")
    
    
    return
}

func NewDescribeConcurrentCountResponse() (response *DescribeConcurrentCountResponse) {
    response = &DescribeConcurrentCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrentCount
// This API is used to obtain the concurrency count.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeConcurrentCount(request *DescribeConcurrentCountRequest) (response *DescribeConcurrentCountResponse, err error) {
    return c.DescribeConcurrentCountWithContext(context.Background(), request)
}

// DescribeConcurrentCount
// This API is used to obtain the concurrency count.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeConcurrentCountWithContext(ctx context.Context, request *DescribeConcurrentCountRequest) (response *DescribeConcurrentCountResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrentCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeConcurrentCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrentCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrentCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrentPackagesRequest() (request *DescribeConcurrentPackagesRequest) {
    request = &DescribeConcurrentPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeConcurrentPackages")
    
    
    return
}

func NewDescribeConcurrentPackagesResponse() (response *DescribeConcurrentPackagesResponse) {
    response = &DescribeConcurrentPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrentPackages
// This API is used to obtain the cloud application concurrency pack list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrentPackages(request *DescribeConcurrentPackagesRequest) (response *DescribeConcurrentPackagesResponse, err error) {
    return c.DescribeConcurrentPackagesWithContext(context.Background(), request)
}

// DescribeConcurrentPackages
// This API is used to obtain the cloud application concurrency pack list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrentPackagesWithContext(ctx context.Context, request *DescribeConcurrentPackagesRequest) (response *DescribeConcurrentPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrentPackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeConcurrentPackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrentPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrentPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrentSummaryRequest() (request *DescribeConcurrentSummaryRequest) {
    request = &DescribeConcurrentSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeConcurrentSummary")
    
    
    return
}

func NewDescribeConcurrentSummaryResponse() (response *DescribeConcurrentSummaryResponse) {
    response = &DescribeConcurrentSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrentSummary
// This API is used to query the concurrency overview.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrentSummary(request *DescribeConcurrentSummaryRequest) (response *DescribeConcurrentSummaryResponse, err error) {
    return c.DescribeConcurrentSummaryWithContext(context.Background(), request)
}

// DescribeConcurrentSummary
// This API is used to query the concurrency overview.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrentSummaryWithContext(ctx context.Context, request *DescribeConcurrentSummaryRequest) (response *DescribeConcurrentSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrentSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeConcurrentSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrentSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrentSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosCredentialRequest() (request *DescribeCosCredentialRequest) {
    request = &DescribeCosCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DescribeCosCredential")
    
    
    return
}

func NewDescribeCosCredentialResponse() (response *DescribeCosCredentialResponse) {
    response = &DescribeCosCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosCredential
// This API is used to query COS key information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeCosCredential(request *DescribeCosCredentialRequest) (response *DescribeCosCredentialResponse, err error) {
    return c.DescribeCosCredentialWithContext(context.Background(), request)
}

// DescribeCosCredential
// This API is used to query COS key information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeCosCredentialWithContext(ctx context.Context, request *DescribeCosCredentialRequest) (response *DescribeCosCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeCosCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DescribeCosCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDestroySessionRequest() (request *DestroySessionRequest) {
    request = &DestroySessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "DestroySession")
    
    
    return
}

func NewDestroySessionResponse() (response *DestroySessionResponse) {
    response = &DestroySessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroySession
// This API is used to terminate a session. If cloud-based streaming has been enabled for this session, the cloud-based streaming will end upon session termination.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"
func (c *Client) DestroySession(request *DestroySessionRequest) (response *DestroySessionResponse, err error) {
    return c.DestroySessionWithContext(context.Background(), request)
}

// DestroySession
// This API is used to terminate a session. If cloud-based streaming has been enabled for this session, the cloud-based streaming will end upon session termination.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"
func (c *Client) DestroySessionWithContext(ctx context.Context, request *DestroySessionRequest) (response *DestroySessionResponse, err error) {
    if request == nil {
        request = NewDestroySessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "DestroySession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroySession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroySessionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationBaseInfoRequest() (request *ModifyApplicationBaseInfoRequest) {
    request = &ModifyApplicationBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "ModifyApplicationBaseInfo")
    
    
    return
}

func NewModifyApplicationBaseInfoResponse() (response *ModifyApplicationBaseInfoResponse) {
    response = &ModifyApplicationBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationBaseInfo
// This API is used to modify basic information of a cloud application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyApplicationBaseInfo(request *ModifyApplicationBaseInfoRequest) (response *ModifyApplicationBaseInfoResponse, err error) {
    return c.ModifyApplicationBaseInfoWithContext(context.Background(), request)
}

// ModifyApplicationBaseInfo
// This API is used to modify basic information of a cloud application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyApplicationBaseInfoWithContext(ctx context.Context, request *ModifyApplicationBaseInfoRequest) (response *ModifyApplicationBaseInfoResponse, err error) {
    if request == nil {
        request = NewModifyApplicationBaseInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "ModifyApplicationBaseInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationBaseInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProjectRequest() (request *ModifyApplicationProjectRequest) {
    request = &ModifyApplicationProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "ModifyApplicationProject")
    
    
    return
}

func NewModifyApplicationProjectResponse() (response *ModifyApplicationProjectResponse) {
    response = &ModifyApplicationProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationProject
// This API is used to modify a cloud application project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyApplicationProject(request *ModifyApplicationProjectRequest) (response *ModifyApplicationProjectResponse, err error) {
    return c.ModifyApplicationProjectWithContext(context.Background(), request)
}

// ModifyApplicationProject
// This API is used to modify a cloud application project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyApplicationProjectWithContext(ctx context.Context, request *ModifyApplicationProjectRequest) (response *ModifyApplicationProjectResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "ModifyApplicationProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationVersionRequest() (request *ModifyApplicationVersionRequest) {
    request = &ModifyApplicationVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "ModifyApplicationVersion")
    
    
    return
}

func NewModifyApplicationVersionResponse() (response *ModifyApplicationVersionResponse) {
    response = &ModifyApplicationVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationVersion
// This API is used to modify the version information of a cloud application.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyApplicationVersion(request *ModifyApplicationVersionRequest) (response *ModifyApplicationVersionResponse, err error) {
    return c.ModifyApplicationVersionWithContext(context.Background(), request)
}

// ModifyApplicationVersion
// This API is used to modify the version information of a cloud application.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyApplicationVersionWithContext(ctx context.Context, request *ModifyApplicationVersionRequest) (response *ModifyApplicationVersionResponse, err error) {
    if request == nil {
        request = NewModifyApplicationVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "ModifyApplicationVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationVersionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConcurrentPackageRequest() (request *ModifyConcurrentPackageRequest) {
    request = &ModifyConcurrentPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "ModifyConcurrentPackage")
    
    
    return
}

func NewModifyConcurrentPackageResponse() (response *ModifyConcurrentPackageResponse) {
    response = &ModifyConcurrentPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConcurrentPackage
// This API is used to modify a cloud application concurrency pack.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyConcurrentPackage(request *ModifyConcurrentPackageRequest) (response *ModifyConcurrentPackageResponse, err error) {
    return c.ModifyConcurrentPackageWithContext(context.Background(), request)
}

// ModifyConcurrentPackage
// This API is used to modify a cloud application concurrency pack.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyConcurrentPackageWithContext(ctx context.Context, request *ModifyConcurrentPackageRequest) (response *ModifyConcurrentPackageResponse, err error) {
    if request == nil {
        request = NewModifyConcurrentPackageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "ModifyConcurrentPackage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConcurrentPackage require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConcurrentPackageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMobileApplicationInfoRequest() (request *ModifyMobileApplicationInfoRequest) {
    request = &ModifyMobileApplicationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "ModifyMobileApplicationInfo")
    
    
    return
}

func NewModifyMobileApplicationInfoResponse() (response *ModifyMobileApplicationInfoResponse) {
    response = &ModifyMobileApplicationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMobileApplicationInfo
// This API is used to modify the mobile application data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyMobileApplicationInfo(request *ModifyMobileApplicationInfoRequest) (response *ModifyMobileApplicationInfoResponse, err error) {
    return c.ModifyMobileApplicationInfoWithContext(context.Background(), request)
}

// ModifyMobileApplicationInfo
// This API is used to modify the mobile application data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyMobileApplicationInfoWithContext(ctx context.Context, request *ModifyMobileApplicationInfoRequest) (response *ModifyMobileApplicationInfoResponse, err error) {
    if request == nil {
        request = NewModifyMobileApplicationInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "ModifyMobileApplicationInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMobileApplicationInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMobileApplicationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewResetConcurrentPackagesRequest() (request *ResetConcurrentPackagesRequest) {
    request = &ResetConcurrentPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "ResetConcurrentPackages")
    
    
    return
}

func NewResetConcurrentPackagesResponse() (response *ResetConcurrentPackagesResponse) {
    response = &ResetConcurrentPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetConcurrentPackages
// This API is used to reset a concurrency pack and disconnect all user connections.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResetConcurrentPackages(request *ResetConcurrentPackagesRequest) (response *ResetConcurrentPackagesResponse, err error) {
    return c.ResetConcurrentPackagesWithContext(context.Background(), request)
}

// ResetConcurrentPackages
// This API is used to reset a concurrency pack and disconnect all user connections.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResetConcurrentPackagesWithContext(ctx context.Context, request *ResetConcurrentPackagesRequest) (response *ResetConcurrentPackagesResponse, err error) {
    if request == nil {
        request = NewResetConcurrentPackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "ResetConcurrentPackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetConcurrentPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetConcurrentPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewSetApplicationVersionOnlineRequest() (request *SetApplicationVersionOnlineRequest) {
    request = &SetApplicationVersionOnlineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "SetApplicationVersionOnline")
    
    
    return
}

func NewSetApplicationVersionOnlineResponse() (response *SetApplicationVersionOnlineResponse) {
    response = &SetApplicationVersionOnlineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetApplicationVersionOnline
// This API is used to launch an application version.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) SetApplicationVersionOnline(request *SetApplicationVersionOnlineRequest) (response *SetApplicationVersionOnlineResponse, err error) {
    return c.SetApplicationVersionOnlineWithContext(context.Background(), request)
}

// SetApplicationVersionOnline
// This API is used to launch an application version.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) SetApplicationVersionOnlineWithContext(ctx context.Context, request *SetApplicationVersionOnlineRequest) (response *SetApplicationVersionOnlineResponse, err error) {
    if request == nil {
        request = NewSetApplicationVersionOnlineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "SetApplicationVersionOnline")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetApplicationVersionOnline require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetApplicationVersionOnlineResponse()
    err = c.Send(request, response)
    return
}

func NewStartPublishStreamRequest() (request *StartPublishStreamRequest) {
    request = &StartPublishStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "StartPublishStream")
    
    
    return
}

func NewStartPublishStreamResponse() (response *StartPublishStreamResponse) {
    response = &StartPublishStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartPublishStream
// This API is used to start pushing your cloud application's video streams in real time to CSS. The codec for the streaming is automatically selected based on the client's (SDK) capabilities, with a default order of H.265, H.264, VP8, and VP9.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) StartPublishStream(request *StartPublishStreamRequest) (response *StartPublishStreamResponse, err error) {
    return c.StartPublishStreamWithContext(context.Background(), request)
}

// StartPublishStream
// This API is used to start pushing your cloud application's video streams in real time to CSS. The codec for the streaming is automatically selected based on the client's (SDK) capabilities, with a default order of H.265, H.264, VP8, and VP9.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) StartPublishStreamWithContext(ctx context.Context, request *StartPublishStreamRequest) (response *StartPublishStreamResponse, err error) {
    if request == nil {
        request = NewStartPublishStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "StartPublishStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartPublishStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartPublishStreamResponse()
    err = c.Send(request, response)
    return
}

func NewStartPublishStreamWithURLRequest() (request *StartPublishStreamWithURLRequest) {
    request = &StartPublishStreamWithURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "StartPublishStreamWithURL")
    
    
    return
}

func NewStartPublishStreamWithURLResponse() (response *StartPublishStreamWithURLResponse) {
    response = &StartPublishStreamWithURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartPublishStreamWithURL
// This API is used to start pushing your cloud application's video streams to a specified URL. The codec for the streaming is automatically selected based on the client's (SDK) capabilities, with a default order of H.265, H.264, VP8, and VP9. This streaming method will be billed separately. For details about the billing method, see [Charging for Streaming to Specified URL](https://intl.cloud.tencent.com/document/product/1547/72168?from_cn_redirect=1#98ac188a-d122-4caf-88be-05268ecefdf6).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) StartPublishStreamWithURL(request *StartPublishStreamWithURLRequest) (response *StartPublishStreamWithURLResponse, err error) {
    return c.StartPublishStreamWithURLWithContext(context.Background(), request)
}

// StartPublishStreamWithURL
// This API is used to start pushing your cloud application's video streams to a specified URL. The codec for the streaming is automatically selected based on the client's (SDK) capabilities, with a default order of H.265, H.264, VP8, and VP9. This streaming method will be billed separately. For details about the billing method, see [Charging for Streaming to Specified URL](https://intl.cloud.tencent.com/document/product/1547/72168?from_cn_redirect=1#98ac188a-d122-4caf-88be-05268ecefdf6).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) StartPublishStreamWithURLWithContext(ctx context.Context, request *StartPublishStreamWithURLRequest) (response *StartPublishStreamWithURLResponse, err error) {
    if request == nil {
        request = NewStartPublishStreamWithURLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "StartPublishStreamWithURL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartPublishStreamWithURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartPublishStreamWithURLResponse()
    err = c.Send(request, response)
    return
}

func NewStopPublishStreamRequest() (request *StopPublishStreamRequest) {
    request = &StopPublishStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "StopPublishStream")
    
    
    return
}

func NewStopPublishStreamResponse() (response *StopPublishStreamResponse) {
    response = &StopPublishStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopPublishStream
// This API is used to stop pushing streams.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) StopPublishStream(request *StopPublishStreamRequest) (response *StopPublishStreamResponse, err error) {
    return c.StopPublishStreamWithContext(context.Background(), request)
}

// StopPublishStream
// This API is used to stop pushing streams.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) StopPublishStreamWithContext(ctx context.Context, request *StopPublishStreamRequest) (response *StopPublishStreamResponse, err error) {
    if request == nil {
        request = NewStopPublishStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "StopPublishStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopPublishStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopPublishStreamResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindConcurrentPackagesFromProjectRequest() (request *UnbindConcurrentPackagesFromProjectRequest) {
    request = &UnbindConcurrentPackagesFromProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("car", APIVersion, "UnbindConcurrentPackagesFromProject")
    
    
    return
}

func NewUnbindConcurrentPackagesFromProjectResponse() (response *UnbindConcurrentPackagesFromProjectResponse) {
    response = &UnbindConcurrentPackagesFromProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindConcurrentPackagesFromProject
// This API is used to unbind a concurrency pack from a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnbindConcurrentPackagesFromProject(request *UnbindConcurrentPackagesFromProjectRequest) (response *UnbindConcurrentPackagesFromProjectResponse, err error) {
    return c.UnbindConcurrentPackagesFromProjectWithContext(context.Background(), request)
}

// UnbindConcurrentPackagesFromProject
// This API is used to unbind a concurrency pack from a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnbindConcurrentPackagesFromProjectWithContext(ctx context.Context, request *UnbindConcurrentPackagesFromProjectRequest) (response *UnbindConcurrentPackagesFromProjectResponse, err error) {
    if request == nil {
        request = NewUnbindConcurrentPackagesFromProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "car", APIVersion, "UnbindConcurrentPackagesFromProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindConcurrentPackagesFromProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindConcurrentPackagesFromProjectResponse()
    err = c.Send(request, response)
    return
}
