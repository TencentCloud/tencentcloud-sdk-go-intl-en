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

package v20230616

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2023-06-16"

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


func NewCancelAIWorkbenchChatRequest() (request *CancelAIWorkbenchChatRequest) {
    request = &CancelAIWorkbenchChatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CancelAIWorkbenchChat")
    
    
    return
}

func NewCancelAIWorkbenchChatResponse() (response *CancelAIWorkbenchChatResponse) {
    response = &CancelAIWorkbenchChatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelAIWorkbenchChat
// Cancel dialogue execution
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CancelAIWorkbenchChat(request *CancelAIWorkbenchChatRequest) (response *CancelAIWorkbenchChatResponse, err error) {
    return c.CancelAIWorkbenchChatWithContext(context.Background(), request)
}

// CancelAIWorkbenchChat
// Cancel dialogue execution
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CancelAIWorkbenchChatWithContext(ctx context.Context, request *CancelAIWorkbenchChatRequest) (response *CancelAIWorkbenchChatResponse, err error) {
    if request == nil {
        request = NewCancelAIWorkbenchChatRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CancelAIWorkbenchChat")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelAIWorkbenchChat require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelAIWorkbenchChatResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIWorkbenchAgentRequest() (request *CreateAIWorkbenchAgentRequest) {
    request = &CreateAIWorkbenchAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAIWorkbenchAgent")
    
    
    return
}

func NewCreateAIWorkbenchAgentResponse() (response *CreateAIWorkbenchAgentResponse) {
    response = &CreateAIWorkbenchAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAIWorkbenchAgent
// This API is used to create an Agent.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAIWorkbenchAgent(request *CreateAIWorkbenchAgentRequest) (response *CreateAIWorkbenchAgentResponse, err error) {
    return c.CreateAIWorkbenchAgentWithContext(context.Background(), request)
}

// CreateAIWorkbenchAgent
// This API is used to create an Agent.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAIWorkbenchAgentWithContext(ctx context.Context, request *CreateAIWorkbenchAgentRequest) (response *CreateAIWorkbenchAgentResponse, err error) {
    if request == nil {
        request = NewCreateAIWorkbenchAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateAIWorkbenchAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAIWorkbenchAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAIWorkbenchAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIWorkbenchTaskRequest() (request *CreateAIWorkbenchTaskRequest) {
    request = &CreateAIWorkbenchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAIWorkbenchTask")
    
    
    return
}

func NewCreateAIWorkbenchTaskResponse() (response *CreateAIWorkbenchTaskResponse) {
    response = &CreateAIWorkbenchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAIWorkbenchTask
// Create a task
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAIWorkbenchTask(request *CreateAIWorkbenchTaskRequest) (response *CreateAIWorkbenchTaskResponse, err error) {
    return c.CreateAIWorkbenchTaskWithContext(context.Background(), request)
}

// CreateAIWorkbenchTask
// Create a task
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAIWorkbenchTaskWithContext(ctx context.Context, request *CreateAIWorkbenchTaskRequest) (response *CreateAIWorkbenchTaskResponse, err error) {
    if request == nil {
        request = NewCreateAIWorkbenchTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateAIWorkbenchTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAIWorkbenchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAIWorkbenchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIWorkbenchAgentRequest() (request *DeleteAIWorkbenchAgentRequest) {
    request = &DeleteAIWorkbenchAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAIWorkbenchAgent")
    
    
    return
}

func NewDeleteAIWorkbenchAgentResponse() (response *DeleteAIWorkbenchAgentResponse) {
    response = &DeleteAIWorkbenchAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAIWorkbenchAgent
// Delete Agent
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAIWorkbenchAgent(request *DeleteAIWorkbenchAgentRequest) (response *DeleteAIWorkbenchAgentResponse, err error) {
    return c.DeleteAIWorkbenchAgentWithContext(context.Background(), request)
}

// DeleteAIWorkbenchAgent
// Delete Agent
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAIWorkbenchAgentWithContext(ctx context.Context, request *DeleteAIWorkbenchAgentRequest) (response *DeleteAIWorkbenchAgentResponse, err error) {
    if request == nil {
        request = NewDeleteAIWorkbenchAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteAIWorkbenchAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAIWorkbenchAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAIWorkbenchAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIWorkbenchTaskRequest() (request *DeleteAIWorkbenchTaskRequest) {
    request = &DeleteAIWorkbenchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAIWorkbenchTask")
    
    
    return
}

func NewDeleteAIWorkbenchTaskResponse() (response *DeleteAIWorkbenchTaskResponse) {
    response = &DeleteAIWorkbenchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAIWorkbenchTask
// This API is used to delete a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAIWorkbenchTask(request *DeleteAIWorkbenchTaskRequest) (response *DeleteAIWorkbenchTaskResponse, err error) {
    return c.DeleteAIWorkbenchTaskWithContext(context.Background(), request)
}

// DeleteAIWorkbenchTask
// This API is used to delete a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAIWorkbenchTaskWithContext(ctx context.Context, request *DeleteAIWorkbenchTaskRequest) (response *DeleteAIWorkbenchTaskResponse, err error) {
    if request == nil {
        request = NewDeleteAIWorkbenchTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteAIWorkbenchTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAIWorkbenchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAIWorkbenchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIWorkbenchAgentRequest() (request *DescribeAIWorkbenchAgentRequest) {
    request = &DescribeAIWorkbenchAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAIWorkbenchAgent")
    
    
    return
}

func NewDescribeAIWorkbenchAgentResponse() (response *DescribeAIWorkbenchAgentResponse) {
    response = &DescribeAIWorkbenchAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIWorkbenchAgent
// Query Agent details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchAgent(request *DescribeAIWorkbenchAgentRequest) (response *DescribeAIWorkbenchAgentResponse, err error) {
    return c.DescribeAIWorkbenchAgentWithContext(context.Background(), request)
}

// DescribeAIWorkbenchAgent
// Query Agent details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchAgentWithContext(ctx context.Context, request *DescribeAIWorkbenchAgentRequest) (response *DescribeAIWorkbenchAgentResponse, err error) {
    if request == nil {
        request = NewDescribeAIWorkbenchAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAIWorkbenchAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIWorkbenchAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIWorkbenchAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIWorkbenchArtifactRequest() (request *DescribeAIWorkbenchArtifactRequest) {
    request = &DescribeAIWorkbenchArtifactRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAIWorkbenchArtifact")
    
    
    return
}

func NewDescribeAIWorkbenchArtifactResponse() (response *DescribeAIWorkbenchArtifactResponse) {
    response = &DescribeAIWorkbenchArtifactResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIWorkbenchArtifact
// Query artifact details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchArtifact(request *DescribeAIWorkbenchArtifactRequest) (response *DescribeAIWorkbenchArtifactResponse, err error) {
    return c.DescribeAIWorkbenchArtifactWithContext(context.Background(), request)
}

// DescribeAIWorkbenchArtifact
// Query artifact details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchArtifactWithContext(ctx context.Context, request *DescribeAIWorkbenchArtifactRequest) (response *DescribeAIWorkbenchArtifactResponse, err error) {
    if request == nil {
        request = NewDescribeAIWorkbenchArtifactRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAIWorkbenchArtifact")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIWorkbenchArtifact require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIWorkbenchArtifactResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIWorkbenchExecutionRequest() (request *DescribeAIWorkbenchExecutionRequest) {
    request = &DescribeAIWorkbenchExecutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAIWorkbenchExecution")
    
    
    return
}

func NewDescribeAIWorkbenchExecutionResponse() (response *DescribeAIWorkbenchExecutionResponse) {
    response = &DescribeAIWorkbenchExecutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIWorkbenchExecution
// Query execution details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchExecution(request *DescribeAIWorkbenchExecutionRequest) (response *DescribeAIWorkbenchExecutionResponse, err error) {
    return c.DescribeAIWorkbenchExecutionWithContext(context.Background(), request)
}

// DescribeAIWorkbenchExecution
// Query execution details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchExecutionWithContext(ctx context.Context, request *DescribeAIWorkbenchExecutionRequest) (response *DescribeAIWorkbenchExecutionResponse, err error) {
    if request == nil {
        request = NewDescribeAIWorkbenchExecutionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAIWorkbenchExecution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIWorkbenchExecution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIWorkbenchExecutionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIWorkbenchSessionRequest() (request *DescribeAIWorkbenchSessionRequest) {
    request = &DescribeAIWorkbenchSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAIWorkbenchSession")
    
    
    return
}

func NewDescribeAIWorkbenchSessionResponse() (response *DescribeAIWorkbenchSessionResponse) {
    response = &DescribeAIWorkbenchSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIWorkbenchSession
// Query session details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchSession(request *DescribeAIWorkbenchSessionRequest) (response *DescribeAIWorkbenchSessionResponse, err error) {
    return c.DescribeAIWorkbenchSessionWithContext(context.Background(), request)
}

// DescribeAIWorkbenchSession
// Query session details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchSessionWithContext(ctx context.Context, request *DescribeAIWorkbenchSessionRequest) (response *DescribeAIWorkbenchSessionResponse, err error) {
    if request == nil {
        request = NewDescribeAIWorkbenchSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAIWorkbenchSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIWorkbenchSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIWorkbenchSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIWorkbenchSkillRequest() (request *DescribeAIWorkbenchSkillRequest) {
    request = &DescribeAIWorkbenchSkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAIWorkbenchSkill")
    
    
    return
}

func NewDescribeAIWorkbenchSkillResponse() (response *DescribeAIWorkbenchSkillResponse) {
    response = &DescribeAIWorkbenchSkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIWorkbenchSkill
// Query skill details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchSkill(request *DescribeAIWorkbenchSkillRequest) (response *DescribeAIWorkbenchSkillResponse, err error) {
    return c.DescribeAIWorkbenchSkillWithContext(context.Background(), request)
}

// DescribeAIWorkbenchSkill
// Query skill details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAIWorkbenchSkillWithContext(ctx context.Context, request *DescribeAIWorkbenchSkillRequest) (response *DescribeAIWorkbenchSkillResponse, err error) {
    if request == nil {
        request = NewDescribeAIWorkbenchSkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAIWorkbenchSkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIWorkbenchSkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIWorkbenchSkillResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNotifyHistoriesRequest() (request *DescribeAlarmNotifyHistoriesRequest) {
    request = &DescribeAlarmNotifyHistoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNotifyHistories")
    
    
    return
}

func NewDescribeAlarmNotifyHistoriesResponse() (response *DescribeAlarmNotifyHistoriesResponse) {
    response = &DescribeAlarmNotifyHistoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNotifyHistories
// Query alarm notification history as needed
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlarmNotifyHistories(request *DescribeAlarmNotifyHistoriesRequest) (response *DescribeAlarmNotifyHistoriesResponse, err error) {
    return c.DescribeAlarmNotifyHistoriesWithContext(context.Background(), request)
}

// DescribeAlarmNotifyHistories
// Query alarm notification history as needed
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlarmNotifyHistoriesWithContext(ctx context.Context, request *DescribeAlarmNotifyHistoriesRequest) (response *DescribeAlarmNotifyHistoriesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNotifyHistoriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmNotifyHistories")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotifyHistories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNotifyHistoriesResponse()
    err = c.Send(request, response)
    return
}

func NewGetAIWorkbenchArtifactDownloadURLRequest() (request *GetAIWorkbenchArtifactDownloadURLRequest) {
    request = &GetAIWorkbenchArtifactDownloadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "GetAIWorkbenchArtifactDownloadURL")
    
    
    return
}

func NewGetAIWorkbenchArtifactDownloadURLResponse() (response *GetAIWorkbenchArtifactDownloadURLResponse) {
    response = &GetAIWorkbenchArtifactDownloadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAIWorkbenchArtifactDownloadURL
// Get the download URL of AI Workbench artifacts.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) GetAIWorkbenchArtifactDownloadURL(request *GetAIWorkbenchArtifactDownloadURLRequest) (response *GetAIWorkbenchArtifactDownloadURLResponse, err error) {
    return c.GetAIWorkbenchArtifactDownloadURLWithContext(context.Background(), request)
}

// GetAIWorkbenchArtifactDownloadURL
// Get the download URL of AI Workbench artifacts.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) GetAIWorkbenchArtifactDownloadURLWithContext(ctx context.Context, request *GetAIWorkbenchArtifactDownloadURLRequest) (response *GetAIWorkbenchArtifactDownloadURLResponse, err error) {
    if request == nil {
        request = NewGetAIWorkbenchArtifactDownloadURLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "GetAIWorkbenchArtifactDownloadURL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAIWorkbenchArtifactDownloadURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAIWorkbenchArtifactDownloadURLResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchAgentsRequest() (request *ListAIWorkbenchAgentsRequest) {
    request = &ListAIWorkbenchAgentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchAgents")
    
    
    return
}

func NewListAIWorkbenchAgentsResponse() (response *ListAIWorkbenchAgentsResponse) {
    response = &ListAIWorkbenchAgentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchAgents
// Query the Agent list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchAgents(request *ListAIWorkbenchAgentsRequest) (response *ListAIWorkbenchAgentsResponse, err error) {
    return c.ListAIWorkbenchAgentsWithContext(context.Background(), request)
}

// ListAIWorkbenchAgents
// Query the Agent list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchAgentsWithContext(ctx context.Context, request *ListAIWorkbenchAgentsRequest) (response *ListAIWorkbenchAgentsResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchAgentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchAgents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchAgentsResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchArtifactsRequest() (request *ListAIWorkbenchArtifactsRequest) {
    request = &ListAIWorkbenchArtifactsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchArtifacts")
    
    
    return
}

func NewListAIWorkbenchArtifactsResponse() (response *ListAIWorkbenchArtifactsResponse) {
    response = &ListAIWorkbenchArtifactsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchArtifacts
// Query the product list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchArtifacts(request *ListAIWorkbenchArtifactsRequest) (response *ListAIWorkbenchArtifactsResponse, err error) {
    return c.ListAIWorkbenchArtifactsWithContext(context.Background(), request)
}

// ListAIWorkbenchArtifacts
// Query the product list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchArtifactsWithContext(ctx context.Context, request *ListAIWorkbenchArtifactsRequest) (response *ListAIWorkbenchArtifactsResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchArtifactsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchArtifacts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchArtifacts require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchArtifactsResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchExecutionsRequest() (request *ListAIWorkbenchExecutionsRequest) {
    request = &ListAIWorkbenchExecutionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchExecutions")
    
    
    return
}

func NewListAIWorkbenchExecutionsResponse() (response *ListAIWorkbenchExecutionsResponse) {
    response = &ListAIWorkbenchExecutionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchExecutions
// Query the execution list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchExecutions(request *ListAIWorkbenchExecutionsRequest) (response *ListAIWorkbenchExecutionsResponse, err error) {
    return c.ListAIWorkbenchExecutionsWithContext(context.Background(), request)
}

// ListAIWorkbenchExecutions
// Query the execution list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchExecutionsWithContext(ctx context.Context, request *ListAIWorkbenchExecutionsRequest) (response *ListAIWorkbenchExecutionsResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchExecutionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchExecutions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchExecutions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchExecutionsResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchMCPsRequest() (request *ListAIWorkbenchMCPsRequest) {
    request = &ListAIWorkbenchMCPsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchMCPs")
    
    
    return
}

func NewListAIWorkbenchMCPsResponse() (response *ListAIWorkbenchMCPsResponse) {
    response = &ListAIWorkbenchMCPsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchMCPs
// Query the MCP list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchMCPs(request *ListAIWorkbenchMCPsRequest) (response *ListAIWorkbenchMCPsResponse, err error) {
    return c.ListAIWorkbenchMCPsWithContext(context.Background(), request)
}

// ListAIWorkbenchMCPs
// Query the MCP list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchMCPsWithContext(ctx context.Context, request *ListAIWorkbenchMCPsRequest) (response *ListAIWorkbenchMCPsResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchMCPsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchMCPs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchMCPs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchMCPsResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchMessagesRequest() (request *ListAIWorkbenchMessagesRequest) {
    request = &ListAIWorkbenchMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchMessages")
    
    
    return
}

func NewListAIWorkbenchMessagesResponse() (response *ListAIWorkbenchMessagesResponse) {
    response = &ListAIWorkbenchMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchMessages
// This API is used to query message list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchMessages(request *ListAIWorkbenchMessagesRequest) (response *ListAIWorkbenchMessagesResponse, err error) {
    return c.ListAIWorkbenchMessagesWithContext(context.Background(), request)
}

// ListAIWorkbenchMessages
// This API is used to query message list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchMessagesWithContext(ctx context.Context, request *ListAIWorkbenchMessagesRequest) (response *ListAIWorkbenchMessagesResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchMessagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchMessages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchMessages require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchResourceInstancesRequest() (request *ListAIWorkbenchResourceInstancesRequest) {
    request = &ListAIWorkbenchResourceInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchResourceInstances")
    
    
    return
}

func NewListAIWorkbenchResourceInstancesResponse() (response *ListAIWorkbenchResourceInstancesResponse) {
    response = &ListAIWorkbenchResourceInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchResourceInstances
// List resource instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchResourceInstances(request *ListAIWorkbenchResourceInstancesRequest) (response *ListAIWorkbenchResourceInstancesResponse, err error) {
    return c.ListAIWorkbenchResourceInstancesWithContext(context.Background(), request)
}

// ListAIWorkbenchResourceInstances
// List resource instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchResourceInstancesWithContext(ctx context.Context, request *ListAIWorkbenchResourceInstancesRequest) (response *ListAIWorkbenchResourceInstancesResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchResourceInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchResourceInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchResourceInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchResourceInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchResourceMapsRequest() (request *ListAIWorkbenchResourceMapsRequest) {
    request = &ListAIWorkbenchResourceMapsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchResourceMaps")
    
    
    return
}

func NewListAIWorkbenchResourceMapsResponse() (response *ListAIWorkbenchResourceMapsResponse) {
    response = &ListAIWorkbenchResourceMapsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchResourceMaps
// Query the list of resource maps
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchResourceMaps(request *ListAIWorkbenchResourceMapsRequest) (response *ListAIWorkbenchResourceMapsResponse, err error) {
    return c.ListAIWorkbenchResourceMapsWithContext(context.Background(), request)
}

// ListAIWorkbenchResourceMaps
// Query the list of resource maps
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchResourceMapsWithContext(ctx context.Context, request *ListAIWorkbenchResourceMapsRequest) (response *ListAIWorkbenchResourceMapsResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchResourceMapsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchResourceMaps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchResourceMaps require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchResourceMapsResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchSessionsRequest() (request *ListAIWorkbenchSessionsRequest) {
    request = &ListAIWorkbenchSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchSessions")
    
    
    return
}

func NewListAIWorkbenchSessionsResponse() (response *ListAIWorkbenchSessionsResponse) {
    response = &ListAIWorkbenchSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchSessions
// Query session list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchSessions(request *ListAIWorkbenchSessionsRequest) (response *ListAIWorkbenchSessionsResponse, err error) {
    return c.ListAIWorkbenchSessionsWithContext(context.Background(), request)
}

// ListAIWorkbenchSessions
// Query session list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchSessionsWithContext(ctx context.Context, request *ListAIWorkbenchSessionsRequest) (response *ListAIWorkbenchSessionsResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchSessionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchSessions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchSessions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchSkillsRequest() (request *ListAIWorkbenchSkillsRequest) {
    request = &ListAIWorkbenchSkillsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchSkills")
    
    
    return
}

func NewListAIWorkbenchSkillsResponse() (response *ListAIWorkbenchSkillsResponse) {
    response = &ListAIWorkbenchSkillsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchSkills
// Query the skill list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchSkills(request *ListAIWorkbenchSkillsRequest) (response *ListAIWorkbenchSkillsResponse, err error) {
    return c.ListAIWorkbenchSkillsWithContext(context.Background(), request)
}

// ListAIWorkbenchSkills
// Query the skill list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchSkillsWithContext(ctx context.Context, request *ListAIWorkbenchSkillsRequest) (response *ListAIWorkbenchSkillsResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchSkillsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchSkills")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchSkills require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchSkillsResponse()
    err = c.Send(request, response)
    return
}

func NewListAIWorkbenchTasksRequest() (request *ListAIWorkbenchTasksRequest) {
    request = &ListAIWorkbenchTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ListAIWorkbenchTasks")
    
    
    return
}

func NewListAIWorkbenchTasksResponse() (response *ListAIWorkbenchTasksResponse) {
    response = &ListAIWorkbenchTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAIWorkbenchTasks
// This API is used to query the task list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchTasks(request *ListAIWorkbenchTasksRequest) (response *ListAIWorkbenchTasksResponse, err error) {
    return c.ListAIWorkbenchTasksWithContext(context.Background(), request)
}

// ListAIWorkbenchTasks
// This API is used to query the task list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAIWorkbenchTasksWithContext(ctx context.Context, request *ListAIWorkbenchTasksRequest) (response *ListAIWorkbenchTasksResponse, err error) {
    if request == nil {
        request = NewListAIWorkbenchTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ListAIWorkbenchTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAIWorkbenchTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAIWorkbenchTasksResponse()
    err = c.Send(request, response)
    return
}

func NewTriggerAIWorkbenchTaskRequest() (request *TriggerAIWorkbenchTaskRequest) {
    request = &TriggerAIWorkbenchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "TriggerAIWorkbenchTask")
    
    
    return
}

func NewTriggerAIWorkbenchTaskResponse() (response *TriggerAIWorkbenchTaskResponse) {
    response = &TriggerAIWorkbenchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TriggerAIWorkbenchTask
// Manually trigger a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) TriggerAIWorkbenchTask(request *TriggerAIWorkbenchTaskRequest) (response *TriggerAIWorkbenchTaskResponse, err error) {
    return c.TriggerAIWorkbenchTaskWithContext(context.Background(), request)
}

// TriggerAIWorkbenchTask
// Manually trigger a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) TriggerAIWorkbenchTaskWithContext(ctx context.Context, request *TriggerAIWorkbenchTaskRequest) (response *TriggerAIWorkbenchTaskResponse, err error) {
    if request == nil {
        request = NewTriggerAIWorkbenchTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "TriggerAIWorkbenchTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TriggerAIWorkbenchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewTriggerAIWorkbenchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAIWorkbenchAgentRequest() (request *UpdateAIWorkbenchAgentRequest) {
    request = &UpdateAIWorkbenchAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateAIWorkbenchAgent")
    
    
    return
}

func NewUpdateAIWorkbenchAgentResponse() (response *UpdateAIWorkbenchAgentResponse) {
    response = &UpdateAIWorkbenchAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAIWorkbenchAgent
// Update an Agent
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateAIWorkbenchAgent(request *UpdateAIWorkbenchAgentRequest) (response *UpdateAIWorkbenchAgentResponse, err error) {
    return c.UpdateAIWorkbenchAgentWithContext(context.Background(), request)
}

// UpdateAIWorkbenchAgent
// Update an Agent
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateAIWorkbenchAgentWithContext(ctx context.Context, request *UpdateAIWorkbenchAgentRequest) (response *UpdateAIWorkbenchAgentResponse, err error) {
    if request == nil {
        request = NewUpdateAIWorkbenchAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateAIWorkbenchAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAIWorkbenchAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAIWorkbenchAgentResponse()
    err = c.Send(request, response)
    return
}
