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

package v20260715

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2026-07-15"

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


func NewAddConsoleUsersRequest() (request *AddConsoleUsersRequest) {
    request = &AddConsoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "AddConsoleUsers")
    
    
    return
}

func NewAddConsoleUsersResponse() (response *AddConsoleUsersResponse) {
    response = &AddConsoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddConsoleUsers
// Add a console user
func (c *Client) AddConsoleUsers(request *AddConsoleUsersRequest) (response *AddConsoleUsersResponse, err error) {
    return c.AddConsoleUsersWithContext(context.Background(), request)
}

// AddConsoleUsers
// Add a console user
func (c *Client) AddConsoleUsersWithContext(ctx context.Context, request *AddConsoleUsersRequest) (response *AddConsoleUsersResponse, err error) {
    if request == nil {
        request = NewAddConsoleUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "AddConsoleUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddConsoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddConsoleUsersResponse()
    err = c.Send(request, response)
    return
}

func NewAddWorkspaceGroupsRequest() (request *AddWorkspaceGroupsRequest) {
    request = &AddWorkspaceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "AddWorkspaceGroups")
    
    
    return
}

func NewAddWorkspaceGroupsResponse() (response *AddWorkspaceGroupsResponse) {
    response = &AddWorkspaceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddWorkspaceGroups
// This API is used to add a user group to a workspace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_TENCENTCLOUDSDKERROR = "FailedOperation.TencentCloudSDKError"
//  INTERNALERROR_INTERFACECALLERROR = "InternalError.InterfaceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JWTAUTHFAILURE = "InvalidParameter.JWTAuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_WORKSPACEGROUPMEMBERQUOTA = "LimitExceeded.WorkspaceGroupMemberQuota"
//  LIMITEXCEEDED_WORKSPACEGROUPQUOTA = "LimitExceeded.WorkspaceGroupQuota"
//  LIMITEXCEEDED_WORKSPACEUSERQUOTA = "LimitExceeded.WorkspaceUserQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddWorkspaceGroups(request *AddWorkspaceGroupsRequest) (response *AddWorkspaceGroupsResponse, err error) {
    return c.AddWorkspaceGroupsWithContext(context.Background(), request)
}

// AddWorkspaceGroups
// This API is used to add a user group to a workspace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_TENCENTCLOUDSDKERROR = "FailedOperation.TencentCloudSDKError"
//  INTERNALERROR_INTERFACECALLERROR = "InternalError.InterfaceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JWTAUTHFAILURE = "InvalidParameter.JWTAuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_WORKSPACEGROUPMEMBERQUOTA = "LimitExceeded.WorkspaceGroupMemberQuota"
//  LIMITEXCEEDED_WORKSPACEGROUPQUOTA = "LimitExceeded.WorkspaceGroupQuota"
//  LIMITEXCEEDED_WORKSPACEUSERQUOTA = "LimitExceeded.WorkspaceUserQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddWorkspaceGroupsWithContext(ctx context.Context, request *AddWorkspaceGroupsRequest) (response *AddWorkspaceGroupsResponse, err error) {
    if request == nil {
        request = NewAddWorkspaceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "AddWorkspaceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddWorkspaceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddWorkspaceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewAddWorkspaceUsersRequest() (request *AddWorkspaceUsersRequest) {
    request = &AddWorkspaceUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "AddWorkspaceUsers")
    
    
    return
}

func NewAddWorkspaceUsersResponse() (response *AddWorkspaceUsersResponse) {
    response = &AddWorkspaceUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddWorkspaceUsers
// This API is used to create user groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_TENCENTCLOUDSDKERROR = "FailedOperation.TencentCloudSDKError"
//  INTERNALERROR_INTERFACECALLERROR = "InternalError.InterfaceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JWTAUTHFAILURE = "InvalidParameter.JWTAuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_WORKSPACEUSERQUOTA = "LimitExceeded.WorkspaceUserQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddWorkspaceUsers(request *AddWorkspaceUsersRequest) (response *AddWorkspaceUsersResponse, err error) {
    return c.AddWorkspaceUsersWithContext(context.Background(), request)
}

// AddWorkspaceUsers
// This API is used to create user groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_TENCENTCLOUDSDKERROR = "FailedOperation.TencentCloudSDKError"
//  INTERNALERROR_INTERFACECALLERROR = "InternalError.InterfaceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JWTAUTHFAILURE = "InvalidParameter.JWTAuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_WORKSPACEUSERQUOTA = "LimitExceeded.WorkspaceUserQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddWorkspaceUsersWithContext(ctx context.Context, request *AddWorkspaceUsersRequest) (response *AddWorkspaceUsersResponse, err error) {
    if request == nil {
        request = NewAddWorkspaceUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "AddWorkspaceUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddWorkspaceUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddWorkspaceUsersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCatalogRequest() (request *CreateCatalogRequest) {
    request = &CreateCatalogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateCatalog")
    
    
    return
}

func NewCreateCatalogResponse() (response *CreateCatalogResponse) {
    response = &CreateCatalogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCatalog
// This API is used to create a data catalog.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateCatalog(request *CreateCatalogRequest) (response *CreateCatalogResponse, err error) {
    return c.CreateCatalogWithContext(context.Background(), request)
}

// CreateCatalog
// This API is used to create a data catalog.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateCatalogWithContext(ctx context.Context, request *CreateCatalogRequest) (response *CreateCatalogResponse, err error) {
    if request == nil {
        request = NewCreateCatalogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateCatalog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCatalog require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCatalogResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsoleGroupRequest() (request *CreateConsoleGroupRequest) {
    request = &CreateConsoleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateConsoleGroup")
    
    
    return
}

func NewCreateConsoleGroupResponse() (response *CreateConsoleGroupResponse) {
    response = &CreateConsoleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsoleGroup
// Create a console user group
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateConsoleGroup(request *CreateConsoleGroupRequest) (response *CreateConsoleGroupResponse, err error) {
    return c.CreateConsoleGroupWithContext(context.Background(), request)
}

// CreateConsoleGroup
// Create a console user group
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateConsoleGroupWithContext(ctx context.Context, request *CreateConsoleGroupRequest) (response *CreateConsoleGroupResponse, err error) {
    if request == nil {
        request = NewCreateConsoleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateConsoleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsoleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsoleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileRequest() (request *CreateFileRequest) {
    request = &CreateFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateFile")
    
    
    return
}

func NewCreateFileResponse() (response *CreateFileResponse) {
    response = &CreateFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFile
// Create a new file (Notebook/SQL/Python, etc.) in the workspace file tree of Studio (unified development IDE), and return the complete meta-information of the file upon success.
//
// 
//
// **Prerequisite**
//
// 1. WorkspaceId corresponds to an existent workspace, and the caller is a space member of this workspace;
//
// 2. The parent folder corresponding to ParentFolderPath must exist, and the caller must have write permission on it (pass `/` for the root directory);
//
// 3. FileName must be unique under the same parent folder, including suffix comparison.
//
// 4. The FileName suffix must be consistent with FileType (`.ipynb`↔`NOTEBOOK_FILE`, `.sql`↔`SQL_FILE`);
//
// 5. To create with file content, pass it through Storage (large files are transferred via COS, and small files are placed in Storage.Content).
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateFile(request *CreateFileRequest) (response *CreateFileResponse, err error) {
    return c.CreateFileWithContext(context.Background(), request)
}

// CreateFile
// Create a new file (Notebook/SQL/Python, etc.) in the workspace file tree of Studio (unified development IDE), and return the complete meta-information of the file upon success.
//
// 
//
// **Prerequisite**
//
// 1. WorkspaceId corresponds to an existent workspace, and the caller is a space member of this workspace;
//
// 2. The parent folder corresponding to ParentFolderPath must exist, and the caller must have write permission on it (pass `/` for the root directory);
//
// 3. FileName must be unique under the same parent folder, including suffix comparison.
//
// 4. The FileName suffix must be consistent with FileType (`.ipynb`↔`NOTEBOOK_FILE`, `.sql`↔`SQL_FILE`);
//
// 5. To create with file content, pass it through Storage (large files are transferred via COS, and small files are placed in Storage.Content).
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateFileWithContext(ctx context.Context, request *CreateFileRequest) (response *CreateFileResponse, err error) {
    if request == nil {
        request = NewCreateFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFolderRequest() (request *CreateFolderRequest) {
    request = &CreateFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateFolder")
    
    
    return
}

func NewCreateFolderResponse() (response *CreateFolderResponse) {
    response = &CreateFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFolder
// Create a folder
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateFolder(request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    return c.CreateFolderWithContext(context.Background(), request)
}

// CreateFolder
// Create a folder
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateFolderWithContext(ctx context.Context, request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    if request == nil {
        request = NewCreateFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSchemaRequest() (request *CreateSchemaRequest) {
    request = &CreateSchemaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateSchema")
    
    
    return
}

func NewCreateSchemaResponse() (response *CreateSchemaResponse) {
    response = &CreateSchemaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSchema
// Create a schema
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCEINUSE_SCHEMAALREADYEXISTS = "ResourceInUse.SchemaAlreadyExists"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) CreateSchema(request *CreateSchemaRequest) (response *CreateSchemaResponse, err error) {
    return c.CreateSchemaWithContext(context.Background(), request)
}

// CreateSchema
// Create a schema
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCEINUSE_SCHEMAALREADYEXISTS = "ResourceInUse.SchemaAlreadyExists"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) CreateSchemaWithContext(ctx context.Context, request *CreateSchemaRequest) (response *CreateSchemaResponse, err error) {
    if request == nil {
        request = NewCreateSchemaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateSchema")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchema require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSchemaResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRequest() (request *CreateWorkflowRequest) {
    request = &CreateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateWorkflow")
    
    
    return
}

func NewCreateWorkflowResponse() (response *CreateWorkflowResponse) {
    response = &CreateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflow
// Create a workflow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_CREATEWORKFLOWFAILED = "FailedOperation.CreateWorkflowFailed"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_WORKFLOWCOUNTLIMIT = "FailedOperation.WorkflowCountLimit"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) CreateWorkflow(request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    return c.CreateWorkflowWithContext(context.Background(), request)
}

// CreateWorkflow
// Create a workflow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_CREATEWORKFLOWFAILED = "FailedOperation.CreateWorkflowFailed"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_WORKFLOWCOUNTLIMIT = "FailedOperation.WorkflowCountLimit"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) CreateWorkflowWithContext(ctx context.Context, request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceRequest() (request *CreateWorkspaceRequest) {
    request = &CreateWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateWorkspace")
    
    
    return
}

func NewCreateWorkspaceResponse() (response *CreateWorkspaceResponse) {
    response = &CreateWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkspace
// This API is used to create a workspace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_CREATEWORKFLOWFAILED = "FailedOperation.CreateWorkflowFailed"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_WORKFLOWCOUNTLIMIT = "FailedOperation.WorkflowCountLimit"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) CreateWorkspace(request *CreateWorkspaceRequest) (response *CreateWorkspaceResponse, err error) {
    return c.CreateWorkspaceWithContext(context.Background(), request)
}

// CreateWorkspace
// This API is used to create a workspace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_CREATEWORKFLOWFAILED = "FailedOperation.CreateWorkflowFailed"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_WORKFLOWCOUNTLIMIT = "FailedOperation.WorkflowCountLimit"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) CreateWorkspaceWithContext(ctx context.Context, request *CreateWorkspaceRequest) (response *CreateWorkspaceResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateWorkspace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceRoleRequest() (request *CreateWorkspaceRoleRequest) {
    request = &CreateWorkspaceRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateWorkspaceRole")
    
    
    return
}

func NewCreateWorkspaceRoleResponse() (response *CreateWorkspaceRoleResponse) {
    response = &CreateWorkspaceRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkspaceRole
// Creates a workspace role
//
// error code that may be returned:
//  LIMITEXCEEDED_CUSTOMROLEQUOTA = "LimitExceeded.CustomRoleQuota"
func (c *Client) CreateWorkspaceRole(request *CreateWorkspaceRoleRequest) (response *CreateWorkspaceRoleResponse, err error) {
    return c.CreateWorkspaceRoleWithContext(context.Background(), request)
}

// CreateWorkspaceRole
// Creates a workspace role
//
// error code that may be returned:
//  LIMITEXCEEDED_CUSTOMROLEQUOTA = "LimitExceeded.CustomRoleQuota"
func (c *Client) CreateWorkspaceRoleWithContext(ctx context.Context, request *CreateWorkspaceRoleRequest) (response *CreateWorkspaceRoleResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateWorkspaceRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaceRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCatalogRequest() (request *DeleteCatalogRequest) {
    request = &DeleteCatalogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteCatalog")
    
    
    return
}

func NewDeleteCatalogResponse() (response *DeleteCatalogResponse) {
    response = &DeleteCatalogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCatalog
// Delete a catalog.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteCatalog(request *DeleteCatalogRequest) (response *DeleteCatalogResponse, err error) {
    return c.DeleteCatalogWithContext(context.Background(), request)
}

// DeleteCatalog
// Delete a catalog.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteCatalogWithContext(ctx context.Context, request *DeleteCatalogRequest) (response *DeleteCatalogResponse, err error) {
    if request == nil {
        request = NewDeleteCatalogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteCatalog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCatalog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCatalogResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsoleGroupsRequest() (request *DeleteConsoleGroupsRequest) {
    request = &DeleteConsoleGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteConsoleGroups")
    
    
    return
}

func NewDeleteConsoleGroupsResponse() (response *DeleteConsoleGroupsResponse) {
    response = &DeleteConsoleGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsoleGroups
// Delete a console user group
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteConsoleGroups(request *DeleteConsoleGroupsRequest) (response *DeleteConsoleGroupsResponse, err error) {
    return c.DeleteConsoleGroupsWithContext(context.Background(), request)
}

// DeleteConsoleGroups
// Delete a console user group
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteConsoleGroupsWithContext(ctx context.Context, request *DeleteConsoleGroupsRequest) (response *DeleteConsoleGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteConsoleGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteConsoleGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsoleGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsoleGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileRequest() (request *DeleteFileRequest) {
    request = &DeleteFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteFile")
    
    
    return
}

func NewDeleteFileResponse() (response *DeleteFileResponse) {
    response = &DeleteFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFile
// Move the file to Recycle Bin (soft deletion) and clear its version history and execution result snapshots.
//
// 
//
// **Prerequisite**
//
// 1. The file corresponding to FileId must exist and be in active status;
//
// 2. The caller has permission to delete the file;
//
// 3. The file hasn't been referenced by a workflow task.
//
// 
//
// **Error codes (Module are all `Studio`)**
//
// 
//
// | Error Code (Code) | InnerCode | Description | Handling Suggestion |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | lack of WorkspaceId | Please input WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | lack of FileId | Please import FileId |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType value not supported | FileType takes FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | File does not exist or has been deleted | Please confirm the FileId |
//
// | `ResourceInUse.FileReferencedByTask` | 1030204 | The file is referenced by a workflow task and cannot be deleted. | Disassociate the task reference first. |
//
// | `UnauthorizedOperation.FileDeleteDenied` | 1030303 | No deletion permission for this file | Contact the file owner or space admin for authorization |
//
// | `InternalError` | 1030900 | Internal service error | Please carry the RequestId and contact support |
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteFile(request *DeleteFileRequest) (response *DeleteFileResponse, err error) {
    return c.DeleteFileWithContext(context.Background(), request)
}

// DeleteFile
// Move the file to Recycle Bin (soft deletion) and clear its version history and execution result snapshots.
//
// 
//
// **Prerequisite**
//
// 1. The file corresponding to FileId must exist and be in active status;
//
// 2. The caller has permission to delete the file;
//
// 3. The file hasn't been referenced by a workflow task.
//
// 
//
// **Error codes (Module are all `Studio`)**
//
// 
//
// | Error Code (Code) | InnerCode | Description | Handling Suggestion |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | lack of WorkspaceId | Please input WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | lack of FileId | Please import FileId |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType value not supported | FileType takes FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | File does not exist or has been deleted | Please confirm the FileId |
//
// | `ResourceInUse.FileReferencedByTask` | 1030204 | The file is referenced by a workflow task and cannot be deleted. | Disassociate the task reference first. |
//
// | `UnauthorizedOperation.FileDeleteDenied` | 1030303 | No deletion permission for this file | Contact the file owner or space admin for authorization |
//
// | `InternalError` | 1030900 | Internal service error | Please carry the RequestId and contact support |
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteFileWithContext(ctx context.Context, request *DeleteFileRequest) (response *DeleteFileResponse, err error) {
    if request == nil {
        request = NewDeleteFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFolderRequest() (request *DeleteFolderRequest) {
    request = &DeleteFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteFolder")
    
    
    return
}

func NewDeleteFolderResponse() (response *DeleteFolderResponse) {
    response = &DeleteFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFolder
// Delete folder
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteFolder(request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    return c.DeleteFolderWithContext(context.Background(), request)
}

// DeleteFolder
// Delete folder
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteFolderWithContext(ctx context.Context, request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    if request == nil {
        request = NewDeleteFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSchemaRequest() (request *DeleteSchemaRequest) {
    request = &DeleteSchemaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteSchema")
    
    
    return
}

func NewDeleteSchemaResponse() (response *DeleteSchemaResponse) {
    response = &DeleteSchemaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSchema
// Delete schema
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
func (c *Client) DeleteSchema(request *DeleteSchemaRequest) (response *DeleteSchemaResponse, err error) {
    return c.DeleteSchemaWithContext(context.Background(), request)
}

// DeleteSchema
// Delete schema
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
func (c *Client) DeleteSchemaWithContext(ctx context.Context, request *DeleteSchemaRequest) (response *DeleteSchemaResponse, err error) {
    if request == nil {
        request = NewDeleteSchemaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteSchema")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSchema require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSchemaResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowRequest() (request *DeleteWorkflowRequest) {
    request = &DeleteWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteWorkflow")
    
    
    return
}

func NewDeleteWorkflowResponse() (response *DeleteWorkflowResponse) {
    response = &DeleteWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkflow
// Deletes a workflow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    return c.DeleteWorkflowWithContext(context.Background(), request)
}

// DeleteWorkflow
// Deletes a workflow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) DeleteWorkflowWithContext(ctx context.Context, request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkspaceRequest() (request *DeleteWorkspaceRequest) {
    request = &DeleteWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteWorkspace")
    
    
    return
}

func NewDeleteWorkspaceResponse() (response *DeleteWorkspaceResponse) {
    response = &DeleteWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkspace
// Delete a workspace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) DeleteWorkspace(request *DeleteWorkspaceRequest) (response *DeleteWorkspaceResponse, err error) {
    return c.DeleteWorkspaceWithContext(context.Background(), request)
}

// DeleteWorkspace
// Delete a workspace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) DeleteWorkspaceWithContext(ctx context.Context, request *DeleteWorkspaceRequest) (response *DeleteWorkspaceResponse, err error) {
    if request == nil {
        request = NewDeleteWorkspaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteWorkspace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkspaceRoleRequest() (request *DeleteWorkspaceRoleRequest) {
    request = &DeleteWorkspaceRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteWorkspaceRole")
    
    
    return
}

func NewDeleteWorkspaceRoleResponse() (response *DeleteWorkspaceRoleResponse) {
    response = &DeleteWorkspaceRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkspaceRole
// Delete a workspace role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) DeleteWorkspaceRole(request *DeleteWorkspaceRoleRequest) (response *DeleteWorkspaceRoleResponse, err error) {
    return c.DeleteWorkspaceRoleWithContext(context.Background(), request)
}

// DeleteWorkspaceRole
// Delete a workspace role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) DeleteWorkspaceRoleWithContext(ctx context.Context, request *DeleteWorkspaceRoleRequest) (response *DeleteWorkspaceRoleResponse, err error) {
    if request == nil {
        request = NewDeleteWorkspaceRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteWorkspaceRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkspaceRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkspaceRoleResponse()
    err = c.Send(request, response)
    return
}

func NewGetCatalogRequest() (request *GetCatalogRequest) {
    request = &GetCatalogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetCatalog")
    
    
    return
}

func NewGetCatalogResponse() (response *GetCatalogResponse) {
    response = &GetCatalogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCatalog
// Query catalog description
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetCatalog(request *GetCatalogRequest) (response *GetCatalogResponse, err error) {
    return c.GetCatalogWithContext(context.Background(), request)
}

// GetCatalog
// Query catalog description
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetCatalogWithContext(ctx context.Context, request *GetCatalogRequest) (response *GetCatalogResponse, err error) {
    if request == nil {
        request = NewGetCatalogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetCatalog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCatalog require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCatalogResponse()
    err = c.Send(request, response)
    return
}

func NewGetFileRequest() (request *GetFileRequest) {
    request = &GetFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetFile")
    
    
    return
}

func NewGetFileResponse() (response *GetFileResponse) {
    response = &GetFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFile
// Get the file's metadata, optionally including file content, and read historical snapshots by version.
//
// 
//
// **Prerequisite**
//
// 1. Pass at least one of FileId and FilePath. If both are passed, FileId takes precedence.
//
// 2. The corresponding file must exist, and the caller must have read permission for the file;
//
// 3. If VersionId is passed, the version must exist.
//
// 
//
// **Error codes (Module: `Studio`)**
//
// 
//
// | Error Code | Inner Code | Description | Handling Suggestion |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | Lack of WorkspaceId | Please input WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | FileId and FilePath cannot be empty at the same time | Pass at least one of FileId and FilePath |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType value unsupported | FileType takes FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | File does not exist or has been deleted | Confirm the FileId or FilePath |
//
// | `ResourceNotFound.FileVersionNotFound` | 1030205 | The specified file version does not exist | Confirm the VersionId or call ListFileVersions to retrieve |
//
// | `UnauthorizedOperation.FileReadDenied` | 1030304 | No read permission for this file | Contact the file owner or space admin for authorization |
//
// | `InternalError` | 1030900 | Internal service error | Please carry the RequestId and contact support |
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetFile(request *GetFileRequest) (response *GetFileResponse, err error) {
    return c.GetFileWithContext(context.Background(), request)
}

// GetFile
// Get the file's metadata, optionally including file content, and read historical snapshots by version.
//
// 
//
// **Prerequisite**
//
// 1. Pass at least one of FileId and FilePath. If both are passed, FileId takes precedence.
//
// 2. The corresponding file must exist, and the caller must have read permission for the file;
//
// 3. If VersionId is passed, the version must exist.
//
// 
//
// **Error codes (Module: `Studio`)**
//
// 
//
// | Error Code | Inner Code | Description | Handling Suggestion |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | Lack of WorkspaceId | Please input WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | FileId and FilePath cannot be empty at the same time | Pass at least one of FileId and FilePath |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType value unsupported | FileType takes FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | File does not exist or has been deleted | Confirm the FileId or FilePath |
//
// | `ResourceNotFound.FileVersionNotFound` | 1030205 | The specified file version does not exist | Confirm the VersionId or call ListFileVersions to retrieve |
//
// | `UnauthorizedOperation.FileReadDenied` | 1030304 | No read permission for this file | Contact the file owner or space admin for authorization |
//
// | `InternalError` | 1030900 | Internal service error | Please carry the RequestId and contact support |
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetFileWithContext(ctx context.Context, request *GetFileRequest) (response *GetFileResponse, err error) {
    if request == nil {
        request = NewGetFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFileResponse()
    err = c.Send(request, response)
    return
}

func NewGetFolderRequest() (request *GetFolderRequest) {
    request = &GetFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetFolder")
    
    
    return
}

func NewGetFolderResponse() (response *GetFolderResponse) {
    response = &GetFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFolder
// Get folder details.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetFolder(request *GetFolderRequest) (response *GetFolderResponse, err error) {
    return c.GetFolderWithContext(context.Background(), request)
}

// GetFolder
// Get folder details.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetFolderWithContext(ctx context.Context, request *GetFolderRequest) (response *GetFolderResponse, err error) {
    if request == nil {
        request = NewGetFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFolderResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowRequest() (request *GetWorkflowRequest) {
    request = &GetWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetWorkflow")
    
    
    return
}

func NewGetWorkflowResponse() (response *GetWorkflowResponse) {
    response = &GetWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkflow
// This API is used to query workflow details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) GetWorkflow(request *GetWorkflowRequest) (response *GetWorkflowResponse, err error) {
    return c.GetWorkflowWithContext(context.Background(), request)
}

// GetWorkflow
// This API is used to query workflow details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) GetWorkflowWithContext(ctx context.Context, request *GetWorkflowRequest) (response *GetWorkflowResponse, err error) {
    if request == nil {
        request = NewGetWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowRunRequest() (request *GetWorkflowRunRequest) {
    request = &GetWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetWorkflowRun")
    
    
    return
}

func NewGetWorkflowRunResponse() (response *GetWorkflowRunResponse) {
    response = &GetWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkflowRun
// Queries workflow running details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWEXECUTIONNOTEXIST = "ResourceNotFound.WorkflowExecutionNotExist"
func (c *Client) GetWorkflowRun(request *GetWorkflowRunRequest) (response *GetWorkflowRunResponse, err error) {
    return c.GetWorkflowRunWithContext(context.Background(), request)
}

// GetWorkflowRun
// Queries workflow running details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWEXECUTIONNOTEXIST = "ResourceNotFound.WorkflowExecutionNotExist"
func (c *Client) GetWorkflowRunWithContext(ctx context.Context, request *GetWorkflowRunRequest) (response *GetWorkflowRunResponse, err error) {
    if request == nil {
        request = NewGetWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowTaskRunRequest() (request *GetWorkflowTaskRunRequest) {
    request = &GetWorkflowTaskRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetWorkflowTaskRun")
    
    
    return
}

func NewGetWorkflowTaskRunResponse() (response *GetWorkflowTaskRunResponse) {
    response = &GetWorkflowTaskRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkflowTaskRun
// Query task execution details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TASKEXECUTIONNOTEXIST = "ResourceNotFound.TaskExecutionNotExist"
func (c *Client) GetWorkflowTaskRun(request *GetWorkflowTaskRunRequest) (response *GetWorkflowTaskRunResponse, err error) {
    return c.GetWorkflowTaskRunWithContext(context.Background(), request)
}

// GetWorkflowTaskRun
// Query task execution details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TASKEXECUTIONNOTEXIST = "ResourceNotFound.TaskExecutionNotExist"
func (c *Client) GetWorkflowTaskRunWithContext(ctx context.Context, request *GetWorkflowTaskRunRequest) (response *GetWorkflowTaskRunResponse, err error) {
    if request == nil {
        request = NewGetWorkflowTaskRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetWorkflowTaskRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkflowTaskRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkflowTaskRunResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkspaceRequest() (request *GetWorkspaceRequest) {
    request = &GetWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetWorkspace")
    
    
    return
}

func NewGetWorkspaceResponse() (response *GetWorkspaceResponse) {
    response = &GetWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkspace
// Queries the details of a workspace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TASKEXECUTIONNOTEXIST = "ResourceNotFound.TaskExecutionNotExist"
func (c *Client) GetWorkspace(request *GetWorkspaceRequest) (response *GetWorkspaceResponse, err error) {
    return c.GetWorkspaceWithContext(context.Background(), request)
}

// GetWorkspace
// Queries the details of a workspace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TASKEXECUTIONNOTEXIST = "ResourceNotFound.TaskExecutionNotExist"
func (c *Client) GetWorkspaceWithContext(ctx context.Context, request *GetWorkspaceRequest) (response *GetWorkspaceResponse, err error) {
    if request == nil {
        request = NewGetWorkspaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetWorkspace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewKillWorkflowRunRequest() (request *KillWorkflowRunRequest) {
    request = &KillWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "KillWorkflowRun")
    
    
    return
}

func NewKillWorkflowRunResponse() (response *KillWorkflowRunResponse) {
    response = &KillWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillWorkflowRun
// Terminate the operation of a workflow.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_NOWORKFLOWEXECUTIONNEEDOPERATE = "FailedOperation.NoWorkflowExecutionNeedOperate"
//  FAILEDOPERATION_WORKFLOWEXECUTIONHASBEDELETE = "FailedOperation.WorkflowExecutionHasBeDelete"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_WORKFLOWEXECUTIONHASREACHEDFINALSTATECANNOTBESTOPPED = "UnsupportedOperation.WorkflowExecutionHasReachedFinalStateCannotBeStopped"
func (c *Client) KillWorkflowRun(request *KillWorkflowRunRequest) (response *KillWorkflowRunResponse, err error) {
    return c.KillWorkflowRunWithContext(context.Background(), request)
}

// KillWorkflowRun
// Terminate the operation of a workflow.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_NOWORKFLOWEXECUTIONNEEDOPERATE = "FailedOperation.NoWorkflowExecutionNeedOperate"
//  FAILEDOPERATION_WORKFLOWEXECUTIONHASBEDELETE = "FailedOperation.WorkflowExecutionHasBeDelete"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_WORKFLOWEXECUTIONHASREACHEDFINALSTATECANNOTBESTOPPED = "UnsupportedOperation.WorkflowExecutionHasReachedFinalStateCannotBeStopped"
func (c *Client) KillWorkflowRunWithContext(ctx context.Context, request *KillWorkflowRunRequest) (response *KillWorkflowRunResponse, err error) {
    if request == nil {
        request = NewKillWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "KillWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewListCatalogsRequest() (request *ListCatalogsRequest) {
    request = &ListCatalogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListCatalogs")
    
    
    return
}

func NewListCatalogsResponse() (response *ListCatalogsResponse) {
    response = &ListCatalogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListCatalogs
// Pagination fetch data catalog list
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListCatalogs(request *ListCatalogsRequest) (response *ListCatalogsResponse, err error) {
    return c.ListCatalogsWithContext(context.Background(), request)
}

// ListCatalogs
// Pagination fetch data catalog list
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListCatalogsWithContext(ctx context.Context, request *ListCatalogsRequest) (response *ListCatalogsResponse, err error) {
    if request == nil {
        request = NewListCatalogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListCatalogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCatalogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCatalogsResponse()
    err = c.Send(request, response)
    return
}

func NewListConsoleGroupUsersRequest() (request *ListConsoleGroupUsersRequest) {
    request = &ListConsoleGroupUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListConsoleGroupUsers")
    
    
    return
}

func NewListConsoleGroupUsersResponse() (response *ListConsoleGroupUsersResponse) {
    response = &ListConsoleGroupUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConsoleGroupUsers
// Queries the console user group member list. This API is a console-level API and is only supported in the central region: use ap-guangzhou for the Chinese site and ap-singapore for the international site. Calls from other regions will return UnsupportedRegion.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleGroupUsers(request *ListConsoleGroupUsersRequest) (response *ListConsoleGroupUsersResponse, err error) {
    return c.ListConsoleGroupUsersWithContext(context.Background(), request)
}

// ListConsoleGroupUsers
// Queries the console user group member list. This API is a console-level API and is only supported in the central region: use ap-guangzhou for the Chinese site and ap-singapore for the international site. Calls from other regions will return UnsupportedRegion.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleGroupUsersWithContext(ctx context.Context, request *ListConsoleGroupUsersRequest) (response *ListConsoleGroupUsersResponse, err error) {
    if request == nil {
        request = NewListConsoleGroupUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListConsoleGroupUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConsoleGroupUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConsoleGroupUsersResponse()
    err = c.Send(request, response)
    return
}

func NewListConsoleGroupsRequest() (request *ListConsoleGroupsRequest) {
    request = &ListConsoleGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListConsoleGroups")
    
    
    return
}

func NewListConsoleGroupsResponse() (response *ListConsoleGroupsResponse) {
    response = &ListConsoleGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConsoleGroups
// Queries the console user group list. This API is a console-level API and is only supported in the central region: use ap-guangzhou for the Chinese site and ap-singapore for the international site. Calls from other regions will return UnsupportedRegion.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleGroups(request *ListConsoleGroupsRequest) (response *ListConsoleGroupsResponse, err error) {
    return c.ListConsoleGroupsWithContext(context.Background(), request)
}

// ListConsoleGroups
// Queries the console user group list. This API is a console-level API and is only supported in the central region: use ap-guangzhou for the Chinese site and ap-singapore for the international site. Calls from other regions will return UnsupportedRegion.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleGroupsWithContext(ctx context.Context, request *ListConsoleGroupsRequest) (response *ListConsoleGroupsResponse, err error) {
    if request == nil {
        request = NewListConsoleGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListConsoleGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConsoleGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConsoleGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListConsoleRolesRequest() (request *ListConsoleRolesRequest) {
    request = &ListConsoleRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListConsoleRoles")
    
    
    return
}

func NewListConsoleRolesResponse() (response *ListConsoleRolesResponse) {
    response = &ListConsoleRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConsoleRoles
// Queries the console role list. This API is a console-level API and is only supported in the central region: use ap-guangzhou for the Chinese site and ap-singapore for the international site. Calls from other regions will return UnsupportedRegion.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleRoles(request *ListConsoleRolesRequest) (response *ListConsoleRolesResponse, err error) {
    return c.ListConsoleRolesWithContext(context.Background(), request)
}

// ListConsoleRoles
// Queries the console role list. This API is a console-level API and is only supported in the central region: use ap-guangzhou for the Chinese site and ap-singapore for the international site. Calls from other regions will return UnsupportedRegion.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleRolesWithContext(ctx context.Context, request *ListConsoleRolesRequest) (response *ListConsoleRolesResponse, err error) {
    if request == nil {
        request = NewListConsoleRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListConsoleRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConsoleRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConsoleRolesResponse()
    err = c.Send(request, response)
    return
}

func NewListConsoleUsersRequest() (request *ListConsoleUsersRequest) {
    request = &ListConsoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListConsoleUsers")
    
    
    return
}

func NewListConsoleUsersResponse() (response *ListConsoleUsersResponse) {
    response = &ListConsoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConsoleUsers
// Queries the console user list. This API is a console-level API and is only supported in the central region: use ap-guangzhou for the Chinese site and ap-singapore for the international site. Calls from other regions will return UnsupportedRegion.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleUsers(request *ListConsoleUsersRequest) (response *ListConsoleUsersResponse, err error) {
    return c.ListConsoleUsersWithContext(context.Background(), request)
}

// ListConsoleUsers
// Queries the console user list. This API is a console-level API and is only supported in the central region: use ap-guangzhou for the Chinese site and ap-singapore for the international site. Calls from other regions will return UnsupportedRegion.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleUsersWithContext(ctx context.Context, request *ListConsoleUsersRequest) (response *ListConsoleUsersResponse, err error) {
    if request == nil {
        request = NewListConsoleUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListConsoleUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConsoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConsoleUsersResponse()
    err = c.Send(request, response)
    return
}

func NewListFilesRequest() (request *ListFilesRequest) {
    request = &ListFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListFiles")
    
    
    return
}

func NewListFilesResponse() (response *ListFilesResponse) {
    response = &ListFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListFiles
// Get the folder and file list
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListFiles(request *ListFilesRequest) (response *ListFilesResponse, err error) {
    return c.ListFilesWithContext(context.Background(), request)
}

// ListFiles
// Get the folder and file list
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListFilesWithContext(ctx context.Context, request *ListFilesRequest) (response *ListFilesResponse, err error) {
    if request == nil {
        request = NewListFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewListFilesResponse()
    err = c.Send(request, response)
    return
}

func NewListSchemasRequest() (request *ListSchemasRequest) {
    request = &ListSchemasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListSchemas")
    
    
    return
}

func NewListSchemasResponse() (response *ListSchemasResponse) {
    response = &ListSchemasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSchemas
// Get schema list
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) ListSchemas(request *ListSchemasRequest) (response *ListSchemasResponse, err error) {
    return c.ListSchemasWithContext(context.Background(), request)
}

// ListSchemas
// Get schema list
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) ListSchemasWithContext(ctx context.Context, request *ListSchemasRequest) (response *ListSchemasResponse, err error) {
    if request == nil {
        request = NewListSchemasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListSchemas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSchemas require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSchemasResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowRunsRequest() (request *ListWorkflowRunsRequest) {
    request = &ListWorkflowRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkflowRuns")
    
    
    return
}

func NewListWorkflowRunsResponse() (response *ListWorkflowRunsResponse) {
    response = &ListWorkflowRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflowRuns
// Workflow run list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTWORKFLOWEXECUTIONS = "FailedOperation.ExistWorkflowExecutions"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  INVALIDPARAMETERVALUE_WORKFLOWENDTIMELESSSTARTTIME = "InvalidParameterValue.WorkflowEndTimeLessStartTime"
//  INVALIDPARAMETERVALUE_WORKFLOWQUERYENDTIMEANDSTARTTIMEEXCEED = "InvalidParameterValue.WorkflowQueryEndTimeAndStartTimeExceed"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
func (c *Client) ListWorkflowRuns(request *ListWorkflowRunsRequest) (response *ListWorkflowRunsResponse, err error) {
    return c.ListWorkflowRunsWithContext(context.Background(), request)
}

// ListWorkflowRuns
// Workflow run list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTWORKFLOWEXECUTIONS = "FailedOperation.ExistWorkflowExecutions"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  INVALIDPARAMETERVALUE_WORKFLOWENDTIMELESSSTARTTIME = "InvalidParameterValue.WorkflowEndTimeLessStartTime"
//  INVALIDPARAMETERVALUE_WORKFLOWQUERYENDTIMEANDSTARTTIMEEXCEED = "InvalidParameterValue.WorkflowQueryEndTimeAndStartTimeExceed"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
func (c *Client) ListWorkflowRunsWithContext(ctx context.Context, request *ListWorkflowRunsRequest) (response *ListWorkflowRunsResponse, err error) {
    if request == nil {
        request = NewListWorkflowRunsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkflowRuns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflowRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowRunsResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowTaskRunsRequest() (request *ListWorkflowTaskRunsRequest) {
    request = &ListWorkflowTaskRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkflowTaskRuns")
    
    
    return
}

func NewListWorkflowTaskRunsResponse() (response *ListWorkflowTaskRunsResponse) {
    response = &ListWorkflowTaskRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflowTaskRuns
// Query the historical operation list of workflow tasks
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListWorkflowTaskRuns(request *ListWorkflowTaskRunsRequest) (response *ListWorkflowTaskRunsResponse, err error) {
    return c.ListWorkflowTaskRunsWithContext(context.Background(), request)
}

// ListWorkflowTaskRuns
// Query the historical operation list of workflow tasks
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListWorkflowTaskRunsWithContext(ctx context.Context, request *ListWorkflowTaskRunsRequest) (response *ListWorkflowTaskRunsResponse, err error) {
    if request == nil {
        request = NewListWorkflowTaskRunsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkflowTaskRuns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflowTaskRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowTaskRunsResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowsRequest() (request *ListWorkflowsRequest) {
    request = &ListWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkflows")
    
    
    return
}

func NewListWorkflowsResponse() (response *ListWorkflowsResponse) {
    response = &ListWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflows
// This API is used to query the list of workflows.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflows(request *ListWorkflowsRequest) (response *ListWorkflowsResponse, err error) {
    return c.ListWorkflowsWithContext(context.Background(), request)
}

// ListWorkflows
// This API is used to query the list of workflows.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflowsWithContext(ctx context.Context, request *ListWorkflowsRequest) (response *ListWorkflowsResponse, err error) {
    if request == nil {
        request = NewListWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkspaceGroupsRequest() (request *ListWorkspaceGroupsRequest) {
    request = &ListWorkspaceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkspaceGroups")
    
    
    return
}

func NewListWorkspaceGroupsResponse() (response *ListWorkspaceGroupsResponse) {
    response = &ListWorkspaceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkspaceGroups
// This API is used to query the list of user groups in a workspace.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkspaceGroups(request *ListWorkspaceGroupsRequest) (response *ListWorkspaceGroupsResponse, err error) {
    return c.ListWorkspaceGroupsWithContext(context.Background(), request)
}

// ListWorkspaceGroups
// This API is used to query the list of user groups in a workspace.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkspaceGroupsWithContext(ctx context.Context, request *ListWorkspaceGroupsRequest) (response *ListWorkspaceGroupsResponse, err error) {
    if request == nil {
        request = NewListWorkspaceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkspaceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkspaceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkspaceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkspaceUsersRequest() (request *ListWorkspaceUsersRequest) {
    request = &ListWorkspaceUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkspaceUsers")
    
    
    return
}

func NewListWorkspaceUsersResponse() (response *ListWorkspaceUsersResponse) {
    response = &ListWorkspaceUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkspaceUsers
// Queries the list of users in a workspace.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkspaceUsers(request *ListWorkspaceUsersRequest) (response *ListWorkspaceUsersResponse, err error) {
    return c.ListWorkspaceUsersWithContext(context.Background(), request)
}

// ListWorkspaceUsers
// Queries the list of users in a workspace.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkspaceUsersWithContext(ctx context.Context, request *ListWorkspaceUsersRequest) (response *ListWorkspaceUsersResponse, err error) {
    if request == nil {
        request = NewListWorkspaceUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkspaceUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkspaceUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkspaceUsersResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkspacesRequest() (request *ListWorkspacesRequest) {
    request = &ListWorkspacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkspaces")
    
    
    return
}

func NewListWorkspacesResponse() (response *ListWorkspacesResponse) {
    response = &ListWorkspacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkspaces
// Queries the workspace list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkspaces(request *ListWorkspacesRequest) (response *ListWorkspacesResponse, err error) {
    return c.ListWorkspacesWithContext(context.Background(), request)
}

// ListWorkspaces
// Queries the workspace list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkspacesWithContext(ctx context.Context, request *ListWorkspacesRequest) (response *ListWorkspacesResponse, err error) {
    if request == nil {
        request = NewListWorkspacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkspaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkspaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkspacesResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveConsoleUsersRequest() (request *RemoveConsoleUsersRequest) {
    request = &RemoveConsoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "RemoveConsoleUsers")
    
    
    return
}

func NewRemoveConsoleUsersResponse() (response *RemoveConsoleUsersResponse) {
    response = &RemoveConsoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveConsoleUsers
// <p>Batch remove console users (up to 10 at a time; pre-verification rejects the entire request if any fails; during the Execution Phase, a single failure does not interrupt subsequent deletion, and success or failure is determined by SuccessUins/FailItems)</p>
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RemoveConsoleUsers(request *RemoveConsoleUsersRequest) (response *RemoveConsoleUsersResponse, err error) {
    return c.RemoveConsoleUsersWithContext(context.Background(), request)
}

// RemoveConsoleUsers
// <p>Batch remove console users (up to 10 at a time; pre-verification rejects the entire request if any fails; during the Execution Phase, a single failure does not interrupt subsequent deletion, and success or failure is determined by SuccessUins/FailItems)</p>
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RemoveConsoleUsersWithContext(ctx context.Context, request *RemoveConsoleUsersRequest) (response *RemoveConsoleUsersResponse, err error) {
    if request == nil {
        request = NewRemoveConsoleUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "RemoveConsoleUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveConsoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveConsoleUsersResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveWorkspaceGroupsRequest() (request *RemoveWorkspaceGroupsRequest) {
    request = &RemoveWorkspaceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "RemoveWorkspaceGroups")
    
    
    return
}

func NewRemoveWorkspaceGroupsResponse() (response *RemoveWorkspaceGroupsResponse) {
    response = &RemoveWorkspaceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveWorkspaceGroups
// Remove a user group from a workspace.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RemoveWorkspaceGroups(request *RemoveWorkspaceGroupsRequest) (response *RemoveWorkspaceGroupsResponse, err error) {
    return c.RemoveWorkspaceGroupsWithContext(context.Background(), request)
}

// RemoveWorkspaceGroups
// Remove a user group from a workspace.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RemoveWorkspaceGroupsWithContext(ctx context.Context, request *RemoveWorkspaceGroupsRequest) (response *RemoveWorkspaceGroupsResponse, err error) {
    if request == nil {
        request = NewRemoveWorkspaceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "RemoveWorkspaceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveWorkspaceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveWorkspaceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewRerunWorkflowRunRequest() (request *RerunWorkflowRunRequest) {
    request = &RerunWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "RerunWorkflowRun")
    
    
    return
}

func NewRerunWorkflowRunResponse() (response *RerunWorkflowRunResponse) {
    response = &RerunWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RerunWorkflowRun
// Rerun a workflow
//
// error code that may be returned:
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_RERUNWORKFLOWFAIL = "FailedOperation.RerunWorkflowFail"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMSERROR = "InvalidParameter.InvalidParamsError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  RESOURCENOTFOUND_WORKFLOWEXECUTIONNOTEXIST = "ResourceNotFound.WorkflowExecutionNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWTASKNOTEXIST = "ResourceNotFound.WorkflowTaskNotExist"
func (c *Client) RerunWorkflowRun(request *RerunWorkflowRunRequest) (response *RerunWorkflowRunResponse, err error) {
    return c.RerunWorkflowRunWithContext(context.Background(), request)
}

// RerunWorkflowRun
// Rerun a workflow
//
// error code that may be returned:
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_RERUNWORKFLOWFAIL = "FailedOperation.RerunWorkflowFail"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMSERROR = "InvalidParameter.InvalidParamsError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  RESOURCENOTFOUND_WORKFLOWEXECUTIONNOTEXIST = "ResourceNotFound.WorkflowExecutionNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWTASKNOTEXIST = "ResourceNotFound.WorkflowTaskNotExist"
func (c *Client) RerunWorkflowRunWithContext(ctx context.Context, request *RerunWorkflowRunRequest) (response *RerunWorkflowRunResponse, err error) {
    if request == nil {
        request = NewRerunWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "RerunWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RerunWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewRerunWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewRunWorkflowRequest() (request *RunWorkflowRequest) {
    request = &RunWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "RunWorkflow")
    
    
    return
}

func NewRunWorkflowResponse() (response *RunWorkflowResponse) {
    response = &RunWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunWorkflow
// Runs the workflow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_RUNWORKFLOWFAIL = "FailedOperation.RunWorkflowFail"
//  FAILEDOPERATION_RUNWORKFLOWFAILEXECUTIONIDEMPTY = "FailedOperation.RunWorkflowFailExecutionIdEmpty"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWTASKNOTEXIST = "ResourceNotFound.WorkflowTaskNotExist"
func (c *Client) RunWorkflow(request *RunWorkflowRequest) (response *RunWorkflowResponse, err error) {
    return c.RunWorkflowWithContext(context.Background(), request)
}

// RunWorkflow
// Runs the workflow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_RUNWORKFLOWFAIL = "FailedOperation.RunWorkflowFail"
//  FAILEDOPERATION_RUNWORKFLOWFAILEXECUTIONIDEMPTY = "FailedOperation.RunWorkflowFailExecutionIdEmpty"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWTASKNOTEXIST = "ResourceNotFound.WorkflowTaskNotExist"
func (c *Client) RunWorkflowWithContext(ctx context.Context, request *RunWorkflowRequest) (response *RunWorkflowResponse, err error) {
    if request == nil {
        request = NewRunWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "RunWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewStartComputeRequest() (request *StartComputeRequest) {
    request = &StartComputeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "StartCompute")
    
    
    return
}

func NewStartComputeResponse() (response *StartComputeResponse) {
    response = &StartComputeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartCompute
// Start up the computational resource.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INTERNALERROR_INVALIDPARAMETER = "InternalError.InvalidParameter"
//  INTERNALERROR_PERMISSIONDENIED = "InternalError.PermissionDenied"
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
func (c *Client) StartCompute(request *StartComputeRequest) (response *StartComputeResponse, err error) {
    return c.StartComputeWithContext(context.Background(), request)
}

// StartCompute
// Start up the computational resource.
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INTERNALERROR_INVALIDPARAMETER = "InternalError.InvalidParameter"
//  INTERNALERROR_PERMISSIONDENIED = "InternalError.PermissionDenied"
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
func (c *Client) StartComputeWithContext(ctx context.Context, request *StartComputeRequest) (response *StartComputeResponse, err error) {
    if request == nil {
        request = NewStartComputeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "StartCompute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCompute require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartComputeResponse()
    err = c.Send(request, response)
    return
}

func NewStopComputeRequest() (request *StopComputeRequest) {
    request = &StopComputeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "StopCompute")
    
    
    return
}

func NewStopComputeResponse() (response *StopComputeResponse) {
    response = &StopComputeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCompute
// Terminate computational resources
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INTERNALERROR_INVALIDPARAMETER = "InternalError.InvalidParameter"
//  INTERNALERROR_PERMISSIONDENIED = "InternalError.PermissionDenied"
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
func (c *Client) StopCompute(request *StopComputeRequest) (response *StopComputeResponse, err error) {
    return c.StopComputeWithContext(context.Background(), request)
}

// StopCompute
// Terminate computational resources
//
// error code that may be returned:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INTERNALERROR_INVALIDPARAMETER = "InternalError.InvalidParameter"
//  INTERNALERROR_PERMISSIONDENIED = "InternalError.PermissionDenied"
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
func (c *Client) StopComputeWithContext(ctx context.Context, request *StopComputeRequest) (response *StopComputeResponse, err error) {
    if request == nil {
        request = NewStopComputeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "StopCompute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCompute require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopComputeResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindWorkflowBundleRequest() (request *UnbindWorkflowBundleRequest) {
    request = &UnbindWorkflowBundleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UnbindWorkflowBundle")
    
    
    return
}

func NewUnbindWorkflowBundleResponse() (response *UnbindWorkflowBundleResponse) {
    response = &UnbindWorkflowBundleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindWorkflowBundle
// Unbind workflow Bundle info
//
// Description: This interface is semantically equivalent to Detach in the standard verb list. The name Unbind is retained for compatibility with existing product forms.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UnbindWorkflowBundle(request *UnbindWorkflowBundleRequest) (response *UnbindWorkflowBundleResponse, err error) {
    return c.UnbindWorkflowBundleWithContext(context.Background(), request)
}

// UnbindWorkflowBundle
// Unbind workflow Bundle info
//
// Description: This interface is semantically equivalent to Detach in the standard verb list. The name Unbind is retained for compatibility with existing product forms.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UnbindWorkflowBundleWithContext(ctx context.Context, request *UnbindWorkflowBundleRequest) (response *UnbindWorkflowBundleResponse, err error) {
    if request == nil {
        request = NewUnbindWorkflowBundleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UnbindWorkflowBundle")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindWorkflowBundle require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindWorkflowBundleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateConsoleGroupRequest() (request *UpdateConsoleGroupRequest) {
    request = &UpdateConsoleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateConsoleGroup")
    
    
    return
}

func NewUpdateConsoleGroupResponse() (response *UpdateConsoleGroupResponse) {
    response = &UpdateConsoleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateConsoleGroup
// Modify a console user group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateConsoleGroup(request *UpdateConsoleGroupRequest) (response *UpdateConsoleGroupResponse, err error) {
    return c.UpdateConsoleGroupWithContext(context.Background(), request)
}

// UpdateConsoleGroup
// Modify a console user group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateConsoleGroupWithContext(ctx context.Context, request *UpdateConsoleGroupRequest) (response *UpdateConsoleGroupResponse, err error) {
    if request == nil {
        request = NewUpdateConsoleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateConsoleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateConsoleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateConsoleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateConsoleUsersRequest() (request *UpdateConsoleUsersRequest) {
    request = &UpdateConsoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateConsoleUsers")
    
    
    return
}

func NewUpdateConsoleUsersResponse() (response *UpdateConsoleUsersResponse) {
    response = &UpdateConsoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateConsoleUsers
// Modify a console user role
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateConsoleUsers(request *UpdateConsoleUsersRequest) (response *UpdateConsoleUsersResponse, err error) {
    return c.UpdateConsoleUsersWithContext(context.Background(), request)
}

// UpdateConsoleUsers
// Modify a console user role
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateConsoleUsersWithContext(ctx context.Context, request *UpdateConsoleUsersRequest) (response *UpdateConsoleUsersResponse, err error) {
    if request == nil {
        request = NewUpdateConsoleUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateConsoleUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateConsoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateConsoleUsersResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFileRequest() (request *UpdateFileRequest) {
    request = &UpdateFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateFile")
    
    
    return
}

func NewUpdateFileResponse() (response *UpdateFileResponse) {
    response = &UpdateFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateFile
// Update file contents and running configuration (computational resource, default catalog/schema, parameters), and return the file metadata after the update.
//
// 
//
// **Prerequisite**
//
// 1. The file corresponding to FileId must exist and be in active status;
//
// 2. The caller has write permission on the file;
//
// 3. To update only the configuration, do not pass in Storage. To update only the content, do not pass in FileConfig.
//
// 4. If FileConfig.ResourceId is not empty, the match between the resource type and file type will be validated.
//
// 
//
// **Error codes (Module: `Studio`)**
//
// 
//
// | Error Code (Code) | InnerCode | Description | Handling Suggestions |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | lack of WorkspaceId | Please input WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | lack of FileId | Please import FileId |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType value not supported | FileType takes FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `InvalidParameterValue.ResourceId` | 1030104 | Compute resource type and file type mismatch | Select a computational resource for Python/Notebook and a data analysis resource for SQL |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | File does not exist or has been deleted | confirm the FileId or call GetFile to check the file status |
//
// | `UnauthorizedOperation.FileWriteDenied` | 1030302 | No write permission for this file | Contact the file owner or space admin for authorization |
//
// | `FailedOperation.FileStorageUpdateFailed` | 1030401 | Failed to write file content to storage | Try again later. If the failure persists, contact support with the RequestId. |
//
// | `InternalError` | 1030900 | Internal service error | Please carry the RequestId and contact support |
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateFile(request *UpdateFileRequest) (response *UpdateFileResponse, err error) {
    return c.UpdateFileWithContext(context.Background(), request)
}

// UpdateFile
// Update file contents and running configuration (computational resource, default catalog/schema, parameters), and return the file metadata after the update.
//
// 
//
// **Prerequisite**
//
// 1. The file corresponding to FileId must exist and be in active status;
//
// 2. The caller has write permission on the file;
//
// 3. To update only the configuration, do not pass in Storage. To update only the content, do not pass in FileConfig.
//
// 4. If FileConfig.ResourceId is not empty, the match between the resource type and file type will be validated.
//
// 
//
// **Error codes (Module: `Studio`)**
//
// 
//
// | Error Code (Code) | InnerCode | Description | Handling Suggestions |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | lack of WorkspaceId | Please input WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | lack of FileId | Please import FileId |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType value not supported | FileType takes FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `InvalidParameterValue.ResourceId` | 1030104 | Compute resource type and file type mismatch | Select a computational resource for Python/Notebook and a data analysis resource for SQL |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | File does not exist or has been deleted | confirm the FileId or call GetFile to check the file status |
//
// | `UnauthorizedOperation.FileWriteDenied` | 1030302 | No write permission for this file | Contact the file owner or space admin for authorization |
//
// | `FailedOperation.FileStorageUpdateFailed` | 1030401 | Failed to write file content to storage | Try again later. If the failure persists, contact support with the RequestId. |
//
// | `InternalError` | 1030900 | Internal service error | Please carry the RequestId and contact support |
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateFileWithContext(ctx context.Context, request *UpdateFileRequest) (response *UpdateFileResponse, err error) {
    if request == nil {
        request = NewUpdateFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateFileResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFolderRequest() (request *UpdateFolderRequest) {
    request = &UpdateFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateFolder")
    
    
    return
}

func NewUpdateFolderResponse() (response *UpdateFolderResponse) {
    response = &UpdateFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateFolder
// Update folder (supports rename and move)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateFolder(request *UpdateFolderRequest) (response *UpdateFolderResponse, err error) {
    return c.UpdateFolderWithContext(context.Background(), request)
}

// UpdateFolder
// Update folder (supports rename and move)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateFolderWithContext(ctx context.Context, request *UpdateFolderRequest) (response *UpdateFolderResponse, err error) {
    if request == nil {
        request = NewUpdateFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateFolderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkflowRequest() (request *UpdateWorkflowRequest) {
    request = &UpdateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateWorkflow")
    
    
    return
}

func NewUpdateWorkflowResponse() (response *UpdateWorkflowResponse) {
    response = &UpdateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkflow
// Updating a Workflow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYELEMENTCOUNTLIMIT = "InvalidParameterValue.LoopDataArrayElementCountLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYJSONINVALID = "InvalidParameterValue.LoopDataArrayJsonInvalid"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAY = "InvalidParameterValue.LoopDataArrayNotJsonArray"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAYLITERAL = "InvalidParameterValue.LoopDataArrayNotJsonArrayLiteral"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUEBLANK = "InvalidParameterValue.LoopDataArrayValueBlank"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUELENGTHLIMIT = "InvalidParameterValue.LoopDataArrayValueLengthLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVARIABLEEXPRESSIONINVALID = "InvalidParameterValue.LoopDataArrayVariableExpressionInvalid"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkflow(request *UpdateWorkflowRequest) (response *UpdateWorkflowResponse, err error) {
    return c.UpdateWorkflowWithContext(context.Background(), request)
}

// UpdateWorkflow
// Updating a Workflow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYELEMENTCOUNTLIMIT = "InvalidParameterValue.LoopDataArrayElementCountLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYJSONINVALID = "InvalidParameterValue.LoopDataArrayJsonInvalid"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAY = "InvalidParameterValue.LoopDataArrayNotJsonArray"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAYLITERAL = "InvalidParameterValue.LoopDataArrayNotJsonArrayLiteral"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUEBLANK = "InvalidParameterValue.LoopDataArrayValueBlank"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUELENGTHLIMIT = "InvalidParameterValue.LoopDataArrayValueLengthLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVARIABLEEXPRESSIONINVALID = "InvalidParameterValue.LoopDataArrayVariableExpressionInvalid"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkflowWithContext(ctx context.Context, request *UpdateWorkflowRequest) (response *UpdateWorkflowResponse, err error) {
    if request == nil {
        request = NewUpdateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkspaceRequest() (request *UpdateWorkspaceRequest) {
    request = &UpdateWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateWorkspace")
    
    
    return
}

func NewUpdateWorkspaceResponse() (response *UpdateWorkspaceResponse) {
    response = &UpdateWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkspace
// Modifies a workspace
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYELEMENTCOUNTLIMIT = "InvalidParameterValue.LoopDataArrayElementCountLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYJSONINVALID = "InvalidParameterValue.LoopDataArrayJsonInvalid"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAY = "InvalidParameterValue.LoopDataArrayNotJsonArray"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAYLITERAL = "InvalidParameterValue.LoopDataArrayNotJsonArrayLiteral"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUEBLANK = "InvalidParameterValue.LoopDataArrayValueBlank"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUELENGTHLIMIT = "InvalidParameterValue.LoopDataArrayValueLengthLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVARIABLEEXPRESSIONINVALID = "InvalidParameterValue.LoopDataArrayVariableExpressionInvalid"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspace(request *UpdateWorkspaceRequest) (response *UpdateWorkspaceResponse, err error) {
    return c.UpdateWorkspaceWithContext(context.Background(), request)
}

// UpdateWorkspace
// Modifies a workspace
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYELEMENTCOUNTLIMIT = "InvalidParameterValue.LoopDataArrayElementCountLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYJSONINVALID = "InvalidParameterValue.LoopDataArrayJsonInvalid"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAY = "InvalidParameterValue.LoopDataArrayNotJsonArray"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAYLITERAL = "InvalidParameterValue.LoopDataArrayNotJsonArrayLiteral"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUEBLANK = "InvalidParameterValue.LoopDataArrayValueBlank"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUELENGTHLIMIT = "InvalidParameterValue.LoopDataArrayValueLengthLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVARIABLEEXPRESSIONINVALID = "InvalidParameterValue.LoopDataArrayVariableExpressionInvalid"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceWithContext(ctx context.Context, request *UpdateWorkspaceRequest) (response *UpdateWorkspaceResponse, err error) {
    if request == nil {
        request = NewUpdateWorkspaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateWorkspace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkspaceRoleRequest() (request *UpdateWorkspaceRoleRequest) {
    request = &UpdateWorkspaceRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateWorkspaceRole")
    
    
    return
}

func NewUpdateWorkspaceRoleResponse() (response *UpdateWorkspaceRoleResponse) {
    response = &UpdateWorkspaceRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkspaceRole
// Updates a workspace role
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYELEMENTCOUNTLIMIT = "InvalidParameterValue.LoopDataArrayElementCountLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYJSONINVALID = "InvalidParameterValue.LoopDataArrayJsonInvalid"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAY = "InvalidParameterValue.LoopDataArrayNotJsonArray"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAYLITERAL = "InvalidParameterValue.LoopDataArrayNotJsonArrayLiteral"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUEBLANK = "InvalidParameterValue.LoopDataArrayValueBlank"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUELENGTHLIMIT = "InvalidParameterValue.LoopDataArrayValueLengthLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVARIABLEEXPRESSIONINVALID = "InvalidParameterValue.LoopDataArrayVariableExpressionInvalid"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceRole(request *UpdateWorkspaceRoleRequest) (response *UpdateWorkspaceRoleResponse, err error) {
    return c.UpdateWorkspaceRoleWithContext(context.Background(), request)
}

// UpdateWorkspaceRole
// Updates a workspace role
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYELEMENTCOUNTLIMIT = "InvalidParameterValue.LoopDataArrayElementCountLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYJSONINVALID = "InvalidParameterValue.LoopDataArrayJsonInvalid"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAY = "InvalidParameterValue.LoopDataArrayNotJsonArray"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAYLITERAL = "InvalidParameterValue.LoopDataArrayNotJsonArrayLiteral"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUEBLANK = "InvalidParameterValue.LoopDataArrayValueBlank"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUELENGTHLIMIT = "InvalidParameterValue.LoopDataArrayValueLengthLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVARIABLEEXPRESSIONINVALID = "InvalidParameterValue.LoopDataArrayVariableExpressionInvalid"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceRoleWithContext(ctx context.Context, request *UpdateWorkspaceRoleRequest) (response *UpdateWorkspaceRoleResponse, err error) {
    if request == nil {
        request = NewUpdateWorkspaceRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateWorkspaceRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkspaceRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkspaceRoleResponse()
    err = c.Send(request, response)
    return
}
