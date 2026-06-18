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

package v20221121

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-11-21"

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


func NewAddNewBindRoleUserRequest() (request *AddNewBindRoleUserRequest) {
    request = &AddNewBindRoleUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "AddNewBindRoleUser")
    
    
    return
}

func NewAddNewBindRoleUserResponse() (response *AddNewBindRoleUserResponse) {
    response = &AddNewBindRoleUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddNewBindRoleUser
// This API is used to add the CAM role of Cloud Security Center (CSC) to the current account. The name of the CAM role is "csip".
func (c *Client) AddNewBindRoleUser(request *AddNewBindRoleUserRequest) (response *AddNewBindRoleUserResponse, err error) {
    return c.AddNewBindRoleUserWithContext(context.Background(), request)
}

// AddNewBindRoleUser
// This API is used to add the CAM role of Cloud Security Center (CSC) to the current account. The name of the CAM role is "csip".
func (c *Client) AddNewBindRoleUserWithContext(ctx context.Context, request *AddNewBindRoleUserRequest) (response *AddNewBindRoleUserResponse, err error) {
    if request == nil {
        request = NewAddNewBindRoleUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "AddNewBindRoleUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNewBindRoleUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNewBindRoleUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessKeyCheckTaskRequest() (request *CreateAccessKeyCheckTaskRequest) {
    request = &CreateAccessKeyCheckTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateAccessKeyCheckTask")
    
    
    return
}

func NewCreateAccessKeyCheckTaskResponse() (response *CreateAccessKeyCheckTaskResponse) {
    response = &CreateAccessKeyCheckTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccessKeyCheckTask
// Detect AK async task.
func (c *Client) CreateAccessKeyCheckTask(request *CreateAccessKeyCheckTaskRequest) (response *CreateAccessKeyCheckTaskResponse, err error) {
    return c.CreateAccessKeyCheckTaskWithContext(context.Background(), request)
}

// CreateAccessKeyCheckTask
// Detect AK async task.
func (c *Client) CreateAccessKeyCheckTaskWithContext(ctx context.Context, request *CreateAccessKeyCheckTaskRequest) (response *CreateAccessKeyCheckTaskResponse, err error) {
    if request == nil {
        request = NewCreateAccessKeyCheckTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateAccessKeyCheckTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccessKeyCheckTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccessKeyCheckTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessKeySyncTaskRequest() (request *CreateAccessKeySyncTaskRequest) {
    request = &CreateAccessKeySyncTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateAccessKeySyncTask")
    
    
    return
}

func NewCreateAccessKeySyncTaskResponse() (response *CreateAccessKeySyncTaskResponse) {
    response = &CreateAccessKeySyncTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccessKeySyncTask
// Trigger an AK asset sync task.
func (c *Client) CreateAccessKeySyncTask(request *CreateAccessKeySyncTaskRequest) (response *CreateAccessKeySyncTaskResponse, err error) {
    return c.CreateAccessKeySyncTaskWithContext(context.Background(), request)
}

// CreateAccessKeySyncTask
// Trigger an AK asset sync task.
func (c *Client) CreateAccessKeySyncTaskWithContext(ctx context.Context, request *CreateAccessKeySyncTaskRequest) (response *CreateAccessKeySyncTaskResponse, err error) {
    if request == nil {
        request = NewCreateAccessKeySyncTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateAccessKeySyncTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccessKeySyncTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccessKeySyncTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainAndIpRequest() (request *CreateDomainAndIpRequest) {
    request = &CreateDomainAndIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateDomainAndIp")
    
    
    return
}

func NewCreateDomainAndIpResponse() (response *CreateDomainAndIpResponse) {
    response = &CreateDomainAndIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDomainAndIp
// This API is used to create an asset with the specific domain/IP.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDomainAndIp(request *CreateDomainAndIpRequest) (response *CreateDomainAndIpResponse, err error) {
    return c.CreateDomainAndIpWithContext(context.Background(), request)
}

// CreateDomainAndIp
// This API is used to create an asset with the specific domain/IP.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDomainAndIpWithContext(ctx context.Context, request *CreateDomainAndIpRequest) (response *CreateDomainAndIpResponse, err error) {
    if request == nil {
        request = NewCreateDomainAndIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateDomainAndIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainAndIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainAndIpResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIaCAccessTokenRequest() (request *CreateIaCAccessTokenRequest) {
    request = &CreateIaCAccessTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateIaCAccessToken")
    
    
    return
}

func NewCreateIaCAccessTokenResponse() (response *CreateIaCAccessTokenResponse) {
    response = &CreateIaCAccessTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIaCAccessToken
// Create an IaC detection integration Token
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateIaCAccessToken(request *CreateIaCAccessTokenRequest) (response *CreateIaCAccessTokenResponse, err error) {
    return c.CreateIaCAccessTokenWithContext(context.Background(), request)
}

// CreateIaCAccessToken
// Create an IaC detection integration Token
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateIaCAccessTokenWithContext(ctx context.Context, request *CreateIaCAccessTokenRequest) (response *CreateIaCAccessTokenResponse, err error) {
    if request == nil {
        request = NewCreateIaCAccessTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateIaCAccessToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIaCAccessToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIaCAccessTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIaCFileExportJobRequest() (request *CreateIaCFileExportJobRequest) {
    request = &CreateIaCFileExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateIaCFileExportJob")
    
    
    return
}

func NewCreateIaCFileExportJobResponse() (response *CreateIaCFileExportJobResponse) {
    response = &CreateIaCFileExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIaCFileExportJob
// Create an IaC detection file export task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateIaCFileExportJob(request *CreateIaCFileExportJobRequest) (response *CreateIaCFileExportJobResponse, err error) {
    return c.CreateIaCFileExportJobWithContext(context.Background(), request)
}

// CreateIaCFileExportJob
// Create an IaC detection file export task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateIaCFileExportJobWithContext(ctx context.Context, request *CreateIaCFileExportJobRequest) (response *CreateIaCFileExportJobResponse, err error) {
    if request == nil {
        request = NewCreateIaCFileExportJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateIaCFileExportJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIaCFileExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIaCFileExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIaCFileReScanTaskRequest() (request *CreateIaCFileReScanTaskRequest) {
    request = &CreateIaCFileReScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateIaCFileReScanTask")
    
    
    return
}

func NewCreateIaCFileReScanTaskResponse() (response *CreateIaCFileReScanTaskResponse) {
    response = &CreateIaCFileReScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIaCFileReScanTask
// Create an IaC detection file rescan task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateIaCFileReScanTask(request *CreateIaCFileReScanTaskRequest) (response *CreateIaCFileReScanTaskResponse, err error) {
    return c.CreateIaCFileReScanTaskWithContext(context.Background(), request)
}

// CreateIaCFileReScanTask
// Create an IaC detection file rescan task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateIaCFileReScanTaskWithContext(ctx context.Context, request *CreateIaCFileReScanTaskRequest) (response *CreateIaCFileReScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateIaCFileReScanTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateIaCFileReScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIaCFileReScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIaCFileReScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRiskCenterScanTaskRequest() (request *CreateRiskCenterScanTaskRequest) {
    request = &CreateRiskCenterScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateRiskCenterScanTask")
    
    
    return
}

func NewCreateRiskCenterScanTaskResponse() (response *CreateRiskCenterScanTaskResponse) {
    response = &CreateRiskCenterScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRiskCenterScanTask
// This API is used to create a risk scan task. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) CreateRiskCenterScanTask(request *CreateRiskCenterScanTaskRequest) (response *CreateRiskCenterScanTaskResponse, err error) {
    return c.CreateRiskCenterScanTaskWithContext(context.Background(), request)
}

// CreateRiskCenterScanTask
// This API is used to create a risk scan task. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) CreateRiskCenterScanTaskWithContext(ctx context.Context, request *CreateRiskCenterScanTaskRequest) (response *CreateRiskCenterScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateRiskCenterScanTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateRiskCenterScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRiskCenterScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRiskCenterScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSkillScanRequest() (request *CreateSkillScanRequest) {
    request = &CreateSkillScanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "CreateSkillScan")
    
    
    return
}

func NewCreateSkillScanResponse() (response *CreateSkillScanResponse) {
    response = &CreateSkillScanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSkillScan
// Upload the Skill ZIP file to trigger asynchronous security detection. After a successful upload, use the returned ContentHash and EngineVersion to poll the DescribeSkillScanResult API to obtain the result. The upload API is idempotent, and uploading the same file with an identical hash will not create a repetition task. The detection result is reserved for 90 days. If overdue, reupload the file for detection.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) CreateSkillScan(request *CreateSkillScanRequest) (response *CreateSkillScanResponse, err error) {
    return c.CreateSkillScanWithContext(context.Background(), request)
}

// CreateSkillScan
// Upload the Skill ZIP file to trigger asynchronous security detection. After a successful upload, use the returned ContentHash and EngineVersion to poll the DescribeSkillScanResult API to obtain the result. The upload API is idempotent, and uploading the same file with an identical hash will not create a repetition task. The detection result is reserved for 90 days. If overdue, reupload the file for detection.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) CreateSkillScanWithContext(ctx context.Context, request *CreateSkillScanRequest) (response *CreateSkillScanResponse, err error) {
    if request == nil {
        request = NewCreateSkillScanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "CreateSkillScan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSkillScan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSkillScanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainAndIpRequest() (request *DeleteDomainAndIpRequest) {
    request = &DeleteDomainAndIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteDomainAndIp")
    
    
    return
}

func NewDeleteDomainAndIpResponse() (response *DeleteDomainAndIpResponse) {
    response = &DeleteDomainAndIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDomainAndIp
// This API is used to delete assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DeleteDomainAndIp(request *DeleteDomainAndIpRequest) (response *DeleteDomainAndIpResponse, err error) {
    return c.DeleteDomainAndIpWithContext(context.Background(), request)
}

// DeleteDomainAndIp
// This API is used to delete assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DeleteDomainAndIpWithContext(ctx context.Context, request *DeleteDomainAndIpRequest) (response *DeleteDomainAndIpResponse, err error) {
    if request == nil {
        request = NewDeleteDomainAndIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteDomainAndIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomainAndIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainAndIpResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIaCAccessTokenRequest() (request *DeleteIaCAccessTokenRequest) {
    request = &DeleteIaCAccessTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteIaCAccessToken")
    
    
    return
}

func NewDeleteIaCAccessTokenResponse() (response *DeleteIaCAccessTokenResponse) {
    response = &DeleteIaCAccessTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIaCAccessToken
// Delete an IaC detection integration Token
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteIaCAccessToken(request *DeleteIaCAccessTokenRequest) (response *DeleteIaCAccessTokenResponse, err error) {
    return c.DeleteIaCAccessTokenWithContext(context.Background(), request)
}

// DeleteIaCAccessToken
// Delete an IaC detection integration Token
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteIaCAccessTokenWithContext(ctx context.Context, request *DeleteIaCAccessTokenRequest) (response *DeleteIaCAccessTokenResponse, err error) {
    if request == nil {
        request = NewDeleteIaCAccessTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteIaCAccessToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIaCAccessToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIaCAccessTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIaCFileRequest() (request *DeleteIaCFileRequest) {
    request = &DeleteIaCFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteIaCFile")
    
    
    return
}

func NewDeleteIaCFileResponse() (response *DeleteIaCFileResponse) {
    response = &DeleteIaCFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIaCFile
// Delete IaC detection file
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteIaCFile(request *DeleteIaCFileRequest) (response *DeleteIaCFileResponse, err error) {
    return c.DeleteIaCFileWithContext(context.Background(), request)
}

// DeleteIaCFile
// Delete IaC detection file
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteIaCFileWithContext(ctx context.Context, request *DeleteIaCFileRequest) (response *DeleteIaCFileResponse, err error) {
    if request == nil {
        request = NewDeleteIaCFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteIaCFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIaCFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIaCFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRiskScanTaskRequest() (request *DeleteRiskScanTaskRequest) {
    request = &DeleteRiskScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DeleteRiskScanTask")
    
    
    return
}

func NewDeleteRiskScanTaskResponse() (response *DeleteRiskScanTaskResponse) {
    response = &DeleteRiskScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRiskScanTask
// This API is used to delete a risk scan task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DeleteRiskScanTask(request *DeleteRiskScanTaskRequest) (response *DeleteRiskScanTaskResponse, err error) {
    return c.DeleteRiskScanTaskWithContext(context.Background(), request)
}

// DeleteRiskScanTask
// This API is used to delete a risk scan task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DeleteRiskScanTaskWithContext(ctx context.Context, request *DeleteRiskScanTaskRequest) (response *DeleteRiskScanTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRiskScanTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DeleteRiskScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRiskScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRiskScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIAgentAssetListRequest() (request *DescribeAIAgentAssetListRequest) {
    request = &DescribeAIAgentAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAIAgentAssetList")
    
    
    return
}

func NewDescribeAIAgentAssetListResponse() (response *DescribeAIAgentAssetListResponse) {
    response = &DescribeAIAgentAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIAgentAssetList
// Retrieve the AI agent asset list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAIAgentAssetList(request *DescribeAIAgentAssetListRequest) (response *DescribeAIAgentAssetListResponse, err error) {
    return c.DescribeAIAgentAssetListWithContext(context.Background(), request)
}

// DescribeAIAgentAssetList
// Retrieve the AI agent asset list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAIAgentAssetListWithContext(ctx context.Context, request *DescribeAIAgentAssetListRequest) (response *DescribeAIAgentAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeAIAgentAssetListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAIAgentAssetList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIAgentAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIAgentAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAKAnalysisDetailRequest() (request *DescribeAKAnalysisDetailRequest) {
    request = &DescribeAKAnalysisDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAKAnalysisDetail")
    
    
    return
}

func NewDescribeAKAnalysisDetailResponse() (response *DescribeAKAnalysisDetailResponse) {
    response = &DescribeAKAnalysisDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAKAnalysisDetail
// Access key alarm record AI analysis result details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAKAnalysisDetail(request *DescribeAKAnalysisDetailRequest) (response *DescribeAKAnalysisDetailResponse, err error) {
    return c.DescribeAKAnalysisDetailWithContext(context.Background(), request)
}

// DescribeAKAnalysisDetail
// Access key alarm record AI analysis result details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAKAnalysisDetailWithContext(ctx context.Context, request *DescribeAKAnalysisDetailRequest) (response *DescribeAKAnalysisDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAKAnalysisDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAKAnalysisDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAKAnalysisDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAKAnalysisDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalCallRecordRequest() (request *DescribeAbnormalCallRecordRequest) {
    request = &DescribeAbnormalCallRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAbnormalCallRecord")
    
    
    return
}

func NewDescribeAbnormalCallRecordResponse() (response *DescribeAbnormalCallRecordResponse) {
    response = &DescribeAbnormalCallRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalCallRecord
// Retrieve the call record list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAbnormalCallRecord(request *DescribeAbnormalCallRecordRequest) (response *DescribeAbnormalCallRecordResponse, err error) {
    return c.DescribeAbnormalCallRecordWithContext(context.Background(), request)
}

// DescribeAbnormalCallRecord
// Retrieve the call record list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAbnormalCallRecordWithContext(ctx context.Context, request *DescribeAbnormalCallRecordRequest) (response *DescribeAbnormalCallRecordResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalCallRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAbnormalCallRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalCallRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalCallRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyAlarmRequest() (request *DescribeAccessKeyAlarmRequest) {
    request = &DescribeAccessKeyAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyAlarm")
    
    
    return
}

func NewDescribeAccessKeyAlarmResponse() (response *DescribeAccessKeyAlarmResponse) {
    response = &DescribeAccessKeyAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyAlarm
// List of alarm records for access keys.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyAlarm(request *DescribeAccessKeyAlarmRequest) (response *DescribeAccessKeyAlarmResponse, err error) {
    return c.DescribeAccessKeyAlarmWithContext(context.Background(), request)
}

// DescribeAccessKeyAlarm
// List of alarm records for access keys.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyAlarmWithContext(ctx context.Context, request *DescribeAccessKeyAlarmRequest) (response *DescribeAccessKeyAlarmResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyAlarmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyAlarm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyAlarmDetailRequest() (request *DescribeAccessKeyAlarmDetailRequest) {
    request = &DescribeAccessKeyAlarmDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyAlarmDetail")
    
    
    return
}

func NewDescribeAccessKeyAlarmDetailResponse() (response *DescribeAccessKeyAlarmDetailResponse) {
    response = &DescribeAccessKeyAlarmDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyAlarmDetail
// Alarm Record Details for Access Key.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyAlarmDetail(request *DescribeAccessKeyAlarmDetailRequest) (response *DescribeAccessKeyAlarmDetailResponse, err error) {
    return c.DescribeAccessKeyAlarmDetailWithContext(context.Background(), request)
}

// DescribeAccessKeyAlarmDetail
// Alarm Record Details for Access Key.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyAlarmDetailWithContext(ctx context.Context, request *DescribeAccessKeyAlarmDetailRequest) (response *DescribeAccessKeyAlarmDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyAlarmDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyAlarmDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyAlarmDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyAlarmDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyAssetRequest() (request *DescribeAccessKeyAssetRequest) {
    request = &DescribeAccessKeyAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyAsset")
    
    
    return
}

func NewDescribeAccessKeyAssetResponse() (response *DescribeAccessKeyAssetResponse) {
    response = &DescribeAccessKeyAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyAsset
// Obtain user access key asset list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyAsset(request *DescribeAccessKeyAssetRequest) (response *DescribeAccessKeyAssetResponse, err error) {
    return c.DescribeAccessKeyAssetWithContext(context.Background(), request)
}

// DescribeAccessKeyAsset
// Obtain user access key asset list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyAssetWithContext(ctx context.Context, request *DescribeAccessKeyAssetRequest) (response *DescribeAccessKeyAssetResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyRiskRequest() (request *DescribeAccessKeyRiskRequest) {
    request = &DescribeAccessKeyRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyRisk")
    
    
    return
}

func NewDescribeAccessKeyRiskResponse() (response *DescribeAccessKeyRiskResponse) {
    response = &DescribeAccessKeyRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyRisk
// Access key risk record list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyRisk(request *DescribeAccessKeyRiskRequest) (response *DescribeAccessKeyRiskResponse, err error) {
    return c.DescribeAccessKeyRiskWithContext(context.Background(), request)
}

// DescribeAccessKeyRisk
// Access key risk record list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyRiskWithContext(ctx context.Context, request *DescribeAccessKeyRiskRequest) (response *DescribeAccessKeyRiskResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyRiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyRisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyRiskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyRiskDetailRequest() (request *DescribeAccessKeyRiskDetailRequest) {
    request = &DescribeAccessKeyRiskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyRiskDetail")
    
    
    return
}

func NewDescribeAccessKeyRiskDetailResponse() (response *DescribeAccessKeyRiskDetailResponse) {
    response = &DescribeAccessKeyRiskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyRiskDetail
// Access Key Risk Record Details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyRiskDetail(request *DescribeAccessKeyRiskDetailRequest) (response *DescribeAccessKeyRiskDetailResponse, err error) {
    return c.DescribeAccessKeyRiskDetailWithContext(context.Background(), request)
}

// DescribeAccessKeyRiskDetail
// Access Key Risk Record Details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyRiskDetailWithContext(ctx context.Context, request *DescribeAccessKeyRiskDetailRequest) (response *DescribeAccessKeyRiskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyRiskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyRiskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyRiskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyRiskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyUserDetailRequest() (request *DescribeAccessKeyUserDetailRequest) {
    request = &DescribeAccessKeyUserDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyUserDetail")
    
    
    return
}

func NewDescribeAccessKeyUserDetailResponse() (response *DescribeAccessKeyUserDetailResponse) {
    response = &DescribeAccessKeyUserDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyUserDetail
// Query the user's account details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyUserDetail(request *DescribeAccessKeyUserDetailRequest) (response *DescribeAccessKeyUserDetailResponse, err error) {
    return c.DescribeAccessKeyUserDetailWithContext(context.Background(), request)
}

// DescribeAccessKeyUserDetail
// Query the user's account details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyUserDetailWithContext(ctx context.Context, request *DescribeAccessKeyUserDetailRequest) (response *DescribeAccessKeyUserDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyUserDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyUserDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyUserDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyUserDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessKeyUserListRequest() (request *DescribeAccessKeyUserListRequest) {
    request = &DescribeAccessKeyUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAccessKeyUserList")
    
    
    return
}

func NewDescribeAccessKeyUserListResponse() (response *DescribeAccessKeyUserListResponse) {
    response = &DescribeAccessKeyUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessKeyUserList
// Query the account list of a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyUserList(request *DescribeAccessKeyUserListRequest) (response *DescribeAccessKeyUserListResponse, err error) {
    return c.DescribeAccessKeyUserListWithContext(context.Background(), request)
}

// DescribeAccessKeyUserList
// Query the account list of a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAccessKeyUserListWithContext(ctx context.Context, request *DescribeAccessKeyUserListRequest) (response *DescribeAccessKeyUserListResponse, err error) {
    if request == nil {
        request = NewDescribeAccessKeyUserListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAccessKeyUserList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessKeyUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessKeyUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertListRequest() (request *DescribeAlertListRequest) {
    request = &DescribeAlertListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAlertList")
    
    
    return
}

func NewDescribeAlertListResponse() (response *DescribeAlertListResponse) {
    response = &DescribeAlertListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertList
// This API is used to list all alarms in the alert center.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAlertList(request *DescribeAlertListRequest) (response *DescribeAlertListResponse, err error) {
    return c.DescribeAlertListWithContext(context.Background(), request)
}

// DescribeAlertList
// This API is used to list all alarms in the alert center.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAlertListWithContext(ctx context.Context, request *DescribeAlertListRequest) (response *DescribeAlertListResponse, err error) {
    if request == nil {
        request = NewDescribeAlertListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAlertList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetProcessListRequest() (request *DescribeAssetProcessListRequest) {
    request = &DescribeAssetProcessListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetProcessList")
    
    
    return
}

func NewDescribeAssetProcessListResponse() (response *DescribeAssetProcessListResponse) {
    response = &DescribeAssetProcessListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetProcessList
// Query the process list of host nodes under the exposed path in cloud boundary analysis.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAssetProcessList(request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
    return c.DescribeAssetProcessListWithContext(context.Background(), request)
}

// DescribeAssetProcessList
// Query the process list of host nodes under the exposed path in cloud boundary analysis.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAssetProcessListWithContext(ctx context.Context, request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetProcessListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAssetProcessList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetProcessList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetProcessListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetRiskListRequest() (request *DescribeAssetRiskListRequest) {
    request = &DescribeAssetRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetRiskList")
    
    
    return
}

func NewDescribeAssetRiskListResponse() (response *DescribeAssetRiskListResponse) {
    response = &DescribeAssetRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetRiskList
// Cloud resource configuration risk list from asset perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAssetRiskList(request *DescribeAssetRiskListRequest) (response *DescribeAssetRiskListResponse, err error) {
    return c.DescribeAssetRiskListWithContext(context.Background(), request)
}

// DescribeAssetRiskList
// Cloud resource configuration risk list from asset perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAssetRiskListWithContext(ctx context.Context, request *DescribeAssetRiskListRequest) (response *DescribeAssetRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAssetRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetViewVulRiskListRequest() (request *DescribeAssetViewVulRiskListRequest) {
    request = &DescribeAssetViewVulRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetViewVulRiskList")
    
    
    return
}

func NewDescribeAssetViewVulRiskListResponse() (response *DescribeAssetViewVulRiskListResponse) {
    response = &DescribeAssetViewVulRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetViewVulRiskList
// Obtain Vulnerability Risk List from Asset's Perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAssetViewVulRiskList(request *DescribeAssetViewVulRiskListRequest) (response *DescribeAssetViewVulRiskListResponse, err error) {
    return c.DescribeAssetViewVulRiskListWithContext(context.Background(), request)
}

// DescribeAssetViewVulRiskList
// Obtain Vulnerability Risk List from Asset's Perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeAssetViewVulRiskListWithContext(ctx context.Context, request *DescribeAssetViewVulRiskListRequest) (response *DescribeAssetViewVulRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetViewVulRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeAssetViewVulRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetViewVulRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetViewVulRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCFWAssetStatisticsRequest() (request *DescribeCFWAssetStatisticsRequest) {
    request = &DescribeCFWAssetStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCFWAssetStatistics")
    
    
    return
}

func NewDescribeCFWAssetStatisticsResponse() (response *DescribeCFWAssetStatisticsResponse) {
    response = &DescribeCFWAssetStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCFWAssetStatistics
// Cloud Defense Asset Center Statistics
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCFWAssetStatistics(request *DescribeCFWAssetStatisticsRequest) (response *DescribeCFWAssetStatisticsResponse, err error) {
    return c.DescribeCFWAssetStatisticsWithContext(context.Background(), request)
}

// DescribeCFWAssetStatistics
// Cloud Defense Asset Center Statistics
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCFWAssetStatisticsWithContext(ctx context.Context, request *DescribeCFWAssetStatisticsRequest) (response *DescribeCFWAssetStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeCFWAssetStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCFWAssetStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCFWAssetStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCFWAssetStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCSIPRiskStatisticsRequest() (request *DescribeCSIPRiskStatisticsRequest) {
    request = &DescribeCSIPRiskStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCSIPRiskStatistics")
    
    
    return
}

func NewDescribeCSIPRiskStatisticsResponse() (response *DescribeCSIPRiskStatisticsResponse) {
    response = &DescribeCSIPRiskStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCSIPRiskStatistics
// Obtain risk center risk overview example.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCSIPRiskStatistics(request *DescribeCSIPRiskStatisticsRequest) (response *DescribeCSIPRiskStatisticsResponse, err error) {
    return c.DescribeCSIPRiskStatisticsWithContext(context.Background(), request)
}

// DescribeCSIPRiskStatistics
// Obtain risk center risk overview example.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCSIPRiskStatisticsWithContext(ctx context.Context, request *DescribeCSIPRiskStatisticsRequest) (response *DescribeCSIPRiskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeCSIPRiskStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCSIPRiskStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCSIPRiskStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCSIPRiskStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCVMAssetInfoRequest() (request *DescribeCVMAssetInfoRequest) {
    request = &DescribeCVMAssetInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCVMAssetInfo")
    
    
    return
}

func NewDescribeCVMAssetInfoResponse() (response *DescribeCVMAssetInfoResponse) {
    response = &DescribeCVMAssetInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCVMAssetInfo
// This API is used to query details of CVM assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCVMAssetInfo(request *DescribeCVMAssetInfoRequest) (response *DescribeCVMAssetInfoResponse, err error) {
    return c.DescribeCVMAssetInfoWithContext(context.Background(), request)
}

// DescribeCVMAssetInfo
// This API is used to query details of CVM assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCVMAssetInfoWithContext(ctx context.Context, request *DescribeCVMAssetInfoRequest) (response *DescribeCVMAssetInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCVMAssetInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCVMAssetInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCVMAssetInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCVMAssetInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCVMAssetsRequest() (request *DescribeCVMAssetsRequest) {
    request = &DescribeCVMAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCVMAssets")
    
    
    return
}

func NewDescribeCVMAssetsResponse() (response *DescribeCVMAssetsResponse) {
    response = &DescribeCVMAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCVMAssets
// Retrieve the cvm list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCVMAssets(request *DescribeCVMAssetsRequest) (response *DescribeCVMAssetsResponse, err error) {
    return c.DescribeCVMAssetsWithContext(context.Background(), request)
}

// DescribeCVMAssets
// Retrieve the cvm list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCVMAssetsWithContext(ctx context.Context, request *DescribeCVMAssetsRequest) (response *DescribeCVMAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeCVMAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCVMAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCVMAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCVMAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallRecordRequest() (request *DescribeCallRecordRequest) {
    request = &DescribeCallRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCallRecord")
    
    
    return
}

func NewDescribeCallRecordResponse() (response *DescribeCallRecordResponse) {
    response = &DescribeCallRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCallRecord
// Retrieve the call record list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCallRecord(request *DescribeCallRecordRequest) (response *DescribeCallRecordResponse, err error) {
    return c.DescribeCallRecordWithContext(context.Background(), request)
}

// DescribeCallRecord
// Retrieve the call record list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCallRecordWithContext(ctx context.Context, request *DescribeCallRecordRequest) (response *DescribeCallRecordResponse, err error) {
    if request == nil {
        request = NewDescribeCallRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCallRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckViewRisksRequest() (request *DescribeCheckViewRisksRequest) {
    request = &DescribeCheckViewRisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeCheckViewRisks")
    
    
    return
}

func NewDescribeCheckViewRisksResponse() (response *DescribeCheckViewRisksResponse) {
    response = &DescribeCheckViewRisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCheckViewRisks
// List of cloud resource configuration risks from a check perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCheckViewRisks(request *DescribeCheckViewRisksRequest) (response *DescribeCheckViewRisksResponse, err error) {
    return c.DescribeCheckViewRisksWithContext(context.Background(), request)
}

// DescribeCheckViewRisks
// List of cloud resource configuration risks from a check perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeCheckViewRisksWithContext(ctx context.Context, request *DescribeCheckViewRisksRequest) (response *DescribeCheckViewRisksResponse, err error) {
    if request == nil {
        request = NewDescribeCheckViewRisksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeCheckViewRisks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckViewRisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckViewRisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAssetsRequest() (request *DescribeClusterAssetsRequest) {
    request = &DescribeClusterAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeClusterAssets")
    
    
    return
}

func NewDescribeClusterAssetsResponse() (response *DescribeClusterAssetsResponse) {
    response = &DescribeClusterAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterAssets
// This example shows you how to obtain the cluster list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeClusterAssets(request *DescribeClusterAssetsRequest) (response *DescribeClusterAssetsResponse, err error) {
    return c.DescribeClusterAssetsWithContext(context.Background(), request)
}

// DescribeClusterAssets
// This example shows you how to obtain the cluster list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeClusterAssetsWithContext(ctx context.Context, request *DescribeClusterAssetsRequest) (response *DescribeClusterAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeClusterAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterPodAssetsRequest() (request *DescribeClusterPodAssetsRequest) {
    request = &DescribeClusterPodAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeClusterPodAssets")
    
    
    return
}

func NewDescribeClusterPodAssetsResponse() (response *DescribeClusterPodAssetsResponse) {
    response = &DescribeClusterPodAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterPodAssets
// This API is used to list cluster pods.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeClusterPodAssets(request *DescribeClusterPodAssetsRequest) (response *DescribeClusterPodAssetsResponse, err error) {
    return c.DescribeClusterPodAssetsWithContext(context.Background(), request)
}

// DescribeClusterPodAssets
// This API is used to list cluster pods.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeClusterPodAssetsWithContext(ctx context.Context, request *DescribeClusterPodAssetsRequest) (response *DescribeClusterPodAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterPodAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeClusterPodAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterPodAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterPodAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigCheckRulesRequest() (request *DescribeConfigCheckRulesRequest) {
    request = &DescribeConfigCheckRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeConfigCheckRules")
    
    
    return
}

func NewDescribeConfigCheckRulesResponse() (response *DescribeConfigCheckRulesResponse) {
    response = &DescribeConfigCheckRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigCheckRules
// Cloud Resource Configuration Risk Rule List Example
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeConfigCheckRules(request *DescribeConfigCheckRulesRequest) (response *DescribeConfigCheckRulesResponse, err error) {
    return c.DescribeConfigCheckRulesWithContext(context.Background(), request)
}

// DescribeConfigCheckRules
// Cloud Resource Configuration Risk Rule List Example
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeConfigCheckRulesWithContext(ctx context.Context, request *DescribeConfigCheckRulesRequest) (response *DescribeConfigCheckRulesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigCheckRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeConfigCheckRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigCheckRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigCheckRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDbAssetInfoRequest() (request *DescribeDbAssetInfoRequest) {
    request = &DescribeDbAssetInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDbAssetInfo")
    
    
    return
}

func NewDescribeDbAssetInfoResponse() (response *DescribeDbAssetInfoResponse) {
    response = &DescribeDbAssetInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDbAssetInfo
// This API is used to query details of a database asset. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDbAssetInfo(request *DescribeDbAssetInfoRequest) (response *DescribeDbAssetInfoResponse, err error) {
    return c.DescribeDbAssetInfoWithContext(context.Background(), request)
}

// DescribeDbAssetInfo
// This API is used to query details of a database asset. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDbAssetInfoWithContext(ctx context.Context, request *DescribeDbAssetInfoRequest) (response *DescribeDbAssetInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDbAssetInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDbAssetInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbAssetInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDbAssetInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDbAssetsRequest() (request *DescribeDbAssetsRequest) {
    request = &DescribeDbAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDbAssets")
    
    
    return
}

func NewDescribeDbAssetsResponse() (response *DescribeDbAssetsResponse) {
    response = &DescribeDbAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDbAssets
// This API is used to list database assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDbAssets(request *DescribeDbAssetsRequest) (response *DescribeDbAssetsResponse, err error) {
    return c.DescribeDbAssetsWithContext(context.Background(), request)
}

// DescribeDbAssets
// This API is used to list database assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDbAssetsWithContext(ctx context.Context, request *DescribeDbAssetsRequest) (response *DescribeDbAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeDbAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDbAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDbAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainAssetsRequest() (request *DescribeDomainAssetsRequest) {
    request = &DescribeDomainAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeDomainAssets")
    
    
    return
}

func NewDescribeDomainAssetsResponse() (response *DescribeDomainAssetsResponse) {
    response = &DescribeDomainAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainAssets
// This API is used to list domain assets. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDomainAssets(request *DescribeDomainAssetsRequest) (response *DescribeDomainAssetsResponse, err error) {
    return c.DescribeDomainAssetsWithContext(context.Background(), request)
}

// DescribeDomainAssets
// This API is used to list domain assets. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDomainAssetsWithContext(ctx context.Context, request *DescribeDomainAssetsRequest) (response *DescribeDomainAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeDomainAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExposeAssetCategoryRequest() (request *DescribeExposeAssetCategoryRequest) {
    request = &DescribeExposeAssetCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeExposeAssetCategory")
    
    
    return
}

func NewDescribeExposeAssetCategoryResponse() (response *DescribeExposeAssetCategoryResponse) {
    response = &DescribeExposeAssetCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExposeAssetCategory
// Cloud boundary analysis asset category.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeExposeAssetCategory(request *DescribeExposeAssetCategoryRequest) (response *DescribeExposeAssetCategoryResponse, err error) {
    return c.DescribeExposeAssetCategoryWithContext(context.Background(), request)
}

// DescribeExposeAssetCategory
// Cloud boundary analysis asset category.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeExposeAssetCategoryWithContext(ctx context.Context, request *DescribeExposeAssetCategoryRequest) (response *DescribeExposeAssetCategoryResponse, err error) {
    if request == nil {
        request = NewDescribeExposeAssetCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeExposeAssetCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExposeAssetCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExposeAssetCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExposePathRequest() (request *DescribeExposePathRequest) {
    request = &DescribeExposePathRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeExposePath")
    
    
    return
}

func NewDescribeExposePathResponse() (response *DescribeExposePathResponse) {
    response = &DescribeExposePathResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExposePath
// Query the node of the cloud boundary analysis path.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeExposePath(request *DescribeExposePathRequest) (response *DescribeExposePathResponse, err error) {
    return c.DescribeExposePathWithContext(context.Background(), request)
}

// DescribeExposePath
// Query the node of the cloud boundary analysis path.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeExposePathWithContext(ctx context.Context, request *DescribeExposePathRequest) (response *DescribeExposePathResponse, err error) {
    if request == nil {
        request = NewDescribeExposePathRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeExposePath")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExposePath require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExposePathResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExposuresRequest() (request *DescribeExposuresRequest) {
    request = &DescribeExposuresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeExposures")
    
    
    return
}

func NewDescribeExposuresResponse() (response *DescribeExposuresResponse) {
    response = &DescribeExposuresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExposures
// Cloud Boundary Analysis Asset List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeExposures(request *DescribeExposuresRequest) (response *DescribeExposuresResponse, err error) {
    return c.DescribeExposuresWithContext(context.Background(), request)
}

// DescribeExposures
// Cloud Boundary Analysis Asset List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeExposuresWithContext(ctx context.Context, request *DescribeExposuresRequest) (response *DescribeExposuresResponse, err error) {
    if request == nil {
        request = NewDescribeExposuresRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeExposures")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExposures require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExposuresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayAssetsRequest() (request *DescribeGatewayAssetsRequest) {
    request = &DescribeGatewayAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeGatewayAssets")
    
    
    return
}

func NewDescribeGatewayAssetsResponse() (response *DescribeGatewayAssetsResponse) {
    response = &DescribeGatewayAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatewayAssets
// Obtain Gateway List
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeGatewayAssets(request *DescribeGatewayAssetsRequest) (response *DescribeGatewayAssetsResponse, err error) {
    return c.DescribeGatewayAssetsWithContext(context.Background(), request)
}

// DescribeGatewayAssets
// Obtain Gateway List
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeGatewayAssetsWithContext(ctx context.Context, request *DescribeGatewayAssetsRequest) (response *DescribeGatewayAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeGatewayAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHighBaseLineRiskListRequest() (request *DescribeHighBaseLineRiskListRequest) {
    request = &DescribeHighBaseLineRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeHighBaseLineRiskList")
    
    
    return
}

func NewDescribeHighBaseLineRiskListResponse() (response *DescribeHighBaseLineRiskListResponse) {
    response = &DescribeHighBaseLineRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHighBaseLineRiskList
// Query the high-risk baseline risk list of host nodes under the exposed path in cloud boundary analysis.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeHighBaseLineRiskList(request *DescribeHighBaseLineRiskListRequest) (response *DescribeHighBaseLineRiskListResponse, err error) {
    return c.DescribeHighBaseLineRiskListWithContext(context.Background(), request)
}

// DescribeHighBaseLineRiskList
// Query the high-risk baseline risk list of host nodes under the exposed path in cloud boundary analysis.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeHighBaseLineRiskListWithContext(ctx context.Context, request *DescribeHighBaseLineRiskListRequest) (response *DescribeHighBaseLineRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeHighBaseLineRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeHighBaseLineRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHighBaseLineRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHighBaseLineRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIaCFileListRequest() (request *DescribeIaCFileListRequest) {
    request = &DescribeIaCFileListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIaCFileList")
    
    
    return
}

func NewDescribeIaCFileListResponse() (response *DescribeIaCFileListResponse) {
    response = &DescribeIaCFileListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIaCFileList
// Retrieve the IaC detection file list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCFileList(request *DescribeIaCFileListRequest) (response *DescribeIaCFileListResponse, err error) {
    return c.DescribeIaCFileListWithContext(context.Background(), request)
}

// DescribeIaCFileList
// Retrieve the IaC detection file list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCFileListWithContext(ctx context.Context, request *DescribeIaCFileListRequest) (response *DescribeIaCFileListResponse, err error) {
    if request == nil {
        request = NewDescribeIaCFileListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIaCFileList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIaCFileList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIaCFileListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIaCFileOverviewRequest() (request *DescribeIaCFileOverviewRequest) {
    request = &DescribeIaCFileOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIaCFileOverview")
    
    
    return
}

func NewDescribeIaCFileOverviewResponse() (response *DescribeIaCFileOverviewResponse) {
    response = &DescribeIaCFileOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIaCFileOverview
// Retrieve IaC detection file overview
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCFileOverview(request *DescribeIaCFileOverviewRequest) (response *DescribeIaCFileOverviewResponse, err error) {
    return c.DescribeIaCFileOverviewWithContext(context.Background(), request)
}

// DescribeIaCFileOverview
// Retrieve IaC detection file overview
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCFileOverviewWithContext(ctx context.Context, request *DescribeIaCFileOverviewRequest) (response *DescribeIaCFileOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeIaCFileOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIaCFileOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIaCFileOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIaCFileOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIaCFileReportRequest() (request *DescribeIaCFileReportRequest) {
    request = &DescribeIaCFileReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIaCFileReport")
    
    
    return
}

func NewDescribeIaCFileReportResponse() (response *DescribeIaCFileReportResponse) {
    response = &DescribeIaCFileReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIaCFileReport
// Retrieve the IaC detection file report
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIaCFileReport(request *DescribeIaCFileReportRequest) (response *DescribeIaCFileReportResponse, err error) {
    return c.DescribeIaCFileReportWithContext(context.Background(), request)
}

// DescribeIaCFileReport
// Retrieve the IaC detection file report
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIaCFileReportWithContext(ctx context.Context, request *DescribeIaCFileReportRequest) (response *DescribeIaCFileReportResponse, err error) {
    if request == nil {
        request = NewDescribeIaCFileReportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIaCFileReport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIaCFileReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIaCFileReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIaCTokenListRequest() (request *DescribeIaCTokenListRequest) {
    request = &DescribeIaCTokenListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeIaCTokenList")
    
    
    return
}

func NewDescribeIaCTokenListResponse() (response *DescribeIaCTokenListResponse) {
    response = &DescribeIaCTokenListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIaCTokenList
// Retrieve the IaC detection integration Token list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCTokenList(request *DescribeIaCTokenListRequest) (response *DescribeIaCTokenListResponse, err error) {
    return c.DescribeIaCTokenListWithContext(context.Background(), request)
}

// DescribeIaCTokenList
// Retrieve the IaC detection integration Token list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIaCTokenListWithContext(ctx context.Context, request *DescribeIaCTokenListRequest) (response *DescribeIaCTokenListResponse, err error) {
    if request == nil {
        request = NewDescribeIaCTokenListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeIaCTokenList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIaCTokenList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIaCTokenListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeySandboxCredentialRequest() (request *DescribeKeySandboxCredentialRequest) {
    request = &DescribeKeySandboxCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeKeySandboxCredential")
    
    
    return
}

func NewDescribeKeySandboxCredentialResponse() (response *DescribeKeySandboxCredentialResponse) {
    response = &DescribeKeySandboxCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKeySandboxCredential
// Query credential details and return credential metadata and masked credential data. The access kind returns an Access array (original Key, masked Value), and the sts kind returns an STS object (original System, masked SecretID and SecretKey).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKeySandboxCredential(request *DescribeKeySandboxCredentialRequest) (response *DescribeKeySandboxCredentialResponse, err error) {
    return c.DescribeKeySandboxCredentialWithContext(context.Background(), request)
}

// DescribeKeySandboxCredential
// Query credential details and return credential metadata and masked credential data. The access kind returns an Access array (original Key, masked Value), and the sts kind returns an STS object (original System, masked SecretID and SecretKey).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKeySandboxCredentialWithContext(ctx context.Context, request *DescribeKeySandboxCredentialRequest) (response *DescribeKeySandboxCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeKeySandboxCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeKeySandboxCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKeySandboxCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKeySandboxCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeySandboxCredentialListRequest() (request *DescribeKeySandboxCredentialListRequest) {
    request = &DescribeKeySandboxCredentialListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeKeySandboxCredentialList")
    
    
    return
}

func NewDescribeKeySandboxCredentialListResponse() (response *DescribeKeySandboxCredentialListResponse) {
    response = &DescribeKeySandboxCredentialListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKeySandboxCredentialList
// Query the voucher list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKeySandboxCredentialList(request *DescribeKeySandboxCredentialListRequest) (response *DescribeKeySandboxCredentialListResponse, err error) {
    return c.DescribeKeySandboxCredentialListWithContext(context.Background(), request)
}

// DescribeKeySandboxCredentialList
// Query the voucher list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKeySandboxCredentialListWithContext(ctx context.Context, request *DescribeKeySandboxCredentialListRequest) (response *DescribeKeySandboxCredentialListResponse, err error) {
    if request == nil {
        request = NewDescribeKeySandboxCredentialListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeKeySandboxCredentialList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKeySandboxCredentialList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKeySandboxCredentialListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerListRequest() (request *DescribeListenerListRequest) {
    request = &DescribeListenerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeListenerList")
    
    
    return
}

func NewDescribeListenerListResponse() (response *DescribeListenerListResponse) {
    response = &DescribeListenerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListenerList
// This API is used to query the list of TCP listeners.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeListenerList(request *DescribeListenerListRequest) (response *DescribeListenerListResponse, err error) {
    return c.DescribeListenerListWithContext(context.Background(), request)
}

// DescribeListenerList
// This API is used to query the list of TCP listeners.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeListenerListWithContext(ctx context.Context, request *DescribeListenerListRequest) (response *DescribeListenerListResponse, err error) {
    if request == nil {
        request = NewDescribeListenerListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeListenerList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListenerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNICAssetsRequest() (request *DescribeNICAssetsRequest) {
    request = &DescribeNICAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeNICAssets")
    
    
    return
}

func NewDescribeNICAssetsResponse() (response *DescribeNICAssetsResponse) {
    response = &DescribeNICAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNICAssets
// Obtain Network Interface Card List
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNICAssets(request *DescribeNICAssetsRequest) (response *DescribeNICAssetsResponse, err error) {
    return c.DescribeNICAssetsWithContext(context.Background(), request)
}

// DescribeNICAssets
// Obtain Network Interface Card List
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNICAssetsWithContext(ctx context.Context, request *DescribeNICAssetsRequest) (response *DescribeNICAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeNICAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeNICAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNICAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNICAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationInfoRequest() (request *DescribeOrganizationInfoRequest) {
    request = &DescribeOrganizationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeOrganizationInfo")
    
    
    return
}

func NewDescribeOrganizationInfoResponse() (response *DescribeOrganizationInfoResponse) {
    response = &DescribeOrganizationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationInfo
// Query Group Account Details
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationInfo(request *DescribeOrganizationInfoRequest) (response *DescribeOrganizationInfoResponse, err error) {
    return c.DescribeOrganizationInfoWithContext(context.Background(), request)
}

// DescribeOrganizationInfo
// Query Group Account Details
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationInfoWithContext(ctx context.Context, request *DescribeOrganizationInfoRequest) (response *DescribeOrganizationInfoResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeOrganizationInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationUserInfoRequest() (request *DescribeOrganizationUserInfoRequest) {
    request = &DescribeOrganizationUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeOrganizationUserInfo")
    
    
    return
}

func NewDescribeOrganizationUserInfoResponse() (response *DescribeOrganizationUserInfoResponse) {
    response = &DescribeOrganizationUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationUserInfo
// Query Group Account User List
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationUserInfo(request *DescribeOrganizationUserInfoRequest) (response *DescribeOrganizationUserInfoResponse, err error) {
    return c.DescribeOrganizationUserInfoWithContext(context.Background(), request)
}

// DescribeOrganizationUserInfo
// Query Group Account User List
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationUserInfoWithContext(ctx context.Context, request *DescribeOrganizationUserInfoRequest) (response *DescribeOrganizationUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationUserInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeOrganizationUserInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOtherCloudAssetsRequest() (request *DescribeOtherCloudAssetsRequest) {
    request = &DescribeOtherCloudAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeOtherCloudAssets")
    
    
    return
}

func NewDescribeOtherCloudAssetsResponse() (response *DescribeOtherCloudAssetsResponse) {
    response = &DescribeOtherCloudAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOtherCloudAssets
// Gets the list of other cloud assets
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOtherCloudAssets(request *DescribeOtherCloudAssetsRequest) (response *DescribeOtherCloudAssetsResponse, err error) {
    return c.DescribeOtherCloudAssetsWithContext(context.Background(), request)
}

// DescribeOtherCloudAssets
// Gets the list of other cloud assets
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOtherCloudAssetsWithContext(ctx context.Context, request *DescribeOtherCloudAssetsRequest) (response *DescribeOtherCloudAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeOtherCloudAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeOtherCloudAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOtherCloudAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOtherCloudAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicIpAssetsRequest() (request *DescribePublicIpAssetsRequest) {
    request = &DescribePublicIpAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribePublicIpAssets")
    
    
    return
}

func NewDescribePublicIpAssetsResponse() (response *DescribePublicIpAssetsResponse) {
    response = &DescribePublicIpAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicIpAssets
// This API is used to query the list of public IP assets.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePublicIpAssets(request *DescribePublicIpAssetsRequest) (response *DescribePublicIpAssetsResponse, err error) {
    return c.DescribePublicIpAssetsWithContext(context.Background(), request)
}

// DescribePublicIpAssets
// This API is used to query the list of public IP assets.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePublicIpAssetsWithContext(ctx context.Context, request *DescribePublicIpAssetsRequest) (response *DescribePublicIpAssetsResponse, err error) {
    if request == nil {
        request = NewDescribePublicIpAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribePublicIpAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicIpAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicIpAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryImageAssetsRequest() (request *DescribeRepositoryImageAssetsRequest) {
    request = &DescribeRepositoryImageAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRepositoryImageAssets")
    
    
    return
}

func NewDescribeRepositoryImageAssetsResponse() (response *DescribeRepositoryImageAssetsResponse) {
    response = &DescribeRepositoryImageAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRepositoryImageAssets
// Repository Image List
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRepositoryImageAssets(request *DescribeRepositoryImageAssetsRequest) (response *DescribeRepositoryImageAssetsResponse, err error) {
    return c.DescribeRepositoryImageAssetsWithContext(context.Background(), request)
}

// DescribeRepositoryImageAssets
// Repository Image List
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRepositoryImageAssetsWithContext(ctx context.Context, request *DescribeRepositoryImageAssetsRequest) (response *DescribeRepositoryImageAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryImageAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRepositoryImageAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepositoryImageAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRepositoryImageAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCallRecordRequest() (request *DescribeRiskCallRecordRequest) {
    request = &DescribeRiskCallRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCallRecord")
    
    
    return
}

func NewDescribeRiskCallRecordResponse() (response *DescribeRiskCallRecordResponse) {
    response = &DescribeRiskCallRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCallRecord
// Retrieve risk call record list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskCallRecord(request *DescribeRiskCallRecordRequest) (response *DescribeRiskCallRecordResponse, err error) {
    return c.DescribeRiskCallRecordWithContext(context.Background(), request)
}

// DescribeRiskCallRecord
// Retrieve risk call record list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskCallRecordWithContext(ctx context.Context, request *DescribeRiskCallRecordRequest) (response *DescribeRiskCallRecordResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCallRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCallRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCallRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCallRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterAssetViewCFGRiskListRequest() (request *DescribeRiskCenterAssetViewCFGRiskListRequest) {
    request = &DescribeRiskCenterAssetViewCFGRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewCFGRiskList")
    
    
    return
}

func NewDescribeRiskCenterAssetViewCFGRiskListResponse() (response *DescribeRiskCenterAssetViewCFGRiskListResponse) {
    response = &DescribeRiskCenterAssetViewCFGRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterAssetViewCFGRiskList
// This API is used to query the list of configuration risks by assets.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskCenterAssetViewCFGRiskList(request *DescribeRiskCenterAssetViewCFGRiskListRequest) (response *DescribeRiskCenterAssetViewCFGRiskListResponse, err error) {
    return c.DescribeRiskCenterAssetViewCFGRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterAssetViewCFGRiskList
// This API is used to query the list of configuration risks by assets.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskCenterAssetViewCFGRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewCFGRiskListRequest) (response *DescribeRiskCenterAssetViewCFGRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewCFGRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterAssetViewCFGRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterAssetViewCFGRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterAssetViewCFGRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterAssetViewPortRiskListRequest() (request *DescribeRiskCenterAssetViewPortRiskListRequest) {
    request = &DescribeRiskCenterAssetViewPortRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewPortRiskList")
    
    
    return
}

func NewDescribeRiskCenterAssetViewPortRiskListResponse() (response *DescribeRiskCenterAssetViewPortRiskListResponse) {
    response = &DescribeRiskCenterAssetViewPortRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterAssetViewPortRiskList
// This API is used to query the list of port risks by assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterAssetViewPortRiskList(request *DescribeRiskCenterAssetViewPortRiskListRequest) (response *DescribeRiskCenterAssetViewPortRiskListResponse, err error) {
    return c.DescribeRiskCenterAssetViewPortRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterAssetViewPortRiskList
// This API is used to query the list of port risks by assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterAssetViewPortRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewPortRiskListRequest) (response *DescribeRiskCenterAssetViewPortRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewPortRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterAssetViewPortRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterAssetViewPortRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterAssetViewPortRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterAssetViewVULRiskListRequest() (request *DescribeRiskCenterAssetViewVULRiskListRequest) {
    request = &DescribeRiskCenterAssetViewVULRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewVULRiskList")
    
    
    return
}

func NewDescribeRiskCenterAssetViewVULRiskListResponse() (response *DescribeRiskCenterAssetViewVULRiskListResponse) {
    response = &DescribeRiskCenterAssetViewVULRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterAssetViewVULRiskList
// This API is used to query the list of vulnerabilities by assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterAssetViewVULRiskList(request *DescribeRiskCenterAssetViewVULRiskListRequest) (response *DescribeRiskCenterAssetViewVULRiskListResponse, err error) {
    return c.DescribeRiskCenterAssetViewVULRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterAssetViewVULRiskList
// This API is used to query the list of vulnerabilities by assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterAssetViewVULRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewVULRiskListRequest) (response *DescribeRiskCenterAssetViewVULRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewVULRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterAssetViewVULRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterAssetViewVULRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterAssetViewVULRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterAssetViewWeakPasswordRiskListRequest() (request *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) {
    request = &DescribeRiskCenterAssetViewWeakPasswordRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewWeakPasswordRiskList")
    
    
    return
}

func NewDescribeRiskCenterAssetViewWeakPasswordRiskListResponse() (response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) {
    response = &DescribeRiskCenterAssetViewWeakPasswordRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterAssetViewWeakPasswordRiskList
// This API is used to query the list of weak passwords by assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterAssetViewWeakPasswordRiskList(request *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) (response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse, err error) {
    return c.DescribeRiskCenterAssetViewWeakPasswordRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterAssetViewWeakPasswordRiskList
// This API is used to query the list of weak passwords by assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterAssetViewWeakPasswordRiskListWithContext(ctx context.Context, request *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) (response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterAssetViewWeakPasswordRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterAssetViewWeakPasswordRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterAssetViewWeakPasswordRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterAssetViewWeakPasswordRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterCFGViewCFGRiskListRequest() (request *DescribeRiskCenterCFGViewCFGRiskListRequest) {
    request = &DescribeRiskCenterCFGViewCFGRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterCFGViewCFGRiskList")
    
    
    return
}

func NewDescribeRiskCenterCFGViewCFGRiskListResponse() (response *DescribeRiskCenterCFGViewCFGRiskListResponse) {
    response = &DescribeRiskCenterCFGViewCFGRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterCFGViewCFGRiskList
// Obtain Configuration Risk List from Configuration's Perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterCFGViewCFGRiskList(request *DescribeRiskCenterCFGViewCFGRiskListRequest) (response *DescribeRiskCenterCFGViewCFGRiskListResponse, err error) {
    return c.DescribeRiskCenterCFGViewCFGRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterCFGViewCFGRiskList
// Obtain Configuration Risk List from Configuration's Perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterCFGViewCFGRiskListWithContext(ctx context.Context, request *DescribeRiskCenterCFGViewCFGRiskListRequest) (response *DescribeRiskCenterCFGViewCFGRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterCFGViewCFGRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterCFGViewCFGRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterCFGViewCFGRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterCFGViewCFGRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterPortViewPortRiskListRequest() (request *DescribeRiskCenterPortViewPortRiskListRequest) {
    request = &DescribeRiskCenterPortViewPortRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterPortViewPortRiskList")
    
    
    return
}

func NewDescribeRiskCenterPortViewPortRiskListResponse() (response *DescribeRiskCenterPortViewPortRiskListResponse) {
    response = &DescribeRiskCenterPortViewPortRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterPortViewPortRiskList
// This API is used to query the list of port risks by ports.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterPortViewPortRiskList(request *DescribeRiskCenterPortViewPortRiskListRequest) (response *DescribeRiskCenterPortViewPortRiskListResponse, err error) {
    return c.DescribeRiskCenterPortViewPortRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterPortViewPortRiskList
// This API is used to query the list of port risks by ports.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterPortViewPortRiskListWithContext(ctx context.Context, request *DescribeRiskCenterPortViewPortRiskListRequest) (response *DescribeRiskCenterPortViewPortRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterPortViewPortRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterPortViewPortRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterPortViewPortRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterPortViewPortRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterServerRiskListRequest() (request *DescribeRiskCenterServerRiskListRequest) {
    request = &DescribeRiskCenterServerRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterServerRiskList")
    
    
    return
}

func NewDescribeRiskCenterServerRiskListResponse() (response *DescribeRiskCenterServerRiskListResponse) {
    response = &DescribeRiskCenterServerRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterServerRiskList
// This API is used to query the list of services in risk.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterServerRiskList(request *DescribeRiskCenterServerRiskListRequest) (response *DescribeRiskCenterServerRiskListResponse, err error) {
    return c.DescribeRiskCenterServerRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterServerRiskList
// This API is used to query the list of services in risk.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTERROR = "InvalidParameter.DecryptError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLOUDAUDIT = "UnauthorizedOperation.CloudAudit"
//  UNAUTHORIZEDOPERATION_COS = "UnauthorizedOperation.Cos"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskCenterServerRiskListWithContext(ctx context.Context, request *DescribeRiskCenterServerRiskListRequest) (response *DescribeRiskCenterServerRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterServerRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterServerRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterServerRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterServerRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterVULViewVULRiskListRequest() (request *DescribeRiskCenterVULViewVULRiskListRequest) {
    request = &DescribeRiskCenterVULViewVULRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterVULViewVULRiskList")
    
    
    return
}

func NewDescribeRiskCenterVULViewVULRiskListResponse() (response *DescribeRiskCenterVULViewVULRiskListResponse) {
    response = &DescribeRiskCenterVULViewVULRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterVULViewVULRiskList
// This API is used to query the list of vulnerabilities by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskCenterVULViewVULRiskList(request *DescribeRiskCenterVULViewVULRiskListRequest) (response *DescribeRiskCenterVULViewVULRiskListResponse, err error) {
    return c.DescribeRiskCenterVULViewVULRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterVULViewVULRiskList
// This API is used to query the list of vulnerabilities by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskCenterVULViewVULRiskListWithContext(ctx context.Context, request *DescribeRiskCenterVULViewVULRiskListRequest) (response *DescribeRiskCenterVULViewVULRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterVULViewVULRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterVULViewVULRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterVULViewVULRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterVULViewVULRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCenterWebsiteRiskListRequest() (request *DescribeRiskCenterWebsiteRiskListRequest) {
    request = &DescribeRiskCenterWebsiteRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterWebsiteRiskList")
    
    
    return
}

func NewDescribeRiskCenterWebsiteRiskListResponse() (response *DescribeRiskCenterWebsiteRiskListResponse) {
    response = &DescribeRiskCenterWebsiteRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCenterWebsiteRiskList
// Obtain Content Risk List
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskCenterWebsiteRiskList(request *DescribeRiskCenterWebsiteRiskListRequest) (response *DescribeRiskCenterWebsiteRiskListResponse, err error) {
    return c.DescribeRiskCenterWebsiteRiskListWithContext(context.Background(), request)
}

// DescribeRiskCenterWebsiteRiskList
// Obtain Content Risk List
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskCenterWebsiteRiskListWithContext(ctx context.Context, request *DescribeRiskCenterWebsiteRiskListRequest) (response *DescribeRiskCenterWebsiteRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCenterWebsiteRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskCenterWebsiteRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCenterWebsiteRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCenterWebsiteRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskDetailListRequest() (request *DescribeRiskDetailListRequest) {
    request = &DescribeRiskDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskDetailList")
    
    
    return
}

func NewDescribeRiskDetailListResponse() (response *DescribeRiskDetailListResponse) {
    response = &DescribeRiskDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskDetailList
// Risk detail list example
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskDetailList(request *DescribeRiskDetailListRequest) (response *DescribeRiskDetailListResponse, err error) {
    return c.DescribeRiskDetailListWithContext(context.Background(), request)
}

// DescribeRiskDetailList
// Risk detail list example
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskDetailListWithContext(ctx context.Context, request *DescribeRiskDetailListRequest) (response *DescribeRiskDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskDetailListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskDetailList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskRuleDetailRequest() (request *DescribeRiskRuleDetailRequest) {
    request = &DescribeRiskRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskRuleDetail")
    
    
    return
}

func NewDescribeRiskRuleDetailResponse() (response *DescribeRiskRuleDetailResponse) {
    response = &DescribeRiskRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskRuleDetail
// Query risk rule details example
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskRuleDetail(request *DescribeRiskRuleDetailRequest) (response *DescribeRiskRuleDetailResponse, err error) {
    return c.DescribeRiskRuleDetailWithContext(context.Background(), request)
}

// DescribeRiskRuleDetail
// Query risk rule details example
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskRuleDetailWithContext(ctx context.Context, request *DescribeRiskRuleDetailRequest) (response *DescribeRiskRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRiskRuleDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskRuleDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskRulesRequest() (request *DescribeRiskRulesRequest) {
    request = &DescribeRiskRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskRules")
    
    
    return
}

func NewDescribeRiskRulesResponse() (response *DescribeRiskRulesResponse) {
    response = &DescribeRiskRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskRules
// Advanced configuration risk rule list illustrative example
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskRules(request *DescribeRiskRulesRequest) (response *DescribeRiskRulesResponse, err error) {
    return c.DescribeRiskRulesWithContext(context.Background(), request)
}

// DescribeRiskRules
// Advanced configuration risk rule list illustrative example
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeRiskRulesWithContext(ctx context.Context, request *DescribeRiskRulesRequest) (response *DescribeRiskRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRiskRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeRiskRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanReportListRequest() (request *DescribeScanReportListRequest) {
    request = &DescribeScanReportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeScanReportList")
    
    
    return
}

func NewDescribeScanReportListResponse() (response *DescribeScanReportListResponse) {
    response = &DescribeScanReportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScanReportList
// This API is used to get the list of scan reports.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeScanReportList(request *DescribeScanReportListRequest) (response *DescribeScanReportListResponse, err error) {
    return c.DescribeScanReportListWithContext(context.Background(), request)
}

// DescribeScanReportList
// This API is used to get the list of scan reports.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeScanReportListWithContext(ctx context.Context, request *DescribeScanReportListRequest) (response *DescribeScanReportListResponse, err error) {
    if request == nil {
        request = NewDescribeScanReportListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeScanReportList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanReportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanReportListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanStatisticRequest() (request *DescribeScanStatisticRequest) {
    request = &DescribeScanStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeScanStatistic")
    
    
    return
}

func NewDescribeScanStatisticResponse() (response *DescribeScanStatisticResponse) {
    response = &DescribeScanStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScanStatistic
// Query the statistical information of cloud boundary analysis scanning results.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeScanStatistic(request *DescribeScanStatisticRequest) (response *DescribeScanStatisticResponse, err error) {
    return c.DescribeScanStatisticWithContext(context.Background(), request)
}

// DescribeScanStatistic
// Query the statistical information of cloud boundary analysis scanning results.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeScanStatisticWithContext(ctx context.Context, request *DescribeScanStatisticRequest) (response *DescribeScanStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeScanStatisticRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeScanStatistic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanTaskListRequest() (request *DescribeScanTaskListRequest) {
    request = &DescribeScanTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeScanTaskList")
    
    
    return
}

func NewDescribeScanTaskListResponse() (response *DescribeScanTaskListResponse) {
    response = &DescribeScanTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScanTaskList
// This API is used to get the list of scan tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeScanTaskList(request *DescribeScanTaskListRequest) (response *DescribeScanTaskListResponse, err error) {
    return c.DescribeScanTaskListWithContext(context.Background(), request)
}

// DescribeScanTaskList
// This API is used to get the list of scan tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeScanTaskListWithContext(ctx context.Context, request *DescribeScanTaskListRequest) (response *DescribeScanTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeScanTaskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeScanTaskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchBugInfoRequest() (request *DescribeSearchBugInfoRequest) {
    request = &DescribeSearchBugInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSearchBugInfo")
    
    
    return
}

func NewDescribeSearchBugInfoResponse() (response *DescribeSearchBugInfoResponse) {
    response = &DescribeSearchBugInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSearchBugInfo
// This API is used to query information of a vulnerability.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSearchBugInfo(request *DescribeSearchBugInfoRequest) (response *DescribeSearchBugInfoResponse, err error) {
    return c.DescribeSearchBugInfoWithContext(context.Background(), request)
}

// DescribeSearchBugInfo
// This API is used to query information of a vulnerability.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSearchBugInfoWithContext(ctx context.Context, request *DescribeSearchBugInfoRequest) (response *DescribeSearchBugInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSearchBugInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSearchBugInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchBugInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchBugInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillScanPayInfoRequest() (request *DescribeSkillScanPayInfoRequest) {
    request = &DescribeSkillScanPayInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSkillScanPayInfo")
    
    
    return
}

func NewDescribeSkillScanPayInfoResponse() (response *DescribeSkillScanPayInfoResponse) {
    response = &DescribeSkillScanPayInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillScanPayInfo
// This API is used to query the billing information of Skill security detection, including order status, total quota, consumed quota, expiry time, and payment mode. It returns zero values (only TimeNow and BetaEndTime included) when no order exists. Trial orders can be claimed via ModifyTrialStatus(Module=9), and official orders are generated by the billing system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSkillScanPayInfo(request *DescribeSkillScanPayInfoRequest) (response *DescribeSkillScanPayInfoResponse, err error) {
    return c.DescribeSkillScanPayInfoWithContext(context.Background(), request)
}

// DescribeSkillScanPayInfo
// This API is used to query the billing information of Skill security detection, including order status, total quota, consumed quota, expiry time, and payment mode. It returns zero values (only TimeNow and BetaEndTime included) when no order exists. Trial orders can be claimed via ModifyTrialStatus(Module=9), and official orders are generated by the billing system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSkillScanPayInfoWithContext(ctx context.Context, request *DescribeSkillScanPayInfoRequest) (response *DescribeSkillScanPayInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSkillScanPayInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSkillScanPayInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillScanPayInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillScanPayInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillScanResultRequest() (request *DescribeSkillScanResultRequest) {
    request = &DescribeSkillScanResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSkillScanResult")
    
    
    return
}

func NewDescribeSkillScanResultResponse() (response *DescribeSkillScanResultResponse) {
    response = &DescribeSkillScanResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillScanResult
// Query the security detection result of Skill. After successfully calling CreateSkillScan, use the returned ContentHash + EngineVersion to poll this interface and obtain the result. For the first time, poll 5 minutes after successful upload. If detection is not completed, poll every 1 minute afterward. The response distinguishes four types of status via the Status field: detection complete (SUCCESS), in-progress detection (SCANNING), no record (NOT_FOUND), and detection failed (FAILED). Note: The detection result is retained for 90 days. It will return NOT_FOUND when overdue.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSkillScanResult(request *DescribeSkillScanResultRequest) (response *DescribeSkillScanResultResponse, err error) {
    return c.DescribeSkillScanResultWithContext(context.Background(), request)
}

// DescribeSkillScanResult
// Query the security detection result of Skill. After successfully calling CreateSkillScan, use the returned ContentHash + EngineVersion to poll this interface and obtain the result. For the first time, poll 5 minutes after successful upload. If detection is not completed, poll every 1 minute afterward. The response distinguishes four types of status via the Status field: detection complete (SUCCESS), in-progress detection (SCANNING), no record (NOT_FOUND), and detection failed (FAILED). Note: The detection result is retained for 90 days. It will return NOT_FOUND when overdue.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSkillScanResultWithContext(ctx context.Context, request *DescribeSkillScanResultRequest) (response *DescribeSkillScanResultResponse, err error) {
    if request == nil {
        request = NewDescribeSkillScanResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSkillScanResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillScanResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillScanResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceIPAssetRequest() (request *DescribeSourceIPAssetRequest) {
    request = &DescribeSourceIPAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSourceIPAsset")
    
    
    return
}

func NewDescribeSourceIPAssetResponse() (response *DescribeSourceIPAssetResponse) {
    response = &DescribeSourceIPAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSourceIPAsset
// This API is used to obtain user access key asset list (source IP perspective).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSourceIPAsset(request *DescribeSourceIPAssetRequest) (response *DescribeSourceIPAssetResponse, err error) {
    return c.DescribeSourceIPAssetWithContext(context.Background(), request)
}

// DescribeSourceIPAsset
// This API is used to obtain user access key asset list (source IP perspective).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSourceIPAssetWithContext(ctx context.Context, request *DescribeSourceIPAssetRequest) (response *DescribeSourceIPAssetResponse, err error) {
    if request == nil {
        request = NewDescribeSourceIPAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSourceIPAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSourceIPAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSourceIPAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubUserInfoRequest() (request *DescribeSubUserInfoRequest) {
    request = &DescribeSubUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSubUserInfo")
    
    
    return
}

func NewDescribeSubUserInfoResponse() (response *DescribeSubUserInfoResponse) {
    response = &DescribeSubUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubUserInfo
// Query the sub-account list of the group
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSubUserInfo(request *DescribeSubUserInfoRequest) (response *DescribeSubUserInfoResponse, err error) {
    return c.DescribeSubUserInfoWithContext(context.Background(), request)
}

// DescribeSubUserInfo
// Query the sub-account list of the group
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSubUserInfoWithContext(ctx context.Context, request *DescribeSubUserInfoRequest) (response *DescribeSubUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSubUserInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSubUserInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetAssetsRequest() (request *DescribeSubnetAssetsRequest) {
    request = &DescribeSubnetAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeSubnetAssets")
    
    
    return
}

func NewDescribeSubnetAssetsResponse() (response *DescribeSubnetAssetsResponse) {
    response = &DescribeSubnetAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubnetAssets
// This API is used to get the list of subnets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSubnetAssets(request *DescribeSubnetAssetsRequest) (response *DescribeSubnetAssetsResponse, err error) {
    return c.DescribeSubnetAssetsWithContext(context.Background(), request)
}

// DescribeSubnetAssets
// This API is used to get the list of subnets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeSubnetAssetsWithContext(ctx context.Context, request *DescribeSubnetAssetsRequest) (response *DescribeSubnetAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeSubnetAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubnetAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubnetAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLogListRequest() (request *DescribeTaskLogListRequest) {
    request = &DescribeTaskLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskLogList")
    
    
    return
}

func NewDescribeTaskLogListResponse() (response *DescribeTaskLogListResponse) {
    response = &DescribeTaskLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskLogList
// This API is used to get the list of scan task reports.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeTaskLogList(request *DescribeTaskLogListRequest) (response *DescribeTaskLogListResponse, err error) {
    return c.DescribeTaskLogListWithContext(context.Background(), request)
}

// DescribeTaskLogList
// This API is used to get the list of scan task reports.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeTaskLogListWithContext(ctx context.Context, request *DescribeTaskLogListRequest) (response *DescribeTaskLogListResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeTaskLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLogURLRequest() (request *DescribeTaskLogURLRequest) {
    request = &DescribeTaskLogURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskLogURL")
    
    
    return
}

func NewDescribeTaskLogURLResponse() (response *DescribeTaskLogURLResponse) {
    response = &DescribeTaskLogURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskLogURL
// This API is used to get the temp download link of a report. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeTaskLogURL(request *DescribeTaskLogURLRequest) (response *DescribeTaskLogURLResponse, err error) {
    return c.DescribeTaskLogURLWithContext(context.Background(), request)
}

// DescribeTaskLogURL
// This API is used to get the temp download link of a report. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeTaskLogURLWithContext(ctx context.Context, request *DescribeTaskLogURLRequest) (response *DescribeTaskLogURLResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLogURLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeTaskLogURL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskLogURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskLogURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopAttackInfoRequest() (request *DescribeTopAttackInfoRequest) {
    request = &DescribeTopAttackInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeTopAttackInfo")
    
    
    return
}

func NewDescribeTopAttackInfoResponse() (response *DescribeTopAttackInfoResponse) {
    response = &DescribeTopAttackInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopAttackInfo
// Query TOP attack information
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeTopAttackInfo(request *DescribeTopAttackInfoRequest) (response *DescribeTopAttackInfoResponse, err error) {
    return c.DescribeTopAttackInfoWithContext(context.Background(), request)
}

// DescribeTopAttackInfo
// Query TOP attack information
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeTopAttackInfoWithContext(ctx context.Context, request *DescribeTopAttackInfoRequest) (response *DescribeTopAttackInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTopAttackInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeTopAttackInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopAttackInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopAttackInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUebaRuleRequest() (request *DescribeUebaRuleRequest) {
    request = &DescribeUebaRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeUebaRule")
    
    
    return
}

func NewDescribeUebaRuleResponse() (response *DescribeUebaRuleResponse) {
    response = &DescribeUebaRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUebaRule
// Query the user behavior analysis policy list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeUebaRule(request *DescribeUebaRuleRequest) (response *DescribeUebaRuleResponse, err error) {
    return c.DescribeUebaRuleWithContext(context.Background(), request)
}

// DescribeUebaRule
// Query the user behavior analysis policy list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeUebaRuleWithContext(ctx context.Context, request *DescribeUebaRuleRequest) (response *DescribeUebaRuleResponse, err error) {
    if request == nil {
        request = NewDescribeUebaRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeUebaRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUebaRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUebaRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserCallRecordRequest() (request *DescribeUserCallRecordRequest) {
    request = &DescribeUserCallRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeUserCallRecord")
    
    
    return
}

func NewDescribeUserCallRecordResponse() (response *DescribeUserCallRecordResponse) {
    response = &DescribeUserCallRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserCallRecord
// Obtain account call record list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeUserCallRecord(request *DescribeUserCallRecordRequest) (response *DescribeUserCallRecordResponse, err error) {
    return c.DescribeUserCallRecordWithContext(context.Background(), request)
}

// DescribeUserCallRecord
// Obtain account call record list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeUserCallRecordWithContext(ctx context.Context, request *DescribeUserCallRecordRequest) (response *DescribeUserCallRecordResponse, err error) {
    if request == nil {
        request = NewDescribeUserCallRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeUserCallRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserCallRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserCallRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVULListRequest() (request *DescribeVULListRequest) {
    request = &DescribeVULListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVULList")
    
    
    return
}

func NewDescribeVULListResponse() (response *DescribeVULListResponse) {
    response = &DescribeVULListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVULList
// Security Center Risk Center - List of Vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVULList(request *DescribeVULListRequest) (response *DescribeVULListResponse, err error) {
    return c.DescribeVULListWithContext(context.Background(), request)
}

// DescribeVULList
// Security Center Risk Center - List of Vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVULListWithContext(ctx context.Context, request *DescribeVULListRequest) (response *DescribeVULListResponse, err error) {
    if request == nil {
        request = NewDescribeVULListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVULList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVULList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVULListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVULRiskAdvanceCFGListRequest() (request *DescribeVULRiskAdvanceCFGListRequest) {
    request = &DescribeVULRiskAdvanceCFGListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVULRiskAdvanceCFGList")
    
    
    return
}

func NewDescribeVULRiskAdvanceCFGListResponse() (response *DescribeVULRiskAdvanceCFGListResponse) {
    response = &DescribeVULRiskAdvanceCFGListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVULRiskAdvanceCFGList
// This API is used to query the advanced configuration of vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVULRiskAdvanceCFGList(request *DescribeVULRiskAdvanceCFGListRequest) (response *DescribeVULRiskAdvanceCFGListResponse, err error) {
    return c.DescribeVULRiskAdvanceCFGListWithContext(context.Background(), request)
}

// DescribeVULRiskAdvanceCFGList
// This API is used to query the advanced configuration of vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVULRiskAdvanceCFGListWithContext(ctx context.Context, request *DescribeVULRiskAdvanceCFGListRequest) (response *DescribeVULRiskAdvanceCFGListResponse, err error) {
    if request == nil {
        request = NewDescribeVULRiskAdvanceCFGListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVULRiskAdvanceCFGList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVULRiskAdvanceCFGList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVULRiskAdvanceCFGListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVULRiskDetailRequest() (request *DescribeVULRiskDetailRequest) {
    request = &DescribeVULRiskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVULRiskDetail")
    
    
    return
}

func NewDescribeVULRiskDetailResponse() (response *DescribeVULRiskDetailResponse) {
    response = &DescribeVULRiskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVULRiskDetail
// Retrieve vulnerability details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVULRiskDetail(request *DescribeVULRiskDetailRequest) (response *DescribeVULRiskDetailResponse, err error) {
    return c.DescribeVULRiskDetailWithContext(context.Background(), request)
}

// DescribeVULRiskDetail
// Retrieve vulnerability details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVULRiskDetailWithContext(ctx context.Context, request *DescribeVULRiskDetailRequest) (response *DescribeVULRiskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVULRiskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVULRiskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVULRiskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVULRiskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcAssetsRequest() (request *DescribeVpcAssetsRequest) {
    request = &DescribeVpcAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVpcAssets")
    
    
    return
}

func NewDescribeVpcAssetsResponse() (response *DescribeVpcAssetsResponse) {
    response = &DescribeVpcAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVpcAssets
// This API is used to get the list of VPCs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVpcAssets(request *DescribeVpcAssetsRequest) (response *DescribeVpcAssetsResponse, err error) {
    return c.DescribeVpcAssetsWithContext(context.Background(), request)
}

// DescribeVpcAssets
// This API is used to get the list of VPCs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVpcAssetsWithContext(ctx context.Context, request *DescribeVpcAssetsRequest) (response *DescribeVpcAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcAssetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVpcAssets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulRiskListRequest() (request *DescribeVulRiskListRequest) {
    request = &DescribeVulRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVulRiskList")
    
    
    return
}

func NewDescribeVulRiskListResponse() (response *DescribeVulRiskListResponse) {
    response = &DescribeVulRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulRiskList
// Query the list of vulnerabilities of host nodes under the exposed path in cloud boundary analysis.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVulRiskList(request *DescribeVulRiskListRequest) (response *DescribeVulRiskListResponse, err error) {
    return c.DescribeVulRiskListWithContext(context.Background(), request)
}

// DescribeVulRiskList
// Query the list of vulnerabilities of host nodes under the exposed path in cloud boundary analysis.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVulRiskListWithContext(ctx context.Context, request *DescribeVulRiskListRequest) (response *DescribeVulRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeVulRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVulRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulViewVulRiskListRequest() (request *DescribeVulViewVulRiskListRequest) {
    request = &DescribeVulViewVulRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "DescribeVulViewVulRiskList")
    
    
    return
}

func NewDescribeVulViewVulRiskListResponse() (response *DescribeVulViewVulRiskListResponse) {
    response = &DescribeVulViewVulRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulViewVulRiskList
// Obtain Vulnerability Risk List from Vulnerability's Perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVulViewVulRiskList(request *DescribeVulViewVulRiskListRequest) (response *DescribeVulViewVulRiskListResponse, err error) {
    return c.DescribeVulViewVulRiskListWithContext(context.Background(), request)
}

// DescribeVulViewVulRiskList
// Obtain Vulnerability Risk List from Vulnerability's Perspective
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeVulViewVulRiskListWithContext(ctx context.Context, request *DescribeVulViewVulRiskListRequest) (response *DescribeVulViewVulRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeVulViewVulRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "DescribeVulViewVulRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulViewVulRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulViewVulRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIaCTokenPeriodRequest() (request *ModifyIaCTokenPeriodRequest) {
    request = &ModifyIaCTokenPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyIaCTokenPeriod")
    
    
    return
}

func NewModifyIaCTokenPeriodResponse() (response *ModifyIaCTokenPeriodResponse) {
    response = &ModifyIaCTokenPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIaCTokenPeriod
// This API is used to modify the storage cycle of the IaC detection integration Token.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyIaCTokenPeriod(request *ModifyIaCTokenPeriodRequest) (response *ModifyIaCTokenPeriodResponse, err error) {
    return c.ModifyIaCTokenPeriodWithContext(context.Background(), request)
}

// ModifyIaCTokenPeriod
// This API is used to modify the storage cycle of the IaC detection integration Token.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyIaCTokenPeriodWithContext(ctx context.Context, request *ModifyIaCTokenPeriodRequest) (response *ModifyIaCTokenPeriodResponse, err error) {
    if request == nil {
        request = NewModifyIaCTokenPeriodRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyIaCTokenPeriod")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIaCTokenPeriod require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIaCTokenPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOrganizationAccountStatusRequest() (request *ModifyOrganizationAccountStatusRequest) {
    request = &ModifyOrganizationAccountStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyOrganizationAccountStatus")
    
    
    return
}

func NewModifyOrganizationAccountStatusResponse() (response *ModifyOrganizationAccountStatusResponse) {
    response = &ModifyOrganizationAccountStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOrganizationAccountStatus
// Modify group account status
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOrganizationAccountStatus(request *ModifyOrganizationAccountStatusRequest) (response *ModifyOrganizationAccountStatusResponse, err error) {
    return c.ModifyOrganizationAccountStatusWithContext(context.Background(), request)
}

// ModifyOrganizationAccountStatus
// Modify group account status
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOrganizationAccountStatusWithContext(ctx context.Context, request *ModifyOrganizationAccountStatusRequest) (response *ModifyOrganizationAccountStatusResponse, err error) {
    if request == nil {
        request = NewModifyOrganizationAccountStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyOrganizationAccountStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOrganizationAccountStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOrganizationAccountStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRiskCenterRiskStatusRequest() (request *ModifyRiskCenterRiskStatusRequest) {
    request = &ModifyRiskCenterRiskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyRiskCenterRiskStatus")
    
    
    return
}

func NewModifyRiskCenterRiskStatusResponse() (response *ModifyRiskCenterRiskStatusResponse) {
    response = &ModifyRiskCenterRiskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRiskCenterRiskStatus
// This API is used to modify the status of a risk. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) ModifyRiskCenterRiskStatus(request *ModifyRiskCenterRiskStatusRequest) (response *ModifyRiskCenterRiskStatusResponse, err error) {
    return c.ModifyRiskCenterRiskStatusWithContext(context.Background(), request)
}

// ModifyRiskCenterRiskStatus
// This API is used to modify the status of a risk. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) ModifyRiskCenterRiskStatusWithContext(ctx context.Context, request *ModifyRiskCenterRiskStatusRequest) (response *ModifyRiskCenterRiskStatusResponse, err error) {
    if request == nil {
        request = NewModifyRiskCenterRiskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyRiskCenterRiskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskCenterRiskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskCenterRiskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRiskCenterScanTaskRequest() (request *ModifyRiskCenterScanTaskRequest) {
    request = &ModifyRiskCenterScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyRiskCenterScanTask")
    
    
    return
}

func NewModifyRiskCenterScanTaskResponse() (response *ModifyRiskCenterScanTaskResponse) {
    response = &ModifyRiskCenterScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRiskCenterScanTask
// Modify Risk Center Scan Task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) ModifyRiskCenterScanTask(request *ModifyRiskCenterScanTaskRequest) (response *ModifyRiskCenterScanTaskResponse, err error) {
    return c.ModifyRiskCenterScanTaskWithContext(context.Background(), request)
}

// ModifyRiskCenterScanTask
// Modify Risk Center Scan Task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) ModifyRiskCenterScanTaskWithContext(ctx context.Context, request *ModifyRiskCenterScanTaskRequest) (response *ModifyRiskCenterScanTaskResponse, err error) {
    if request == nil {
        request = NewModifyRiskCenterScanTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyRiskCenterScanTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskCenterScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskCenterScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUebaRuleSwitchRequest() (request *ModifyUebaRuleSwitchRequest) {
    request = &ModifyUebaRuleSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "ModifyUebaRuleSwitch")
    
    
    return
}

func NewModifyUebaRuleSwitchResponse() (response *ModifyUebaRuleSwitchResponse) {
    response = &ModifyUebaRuleSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUebaRuleSwitch
// Refresh the custom policy switch
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) ModifyUebaRuleSwitch(request *ModifyUebaRuleSwitchRequest) (response *ModifyUebaRuleSwitchResponse, err error) {
    return c.ModifyUebaRuleSwitchWithContext(context.Background(), request)
}

// ModifyUebaRuleSwitch
// Refresh the custom policy switch
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) ModifyUebaRuleSwitchWithContext(ctx context.Context, request *ModifyUebaRuleSwitchRequest) (response *ModifyUebaRuleSwitchResponse, err error) {
    if request == nil {
        request = NewModifyUebaRuleSwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "ModifyUebaRuleSwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUebaRuleSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUebaRuleSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewStopRiskCenterTaskRequest() (request *StopRiskCenterTaskRequest) {
    request = &StopRiskCenterTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "StopRiskCenterTask")
    
    
    return
}

func NewStopRiskCenterTaskResponse() (response *StopRiskCenterTaskResponse) {
    response = &StopRiskCenterTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopRiskCenterTask
// This API is used to stop a scan task. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) StopRiskCenterTask(request *StopRiskCenterTaskRequest) (response *StopRiskCenterTaskResponse, err error) {
    return c.StopRiskCenterTaskWithContext(context.Background(), request)
}

// StopRiskCenterTask
// This API is used to stop a scan task. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) StopRiskCenterTaskWithContext(ctx context.Context, request *StopRiskCenterTaskRequest) (response *StopRiskCenterTaskResponse, err error) {
    if request == nil {
        request = NewStopRiskCenterTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "StopRiskCenterTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRiskCenterTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRiskCenterTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAccessKeyAlarmStatusRequest() (request *UpdateAccessKeyAlarmStatusRequest) {
    request = &UpdateAccessKeyAlarmStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "UpdateAccessKeyAlarmStatus")
    
    
    return
}

func NewUpdateAccessKeyAlarmStatusResponse() (response *UpdateAccessKeyAlarmStatusResponse) {
    response = &UpdateAccessKeyAlarmStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAccessKeyAlarmStatus
// Tag the risk or Alarm as processed/ignored.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) UpdateAccessKeyAlarmStatus(request *UpdateAccessKeyAlarmStatusRequest) (response *UpdateAccessKeyAlarmStatusResponse, err error) {
    return c.UpdateAccessKeyAlarmStatusWithContext(context.Background(), request)
}

// UpdateAccessKeyAlarmStatus
// Tag the risk or Alarm as processed/ignored.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) UpdateAccessKeyAlarmStatusWithContext(ctx context.Context, request *UpdateAccessKeyAlarmStatusRequest) (response *UpdateAccessKeyAlarmStatusResponse, err error) {
    if request == nil {
        request = NewUpdateAccessKeyAlarmStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "UpdateAccessKeyAlarmStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAccessKeyAlarmStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAccessKeyAlarmStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAccessKeyRemarkRequest() (request *UpdateAccessKeyRemarkRequest) {
    request = &UpdateAccessKeyRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "UpdateAccessKeyRemark")
    
    
    return
}

func NewUpdateAccessKeyRemarkResponse() (response *UpdateAccessKeyRemarkResponse) {
    response = &UpdateAccessKeyRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAccessKeyRemark
// Edit access key/Source IP remark.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) UpdateAccessKeyRemark(request *UpdateAccessKeyRemarkRequest) (response *UpdateAccessKeyRemarkResponse, err error) {
    return c.UpdateAccessKeyRemarkWithContext(context.Background(), request)
}

// UpdateAccessKeyRemark
// Edit access key/Source IP remark.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) UpdateAccessKeyRemarkWithContext(ctx context.Context, request *UpdateAccessKeyRemarkRequest) (response *UpdateAccessKeyRemarkResponse, err error) {
    if request == nil {
        request = NewUpdateAccessKeyRemarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "UpdateAccessKeyRemark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAccessKeyRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAccessKeyRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAlertStatusListRequest() (request *UpdateAlertStatusListRequest) {
    request = &UpdateAlertStatusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csip", APIVersion, "UpdateAlertStatusList")
    
    
    return
}

func NewUpdateAlertStatusListResponse() (response *UpdateAlertStatusListResponse) {
    response = &UpdateAlertStatusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAlertStatusList
// Batch alarm status process API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) UpdateAlertStatusList(request *UpdateAlertStatusListRequest) (response *UpdateAlertStatusListResponse, err error) {
    return c.UpdateAlertStatusListWithContext(context.Background(), request)
}

// UpdateAlertStatusList
// Batch alarm status process API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) UpdateAlertStatusListWithContext(ctx context.Context, request *UpdateAlertStatusListRequest) (response *UpdateAlertStatusListResponse, err error) {
    if request == nil {
        request = NewUpdateAlertStatusListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "csip", APIVersion, "UpdateAlertStatusList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAlertStatusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAlertStatusListResponse()
    err = c.Send(request, response)
    return
}
