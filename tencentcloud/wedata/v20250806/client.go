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

package v20250806

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2025-08-06"

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


func NewCreateCodeFileRequest() (request *CreateCodeFileRequest) {
    request = &CreateCodeFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateCodeFile")
    
    
    return
}

func NewCreateCodeFileResponse() (response *CreateCodeFileResponse) {
    response = &CreateCodeFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCodeFile
// This API is used to create a code file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodeFile(request *CreateCodeFileRequest) (response *CreateCodeFileResponse, err error) {
    return c.CreateCodeFileWithContext(context.Background(), request)
}

// CreateCodeFile
// This API is used to create a code file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodeFileWithContext(ctx context.Context, request *CreateCodeFileRequest) (response *CreateCodeFileResponse, err error) {
    if request == nil {
        request = NewCreateCodeFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateCodeFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCodeFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCodeFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCodeFolderRequest() (request *CreateCodeFolderRequest) {
    request = &CreateCodeFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateCodeFolder")
    
    
    return
}

func NewCreateCodeFolderResponse() (response *CreateCodeFolderResponse) {
    response = &CreateCodeFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCodeFolder
// This API is used to create a code folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodeFolder(request *CreateCodeFolderRequest) (response *CreateCodeFolderResponse, err error) {
    return c.CreateCodeFolderWithContext(context.Background(), request)
}

// CreateCodeFolder
// This API is used to create a code folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodeFolderWithContext(ctx context.Context, request *CreateCodeFolderRequest) (response *CreateCodeFolderResponse, err error) {
    if request == nil {
        request = NewCreateCodeFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateCodeFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCodeFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCodeFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataBackfillPlanRequest() (request *CreateDataBackfillPlanRequest) {
    request = &CreateDataBackfillPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateDataBackfillPlan")
    
    
    return
}

func NewCreateDataBackfillPlanResponse() (response *CreateDataBackfillPlanResponse) {
    response = &CreateDataBackfillPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataBackfillPlan
// This API is used to create a data backfill plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDataBackfillPlan(request *CreateDataBackfillPlanRequest) (response *CreateDataBackfillPlanResponse, err error) {
    return c.CreateDataBackfillPlanWithContext(context.Background(), request)
}

// CreateDataBackfillPlan
// This API is used to create a data backfill plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDataBackfillPlanWithContext(ctx context.Context, request *CreateDataBackfillPlanRequest) (response *CreateDataBackfillPlanResponse, err error) {
    if request == nil {
        request = NewCreateDataBackfillPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateDataBackfillPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataBackfillPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataBackfillPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpsAlarmRuleRequest() (request *CreateOpsAlarmRuleRequest) {
    request = &CreateOpsAlarmRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateOpsAlarmRule")
    
    
    return
}

func NewCreateOpsAlarmRuleResponse() (response *CreateOpsAlarmRuleResponse) {
    response = &CreateOpsAlarmRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOpsAlarmRule
// This API is used to set alarm rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) CreateOpsAlarmRule(request *CreateOpsAlarmRuleRequest) (response *CreateOpsAlarmRuleResponse, err error) {
    return c.CreateOpsAlarmRuleWithContext(context.Background(), request)
}

// CreateOpsAlarmRule
// This API is used to set alarm rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) CreateOpsAlarmRuleWithContext(ctx context.Context, request *CreateOpsAlarmRuleRequest) (response *CreateOpsAlarmRuleResponse, err error) {
    if request == nil {
        request = NewCreateOpsAlarmRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateOpsAlarmRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpsAlarmRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpsAlarmRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceFileRequest() (request *CreateResourceFileRequest) {
    request = &CreateResourceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateResourceFile")
    
    
    return
}

func NewCreateResourceFileResponse() (response *CreateResourceFileResponse) {
    response = &CreateResourceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResourceFile
// This API is used to create a resource file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceFile(request *CreateResourceFileRequest) (response *CreateResourceFileResponse, err error) {
    return c.CreateResourceFileWithContext(context.Background(), request)
}

// CreateResourceFile
// This API is used to create a resource file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceFileWithContext(ctx context.Context, request *CreateResourceFileRequest) (response *CreateResourceFileResponse, err error) {
    if request == nil {
        request = NewCreateResourceFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateResourceFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourceFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceFolderRequest() (request *CreateResourceFolderRequest) {
    request = &CreateResourceFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateResourceFolder")
    
    
    return
}

func NewCreateResourceFolderResponse() (response *CreateResourceFolderResponse) {
    response = &CreateResourceFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResourceFolder
// This API is used to create a resource file folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceFolder(request *CreateResourceFolderRequest) (response *CreateResourceFolderResponse, err error) {
    return c.CreateResourceFolderWithContext(context.Background(), request)
}

// CreateResourceFolder
// This API is used to create a resource file folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceFolderWithContext(ctx context.Context, request *CreateResourceFolderRequest) (response *CreateResourceFolderResponse, err error) {
    if request == nil {
        request = NewCreateResourceFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateResourceFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourceFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourceFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSQLFolderRequest() (request *CreateSQLFolderRequest) {
    request = &CreateSQLFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateSQLFolder")
    
    
    return
}

func NewCreateSQLFolderResponse() (response *CreateSQLFolderResponse) {
    response = &CreateSQLFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSQLFolder
// This API is used to create an SQL folder
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSQLFolder(request *CreateSQLFolderRequest) (response *CreateSQLFolderResponse, err error) {
    return c.CreateSQLFolderWithContext(context.Background(), request)
}

// CreateSQLFolder
// This API is used to create an SQL folder
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSQLFolderWithContext(ctx context.Context, request *CreateSQLFolderRequest) (response *CreateSQLFolderResponse, err error) {
    if request == nil {
        request = NewCreateSQLFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateSQLFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSQLFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSQLFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSQLScriptRequest() (request *CreateSQLScriptRequest) {
    request = &CreateSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateSQLScript")
    
    
    return
}

func NewCreateSQLScriptResponse() (response *CreateSQLScriptResponse) {
    response = &CreateSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSQLScript
// This API is used to add an SQL script.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSQLScript(request *CreateSQLScriptRequest) (response *CreateSQLScriptResponse, err error) {
    return c.CreateSQLScriptWithContext(context.Background(), request)
}

// CreateSQLScript
// This API is used to add an SQL script.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSQLScriptWithContext(ctx context.Context, request *CreateSQLScriptRequest) (response *CreateSQLScriptResponse, err error) {
    if request == nil {
        request = NewCreateSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTask
// This API is used to create a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    return c.CreateTaskWithContext(context.Background(), request)
}

// CreateTask
// This API is used to create a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRequest() (request *CreateWorkflowRequest) {
    request = &CreateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateWorkflow")
    
    
    return
}

func NewCreateWorkflowResponse() (response *CreateWorkflowResponse) {
    response = &CreateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflow
// This API is used to create workflow.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflow(request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    return c.CreateWorkflowWithContext(context.Background(), request)
}

// CreateWorkflow
// This API is used to create workflow.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflowWithContext(ctx context.Context, request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowFolderRequest() (request *CreateWorkflowFolderRequest) {
    request = &CreateWorkflowFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateWorkflowFolder")
    
    
    return
}

func NewCreateWorkflowFolderResponse() (response *CreateWorkflowFolderResponse) {
    response = &CreateWorkflowFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflowFolder
// This API is used to create a workflow folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateWorkflowFolder(request *CreateWorkflowFolderRequest) (response *CreateWorkflowFolderResponse, err error) {
    return c.CreateWorkflowFolderWithContext(context.Background(), request)
}

// CreateWorkflowFolder
// This API is used to create a workflow folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateWorkflowFolderWithContext(ctx context.Context, request *CreateWorkflowFolderRequest) (response *CreateWorkflowFolderResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateWorkflowFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflowFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCodeFileRequest() (request *DeleteCodeFileRequest) {
    request = &DeleteCodeFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteCodeFile")
    
    
    return
}

func NewDeleteCodeFileResponse() (response *DeleteCodeFileResponse) {
    response = &DeleteCodeFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCodeFile
// This API is used to delete a code file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodeFile(request *DeleteCodeFileRequest) (response *DeleteCodeFileResponse, err error) {
    return c.DeleteCodeFileWithContext(context.Background(), request)
}

// DeleteCodeFile
// This API is used to delete a code file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodeFileWithContext(ctx context.Context, request *DeleteCodeFileRequest) (response *DeleteCodeFileResponse, err error) {
    if request == nil {
        request = NewDeleteCodeFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteCodeFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCodeFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCodeFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCodeFolderRequest() (request *DeleteCodeFolderRequest) {
    request = &DeleteCodeFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteCodeFolder")
    
    
    return
}

func NewDeleteCodeFolderResponse() (response *DeleteCodeFolderResponse) {
    response = &DeleteCodeFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCodeFolder
// This API is used to delete folders in data exploration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodeFolder(request *DeleteCodeFolderRequest) (response *DeleteCodeFolderResponse, err error) {
    return c.DeleteCodeFolderWithContext(context.Background(), request)
}

// DeleteCodeFolder
// This API is used to delete folders in data exploration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodeFolderWithContext(ctx context.Context, request *DeleteCodeFolderRequest) (response *DeleteCodeFolderResponse, err error) {
    if request == nil {
        request = NewDeleteCodeFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteCodeFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCodeFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCodeFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOpsAlarmRuleRequest() (request *DeleteOpsAlarmRuleRequest) {
    request = &DeleteOpsAlarmRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteOpsAlarmRule")
    
    
    return
}

func NewDeleteOpsAlarmRuleResponse() (response *DeleteOpsAlarmRuleResponse) {
    response = &DeleteOpsAlarmRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOpsAlarmRule
// Deletes alarm rules
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteOpsAlarmRule(request *DeleteOpsAlarmRuleRequest) (response *DeleteOpsAlarmRuleResponse, err error) {
    return c.DeleteOpsAlarmRuleWithContext(context.Background(), request)
}

// DeleteOpsAlarmRule
// Deletes alarm rules
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteOpsAlarmRuleWithContext(ctx context.Context, request *DeleteOpsAlarmRuleRequest) (response *DeleteOpsAlarmRuleResponse, err error) {
    if request == nil {
        request = NewDeleteOpsAlarmRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteOpsAlarmRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOpsAlarmRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOpsAlarmRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceFileRequest() (request *DeleteResourceFileRequest) {
    request = &DeleteResourceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteResourceFile")
    
    
    return
}

func NewDeleteResourceFileResponse() (response *DeleteResourceFileResponse) {
    response = &DeleteResourceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceFile
// This API is used to delete a resource file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceFile(request *DeleteResourceFileRequest) (response *DeleteResourceFileResponse, err error) {
    return c.DeleteResourceFileWithContext(context.Background(), request)
}

// DeleteResourceFile
// This API is used to delete a resource file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceFileWithContext(ctx context.Context, request *DeleteResourceFileRequest) (response *DeleteResourceFileResponse, err error) {
    if request == nil {
        request = NewDeleteResourceFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteResourceFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceFolderRequest() (request *DeleteResourceFolderRequest) {
    request = &DeleteResourceFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteResourceFolder")
    
    
    return
}

func NewDeleteResourceFolderResponse() (response *DeleteResourceFolderResponse) {
    response = &DeleteResourceFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceFolder
// This API is used to delete a resource folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceFolder(request *DeleteResourceFolderRequest) (response *DeleteResourceFolderResponse, err error) {
    return c.DeleteResourceFolderWithContext(context.Background(), request)
}

// DeleteResourceFolder
// This API is used to delete a resource folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceFolderWithContext(ctx context.Context, request *DeleteResourceFolderRequest) (response *DeleteResourceFolderResponse, err error) {
    if request == nil {
        request = NewDeleteResourceFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteResourceFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSQLFolderRequest() (request *DeleteSQLFolderRequest) {
    request = &DeleteSQLFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteSQLFolder")
    
    
    return
}

func NewDeleteSQLFolderResponse() (response *DeleteSQLFolderResponse) {
    response = &DeleteSQLFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSQLFolder
// This API is used to delete a SQL folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSQLFolder(request *DeleteSQLFolderRequest) (response *DeleteSQLFolderResponse, err error) {
    return c.DeleteSQLFolderWithContext(context.Background(), request)
}

// DeleteSQLFolder
// This API is used to delete a SQL folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSQLFolderWithContext(ctx context.Context, request *DeleteSQLFolderRequest) (response *DeleteSQLFolderResponse, err error) {
    if request == nil {
        request = NewDeleteSQLFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteSQLFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSQLFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSQLFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSQLScriptRequest() (request *DeleteSQLScriptRequest) {
    request = &DeleteSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteSQLScript")
    
    
    return
}

func NewDeleteSQLScriptResponse() (response *DeleteSQLScriptResponse) {
    response = &DeleteSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSQLScript
// This API is used to delete an SQL script.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSQLScript(request *DeleteSQLScriptRequest) (response *DeleteSQLScriptResponse, err error) {
    return c.DeleteSQLScriptWithContext(context.Background(), request)
}

// DeleteSQLScript
// This API is used to delete an SQL script.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSQLScriptWithContext(ctx context.Context, request *DeleteSQLScriptRequest) (response *DeleteSQLScriptResponse, err error) {
    if request == nil {
        request = NewDeleteSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
    request = &DeleteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteTask")
    
    
    return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
    response = &DeleteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTask
// This API is used to delete an orchestration space task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    return c.DeleteTaskWithContext(context.Background(), request)
}

// DeleteTask
// This API is used to delete an orchestration space task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteTaskWithContext(ctx context.Context, request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowRequest() (request *DeleteWorkflowRequest) {
    request = &DeleteWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteWorkflow")
    
    
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    return c.DeleteWorkflowWithContext(context.Background(), request)
}

// DeleteWorkflow
// Deletes a workflow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteWorkflowWithContext(ctx context.Context, request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowFolderRequest() (request *DeleteWorkflowFolderRequest) {
    request = &DeleteWorkflowFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteWorkflowFolder")
    
    
    return
}

func NewDeleteWorkflowFolderResponse() (response *DeleteWorkflowFolderResponse) {
    response = &DeleteWorkflowFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkflowFolder
// This API is used to delete a data development folder
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWorkflowFolder(request *DeleteWorkflowFolderRequest) (response *DeleteWorkflowFolderResponse, err error) {
    return c.DeleteWorkflowFolderWithContext(context.Background(), request)
}

// DeleteWorkflowFolder
// This API is used to delete a data development folder
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWorkflowFolderWithContext(ctx context.Context, request *DeleteWorkflowFolderRequest) (response *DeleteWorkflowFolderResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteWorkflowFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflowFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowFolderResponse()
    err = c.Send(request, response)
    return
}

func NewGetAlarmMessageRequest() (request *GetAlarmMessageRequest) {
    request = &GetAlarmMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetAlarmMessage")
    
    
    return
}

func NewGetAlarmMessageResponse() (response *GetAlarmMessageResponse) {
    response = &GetAlarmMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAlarmMessage
// This API is used to query alert information details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAlarmMessage(request *GetAlarmMessageRequest) (response *GetAlarmMessageResponse, err error) {
    return c.GetAlarmMessageWithContext(context.Background(), request)
}

// GetAlarmMessage
// This API is used to query alert information details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAlarmMessageWithContext(ctx context.Context, request *GetAlarmMessageRequest) (response *GetAlarmMessageResponse, err error) {
    if request == nil {
        request = NewGetAlarmMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetAlarmMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAlarmMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAlarmMessageResponse()
    err = c.Send(request, response)
    return
}

func NewGetCodeFileRequest() (request *GetCodeFileRequest) {
    request = &GetCodeFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetCodeFile")
    
    
    return
}

func NewGetCodeFileResponse() (response *GetCodeFileResponse) {
    response = &GetCodeFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCodeFile
// This API is used to view code file details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetCodeFile(request *GetCodeFileRequest) (response *GetCodeFileResponse, err error) {
    return c.GetCodeFileWithContext(context.Background(), request)
}

// GetCodeFile
// This API is used to view code file details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetCodeFileWithContext(ctx context.Context, request *GetCodeFileRequest) (response *GetCodeFileResponse, err error) {
    if request == nil {
        request = NewGetCodeFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetCodeFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCodeFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCodeFileResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsAlarmRuleRequest() (request *GetOpsAlarmRuleRequest) {
    request = &GetOpsAlarmRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsAlarmRule")
    
    
    return
}

func NewGetOpsAlarmRuleResponse() (response *GetOpsAlarmRuleResponse) {
    response = &GetOpsAlarmRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsAlarmRule
// This API is used to query alert rule information based on alarm rule id or name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOpsAlarmRule(request *GetOpsAlarmRuleRequest) (response *GetOpsAlarmRuleResponse, err error) {
    return c.GetOpsAlarmRuleWithContext(context.Background(), request)
}

// GetOpsAlarmRule
// This API is used to query alert rule information based on alarm rule id or name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOpsAlarmRuleWithContext(ctx context.Context, request *GetOpsAlarmRuleRequest) (response *GetOpsAlarmRuleResponse, err error) {
    if request == nil {
        request = NewGetOpsAlarmRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsAlarmRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsAlarmRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsAlarmRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsAsyncJobRequest() (request *GetOpsAsyncJobRequest) {
    request = &GetOpsAsyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsAsyncJob")
    
    
    return
}

func NewGetOpsAsyncJobResponse() (response *GetOpsAsyncJobResponse) {
    response = &GetOpsAsyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsAsyncJob
// This API is used to query async operation details in Ops center.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetOpsAsyncJob(request *GetOpsAsyncJobRequest) (response *GetOpsAsyncJobResponse, err error) {
    return c.GetOpsAsyncJobWithContext(context.Background(), request)
}

// GetOpsAsyncJob
// This API is used to query async operation details in Ops center.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetOpsAsyncJobWithContext(ctx context.Context, request *GetOpsAsyncJobRequest) (response *GetOpsAsyncJobResponse, err error) {
    if request == nil {
        request = NewGetOpsAsyncJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsAsyncJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsAsyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsAsyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsTaskRequest() (request *GetOpsTaskRequest) {
    request = &GetOpsTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsTask")
    
    
    return
}

func NewGetOpsTaskResponse() (response *GetOpsTaskResponse) {
    response = &GetOpsTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsTask
// Obtaining Task Details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetOpsTask(request *GetOpsTaskRequest) (response *GetOpsTaskResponse, err error) {
    return c.GetOpsTaskWithContext(context.Background(), request)
}

// GetOpsTask
// Obtaining Task Details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetOpsTaskWithContext(ctx context.Context, request *GetOpsTaskRequest) (response *GetOpsTaskResponse, err error) {
    if request == nil {
        request = NewGetOpsTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsTaskResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsTaskCodeRequest() (request *GetOpsTaskCodeRequest) {
    request = &GetOpsTaskCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsTaskCode")
    
    
    return
}

func NewGetOpsTaskCodeResponse() (response *GetOpsTaskCodeResponse) {
    response = &GetOpsTaskCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsTaskCode
// This API is used to retrieve task code.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetOpsTaskCode(request *GetOpsTaskCodeRequest) (response *GetOpsTaskCodeResponse, err error) {
    return c.GetOpsTaskCodeWithContext(context.Background(), request)
}

// GetOpsTaskCode
// This API is used to retrieve task code.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetOpsTaskCodeWithContext(ctx context.Context, request *GetOpsTaskCodeRequest) (response *GetOpsTaskCodeResponse, err error) {
    if request == nil {
        request = NewGetOpsTaskCodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsTaskCode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsTaskCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsTaskCodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsWorkflowRequest() (request *GetOpsWorkflowRequest) {
    request = &GetOpsWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsWorkflow")
    
    
    return
}

func NewGetOpsWorkflowResponse() (response *GetOpsWorkflowResponse) {
    response = &GetOpsWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsWorkflow
// This API is used to obtain workflow scheduling details based on the workflow id.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOpsWorkflow(request *GetOpsWorkflowRequest) (response *GetOpsWorkflowResponse, err error) {
    return c.GetOpsWorkflowWithContext(context.Background(), request)
}

// GetOpsWorkflow
// This API is used to obtain workflow scheduling details based on the workflow id.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOpsWorkflowWithContext(ctx context.Context, request *GetOpsWorkflowRequest) (response *GetOpsWorkflowResponse, err error) {
    if request == nil {
        request = NewGetOpsWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewGetResourceFileRequest() (request *GetResourceFileRequest) {
    request = &GetResourceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetResourceFile")
    
    
    return
}

func NewGetResourceFileResponse() (response *GetResourceFileResponse) {
    response = &GetResourceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetResourceFile
// This API is used to obtain resource file details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResourceFile(request *GetResourceFileRequest) (response *GetResourceFileResponse, err error) {
    return c.GetResourceFileWithContext(context.Background(), request)
}

// GetResourceFile
// This API is used to obtain resource file details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResourceFileWithContext(ctx context.Context, request *GetResourceFileRequest) (response *GetResourceFileResponse, err error) {
    if request == nil {
        request = NewGetResourceFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetResourceFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetResourceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetResourceFileResponse()
    err = c.Send(request, response)
    return
}

func NewGetSQLScriptRequest() (request *GetSQLScriptRequest) {
    request = &GetSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetSQLScript")
    
    
    return
}

func NewGetSQLScriptResponse() (response *GetSQLScriptResponse) {
    response = &GetSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSQLScript
// This API is used to query script details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetSQLScript(request *GetSQLScriptRequest) (response *GetSQLScriptResponse, err error) {
    return c.GetSQLScriptWithContext(context.Background(), request)
}

// GetSQLScript
// This API is used to query script details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetSQLScriptWithContext(ctx context.Context, request *GetSQLScriptRequest) (response *GetSQLScriptResponse, err error) {
    if request == nil {
        request = NewGetSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskRequest() (request *GetTaskRequest) {
    request = &GetTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTask")
    
    
    return
}

func NewGetTaskResponse() (response *GetTaskResponse) {
    response = &GetTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTask
// This API is used to retrieve task details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTask(request *GetTaskRequest) (response *GetTaskResponse, err error) {
    return c.GetTaskWithContext(context.Background(), request)
}

// GetTask
// This API is used to retrieve task details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskWithContext(ctx context.Context, request *GetTaskRequest) (response *GetTaskResponse, err error) {
    if request == nil {
        request = NewGetTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskCodeRequest() (request *GetTaskCodeRequest) {
    request = &GetTaskCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTaskCode")
    
    
    return
}

func NewGetTaskCodeResponse() (response *GetTaskCodeResponse) {
    response = &GetTaskCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskCode
// This API is used to obtain task code.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskCode(request *GetTaskCodeRequest) (response *GetTaskCodeResponse, err error) {
    return c.GetTaskCodeWithContext(context.Background(), request)
}

// GetTaskCode
// This API is used to obtain task code.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskCodeWithContext(ctx context.Context, request *GetTaskCodeRequest) (response *GetTaskCodeResponse, err error) {
    if request == nil {
        request = NewGetTaskCodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTaskCode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskCodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskInstanceRequest() (request *GetTaskInstanceRequest) {
    request = &GetTaskInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTaskInstance")
    
    
    return
}

func NewGetTaskInstanceResponse() (response *GetTaskInstanceResponse) {
    response = &GetTaskInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskInstance
// This API is used to query the details of a scheduling instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskInstance(request *GetTaskInstanceRequest) (response *GetTaskInstanceResponse, err error) {
    return c.GetTaskInstanceWithContext(context.Background(), request)
}

// GetTaskInstance
// This API is used to query the details of a scheduling instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskInstanceWithContext(ctx context.Context, request *GetTaskInstanceRequest) (response *GetTaskInstanceResponse, err error) {
    if request == nil {
        request = NewGetTaskInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTaskInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskInstanceLogRequest() (request *GetTaskInstanceLogRequest) {
    request = &GetTaskInstanceLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTaskInstanceLog")
    
    
    return
}

func NewGetTaskInstanceLogResponse() (response *GetTaskInstanceLogResponse) {
    response = &GetTaskInstanceLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskInstanceLog
// This API is used to obtain instance lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskInstanceLog(request *GetTaskInstanceLogRequest) (response *GetTaskInstanceLogResponse, err error) {
    return c.GetTaskInstanceLogWithContext(context.Background(), request)
}

// GetTaskInstanceLog
// This API is used to obtain instance lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskInstanceLogWithContext(ctx context.Context, request *GetTaskInstanceLogRequest) (response *GetTaskInstanceLogResponse, err error) {
    if request == nil {
        request = NewGetTaskInstanceLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTaskInstanceLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskInstanceLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskInstanceLogResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskVersionRequest() (request *GetTaskVersionRequest) {
    request = &GetTaskVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTaskVersion")
    
    
    return
}

func NewGetTaskVersionResponse() (response *GetTaskVersionResponse) {
    response = &GetTaskVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskVersion
// This API is used to list task versions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskVersion(request *GetTaskVersionRequest) (response *GetTaskVersionResponse, err error) {
    return c.GetTaskVersionWithContext(context.Background(), request)
}

// GetTaskVersion
// This API is used to list task versions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskVersionWithContext(ctx context.Context, request *GetTaskVersionRequest) (response *GetTaskVersionResponse, err error) {
    if request == nil {
        request = NewGetTaskVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTaskVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowRequest() (request *GetWorkflowRequest) {
    request = &GetWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetWorkflow")
    
    
    return
}

func NewGetWorkflowResponse() (response *GetWorkflowResponse) {
    response = &GetWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkflow
// This API is used to retrieve workflow information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetWorkflow(request *GetWorkflowRequest) (response *GetWorkflowResponse, err error) {
    return c.GetWorkflowWithContext(context.Background(), request)
}

// GetWorkflow
// This API is used to retrieve workflow information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetWorkflowWithContext(ctx context.Context, request *GetWorkflowRequest) (response *GetWorkflowResponse, err error) {
    if request == nil {
        request = NewGetWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewKillTaskInstancesAsyncRequest() (request *KillTaskInstancesAsyncRequest) {
    request = &KillTaskInstancesAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "KillTaskInstancesAsync")
    
    
    return
}

func NewKillTaskInstancesAsyncResponse() (response *KillTaskInstancesAsyncResponse) {
    response = &KillTaskInstancesAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillTaskInstancesAsync
// This API is used to terminate instances in batches asynchronously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) KillTaskInstancesAsync(request *KillTaskInstancesAsyncRequest) (response *KillTaskInstancesAsyncResponse, err error) {
    return c.KillTaskInstancesAsyncWithContext(context.Background(), request)
}

// KillTaskInstancesAsync
// This API is used to terminate instances in batches asynchronously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) KillTaskInstancesAsyncWithContext(ctx context.Context, request *KillTaskInstancesAsyncRequest) (response *KillTaskInstancesAsyncResponse, err error) {
    if request == nil {
        request = NewKillTaskInstancesAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "KillTaskInstancesAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillTaskInstancesAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillTaskInstancesAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewListAlarmMessagesRequest() (request *ListAlarmMessagesRequest) {
    request = &ListAlarmMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListAlarmMessages")
    
    
    return
}

func NewListAlarmMessagesResponse() (response *ListAlarmMessagesResponse) {
    response = &ListAlarmMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAlarmMessages
// This API is used to search the alarm information list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAlarmMessages(request *ListAlarmMessagesRequest) (response *ListAlarmMessagesResponse, err error) {
    return c.ListAlarmMessagesWithContext(context.Background(), request)
}

// ListAlarmMessages
// This API is used to search the alarm information list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAlarmMessagesWithContext(ctx context.Context, request *ListAlarmMessagesRequest) (response *ListAlarmMessagesResponse, err error) {
    if request == nil {
        request = NewListAlarmMessagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListAlarmMessages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAlarmMessages require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAlarmMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewListCodeFolderContentsRequest() (request *ListCodeFolderContentsRequest) {
    request = &ListCodeFolderContentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListCodeFolderContents")
    
    
    return
}

func NewListCodeFolderContentsResponse() (response *ListCodeFolderContentsResponse) {
    response = &ListCodeFolderContentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListCodeFolderContents
// This API is used to get folder content.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCodeFolderContents(request *ListCodeFolderContentsRequest) (response *ListCodeFolderContentsResponse, err error) {
    return c.ListCodeFolderContentsWithContext(context.Background(), request)
}

// ListCodeFolderContents
// This API is used to get folder content.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCodeFolderContentsWithContext(ctx context.Context, request *ListCodeFolderContentsRequest) (response *ListCodeFolderContentsResponse, err error) {
    if request == nil {
        request = NewListCodeFolderContentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListCodeFolderContents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCodeFolderContents require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCodeFolderContentsResponse()
    err = c.Send(request, response)
    return
}

func NewListDataBackfillInstancesRequest() (request *ListDataBackfillInstancesRequest) {
    request = &ListDataBackfillInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDataBackfillInstances")
    
    
    return
}

func NewListDataBackfillInstancesResponse() (response *ListDataBackfillInstancesResponse) {
    response = &ListDataBackfillInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDataBackfillInstances
// This API is used to retrieve all instances of a single backfill.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDataBackfillInstances(request *ListDataBackfillInstancesRequest) (response *ListDataBackfillInstancesResponse, err error) {
    return c.ListDataBackfillInstancesWithContext(context.Background(), request)
}

// ListDataBackfillInstances
// This API is used to retrieve all instances of a single backfill.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDataBackfillInstancesWithContext(ctx context.Context, request *ListDataBackfillInstancesRequest) (response *ListDataBackfillInstancesResponse, err error) {
    if request == nil {
        request = NewListDataBackfillInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDataBackfillInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDataBackfillInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDataBackfillInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListDownstreamOpsTasksRequest() (request *ListDownstreamOpsTasksRequest) {
    request = &ListDownstreamOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDownstreamOpsTasks")
    
    
    return
}

func NewListDownstreamOpsTasksResponse() (response *ListDownstreamOpsTasksResponse) {
    response = &ListDownstreamOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDownstreamOpsTasks
// This API is used to retrieve task direct downstream details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamOpsTasks(request *ListDownstreamOpsTasksRequest) (response *ListDownstreamOpsTasksResponse, err error) {
    return c.ListDownstreamOpsTasksWithContext(context.Background(), request)
}

// ListDownstreamOpsTasks
// This API is used to retrieve task direct downstream details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamOpsTasksWithContext(ctx context.Context, request *ListDownstreamOpsTasksRequest) (response *ListDownstreamOpsTasksResponse, err error) {
    if request == nil {
        request = NewListDownstreamOpsTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDownstreamOpsTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDownstreamOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDownstreamOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListDownstreamTaskInstancesRequest() (request *ListDownstreamTaskInstancesRequest) {
    request = &ListDownstreamTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDownstreamTaskInstances")
    
    
    return
}

func NewListDownstreamTaskInstancesResponse() (response *ListDownstreamTaskInstancesResponse) {
    response = &ListDownstreamTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDownstreamTaskInstances
// This API is used to get the direct upstream of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamTaskInstances(request *ListDownstreamTaskInstancesRequest) (response *ListDownstreamTaskInstancesResponse, err error) {
    return c.ListDownstreamTaskInstancesWithContext(context.Background(), request)
}

// ListDownstreamTaskInstances
// This API is used to get the direct upstream of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamTaskInstancesWithContext(ctx context.Context, request *ListDownstreamTaskInstancesRequest) (response *ListDownstreamTaskInstancesResponse, err error) {
    if request == nil {
        request = NewListDownstreamTaskInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDownstreamTaskInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDownstreamTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDownstreamTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListDownstreamTasksRequest() (request *ListDownstreamTasksRequest) {
    request = &ListDownstreamTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDownstreamTasks")
    
    
    return
}

func NewListDownstreamTasksResponse() (response *ListDownstreamTasksResponse) {
    response = &ListDownstreamTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDownstreamTasks
// This API is used to retrieve the direct downstream task details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamTasks(request *ListDownstreamTasksRequest) (response *ListDownstreamTasksResponse, err error) {
    return c.ListDownstreamTasksWithContext(context.Background(), request)
}

// ListDownstreamTasks
// This API is used to retrieve the direct downstream task details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamTasksWithContext(ctx context.Context, request *ListDownstreamTasksRequest) (response *ListDownstreamTasksResponse, err error) {
    if request == nil {
        request = NewListDownstreamTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDownstreamTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDownstreamTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDownstreamTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListOpsAlarmRulesRequest() (request *ListOpsAlarmRulesRequest) {
    request = &ListOpsAlarmRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListOpsAlarmRules")
    
    
    return
}

func NewListOpsAlarmRulesResponse() (response *ListOpsAlarmRulesResponse) {
    response = &ListOpsAlarmRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOpsAlarmRules
// This API is used to query the alarm rule list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListOpsAlarmRules(request *ListOpsAlarmRulesRequest) (response *ListOpsAlarmRulesResponse, err error) {
    return c.ListOpsAlarmRulesWithContext(context.Background(), request)
}

// ListOpsAlarmRules
// This API is used to query the alarm rule list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListOpsAlarmRulesWithContext(ctx context.Context, request *ListOpsAlarmRulesRequest) (response *ListOpsAlarmRulesResponse, err error) {
    if request == nil {
        request = NewListOpsAlarmRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListOpsAlarmRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOpsAlarmRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOpsAlarmRulesResponse()
    err = c.Send(request, response)
    return
}

func NewListOpsTasksRequest() (request *ListOpsTasksRequest) {
    request = &ListOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListOpsTasks")
    
    
    return
}

func NewListOpsTasksResponse() (response *ListOpsTasksResponse) {
    response = &ListOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOpsTasks
// This API is used to obtain the task list based on the project id.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListOpsTasks(request *ListOpsTasksRequest) (response *ListOpsTasksResponse, err error) {
    return c.ListOpsTasksWithContext(context.Background(), request)
}

// ListOpsTasks
// This API is used to obtain the task list based on the project id.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListOpsTasksWithContext(ctx context.Context, request *ListOpsTasksRequest) (response *ListOpsTasksResponse, err error) {
    if request == nil {
        request = NewListOpsTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListOpsTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListOpsWorkflowsRequest() (request *ListOpsWorkflowsRequest) {
    request = &ListOpsWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListOpsWorkflows")
    
    
    return
}

func NewListOpsWorkflowsResponse() (response *ListOpsWorkflowsResponse) {
    response = &ListOpsWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOpsWorkflows
// Get Workflows under a Project by Project ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListOpsWorkflows(request *ListOpsWorkflowsRequest) (response *ListOpsWorkflowsResponse, err error) {
    return c.ListOpsWorkflowsWithContext(context.Background(), request)
}

// ListOpsWorkflows
// Get Workflows under a Project by Project ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListOpsWorkflowsWithContext(ctx context.Context, request *ListOpsWorkflowsRequest) (response *ListOpsWorkflowsResponse, err error) {
    if request == nil {
        request = NewListOpsWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListOpsWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOpsWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOpsWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewListResourceFilesRequest() (request *ListResourceFilesRequest) {
    request = &ListResourceFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListResourceFiles")
    
    
    return
}

func NewListResourceFilesResponse() (response *ListResourceFilesResponse) {
    response = &ListResourceFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListResourceFiles
// This API is used to view resource file list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListResourceFiles(request *ListResourceFilesRequest) (response *ListResourceFilesResponse, err error) {
    return c.ListResourceFilesWithContext(context.Background(), request)
}

// ListResourceFiles
// This API is used to view resource file list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListResourceFilesWithContext(ctx context.Context, request *ListResourceFilesRequest) (response *ListResourceFilesResponse, err error) {
    if request == nil {
        request = NewListResourceFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListResourceFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListResourceFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewListResourceFilesResponse()
    err = c.Send(request, response)
    return
}

func NewListResourceFoldersRequest() (request *ListResourceFoldersRequest) {
    request = &ListResourceFoldersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListResourceFolders")
    
    
    return
}

func NewListResourceFoldersResponse() (response *ListResourceFoldersResponse) {
    response = &ListResourceFoldersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListResourceFolders
// This API is used to query the resource file folder list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListResourceFolders(request *ListResourceFoldersRequest) (response *ListResourceFoldersResponse, err error) {
    return c.ListResourceFoldersWithContext(context.Background(), request)
}

// ListResourceFolders
// This API is used to query the resource file folder list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListResourceFoldersWithContext(ctx context.Context, request *ListResourceFoldersRequest) (response *ListResourceFoldersResponse, err error) {
    if request == nil {
        request = NewListResourceFoldersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListResourceFolders")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListResourceFolders require credential")
    }

    request.SetContext(ctx)
    
    response = NewListResourceFoldersResponse()
    err = c.Send(request, response)
    return
}

func NewListSQLFolderContentsRequest() (request *ListSQLFolderContentsRequest) {
    request = &ListSQLFolderContentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListSQLFolderContents")
    
    
    return
}

func NewListSQLFolderContentsResponse() (response *ListSQLFolderContentsResponse) {
    response = &ListSQLFolderContentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSQLFolderContents
// This API is used to retrieve the content list of an sql folder
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSQLFolderContents(request *ListSQLFolderContentsRequest) (response *ListSQLFolderContentsResponse, err error) {
    return c.ListSQLFolderContentsWithContext(context.Background(), request)
}

// ListSQLFolderContents
// This API is used to retrieve the content list of an sql folder
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSQLFolderContentsWithContext(ctx context.Context, request *ListSQLFolderContentsRequest) (response *ListSQLFolderContentsResponse, err error) {
    if request == nil {
        request = NewListSQLFolderContentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListSQLFolderContents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSQLFolderContents require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSQLFolderContentsResponse()
    err = c.Send(request, response)
    return
}

func NewListSQLScriptRunsRequest() (request *ListSQLScriptRunsRequest) {
    request = &ListSQLScriptRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListSQLScriptRuns")
    
    
    return
}

func NewListSQLScriptRunsResponse() (response *ListSQLScriptRunsResponse) {
    response = &ListSQLScriptRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSQLScriptRuns
// This API is used to query SQL run history.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSQLScriptRuns(request *ListSQLScriptRunsRequest) (response *ListSQLScriptRunsResponse, err error) {
    return c.ListSQLScriptRunsWithContext(context.Background(), request)
}

// ListSQLScriptRuns
// This API is used to query SQL run history.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSQLScriptRunsWithContext(ctx context.Context, request *ListSQLScriptRunsRequest) (response *ListSQLScriptRunsResponse, err error) {
    if request == nil {
        request = NewListSQLScriptRunsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListSQLScriptRuns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSQLScriptRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSQLScriptRunsResponse()
    err = c.Send(request, response)
    return
}

func NewListTaskInstanceExecutionsRequest() (request *ListTaskInstanceExecutionsRequest) {
    request = &ListTaskInstanceExecutionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTaskInstanceExecutions")
    
    
    return
}

func NewListTaskInstanceExecutionsResponse() (response *ListTaskInstanceExecutionsResponse) {
    response = &ListTaskInstanceExecutionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTaskInstanceExecutions
// This API is used to query the details of a scheduling instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskInstanceExecutions(request *ListTaskInstanceExecutionsRequest) (response *ListTaskInstanceExecutionsResponse, err error) {
    return c.ListTaskInstanceExecutionsWithContext(context.Background(), request)
}

// ListTaskInstanceExecutions
// This API is used to query the details of a scheduling instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskInstanceExecutionsWithContext(ctx context.Context, request *ListTaskInstanceExecutionsRequest) (response *ListTaskInstanceExecutionsResponse, err error) {
    if request == nil {
        request = NewListTaskInstanceExecutionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTaskInstanceExecutions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTaskInstanceExecutions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTaskInstanceExecutionsResponse()
    err = c.Send(request, response)
    return
}

func NewListTaskInstancesRequest() (request *ListTaskInstancesRequest) {
    request = &ListTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTaskInstances")
    
    
    return
}

func NewListTaskInstancesResponse() (response *ListTaskInstancesResponse) {
    response = &ListTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTaskInstances
// This API is used to obtain instance lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskInstances(request *ListTaskInstancesRequest) (response *ListTaskInstancesResponse, err error) {
    return c.ListTaskInstancesWithContext(context.Background(), request)
}

// ListTaskInstances
// This API is used to obtain instance lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskInstancesWithContext(ctx context.Context, request *ListTaskInstancesRequest) (response *ListTaskInstancesResponse, err error) {
    if request == nil {
        request = NewListTaskInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTaskInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListTaskVersionsRequest() (request *ListTaskVersionsRequest) {
    request = &ListTaskVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTaskVersions")
    
    
    return
}

func NewListTaskVersionsResponse() (response *ListTaskVersionsResponse) {
    response = &ListTaskVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTaskVersions
// This API is used to list saved task versions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskVersions(request *ListTaskVersionsRequest) (response *ListTaskVersionsResponse, err error) {
    return c.ListTaskVersionsWithContext(context.Background(), request)
}

// ListTaskVersions
// This API is used to list saved task versions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskVersionsWithContext(ctx context.Context, request *ListTaskVersionsRequest) (response *ListTaskVersionsResponse, err error) {
    if request == nil {
        request = NewListTaskVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTaskVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTaskVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTaskVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewListTasksRequest() (request *ListTasksRequest) {
    request = &ListTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTasks")
    
    
    return
}

func NewListTasksResponse() (response *ListTasksResponse) {
    response = &ListTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTasks
// This API is used to query pagination information of tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTasks(request *ListTasksRequest) (response *ListTasksResponse, err error) {
    return c.ListTasksWithContext(context.Background(), request)
}

// ListTasks
// This API is used to query pagination information of tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTasksWithContext(ctx context.Context, request *ListTasksRequest) (response *ListTasksResponse, err error) {
    if request == nil {
        request = NewListTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListUpstreamOpsTasksRequest() (request *ListUpstreamOpsTasksRequest) {
    request = &ListUpstreamOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListUpstreamOpsTasks")
    
    
    return
}

func NewListUpstreamOpsTasksResponse() (response *ListUpstreamOpsTasksResponse) {
    response = &ListUpstreamOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUpstreamOpsTasks
// This API is used to retrieve task direct upstream.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamOpsTasks(request *ListUpstreamOpsTasksRequest) (response *ListUpstreamOpsTasksResponse, err error) {
    return c.ListUpstreamOpsTasksWithContext(context.Background(), request)
}

// ListUpstreamOpsTasks
// This API is used to retrieve task direct upstream.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamOpsTasksWithContext(ctx context.Context, request *ListUpstreamOpsTasksRequest) (response *ListUpstreamOpsTasksResponse, err error) {
    if request == nil {
        request = NewListUpstreamOpsTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListUpstreamOpsTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUpstreamOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUpstreamOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListUpstreamTaskInstancesRequest() (request *ListUpstreamTaskInstancesRequest) {
    request = &ListUpstreamTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListUpstreamTaskInstances")
    
    
    return
}

func NewListUpstreamTaskInstancesResponse() (response *ListUpstreamTaskInstancesResponse) {
    response = &ListUpstreamTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUpstreamTaskInstances
// This API is used to get the direct upstream of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamTaskInstances(request *ListUpstreamTaskInstancesRequest) (response *ListUpstreamTaskInstancesResponse, err error) {
    return c.ListUpstreamTaskInstancesWithContext(context.Background(), request)
}

// ListUpstreamTaskInstances
// This API is used to get the direct upstream of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamTaskInstancesWithContext(ctx context.Context, request *ListUpstreamTaskInstancesRequest) (response *ListUpstreamTaskInstancesResponse, err error) {
    if request == nil {
        request = NewListUpstreamTaskInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListUpstreamTaskInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUpstreamTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUpstreamTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListUpstreamTasksRequest() (request *ListUpstreamTasksRequest) {
    request = &ListUpstreamTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListUpstreamTasks")
    
    
    return
}

func NewListUpstreamTasksResponse() (response *ListUpstreamTasksResponse) {
    response = &ListUpstreamTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUpstreamTasks
// This API is used to retrieve the direct upstream task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamTasks(request *ListUpstreamTasksRequest) (response *ListUpstreamTasksResponse, err error) {
    return c.ListUpstreamTasksWithContext(context.Background(), request)
}

// ListUpstreamTasks
// This API is used to retrieve the direct upstream task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamTasksWithContext(ctx context.Context, request *ListUpstreamTasksRequest) (response *ListUpstreamTasksResponse, err error) {
    if request == nil {
        request = NewListUpstreamTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListUpstreamTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUpstreamTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUpstreamTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowFoldersRequest() (request *ListWorkflowFoldersRequest) {
    request = &ListWorkflowFoldersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListWorkflowFolders")
    
    
    return
}

func NewListWorkflowFoldersResponse() (response *ListWorkflowFoldersResponse) {
    response = &ListWorkflowFoldersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflowFolders
// This API is used to query the folder list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListWorkflowFolders(request *ListWorkflowFoldersRequest) (response *ListWorkflowFoldersResponse, err error) {
    return c.ListWorkflowFoldersWithContext(context.Background(), request)
}

// ListWorkflowFolders
// This API is used to query the folder list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListWorkflowFoldersWithContext(ctx context.Context, request *ListWorkflowFoldersRequest) (response *ListWorkflowFoldersResponse, err error) {
    if request == nil {
        request = NewListWorkflowFoldersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListWorkflowFolders")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflowFolders require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowFoldersResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowsRequest() (request *ListWorkflowsRequest) {
    request = &ListWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListWorkflows")
    
    
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflowsWithContext(ctx context.Context, request *ListWorkflowsRequest) (response *ListWorkflowsResponse, err error) {
    if request == nil {
        request = NewListWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewPauseOpsTasksAsyncRequest() (request *PauseOpsTasksAsyncRequest) {
    request = &PauseOpsTasksAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "PauseOpsTasksAsync")
    
    
    return
}

func NewPauseOpsTasksAsyncResponse() (response *PauseOpsTasksAsyncResponse) {
    response = &PauseOpsTasksAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseOpsTasksAsync
// This API is used to pause tasks in batches asynchronously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PauseOpsTasksAsync(request *PauseOpsTasksAsyncRequest) (response *PauseOpsTasksAsyncResponse, err error) {
    return c.PauseOpsTasksAsyncWithContext(context.Background(), request)
}

// PauseOpsTasksAsync
// This API is used to pause tasks in batches asynchronously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PauseOpsTasksAsyncWithContext(ctx context.Context, request *PauseOpsTasksAsyncRequest) (response *PauseOpsTasksAsyncResponse, err error) {
    if request == nil {
        request = NewPauseOpsTasksAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "PauseOpsTasksAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseOpsTasksAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseOpsTasksAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewRerunTaskInstancesAsyncRequest() (request *RerunTaskInstancesAsyncRequest) {
    request = &RerunTaskInstancesAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RerunTaskInstancesAsync")
    
    
    return
}

func NewRerunTaskInstancesAsyncResponse() (response *RerunTaskInstancesAsyncResponse) {
    response = &RerunTaskInstancesAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RerunTaskInstancesAsync
// This API is used to batch rerun instances asynchronously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunTaskInstancesAsync(request *RerunTaskInstancesAsyncRequest) (response *RerunTaskInstancesAsyncResponse, err error) {
    return c.RerunTaskInstancesAsyncWithContext(context.Background(), request)
}

// RerunTaskInstancesAsync
// This API is used to batch rerun instances asynchronously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunTaskInstancesAsyncWithContext(ctx context.Context, request *RerunTaskInstancesAsyncRequest) (response *RerunTaskInstancesAsyncResponse, err error) {
    if request == nil {
        request = NewRerunTaskInstancesAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "RerunTaskInstancesAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RerunTaskInstancesAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewRerunTaskInstancesAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewRunSQLScriptRequest() (request *RunSQLScriptRequest) {
    request = &RunSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RunSQLScript")
    
    
    return
}

func NewRunSQLScriptResponse() (response *RunSQLScriptResponse) {
    response = &RunSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunSQLScript
// This API is used to run an SQL script.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunSQLScript(request *RunSQLScriptRequest) (response *RunSQLScriptResponse, err error) {
    return c.RunSQLScriptWithContext(context.Background(), request)
}

// RunSQLScript
// This API is used to run an SQL script.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunSQLScriptWithContext(ctx context.Context, request *RunSQLScriptRequest) (response *RunSQLScriptResponse, err error) {
    if request == nil {
        request = NewRunSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "RunSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewSetSuccessTaskInstancesAsyncRequest() (request *SetSuccessTaskInstancesAsyncRequest) {
    request = &SetSuccessTaskInstancesAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SetSuccessTaskInstancesAsync")
    
    
    return
}

func NewSetSuccessTaskInstancesAsyncResponse() (response *SetSuccessTaskInstancesAsyncResponse) {
    response = &SetSuccessTaskInstancesAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetSuccessTaskInstancesAsync
// This API is used to batch configure instances asynchronously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SetSuccessTaskInstancesAsync(request *SetSuccessTaskInstancesAsyncRequest) (response *SetSuccessTaskInstancesAsyncResponse, err error) {
    return c.SetSuccessTaskInstancesAsyncWithContext(context.Background(), request)
}

// SetSuccessTaskInstancesAsync
// This API is used to batch configure instances asynchronously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SetSuccessTaskInstancesAsyncWithContext(ctx context.Context, request *SetSuccessTaskInstancesAsyncRequest) (response *SetSuccessTaskInstancesAsyncResponse, err error) {
    if request == nil {
        request = NewSetSuccessTaskInstancesAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "SetSuccessTaskInstancesAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetSuccessTaskInstancesAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetSuccessTaskInstancesAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewStopOpsTasksAsyncRequest() (request *StopOpsTasksAsyncRequest) {
    request = &StopOpsTasksAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "StopOpsTasksAsync")
    
    
    return
}

func NewStopOpsTasksAsyncResponse() (response *StopOpsTasksAsyncResponse) {
    response = &StopOpsTasksAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopOpsTasksAsync
// This API is used to asynchronously deactivate tasks in batch.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopOpsTasksAsync(request *StopOpsTasksAsyncRequest) (response *StopOpsTasksAsyncResponse, err error) {
    return c.StopOpsTasksAsyncWithContext(context.Background(), request)
}

// StopOpsTasksAsync
// This API is used to asynchronously deactivate tasks in batch.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopOpsTasksAsyncWithContext(ctx context.Context, request *StopOpsTasksAsyncRequest) (response *StopOpsTasksAsyncResponse, err error) {
    if request == nil {
        request = NewStopOpsTasksAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "StopOpsTasksAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopOpsTasksAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopOpsTasksAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewStopSQLScriptRunRequest() (request *StopSQLScriptRunRequest) {
    request = &StopSQLScriptRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "StopSQLScriptRun")
    
    
    return
}

func NewStopSQLScriptRunResponse() (response *StopSQLScriptRunResponse) {
    response = &StopSQLScriptRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopSQLScriptRun
// This API is used to stop running an SQL script.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopSQLScriptRun(request *StopSQLScriptRunRequest) (response *StopSQLScriptRunResponse, err error) {
    return c.StopSQLScriptRunWithContext(context.Background(), request)
}

// StopSQLScriptRun
// This API is used to stop running an SQL script.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopSQLScriptRunWithContext(ctx context.Context, request *StopSQLScriptRunRequest) (response *StopSQLScriptRunResponse, err error) {
    if request == nil {
        request = NewStopSQLScriptRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "StopSQLScriptRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopSQLScriptRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopSQLScriptRunResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTaskRequest() (request *SubmitTaskRequest) {
    request = &SubmitTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitTask")
    
    
    return
}

func NewSubmitTaskResponse() (response *SubmitTaskResponse) {
    response = &SubmitTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTask
// This API is used to submit a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitTask(request *SubmitTaskRequest) (response *SubmitTaskResponse, err error) {
    return c.SubmitTaskWithContext(context.Background(), request)
}

// SubmitTask
// This API is used to submit a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitTaskWithContext(ctx context.Context, request *SubmitTaskRequest) (response *SubmitTaskResponse, err error) {
    if request == nil {
        request = NewSubmitTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "SubmitTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCodeFileRequest() (request *UpdateCodeFileRequest) {
    request = &UpdateCodeFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateCodeFile")
    
    
    return
}

func NewUpdateCodeFileResponse() (response *UpdateCodeFileResponse) {
    response = &UpdateCodeFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCodeFile
// This API is used to update a code file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateCodeFile(request *UpdateCodeFileRequest) (response *UpdateCodeFileResponse, err error) {
    return c.UpdateCodeFileWithContext(context.Background(), request)
}

// UpdateCodeFile
// This API is used to update a code file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateCodeFileWithContext(ctx context.Context, request *UpdateCodeFileRequest) (response *UpdateCodeFileResponse, err error) {
    if request == nil {
        request = NewUpdateCodeFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateCodeFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCodeFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCodeFileResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCodeFolderRequest() (request *UpdateCodeFolderRequest) {
    request = &UpdateCodeFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateCodeFolder")
    
    
    return
}

func NewUpdateCodeFolderResponse() (response *UpdateCodeFolderResponse) {
    response = &UpdateCodeFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCodeFolder
// This API is used to rename a code folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateCodeFolder(request *UpdateCodeFolderRequest) (response *UpdateCodeFolderResponse, err error) {
    return c.UpdateCodeFolderWithContext(context.Background(), request)
}

// UpdateCodeFolder
// This API is used to rename a code folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateCodeFolderWithContext(ctx context.Context, request *UpdateCodeFolderRequest) (response *UpdateCodeFolderResponse, err error) {
    if request == nil {
        request = NewUpdateCodeFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateCodeFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCodeFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCodeFolderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOpsAlarmRuleRequest() (request *UpdateOpsAlarmRuleRequest) {
    request = &UpdateOpsAlarmRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateOpsAlarmRule")
    
    
    return
}

func NewUpdateOpsAlarmRuleResponse() (response *UpdateOpsAlarmRuleResponse) {
    response = &UpdateOpsAlarmRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOpsAlarmRule
// Modifies alarm rules
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsAlarmRule(request *UpdateOpsAlarmRuleRequest) (response *UpdateOpsAlarmRuleResponse, err error) {
    return c.UpdateOpsAlarmRuleWithContext(context.Background(), request)
}

// UpdateOpsAlarmRule
// Modifies alarm rules
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsAlarmRuleWithContext(ctx context.Context, request *UpdateOpsAlarmRuleRequest) (response *UpdateOpsAlarmRuleResponse, err error) {
    if request == nil {
        request = NewUpdateOpsAlarmRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateOpsAlarmRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOpsAlarmRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOpsAlarmRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOpsTasksOwnerRequest() (request *UpdateOpsTasksOwnerRequest) {
    request = &UpdateOpsTasksOwnerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateOpsTasksOwner")
    
    
    return
}

func NewUpdateOpsTasksOwnerResponse() (response *UpdateOpsTasksOwnerResponse) {
    response = &UpdateOpsTasksOwnerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOpsTasksOwner
// This API is used to modify the task owner.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsTasksOwner(request *UpdateOpsTasksOwnerRequest) (response *UpdateOpsTasksOwnerResponse, err error) {
    return c.UpdateOpsTasksOwnerWithContext(context.Background(), request)
}

// UpdateOpsTasksOwner
// This API is used to modify the task owner.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsTasksOwnerWithContext(ctx context.Context, request *UpdateOpsTasksOwnerRequest) (response *UpdateOpsTasksOwnerResponse, err error) {
    if request == nil {
        request = NewUpdateOpsTasksOwnerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateOpsTasksOwner")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOpsTasksOwner require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOpsTasksOwnerResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateResourceFileRequest() (request *UpdateResourceFileRequest) {
    request = &UpdateResourceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateResourceFile")
    
    
    return
}

func NewUpdateResourceFileResponse() (response *UpdateResourceFileResponse) {
    response = &UpdateResourceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateResourceFile
// This API is used to update a resource file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceFile(request *UpdateResourceFileRequest) (response *UpdateResourceFileResponse, err error) {
    return c.UpdateResourceFileWithContext(context.Background(), request)
}

// UpdateResourceFile
// This API is used to update a resource file.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceFileWithContext(ctx context.Context, request *UpdateResourceFileRequest) (response *UpdateResourceFileResponse, err error) {
    if request == nil {
        request = NewUpdateResourceFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateResourceFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateResourceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateResourceFileResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateResourceFolderRequest() (request *UpdateResourceFolderRequest) {
    request = &UpdateResourceFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateResourceFolder")
    
    
    return
}

func NewUpdateResourceFolderResponse() (response *UpdateResourceFolderResponse) {
    response = &UpdateResourceFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateResourceFolder
// This API is used to update a resource folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceFolder(request *UpdateResourceFolderRequest) (response *UpdateResourceFolderResponse, err error) {
    return c.UpdateResourceFolderWithContext(context.Background(), request)
}

// UpdateResourceFolder
// This API is used to update a resource folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceFolderWithContext(ctx context.Context, request *UpdateResourceFolderRequest) (response *UpdateResourceFolderResponse, err error) {
    if request == nil {
        request = NewUpdateResourceFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateResourceFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateResourceFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateResourceFolderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSQLFolderRequest() (request *UpdateSQLFolderRequest) {
    request = &UpdateSQLFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateSQLFolder")
    
    
    return
}

func NewUpdateSQLFolderResponse() (response *UpdateSQLFolderResponse) {
    response = &UpdateSQLFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSQLFolder
// This API is used to rename an SQL folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateSQLFolder(request *UpdateSQLFolderRequest) (response *UpdateSQLFolderResponse, err error) {
    return c.UpdateSQLFolderWithContext(context.Background(), request)
}

// UpdateSQLFolder
// This API is used to rename an SQL folder.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateSQLFolderWithContext(ctx context.Context, request *UpdateSQLFolderRequest) (response *UpdateSQLFolderResponse, err error) {
    if request == nil {
        request = NewUpdateSQLFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateSQLFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSQLFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSQLFolderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSQLScriptRequest() (request *UpdateSQLScriptRequest) {
    request = &UpdateSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateSQLScript")
    
    
    return
}

func NewUpdateSQLScriptResponse() (response *UpdateSQLScriptResponse) {
    response = &UpdateSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSQLScript
// This API is used to save the script content for data exploration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateSQLScript(request *UpdateSQLScriptRequest) (response *UpdateSQLScriptResponse, err error) {
    return c.UpdateSQLScriptWithContext(context.Background(), request)
}

// UpdateSQLScript
// This API is used to save the script content for data exploration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateSQLScriptWithContext(ctx context.Context, request *UpdateSQLScriptRequest) (response *UpdateSQLScriptResponse, err error) {
    if request == nil {
        request = NewUpdateSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTaskRequest() (request *UpdateTaskRequest) {
    request = &UpdateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateTask")
    
    
    return
}

func NewUpdateTaskResponse() (response *UpdateTaskResponse) {
    response = &UpdateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTask
// This API is used to update a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTask(request *UpdateTaskRequest) (response *UpdateTaskResponse, err error) {
    return c.UpdateTaskWithContext(context.Background(), request)
}

// UpdateTask
// This API is used to update a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTaskWithContext(ctx context.Context, request *UpdateTaskRequest) (response *UpdateTaskResponse, err error) {
    if request == nil {
        request = NewUpdateTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkflowRequest() (request *UpdateWorkflowRequest) {
    request = &UpdateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateWorkflow")
    
    
    return
}

func NewUpdateWorkflowResponse() (response *UpdateWorkflowResponse) {
    response = &UpdateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkflow
// This API is used to update a workflow including basic information and workflow parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateWorkflow(request *UpdateWorkflowRequest) (response *UpdateWorkflowResponse, err error) {
    return c.UpdateWorkflowWithContext(context.Background(), request)
}

// UpdateWorkflow
// This API is used to update a workflow including basic information and workflow parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateWorkflowWithContext(ctx context.Context, request *UpdateWorkflowRequest) (response *UpdateWorkflowResponse, err error) {
    if request == nil {
        request = NewUpdateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkflowFolderRequest() (request *UpdateWorkflowFolderRequest) {
    request = &UpdateWorkflowFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateWorkflowFolder")
    
    
    return
}

func NewUpdateWorkflowFolderResponse() (response *UpdateWorkflowFolderResponse) {
    response = &UpdateWorkflowFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkflowFolder
// This API is used to update a workflow folder
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateWorkflowFolder(request *UpdateWorkflowFolderRequest) (response *UpdateWorkflowFolderResponse, err error) {
    return c.UpdateWorkflowFolderWithContext(context.Background(), request)
}

// UpdateWorkflowFolder
// This API is used to update a workflow folder
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateWorkflowFolderWithContext(ctx context.Context, request *UpdateWorkflowFolderRequest) (response *UpdateWorkflowFolderResponse, err error) {
    if request == nil {
        request = NewUpdateWorkflowFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateWorkflowFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkflowFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkflowFolderResponse()
    err = c.Send(request, response)
    return
}
